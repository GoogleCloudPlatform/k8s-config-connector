# AIPlatformRAGCorpus Journal

- **Service**: AIPlatform (`google.cloud.aiplatform.v1`)
- **Kind**: `AIPlatformRAGCorpus`
- **Group**: `aiplatform.cnrm.cloud.google.com`
- **Proto**: `google.cloud.aiplatform.v1.RagCorpus`

## Observations & Learnings

1. **Identity & ID Generation**:
   - `CreateRagCorpusRequest` in `google.cloud.aiplatform.v1` does not accept a client-provided `rag_corpus_id` parameter.
   - Therefore, `AIPlatformRAGCorpus` uses `identity.ServerGeneratedIdentity`.

2. **Empty Struct CRD Validation (`TestCRDObjectTypes`)**:
   - `RagVectorDbConfig_RagManagedDb_Knn` corresponds to `message KNN {}` in proto which has no fields.
   - In Kubernetes OpenAPI CRD schemas, empty structs fail object property validation (`missing properties, additionalProperties, or x-kubernetes-preserve-unknown-fields`).
   - Per Greenfield guidelines, added `aiplatformragcorpuses.aiplatform.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with comment `// GCP known pattern: Empty oneof option (spec.vectorDbConfig.ragManagedDb.knn)`.

3. **Reference Conversion**:
   - `APIAuth_APIKeyConfig.apiKeySecretVersion` was converted to `APIKeySecretVersionRef` (`*refsv1beta1.SecretManagerSecretVersionRef`).
   - Fields such as `endpoint`, `index`, `indexEndpoint`, and `servingConfig` are either polymorphic or point to unrepresented GCP resource types, so they remain string fields and are recorded in `refs_not_representable.txt`.

4. **Resource Documentation Link**:
   - Updated source link in `aiplatformragcorpus_types.go` to REST reference `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.ragCorpora`.
