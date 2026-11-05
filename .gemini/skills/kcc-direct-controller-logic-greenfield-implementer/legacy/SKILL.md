---
name: kcc-direct-controller-logic-greenfield-implementer-legacy
description: Legacy instructions for in-flight Greenfield Phase 2 PRs that implemented the direct controller and fuzzer by hand before deterministic controller/fuzzer generation (`--prepopulate-resource`).
---

# KCC Direct Controller Logic Greenfield Implementer (Legacy / In-Flight)

Use this legacy guide **only** when iterating on an in-flight PR that already has a handwritten direct controller (`pkg/controller/direct/<service>/*_controller.go` without `// +generated:controller`) and/or a handwritten fuzzer (`*_fuzzer.go`). For all new greenfield controller implementations, use [`generated/SKILL.md`](../generated/SKILL.md).

## Inputs
- `resource_kind`: The KCC Kind (`{kind}`).
- `package_path`: The isolated package directory (e.g., `pkg/controller/direct/vertexai/examplestore/`).
- `proto_package`: The GCP proto package (e.g., `google.cloud.aiplatform.v1`).
- `service_name`: The GCP service name (short, e.g., `apigee`).
- `api_version`: The KCC API version.

## Workflow

1.  **Package Isolation**:
    Keep all controller-related logic in the existing `package_path`.

2.  **Implement Adapter Logic & Controller Patterns**:
    Update `pkg/controller/direct/<service_name>/<resource_lower>_controller.go` to implement `Find`, `Create`, `Update`, and `Delete`. Ensure correct error handling (e.g., handling 404s in `Find`), and adhere to these structural patterns:
    - **Proto Format Desired State**: Convert the KRM Spec to its Proto representation once in `AdapterForObject` and store it as a proto struct pointer (e.g., `*pb.MyResource` or `desired`) in the adapter, rather than duplicating conversion logic in both `Create` and `Update`.
    - **Handling Non-API / KRM-only Spec Fields**: If the KRM Spec has fields that are not represented in the GCP resource's proto message, parse and store them as separate fields on the adapter struct.
    - **NormalizeReferences**: Always call `common.NormalizeReferences` in `AdapterForObject` to resolve any resource references:
      ```go
      if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
          return nil, fmt.Errorf("normalizing references: %w", err)
      }
      ```
    - **Identity Parent Paths**: Use `ParentString()` on the resource's identity type instead of constructing formatting string patterns manually inside the controller.
    - **Client Creation Options & Go Client Preference**:
      * **Prefer GAPIC REST Client over gRPC (`go-client`)**: Prefer the REST variant (`NewFooRESTClient`) initialized with `m.config.RESTClientOptions()`.
      * **Fallback Order & Troubleshooting**:
        1. **GAPIC REST Client (`NewFooRESTClient`)**: Primary preference.
        2. **GAPIC gRPC Client (`NewFooClient`)**: Try with `m.config.GRPCClientOptions()` if the REST constructor is missing or fails during live GCP testing.
        3. **Raw Protobuf gRPC (`pb.NewFooClient(conn)`)**: Only fall back if no GAPIC Go client library exists.
    - **Diff Comparison & Structured Diff**: Prefer `common.CompareGreenfieldSpec(ctx, a.desired, actual, krm.<Kind>Spec_FromProto)` or `tags.DiffForTopLevelFields`.
    - **Reconciling Empty or Incomplete LRO Responses**: Always perform a GET operation (`Get<Resource>`) immediately after a Create or Update LRO completes to fetch the fully-populated resource before calling `updateStatus`.
    - **Structured Reporting & updateStatus**: In `Update`, report diffs via `structuredreporting.ReportDiff(ctx, report)` and always call `updateStatus` at the end of both `Create` and `Update`.
    - **Delete Idempotency Check**: In `Delete`, check `direct.IsNotFound(err)` and return `false, nil` (or `true, nil` if already deleted) appropriately.

3.  **Mappers**:
    Verify `mapper.generated.go` and manual mappers. Ensure all references use the standard `Ref` pattern.

4.  **Fuzzer**:
    Verify `<resource_lower>_fuzzer.go` and register it with `fuzztesting.RegisterKRMFuzzer`.
    - Prefer `f.SpecField(".foo")`, `f.StatusField(".bar")`, and `f.Unimplemented_NotYetTriaged(".baz")` / `f.Unimplemented_Identity(".name")` / `f.Unimplemented_Etag()`.

5.  **Create Minimal & Maximal Fixtures**:
    Create `pkg/test/resourcefixture/testdata/basic/<service_name>/<api_version>/<resource_lower>/<resource_lower>-minimal/` and `<resource_lower>-maximal/` with `create.yaml`, `update.yaml` (if mutable), and `dependencies.yaml`.

6.  **Remove from Ratcheting Exclusions (MANDATORY)**:
    Remove the target resource's `GroupKind` from `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go`.

7.  **MANDATORY: Record Golden Files Against Real GCP (`hack/record-gcp`)**:
    Run `RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<testname>$"` against real GCP:
    ```bash
    RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<resource_lower>-minimal$"
    RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<resource_lower>-maximal$"
    ```

8.  **Validation & Final Generation**:
    - `dev/ci/presubmits/fuzz-roundtrippers`
    - `go test ./tests/apichecks/... ./pkg/controller/direct/<service>/...` (don't run the fixtures against MockGCP here: MockGCP comes in Step 3)
    - `WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha`
    - `make ready-pr`
