# GET /api/v1/links

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Query Params
- `page` default `1`
- `limit` default `10`, max `100`
- `search` optional (`shortCode` or `longUrl`)

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "items": [],
    "page": 1,
    "limit": 10,
    "total": 0
  }
}
```
