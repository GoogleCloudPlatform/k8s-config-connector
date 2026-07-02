# VertexAINotebookRuntime Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Mapping to GCP Resource**:
   - `NotebookRuntime` is a resource of `aiplatform.googleapis.com` (Vertex AI API). Following the established naming pattern under `apis/aiplatform/v1alpha1`, we mapped this to the KCC Kind `VertexAINotebookRuntime:NotebookRuntime`.

2. **Handwritten Mapper Support for References**:
   - Because `NotebookRuntime` requires a template reference during assignment, we defined `NotebookRuntimeTemplateRef` in its own file `vertexainotebookruntime_reference.go`.
   - The code generator maps `notebook_runtime_template_ref` in proto to our custom Go type but expects handwritten mappers `NotebookRuntimeTemplateRef_FromProto` and `NotebookRuntimeTemplateRef_ToProto` to avoid generation errors.
   - We created `notebookruntime_mapper.go` under `pkg/controller/direct/aiplatform/` to bridge this conversion.

3. **Verification**:
   - Built the packages and ran unit tests to ensure `VertexAINotebookRuntimeIdentity` parses and formats the GCP resource URL (`projects/{project}/locations/{location}/notebookRuntimes/{notebookRuntime}`) successfully.
   - Verified that the `gcpurls` registration tests successfully validate our newly added identity template.
