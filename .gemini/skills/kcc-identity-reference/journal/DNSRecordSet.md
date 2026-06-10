### [2026-06-10] DNSRecordSet Identity & Refs
- **Context**: Moving DNSRecordSet to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. DNSRecordSet was entirely missing identity or reference files.
    2. The GCP URL template for `ResourceRecordSets` (`projects/{project}/managedZones/{managedZone}/rrsets/{name}/{type}`) has a custom shape at the end (`{name}/{type}`) that does not directly match the standard Cloud Asset Inventory (CAI) name format where the last segment is a single variable (`{{RESOURCE_RECORD_SET}}`), requiring an exclusion in registry tests.
- **Solution**:
    1. Implemented `apis/dns/v1beta1/dnsrecordset_identity.go` utilizing `gcpurls.Template` with the correct segments mapping.
    2. Implemented `apis/dns/v1beta1/dnsrecordset_reference.go` implementing `refs.Ref` with appropriate registration and `Normalize` fallback delegating to `getIdentityFromDNSRecordSetSpec`.
    3. Added comprehensive unit tests in `dnsrecordset_identity_test.go` asserting parsing behavior with both relative and absolute external GCP reference URLs.
    4. Added the template to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to prevent verification failure on CAI format shape mismatch.
- **Impact**: DNSRecordSet now follows modern KCC conventions for identity and references representation, fully prepared for any direct controller migration.
