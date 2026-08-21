# Journal: KMSCryptoKey Fuzzer Integration

This journal documents the specific details and learnings from implementing and validating the round-trip KRM fuzzer for `KMSCryptoKey`.

## Context & Structure
- **Resource:** `KMSCryptoKey`
- **Location of Controller & Mappers:** `pkg/controller/direct/kms/mapper.generated.go` and `pkg/controller/direct/kms/mappers.go`.
- **Location of Fuzzer:** `pkg/controller/direct/kms/cryptokey_fuzzer.go`
- **GCP Proto Type:** `google.cloud.kms.v1.CryptoKey` (represented in Go as `*kmspb.CryptoKey` / `*pb.CryptoKey`)

## Field Configuration in Fuzzer
- **Spec Fields:**
  - `.purpose` is registered via `f.SpecField(".purpose")`.
  - `.rotation_period` is registered via `f.SpecField(".rotation_period")`. Note that in the Go struct `pb.CryptoKey`, this resides under `RotationSchedule` (a Go structure for the protobuf `oneof rotation_schedule`), but at the proto descriptor level, it is visited as `.rotation_period`.
  - `.version_template` (and its nested `.version_template.algorithm`, `.version_template.protection_level`) are registered as SpecFields.
  - `.import_only` is registered via `f.SpecField(".import_only")`.
  - `.destroy_scheduled_duration` is registered via `f.SpecField(".destroy_scheduled_duration")`.
- **Identity Fields:**
  - `.name` is the resource name (identity) in GCP, registered via `f.Unimplemented_Identity(".name")`.
- **Unimplemented/Not Yet Triaged Fields:**
  - `.primary` (the current primary version, which is output-only).
  - `.create_time` (the creation timestamp, which is output-only).
  - `.next_rotation_time` (the next rotation timestamp, which is output-only).
  - `.labels` (not yet supported by KMSCryptoKey in direct).
  - `.crypto_key_backend` (not yet supported by KMSCryptoKey in direct).
  - `.key_access_justifications_policy` (not yet supported by KMSCryptoKey in direct).

## Learnings & Observations
- **KRM-only Fields:** `KMSCryptoKeySpec` contains fields like `skipInitialVersionCreation` and `keyRingRef` that do not exist as distinct fields in `pb.CryptoKey` (parent information is embedded entirely within `.name`). Since the fuzzer operates by fuzzing the GCP proto, mapping to KRM, and mapping back to GCP proto, these KRM-only fields are naturally not fuzzed on the proto side, preventing any round-trip issues.
