---
name: add-missing-field
description: Guides through adding a missing field to a GCP resource managed by KCC across all controller archetypes (Direct, Terraform, DCL, and Dual-Controller).
---

# Skill: Add Missing Field

This skill guides an automated agent through adding a missing field to a GCP resource managed by Config Connector (KCC).

## Pre-requisites
- Identify the target KCC API Group and Resource Kind.
- Identify the missing field from the issue description, GCP protobuf definitions, or fuzzer output (e.g. `f.Unimplemented_NotYetTriaged(".cross_instance_replication_config")`).

---

## Step 1: Controller Archetype Triage (MANDATORY FIRST STEP)

Before modifying any files, you **MUST** determine which controller engine manages the resource by inspecting `pkg/controller/resourceconfig/static_config.go` for the resource's `{Group, Kind}` entry under `ControllerConfigStatic`.

Also check CRD labels in `config/crds/resources/` as a secondary check:
- `cnrm.cloud.google.com/tf2crd: "true"` indicates a Terraform-based controller.
- `cnrm.cloud.google.com/dcl2crd: "true"` indicates a DCL-based controller.

### Routing Decision

1. **Dual-Controller Resources (`SupportedControllers` contains both `Direct` and `Terraform`/`DCL`)**:
   - **MANDATORY**: You **MUST** update **BOTH** the legacy patch (Terraform or DCL flow) **AND** the direct controller files (Direct flow).
   - Dual-controller resources run both reconcilers side-by-side during CI testing. Omitting updates from either side causes diff failures (`_http.diff`, `_final_object.diff`) or test breakages.
   - Follow the **Terraform Flow** (or **DCL Flow**) for legacy provider patches, AND follow the **Direct Flow** for KRM types, mappers, fuzzers, and controller updates.

2. **Direct Controller (`DefaultController: Direct` or pure Direct)**:
   - Follow the **Direct Flow** below (Steps 2–8).

3. **Terraform-Based Controller (`DefaultController: Terraform` without Direct in `SupportedControllers`)**:
   - Follow the **Terraform Flow** below.

4. **DCL-Based Controller (`DefaultController: DCL` without Direct in `SupportedControllers`)**:
   - Follow the **DCL Flow** below.

---

## Direct Flow

### 2. Add the field to the `_types.go` file
- Locate the `_types.go` file for the resource (e.g., `apis/<group>/<version>/<resource>_types.go`).
- Add the field to the `Spec` (for inputs) or `ObservedState` (for outputs) struct as appropriate.
- If the type relies on autogeneration, you can use `types.generated.go` as a reference. Note that `types.generated.go` includes generated code for all types if `--include-skipped-output` is passed to the generate commands.
- **Important:** If the field represents a GCP URI or URI fragment (e.g., the path to another resource), you should use a KCC reference field (e.g., `InstanceRef *refs.MemorystoreInstanceRef`) instead of a simple string.
- Follow KRM naming conventions (camelCase in Go, etc.).

### 3. Update the Fuzzer
- Find the fuzzer file (e.g., `pkg/controller/direct/<group>/<resource>_fuzzer.go`).
- Remove the missing field from `f.UnimplementedFields` or `f.Unimplemented_NotYetTriaged`.
- Register the field as a spec or status field in the fuzzer (e.g., `f.SpecField(".cross_instance_replication_config")` or `f.StatusField(".psc_attachment_details")`).
- If there are subfields you aren't implementing yet, use `f.Unimplemented_NotYetTriaged` to ignore them for now.

### 4. Update the Mappers
- If the resource uses hand-written mappers (e.g., `mapper.go` or `<resource>_mappings.go`), add the appropriate logic to convert between the Kubernetes resource (`KRM`) and the GCP API (`Proto/API`).
- Note that if the root `Spec` or `ObservedState` has a handwritten mapping function (e.g., `_FromProto` and `_ToProto`), you'll need to manually add the mapping for the new field there, even if it's just a top-level field.
- Check `mapper.generated.go` for blocks starting with `/* found existing non-generated mapping function ... */`. If the generator skipped a parent mapper, it may also skip nested types and leave comments like `// MISSING: <Type>`. You will need to manually write `_FromProto` and `_ToProto` functions for these missing nested types.
- **Important:** Always update the `ToProto` and `FromProto` methods "symmetrically", so that they round trip. If you map a field in `Spec_ToProto`, make sure to also map it in `Spec_FromProto`.
- Never edit `mapper.generated.go` directly as it is overwritten by code generators.

### 5. Run `generate.sh` and Verify Fuzzer
- Run the code generator script (e.g., `./apis/<group>/<version>/generate.sh` or the global `dev/tasks/generate-types-and-mappers`) to regenerate CRDs and other boilerplate.
- To quickly verify that your mappers and fuzzer mapping logic are correct and round-trip successfully, run focused fuzzer tests using the `FOCUS` environment variable:
  ```bash
  FOCUS=MyResourceKind go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```
  *(For example: `FOCUS=ComputeNetwork go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers`)*

### 6. Create or Update Tests
- Find existing tests under `pkg/test/resourcefixture/testdata/basic/<group>/<version>/<kind>/`.
- If the new field can be added to an existing test without conflict, do so.
- If the new field requires a specific setup (e.g., cross-instance replication requires two instances), create a new test folder with `create.yaml` (and `dependencies.yaml` if needed).

### 7. Remove from Ratcheting Exclusions (MANDATORY)
- Before running test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental invariant KCC resources must support.
  1. Open `tests/e2e/ratcheting.go`.
  2. Locate the function `ShouldTestRereconiliation`.
  3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
  4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

### 8. Regenerate Golden Output
- Run the mock compare script to generate new golden output (you must set `WRITE_GOLDEN_OUTPUT=1`):
  ```bash
  WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<group>/<version>/<kind>/[testname]/
  ```
- Note that test paths may use the full kind name (e.g., `memorystoreinstance` instead of `instance`).
- Review the generated `_http.log` and `_generated_object...golden.yaml` for correctness.
- If mockgcp doesn't support the field yet, implement the necessary mock handling in `mockgcp/mock<service>/`.

---

## Terraform Flow

For resources managed by the Terraform provider:

1. **Locate Schema and Implement Patch**:
   - Target files are located under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/resource_<service>_<resource>.go`.
   - Add the new field schema to the Terraform schema map (`map[string]*schema.Schema`) with appropriate type, `Optional`, `Computed`, and `ForceNew` properties.
   - Implement expander logic (Terraform map -> GCP SDK client struct).
   - Implement flattener logic (GCP SDK response -> Terraform map).
   - **The Empty Block Flattener Invariant (`ForceNew: true`)**: For optional nested blocks marked `ForceNew: true`, if all attributes within the block are empty in GCP's API response, the flattener **must return `nil` or empty slice**, not a map with zero-values. Returning a zero-value map causes Terraform to compute a permanent diff, triggering endless resource recreation cycles in KCC.

2. **Configure ServiceMapping (`config/servicemappings/<service>.yaml`)**:
   - Remove the field from `ignoredFields` if present.
   - If the field references another resource, declare it under `resourceReferences`.
   - If server-computed, declare it under `observedFields`.

3. **Generate Manifests and Verify**:
   ```bash
   make manifests
   make resource-docs
   go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
   make ready-pr
   ```

---

## DCL Flow

For resources managed by the Declarative Resource Client Library:

1. **Upgrade DCL or Apply Patch**:
   - Run `make upgrade-dcl` if the field is supported in newer upstream DCL releases.
   - If patching locally under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/`, apply schema and implementation changes across **all supported API versions** for that resource.

2. **Regenerate CRDs and Verify**:
   ```bash
   make ready-pr
   git diff config/crds/resources/
   ```

---

## Reference Mapping Tips & Timing Rules

- **Use capital KMS**: If you add a reference like `KMSKeyRef`, naming it all-caps `KMSKeyRef *refs.KMSCryptoKeyRef` instead of `KmsKeyRef` allows the controller builder to automatically identify it as a reference field and generate its mapping logic.
- **Annotate reference fields**: Always add the `// +kcc:proto:field=<GCP_PROTO_PATH>` annotation to reference fields in `_types.go` (e.g. `// +kcc:proto:field=google.cloud.redis.cluster.v1.Cluster.kms_key`) so the generator knows exactly how to map it to the underlying protobuf schema.
- **Pointers in Proto Binding**: In proto3, optional string fields generated by protoc map to `*string` Go pointers. Since KCC's generator has a hardcoded list of services that use pointers (`usesPointersInProtoBinding`), if the service you are updating is not on that list, the generator will attempt to map the reference field as a direct string value (resulting in `out.KmsKey = in.KMSKeyRef.External`, which fails to compile because it expects a `*string`). To resolve this, write a custom handwritten spec mapper (`Spec_ToProto` and `Spec_FromProto`) inside `cluster_mapper.go` or similar to handle the pointer assignment (`out.KmsKey = &in.KMSKeyRef.External`) correctly.

### Reference Resolution Timing (CRITICAL INVARIANT)

- **Strict Placement in `Create()` and `Update()`**: Always resolve external references (`*refs.<Kind>Ref` or `common.NormalizeReferences`) on **line 1 of `Create()` and `Update()`** inside `pkg/controller/direct/<service>/<resource>_controller.go` (or inside a `buildDesired()` helper invoked exclusively by `Create()` and `Update()`).
- **NEVER in `AdapterForObject()`**: You **MUST NEVER** resolve references inside `AdapterForObject()`.
- **Rationale (Preventing Deletion Deadlocks per PR #13573)**:
  During object deletion, the controller calls `AdapterForObject()` before invoking `Delete()`. If a referenced dependency has already been deleted, attempting to resolve external references inside `AdapterForObject()` fails with a reference resolution error. This blocks controller initialization, prevents `Delete()` from ever executing, and leaves the resource permanently wedged with its finalizer intact (unrecoverable deletion deadlock). `Delete()` must never attempt to resolve external references.

---

## SOUL.md Invariant 7 Compliance

When creating commits, PR descriptions, comments, or documentation:
- **No Internal Buganizer IDs**: **NEVER** include `b/<id>` or Buganizer URLs (such as `https://buganizer.corp.google.com/...`) anywhere in commit messages, PR descriptions, code comments, test fixtures, or release notes.
- **No Customer Names**: **NEVER** include customer names or private account identifiers.
- **GitHub Reference Linking**: Link strictly to public GitHub issues or pull requests (e.g., `Fixes #<id>` or `Part of #<id>`).
