# DELETE /api/v1/links/:id

## Auth
- Required: `Authorization: Bearer <jwt>`.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "message": "link deleted"
  }
}
```

## Errors
- `400 BAD_REQUEST`
- `401 UNAUTHORIZED`
- `403 FORBIDDEN`
- `404 NOT_FOUND`
