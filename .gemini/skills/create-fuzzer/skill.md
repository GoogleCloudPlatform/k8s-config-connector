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

4. **Map Fuzzed Proto Fields**
   - Tell the fuzzer how fields map between Spec and Status in the proto:
     - **Spec Fields:** Call `f.SpecField(".proto_field_name")` for fields mapped in the Spec. The fuzzer will ignore them during the Status round-trip.
     - **Status Fields:** Call `f.StatusField(".proto_field_name")` for fields mapped in Status/Observed State. The fuzzer will ignore them during the Spec round-trip.
     - **Unimplemented/Identity Fields:** Call `f.Unimplemented_Identity(".name")` for the resource name/ID GCP fields that are not part of the standard KCC resource's Spec or Status.
     - **Not Yet Triaged/Unimplemented Fields:** Call `f.Unimplemented_NotYetTriaged(".field_name")` or other unimplemented helpers for fields that are not yet implemented by KCC.

5. **Verify with Fuzzer Tests**
   - Run the fuzz tests quickly using:
     ```bash
     go test -count=1 -v ./pkg/fuzztesting/fuzztests/
     ```
   - Ensure the test suite completes successfully. If any fields fail round-tripping, examine the `<hint_for_agent>` in the test output to identify which fields need to be registered or marked as unimplemented.
