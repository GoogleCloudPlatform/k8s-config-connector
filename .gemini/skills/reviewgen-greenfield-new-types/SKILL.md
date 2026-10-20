---
name: reviewgen-greenfield-new-types
description: Entry point for reviewing KCC Greenfield types PRs. Routes between Step 1a (`prepopulate/SKILL.md` for initial mechanical `--prepopulate-resource` scaffolding PRs), Step 1b (`judgement/SKILL.md` for `judgement_queue.yaml` triage and finalization PRs), and `legacy/SKILL.md` (for in-flight PRs generated without `judgement_queue.yaml`).
---

# Review Guide for KCC Greenfield New Types PRs (Router)

This skill is the entry point for reviewing **Greenfield Step 1a (Prepopulate)**, **Greenfield Step 1b (Judgement Queue Triage)**, and **In-Flight Legacy Greenfield Types** PRs. It routes to one of three review checklists depending on whether `apis/<service>/judgement_queue.yaml` contains `<Kind>` and whether its `untriaged-bulk-generation` entry is `status: open` (Step 1a) or `status: resolved` (Step 1b).

## Detect Mode and Route

Inspect `apis/<service>/judgement_queue.yaml` for the target `<Kind>`:

```bash
awk -v kind="<Kind>" '
  $1 == "-" && $2 == "kind:" && $3 == kind { in_kind=1; next }
  in_kind && $1 == "-" && $2 == "kind:" { in_kind=0 }
  in_kind && $1 == "reason:" && $2 == "untriaged-bulk-generation" { found=1; next }
  found && $1 == "status:" { print $2; exit }
' "apis/<service>/judgement_queue.yaml" 2>/dev/null
```

- **If `untriaged-bulk-generation` is `open` (Step 1a / `step/gen-types`: Mechanical Prepopulate PR)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-new-types/prepopulate/SKILL.md`](.gemini/skills/reviewgen-greenfield-new-types/prepopulate/SKILL.md)**.
- **If `untriaged-bulk-generation` is `resolved` (Step 1b / `step/types-judgement`: Judgement Queue Triage & Finalize Types PR)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-new-types/judgement/SKILL.md`](.gemini/skills/reviewgen-greenfield-new-types/judgement/SKILL.md)**.
- **If `apis/<service>/judgement_queue.yaml` does not exist or has no entry for `<Kind>` (In-Flight Legacy Greenfield Types PR)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-new-types/legacy/SKILL.md`](.gemini/skills/reviewgen-greenfield-new-types/legacy/SKILL.md)**.

---

## SOUL.md Invariant 7: Internal Information Sanitization (Blocking Finding)

Across all greenfield new types reviews:
*   **Blocking Finding**: Treat any Buganizer ID (`b/<id>`), internal bug URL (`buganizer.corp.google.com`), or customer name anywhere on GitHub (PR title, description, commit messages, code comments, or release notes) as an absolute **BLOCKING FINDING**. Request changes until removed.
*   **Reviewer Hygiene**: Reviewer comments must **NEVER** contain Buganizer IDs or customer names. Refer strictly to public GitHub issues or PRs.
