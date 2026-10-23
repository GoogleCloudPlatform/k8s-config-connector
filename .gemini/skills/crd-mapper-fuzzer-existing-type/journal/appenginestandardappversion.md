# AppEngineStandardAppVersion Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Kind and Proto Mapping**:
   - Resource Kind `AppEngineStandardAppVersion` maps to proto message `Version` in service `google.appengine.v1`.
   - In `apis/appengine/generate.sh`, added `--resource AppEngineStandardAppVersion:Version` to `generate-types`.

2. **Strict Schema Compatibility**:
   - Retained metadata labels: `cnrm.cloud.google.com/managed-by-kcc: "true"`, `cnrm.cloud.google.com/stability-level: alpha`, `cnrm.cloud.google.com/system: "true"`, `cnrm.cloud.google.com/tf2crd: "true"`.
   - Omitted doc comments on `AppEngineStandardAppVersionSpec` and `AppEngineStandardAppVersionStatus` to preserve `_no_description.spec` and `_no_description.status`.
   - Hand-coded `AppEngineServiceRef` locally in `appenginestandardappversion_types.go` with fields `external`, `name`, `namespace` to trigger the OpenAPI `oneOf` reference constraint validation block and match baseline CRD.
   - Preserved `name` in `AppEngineStandardAppVersionStatus` with its original description (`Full path to the Version resource in the API. Example, "v1".`), and did not include `observedState` or `externalRef` in Status.
   - Defined empty `AppEngineStandardAppVersionObservedState` struct with `// +kcc:observedstate:proto=google.appengine.v1.Version` without referencing it in Status.
   - Kept required scalar primitive types (`runtime`, `shell`, `scriptPath`, `name`) as non-pointer values and optional primitives as pointers, adhering to `reviewgen-brownfield-new-types`.

3. **Field Naming & Acronym Alignment**:
   - Renamed Go struct fields to match proto field casing and acronyms while preserving exact original JSON tags:
     - `RuntimeAPIVersion *string` with `json:"runtimeApiVersion,omitempty"`
     - `TargetCPUUtilization *float64` with `json:"targetCpuUtilization,omitempty"`
     - `HTTPHeaders map[string]string` with `json:"httpHeaders,omitempty"`
     - `URLRegex *string` with `json:"urlRegex,omitempty"`
     - `RedirectHTTPResponseCode *string` with `json:"redirectHttpResponseCode,omitempty"`
     - `VPCAccessConnector *VPCAccessConnector` with `json:"vpcAccessConnector,omitempty"`
     - Kept `AppEngineApis` (lowercase `pis`) to match proto getter `GetAppEngineApis()`.

4. **Handling Integer Fields Without OpenAPI Format**:
   - The baseline CRD schema defined numeric integer fields with `type: integer` without `format: int32`.
   - Using Go `int32` / `*int32` types causes `controller-gen` to emit `format: int32`, which trips `TestCRDBackwardCompatibility` with `field type changed: ... (integer -> int32)`.
   - Using Go `int` / `*int` produces `type: integer` without `format: int32`, maintaining 100% backward compatibility with the baseline schema.
   - Hand-coded mappers for `AutomaticScaling`, `StandardSchedulerSettings`, `BasicScaling`, `ManualScaling`, and `Deployment` cleanly cast between Go `int` and protobuf `int32`.

5. **Handcoded Mappers**:
   - In protobuf `Deployment`, files are modeled as `map<string, *FileInfo>`, whereas in KRM they are modeled as `[]DeploymentFile`. Implemented handcoded `Deployment_v1alpha1_FromProto` and `Deployment_v1alpha1_ToProto` in `pkg/controller/direct/appengine/mappers.go`.
   - Implemented helper and conversion functions in `mappers.go`: `BasicScaling_v1alpha1_ToProto`, `ManualScaling_v1alpha1_ToProto`, `Entrypoint_v1alpha1_ToProto`, `Entrypoint_Shell_ToProto`, `ScriptHandler_ScriptPath_ToProto`, `VPCAccessConnector_Name_ToProto`, and `AppEngineStandardAppVersionSpec_Runtime_ToProto` to properly convert required scalar non-pointer fields.

6. **Verification**:
   - Added comprehensive unit tests in `pkg/controller/direct/appengine/mapper_test.go` covering round-trip conversions across automatic scaling, basic scaling, manual scaling, handlers, and deployments.
   - Verified with `dev/tasks/diff-crds` that the generated CRD schema has zero unintended diffs against the baseline.
   - Ran `go test ./tests/apichecks/...` and verified all tests pass, including `TestCRDBackwardCompatibility`.
