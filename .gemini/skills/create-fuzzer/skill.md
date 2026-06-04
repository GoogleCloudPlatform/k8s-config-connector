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
   - **CRITICAL - Helper Wrappers Only**: Always use the wrapper helper functions (such as `f.SpecField()`, `f.StatusField()`, `f.Unimplemented_Identity()`) instead of directly manipulating the fields sets via `.Insert()` (e.g. do NOT use `f.SpecFields.Insert()`, `f.StatusFields.Insert()`, or `f.UnimplementedFields.Insert()`). This maintains clean api patterns and safety checks.

5. **Verify with Fuzzer Tests**
   - Ensure the package is registered centrally by adding an import of the package in `pkg/controller/direct/register/register.go`.
   - **Do NOT** create a separate test file (e.g. `<resource>_fuzzer_test.go`) for the fuzzer. All fuzzers are registered and executed centrally.
   - Run the fuzz tests quickly to verify your implementation:
     ```bash
     go test -count=1 -v ./pkg/fuzztesting/fuzztests/
     ```
   - Ensure the test suite completes successfully. If any fields fail round-tripping, examine the `<hint_for_agent>` in the test output to identify which fields need to be registered or marked as unimplemented.

## Tips & Troubleshooting

### Missing Proto Messages during Code Generation
If running `generate.sh` fails with `failed to find the proto message ...: proto: not found`, it is likely because the proto directory (e.g. `google/identity/...`) is not being matched and compiled by `dev/tools/controllerbuilder/generate-proto.sh`. Add the necessary matching pattern (e.g. `${THIRD_PARTY}/googleapis/google/identity/*/*/*.proto`) to `generate-proto.sh` and remove any cached `.pb` files in `.build/` to force reconstruction.

### Spec Fields Expecting Custom Map Functions
If compilation of generated mappers fails due to undefined mapper functions like `SomeSpec_SomeField_ToProto`, check the Go type of that field. In KCC, the `mappergenerator` expects spec string fields to be declared as pointer types (`*string`). If declared as non-pointers (e.g. `string`), it expects a custom map function. Switching the field to a pointer (e.g. `*string`) allows automatic mapper generation.

### Parent Resource References (Organization / Folder / Project)
By default, the template generator scaffolds a `projectRef` parent reference. If the resource is parented by an Organization or a Folder, manually replace `projectRef *refsv1beta1.ProjectRef` with `organizationRef *refsv1beta1.OrganizationRef` or `folderRef *refsv1beta1.FolderRef` in the `_types.go` spec, and make sure to remove any irrelevant `location` fields if the resource is global.

