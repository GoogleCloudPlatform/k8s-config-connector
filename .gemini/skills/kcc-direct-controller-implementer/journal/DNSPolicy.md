# DNS Service Journal

### 2026-06-08 DNSPolicy Direct Controller Implementation
- **Context**: Migration of `DNSPolicy` under `dns.cnrm.cloud.google.com/v1beta1` to the Direct controller model and addition of `mockgcp` DNS Policies support.
- **Problem**: 
  1. The underlying REST-based client `google.golang.org/api/dns/v1` doesn't use protobufs, making standard gRPC or pb-based diffing tools (`tags.DiffForTopLevelFields`) non-applicable directly.
  2. During resource deletion, the legacy/old controller (the Terraform-based reconciler tested in parallel) unbinds/clears networks by sending a `PATCH` request to GCP. However, because names are excluded from the update request body, a naive `mockgcp` implementation that blindly clones the request body would overwrite/clear the stored policy's `Name` with an empty string, causing subsequent `tpgresource.ReplaceVars` interpolation to produce an invalid delete URL path `projects/{{project}}/policies/` and a 404 response.
- **Solution**:
  1. Utilized a dedicated `comparePolicy` helper function which round-trips actual policy fields via `DNSPolicySpec_FromAPI`/`DNSPolicySpec_ToAPI` and employs the JSON-based `diffs.GoogleAPI.Diff` tool.
  2. Implemented `mockgcp/mockdns/policy.go` from scratch.
  3. Ensured `mockgcp` update/patch handlers explicitly preserve the policy's name field (`updated.Name = existing.Name`) if it is omitted or empty in the update/patch request body, ensuring no ID leaks/corruptions.
- **Impact**: Flawless execution of E2E fixture tests in parallel for both the legacy Terraform controller and the direct controller, resulting in a perfect PASS.
