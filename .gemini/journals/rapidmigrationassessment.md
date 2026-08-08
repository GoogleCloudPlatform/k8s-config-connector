# Rapid Migration Assessment Journal

## Quirks & Learnings

### 1. Proto Type Namespace Resolution Issue (v1main vs v1)
The GCP `RapidMigrationAssessmentCollector` API returns an LRO response/metadata containing `"type.googleapis.com/google.cloud.rapidmigrationassessment.v1main.Collector"` instead of `"google.cloud.rapidmigrationassessment.v1.Collector"`.
The canonical Go `google.golang.org/protobuf` registry does not have the `"v1main"` version of the message registered. As a result, calling `op.Wait(ctx)` during Create, Update, or Delete operations always fails with:
`proto: unable to resolve "type.googleapis.com/google.cloud.rapidmigrationassessment.v1main.Collector"`

**Resolution:**
We implemented an `isProtoResolutionError(err)` helper to identify and ignore these unmarshalling failures. Since we always perform a standard `GetCollector` GET operation immediately after `Wait` completes to fetch the fully-populated resource, ignoring this client-side unmarshalling error is completely safe and robust.

### 2. Service Account Reference Requirement
Creating a collector on the GCP backend requires a valid service account. If omitted or empty, the API returns:
`rpc error: code = InvalidArgument desc = invalid SA`

**Resolution:**
We explicitly populated the `serviceAccountRef.external` field in our test fixtures using the KCC test project's existing compute/tester service account:
`overseer-kcc-tester@${projectId}.iam.gserviceaccount.com`

### 3. Defaulting Required Integer Fields
Fields like `expectedAssetCount` and `collectionDays` are required by the GCP API to be positive integers. In KRM, these are optional pointers, but because they map to raw integer fields in proto, omitting them in KRM sends a `0` value, which is rejected by the GCP API.

**Resolution:**
We added a `populateProtoDefaults` helper inside the direct controller to inject sensible defaults (`expectedAssetCount = 100` and `collectionDays = 30`) before sending the create or update payload.

### 4. Immutable Labels
The `labels` field on the Rapid Migration Assessment Collector is completely immutable once created. Attempting to update labels via the patch request mask returns:
`rpc error: code = InvalidArgument desc = Invalid update mask 'labels'. Field cannot be modified in location 5.`

**Resolution:**
We left `labels` unchanged in the `update.yaml` of the maximal test fixture to ensure successful in-place updates of other mutable fields.
