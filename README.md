# Kairo

A modern task management application built with a Go backend and a TypeScript frontend.

Kairo is designed around fast task management, background processing, and a clean backend architecture. It uses PostgreSQL for persistent data, Valkey for caching and background jobs, and Go for the API and application services.

## Features

* **Task Management** — Create, update, complete, and delete tasks
* **Task Organization** — Organize tasks and manage their lifecycle
* **REST API** — HTTP API built with Go and Echo
* **PostgreSQL** — Persistent storage with connection pooling and migrations
* **Valkey** — Redis-compatible datastore for caching and background jobs
* **Background Jobs** — Asynchronous task processing with Asynq
* **Authentication** — User authentication and authorization
* **Structured Logging** — Application logging with Zerolog
* **Observability** — New Relic integration
* **API Documentation** — OpenAPI/Swagger support
* **Security** — CORS, rate limiting, secure headers, and request validation
* **Containerized Development** — PostgreSQL and Valkey managed with Podman

## Tech Stack

| Layer             | Technology        |
| ----------------- | ----------------- |
| Backend           | Go                |
| HTTP Framework    | Echo              |
| Frontend          | TypeScript        |
| Database          | PostgreSQL        |
| Cache             | Valkey            |
| Background Jobs   | Asynq             |
| Logging           | Zerolog           |
| Observability     | New Relic         |
| API Documentation | OpenAPI / Swagger |
| Containers        | Podman            |
| Migrations        | golang-migrate    |

## Project Structure

```text
kairo/
├── apps/
│   ├── backend/
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── config/
│   │   │   ├── database/
│   │   │   ├── handlers/
│   │   │   ├── middleware/
│   │   │   ├── repositories/
│   │   │   ├── services/
│   │   │   └── ...
│   │   ├── migrations/
│   │   ├── .env.example
│   │   ├── go.mod
│   │   └── Taskfile.yml
│   │
│   └── frontend/
│
├── compose.yml
├── go.work
└── README.md
```

## Architecture

Kairo follows a layered backend architecture:

```text
                    ┌──────────────┐
                    │   Frontend   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   HTTP API   │
                    │    Echo      │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   Services   │
                    │ Business     │
                    │ Logic        │
                    └──────┬───────┘
                           │
                    ┌──────┴───────┐
                    ▼              ▼
              ┌───────────┐  ┌───────────┐
              │PostgreSQL │  │  Valkey   │
              │           │  │           │
              └───────────┘  └─────┬─────┘
                                    │
                                    ▼
                              ┌───────────┐
                              │   Asynq   │
                              │  Workers  │
                              └───────────┘
```

### Handlers

Handle HTTP-specific concerns:

* Request parsing
* Validation
* Authentication
* HTTP responses
* Error handling

### Services

Contain application and task-management logic:

* Creating tasks
* Updating tasks
* Completing tasks
* Deleting tasks
* User-specific operations
* Background job orchestration

### Repositories

Handle persistence:

* PostgreSQL queries
* Transactions
* Task persistence
* User persistence

### Background Workers

Asynq workers process operations that don't need to block an HTTP request.

Valkey provides the Redis-compatible backend used by Asynq.

## Local Development

### Prerequisites

Install:

* Go 1.27+
* Node.js 22+
* Podman
* Podman Compose
* Task
* Git

Verify:

```bash
go version
node --version
podman --version
podman compose version
task --version
```

## Quick Start

### 1. Clone

```bash
git clone https://github.com/Venkat1abhinav/kairo.git
cd kairo
```

### 2. Start PostgreSQL and Valkey

Kairo uses Podman Compose for local infrastructure.

```bash
podman compose up -d
```

Check the containers:

```bash
podman compose ps
```

The development environment provides:

```text
PostgreSQL → localhost:5432
Valkey     → localhost:6379
```

### 3. Verify PostgreSQL

```bash
podman exec -it kairo-postgres pg_isready -U postgres
```

Connect to the database:

```bash
podman exec -it kairo-postgres \
  psql -U postgres -d KAIRO
```

### 4. Verify Valkey

```bash
podman exec -it kairo-valkey valkey-cli ping
```

Expected:

```text
PONG
```

### 5. Configure Environment

Copy the example configuration:

```bash
cp apps/backend/.env.example apps/backend/.env
```

For local PostgreSQL:

```env
KAIRO_DATABASE.HOST=localhost
KAIRO_DATABASE.PORT=5432
KAIRO_DATABASE.USER=postgres
KAIRO_DATABASE.PASSWORD=postgres
KAIRO_DATABASE.NAME=KAIRO
KAIRO_DATABASE.SSL_MODE=disable

KAIRO_DATABASE.MAX_OPEN_CONNS=25
KAIRO_DATABASE.MAX_IDLE_CONNS=25
KAIRO_DATABASE.CONN_MAX_LIFETIME=300
KAIRO_DATABASE.CONN_MAX_IDLE_TIME=300
```

For Valkey:

```env
KAIRO_REDIS.ADDRESS=redis://localhost:6379
```

Valkey is compatible with the Redis protocol, so Redis connection URLs can be used by compatible clients.

### 6. Install Dependencies

Backend:

```bash
cd apps/backend
go mod download
```

Frontend:

```bash
cd apps/frontend
npm install
```

### 7. Run Migrations

From the backend directory:

```bash
task migrations:up
```

### 8. Start Kairo

Backend:

```bash
task run
```

The API will be available at:

```text
http://localhost:8080
```

Start the frontend using the project's frontend development command.

## Database

Kairo uses PostgreSQL as its primary persistent datastore.

PostgreSQL stores application data such as:

* Users
* Tasks
* Task state
* Task metadata
* Other persistent application data

Database migrations are version controlled and should be applied whenever the database schema changes.

Create a migration:

```bash
task migrations:new
```

Apply migrations:

```bash
task migrations:up
```

Rollback migrations:

```bash
task migrations:down
```

## Background Jobs

Kairo uses **Asynq** for background job processing.

Valkey provides the Redis-compatible datastore used by Asynq.

```text
             Kairo API
                 │
                 │ enqueue
                 ▼
              Valkey
                 │
                 │
                 ▼
            Asynq Worker
                 │
                 ▼
          Background Task
```

This allows longer-running operations to execute asynchronously without blocking API requests.

## Configuration

Configuration is provided through environment variables.

Main configuration groups include:

```text
KAIRO_DATABASE_*
KAIRO_REDIS_*
KAIRO_SERVER_*
KAIRO_AUTH_*
KAIRO_EMAIL_*
KAIRO_OBSERVABILITY_*
```

See:

```text
apps/backend/.env.example
```

for the complete configuration.

Do not commit `.env` files or production secrets.

## Development Commands

Run commands from:

```bash
cd apps/backend
```

Show available tasks:

```bash
task help
```

Start the backend:

```bash
task run
```

Run tests:

```bash
task test
```

Create a migration:

```bash
task migrations:new
```

Apply migrations:

```bash
task migrations:up
```

Rollback migrations:

```bash
task migrations:down
```

Format and tidy dependencies:

```bash
task tidy
```

## Testing

Run all Go tests:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run a specific package:

```bash
go test ./path/to/package
```

## Local Infrastructure

The repository includes a `compose.yml` containing the services required for local development.

```yaml
PostgreSQL
    │
    └── localhost:5432

Valkey
    │
    └── localhost:6379
```

Start:

```bash
podman compose up -d
```

Stop:

```bash
podman compose down
```

The named volumes preserve database and Valkey data.

To remove the containers and their volumes:

```bash
podman compose down -v
```

> `podman compose down -v` permanently deletes the local PostgreSQL and Valkey data.

## Production Considerations

Before deploying Kairo:

1. Use strong database credentials.
2. Store secrets in a dedicated secret-management system.
3. Use production-specific environment configuration.
4. Configure PostgreSQL connection pooling based on workload.
5. Enable TLS for external connections.
6. Restrict CORS to trusted origins.
7. Configure rate limiting.
8. Configure application monitoring and alerting.
9. Set up PostgreSQL backups.
10. Configure appropriate Valkey persistence and resource limits.
11. Run migrations as part of the deployment process.
12. Use a reverse proxy such as Caddy or Nginx.

## Contributing

1. Fork the repository.
2. Create a feature branch:

```bash
git checkout -b feature/my-feature
```

3. Make your changes.
4. Run the test suite:

```bash
go test ./...
```

5. Commit your changes:

```bash
git commit -m "feat: add my feature"
```

6. Push your branch:

```bash
git push origin feature/my-feature
```

7. Open a Pull Request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
