---
name: add-missing-field
description: Guides adding missing fields already supported in GCP API but not yet supported in KCC CRD to Config Connector resources.
---

# Skill: Add Missing Field

This skill guides an automated agent or contributor through adding a missing field to a GCP resource managed by KCC using the "direct" controller approach. A "missing field" is a field already supported/exposed in the underlying GCP service/API, but not yet exposed in KCC's KRM CRD and resource types.

---

## 1. End-to-End Workflow

Follow these steps to implement and verify a new field in a direct KCC controller:

### Step 1: Proto Identification
- Locate and identify the protobuf field definition of the target GCP service.
- Search within the repository's local copy of Google APIs under `./.build/third_party/googleapis/`.
  - For example, if you are adding DNS endpoints support to `ContainerCluster`, locate the `.proto` file (e.g., `.build/third_party/googleapis/google/container/v1/cluster_service.proto`) and identify the exact protobuf field definition:
    ```protobuf
    DNSEndpointConfig dns_endpoint_config = 1;
    ```
- Take note of the exact protobuf field name (usually `snake_case`) and its fully qualified path (e.g., `google.cloud.apigateway.v1.Api.state`).

### Step 2: KRM API Type Scaffolding
- Locate the Go types file for the resource under `apis/<service>/<version>/<message>_types.go` or similar (e.g., `apis/apigateway/v1alpha1/apigatewayapi_types.go`).
- Add the corresponding field to the appropriate Go struct (`Spec` for configuration inputs, or `ObservedState` / `Status` for status/output-only fields) following KRM camelCase conventions.
- Make sure to write proper Go doc comments.
- **Annotations Required:**
  1. Add the KRM optional annotation: `// +optional`.
  2. Add the KCC proto mapping annotation to allow automated mappings to resolve the exact proto schema mapping: `// +kcc:proto:field=<GCP_PROTO_PATH>` (e.g., `// +kcc:proto:field=google.cloud.apigateway.v1.Api.state`).
- **Example:**
  ```go
  // The state of the API.
  // +optional
  // +kcc:proto:field=google.cloud.apigateway.v1.Api.state
  State *string `json:"state,omitempty"`
  ```

### Step 3: Code & CRD Regeneration
- Run the centralized generation script to rebuild the Custom Resource Definitions (CRDs) and regenerate deepcopy helpers:
  ```bash
  dev/tasks/generate-types-and-mappers
  ```
- Alternatively, you can run the service-specific generator if one is present (e.g., `./apis/<service>/<version>/generate.sh`).
- Confirm that the newly added field appears in the generated `types.generated.go` or `config/crds/resources/...yaml` files.

### Step 4: Direct Controller Support
Update the reconciliation logic for the direct resource:
1. **Mappings:** Locate the conversion/mappings file (typically under `pkg/controller/direct/<service>/<resource>_mappings.go`, `mapper.go`, or similar). Update both `KRMToGCP` (desired state mapping) and `GCPToKRM` (observed/actual state mapping) functions to map the new field symmetrically.
2. **Handle Optional Zero-Value Fields (`ForceSendFields`):**
   - For optional fields (especially booleans or numerical zero-values like `false` or `0`), if the user explicitly configures a zero-value, standard Go JSON marshaling might omit it if annotated with `omitempty`.
   - To prevent this, ensure you configure `ForceSendFields` on the GCP API payload struct.
   - **Example:**
     ```go
     if in.SomeBoolean != nil {
         out.SomeBoolean = in.SomeBoolean
         out.ForceSendFields = append(out.ForceSendFields, "SomeBoolean")
     }
     ```
3. **Equality / Diff Checks:** Locate the equality checking file (e.g., `pkg/controller/direct/<service>/<resource>_equality.go`). Update the manual diff logic (e.g., `DiffInstances` or similar) to compare the desired and actual states of the newly added field, appending differences to the `structuredreporting.Diff` struct.

### Step 5: Fuzzer Support
To ensure the fuzzer doesn't panic or complain about unmapped fields:
- Locate the fuzzer definition file (typically `<resource>_fuzzer.go` or `<resource>_legacy_fuzzer.go` under `pkg/controller/direct/<service>/`).
- Find the list of ignored/unimplemented fields (look for `f.UnimplementedFields` or `f.Unimplemented_NotYetTriaged`).
- Promote your newly supported field path by removing it from the unimplemented list, and instead registering it:
  - Use `f.SpecField(...)` for spec fields (e.g., `f.SpecField(".cross_instance_replication_config")`).
  - Use `f.StatusField(...)` for status or observedState fields.
- To verify that your mappers and fuzzer mapping logic are correct and round-trip successfully, run focused fuzzer tests:
  ```bash
  FOCUS=MyResourceKind go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```

### Step 6: Test Fixture Coverage
Verify the end-to-end functionality of the new field:
- Locate or create E2E test fixtures under `pkg/test/resourcefixture/testdata/basic/<service>/...` (e.g., `pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance/`).
- Add the new field to an existing test (e.g., `create.yaml` and `update.yaml`) or create a dedicated sub-directory test fixture if it requires specific configuration.
- To run and generate/update the golden HTTP log outputs and object files, run the comparison tool with `WRITE_GOLDEN_OUTPUT=1`:
  ```bash
  WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/[testname]/
  ```
- If mockgcp is missing support for the new field, implement minimal stub/mock behavior under `mockgcp/` to satisfy the test pipeline (refer to `mockgcp/GEMINI.md` for details).

### Step 7: Remove from Ratcheting Exclusions (MANDATORY)
- Before executing your tests, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which validates that KCC does not perform redundant updates.
- In `tests/e2e/ratcheting.go`, locate the `ShouldTestRereconiliation` function, find the `switch` statement checking `primaryResource.GroupVersionKind()`, and remove the `case` line matching your target resource GVK/Kind.

---

## 2. Backwards Compatibility & Defaulting Guidance

When supporting a field that was already active/supported in the underlying GCP service but newly exposed in the KCC CRD, you must strictly adhere to the following principles to avoid disrupting existing users:

### Backwards Compatibility for Existing Resources
- Adding a new field to a CRD must **never** break or modify existing, live GCP resources that were created prior to the field being exposed in the CRD.
- Pre-existing KRM manifests will not specify the new field, so the field will be parsed as `nil` (for pointer types). Your controller must handle this gracefully.

### Avoid Destructive Static Defaulting
- Do **NOT** blindly default a newly supported field to a static default value when it is omitted from the KRM spec if that field is already set to some live value on GCP.
- Doing so can inadvertently overwrite the existing user configuration on GCP, trigger disruptive side-effects (e.g., unexpected read replica restarts, network configuration resets, or performance degradation), or cause infinite reconciliation loops.

### Preserve Live State (`actual` / `observed`)
- If a field is omitted from the KRM spec but already has an active, non-zero state on GCP (the `actual` live state), your controller **MUST preserve the live value during updates**.
- In your mapping or reconciliation update path, explicitly copy the value from the live/actual resource to the desired resource payload when the KRM spec does not specify it:
  ```go
  // Preserve live GCP state when KRM spec is unspecified
  if in.MyNewField == nil && actual.MyNewField != nil {
      out.MyNewField = actual.MyNewField
  }
  ```
- This ensures that KCC remains non-destructive and doesn't wipe out existing cloud configurations.

---

## 3. Troubleshooting & Common Pitfalls

- **Pointer mismatches in Protobuf Mapping:** In proto3, optional fields map to `*string` or other pointer types in the Go SDK. If the target GCP service is not in the generator's default pointer list (`usesPointersInProtoBinding`), you may need to write a custom, handwritten mapper function (`Spec_ToProto` and `Spec_FromProto`) to handle the pointer references correctly (e.g., `out.SomeField = &in.SomeField`).
- **Missing Symmetrical Mappings:** Always ensure mapping between KRM and GCP is completely symmetrical. If `ToProto` maps a field, `FromProto` must also map it back to KRM, otherwise the fuzzer or re-reconciliation tests will report a drift.
- **Ratcheting/Re-reconciliation failure:** If the resource undergoes infinite reconcile loops during tests, verify that `ForceSendFields` is correctly set, or that the equality checks in `<resource>_equality.go` are ignoring empty/default differences that the GCP API may return.
