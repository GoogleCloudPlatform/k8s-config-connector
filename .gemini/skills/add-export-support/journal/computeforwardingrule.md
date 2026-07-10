# Export Support Journal: ComputeForwardingRule

During the implementation of export support for `ComputeForwardingRule`, we made the following observations:

1. **Alignment on IP Address:**
   - In `FromProto`, we aligned the mapping for `ipAddress`. If the IP address returned from GCP starts with `projects/` or contains `/addresses/`, it represents a reference to a `ComputeAddress` resource rather than a raw IP string. We updated `ComputeForwardingRuleSpec_IpAddress_FromProto` to map these reference paths to `AddressRef` instead of `Ip`. This resolved differences in exported objects where the old controller mapped the reference path to `addressRef.external` but the direct controller initially mapped it to `ip`.

2. **Legacy Controller Target Mapping Bug:**
   - There remains a minor, correct difference in the target reference field for some test cases (such as TCP and gRPC forwarding rules). The legacy controller erroneously exported target TCP proxy and gRPC proxy references under `targetHTTPProxyRef` instead of `targetTCPProxyRef` and `targetGRPCProxyRef` respectively. The direct controller correctly maps these to `targetTCPProxyRef` and `targetGRPCProxyRef` based on the resource URI type, resulting in a cleaner and more accurate KRM representation.
