# GET /api/v1/links/:id/analytics/summary

## Auth
- Required: `Authorization: Bearer <jwt>`.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "totalClicks": 1234,
    "uniqueVisitors": 932,
    "clicksLast24h": 50,
    "clicksLast7d": 300
  }
}
```
