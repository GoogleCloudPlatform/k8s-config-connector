# Fuzzer Journal: FirebaseHostingSite (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseHostingSite` (`pb.Site`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Hosting is managed via REST / OpenAPI endpoints (`firebasehosting.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase Hosting is `google.golang.org/api/firebasehosting/v1beta1`, where the API representation of a Hosting site is `Site`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions:
   - `FirebaseHostingSiteSpec_FromAPI` / `FirebaseHostingSiteSpec_ToAPI`: maps `AppID` <-> `AppId`.
   - `FirebaseHostingSiteStatus_FromAPI` / `FirebaseHostingSiteStatus_ToAPI`: maps `DefaultURL` <-> `DefaultUrl` and `Name` <-> `Name`.
2. Created `firebasehostingsite_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Spec fields (`.AppId`)
   - Configured Status fields (`.DefaultUrl`, `.Name`)
   - Marked untriaged/unimplemented API fields (`.Labels`, `.Type`) via `f.Unimplemented_NotYetTriaged`
3. Registered `pkg/controller/direct/firebasehosting` in `pkg/controller/direct/register/register.go`.
4. Verified that `FOCUS=FirebaseHostingSite go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.

### 3. Golden CAIS Identities Alignment
Registering `pkg/controller/direct/firebasehosting` into `pkg/controller/direct/register/register.go` exposes `FirebaseHostingSite` to `kccscheme`, which activates identity resolution in the CAIS powertool tests (`TestGoldenIdentitiesYamlFiles`). Updated the autogen test fixture `firebasehostingsiteautogen` to reference `projectRef.external: ${projectId}` and updated `_identities.yaml` with the resolved CAIS URL.
