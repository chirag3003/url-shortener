# PATCH /api/v1/links/:id

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Request
```json
{
  "longUrl": "https://example.com/new-destination",
  "expiresAt": "2026-12-31T00:00:00Z",
  "redirectType": 301,
  "isActive": true
}
```

## Success Response
- `200 OK` with updated link payload.

## Errors
- `400 BAD_REQUEST`
- `401 UNAUTHORIZED`
- `403 FORBIDDEN`
- `404 NOT_FOUND`
- `422 VALIDATION_ERROR`
