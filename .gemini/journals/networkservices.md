# Network Services Journal

### [2026-06-15] NetworkServicesAuthzExtension Scaffolding and References
- **Context**: Scaffolding direct types, IdentityV2, and Reference for `NetworkServicesAuthzExtension` (v1alpha1).
- **Problem**: The `service` field in `AuthzExtension` is a URL referencing a backend service (`projects/{project}/regions/{region}/backendServices/{backendService}` or `projects/{project}/global/backendServices/{backendService}`).
- **Solution**: Implemented this field as a proper KCC reference using `computev1beta1.ComputeBackendServiceRef` and named it `ServiceRef` in the Spec to follow the standard `Ref` suffix naming convention. Alphabetized the new missing-field exceptions in `alpha-missingfields.txt` to satisfy the API checks during PR validation before the controller is implemented.
- **Impact**: Ensures that the `NetworkServicesAuthzExtension` conforms to the KCC design standards on references (avoiding the missing references check failure) and allows future agents to seamlessly implement the reconciliation logic.
