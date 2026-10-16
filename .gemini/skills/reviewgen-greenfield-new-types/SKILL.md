---
name: reviewgen-greenfield-new-types
description: Provides clear review criteria for reviewing PRs that add or finalize KCC types, identity, and references for Greenfield resources.
---

# Review Guide for KCC Greenfield New Types PRs (Router)

This skill is the entry point for reviewing **Greenfield Step 1** PRs (KRM types, identity, references, mappers, and CRDs). It routes to one of two review checklists depending on whether the PR involves `judgement_queue.yaml`.

## Step 1: Detect Mode and Route

Check whether the PR diff modifies `apis/<service>/judgement_queue.yaml`, or modifies a `<Kind>` that has entries in `apis/<service>/judgement_queue.yaml`:

```bash
grep -q "kind: <Kind>$" "apis/<service>/judgement_queue.yaml" 2>/dev/null
```

- **If YES (Bulk-Generated / Judgement Queue PR)**:
  - Read and follow **`.gemini/skills/reviewgen-greenfield-new-types/JUDGEMENT_QUEUE.md`** using `view_file`.
- **If NO (Vanilla Greenfield PR)**:
  - Read and follow **`.gemini/skills/reviewgen-greenfield-new-types/VANILLA.md`** using `view_file`.
