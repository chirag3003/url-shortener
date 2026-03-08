# GET /api/v1/user/me

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "id": "739182001234567890",
    "name": "Alice",
    "email": "alice@example.com",
    "phoneNo": "+1234567890"
  }
}
```

## Errors
- `401 UNAUTHORIZED` missing or invalid token.
- `404 NOT_FOUND` user not found.
- `500 INTERNAL_ERROR` unexpected server error.
