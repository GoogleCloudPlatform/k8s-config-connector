# Pushing a Release Tag

This document guides the Gemini CLI to find the correct release commit on an existing branch and execute the `push-tag-github.sh` script.

## Gemini, please follow these instructions:

When a user asks to "push a release tag", you must follow this multi-step process precisely.

### Step 1: Gather Information and Prepare the Repository

1.  Ask the user for the **major.minor** version of the release (e.g., `1.133`).
2.  Construct the release branch name from the user's input (e.g., `release-1.133`).
3.  Check out the corresponding release branch.

    ```bash
    git checkout upstream/release-MAJOR.MINOR
    ```

### Step 2: Find the Release Commit

1.  Use `git log` to find the specific release commit. The command below will search the current history for the latest commit with a message **exactly** matching the format "Release MAJOR.MINOR.PATCH", excluding any merge commits. It will then output the commit hash and the full commit message.

    ```bash
    git log --no-merges --grep="^Release MAJOR.MINOR." --pretty=format:'%H %s' -n 1
    ```
2.  From the output, extract the full version string (e.g., `1.133.0`) from the commit message for the `VERSION` variable, and the commit hash for the `GIT_COMMIT` variable.

3. **Important:** Return to the master branch before executing the script to ensure a clean state and rebase to ensure that branches do not diverge.
    ```bash
    git checkout master
    ```

### Step 3: Get the Repository Path

1.  Run the `pwd` command to get the absolute `REPO_PATH`.

### Step 4: Verify Variables and Execute the Script

1.  Print the values for `VERSION`, `GIT_COMMIT`, and `REPO_PATH` for the user to confirm.
2.  Execute the `push-tag-github.sh` script with the collected variables. The script will perform a dry run first.

    ```bash
    cd dev/tasks
    VERSION=${VERSION} GIT_COMMIT=${GIT_COMMIT} REPO_PATH=${REPO_PATH} ./push-tag-github.sh
    ```

### Step 5: Analyze Dry Run Output

1.  Examine the `stdout` and `stderr` from the previous command.
2.  **You must verify the following:**
    *   `stderr` is empty.
    *   `stdout` does **not** contain the words "ERROR", "failed", or "Aborted".
    *   The "Preparing local repository" section shows the correct `GIT_COMMIT` and `VERSION`.
    *   The "The following command will be executed for a dry run" section shows a `go run` command with the correct branch name (`--branch release_MAJOR.MINOR`) and other parameters.
3.  If the output looks correct and free of errors, proceed to the final step. Otherwise, report the errors to the user and stop.

### Step 6: Execute the Final Push

1.  If the dry run analysis was successful, execute the `go run` command from the dry run output with the `--yes=true` flag.

    ```bash
    cd dev/tasks
    # Example command extracted from dry run output
    go run . --remote upstream --branch release_1.133 --version-file version/VERSION --source /usr/local/google/home/${user}$/gitonborg/cnrm --push-options  -v=2 --yes=true
    ```
2.  Confirm to the user that the tag was pushed successfully.