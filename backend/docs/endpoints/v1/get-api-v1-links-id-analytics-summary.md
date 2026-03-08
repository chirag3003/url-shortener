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
    "totalClicksChange": 12.5,
    "uniqueVisitors": 932,
    "uniqueVisitorsChange": 8.2,
    "clicksLast24h": 50,
    "clicksLast24hChange": -3.1,
    "clicksLast7d": 300,
    "clicksLast7dChange": 15.0
  }
}
```
