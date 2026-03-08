# POST /api/v1/links

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Request
Same payload as quick endpoint.

## Success Response
- `201 Created`
- Returns link object (same format as quick endpoint), including `userId`.

## Errors
- `400 BAD_REQUEST`
- `401 UNAUTHORIZED`
- `409 ALIAS_TAKEN`
- `422 VALIDATION_ERROR`
- `500 INTERNAL_ERROR`
