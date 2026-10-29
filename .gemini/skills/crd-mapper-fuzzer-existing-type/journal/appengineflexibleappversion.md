# AppEngineFlexibleAppVersion Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Proto & Resource Mapping**:
   - The proto message corresponding to `AppEngineFlexibleAppVersion` is `google.appengine.v1.Version`.
   - In `apis/appengine/generate.sh`, added `--resource AppEngineFlexibleAppVersion:Version` to `generate-types`.

2. **Strict Schema Compatibility & Omission of Top-Level Doc Comments**:
   - The baseline CRD for `AppEngineFlexibleAppVersion` has no top-level description for `spec` or `status`. To preserve `_no_description.spec` and `_no_description.status`, doc comments were omitted from `AppEngineFlexibleAppVersionSpec` and `AppEngineFlexibleAppVersionStatus`.
   - Baseline status only contains `conditions`, `name`, and `observedGeneration`. No `externalRef` or `observedState` fields were added to `AppEngineFlexibleAppVersionStatus`.
   - An empty `AppEngineFlexibleAppVersionObservedState` struct was added with `// +kcc:observedstate:proto=google.appengine.v1.Version` to allow `generate-mapper` to output observed state mappers without polluting the CRD schema.

3. **Handling Discrepancies between Proto and Baseline CRD**:
   - The baseline CRD schema contains several field-level type and structure differences from `google.appengine.v1.Version`:
     - Fields defined as unformatted `type: integer` in the baseline CRD (e.g. `resources.cpu`, `resources.diskGb`, `resources.volumes[].sizeGb`, `manualScaling.instances`, `deployment.zip.filesCount`, `automaticScaling.*`) use Go `int` / `*int` rather than `int32` / `*int32`, ensuring `controller-gen` generates `type: integer` without `format: int32` to maintain strict schema backward compatibility with `TestCRDBackwardCompatibility`.
     - `resources.cpu` and `resources.diskGb` are integers (`*int`) in CRD but `double` (float64) in proto.
     - `resources.volumes[].sizeGb` is an integer (`int`) in CRD but `double` in proto.
     - `automaticScaling.requestUtilization.targetRequestCountPerSecond` is `string` in CRD but `int32` in proto.
     - `automaticScaling.requestUtilization.targetConcurrentRequests` is `number` (`*float64`) in CRD but `int32` in proto.
     - `deployment.files` is an array of `{name, sha1Sum, sourceUrl}` in CRD, but `map<string, FileInfo>` in proto.
     - `network.subnetwork` in CRD maps to `subnetwork_name` in proto.
     - `entrypoint.shell` in CRD maps to `command` oneof in proto (`Entrypoint_Shell`).
     - Standard environment-only fields such as `standardSchedulerSettings` and `apiEndpoint` exist in proto but are not present in the flexible app version CRD.
     - `serviceRef` is a reference with signature `"external,name,namespace"`, matching `refRuleWithoutKind` in `scripts/add-validation-to-crds`.
   - Hand-coded mapper functions (`AppEngineFlexibleAppVersionSpec_v1alpha1_FromProto`, `AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto`, `AppEngineFlexibleAppVersionStatus_v1alpha1_FromProto`, and `AppEngineFlexibleAppVersionStatus_v1alpha1_ToProto`) were implemented in `pkg/controller/direct/appengine/mappers.go`. `generate-mapper` automatically skips generating duplicates when these exist.

4. **Validation & Verification**:
   - `dev/tasks/diff-crds --base upstream/master` produces completely empty output, verifying zero schema changes.
   - Comprehensive unit tests were added in `pkg/controller/direct/appengine/mapper_test.go` covering round-trip conversions between proto and KRM types.
   - All tests and custom linters pass cleanly.
