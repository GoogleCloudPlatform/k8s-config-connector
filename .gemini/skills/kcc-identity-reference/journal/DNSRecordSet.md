### [2026-06-10] DNSRecordSet Identity & Refs
- **Context**: Moving DNSRecordSet to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. DNSRecordSet was entirely missing identity or reference files.
    2. The standard REST API resource URL contains `{name}/{type}` at the end, but per reviewer feedback we dropped the `{type}` segment from the identity/reference template format so that it stays standard and simpler (e.g. `projects/{project}/managedZones/{managedZone}/rrsets/{name}`).
    3. Even without type, the template doesn't include the `/locations/{location}` segment present in the CAI format, so it still requires an exclusion in the registry test.
- **Solution**:
    1. Implemented `apis/dns/v1beta1/dnsrecordset_identity.go` utilizing `gcpurls.Template` with the projects/project/managedZones/managedZone/rrsets/name pattern (dropping `{type}`).
    2. Implemented `apis/dns/v1beta1/dnsrecordset_reference.go` implementing `refs.Ref` with appropriate registration and `Normalize` fallback delegating to `getIdentityFromDNSRecordSetSpec`.
    3. Added comprehensive unit tests in `dnsrecordset_identity_test.go` asserting parsing behavior with both relative and absolute external GCP reference URLs.
    4. Added the template `//dns.googleapis.com/projects/{}/managedZones/{}/rrsets/{}` to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to prevent verification failure on CAI format location mismatch.
- **Impact**: DNSRecordSet now follows modern KCC conventions for identity and references representation, fully prepared for any direct controller migration.
