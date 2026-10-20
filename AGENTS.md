# AI Agent Guidelines (AGENTS.md)

Welcome to Kubernetes Config Connector (KCC). This document provides direct entry-point instructions and architectural mandates for autonomous coding agents (including Claude Code, Antigravity, and Gemini).

---

## 1. Controller Archetype Triage Matrix (MANDATORY FIRST STEP)

Before writing or editing any code, determine which controller archetype manages the target resource by running this deterministic triage sequence:

| Check | Location / Command | Controller Archetype | Development Playbook |
| :--- | :--- | :--- | :--- |
| **Step 1: Check `static_config.go`** | Check `{Group, Kind}` in `pkg/controller/resourceconfig/static_config.go` | **Direct** (if `DefaultController: Direct`)<br>**Dual** (if `SupportedControllers` has `Direct` + `TF`/`DCL`)<br>**Terraform** (if `DefaultController: Terraform`)<br>**DCL** (if `DefaultController: DCL`) | See Archetype CLI Fast Paths below. |
| **Step 2: Inspect CRD Labels** | Check YAML under `config/crds/resources/` for:<br>• `cnrm.cloud.google.com/tf2crd: "true"`<br>• `cnrm.cloud.google.com/dcl2crd: "true"` | **Terraform-Based** if `tf2crd=true`<br>**DCL-Based** if `dcl2crd=true`<br>**Direct** if neither label is present | Confirms runtime engine if `static_config.go` is not yet updated. |
| **Step 3: Greenfield Resource** | New GCP API not yet represented in KCC | **Direct Controller** (Mandatory default) | Greenfield 3-Phase Multi-PR Pipeline. |

---

## 2. Archetype CLI Fast Paths

### Direct Controllers
```bash
apis/<service>/generate.sh
make manifests
go test -v ./pkg/controller/direct/... -run TestFocusedMappers
hack/compare-mock <service> <resource>
make ready-pr
```

### Terraform-Based Controllers
```bash
# Check TPG compiles
make vet
# Generate CRDs from updated TPG schema and ServiceMapping
make manifests
make resource-docs
# Run integration tests
go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
make ready-pr
```

### DCL-Based Controllers
```bash
make upgrade-dcl
make ready-pr
git diff config/crds/resources/
```

### Dual-Controller Resources
When modifying a resource with both Direct and Terraform/DCL enabled in `static_config.go`, you must update **both** the legacy patch and the direct mappers/types symmetrically.

---

## 3. Six Architectural Invariants

1. **Zero Defaulting in Conversion Mappers**: Mappers perform pure 1:1 translation between KRM and GCP proto types. Never inject default values in conversion mappers; defaulting belongs to GCP APIs or explicit defaulting webhooks.
2. **KRM Source of Truth Preservation (Two-Payload Pattern)**: Distinguish between `rawDesiredPb` (constructed purely from user-specified KRM `spec` via `buildDesired()`) and `mergedDesiredPb` (incorporating current server state `actualPb` to preserve unmanaged attributes).
3. **Structured References**: Use `*refs.<Kind>Ref` following the 3-tier contract (`name`, `namespace`, `external`).
4. **Reference Resolution Timing**:
   - Resolve references on line 1 of `Create()` and `Update()` in `_controller.go`.
   - **NEVER resolve references in `AdapterForObject()`**. `Delete()` must never attempt to resolve external references, preventing deletion deadlocks when dependencies are deleted first (PR #13573).
5. **Zero Test Hacks in MockGCP**: Implement authentic GCP API semantics in `mockgcp/mock<service>/` without synthetic shortcuts.
6. **0-Op Steady-State GitOps Reconciliation**: A reconciled resource must produce zero diff and zero API mutation calls on subsequent reconcile loops.
7. **100% Field Parity on Migrations**: When migrating a legacy resource to Direct, the direct types must support **every single field** present in the legacy CRD. Dropping a field causes KCC to wipe configuration from GCP.

---

## 4. PR Routing & Tracking Protocol

- **Small changes / bugfixes (<400 lines)**: Single PR via Atomic Commits.
- **Large changes / Greenfield (>400 lines)**: 3-Phase Multi-PR pipeline:
  - **Phase 1**: API Types, CRDs, Identity & References (`Part of #<id>`, `release-note: NONE`)
  - **Phase 2**: Direct Controller, Mappers, Registration in `register.go` & Fuzzer (`Part of #<id>`, `release-note: NONE`)
  - **Phase 3**: Resource Fixtures, MockGCP HTTP Log Alignment & Docs (`Fixes #<id>`, customer release note)
- **Keyword rules**:
  - Use `Part of #<id>` and ````release-note\nNONE\n```` on intermediate PRs.
  - Use `Fixes #<id>` and customer-facing release note on the terminal PR.

---

## 5. Public Issue & PR Hygiene (SOUL.md Invariant 7)

- The bugs should have the Issues, but the issues/PRs should never show the bug ID or customer names.
- Agents must **never** write `b/<id>`, Buganizer URLs, or customer names in issues, PR titles, PR descriptions, comments, commit messages, code comments, test fixtures, or release notes.
- Always add the public GitHub issue/PR URL to the internal Buganizer bug instead.
