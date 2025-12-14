# 2. New Field

### Recommended Usage (Starting Prompt)
> "Read `scenarios/new-field.md` and follow the prompt to add the `<Field>` field to `<Resource>`."


**Use Case:** Adding a missing field to an existing **Direct** resource.

**Reference Example:** [PR #5646 (Firestore TTL)](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5646) is a perfect example of a standard "Add Field" PR. It demonstrates the expected file changes (API types, controller mapping, and verification with real GCP golden logs).

**Prompt:**
> "Add the `<fieldName>` field to the `<Kind>` resource in `<group>/<version>`.
>
> **Constraints:**
> *   **Documentation:** Refer to `docs/develop-resources/ai/scenarios/new-field.md`.
> *   **Troubleshooting:** Refer to `scenarios/troubleshooting.md` for common errors (Proto Noise, lint failures, Envtest).
> *   **Git Author:** Use `git commit --author="<User Name> <user@google.com>"`.
> *   **DeepCopy Safety:** Do NOT use `interface{}` or `any`. Use `apiextensionsv1.JSON` if needed.
> *   **NOISE CONTROL (Critical):** Do NOT include unrelated `mockgcp/` proto updates. If `make generate` adds thousands of lines of unrelated proto changes (e.g., from other services), **REVERT** those changes. Only commit changes relevant to your resource.
> *   **Field Naming:** Ensure JSON tag matches the GCP field name (camelCase).
> *   **Atomic PRs (CRITICAL):** Submit **ONE PR per field**. Do NOT bundle multiple fields unless they are strictly dependent. This dramatically speeds up review and verification.
> *   **License Headers:** Ensure all new files have the standard Apache 2.0 license header with the current year (2025).
> *   **Code Cleanliness:** Address all linting issues (including `typecheck`) by fixing the root cause. Do NOT use `//nolint` directives to bypass errors unless absolutely necessary and documented.
>
> **Workflow:**
>     1.  **Update API:** Run `update-types` or manually add field to `_types.go`.
>     2.  **Verify Gen:** Run `make generate` and `make generate-go-client`.
>         *   **CHECK:** If unrelated files (e.g., `mockgcp/.../other_service.pb.go`) changed, revert them (`git checkout -- <file>`).
>         *   **File Count Check:** Expect **< 15 files**. If > 20, you likely have unrelated "noise".
>     3.  **Update Mapper:** Regenerate mapper or add manual mapping.
>     4.  **Update Mock:** Ensure MockGCP handles the field (if applicable).
>     5.  **Implement Logic:**
>         *   **TF-Based Resources:** Explicitly implement `Create`, `Read`, and `Update` logic in the resource file (e.g., `resource_*.go`).
>         *   **Update Logic:** For boolean fields or partial updates, ensure you use `d.HasChange("field_name")` checks to prevent overwriting other fields or missing updates.
>         *   **Nested Fields:** Verify expansion/flattening logic handles the new field correctly (e.g., inside `expandManagement` or similar helpers).
>     6.  **Verify:** Update `create.yaml`/`update.yaml` to test the field. Run E2E tests:
>         ```bash
>         # Real GCP (Capture Log)
>         export E2E_GCP_TARGET=real WRITE_GOLDEN_OUTPUT=1
>         go test -v -run TestAllInSeries/... ./tests/e2e
>         ```
>     7.  **Record:** Re-record golden files (`_http.log`) and commit them.
