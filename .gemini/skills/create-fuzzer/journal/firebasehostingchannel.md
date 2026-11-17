# Fuzzer Journal: FirebaseHostingChannel (NoProto KRM Fuzzer)

This journal captures learnings from implementing the round-trip KRM fuzzer for `FirebaseHostingChannel` (`pb.Channel`).

## Observations and Learnings

### 1. OpenAPI / REST-based Service Client (`google.golang.org/api`)
Firebase Hosting is managed via REST / OpenAPI endpoints (`firebasehosting.googleapis.com`), and does not have protobuf definitions under `googleapis`. The official Google Cloud Go client for Firebase Hosting is `google.golang.org/api/firebasehosting/v1beta1`, where the API representation of a Channel is `Channel`.

### 2. Implementation of NoProto KRM Fuzzer
Following the `create-fuzzer` guidelines for No-Proto (OpenAPI) resources:
1. Implemented mapping functions in `pkg/controller/direct/firebasehosting/firebasehostingchannel_mappings.go`:
   - `FirebaseHostingChannelSpec_FromAPI` / `FirebaseHostingChannelSpec_ToAPI`: maps `ExpireTime`, `RetainedReleaseCount`, `TTL` <-> `Ttl`.
   - `FirebaseHostingChannelStatus_FromAPI` / `FirebaseHostingChannelStatus_ToAPI`: maps `Name`.
2. Created `firebasehostingchannel_fuzzer.go` registering the fuzzer via `fuzztesting.RegisterKRMFuzzer_NoProto`:
   - Configured Spec fields (`.ExpireTime`, `.RetainedReleaseCount`, `.Ttl`)
   - Configured Status fields (`.Name`)
   - Marked untriaged/unimplemented API fields (`.CreateTime`, `.Labels`, `.Release`, `.UpdateTime`, `.Url`) via `f.Unimplemented_NotYetTriaged`
   - Configured JSON bookkeeping fields (`.ForceSendFields`, `.NullFields`, `.ServerResponse`)
3. Verified that `FOCUS=FirebaseHostingChannel go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` passes.
