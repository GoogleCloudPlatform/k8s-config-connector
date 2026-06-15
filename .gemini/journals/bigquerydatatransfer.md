### [2026-05-17] bigquerydatatransfer migration
- **Context**: Moving mockgcp bigquerydatatransfer to httptogrpc (PR 8302).
- **Problem**: Previously used RewriteError for 404s to clear error.Errors.
- **Solution**: Removed as httptogrpc does not support it.
- **Impact**: It is important to remember that RewriteError is not supported in httptogrpc.
