# ComputeSSLPolicy Identity and Reference Journal

## Observations & Learnings

- **Global and Regional Dual Identity Templates:** ComputeSSLPolicy supports both global and regional GCP URL formats (`global/sslPolicies/...` and `regions/{{region}}/sslPolicies/...`). We successfully modeled both formats with two distinct `gcpurls.Template` instances and dynamically matched the incoming format in `FromExternal`.
- **Global as Default Spec Extraction:** In KCC, the `ComputeSSLPolicy` Spec struct does not contain a location or region field. As a result, when resolving the identity from the spec/metadata, it is modeled as global by default, while external references still allow both global and regional formats to be fully resolved.
- **Normalize Integration:** We updated the `Normalize` function on `ComputeSSLPolicyRef` to use `refs.NormalizeWithFallback`. This provides backward compatibility by fetching and parsing the `status.selfLink` if present, and cleanly falls back to resolving the resource ID from the spec if not.
- **Compilation & Test Coverage:** We added unit tests to verify standard external references (both global/regional and full HTTP URLs/partial paths) and updated existing reference tests to expect standard normalized output.
