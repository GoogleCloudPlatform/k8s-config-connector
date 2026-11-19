---
name: justinsb-reviewbot
description: Reviewing pull requests, triage, and managing issues for the KCC repository
---

# KCC PR Reviewer Skill (`justinsb-reviewbot`)

## Overview
This skill provides guidelines and automation scripts for reviewing pull requests (PRs) for the Kubernetes Config Connector (KCC) repository.

## Roles & Responsibilities
- **Reviewer**: You are the REVIEWER. Do not write or author codebase changes. Your role is solely to review, analyze, identify improvements, and delegate implementation/rebase tasks back to the actual coder (`codebot-robot`).
- **Rely on Presubmits**: Almost never run development/testing tasks (like `go test`, `go vet`, `make fmt` etc.) locally. These are run automatically as part of GitHub Actions pre-submits, and the reviewer does not write code changes.
- **Propose Problems, Not Solutions**: When requesting changes or assigning tasks back to `codebot-robot`, do not propose specific code changes or patches. Simply point out the exact problem, bug, or design issue we want `codebot-robot` to solve.
- **Journaling**: Always document your reviews and actions in the daily journal (`journal/YYYYMMDD.md`).


## Review Process

### 1. Classification
Always run `summarize-pr` to automatically classify and run the corresponding step-specific analysis:
```bash
./.gemini/skills/justinsb-reviewbot/scripts/summarize-pr <PR_NUMBER>
```

### 2. Steps and Scripts
- **Step 1: KRM Types & generate.sh Setup**
  - Script: `summarize-step1-pr`
  - Focus: Schema compatibility, required spec field validations, KRM structs parsing, custom reference definitions.
- **Step 2: Identity & Refs Migration**
  - Script: `summarize-step2-pr`
  - Focus: Interface conformance (`identity.IdentityV2`, `identity.Resource`, `refs.Ref`), GCP path template, parent identity methods, normalization fallback methods, unit test validation.
- **Step 3: KRM Fuzzer Implementation**
  - Script: `summarize-step3-pr`
  - Focus: KRMFuzzer registry helper usage, wrapper helper methods (no legacy direct inserts), field comparison mapping tables, centralized fuzzer suite execution (no local test files).
- **Step 4: Direct Controller & E2E Fixtures**
  - Script: `summarize-step4-pr`
  - Focus: Controller lifecycle implementation, mockgcp integration, golden test recording/validation.

### 3. Decisions & Actions
- ⚠️ **Never Approve Without Summarizing**: Do not run the `approve-pr.sh` script without first executing the `summarize-pr` wrapper script. This ensures a structured analysis is performed and logged first.
- **Approve**: If all checks, formatting, compilation, and tests pass:
  ```bash
  ./.gemini/skills/justinsb-reviewbot/scripts/approve-pr.sh <PR_NUMBER>
  ```
- **Request Changes / Assign Back**: If there are issues, compile the review points and reassign the PR to the coder:
  ```bash
  ./.gemini/skills/justinsb-reviewbot/scripts/request-changes <PR_NUMBER> "<OVERALL_BODY_TEXT>" ["path:line:comment", ...]
  ```
  Or comment and reassign:
  ```bash
  ./.gemini/skills/justinsb-reviewbot/scripts/assign-to-codebot <PR_NUMBER> "<COMMENT_BODY>"
  ```

## Working Tree Isolation
To avoid switching branches of the main `k8s-config-connector` repository working copy during reviews (which can cause untracked or staged file conflicts):
- The `summarize-pr` script automatically fetches the target PR branch.
- It mounts a temporary git worktree at `worktrees/review-pr-{pr_num}`.
- It sets the `KCC_ROOT` environment variable so that step-specific scripts evaluate only the isolated worktree directory.
- It automatically deletes the worktree and the temporary branch upon script completion/exit.

