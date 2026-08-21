# PubSubTopic Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Strict Schema Compatibility for References**:
   - The baseline `PubSubTopic` CRD does not define `projectRef` in its spec. When migrating, we added an optional `ProjectRef *refs.ProjectRef` pointer to provide the standard reference field without breaking backward compatibility.
   - For references like `kmsKeyRef` and `schemaRef` under `schemaSettings`, the baseline CRD lacks a `kind` field. We hand-coded custom reference types `KMSCryptoKeyRef` and `SchemaRef` directly in `topic_types.go`.
   - By matching their signatures as `"external,name,namespace"`, the OpenAPI `oneOf` reference constraints are automatically appended by the CRD generator to match the original baseline constraints.

2. **Removal of Unused Status Fields**:
   - The baseline `PubSubTopic` CRD status only includes `conditions` and `observedGeneration`. We excluded `externalRef` and `observedState` from `PubSubTopicStatus` in `topic_types.go` to keep the schema strictly compatible.

3. **Hand-Coded Mapper for Nested and Unreachable Types**:
   - Since `Topic` in `types.generated.go` is skipped by `generate-types` (due to being a direct resource type), nested fields like `MessageStoragePolicy` and `SchemaSettings` were commented out as unreachable types.
   - To resolve this and prevent any build failures in generated mappers, we implemented a custom, handcoded `topic_mapper.go` in `pkg/controller/direct/pubsub/`.
   - The generator successfully detected our custom `PubSubTopicSpec_FromProto` and `PubSubTopicSpec_ToProto` and safely omitted generating redundant/conflicting mapper code.

## Verification

- Successfully ran `dev/tasks/diff-crds` showing that only the expected optional `projectRef` addition was introduced to the baseline schema.
- Successfully ran `make manifests`, `make generate-go-client`, `make ensure`, and `go vet ./...` without compilation or lint issues.
