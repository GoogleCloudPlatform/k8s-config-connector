### [2026-05-28] Vendoring missing protos for SecurityProfileGroup
- **Context**: Implementing types for `NetworkSecuritySecurityProfileGroup` (Issue #8738).
- **Problem**: The proto messages for `SecurityProfileGroup` and its helper modules (`security_profile_group_intercept.proto`, `security_profile_group_mirroring.proto`, `security_profile_group_threatprevention.proto`, `security_profile_group_urlfiltering.proto`, `security_profile_group_service.proto`) were missing from the pinned `googleapis` SHA (`731d7f2ab6`).
- **Solution**: 
  1. Extracted all `security_profile_group*.proto` files from the `googleapis` master branch and placed them in `mockgcp/apis/google/cloud/networksecurity/v1/`.
  2. Regenerated `.build/googleapis.pb` with `generate-proto.sh` and ran `generate-types` for `NetworkSecuritySecurityProfileGroup`.
  3. Added `NetworkSecuritySecurityProfileRef` under `apis/refs/v1beta1/networksecurityrefs.go` to handle reference fields.
- **Impact**: Enables smooth generation of types and CRDs for resources added to Google APIs after the pinned SHA `731d7f2ab6`.

### [2026-05-27] Missing proto in pinned googleapis SHA
- **Context**: Implementing types for `NetworkSecurityInterceptDeployment` (Issue #8726).
- **Problem**: The proto message `InterceptDeployment` does not exist in the pinned `googleapis` SHA (`731d7f2ab6`). The generator (`generate-types`) failed with `proto: not found`.
- **Solution**: Downloaded the missing `intercept.proto` from the `googleapis` master branch into `mockgcp/apis/google/cloud/networksecurity/v1/intercept.proto` and modified `dev/tools/controllerbuilder/generate-proto.sh` to include `mockgcp/apis/google/cloud/networksecurity/*/*.proto` during `.build/googleapis.pb` generation. Manually added fields to the spec/status structs as `generate-types` skips fields if it cannot map them perfectly during initial scaffolding.
- **Impact**: Future agents working on `NetworkSecurityInterceptDeployment` or `NetworkSecurityInterceptDeploymentGroup` should know that their protos are sourced from `mockgcp/apis` instead of the pinned `googleapis` SHA. If updating `googleapis` in the future, we can safely remove these from `mockgcp/apis`.

### [2026-05-27] Vendoring missing protos and adding custom reference struct
- **Context**: Implementing `NetworkSecurityInterceptEndpointGroup` (Issue #8728).
- **Problem**: The proto `intercept.proto` for `InterceptEndpointGroup` was introduced in a newer commit and was missing from the pinned `googleapis` SHA (`731d7f2ab6`). The `generate-types` tool failed because it could not find the message. Also, `intercept_deployment_group` required a `NetworkSecurityInterceptDeploymentGroupRef` which didn't exist in KCC.
- **Solution**: 
  1. Fetched `intercept.proto` directly from the `googleapis` master branch and vendored it into `mockgcp/apis/google/cloud/networksecurity/v1/intercept.proto`.
  2. Updated `dev/tools/controllerbuilder/generate-proto.sh` to include `${REPO_ROOT}/mockgcp/apis/google/cloud/networksecurity/*/*.proto`.
  3. Created `NetworkSecurityInterceptDeploymentGroupRef` in `apis/refs/v1beta1/networksecurity_refs.go`.
- **Impact**: Future agents working on Greenfield resources where protos are missing from the pinned SHA should vendor the proto files into `mockgcp/apis/google/cloud/...` and update `generate-proto.sh` instead of attempting to bump the `apis/git.versions` SHA.

### [2026-06-03] Protobuf SHA Override in generate.sh for NetworkSecurity v1alpha1
- **Context**: Implementing types for `NetworkSecuritySACRealm` (Issue #8736).
- **Problem**: The main pinned `googleapis` SHA (`731d7f2ab6`) does not contain `BackendAuthenticationConfig` or `SACRealm`. When attempting to regenerate the protobuf with `731d7f2ab6`, the generation fails because `BackendAuthenticationConfig` is not found.
- **Solution**: Reverted the generator's `PROTO_SHA` in `apis/networksecurity/v1alpha1/generate.sh` to its custom override commit `cdc919ff596e263f2cc55a9780d2f74633da1ced`. This commit contains the protobufs for both `BackendAuthenticationConfig` and `SACRealm`, allowing the code generation to succeed.
- **Impact**: When adding new resources in `networksecurity`, keep in mind that `v1alpha1/generate.sh` uses a specific protobuf SHA override `cdc919ff596e263f2cc55a9780d2f74633da1ced` to fetch newer APIs that are not yet available in the repo-wide default SHA.

### [2026-06-03] Creating target MirroringEndpointGroup and InterceptEndpointGroup references
- **Context**: Implementing types and Identity for `NetworkSecuritySecurityProfile` (Issue #8737).
- **Problem**: `NetworkSecuritySecurityProfile` contains two references to endpoint groups (`mirroringEndpointGroup` and `interceptEndpointGroup`), which did not have pre-existing reference structs in `apis/refs/v1beta1/networksecurityrefs.go`.
- **Solution**: Defined `NetworkSecurityMirroringEndpointGroupRef` and `NetworkSecurityInterceptEndpointGroupRef` in `apis/refs/v1beta1/networksecurityrefs.go` to provide structured validation for endpoint group references, and used them in `CustomMirroringProfile` and `CustomInterceptProfile` respectively.
- **Impact**: Enables strict validation and clean reference resolution for endpoint group fields within a SecurityProfile definition.
