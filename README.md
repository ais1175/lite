# Fivemanage Lite

Fivemanage Lite is an open-source, lightweight management service designed for gaming communities. It provides essential features for file storage, structured logging, and community organization.

## Features

- Multi-tenant organization support.
- File storage with S3-compatible providers (AWS S3, Cloudflare R2, MinIO).
- High-performance structured logging powered by ClickHouse.
- Built-in authentication and session management.
- OpenTelemetry integration for tracing (Jaeger).
- Modern React-based administrative dashboard.

---

## Production Setup

The recommended way to run Fivemanage Lite in production is using Docker.

### Prerequisites

- Docker and Docker Compose installed on your system.
- A PostgreSQL database (Metadata storage).
- A ClickHouse instance (Logging storage).
- An S3-compatible storage bucket (File storage).

### Deployment Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/fivemanage/fivemanage-lite.git
   cd fivemanage-lite
   ```

2. **Configure Environment Variables:**
   Copy the template environment file and update it with your production values.
   ```bash
   cp .env.template .env
   ```

3. **Start Infrastructure Services:**
   You can use the provided Docker Compose file to start required services (PostgreSQL, ClickHouse, MinIO, Jaeger).
   ```bash
   docker compose -f deployments/docker-compose.yml up -d
   ```

4. **Run the Application:**
   Build and run the Fivemanage Lite container using the provided Dockerfile.
   ```bash
   docker build -t fivemanage-lite -f build/package/Dockerfile .
   docker run -d --name fivemanage-lite -p 8080:8080 --env-file .env fivemanage-lite
   ```

The application will be accessible at `http://localhost:8080`. Database migrations are handled automatically on startup.

### Initial Login

After starting the application for the first time, you can log in to the administrative dashboard using the following credentials:

- **Username:** `admin`
- **Password:** The value you set for `ADMIN_PASSWORD` in your `.env` file.

---

## Configuration

The application is configured via environment variables.

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Port the server listens on | `8080` |
| `DSN` | PostgreSQL connection string | - |
| `ADMIN_PASSWORD` | Initial password for the 'admin' user | `password` |
| `API_TOKEN_HMAC_SECRET` | 32-byte secret for signing API tokens | - |
| `S3_PROVIDER` | S3 provider (`s3`, `r2`, or `minio`) | `minio` |
| `AWS_ACCESS_KEY_ID` | S3 access key ID | - |
| `AWS_SECRET_ACCESS_KEY` | S3 secret access key | - |
| `AWS_ENDPOINT` | S3 endpoint URL | - |
| `AWS_BUCKET` | S3 bucket name | - |
| `AWS_REGION` | S3 region | - |
| `CLICKHOUSE_HOST` | ClickHouse host address | `localhost:19000` |
| `CLICKHOUSE_USERNAME` | ClickHouse username | `default` |
| `CLICKHOUSE_PASSWORD` | ClickHouse password | `password` |
| `CLICKHOUSE_DATABASE` | ClickHouse database name | `default` |
| `ENV` | Environment mode (`production` or `dev`) | `dev` |

---

## Development Setup

If you want to contribute to the project or run it locally for development, follow these instructions.

### Prerequisites

- **Go**: 1.24 or later
- **Node.js**: 22.x or later
- **pnpm**: 9.x or later
- **Air**: For backend hot-reloading (recommended)

### Initial Setup

1. **Install Frontend Dependencies:**
   ```bash
   cd web
   pnpm install
   cd ..
   ```

2. **Download Backend Dependencies:**
   ```bash
   go mod download
   ```

3. **Start Development Infrastructure:**
   ```bash
   docker compose -f deployments/docker-compose.yml up -d
   ```

4. **Configure Local Environment:**
   Ensure your `.env` file points to the local services (the default values in `.env.template` are pre-configured for Docker Compose).

### Running the Application

To develop efficiently, run the backend and frontend in separate terminals.

1. **Start the Backend:**
   Using Air (hot-reloading):
   ```bash
   air
   ```
   Or standard Go:
   ```bash
   go run cmd/lite/lite.go
   ```

2. **Start the Frontend:**
   ```bash
   cd web
   pnpm dev
   ```

The frontend development server runs on `http://localhost:5173` and proxies API requests to the backend.

---

## Contributing

We welcome contributions. Please follow these guidelines:

1. **Branching:** Use descriptive branch names like `feature/new-feature` or `fix/bug-description`.
2. **Formatting:** Use `go fmt` for Go and `pnpm lint` for React code.
3. **Commits:** Provide concise commit messages that explain the intent of the change.
4. **Pull Requests:** Ensure your PR has a clear title and description of the changes made.

## Project Structure

- `cmd/lite`: Main entry point and CLI configuration.
- `internal/`: Application logic, API handlers, and services.
- `pkg/`: Reusable packages for storage, logging, and caching.
- `web/`: Frontend React application.
- `deployments/`: Docker Compose and deployment configurations.
- `migrate/`: PostgreSQL schema migrations.
- `build/`: Dockerfile and build scripts.

## License

Fivemanage Lite is released under the [MIT License](LICENSE.md).
