# AIPlatformModelMonitor Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Non-GA Proto Service (`v1beta1`) & Invocation Order in `generate.sh`**:
   - `ModelMonitor` is defined in `google.cloud.aiplatform.v1beta1.ModelMonitor`.
   - Running `generate-types` for `v1beta1` alone or after `v1` would cause `types.generated.go` to replace `v1` messages (like `ExplanationSpec`, `MachineSpec`, `GCSDestination`) with their `v1beta1` counterparts, altering unrelated CRDs (e.g. `aiplatformmodels`, `aiplatformcustomjobs`).
   - By running the `v1beta1` `ModelMonitor` generation first, followed by the canonical `v1` resources, any shared type names in `types.generated.go` are safely preserved as `v1`, while `ModelMonitor`-specific types (`ModelMonitoringInput`, `ModelMonitoringNotificationSpec`, etc.) are cleanly added without mutating any existing CRDs.
   - Required passing `--allow-non-ga-proto` to `generate-types` since `v1beta1` is a beta package.

2. **Custom Mappers in `aiplatformmodelmonitor_mapper.go`**:
   - To resolve cross-package protobuf differences between `cloud.google.com/go/aiplatform/apiv1/aiplatformpb` (v1) and `cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb` (v1beta1), custom mapper functions were created in `pkg/controller/direct/aiplatform/aiplatformmodelmonitor_mapper.go`:
     - `AIPlatformModelMonitorSpec_FromProto` / `ToProto`: maps `AIPlatformModelMonitorSpec` to `aiplatformpb.ModelMonitor`.
     - `ExplanationSpecWithRequired_FromProto` / `ToProto`: translates between `krm.ExplanationSpecWithRequired` (backed by v1 structs) and `aiplatformpb.ExplanationSpec` (v1beta1).
     - `BatchDedicatedResources_FromProto` / `ToProto`: translates `MachineSpec` between v1 and v1beta1.
     - `ModelMonitoringObjectiveSpec_DataDriftSpec_ToProto` & `FeatureAttributionSpec_ToProto`: ensures `FeatureAlertConditions` map uses `*aiplatformpb.ModelMonitoringAlertCondition`.
     - `ModelMonitoringOutputSpec_FromProto` / `ToProto`: maps `GcsDestination` cleanly.
     - `Interval_FromProto` / `ToProto`: implements mappings for `google.type.Interval`.
   - Because `mappergenerator` skips emitting functions that already exist in handwritten files in the same directory, `mapper.generated.go` compiles cleanly with zero type errors.

3. **Deterministic Identity Generation**:
   - Generated `IdentityV2` and `refs.Ref` using `${CONTROLLERBUILDER} generate-identity` in `apis/aiplatform/generate.sh` based on the pattern `projects/{project}/locations/{location}/modelMonitors/{model_monitor}`.
   - Conformance test `TestAIPlatformModelMonitorIdentity_Conformance` passed.

4. **Release Bundle Blocklist**:
   - Added `"AIPlatformModelMonitor"` in alphabetical order to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.

5. **Apichecks Golden Files**:
   - Ran `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` to record untriaged entries in `alpha-missingfields.txt` and `refs_not_representable.txt`.
