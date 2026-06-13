# ComputeAddress Journal

## Field Name Acronym Casing (Proto vs KRM)

When the controller builder maps KRM types to GCP Proto types, Go protobuf's acronym upper-casing (e.g., `IPVersion` instead of `IpVersion` and `IPV6EndpointType` instead of `Ipv6EndpointType`) can result in near-miss warnings. By matching KRM's Go field names with the protobuf compiler's capitalized field names while preserving the lowerCamelCase JSON tags, the generator maps fields automatically without warnings.

## Backward Compatibility with PrefixLength (*int32 vs *int64)

In the standard GCP Compute proto, `PrefixLength` is modeled as an `int32` pointer. However, the existing KCC client-generated CRD represents this as `integer` with an `int64` format (mapping to `*int64` in Go). To ensure 100% backward compatibility of the CRD schema, we:
1. Retained `PrefixLength *int64` in `address_types.go`.
2. Implemented hand-written mapper functions (`ComputeAddressSpec_v1beta1_FromProto` and `ComputeAddressSpec_v1beta1_ToProto`) inside a new file `pkg/controller/direct/compute/address_mapper.go` to explicitly convert the type pointer from/to `int32` and `int64`.
3. The generator automatically detected these manual functions and skipped generation of conflicting stubs in `mapper.generated.go`.
