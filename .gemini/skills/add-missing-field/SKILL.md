# Skill: Add Missing Field

This skill guides an automated agent through adding a missing field to a GCP resource managed by KCC using the "direct" controller approach.

## Pre-requisites
- Know the KCC API Group and Resource name.
- Identify the missing field from the issue description or fuzzer output (e.g. `f.UnimplementedFields.Insert(".cross_instance_replication_config")`).

## Steps

1. **Add the field to the `_types.go` file**
   - Locate the `_types.go` file for the resource (e.g., `apis/group/version/resource_types.go`).
   - Add the field to the `Spec` (for inputs) or `ObservedState` (for outputs) struct as appropriate.
   - If the type relies on autogeneration, you can use `types.generated.go` as a reference. Note that `types.generated.go` includes generated code for all types if `--include-skipped-output` is passed to the generate commands.
   - **Important:** If the field represents a GCP URI or URI fragment (e.g., the path to another resource), you should use a KCC reference field (e.g., `InstanceRef *refs.MemorystoreInstanceRef`) instead of a simple string.
   - Follow KRM naming conventions (camelCase in Go, etc.).

2. **Update the Fuzzer**
   - Find the fuzzer file (e.g., `pkg/controller/direct/group/resource_fuzzer.go`).
   - Remove the missing field from `f.UnimplementedFields` or `f.Unimplemented_NotYetTriaged`.
   - Register the field as a spec or status field in the fuzzer (e.g., `f.SpecField(".cross_instance_replication_config")` or `f.StatusField(".psc_attachment_details")`).
   - If there are subfields you aren't implementing yet, use `f.Unimplemented_NotYetTriaged` to ignore them for now.

3. **Update the Mappers**
   - If the resource uses hand-written mappers (e.g., `mapper.go`), add the appropriate logic to convert between the Kubernetes resource (`KRM`) and the GCP API (`API`). 
   - Note that if the root `Spec` or `ObservedState` has a handwritten mapping function (e.g., `_FromProto` and `_ToProto`), you'll need to manually add the mapping for the new field there, even if it's just a top-level field.
   - Check `mapper.generated.go` for blocks starting with `/* found existing non-generated mapping function ... */`. If the generator skipped a parent mapper, it may also skip nested types and leave comments like `// MISSING: <Type>`. You will need to manually write `_FromProto` and `_ToProto` functions for these missing nested types.
   - **Important:** We almost always want to update the `ToProto` and `FromProto` methods "symmetrically", so that they round trip. If you map a field in `Spec_ToProto`, make sure to also map it in `Spec_FromProto`.
   - Again, `mapper.generated.go` can provide a good reference if `generate.sh --include-skipped-output` is used.

4. **Run `generate.sh`**
   - Run the code generator script (e.g., `./apis/group/vX/generate.sh` or the global `dev/tasks/generate-types-and-mappers`) to regenerate CRDs and other boilerplate.

5. **Create a Test**
   - Find existing tests under `pkg/test/resourcefixture/testdata/basic/group/version/kind/`.
   - If the new field can be added to an existing test without conflict, do so.
   - If the new field requires a specific setup (e.g., cross-instance replication requires two instances), create a new test folder with `create.yaml` (and `dependencies.yaml` if needed).

6. **Regenerate Golden Output**
   - Run the mock compare script to generate new golden output (you must set `WRITE_GOLDEN_OUTPUT=1`):
     `WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/group/version/kind/[testname]/`
   - Note that test paths may use the full kind name (e.g., `memorystoreinstance` instead of `instance`).
   - Also note that the test will still report a failure if the golden files were modified (this is expected behavior to highlight the diff).
   - Review the generated `_http.log` and `_generated_object...golden.yaml` for correctness.

7. **Update mockgcp (If Needed)**
   - If the tests fail because mockgcp doesn't support the field, you might need to implement a stub in mockgcp.
   - This stub just needs to behave reasonably so tests pass; full realgcp parity will be checked in separate E2E testing.

8. **Regenerate Go Client**
   - Please run 'make ready-pr' or 'make generate-go-client ensure fmt'.
