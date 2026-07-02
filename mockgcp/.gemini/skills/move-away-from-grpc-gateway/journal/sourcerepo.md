# sourcerepo migration notes

The migration of the `sourcerepo` service away from `grpc-gateway` to `httptogrpc` is blocked.

Under the updated skill guidelines (Step 0: Verify gRPC client library existence), a migration can only proceed if an official gRPC Go client and protobuf module exists for the service under `cloud.google.com/go/`.

For Cloud Source Repositories (sourcerepo), which is a deprecated service, no such modern client library exists (e.g., under `cloud.google.com/go/sourcerepo/apiv1`). Since the modern Go client library package is not available, the migration is blocked and we must not perform it.
