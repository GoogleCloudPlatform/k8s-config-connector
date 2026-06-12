# Skill: Create Fuzzer for Direct Controller

This skill guides an automated agent through the process of implementing a round-trip KRM fuzzer for a direct Config Connector resource controller. Round-trip fuzz testing verifies the mapping functions can accurately convert KRM objects to GCP protos and back without losing or misrepresenting data.

## Steps

1. **Find or Implement Observed State ToProto Mapping**
   - In the controller's mapper file (typically `<resource>_mapper.go` or `mapper.generated.go`), ensure that both `FromProto` and `ToProto` functions exist for both the Spec and the Status (Observed State).
   - If `StatusObservedState_ToProto` (or similar) does not exist, implement it manually. For example:
     ```go
     func KMSKeyHandleStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyHandleObservedState) *pb.KeyHandle {
         if in == nil {
             return nil
         }
         out := &pb.KeyHandle{}
         if in.KMSKey != nil {
             out.KmsKey = *in.KMSKey
         }
         return out
     }
     ```

2. **Create the Fuzzer File**
   - Create a new fuzzer file inside the resource's direct controller package: `pkg/controller/direct/<service>/<resource>_fuzzer.go` (or `pkg/controller/direct/<service>/<resource>/<resource>_fuzzer.go`).
   - Use the appropriate package name matching the surrounding controller files.
   - Annotate the file with the appropriate `+tool:fuzz-gen`, `proto.message`, and `api.group` metadata.

3. **Register the Fuzzer in `init()`**
   - Register the fuzzer using `fuzztesting.RegisterKRMFuzzer` inside the file's `init()` function.
   - Instantiate the fuzzer using `fuzztesting.NewKRMTypedFuzzer`, passing:
     - A pointer to the zero-value proto struct (e.g. `&pb.KeyHandle{}`)
     - The Spec FromProto mapper
     - The Spec ToProto mapper
     - The Status ObservedState FromProto mapper
     - The Status ObservedState ToProto mapper
   - Ensure the direct package is registered in `pkg/controller/direct/register/register.go` by adding an anonymous import in alphabetical order (e.g., `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/<service>"`), otherwise its `init()` function will not be executed by the fuzz test suite.

4. **Map Fuzzed Proto Fields**
   - Tell the fuzzer how fields map between Spec and Status in the proto:
     - **Spec Fields:** Call `f.SpecField(".proto_field_name")` for fields mapped in the Spec. The fuzzer will ignore them during the Status round-trip.
     - **Status Fields:** Call `f.StatusField(".proto_field_name")` for fields mapped in Status/Observed State. The fuzzer will ignore them during the Spec round-trip.
     - **Unimplemented/Identity Fields:** Call `f.Unimplemented_Identity(".name")` for the resource name/ID GCP fields that are not part of the standard KCC resource's Spec or Status.
     - **Not Yet Triaged/Unimplemented Fields:** Call `f.Unimplemented_NotYetTriaged(".field_name")` or other unimplemented helpers for fields that are not yet implemented by KCC.
   - **Document Field Comparison**: Always add detailed comparison comments in the fuzzer file (just above `f.SpecField` calls) that explicitly compare the KRM Spec type fields with their corresponding fuzzer/proto field mapping status. This helps reviewers verify that every field in the KRM Spec was accounted for, especially when fields map to nested untriaged proto structures (e.g. `lifecycleRule` maps to `f.Unimplemented_NotYetTriaged(".lifecycle")`).
   - **CRITICAL - Helper Wrappers Only**: Always use the wrapper helper functions (such as `f.SpecField()`, `f.StatusField()`, `f.Unimplemented_Identity()`) instead of directly manipulating the fields sets via `.Insert()` (e.g. do NOT use `f.SpecFields.Insert()`, `f.StatusFields.Insert()`, or `f.UnimplementedFields.Insert()`). This maintains clean api patterns and safety checks.

5. **Verify with Fuzzer Tests**
   - Ensure the package is registered centrally by adding an import of the package in `pkg/controller/direct/register/register.go`.
   - **CRITICAL / DO NOT**: Never create a separate test file (e.g. `<resource>_fuzzer_test.go` or any `*_test.go` file inside the direct controller directories) for individual fuzzers. Creating custom, single-fuzzer unit tests is strictly discouraged and will fail reviews. All fuzzers are registered via `init()` and executed centrally under the central fuzz test suite.
   - Run the central fuzz tests quickly to verify your implementation:
     ```bash
     go test -count=1 -v ./pkg/fuzztesting/fuzztests/
     ```
   - Ensure the test suite completes successfully. If any fields fail round-tripping, examine the `<hint_for_agent>` in the test output to identify which fields need to be registered or marked as unimplemented.

## Creating a Fuzzer for No-Proto (OpenAPI) Resources

When a direct controller manages a resource that does not have a proto schema (e.g., OpenAPI-based or Discovery API-based resources such as DNS ManagedZone), use the NoProto KRM Fuzzer framework and automatic OpenAPI mapper generation.

### 1. Auto-Generate Mappers with `generate-mapper`
Use the generator tool to automatically generate mappers between the KRM types and the OpenAPI Go SDK types:
```bash
go run "${REPO_ROOT}/dev/tools/openapi-to-krm/cmd/generate-mapper/main.go" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/<service>/<version>/<KRMTypeSpec>::google.golang.org/api/<service>/<version>/<APIType>" \
  --mapper "github.com/GoogleCloudPlatform/k8s-config-connector/apis/<service>/<version>/<KRMTypeStatus>::google.golang.org/api/<service>/<version>/<APIType>" \
  --output-file "${REPO_ROOT}/pkg/controller/direct/<service>/zz_generated.mappers.go"
```
The generator will:
- Recursively find and map nested child structures using a BFS type traversal.
- Match fields case-insensitively.
- Automatically recognize Reference fields (`Refs`) that end with `Ref` (e.g. `NetworkRef`) and map them to/from the corresponding API fields (e.g. `NetworkUrl` or `GkeClusterName`).
- Detect pre-existing handcoded functions in the package (e.g., in `mappers.go`) and automatically comment out their generated equivalents.

### 2. Implement Custom Mappers in `mappers.go`
For any field conversions requiring custom translation logic (such as different field names or non-standard types), implement them manually in `pkg/controller/direct/<service>/mappers.go`. The generator will automatically notice them and comment out its generated equivalents, leaving your custom handcoded ones active.

### 3. Create the Fuzzer File
Create `pkg/controller/direct/<service>/<resource>_fuzzer.go` using package name `<service>`. Rather than registering with `RegisterKRMFuzzer` (which expects protos), register it with `RegisterKRMFuzzer_NoProto`:
```go
package <service>

import (
	api "google.golang.org/api/<service>/<version>"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(dnsManagedZoneFuzzer())
}

func dnsManagedZoneFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.ManagedZone{},
		DNSManagedZoneSpec_FromAPI, DNSManagedZoneSpec_ToAPI,
		DNSManagedZoneStatus_FromAPI, DNSManagedZoneStatus_ToAPI,
	)

	// Register fields...
	f.SpecField(".CloudLoggingConfig")
	f.StatusField(".CreationTime")
	f.IdentityField(".Name")

	return f
}
```

### 4. Register the Direct Package
Ensure `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/<service>"` is imported in `pkg/controller/direct/register/register.go` so the fuzzer registers itself centrally.

### 5. Run and Verify Fuzz Tests
Run the central fuzz test suite:
```bash
go test -count=1 -v ./pkg/fuzztesting/fuzztests/
```

