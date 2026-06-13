# ComputeInstance Direct KRM Types Implementation Journal

## Resource Details
- **Kind**: `ComputeInstance`
- **GCP Resource**: `Instance`
- **Service**: `google.cloud.compute.v1`
- **KCC API Version**: `v1beta1`

## Implementation Observations & Key Decisions

### 1. Schema Compatibility & Reference Fields (`InstanceResourceRef`)
- **Problem**: The standard `k8sv1alpha1.ResourceRef` declares a `kind` property. However, the pre-existing baseline `ComputeInstance` CRD (generated previously via Terraform) did not have a `kind` property for its reference fields (except for `NetworkIpRef`). Using the standard `ResourceRef` would alter the schema and break backward compatibility, failing the `dev/tasks/diff-crds` presubmit check.
- **Solution**: We implemented a custom reference type `InstanceResourceRef` in `apis/compute/v1beta1/instance_types.go` that does not declare a `kind` property. This custom type is used for the 11 reference fields that had no `kind` defined, while `NetworkIpRef` continues to use `k8sv1alpha1.ResourceRef` to align precisely with the baseline CRD.

### 2. Manual Mapper Stubs
- **Problem**: The automatic mapper generator requires matching signatures for complex types, and could not resolve them for our customized structs without manual intervention.
- **Solution**: We hand-coded the mapper function stubs in `pkg/controller/direct/compute/instance_mapper.go`. This includes the main `FromProto` and `ToProto` specs/status stubs, and sub-struct stubs like `InstanceMetadata_v1beta1_FromProto`. This ensures that `mapper.generated.go` compiles cleanly and the direct controller can build correctly.

### 3. Verification & CI Validation
- We successfully validated the schema with `dev/tasks/diff-crds` to ensure zero schema drift, and ran `go vet ./...` and `make fmt` to verify formatting and compilation.
