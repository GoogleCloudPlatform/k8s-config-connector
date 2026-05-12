### [2026-05-12] NetworkConnectivityRegionalEndpoint lack of mutable fields
- **Context**: Implementing `NetworkConnectivityRegionalEndpoint` using direct controller pattern.
- **Problem**: The GCP API definition for `RegionalEndpoint` specifies no input fields other than the resource ID and parent URI. `Update`/`Patch` operations are completely omitted from the API surface.
- **Solution**: The `Update()` method in the direct controller adapter was simplified to skip any GCP API `Patch` calls since no such method exists and no fields can be updated. It just updates the status observed state.
- **Impact**: Future agents working on networkconnectivity or similar singleton-like resources should be aware that missing `Patch` or mutable fields in the API definition necessitates an `Update()` method that only performs status updates.
