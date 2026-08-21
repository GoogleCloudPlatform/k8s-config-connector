# ComputeAddress Identity and Reference Journal

## Observations & Learnings

- **Global and Regional Dual Identity Templates:** ComputeAddress, like ForwardingRule, supports both global and regional GCP URL paths (under `global/addresses/...` and `regions/{{region}}/addresses/...`). We successfully modeled both paths using two distinct `gcpurls.Template` instances and dynamically matched the incoming reference format in `FromExternal`.
- **Package Integration & Dependency Cleanup:** Prior to this work, multiple external APIs (such as DNS RecordSet and NetworkConnectivity RegionalEndpoint) depended on a legacy, hand-rolled `ComputeAddressRef` in `apis/refs/v1beta1/computerefs.go` that did not support full `refs.Ref` or `identity.IdentityV2` integration. We moved the reference type definition to `apis/compute/v1beta1/computeaddress_reference.go`, implemented `refs.Ref`, and fully updated DNS, ForwardingRule, and NetworkConnectivity to import the standard `computev1beta1.ComputeAddressRef` type.
- **Normalization Replaced Manual Resolution:** By implementing the `refs.Ref` interface and standard `Normalize()` flow on `ComputeAddressRef`, we simplified dependency resolution in direct controllers. In NetworkConnectivity, we replaced manual `ResolveComputeAddress` invocations with a unified `Normalize` call.
- **No Schema Modifications:** The schema of all modified resources was preserved completely untouched, keeping backward compatibility and adhering strictly to KCC engineering mandates.
