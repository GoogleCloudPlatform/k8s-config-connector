# Journal: PubSubSchema Fuzzer Implementation

## Learnings & Patterns

### 1. Spec-Only Fuzzing with `NewKRMTypedSpecFuzzer`
Since `PubSubSchema` does not have an `ObservedState` in its `PubSubSchemaStatus` KRM representation, we implemented a spec-only fuzzer using `fuzztesting.NewKRMTypedSpecFuzzer`. This takes only the Spec mapping functions `PubSubSchemaSpec_FromProto` and `PubSubSchemaSpec_ToProto` and uses `NoStatus` as a placeholder for the missing Status mapping functions.

### 2. Matching Fields & Unimplemented/Identity Fields
We analyzed the fields in `pb.Schema` (representing `google.pubsub.v1.Schema`) and correctly categorized them:
- `.name`: Marked as an identity field using `f.Unimplemented_Identity(".name")`.
- `.type`: Fuzzed under Spec using `f.SpecField(".type")`.
- `.definition`: Fuzzed under Spec using `f.SpecField(".definition")`.
- `.revision_id`: Marked as a status field using `f.StatusField(".revision_id")`.
- `.revision_create_time`: Marked as a status field using `f.StatusField(".revision_create_time")`.

This allows the fuzzer to cleanly verify the round-trip mappings between KRM types and protobuf representations without errors.
