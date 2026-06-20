### [2026-06-20] ComputeURLMap Identity & Refs
- **Context**: Moving ComputeURLMap to `identity.IdentityV2` and `refs.Ref` patterns.
- **Problem**: 
    1. ComputeURLMap did not have an identity file implementing `identity.IdentityV2` backed by `gcpurls.Template`.
    2. `urlmap_reference.go` contained a manual lookup implementation of reference normalization and validation.
- **Solution**:
    1. Implemented `apis/compute/v1beta1/computeurlmap_identity.go` to use `identity.IdentityV2` backed by two distinct `gcpurls.Template` instances supporting both global and regional GCP URL paths (under `global/urlMaps/{urlmap}` and `regions/{region}/urlMaps/{urlmap}`).
    2. Completely replaced `urlmap_reference.go` with `computeurlmap_reference.go` supporting standard `Normalize` fallback flow and `ParseExternalToIdentity()`.
    3. Added `apis/compute/v1beta1/computeurlmap_identity_test.go` to verify the standard template-based URL parsing, ValidateExternal, and Normalize logic.
- **Impact**: ComputeURLMap fully matches the modern direct controller identity and references conventions.
