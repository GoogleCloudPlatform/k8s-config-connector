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

### [2026-07-06] DialogflowTool Initial Types and Identity Scaffolding
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DialogflowTool` (Issue #9290).
- **Problem**:
  - The issue description requested using service `google.cloud.dialogflow.v1` for the `DialogflowTool` resource. However, there is no `Tool` proto message under `google.cloud.dialogflow.v1`. It is defined under `google.cloud.dialogflow.cx.v3beta1`.
  - The GCP URL template given in the instructions is `projects/{project}/locations/{location}/tools/{tool}` which omits the `agents/{agent}` parent path segment. However, Dialogflow CX Tools are subresources of agents (`projects/{project}/locations/{location}/agents/{agent}/tools/{tool}`).
- **Solution**:
  - Configured `generate.sh` to run the scaffolder for `DialogflowTool:Tool` under `--service google.cloud.dialogflow.cx.v3beta1` and rename `types.generated.go` to `tool_types.generated.go`.
  - Defined `DialogflowCXAgentRef` in `apis/dialogflow/v1alpha1/dialogflowcxagent_reference.go` to act as the parent reference type for the tool.
  - Implemented the identity parser and validation in `apis/dialogflow/v1alpha1/dialogflowtool_identity.go` with the correct nested `projects/{project}/locations/{location}/agents/{agent}/tools/{tool}` template.
- **Impact**: Sets up correct types, CRD, identity, and reference mapping for DialogflowTool, providing a clean path for the subsequent direct controller implementation.
