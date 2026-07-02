# IAM Service Journal

### [2026-06-29] IamAccessPolicy Types and gcpurls.Template Matching
- **Context**: Implementing Greenfield direct types and IdentityV2 for `IamAccessPolicy` in package `v1alpha1`.
- **Problem**:
  1. `generate-types` writes auxiliary types to `types.generated.go` on a per-package basis. Because `apis/iam/v1alpha1` merges resources from `google.iam.v2` (`IAMDenyPolicy`) and `google.identity.accesscontextmanager.v1` (`IamAccessPolicy`), running the generator for `IamAccessPolicy` overwrote `types.generated.go` and removed the nested type definitions (like `PolicyRule`) needed by `IAMDenyPolicy`, leading to compilation errors.
  2. The `gcpurls.Template` URL parser converted the `{access_policy}` segment to lowercase (`access_policy`) and tried to find it in the Go struct field `AccessPolicy` (which was converted to lowercase `accesspolicy`). Since they did not match, it caused a runtime panic during test initialization.
- **Solution**:
  1. In `apis/iam/v1alpha1/generate.sh`, we sequenced the `generate-types` runs so that `IamAccessPolicy:AccessPolicy` runs first, and `IAMDenyPolicy:Policy` runs last. Since `AccessPolicy` is a root message and does not reference any nested types that must go in `types.generated.go`, running `IAMDenyPolicy` last correctly restored all auxiliary types needed by the package.
  2. We changed the placeholder in the template from `{access_policy}` to `{accessPolicy}` to match the Go struct field `AccessPolicy` case-insensitively, resolving the panic.
- **Impact**: Enables smooth multi-service type merging in single KCC packages and prevents `gcpurls.Template` panics when implementing multi-word resource names.
