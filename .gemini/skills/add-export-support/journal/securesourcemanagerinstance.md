# Export Support Journal for SecureSourceManagerInstance

## Observations
- Discovered and fixed a critical bug in the pre-existing `Export` implementation where `u.Object = uObj` was assigned *after* `SetName` and `SetGroupVersionKind`, thereby overwriting/wiping out the serialization of `name`, `apiVersion`, and `kind` in the returned `unstructured.Unstructured` object. This was causing downstream export clients like `SupportsIAM` to fail with "unable to get service mapping: no mapping with name '' found" because the Group and Kind of the resource were received as empty strings.
- Moving the `u.Object = uObj` assignment *before* `SetName` and `SetGroupVersionKind` successfully resolved the issue and generated the correct `_exported.yaml` files.
- `SecureSourceManagerInstance` uses reference-based project binding via its custom `v1alpha1.ResourceRef` fields, meaning we should populate `obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}` rather than setting the project ID annotation.
- Resolved integration testing setup issues on headless containers where etcd is not natively in `/usr/local/kubebuilder/bin` by running the tests using the correct dynamically determined local asset path under `KUBEBUILDER_ASSETS=/root/.local/share/kubebuilder-envtest/k8s/1.36.0-linux-amd64` when `E2E_KUBE_TARGET=envtest`.
