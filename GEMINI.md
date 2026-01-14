# KCC Development Guide (GEMINI.md)

> **Mission**: To achieve "Zero Friction" development by internalizing the project's unwritten rules, avoiding known pitfalls (like the "Generator Trap"), and strictly adhering to the "Golden PR" standard.

## 1. Scope & Context
*   **Terraform-Based Resources** (e.g., `ContainerCluster`): This guide primarily addresses the maintenance of these "Legacy" resources, where the "Surgical Workflow" is critical to avoid destructive generation.
*   **Direct Resources**: For modern "Direct" resources, prioritize the official [AI-Assisted Workflow](docs/develop-resources/ai/README.md).

## 2. The "Golden PR" Standard
Based on "Golden" PRs (#4941, #5815) and [Official MockGCP Standards](docs/develop-resources/deep-dives/1-add-mockgcp-tests.md), a PR must follow this structure to minimize review friction.

### 2.1 The "Atomic Scope" Rule (One Field Per PR)
*   **Rule**: Do not bundle multiple unrelated feature additions into a single PR.
*   **Why**: Large PRs get stale and are harder to review. Atomic PRs merge faster.
*   **Exception**: If fields are strictly dependent on each other, they can be grouped.

### 2.2 The "Contiguous Commit" Rule
Reviewers (**and the official docs**) explicitly request that Real GCP logs and Mock GCP logs are committed sequentially to facilitate diffing.

**Ideal Commit Sequence:**
1.  `fix(tf): patch resource_container_cluster.go` (Logic)
2.  `feat(mock): update mockgcp/mockcontainer` (Simulation)
3.  `fix(crd): surgical update to crd` (Schema)
4.  `test(real): update goldens from Real GCP` (Evidence A)
5.  `test(mock): update goldens from Mock GCP` (Evidence B - Should match A)

### 2.3 The PR Description Template
*   **Title**: `feat(resource): add support for <field>`
*   **Release Note**: Clear, user-facing summary.
*   **Tests**: Explicitly listed.
    *   *Correct*: "Ran `hack/record-gcp` against Real and `hack/compare-mock` against Mock."

## 3. The "Surgical" Workflow (Anti-Friction)
The standard tools (`make manifests`) are "blunt instruments" that can destroy manual fixes in Terraform-based resources.

### 3.1 The Generator Trap
*   **Danger**: `scripts/generate-crds` often fails to pick up custom Terraform patches. Running `make manifests` will **wipe out** your manual CRD edits.
*   **Solution**: Patch the CRD manually (`config/crds/resources/...`) unless you are certain the generator is configured for your new fields.

### 3.2 The "Surgical" Loop
1.  **Edit**: Modify `third_party/...` (Go logic).
2.  **Lint (Early)**: Run `dev/tasks/run-linters` immediately.
3.  **Patch CRD**: Manually add new fields to the CRD YAML.
4.  **Verify**: Run `hack/compare-mock` immediately.

## 4. The 5-Point Verification Protocol
Pass these gates locally before pushing.

1.  **Mock Verification**:
    ```bash
    hack/compare-mock fixtures/<category-or-path>
    ```
    *   *Ref*: [Add MockGCP Tests](docs/develop-resources/deep-dives/1-add-mockgcp-tests.md)

2.  **Real Verification**:
    ```bash
    hack/record-gcp fixtures/<category-or-path>
    ```
    *   *Goal*: Generate the definitive `_http.log`.

3.  **Unit Tests & API Checks**:
    ```bash
    ./scripts/unit-test.sh
    ```
    *   *Goal*: Pass `TestCRDsAcronyms` and `TestCRDFieldPresenceInTests`.
    *   *Ref*: [API Conventions](docs/develop-resources/api-conventions)

4.  **CRD Schema Check**:
    *   Ensure Go types (`pkg/clients/generated`), CRD YAML, and Mock implementations align.
    *   *Rule*: Use `*string` (pointers) for optional fields, not raw types.

5.  **Lint Check**:
    *   `dev/tasks/run-linters`

## 5. Official Resource Map
Consult these official documents for deep dives:

| Topic | Doc Path | Why Read? |
| :--- | :--- | :--- |
| **New Fields** | [scenarios/new-field.md](docs/develop-resources/scenarios/new-field.md) | Standard procedure for adding fields (Focuses on "Direct"). |
| **Validation** | [api-conventions/validations.md](docs/develop-resources/api-conventions/validations.md) | Rules for `+required`, `omitempty`, and CEL validation. |
| **MockGCP** | [deep-dives/1-add-mockgcp-tests.md](docs/develop-resources/deep-dives/1-add-mockgcp-tests.md) | **Critical**: Defining the `create.yaml` / `update.yaml` suite structure. |
| **Pointers** | [reviewing/REVIEWAGENT.md](docs/reviewing/REVIEWAGENT.md) | Why `*int` is preferred over `int`. |

## 6. Friction Removal Cheatsheet

| Friction Point | Solution |
| :--- | :--- |
| **"make manifests" deleted my code** | Use **Surgical Mode**: Edit CRD directly. |
| **"unknown field" in E2E logs** | Missed CRD patch. Check `config/crds/...`. |
| **CI reports "Unapproved"** | You need `lgtm` labels. |
| **Diff between Real/Mock** | **Never** change `create.yaml` input between runs. |
| **Linter "import order"** | Use `goimports` locally. |

## 7. CI Troubleshooting / Advanced Diagnosis

### 7.1 Retrieving Logs for "In Progress" Runs
When a CI run is "In Progress", `gh run view --log` is often blocked. Use `gh api` to bypass this and fetch logs for specific failed jobs.

**Protocol:**
1.  **List Jobs** (Find the `databaseId` of the failed job):
    ```bash
    gh run view <RUN_ID> --json jobs > ci_jobs.json
    # Inspect JSON to find the job with "conclusion": "failure"
    ```
2.  **Fetch Raw Logs** (Using the Job ID):
    ```bash
    gh api repos/GoogleCloudPlatform/k8s-config-connector/actions/jobs/<JOB_ID>/logs > job_failure.log
    ```
3.  **Analyze**:
    ```bash
    grep -C 5 "FAIL" job_failure.log
    ```

### 7.2 Common Failure Signatures
| Error Signature | Likely Cause | Fix |
| :--- | :--- | :--- |
| `FAIL: unexpected diff in testdata/exceptions/acronyms.txt` | Field naming violation (e.g., `Dns` vs `DNS`, `Id` vs `ID`). | Rename field in Go struct to match K8s conventions (e.g., `ViaDNS`). |
| `Manifests must be regenerated` | `config/crds` out of sync with Go types. | Run `make manifests` (Be careful of "Surgical" mode!). |
| `field not declared in schema` | Missing field in CRD. | Patch CRD YAML manually or fix generator config. |
| `error retrieving userinfo... scope?` | Missing OAuth scopes in test environment. | Ensure `hack/record-gcp` has correct credentials/scopes. |

