### 2026-06-05 Implement Direct Types, CRD, and IdentityV2 for DataLabelingDataset
- **Context**: Implementing the initial KRM types, CRD, and IdentityV2 for `DataLabelingDataset` (Issue #9268).
- **Problem**: The issue description mentioned using `google.cloud.datalabeling.v1` as the service name, but only `v1beta1` exists in the `googleapis` proto repository.
- **Solution**: Used `google.cloud.datalabeling.v1beta1` in the `generate.sh` configuration to successfully compile descriptor files and scaffold KRM Go types.
- **Impact**: The correct service version must be `v1beta1` for all subsequent reconciliation and mapper controller development.
- **Location**: Note that the `DataLabelingDataset` resource does not have a location/region field, as the Data Labeling service organizes datasets globally/directly under a project (`projects/{project}/datasets/{dataset}`).

### [2026-06-05] Greenfield Type Generation and Identity
- **Context**: Implementing direct types, CRD, and IdentityV2 for `DataLabelingAnnotationSpecSet` under the `datalabeling.cnrm.cloud.google.com/v1alpha1` group.
- **Problem**:
  1. The task description requested `google.cloud.datalabeling.v1` as the proto service name, but the canonical Google APIs repository only includes the `v1beta1` package (`google.cloud.datalabeling.v1beta1`).
  2. The generated spec code had a `Location string` field, but the identity URL structure is project-scoped without regional/location variables (`projects/{project}/annotationSpecSets/{annotation_spec_set}`).
- **Solution**:
  1. Updated the service name in the generation script `generate.sh` to use `google.cloud.datalabeling.v1beta1`.
  2. Stripped out the `Location` field from `DataLabelingAnnotationSpecSetSpec` to properly reflect its global (project-scoped) nature.
  3. Ran `dev/tasks/generate-ci-cd-jobs` to generate the new e2e fixtures presubmit and periodic scripts for the new `datalabeling` service.
- **Impact**: The next agent implementing controller logic for `datalabeling` can immediately rely on the generated types, CRDs, and compiled clients.

### 2026-06-05 Implement direct types for DataLabelingEvaluationJob
- **Context**: Implementing the initial KRM types, CRD, and IdentityV2 for `DataLabelingEvaluationJob` (datalabeling.cnrm.cloud.google.com/v1alpha1).
- **Problem**: 
  1. The instructions requested `google.cloud.datalabeling.v1`, but the only version available under googleapis is `google.cloud.datalabeling.v1beta1`.
  2. The GCP URL template is `projects/{project}/evaluationJobs/{evaluation_job}`. The `gcpurls.Template` requires that variable names (lowercased) match the CamelCased Go struct fields (lowercased) exactly. If we used `{evaluation_job}` in the template, it would fail to match `EvaluationJob` (lowercased `evaluationjob`) and cause a panic.
- **Solution**: 
  1. Changed the generator service to `google.cloud.datalabeling.v1beta1`.
  2. Used the variable `{evaluationjob}` (without underscore) in the template to map to `EvaluationJob` struct field correctly.
- **Impact**: Prevented runtime panics on template initialization and ensured successful type generation.
