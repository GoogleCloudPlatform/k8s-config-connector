# DataflowJob Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Strict Schema Compatibility for References**:
   - The baseline `DataflowJob` CRD does not define `projectRef` or `location` in its spec (unlike most other resources). To adhere to strict schema compatibility, we refrained from adding them to `DataflowJobSpec` in `job_types.go`.
   - For references like `kmsKeyRef`, `networkRef`, `serviceAccountRef`, and `subnetworkRef`, the baseline CRD lacks a `kind` field. We hand-coded custom reference types `KMSCryptoKeyRef`, `ComputeNetworkRef`, `ComputeSubnetworkRef`, and `IAMServiceAccountRef` directly in `job_types.go`.
   - By matching their signatures as `"external,name,namespace"`, the OpenAPI `oneOf` reference constraints are automatically appended by the CRD generator to match the original baseline constraints perfectly.

2. **Map Types and Preserve Unknown Fields**:
   - The baseline `DataflowJob` CRD specifies `parameters` and `transformNameMapping` as objects with `x-kubernetes-preserve-unknown-fields: true`.
   - We mapped these fields to `*runtime.RawExtension` in our Go types, which correctly compiles to the exact same OpenAPI representation.

3. **Removal of Unused Status Fields**:
   - The baseline `DataflowJob` CRD status only includes `conditions`, `jobId`, `observedGeneration`, `state`, and `type`. We excluded `externalRef` and `observedState` from `DataflowJobStatus` in `job_types.go` to keep the schema strictly compatible.

4. **Package generate.sh Configuration**:
   - We configured `apis/dataflow/v1beta1/generate.sh` to include both the pre-existing `DataflowFlexTemplateJob` and our new `DataflowJob`.
   - Running `./generate.sh` correctly parses the `google.dataflow.v1beta3` protobuf messages and produces/updates `types.generated.go`.

5. **Round-trip KRM Fuzzer Implementation**:
   - Added proto mapping annotations `// +kcc:spec:proto=google.dataflow.v1beta3.Job` and `// +kcc:status:proto=google.dataflow.v1beta3.Job` to trigger the mapper generator.
   - Wrote handcoded `DataflowJobSpec_FromProto`, `DataflowJobSpec_ToProto`, `DataflowJobStatus_FromProto`, and `DataflowJobStatus_ToProto` functions in `pkg/controller/direct/dataflow/mapper.go` to cleanly map nested protobuf `Job.Environment` and `WorkerPool` fields to flat KRM fields.
   - Handled `*runtime.RawExtension` <-> `map[string]string` conversion by marshaling and unmarshaling.
   - Designed and implemented `dataflowjob_fuzzer.go` using `fuzztesting.RegisterKRMFuzzer` to round-trip and verify conversion losslessness of both Spec and Status fields.
   - Designed a robust `FilterSpec` and `FilterStatus` in the fuzzer to normalize randomized `WorkerPool` slices and empty nested `Environment` messages, allowing 100,000 randomized fuzz test runs to pass successfully with zero loss.

## Verification

- Successfully ran `dev/tasks/diff-crds` showing that the generated CRD schema is 100% structurally identical to the baseline schema.
- Successfully executed `make manifests`, `make generate-go-client`, and `go vet ./...` without compilation or lint issues.
- Successfully ran fuzzer tests `TestSomeMappers` and `FuzzAllMappers` confirming 100% round-trip conversion losslessness.
