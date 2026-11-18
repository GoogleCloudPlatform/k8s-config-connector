
### [2026-05-26] Direct Types for CloudBuildConnection: Reference Overrides and Skipped Types
- **Context**: Scaffolding direct types for `CloudBuildConnection:Connection` (v1alpha1).
- **Problem**: 
  1. The proto definition for `Connection` includes multiple nested configuration structures (e.g., `BitbucketCloudConfig`, `GitHubEnterpriseConfig`) containing `webhook_secret_secret_version` and similar fields that map to `SecretManagerSecretVersion`. These require `refsv1beta1.SecretManagerSecretVersionRef`.
  2. The proto definition for `InstallationState` only contains output fields, meaning the generator only creates `InstallationStateObservedState` and skips `InstallationState` entirely.
- **Solution**: 
  1. Copied the nested configuration structs from `types.generated.go` into `connection_types.go` and explicitly overridden the fields to use `*refsv1beta1.SecretManagerSecretVersionRef` (and `json:"...Ref,omitempty"`). Reran `make generate` to prune the string-typed structs from `types.generated.go`.
  2. Used `*InstallationStateObservedState` directly in `CloudBuildConnectionObservedState` within `connection_types.go`.
  3. Added the skipped `*ObservedState` corresponding structs to `connection_types.go` as overriding a base type forces the prune-types script to skip its `ObservedState` generation as well.
- **Impact**: Future agents implementing resources with nested references should know they need to manually copy the structs to `<kind>_types.go`, override the reference fields with `ref` structs, AND also provide the `ObservedState` structs for those overridden types to ensure everything compiles correctly.

### [2026-11-17] Prepopulate Direct Types, Identity, and Judgement Queue for CloudBuildBuild
- **Context**: Greenfield Step 1a prepopulation for `CloudBuildBuild:Build` (`cloudbuild.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Decisions**:
  1. `Build` in `google.devtools.cloudbuild.v1` defines `CreateBuild` without `UpdateBuild` (one-shot execution), requiring `--allow-non-declarative` in `controllerbuilder generate-types`.
  2. `google.devtools.cloudbuild.v1.Build` declares two resource patterns: `projects/{project}/builds/{build}` and `projects/{project}/locations/{location}/builds/{build}`. Following CAI (`cloudassetinventory_names.jsonl`) and modern regional URL conventions, explicitly locked in `--pattern CloudBuildBuild=projects/{project}/locations/{location}/builds/{build}` in `generate-identity`.
  3. Added `CloudBuildBuild` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Ran `apis/cloudbuild/generate.sh` and confirmed pure mechanical output and idempotency. Recorded expected untriaged missing fields and non-representable refs in `tests/apichecks` exceptions via golden update.

### [2026-11-17] Triage Judgement Queue and Finalize Direct Types for CloudBuildBuild
- **Context**: Greenfield Step 1b judgement queue triage and type finalization for `CloudBuildBuild:Build` (`cloudbuild.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Decisions**:
  1. **Server-Generated ID & Parent Ref**: Cloud Build assigns build IDs on creation (`Build.id` and `Build.name` are output-only, and `CreateBuildRequest` has no ID parameter). Added server-generated ID comment to `spec.resourceID` so `generate-identity` emits `_ identity.ServerGeneratedIdentity = &CloudBuildBuildIdentity{}`. Replaced `projectRef` with inlined `*parent.ProjectAndLocationRef`.
  2. **Nested Reference Overrides**: Converted `logsBucket` to `*storagev1beta1.StorageBucketRef`, `serviceAccount` to `*refsv1beta1.IAMServiceAccountRef`, `options.pool.name` to `*cloudbuildv1beta1.CloudBuildWorkerPoolRef`, `secrets[].kmsKeyName` and `availableSecrets.inline[].kmsKeyName` to `*kmsv1beta1.KMSCryptoKeyRef`, `availableSecrets.secretManager[].versionName` and `gitConfig.http.proxySecretVersionName` to `*secretmanagerv1beta1.SecretManagerSecretVersionRef`, and `artifacts.goModules[].repositoryProjectID` to `*refsv1beta1.ProjectRef`.
  3. **Forked Spec Structs to Prevent Status Ref Leak**: `Source` (`StorageSource`, `RepoSource`, `StorageSourceManifest`) is shared between `Spec.Source` and `Status.ObservedState.SourceProvenance`. Forked the spec structs as `Build_Source`, `Build_StorageSource`, `Build_RepoSource`, and `Build_StorageSourceManifest` with Ref types in `cloudbuildbuild_types.go`, leaving unreferenced scalar versions in `types.generated.go` for status so `TestNoRefsInStatus` passes.
  4. **Oneof Reference Support in Mapper Generator**: `dependencies[].gitSource.repository.developerConnect` is inside a proto oneof (`repotype`). When mapped to `DeveloperConnectRef`, `mappergenerator.go` previously generated field assignment directly on the struct instead of instantiating the oneof wrapper struct `&pb.Dependency_GitSourceRepository_DeveloperConnect{...}`. Updated `mappergenerator.go` to wrap oneof references.

### [2026-11-18] Direct Controller, E2E Fixtures, and Fuzzer for CloudBuildBuild
- **Context**: Greenfield Step 2 direct controller, fuzzer, and E2E fixtures recorded against real GCP for `CloudBuildBuild:Build` (`cloudbuild.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Decisions**:
  1. **gRPC Protocol and Regional Endpoints**: Cloud Build `Build` operations use regional endpoints (`${location}-cloudbuild.googleapis.com:443`). The GAPIC REST client has path generation limitations with `CreateBuildRequest` for regional endpoints, so `gcp.NewClient` with `m.config.GRPCClientOptions()` is used.
  2. **User-Agent Sanitization**: The Cloud Build API gateway crashes with HTTP 500 / gRPC Internal when the `User-Agent` contains duplicate `kcc/` product prefixes (specifically `kcc/... kcc/controller-manager/...`). Sanitizing `config.UserAgent` by replacing `kcc/controller-manager` with `controller-manager` resolves the server-side error while preserving attribution.
  3. **Non-Deletable GCP Resource**: Cloud Build builds cannot be deleted via the GCP API (no `DeleteBuild` RPC). The controller safely removes the finalizer without attempting a delete, and `CloudBuildBuild` is marked as non-deletable in `pkg/test/httplogreport/fixture/loader.go`.
  4. **BuildStepOutputs Byte Formatting**: When steps produce no builder output, Cloud Build returns empty byte slices `[]byte{}` in `build_step_outputs`. Because `[][]byte` serializes in JSON as `[""]`, kube-apiserver OpenAPI schema validation rejects empty strings for `format: byte`. Added `cleanBuildStatus` in the controller to drop empty byte slices from `status.ObservedState.Results.BuildStepOutputs`.
  5. **Operation ID Extraction in E2E Harness**: Cloud Build long-running operations use the format `operations/build/{project}/{operationId}`. Previously, the test harness treated `build` as the operation ID, replacing the word `build` with `${operationID}` throughout proto @types and endpoints. Updated `replacements.go` and `httplog.go` to ignore `build` as an ID and extract the true operation ID token.

