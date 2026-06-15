### [2026-06-11] DNSResponsePolicy Identity & Refs
- **Context**: Moving DNSResponsePolicy to the new identity and reference patterns using multiple `gcpurls.Template` formats.
- **Problem**:
    1. DNSResponsePolicy was previously implemented in `apis/dns/v1alpha1` but only supported a single `locations/{location}/responsePolicies/{responsePolicy}` format.
    2. This led to failures when standard GCP reference URLs (which omit the location segment for global response policies) were passed as external references, as they were treated as invalid by unit tests.
- **Solution**:
    1. Refactored `apis/dns/v1alpha1/dnsresponsepolicy_identity.go` to utilize two `gcpurls.Template` vars: a primary location-based format (`projects/{project}/locations/{location}/responsePolicies/{responsePolicy}`) and a fallback location-less format (`projects/{project}/responsePolicies/{responsePolicy}`).
    2. Updated `apis/dns/v1alpha1/dnsresponsepolicy_identity_test.go` unit tests to verify that standard location-less GCP response policies are parsed correctly into the fallback template representation.
    3. Registered the fallback template exception `//dns.googleapis.com/projects/{}/responsePolicies/{}` in `pkg/gcpurls/registry_test.go` so the CAI alignment validation tests continue to pass successfully.
- **Impact**: DNSResponsePolicy is now fully compliant with modern KCC identity and reference patterns, supporting both CAIS-compliant formats and direct location-less GCP API references.
