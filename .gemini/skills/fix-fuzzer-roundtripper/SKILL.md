---
name: fix-fuzzer-roundtripper
description: Fix fuzz-roundtrippers failing tests in k8s-config-connector by fetching the test logs, extracting the hint_for_agent snippet, adding f.Unimplemented_NotYetTriaged rules to the fuzzer file, and sending a PR. Use this when the user mentions a failing "fuzz-roundtrippers" test or provides a test URL to fix.
---

# Fix Fuzzer Roundtripper

## Overview
This skill provides the workflow to fix failing `fuzz-roundtrippers` tests in the k8s-config-connector project. These test failures typically occur when a new field is added to the GCP proto but is not yet supported or triaged in the Config Connector fuzzer. The test logs provide a `<hint_for_agent>` that tells us exactly what to ignore.

## Workflow

When asked to fix a `fuzz-roundtrippers` test, follow these steps:

### 1. Fetch the Test Logs
- The user will typically provide a URL to a failing GitHub Actions job (e.g., `https://github.com/GoogleCloudPlatform/k8s-config-connector/actions/runs/...`).
- Use the `web_fetch` tool to retrieve the content of the test logs from the provided URL.
- To ensure you find the hint, you might instruct `web_fetch` with a specific prompt: "Fetch the URL and extract the text within `<hint_for_agent>...</hint_for_agent>`."

### 2. Extract the Hint
- The logs will contain a hint in this format:
  `<hint_for_agent>Add \`f.Unimplemented_NotYetTriaged(".some.field.path")\` to the fuzzer for the proto type google.cloud.service.v1.Resource to mark this field as not yet triaged.</hint_for_agent>`
- Identify the exact code snippet to add (e.g., `f.Unimplemented_NotYetTriaged(".some.field.path")`).
- Identify the proto type (e.g., `google.cloud.networkmanagement.v1.ConnectivityTest`).

### 3. Locate the Fuzzer File
- Search the workspace for the fuzzer file corresponding to the proto type.
- These files are typically located in `pkg/controller/direct/<service>/<kind>_fuzzer.go` (e.g., `pkg/controller/direct/networkmanagement/connectivitytest_fuzzer.go`).
- To be absolutely sure, use `grep_search` to find the file containing the annotation comment: `// proto.message: <proto_type>`.
  - Example search pattern: `// proto.message: google.cloud.networkmanagement.v1.ConnectivityTest`

### 4. Update the Fuzzer File
- Read the located fuzzer file.
- Inside the fuzzer function (e.g., `func fuzz...`), look for existing `f.Unimplemented_NotYetTriaged` calls or similar fuzzer rules.
- Add the new `f.Unimplemented_NotYetTriaged(...)` line provided by the hint.
- Place it neatly with the other ignored fields.
- Ensure the code is syntactically valid Go code.

### 5. Commit and Create PR
- Create a new Git branch for the fix: `git checkout -b fix-fuzzer-<kind>` (e.g., `git checkout -b fix-fuzzer-connectivitytest`).
- Commit the changes. The commit message must include a link to the failing test URL provided by the user.
  - Example commit message:
    ```
    Fix fuzz-roundtripper for <kind>

    Adds f.Unimplemented_NotYetTriaged for <field_path>.
    Fixes failing test: <test_url>
    ```
- Use the `run_shell_command` tool to run `git commit -a -m "..."`.
- Push the branch to the repository using `git push -u origin HEAD`.
- Use the GitHub CLI to create a Pull Request. Run a shell command like:
  `gh pr create --title "Fix fuzz-roundtripper for <kind>" --body "Fixes failing test: <test_url>"`
- Ensure you verify the Git status and diffs along the way to avoid committing unrelated changes.
