# ComputeBackendServiceSignedURLKey Implementation Journal

## Observations
- **GCP Subresource Mapping**: `ComputeBackendServiceSignedURLKey` is not a standalone top-level REST resource in Compute Engine, but rather a subresource managed via `addSignedUrlKey` and `deleteSignedUrlKey` methods on a `BackendService`. In proto definitions, it is represented by the `SignedUrlKey` message.
- **Generator defaults**: The direct type generator default scaffolding includes `Location` field in `Spec` for `google.cloud.compute.v1` types. However, since the existing TF-based CRD has no `Location` field (backend service signed URL keys are global), we manually removed `Location` to guarantee strict schema compatibility and avoid breaking existing users.
- **Secret references**: The `keyValue` field contains a secret-key reference (`valueFrom.secretKeyRef`). In direct resources, we map this to `k8sv1alpha1.SecretKeyRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1` to ensure correct validation and deepcopy generation behavior.
