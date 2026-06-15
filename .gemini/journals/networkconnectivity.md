### [2026-05-12] NetworkConnectivityRegionalEndpoint lack of mutable fields
- **Context**: Implementing `NetworkConnectivityRegionalEndpoint` using direct controller pattern.
- **Problem**: The GCP API definition for `RegionalEndpoint` specifies no input fields other than the resource ID and parent URI. `Update`/`Patch` operations are completely omitted from the API surface.
- **Solution**: The `Update()` method in the direct controller adapter was simplified to skip any GCP API `Patch` calls since no such method exists and no fields can be updated. It just updates the status observed state.
- **Impact**: Future agents working on networkconnectivity or similar singleton-like resources should be aware that missing `Patch` or mutable fields in the API definition necessitates an `Update()` method that only performs status updates.

### [2026-06-15] Unsupported Map Types in Controller Generator
- **Context**: Implementing direct types for `NetworkConnectivityMulticloudDataTransferConfig`.
- **Problem**: The GCP resource definition has a `map<string, StateTimeline>` field called `services`. The generator outputs an "unsupported map type with key string and value message" warning and comments it out in `types.generated.go`.
- **Solution**: Since all fields of `StateTimeline` and its sub-message `StateMetadata` are output-only (annotated as `Output only`), they do not need to be represented as input spec fields. Therefore, the field is safely skipped from the Spec, while standard fields like `description` and `labels` are mapped to KRM `Spec` and output-only metadata to `ObservedState`.
- **Impact**: When designing KRM spec and status, verify if fields skipped by the generator due to map types contain any writable input fields. If they are purely output-only, they can be omitted or handled in ObservedState.
