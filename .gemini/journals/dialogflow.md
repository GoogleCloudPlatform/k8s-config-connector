### [2026-06-05] Dialogflow Direct Types Scaffolding and Multi-Hierarchy Identity
- **Context**: Implementing initial direct types, CRD, and IdentityV2 for `DialogflowKnowledgeBase` under `apis/dialogflow/v1alpha1` (Issue #9287).
- **Problem**: 
  1. The issue instructions suggested using the `google.cloud.dialogflow.v1` service package; however, no `v1` of Dialogflow containing the `KnowledgeBase` resource exists in the `googleapis` proto definitions. The `KnowledgeBase` resource is located under `google.cloud.dialogflow.v2`.
  2. Dialogflow `KnowledgeBase` supports both Global and Regional resource structures (`projects/{project}/knowledgeBases/{knowledge_base}` vs. `projects/{project}/locations/{location}/knowledgeBases/{knowledge_base}`).
- **Solution**:
  1. Configured `generate.sh` and `doc.go` to use `google.cloud.dialogflow.v2` for generating the core schemas.
  2. Implemented a dual-template hierarchical Identity in `knowledgebase_identity.go` to parse, stringify, and validate both global and regional formats dynamically based on the presence of the `Spec.Location` field.
- **Impact**: The next agent building the Dialogflow direct controllers can seamlessly reuse this multi-parent identity logic without fearing validation failures or schema discrepancies.
