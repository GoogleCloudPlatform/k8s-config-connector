# Journal: FirebaseProject Transition to Direct KRM Types

## Summary
- Resource: `FirebaseProject`
- API Group: `firebase.cnrm.cloud.google.com`
- API Version: `v1alpha1`
- Underlying Service: Firebase Management API (`firebase.googleapis.com`)

## Observations & Implementation Decisions

1. **Absence of Proto Definition in googleapis**:
   - Similar to `FirebaseHostingSite`, `FirebaseProject` is part of the Firebase Management API, which does not have gRPC proto definitions in `googleapis/googleapis`. The service is defined via Google API Discovery / `google-api-go-client` (`google.golang.org/api/firebase/v1beta1`).
   - Consequently, `controllerbuilder`'s `generate-types` and `generate-mapper` are not applicable.
   - The direct KRM types and reference types were hand-coded directly in `apis/firebase/v1alpha1/`, matching the approach taken by `FirebaseHostingSite` (`apis/firebasehosting/`).

2. **generate.sh Script**:
   - Created `apis/firebase/generate.sh` following the consolidated `generate.sh` pattern.
   - Sourced `goimports.sh` and executed `dev/tasks/generate-crds` and goimports on the direct controller package when available.

3. **Strict Schema Compatibility**:
   - `Spec`:
     - `ProjectRef` (`refs.ProjectRef` required): The parent project reference matches the baseline CRD schema and automatically gets `oneOf` constraints injected by `scripts/add-validation-to-crds`.
     - `ResourceID` (`*string` optional): Matches the existing CRD definition.
   - `Status`:
     - `Conditions` (`[]v1alpha1.Condition`)
     - `DisplayName` (`*string` optional)
     - `ObservedGeneration` (`*int64` optional)
     - `ProjectNumber` (`*string` optional)
   - Verified with `dev/tasks/diff-crds --base master`, which returned completely empty (0 diff).

4. **Identity & Reference**:
   - `FirebaseProjectIdentity`: Formatted as `//firebase.googleapis.com/projects/{project}` using `gcpurls.Template[FirebaseProjectIdentity]("firebase.googleapis.com", "projects/{project}")`.
   - `FirebaseProjectRef`: Implemented `refs.Ref` interface and registered via `refs.Register`.
   - Unit tests created in `firebaseproject_identity_test.go` verifying URL formatting, relative path parsing, and external reference validation.
