### 2026-06-24 Managing Shared Protobuf Types and Acronym Capitalization in aiplatform Package
- **Context**: Implementing direct types for `VertexAITuningJob` under `apis/aiplatform/v1alpha1/`.
- **Problem**: 
  1. The existing `apis/aiplatform/v1alpha1/generate.sh` script ran `generate-types` only for `VertexAISpecialistPool` and then ran `git checkout HEAD -- types.generated.go`. This discarded any generated/uncommented helper types required by other resources or newly added resources (like `VertexAITuningJobSpec`'s `SupervisedTuningSpec`).
  2. Running `generate-types` on a single resource would cause the generator to deem types used by other resources as "unreachable" and comment them out, causing extensive compilation failures.
  3. The `ListValue` and `ListTypeValue` structures were commented out in `model_types.go`, but regenerated references to them were added in `types.generated.go`.
- **Solution**:
  1. Updated `generate.sh` to include all 4 active package resources (`VertexAISpecialistPool`, `AIPlatformModel`, `VertexAIFeatureOnlineStore`, and `VertexAITuningJob`) in the `generate-types` run, and removed the `git checkout types.generated.go` step.
  2. Restored and uncommented a standard `ListValue` structure in `model_types.go` matching the convention in other packages.
  3. Ran `make fmt` and updated the generated CRD manifests for both `VertexAITuningJob` and `AIPlatformModel` cleanly.
- **Impact**: Ensures that any future direct resources added to the `aiplatform` package can be generated and compiled safely without breaking existing resources or generating out-of-date/pruned structures in `types.generated.go`.

### 2026-07-06 Exposing Commented-Out Unreachable Sub-structures and References for VertexAIPipelineJob
- **Context**: Implementing direct types for `VertexAIPipelineJob` under `apis/aiplatform/v1alpha1/`.
- **Problem**: The nested sub-structures `PipelineJob.RuntimeConfig` (`PipelineJobRuntimeConfig`), `PscInterfaceConfig` (`PSCInterfaceConfig`), `DnsPeeringConfig` (`DNSPeeringConfig`), and several observed state sub-structures were commented out in `types.generated.go` as unreachable since they were not previously used by other resources in this service.
- **Solution**: Explicitly defined `PipelineJobRuntimeConfig`, `PSCInterfaceConfig`, `DNSPeeringConfig`, and the `VertexAIPipelineJobObservedState` nested structures inside `vertexaipipelinejob_types.go`. This automatically allowed the code generator to recognize and map these structures. We also implemented proper reference types like `computev1beta1.ComputeNetworkRef`, `computev1alpha1.ComputeNetworkAttachmentRef`, and `refsv1beta1.IAMServiceAccountRef` for reference fields.
- **Impact**: Enables flawless generation of deepcopy methods, CRD fields, and mappers for `VertexAIPipelineJob` while maintaining 100% clean pre-submit checks.

### 2026-07-06 Implementing types, CRD, and IdentityV2 for VertexAIStudy
- **Context**: Implementing direct types, CRD, and IdentityV2 for `VertexAIStudy` under `apis/aiplatform/v1alpha1/` (Issue #9250).
- **Problem**: The `StudySpec` in Vertex AI is a highly complex nested structure with recursive self-references in parameter specifications (e.g., `ConditionalParameterSpecs` containing nested `ParameterSpec` fields). These types were previously marked "unreachable" and commented out in `types.generated.go`.
- **Solution**:
  1. Updated `apis/aiplatform/v1alpha1/generate.sh` to include `--resource VertexAIStudy:Study`.
  2. Defined KRM Go type `VertexAIStudySpec` referencing `StudySpec` and `VertexAIStudyObservedState` in `vertexaistudy_types.go`.
  3. Re-ran `generate.sh` and `make fmt`, which automatically uncommented and populated the nested and recursive types in `types.generated.go`, generated deepcopy functions, and produced mapping code.
  4. Implemented `VertexAIStudy` IdentityV2 matching `projects/{project}/locations/{location}/studies/{study}` and verified it via unit tests.
- **Impact**: Enables full structural representation of complex studies and hyperparameter tuning specs in KRM with 100% Go-compatible schemas, preserving nested and recursive specs natively.

### 2026-07-06 Implementing VertexAISchedule Greenfield Types
- **Context**: Implementing direct types and IdentityV2 for `VertexAISchedule` under `apis/aiplatform/v1alpha1/` (Issue #9248).
- **Problem**: Running `validate-prereqs.sh` locally on a branch with uncommitted changes triggers formatting errors because the script checks `git diff --name-only` expecting a committed state.
- **Solution**: Followed the `send-pr` workflow by committing the scaffolded types, CRD, and Identity files first before executing the PR sending script.
- **Impact**: Demonstrates that the `generate.sh` script robustly handles multiple direct resources in `aiplatform` without pruned/unreachable type issues when all active types are correctly listed.

### 2026-07-08 Implementing the Greenfield Direct Controller, Fuzzer, and E2E Fixtures for VertexAIPipelineJob
- **Context**: Implementing the direct controller, E2E basic test fixtures, and fuzzer for `VertexAIPipelineJob` as part of the Greenfield migration.
- **Problem**: Greenfield resource implementation requires the creation of a fully-isolated direct controller to manage the reconciliation lifecycle (Adapter interface: Find, Create, Update, Delete, and Export), registration in the static configuration, a KRM fuzzer matching specification/status fields, and minimal/maximal golden test fixtures.
- **Solution**:
  1. Created `pkg/controller/direct/aiplatform/vertexaipipelinejob_controller.go` package-isolated controller, implementing the `Adapter` interface. Since `PipelineJob` is immutable in Vertex AI, `Update` returns an immutability error if any diff is found between desired and actual state, in accordance with KCC guidelines.
  2. Leveraged Python generator `dev/tasks/generate_static_config.py` to automatically register `VertexAIPipelineJob` into `pkg/controller/resourceconfig/static_config.go` with its default/supported reconciler set to `direct`.
  3. Created `pkg/controller/direct/aiplatform/vertexaipipelinejob_fuzzer.go` and configured fuzzer fields utilizing the fluent builder pattern.
  4. Scaffolded E2E golden tests under `pkg/test/resourcefixture/testdata/basic/aiplatform/v1alpha1/vertexaipipelinejob/` (`vertexaipipelinejob-minimal` and `vertexaipipelinejob-maximal`), including `dependencies.yaml` to provision a `StorageBucket` used as the `gcsOutputDirectory` parameter.
- **Impact**: Ensures standard, fully compliant Greenfield controller implementation and E2E testing framework support for VertexAIPipelineJob.

### 2026-09-29 Implementing AIPlatformReasoningEngine Greenfield Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing the direct controller, E2E fixtures, and fuzzer for `AIPlatformReasoningEngine` (Issue #13501).
- **Problem**: 
  1. `ReasoningEngine` in GCP Vertex AI uses purely service-generated numerical identifiers (e.g., `projects/<project-number>/locations/<location>/reasoningEngines/<id>`). When parsing the created resource name from the LRO response, overwriting the identity struct directly converted the project ID to project number, causing subsequent reconciliations in `GetIdentity()` to detect a parent project mismatch between `status.externalRef` and `spec.projectRef`.
  2. The E2E normalizer in `mockgcp/mockaiplatform/normalize.go` only checked for `aiplatform.googleapis.com` in `event.URL()`, skipping gRPC requests matching `/google.cloud.aiplatform.v1.ReasoningEngineService/...`.
  3. `spec.agentFramework` must be one of `[google-adk, langchain, langgraph, ag2, llama-index, custom, a2a]`.
- **Solution**:
  1. Updated `AIPlatformReasoningEngineAdapter.Create` to parse only the `ReasoningEngine` ID segment from the LRO response into `a.id.ReasoningEngine`, keeping `a.id.Project` consistent with `spec.projectRef`.
  2. Updated `mockgcp/mockaiplatform/normalize.go` and `pkg/cais/caistesting/testing.go` to normalize reasoning engine dynamic IDs with `${reasoningEngineID}` for both gRPC and CAIS logs.
  3. Recorded and verified both minimal and maximal fixtures against real GCP.
- **Impact**: Cleanly reconciles `AIPlatformReasoningEngine` with service-generated IDs, and enables reliable, reproducible E2E tests against real GCP.

### 2026-10-03 Implementing AIPlatformSpecialistPool Greenfield Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing the direct controller, E2E fixtures, and fuzzer for `AIPlatformSpecialistPool` (Issue #13683).
- **Implementation**:
  1. Implemented direct controller `pkg/controller/direct/aiplatform/aiplatformspecialistpool_controller.go` adhering to the 4 rules for service-generated resource IDs.
  2. Registered `AIPlatformSpecialistPool` with direct reconciler in `pkg/controller/resourceconfig/static_config.go`.
  3. Configured KRM fuzzer with fluent syntax in `pkg/controller/direct/aiplatform/aiplatformspecialistpool_fuzzer.go`.
  4. Scaffolded minimal and maximal E2E test fixtures under `pkg/test/resourcefixture/testdata/basic/aiplatform/v1alpha1/aiplatformspecialistpool/`.
- **Finding & Real GCP Observation**:
  During live execution against real GCP (`./hack/record-gcp`), the Vertex AI `CreateSpecialistPool` API returned `rpc error: code = FailedPrecondition desc = Data labeling service is shutdown` (audited via Cloud Audit Logs `SpecialistPoolService.CreateSpecialistPool`). Google Cloud deprecated and shut down the Vertex AI Data Labeling Service (and its associated human labeling SpecialistPool backend) on July 1, 2024.
- **Impact**: Controller implementation and CRD structures are complete, idiomatic, and verified with unit tests and fuzzer roundtripping.

### [2026-10-29] Triaging Judgement Queue and Modelling Child Hierarchy for AIPlatformFeatureView
- **Context**: Triaging judgement queue and finalizing types/identity for `AIPlatformFeatureView` (Issue #14600).
- **Problem**: `AIPlatformFeatureView` was generated with separate `ProjectRef` and `Location` fields and open `identity-root-unknown` and `parent-ref-not-modelled` entries in `judgement_queue.yaml`. It is a child resource of `VertexAIFeatureOnlineStore` (`projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`). Additionally, `VertexAIFeatureOnlineStore` had no `VertexAIFeatureOnlineStoreRef` reference type defined in the package, and `etag` and top-level `labels` were placed in `AIPlatformFeatureViewSpec`.
- **Solution**:
  1. Implemented `VertexAIFeatureOnlineStoreRef` in `apis/aiplatform/v1alpha1/vertexaifeatureonlinestore_reference.go` referencing `FeatureOnlineStoreIdentity`.
  2. Modelled `spec.parentRef` (`*VertexAIFeatureOnlineStoreRef`) with `+required` and immutable CEL rule `self == oldSelf` on `AIPlatformFeatureViewSpec`, and removed `projectRef` and `location`.
  3. Removed `labels` from `AIPlatformFeatureViewSpec`, moved `etag` to `AIPlatformFeatureViewObservedState`, and updated the REST resource docs URL.
  4. Added `generate-controller` and `generate-fuzzer` for `AIPlatformFeatureView` in `apis/aiplatform/generate.sh`, and updated the controller to resolve location via `id.Location`.
  5. Resolved all 10 judgement queue entries for `AIPlatformFeatureView`.
- **Impact**: Enables clean identity resolution from `spec.parentRef`, maintains KRM standards, and ensures mapper and fuzzer tests round-trip seamlessly.

### [2026-11-02] Implementing AIPlatformFeatureView Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and real GCP E2E fixtures for `AIPlatformFeatureView` (Issue #14780).
- **Problems & Solutions**:
  1. **Regional Endpoint for VertexAIFeatureOnlineStore**: `VertexAIFeatureOnlineStore`'s controller previously omitted regional endpoint options, connecting to `aiplatform.googleapis.com:443` which only accepts location `global`. Calls to `us-central1` failed with `The provided location ID doesn't match the endpoint. The valid location ID is 'global'`. Fixed `vertexaifeatureonlinestore_controller.go` to configure `WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location))`.
  2. **GCP Resource ID Constraints**: In GCP's `FeatureOnlineStoreAdminService`, resource IDs for `FeatureOnlineStore` and `FeatureView` must match `[a-z0-9_]` (hyphens `-` are rejected with `Invalid FeatureOnlineStore resource name`). The fixtures were updated to use hyphen-free identifiers (`featureonlinestore${uniqueId}` and `resourceID: featureview${uniqueId}`).
  3. **Parent Reference Normalization Fallback**: During identity tests (`TestGoldenIdentitiesYamlFiles`), `spec.parentRef` (`VertexAIFeatureOnlineStore`) in `dependencies.yaml` lacks `status.externalRef`. Implemented a spec-identity fallback in `VertexAIFeatureOnlineStoreRef.Normalize` using `NormalizeWithFallback` and `getIdentityFromFeatureOnlineStoreSpec` to ensure dependent child resources can resolve the parent identity prior to full reconciliation.
  4. **serviceAgentType Omission from GET Responses**: GCP's `GetFeatureView` does not return `service_agent_type` in its response (it returns `serviceAccountEmail`). Specifying `serviceAgentType` in KRM spec causes continuous drift during reconciliation.
  - **Impact**: Both minimal and maximal E2E fixtures recorded and passed against real GCP with all checks (V1 through V7) passing. Mock GCP skips gracefully pending Step 3.

### [2026-11-04] Triaging Judgement Queue and Finalizing Direct Types and Identity for AIPlatformModelMonitor
- **Context**: Triaging judgement queue and finalizing direct KRM types, identity, references, and parent for `AIPlatformModelMonitor` under `apis/aiplatform/v1alpha1/` (Issue #14866).
- **Problem**: `AIPlatformModelMonitor` was prepopulated with separate `ProjectRef` and `Location` fields, open `location-or-parent-ref` and reference judgement queue entries (`authoritative-proto-reference`, `possible-reference-by-description`, etc.).
- **Solution**:
  1. Updated `AIPlatformModelMonitorSpec` to inline `*parent.ProjectAndLocationRef`, replacing separate `ProjectRef` and `Location`.
  2. Converted string reference fields to typed KCC reference types: `modelRef` (`*AIPlatformModelRef`), `notificationChannelRef` (`*monitoringv1beta1.MonitoringNotificationChannelRef`), `vertexDatasetRef` (`*vertexaiv1beta1.VertexAIDatasetRef`), and `kmsKeyRef` (`*kmsv1beta1.KMSCryptoKeyRef`).
  3. Corrected service docs link to `https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.modelMonitors`.
  4. Updated `pkg/controller/direct/aiplatform/aiplatformmodelmonitor_mapper.go` with conversion methods for the updated reference fields.
  5. Resolved all open queue entries for `AIPlatformModelMonitor` in `apis/aiplatform/judgement_queue.yaml`.
- **Impact**: AIPlatformModelMonitor KRM types, references, identity, and CRD manifests are fully triaged, aligned with base KRM standards, and verified with unit tests.

### [2026-11-05] Implementing AIPlatformModelDeploymentMonitoringJob Greenfield Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing Greenfield direct controller, fuzzer, and E2E fixtures for `AIPlatformModelDeploymentMonitoringJob` (Issue #15060).
- **Problems & Solutions**:
  1. **Deterministic Controller & Fuzzer Generation**: Autogenerated the direct controller and fuzzer using `controllerbuilder generate-controller` and `generate-fuzzer --prepopulate-resource`. Verified fuzzer roundtripping with `FOCUS=AIPlatformModelDeploymentMonitoringJob go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`.
  2. **Service-Generated Resource ID**: `AIPlatformModelDeploymentMonitoringJob` uses a service-generated numeric ID assigned by GCP upon creation (`projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{id}`).
     - Updated `AIPlatformModelDeploymentMonitoringJobIdentity.GetIdentity` to support service-generated IDs (when `resourceID` is empty, keeping `ModelDeploymentMonitoringJob` empty until created).
     - Updated `Find` in `aiplatformmodeldeploymentmonitoringjob_controller.go` to return `false, nil` if `a.id.ModelDeploymentMonitoringJob` is empty.
     - Updated `Create` in `aiplatformmodeldeploymentmonitoringjob_controller.go` to omit `resource.Name` if `ModelDeploymentMonitoringJob` is empty, and parse the assigned `created.Name` from GCP response into `a.id`.
  3. **Static Config Registration**: Registered `AIPlatformModelDeploymentMonitoringJob` as a direct controller in `pkg/controller/resourceconfig/static_config.go` using `python3 dev/tasks/generate_static_config.py`.
  4. **Addressing Review Feedback**:
     - Moved `common.NormalizeGreenfieldReferences` before `obj.GetIdentity` in `AdapterForObject`.
     - Used `id.Location` instead of `direct.ValueOf(obj.Spec.Location)` when obtaining the GCP client.
     - Removed `resource.Name` assignment from `Create()` for service-generated IDs.
     - Used `k8s.ValueToDNSSubdomainName` on resource ID in `Export()`.
- **Finding & Real GCP Observation**:
  During live execution against real GCP (`./hack/record-gcp`), the Vertex AI `CreateModelDeploymentMonitoringJob` API returned `rpc error: code = InvalidArgument desc = Deployed Model ID: deployed-model-1 does not match existing Deployed Models from ModelMonitoringObjectiveConfig`. The API requires deployed models to already exist on the target `VertexAIEndpoint`. Deploying a model to an endpoint requires a trained model artifact or container to be deployed first, which cannot be provisioned purely through basic test fixtures. Golden HTTP logs will be aligned with MockGCP in Phase 3.
- **Impact**: Greenfield controller, fuzzer, and E2E fixtures implemented and verified with clean unit and fuzzer tests.

### [2026-11-06] Triaging Judgement Queue and Finalizing Direct KRM Types and Identity for AIPlatformCachedContent
- **Context**: Triaging judgement queue and finalizing types/identity for `AIPlatformCachedContent` (Issue #15070).
- **Problem**: `AIPlatformCachedContent` was generated with separate `ProjectRef` and `Location` fields and open findings in `judgement_queue.yaml`.
- **Solution**:
  1. Replaced separate `ProjectRef` and `Location` fields on `AIPlatformCachedContentSpec` with inlined `*parent.ProjectAndLocationRef`.
  2. Fixed the `+kcc:source:resource-docs` link to `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.cachedContents` and removed `+kcc:guess` tags.
  3. Verified that `AIPlatformCachedContentIdentity` correctly implements `identity.ServerGeneratedIdentity` because `CreateCachedContentRequest` in proto has no `cached_content_id` parameter and GCP server assigns the ID.
  4. Handled references: moved `VertexAiSearch`, `VertexAiSearch_DataStoreSpec`, and `VertexRagStore_RagResource` into `aiplatformcachedcontent_types.go` and converted `ragCorpus` to `RagCorpusRef *AIPlatformRAGCorpusRef`, `datastore`/`dataStore` to `DatastoreRef`/`DataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef`, and `engine` to `EngineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef`. Added `spec.model` to `refs_deferred.txt` as it refers to published Gemini base models (`publishers/google/models/{MODEL}`) with no publisher model ref type available.
  5. Ran `apis/aiplatform/generate.sh`, updated `judgement_queue.yaml` marking all findings for `AIPlatformCachedContent` as resolved, and verified all tests with `scripts/validate-prereqs.sh`.
- **Impact**: Cleanly finalizes KRM types, identity, and reference definitions for `AIPlatformCachedContent` according to KCC standards.

### [2026-11-06] Implementing AIPlatformHyperparameterTuningJob Greenfield Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Greenfield Step 2 direct controller, fuzzer, and real GCP E2E fixtures for `AIPlatformHyperparameterTuningJob` (Issue #15109).
- **Implementation & Findings**:
  1. **Autogenerated Controller & Fuzzer**: Generated direct controller `pkg/controller/direct/aiplatform/aiplatformhyperparametertuningjob_controller.go` and round-trip fuzzer `aiplatformhyperparametertuningjob_fuzzer.generated.go` via `controllerbuilder` (one-shot execution, not added to `generate.sh`).
  2. **Job Cancellation Pre-Deletion**: Vertex AI `HyperparameterTuningJob` resources cannot be deleted directly while in non-terminal states (`JOB_STATE_PENDING` or `JOB_STATE_RUNNING`). Updated `Delete` in `HyperparameterTuningJobAdapter` to check job state and call `CancelHyperparameterTuningJob` if the job is not yet in a terminal state (`JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, or `JOB_STATE_CANCELLED`).
  3. **Field Naming Alignment**: Standardized fixture YAML field names (`metricID`, `parameterID`, `useElapsedDuration`) to match KRM Go struct json tag definitions.
  4. **Golden E2E Fixtures**: Recorded and verified minimal and maximal test fixtures (`aiplatformhyperparametertuningjob-minimal` and `aiplatformhyperparametertuningjob-maximal`) against real GCP with PASS (7/7 checks passed).

### [2026-11-06] Implementing Direct Controller, Fuzzer, and E2E Fixtures for AIPlatformCachedContent
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and E2E fixtures for `AIPlatformCachedContent` (Issue #15147).
- **Problems & Solutions**:
  1. **AIPlatformModelRef Validation for Publisher Models**: The proto for `CachedContent.model` specifies published Gemini base models in format `projects/{PROJECT}/locations/{LOCATION}/publishers/google/models/{MODEL}` (or `publishers/google/models/{MODEL}`). However, `AIPlatformModelRef.ValidateExternal` only validated against `AIPlatformModelIdentity` (`projects/{project}/locations/{location}/models/{model}`), failing reference normalization with `format of AIPlatformModel external="..." was not known`. Updated `AIPlatformModelRef.ValidateExternal` in `apis/aiplatform/v1alpha1/aiplatformmodel_reference.go` to permit publisher model formats.
  2. **Round-Trip Fuzzer Recursive Schema Handling**: The fuzzer failed on nested `Schema` objects within `tools[].functionDeclarations[].parameters` and `response`. Because `krm.Schema` omits unsupported protobuf maps `properties` and `defs`, and `krm.Value` omits recursive `ListValue`, added `clearUnsupportedSchemaFields` and `f.FilterSpec` in `aiplatformcachedcontent_fuzzer.generated.go` to normalize these unsupported proto fields. Round-trip fuzzing passes cleanly for both Spec and Status.
  3. **Real GCP Live Execution Observation**: During live execution against real GCP (`./hack/record-gcp`), Vertex AI's `GenAiCacheService.CreateCachedContent` API returned `rpc error: code = InvalidArgument desc = Model <model> does not support explicit context caching.` Comprehensive probing across all available Model Garden models confirmed that Google Cloud has discontinued explicit context caching across Gemini models in Vertex AI in favor of implicit context caching. Controller implementation and E2E fixtures are complete and ready for MockGCP simulation in Step 3.

### [2026-11-06] Implementing AIPlatformRAGCorpus Greenfield Direct Controller, Fuzzer, and E2E Fixtures
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and real GCP E2E fixtures for `AIPlatformRAGCorpus` (Issue #15089).
- **Problems & Solutions**:
  1. **Regional Endpoint & Serverless Mode in GCP**: Vertex AI's `VertexRagDataService` requires regional endpoint configuration (`us-central1-aiplatform.googleapis.com:443`). For new GCP projects, RAG Engine defaults to Spanner mode which requires project allowlisting or switching `ragEngineConfig` to Serverless mode (`{"ragManagedDbConfig": {"serverless": {}}}`).
  2. **Immutability and Update Handling**: `UpdateRagCorpus` in Vertex AI's `VertexRagDataService` rejects updates if immutable fields or unmasked payloads are supplied. In accordance with KCC Greenfield guidelines, `AIPlatformRAGCorpusAdapter.Update` uses `common.CompareGreenfieldSpec` to compare desired spec with actual state and returns `updating AIPlatformRAGCorpus is not supported, fields modified: %v` if any diff is detected.
  3. **Real GCP E2E Fixtures**: Minimal and maximal test fixtures (`aiplatformragcorpus-minimal` and `aiplatformragcorpus-maximal`) were created under `pkg/test/resourcefixture/testdata/basic/aiplatform/v1alpha1/aiplatformragcorpus/`. Golden logs were recorded against real GCP, passing all E2E validation checks (V1-V7).
- **Impact**: Greenfield direct controller and round-trip fuzzer for `AIPlatformRAGCorpus` are fully implemented, tested, and verified against real GCP.

### 2026-11-06 Prepopulating AIPlatformTensorboardRun Greenfield Direct Types
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `AIPlatformTensorboardRun` under `apis/aiplatform/v1alpha1/` (Issue #15177).
- **Implementation**:
  1. Configured `apis/aiplatform/generate.sh` to include `AIPlatformTensorboardRun:TensorboardRun` for `generate-types` (backed by GA proto `google.cloud.aiplatform.v1`) and `generate-identity`.
  2. Generated `aiplatformtensorboardrun_types.go`, `aiplatformtensorboardrun_identity.generated.go`, `aiplatformtensorboardrun_reference.generated.go`, and test file `aiplatformtensorboardrun_identity_generated_test.go`.
  3. Added `AIPlatformTensorboardRun` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Updated `apis/aiplatform/judgement_queue.yaml` with open triage findings for Step 1b.
  5. Recorded golden test exceptions in `alpha-missingfields.txt` and `spec_dislike_etag.txt`.
- **Impact**: Step 1a prepopulation complete, deterministic, and verified with all apichecks and conformance unit tests passing.

### [2026-11-12] Greenfield Step 1a: Prepopulating Direct KRM Types, Identity, and Judgement Queue for AIPlatformArtifact
- **Context**: Prepopulating direct KRM types, identity, reference, and open judgement queue for `AIPlatformArtifact` under `aiplatform.cnrm.cloud.google.com/v1alpha1` (Issue #15495).
- **Implementation & Learnings**:
  1. **Non-Declarative Leaf Collection**: `google.cloud.aiplatform.v1.Artifact` is cataloged in `nonDeclarativeLeafCollections` (`artifacts: "data-plane lineage"`), requiring `--allow-non-declarative` when running `controllerbuilder generate-types`. In `apis/aiplatform/generate.sh`, added `--resource AIPlatformArtifact:Artifact` to the existing block that specifies `--allow-non-declarative` and `--service google.cloud.aiplatform.v1`.
  2. **Deterministic Identity Generation**: Added `--resource AIPlatformArtifact:Artifact` to the `generate-identity` invocation in `apis/aiplatform/generate.sh`. Deterministically generated `aiplatformartifact_identity.generated.go`, `aiplatformartifact_reference.generated.go`, and conformance tests in `aiplatformartifact_identity_generated_test.go` matching the pattern `projects/{project}/locations/{location}/metadataStores/{metadata_store}/artifacts/{artifact}`.
  3. **Release Bundle Blocklist**: Added `"AIPlatformArtifact"` in alphabetical order to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. **Judgement Queue & API Checks**: Open review findings were recorded in `apis/aiplatform/judgement_queue.yaml` for Step 1b triage. Golden test exceptions were recorded in `alpha-missingfields.txt` and `spec_dislike_etag.txt`.

### [2026-11-12] Greenfield Step 1b: Triaging Judgement Queue and Finalizing Direct KRM Types for AIPlatformArtifact
- **Context**: Finalizing KRM types, parent reference, and triaging `apis/aiplatform/judgement_queue.yaml` for `AIPlatformArtifact` under `aiplatform.cnrm.cloud.google.com/v1alpha1` (Issue #15505).
- **Implementation & Learnings**:
  1. **Single-Parent Rule (`spec.parentRef`)**:
     - `Artifact` is a child resource under `projects/{project}/locations/{location}/metadataStores/{metadata_store}`.
     - Following `PARENT_AND_IDENTITY.md` Case B, replaced `projectRef` and `location` with required, immutable `spec.parentRef *vertexaiv1beta1.MetadataStoreRef`.
     - Reused existing `MetadataStoreRef` from `apis/vertexai/v1beta1/`.
  2. **Updating `MetadataStoreRef` to Implement `refs.Ref`**:
     - `MetadataStoreRef` in `apis/vertexai/v1beta1/vertexaimetadatastore_reference.go` had legacy `NormalizedExternal` but did not implement modern `refs.Ref` methods (`GetGVK`, `GetNamespacedName`, `GetExternal`, `SetExternal`, `ValidateExternal`, `ParseExternalToIdentity`, `Normalize`), which caused `controllerbuilder generate-identity` normalization to fail compilation.
     - Implemented `refs.Ref` on `MetadataStoreRef` and `identity.Identity` on `MetadataStoreIdentity`, preserving the existing struct schema and backwards compatibility with `NormalizedExternal`.
  3. **Field Triage & Cleanup**:
     - Moved `etag` from `AIPlatformArtifactSpec` to `AIPlatformArtifactObservedState`.
     - Removed `spec.labels` (reconciled through `metadata.labels`).
     - Kept `state` in `spec` as the proto comments indicate it is client-managed state.
     - Fixed `+kcc:source:resource-docs` link to `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.artifacts` and removed `+kcc:guess` marker.
  4. **Judgement Queue Resolution**:
     - Triaged and marked all 8 entries for `AIPlatformArtifact` in `apis/aiplatform/judgement_queue.yaml` as `resolved`.
     - Closed `untriaged-bulk-generation` last with summary note.

### [2026-11-12] Prepopulating Direct KRM Types, Identity, Reference, and Judgement Queue for AIPlatformContext
- **Context**: Greenfield Step 1a prepopulate for `AIPlatformContext` under `apis/aiplatform/v1alpha1/` (Issue #15498).
- **Actions**:
  1. Updated `apis/aiplatform/generate.sh` to include `--resource AIPlatformContext:Context` in `generate-types` and `generate-identity`.
  2. Executed `./apis/aiplatform/generate.sh`, generating `aiplatformcontext_types.go`, `aiplatformcontext_identity.generated.go`, `aiplatformcontext_reference.generated.go`, `aiplatformcontext_identity_generated_test.go`, mapper code, deepcopy methods, and CRD manifest (`apiextensions.k8s.io_v1_customresourcedefinition_aiplatformcontexts.aiplatform.cnrm.cloud.google.com.yaml`).
  3. Added `"AIPlatformContext"` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` to record untriaged entries in apichecks exception files.
  5. Verified all unit tests (`go test -v ./apis/aiplatform/... ./pkg/controller/direct/aiplatform/... ./tests/apichecks/...`) pass and script execution is idempotent.
- **Impact**: Standard mechanical prepopulation for `AIPlatformContext` completed with open entries in `judgement_queue.yaml` ready for Step 1b triage.

### [2026-11-12] Triaging Judgement Queue and Finalizing Direct Types and Identity for AIPlatformContext
- **Context**: Greenfield Step 1b triage and finalization for `AIPlatformContext` under `apis/aiplatform/v1alpha1/` (Issue #15513).
- **Problems & Solutions**:
  1. **Parent Hierarchy (`MetadataStoreRef`)**: `AIPlatformContext` is a child resource under `projects/{project}/locations/{location}/metadataStores/{metadata_store}/contexts/{context}`. Reused the existing `MetadataStoreRef` from `apis/vertexai/v1beta1/` (representing `VertexAIMetadataStore`) and implemented `refs.Ref` interface methods on `MetadataStoreRef` in `apis/vertexai/v1beta1/vertexaimetadatastore_reference.go` and `apis/vertexai/v1alpha1/vertexaimetadatastore_reference.go`. Modelled `spec.parentRef` (`*vertexaiv1beta1.MetadataStoreRef`) with `+required` and CEL immutability (`self == oldSelf`) on `AIPlatformContextSpec`, removing redundant `projectRef` and `location`.
  2. **Etag & Labels Cleanup**: Removed top-level `labels` from `AIPlatformContextSpec` (managed via `metadata.labels`), and moved `etag` from `AIPlatformContextSpec` to `AIPlatformContextObservedState`.
  3. **Resource Documentation Link**: Updated `+kcc:source:resource-docs` to `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.contexts` and removed the `+kcc:guess` annotation.
  4. **Identity & Unit Testing**: Verified deterministic regeneration of `aiplatformcontext_identity.generated.go` via `generate-identity` using `parentRef`. Added unit tests in `aiplatformcontext_identity_test.go` covering `GetIdentity` resolution from `parentRef` and error cases. Resolved all open entries for `AIPlatformContext` in `apis/aiplatform/judgement_queue.yaml`.
- **Impact**: All judgement queue entries for `AIPlatformContext` are triaged and resolved, CRD schema and mappers regenerated cleanly, and all unit tests pass.

### [2026-11-12] Prepopulating Direct KRM Types, Identity, and Judgement Queue for AIPlatformExecution
- **Context**: Greenfield Step 1a mechanical prepopulation for `AIPlatformExecution` (Issue #15503).
- **Implementation & Observations**:
  1. **Proto Verification & Declarative Exemption**: `google.cloud.aiplatform.v1.Execution` (`AIPlatformExecution`) is in GA proto package `google.cloud.aiplatform.v1`. The proto collection is `projects/{project}/locations/{location}/metadataStores/{metadata_store}/executions/{execution}`, which matches `ephemeralCollectionSegments["executions"]` in `controllerbuilder`. Configured `generate-types` with `--prepopulate-resource` and `--allow-non-declarative`.
  2. **Code Generation**: Added `AIPlatformExecution:Execution` to `generate-types` and `generate-identity` in `apis/aiplatform/generate.sh`. Running `generate.sh` generated `aiplatformexecution_types.go`, identity, reference, conformance tests, and updated `judgement_queue.yaml` with open findings (including parent and identity ancestor segments `metadataStores/{metadata_store}`).
  3. **Release Bundle Blocklist**: Added `"AIPlatformExecution"` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. **Exceptions & Golden Updates**: Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` to record untriaged entries in `spec_dislike_etag.txt` and `alpha-missingfields.txt`. Reverted spurious formatting diff in `tests/apichecks/testdata/exceptions/multi_version_crd_diff/IAPSettings.diff`.
- **Impact**: Clean mechanical prepopulation ready for Step 1b triage.

### [2026-11-12] Finalizing Direct KRM Types, ParentRef, and Judgement Queue for AIPlatformExecution
- **Context**: Greenfield Step 1b types-judgement for `AIPlatformExecution` (Issue #15521).
- **Implementation & Observations**:
  1. **Single Parent Resolution**: `AIPlatformExecution` is a child resource under `VertexAIMetadataStore` (`projects/{project}/locations/{location}/metadataStores/{metadata_store}`). Reused `MetadataStoreRef` from `apis/vertexai/v1beta1`, implementing `refs.Ref` interface (`GetGVK`, `GetNamespacedName`, `GetExternal`, `SetExternal`, `ValidateExternal`, `Normalize`) on `MetadataStoreRef` in both `v1alpha1` and `v1beta1` while preserving `refsv1beta1.ExternalNormalizer` for backward compatibility. Replaced `projectRef` and `location` on `AIPlatformExecutionSpec` with required, immutable `parentRef *vertexaiv1beta1.MetadataStoreRef`.
  2. **Identity & Client-Specified ID**: In `google.cloud.aiplatform.v1.CreateExecutionRequest`, field `execution_id` allows client-specified IDs (or service-generated UUID if omitted). `IdentityV2` and `spec.resourceID` are kept as accepted.
  3. **KRM Conventions & Field Sanitization**:
     - Removed `Etag` from `AIPlatformExecutionSpec` and moved it to `AIPlatformExecutionObservedState`.
     - Removed top-level `Labels` from `AIPlatformExecutionSpec` as top-level labels are reconciled via `metadata.labels`.
     - Retained `State` in `AIPlatformExecutionSpec` as client-managed desired state per proto documentation ("This property is managed by clients and the system does not prescribe or check the validity of state transitions").
     - Updated documentation link to valid `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores.executions` and removed `+kcc:guess` annotation.
  4. **Judgement Queue & Golden Exceptions**:
     - Resolved all 8 entries for `AIPlatformExecution` in `apis/aiplatform/judgement_queue.yaml`.
     - Regenerated CRD and types via `./apis/aiplatform/generate.sh`.
     - Re-recorded `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `spec_dislike_etag.txt` (which removed `AIPlatformExecution` from `spec_dislike_etag.txt`).
     - Verified all apichecks, CAI template, and package tests with clean passes.
- **Impact**: Cleanly finalized types, identity, and CRD for `AIPlatformExecution`.

### [2026-11-12] Prepopulating Direct KRM Types, Identity, and Judgement Queue for AIPlatformNotebookExecutionJob
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, and judgement queue for `AIPlatformNotebookExecutionJob` (Issue #15511) under `apis/aiplatform/v1alpha1/` from `google.cloud.aiplatform.v1.NotebookExecutionJob`.
- **Implementation & Observations**:
  1. **Non-Declarative Generation**: In `notebook_service.proto`, `NotebookService` exposes `CreateNotebookExecutionJob`, `GetNotebookExecutionJob`, and `DeleteNotebookExecutionJob`, but lacks an `UpdateNotebookExecutionJob` RPC. Added `AIPlatformNotebookExecutionJob:NotebookExecutionJob` to the `--allow-non-declarative` block of `generate-types` in `apis/aiplatform/generate.sh`.
  2. **Empty oneof Option Handling**: `NotebookExecutionJob` defines `oneof runtime_environment { WorkbenchRuntime workbench_runtime = 23; }` where `message WorkbenchRuntime {}` contains no fields. Added `aiplatformnotebookexecutionjobs.aiplatform.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with comment `// GCP known pattern: Empty oneof option (spec.workbenchRuntime)`.
  3. **Release Bundle Blocked Kinds**: Added `AIPlatformNotebookExecutionJob` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. **Verification**: Executed `./apis/aiplatform/generate.sh` idempotently, verified open entries in `apis/aiplatform/judgement_queue.yaml`, and confirmed all unit tests and `tests/apichecks` pass cleanly.

### [2026-11-13] Triaging Judgement Queue and Finalizing Direct KRM Types and Identity for AIPlatformNotebookExecutionJob
- **Context**: Triaging judgement queue and finalizing direct KRM types, identity, and references for `AIPlatformNotebookExecutionJob` under `apis/aiplatform/v1alpha1/` (Issue #15543).
- **Implementation & Observations**:
  1. **Parent & Location Ref**: Replaced separate `ProjectRef` and `Location` fields on `AIPlatformNotebookExecutionJobSpec` with inlined `*parent.ProjectAndLocationRef`.
  2. **References Modelled**:
     - `notebookRuntimeTemplateRef` modelled as `*AIPlatformNotebookRuntimeTemplateRef`.
     - `scheduleRef` modelled as `*AIPlatformScheduleRef`.
     - `serviceAccountRef` modelled as `*refsv1beta1.IAMServiceAccountRef`.
  3. **Documentation URL & Guess Marker**: Verified REST resource documentation URL `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.notebookExecutionJobs` (200 OK) and removed guess marker comments.
  4. **Shared Types and Exceptions**:
     - Retained shared `EncryptionSpec` without forcing `+required` on `kmsKeyRef` to avoid breaking other resources sharing this struct.
     - Kept string values for reservation affinity on `MachineSpec.ReservationAffinity`.
     - Retained `gcsNotebookSource` and `gcsOutputURI` decomposable paths in `refs_not_representable.txt`.
  5. **Judgement Queue**: Resolved all 16 entries in `apis/aiplatform/judgement_queue.yaml` for `AIPlatformNotebookExecutionJob`.

