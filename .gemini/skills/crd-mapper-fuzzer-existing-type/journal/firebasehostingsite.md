# Journal: FirebaseHostingSite Transition to Direct KRM Types

## Learnings & Observations

### 1. Resource Without Proto Message in googleapis
- `FirebaseHostingSite` represents a Firebase Hosting site (`projects/{project}/sites/{site}`).
- The Firebase Hosting API is a REST-based Google API service (`google.golang.org/api/firebasehosting/v1beta1`) rather than a gRPC/protobuf API defined in `googleapis/googleapis`.
- Because there is no proto definition in `googleapis`, controllerbuilder's `generate-types` and `generate-mapper` tools are not used. Instead, the direct KRM types are defined directly in `apis/firebasehosting/v1alpha1/firebasehostingsite_types.go`, matching the approach taken by other non-proto resources (such as `ComputeSharedVPCServiceProject` and `KMSSecretCiphertext`).
- `apis/firebasehosting/generate.sh` is configured to run `./generate-proto.sh` and `dev/tasks/generate-crds`, standardizing CRD and deepcopy generation.

### 2. Acronyms and Field Naming in KRM Types
- The baseline CRD schema defined field names with camelCase acronyms: `spec.appId` and `status.defaultUrl`.
- Following Go naming conventions and repository acronym guidelines, Go struct fields are named `AppID` and `DefaultURL`, while setting the JSON struct tags to `json:"appId,omitempty"` and `json:"defaultUrl,omitempty"`. This satisfies both Go acronym conventions and preserves 100% strict schema compatibility with the existing CRD.

### 3. Pointers and Reference Types
- Optional scalar primitive fields (`AppID`, `ResourceID`, `DefaultURL`, `Name`, `ObservedGeneration`) are typed as Go pointers (`*string`, `*int64`).
- Required parent project reference uses `refs.ProjectRef` value type (`ProjectRef refs.ProjectRef \`json:"projectRef"\``), which matches the baseline CRD schema and enables automatic `oneOf` constraint injection via `scripts/add-validation-to-crds`.
- Status conditions use `[]v1alpha1.Condition \`json:"conditions,omitempty"\`` from `pkg/apis/k8s/v1alpha1`.

### 4. Identity & Reference Implementation
- `FirebaseHostingSiteIdentity` implements `identity.IdentityV2` using `gcpurls.Template[FirebaseHostingSiteIdentity]("firebasehosting.googleapis.com", "projects/{project}/sites/{site}")`.
- `GetIdentity` handles cross-checking `spec` identity against `status.name`.
- `FirebaseHostingSiteRef` implements `refs.Ref` with fallback normalization.
- Since Firebase Hosting URL format is not present in Cloud Asset Inventory (`cloudassetinventory_names.jsonl`), `"//firebasehosting.googleapis.com/projects/{}/sites/{}"` was added to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to ensure `TestRegisteredTemplatesMatchCAI` passes.

### 5. Verification
- `dev/tasks/diff-crds` confirmed zero schema diff against the baseline CRD.
- Ran `dev/tasks/generate-crds` and `make fmt`.
- Updated `docs/reports/crd_report.csv` and `docs/reports/crd_report.md` to reflect `has_go_type = True`.
