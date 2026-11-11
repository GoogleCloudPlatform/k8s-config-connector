# Fuzzer Journal: FirebaseWebApp (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseWebApp` (`pb.WebApp`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Management API is managed via REST / OpenAPI endpoints (`firebase.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase is `google.golang.org/api/firebase/v1beta1`, where the API representation of a Web app is `WebApp`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebase/firebasewebapp_mappings.go`:
   - `FirebaseWebAppSpec_FromAPI` / `FirebaseWebAppSpec_ToAPI`: maps `ApiKeyId`, `DisplayName`.
   - `FirebaseWebAppStatus_FromAPI` / `FirebaseWebAppStatus_ToAPI`: maps `AppId`, `AppUrls`, `Name`.
2. Created `firebasewebapp_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Spec fields (`.ApiKeyId`, `.DisplayName`)
   - Configured Status fields (`.AppId`, `.AppUrls`, `.Name`)
   - Configured Identity fields (`.ProjectId`)
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
   - Marked untriaged/unimplemented API fields (`.Etag`, `.ExpireTime`, `.State`, `.WebId`) via `f.Unimplemented_NotYetTriaged`
3. Registered `pkg/controller/direct/firebase` in `pkg/controller/direct/register/register.go`.
4. Verified that `FOCUS=FirebaseWebApp go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.
