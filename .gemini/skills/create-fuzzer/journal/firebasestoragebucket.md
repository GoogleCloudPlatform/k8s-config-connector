# Fuzzer Journal: FirebaseStorageBucket (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseStorageBucket` (`pb.Bucket`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Storage is managed via REST / OpenAPI endpoints (`firebasestorage.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase Storage is `google.golang.org/api/firebasestorage/v1beta`, where the API representation of a Bucket is `Bucket`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebasestorage/firebasestoragebucket_mappings.go`:
   - `FirebaseStorageBucketSpec_FromAPI` / `FirebaseStorageBucketSpec_ToAPI`: empty spec mapper (spec fields `projectRef` and `resourceID` are standard KCC reference and identity fields).
   - `FirebaseStorageBucketStatus_FromAPI` / `FirebaseStorageBucketStatus_ToAPI`: maps `Name`.
2. Created `firebasestoragebucket_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Status fields (`.Name`)
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
3. Registered `pkg/controller/direct/firebasestorage` in `pkg/controller/direct/register/register.go` in alphabetical order.
4. Verified that `FOCUS=FirebaseStorageBucket go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.
