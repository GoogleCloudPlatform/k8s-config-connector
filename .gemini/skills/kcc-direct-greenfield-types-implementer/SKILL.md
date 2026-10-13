---
name: kcc-direct-greenfield-types-implementer
description: Automate the initial scaffolding or judgement-queue triage of a KCC "direct" greenfield resource, including CRD types, identity, references, and generation scripts. Use this when implementing Step 1 (types/CRD) for a new "direct" GCP resource.
---

# KCC Direct Greenfield Types Implementer (Router)

This skill is the entry point for **Greenfield Step 1** (KRM types, identity, references, mappers, and CRDs). It routes to one of two workflows depending on whether `<Kind>` has already been bulk-scaffolded with a `judgement_queue.yaml`.

## Prerequisites
You **must** always apply the baseline KRM standards in:
- [`.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-base-types-implementer/SKILL.md)
- [`.gemini/skills/kcc-identity-reference/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-identity-reference/SKILL.md)

## Step 1: Detect Mode and Route

Check whether `apis/<service_short>/judgement_queue.yaml` exists and contains entries for `<Kind>`:

```bash
grep -q "kind: <Kind>$" "apis/<service_short>/judgement_queue.yaml" 2>/dev/null
```

- **If `apis/<service_short>/judgement_queue.yaml` contains entries for `<Kind>` (Bulk-Generated Resource)**:
  - Read and follow **`.gemini/skills/kcc-direct-greenfield-types-implementer/JUDGEMENT_QUEUE.md`** using `view_file`.
- **If `apis/<service_short>/judgement_queue.yaml` does NOT exist or has NO entries for `<Kind>` (Vanilla Greenfield Resource)**:
  - Read and follow **`.gemini/skills/kcc-direct-greenfield-types-implementer/VANILLA.md`** using `view_file`.
