### [2026-06-28] OSConfigGuestPolicy Identity & Refs
- **Context**: Moving OSConfigGuestPolicy to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. OSConfigGuestPolicy was using legacy reconciler configuration, but didn't have canonical `IdentityV2` or `refs.Ref` implementations.
    2. OSConfigGuestPolicy is a project-scoped resource (no location segment in its URL: `projects/{project}/guestPolicies/{osConfigGuestPolicy}`).
    3. OSConfigGuestPolicy is missing from Cloud Asset Inventory (CAIS) mapping, which causes template registry tests to fail unless explicitly ignored.
- **Solution**:
    1. Implemented `apis/osconfig/v1beta1/osconfigguestpolicy_identity.go` utilizing standard `gcpurls.Template`.
    2. Implemented `apis/osconfig/v1beta1/osconfigguestpolicy_reference.go` utilizing standard `refs.Ref` and `common.ToStructuredType` inside the fallback of its `Normalize` implementation.
    3. Added `apis/osconfig/v1beta1/osconfigguestpolicy_identity_test.go` to test external URI parsing.
    4. Ignored the registered template `//osconfig.googleapis.com/projects/{}/guestPolicies/{}` in `pkg/gcpurls/registry_test.go` as it is not present in CAIS definitions.
- **Impact**: OSConfigGuestPolicy is fully aligned with modern direct controller identity and reference conventions.
