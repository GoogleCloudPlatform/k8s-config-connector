---
name: send-pr
description: CRITICAL: MUST be used whenever sending a PR, updating an existing PR, or pushing any commit to a remote branch. Provides instructions and pre-push checks (`dev/tasks/validate-and-push`) before pushing or interacting with GitHub via `gh`.
---

# Send PR & Safe Remote Push

## Overview

This skill mandates strict pre-push presubmit validation (`dev/tasks/validate-and-push`) whenever creating or updating a Pull Request (PR) or pushing any commit to a remote branch in `k8s-config-connector`.

## Workflow: Updating an Existing PR or Pushing to a Remote Branch

Before pushing to any remote branch (`git push ...`), you MUST run the automated pre-push validation script:
```bash
./dev/tasks/validate-and-push origin <branch-name>
```
`validate-and-push` sequentially runs `make fmt`, `go vet ./...`, `./dev/ci/presubmits/validate-generated-files`, and `./dev/ci/presubmits/unit-tests`. It will only execute `git push` if all local authoritative presubmits exit cleanly (`0`).

1. **Write the PR Body**:
   Create a temporary markdown file containing the body/description of your PR. Make sure to reference the issue you are solving (e.g., `Fixes #1234`).
   **CRITICAL**: Do NOT check in this file. It is recommended to use the project's temporary directory to avoid accidentally committing it.
   ```bash
   cat << 'EOF' > /root/.gemini/tmp/k8s-config-connector/pr-body.txt
   Description of the PR.

   Fixes #1234
   EOF
   ```

2. **Run the `send-pr.sh` Script**:
   Execute the script by providing the PR title, the path to the body file you just created, and optionally a comma-separated list of labels.
   ```bash
   ./.gemini/skills/send-pr/scripts/send-pr.sh --title "Brief PR Title" --body /root/.gemini/tmp/k8s-config-connector/pr-body.txt --labels "overseer,area/direct"
   ```

3. **MANDATORY PRE-PUSH VALIDATION**:
   - Before running `send-pr.sh` or pushing any commit, always verify your changes locally using the authoritative presubmit scripts:
     1. Run `make fmt && go vet ./...`.
     2. Run `./dev/ci/presubmits/unit-tests` (validates all package unit tests and golden log alignment).
     3. Run `./dev/ci/presubmits/tests-e2e-fixtures-<service>` when modifying a resource controller or fixture.

4. **What the Script Does**:
   - Runs `make fmt` to ensure the code is properly formatted.
   - Checks if `make fmt` introduced any git diffs or if there are any uncommitted changes. If there are, it will *stop the push* and require you to commit the formatting changes first.
   - Pushes the current branch to `origin` (`git push --set-upstream origin <branch>`).
   - Uses the `gh` tool to create a new PR or update the existing PR if one is already open for this branch.

4. **Handling Failures**:
   - If the script fails because `make fmt` introduced changes, you should add and commit those changes:
     ```bash
     git add .
     git commit -m "chore: make fmt"
     ```
     Then, run `send-pr.sh` again.

## Resources

### scripts/send-pr.sh
A bash script that coordinates `make fmt`, `git push`, and `gh pr create`/`gh pr edit`. It ensures good PR hygiene by checking for formatting changes before sending.
