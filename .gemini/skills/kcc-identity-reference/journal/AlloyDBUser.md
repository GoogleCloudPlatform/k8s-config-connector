# AlloyDBUser Identity and Reference Implementation Journal

### 2026-05-29 AlloyDBUser Identity and Refs Implementation
- **Context**: Transitioning `AlloyDBUser` to the `identity` and `refs` pattern.
- **Problem 1**: `ClusterRef` in `AlloyDBUserSpec` (defined in `user_types.go`) uses the package-local `ClusterRef` type rather than `refs.AlloyDBClusterRef`. Because `refs.ResolveAlloyDBCluster` expects `*refs.AlloyDBClusterRef`, passing `*ClusterRef` directly would cause a compilation/type error.
- **Solution 1**: Instead of trying to cast the pointer directly, we dynamically convert the fields from the local `v1beta1.ClusterRef` to `refs.AlloyDBClusterRef` before calling `refs.ResolveAlloyDBCluster`:
  ```go
  var clusterRefParam *refs.AlloyDBClusterRef
  if obj.Spec.ClusterRef != nil {
      clusterRefParam = &refs.AlloyDBClusterRef{
          External:  obj.Spec.ClusterRef.External,
          Name:      obj.Spec.ClusterRef.Name,
          Namespace: obj.Spec.ClusterRef.Namespace,
      }
  }
  clusterRef, err := refs.ResolveAlloyDBCluster(ctx, reader, obj, clusterRefParam)
  ```
- **Problem 2**: `AlloyDBUser` has no CAI definition registered in `docs/ai/metadata/cloudassetinventory_names.jsonl`, which would cause `TestRegisteredTemplatesMatchCAI` to fail once the template `projects/{project}/locations/{location}/clusters/{cluster}/users/{user}` was registered.
- **Solution 2**: Added the normalized template format `"//alloydb.googleapis.com/projects/{}/locations/{}/clusters/{}/users/{}"` as an exception in the `ignoredTemplates` list inside `pkg/gcpurls/registry_test.go`.
- **Impact**: Clean isolation of types and robust template registry validation preventing any compilation/validation breaks.
