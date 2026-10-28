# AppEngineStandardAppVersion Fuzzer Journal

## Observations & Implementation Details

1. **Proto & KRM Mapping Details:**
   - `AppEngineStandardAppVersion` maps to protobuf message `google.appengine.v1.Version` (`cloud.google.com/go/appengine/apiv1/appenginepb`).
   - Mappers for `AppEngineStandardAppVersion` already exist in `pkg/controller/direct/appengine/mapper.generated.go` and `mappers.go`.
   - `AppEngineStandardAppVersionSpec` maps standard version configurations, including scaling (`automatic_scaling`, `basic_scaling`, `manual_scaling`), `deployment`, `entrypoint`, `env_variables`, `handlers`, `inbound_services`, `instance_class`, `libraries`, `runtime`, `runtime_api_version`, `service_account`, `threadsafe`, and `vpc_access_connector`.
   - `AppEngineStandardAppVersionObservedState` is currently an empty struct with no status fields.
   - The resource identity is server-generated/URL-based (`apps/{app}/services/{service}/versions/{version}`). `.name` and `.id` are marked with `f.Unimplemented_Identity`.

2. **Nested Map Field Wildcards (`ClearFields`):**
   - In `google.appengine.v1.Deployment`, the `files` field is defined as a `map[string]*FileInfo` where `FileInfo` has fields `source_url`, `sha1_sum`, and `mime_type`.
   - KRM's `DeploymentFile` maps `source_url` and `sha1_sum`, but does not support `mime_type`.
   - When fuzzing maps of messages with random keys, `Visit` visits children using `mapPath := path + "[" + k.String() + "]"`.
   - Enhanced `ClearFields` in `pkg/test/fuzz/generate.go` to support wildcard bracket matching `[]` across map keys (`[key]` matches `[]`). This allows `f.Unimplemented_NotYetTriaged(".deployment.files[].mime_type")` to properly clear the untriaged nested field in map elements during fuzz testing.

3. **Scaling and Oneof Fields:**
   - `Version` proto defines a oneof `scaling` for `automatic_scaling`, `basic_scaling`, and `manual_scaling`.
   - `automatic_scaling` proto contains several fields not exposed in KRM (such as `cool_down_period`, `cpu_utilization`, `max_total_instances`, `min_total_instances`, `request_utilization`, `disk_utilization`, `network_utilization`), which are marked with `f.Unimplemented_NotYetTriaged`.
   - `UrlMap` proto has a `handler_type` oneof containing `static_files`, `script`, and `api_endpoint`. Since `api_endpoint` is not exposed in KRM, it is marked with `f.Unimplemented_NotYetTriaged(".handlers[].api_endpoint")`.

4. **Verification:**
   - Tested specifically with `FOCUS=AppEngineStandardAppVersion go test -count=100 ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`, passing 100 iterations.
   - Verified that `AppEngineStandardAppVersion` is removed from `TestListTypesWithoutFuzzers` (Kinds with Fuzzers increased from 560 to 561).
   - Verified unit tests in `pkg/test/fuzz/...` and `pkg/controller/direct/appengine/...`.
   - Verified central fuzz suite with `go test ./pkg/fuzztesting/fuzztests/`.
