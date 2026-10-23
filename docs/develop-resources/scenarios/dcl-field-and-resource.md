# Playbook: DCL-Based Field and Resource Additions & Maintenance

> [!WARNING]
> **Prioritize the Direct Approach**: Config Connector is actively migrating away from legacy controllers (Terraform and DCL) towards the [Direct approach](../README.md). New resources and major new features should be implemented as Direct controllers whenever possible. Refer to [Migrate a Terraform/DCL-based Resource to Direct](./migrate-tf-resource-beta.md) for migration workflows.
> 
> This playbook documents the workflows for maintaining existing DCL-backed resources, modifying in-tree DCL code under `third_party/`, embedding DCL schemas, and verifying changes.

---

## 1. DCL In-Tree Architecture & Patch Flow

Config Connector maintains an in-tree fork of the DCL library located at:
```
third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/
```
In `go.mod`, a `replace` directive points `github.com/GoogleCloudPlatform/declarative-resource-client-library` to `./third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library`. Therefore, running `go get` does not affect the code being built. To introduce new fields, bug fixes, or behavioral adjustments to DCL-backed resources, you must patch the in-tree code under `third_party/`.

### Multi-Version Requirements for DCL Code Changes
When patching or adding a field in DCL directly under `third_party/`:

1. **All Supported API Versions Must Be Updated**:
   - DCL often provides multiple versions for a service (e.g., `services/google/<service>/`, `services/google/<service>/beta/`, and `services/google/<service>/alpha/`).
   - If a field or patch is modified in one version, it **must** be applied across all supported API versions of that resource to prevent regressions when bumping or reconciling versions.
2. **Components of a DCL Resource Patch**:
   When patching or adding a field in DCL:
   - **Schema Definition**: Update `[resource].yaml` to include the new field definition and OpenAPI validation types.
   - **Go Struct Types**: Update the resource Go struct in `[resource].go`.
   - **Expanding Logic**: Update expansion helpers in `[resource]_internal.go` (converting the Go struct into HTTP/JSON API requests).
   - **Flattening Logic**: Update flattening helpers in `[resource]_internal.go` (parsing HTTP/JSON API responses back into the Go struct).
   - **Diffing Logic**: Verify if the field triggers an update operation or requires custom difference handling (`[resource]_internal.go`).

### Step-by-Step Patch Workflow
1. Apply the edits across the schema (`.yaml`), structs, and conversion logic (`.go` and `_internal.go`) under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/<service>/`.
2. Regenerate embedded virtual filesystem assets and manifests:
   ```bash
   make generate-including-dcl
   make ready-pr
   ```
3. Inspect generated CRD diffs:
   ```bash
   git diff config/crds/resources/
   ```
   Confirm the intended changes appear in the CRD schemas and that no unintended breaking changes or schema regressions were introduced.

---

## 2. DCL Schema Architecture

Config Connector generates Kubernetes Custom Resource Definitions (CRDs) for DCL-based resources by parsing OpenAPI schemas exported by the DCL library.

### Schema Embedding Mechanism
1. **Source Schemas**: The in-tree DCL library ships OpenAPI v2/v3 schema files (`.yaml`) for each supported resource alongside Go service packages under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/<service>/`.
2. **Schema Embedding Generator (`pkg/dcl/schema/main/assets_generate.go`)**:
   - Uses `vfsgen` to embed OpenAPI schema `.yaml` files into Go code.
   - Filters schema assets based on registered metadata so that only schemas corresponding to supported services, versions, and resources are packaged.
   - Outputs the virtual filesystem to `pkg/dcl/schema/embed/dcl_assets_vfsdata.go`.
3. **Triggering Generation (`pkg/dcl/schema/embed/generate.go`)**:
   - Contains the directive `//go:generate go run -tags=integration ../main/assets_generate.go`.
   - Executed via `make generate-including-dcl`. Under the hood, this target temporarily stages DCL schemas into a scratch `temp-vendor/` directory, runs `go generate ./pkg/dcl/schema/...`, and cleans up.
4. **Schema Loader (`pkg/dcl/schema/dclschemaloader/dclschemaloader.go`)**:
   - Reads OpenAPI schemas at runtime or during CRD generation (`scripts/generate-crds/main.go`) from the embedded `embed.Assets` virtual filesystem.
   - Supplies schema models for CRD generation (`crdgeneration.New`) and runtime conversions (`pkg/dcl/conversion`).

### Metadata Registration (`pkg/dcl/metadata/metadata.go`)
Config Connector maintains a centralized registry of all DCL services and resources in `pkg/dcl/metadata/metadata.go`. Each service entry configures:
- `Name`: Canonical service name (e.g., `BinaryAuthorization`, `CloudScheduler`, `Compute`).
- `APIVersion`: The KCC KRM API version (typically `v1beta1` or `v1alpha1`).
- `DCLVersion`: Target DCL API version (e.g., `ga`, `beta`, or `alpha`).
- `Resources`: A list of registered resources under this service:
  - `Kind`: KRM Kind (e.g., `CloudSchedulerJob`).
  - `Releasable`: Boolean indicating whether the resource is officially supported and released in KCC. Resources set to `Releasable: false` are loaded for schema reference resolution only.
  - `SupportsHierarchicalReferences`: Flag indicating support for hierarchical parent references (such as `projectRef`, `folderRef`, or `organizationRef`).
  - `SupportsContainerAnnotations`: Flag indicating whether the resource supports hierarchical container annotations (`cnrm.cloud.google.com/project-id`, etc.).

When adding a new DCL-based resource or enabling a new version, the resource **must** be registered in `pkg/dcl/metadata/metadata.go` before schemas can be embedded and CRDs can be generated.

---

## 3. Upstream Sync and Dependency Upgrade

In rare cases where Config Connector's in-tree DCL fork is synced against the canonical upstream [declarative-resource-client-library](https://github.com/GoogleCloudPlatform/declarative-resource-client-library) repository:

1. Update the dependency version or merge upstream commits into `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/`.
2. Run `make upgrade-dcl` (which runs `go get github.com/GoogleCloudPlatform/declarative-resource-client-library` and `make ensure`).
3. Re-embed schemas and regenerate manifests:
   ```bash
   make generate-including-dcl
   make ready-pr
   ```
4. Review generated CRD diffs (`git diff config/crds/resources/`) and cross-reference them against upstream release notes.

---

## 4. Test Verification for DCL Changes

All changes to DCL-backed resources (field additions, behavior fixes, or schema updates) must be thoroughly verified using automated test fixtures.

### Running DCL Fixture Tests
Test fixtures are located under:
```
pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/
```
Run tests locally using the unified test runner:
```bash
# Run against MockGCP (default hermetic mock environment)
RUN_E2E=1 E2E_GCP_TARGET=mock go test ./tests/e2e -run 'TestAllInSeries/fixtures/<testname>'
```

### Canonical Test Workflow: `hack/record-gcp` and `hack/compare-mock`
The canonical scripts for running fixture tests and aligning golden HTTP logs are:
- **`hack/record-gcp`**: Records golden HTTP traffic (`_http.log`) and golden KRM object state against real GCP.
- **`hack/compare-mock`**: Runs tests against MockGCP and checks for behavioral differences against golden logs.

#### Recording Golden HTTP Logs Against Real GCP
To run a fixture test against real GCP and record baseline golden files:
```bash
hack/record-gcp "fixtures/^<testname>$"
```
This sets `RUN_E2E=1`, `E2E_GCP_TARGET=real`, and `WRITE_GOLDEN_OUTPUT=1`, writing `_http.log` and `_generated_object_<resource>.golden.yaml`.

#### Aligning MockGCP
To run the fixture test against MockGCP:
```bash
hack/compare-mock "fixtures/^<testname>$"
```
Inspect any differences in `git diff <fixture-path>`. If MockGCP differs from real GCP behavior:
1. Update MockGCP handlers under `mockgcp/` to match real GCP behavior.
2. If discrepancies are due to volatile fields (timestamps, dynamic tokens), update normalizers in `tests/e2e/normalize.go` or the service's mock controller.
3. Re-run `hack/compare-mock` until no unintended diffs remain.

For single-controller mock tests without running against real GCP, you can also record mock golden output directly:
```bash
RUN_E2E=1 E2E_GCP_TARGET=mock WRITE_GOLDEN_OUTPUT=1 go test ./tests/e2e -run 'TestAllInSeries/fixtures/<testname>'
```

### Coverage Requirements for PRs
1. **Spec Coverage**: The new or modified field must be tested in `create.yaml` and/or `update.yaml` (if the field is mutable).
2. **Golden Object Verification**: Ensure `_generated_object_<resource>.golden.yaml` contains the new field in status or spec as expected.
3. **No Unintended Regressions**: Existing tests for the resource must continue passing without unexpected diffs in `_http.log`.
