# Journal: FirebaseAndroidApp, FirebaseWebApp, and FirebaseDatabaseInstance Transition to Direct KRM Types

## Summary
Defined direct KRM types for:
- `FirebaseAndroidApp` in `apis/firebase/v1alpha1`
- `FirebaseWebApp` in `apis/firebase/v1alpha1`
- `FirebaseDatabaseInstance` in `apis/firebasedatabase/v1alpha1`

And set up `generate.sh` in both `apis/firebase` and `apis/firebasedatabase`.

## Observations & Implementation Details

1. **Groups and Versions**:
   - `FirebaseAndroidApp` and `FirebaseWebApp` belong to API group `firebase.cnrm.cloud.google.com` (version `v1alpha1`).
   - `FirebaseDatabaseInstance` belongs to API group `firebasedatabase.cnrm.cloud.google.com` (version `v1alpha1`).
   - `apis/firebase` already existed with `FirebaseProject`, so `generate.sh` was updated / verified, and types were added alongside `firebaseproject_types.go`.
   - `apis/firebasedatabase/v1alpha1` was created from scratch with `doc.go`, `groupversion_info.go`, `firebasedatabaseinstance_types.go`, `firebasedatabaseinstance_identity.go`, `firebasedatabaseinstance_reference.go`, and `apis/firebasedatabase/generate.sh`.

2. **Schema Parity**:
   - For `FirebaseAndroidApp` and `FirebaseDatabaseInstance`, `spec.projectRef` was in the baseline CRD, so `refs.ProjectRef` was used.
   - For `FirebaseWebApp`, `spec.project` (a plain string pointer) was in the baseline CRD rather than `projectRef`. Strict schema parity was preserved by keeping `Project *string `json:"project,omitempty"`` without adding `projectRef`.
   - `ExternalRef *string `json:"externalRef,omitempty"`` was added to each resource's `Status` struct.
   - Kubebuilder tags `shortName=...` and labels `cnrm.cloud.google.com/stability-level=alpha`, `system=true`, `tf2crd=true`, and `managed-by-kcc=true` were included to match existing CRD metadata and avoid diffs.
   - `dev/tasks/diff-crds --allow-status-external-ref-addition` produced zero diffs.

3. **Identity & References**:
   - `FirebaseAndroidApp`: `projects/{project}/androidApps/{app}`
   - `FirebaseWebApp`: `projects/{project}/webApps/{app}`
   - `FirebaseDatabaseInstance`: `projects/{project}/locations/{location}/instances/{instance}`
   - Implemented `identity.IdentityV2` and `identity.Resource` on each type, and `refs.Ref` for `FirebaseDatabaseInstanceRef`.
