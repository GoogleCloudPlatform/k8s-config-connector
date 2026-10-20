# DCL-Based Field and Resource Additions Playbook

This playbook provides guidance for maintaining and updating resources managed by the Declarative Resource Client Library (DCL) in Config Connector.

---

## 1. Controller Archetype Overview

Approximately 20% of legacy Config Connector resources are backed by the [declarative-resource-client-library (DCL)](https://github.com/GoogleCloudPlatform/declarative-resource-client-library).

A resource is DCL-based if:
- `pkg/controller/resourceconfig/static_config.go` specifies `DefaultController: DCL`, or
- The CRD schema under `config/crds/resources/` contains the label `cnrm.cloud.google.com/dcl2crd: "true"`.

> **Note**: For new features or resources, Direct Controllers are strongly preferred over DCL additions. However, when maintaining existing DCL resources or updating upstream schemas, follow this workflow.

---

## 2. DCL Dependency Upgrade Flow

To update DCL dependency or incorporate newly released upstream fields:

### Step 1: Upgrade DCL Dependency
```bash
make upgrade-dcl
```
This updates `github.com/GoogleCloudPlatform/declarative-resource-client-library` in `go.mod` and runs `make ensure`.

### Step 2: Regenerate CRDs and Static Configs
```bash
make ready-pr
```

### Step 3: Inspect Generated CRD Diffs
Inspect the generated changes in CRD definitions:
```bash
git diff config/crds/resources/
```
Verify that the intended fields are present and backwards compatibility is preserved.

---

## 3. DCL Schema Architecture

### Schema Embedding
DCL schemas are embedded as virtual assets into the binary:
1. `pkg/dcl/schema/main/assets_generate.go` reads DCL JSON schemas and generates `pkg/dcl/schema/embed/dcl_assets_vfsdata.go`.
2. The KCC dynamic DCL controller parses these schemas at runtime to dynamically construct Kubernetes CRDs and validate object states.

### Metadata Registration
DCL resource types are registered in `pkg/dcl/metadata/metadata.go`. This metadata configures:
- Group, Version, and Kind mappings.
- Specification of parent fields and resource reference resolvers.
- ID formatting and container type definitions (Project, Folder, Organization).

---

## 4. Multi-Version Schema Requirements

When modifying DCL patches directly under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/`:
- If multiple versions exist for the resource (e.g. `v1alpha`, `v1beta`, `ga`), changes **must be applied across all supported API versions** for that resource.
- Schema definitions in `[typename].yaml` as well as expanding and flattening logic in `[typename]_internal.go` must be updated symmetrically across versions.

---

## 5. Test Verification for DCL Changes

1. **Run DCL fixture integration tests**:
   ```bash
   go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
   ```
2. **Update Golden HTTP Logs**:
   - Compare real logs (`_http.log`) or mock logs (`_http_mock.log`) using `hack/compare-mock`.
   - Ensure the new fields are exercised in `create.yaml` and/or `update.yaml`.

---

## 6. Public Issue/PR Hygiene (SOUL.md Invariant 7)

Never mention Buganizer IDs (`b/<id>`), internal Buganizer URLs, or customer names in issue descriptions, PR titles, PR bodies, commit messages, code comments, test fixtures, or release notes.
