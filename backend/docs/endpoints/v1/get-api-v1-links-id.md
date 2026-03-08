# GET /api/v1/links/:id

## Auth
- Required: `Authorization: Bearer <jwt>`.

## Success Response
- `200 OK` with link payload.

## Errors
- `400 BAD_REQUEST` invalid id.
- `401 UNAUTHORIZED`
- `403 FORBIDDEN` not owner.
- `404 NOT_FOUND` not found.
