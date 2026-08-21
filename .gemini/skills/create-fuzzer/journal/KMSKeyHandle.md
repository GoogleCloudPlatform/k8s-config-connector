# Journal: KMSKeyHandle Fuzzer Integration

This journal documents the specific details and learnings from implementing and validating the round-trip KRM fuzzer for `KMSKeyHandle`.

## Context & Structure
- **Resource:** `KMSKeyHandle`
- **Location of Controller & Mappers:** `pkg/controller/direct/kms/keyhandle/keyhandle_controller.go` and `pkg/controller/direct/kms/mapper.generated.go`.
- **Location of Fuzzer:** `pkg/controller/direct/kms/keyhandle/keyhandle_fuzzer.go`
- **GCP Proto Type:** `google.cloud.kms.v1.KeyHandle` (represented in Go as `*kmspb.KeyHandle` / `*pb.KeyHandle`)

## Field Configuration in Fuzzer
- **Spec Fields:**
  - `.resource_type_selector` is registered via `f.SpecField(".resource_type_selector")`.
- **Status Fields:**
  - `.kms_key` is registered via `f.StatusField(".kms_key")`.
- **Identity Fields:**
  - `.name` is the resource name (identity) in GCP, registered via `f.Unimplemented_Identity(".name")`.

## Learnings & Observations
- **KRM-only Fields:** `KMSKeyHandleSpec` contains fields like `resourceID`, `projectRef`, and `location` that do not exist as distinct fields in `pb.KeyHandle` because they are part of the resource's KCC identity (URL/URI) used to construct the GCP resource path (`.name`), and are thus not part of the standard object payload.
- **Minimal Schema:** `pb.KeyHandle` is an extremely minimal protobuf with only three public fields (`name`, `kms_key`, and `resource_type_selector`), so no additional untriaged or unimplemented fields need to be handled.
