
# KCC Release Workflow

This document outlines the automated steps to prepare a new release for KCC.

## PHASE 1: PREPARATION

### Step 1.1: Determine Release Versions

**Description:** Identify the version for the new release and the most recent existing version.

**Actions:**
1.  To find the `STALE_VERSION`, run this command and use the latest version from the output:
    ```bash
    git tag | sort -V | tail -n 1
    ```
2.  The `VERSION` for the new release should be the `STALE_VERSION` with the minor version incremented (e.g., if stale is `v1.132.1`, the new version is `1.133.0`).

### Step 1.2: Create Release Branch

**Description:** Create and check out a new release branch.

**Actions:**
1.  Execute this command, replacing `${VERSION}` with the value from Step 1.1:
    ```bash
    git checkout -b release-${VERSION}
    ```
2.  **Success Condition:** The command should execute with exit code 0.

## PHASE 2: CODE & MANIFEST GENERATION

### Step 2.1: Propose Tag and Update Manifests

**Description:** Run the `propose-tag` script to update version numbers and manifests across the codebase.

**Actions:**
1.  Execute this command, using the variables from Step 1.1:
    ```bash
    VERSION=${VERSION} STALE_VERSION=${STALE_VERSION} ./dev/tasks/propose-tag
    ```
2.  Commit the changes using the exact commit message format below:
    ```bash
    git add .
    git commit -m "Release ${VERSION}"
    ```
3.  **Success Condition:** Both commands should execute with exit code 0.

## PHASE 3: TESTING & VERIFICATION

### Step 3.1: Run Unit Tests

**Description:** Run the unit tests for the operator controllers to ensure the changes haven't introduced regressions.

**Actions:**
1.  Execute this command:
    ```bash
    cd operator && go test ./pkg/controllers/...
    ```
2.  **Success Condition:** The test command exits with code 0.
3.  **Failure Handling:** If the test command fails with an exit code of 1 and the output contains "unexpected diff between actual and expected YAML", it means the golden test files are out of date. In this case, perform the following remediation steps:
    a. Run this command to update the golden files:
       ```bash
       cd operator && WRITE_GOLDEN_OUTPUT="true" go test ./pkg/controllers/...
       ```
    b. Commit the updated files with this exact message:
       ```bash
       git add .
       git commit -m "Update golden files for operator controllers"
       ```
    c. Rerun the original test command from Step 3.1.1 to confirm it now passes.

## PHASE 4: GENERATE RELEASE SUMMARY

**Description:** Analyze the three release commits to create a detailed summary for the pull request body. This process filters out noisy changes like simple version bumps to highlight the most important changes.

**Actions:**

1.  **Get Commit Information:**
    *   Run the following command to get the hash and subject of the last three commits and save them to `commits.log`:
        ```bash
        git log -n 3 --pretty="format:%H %s" > commits.log
        ```
    *   Inspect `commits.log` to identify the commit hashes for the "Release", "Update alpha CRDs", and "Update golden files" commits.

2.  **Generate and Filter Diffs:**
    *   For each of the three commits, perform the following steps. Replace `<commit_hash>` with the hash of the commit and `<commit_name>` with a descriptive name (e.g., `release`, `crd_update`, `golden_file_update`).

    *   **Generate Raw Diff:**
        ```bash
        git show <commit_hash> > <commit_name>_diff.txt
        ```

    *   **Filter Diff:**
        *   The following command filters the raw diff, removing file-specific diffs that only contain changes to the `cnrm.cloud.google.com/version` annotation. This helps to focus on more substantial changes.
        ```bash
        cat <commit_name>_diff.txt | perl -pe 'BEGIN{$/="\ndiff --git "} s/^.*//s if !grep {/^[+-]/ && !/---|\+\+\+/ && !/cnrm\.cloud\.google\.com\/version/} split/\n/' > <commit_name>_filtered_diff.txt
        ```
        *   After running this for all three commits, you will have three filtered diff files (e.g., `release_filtered_diff.txt`, `crd_update_filtered_diff.txt`, `golden_file_update_filtered_diff.txt`).

3.  **Assemble `release-summary.md`:**
    *   Create a new file named `release-summary.md`.
    *   Review the filtered diff files to identify the key changes for the release.
    *   Write a summary of the changes in `release-summary.md`. Your summary should describe the significant updates, such as new features, bug fixes, or important structural changes, using the filtered diffs as a reference.


## PHASE 5: PUSH & FINALIZE

### Step 5.1: Push Branch to Remote

**Description:** Push the release branch to the `origin` remote to prepare for a Pull Request.

**Actions:**
1.  Execute this command:
    ```bash
    git push origin release-${VERSION}
    ```
2.  **Success Condition:** The command should execute with exit code 0.
3.  **Final Output:** Provide the user with the URL to create a new pull request.
