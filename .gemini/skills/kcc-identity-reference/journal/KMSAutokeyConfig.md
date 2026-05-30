# KMSAutokeyConfig Identity and Reference Implementation

- The resource is a singleton under a folder: `folders/{folder}/autokeyConfig`.
- It does not have a `ResourceID` field in its spec.
- The `gcpurls.Template` used is `folders/{folder}/autokeyConfig`.
- Both `v1alpha1` and `v1beta1` were updated for consistency.
- A local `valueOf` helper was required in `autokeyconfig_reference.go` because it's used by other resources in the same package (like `KMSKeyHandle`) and I removed it from its original location. It's better to use `common.ValueOf` but that would require updating other resources which might be out of scope.
- `mapper.generated.go` required manual update because it was using the removed `ParseKMSAutokeyConfigExternal` function.
