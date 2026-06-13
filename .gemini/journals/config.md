### [2026-06-03] Proto Overlay for ConfigDeploymentGroup due to Pinned googleapis SHA
- **Context**: Implementing `ConfigDeploymentGroup` direct KRM types and CRD under `config.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The proto message `DeploymentGroup` was introduced in a newer `googleapis` commit (`f5cb7afc40b`) than the pinned SHA (`731d7f2ab6`) in `apis/git.versions`. As a result, the code generator (`generate-types`) failed because it could not find `google.cloud.config.v1.DeploymentGroup`.
- **Solution**: 
  1. Extracted the updated `config.proto` file from the `f5cb7afc40b63d52f43bc306cb9b64a87b681aea` commit of `googleapis`.
  2. Placed it as a custom overlay under `mockgcp/apis/google/cloud/config/v1/config.proto`.
  3. Modified `dev/tools/controllerbuilder/generate-proto.sh` to swap `-I` priorities (putting `${REPO_ROOT}/mockgcp/apis` before `${THIRD_PARTY}/googleapis/`) and explicitly list `google/cloud/config/v1/*.proto` to compile.
  4. Added a `rm` step in `generate-proto.sh` to delete the old `config.proto` from the `googleapis` clone to prevent `protoc` from failing due to input shadowing (since the wildcard matches both the clone and the mock overlay).
- **Impact**: Future agents working on config service resources can safely leverage custom overlays in `mockgcp/apis/` without waiting for a global repository-wide update of the pinned `googleapis` SHA.
