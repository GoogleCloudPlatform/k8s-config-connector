# Fuzzer Journal: FirebaseAndroidApp (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseAndroidApp` (`pb.AndroidApp`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Management API is managed via REST / OpenAPI endpoints (`firebase.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase is `google.golang.org/api/firebase/v1beta1`, where the API representation of an Android app is `AndroidApp`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebase/firebaseandroidapp_mappings.go`:
   - `FirebaseAndroidAppSpec_FromAPI` / `FirebaseAndroidAppSpec_ToAPI`: maps `ApiKeyId`, `DisplayName`, `PackageName`, `Sha1Hashes`, `Sha256Hashes`.
   - `FirebaseAndroidAppStatus_FromAPI` / `FirebaseAndroidAppStatus_ToAPI`: maps `AppId`, `Etag`, `Name`.
2. Created `firebaseandroidapp_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Spec fields (`.ApiKeyId`, `.DisplayName`, `.PackageName`, `.Sha1Hashes`, `.Sha256Hashes`)
   - Configured Status fields (`.AppId`, `.Etag`, `.Name`)
   - Configured Identity fields (`.ProjectId`)
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
   - Marked untriaged/unimplemented API fields (`.ExpireTime`, `.State`) via `f.Unimplemented_NotYetTriaged`
3. Registered `pkg/controller/direct/firebase` in `pkg/controller/direct/register/register.go`.
4. Verified that `FOCUS=FirebaseAndroidApp go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.
