---
name: test-terraform-fields
description: Guides through validating, testing, generating golden files (HTTP logs and KRM objects), running tests against mockgcp or real GCP, and aligning mock behavior.
---

# KCC Test Terraform Fields (Agentic-Friendly Guide)

This skill guides an automated agent or developer through testing, validating, and generating golden test assets (HTTP logs and KRM object configurations) when adding or modifying fields on Config Connector resources. It covers running tests against real GCP, recording the baseline behavior, running against MockGCP, and resolving mock discrepancies.

---

## 1. Structure of Resource Fixture Tests

KCC uses a golden file testing strategy for end-to-end (E2E) validation. The tests are defined under `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/`.

A complete test fixture directory contains:
- **`create.yaml`**: The primary KRM resource definition (what gets created first). Ensure the resource has unique labels/names.
- **`dependencies.yaml` (optional)**: Supporting resources (e.g. IAM policies, networks, service accounts) that the primary resource depends on.
- **`update.yaml` (optional)**: The KRM resource definition with updates applied after initial creation.
- **`_http.log`**: Golden HTTP/gRPC request/response traffic log generated during E2E reconciliation.
- **`_generated_object_[testname].golden.yaml`**: Golden file representing the final KRM object status/spec in the Kube API server.
- **`_generated_export_[testname].golden` (optional)**: Golden exported KRM representation.

---

## 2. Setting Up Test Cases

1. **Create Directory**: Create the directory `pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/`.
2. **Define KRM files**: Add `create.yaml`, and optionally `dependencies.yaml` and `update.yaml`.
3. **Verify Yaml Validation**: Ensure YAML files are valid Kubernetes manifests. Avoid any hardcoded project IDs or dynamic IDs (use placeholders if necessary, but KCC test runner resolves them).

---

## 3. Recording Ground Truth against Real GCP

To establish a baseline or update golden logs, you must run the tests against a real GCP project. This records actual HTTP/gRPC API interactions into the `_http.log` file.

### A. Authentication and Credentials
1. **Application Default Credentials (ADC)**:
   Ensure your local environment is authenticated with Google Cloud:
   ```bash
   gcloud auth application-default login
   ```
2. **IAM Privileges**:
   The active credentials must have sufficient privileges in the target GCP project. Depending on the resource type, permissions like `Project Owner`, `Storage Admin`, `DNS Administrator`, or custom resource-specific roles are required.
3. **Resource Quotas**:
   Verify that your target project has sufficient quota remaining for the resources you intend to create (e.g., compute CPUs, IP addresses, database instances).

### B. Environment Variables Configuration
Configure the following environment variables to control project routing, billing, and resource scopes:

| Environment Variable | Description | Default / Fallback |
| :--- | :--- | :--- |
| `GCP_PROJECT_ID` | The target GCP project to provision resources in. | Output of `gcloud config get-value project` |
| `TEST_BILLING_ACCOUNT_ID` | The billing account to associate with billing resources or projects created during E2E. | Dynamically queried from `GCP_PROJECT_ID` billing config |
| `TEST_FOLDER_ID` | The folder ID under which to run folder-level resource tests. | Parent folder ID of `GCP_PROJECT_ID` |
| `TEST_ORG_ID` | The organization ID under which organization-level tests run. | Ancestor organization of `GCP_PROJECT_ID` |

> [!NOTE]
> KCC tests dynamically substitute namespace project references with the actual target GCP project specified by these variables during execution.

### C. Running E2E Fixture Test Recordings
For standard E2E fixture tests under `pkg/test/resourcefixture/testdata/basic/`, use the `hack/record-gcp` script:
1. Run the script passing either the test name suffix or the package path:
   - **Using test name suffix**:
     ```bash
     hack/record-gcp <test_name>
     ```
   - **Using full test package path**:
     ```bash
     hack/record-gcp pkg/test/resourcefixture/testdata/basic/dns/v1beta1/<test_name>
     ```
     - **Example: Using test name suffix**:
     ```bash
     hack/record-gcp fixtures/dnsrecordsetbasic
     ```
   - **Example: Using full test package path**:
     ```bash
     hack/record-gcp pkg/test/resourcefixture/testdata/basic/dns/v1beta1/dnsrecordset
     ```
2. The script executes the tests with `E2E_GCP_TARGET=real`, `WRITE_GOLDEN_OUTPUT=1`, and records the traffic to `_http.log`.
3. **If the script fails** (e.g. due to permissions or an invalid default project ID like `foobar`), DO NOT skip this step. Ask the user for a valid GCP project ID to test against, and then run:
   ```bash
   GCP_PROJECT_ID=<project-id> hack/record-gcp <test_name>
   ```

### D. Running MockGCP Script Test Recordings
For MockGCP-specific script tests (located under `mockgcp/mockgcptests/`), use the python recording task:
1. Run the task specifying the mock test path:
   ```bash
   mockgcp/dev/tasks/record-gcp mocksql/testdata/instance/dr-replica
   ```
2. This runs the mock tests targeting real GCP to refresh golden script logs.

### E. Committing E2E Baseline
Once the recording successfully completes, stage and commit the generated `_http.log` and `_generated_object_*.golden.yaml` files:
```bash
git add pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/
git commit -m "Establish clean GCP golden logs for <testname>"
```

---

## 4. Running against MockGCP and Checking Diffs

To verify the mock implementation matches real GCP behavior:

1. **Run Mock Test**:
   - Run `hack/compare-mock <testname>`. 
     - **Example:** `hack/compare-mock tests-e2e-sql`
   - The test will execute using the MockGCP control plane.
2. **Check for Differences**:
   - Check if the command fails or if `git status` shows modifications to the golden files.
   - Run `git diff pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/`.
3. **Commit the Baseline Updates**:
   - If there are any differences (such as `selfLink` removal or minor formatting alignment), stage and commit these updates:
     ```bash
     git add pkg/test/resourcefixture/testdata/basic/<service>/<version>/<kind>/<testname>/
     git commit -m "Update golden logs for <testname> after mock comparison"
     ```

---

## 5. Aligning MockGCP with Real GCP (Troubleshooting & Alignment)

If the mock test fails or produces a diff against the baseline golden logs, apply the following strategies:

### A. Enum Mismatch (e.g., Short vs. Proto Enum Names)
- **Problem**: GCP REST API returns short enum values (e.g., `"OS_2022"`), but MockGCP protobuf definition expects full enum values (e.g., `"OS_VERSION_LTSC2022"`).
- **Solution**: Intercept the request and response bodies in the mock HTTP Mux (`mockgcp/mock<service>/service.go`) to translate these names.
- **Example in `mockgcp/mockcontainer/service.go`**:
  ```go
  // Intercept Request Body: OS_2022 -> OS_VERSION_LTSC2022
  if r.Body != nil {
      bodyBytes, err := io.ReadAll(r.Body)
      if err == nil {
          bodyBytes = bytes.ReplaceAll(bodyBytes, []byte(`"OS_2022"`), []byte(`"OS_VERSION_LTSC2022"`))
          r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
          r.ContentLength = int64(len(bodyBytes))
      }
  }
  ```

### B. Output-Only or Server-Generated Fields
- **Problem**: Fields generated by real GCP (e.g. project numbers, server-assigned IDs, URLs) differ in the mock response.
- **Solution**: Implement default value generation in the mock service CRUD handlers (`mockgcp/mock<service>/<resource>.go`). Ensure values match patterns expected by KCC (e.g., `projects/<project-id>/locations/<location>/...`).

### C. Normalizing Volatile Fields (timestamps, UUIDs, IPs)
- **Problem**: Dynamic values change on every execution.
- **Solution**: Add normalization rules in `mockgcp/mock<service>/normalize.go`.
- **CRITICAL**: The `Previsit` normalizer runs globally. To prevent corrupting golden files of unrelated services, **always scope the normalization rule to the specific service domain**:
  ```go
  func (s *MockService) Previsit(event mockgcpregistry.Event, replacements mockgcpregistry.NormalizingVisitor) {
      if !strings.Contains(event.URL(), "myservice.googleapis.com") {
          return
      }
      // Apply replacements here...
      // e.g. replacements.Replace(uuidRegex, "uuid-placeholder")
  }
  ```

### D. Duplicate Normalization Conflict (Zeros Placeholder Conflict)
- **Problem**: When running legacy E2E tests, the test runner applies a global normalizer that replaces server-generated identifiers in JSON response bodies with a string of zeros (e.g., `"000000000000000000000"`). If the per-service normalizer (`Previsit`) tries to match or replace this path, it can cause conflict errors or map multiple distinct resources to the same zeros placeholder.
- **Solution**:
  1. In the mock service's `normalize.go`, guard your `ReplaceStringValue` calls to ignore the placeholder:
     ```go
     if val == "000000000000000000000" {
         return
     }
     ```
  2. Extract resource IDs from the request URL path (e.g. matching `/changes/<id>`) instead of response bodies, as URL paths are not mutated by the legacy normalizer.

---

## 6. Pre-Submit Checks

Before finishing the task or proposing a PR, the agent must run formatting, generation, static analysis, and full CI validation checks locally to ensure zero CI/CD failures:

1. **Prepare PR and Regenerate Code**:
   - Run `make ready-pr` to ensure all manifests, Go client types, and code formatting are up to date:
     ```bash
     make ready-pr
     ```
2. **Mandatory CI/CD Presubmit Verification (CRITICAL)**:
   - To guarantee generated PRs pass GitHub Actions CI/CD checks cleanly, execute the primary validation scripts locally:
     ```bash
     dev/ci/presubmits/validate-generated-files
     scripts/validate-prereqs.sh
     ```
   - *(Note: `validate-generated-files` runs GitHub Actions workflow codegen, static config generation, CRDs, and mappers. `validate-prereqs.sh` validates formatting and generation).*
3. **Go Vet**:
   ```bash
   go vet ./...
   ```
4. **Verify Local Control Plane Webhooks**:
   - If envtest webhook startup fails with validation errors under new Kubernetes control plane versions, ensure `admissionReviewVersions` in `pkg/webhook/manifests.go` includes both `"v1"` and `"v1beta1"`.
5. **Verify CRD Field Coverage Checks**:
   - Run the API checks tests to ensure all new fields are either tested in the fixture tests. WARNING: New fields, in the vast majority of cases, should not be added to the  the exceptions list; this is verified by KCC testing as well. If a field needs to be added to the exceptions list you **must** provide a clear explanation as to why this is the case:
     - For **Beta** resources:
       ```bash
       WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTests
       ```
     - For **Alpha** resources:
       ```bash
       WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/... -run TestCRDFieldPresenceInTestsForAlpha
       ```
6. **Run CI/CD Group Presubmit Tests Locally**:
   - Locate and run the presubmit script under `dev/ci/presubmits/tests-e2e-fixtures-<service_name>` matching the resource's service name (e.g., `dev/ci/presubmits/tests-e2e-fixtures-container`) to ensure everything reconciles cleanly:
     ```bash
     dev/ci/presubmits/tests-e2e-fixtures-<service_name>
     ```
7. **Commit All Updated Artifacts and Generated Changes**:
   - Verify if any generated files (such as `mapper.generated.go`, GitHub Actions YAMLs, CRDs, Go clients, or exceptions) are modified using `git status` or `git diff`. Stage and commit them:
     ```bash
     git add -A
     git commit -m "chore: ensure pristine generated state and formatting to pass CI/CD presubmits"
     ```
8. **CI/CD & Golden File Traps (Gotchas)**:
   - **Accidental Binary Profile Artifacts (`heap.prof`)**: When running recorder or memory profile footprint tests (e.g., `TestProfileRecorderFootprint`), binary profile outputs like `heap.prof` may be written to the test directory (`cmd/recorder/pprof/.../heap.prof`). Always ensure these binary files are deleted and never committed to git.
   - **Selective Presubmit Harness & `WRITE_GOLDEN_OUTPUT=1` Trap**: Running specialized presubmit subsets (e.g., `test-pause`) with `WRITE_GOLDEN_OUTPUT=1` can inadvertently delete golden files belonging to skipped phases (e.g., `_generated_export_*.golden`). Ensure you only regenerate golden files using the appropriate comprehensive test scope, or inspect git diffs to revert unintended deletions.



