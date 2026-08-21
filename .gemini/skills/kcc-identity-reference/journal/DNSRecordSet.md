### [2026-06-10] DNSRecordSet Identity & Refs
- **Context**: Moving DNSRecordSet to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. DNSRecordSet was entirely missing identity or reference files.
    2. Per reviewer feedback, we dropped the `{type}` segment from the identity/reference template format to simplify it.
    3. Per further reviewer feedback, we added the primary CAIS-compliant format containing locations (e.g. `projects/{project}/locations/{location}/managedZones/{managedZone}/rrsets/{name}`) while retaining a fallback format without location (`projects/{project}/managedZones/{managedZone}/rrsets/{name}`) where `Location` is left empty/unset.
- **Solution**:
    1. Implemented `apis/dns/v1beta1/dnsrecordset_identity.go` utilizing two `gcpurls.Template` vars (primary and fallback).
    2. Implemented `apis/dns/v1beta1/dnsrecordset_reference.go` implementing `refs.Ref` with appropriate registration and `Normalize` fallback delegating to `getIdentityFromDNSRecordSetSpec`.
    3. Added comprehensive unit tests in `dnsrecordset_identity_test.go` asserting parsing behavior with both primary (with location) and fallback (without location) reference URLs.
    4. Kept the fallback template `//dns.googleapis.com/projects/{}/managedZones/{}/rrsets/{}` in `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to prevent verification failure since CAIS only knows the primary format.
- **Impact**: DNSRecordSet now follows modern KCC conventions for identity and references representation, fully prepared for any direct controller migration.
