# Fivemanage Lite

## Deployment
To run Fivemanage Lite, you need to set up a MySQL or PostgreSQL database and an object storage service (S3 compatible).

### Database
You can use either MySQL or PostgreSQL. The default is MySQL.

### Object Storage
Currently only S3 or compatible object storage is supported.
- s3
- r2
- minio

Azure and GCP are not supported yet.

### Docker
You can have a look at the `docker-compose.yml` file in the `deployments` folder.


### Environment variables
You can copy the `.env.template` file to `.env` and set the values.

```env
ADMIN_PASSWORD=verysecurepassword
DB_DRIVER=mysql
DSN="root:root@tcp(localhost:3306)/fivemanage-lite"

AWS_ACCESS_KEY_ID=xxxx
AWS_SECRET_ACCESS_KEY=xxxx
AWS_ENDPOINT=
AWS_BUCKET=
AWS_REGION=
```

## Development

### Prerequisites

- Go: 1.23 or later
- Node.js: 20.x or later
- pnpm: 9.11.x or later
- air: hot reload for Go (https://github.com/air-verse/air)

### Setup

1. Install node_modules in `/web`.
2. Either run `go mod download` in the root directory, or just let `air` handle it for you, but simply running `air` in the root directory.
3. Set up docker.
   1. Run `docker compose -f deployments/docker-compose.yml -d`
4. Set up environment variables

   ```env
   ADMIN_PASSWORD=password
   DB_DRIVER=mysql
   DSN="root:root@tcp(localhost:3306)/fivemanage-lite-dev"
   ```

   - `DB_DRIVER`
     Can either be `mysql` or `pg`. Only mysql is properly tested.

### Running the application

First run all migrations.

1. Init the migration: `go cmd/lite/lite.go db init`
2. Run the migrations: `go cmd/lite/lite.go db migrate`

Start the actual app: 3. Run `air` or `go cmd/lite/lite.go` in the root directory to start the Go application 4. In `web/`, run `pnpm dev` to start the React application.
