# Journal: ApigeeSyncAuthorization MockGCP Implementation

This journal captures the details and findings of implementing the `ApigeeSyncAuthorization` mock in MockGCP.

## Observations

1. **Service Proto and Generation**:
   The `service.proto` for `google.cloud.apigee.v1` already generates the RPC methods `GetSyncAuthorizationOrganization` and `SetSyncAuthorizationOrganization` on `OrganizationsServer`.
   The endpoints correspond to:
   - `POST /v1/organizations/{name}:getSyncAuthorization`
   - `POST /v1/organizations/{name}:setSyncAuthorization`

2. **Mock Implementation**:
   - Implemented `GetSyncAuthorizationOrganization` and `SetSyncAuthorizationOrganization` on `organizationsServer` in `mockgcp/mockapigee/syncauthorization.go`.
   - Utilizes `s.storage` to store and retrieve `GoogleCloudApigeeV1SyncAuthorization` under key `organizations/{org}/syncAuthorization`.
   - Computes MD5 etags from proto byte serialization.
   - If not found on `Get`, returns empty identities with computed etag (matching GCP API behavior where syncAuthorization is initialized on the organization).

3. **Re-reconciliation Testing**:
   - `ApigeeSyncAuthorization` uses a POST HTTP method for reading sync authorization state (`POST /v1/organizations/{org}:getSyncAuthorization`).
   - The generic re-reconciliation checker flags any non-GET/non-gRPC method during touch as a write.
   - Added an entry in `tests/e2e/ratcheting.go` explaining that ApigeeSyncAuthorization uses POST for its read endpoint and returns false.

4. **Real GCP Recording Attempt & Fallback**:
   - Attempted recording against real GCP (`cnrm-barni-4`) using `./hack/record-gcp "fixtures/apigeesyncauthorization-basic"`.
   - Encountered 400 FAILED_PRECONDITION:
     ```
     GetSyncAuthorization is not allowed for organization with runtimeType CLOUD
     ```
   - Apigee Organizations only allow `getSyncAuthorization` and `setSyncAuthorization` when `runtimeType` is `HYBRID`.
   - Creating a `HYBRID` Apigee organization requires Apigee Hybrid entitlements/licensing not enabled on PAYG GCP projects (`runtimeType HYBRID is not supported for PAYG_V2, use CLOUD instead`).
   - Consequently, real GCP recording could not proceed.
   - As required by the guidelines, the fallback approach was used: golden logs and object definitions were generated and verified hermetically through MockGCP using `hack/compare-mock`.
