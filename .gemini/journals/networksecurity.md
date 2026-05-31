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
### [2026-05-27] Missing UrlList Proto in Pinned googleapis SHA
- **Context**: Implementing Greenfield direct types for `NetworkSecurityUrlList` (issue #8740).
- **Problem**: The proto message `google.cloud.networksecurity.v1.UrlList` was missing in the pinned `googleapis` SHA `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb`.
- **Solution**: Downloaded the `url_list.proto` from the `googleapis` master branch, placed it under `mockgcp/apis/google/cloud/networksecurity/v1/url_list.proto`, and modified `dev/tools/controllerbuilder/generate-proto.sh` to include `mockgcp/apis/google/cloud/networksecurity/*/*.proto`. This allowed the proto to be successfully compiled into `.build/googleapis.pb` and read by `generate-types`.
- **Impact**: When scaffolding new direct resources, if the proto does not exist in the pinned `googleapis` version, you can inject it via `mockgcp/apis/` and update `generate-proto.sh` to include that path instead of relying solely on `googleapis` or overlays.

### [2026-05-27] Scaffolded Location Field as Pointer
- **Context**: Reviewing auto-generated KRM types for `NetworkSecurityUrlList` according to Greenfield Review rules.
- **Problem**: The auto-generated type scaffolded `Location` as a `string` (scalar primitive), which violates the AI reviewer rule that all Go scalar primitive types must be pointers.
- **Solution**: Changed `Location string` to `Location *string` in `NetworkSecurityUrlListSpec`. The identity parser (`refs.GetLocation(obj)`) supports both `string` and `*string` via reflection.
- **Impact**: All scalar primitives, even required ones like `Location`, must be pointers in the KRM spec to pass the Greenfield review scripts.
