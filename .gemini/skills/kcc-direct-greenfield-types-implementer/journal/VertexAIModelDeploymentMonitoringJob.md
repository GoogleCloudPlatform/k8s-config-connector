# Journal: VertexAIModelDeploymentMonitoringJob

This journal chronicles observations and learnings from implementing the direct KRM types, identity, references, and generate.sh for `VertexAIModelDeploymentMonitoringJob`.

## Key Observations & Learnings

### 1. Handling of `google.protobuf.Value` in Greenfield Direct Controllers
In Vertex AI Model Deployment Monitoring Job proto, the field `sample_predict_instance` is a `google.protobuf.Value` (which corresponds to `*structpb.Value` in Go).
- **Initial spec generation**: The generator did not map this field natively and commented it out as unreachable.
- **Mapping strategy**: For KRM APIs, `google.protobuf.Value` fields are best represented using `*apiextensionsv1.JSON`. By manually specifying `SamplePredictInstance *apiextensionsv1.JSON` in `vertexaimodeldeploymentmonitoringjob_types.go`, we align the API schema with typical Kubernetes and Config Connector standards.
- **Mapper compilation**: Changing the type to `*apiextensionsv1.JSON` allowed the generator to emit standard calls to `apiextensionsv1.JSON_v1alpha1_FromProto` / `ToProto` inside the generated mapper.

### 2. Bypassing Go Package Function Boundaries with Type Aliasing
The generated mapper `mapper.generated.go` references `apiextensionsv1.JSON_v1alpha1_FromProto`, but since `apiextensionsv1` is an external Kubernetes package, we cannot define functions on it.
- **Solution**: We created a custom package `pkg/controller/direct/vertexai/apiextensionsv1` which defines type alias `type JSON = v1.JSON` and implements `JSON_v1alpha1_FromProto` and `JSON_v1alpha1_ToProto`. By importing this package as `apiextensionsv1` in `pkg/controller/direct/vertexai/mappers.go`, we satisfied the generated mapper's calls cleanly without modifying the code generator.

### 3. Reachability-Based Pruning in `types.generated.go`
The direct types generator comments out complex structures under `types.generated.go` wrapped in `/* unreachable type ... */` to keep binary sizes lean.
- To resolve compilation issues with unused or unreachable types, we manually declared our Spec/ObservedState fields in `vertexaimodeldeploymentmonitoringjob_types.go` referencing them, which automatically marked them as reachable and uncommented them in `types.generated.go` on subsequent generator runs.

### 4. Aligning Unrelated Promoted Status Fields
A peer direct mapper `aiplatform/vertexaispecialistpool_mapper.go` was failing to compile due to `krm.Status` being undefined. This occurred because `Status` was moved from `apis/vertexai/v1alpha1` to `apis/common` (`common.Status`). We updated the specialistpool mapper to import and use `common.Status` directly, resolving the repository-wide compilation breakages.
