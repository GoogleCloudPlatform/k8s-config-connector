# AppEngineDomainMapping Fuzzer Journal

## Observations & Learnings

- **Proto & KRM Mapping Details:**
  - `AppEngineDomainMapping` maps to protobuf message `google.appengine.v1.DomainMapping` (`cloud.google.com/go/appengine/apiv1/appenginepb.DomainMapping`).
  - In KRM, `AppEngineDomainMappingSpec` maps:
    - `sslSettings`: maps to/from `.ssl_settings` (`DomainmappingSslSettings` in KRM, `SslSettings` in proto). Handcoded in `mappers.go` because `SSLManagementType` is a required non-pointer string in KRM but an enum in proto.
    - `overrideStrategy`: KRM spec-only creation parameter (not part of `pb.DomainMapping`).
    - `project`: KRM spec-only GCP project reference.
    - `resourceID`: maps to/from `.id` (relative domain name, e.g., `example.com`).
  - In KRM, `AppEngineDomainMappingStatus` maps:
    - `name`: maps to/from `.name` (full resource path).
    - `resourceRecords`: maps to/from `.resource_records` (repeated `DomainmappingResourceRecords` in KRM, repeated `ResourceRecord` in proto).
  - Both Spec and Status mappers were already implemented in `mapper.generated.go` and `mappers.go`.

- **Fuzzer Implementation & Registration:**
  - Created `pkg/controller/direct/appengine/appenginedomainmapping_fuzzer.go` in package `appengine`.
  - Annotated with `// +tool:fuzz-gen`, `// proto.message: google.appengine.v1.DomainMapping`, and `// api.group: appengine.cnrm.cloud.google.com`.
  - Registered via `init()` using `fuzztesting.RegisterKRMFuzzer(appEngineDomainMappingFuzzer())`.
  - The direct package `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/appengine"` was already imported in `pkg/controller/direct/register/register.go`.
  - Fuzzer configures:
    - `f.SpecField(".ssl_settings")`
    - `f.StatusField(".name")`
    - `f.StatusField(".resource_records")`
    - `f.Unimplemented_Identity(".id")`

- **Verification:**
  - Verified with `FOCUS=AppEngineDomainMapping go test -count=10 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers` (both Spec and Status subtests pass).
  - Verified with `go test -v ./pkg/fuzztesting/fuzztests/ -run TestListTypesWithoutFuzzers`, confirming `AppEngineDomainMapping` is recognized as covered and not in the missing list.
  - Verified with `go test -v ./pkg/controller/direct/appengine/...`.
