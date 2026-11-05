### [2026-10-23] ApigeeSyncAuthorization Direct Controller Implementation
- **Context**: Migrating `ApigeeSyncAuthorization` from legacy Terraform controller to direct controller (issue #14363).
- **Observations & Architecture**:
  1. `ApigeeSyncAuthorization` manages the synchronization authorization identities on an Apigee Organization (`organizations/{org}/syncAuthorization`).
  2. The Google Cloud Apigee API does not use standard REST GET/PUT endpoints or standard CRUD. Instead, it exposes custom verbs `POST organizations/{organization}:getSyncAuthorization` and `POST organizations/{organization}:setSyncAuthorization`.
  3. `GetSyncAuthorization` returns 200 OK with an initial `etag` even before any identities have been set, rather than returning a 404 Not Found error.
  4. GCP Apigee API does not have a delete endpoint for SyncAuthorization. In legacy Terraform, deleting `google_apigee_sync_authorization` only removes it from state without modifying the backend. The direct controller `Delete` method similarly treats deletion as a no-op returning `(true, nil)`.
  5. The `spec.identities` slice contains service account email strings. To prevent unnecessary API calls and flapping diffs, `equalIdentities` sorts copies of both desired and actual identities before comparison.
