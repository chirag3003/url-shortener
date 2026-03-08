# POST /api/v1/auth/register

## Auth
- Public.

## Request
```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "securepass123"
}
```

## Success Response
- `201 Created`
```json
{
  "success": true,
  "data": {
    "message": "registration successful"
  }
}
```

## Errors
- `400 BAD_REQUEST` invalid body.
- `409 USER_ALREADY_EXISTS` duplicate email.
- `422 VALIDATION_ERROR` invalid fields.
- `500 INTERNAL_ERROR` unexpected server error.
