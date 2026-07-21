### [2026-05-12] NetworkConnectivityRegionalEndpoint lack of mutable fields
- **Context**: Implementing `NetworkConnectivityRegionalEndpoint` using direct controller pattern.
- **Problem**: The GCP API definition for `RegionalEndpoint` specifies no input fields other than the resource ID and parent URI. `Update`/`Patch` operations are completely omitted from the API surface.
- **Solution**: The `Update()` method in the direct controller adapter was simplified to skip any GCP API `Patch` calls since no such method exists and no fields can be updated. It just updates the status observed state.
- **Impact**: Future agents working on networkconnectivity or similar singleton-like resources should be aware that missing `Patch` or mutable fields in the API definition necessitates an `Update()` method that only performs status updates.

### [2026-07-02] Handling Underscored URL Placeholders in gcpurls.Template Matching
- **Context**: Implementing IdentityV2 for `NetworkConnectivityMulticloudDataTransferConfig` (Issue #10290).
- **Problem**: The GCP URL template has a snake_case variable name: `projects/{project}/locations/{location}/multicloudDataTransferConfigs/{multicloud_data_transfer_config}`. When `gcpurls.Template` initializes, it does case-insensitive, alphanumeric matching against the Go Identity struct's fields. A simple field name like `MulticloudDataTransferConfig` resulted in matching `"multiclouddatatransferconfig"`, which failed to match the key `"multicloud_data_transfer_config"` because of the underscores, causing a runtime panic during package initialization.
- **Solution**: Named the field in the Identity Go struct with underscores matching the URL template (e.g. `Multicloud_data_transfer_config string`). This ensures that `strings.ToLower()` matches `"multicloud_data_transfer_config"` exactly, avoiding the initialization panic.
- **Impact**: When implementing IdentityV2 with URL templates, ensure that the fields in the Identity struct exactly match any snake_case or underscored variables present in the GCP API URL path placeholder.
