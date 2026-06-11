### [2026-06-11] DNSResponsePolicy Identity & Refs

- **Context**: Moving `DNSResponsePolicy` and `DNSResponsePolicyRule` to the identity and refs pattern.
- **Observations & Changes**:
  1. `DNSResponsePolicy` and `DNSResponsePolicyRule` were originally in the autogen allowlist (`scripts/resource-autogen/allowlist/allowlist.go`).
  2. We moved both resources out of the allowlist into `/workspaces/k8s-config-connector/config/servicemappings/dns.yaml` so they are statically managed.
  3. `DNSResponsePolicy` has custom Go types under `apis/dns/v1alpha1/dnsresponsepolicy_types.go`, while `DNSResponsePolicyRule` has generated Go types.
  4. In `/workspaces/k8s-config-connector/config/servicemappings/dns.yaml`, we configured `DNSResponsePolicyRule` with `responsePolicyRef` referencing `DNSResponsePolicy`.
  5. We updated `apis/dns/v1alpha1/dnsresponsepolicy_types.go` to support `networkRef` and `gkeClusterNameRef`, making the original string fields optional.
  6. We also updated `config/servicemappings/dns.yaml` to specify the reference mappings for `networks.network_url` -> `networkRef` and `gke_clusters.gke_cluster_name` -> `gkeClusterNameRef`.
  7. We consolidated/updated the static `zz_generated.v1alpha1.mappers.go` mapper file to support pointers `*string` using `direct.LazyPtr` and `direct.ValueOf`.
  8. We ran `make manifests`, `make generate-go-client`, and verified that all API schema checks in `tests/apichecks` pass successfully.
