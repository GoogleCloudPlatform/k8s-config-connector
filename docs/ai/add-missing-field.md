# How to Add a Missing Field

This guide outlines the complete process for adding a missing field to a resource managed by Config Connector (KCC).

---

## 1. Controller Archetype Triage (Mandatory First Step)

Before editing any code, determine which controller archetype manages the resource by checking `pkg/controller/resourceconfig/static_config.go` for the resource's `{Group, Kind}` entry under `ControllerConfigStatic`.

Also check CRD YAML labels in `config/crds/resources/`:
- `cnrm.cloud.google.com/tf2crd: "true"` (Terraform-based)
- `cnrm.cloud.google.com/dcl2crd: "true"` (DCL-based)

### Triage Matrix

| Controller Archetype | Identification | Workflow |
| :--- | :--- | :--- |
| **Direct Controller** | `DefaultController: Direct` (or purely direct) | Update `apis/` types, mappers, fuzzer, and golden fixtures. |
| **Dual-Controller Resource** | `SupportedControllers` contains both `Direct` and `Terraform`/`DCL` | **MANDATORY**: Update **BOTH** the legacy patch and direct mappers/fuzzers/types. |
| **Terraform-Based Controller** | `DefaultController: Terraform` | Update TPG Beta schema/expanders/flatteners, and `config/servicemappings/<service>.yaml`. |
| **DCL-Based Controller** | `DefaultController: DCL` | Run `make upgrade-dcl` or update DCL schema across all supported versions. |

---

## 2. Direct Controller Flow

### 2.1 Update Go Types (`apis/`)
- Locate `apis/<service>/<version>/<resource>_types.go`.
- Add the missing field to `Spec` (user-configurable inputs) or `ObservedState` (server-computed outputs).
- Use reference types (e.g. `*refs.<Kind>Ref`) for fields that reference other GCP resources rather than plain strings. Add the structured tag comment `// +kcc:proto:field=<GCP_PROTO_PATH>`.

### 2.2 Update Mappers and Fuzzers
- Update `pkg/controller/direct/<service>/<resource>_fuzzer.go`: remove the field from `UnimplementedFields` or `Unimplemented_NotYetTriaged`, and register it via `f.SpecField(...)` or `f.StatusField(...)`.
- Update handwritten mappers (`_mapper.go` or `<resource>_mappings.go`) symmetrically (`Spec_ToProto` and `Spec_FromProto`). Do NOT edit `mapper.generated.go` directly.

### 2.3 Regenerate Manifests and Test
- Run `apis/<service>/<version>/generate.sh` or `dev/tasks/generate-types-and-mappers`.
- Run focused fuzzer validation:
  ```bash
  FOCUS=<Kind> go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```
- Remove the resource from the ratcheting exclusion list in `tests/e2e/ratcheting.go` (`ShouldTestRereconiliation`).
- Regenerate golden HTTP logs and KRM object YAML:
  ```bash
  WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/[testname]/
  ```

---

## 3. Terraform Provider Flow

1. **Modify TPG Beta Schema**:
   - Locate `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/resource_<service>_<resource>.go`.
   - Add field definition to `schema.Schema`.
   - Implement expander and flattener functions.
   - **The Empty Block Flattener Invariant (`ForceNew: true`)**: For optional nested blocks marked `ForceNew: true`, if all attributes within the block are empty in GCP's API response, the flattener **must return `nil` or empty slice**, not a map with zero-values. Returning a zero-value map causes Terraform to compute a permanent diff, triggering endless resource recreation cycles in KCC.

2. **Configure ServiceMapping**:
   - In `config/servicemappings/<service>.yaml`, remove the field from `ignoredFields`, and add reference or observed field declarations if appropriate.

3. **Regenerate and Verify**:
   ```bash
   make manifests
   make resource-docs
   go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
   make ready-pr
   ```

---

## 4. DCL Flow

1. Update DCL dependency via `make upgrade-dcl` or apply patches under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/<service>/` across all supported API versions.
2. Regenerate CRDs via `make ready-pr` and verify `git diff config/crds/resources/`.

---

## 5. Dual-Controller Synchronization (Mandatory Rule)

If a resource is configured with dual controllers (`SupportedControllers` includes both `Direct` and `Terraform` or `DCL`), you **MUST** update both engines:
1. Update the legacy Terraform/DCL patch and ServiceMapping.
2. Update the Direct KRM types, mappers, fuzzer, and controller.
Both reconcilers run side-by-side during CI tests; updating only one engine will result in test and diff validation failures.

---

## 6. Reference Resolution Timing (CRITICAL INVARIANT)

When implementing or updating reference resolution logic:
- **Where to resolve**: Reference resolution must take place on **line 1 of `Create()` and `Update()`** inside `pkg/controller/direct/<service>/<resource>_controller.go` (or inside a `buildDesired()` method invoked strictly by `Create()` and `Update()`).
- **NEVER in `AdapterForObject()`**: Resolving references (`*refs.<Kind>Ref` or `common.NormalizeReferences`) inside `AdapterForObject()` is strictly prohibited.
- **Rationale (Preventing Deletion Deadlocks per PR #13573)**:
  `AdapterForObject()` is called during controller setup for all operations, including `Delete()`. If a referenced dependency is deleted before the primary resource, calling reference resolution inside `AdapterForObject()` fails with a reference resolution error. This aborts controller initialization, preventing `Delete()` from running and permanently trapping the resource with its finalizer intact (unrecoverable deletion deadlock). `Delete()` must never attempt to resolve external references.

---

## 7. SOUL.md Invariant 7 Compliance

- **No Internal Trackers**: Never include Buganizer IDs (`b/<id>`), internal bug links (`buganizer.corp.google.com`), or customer names in commit messages, PR titles, descriptions, code comments, or release notes.
- **Public GitHub Links Only**: Link issues using GitHub syntax (`Fixes #<id>` or `Part of #<id>`).
