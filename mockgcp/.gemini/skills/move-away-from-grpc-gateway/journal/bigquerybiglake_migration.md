## 2026-05-17 - bigquerybiglake migration
- When using `mux.AddOperationsPath` in `NewHTTPMux`, ensure that `s.operations.RegisterGRPCServices(grpcServer)` is called in the `Register` method. This ensures the operations service is available to the client used by `AddOperationsPath`.
