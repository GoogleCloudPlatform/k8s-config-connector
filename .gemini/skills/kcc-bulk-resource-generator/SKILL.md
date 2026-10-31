---
name: kcc-bulk-resource-generator
description: Orchestrator workflow for bulk-generating KRM types, identities, references, mappers, CRDs, and judgement queues across multiple greenfield GCP resources, splitting them into PR categories, and verifying each PR via an isolated subagent.
---

# KCC Bulk Resource Generator (Orchestrator)

This skill orchestrates the two-phase workflow for bulk-scaffolding multiple greenfield KCC resources (`--prepopulate-resource`), splitting the output into three reviewable PR categories, and verifying each category branch in an isolated subagent context.

---

## Phase 1: Candidate Sourcing, Bulk Generation & Category Splitting

Read and follow `.gemini/skills/kcc-bulk-resource-generator/generate/SKILL.md` (`IsSkillFile: true`) to:
0. Source missing manageable GA declarative GCP resources using `hack/tools/greenfield/calculate_coverage.py LOCAL LOCAL 50 --resources-list` (cross-checked against `hack/tools/greenfield/{coverage_skip.json,deprecated_services.json,services_acronyms.json}`, `.build/googleapis.pb`, and `docs/ai/declarative-resource-criteria.md`).
1. Configure `apis/<service>/generate.sh` with `generate-types --prepopulate-resource`, `generate-identity`, and `generate-mapper --prepopulate-resource` (strictly omitting `--allow-non-ga-proto`).
2. Run batch code and CRD generation across all candidate services.
3. Classify and split the resulting services into committed branches by category:
   - **Category 1 (`Zero-Touch`):** Adding `--prepopulate-resource` types, identities, and mappers compiled and passed `apichecks` with no extra setup.
   - **Category 2 (`Initial Setup`):** Services requiring initial `go.mod`/`go.sum` dependencies, new service package scaffolding (`doc.go`, `groupversion_info.go`), or shared proto packages (e.g., `,google.type`) in `generate-mapper --service`.
   - **Category 3 (`Major Surgery`):** Services requiring manual mapper stubs/fixes, existing controller caller updates, `TestCRDObjectTypes` empty `oneof` workarounds, or multi-stage `generate.sh` scripts.

---

## Phase 2: Isolated Subagent Verification

To avoid context bloat and self-review bias, **do not** run the final verification checklist in the same agent context that performed Phase 1.

For each committed category branch produced in Phase 1:
1. **Prevent Git Working-Tree Collisions:**
   - If spawning verifier subagents **concurrently**, set `Workspace: "share"` (or `"branch"`) on `invoke_subagent` so each subagent gets its own isolated working tree and can `git checkout <category-branch>` without colliding.
   - If using `Workspace: "inherit"`, run only **one** verifier subagent at a time after checking out `<category-branch>`.
2. **Spawn a dedicated verifier subagent** (`invoke_subagent` with `TypeName: "self"`) using a clear prompt:
   ```text
   Check out branch `<category-branch>` and verify it as `<Category 1 | Category 2 | Category 3>` by reading and executing `.gemini/skills/kcc-bulk-resource-generator/verify/SKILL.md`. Return the structured Verification Report.
   ```
3. Wait for each verifier subagent's report and fix any flagged violations before pushing or opening the draft PRs.

- **Declarative Resource Filtering:** Before bulk-generating types or controllers, check candidate resources against [`docs/ai/declarative-resource-criteria.md`](../../../docs/ai/declarative-resource-criteria.md) (or `../../../../docs/ai/declarative-resource-criteria.md`) to separate read-only/ephemeral/data-plane resources from declarative control-plane resources.
