# ComputeFirewall Direct Controller Implementation Journal

## Observations
1. **MockGCP HTTP Handler Registration**: When implementing a mock for a direct controller (which uses HTTP REST / JSON), simply registering the service with `pb.RegisterFirewallsServer` on the gRPC server is not enough. You must also register the HTTP gateway handler via `pb.RegisterFirewallsHandler(ctx, mux.ServeMux, conn)` in `NewHTTPMux` in `mockgcp/mockcompute/service.go`.
2. **Harness Enablement**: In the test harness (`config/tests/samples/create/harness.go`), you must explicitly add a case for `ComputeFirewall` under `schema.GroupKind` to prevent it from being skipped as "not supported by mock gcp".
3. **Global Scope Identity**: Like ComputeNetwork, ComputeFirewall is a global resource under `projects/<project>/global/firewalls/<firewall-name>`.
