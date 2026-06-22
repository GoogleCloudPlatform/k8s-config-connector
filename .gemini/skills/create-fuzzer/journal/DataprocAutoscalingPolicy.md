# Journal: DataprocAutoscalingPolicy Fuzzer Integration

This journal documents the specific details and learnings from implementing and validating the round-trip KRM fuzzer for `DataprocAutoscalingPolicy`.

## Context & Structure
- **Resource:** `DataprocAutoscalingPolicy`
- **Location of Controller & Mappers:** `pkg/controller/direct/dataproc/autoscalingpolicy_mappings.go` and `pkg/controller/direct/dataproc/mapper.generated.go`.
- **Location of Fuzzer:** `pkg/controller/direct/dataproc/dataprocautoscalingpolicy_fuzzer.go`
- **GCP Proto Type:** `google.cloud.dataproc.v1.AutoscalingPolicy` (represented in Go as `*pb.AutoscalingPolicy`)

## Field Configuration in Fuzzer
- **Spec Fields:**
  - `.basic_algorithm` is registered via `f.SpecField(".basic_algorithm")`.
  - `.worker_config` is registered via `f.SpecField(".worker_config")`.
  - `.secondary_worker_config` is registered via `f.SpecField(".secondary_worker_config")`.
- **Identity Fields:**
  - `.id` and `.name` are the resource identity fields, registered via `f.Unimplemented_Identity(".id")` and `f.Unimplemented_Identity(".name")`.
- **Unimplemented / Not Yet Triaged Fields:**
  - `.labels` is registered via `f.Unimplemented_NotYetTriaged(".labels")`.
  - `.cluster_type` is registered via `f.Unimplemented_NotYetTriaged(".cluster_type")`.

## Learnings & Observations
- **KRM-only Fields:** Fields like `location`, `projectRef`, and `resourceID` do not exist as standard spec fields in GCP `pb.AutoscalingPolicy` because they are part of the resource's KCC identity (URL/URI).
- **Compilation performance:** Running `go test` on `./pkg/fuzztesting/fuzztests/` directly is normally slow because Go runs vetting on all dependent direct controllers. Using the `-vet=off` flag enables instant compilation of the test binary (`go test -vet=off -c`), which can then be run as an executable for extremely fast feedback loop.
