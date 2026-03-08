# POST /api/v1/auth/login

## Auth
- Public.

## Request
```json
{
  "email": "alice@example.com",
  "password": "securepass123"
}
```

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "token": "<jwt>",
    "user": {
      "id": "739182001234567890",
      "name": "Alice",
      "email": "alice@example.com"
    }
  }
}
```

## Errors
- `400 BAD_REQUEST` invalid body.
- `401 INVALID_CREDENTIALS` invalid email/password.
- `422 VALIDATION_ERROR` invalid fields.
- `500 INTERNAL_ERROR` unexpected server error.
