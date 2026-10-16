# ComputeNetworkFirewallPolicy Fuzzer Journal

## Observations & Learnings

- **Fuzzer Implementation Details:** The KRM fuzzer for `ComputeNetworkFirewallPolicy` is located in `pkg/controller/direct/compute/computenetworkfirewallpolicy_fuzzer.go`.
- **Shared Protobuf Message:** `ComputeNetworkFirewallPolicy` uses `google.cloud.compute.v1.FirewallPolicy` (`pb.FirewallPolicy`), which is also shared with `ComputeFirewallPolicy` (hierarchical firewall policies).
- **Field Segregation:** Fields such as `.associations` and `.rules` are managed by separate child resources (`ComputeNetworkFirewallPolicyAssociation` and `ComputeNetworkFirewallPolicyRule`), and hierarchical/regional fields (`.display_name`, `.short_name`, `.region`, `.packet_mirroring_rules`, `.policy_type`) are safely marked as untriaged/unimplemented for network firewall policies.
- **Central Fuzz Testing Integration:** Confirmed that `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/compute"` is registered and imported under `pkg/controller/direct/register/register.go` so the fuzzer runs in the central fuzz test suite.
- **Detailed Comparisons Documentation:** Added structured ASCII table documentation inside the fuzzer file explicitly mapping each KRM Spec and Status field to its Protobuf equivalent and fuzzer status.
- **Testing Verification:** Executed the centralized fuzz testing framework targeting `ComputeNetworkFirewallPolicy` (`FOCUS=ComputeNetworkFirewallPolicy go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`), and confirmed that the Spec and Status round-trips pass flawlessly.
