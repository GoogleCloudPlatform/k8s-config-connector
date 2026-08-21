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
