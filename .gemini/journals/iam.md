### [2026-06-15] Greenfield types generation for IamAccessPolicy under iam package
- **Context**: Implementing the direct KRM types, CRD, and IdentityV2 for the new greenfield resource `IamAccessPolicy` in `iam.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: 
  1. The canonical GCP `AccessPolicy` proto definition belongs to the `google.iam.v3beta` package which was not included in KCC's standard `dev/tools/controllerbuilder/generate-proto.sh` compilation script.
  2. Running a second `generate-types` command under the same group/version overwrites and comments out the existing generated Go types for other resources (e.g. `IAMDenyPolicy`).
  3. The `conditions` field uses an unsupported proto map with message values (`map<string, google.type.Expr>`), causing the generator to comment it out as unreachable.
- **Solution**:
  1. Copied `google/iam/v3beta/access_policy_resources.proto` from the latest googleapis repo to `mockgcp/apis/google/iam/v3beta/access_policy_resources.proto` as a hermetic overlay.
  2. Updated `generate-proto.sh` to compile our overlay protos under `mockgcp/apis/google/iam/v3beta/*.proto`.
  3. Combined the generation of both `IAMDenyPolicy` and `IamAccessPolicy` in a single `generate-types` call within `apis/iam/v1alpha1/generate.sh` using fully-qualified proto namespaces:
     ```bash
     go run . generate-types \
       --service google.iam.v2 \
       --api-version iam.cnrm.cloud.google.com/v1alpha1 \
       --resource IAMDenyPolicy:Policy \
       --resource IamAccessPolicy:google.iam.v3beta.AccessPolicy
     ```
  4. Expanded the scaffolded `IamAccessPolicySpec` to reference nested types like `AccessPolicyDetails` to keep them uncommented and reachable.
  5. Hand-coded `iamaccesspolicy_identity.go` utilizing three `gcpurls.Template` formats to handle hierarchical project, folder, and organization scopes seamlessly.
- **Impact**: Any direct controller implementation that shares a package or group with other resources must combine `generate-types` into a single invocation and use fully-qualified proto package paths for resources defined across different API levels.
