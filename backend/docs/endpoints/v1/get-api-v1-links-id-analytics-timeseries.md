# GET /api/v1/links/:id/analytics/timeseries

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Query Params
- `window`: `24h`, `7d`, `30d` (default `30d`).

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": [
    {"bucket": "2026-03-08T10:00:00Z", "clicks": 12}
  ]
}
```
