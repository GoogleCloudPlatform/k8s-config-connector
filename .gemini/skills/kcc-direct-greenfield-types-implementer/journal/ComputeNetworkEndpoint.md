# ComputeNetworkEndpoint Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Compatibility**:
   - The initial generator output produced `Location string` in `ComputeNetworkEndpointSpec` to map the zonal parent. However, the existing `ComputeNetworkEndpoint` CRD specifies `zone: string` and does not define a `location` field. To ensure backward-compatibility and avoid breaking changes to the Custom Resource Definition schema, we hand-coded `Zone string` (along with `+required` and appropriate json tags) in place of `location`.

2. **Reference Implementation**:
   - `ComputeNetworkEndpoint` relies on references to `ComputeNetworkEndpointGroup` (`networkEndpointGroupRef`) and `ComputeInstance` (`instanceRef`).
   - While `InstanceRef` was already defined in the `computev1beta1` package, `ComputeNetworkEndpointGroupRef` was missing.
   - We created `computenetworkendpointgroup_identity.go` and `computenetworkendpointgroup_reference.go` under `apis/compute/v1beta1/` to properly expose `ComputeNetworkEndpointGroupRef` using `identity.IdentityV2` and `refs.Ref` templates, including fallback parsing logic for the TF/DCL-based `ComputeNetworkEndpointGroup` resource.
   - We verified reference resolution through unit tests in `computenetworkendpointgroup_identity_test.go` confirming both canonical relative resource paths and full resource URLs are successfully normalized.
