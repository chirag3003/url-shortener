# GET /:code

## Auth
- Public.

## Behavior
- Resolves short code.
- If active and not expired, redirects using configured `redirectType` (`301` or `302`).
- Records click analytics asynchronously.

## Responses
- `301` or `302` redirect to `longUrl`.
- `404 NOT_FOUND` if code does not exist, is inactive, or expired.
