# DialogflowCXFlow Fuzzer Implementation Journal

## Observations & Learnings

### 1. Proto and KRM Mapping Details
- `DialogflowCXFlow` maps to protobuf message `google.cloud.dialogflow.cx.v3.Flow` (`cloud.google.com/go/dialogflow/cx/apiv3/cxpb.Flow`).
- KRM group is `dialogflowcx.cnrm.cloud.google.com/v1alpha1`.
- The KRM Spec contains `parent`, `languageCode`, and `resourceID`, which are not present on `pb.Flow` (`parent` and `languageCode` are request parameters in `CreateFlowRequest`/`UpdateFlowRequest`, and `resourceID` is KCC's acquisition identifier).
- The KRM Status does not contain an `observedState` in its CRD schema (only standard `conditions`, `name`, `observedGeneration`, `externalRef`), though `DialogflowCXFlowObservedState` is defined as an empty struct in KRM types.
- Following the convention in `dialogflowcxwebhook` and `create-fuzzer` skill instructions for resources without observed state fields, `fuzztesting.RegisterKRMSpecFuzzer` was used with `fuzztesting.NewKRMTypedSpecFuzzer` to register a Spec-only fuzzer.

### 2. Complex Fields and JSON Serialization
- Several fields in KRM are JSON-encoded strings mapping to protobuf structs or messages:
  - `SetParameterActions[].value`: maps between `*string` and `*structpb.Value`.
  - `ConditionalCases[].cases`: maps between `*string` and `[]*Fulfillment_ConditionalCases_Case`.
  - `Messages[].payload`: maps between `*string` and `*structpb.Struct`.
  - `Messages[].conversation_success.metadata`: maps between `*string` and `*structpb.Struct`.
  - `Messages[].live_agent_handoff.metadata`: maps between `*string` and `*structpb.Struct`.
- `protojson` was used for marshaling and unmarshaling these fields.

### 3. Protobuf Oneof Invariants & Normalization (FilterSpec)
- For `structpb.Value`, protobuf random fuzzing can produce instances where `Kind == nil` (none of the oneof fields is set). `protojson.Marshal` rejects these with `proto: google.protobuf.Value: none of the oneof fields is set`.
- To preserve round-trip equivalence and prevent marshaling errors, `normalizeValue` sets unset oneof kinds to `structpb.Value_NullValue`.
- Similarly, `normalizeStruct` and `normalizeConditionalCasesCase` ensure nested `structpb.Value` and `structpb.Struct` instances are normalized before marshaling.
- `f.FilterSpec` runs these normalizers on `p1` (the incoming fuzzed proto) so the initial proto matches the reconstructed proto.

### 4. Unimplemented Fields
- Proto fields not represented in KRM Spec:
  - Identity: `.name` marked via `f.Unimplemented_Identity`.
  - Flow-level: `.advanced_settings`, `.knowledge_connector_settings`, `.input_parameter_definitions`, `.output_parameter_definitions`, `.multi_language_settings`, `.locked`.
  - Route / Handler-level: `.transition_routes[].description`, `.event_handlers[].target_playbook`.
  - Fulfillment-level: `.trigger_fulfillment.advanced_settings`, `.trigger_fulfillment.enable_generative_fallback`, `.trigger_fulfillment.generators`.
  - ResponseMessage-level: `.messages[].response_type`, `.messages[].end_interaction`, `.messages[].mixed_audio`, `.messages[].knowledge_info_card`, `.messages[].tool_call`.
- All marked via `f.Unimplemented_NotYetTriaged`.

### 5. Verification
- `FOCUS=DialogflowCXFlow go test -count=10 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passed.
- `go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFuzzersAttributedToService` passed.
- `go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestListTypesWithoutFuzzers` confirms `DialogflowCXFlow` is no longer on the missing fuzzers list.
