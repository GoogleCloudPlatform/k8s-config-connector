# Fuzzing ComputeHTTPSHealthCheck

The `ComputeHTTPSHealthCheck` resource represents the legacy GCP HTTPS Health Check API. Under the hood of the direct controller, this is mapped to the modern `pb.HealthCheck` protobuf type from Compute Engine.

## Observations & Implementation Details

1. **Mapping Nested Structures:**
   The KRM type `ComputeHTTPSHealthCheckSpec` declares several fields at its top level, such as `Host`, `Port`, and `RequestPath`. However, in the GCE protobuf representation `pb.HealthCheck`, these fields are nested inside a sub-struct `pb.HTTPSHealthCheck` via the `HttpsHealthCheck` field.
   Because the code generator is flat and greedy, it missed this nesting and ignored those fields entirely. We solved this by implementing hand-coded mapper functions (`ComputeHTTPSHealthCheckSpec_v1beta1_FromProto`, `ComputeHTTPSHealthCheckSpec_v1beta1_ToProto`, etc.) in `computehttpshealthcheck_mapper.go` which handle this nesting manually.

2. **Handling Hardcoded Default Fields (`Type`):**
   In the legacy API, the health check is always of type `"HTTPS"`. Hence, in `ToProto` we hardcode `out.Type = direct.PtrTo("HTTPS")`. During round-trip fuzzing, this can cause a mismatch because the generated input `p1` might have `Type` unset, whereas the mapped/unmapped `p2` always has `Type` set to `"HTTPS"`.
   We solved this by implementing a `FilterSpec` function that sets `in.Type = direct.PtrTo("HTTPS")` on the incoming proto.

3. **Handling Empty Nested Structs (`https_health_check: {}`):**
   The fuzzer randomly populates fields on the sub-struct, but since unmapped fields are ignored (cleared), they are removed from `p1`. If all three mapped fields (`host`, `port`, `request_path`) are empty, the parent `.https_health_check` field of `p1` is left as an empty struct `{}`.
   However, `ToProto` does not create `pb.HTTPSHealthCheck` at all if all three KRM fields are nil, meaning `p2.HttpsHealthCheck` is `nil`.
   This results in a mismatch between `p1` (empty struct `{}`) and `p2` (nil).
   We resolved this by writing a check in `FilterSpec` to clear `in.HttpsHealthCheck` to `nil` if all mapped nested fields on it are empty.
