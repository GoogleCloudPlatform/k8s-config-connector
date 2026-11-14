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

### [2026-11-05] DialogflowTool Greenfield Types and Identity Prepopulation
- **Context**: Prepopulating initial direct KRM types, identity, and `judgement_queue.yaml` for `DialogflowTool` under `apis/dialogflow/v1alpha1` (Issue #14945).
- **Problem**: 
  - `Tool_DataStoreTool_FallbackPrompt` in `Tool` proto is an empty struct message, causing `TestCRDObjectTypes` in `apichecks` to fail with "object at v1alpha1.spec.dataStoreSpec.fallbackPrompt is missing properties, additionalProperties, or x-kubernetes-preserve-unknown-fields".
- **Solution**:
  - Configured `apis/dialogflow/generate.sh` to include `DialogflowTool:Tool` under `google.cloud.dialogflow.cx.v3`.
  - Added `"dialogflowtools.dialogflow.cnrm.cloud.google.com": true` with comment `// GCP known pattern: Empty oneof option (spec.dataStoreSpec.fallbackPrompt)` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go`.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` via `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks -run TestCRDFieldPresenceInTestsForAlpha`.
- **Impact**: Provides deterministic type generation, IdentityV2, reference definitions, CRD, and judgement queue entry for `DialogflowTool`, completing Greenfield Step 1a.

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
