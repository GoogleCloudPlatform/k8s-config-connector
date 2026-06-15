---
name: send-pr
description: Provides tools and instructions for sending Pull Requests (PRs), particularly from automated agents like codebot-robot. Use this skill when you need to send or update a PR, format code before pushing, or interact with github via the `gh` tool.
---

# Send PR

## Overview

This skill simplifies the process of creating and updating Pull Requests (PRs) in the `k8s-config-connector` repository. It provides a standard script that handles formatting checks, git pushing, and the `gh` tool interaction to create or edit PRs safely.

## Workflow: Creating or Updating a PR

When you are ready to send your changes as a Pull Request, use the provided `send-pr.sh` script.

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

3. **What the Script Does**:
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
