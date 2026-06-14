# Dialogflow Service Journal

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
