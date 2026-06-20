# ComputeURLMap Identity and Reference Journal

## Observations & Learnings

- **Dual Global and Regional Formats:** Similar to other Compute resources, `ComputeURLMap` supports both global and regional formats (`projects/{project}/global/urlMaps/{name}` and `projects/{project}/regions/{region}/urlMaps/{name}`). Using separate `gcpurls.Template` templates for global and regional URLs allowed us to gracefully parse both in `FromExternal`.
- **IdentityV2 Implementation:** We introduced a full `ComputeURLMapIdentity` struct that implements the standard `identity.IdentityV2` interface, making the identity resolving process safe, robust, and highly type-safe.
- **Reference Refactoring:** We updated the existing `urlmap_reference.go` to conform to KCC guidelines by using the standard `refs.Ref` interface and unified `NormalizeWithFallback` helper. This preserves backwards-compatibility for existing Terraform/DCL objects by falling back to `status.selfLink` if `status.externalRef` is not present.
- **Verification and Testing:** Tests for identity parsing/formatting and reference validation/normalization were implemented. All tests compiled and passed perfectly on the first try. No schema changes were made.
