# Phase 1 validation

## Full stack

```bash
docker compose up --build -d
docker compose ps
```

Expected: postgres healthy, backend running, frontend running.

## API

```bash
curl --fail http://localhost:8080/health
curl --fail http://localhost:8080/api/v1/servers
curl --fail http://localhost:8080/metrics
./tests/api-smoke.sh
```

## Database

```bash
docker exec cloud-platform-postgres psql -U cloud_platform -d cloud_platform -c '\dt'
```

Expected tables: `schema_migrations`, `servers`.

## Frontend

Open `http://localhost:5173`, create a server, refresh, then delete it.

## Persistence

Create a server, restart the backend, then query the API again. The record must remain.

## Automated test

```bash
go test ./...
```

## Cleanup

```bash
docker compose down
```
