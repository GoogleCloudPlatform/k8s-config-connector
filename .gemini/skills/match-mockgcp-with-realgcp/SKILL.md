# Skill: Match Mock behaviour with real GCP api

This skill provides a structured workflow for matching the mock {service}{resource} behaviour with the real GCP.

When the golden tests for K8s Config Connector mock output diverge from real GCP output, we need to inspect the discrepancies and fix either the mock implementation or the normalizers. This ensures that `hack/compare-mock` runs cleanly and accurately represents GCP API contracts.

## Critical Rules

*   **Real GCP Baseline Required**: You must always generate the initial baseline `_http.log` by running `hack/record-gcp` against a real GCP project and commit it first.
*   **Basic Test Case Naming**: Ensure that the basic test case directory/folder names always contain the resource kind name (e.g., use `computesubnetwork-reservedinternalrange` instead of `reservedinternalrange`).
*   **Handling record-gcp Failures**: If `record-gcp` fails, the bot should attempt to debug and fix the test case configuration (e.g., resolving IP address range conflicts, API/service enablement, or configuration schema issues). If it still fails after 3 distinct retry attempts, the bot must halt, escalate the failure to the human reviewer, and clearly explain the blocker/error in a comment.
*   **No New Missing Fields**: We should not add new missing fields to the exceptions file (`tests/apichecks/testdata/exceptions/missingfields.txt`). If fixing or renaming an existing test case exposes/causes more missing fields, you must extend the test case configuration or add a new test case (or more than one if needed) to cover those missing fields.
*   **Do NOT Generate Golden Logs from Mock**: You are strictly forbidden from generating or updating `_http.log` files against the mock in this step. You MUST first establish a baseline against `real` GCP with `E2E_GCP_TARGET=real E2E_KUBE_TARGET=envtest` with `WRITE_GOLDEN_OUTPUT=1`.
*   **Do NOT Manually Edit HTTP Logs**: You are strictly forbidden from using text-editing tools (like `replace`, `write_file`, or `sed`) to modify any `_http.log` files. The ONLY valid way to update these files is by executing the test scripts (`hack/record-gcp` or `hack/compare-mock`) via the shell.
*   **Provide Timestamps**: When starting and finishing running `hack/record-gcp` command, you must capture and provide the exact timestamps in PDT in the summary comment after addressing the comments so that the reviewer can check the GCP logs to verify the work.
*   **Do NOT Modify Test Cases Post-Recording**: Do not update/modify the test case configuration (e.g., changing location/zone) after you have committed the real GCP logs. If you must change the test case, you must run `hack/record-gcp` again to record the real GCP logs.
*   **WHENEVER A TEST CASE IS UPDATED, WE MUST RECORD REAL GCP LOGS AGAIN**: If you make any modifications to a test case configuration, manifest files (such as `create.yaml`, `update.yaml`, or `dependencies.yaml`), or the controller's runtime mapping configuration, you **MUST** run the test case against real GCP (`hack/record-gcp`) to regenerate the authentic `_http.log` baseline before comparing or committing any mock log changes. Do not attempt to manually edit the logs or bypass recording live traffic.
*   **Record and Align All Test Fixtures**: Do not record/align only a single test case/fixture (e.g., `containernodepool`). You must discover and process all test fixtures under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` (including variants like `containernodepool-*`, `containernodepoolremovetaint`, etc.) to prevent partial coverage or out-of-sync logs.
*   **Do NOT Update Legacy Controller**: Do not update or modify the legacy/Terraform-based controller (such as code under `pkg/krmtotf/references.go`) unless absolutely necessary. Keep the legacy controller logic unchanged and avoid modifications to common or shared legacy framework code.
*   **Focus on Target CRUD Requests**: When aligning the HTTP logs, we only care about the requests made to the CRUD endpoint of the target resource. Non-target/dependent resource requests in the log do not need to be matched as strictly.

## Workflow

### Step 0: Remove from Ratcheting Exclusions (MANDATORY)

Before running the test cases against real or mock GCP, you **MUST** ensure the target resource is removed from the ratcheting exclusion list in `tests/e2e/ratcheting.go`. This enables the re-reconciliation test step, which is a fundamental use case KCC resources must support.

1. Open `tests/e2e/ratcheting.go`.
2. Locate the function `ShouldTestRereconiliation`.
3. Locate the `switch` statement that checks `primaryResource.GroupVersionKind()`.
4. If there is a `case` block for your target resource's `GroupKind`, remove that `case` line from the switch statement.

### Step 1: Discover and Record GCP logs

1.  **Discover All Test Fixtures for Resource**:
    Locate all test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` (e.g. `containernodepool`, `containernodepool-update`, `containernodepoolremovetaint`, etc.).

2.  **Fast vs. Slow Resource Execution Strategy**:
    Determine whether your target resource is fast or slow to reconcile against real GCP. Resources like `ContainerNodePool`, `ContainerCluster`, `SQLInstance`, and `ComposerEnvironment` are slow resources (>10 minutes per test), whereas most others are fast (<10 min per test).
    - **Fast resources (<10 min per test)**:
      - Run and record all fixtures for the resource in a single command using:
        ```bash
        hack/record-gcp "fixtures/<kind_lowercase>"
        ```
    - **Slow resources (>10 min per test)**:
      - Do NOT run all test fixtures sequentially in a single command.
      - Execute each individual test fixture in parallel (e.g., triggering parallel background processes or subagents with `hack/record-gcp "fixtures/^<testname>$"`).
      - **Parallel Directory Collisions**: `hack/record-gcp` cleans and writes to `artifactz/realgcp` by default. When running multiple invocations in parallel, you **MUST** provide a unique `ARTIFACTS` path for each invocation (e.g., `ARTIFACTS=artifactz/realgcp-<testname> hack/record-gcp "fixtures/^<testname>$"`) to prevent folder collision/cleanup issues.
    *   **Troubleshooting Service Not Enabled**: If `hack/record-gcp` fails because a GCP service is not enabled (e.g., error mentions that the API is disabled or has not been used in the project before), enable the service using `gcloud` and try again:
        ```bash
        gcloud services enable <service-name>.googleapis.com
        ```
        *(For example: `gcloud services enable compute.googleapis.com` or `gcloud services enable run.googleapis.com`)*

3.  **Verification Criteria for Real GCP Recording (MANDATORY)**:
    Checking `git status` or top-level `PASS`/`FAIL` exit codes can be misleading (e.g. 0-diff is valid for unaffected fixtures, unmatched regex gives a false `PASS`, and a successful recording can mark the subtest as `FAIL` because it wrote a new golden log).
    Instead, verify that each discovered fixture completed its full lifecycle by checking stdout for:
    1. `=== RUN   TestAllInSeries/fixtures/<testname>` (confirms the subtest matched and started).
    2. Resource reached `Ready` (`status.condition.status: True`) for create (and update, if present).
    3. Deletion completed (`Done waiting for resource to delete`).
    4. `wrote updated golden output to .../_http.log` was logged (or `_http.log` was updated/written on disk).
    5. No fatal errors (`t.Fatalf` / timeout / permission errors) aborted the test before cleanup.

4.  Please commit the real gcp logs (`_http.log`) generated by running the command.

Important:
* It is important to commit the files modified by running realgcp tests in its own commit.
* This is for the human reviewer to compare the diff in the test artifacts when running real and mockgcp.


### Step 2: Match mockgcp behaviour with real GCP
1.  **Run compare-mock**:
    Run the mock comparison tests for all fixtures of the resource in a single command (MockGCP runs in-memory and completes in seconds, so no parallel execution split or Fast/Slow execution strategy is needed for mock tests):
    ```bash
    hack/compare-mock "fixtures/<kind_lowercase>"
    ```
    This command will execute the tests against MockGCP and automatically generate/update the golden files (such as `_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) under the test case directories.

2.  Iteratively fix discrepancies in the mock implementation or `normalize.go`.

Tips for fixing discrepancies:
1. Diff `_http.log` (real) against `_http_mock.log` (mock) for the target fixture (e.g. using `diff` or `go test ./pkg/test/resourcefixture -run TestGoldenLogAlignment`).
2. **Missing Default Values (Field in Real but NOT in Mock)**:
   - For missing default values (e.g. `<missing> => REGIONAL`), add or update `populateDefaultsFor<Resource>` in `mockgcp/mock<service>/<kind>.go`. Ensure it is called on both mutation (`Create`/`Insert`) and retrieval (`Get`/`List`) methods.
   - If the field is **missing from the proto**, leave a `TODO` comment noting the missing proto field instead.
3. **Unexpected Values (Field in Mock but NOT in Real)**:
   - Check if the field was added to align with another test case. If real GCP includes it in other scenarios but omits it here, it is possibly **conditionally defaulted** (or was accidentally hardcoded).
   - If verified to be conditional, implement dynamic/conditional defaulting in MockGCP and add a comment explaining the condition. Otherwise, remove or fix the hardcoded default.
   - If the field is never returned by real GCP in any scenario, remove/omit it from MockGCP.
4. **Empty / Zero-Value Fields in Real**: If real GCP outputs an empty value (e.g. `""`, `0`, `false`, `[]`) that MockGCP drops due to proto3 serialization (unset vs zero-value equivalence):
   - Try explicitly setting the field in MockGCP.
   - If proto3 JSON marshaling still omits it, leave a `TODO` comment clarifying that it is an expected proto3 zero-value limitation.
5. **Input vs Output Canonicalization**:
   - When real GCP accepts short names or unformatted references but returns canonicalized formats (e.g. converting `/zones/` to `/locations/`, expanding relative names to full URIs, or normalizing CIDRs), implement the transformation in MockGCP during `Create`/`Insert` so the stored state matches GCP's canonical output.
6. **Volatile / Dynamic Values in Self-Links & Long Names**:
   - For dynamic segments embedded in resource names, self-links, or generated resources (e.g., auto-generated hashes in `/instanceGroupManagers/gke-.*-grp`, `/forwardingRules/${forwardingRuleID}`, `${uniqueId}`, `${projectNumber}`), normalize them to standard placeholders in `mockgcp/mock<service>/normalize.go`.
   - **Critical Rule**: Always scope `Previsit` normalization in `normalize.go` to the service URL (e.g. `strings.Contains(event.URL(), "<service>.googleapis.com")`).
7. **Exempt Specific URL Prefixes**:
   - If real GCP returns full URL prefixes (e.g. `//certificatemanager.googleapis.com/` or `https://www.googleapis.com/...`) while MockGCP uses relative paths, strip the prefix in the service's normalizer, keeping the rule strictly scoped to that specific resource/field.
8. **Exempt API Version Differences (Specific / Hardcoded Only)**:
   - If real GCP and MockGCP use different API versions for a specific service (e.g., `aiplatform.v1beta1` vs `aiplatform.v1`), normalize it using explicit, hardcoded string replacements for that specific service only. Never use broad wildcards that could mask unintended API version mismatches.
9. **Iterate and Verify**: Re-run `hack/compare-mock "fixtures/<kind_lowercase>"` (or `hack/compare-mock "fixtures/^<testname>$"`) to update `_http_mock.log` and the golden object files (`_generated_object_*.golden.yaml`), and verify log alignment against `_http.log`. Repeat until all discrepancies are resolved and the tests pass.
10. **Operation Metadata**: Minor operation polling metadata (e.g. intermediate LRO statuses like `done: false`) can be simplified in MockGCP as long as the final resource state matches.

### Step 3: Multi-Fixture Verification Requirement (MANDATORY)

You **MUST** verify and ensure that all test directories associated with your target resource are completely verified.
1. Ensure that all discovered test subdirectories under `pkg/test/resourcefixture/testdata/basic/<group>/<api_version>/<kind_lowercase>/` have their golden files (`_generated_object_*.golden.yaml`) and mock HTTP logs (`_http_mock.log`) successfully generated. Note that a 0-diff in `git status` is perfectly valid if the resource doesn't contain any service-generated, randomized values in the responses or our normalizers have overwritten those values.
2. Commit and push all regenerated and matched log/golden files for every discovered fixture in your Pull Request.



