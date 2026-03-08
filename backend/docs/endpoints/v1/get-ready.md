# GET /ready

## Auth
- Public.

## Behavior
- Performs a PostgreSQL ping check.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "status": "ready"
  }
}
```

## Errors
- `503 NOT_READY` database not available.
