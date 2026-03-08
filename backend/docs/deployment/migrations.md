# Database Migrations (Goose)

This backend uses `goose` with SQL migration files stored in `backend/db/migrations`.

## Files
- `backend/db/migrations/000001_init.sql` - tables and foreign keys.
- `backend/db/migrations/000002_indexes.sql` - indexes.

## Commands
- Apply all pending migrations:
  - `make migrate-up`
- Roll back one migration:
  - `make migrate-down`
- Show migration status:
  - `make migrate-status`

All commands require `DATABASE_URL` to be set.

## Runtime Behavior
- App startup executes migrations automatically via `db.RunMigrations` before serving traffic.
- CI also runs migrations before tests.

## Adding a New Migration
1. Create a new SQL file in `backend/db/migrations` with an incremented numeric prefix.
2. Include both sections:
   - `-- +goose Up`
   - `-- +goose Down`
3. Keep migrations backward-safe and idempotent where possible.
