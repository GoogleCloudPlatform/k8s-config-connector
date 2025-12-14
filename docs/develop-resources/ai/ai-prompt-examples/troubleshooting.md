# 5. Troubleshooting & Tips (Lessons Learned)

### Recommended Usage
> "Read `scenarios/troubleshooting.md` if you encounter errors with Envtest, GKE resources, or client generation."


### Test Environment
*   **Missing kubebuilder assets (etcd/apiserver):** If tests fail with "unable to start control plane itself", set `KUBEBUILDER_ASSETS` to the bin path (e.g., `.../bin/k8s/1.29.5-linux-amd64`).
*   **Golden Logs:** To capture `_http.log`, set `WRITE_GOLDEN_OUTPUT=1` and `E2E_GCP_TARGET=real`. The file is written only when the test destroys the resource successfully.

### Resource Specifics
*   **ContainerCluster (GKE):**
    *   **Auto Upgrade Requirement:** If `releaseChannel: REGULAR` is set (often default in tests), and you enable `nodeAutoProvisioning`, you **MUST** also set `autoProvisioningDefaults.management.autoUpgrade: true`. Failing to do so causes a `400 Bad Request`.
    *   **NetworkPolicy:** If enabling `enableCiliumClusterwideNetworkPolicy`, you must also ensure `datapathProvider` is set to `ADVANCED_DATAPATH`.
    *   **K8s Tokens Via DNS:** If enabling `enableK8sTokensViaDNS`, you must also ensure the DNS endpoint is enabled (e.g., via `controlPlaneEndpointsConfig`). Failure to do so results in "K8s token via DNS cannot be enabled without enabling DNS endpoint".

### PR Hygiene
*   **Bloated PRs (File Count Watchdog):**
    *   **New Field:** Should be **< 15 files**.
    *   **New Resource:** Should be **~20 files**.
    *   **Promotion:** Should be **~30-40 files**.
    *   **Action:** If your PR has 200+ files, **STOP**. You have accidentally committed generated noise. Revert unrelated files.
*   **Unrelated Lint Errors:** `make ready-pr` might fail due to lint errors in files you didn't touch (e.g., auto-generated files).
    *   **Fix:** `git restore <unrelated_file>` or `git checkout -- <unrelated_file>` to revert them. DO NOT try to fix lint errors in files you don't own.
*   **Proto Noise:** If `make generate` produces 10k+ lines of changes in `mockgcp/` for services you aren't touching, **REVERT THEM**. This is "noise" and will block review.

### Client Generation Failures
*   **Panic in `deepcopy-gen`:** Often caused by using `interface{}` or `any` in API types.
    *   **Fix:** Use `apiextensionsv1.JSON` instead of `any`.
*   **Script Failures / Opaque Errors:** If `make generate-go-client` fails without a clear error, it is often due to syntax errors in your `_types.go` file (missing imports, typos).
    *   **Fix:** **ALWAYS** check that your code compiles *before* generating. Run `go build ./apis/<service>/<version>/...` to find the syntax error.
*   **Missing Imports:** `generate-types-file` might miss imports. Manually add them to `_types.go` if compilation fails.

### General
*   **Split PRs (Atomic Commits):** If your PR review is stuck or complex, split it! One field per PR is much faster to review than a bundle of 5 fields.
*   **Hidden Fields:** Some fields in Terraform provider might be defined in schema but ignored in logic. Always check `resource<Name>Create` and `Read` functions to ensure mapping exists.
