---
name: Field and Resource Addition Playbook
description: Playbook for adding fields and resources across Direct, Terraform, DCL, and Dual-Controller archetypes.
schedule: "never"
skipPR: true
---

<!--
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Field and Resource Addition Playbook

This playbook defines the required workflow for autonomous agents and contributors adding fields or resources to Kubernetes Config Connector (KCC) across all controller archetypes: **Direct Controllers**, **Terraform-Based Controllers**, **DCL-Based Controllers**, and **Dual-Controller Resources**.

---

## 1. Controller Archetype Triage Matrix (Mandatory Step 1)

Before inspecting code or proposing changes, you **MUST** identify the target controller archetype:

| Check | Target / Location | Controller Archetype | Development Playbook |
| :--- | :--- | :--- | :--- |
| **Check 1: Static Config** | Check `{Group, Kind}` in `pkg/controller/resourceconfig/static_config.go` | **Direct** (if `DefaultController: Direct`)<br>**Dual** (if `SupportedControllers` contains both `Direct` and `Terraform`/`DCL`)<br>**Terraform** (if `DefaultController: Terraform`)<br>**DCL** (if `DefaultController: DCL`) | See respective section below. |
| **Check 2: CRD Labels** | Inspect CRD YAML under `config/crds/resources/` for:<br>• `cnrm.cloud.google.com/tf2crd: "true"`<br>• `cnrm.cloud.google.com/dcl2crd: "true"` | **Terraform-Based** if `tf2crd=true`<br>**DCL-Based** if `dcl2crd=true`<br>**Direct** if neither is present | Confirms runtime engine if `static_config.go` is not yet regenerated. |
| **Check 3: Greenfield Resource** | New GCP API not yet represented in KCC | **Direct Controller** (Mandatory default for all new resources) | Greenfield 3-Phase Multi-PR Pipeline. |

---

## 2. Direct Controller Flow

### 2.1 API Types (`apis/<service>/<version>/`)
- Add fields to `Spec` or `ObservedState` in `<resource>_types.go`.
- Use structured reference types (e.g. `*refs.<Kind>Ref`) for references to other GCP resources.
- Use `// +kcc:proto:field=<GCP_PROTO_PATH>` annotations on all fields.

### 2.2 Mappers and Fuzzers (`pkg/controller/direct/<service>/`)
- Register the field in `<resource>_fuzzer.go` using `f.SpecField(...)` or `f.StatusField(...)`. Remove from `f.UnimplementedFields` or `f.Unimplemented_NotYetTriaged`.
- Update handwritten mappers symmetrically in `<resource>_mappings.go` or `mapper.go`. Never edit `mapper.generated.go`.

### 2.3 Reference Resolution Timing Invariant
- **Rule**: Always resolve external references (`*refs.<Kind>Ref` or `common.NormalizeReferences`) on **line 1 of `Create()` and `Update()`** in `pkg/controller/direct/<service>/<resource>_controller.go` (or in a `buildDesired()` helper invoked strictly from `Create()` and `Update()`).
- **NEVER resolve references in `AdapterForObject()`**.
- **Rationale (Preventing Deletion Deadlocks per PR #13573)**:
  `AdapterForObject()` runs during controller initialization for every operation, including `Delete()`. If a referenced dependency is deleted first, resolving references in `AdapterForObject()` fails with a reference resolution error, preventing `Delete()` from running and permanently trapping the resource with its finalizer intact.

### 2.4 Testing and Ratcheting
- Remove the resource from the ratcheting exclusion list in `tests/e2e/ratcheting.go` (`ShouldTestRereconiliation`).
- Run focused fuzzer checks:
  ```bash
  FOCUS=<Kind> go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```
- Re-record golden fixtures:
  ```bash
  WRITE_GOLDEN_OUTPUT=1 hack/compare-mock pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/[testname]/
  ```

---

## 3. Terraform-Based Controller Flow

### 3.1 TPG Beta Subtree Patches
- Location: `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/<service>/resource_<service>_<resource>.go`.
- Add fields to schema with appropriate type, `Optional`, `Computed`, and `ForceNew` properties.
- Implement expanders (Terraform map -> GCP SDK client struct) and flatteners (GCP SDK response -> Terraform map).
- **The Empty Block Flattener Invariant (`ForceNew: true`)**: For optional nested blocks marked `ForceNew: true`, if all attributes within the block are empty in GCP's API response, the flattener **must return `nil` or empty slice**, not a map with zero-values. Returning a zero-value map causes Terraform to compute a permanent diff, triggering endless resource recreation cycles in KCC.

### 3.2 ServiceMapping Configuration
- In `config/servicemappings/<service>.yaml`:
  - Remove new fields from `ignoredFields`.
  - Declare reference fields under `resourceReferences`.
  - Declare server-computed fields under `observedFields`.

### 3.3 Manifest Generation and Verification
```bash
make manifests
make resource-docs
go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
make ready-pr
```

---

## 4. DCL-Based Controller Flow

1. Update DCL dependency via `make upgrade-dcl` or patch under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/<service>/` across all supported API versions.
2. Regenerate CRDs via `make ready-pr` and verify `git diff config/crds/resources/`.

---

## 5. Dual-Controller Synchronization (Mandatory)

For resources configured with dual controllers (`SupportedControllers` includes both `Direct` and `Terraform`/`DCL`):
- **You MUST update BOTH engines**:
  1. The legacy Terraform/DCL patch, flatteners, expanders, and ServiceMapping.
  2. The Direct KRM types, mappers, fuzzers, and controller.
- Dual-controller resources run both reconcilers in parallel in CI test suites. Updating only one engine causes schema drift and fixture diff failures.

---

## 6. Brownfield Migration: 100% Field Parity Invariant

When migrating an existing Terraform or DCL resource to a Direct Controller:
- **Mandatory 100% Field Parity**: The Direct Controller must support **every single field** present in the legacy CRD.
- **Consequence of Dropped Fields**: Dropping any field causes KCC to wipe that configuration from GCP on the very first reconcile loop after upgrade, leading to severe data loss or outage.
- API types, mappers, and fuzzers must cover 100% of the existing CRD fields before cutover.

---

## 7. Multi-PR Staging Protocol

- **Small Changes (< 400 lines)**: Deliver as a single atomic PR following the 4-commit pattern.
- **Greenfield / Large Changes (> 400 lines)**: Deliver via the Greenfield 3-Phase Multi-PR Pipeline:
  - Phase 1: API Types, CRDs, Identity & References (`Part of #<id>`, `release-note: NONE`).
  - Phase 2: Direct Controller, Mappers, Registration & Fuzzer (`Part of #<id>`, `release-note: NONE`).
  - Phase 3: Resource Fixtures, MockGCP HTTP Log Alignment & Docs (`Fixes #<id>`, customer release note).

---

## 8. SOUL.md Invariant 7 Compliance

- **No Internal Buganizer IDs**: Agents must **NEVER** write `b/<id>` or Buganizer URLs anywhere on GitHub (issue bodies, PR titles, descriptions, comments, commit messages, code comments, or release notes).
- **No Customer Names**: Agents must **NEVER** write customer names or private account details.
- Reference public GitHub issues and PRs only (`Fixes #<id>` or `Part of #<id>`).
