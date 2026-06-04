---
name: move-pr-forwards
description: Provides instructions and tools for an automated agent to move a Pull Request (PR) forwards towards merging. It analyzes PR status, checks for rebasing needs, filters feedback from KCC project approvers, and inspects failing CI logs for actionable agent hints.
---

# Move PR Forwards

## Overview

This skill guides an automated agent through the process of unblocking a Pull Request (PR) and moving it forwards to get merged. In a mature repository like Kubernetes Config Connector (KCC), PRs can get stalled due to CI failures, merge conflicts, or pending feedback from project approvers. 

This skill provides a helper Python script (`check-pr.py`) that inspects the PR status, parses comments from core maintainers, identifies failing CI checks, and deep-dives into GitHub Actions logs to find `<hint_for_agent>` directives, `FAIL:` messages, or `ERROR` strings.

## Workflow: Moving a PR Forwards

When tasked with moving a PR forwards, follow these steps:

### 1. Run the `check-pr.py` Helper Script

Execute the helper script on the target PR to gather all necessary context in a clean, structured summary without cluttering your context window.

```bash
./.gemini/skills/move-pr-forwards/scripts/check-pr.py --pr <PR_NUMBER>
```
*(Example: `./.gemini/skills/move-pr-forwards/scripts/check-pr.py --pr 8426`)*

### 2. Analyze the Output and Determine Next Steps

The script breaks down the PR state into three key areas. Address them in the following priority:

#### A. Failing CI Checks
The script fetches the exact job logs for any failing CI checks and scans for `<hint_for_agent>...</hint_for_agent>`.
- **Strong Hint Found**: If the log contains a strong hint (e.g., `Run 'dev/ci/presubmits/tests-e2e-fixtures-dataform' to fix.`), execute that exact command in the repository root. These presubmit scripts are designed to automatically regenerate golden files, fix formatting, or update mappers.  Typically you will run the test to generate updated golden output;
the test will fail the first time, but then should pass the second time (with the updated golden output).  Normally you will commit the updated golden output with a message like "autogen: update golden output".
- **Fallback Analysis**: If no explicit hint is found, review the `FAIL:` or `ERROR:` lines summarized by the script to diagnose the root cause (e.g., compilation error, linter failure, or unexpected golden diff).

#### B. Approver Feedback
The script filters comments and reviews specifically from KCC project approvers in `OWNERS` (`acpana`, `barney-s`, `cheftako`, `fedebongio`, `justinsb`, `ldanielmadariaga`, `maqiuyujoyce`, `xiaoweim`).
- Carefully read their feedback.
- Make the requested code changes or clarifications.

#### C. Rebasing & Merge Conflicts
If the script warns that the PR has merge conflicts (`CONFLICTING`) or is behind the base branch (`BEHIND`/`DIRTY`), rebase the branch:
```bash
gh pr checkout <PR_NUMBER>
git fetch origin master
git rebase origin/master
# Resolve conflicts if any, then:
git push --force-with-lease origin <HEAD_BRANCH>
```

### 3. Apply Fixes, Verify, and Push

1. **Checkout the PR Branch**:
   ```bash
   gh pr checkout <PR_NUMBER>
   ```
2. **Apply the Fix**: Run the presubmit script, edit the code to address approver comments, or complete the rebase.
3. **Validate**: Always run `make fmt` and `go vet ./...` to ensure formatting and basic compilation are clean.
4. **Commit and Push**:
   ```bash
   git add .
   git commit -m "chore: address CI failures and approver feedback"
   git push --force-with-lease origin <HEAD_BRANCH>
   ```
   *(Alternatively, you can use the `send-pr` skill if appropriate).*

## Resources

### scripts/check-pr.py
A robust Python CLI tool that uses `gh` to fetch PR metadata, filter comments from KCC OWNERS, identify failing CI jobs, fetch their full logs, and extract actionable `<hint_for_agent>` directives.
