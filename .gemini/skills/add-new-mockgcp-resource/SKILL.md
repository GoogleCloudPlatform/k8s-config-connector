# Skill: Add new mockgcp resource

When adding a new mockgcp service/resource:

1. Look at an existing mock service like `mockgcp/mockfirestore/service.go` for inspiration.
2. If there is a KRM resource (e.g. MemcacheInstance), there is no need to create a gcloud test. The goal is to have mockgcp produce plausible output that allows the KRM tests to run. Then we will run it against real GCP in periodic tests and sync up the exact behavior of the service by comparing HTTP logs.
3. Use `httptogrpc` to avoid having to compile protos. You should be able to use the pre-compiled Go protos from `cloud.google.com/go/` (e.g., `cloud.google.com/go/memcache/apiv1/memcachepb`).
4. In `mockgcp/mock<service>/service.go`:
   - Initialize the `MockService` with `mockgcpregistry.Register(New)`.
   - Register the GRPC server in `Register(grpcServer *grpc.Server)`.
   - Implement `NewHTTPMux` using `httptogrpc.NewGRPCMux(conn)` and `grpcMux.AddService`. Include operation path routing if the API uses standard long-running operations.
   - Include magic comments for generation (e.g., `// +tool:mockgcp-service`).
5. Add the primary resource to the list of resources in `config/tests/samples/create/harness.go`, so that we actually run the test.
6. Make sure the service is in the list of services in `mockgcp/mockserviceusage/knownservices.go`.
7. You can run `dev/ci/presubmits/tests-e2e-fixtures-<service>` to run the test.

Use `mockgcp/.gemini/skills/add-new-mockgcp-resource/journal.md` for keeping notes and temporary insights when adding new resources.