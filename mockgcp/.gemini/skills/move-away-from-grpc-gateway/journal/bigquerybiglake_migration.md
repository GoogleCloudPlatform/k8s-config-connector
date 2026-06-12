## 2026-05-17 - bigquerybiglake migration
- When using `mux.AddOperationsPath` in `NewHTTPMux`, just add the path. Do NOT call `s.operations.RegisterGRPCServices(grpcServer)` in the `Register` method, as this causes duplicate gRPC service registration panics when multiple mock services are loaded.
