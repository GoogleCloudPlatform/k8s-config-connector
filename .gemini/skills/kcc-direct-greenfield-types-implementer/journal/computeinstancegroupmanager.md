# ComputeInstanceGroupManager Greenfield Types Implementation Journal

## Observations & Implementation Details

1. **Map Type Support in Direct Generator**:
   - The direct code generator does not support `map[string]message` fields out of the box, noting `TODO: unsupported map type with key string and value message`.
   - To resolve this for `StatefulPolicy` and `StatefulPolicyPreservedState`, we hand-wrote the types using standard Go `map[string]T` maps and wrote custom `FromProto` and `ToProto` mapping functions in `pkg/controller/direct/compute/instancegroupmanager_mapper.go`.

2. **Acronym Linter vs. CRD Backward Compatibility**:
   - The acronym linter requires uppercase IPs (e.g., `ExternalIPs`, `InternalIPs`), but the existing DCL CRD schema previously used `externalIps` and `internalIps` in JSON.
   - We achieved seamless compliance and compatibility by naming the Go struct fields `ExternalIPs` and `InternalIPs` while overriding the JSON tags to `externalIps` and `internalIps` (e.g. `json:"externalIps,omitempty"`).

3. **OpenAPI Structural Validation**:
   - The generator initially produced an empty struct for `InstanceGroupManagerInstanceFlexibilityPolicy` because it only contained unsupported map fields.
   - OpenAPI validation in Kubernetes rejects empty structs with no properties unless `x-kubernetes-preserve-unknown-fields` is set. To ensure CRD validity, we removed `InstanceFlexibilityPolicy` from the Spec for now, as it was not supported by the old KCC CRD either.
