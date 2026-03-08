# API Endpoint Index

## Health
- `GET /health` - Liveness check.
- `GET /ready` - Readiness check against PostgreSQL.

## Auth
- `POST /api/v1/auth/register` - Register with email/password.
- `POST /api/v1/auth/login` - Login with email/password and receive JWT.

## User
- `GET /api/v1/user/me` - Get authenticated user profile.

## Links
- `POST /api/v1/links/quick` - Public quick shortener (guest flow).
- `GET /api/v1/links/alias-availability` - Check alias availability.
- `POST /api/v1/links` - Create short link (authenticated).
- `GET /api/v1/links` - List current user's links.
- `GET /api/v1/links/:id` - Get a single link by ID.
- `PATCH /api/v1/links/:id` - Update link fields.
- `DELETE /api/v1/links/:id` - Delete a link.
- `GET /:code` - Public short-code redirect.

## Analytics
- `GET /api/v1/links/:id/analytics/summary` - Summary cards metrics.
- `GET /api/v1/links/:id/analytics/timeseries` - Time-series clicks data.
- `GET /api/v1/links/:id/analytics/:kind` - Breakdown (`referrers`, `devices`, `browsers`, `geography`).

## API Keys
- `POST /api/v1/api-keys` - Create an API key (shown once).
- `GET /api/v1/api-keys` - List API keys.
- `DELETE /api/v1/api-keys/:id` - Revoke API key.

## Media
- `POST /api/v1/media/upload` - Upload image files to S3.

## Database Ops
- Versioned SQL migrations are managed with Goose. See `backend/docs/deployment/migrations.md`.
