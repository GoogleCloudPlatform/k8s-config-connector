# Migrate Terraform/DCL-based Alpha to direct Alpha

When migrating an existing Terraform or DCL-based Alpha resource to a Direct Controller, preserving existing user configurations is critical.

---

## The 100% Field Parity Invariant (Mandatory Migration Invariant)

> [!CAUTION]
> **A Direct Controller migrating an existing TF or DCL resource must support every single field present in the legacy CRD.**
> While greenfield Alpha resources generally allow breaking API changes, **brownfield migrations do NOT**.
> **Dropping a field causes KCC to wipe that configuration from GCP on the very first reconcile loop after upgrade**, resulting in silent data loss or service disruption for existing users.
> Every field in the legacy CRD must be mapped and supported in the direct types, mappers, and controller.

---

## Add MockGCP tests
 
The MockGCP provides the test coverage and helps you develop the direct resource in the follow up steps. *The reason we need the Terraform/DCL-based Alpha resource here is that it is a pre-existing option to help you develop the Mock GCP faster; you don't need or rely on the Terraform/DCL based Alpha resource in the follow up steps.* 

Follow [deep-dives Step 1](../deep-dives/1-add-mockgcp-tests.md)

### PR Reviews

* The 1st git-commit shall contain `create.yaml` and `update.yaml` files with exactly the same config. 
* The 2nd git-commit shall have `update.yaml` changing all the mutable fields, reflected in the git-diff based on the first commit.
* The 3rd git-commit (or N-th git-commit) shall run against real GCP. The generated `_generated_object_<resource>.golden.yaml` and `_http.log` is the golden log reflecting the real GCP. The `_generated_object_<resource>.golden.yaml` matches the `spec` in `update.yaml`.
* The 4th git-commit (or N+1-th git-commit) shall run against Mock GCP. The git-diff shows `_generated_object_<resource>.golden.yaml` is unchanged, and the `_http.log` is as much like the real GCP as possible. 

* If the resource has dependencies (`dependencies.yaml`), we suggest covering all the referenced fields. If the dependency resource does not have mockGCP, you need to implement those dependencies' MockGCP methods as well.

---

## Add API

Follow [deep-dives Step 2](../deep-dives/2-define-apis.md)

The PR shall contain the types and deepcopy codes. It shall follow the direct resource [recommended styles and conventions](../api-conventions).

* **Enforce 100% Field Parity**: Ensure that every single field present in the existing legacy CRD is supported. Do not remove or omit existing fields.
* Make sure the generated code is placed under `./apis/<service>/v1alpha1`. If not, please check your flags.
* Add `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` to the API tag, to continue using Terraform-based controllers while developing. *This relies on the existing Terraform-based controller to validate your API.*

### PR Reviews

* If the config/crds is not changed after running `dev/tasks/generate-crds`, we require `the _generated_object_<resource>.golden.yaml` unchanged.
* If the API is updated, we require the `create.yaml` and `update.yaml` changes. Also, the `config/crds` and `_generated_object_<resource>.golden.yaml` shall match the corresponding changes after running.

---

## Add the mapper

Follow [deep-dives Step 3](../deep-dives/3-add-mapper.md)

This PR adds the direct mapper. You can do this together with the previous step or the next step if no additional manual changes are needed.

* Ensure symmetric mapping for 100% of the CRD fields (`Spec_ToProto` and `Spec_FromProto`).

---

## Add the controller 

Follow [deep-dives Step 4](../deep-dives/4-add-controller.md).

* **Reference Resolution Timing**: Always resolve external references (`*refs.<Kind>Ref` or `common.NormalizeReferences`) on line 1 of `Create()` and `Update()` inside `pkg/controller/direct/<service>/<resource>_controller.go`. **NEVER resolve references in `AdapterForObject()`**. `Delete()` must never attempt to resolve external references; doing so in `AdapterForObject()` triggers an unrecoverable deletion deadlock if a referenced resource was deleted first (per PR #13573).
* Remove the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` from the API tag, and rerun `dev/tasks/generate-crds`.

Now, everything is on the direct reconciliation and the resource no longer depends on the Terraform or DCL based controllers. To verify the direct controller, `the _generated_object_<resource>.golden.yaml` should be unchanged.

### PR Reviews

* We suggest to have the roundtrip fuzz tests to cover all the fields in `spec` and `status.observedState` fields.
* We require removing the `cnrm.cloud.google.com/dcl2crd: "true"` or `cnrm.cloud.google.com/tf2crd: "true"` tags from the API and the presubmit test passes without `the _generated_object_<resource>.golden.yaml` change.
