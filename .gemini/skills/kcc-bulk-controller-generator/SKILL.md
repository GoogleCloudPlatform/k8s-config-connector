---
name: kcc-bulk-controller-generator
description: Orchestrator workflow for bulk-generating deterministic direct controllers (`generate-controller --prepopulate-resource`) and round-trip fuzzers (`generate-fuzzer --prepopulate-resource`) across multiple greenfield GCP resources (either post-judgement after Step 1b or pre-judgement right after Step 1a), splitting them into PR categories, and verifying each PR via an isolated subagent.
---

# KCC Bulk Controller & Fuzzer Generator (Orchestrator)

This skill orchestrates the two-phase workflow for bulk-generating deterministic direct controllers (`<proto_lower>_controller.go`) and round-trip KRM fuzzers (`<kind_lower>_fuzzer.generated.go`) across multiple greenfield KCC resources, splitting the output into three reviewable PR categories, and verifying each category branch in an isolated subagent context.

---

## Supported Pipeline Modes

1. **Standard Mode — Post-Judgement (After Step 1a + Step 1b Complete)**:
   - Target resources whose `untriaged-bulk-generation` entry in `apis/<service>/judgement_queue.yaml` is `status: resolved` (or finalized greenfield types without open queue entries) and that do not yet have a direct controller/fuzzer.
   - KRM types, identities, references, top-level `labels` removal, and mappers are already finalized, so the generated controllers and fuzzers reflect the final triaged schema.
2. **Ad-Hoc Campaign Mode Only — Pre-Judgement (Immediately After Step 1a, Before Step 1b)**:
   - **Strictly for ad-hoc human-invoked bulk generation campaigns — NEVER for systemic per-resource automated workflows.**
   - Targets resources that were bulk-prepopulated in Step 1a (`untriaged-bulk-generation` is `status: open` in `apis/<service>/judgement_queue.yaml`) when doing a one-off repository-wide bulk generation pass.
   - Later, when Step 1b (`step/types-judgement`) runs on each resource, [`.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/SKILL.md`](../kcc-direct-greenfield-types-implementer/judgement/SKILL.md) updates the CRD/KRM types and regenerates the controller (`// +generated:controller`) and fuzzer (`// +generated:fuzzer`) in lockstep.


---

## Phase 1: Bulk Controller & Fuzzer Generation + Category Splitting

Read and follow [`.gemini/skills/kcc-bulk-controller-generator/generate/SKILL.md`](generate/SKILL.md) (`IsSkillFile: true`) to:
1. Configure `apis/<service>/generate.sh` with `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` for the target greenfield `--resource <Kind>:<ProtoMessage>` entries (strictly protecting existing handwritten controllers and fuzzers).
2. Run batch controller and fuzzer generation across all candidate services and verify compilation and round-trip fuzzing (`FuzzAllMappers`).
3. Classify and split the resulting services into committed branches by category:
   - **Category 1 (`Zero-Touch`):** Adding `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` compiled cleanly and passed `FuzzAllMappers` and `TestDirectResourceFileNaming` with no extra setup.
   - **Category 2 (`Initial Setup`):** Services requiring GAPIC client `go.mod`/`go.sum` upgrades or additions (e.g., `go get cloud.google.com/go/<service>/apiv1@...`).
   - **Category 3 (`Major Surgery`):** Services whose proto RPC signatures are non-standard (missing standard CRUD RPCs, custom singleton RPCs, or requiring a `FilterSpec`/`FilterStatus` fuzzer hook) that require targeted adjustments.

---

## Phase 2: Isolated Subagent Verification

To avoid context bloat and self-review bias, **do not** run the final verification checklist in the same agent context that performed Phase 1.

For each committed category branch produced in Phase 1:
1. **Prevent Git Working-Tree Collisions:**
   - If spawning verifier subagents **concurrently**, set `Workspace: "share"` (or `"branch"`) on `invoke_subagent` so each subagent gets its own isolated working tree and can `git checkout <category-branch>` without colliding.
   - If using `Workspace: "inherit"`, run only **one** verifier subagent at a time after checking out `<category-branch>`.
2. **Spawn a dedicated verifier subagent** (`invoke_subagent` with `TypeName: "self"`) using a clear prompt:
   ```text
   Check out branch `<category-branch>` and verify it as `<Category 1 | Category 2 | Category 3>` by reading and executing `.gemini/skills/kcc-bulk-controller-generator/verify/SKILL.md`. Return the structured Verification Report.
   ```
3. Wait for each verifier subagent's report and fix any flagged violations before pushing or opening the draft PRs.
