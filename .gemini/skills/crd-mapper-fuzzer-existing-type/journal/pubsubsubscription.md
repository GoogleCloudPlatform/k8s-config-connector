# PubSubSubscription KRM Types Journal

## Observations & Learnings

1. **Leveraging Automated Mapper Generation via Proto Annotations**:
   By annotating nested structs (like `BigQueryConfig`, `CloudStorageConfig`, `PushConfig`, `RetryPolicy`, etc.) with their canonical GCP protobuf counterpart messages, e.g., `// +kcc:proto=google.pubsub.v1.BigQueryConfig`, the generator automatically generated all required `FromProto`/`ToProto` conversion functions. This saved us from writing hundreds of lines of hand-coded mappers.

2. **Resolving Proto and KRM Schema Field Name Deviations**:
   The generator is sensitive to naming conventions. When a field is defined under KRM with custom names (like `tableRef`), mapping functions are expected to match. In our case, registering the sub-structs allowed the automated mapper generation to seamlessly map KRM's nested objects without requiring manual mapper intervention.

3. **Strict Schema Validation & CRD Compatibility**:
   Running `dev/tasks/diff-crds` on the newly generated KRM types verified that the schema compiled identically to the baseline CRD. Removing `observedState` and standard status fields kept the generated OpenAPI schema 100% compliant with the original definition.

4. **Fuzzer Customization**:
   Since Pub/Sub includes several advanced GCP features (such as tags, detached status, and analytical hub integrations) that are not currently exposed in KRM, we utilized `f.Unimplemented_NotYetTriaged` on these proto paths to successfully round-trip and bypass fuzzer mismatches.
