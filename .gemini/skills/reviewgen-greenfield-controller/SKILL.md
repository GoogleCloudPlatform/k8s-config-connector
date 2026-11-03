---
name: reviewgen-greenfield-controller
description: Entry point for reviewing KCC Greenfield controller PRs. Routes between `generated/SKILL.md` (for deterministic `--prepopulate-resource` controllers and fuzzers) and `legacy/SKILL.md` (for in-flight handwritten controller PRs).
---

# Review Guide for KCC Greenfield Controller PRs (Router)

This skill is the entry point for reviewing **Greenfield Step 2 (Direct Controller, Fuzzer, and E2E Fixtures)** PRs. It routes between the deterministic controller/fuzzer review checklist (`generated/SKILL.md`) and the legacy review checklist (`legacy/SKILL.md`) so in-flight PRs written in the old style are not disrupted.

## Detect Mode and Route

Inspect the controller file for `<Kind>` in `pkg/controller/direct/<service>/`:

```bash
grep -rn "+generated:controller" pkg/controller/direct/<service>/ 2>/dev/null
```

*(`apis/<service>/generate.sh` is not a routing signal: greenfield controllers and fuzzers are generated with one-shot `controllerbuilder` commands, and `generate.sh` must not invoke `generate-controller` / `generate-fuzzer`.)*

- **If the controller file for `<Kind>` contains `// +generated:controller` (or if the PR is a new Greenfield Step 2 PR created after deterministic controller generation)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-controller/generated/SKILL.md`](generated/SKILL.md)** using `view_file`.
- **If the PR is an in-flight handwritten controller PR whose `*_controller.go` does NOT contain `// +generated:controller`**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-controller/legacy/SKILL.md`](legacy/SKILL.md)** using `view_file`.
