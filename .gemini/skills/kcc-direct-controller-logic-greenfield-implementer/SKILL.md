---
name: kcc-direct-controller-logic-greenfield-implementer
description: Guides the implementation of Phase 2 (direct controller logic, E2E fixtures, and fuzzer) for KCC Greenfield resources, with a focus on real GCP recording and verification.
---

# KCC Direct Controller Logic Greenfield Implementer

Please follow the instructions below to implement the direct controller and record/verify E2E fixtures for `{kind}`.

The direct controller must be implemented to manage reconciliation logic (Adapter: Find, Create, Update, Delete) and E2E fixtures must be created and recorded against real GCP to verify functionality.

## Inputs
- `resource_kind`: The KCC Kind (`{kind}`).
- `package_path`: The isolated package directory (e.g., `pkg/controller/direct/vertexai/examplestore/`).
- `proto_package`: The GCP proto package (e.g., `google.cloud.aiplatform.v1`).
- `service_name`: The GCP service name (short, e.g., `apigee`).
- `api_version`: The KCC API version.

## Workflow

1.  **Package Isolation**:
    You MUST implement all controller-related logic in the provided `package_path`. This prevents symbol collisions in `mapper.generated.go`.

2.  **Implement Adapter Logic & Controller Patterns**:
    Update `pkg/controller/direct/<service_name>/<resource_lower>_controller.go` to implement `Find`, `Create`, `Update`, and `Delete`. Ensure correct error handling (e.g., handling 404s in `Find`), and adhere to these structural patterns:
    - **Proto Format Desired State**: Convert the KRM Spec to its Proto representation once in `AdapterForObject` and store it as a proto struct pointer (e.g., `*pb.MyResource` or `desired`) in the adapter, rather than duplicating conversion logic in both `Create` and `Update`. The adapter should avoid holding references to raw KRM objects for desired state to keep the interfaces clean, consistent with `actual` (which is also of proto type), and avoid redundant conversions. This ensures that `desired` is stored in the same proto format as `actual`.
    - **Handling Non-API / KRM-only Spec Fields**: If the KRM Spec has fields that are not represented in the GCP resource's proto message (such as client-side behavioral options or custom installation flags, e.g., `SkipInitialVersionCreation`), do NOT mix them into the proto. Instead, parse and store them as separate, explicitly-named fields on the adapter struct (e.g., `desiredSkipInitialVersionCreation bool`).
    - **NormalizeReferences**: Always call `common.NormalizeReferences` in `AdapterForObject` to resolve any resource references:
      ```go
      if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
          return nil, fmt.Errorf("normalizing references: %w", err)
      }
      ```
    - **Identity Parent Paths**: Create a `ParentString()` method on the resource's identity type (e.g., `KMSCryptoKeyIdentity`) instead of constructing formatting string patterns manually inside the controller. This keeps parent paths canonical and reusable.
    - **Client Creation Options & Go Client Preference**:
      * **Prefer GAPIC REST Client over gRPC (`go-client`)**: When building the GCP client in `m.client(ctx)`, ALWAYS prefer using the official Google Cloud Go client library (GAPIC client, e.g., `cloud.google.com/go/<service>/apivX`, or discovery client `google.golang.org/api/...`). Many GAPIC libraries provide both REST (`NewFooRESTClient`) and gRPC (`NewFooClient`) constructors. **You MUST prefer the REST variant (`NewFooRESTClient`)** initialized with `m.config.RESTClientOptions()`.
      * Why: REST GAPIC clients integrate cleanly with HTTP golden traffic recording (`_http.log`), mock layer HTTP alignment, and avoid gRPC transport/auth quirks, while still automatically wrapping Long-Running Operations (`*service.FooOperation`), OAuth scopes, and retry logic.
      * **Client Construction Example (GAPIC REST Primary Choice)**: Retrieve configuration options using `m.config.RESTClientOptions()` and construct the REST client:
        ```go
        var opts []option.ClientOption
        opts, err := m.config.RESTClientOptions()
        if err != nil {
            return nil, err
        }
        // Always prefer GAPIC New*RESTClient over New*Client (gRPC) or raw pb clients
        gcpClient, err := gcp.NewMyResourceRESTClient(ctx, opts...)
        if err != nil {
            return nil, fmt.Errorf("building MyResource REST client: %w", err)
        }
        ```
      * **Fallback Order & Troubleshooting**:
        1. **GAPIC REST Client (`NewFooRESTClient`)**: Primary preference.
        2. **GAPIC gRPC Client (`NewFooClient`)**: Try with `m.config.GRPCClientOptions()` if the REST constructor is missing or fails during live GCP testing.
        3. **Raw Protobuf gRPC (`pb.NewFooClient(conn)`)**: Only fall back to dialing raw gRPC conns (`grpc.Dial`) if no GAPIC Go client library struct/constructor exists for the API.
      * **Try Different Client Variants on Issues During Testing**: When testing against real GCP (`./hack/record-gcp` or API checks) and encountering `403 Forbidden`, `404 Not Found`, or LRO polling errors, actively switch between the client variants (REST vs gRPC) to find the working transport.
    - **Diff Comparison & Structured Diff (`tags.DiffForTopLevelFields`)**: Always prefer using top-level field tags-based diff comparison via `tags.DiffForTopLevelFields` over recursive/magical comparison functions (such as `common.CompareProtoMessageStructuredDiff` or `common.CompareProtoMessage` which are deprecated/discouraged due to unpredictable behaviors in `BasicDiff`).
      ```go
      func compareResource(ctx context.Context, actual, desired *pb.MyResource) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
          maskedActual, err := mappers.OnlySpecFields(actual, MyResourceSpec_FromProto, MyResourceSpec_ToProto)
          if err != nil {
              return nil, nil, err
          }
          maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

          clonedDesired := proto.CloneOf(desired)

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
    - **Reconciling Empty or Incomplete LRO Responses**: Many GCP APIs (such as Dataproc's `UpdateCluster` LRO) return an empty response (`google.protobuf.Empty`), or do not fully populate read-only status fields (such as state, metrics, or instance names) during resource creation. If you map status directly from such incomplete/empty LRO responses, you will inadvertently clear status fields in Kubernetes.
      * **Rule**: Always perform a GET operation (`Get<Resource>`) immediately after a Create or Update LRO successfully completes to fetch the fully-populated resource before calling `updateStatus`.
    - **Propagating KRM Metadata Labels**: Metadata labels (such as `managed-by-cnrm: true` or custom user-supplied labels) must be explicitly mapped and propagated on both `Create` and `Update` operations so they are correctly synchronized to GCP.
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
    - **Delete Idempotency Check**: In the `Delete` method, check if the resource has already been deleted using `direct.IsNotFound(err)` to ensure idempotency, returning `true, nil` to gracefully exit:
      ```go
      if err != nil {
          if direct.IsNotFound(err) {
              return true, nil
          }
          return false, err
      }
      ```
    - **Immutable Resources**: If a direct resource is completely immutable in GCP (meaning no fields can be updated once created), the `Update` method must STILL perform the comparison check on spec fields. If a diff is detected, return a descriptive error (e.g. `fmt.Errorf("<Kind> is immutable and cannot be updated")`) so that the error/diff is surfaced on the resource status rather than silently doing nothing. Also, register the model using `registry.CannotBeDeleted()` if deletion is not supported.
    - **Mutable-but-Unreadable Fields**: If certain spec fields (such as keys, passwords, or specific metadata/launchStage fields) are mutable but cannot be read back from the GCP API (meaning they are write-only or missing from the GET response), never check `desired` inside `populateDefaults`. Instead:
      1. Check the resource's `updateTime` or change cookie to see if the GCP side is fully reconciled and unchanged (meaning `generation == observedGeneration` and GCP `updateTime` matches status `updateTime` on a successful/Ready status).
      2. If the update time matches and there are no external modifications, copy those mutable-but-unreadable fields from `desired` to `maskedActual` in the comparison function to avoid false diffs.
      3. If the update time does not match (or the resource is not ready), do NOT copy them. This treats the unreadable fields as having changed (since we cannot prove they didn't), correctly forcing an update.

3.  **Mappers**:
    Verify `mapper.generated.go` and manual mappers. Ensure all references use the standard `Ref` pattern.

4.  **Fuzzer**:
    Implement `<resource_lower>_fuzzer.go` and register it with `fuzztesting.RegisterKRMFuzzer`.
    - **Fuzzer Implementation Guidelines**:
      * ALWAYS prefer and encourage using the fluent `f.SpecField(".foo")` and `f.StatusField(".bar")` methods rather than inserting directly into `SpecFields`/`StatusFields` sets (e.g. avoid `f.SpecFields.Insert...`).
      * For unhandled or unimplemented fields, prefer highly descriptive helper methods (such as `f.Unimplemented_NotYetTriaged(".baz")`, `f.Unimplemented_Identity(".id")`, or other specific variants) rather than standard inserts into `UnimplementedFields` sets. This categorizes why we are not handling specific fields.

5.  **Create Minimal Fixture**:
    Create directory `pkg/test/resourcefixture/testdata/basic/<service_name>/<api_version>/<resource_lower>/<resource_lower>-minimal/`.
    - Add `create.yaml`: Use the bare minimum **Required** fields.
    - Use `${uniqueId}` for resource names.

6.  **Create Maximal Fixture**:
    Create directory `pkg/test/resourcefixture/testdata/basic/<service_name>/<api_version>/<resource_lower>/<resource_lower>-maximal/`.
    - Add `create.yaml`: Include **every supported field** in the Spec.
    - Add `update.yaml`: Update all **mutable** fields.
    - Add `dependencies.yaml` if the resource requires other KCC resources to exist first.

6.5. **Remove from Ratcheting Exclusions (MANDATORY)**:
    Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.
    1. Open `tests/e2e/ratcheting.go`.
    2. Locate the function `ShouldTestRereconiliation`.
    3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
    4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

7.  **MANDATORY: Record Golden Files Against Real GCP (`hack/record-gcp`)**:
    - **CRITICAL OVERRIDE OF GEMINI.md**: For this Greenfield task, **ignore any instructions in `GEMINI.md`** (or `mockgcp/GEMINI.md`) regarding `mockgcp`, `E2E_GCP_TARGET=mock`, or `hack/compare-mock`. Those global instructions only apply to legacy brownfield resources.
    - **STRICT GUARDRAIL — DO NOT USE MOCKGCP OR OFFLINE MOCKS**: You MUST record golden files directly against real GCP using `./hack/record-gcp`. Do **NOT** attempt to use `mockgcp`, do NOT search for or try to implement mockgcp services, and do NOT use `hack/compare-mock` or `E2E_GCP_TARGET=mock`.
    - **MANDATORY EXECUTION**: You are explicitly required to run `./hack/record-gcp` on both your minimal and maximal fixtures before running any validation tests in Step 8 or preparing the PR in Step 9. Skipping this step or attempting to substitute mock tests is considered a critical failure.
    Run `RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<testname>$"` to capture real GCP behavior and record traffic, object state, and live REST GET probe audit receipt (`_audit_probe.log`):
    ```bash
    # Run from the repository root for both minimal and maximal fixtures
    RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<resource_lower>-minimal$"
    RECORD_AUDIT_PROBE=1 ./hack/record-gcp "fixtures/^<resource_lower>-maximal$"
    ```
    - **Troubleshooting Real GCP Errors (Do Not Give Up / No Mock Fallback)**:
      * **Missing Dependent Resources**: If real GCP returns an error indicating that a referenced resource does not exist (e.g., `ForwardingRule non-existent-rule does not exist`, `Network default does not exist`, `ServiceAccount not found`), you MUST create that prerequisite resource in `dependencies.yaml` (e.g. a `ComputeForwardingRule`, `ComputeNetwork`, `IAMServiceAccount`, etc.) so the test harness provisions it in GCP before testing your resource.
      * **Invalid Parameters / Constraint Violations**: If real GCP returns `The request was invalid`, `InvalidArgument`, or `an internal error has occurred` due to invalid or conflicting spec fields, inspect the GCP API reference, fix `create.yaml` or `update.yaml` to specify valid configuration, and re-run.
      * **Disabled GCP APIs**: If `hack/record-gcp` fails because a GCP service is not enabled, enable the service using `gcloud services enable <service-name>.googleapis.com` and try again.
      * **Strict Rule**: NEVER fall back to `compare-mock` or mock logs. Assume mocks do not exist. Iterate on `dependencies.yaml`, `create.yaml`, `update.yaml`, and controller code until `./hack/record-gcp` passes cleanly against real GCP.

    - Using the `hack/record-gcp` wrapper ensures a sufficient timeout (e.g., 30-60 minutes) is already configured, automatically handling slow GCP resource creation. There is no need to specify additional timeout flags when using this helper.

8.  **Validation & Last-Mile Tests**:
    Run the following tests to ensure CI compliance and verify field coverage:
    - **Fuzzing**: `dev/ci/presubmits/fuzz-roundtrippers`
    - **E2E Scaffolding**: `dev/ci/presubmits/tests-e2e-fixtures-direct`
    - **Schema Integrity**: `go test ./pkg/crd/template/...`
    - **API Field Coverage**: `go test ./tests/apichecks/...`. 
      - For alpha: `WRITE_GOLDEN_OUTPUT=1 go test -v ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha`
      - Verify that your "Maximal" test reduces the number of missing fields in the exceptions file. If `TestCRDFieldPresenceInTestsForAlpha` fails, running with `WRITE_GOLDEN_OUTPUT=1` will regenerate the exceptions file.
    - **Iterative Refinement Loop (Steps 7 & 8)**: Use any failures, unexpected diffs, or missing field coverage found during steps 7 and 8 as your critical debugging feedback loop. Actively refine your controller logic (Step 2), mappers (Step 3), fuzzer (Step 4), and fixtures (Steps 5 & 6), re-running steps 7 and 8 until all E2E recordings and validation suites execute and pass without error before proceeding to step 9.

9.  **Final Generation & Reporting**:
    Run `make ready-pr` from the repository root. This is a critical step that:
    - Runs `make fmt` and `make vet`.
    YOU MUST COMMIT ALL RESULTING CHANGES.

    **GCP RECORDING PR REPORTING REQUIREMENT**:
    In the created PR description, you MUST explicitly document that `record-gcp` was successfully run against real GCP and state the GCP project used. Mock logs (`_http_mock.log`) and mock fallback modes are strictly prohibited for greenfield PRs.

    **AUDIT LOG REPORTING REQUIREMENT**:
    As part of the final verification of development testing, you MUST retrieve Google Cloud audit logs demonstrating that the resource API was exercised.
    Use the `gcloud logging read` command to fetch the relevant audit entries to prove that the controller successfully hit the real GCP endpoint.
    
    For example:
    ```bash
    gcloud logging read 'logName:"audit" AND "<service_name>"' --freshness=3h --limit=50
    ```
    *(Adjust the search terms, e.g., replacing `<service_name>` with the actual service/API name such as `discoveryengine` or `vertexai` to filter for correct audit entries.)*
    After creating the PR, add these retrieved audit logs as a separate comment on the PR to serve as proof of real-GCP testing.

## Journaling
Append any reconciliation hurdles, GCP SDK quirks, or other controller issues to `.gemini/journals/<service_name>.md` using the format described in the `kcc-agentic-journaler` skill.

## Review Instructions
Use .gemini/skills/reviewgen-greenfield-controller/SKILL.md to conduct the review.
