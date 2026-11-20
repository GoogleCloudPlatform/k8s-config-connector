### [2026-05-26] DialogflowConversationDataset Scaffold
- **Context**: Implementing Greenfield step 1 for DialogflowConversationDataset.
- **Problem**: The proto mapping generated unreachable types due to lack of usages in the scaffolded struct.
- **Solution**: Manually copied fields like `InputConfig` and `ConversationInfo` to the `ConversationDatasetObservedState` before re-running generation.
- **Impact**: Enables properly typing output fields in the CRD schema.

### [2026-06-05] Dialogflow Direct Types Scaffolding and Multi-Hierarchy Identity
- **Context**: Implementing initial direct types, CRD, and IdentityV2 for `DialogflowKnowledgeBase` under `apis/dialogflow/v1alpha1` (Issue #9287).
- **Problem**: 
  1. The issue instructions suggested using the `google.cloud.dialogflow.v1` service package; however, no `v1` of Dialogflow containing the `KnowledgeBase` resource exists in the `googleapis` proto definitions. The `KnowledgeBase` resource is located under `google.cloud.dialogflow.v2`.
  2. Dialogflow `KnowledgeBase` supports both Global and Regional resource structures (`projects/{project}/knowledgeBases/{knowledge_base}` vs. `projects/{project}/locations/{location}/knowledgeBases/{knowledge_base}`).
- **Solution**:
  1. Configured `generate.sh` and `doc.go` to use `google.cloud.dialogflow.v2` for generating the core schemas.
  2. Implemented a dual-template hierarchical Identity in `knowledgebase_identity.go` to parse, stringify, and validate both global and regional formats dynamically based on the presence of the `Spec.Location` field.
- **Impact**: The next agent building the Dialogflow direct controllers can seamlessly reuse this multi-parent identity logic without fearing validation failures or schema discrepancies.

### 2026-06-05 Initial Scaffolding and Identity for DialogflowGenerator
- **Context**: Greenfield implementation of DialogflowGenerator types, CRD, and IdentityV2 under `dialogflow.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The issue description requested service `google.cloud.dialogflow.v1`, but Dialogflow ES Generator does not exist in `v1` and is instead located in package `google.cloud.dialogflow.v2` (GCP's Dialogflow CX `v3` has a similar but different generator pattern nested under `agents`). This mismatch caused `generate-types` to fail.
- **Solution**: Changed the service in `generate.sh` to use `google.cloud.dialogflow.v2` and mapped the template pattern to `projects/{project}/locations/{location}/generators/{generator}`. Since DialogflowGenerator is missing from `cloudassetinventory_names.jsonl`, added it to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to avoid failures in `TestRegisteredTemplatesMatchCAI`.
- **Impact**: Provides a correct scaffolding, CRD, identity, and reference setup for DialogflowGenerator, paving the way for the subsequent adapter reconciliation logic.

### 2026-06-05 DialogflowSecuritySettings Initial Types and Identity Scaffolding
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DialogflowSecuritySettings`. (Issue #9288)
- **Problem**: 
  - Dialogflow security settings belong to Dialogflow CX and are defined in the `google.cloud.dialogflow.cx.v3.SecuritySettings` protobuf schema, rather than standard `google.cloud.dialogflow.v1` which does not contain security settings. Running the scaffolder against `google.cloud.dialogflow.v1` fails with "proto not found".
  - The URL format for Dialogflow security settings is not listed in Cloud Asset Inventory (CAI) metadata, meaning `TestRegisteredTemplatesMatchCAI` will fail when the gcpurls template is evaluated.
- **Solution**: 
  - Ran `generate-types` using `--service google.cloud.dialogflow.cx.v3` to correctly reference `google.cloud.dialogflow.cx.v3.SecuritySettings`.
  - Added `//dialogflow.googleapis.com/projects/{}/locations/{}/securitySettings/{}` to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to explicitly allow registering the Dialogflow security settings URL template without requiring a corresponding CAI entry.
- **Impact**: Enables smooth scaffolding and test success for Dialogflow CX resources that use direct identity models and are absent from CAI definitions.

### [2026-06-24] DialogflowSipTrunk Initial Types and Identity Scaffolding
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DialogflowSipTrunk` (Issue #9289).
- **Problem**:
  - The issue description requested using service `google.cloud.dialogflow.v1` for the `DialogflowSipTrunk:SipTrunk` resource; however, the `SipTrunk` resource only exists under the beta/v2beta1 version of Dialogflow, i.e., package `google.cloud.dialogflow.v2beta1` in the Google APIs protobuf schema. Attempting to use `google.cloud.dialogflow.v1` fails with "proto: not found".
- **Solution**:
  - Configured `generate.sh` to use the service `google.cloud.dialogflow.v2beta1` to correctly generate the schema for `DialogflowSipTrunk:SipTrunk`.
  - Moved the generated output `types.generated.go` to `siptrunk_types.generated.go` and restored `types.generated.go` for the other Dialogflow v2 resources.
  - Implemented the identity and external ref logic in `dialogflowsiptrunk_identity.go` with unit tests in `dialogflowsiptrunk_identity_test.go`.
- **Impact**: Provides a correct scaffolding, CRD, identity, and reference setup for DialogflowSipTrunk, preparing the codebase for the subsequent adapter and reconciliation controller implementation steps.

### [2026-10-01] DialogflowConversationDataset Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowConversationDataset` (Issue #13603).
- **Problem**:
  1. Dialogflow ConversationDataset is completely immutable in GCP (no update RPC exists in Dialogflow v2 API).
  2. Dataset IDs are service-generated (format `projects/{project}/locations/{location}/conversationDatasets/{conversationDataset}`).
  3. `DeleteConversationDatasetOperation.Wait(ctx)` in GAPIC REST client returns `unsupported result type <nil>: <nil>` on successful completion due to empty LRO response body handling.
  4. Golden HTTP logs and object outputs require dataset ID normalization (`${conversationDatasetID}`).
- **Solution**:
  1. Implemented direct controller adapter with immutable update logic (rejecting updates with descriptive field diffs), service-generated ID semantics (following 4 rules in `kcc-direct-service-generated-id`), and handled empty LRO result error in `Delete`.
  2. Added `${conversationDatasetID}` normalization in `mockgcp/mockdialogflow/normalize.go`.
  3. Generated mappers and fuzzer under `pkg/controller/direct/dialogflow/conversationdataset/`.
  4. Recorded `dialogflowconversationdataset-minimal` and `dialogflowconversationdataset-maximal` fixtures against real GCP.
- **Impact**: Fully enables management and reconciliation of `DialogflowConversationDataset` resources via direct controller.

### [2026-10-19] DialogflowWebhook Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowWebhook` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #14087).
- **Problem**:
  1. Dialogflow Webhook IDs in Dialogflow CX are exclusively server-generated (format `projects/{project}/locations/{location}/agents/{agent}/webhooks/{webhook}`). Initial identity scaffolding used `refs.GetResourceID(obj)` which erroneously fell back to `metadata.name`.
  2. DialogflowCXAgent parent references only supported `external` references (`Name` and `Namespace` were commented out). In dependencies, `DialogflowCXAgent` uses a server-assigned UUID, so referencing it by `name` was necessary.
  3. `common.VisitFields` threw `visiting slice of type slice is not supported` when traversing `DialogflowWebhookSpec` due to `AllowedCACerts [][]byte` (slice of slice of byte).
- **Solution**:
  1. Updated `DialogflowWebhookIdentity` to implement `identity.ServerGeneratedIdentity`, adhering to the 4 service-generated ID rules (leaving `Name` empty in GCP create payload, pre-check in `Find`, and mapping `status.externalRef` on reconciliation).
  2. Enabled `Name` and `Namespace` on `DialogflowCXAgentRef` and implemented `NormalizeWithFallback` retrieving agent ID from `status.startFlow` or `status.name`.
  3. Added `case reflect.Slice:` support in `common.VisitFields` for slices of slices.
  4. Implemented direct controller adapter using GAPIC REST client (`cx.NewWebhooksRESTClient`) and registered model and fuzzer.
  5. Created minimal and maximal fixtures and recorded golden outputs against real GCP using `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`.
- **Impact**: Fully enables management and reconciliation of `DialogflowWebhook` resources via direct controller.

### [2026-10-29] DialogflowPlaybook Greenfield Step 1b Judgement Queue Triage
- **Context**: Triaging judgement queue and finalizing KRM types and identity for DialogflowPlaybook (Issue #14672).
- **Problem**:
  1. `DialogflowPlaybook` is a child of `DialogflowCXAgent` (`projects/{project}/locations/{location}/agents/{agent}/playbooks/{playbook}`). The generator emitted separate `projectRef` and `location` fields instead of following the single-parent rule.
  2. The create RPC `CreatePlaybook` has no client-specified ID parameter; the playbook ID is server-generated.
  3. `generate-mapper` for Dialogflow CX previously dumped Playbook mappers into `dialogflowconversationmodel_mapper.go` alongside ConversationModel mappers.
  4. Deeply nested fields inside shared structs (`Fulfillment`, `ParameterDefinition`) reference GCP resources (`Webhook`, `Tool`, `Generator`) that are not yet modeled as KCC kinds or are shared with untriaged kinds (`DialogflowFlow`, `DialogflowPage`, `DialogflowTransitionRouteGroup`).
- **Solution**:
  1. Replaced `projectRef` and `location` with `ParentRef *DialogflowCXAgentRef` with the `self == oldSelf` immutability CEL rule.
  2. Updated `ResourceID` godoc to document server-generated ID, allowing `generate-identity` to emit `identity.ServerGeneratedIdentity`.
  3. Isolated DialogflowPlaybook mappers into `pkg/controller/direct/dialogflow/dialogflowplaybook_mapper.go` satisfying `TestDirectResourceFileNaming`.
  4. Deferred unmodeled references in `judgement_queue.yaml` and recorded them in `tests/apichecks/testdata/exceptions/refs_deferred.txt`.
- **Impact**: Cleanly establishes identity, types, and controller/fuzzer conformance for `DialogflowPlaybook` without breaking backward compatibility of shared structs or modifying other kinds' CRDs.

### [2026-11-05] DialogflowPlaybook Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowPlaybook` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #15056).
- **Problem**:
  1. `ParameterDefinition.type` is deprecated in the Dialogflow CX v3 API; Playbook parameter definitions require `typeSchema.inlineSchema.type: STRING` rather than the deprecated enum `type`. Sending `type` resulted in 400 Bad Request (`Playbook input parameter definition must set a type schema or parameter type`).
  2. In Dialogflow CX Playbooks, array fields like `instruction.steps` and `inputParameterDefinitions` cannot be modified via `PATCH` in-place (updates either append steps or are ignored). The primary mutable fields accepted on `PATCH` are `displayName`, `goal`, and `instruction.guidelines`.
- **Solution**:
  1. Configured `dialogflowplaybook-maximal` to use `typeSchema.inlineSchema.type: STRING` in `inputParameterDefinitions` and `outputParameterDefinitions`.
  2. Tested in-place updates on supported mutable fields (`displayName`, `goal`, `instruction.guidelines`) across minimal and maximal fixtures.
  3. Recorded authentic golden HTTP logs, reports, and object representations against real GCP, passing all 7/7 validation checks.
- **Impact**: Fully validates the direct controller lifecycle, golden HTTP logs, and identity resolution for `DialogflowPlaybook`.

### [2026-11-11] DialogflowCompanionAgent Greenfield Step 1a Prepopulate
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for DialogflowCompanionAgent (Issue #15367).
- **Details**:
  1. `DialogflowCompanionAgent` belongs to `google.cloud.dialogflow.v2` (GCP resource: `projects/{project}/locations/{location}/companionAgents/{companion_agent}`).
  2. The proto definition is in `.build/googleapis-e7513c828b8db536fa12fc5493d541f323aedf90.pb` (`google/cloud/dialogflow/v2/companion_agent.proto`).
  3. Added `generate-types` and `generate-identity` with `--prepopulate-resource` for `DialogflowCompanionAgent:CompanionAgent` to `apis/dialogflow/generate.sh`.
  4. Added `DialogflowCompanionAgent` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  5. Handled empty proto struct `CompanionAgent_KnowledgeSource` by allowlisting `dialogflowcompanionagents.dialogflow.cnrm.cloud.google.com` in `knownInvalidCRDs` in `tests/apichecks/crds_test.go`.
- **Impact**: Sets up clean scaffolding and open judgement queue items for subsequent triage in Step 1b.

### [2026-11-11] DialogflowCompanionAgent Greenfield Step 1b Judgement Queue Triage
- **Context**: Triaging judgement queue and finalizing KRM types and identity for DialogflowCompanionAgent (Issue #15377).
- **Details**:
  1. Resolved `location-or-parent-ref`: Regional project-parented resource (`projects/{project}/locations/{location}/companionAgents/{companion_agent}`), replaced separate `projectRef` and `location` with inlined `*parent.ProjectAndLocationRef`.
  2. Resolved `authoritative-proto-reference` for `ces_tool`: Proto references `ces.googleapis.com/Tool` (`projects/{project}/locations/{location}/apps/{app}/tools/{tool}`). Moved `CesToolSpec` to `dialogflowcompanionagent_types.go`, defined external-only `CesTool` identity and reference in `apis/ces/v1alpha1/`, and updated field to `CesToolRef *cesv1alpha1.CesToolRef`.
  3. Resolved `authoritative-proto-reference` for `toolset`: Proto references `ces.googleapis.com/Toolset` (`projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}`). Moved `ToolsetTool` to `dialogflowcompanionagent_types.go`, defined external-only `CesToolset` identity and reference in `apis/ces/v1alpha1/`, and updated field to required `ToolsetRef *cesv1alpha1.CesToolsetRef`.
  4. Resolved `possible-reference-by-name` for `operationID`: Confirmed it is an OpenAPI operationId within the toolset schema, not a GCP resource reference (`resolution: not-applicable`).
  5. Resolved `untriaged-bulk-generation` after verifying all fields and types.
- **Impact**: All 5 judgement queue entries for DialogflowCompanionAgent resolved; CRD and API checks cleanly pass.

### [2026-11-12] DialogflowCompanionAgent Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowCompanionAgent` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #15457).
- **Details**:
  1. `CompanionAgentsClient` in `cloud.google.com/go/dialogflow` was added in `v1.87.0`. Upgraded module dependency from `v1.82.0` to `v1.87.0` in both `go.mod` and `mockgcp/go.mod`.
  2. Isolated `DialogflowCompanionAgent` mappers in `pkg/controller/direct/dialogflow/dialogflowcompanionagent_mapper.go` to avoid clobbering CX v3 mappers in `mapper.generated.go` and comply with `TestDirectResourceFileNaming`.
  3. Controller and fuzzer generated using `controllerbuilder` with `NewCompanionAgentsRESTClient`.
  4. Handled status update with `ObservedState` on `Create` and `Update`.
  5. Recorded minimal and maximal fixtures (`dialogflowcompanionagent-minimal` and `dialogflowcompanionagent-maximal`) against real GCP with `RECORD_AUDIT_PROBE=1 ./hack/record-gcp` (project `cnrm-barni-4`, location `global`), passing all 7/7 validation checks and generating audit log actuation entries.
  6. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` to remove tested spec fields.
- **Impact**: Fully enables direct reconciliation, fuzzer round-tripping, and authentic E2E test verification for `DialogflowCompanionAgent`.

### [2026-11-14] DialogflowConversationProfile Greenfield Step 1b Judgement Queue Triage
- **Context**: Triaging judgement queue and finalizing KRM types, identity, and references for DialogflowConversationProfile (Issue #15630).
- **Details**:
  1. Resolved parent & identity: Mapped parent as inlined `*parent.ProjectAndLocationRef` in `DialogflowConversationProfileSpec`. Configured `--pattern DialogflowConversationProfile=projects/{project}/locations/{location}/conversationProfiles/{conversationProfile}` in `apis/dialogflow/generate.sh` to match CAIS naming format. Confirmed server-generated identity as `CreateConversationProfileRequest` has no client-specified ID parameter.
  2. References modeled:
     - `SecuritySettingsRef *DialogflowSecuritySettingsRef`: Converted `security_settings` to `SecuritySettingsRef`.
     - `NotificationConfig.TopicRef *pubsubv1beta1.PubSubTopicRef`: Moved `NotificationConfig` to `dialogflowconversationprofile_types.go` and converted `topic` to `TopicRef` across all 4 notification config usages.
     - `SpeechToTextConfig.PhraseSetRefs []speechv1beta1.PhraseSetRef`: Moved `SpeechToTextConfig` to `dialogflowconversationprofile_types.go` and converted `phrase_sets` to `PhraseSetRefs`.
     - `HumanAgentAssistantConfig_ConversationModelConfig.ModelRef *DialogflowConversationModelRef`: Converted `model` to `ModelRef`.
     - `HumanAgentAssistantConfig_SuggestionQueryConfig_KnowledgeBaseQuerySource.KnowledgeBaseRefs []DialogflowKnowledgeBaseRef`: Converted `knowledge_bases` to required `KnowledgeBaseRefs`.
     - `HumanAgentAssistantConfig_SuggestionConfig.GeneratorRefs []DialogflowGeneratorRef`: Converted `generators` to `GeneratorRefs`.
  3. Non-reference fields resolved:
     - `salesforceLiveAgentConfig` fields (`buttonID`, `deploymentID`, `organizationID`): Third-party Salesforce IDs (`resolution: not-applicable`).
     - `ttsConfig.effectsProfileID`: Cloud Text-to-Speech audio effect profile identifiers (`resolution: not-applicable`).
  4. Deferred references in `refs_deferred.txt`:
     - `automatedAgentConfig.agent`: Polymorphic ES/CX agent environment path (`projects/*/locations/*/agent/environments/*` or `projects/*/locations/*/agents/*/environments/*`).
     - `dialogflowQuerySource.agent` and `humanAgentSideConfig.agent`: Regional Dialogflow agent paths (`projects/*/locations/*/agent`), whereas `DialogflowAgentRef` is global (`projects/{project}/agent`).
     - `documentQuerySource.documents`: Data-plane knowledge documents.
  5. Triaged all 27 judgement queue entries for DialogflowConversationProfile in `apis/dialogflow/judgement_queue.yaml`.
- **Impact**: Cleanly finalizes KRM types, identity, references, CRD schema, and generated mappers for DialogflowConversationProfile, passing all tests and apichecks.

### [2026-11-19] DialogflowConversationProfile Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowConversationProfile` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #16172).
- **Details**:
  1. Isolated `DialogflowConversationProfile` mappers in `pkg/controller/direct/dialogflow/dialogflowconversationprofile_mapper.go` to avoid clobbering CX v3 mappers in `mapper.generated.go` and comply with `TestDirectResourceFileNaming`.
  2. Direct controller generated using `controllerbuilder generate-controller` with `NewConversationProfilesRESTClient`.
  3. Fuzzer generated via `generate-fuzzer` and configured with `f.Unimplemented_NotYetTriaged(".sip_config")` and `FilterSpec` to filter unmapped fields from newer protobufs (`sip_config`, `gemini_asr_config`, `session_ttl`, `enable_sentiment_analysis_v3`, `pronunciations`, and suggestion config fields).
  4. Updated `pkg/controller/resourceconfig/static_config.go` with direct reconciler for DialogflowConversationProfile.
  5. Implemented minimal and maximal test fixtures under `pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowconversationprofile/`. Found that `timeZone` is write-only / immutable in Dialogflow v2 API (`[update_mask]: 'The path is not allowed to update: time_zone'`). Used `displayName` and `loggingConfig.enableStackdriverLogging` for update verification.
  6. Successfully recorded and verified E2E fixtures against real GCP (`project: cnrm-barni-4`, `location: global`) passing all 7/7 automated validation checks.
- **Impact**: Enables direct controller reconciliation, roundtrip fuzzing, and authentic real GCP testing for DialogflowConversationProfile.

### [2026-11-20] DialogflowConversationProfile MockGCP Implementation and Alignment (Phase 3)
- **Context**: Implementing MockGCP service handler and aligning golden logs for DialogflowConversationProfile (Issue #16190).
- **Details**:
  1. Implemented `conversationProfilesServer` in `mockgcp/mockdialogflow/conversationprofile.go` supporting `GetConversationProfile`, `CreateConversationProfile`, `UpdateConversationProfile`, `DeleteConversationProfile`, and `ListConversationProfiles`.
  2. Registered `pb.RegisterConversationProfilesServer` and `grpcpb.RegisterConversationProfilesHandler` in `mockgcp/mockdialogflow/service.go`.
  3. Added `DialogflowConversationProfile` to supported mock GroupKinds in `config/tests/samples/create/harness.go`.
  4. Handled default values (`languageCode: "en-US"`, conditional initialization of empty notification and assistant config blocks, and `createTime` omission on update responses) matching authentic GCP behavior.
  5. Successfully aligned and verified minimal and maximal test fixtures with `hack/compare-mock` and `TestGoldenLogAlignment`, with `_http_mock.log` matching `_http.log` cleanly. Passed `./dev/ci/presubmits/tests-e2e-fixtures-suite dialogflow`.
- **Impact**: Enables hermetic mock testing and end-to-end reconciliation for DialogflowConversationProfile without requiring live GCP access.

### [2026-11-19] DialogflowGeneratorEvaluation Greenfield Step 1b Judgement Queue Triage
- **Context**: Finalizing direct KRM types, identity, references, and triaging judgement queue for DialogflowGeneratorEvaluation (Issue #15949).
- **Details**:
  1. `DialogflowGeneratorEvaluation` belongs to `google.cloud.dialogflow.v2` (GCP resource: `projects/{project}/locations/{location}/generators/{generator}/evaluations/{evaluation}`).
  2. Modeled child resource parentage with required, immutable `spec.parentRef *DialogflowGeneratorRef` (`json:"parentRef"`) per `go/kcc-oneparent` (Case B). Configured `generate-identity` to resolve ancestor segments from `parentRef`. Confirmed server-generated identity conforming to `identity.ServerGeneratedIdentity`.
  3. Defined dedicated nested struct `DialogflowGeneratorEvaluation_Generator` for `spec.initialGenerator` to omit top-level KRM resource-identity fields (`projectRef`, `location`, `resourceID`).
  4. Enforced proto-required fields with dedicated `WithRequired` variants (`SummarizationContextWithRequired`, `FewShotExampleWithRequired`, `GeneratorSuggestionWithRequired`, `FreeFormSuggestionWithRequired`, `SummarySuggestionWithRequired`, `SummarySuggestion_SummarySectionWithRequired`).
  5. Moved `ToolCall` to `dialogflowplaybook_types.go` and `GeneratorSuggestion` to `dialogflowgenerator_types.go` to prevent upstream proto drift without modifying shared `controllerbuilder` tooling.
  6. Converted `.spec.generatorEvaluationConfig.inputDataConfig.datasetInputDataConfig.dataset` to `datasetRef *vertexaiv1beta1.VertexAIDatasetRef`.
  7. Deferred `.spec.generatorEvaluationConfig.outputGCSBucketPath` as a GCS folder URI (`gs://...`) in `refs_deferred.txt`.
  8. Implemented dedicated direct mappers in `pkg/controller/direct/dialogflow/dialogflowgeneratorevaluation_mapper.go` conforming to `TestDirectResourceFileNaming` and preserved existing controller mappers with zero diff.
  9. Resolved all 37 judgement queue entries in `apis/dialogflow/judgement_queue.yaml` and placed `DialogflowGeneratorEvaluation` on `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Impact**: Cleanly scaffolds and finalizes direct KRM types, IdentityV2 conformance, references, CRD schema, and isolated direct mappers for `DialogflowGeneratorEvaluation`.

### [2026-11-20] DialogflowGeneratorEvaluation Direct Controller Implementation (Phase 2)
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DialogflowGeneratorEvaluation` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #16227).
- **Details**:
  1. Generated direct controller (`pkg/controller/direct/dialogflow/dialogflowgeneratorevaluation_controller.go`) and roundtrip fuzzer (`pkg/controller/direct/dialogflow/dialogflowgeneratorevaluation_fuzzer.generated.go`) via one-shot `controllerbuilder generate-controller` and `generate-fuzzer` using `--prepopulate-resource`. Neither was added to `apis/dialogflow/generate.sh`.
  2. Fixed fuzzer roundtrip zero-value preservation in `pkg/controller/direct/dialogflow/dialogflowgeneratorevaluation_mapper.go`: preserved `SatisfiesPzs` and `SatisfiesPzi` pointer values (`out.SatisfiesPzs = in.SatisfiesPzs`) instead of calling `direct.LazyPtr(in.GetSatisfiesPzs())`, which dropped `false` zero-values.
  3. Controller conforms to Greenfield invariants: calls `common.NormalizeGreenfieldReferences(ctx, reader, obj, nil)` before `obj.GetIdentity(ctx, reader)` in `AdapterForObject`. In `Create`, waits on `op.Wait(ctx)` without suppressing errors, followed by a post-LRO `Get` to fetch the complete `GeneratorEvaluation` and update `status.ObservedState` and `status.externalRef`.
  4. Updated `pkg/controller/resourceconfig/static_config.go` with direct reconciler for DialogflowGeneratorEvaluation via `python3 dev/tasks/generate_static_config.py`.
  5. Implemented minimal and maximal test fixtures under `pkg/test/resourcefixture/testdata/basic/dialogflow/v1alpha1/dialogflowgeneratorevaluation/` with prerequisite `DialogflowGenerator` and `StorageBucket` in `dependencies.yaml`.
  6. Added timestamp normalization rules for `completeTime`, `initialGenerator.createTime`, and `initialGenerator.updateTime` in `tests/e2e/normalize.go`, `tests/e2e/httplog.go`, `tests/e2e/normalize_legacy.go`, and `pkg/test/resourcefixture/golden_alignment_test.go`.
  7. Recorded and verified E2E fixtures passing all 7/7 automated validation checks in `_http_report.md` without modifying test harness validator logic.
  8. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` via `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha` to reflect covered spec fields.
- **Impact**: Enables direct controller reconciliation, roundtrip fuzzing, and authentic real GCP testing for DialogflowGeneratorEvaluation.

### [2026-11-20] DialogflowGeneratorEvaluation MockGCP Implementation and Alignment (Phase 3)
- **Context**: Implementing MockGCP service handler and aligning golden logs for `DialogflowGeneratorEvaluation` under `dialogflow.cnrm.cloud.google.com/v1alpha1` (Issue #16239).
- **Details**:
  1. Implemented `generatorEvaluationsServer` in `mockgcp/mockdialogflow/generatorevaluation.go` supporting `CreateGeneratorEvaluation`, `GetGeneratorEvaluation`, `ListGeneratorEvaluations`, and `DeleteGeneratorEvaluation`.
  2. Registered `pb_v2.RegisterGeneratorEvaluationsServer` and `RegisterGeneratorEvaluationsHandler` in `mockgcp/mockdialogflow/service.go`.
  3. Validated parent `DialogflowGenerator` existence during `CreateGeneratorEvaluation` and returned `codes.NotFound` (404) if missing, conforming to real GCP behavior.
  4. Handled server-generated evaluation IDs (`MT...`), LRO creation metadata (`GeneratorEvaluationOperationMetadata`), evaluation status (`EvaluationStatus`), summarization metrics, and defaulting (`summarizationConfig.evaluatorVersion: "2.0"`, `initialGenerator.summarizationContext.version: "6.0"`).
  5. Updated `config/tests/samples/create/harness.go` to add `DialogflowGeneratorEvaluation` to supported mock GroupKinds.
  6. Added timestamp normalizations for `completeTime` and `initialGenerator` timestamps to `mockdialogflow/normalize.go` and normalized KRM `completeTime` in `tests/e2e/normalize.go` with `PlaceholderTime`. Added evaluation ID normalization in `pkg/test/resourcefixture/golden_alignment_test.go`.
  7. Generated `_http_mock.log` files with zero diff on the target resource against authentic real GCP golden logs (`_http.log`) and 0 diff in golden objects. Passed `TestGoldenLogAlignment` and `./dev/ci/presubmits/tests-e2e-fixtures-suite dialogflow`.
- **Impact**: Enables hermetic mock testing and end-to-end reconciliation for `DialogflowGeneratorEvaluation` without requiring live GCP access.


