# POST /api/v1/links/quick

## Auth
- Public (guest users supported).

## Request
```json
{
  "longUrl": "https://example.com/article",
  "customAlias": "myalias",
  "expiresAt": "2026-12-31T00:00:00Z",
  "redirectType": 302
}
```

`customAlias`, `expiresAt`, and `redirectType` are optional.

## Success Response
- `201 Created`
```json
{
  "success": true,
  "data": {
    "id": "739182001234567891",
    "longUrl": "https://example.com/article",
    "shortCode": "myalias",
    "shortUrl": "http://localhost:5000/myalias",
    "redirectType": 302,
    "isActive": true,
    "createdAt": "2026-03-08T12:00:00Z",
    "updatedAt": "2026-03-08T12:00:00Z"
  }
}
```

## Errors
- `400 BAD_REQUEST` invalid body or timestamp.
- `409 ALIAS_TAKEN` alias already exists.
- `422 VALIDATION_ERROR` invalid fields.
- `500 INTERNAL_ERROR` unexpected server error.
