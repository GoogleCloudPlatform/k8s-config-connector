### [2026-05-12] NetworkConnectivityRegionalEndpoint lack of mutable fields
- **Context**: Implementing `NetworkConnectivityRegionalEndpoint` using direct controller pattern.
- **Problem**: The GCP API definition for `RegionalEndpoint` specifies no input fields other than the resource ID and parent URI. `Update`/`Patch` operations are completely omitted from the API surface.
- **Solution**: The `Update()` method in the direct controller adapter was simplified to skip any GCP API `Patch` calls since no such method exists and no fields can be updated. It just updates the status observed state.
- **Impact**: Future agents working on networkconnectivity or similar singleton-like resources should be aware that missing `Patch` or mutable fields in the API definition necessitates an `Update()` method that only performs status updates.

### [2026-07-23] Services map field in MulticloudDataTransferConfig ObsState
- **Context**: Implementing `NetworkConnectivityMulticloudDataTransferConfig` types, identity, and mapping.
- **Problem**: The GCP API definition for `MulticloudDataTransferConfig` has a map field `services` which maps service names to `StateTimeline` nested messages. This resulted in undefined `Services_FromProto` and `Services_ToProto` helper functions in `mapper.generated.go`.
- **Solution**: Hand-wrote `Services_FromProto`, `Services_ToProto`, `StateMetadata_FromProto`, and `StateMetadata_ToProto` mapping helper functions at the end of `mapper.go`.
- **Impact**: Map fields pointing to nested messages often require manually implemented mapper helpers in `mapper.go` to avoid compilation/vet errors due to undefined generated symbols.

### [2026-09-25] NetworkConnectivityTransport mutable fields and defaulting
- **Context**: Implementing `NetworkConnectivityTransport` direct controller and recording E2E fixtures against real GCP.
- **Problem**:
  1. GCP server returns `bandwidth can only be updated for ACTIVE transports` when attempting to patch `bandwidth` while the transport is in `PENDING_KEY` state.
  2. GCP server defaults `stack_type` to `IPV4_ONLY` if unspecified in spec.
  3. `generatedActivationKey` (provider key) and `peeringNetwork` (generated VPC peering URI) are volatile server-generated values.
- **Solution**:
  1. In E2E update tests, updated `description` and `labels` rather than `bandwidth`.
  2. Populated default `StackType = "IPV4_ONLY"` in `populateDefaults` inside `compareTransport`.
  3. Normalized `generatedActivationKey` and `peeringNetwork` in `mockgcp/mocknetworkconnectivity/normalize.go` and `tests/e2e/normalize.go`.
- **Impact**: When testing or updating transports, only mutable fields valid in `PENDING_KEY` state (`description`, `labels`) can be modified prior to provider activation. Defaulted fields like `stackType` must be aligned in `compareTransport`.

