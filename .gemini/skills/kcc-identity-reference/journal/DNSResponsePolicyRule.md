### [2026-06-18] DNSResponsePolicyRule Identity & Refs
- **Context**: Refactoring DNSResponsePolicyRule to follow the canonical Identity and Reference patterns using `gcpurls.Template`.
- **Observations & Decisions**:
    1. **Parent Reference Resolving**: The resource spec for `DNSResponsePolicyRule` has a string-typed field `responsePolicy` rather than a dedicated `ResponsePolicyRef` struct. To keep the schema backwards-compatible and avoid changing KRM types prematurely:
        - In `getIdentityFromDNSResponsePolicyRuleSpec`, we first check if the `responsePolicy` string parses successfully as an external reference.
        - If it does not parse as an external reference, we assume it represents a local KRM `DNSResponsePolicy` name in the same namespace, fetch it using the `client.Reader`, and resolve its identity to construct the correct rule identity.
    2. **Multi-Format Templates**: Defined two `gcpurls.Template` formats:
        - Primary location-based format: `projects/{project}/locations/{location}/responsePolicies/{responsePolicy}/rules/{rule}`
        - Fallback location-less format: `projects/{project}/responsePolicies/{responsePolicy}/rules/{rule}`
    3. **Registry Exception Registration**: Added `"//dns.googleapis.com/projects/{}/responsePolicies/{}/rules/{}"` to the ignored templates list in `pkg/gcpurls/registry_test.go` to satisfy CAI alignment validation checking.
- **Impact**: Enables standard identity extraction and reference support for DNSResponsePolicyRule, unblocking future migrations and direct controller work while maintaining full schema compatibility.
