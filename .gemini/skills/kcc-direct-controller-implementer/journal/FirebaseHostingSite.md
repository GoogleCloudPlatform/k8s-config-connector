# FirebaseHostingSite Journal

### 2026-10-14 FirebaseHostingSite Direct Controller Implementation
- **Context**: Brownfield migration of `FirebaseHostingSite` under `firebasehosting.cnrm.cloud.google.com/v1alpha1` to the Direct controller model.
- **Discovery Client Usage**: Firebase Hosting has no proto definition in googleapis; it uses the Google Cloud discovery client `google.golang.org/api/firebasehosting/v1beta1`. Client is instantiated via `m.config.RESTClientOptions()` with `api.NewService(ctx, opts...)`.
- **Operations & Mutability**:
  - `Create`: synchronous call `Projects.Sites.Create(parent, desired).SiteId(siteId)`.
  - `Update`: the only mutable field is `AppId`, updated via `Projects.Sites.Patch(name, desired).UpdateMask("appId")`. When `desired.AppId == ""`, `ForceSendFields: []string{"AppId"}` ensures clearing the appId.
  - `Delete`: `Projects.Sites.Delete(name)`. Handled with `direct.IsNotFound(err)` for idempotency.
  - `Find`: `Projects.Sites.Get(name)`.
- **Naming Constraints & Real GCP vs MockGCP**:
  - Firebase Hosting site IDs have a strict 3-30 character length requirement (`^[a-z0-9][a-z0-9-]{1,28}[a-z0-9]$`) and require the project to be a Firebase-enabled project (`projects.addFirebase`).
  - Pre-existing autogen mock fixture used `resourceID: firebasehostingsite${uniqueId}` which worked under MockGCP because MockGCP did not enforce the length limit.
- **Migration & Golden Testing**:
  - Dual-controller support enabled in `pkg/controller/resourceconfig/static_config.go`.
  - Golden files and migration logs (`_http_migration_phase1-4`, `_migration_diffs.json`, `_final_object.diff`) verified and passing with 0 re-reconciliation diffs.
