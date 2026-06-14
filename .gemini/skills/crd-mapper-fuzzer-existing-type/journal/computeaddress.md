# ComputeAddress Direct KRM Transition Journal

## Learnings & Design Decisions

- **Spelling Alignment of Acronym Fields**:
  The KRM CRD fields were spelled `ipVersion` and `ipv6EndpointType`.
  The original proto definitions in `pb.Address` mapped these as `IPVersion` and `IPV6EndpointType`.
  By defining Go struct fields exactly as `IPVersion` and `IPV6EndpointType` with json tags `json:"ipVersion,omitempty"` and `json:"ipv6EndpointType,omitempty"`, we managed to:
  1. Have the mapper generator auto-map these fields perfectly without any handcoded mapping code needed.
  2. Maintain 100% schema compatibility because the JSON schema serialization uses the json tag names.

- **PrefixLength Type Alignment**:
  In the baseline CRD schema, `prefixLength` had type `integer`.
  In Go types, we initially declared it as `*int64`, which is the typical default for `type: integer` in KCC. However, the protobuf-generated field `PrefixLength` on `pb.Address` was typed as `*int32`.
  Changing our KRM Go type to `*int32` resolved all Go compiler type mismatch issues, while still generating an identical schema without any functional schema differences.

- **ObservedState Preservation**:
  The baseline CRD already contained `observedState` in `status`.
  To ensure strict schema-compatibility, we included `ComputeAddressObservedState` and nested `Address` under it.

- **Omission of ExternalRef**:
  Since `externalRef` was not in the baseline CRD, we omitted `ExternalRef` from the KRM `ComputeAddressStatus` Go struct to preserve 100% identical CRD schema.
