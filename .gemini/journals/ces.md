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
  5. Audit Probe & E2E Validation: Golden HTTP logs and audit probes were recorded against real GCP project `cnrm-barni-4` for `cesapp-minimal` and `cesapp-maximal`, with all 7/7 automated lifecycle checks passing cleanly.

  ### [2026-11-15] Prepopulating CESDeployment
  - **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CESDeployment` using `google.cloud.ces.v1` and `--prepopulate-resource`.
  - **Learnings & Observations**:
  1. `Deployment` in `google.cloud.ces.v1` maps cleanly to `CESDeployment`.
  2. Identity pattern follows `projects/{project}/locations/{location}/apps/{app}/deployments/{deployment}`, with parent app resource `projects/{project}/locations/{location}/apps/{app}`.
  3. Generated `cesdeployment_identity.generated.go`, `cesdeployment_reference.generated.go`, and conformance tests cleanly.
  4. Updated `dev/tasks/build-release-bundle` with `CESDeployment` in `BLOCKED_KINDS` and updated `alpha-missingfields.txt` and `kind_name_violations.txt`.

### [2026-11-15] Triaging Judgement Queue and Finalizing CESDeployment Direct Types
- **Context**: Triaging `apis/ces/judgement_queue.yaml` entries and finalizing direct KRM types and identity for `CESDeployment` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Decisions**:
  1. Replaced separate `projectRef` and `location` fields on `CESDeploymentSpec` with required immutable `parentRef *CESAppRef` following `go/kcc-oneparent` and Case B of `PARENT_AND_IDENTITY.md`.
  2. Converted GCP reference `appVersion` (`ces.googleapis.com/AppVersion`) to `appVersionRef *CesAppVersionRef`, backed by external-only `CesAppVersionIdentity` and `CesAppVersionRef` in `apis/ces/v1alpha1/cesappversion_identity.go` and `apis/ces/v1alpha1/cesappversion_reference.go`.
  3. Resolved `channelProfile.profileID` as `not-applicable` because it is an inline profile identifier rather than an external GCP resource reference.
  4. Resolved all open `CESDeployment` entries in `apis/ces/judgement_queue.yaml`.
  5. Regenerated CRDs, mappers, deepcopy methods, and updated `alpha-missingfields.txt` and `naming_violations.txt` exceptions.

### [2026-11-16] Direct Controller, Fuzzer, and E2E Fixtures for CESDeployment
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and recording real-GCP E2E fixtures for `CESDeployment` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Service Behaviors**:
  1. **Resource ID Constraints**: Deployment ID in `projects/{project}/locations/{location}/apps/{app}/deployments/{deployment}` must match regular expression `[a-zA-Z0-9][a-zA-Z0-9-_]{4,35}` (5 to 36 characters). Using default `${metadata.name}` with `${uniqueId}` (e.g. `cesdeployment-minimal-${uniqueId}`, 38 chars) exceeds this limit and fails with HTTP 400. Setting explicit `spec.resourceID` (such as `cesdep-min-${uniqueId}` and `cesdep-max-${uniqueId}`) satisfies the length constraint.
  2. **Synchronous REST CRUD**: Unlike `CESApp` where creation and deletion are LROs, `CESDeployment` CRUD operations (`CreateDeployment`, `UpdateDeployment`, `GetDeployment`, `DeleteDeployment`) are synchronous REST calls returning `*cespb.Deployment` or `*emptypb.Empty` directly without LRO handling.
  3. **Draft App Version**: `spec.appVersionRef` points to `projects/{project}/locations/{location}/apps/{app}/versions/{version}`, and `versions/-` targets the draft app version.
  4. **Field Mutability**: `displayName` and `channelProfile` (and its subfields `personaProperty`, `disableDtmf`, `disableBargeInControl`, `noiseSuppressionLevel`, and `webWidgetConfig`) update cleanly in-place via `UpdateDeployment` with field mask.
  5. **Audit Probe & E2E Validation**: Golden HTTP logs and audit probes were recorded against real GCP project `cnrm-barni-4` for `cesdeployment-minimal` and `cesdeployment-maximal`, with all 7/7 automated lifecycle checks passing cleanly.

### [2026-11-16] MockGCP Implementation and Alignment for CESDeployment
- **Context**: Implementing Greenfield Step 3 MockGCP service and aligning golden logs for `CESDeployment` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Service Behaviors**:
  1. **Parent Validation**: Validated parent `projects/{project}/locations/{location}/apps/{app}` in `CreateDeployment` and returned `codes.NotFound` if parent app does not exist.
  2. **Etag Omission on GET/CREATE**: In real GCP, `CreateDeployment` and `GetDeployment` do not include an `etag` in their JSON responses (left empty string, omitted in proto3 serialization), whereas `UpdateDeployment` returns an `etag`. In MockGCP, `Etag` is left empty in `GetDeployment` and `CreateDeployment` and populated on `UpdateDeployment` (normalized to `abcdef0123A=`).
  3. **Synchronous CRUD**: Implemented `GetDeployment`, `CreateDeployment`, `UpdateDeployment`, `DeleteDeployment`, and `ListDeployments` in `mockces/deployment.go` returning `*cespb.Deployment` or `*emptypb.Empty` directly.
  4. **Test Harness GroupKind**: Enabled `schema.GroupKind{Group: "ces.cnrm.cloud.google.com", Kind: "CESDeployment"}` in `config/tests/samples/create/harness.go`.
  5. **Alignment & CI**: Verified `TestGoldenLogAlignment/ces/v1alpha1/cesdeployment` and `./dev/ci/presubmits/tests-e2e-fixtures-suite ces` with 100% passing results and zero diff.

### [2026-11-16] Prepopulating CESGuardrail
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CESGuardrail` using `google.cloud.ces.v1` and `--prepopulate-resource`.
- **Learnings & Observations**:
  1. `Guardrail` in `google.cloud.ces.v1` maps cleanly to `CESGuardrail`.
  2. Identity pattern is `projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}`, with parent app resource `projects/{project}/locations/{location}/apps/{app}`.
  3. Generated `cesguardrail_identity.generated.go`, `cesguardrail_reference.generated.go`, and conformance tests in `cesguardrail_identity_generated_test.go`.
  4. Retained external-only `CesGuardrailRef` and `CesGuardrailIdentity` in `cesguardrail_reference.go` and `cesguardrail_identity.go` for backward compatibility with `CESApp.spec.guardrailRefs`.
  5. Implemented `TriggerAction_Response_FromProto` and `TriggerAction_Response_ToProto` mapper helpers in `pkg/controller/direct/ces/mapper.go` to support nested action response mapping.
  6. Added `cesguardrails.ces.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` due to empty oneof option `spec.llmPromptSecurity.defaultSettings` (GCP proto pattern where `DefaultSecuritySettings` contains only output-only fields).
  7. Added `CESGuardrail` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` and synchronized golden exception files (`alpha-missingfields.txt`, `spec_dislike_etag.txt`, `kind_name_violations.txt`, and resolved entries in `naming_violations.txt`).

### [2026-11-16] Triaging Judgement Queue and Finalizing CESGuardrail Direct Types
- **Context**: Triaging `apis/ces/judgement_queue.yaml` entries and finalizing direct KRM types and identity for `CESGuardrail` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Decisions**:
  1. Replaced separate `projectRef` and `location` fields on `CESGuardrailSpec` with required immutable `parentRef *CESAppRef` following `go/kcc-oneparent` and Case B of `PARENT_AND_IDENTITY.md`.
  2. Moved `etag` from `CESGuardrailSpec` to `CESGuardrailObservedState` per `etag-in-spec` standards, and updated `spec_dislike_etag.txt`.
  3. Converted GCP reference `.spec.action.transferAgent.agent` (`ces.googleapis.com/Agent`) to `agentRef *CesAgentRef` inside `TriggerAction_TransferAgent`, moved `TriggerAction_TransferAgent` from `types.generated.go` to `cesguardrail_types.go`, and verified automatic proto mapper generation.
  4. Resolved all 6 open `CESGuardrail` entries in `apis/ces/judgement_queue.yaml`.
  5. Regenerated CRDs, mappers, deepcopy implementations, and updated `alpha-missingfields.txt` and `spec_dislike_etag.txt` testdata exceptions.

### [2026-11-17] Direct Controller, Fuzzer, and E2E Fixtures for CESGuardrail
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and recording real-GCP E2E fixtures for `CESGuardrail` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Service Behaviors**:
  1. **Synchronous REST CRUD**: Similar to `CESDeployment`, `CESGuardrail` CRUD operations (`CreateGuardrail`, `UpdateGuardrail`, `GetGuardrail`, `DeleteGuardrail`) are synchronous REST calls returning `*cespb.Guardrail` or `*emptypb.Empty` directly without LRO handling.
  2. **Oneof Structure**: `GuardrailType` is a proto oneof (`content_filter`, `llm_prompt_security`, `llm_policy`, `model_safety`, `code_callback`), and `Action` in `TriggerAction` is also a oneof (`respond_immediately`, `transfer_agent`, `generative_answer`). E2E fixtures test `contentFilter` with `respondImmediately` actions.
  3. **Field Mutability**: `displayName`, `description`, `enabled`, `action.respondImmediately.responses`, and `contentFilter` (including `matchType`, `disregardDiacritics`, `bannedContents`, `bannedContentsInUserInput`, and `bannedContentsInAgentResponse`) update in-place cleanly via `UpdateGuardrail` with field mask.
  4. **Audit Probe & E2E Validation**: Golden HTTP logs and audit probes were recorded against real GCP project `cnrm-barni-4` for `cesguardrail-minimal` and `cesguardrail-maximal`, with all 7/7 automated lifecycle checks passing cleanly.

### [2026-11-17] MockGCP Implementation and Alignment for CESGuardrail
- **Context**: Implementing Greenfield Step 3 (MockGCP and Alignment) for `CESGuardrail` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Implementation**:
  1. **Mock Service**: Added `mockgcp/mockces/guardrail.go` and `guardrailName` in `mockgcp/mockces/names.go`. Implemented synchronous CRUD (`CreateGuardrail`, `GetGuardrail`, `UpdateGuardrail`, `DeleteGuardrail`, `ListGuardrails`) for `AgentServiceV1`.
  2. **Etag Behavior**: Verified that `CreateGuardrail` returns the created guardrail without an `etag`, while `GetGuardrail` and `UpdateGuardrail` return with `etag`. Storage holds the etag and returns it on read/update, matching real GCP.
  3. **Error Alignment**: Aligned 404 message to `com.google.apps.framework.request.NotFoundException: Guardrail %s does not exist in the app.` matching real GCP response format.
  4. **Harness & Verification**: Added `CESGuardrail` to mock-supported GroupKinds in `config/tests/samples/create/harness.go`. Generated `_http_mock.log` for both `cesguardrail-minimal` and `cesguardrail-maximal`, achieving 0 discrepancies for Guardrail requests against real GCP golden logs. Re-reconciliation tests pass without ratcheting exemptions.








### [2026-11-16] Prepopulating CESTool Direct KRM Types, Identity, and Judgement Queue
- **Context**: Prepopulating direct KRM types, identity, reference, and judgement queue for `CESTool` (`ces.cnrm.cloud.google.com/v1alpha1`) using `google.cloud.ces.v1` and `--prepopulate-resource`.
- **Learnings & Decisions**:
  1. **Scaffolded Types & Identity**: Added `--resource CESTool:Tool` to `generate-types` and `generate-identity` in `apis/ces/generate.sh`. Scaffolds `CESTool` KRM types, `CESToolIdentity`, `CESToolRef`, and adds `CESTool` entries to `apis/ces/judgement_queue.yaml` (with `untriaged-bulk-generation` set to `status: open`).
  2. **Release Bundle Blocklist**: Added `CESTool` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  3. **Empty Oneof Handling**: `ServiceAgentIDTokenAuthConfig` is a parameterless message representing an empty option in the `auth_config` oneof. Added `cestools.ces.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` per KCC guidelines.
  4. **Pre-existing Hand-written References**: Pre-existing `CesToolRef` and `CesToolIdentity` stubs in `apis/ces/v1alpha1/` (added earlier for `DialogflowCompanionAgent`) are retained for backwards compatibility, resolving the temporary entries in `testdata/exceptions/naming_violations.txt`.

### [2026-11-16] Triaging Judgement Queue and Finalizing CESTool Direct Types
- **Context**: Triaging `apis/ces/judgement_queue.yaml` entries and finalizing direct KRM types and identity for `CESTool` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Decisions**:
  1. **Single Parent**: Replaced `projectRef` and `location` fields on `CESToolSpec` with required immutable `parentRef *CESAppRef` following `go/kcc-oneparent` and Case B of `PARENT_AND_IDENTITY.md`.
  2. **Converted References**:
     - `spec.agentTool.rootAgent` -> `rootAgentRef *CesAgentRef` (`apis/ces/v1alpha1`)
     - `spec.connectorTool.connection` -> `connectionRef *connectorsv1alpha1.ConnectorsConnectionRef` (`apis/connectors/v1alpha1`)
     - `spec.dataStoreTool.boostSpecs[].dataStores` -> `dataStoreRefs []discoveryenginev1alpha1.DiscoveryEngineDataStoreRef` (`apis/discoveryengine/v1alpha1`)
     - `spec.dataStoreTool.dataStoreSource.dataStore.name` -> `dataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef` (`apis/discoveryengine/v1alpha1`)
     - `spec.dataStoreTool.engineSource.dataStoreSources[].dataStore.name` -> `dataStoreRef *discoveryenginev1alpha1.DiscoveryEngineDataStoreRef` (`apis/discoveryengine/v1alpha1`)
     - `spec.dataStoreTool.engineSource.engine` -> `engineRef *discoveryenginev1alpha1.DiscoveryEngineEngineRef` (`apis/discoveryengine/v1alpha1`)
     - `spec.fileSearchTool.fileCorpus` -> `fileCorpusRef *aiplatformv1alpha1.AIPlatformRAGCorpusRef` (`apis/aiplatform/v1alpha1`)
     - `spec.mcpTool.apiAuthentication.apiKeyConfig.apiKeySecretVersion` and `spec.openAPITool.apiAuthentication.apiKeyConfig.apiKeySecretVersion` -> `apiKeySecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef` (`apis/refs/v1beta1`)
     - `spec.mcpTool.apiAuthentication.oauthConfig.clientSecretVersion` and `spec.openAPITool.apiAuthentication.oauthConfig.clientSecretVersion` -> `clientSecretVersionRef *refsv1beta1.SecretManagerSecretVersionRef` (`apis/refs/v1beta1`)
     - `spec.mcpTool.apiAuthentication.serviceAccountAuthConfig.serviceAccount` and `spec.openAPITool.apiAuthentication.serviceAccountAuthConfig.serviceAccount` -> `serviceAccountRef *refsv1beta1.IAMServiceAccountRef` (`apis/refs/v1beta1`)
     - `spec.mcpTool.serviceDirectoryConfig.service` and `spec.openAPITool.serviceDirectoryConfig.service` -> `serviceRef *servicedirectoryv1beta1.ServiceDirectoryServiceRef` (`apis/servicedirectory/v1beta1`)
  3. **Deferred and Non-GCP Fields**:
     - `spec.connectorTool.action.connectionActionID` and `spec.connectorTool.action.entityOperation.entityID`: marked `not-applicable` as they are connector-specific entity/action IDs, not GCP resources.
     - `spec.mcpTool.apiAuthentication.oauthConfig.clientID` and `spec.openAPITool.apiAuthentication.oauthConfig.clientID`: marked `not-applicable` as they are third-party OAuth provider client IDs.
     - `spec.widgetTool.dataMapping.sourceToolName`: marked `deferred` and recorded in `tests/apichecks/testdata/exceptions/refs_deferred.txt` as `dialogflow.googleapis.com/Tool` has no KCC resource yet.
  4. **Etag & Deprecated Fields**:
     - Moved `etag` from `CESToolSpec` to `CESToolObservedState` and removed the entry from `tests/apichecks/testdata/exceptions/spec_dislike_etag.txt`.
     - Removed unshared deprecated field `pythonScript` from `WidgetTool_DataMapping`.
  5. **Resolved Judgement Queue**: All 41 `CESTool` entries in `apis/ces/judgement_queue.yaml` are triaged and resolved.

### [2026-11-16] Prepopulating CESToolset Direct Types and Identity
- **Context**: Greenfield Step 1a for `CESToolset` under `ces.cnrm.cloud.google.com/v1alpha1` using `google.cloud.ces.v1` (`Toolset` proto).
- **Observations & Decisions**:
  1. **Existing External References**: `CesToolsetRef` and `CesToolsetIdentity` were previously added as handwritten external-only references in `apis/ces/v1alpha1/` (for `DialogflowCompanionAgent`), flagged in `naming_violations.txt`. With `CESToolset` introduced as a full KRM resource, `cestoolset` is now a valid kind prefix, allowing removal of the obsolete exception from `naming_violations.txt` while keeping the legacy `CesToolsetRef` intact to avoid breaking consumers.
  2. **Empty oneof variant**: Proto message `ServiceAgentIDTokenAuthConfig` in `McpToolset.ApiAuthentication` and `OpenApiToolset.ApiAuthentication` is an empty parameterless struct. `TestCRDObjectTypes` flagged `cestoolsets.ces.cnrm.cloud.google.com` as an object missing properties. Added `cestoolsets.ces.cnrm.cloud.google.com` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` following the standard `// GCP known pattern: Empty oneof option` convention.
  3. **Release Bundle Blocked Kinds**: Added `"CESToolset"` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Impact**: All unit tests in `apis/ces/...` and `tests/apichecks` pass cleanly with fully idempotent code generation.

### [2026-11-17] Direct Controller, Fuzzer, and E2E Fixtures for CESTool
- **Context**: Implementing Greenfield Step 2 direct controller, fuzzer, and recording real-GCP E2E fixtures for `CESTool` (`ces.cnrm.cloud.google.com/v1alpha1`).
- **Learnings & Service Behaviors**:
  1. **Synchronous REST CRUD**: Similar to `CESDeployment`, `CESTool` CRUD operations (`CreateTool`, `UpdateTool`, `GetTool`, `DeleteTool`) in `ces.googleapis.com/v1` are synchronous REST calls returning `*cespb.Tool` or `*emptypb.Empty` directly without LRO handling.
  2. **Tool Types**: `Tool.tool_type` is a protobuf `oneof`. `GoogleSearchTool` provides a self-contained tool configuration without external GCP resource dependencies, supporting `name`, `description`, `contextURLs`, and `preferredDomains`.
  3. **Common Tool Fields**: Top-level `executionType` (`SYNCHRONOUS`) and `toolFakeConfig` (`enableFakeMode`, `codeBlock.pythonCode`) are supported on Tool and mutable in-place via `UpdateTool` with `updateMask: googleSearchTool, toolFakeConfig`.
  4. **Audit Probe & E2E Validation**: Golden HTTP logs and audit probes were recorded against real GCP project `cnrm-barni-4` for `cestool-minimal` and `cestool-maximal`, with all 7/7 automated lifecycle checks passing cleanly.

### [2026-11-17] Triaging Judgement Queue and Finalizing Direct KRM Types for CESToolset
- **Context**: Greenfield Step 1b for `CESToolset` under `ces.cnrm.cloud.google.com/v1alpha1` using `google.cloud.ces.v1` (`Toolset` proto).
- **Observations & Decisions**:
  1. **Parent Modeling**: `Toolset` is a child resource of `CESApp` (`projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}`). Replaced `ProjectRef` and `Location` with a single immutable, required `ParentRef *CESAppRef`. `controllerbuilder generate-identity` generated identity conformance helpers parsing `parentRef`.
  2. **References in Shared Types**:
     - `spec.connectorToolset`: Inlined `ConnectorToolset` struct into `cestoolset_types.go` and converted `Connection` (`string`) to `ConnectionRef *connectorsv1alpha1.ConnectorsConnectionRef`. `controllerbuilder` automatically pruned `ConnectorToolset` from `types.generated.go`.
     - `spec.mcpToolset` & `spec.openAPIToolset`: Shared with `CESTool`, already converted to typed KCC references in `types.generated.go` during `CESTool` Step 1b (`APIKeySecretVersionRef`, `ClientSecretVersionRef`, `ServiceAccountRef`, `ServiceRef`).
  3. **Non-GCP Fields**:
     - `spec.connectorToolset.connectorActions[].connectionActionID` and `spec.connectorToolset.connectorActions[].entityOperation.entityID`: marked `not-applicable` as connector/entity IDs within third-party connections.
     - `spec.mcpToolset.apiAuthentication.oauthConfig.clientID` and `spec.openAPIToolset.apiAuthentication.oauthConfig.clientID`: marked `not-applicable` as third-party OAuth client IDs.
  4. **Etag**: Moved `etag` from `CESToolsetSpec` to `CESToolsetObservedState` and removed `cestoolsets.ces.cnrm.cloud.google.com` from `tests/apichecks/testdata/exceptions/spec_dislike_etag.txt`.
  5. **Resolved Judgement Queue**: All 18 `CESToolset` entries in `apis/ces/judgement_queue.yaml` are triaged and resolved.
- **Impact**: All unit tests in `apis/ces/...` and `tests/apichecks` pass cleanly.


