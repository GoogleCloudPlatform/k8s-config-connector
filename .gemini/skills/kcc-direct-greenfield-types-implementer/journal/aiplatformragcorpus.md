# Greenfield Implementation of AIPlatformRAGCorpus

## Observations & Design Decisions

### 1. Breaking Import Cycles for Nested References
The `RagVectorDbConfig_VertexVectorSearch` struct inside the GCP proto of `RagCorpus` references `Index` and `IndexEndpoint`, which are legacy Terraform-controlled resources (`VertexAIIndex` and `VertexAIIndexEndpoint` in KCC).
We wanted to define clean reference types `VertexAIIndexRef` and `VertexAIIndexEndpointRef` following KCC guidelines.
Normally, references for `vertexai` resources would live in the `apis/vertexai/v1alpha1` package. However, the `apis/vertexai/v1alpha1` package already imports `apis/aiplatform/v1alpha1` for other resources (such as `vertexaicustomjob_types.go`).
If `apis/aiplatform/v1alpha1` tried to import `apis/vertexai/v1alpha1`, it would create an **import cycle**.
To cleanly break this import cycle, we declared the `VertexAIIndexRef` and `VertexAIIndexEndpointRef` reference types, along with their identity helpers, directly within the `apis/aiplatform/v1alpha1` package while keeping their registered `GroupVersionKind` pointing to `vertexai.cnrm.cloud.google.com`. This successfully broke the dependency loop and preserved correct resource referencing.

### 2. DisplayName Type Mapping
Standard KCC fields (especially `DisplayName`) are mapped to pointers `*string` rather than plain `string` to ensure proper serialization and fallback behaviors. If a plain string is used, the generator expects a custom `AIPlatformRAGCorpusSpec_DisplayName_ToProto` helper function. Mapping `DisplayName` as `*string` and utilizing the `// +required` comment keeps the types idiomatic and uses the standard generated mappers.
