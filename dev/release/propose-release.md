
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

### Step 2.2: Synchronize CRDs

**Description:** Run the `sync-crds-folder.sh` script to update the CRDs in the `/crds` directory.

**Actions:**
1.  Execute this command:
    ```bash
    VERSION=${VERSION} ./dev/tasks/sync-crds-folder.sh
    ```
2.  Commit the changes using the exact commit message format below:
    ```bash
    git add .
    git commit -m "Update alpha CRDs for Release ${VERSION}"
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

**Description:** Analyze the three release commits to create a detailed summary for the pull request body, following the specified format. Save the output to `release-summary.md`.

**Actions:**

1.  **Get Commit Hashes:**
    *   Run `git log -n 3 --pretty="format:%H"` to get the hashes of the last three commits. Identify which hash corresponds to which commit (Release, CRDs, Golden Files).

2.  **For each commit, gather the following details:**
    *   **Statistics:** Run `git show --stat <commit_hash>` to get the number of files changed, insertions, and deletions.
    *   **New Files:** Run `git show --summary <commit_hash> | grep "create mode"` to list newly created files.
    *   **Renamed Files/Dirs:** Run `git show --summary <commit_hash> | grep "rename"` to list renames.

3.  **Assemble `release-summary.md`:**
    *   Create a new file named `release-summary.md`.
    *   Using the information gathered above, construct the summary by formatting the text exactly as specified in the example below.
    *   For Commit 2, calculate the number of "Modified" files by subtracting the number of "Added" and "Renamed" files from the total number of "files changed".

**Example Output Format:**

```text
Commit 1: Release 1.133.0

This commit handled the main version bump for the release.

Summary of Changes: 19 files changed, with 4,060 additions and 1,576 deletions.
Added Directories: The script renamed versioned directories from 1.132.1 to 1.133.0.
operator/autopilot-channels/packages/configconnector/1.133.0
operator/channels/packages/configconnector/1.133.0
Content Changes: The vast majority of the 4,000+ line changes in the modified files were the replacement of the version string
1.132.1 with 1.133.0 and updating the corresponding image tags.
---
Commit 2: Update alpha CRDs for Release 1.133.0

This commit updated the CRD files in the /crds directory based on the new release version. The changes reflect structural updates
to the CRDs themselves, not just version number changes.

Summary of Changes: 440 files changed, with 1,688 additions and 618 deletions.
File Changes:
Added: 1 new CRD file was added.
crds/apigateway_v1beta1_apigatewayapi.yaml
Renamed: 2 CRD files were renamed, reflecting a promotion in their API version from v1alpha1 to v1beta1.
crds/{apphub_v1alpha1_apphubapplication.yaml => apphub_v1beta1_apphubapplication.yaml}
crds/{apigateway_v1alpha1_apigatewayapi.yaml => bigtable_v1alpha1_bigtablelogicalview.yaml}
Modified: The remaining 437 files were modified with various structural changes to the CRD definitions as a result of the
sync-crds-folder.sh script.
---
Commit 3: Update golden files for operator controllers (ac75bf2)

This commit updated the golden test files to align with the changes from the previous two commits, ensuring the tests pass with
the new version and CRD structures.

Summary of Changes: 2 files changed, with 63 additions and 53 deletions.
Files Changed: 2 golden test files were modified.
operator/pkg/controllers/configconnector/testdata/golden/simple/_expected.yaml
operator/pkg/controllers/configconnectorcontext/testdata/golden/simple/_expected.yaml
Content Changes: The modifications inside these files align the test expectations with the changes from the previous two
commits, including the version bump to 1.133.0 and the addition/modification of CRD definitions in the test output.
```

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
