# POST /api/v1/media/upload

## Auth
- Required: `Authorization: Bearer <jwt-or-api-key>`.

## Request
- `multipart/form-data` with `files` field containing one or more files.

## Success Response
- `200 OK`
```json
{
  "success": true,
  "data": {
    "message": "upload successful",
    "files": [
      "https://bucket.s3.region.amazonaws.com/images/abc.webp"
    ]
  }
}
```

## Errors
- `400 BAD_REQUEST` invalid form or missing files.
- `401 UNAUTHORIZED` missing/invalid auth.
- `500 INTERNAL_ERROR` upload or persistence failure.
