# Journal: DialogflowEntityType Fuzzer Implementation

## Learnings & Patterns

### 1. Spec-Only Fuzzing for Resources without ObservedState
`DialogflowEntityType` does not have an `ObservedState` struct in its `DialogflowEntityTypeStatus` KRM representation. As outlined in `create-fuzzer` SKILL.md, we use `fuzztesting.NewKRMTypedSpecFuzzer` and register it with `fuzztesting.RegisterKRMSpecFuzzer`. This avoids registering a status fuzzing function that would otherwise be skipped during tests.

### 2. Service Subpackages vs Root Package in Dialogflow
The `dialogflow` service directory (`pkg/controller/direct/dialogflow`) contains both direct controller implementations for Dialogflow CX (such as Webhook, Playbook) and Dialogflow v2 (such as ConversationModel), while some v2 resources are separated into dedicated subpackages (e.g. `generator`, `conversationdataset`). Since `DialogflowEntityType`'s fuzzer and mapper are placed in `pkg/controller/direct/dialogflow/`, the package is already registered in `pkg/controller/direct/register/register.go`.

### 3. Field Categorization
We mapped the fields of `google.cloud.dialogflow.v2.EntityType`:
- `.display_name`: Spec field mapped to `spec.displayName`.
- `.enable_fuzzy_extraction`: Spec field mapped to `spec.enableFuzzyExtraction`.
- `.entities`: Repeated message field mapped to `spec.entities`.
- `.entities[].value`: Leaf string field within entity entry.
- `.entities[].synonyms`: Leaf repeated string field within entity entry.
- `.kind`: Spec enum field mapped to `spec.kind`.
- `.name`: Identity field marked with `f.Unimplemented_Identity(".name")`.
- `.auto_expansion_mode`: GCP proto enum field not exposed in KRM Spec/Status, marked with `f.Unimplemented_NotYetTriaged(".auto_expansion_mode")`.
