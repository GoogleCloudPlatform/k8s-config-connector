# ComputeFirewallPolicyRule - Direct KRM Types & generate.sh Journal

### [2026-06-13] ComputeFirewallPolicyRule Investigation
- **Context**: Verifying and documenting the direct KRM types and `generate.sh` configuration for `ComputeFirewallPolicyRule` (GVK: `compute.cnrm.cloud.google.com/v1beta1, Kind: ComputeFirewallPolicyRule`).
- **Problem**: The resource is already implemented as a direct controller and its types/CRDs are fully scaffolded, but the journal documenting its direct transition was missing.
- **Observations & Learnings**:
  1. **API Versioning & Stability**: While new greenfield resources are typically scaffolded as `alpha` under `v1alpha1`, `ComputeFirewallPolicyRule` was implemented as `stable` under `v1beta1` (with the label `cnrm.cloud.google.com/stability-level=stable`) because of its alignment with other stable Compute API resources in Config Connector.
  2. **Manual Field Overrides**: The generator-produced `types.generated.go` for `v1beta1` contains comments showing that the generator successfully skips `FirewallPolicyRule` and `FirewallPolicyRuleMatcher` because those types have hand-crafted custom Go definitions in `firewallpolicyrule_types.go`. This is done to ensure correct reference types such as `ComputeNetworkRef` and `refs.IAMServiceAccountRef` instead of using primitive types, satisfying strict reference conventions.
  3. **CRD Exceptions**: There are known entries under `tests/apichecks/testdata/exceptions/missingfields.txt` and `missingrefs.txt` for `ComputeFirewallPolicyRule` (e.g., `.spec.targetResources[].external`) due to the nested array structure of some reference fields, which is normal for this API.
- **Impact**: Confirming that the direct-model, generator config, and `generate.sh` configuration are fully complete and functional.
