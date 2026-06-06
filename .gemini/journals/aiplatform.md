### [2026-06-06] Consolidating generate-types for Multi-Service Packages & Addressing Status Types
- **Context**: Implementing `VertexAIReasoningEngine` types and IdentityV2 in `apis/vertexai/v1alpha1`.
- **Problem**: When a KRM api directory manages resources that span multiple proto services (e.g. `google.cloud.aiplatform.v1beta1` and `google.cloud.aiplatform.v1`), having separate sequential runs of `generate-types` in `generate.sh` causes the later runs to overwrite `types.generated.go`, commenting out necessary types for previous resources as unreachable and failing CRD compilation (like `unknown type Status`). Additionally, older handwritten files may refer to the old local `Status` instead of `common.Status` imported from `apis/common`, which the new generator expects.
- **Solution**:
  1. Consolidated multiple `generate-types` runs in `generate.sh` into a single invocation. We achieved this by using the primary service (e.g., `--service google.cloud.aiplatform.v1beta1`) and referencing other services' resources with their fully-qualified names:
     ```bash
     --resource VertexAIDataLabelingJob:google.cloud.aiplatform.v1.DataLabelingJob \
     --resource VertexAIReasoningEngine:google.cloud.aiplatform.v1.ReasoningEngine
     ```
  2. Updated `datalabelingjob_types.go` to import `github.com/GoogleCloudPlatform/k8s-config-connector/apis/common` as `common` and map `Error` to `*common.Status`.
- **Impact**: Eliminates compilation issues when regenerating types for services with resources in multiple proto service versions. It also ensures proper use of the centralized `common.Status` type.
