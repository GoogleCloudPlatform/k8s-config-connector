# Journal: FirebaseProject MockGCP Implementation

This journal captures the details and findings of implementing the `FirebaseProject` mock in MockGCP.

## Observations

1. **Service Proto and Generation**:
   The `firebase` service proto is generated from the discovery document (`firebase-api.json` v1beta1) using `mockgcp/tools/gapic`.
   - The Discovery Doc includes a `deprecated: true` property on some resources, which required adding `Deprecated bool` support to `openapi.Resource`.
   - The method `AddFirebaseProject` had a path pattern `{project=projects/*}:addFirebase` which caused `protoc-gen-grpc-gateway` to fail because `project` conflicted with a proto message type in `AddFirebaseProjectRequest`. A replacement in `apply-proto-patches.sh` changed this to `{name=projects/*}:addFirebase`.

2. **Field Semantics and Defaulting**:
   - `name`: Formatted as `projects/<projectId>`.
   - `projectId`: Equal to project ID.
   - `projectNumber`: Equal to project number (int64 formatted as string).
   - `displayName`: Equal to project ID by default.
   - `state`: Default is `ACTIVE`.
   - `etag`: Base64 encoded hash string (`abcdef0123A=`).
   - `resources`: Contains `hostingSite` set to project ID by default.

3. **GCP Real Recording**:
   - `hack/record-gcp` was successfully run against real GCP using project `cnrm-barni-4`.
   - The test fixture `firebaseprojectautogen` previously had an unreferenced `dependencies.yaml` specifying a `Project` under an organization that lacked permissions. Since `FirebaseProject` in the fixture only references `projectRef.external: ${projectId}`, removing the unnecessary `dependencies.yaml` allowed the test to run directly against the test project `cnrm-barni-4`.
   - Golden HTTP logs and golden object YAML were recorded and committed in baseline commit.
   - Re-reconciliation test succeeded after removing `FirebaseProject` from the ratcheting exclusions in `tests/e2e/ratcheting.go`.
