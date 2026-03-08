# GET /api/v1/links/:id/analytics/:kind

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Path Params
- `kind`: `referrers`, `devices`, `browsers`, `geography`

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": [
    {"key": "google.com", "count": 120}
  ]
}
```
