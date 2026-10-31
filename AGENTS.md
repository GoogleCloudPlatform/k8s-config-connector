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

# KCC Universal Agent Guide (Autonomous Coding Harnesses)

This document provides root instructions, operational boundaries, and development workflows for autonomous coding harnesses (such as `antigravity`, `claude-code`, and Gemini CLI) working in Google Cloud Config Connector (`k8s-config-connector`).

## Authoritative Documentation

Before designing or modifying code, autonomous agents should review the core guides:
- [CONTRIBUTING.md](CONTRIBUTING.md) — Contributor agreement requirements, branching model, and PR hygiene.
- [GEMINI.md](GEMINI.md) — Architectural overview, controller routing precedence, and cluster/namespace execution modes.
- [docs/develop-resources/README.md](docs/develop-resources/README.md) — Scenario guides for Direct, Terraform, DCL, and Brownfield migrations.

---

## Controller Archetype Triage Matrix

Config Connector resources are powered by three distinct controller engines (or a dual-controller transitional state). **Autonomous agents must never blindly treat a resource as a Direct Controller.** Before executing commands or modifying code, execute this triage sequence:

| Archetype | Description | Primary Code Path |
| :--- | :--- | :--- |
| **Direct Controller** | Modern Go SDK reconcilers with MockGCP hermetic test suites | `apis/<service>/`, `pkg/controller/direct/<service>/` |
| **Terraform-Based** | Vendored Terraform Google Beta Provider (`tf2crd`) reconcilers (~70% legacy) | `third_party/github.com/hashicorp/terraform-provider-google-beta/`, `config/servicemappings/` |
| **DCL-Based** | Declarative Resource Client Library (`dcl2crd`) reconcilers (~20% legacy) | `pkg/dcl/`, upstream DCL library |
| **Dual-Controller** | Migrating resources supporting both Direct and legacy engines | Both paths above (must be updated simultaneously) |

### Deterministic Triage Sequence

1. **Step 1: Inspect `pkg/controller/resourceconfig/static_config.go`**
   Search for `{Group: "<service>.cnrm.cloud.google.com", Kind: "<Kind>"}`:
   - If `DefaultController` is `Direct` (or `SupportedControllers` contains only `Direct`): Follow **Direct Controller** playbook.
   - If `DefaultController` is `Terraform`: Follow **Terraform-Based** playbook.
   - If `DefaultController` is `DCL`: Follow **DCL-Based** playbook.
   - If `SupportedControllers` contains both `Direct` and `Terraform` or `DCL`: Follow **Dual-Controller** playbook.

2. **Step 2: Inspect CRD YAML under `config/crds/resources/`**
   Check labels in `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<service>.cnrm.cloud.google.com.yaml`:
   - `cnrm.cloud.google.com/tf2crd: "true"`: Terraform-backed CRD.
   - `cnrm.cloud.google.com/dcl2crd: "true"`: DCL-backed CRD.
   - Neither label present (types defined in `apis/<service>/`): Direct Controller CRD.

3. **Step 3: Greenfield (New Resource)**
   - New GCP API not yet present in KCC: Always implement as a **Direct Controller** (`apis/<service>/`, `pkg/controller/direct/<service>/`) using the Greenfield 3-Phase Multi-PR pipeline.

---

## Archetype CLI Fast Paths

### 1. Direct Controller Fast Path
Used for Greenfield resources and migrated Direct resources.

```bash
# 1. Regenerate KRM types and schema
apis/<service>/generate.sh
make manifests

# 2. Run unit and mapper tests
FOCUS=<Kind> go test ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers

# 3. Align MockGCP simulation with real GCP logs (takes test name or fixture path)
hack/compare-mock <test-name-or-fixture-path>

# 4. Verify all linters, formatting, and static configs
make ready-pr
```

### 2. Terraform-Based Fast Path (`tf2crd`)
Used for existing Terraform-backed resources.

```bash
# 1. Verify TPG beta compilation
make vet

# 2. Generate CRDs and docs from updated TPG schema and ServiceMapping
make manifests
make resource-docs

# 3. Run integration tests (via TestAllInSeries using mock or real GCP)
hack/compare-mock <test-name-or-fixture-path>

# 4. Verify PR readiness
make ready-pr
```

### 3. DCL-Based Fast Path (`dcl2crd`)
Used for existing DCL-backed resources.

```bash
# 1. Update DCL dependency and compile schemas
make upgrade-dcl

# 2. Verify PR readiness and static configuration
make ready-pr

# 3. Inspect generated CRD diffs
git diff config/crds/resources/
```

### 4. Dual-Controller Resources
When modifying a resource with `SupportedControllers` containing both `Direct` and `Terraform`/`DCL`, you must update **both** engines:
- Update the legacy TPG/DCL patch, flattener/expander, and `config/servicemappings/`.
- Update the Direct KRM types in `apis/<service>/`, mappers in `pkg/controller/direct/<service>/`, and fuzzers.
- Run both integration and mapper test suites.

---

## Mandatory Direct Controller Registration

When introducing a new Direct Controller package under `pkg/controller/direct/<service>/`, you **must** register it in:
`pkg/controller/direct/register/register.go`

Add an anonymous blank import:
```go
import (
    ...
    _ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/<service>"
)
```

**Why**: Direct controllers self-register their GVK models via `init()` invoking `RegisterModel`. Without this blank import, the controller package is omitted by the Go linker and the controller will never reconcile instances.

---

## Architectural Invariants

All agents must strictly preserve the following architectural invariants:

1. **Zero Defaulting in Conversion Mappers**: Conversion mappers must only translate values explicitly set in KRM or protobuf payloads. Never inject API or SDK defaults in mapper functions.
2. **KRM Source of Truth Preservation (Two-Payload Pattern)**: Preserve user-specified intent (`rawDesiredPb`) separately from the server-merged state (`mergedDesiredPb`). Server defaults must not pollute the desired specification or trigger unnecessary mutations.
3. **Structured References**: All resource references must use canonical reference structs (`*refs.<Kind>Ref` with `name`, `namespace`, and `external`), never plain string resource names. References must be resolved in `Create()` / `Update()`, never in `AdapterForObject()` (to avoid deletion deadlocks).
4. **Zero Test Hacks in MockGCP**: Mock services (`mockgcp/mock<service>/`) must faithfully simulate real GCP API behaviors, error codes, and field semantics. Never add test-only bypasses or shortcuts.
5. **0-Op Steady-State GitOps Reconciliation**: Once a resource reaches `Ready` state, subsequent reconcile loops must produce zero mutations and zero GCP API update/patch operations (zero drift).
6. **Backward Compatibility & Absolute Optionality**: No breaking schema changes are permitted on released API versions (v1alpha1, v1beta1). All new fields added to existing CRDs must be optional (`omitempty`).
7. **Confidentiality & One-Way Tracker Linking**:
   - Public GitHub issues, PRs, comments, commit messages, code comments, fixtures, and release notes must **NEVER** contain Buganizer IDs (`b/<id>`), Buganizer URLs, or internal customer names.
   - Always link GitHub issues to the internal bug instead of exposing internal trackers publicly.

---

## PR Routing & Staging Protocol

### PR Size and Commit Guidelines
- **Small Fields / Bugfixes (<400 lines)**: Single PR using the **Atomic 4-Commit Pattern**:
  1. Commit 1: Types & CRD generation (`apis/`, `config/crds/resources/`)
  2. Commit 2: Mappers & Fuzzer
  3. Commit 3: Controller & Registration (`register.go`)
  4. Commit 4: Tests, Fixtures, MockGCP alignment logs, and Docs
- **Greenfield Resources / Large Migrations (>400 lines)**: Use the **Greenfield 3-Phase Multi-PR Pipeline**:
  - **Phase 1 PR**: API Types, CRDs, Identity & References (`apis/<service>/<version>/`).
  - **Phase 2 PR**: Direct Controller, Mappers, Registration (`register.go`), and Round-Trip Fuzzer.
  - **Phase 3 PR**: E2E Resource Fixtures, MockGCP Implementation, HTTP Log Alignment (`_http.log`), and Documentation.

### Tracking Keywords and Release Notes
- **Intermediate PRs (Phase 1, Phase 2)**:
  - Header: `Part of #<issue_id>`
  - Do **NOT** use `Fixes #<issue_id>` or `Closes #<issue_id>` (prevents premature issue closing).
  - Release Note block:
    ```release-note
    NONE
    ```
- **Terminal PR (Phase 3 or single-PR additions)**:
  - Header: `Fixes #<issue_id>`
  - Release Note block: Customer-facing release note summarizing the new feature or fix.
