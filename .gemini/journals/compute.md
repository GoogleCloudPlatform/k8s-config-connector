### 2026-06-02 Implementing ComputeSecurityPolicy Direct Controller
- **Context**: Implemented ComputeSecurityPolicy direct controller and E2E tests (minimal and maximal) under Issue #8981.
- **Problem**: Moving ComputeSecurityPolicy manual mapping functions to the computesecuritypolicy subpackage caused compilation failures in the parent `compute` package's `mapper.generated.go`. Also, the API spec lists the rule field as `rule` (JSON key `rule,omitempty`) instead of `rules`.
- **Solution**: Created a thin delegation bridge in the parent package `compute` to forward `SecurityPolicy` mappings to the new computesecuritypolicy subpackage, and updated the maximal E2E fixtures to use the correct `rule` field and `redirectSiteKeyRef` representation.
- **Impact**: Clean isolation of ComputeSecurityPolicy controller while maintaining complete backward compatibility with the parent package's generated mappers and API types.
