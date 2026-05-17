## 2026-05-17 - compute migration

- Migrated `mockcompute` to `httptogrpc`.
- Had to update `httptogrpc` to ignore `[google.cloud.operation_service]` and `[google.cloud.operation_polling_method]` annotations found in compute protos.
- Encountered ambiguous service matching because many compute services share the same set of standard methods (Get, List, Insert, etc.). Solved by using `httptogrpc.WithServiceName` for all services in `mockcompute/service.go`.
- Updated `mockcontainer/nodepool.go` as it also used compute proto types.