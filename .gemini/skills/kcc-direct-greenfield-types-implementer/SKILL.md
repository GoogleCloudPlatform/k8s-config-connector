---
name: kcc-direct-greenfield-types-implementer
description: Entry point for KCC greenfield types generation and judgement-queue triage. Routes between Step 1a (`prepopulate/SKILL.md` for initial mechanical `--prepopulate-resource` scaffolding) and Step 1b (`judgement/SKILL.md` for `judgement_queue.yaml` triage and finalization).
---

# KCC Direct Greenfield Types Implementer (Router)

This skill is the entry point for **Greenfield Step 1a (Prepopulate)** and **Greenfield Step 1b (Judgement Queue Triage)**. It routes to one of two sub-workflows depending on whether `<Kind>` has already been prepopulated in `master` (`apis/<service>/judgement_queue.yaml` and `apis/<service>/<version>/<kind_lowercase>_types.go`).

## Prerequisites
You **must** always apply the baseline KRM standards in:
- [`.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-base-types-implementer/SKILL.md)
- [`.gemini/skills/kcc-identity-reference/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-identity-reference/SKILL.md)

## Detect Mode and Route

Check whether `<Kind>` already has an entry in `apis/<service>/judgement_queue.yaml` and `apis/<service>/<version>/<kind_lowercase>_types.go` on `master` (or if the issue is labeled `step/types-judgement`):

```bash
git fetch origin master && git grep -n "kind: <Kind>$" origin/master -- 'apis/*/judgement_queue.yaml'
```

(This searches every group, so it also finds kinds that were bulk-prepopulated under a group prefix you did not expect.)

- **If `<Kind>` is NOT yet in `apis/<service>/judgement_queue.yaml` on `master` (Step 1a / `step/gen-types`: Prepopulate Resource & Judgement Queue)**:
  - Read and follow **[`.gemini/skills/kcc-direct-greenfield-types-implementer/prepopulate/SKILL.md`](prepopulate/SKILL.md)** using `view_file`.
- **If `<Kind>` ALREADY exists in `apis/<service>/judgement_queue.yaml` on `master` (Step 1b / `step/types-judgement`: Triage Judgement Queue & Finalize KRM Types)**:
  - Read and follow **[`.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/SKILL.md`](judgement/SKILL.md)** using `view_file`.

