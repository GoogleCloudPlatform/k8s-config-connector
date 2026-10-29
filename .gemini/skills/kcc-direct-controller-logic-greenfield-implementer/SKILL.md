---
name: kcc-direct-controller-logic-greenfield-implementer
description: Entry point for implementing Phase 2 (direct controller, fuzzer, and E2E fixtures) for KCC Greenfield resources. Routes between `generated/SKILL.md` (deterministic controller & fuzzer generation via `--prepopulate-resource`) and `legacy/SKILL.md` (for in-flight handwritten controller PRs).
---

# KCC Direct Controller Logic Greenfield Implementer (Router)

This skill is the entry point for **Greenfield Step 2 (Direct Controller, Fuzzer, and E2E Fixtures)**. It routes to either the deterministic generator workflow (`generated/SKILL.md`) or the legacy handwritten workflow (`legacy/SKILL.md`) so in-flight PRs written in the old style are not disrupted.

> [!IMPORTANT]
> **Hard Prerequisite — Post-Judgement Only:** For systemic per-resource workflows, Step 1b (`apis/<service>/judgement_queue.yaml` triage) **must** already be completed (`untriaged-bulk-generation` and all actionable entries for `<Kind>` marked `status: resolved`) before running Phase 2. Pre-judgement controller generation is strictly reserved for ad-hoc bulk campaigns (`kcc-bulk-controller-generator`).

## Detect Mode and Route

Check whether an in-flight handwritten controller (without `// +generated:controller`) already exists on the current branch for `<Kind>`:

```bash
grep -rn "RegisterModel(.*<Kind>GVK" pkg/controller/direct/<service>/ 2>/dev/null
```

- **If no controller exists yet for `<Kind>`, OR the existing controller file contains `// +generated:controller` (Default for all new Step 2 PRs and bulk-pregenerated controllers)**:
  - Read and follow **[`.gemini/skills/kcc-direct-controller-logic-greenfield-implementer/generated/SKILL.md`](generated/SKILL.md)** using `view_file`.
- **If a controller file for `<Kind>` already exists on the branch and does NOT contain `// +generated:controller` (In-Flight Legacy Handwritten Controller PR)**:
  - Read and follow **[`.gemini/skills/kcc-direct-controller-logic-greenfield-implementer/legacy/SKILL.md`](legacy/SKILL.md)** using `view_file`.

