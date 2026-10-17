---
name: reviewgen-greenfield-new-types
description: Entry point for reviewing KCC Greenfield types PRs. Routes between Step 1 (`prepopulate/SKILL.md` for initial mechanical `--prepopulate-resource` scaffolding PRs) and Step 2 (`judgement/SKILL.md` for `judgement_queue.yaml` triage and finalization PRs).
---

# Review Guide for KCC Greenfield New Types PRs (Router)

This skill is the entry point for reviewing **Greenfield Step 1 (Prepopulate)** and **Greenfield Step 2 (Judgement Queue Triage)** PRs. It routes to one of two review checklists depending on whether `<Kind>`'s `untriaged-bulk-generation` entry in `apis/<service>/judgement_queue.yaml` is still `status: open` (Step 1) or `status: resolved` (Step 2).

## Detect Mode and Route

Inspect `apis/<service>/judgement_queue.yaml` for the target `<Kind>`:

```bash
awk -v kind="<Kind>" '
  $1 == "-" && $2 == "kind:" && $3 == kind { in_kind=1; next }
  in_kind && $1 == "-" && $2 == "kind:" { in_kind=0 }
  in_kind && $1 == "reason:" && $2 == "untriaged-bulk-generation" { found=1; next }
  found && $1 == "status:" { print $2; exit }
' "apis/<service>/judgement_queue.yaml"
```

- **If `untriaged-bulk-generation` is `open` (Step 1: Mechanical Prepopulate PR)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-new-types/prepopulate/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/reviewgen-greenfield-new-types/prepopulate/SKILL.md)** using `view_file`.
- **If `untriaged-bulk-generation` is `resolved` (Step 2: Judgement Queue Triage & Finalize Types PR)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-new-types/judgement/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/reviewgen-greenfield-new-types/judgement/SKILL.md)** using `view_file`.
