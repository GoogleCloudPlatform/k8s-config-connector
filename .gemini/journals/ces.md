### [2026-06-03] Handling Greenfield Protos Added After Pinned Commit
- **Context**: Implementing `CESApp` types (Google Cloud CES v1beta) under `ces.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The `google/cloud/ces` proto files did not exist in the pinned googleapis SHA `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb`. Therefore, compiling `googleapis.pb` with the standard script left out the CES package, causing `generate-types` to fail.
- **Solution**: Modified `dev/tools/controllerbuilder/generate-proto.sh` to checkout `google/cloud/ces/v1beta` from `origin/master` after resetting to the pinned SHA, and added the path explicitly to the protoc compile command. This allowed us to successfully compile the `googleapis.pb` descriptor containing the ces service definitions.
- **Impact**: Any subsequent developer or automated pipeline running `./apis/ces/v1alpha1/generate.sh` can successfully regenerate types without missing proto errors.

### [2026-11-11] Prepopulating CESApp with GA Proto and Recursive Schema Handling
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CESApp` using `google.cloud.ces.v1` and `--prepopulate-resource`.
- **Problem**:
  1. `google.cloud.ces.v1.App` depends on `google.protobuf.Value`, requiring external proto mapping. In generated mappers, `direct.Struct_ToProto` expects `*apiextensionsv1.JSON` while `Value.StructValue` is non-pointer `apiextensionsv1.JSON`.
  2. `App.VariableDeclaration.Schema` is a recursive JSON schema message (`SchemaWithRequired` / `SchemaObservedState`) that causes `controller-gen` to emit empty `{}` OpenAPI schemas, failing `envtest` CRD validation in `TestPreview`.
- **Solution**:
  1. Added `,google.protobuf` to `--service` in `apis/ces/generate.sh` and implemented a minimal handwritten `Value_ToProto` helper in `pkg/controller/direct/ces/mapper.go`.
  2. Defined `App_VariableDeclaration` and `App_VariableDeclarationObservedState` in `cesapp_types.go` with `+kubebuilder:validation:Schemaless`, `+kubebuilder:validation:XPreserveUnknownFields`, and `+kubebuilder:validation:Type=object` on the `Schema` field to terminate the recursion cycle safely in OpenAPI v3 schema.
- **Impact**: Clean CRD generation, passing envtest, and full idempotency with `apis/ces/generate.sh`.

### [2026-11-11] Triaging Judgement Queue and Finalizing CESApp Direct Types
- **Context**: Triaging `apis/ces/judgement_queue.yaml` entries and finalizing direct KRM types and identity for `CESApp` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Decisions**:
  1. Replaced separate `projectRef` and `location` fields on `CESAppSpec` with inlined `*parent.ProjectAndLocationRef`.
  2. Converted GCP references to KCC types:
     - `rootAgent` -> `rootAgentRef` (`*CesAgentRef`)
     - `clientCertificateSettings.privateKey` -> `privateKeyRef` (`*refsv1beta1.SecretManagerSecretVersionRef`, required)
     - `clientCertificateSettings.passphrase` -> `passphraseRef` (`*refsv1beta1.SecretManagerSecretVersionRef`)
     - `loggingSettings.audioRecordingConfig.gcsBucket` and `evaluationAudioRecordingConfig.gcsBucket` -> `gcsBucketRef` (`*storagev1beta1.StorageBucketRef`)
     - `loggingSettings.bigqueryExportSettings.project` -> `projectRef` (`*refsv1beta1.ProjectRef`)
     - `loggingSettings.redactionConfig.inspectTemplate` -> `inspectTemplateRef` (`*dlpv1alpha1.DLPInspectTemplateRef`)
     - `loggingSettings.redactionConfig.deidentifyTemplate` -> `deidentifyTemplateRef` (`*dlpv1alpha1.DLPDeidentifyTemplateRef`)
     - `guardrails` -> `guardrailRefs` (`[]CesGuardrailRef`), implemented via external-only `CesGuardrailIdentity` and `CesGuardrailRef`.
  3. Pruned unshared deprecated fields: `PrebuiltAmbientNoise` in `AmbientSoundConfig`, `DisableBargeIn` in `BargeInConfig`, `HallucinationMetricBehavior` in `EvaluationMetricsThresholds`, and `FallbackAction` in `LanguageSettings`.
  4. Resolved all 24 open `CESApp` entries in `apis/ces/judgement_queue.yaml`.
- **Impact**: All unit and apichecks tests pass cleanly without naming or CRD backward-compatibility regressions.

### [2026-11-13] Direct Controller, Fuzzer, and E2E Fixtures for CESApp
- **Context**: Implementing Greenfield Step 2 for `CESApp` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Service Behaviors**:
  1. **Supported Locations**: The Customer Engagement Suite (CES) API (`ces.googleapis.com/v1`) only supports multi-region locations `us` and `eu`. Requests directed to `global` return HTTP 403 Forbidden. Test fixtures must specify `location: us` (or `eu`).
  2. **LRO Behavior**: `CreateApp` and `DeleteApp` return Long-Running Operations (LROs), handled via `op.Wait(ctx)`. `UpdateApp` is synchronous and returns `*cespb.App`.
  3. **Language Code Normalization**: The CES API automatically normalizes short language codes like `en` or `es` to full BCP-47 tags (`en-US`, `es-ES`). Fixture YAMLs should declare canonical BCP-47 tags in `languageSettings` to avoid spurious diffs.
  4. **VariableDeclarations Schema**: `App_VariableDeclaration.Schema` is an OpenAPI 3.0 recursive schema object with `type Schema struct {}` in KRM (`Schemaless` and `PreserveUnknownFields`). Since `Schema_ToProto` in `mapper.go` is stubbed and `schema.type` is required by GCP when variable declarations are present, `variableDeclarations` is omitted from maximal fixtures until full schema serialization is implemented.
  5. **Audit Probe & E2E Validation**: Golden HTTP logs and audit probes were recorded against real GCP project `cnrm-barni-4` for `cesapp-minimal` and `cesapp-maximal`, with all 7/7 automated lifecycle checks passing cleanly.


