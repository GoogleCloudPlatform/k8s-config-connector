# networkservices migration

- Similar to datastream, `networkservices` migration involved removing the `RewriteError` logic.
- `networkservices` protos were located in the `generate-grpc-for-google-protos` target in `mockgcp/Makefile`, not `gen-proto-no-fixup`.
