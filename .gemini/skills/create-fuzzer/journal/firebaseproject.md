# Fuzzer Journal: FirebaseProject (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseProject` (`api.FirebaseProject`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Management API is managed via REST / OpenAPI endpoints (`firebase.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase is `google.golang.org/api/firebase/v1beta1`, where the API representation of a project is `FirebaseProject`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebase/firebaseproject_mappings.go`:
   - `FirebaseProjectSpec_FromAPI` / `FirebaseProjectSpec_ToAPI`: The KRM Spec contains only `ProjectRef` (parent GCP Project reference) and `ResourceID` (resource identity), with no direct API Spec fields. The mappers return empty structs.
   - `FirebaseProjectStatus_FromAPI` / `FirebaseProjectStatus_ToAPI`:
     - Maps `DisplayName` <-> `DisplayName` (using `direct.LazyPtr` and `direct.ValueOf`).
     - Maps `ProjectNumber` (`*string`) <-> `ProjectNumber` (`int64`) using `strconv.FormatInt` and `strconv.ParseInt`.
2. Created `firebaseproject_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Status fields (`.DisplayName`, `.ProjectNumber`)
   - Configured Identity fields (`.Name`, `.ProjectId`)
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
   - Marked untriaged/unimplemented API fields (`.Annotations`, `.Etag`, `.Resources`, `.State`) via `f.Unimplemented_NotYetTriaged`
3. Registered `pkg/controller/direct/firebase` in `pkg/controller/direct/register/register.go` in alphabetical order.
4. Verified that `FOCUS=FirebaseProject go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.

### 3. Golden CAIS Identities Alignment
Registering `pkg/controller/direct/firebase` into `pkg/controller/direct/register/register.go` exposes `FirebaseProject` to `kccscheme`, which activates identity resolution in the CAIS powertool tests (`TestGoldenIdentitiesYamlFiles`). Updated the autogen test fixture `_identities.yaml` with the resolved CAIS URL (`//firebase.googleapis.com/projects/${projectId}`).
