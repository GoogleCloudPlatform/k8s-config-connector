# DialogflowAgent Fuzzer Journal

## Observations & Implementation Details

1. **Resource & Proto Mapping**:
   - `DialogflowAgent` maps to protobuf message `google.cloud.dialogflow.v2.Agent` (`cloud.google.com/go/dialogflow/apiv2/dialogflowpb.Agent`).
   - The proto message has 12 fields total:
     - `.parent`: The GCP project resource path (`projects/<Project ID>`), marked with `f.Unimplemented_Identity(".parent")`.
     - 11 Spec fields:
       - `.api_version`: Enum `Agent_ApiVersion`, mapped to/from `spec.apiVersion` via `direct.Enum_FromProto` / `direct.Enum_ToProto`.
       - `.avatar_uri`: Mapped to/from `spec.avatarUri`.
       - `.classification_threshold`: Mapped to/from `spec.classificationThreshold`.
       - `.default_language_code`: Mapped to/from `spec.defaultLanguageCode`.
       - `.description`: Mapped to/from `spec.description`.
       - `.display_name`: Mapped to/from `spec.displayName`.
       - `.enable_logging`: Mapped to/from `spec.enableLogging`.
       - `.match_mode`: Enum `Agent_MatchMode`, mapped to/from `spec.matchMode` via `direct.Enum_FromProto` / `direct.Enum_ToProto`.
       - `.supported_language_codes`: Mapped to/from `spec.supportedLanguageCodes`.
       - `.tier`: Enum `Agent_Tier`, mapped to/from `spec.tier` via `direct.Enum_FromProto` / `direct.Enum_ToProto`.
       - `.time_zone`: Mapped to/from `spec.timeZone`.
   - `spec.resourceID` is a KRM-only field representing the resource ID / name and has no direct proto body counterpart.

2. **Status and Spec-Only Fuzzing**:
   - `DialogflowAgentStatus` contains standard KCC status fields (`conditions`, `observedGeneration`, `externalRef`) and a legacy `avatarUriBackend` field (which is an output-only duplicate of `avatar_uri`), without an `ObservedState` sub-struct.
   - Because `avatar_uri` is already mapped in Spec, and Status has no distinct observed state fields, `DialogflowAgent` is configured as a spec-only fuzzer using `fuzztesting.NewKRMTypedSpecFuzzer` and registered with `fuzztesting.RegisterKRMSpecFuzzer`.

3. **Mapper & Fuzzer Location**:
   - Implemented mappers in `pkg/controller/direct/dialogflow/dialogflowagent_mapper.go`.
   - Implemented and registered fuzzer in `pkg/controller/direct/dialogflow/dialogflowagent_fuzzer.go` in package `dialogflow`.
   - Package `dialogflow` is already registered in `pkg/controller/direct/register/register.go`.

4. **Verification**:
   - Verified round-trip fuzz testing using `FOCUS=DialogflowAgent go test -count=20 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`.
   - Verified `TestFuzzersAttributedToService` in `./pkg/fuzztesting/fuzztests/`.
