# Weather API (Go)

HTTP backend for the Weather app. Built with
[chi](https://github.com/go-chi/chi), PostgreSQL, Redis and OpenTelemetry.

## Running locally

```bash
# from the repo root: start Postgres, Redis and Jaeger
docker compose up -d

cd src/server
cp .env.example .env    # fill in OPEN_WEATHER_API_KEY and MAILTRAP_*
make migrate-up
air                     # live reload (or `make run`)
```

All variables are documented in [`.env.example`](.env.example). The only
required ones are `OPEN_WEATHER_API_KEY`, `MAILTRAP_USERNAME` and
`MAILTRAP_PASSWORD`; everything else defaults to the docker-compose services.
Set `REDIS_ENABLED=false` to run without Redis.

## Make targets

| Target                    | Description                       |
| ------------------------- | --------------------------------- |
| `make run` / `make build` | Run or build the API              |
| `make test` / `make lint` | `go test` / `go vet`              |
| `make migrate-up`         | Apply migrations (uses `DB_ADDR`) |
| `make migrate-down <n>`   | Roll back `n` migrations          |
| `make migration <name>`   | Create a new migration pair       |
| `make gen-docs`           | Regenerate Swagger in `docs/`     |
| `make docker-build`       | Build the Docker image            |

## Behaviour notes

- **Units:** upstream requests use `units=metric`, so responses are in °C and m/s.
- **Caching:** weather responses are cached for 1 hour. Coordinate keys are
  rounded to 2 decimals (~1 km) and city keys are case-insensitive, so a
  request and the upstream echo map to the same entry.
- **Upstream errors:** unknown cities return `404`; 5xx/network errors are
  retried 3 times with linear backoff; other 4xx errors fail fast.
- **Settings:** `/users/{id}/settings` is accessible only to that user (or an
  `admin`). `PATCH` is partial—omitted fields keep their value. Defaults
  (°C, 24 h, km/h) are created on first access.
- **Rate limiting:** 20 requests per 5 s per client IP by default
  (`RATELIMITER_REQUESTS_COUNT`, `RATE_LIMITER_ENABLED`). Exceeding it returns
  `429` with a `Retry-After` header in seconds.

## Docker

```bash
docker build -t weather-api .
docker run --rm -p 8080:8080 --env-file .env weather-api
```

> `--env-file` does not understand the `export` prefix used in `.env`; strip it
> (e.g. `sed 's/^export //' .env > .env.docker`) or pass variables with `-e`.
> When the API runs in a container, point `DB_ADDR`/`REDIS_ADDR`/`JAEGER_ADDR`
> at hosts reachable from the container (e.g. `host.docker.internal`).

CI publishes the image to `ghcr.io/<owner>/weather-api` on every push to `main`.
Migrations are not run by the image—apply them with `make migrate-up` against
the target database.
