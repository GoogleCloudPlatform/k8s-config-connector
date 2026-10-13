# Journal: FirebaseHostingSite MockGCP Implementation

This journal captures the details and findings of implementing the `FirebaseHostingSite` mock in MockGCP.

## Observations

1. **Service Proto and Generation**:
   The `firebasehosting` service proto is generated from the discovery document (`firebasehosting-api.json` v1beta1) using `mockgcp/tools/gapic`. The gRPC services generated include `ProjectsSitesServer` which serves CRUD endpoints for `projects/*/sites/*`:
   - `GetProjectsSite` (`GET /v1beta1/{name=projects/*/sites/*}`)
   - `CreateProjectsSite` (`POST /v1beta1/{parent=projects/*}/sites?siteId={site_id}`)
   - `DeleteProjectsSite` (`DELETE /v1beta1/{name=projects/*/sites/*}`)
   - `PatchProjectsSite` (`PATCH /v1beta1/{name=projects/*/sites/*}?updateMask={updateMask}`)

2. **Field Semantics and Defaulting**:
   - `name`: Formatted as `projects/<projectId>/sites/<siteId>`.
   - `defaultUrl`: Defaults to `https://<siteId>.web.app`.
   - `type`: Defaults to `USER_SITE` (since the default project site is `DEFAULT_SITE` provisioned automatically by Firebase).

3. **GCP Real Recording Blocker & Fallback**:
   When attempting to record against real GCP (`hack/record-gcp`), the active service account `cnrm-barni-1.svc.id.goog` lacked permissions to enable the API (`firebasehosting.googleapis.com`) and could not query project details via Cloud Resource Manager API (`Error 403: The caller does not have permission`).
   Consequently, real GCP recording could not proceed. As required by the guidelines, the mock implementation and golden files were generated and validated through MockGCP using `hack/compare-mock`.
