## 2026-05-17 - batch migration

- When migrating `mockbatch` away from `grpc-gateway`, I found that the service was also listed in the `generate-grpc-for-google-protos` target in `mockgcp/Makefile`. It should be removed from both `gen-proto-no-fixup` and `generate-grpc-for-google-protos`.
- The migration was straightforward and the official `batchpb` types were compatible with the existing mock implementation.