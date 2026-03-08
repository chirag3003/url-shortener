# API Changelog

## 2026-03-08
- Added Goose-based versioned SQL migrations and migration CLI/Makefile targets.
- Migrated backend data layer from MongoDB to PostgreSQL.
- Added Hyperflake-based `BIGINT` ID generation for persisted entities.
- Added URL shortener endpoints for create/list/get/update/delete.
- Added public quick-shortener and redirect endpoint.
- Added analytics summary, timeseries, and breakdown endpoints.
- Added API key management endpoints (create/list/revoke).
- Updated auth middleware to support JWT and API key bearer tokens.
