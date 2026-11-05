# AppEngineFlexibleAppVersion Fuzzer Journal

## Observations & Implementation Details

1. **Proto & KRM Mapping Details:**
   - `AppEngineFlexibleAppVersion` maps to protobuf message `google.appengine.v1.Version` (`cloud.google.com/go/appengine/apiv1/appenginepb`).
   - Group: `appengine.cnrm.cloud.google.com`, Version: `v1alpha1`.
   - The KRM spec and status mappers `AppEngineFlexibleAppVersionSpec_v1alpha1_{From,To}Proto` and `AppEngineFlexibleAppVersionObservedState_v1alpha1_{From,To}Proto` were already implemented in `pkg/controller/direct/appengine/mappers.go` and `mapper.generated.go`.

2. **Oneof Handling in Entrypoint:**
   - `pb.Entrypoint.Command` is a protobuf `oneof` containing `*Entrypoint_Shell`.
   - In `AppEngineFlexibleAppVersionSpec_v1alpha1_ToProto`, `out.Entrypoint.Command` was unconditionally set to `&pb.Entrypoint_Shell{Shell: in.Entrypoint.Shell}` whenever `in.Entrypoint != nil`.
   - When a proto has an empty `Entrypoint{}` (no oneof branch set), `in.GetEntrypoint().GetShell()` returns `""`. If `ToProto` unconditionally instantiates `Entrypoint_Shell`, it clobbers the unset oneof with `shell: ""`.
   - Added `if in.Entrypoint.Shell != ""` check before assigning `Command`, matching `Entrypoint_v1alpha1_ToProto` in the same package.

3. **Integer vs Float Invariants in Resources (FilterSpec):**
   - In protobuf `google.appengine.v1.Resources`, `cpu` and `disk_gb` are `float64`. In `google.appengine.v1.Volume`, `size_gb` is `float64`.
   - In KRM `AppEngineFlexibleAppVersionResources`, `cpu` and `diskGb` are `*int`, and `volumes[].sizeGb` is `int`.
   - Round-tripping arbitrary fractional floats through integer KRM types causes precision loss (e.g., `-1.849` -> `-1`).
   - Configured `f.FilterSpec` to truncate `in.Resources.Cpu`, `in.Resources.DiskGb`, and `v.SizeGb` to integer float values before round-tripping.

4. **Unimplemented / Untriaged Fields:**
   - `.name` and `.id` are marked as `Unimplemented_Identity`.
   - Untriaged fields in `pb.Version` not present in flexible version KRM:
     - `.basic_scaling`
     - `.zones`
     - `.threadsafe`
     - `.vm`
     - `.app_engine_apis`
     - `.env`
     - `.created_by`
     - `.create_time`
     - `.disk_usage_bytes`
     - `.error_handlers`
     - `.libraries`
     - `.build_env_variables`
     - `.health_check`
     - `.version_url`
   - Nested untriaged fields:
     - `.automatic_scaling.standard_scheduler_settings`
     - `.deployment.files[].mime_type`
     - `.handlers[].api_endpoint`
     - `.resources.kms_key_reference`
     - `.vpc_access_connector.egress_setting`

5. **Test Results:**
   - Successfully verified against `TestListTypesWithoutFuzzers` (Kinds with Fuzzers increased from 573 to 574).
   - Verified unit tests in `pkg/controller/direct/appengine/...`.
   - Verified central fuzz suite with `go test -count=1 ./pkg/fuzztesting/fuzztests/` and focused mapper test with 100 iterations.
