---
name: reviewgen-greenfield-controller
description: Entry point for reviewing KCC Greenfield controller PRs. Routes between `generated/SKILL.md` (for deterministic `--prepopulate-resource` controllers and fuzzers) and `legacy/SKILL.md` (for in-flight handwritten controller PRs).
---

# Review Guide for KCC Greenfield Controller PRs (Router)

This skill is the entry point for reviewing **Greenfield Step 2 (Direct Controller, Fuzzer, and E2E Fixtures)** PRs. It routes between the deterministic controller/fuzzer review checklist (`generated/SKILL.md`) and the legacy review checklist (`legacy/SKILL.md`) so in-flight PRs written in the old style are not disrupted.

## Detect Mode and Route

Inspect the controller file for `<Kind>` in `pkg/controller/direct/<service>/` and `apis/<service>/generate.sh`:

```bash
grep -rn "+generated:controller" pkg/controller/direct/<service>/ 2>/dev/null
```

- **If the controller file for `<Kind>` contains `// +generated:controller` OR `apis/<service>/generate.sh` invokes `generate-controller --prepopulate-resource` (or if the PR is a new Greenfield Step 2 PR created after deterministic controller generation)**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-controller/generated/SKILL.md`](generated/SKILL.md)** using `view_file`.
- **If the PR is an in-flight handwritten controller PR whose `*_controller.go` does NOT contain `// +generated:controller` and `apis/<service>/generate.sh` does not invoke `generate-controller`**:
  - Read and follow **[`.gemini/skills/reviewgen-greenfield-controller/legacy/SKILL.md`](legacy/SKILL.md)** using `view_file`.

---

## SOUL.md Invariant 7: Internal Information Sanitization (Blocking Finding)

Across all greenfield controller reviews:
*   **Blocking Finding**: Treat any Buganizer ID (`b/<id>`), internal bug URL (`buganizer.corp.google.com`), or customer name anywhere on GitHub (PR title, description, commit messages, code comments, or release notes) as an absolute **BLOCKING FINDING**. Request changes until removed.
*   **Reviewer Hygiene**: Reviewer comments must **NEVER** contain Buganizer IDs or customer names. Refer strictly to public GitHub issues or PRs.
