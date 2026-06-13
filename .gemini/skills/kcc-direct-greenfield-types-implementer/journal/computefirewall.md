# ComputeFirewall Greenfield KRM Types Journal

## Observations & Learnings
- **Global Resource Multi-hierarchy**: `ComputeFirewall` is a global resource in GCP under `projects/{project}/global/firewalls/{firewall}`. Unlike regional compute resources, we omitted the `Location` field from `ComputeFirewallSpec` and did not specify it in the URI template format, aligning perfectly with its global scope.
- **Reference & Slice Hand-written Mappers**: Ported KRM types `allow` and `deny` to slice-backed `FirewallAllow` and `FirewallDeny` structs respectively, to match the original JSON fields and maintain backward-compatibility. Since `Allowed` and `Denied` map to proto `Allowed` and `Denied` messages, custom slice mapping methods (`FirewallAllowSlice_v1beta1_FromProto` / `ToProto`) were implemented in a dedicated hand-written `firewall_mapper.go` file.
- **Service Account References**: Implemented standard `IAMServiceAccountRef` references for `sourceServiceAccounts` and `targetServiceAccounts` fields. These map from slice of references in KRM to string slice in GCP protos.
- **Nested Field Mapping**: Leveraged `controllerbuilder`'s nested-path mapping capability for `enableLogging` pointing to `Firewall.log_config.enable` which maps to proto `log_config` nested sub-structs smoothly.
