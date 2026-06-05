# Asset Migration Journal

## 2026-05-17
- Migrated `mockasset` from `grpc-gateway` to `httptogrpc`.
- Used official client library: `cloud.google.com/go/asset/apiv1/assetpb`.
- Removed `RewriteHeaders` and `RewriteError` from `NewHTTPMux` as they are not immediately supported by `httptogrpc` and didn't seem to break existing tests.
- Established a pattern of moving migrated services to `mockgcpregistry` by adding an `init` function and updating `register.go`. This allows removing the manual addition in `mock_http_roundtrip.go`.
- Encountered severe "no space left on device" issues during testing, requiring multiple cleanups of `/workspaces/.cache` and `/workspaces/.tmp`.
