# 3. Promote to Beta

### Recommended Usage
> "Read `scenarios/promote-beta.md` and follow the prompt to promote `<Resource>` to v1beta1."


**Use Case:** Promoting a **Direct** resource from `v1alpha1` to `v1beta1`.

**Reference Example:** [PR #4897 (BigQueryReservationReservation)](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/4897) demonstrates the promotion workflow, including version copying, storage version marking, and test updates.

**Prompt:**
> "Promote the `<Kind>` resource from `v1alpha1` to `v1beta1`.
>
> **Constraints:**
> *   **Documentation:** Refer to `docs/develop-resources/ai/scenarios/alpha-to-beta.md` and `docs/ai/how-to-promote-resource.md`.
> *   **Troubleshooting:** Refer to `scenarios/troubleshooting.md` for common errors (Proto Noise, lint failures, Envtest).
> *   **Git Author:** Use `git commit --author="<User Name> <user@google.com>"`.
>
> **Workflow:**
>     1.  **Copy Version:** Copy `v1alpha1` API directory to `v1beta1`.
>     2.  **Storage Version:** Mark `v1beta1` as `// +kubebuilder:storageversion`.
>     3.  **Update Controller:** Point controller to `v1beta1` types.
>     4.  **Update Tests:** Update fixtures to `apiVersion: .../v1beta1`.
>     5.  **Regenerate:** Run `make manifests` (generate CRDs) and `make generate-go-client`.
>         *   **File Count Check:** Expect **~30-40 files** (new API version + conversion + registration). If > 100, checking for massive unrelated proto regeneration.
>     6.  **Verify:** Run E2E tests and `dev/tasks/generate-resource-report`.
>         ```bash
>         export E2E_GCP_TARGET=real WRITE_GOLDEN_OUTPUT=1
>         go test -v -run TestAllInSeries/... ./tests/e2e
>         ```"
