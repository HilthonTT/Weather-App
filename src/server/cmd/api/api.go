package main

import (
	"context"
	"errors"
	"expvar"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hilthontt/weather/internal/auth"
	"github.com/hilthontt/weather/internal/mailer"
	"github.com/hilthontt/weather/internal/ratelimiter"
	"github.com/hilthontt/weather/internal/store"
	"github.com/hilthontt/weather/internal/store/cache"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/hilthontt/weather/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// application wires together every dependency an HTTP handler may need.
type application struct {
	config        config
	store         store.Storage
	cacheStorage  cache.Storage
	logger        *zap.SugaredLogger
	authenticator auth.Authenticator
	rateLimiter   ratelimiter.Limiter
	mailer        mailer.Client
}

// mount builds the router with all middleware and routes.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Trace every request, including the ones rejected by the rate limiter.
	r.Use(NewTelemetryMiddleware(otel.Tracer("dynamic-endpoints")).Middleware)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	if app.config.rateLimiter.Enabled {
		r.Use(app.RateLimiterMiddleware)
	}

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.With(app.BasicAuthMiddleware()).Get("/debug/vars", expvar.Handler().ServeHTTP)

		// Relative URL: resolves to /v1/swagger/doc.json from the UI page.
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

		r.Route("/authentication", func(r chi.Router) {
			r.Post("/login", app.loginUserHandler)
			r.Post("/register", app.registerUserHandler)
		})

		r.Route("/weather", func(r chi.Router) {
			r.Get("/{city}", app.handleGetWeatherByCity)
			r.Get("/coords/{latitude}/{longitude}", app.handleGetWeatherByCoordinates)

			r.Get("/forecast/{city}", app.handleGetForecast)
			r.Get("/forecast/coords/{latitude}/{longitude}", app.handleGetForecastByCoordinates)

			r.Get("/open-meteo/coords/{latitude}/{longitude}", app.handleGetOpenMeteoByCoordinates)
		})

		r.Route("/users", func(r chi.Router) {
			// Public: the token is the credential.
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)

				r.Get("/me", app.getCurrentUserHandler)

				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", app.getUserHandler)

					r.Route("/settings", func(r chi.Router) {
						r.Use(app.settingsContextMiddleware)

						r.Get("/", app.getSettingsHandler)
						r.Patch("/", app.updateSettingsHandler)
					})
				})
			})
		})
	})

	return r
}

// run starts the HTTP server and blocks until it is shut down gracefully on
// SIGINT/SIGTERM.
func (app *application) run(mux http.Handler) error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
