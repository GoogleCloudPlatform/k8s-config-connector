# Add Missing Field

This guide outlines how to add a missing field to a Config Connector resource across the three controller archetypes.

## Step 0: Controller Archetype Triage (MANDATORY)

Before modifying code, check `pkg/controller/resourceconfig/static_config.go` for `{Group, Kind}`:

1. **Direct Controller** (`DefaultController: Direct`):
   - Add field to `apis/<service>/<version>/<resource>_types.go`.
   - Update fuzzer in `pkg/controller/direct/<service>/<resource>_fuzzer.go`.
   - Update mappers in `pkg/controller/direct/<service>/<resource>_mappings.go`.
   - Regenerate: run `apis/<service>/generate.sh` and `make manifests`.
   - Test mappers: `go test -v ./pkg/controller/direct/... -run TestFocusedMappers`.
   - Run MockGCP: `hack/compare-mock pkg/test/resourcefixture/testdata/basic/...`.
2. **Terraform-Based Controller** (`DefaultController: Terraform`):
   - Patch vendored TPG Beta schema, expander, and flattener in `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/`.
   - Remove field from `ignoredFields` (or declare under `observedFields`/`resourceReferences`) in `config/servicemappings/<service>.yaml`.
   - Regenerate: `make manifests` and `make resource-docs`.
   - Run integration tests: `go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...`.
3. **DCL-Based Controller** (`DefaultController: DCL`):
   - Upgrade DCL dependency via `make upgrade-dcl` (or patch schemas under `third_party/...`).
   - Run `make ready-pr` and verify `git diff config/crds/resources/`.
4. **Dual-Controller Resource** (`SupportedControllers` includes both `Direct` and `Terraform`/`DCL`):
   - Update both the legacy patch and direct mapper/fuzzer/type files.

---

## Direct Controller Workflow Details

### 1. Update Types
- Add the field to `Spec` (user-configurable) or `ObservedState` (output-only) in `apis/<service>/<version>/<resource>_types.go`.
- For references to GCP resources, use `*refs.<TargetKind>Ref`.
- Annotate reference fields with `// +kcc:proto:field=<GCP_PROTO_PATH>`.

### 2. Update Fuzzer
- In `pkg/controller/direct/<service>/<resource>_fuzzer.go`, remove the field from `UnimplementedFields` or `Unimplemented_NotYetTriaged` and register as `f.SpecField(...)` or `f.StatusField(...)`.

### 3. Update Mappers
- Implement symmetric mappings (`_ToProto` and `_FromProto`) in `<resource>_mappings.go`.
- **Zero Defaulting Invariant**: Never inject default values in mappers.

### 4. Reference Resolution Timing in Controllers
- **CRITICAL**: Resolve references on line 1 of `Create()` and `Update()` inside `pkg/controller/direct/<service>/<resource>_controller.go`.
- **NEVER resolve references in `AdapterForObject()`**: `Delete()` must never attempt to resolve external references. If a referenced dependency is deleted first, resolving references in `AdapterForObject` blocks finalizer cleanup and causes an unrecoverable deletion deadlock (see PR #13573).

---

## Public Issue/PR Hygiene (SOUL.md Invariant 7)

- Never include Buganizer IDs (`b/<id>`), internal Buganizer URLs, or customer names in issue descriptions, PR titles, PR bodies, commit messages, code comments, test fixtures, or release notes.
- If tracking internal bugs, link the public GitHub issue/PR inside the Buganizer tracker instead.
