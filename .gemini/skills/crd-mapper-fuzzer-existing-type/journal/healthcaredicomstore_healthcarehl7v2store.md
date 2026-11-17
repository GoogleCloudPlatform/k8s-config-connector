# Journal: HealthcareDICOMStore and HealthcareHL7V2Store Transition to Direct KRM Types

## Learnings & Observations

### 1. Resources Without Proto Message in googleapis
- Like `HealthcareConsentStore`, both `HealthcareDICOMStore` and `HealthcareHL7V2Store` belong to the Google Cloud Healthcare API, which is a REST/Discovery API (`google.golang.org/api/healthcare/v1`).
- Proto definitions in `googleapis/googleapis` only contain logging events (`google/cloud/healthcare/logging/*`), with no service protos for store resources.
- As with `HealthcareConsentStore`, `controllerbuilder generate-types` and `generate-mapper` are not applicable. Direct KRM types were written directly in `apis/healthcare/v1alpha1/healthcaredicomstore_types.go` and `apis/healthcare/v1alpha1/healthcarehl7v2store_types.go`.
- `apis/healthcare/generate.sh` was updated to document that `HealthcareDICOMStore` and `HealthcareHL7V2Store` are hand-coded, while retaining the call to `dev/tasks/generate-crds`.

### 2. Strict Schema Compatibility
- The existing CRD for `HealthcareDICOMStore` defines:
  - `spec.dataset` (required string)
  - `spec.notificationConfig` (optional object with required `pubsubTopic` string)
  - `spec.resourceID` (optional string)
  - `spec.streamConfigs` (optional array of objects with required `bigqueryDestination` object with required `tableUri` string)
  - `status.conditions`
  - `status.observedGeneration` (optional integer)
  - `status.selfLink` (optional string)
  - `status.externalRef` (added per standard practice, permitted by `diff-crds --allow-status-external-ref-addition`)

- The existing CRD for `HealthcareHL7V2Store` defines:
  - `spec.dataset` (required string)
  - `spec.notificationConfig` (optional object with required `pubsubTopic` string)
  - `spec.notificationConfigs` (optional array of objects with optional `filter` string and required `pubsubTopic` string)
  - `spec.parserConfig` (optional object with optional `allowNullHeader` bool, `schema` string, `segmentTerminator` string, `version` string)
  - `spec.resourceID` (optional string)
  - `status.conditions`
  - `status.observedGeneration` (optional integer)
  - `status.selfLink` (optional string)
  - `status.externalRef` (added per standard practice, permitted by `diff-crds --allow-status-external-ref-addition`)

- Ran `dev/tasks/diff-crds --allow-status-external-ref-addition`, which yielded 0 diffs.

### 3. Pointers and Acronyms
- All optional primitive fields are typed as pointers (e.g. `*string`, `*bool`, `*int64`).
- All required primitive fields are typed as values (e.g. `Dataset string`, `PubsubTopic string`, `TableURI string`).
- Slices are not pointers (e.g. `[]DICOMStoreStreamConfig`, `[]HL7V2StoreNotificationConfigs`).
- Nested struct fields within `Spec` are defined as pointers to allow automatic serialization/mapping.

### 4. IdentityV2 & Reference Implementation
- `HealthcareDICOMStore`:
  - CAI format: `//healthcare.googleapis.com/projects/{{PROJECT_NUMBER}}/locations/{{LOCATION}}/datasets/{{DATASET}}/dicomStores/{{DICOMSTORE}}`
  - Template: `"projects/{project}/locations/{location}/datasets/{dataset}/dicomStores/{dicomStore}"`
- `HealthcareHL7V2Store`:
  - CAI format: `//healthcare.googleapis.com/projects/{{PROJECT_NUMBER}}/locations/{{LOCATION}}/datasets/{{DATASET}}/hl7V2Stores/{{HL7V2STORE}}`
  - Template: `"projects/{project}/locations/{location}/datasets/{dataset}/hl7V2Stores/{hl7V2Store}"`
- Both resources parse parent dataset information via `DatasetIdentityFormat` and resolve resource ID via `refs.GetResourceID`.
- `GetIdentity` cross-checks the parsed identity against `status.externalRef` if populated.
- `HealthcareDICOMStoreRef` and `HealthcareHL7V2StoreRef` implement `refs.Ref`, registered in `init()`, and use `refs.NormalizeWithFallback` which checks for readiness (`Ready == True`) before resolving identity.

### 5. Verification
- `apis/healthcare/v1alpha1` unit tests pass completely.
- `pkg/gcpurls` tests pass with template CAI matching.
- `tests/apichecks` pass completely.
- `dev/tasks/diff-crds --allow-status-external-ref-addition` reports empty diff.
- Updated `docs/reports/crd_report.csv` and `docs/reports/crd_report.md` with `has_go_type = True`.
