---
name: kcc-direct-greenfield-types-implementer
description: Entry point for KCC greenfield types generation and judgement-queue triage. Routes between Step 1 (`prepopulate/SKILL.md` for initial mechanical `--prepopulate-resource` scaffolding) and Step 2 (`judgement/SKILL.md` for `judgement_queue.yaml` triage and finalization).
---

# KCC Direct Greenfield Types Implementer (Router)

This skill is the entry point for **Greenfield Step 1 (Prepopulate)** and **Greenfield Step 2 (Judgement Queue Triage)**. It routes to one of two sub-workflows depending on whether `<Kind>` has already been bulk-scaffolded (`apis/<service>/<version>/<kind>_types.go`).

## Prerequisites
You **must** always apply the baseline KRM standards in:
- [`.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-base-types-implementer/SKILL.md)
- [`.gemini/skills/kcc-identity-reference/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-identity-reference/SKILL.md)

## Detect Mode and Route

Check whether `apis/<service>/<version>/<kind_lowercase>_types.go` already exists in the repository:

```bash
find apis -maxdepth 3 -iname "<kind_lowercase>_types.go"
```

- **If `apis/<service>/<version>/<kind_lowercase>_types.go` does NOT exist yet (Step 1: Prepopulate Resource & Judgement Queue)**:
  - Read and follow **[`.gemini/skills/kcc-direct-greenfield-types-implementer/prepopulate/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/prepopulate/SKILL.md)** using `view_file`.
- **If `apis/<service>/<version>/<kind_lowercase>_types.go` ALREADY exists (Step 2: Triage Judgement Queue & Finalize KRM Types)**:
  - Read and follow **[`.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/SKILL.md)** using `view_file`.
