---
name: kcc-direct-controller-implementer
description: Implement the controller, mappers, and fuzzer for a direct KCC resource, ensuring package isolation and CI compliance. Use this when implementing the main reconciliation logic for a "direct" resource.
---

# KCC Direct Controller Implementer

This skill guides the implementation of the controller, mappers, and fuzzer for direct KCC resources, with a focus on package isolation to prevent symbol collisions and rigorous CI validation.

## Inputs
- `resource_kind`: The KCC Kind.
- `package_path`: The isolated package directory (e.g., `pkg/controller/direct/vertexai/examplestore/`).
- `proto_package`: The GCP proto package (e.g., `google.cloud.aiplatform.v1`).

## Workflow

1.  **Package Isolation**:
    You MUST implement all controller-related logic in the provided `package_path`. This prevents symbol collisions in `mapper.generated.go`.

2.  **Brownfield Migrations**:
    For brownfield resources (with existing terraform/DCL controllers), we do NOT immediately change the default controller to `direct` in Go type labels or static config map defaults.
    - Keep the default controller as legacy (TF/DCL), so users can opt-in gradually.
    - Both direct and legacy controllers are now automatically tested dynamically if a direct controller is available in `static_config.go` (under `SupportedControllers`), so there is no longer any need to manually modify `tests/e2e/unified_test.go`.

3.  **Controller Structure & Patterns**:
    - **Proto Format Desired State**: Convert the KRM Spec to its Proto representation once in `AdapterForObject` and store it as a proto struct pointer (e.g., `*pb.MyResource` or `desired`) in the adapter, rather than duplicating conversion logic in both `Create` and `Update`. The adapter should avoid holding references to raw KRM objects for desired state to keep the interfaces clean, consistent with `actual` (which is also of proto type), and avoid redundant conversions. This ensures that `desired` is stored in the same proto format as `actual`.
    - **Handling Non-API / KRM-only Spec Fields**: If the KRM Spec has fields that are not represented in the GCP resource's proto message (such as client-side behavioral options or custom installation flags, e.g., `SkipInitialVersionCreation`), do NOT mix them into the proto. Instead, parse and store them as separate, explicitly-named fields on the adapter struct (e.g., `desiredSkipInitialVersionCreation bool`).
    - **NormalizeReferences**: Always call `common.NormalizeReferences` in `AdapterForObject` to resolve any resource references:
      ```go
      if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
          return nil, fmt.Errorf("normalizing references: %w", err)
      }
      ```
    - **Identity Parent Paths**: Create a `ParentString()` method on the resource's identity type (e.g., `KMSCryptoKeyIdentity`) instead of constructing formatting string patterns manually inside the controller. This keeps parent paths canonical and reusable.
    - **Client Creation Options**: Do NOT build the authenticated HTTP client manually. Instead, retrieve configuration options using `RESTClientOptions()` and construct the REST client:
      ```go
      var opts []option.ClientOption
      opts, err := m.config.RESTClientOptions()
      if err != nil {
          return nil, err
      }
      gcpClient, err := gcp.NewConfigRESTClient(ctx, opts...)
      ```
    - **Diff Comparison (`compare<Resource>` pattern)**: Implement a dedicated `compare<Resource>` helper function to compare the spec fields of actual/desired proto structures by round-tripping actual via its KRM Spec using `mappers.OnlySpecFields` and calling `tags.DiffForTopLevelFields`. Define a `populateDefaults := func(obj *pb.MyResource)` function inside `compare<Resource>` and call it for both `maskedActual` and `desired` clone before comparison to populate any defaults (e.g., fields that default to specific values when nil / omitted):
      ```go
      func compareResource(ctx context.Context, actual, desired *pb.MyResource) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
          maskedActual, err := mappers.OnlySpecFields(actual, MyResourceSpec_FromProto, MyResourceSpec_ToProto)
          if err != nil {
              return nil, nil, err
          }
          maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

          clonedDesired := proto.Clone(desired).(*pb.MyResource)

          populateDefaults := func(obj *pb.MyResource) {
              // Even if empty, it's a good pattern to define and populate GCP/server defaults here
              if obj.SomeDefaultedField == nil {
                  obj.SomeDefaultedField = ...
              }
          }
          populateDefaults(maskedActual)
          populateDefaults(clonedDesired)

          diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
          if err != nil {
              return nil, nil, err
          }
          return diffs, updateMask, nil
      }
      ```
    - **Structured Reporting & updateStatus**: In the `Update` method, use `diffs.HasDiff()` to report exact diffs back to the user via `structuredreporting.ReportDiff`. Always call a helper `updateStatus` function to update the Kube status at the end of both `Create` and `Update` reconciliation paths:
      ```go
      func (a *MyResourceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MyResource) error {
          mapCtx := &direct.MapContext{}
          status := MyResourceStatus_FromProto(mapCtx, latest)
          if mapCtx.Err() != nil {
              return mapCtx.Err()
          }
          return op.UpdateStatus(ctx, status, nil)
      }
      ```
    - **Immutable Resources**: If a direct resource is completely immutable in GCP (meaning no fields can be updated once created), the `Update` method must STILL perform the `compare<Resource>` comparison check on spec fields. If a diff is detected, return a descriptive error (e.g. `fmt.Errorf("<Kind> is immutable and cannot be updated")`) so that the error/diff is surfaced on the resource status rather than silently doing nothing. Also, register the model using `registry.CannotBeDeleted()` if deletion is not supported.

4.  **Mappers**:
    Verify `mapper.generated.go` and manual mappers. Ensure all references use the standard `Ref` pattern.

5.  **Fuzzer**:
    Implement `<resource_lower>_fuzzer.go` and register it with `fuzztesting.RegisterKRMFuzzer`.
    - **Fuzzer Implementation Guidelines**:
      * ALWAYS prefer and encourage using the fluent `f.SpecField(".foo")` and `f.StatusField(".bar")` methods rather than inserting directly into `SpecFields`/`StatusFields` sets (e.g. avoid `f.SpecFields.Insert...`).
      * For unhandled or unimplemented fields, prefer highly descriptive helper methods (such as `f.Unimplemented_NotYetTriaged(".baz")`, `f.Unimplemented_Identity(".id")`, or other specific variants) rather than standard inserts into `UnimplementedFields` sets. This categorizes why we are not handling specific fields.

6.  **Final Generation & Reporting**:
    Run `make ready-pr` from the repository root. This is a critical step that:
    - Regenerates all Go clients in `pkg/clients/generated`.
    - Updates global CRD reports (`docs/reports/crd_report.md` and `.csv`).
    - Runs `make fmt` and `make vet`.
    YOU MUST COMMIT ALL RESULTING CHANGES.

7.  **Validation & Last-Mile Tests**:
    Run the following tests to ensure CI compliance:
    - **Fuzzing**: `dev/ci/presubmits/fuzz-roundtrippers`
    - **E2E Scaffolding**: `dev/ci/presubmits/tests-e2e-fixtures-direct`
    - **Schema Integrity**: `go test ./pkg/crd/template/...`
    - **API Field Coverage**: `go test ./tests/apichecks/...`. 
      - If `TestCRDFieldPresenceInTestsForAlpha` fails, you may need to regenerate the exceptions file by running it with `WRITE_GOLDEN_OUTPUT=1`.

## Journaling
Append any controller/reconciliation quirks or API check hurdles to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
