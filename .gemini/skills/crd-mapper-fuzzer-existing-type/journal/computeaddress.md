# ComputeAddress Journal

## Observations
- Configured `generate.sh` under `apis/compute/v1beta1/generate.sh` to include `ComputeAddress:Address`.
- Defined KRM types in `apis/compute/v1beta1/address_types.go` that strictly mirror the original OpenAPI CRD schema for `ComputeAddress` to maintain 100% backward compatibility.
- Excluded `ProjectRef` from `ComputeAddressSpec` because the baseline CRD did not contain `spec.projectRef`. This ensures zero schema additions.
- Configured correct kubebuilder labels on the `ComputeAddress` struct, including `"cnrm.cloud.google.com/tf2crd=true"` and `"cnrm.cloud.google.com/stability-level=stable"` to preserve legacy controller registration compatibility.
- Fixed shortNames mapping to exactly match the original CRD by specifying `gcpcomputeaddress;gcpcomputeaddresses` instead of the default singular-pluralized `gcpcomputeaddresss`.
- Annotated `ComputeAddressStatus` with `// +kcc:status:proto=google.cloud.compute.v1.Address` to automatically generate `FromProto` and `ToProto` status mappings under `pkg/controller/direct/compute/mapper.generated.go`.
- Implemented round-trip fuzzer for `ComputeAddress` in `pkg/controller/direct/compute/computeaddress_fuzzer.go` following the requested type-safe fuzzer helper methods.
- Validated CRD schema compatibility using `dev/tasks/diff-crds`, yielding an empty diff (perfect schema match!).
- Verified clean compilation with `go vet ./...`.
