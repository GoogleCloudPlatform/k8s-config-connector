# Journal - Move away from grpc-gateway

## 2026-05-16 - NetApp

- NetApp was already partially moved to use `cloud.google.com/go/netapp/apiv1/netapppb` and `httptogrpc`.
- The cleanup of `mockgcp/Makefile` and deletion of `mockgcp/generated/mockgcp/cloud/netapp` was missing.
- No additional complexities were found during the migration.
