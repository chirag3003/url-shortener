# PATCH /api/v1/user/me

## Auth
- Required: `Authorization: Bearer <jwt>`.

## Request
```json
{
  "name": "Alice Cooper",
  "avatarUrl": "https://example.com/new-avatar.webp"
}
```

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "id": "739182001234567890",
    "name": "Alice Cooper",
    "email": "alice@example.com",
    "phoneNo": "+1234567890",
    "avatarUrl": "https://example.com/new-avatar.webp"
  }
}
```

## Errors
- `400 BAD_REQUEST` invalid request body.
- `401 UNAUTHORIZED` missing or invalid token.
- `422 VALIDATION_ERROR` validation failed (e.g. name too short).
- `500 INTERNAL_ERROR` unexpected server error.
