# IdentityPlatformConfig Direct Type Implementation Journal

## Observations & Learnings

1. **Protocol Buffer Source and OpenAPI/Discovery Choice**:
   - For IdentityPlatformConfig, there is no public proto file in Google's official `googleapis` repository on GitHub.
   - We did not compile or import from DCL protobuf definitions. Instead, we generated the KRM Go types entirely from the published Discovery API JSON specification (`identitytoolkit-api.json`) using the `openapi-to-krm` tool.

2. **Schema Compatibility and Custom Reference Types**:
   - To achieve 100% schema compatibility with the baseline CRD, we copied the original client-go API types to `identityplatformconfig_types.go` (preserving lowercase proto message file name convention).
   - Hand-coded custom local reference structures (`ConfigFunctionUriRef`, `ConfigDefaultTenantLocationRef`) were implemented to match the original properties precisely and preserve exact OpenAPI validations (e.g. `oneOf` rules).
   - Standard `refs.ProjectRef` from `apis/refs/v1beta1` was used to represent the parent `projectRef` correctly.

## Verification

- `dev/tasks/diff-crds` output was down to only a single intended line: `+ spec.projectRef.kind=string`, proving 100% backward schema compatibility.
- Executed `make lint-custom manifests generate-go-client ensure fmt` with zero failures/compilation errors.
- Executed `go vet ./...` with zero issues.
- Ran `dev/tasks/generate-resource-report` to successfully update the CRD resource reports.
