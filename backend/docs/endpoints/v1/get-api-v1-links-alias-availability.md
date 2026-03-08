# GET /api/v1/links/alias-availability

## Auth
- Public.

## Query Params
- `alias` (required): candidate custom alias.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "alias": "myalias",
    "available": true
  }
}
```

## Errors
- `400 BAD_REQUEST` missing alias.
- `500 INTERNAL_ERROR` unexpected server error.
