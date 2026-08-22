# VertexAICachedContent Types Journal

## Shortcomings & Solutions

### Unreachable standard types (Value, LatLng) becoming reachable

* **Context:** In `VertexAICachedContent`, nested structs like `Tool` -> `Schema` -> `Value` and `Retrieval` -> `VertexRagStore` -> `RagRetrievalConfig` -> `LatLng` reference standard protobuf types like `google.protobuf.Value` and `google.type.LatLng`.
* **Observation:** Previously, these standard protobuf types were commented out as "unreachable" in `types.generated.go`. Once `VertexAICachedContent` was added, they became reachable, which automatically uncommented them in `types.generated.go`. However, the generated mapper expected mapping functions like `Value_v1alpha1_FromProto` and `LatLng_v1alpha1_FromProto` to exist, but they were not automatically generated because the generator skips standard library / custom types, causing compilation errors.
* **Solution:** Manually implement `Value_v1alpha1_FromProto`, `Value_v1alpha1_ToProto`, `ListValue_v1alpha1_FromProto`, `ListValue_v1alpha1_ToProto`, `LatLng_v1alpha1_FromProto`, and `LatLng_v1alpha1_ToProto` in `pkg/controller/direct/vertexai/mappers.go`. This resolves the compilation error while ensuring robust serialization and translation of dynamic/structured values.
