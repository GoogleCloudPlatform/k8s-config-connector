# DocumentAIProcessorDefaultVersion Fuzzer Journal

## Overview
Implemented round-trip KRM fuzzer for `DocumentAIProcessorDefaultVersion` (`documentai.cnrm.cloud.google.com/v1alpha1`) targeting `google.cloud.documentai.v1.SetDefaultProcessorVersionRequest`.

## Learnings & Observations

### 1. RPC Request Message & Mapper Generation
- `DocumentAIProcessorDefaultVersion` maps to the protobuf RPC request `google.cloud.documentai.v1.SetDefaultProcessorVersionRequest` rather than a standard resource message.
- `controllerbuilder generate-mapper` intentionally skips proto messages ending in `*Request` (defined in `dev/tools/controllerbuilder/pkg/commands/generatemapper/generatemappercommand.go`), meaning mappers for request-based resources are not auto-generated into `mapper.generated.go`.
- Furthermore, field names between KRM and Proto do not match directly:
  - KRM `spec.version` maps to proto `.default_processor_version`.
  - KRM `spec.resourceID` represents the processor identity.
- Therefore, handwritten `DocumentAIProcessorDefaultVersionSpec_v1alpha1_FromProto` and `DocumentAIProcessorDefaultVersionSpec_v1alpha1_ToProto` functions were implemented in a dedicated mapper file `documentaiprocessordefaultversion_mapper.go` rather than inside the fuzzer file.

### 2. Spec-Only Fuzzing with `NewKRMTypedSpecFuzzer`
- `DocumentAIProcessorDefaultVersion` is a brownfield Terraform-migrated resource whose KRM Status contains only `conditions` and `observedGeneration`, with no `observedState` sub-struct.
- We used `fuzztesting.NewKRMTypedSpecFuzzer` and registered it using `fuzztesting.RegisterKRMSpecFuzzer` in `init()`, which skips status fuzzing while rigorously validating spec round-tripping.

### 3. Field Categorization
- `.default_processor_version`: Marked as a spec field with `f.SpecField(".default_processor_version")`.
- `.processor`: The canonical GCP identity for the processor (`projects/{project}/locations/{location}/processors/{processor}`), handled via KRM resource identity rather than a direct spec field. Marked as `f.Unimplemented_Identity(".processor")`.
- Both fields in `pb.SetDefaultProcessorVersionRequest` are fully accounted for, ensuring a 100% round-trip validation.
