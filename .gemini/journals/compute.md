### [2026-06-13] Implementation of direct KRM types for ComputeExternalVPNGateway
- **Context**: Implementing direct KRM types for `ComputeExternalVPNGateway` under the `apis/compute/v1beta1` package.
- **Problem**: The resource is already a `v1beta1` resource in the terraform-based era. Generating it in `v1alpha1` (the default api-version for greenfield resources) would downgrade the API version for existing users, causing regression/breakages.
- **Solution**: Implemented the types under `apis/compute/v1beta1` instead of `v1alpha1`, ensuring backward compatibility for `v1beta1` and allowing the generator to clean up and migrate the CRD correctly. Added all nested proto structures like `ExternalVPNGatewayInterface` by declaring them in the spec and status.
- **Impact**: Enables migrating the `ComputeExternalVPNGateway` controller to direct while preserving the `v1beta1` schema compatibility and stability level (`stable`).
