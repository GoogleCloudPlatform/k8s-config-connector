# AI Agent Playbook: Field and Resource Additions Across Controller Archetypes

This document governs autonomous coding agents and human contributors adding fields or resources to Config Connector.

---

## 1. Controller Archetype Triage Matrix

Before planning or executing code modifications, run the following triage sequence:

| Step | Action | Controller Archetype | Action Path |
| :--- | :--- | :--- | :--- |
| **1** | Inspect `pkg/controller/resourceconfig/static_config.go` for `{Group, Kind}` | **Direct** (`DefaultController: Direct`) | Direct controller workflow |
| | | **Dual** (`SupportedControllers` has `Direct` + `TF`/`DCL`) | Synchronized updates to both legacy patch and direct files |
| | | **Terraform** (`DefaultController: Terraform`) | Terraform provider patch workflow |
| | | **DCL** (`DefaultController: DCL`) | DCL library upgrade/patch workflow |
| **2** | Check CRD labels in `config/crds/resources/` | `tf2crd: "true"` -> Terraform<br>`dcl2crd: "true"` -> DCL<br>Neither -> Direct | Runtime controller confirmation |
| **3** | Greenfield Resource (New GCP Service/Resource) | **Direct Controller** | Mandatory default: follow 3-phase greenfield pipeline |

---

## 2. Architectural Invariants

1. **Zero Defaulting Invariant**: Mappers perform strict 1:1 translation between KRM and GCP proto types. Never inject defaults in conversion mappers.
2. **KRM Source of Truth & Two-Payload Pattern**: Preserve user intent using `rawDesiredPb` vs `mergedDesiredPb` via `buildDesired()`.
3. **Structured References**: Use `*refs.<Kind>Ref` following the 3-tier contract (`name`, `namespace`, `external`).
4. **Reference Resolution Timing**:
   - Resolve references on line 1 of `Create()` and `Update()` in `_controller.go`.
   - **NEVER resolve references in `AdapterForObject()`**. `Delete()` must never resolve external references to avoid deletion deadlocks when dependencies are deleted first (PR #13573).
5. **Zero Test Hacks in MockGCP**: Implement genuine GCP REST/gRPC behavior in `mockgcp/` without synthetic shortcuts.
6. **0-Op Steady-State Reconciliation**: A reconciled resource must produce zero diff and zero API mutation calls on subsequent reconcile loops.
7. **100% Field Parity on Migrations**: When migrating a TF or DCL resource to Direct, the direct types must support **every single field** present in the legacy CRD. Dropping fields causes KCC to wipe that configuration in GCP on the first post-migration reconcile.
8. **SOUL.md Invariant 7 (Public Issue & PR Hygiene)**:
   - The bugs should have the Issues, but the issues/PRs should never show the bug ID or customer names.
   - Agents must never write `b/<id>`, Buganizer URLs, or customer names in issues, PR titles, PR descriptions, comments, commit messages, code comments, test fixtures, or release notes.
   - Add the public GitHub issue/PR URL to the internal Buganizer bug instead.

---

## 3. Fast Paths by Archetype

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
make vet
make manifests
make resource-docs
go test -v -tags=integration ./pkg/test/resourcefixture/testdata/basic/<service>/...
make ready-pr
```

### DCL-Based Controllers
```bash
make upgrade-dcl
make ready-pr
git diff config/crds/resources/
```

---

## 4. Multi-PR Staging Protocol

- **Small changes (<400 lines)**: Single PR with atomic commits.
- **Large changes / Greenfield (>400 lines)**:
  - Phase 1: API Types, CRDs, Identity & References (`Part of #<id>`, `release-note: NONE`)
  - Phase 2: Direct Controller, Mappers, Registration & Fuzzer (`Part of #<id>`, `release-note: NONE`)
  - Phase 3: Resource Fixtures, MockGCP HTTP Log Alignment & Docs (`Fixes #<id>`, customer release note)
