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

#### Reference Field Handling & Resolution
If the new field represents a GCP URI or URI fragment pointing to another GCP resource (e.g., a service account, a KMS key, a network), follow these guidelines instead of using a raw `*string`:
- **Use Reference Types:** Prefer a KCC reference type (e.g., `*refs.<Kind>Ref`) over a raw `*string`. For example, use `KMSKeyRef *refs.KMSCryptoKeyRef` instead of `KmsKey *string`.
- **Acronym Naming:** Capitalize acronyms in field names (e.g., use `KMSKeyRef *refs.KMSCryptoKeyRef` rather than `KmsKeyRef`) so that the code generator can automatically detect, parse, and map the reference.
- **Reference Annotations:** Always add the `// +kcc:proto:field=<GCP_PROTO_PATH>` annotation to reference fields (e.g. `// +kcc:proto:field=google.cloud.redis.cluster.v1.Cluster.kms_key`) so the generator knows exactly how to map it to the underlying protobuf schema.
- **Automatic vs. Manual Reference Resolution:**
  - If the controller calls `common.NormalizeReferences(ctx, reader, obj, nil)` in `AdapterForObject` (before converting KRM to Proto), reference resolution is handled automatically by the direct controller framework—no manual `refs.Resolve...` call is needed in the controller or mapper logic.
  - Manual `refs.Resolve...` calls (e.g., `refs.ResolveKMSCryptoKeyRef`) inside the controller are only required if the controller does not employ `common.NormalizeReferences` in its reconciliation loop.

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

   **Handwritten & Generated Mapper Nuances:**
   - **Root Type Mapper Bypass:** If the root `Spec` or `ObservedState` has a handwritten mapper (e.g., custom `_FromProto` and `_ToProto` functions), the generator will skip auto-mapping any new top-level or nested fields. You must manually map the new fields within these custom mapping functions.
   - **Cascading Skipped Mappers (`// MISSING: <Type>`):** When the generator detects `/* found existing non-generated mapping function ... */`, it skips generating child mappers. If the generator skipped a parent mapper, it may also skip nested child types, leaving comments like `// MISSING: <Type>` in `mapper.generated.go`. You must write custom `_FromProto` and `_ToProto` functions for these missing nested types.
   - **Inspection Helper:** Note that passing `--include-skipped-output` to the generator script (e.g., `./apis/<service>/<version>/generate.sh` or through the global generator if supported) prints the skipped generated mapping code directly into `types.generated.go` and `mapper.generated.go`. This serves as an excellent reference/starting point for your handwritten mappers.
   - **Symmetrical Mappings:** Always ensure mapping between KRM and GCP is completely symmetrical. If `ToProto` maps a field, `FromProto` must also map it back to KRM, otherwise the fuzzer or re-reconciliation tests will report a drift.

2. **Handle Optional Zero-Value Fields (`ForceSendFields` vs. Proto3 Pointers):**
   - **Scope of `ForceSendFields` (Proto2 / REST Clients):** `ForceSendFields` (and `NullFields`) only applies to **Proto2 / Discovery-based REST SDKs** (such as Cloud SQL `google.golang.org/api/sql/v1beta4` or Compute) where Go struct fields are direct value types (like `bool`, `int64`) rather than pointers, and are tagged with `omitempty`. For these fields, if a user explicitly configures a zero-value (like `false` or `0`), standard Go JSON marshaling will omit it unless you add the field name to `ForceSendFields`.
     **Example:**
     ```go
     if in.SomeBoolean != nil {
         out.SomeBoolean = *in.SomeBoolean
         out.ForceSendFields = append(out.ForceSendFields, "SomeBoolean")
     }
     ```
   - **Proto3 / gRPC SDKs (Pointer-based Presence):** For **Proto3 / gRPC-based SDKs** (which represent the majority of direct controllers), field presence is natively modeled via Go pointers (e.g., `*bool`, `*int32`). Non-nil zero-values (e.g. pointer to `false` or `0`) are automatically serialized and sent over the wire by the SDK, so `ForceSendFields` is **NOT** needed or used in Proto3 clients.

3. **Equality / Diff Checks:** Locate the equality checking file (e.g., `pkg/controller/direct/<service>/<resource>_equality.go`). Update the manual diff logic (e.g., `DiffInstances` or similar) to compare the desired and actual states of the newly added field, appending differences to the `structuredreporting.Diff` struct.

### Step 5: Fuzzer Support
To ensure the fuzzer doesn't panic or complain about unmapped fields:
- Locate the fuzzer definition file (typically `<resource>_fuzzer.go` or `<resource>_legacy_fuzzer.go` under `pkg/controller/direct/<service>/`).
- Find the list of ignored/unimplemented fields (look for `f.UnimplementedFields` or `f.Unimplemented_NotYetTriaged`).
- **Promoting & Granular Implementation:**
  - Promote your newly supported field path by removing it from the unimplemented list, and instead registering it.
  - Use `f.SpecField(...)` for spec fields (e.g., `f.SpecField(".cross_instance_replication_config")`).
  - Use `f.StatusField(...)` for status or observedState fields.
  - **Fuzzer Granularity & Partial Support:** If you are introducing a complex struct but deferring support for some nested subfields to a later PR, you should specify partial fuzzer support by registering the parent field and specifically marking the deferred subfields as unimplemented using `f.Unimplemented_NotYetTriaged(".parent.subfield")` (e.g., `f.Unimplemented_NotYetTriaged(".cross_instance_replication_config.subfield")`) to avoid fuzzer failures.
- To verify that your mappers and fuzzer mapping logic are correct and round-trip successfully, run focused fuzzer tests using the `FOCUS` environment variable:
  ```bash
  FOCUS=MyResourceKind go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```

### Step 6: Test Fixture Coverage
Verify the end-to-end functionality of the new field:
- Locate or create E2E test fixtures under `pkg/test/resourcefixture/testdata/basic/<service>/...` (e.g., `pkg/test/resourcefixture/testdata/basic/sql/v1beta1/sqlinstance/`).
- Add the new field to an existing test (e.g., `create.yaml` and `update.yaml`) or create a dedicated sub-directory test fixture if it requires specific configuration.

### Step 7: Remove from Ratcheting Exclusions (MANDATORY)
- Before executing your tests, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which validates that KCC does not perform redundant updates.
- In `tests/e2e/ratcheting.go`, locate the `ShouldTestRereconiliation` function, find the `switch` statement checking `primaryResource.GroupVersionKind()`, and remove the `case` line matching your target resource GVK/Kind.

### Step 8: Recording & Aligning Mock with Real GCP (Skill Delegation)
To execute tests, record HTTP logs from real GCP, and align mock behavior:
- Rather than running manually or updating logs by hand, delegate to the **`match-mockgcp-with-realgcp`** developer skill.
- This involves running `hack/record-gcp` to record authentic traffic against real GCP, followed by `hack/compare-mock` to match the behavior under mockgcp.
- If mockgcp is missing support for the new field, implement minimal stub/mock behavior under `mockgcp/` to satisfy the test pipeline (refer to `mockgcp/GEMINI.md` for details).

### Step 9: Run Pre-PR Verification
Before committing and submitting your PR, you **MUST** run the validation suite to ensure that all generated files are up to date and formatting/vet constraints pass:
```bash
make ready-pr
```

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

- **Proto Pointer Binding (`usesPointersInProtoBinding`):** In proto3, fields declared as `optional string` map to `*string` in the generated Go SDK. Since KCC's generator has a hardcoded list of services that use pointers (`usesPointersInProtoBinding`), if the service you are updating is not on that list, the generator will attempt to map the reference field as a direct string value, emitting a compiled assignment like `out.KmsKey = in.KMSKeyRef.External`. This fails compilation because the proto field expects a pointer `*string`. To resolve this, you must write a custom handwritten spec mapper (`Spec_ToProto` and `Spec_FromProto`) inside `<resource>_mapper.go` or similar to assign the pointer correctly (e.g., `out.KmsKey = &in.KMSKeyRef.External`).
- **Missing Symmetrical Mappings:** Always ensure mapping between KRM and GCP is completely symmetrical. If `ToProto` maps a field, `FromProto` must also map it back to KRM, otherwise the fuzzer or re-reconciliation tests will report a drift.
- **Ratcheting/Re-reconciliation failure:** If the resource undergoes infinite reconcile loops during tests, verify that `ForceSendFields` is correctly set, or that the equality checks in `<resource>_equality.go` are ignoring empty/default differences that the GCP API may return.
