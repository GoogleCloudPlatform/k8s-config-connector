### [2026-09-29] Connectors Direct Controller Implementation
- **Context**: Greenfield implementation of direct controller, E2E fixtures, and fuzzer for `ConnectorsConnection` (connectors.cnrm.cloud.google.com/v1alpha1).
- **Problem**:
  1. The Google Cloud Connectors API gRPC endpoint (`connectors.googleapis.com:443`) returned `PermissionDenied: Read access to project was denied` for regional connection resources.
  2. The Connectors API Go client (`google.golang.org/api/connectors/v1`) uses camelCase json field names (e.g., `connectorVersion`, `serviceAccount`, `configVariables`), so `protojson.MarshalOptions{UseProtoNames: false}` is required when translating protobuf structs to REST structs.
  3. Connectors automatically populates server defaults for `nodeConfig`, `lockConfig`, and `authConfig` on creation, which required defaulting in the diff `compare` method to avoid false update diffs during re-reconciliation.
- **Solution**:
  1. Used `google.golang.org/api/connectors/v1` REST discovery client (`NewService`) with `m.config.RESTClientOptions()`, which handles regional resources and generates HTTP golden traffic recording cleanly.
  2. Set `UseProtoNames: false` in `ProtoToREST` helper.
  3. Defaulted unconfigured spec fields (`NodeConfig`, `LockConfig`, `AuthConfig`) from `maskedActual` in `compare()`.
- **Impact**: Provides clear guidance on client choice and JSON translation for future Connectors resource implementations.
