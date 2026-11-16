# Fuzzer Journal: FirebaseDatabaseInstance (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseDatabaseInstance` (`pb.DatabaseInstance`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Realtime Database Management API is managed via REST / OpenAPI endpoints (`firebasedatabase.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase Realtime Database is `google.golang.org/api/firebasedatabase/v1beta`, where the API representation of an instance is `DatabaseInstance`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebasedatabase/firebasedatabaseinstance_mappings.go`:
   - `FirebaseDatabaseInstanceSpec_FromAPI` / `FirebaseDatabaseInstanceSpec_ToAPI`: maps `Type`.
   - `FirebaseDatabaseInstanceStatus_FromAPI` / `FirebaseDatabaseInstanceStatus_ToAPI`: maps `DatabaseUrl`, `Name`, `State`.
2. Created `firebasedatabaseinstance_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Spec fields (`.Type`)
   - Configured Status fields (`.DatabaseUrl`, `.Name`, `.State`)
   - Configured Identity fields (`.Project`)
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
3. Registered `pkg/controller/direct/firebasedatabase` in `pkg/controller/direct/register/register.go` in alphabetical order.
4. Verified that `FOCUS=FirebaseDatabaseInstance go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.
