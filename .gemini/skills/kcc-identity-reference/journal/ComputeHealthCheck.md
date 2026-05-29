When implementing ComputeHealthCheck, I found:
1. `ComputeHealthCheck` can be either regional or global, requiring two separate `gcpurls.Template` formats: `GlobalComputeHealthCheckIdentityFormat` and `RegionalComputeHealthCheckIdentityFormat`.
2. Both templates can share the single `ComputeHealthCheckIdentity` struct because `gcpurls.Template` only panics if a template placeholder is missing from the struct (having extra fields in the struct is perfectly fine).
3. In `FromExternal`, we first attempt to parse with the regional template, and fall back to the global template, setting `Region = "global"` if the global template matches.
4. Unit tests (`computehealthcheck_identity_test.go` and `computehealthcheck_reference_test.go`) were added to thoroughly test both regional and global configurations.
