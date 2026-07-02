PR #7225 (chore(deps): bump github.com/hashicorp/terraform-plugin-sdk/v2 from 2.24.0 to 2.35.0) is failing presubmit checks.

# Task
1. **Update Dependency**:
    - Do NOT push changes to the Dependabot branch directly. Instead, create a new branch from `master`, apply the dependency update, and fix the test failures in your new branch.
2. **Verify**:
    - Please verify by running `make all-binary`.
    - Please verify the fix by going to the `mockgcp` directory and run `make all`.
    - Please fix any problem identified then re-verify.
    - Loop up to 10 times to try to resolve any issues found.
3. **Push & PR**:
    - Commit the changes:
      ```bash
      git add go.mod go.sum
      git commit -m "Update github.com/hashicorp/terraform-plugin-sdk/v2 to 2.35.0" -m "Fixes build failure." -m "TAG=agy"
      ```
    - Push the branch to your fork and create a Pull Request.
