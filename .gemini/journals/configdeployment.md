# ConfigDeployment Service Journal

### 2026-07-07 ConfigDeploymentGroup Types & Identity Implementation
- **Context**: Implementing KRM types and IdentityV2 for `ConfigDeploymentGroup` under `apis/configdeployment/v1alpha1/` (Issue #9020).
- **Problem**: The googleapis commit `731d7f2ab6` specified/templated in the issue description did not contain the `DeploymentGroup` proto definition (it was added later).
- **Solution**: We updated `apis/git.versions` to use the latest googleapis commit `1a91fbee8fd1c7f9bdc3d585ff8032b1619a2051` which contains `DeploymentGroup` under `google.cloud.config.v1`. This enabled `generate.sh` to successfully pull the protos and compile/generate the corresponding types, CRDs, and deepcopy methods.
- **Impact**: Any future `configdeployment` service implementations (such as `ConfigDeployment` or `ConfigRevision`) must use a googleapis commit at or after `1a91fbee8fd1c7f9bdc3d585ff8032b1619a2051` to ensure the proto definitions are present in the build cache.
