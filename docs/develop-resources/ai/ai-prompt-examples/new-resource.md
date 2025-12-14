# 1. New Resource (Direct)

### Recommended Usage
> "Read `scenarios/new-resource.md` and follow the prompt to implement the `<Resource>` resource."


**Use Case:** Adding a brand new resource or migrating a `v1alpha1` Terraform/DCL resource.

**Reference Example:** [PR #5636 (TagsTagBinding)](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/5636) demonstrates creating a new Direct resource using the Standard SDK strategy.

**Strategies:**
*   **Standard SDK:** Use if the resource exists in `google.golang.org/api` (GCP Go SDK).
*   **Isolated Mock:** Use if the resource is missing from the SDK (requires generated protos).

**Prompt:**
> "I need to implement the `<Resource>` resource for the `<Service>` CIG.
>
> **Constraints:**
> *   **Documentation:** Refer to the [AI Development Guide](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/docs/develop-resources/ai/README.md).
> *   **Troubleshooting:** Refer to `scenarios/troubleshooting.md` for common errors (Proto Noise, lint failures, Envtest).
> *   **Strategy:** Choose one:
>     *   `standard` (Use GCP Go SDK).
>     *   `mockgcp` (Use Isolated Mock if SDK support is missing).
> *   **Stacked PRs:** Split execution into 3 stacked PRs:
>     1.  **PR 1:** API Types & MockGCP (Phase 1).
>     2.  **PR 2:** Mapper & Controller (Phase 2).
>     3.  **PR 3:** Final Implementation & Tests (Phase 3 & 4).
> *   **Git Author:** Use `git commit --author="<User Name> <user@google.com>"`.
> *   **PR Status:** Mark all PRs as **Draft** initially. Do not request review until automated tests (CI) pass.
> *   **Pre-Review Checks:** Run `make ready-pr` (includes `lint`, `manifests`, `ensure`, `fmt`) locally.
>
> **Critical Implementation Rules:**
> *   **DeepCopy Safety:** Do NOT use `interface{}` or `any`. Use `apiextensionsv1.JSON` if needed.
> *   **Hidden Fields:** Verify if Terraform fields exist in schema but are ignored in `Create`/`Read` logic. Manually map them if needed.
> *   **GKE Constraints:** If copying `nodeAutoProvisioning`, ensure `autoUpgrade: true` is set for REGULAR release channel.
> *   **Client Generation:** If `make generate-go-client` fails or panics, it is often due to syntax errors in `_types.go` or missing imports. Verify the API types compile *before* running generation.
> *   **Field Naming:** JSON tags MUST be `camelCase`. Watch out for acronyms (e.g., `enableNestedVirtualization`, NOT `enableNestedvirtualization` or `EnableNestedVirtualization`).
> *   **License Headers:** Ensure all new files have the standard Apache 2.0 license header.
> *   **Code Cleanliness:** Fix all lint errors. Do NOT use `//nolint` directives.
>
> **Workflow:** Execute the following 4-Phase Plan. **Create Draft PRs after Phase 1 and 2.**
>
> **Phase 1: Foundation (API & Mock)**
> 1.  (MockGCP Strategy Only) Import/Patch Protos and Generate Mock Service in `mockgcp/`.
> 2.  Generate KRM API scaffolding in `apis/<service>/<version>/<resource>_types.go`.
> 3.  **VERIFY:** Run `make generate` AND `make generate-go-client` to ensure types are valid.
>
> **Phase 2: Implementation (Mapper & Controller)**
> 1.  Generate Mapper using `controllerbuilder`.
> 2.  Implement Direct Controller in `pkg/controller/direct/<service>/`.
> 3.  Register controller in `pkg/controller/direct/register/register.go`.
>
> **Phase 3: Testing**
> 1.  Create `create.yaml` test fixture.
> 2.  Run E2E tests against MockGCP.
>
> **Phase 4: Verification**
> 1.  Run `make ready-pr`.
>     *   **File Count Check:** Expect **~15-25 files** in total. If you have >30, check for unrelated proto changes or accidental vendor updates.
> 2.  **Real GCP Verification:**
>     ```bash
>     export E2E_GCP_TARGET=real WRITE_GOLDEN_OUTPUT=1
>     go test -v -run TestAllInSeries/fixtures/basic/<service>/<version>/<resource> ./tests/e2e
>     ```
> 3.  Verify DeepCopy, Lint, and Compilation.
>
> Please start with **Phase 1**."

---

## 4. Migrate Beta Resource (Complex)

**Use Case:** Migrating a `v1beta1` Terraform/DCL resource to Direct.
**Note:** This requires strict backward compatibility preservation. Please file a GitHub issue or consult the team before proceeding.
