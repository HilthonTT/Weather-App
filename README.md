# 🌤️ Weather App

[![CI/CD](https://github.com/HilthonTT/Weather-App/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/HilthonTT/Weather-App/actions/workflows/ci-cd.yml)

A mobile weather application with a **Flutter** frontend and a **Go** backend.
The backend proxies [OpenWeather](https://openweathermap.org/api) and
[Open-Meteo](https://open-meteo.com/), caches responses in Redis, and manages
users and their display settings in PostgreSQL.

![Home screen](https://github.com/user-attachments/assets/b7a40efc-308f-4238-8867-6fbf620e143d)

![Forecast screen](https://github.com/user-attachments/assets/c5b1dddd-7149-428c-8c3d-f40f1110ad84)

---

## 🚀 Features

- Current weather, hourly (3 h) forecast and 7-day forecast for your location
- City search with live weather tiles
- Unit preferences: °C / °F, km/h / mph, 12 h / 24 h
- Go API with Redis read-through cache, rate limiting, JWT auth and Swagger docs
- OpenTelemetry tracing (Jaeger) and expvar metrics
- Docker Compose for local infrastructure, GitHub Actions for CI/CD

---

## 🧱 Architecture

```
┌──────────────┐  HTTP/JSON  ┌──────────────────┐      ┌──────────────────────┐
│ Flutter app  │ ──────────▶ │  Go API (:8080)  │ ───▶ │ OpenWeather / Meteo  │
│ (Android/iOS)│             │  chi router      │      └──────────────────────┘
└──────────────┘             │                  │ ───▶ PostgreSQL (users, settings)
                             │                  │ ───▶ Redis (cache)
                             └──────────────────┘ ───▶ Jaeger (traces)
```

The backend always returns **metric** values (°C, m/s); the app converts them
to the user's preferred units.

## 📁 Project structure

```
.
├── .github/workflows/ci-cd.yml   # CI (lint/test) + CD (Docker image, APK)
├── docker-compose.yml            # Postgres, Redis, Jaeger for local dev
└── src/
    ├── server/                   # Go backend  → see src/server/README.md
    │   ├── cmd/api/              #   HTTP handlers, middleware, config, main
    │   ├── cmd/migrate/          #   SQL migrations
    │   ├── docs/                 #   Generated Swagger spec (make gen-docs)
    │   └── internal/             #   auth, db, env, mailer, ratelimiter,
    │                             #   store (+cache), tracing
    └── client/weather_app/       # Flutter app → see src/client/weather_app/README.md
        └── lib/
            ├── app/              #   App root & bottom-nav shell
            ├── common/           #   Shared constants, DI, widgets, utils
            └── modules/          #   Feature modules (clean architecture)
                ├── settings/     #     data / domain / presentation
                └── weather/      #     data / domain / presentation
```

---

## 🔧 Prerequisites

- [Flutter](https://flutter.dev/docs/get-started/install) (stable, Dart ≥ 3.7)
- [Go](https://go.dev/dl/) 1.24+
- [Docker](https://www.docker.com/products/docker-desktop)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI (`migrate`)
- Optional: [air](https://github.com/air-verse/air) (live reload), [Make](https://www.gnu.org/software/make/), [swag](https://github.com/swaggo/swag) (API docs)

You also need an **OpenWeather API key** and **Mailtrap** SMTP credentials
(used for the registration email).

---

## ⚡ Quick start

### 1. Infrastructure

From the **repository root**:

```bash
docker compose up -d        # Postgres :5432, Redis :6379, Jaeger UI :16686
```

### 2. Backend

```bash
cd src/server
cp .env.example .env        # then fill in OPEN_WEATHER_API_KEY and MAILTRAP_*
make migrate-up             # create the database schema
air                         # or: make run
```

Check it works: <http://localhost:8080/v1/health> ·
Swagger UI: <http://localhost:8080/v1/swagger/index.html>

### 3. App

```bash
cd src/client/weather_app
flutter pub get
flutter run --dart-define=API_URL=http://<your-computer-ip>:8080
```

`API_URL` defaults to `http://10.0.2.2:8080`, which is how the **Android
emulator** reaches your machine's `localhost`. Physical devices need your
computer's LAN IP.

---

## 📚 API overview

All routes are prefixed with `/v1`. Successful responses are wrapped as
`{"data": ...}`, errors as `{"error": "..."}`.

| Method | Path | Auth | Description |
| ------ | ---- | ---- | ----------- |
| GET | `/health` | – | Health check |
| GET | `/weather/{city}` | – | Current weather by city |
| GET | `/weather/coords/{lat}/{lon}` | – | Current weather by coordinates |
| GET | `/weather/forecast/{city}` | – | 5-day / 3-hour forecast by city |
| GET | `/weather/forecast/coords/{lat}/{lon}` | – | 5-day / 3-hour forecast by coordinates |
| GET | `/weather/open-meteo/coords/{lat}/{lon}` | – | 7-day daily + hourly forecast |
| POST | `/authentication/register` | – | Register (sends a verification email) |
| POST | `/authentication/login` | – | Returns a JWT |
| PUT | `/users/activate/{token}` | – | Verify email with the emailed token |
| GET | `/users/me` | Bearer | Current user |
| GET | `/users/{id}` | Bearer | User profile |
| GET / PATCH | `/users/{id}/settings` | Bearer (owner or admin) | Read / partially update unit settings |
| GET | `/debug/vars` | Basic | expvar metrics |

See the Swagger UI for request/response schemas.

---

## 🧪 Testing

```bash
cd src/server && go test ./...
cd src/client/weather_app && flutter test
```

## 🔄 CI/CD

[`.github/workflows/ci-cd.yml`](.github/workflows/ci-cd.yml) runs on every
push and pull request:

| Job | What it does |
| --- | ------------ |
| **Server · lint & test** | `gofmt`, `go mod tidy` check, `go vet`, `go test -race`, build |
| **Server · Docker build** | Builds `src/server/Dockerfile` |
| **Client · analyze & test** | `flutter analyze`, `flutter test` |

On pushes to `main` and on `v*` tags, once CI is green:

| Job | What it does |
| --- | ------------ |
| **Publish server image** | Pushes `ghcr.io/<owner>/weather-api` (`latest`, `sha-…`, semver tags) |
| **Build Android APK** | Builds a release APK, uploads it as a workflow artifact, and attaches it to the GitHub Release for tags |

Set the repository **variable** `API_URL` (Settings → Secrets and variables →
Actions → Variables) to the public URL of your deployed backend so the APK
points at it.

To cut a release: `git tag v1.0.0 && git push origin v1.0.0`.

---

## Credits

- CodeHQ (UI Design): https://www.youtube.com/watch?v=msoKuk-5QFg

## License

[MIT](LICENSE)
