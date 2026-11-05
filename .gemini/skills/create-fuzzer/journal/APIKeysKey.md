# Journal: APIKeysKey Fuzzer Implementation

This journal documents the specific details and learnings from implementing and validating the round-trip KRM fuzzer for `APIKeysKey`.

## Context & Structure
- **Resource Kind:** `APIKeysKey`
- **Package:** `pkg/controller/direct/apikeys`
- **Fuzzer File:** `pkg/controller/direct/apikeys/apikeyskey_fuzzer.go`
- **GCP Proto Message:** `google.api.apikeys.v2.Key` (`cloud.google.com/go/apikeys/apiv2/apikeyspb`)
- **KRM API Group/Version:** `apikeys.cnrm.cloud.google.com/v1alpha1`

## Field Configuration in Fuzzer
- **Spec Fields:**
  - `displayName` maps to `.display_name`.
  - `restrictions` maps to `.restrictions`.
  - Nested fields in `.restrictions`:
    - `browserKeyRestrictions` -> `.restrictions.browser_key_restrictions.allowed_referrers`
    - `serverKeyRestrictions` -> `.restrictions.server_key_restrictions.allowed_ips`
    - `androidKeyRestrictions` -> `.restrictions.android_key_restrictions.allowed_applications`
    - `iosKeyRestrictions` -> `.restrictions.ios_key_restrictions.allowed_bundle_ids`
    - `apiTargets` -> `.restrictions.api_targets`
- **ObservedState / Status Fields:**
  - `uid` maps to `.uid`.
- **Identity Fields:**
  - `name` represents the GCP resource name (`projects/<project>/locations/global/keys/<key_id>`) derived from `projectRef` and `resourceID`/`metadata.name`, marked via `f.Unimplemented_Identity(".name")`.
- **Unimplemented / Output-only Fields:**
  - `.key_string`: Marked via `f.Unimplemented_NotYetTriaged(".key_string")` (output-only sensitive value).
  - `.create_time`: Marked via `f.Unimplemented_NotYetTriaged(".create_time")`.
  - `.update_time`: Marked via `f.Unimplemented_NotYetTriaged(".update_time")`.
  - `.delete_time`: Marked via `f.Unimplemented_NotYetTriaged(".delete_time")`.
  - `.annotations`: Marked via `f.Unimplemented_LabelsAnnotations(".annotations")`.
  - `.etag`: Marked via `f.Unimplemented_Etag()`.

## Learnings & Observations
- **Protobuf Oneof in Nested Struct:** In `pb.Restrictions`, the `client_restrictions` field is a oneof (`browser_key_restrictions`, `server_key_restrictions`, `android_key_restrictions`, `ios_key_restrictions`). In KRM, these are separate nullable pointer fields under `Restrictions`. When protobuf random fuzzing populates the message, at most one of the oneof fields is set. Round-tripping via `Restrictions_FromProto` and `Restrictions_ToProto` preserves whichever oneof was selected without collisions or overwrites, so no custom `FilterSpec` is required.
- **Pre-existing Package Registration:** `pkg/controller/direct/apikeys` was already imported in `pkg/controller/direct/register/register.go`, so adding the fuzzer in `pkg/controller/direct/apikeys/apikeyskey_fuzzer.go` immediately wired it into `fuzztesting` without requiring changes to `register.go`.
