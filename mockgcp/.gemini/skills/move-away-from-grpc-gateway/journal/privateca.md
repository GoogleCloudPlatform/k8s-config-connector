# 2026-05-18
- Moved mockgcp privateca to httptogrpc.
- Official proto import: `cloud.google.com/go/security/privateca/apiv1/privatecapb`.
- Replaced `httpmux` with `httptogrpc`.
- Verified compilation of `mockprivateca` package.
- Removed `RewriteError` as it's not supported by `httptogrpc` and the logic was for 404 errors which `httptogrpc` should handle.
