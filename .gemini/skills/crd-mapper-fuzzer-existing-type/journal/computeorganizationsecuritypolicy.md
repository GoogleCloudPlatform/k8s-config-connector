# ComputeOrganizationSecurityPolicy Journal

## Implementation Summary

`ComputeOrganizationSecurityPolicy` in GCP is represented under the hood by the `SecurityPolicy` protobuf message (from package `google.cloud.compute.v1`). This was already utilized by the `ComputeSecurityPolicy` (v1beta1) resource, making this a multi-resource mapping scenario where two separate KRM resources (`ComputeSecurityPolicy` and `ComputeOrganizationSecurityPolicy`) map to the exact same proto message type (`SecurityPolicy`).

## Key Learnings & Decisions

### 1. Version Scoping & File Naming
- Since `ComputeSecurityPolicy` resides in `v1beta1/securitypolicy_types.go`, we placed the new `ComputeOrganizationSecurityPolicy` types inside `v1alpha1/securitypolicy_types.go` matching the lowercase of the proto message `SecurityPolicy`. This maintains packages isolation and complies with code generation expectations.

### 2. Handcoded Mappers
- `ComputeOrganizationSecurityPolicySpec` features custom fields `displayName` and `parent` which map to the proto fields `ShortName` and `Parent` respectively. Because these names do not match, we hand-coded the mapping functions `ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto` and `ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto` in `pkg/controller/direct/compute/mappers.go`.
- The code generator recognized these and correctly commented out the generated versions inside `mapper.generated.go`.

### 3. Fuzzer
- We implemented `pkg/controller/direct/compute/computeorganizationsecuritypolicy_fuzzer.go` using type-safe helpers like `f.SpecField` and `f.Unimplemented_Identity`, successfully matching the fuzzer standards of the repository.

### 4. Strict CRD Schema Validation
- The baseline CRD was compared using `dev/tasks/diff-crds`. We ensured exact alignment by omitting `ProjectRef` (which wasn't in the baseline CRD) and adding `// +required` annotation on the `Spec` field. The schema comparison results are 100% clean and identical.
