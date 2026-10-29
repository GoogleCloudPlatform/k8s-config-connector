# AppEngineStandardAppVersion Migration Journal

## Overview
Migrated `AppEngineStandardAppVersion` (`appengine.cnrm.cloud.google.com/v1alpha1`) from the legacy Terraform-based reconciler to the `direct` reconciler architecture, with dual-controller testing against both real GCP and MockGCP.

## Key Observations and Decisions
1. **GAPIC REST Client (`cloud.google.com/go/appengine/apiv1`)**:
   - Used `api.NewVersionsRESTClient` initialized with `m.config.RESTClientOptions()`.
   - Used `pb.VersionView_FULL` in `GetVersionRequest` to ensure all fields (including `handlers` and `deployment`) are returned by the API.

2. **Patch FieldMask Specifics in App Engine Admin API**:
   - The GCP App Engine Admin API Patch (`UpdateVersion`) endpoint does not support top-level `automatic_scaling` in `updateMask`. It returns `400 Bad Request` specifying that only specific subfields are supported:
     `[automatic_scaling.cool_down_period_sec, automatic_scaling.cpu_utilization.target_utilization, automatic_scaling.max_idle_instances, automatic_scaling.max_total_instances, automatic_scaling.min_idle_instances, automatic_scaling.min_total_instances, automatic_scaling.standard_scheduler_settings.max_instances, automatic_scaling.standard_scheduler_settings.min_instances, automatic_scaling.standard_scheduler_settings.target_cpu_utilization, automatic_scaling.standard_scheduler_settings.target_throughput_utilization, instance_class, manual_scaling.instances, serving_status]`.
   - Adopted `common.CompareSpecifiedSpec(..., fineGrained: true)`. With `fineGrained: true`, `DiffForFields` automatically generates the exact fine-grained subfield paths expected by App Engine (e.g. `automatic_scaling.min_idle_instances`, `automatic_scaling.max_idle_instances`), while `diffRes.MergedDesired` preserves unmanaged fields.

3. **Diff Comparison and Normalization**:
   - `entrypoint`: The GCP App Engine API `GetVersion` endpoint does not return `entrypoint` (it is write-only upon creation). In `normalizeStandardAppVersionForDiff`, `pb.Entrypoint` is set to nil to prevent false diffs against the unreadable field.
   - `deployment.files`: The server computes and returns `sha1_sum` for deployment files, whereas KRM specs typically only specify file `name` and `source_url`. Normalized `f.Sha1Sum = ""` in `normalizeStandardAppVersionForDiff`.

4. **Dual-Controller Testing and Ratcheting Exclusion**:
   - Maintained `DefaultController: k8s.ReconcilerTypeTerraform` and added `k8s.ReconcilerTypeDirect` to `SupportedControllers` in `pkg/controller/resourceconfig/static_config.go`.
   - Retained `AppEngineStandardAppVersion` in `tests/e2e/ratcheting.go` exclusion list for `ShouldTestRereconiliation` because the fallback Terraform controller fails re-reconciliation due to Terraform state omission of `deployment.files` (which causes Terraform to trigger a duplicate version creation on re-reconciliation). The direct controller handles re-reconciliation cleanly with 0 diffs.

5. **Behavioral Alignment**:
   - Recorded authentic real GCP baseline via `hack/record-gcp`.
   - Aligned MockGCP via `hack/compare-mock`.
   - Passed `TestGoldenLogAlignment` with 0 diffs.

6. **Identity and Resource ID**:
   - `AppEngineStandardAppVersion` supports user-specified IDs. When `spec.resourceID` is unset, it defaults to `metadata.name`.
   - Implemented `identity.IdentityV2` instead of `identity.ServerGeneratedIdentity`, and resolved the finding in `apis/appengine/judgement_queue.yaml`.
