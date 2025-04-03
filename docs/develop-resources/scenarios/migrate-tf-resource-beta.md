# Migrate Terraform/DCL-based Beta to direct Beta

Config Connector requires the Beta resource to be backward compatible. We
require the Mock GCP to reflect the real GCP  cases where the current Terraform
or DCL based approach uses or depend on to make sure the migration is backward
compatible. Example things to look out for:

*   assigning defaulting value to a field
*   having the same HTTP errors for matching

In the API and mapper step, we require the migration to happen in two stages:

*   The first stage is to make sure all APIs and their behavior are exactly the
    same after migration (verifying via  the golden logs).
*   The second stage is to add new fields to make the resource up-to-date with
    what the current GCP service supports.

In the controller step, we do not require the direct controller and the
Terraform/DCL based controller  to behave exactly the same, but fixing the
legacy or non Kubernetes-native  problems like merging config back to `spec`
(`state-into-spec:merge`), setting defaulting on-behalf of users.

To help the developer better manage the migration, we design the following steps
to be self-validating and not restricted by the Config Connector release cycles.
In other words, it is safe for a migration to happen across multiple Config
Connector releases.

## Add MockGCP tests

Follow [deep-dives Step 1](../deep-dives/1-add-mockgcp-tests.md). We use the
Terraform or DCL based controller to help develop the Mock GCP.

*   The 1st git-commit shall contain `create.yaml` and `update.yaml` files with
    exactly the same config.
*   The 2nd git-commit shall have `update.yaml` changing all the mutable fields,
    reflected in the git-diff based on the first commit.
*   The 3rd git-commit (or N-th git-commit) shall run against real GCP. The
    generated `_generated_object_<resource>.golden.yaml` and `_http.log` is the
    golden log reflecting the real GCP. The
    `_generated_obeject_<resource>.golden.yaml` matches the `spec` in
    `update.yaml`
*   The 4rd git-commit (or N+1-th git-commit) shall run against the Mock GCP.
    The git-diff shows `_generated_object_<resource>.golden.yaml` is unchanged,
    and the `_http.log` is as much like the real GCP as possible.

*   If the resource has dependencies (`dependencies.yaml`), we should cover all
    the referenced fields. If the dependency resource does not have mockGCP, you
    need to implement those dependencies' MockGCP methods as well.

    If the test suite has been added, make sure all field behaviors (OPTIONAL,
    REQUIRED, (IM)MUTABLE) are covered in the basic and full test suites
    `create.yaml` and `update.yaml`.

### PR Reviews

*   We require basic test suites and full test suites.
*   We require two contiguous git-commits reflecting the real GCP and Mock GCP
    logs. The Mock log should be as much like the real GCP log as possible.
*   We require two contiguous git-commits reflecting the mutable fields for
    `update.yaml` and matching changes in the generated
    `_generated_object_<resource>.golden.yaml` and `_http.log`
*   We require the `dependencies.yaml` to cover all referenced fields, and the
    `_http.log` showing the Cloud requests. You need to implement those
    dependencies' MockGCP methods as well.

## Add the backward compatible API

Follow [deep-dives Step 2](../deep-dives/2-define-apis.md). We use the Mock GCP
and Terraform or DCL based controller to verify the API migration is backward
compatible.

The PR shall contain the types and deepcopy codes. It shall make modifications
to make sure the CRDs are the same, because Beta resource has to be backward
compatible.

*   You may need to modify the auto-generated `types.go` to keep the existing
    fields the same (even if it is not following the recommended styles and
    conventions). You can run `dev/tasks/generate-crds` (repeatedly) to make
    sure the CRD are the same. Comment changes to match the latest protobuf are
    recommended as long as the content doesn't contain breaking changes. An
    example of a breaking change would be the comment explicitly says a
    previously optional field `required`.

*   Add `cnrm.cloud.google.com/dcl2crd: "true"` or
    `cnrm.cloud.google.com/tf2crd: "true"` to the API tag
    [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/v1.128.0/apis/alloydb/v1beta1/instance_types.go#L164),
    to continue using DCL-based or Terraform-based controllers.

*   You may see some new fields added to the CRD. These are expected since the
    Terraform/DCL based resources could be out of date (and users are looking
    forward to these new fields!). You **shall comment out** those new fields
    using `/*NOTYET ..*/` in this PR if they are not supported in the
    Terraform-based controller yet (we will add them later).

Note: *Do not use* `excluded_resources`, we want the presubmit to validate the
existing APIs are unchanged and the behaviors are the same.

### PR Reviews

*   We require the `_http.log`, `_generated_object_<resource>.golden.yaml`,
    `create.yaml` and `update.yaml` unchanged.

## Add the backward compatible mapper

Follow [deep-dives Step 3](../deep-dives/3-add-mapper.md)

The PR adds the direct mapper. You can do this together with the previous step
or the next step if no additional manual changes are needed. Using `/*NOTYET ..
*/` to comment out new functions, same as the last step.

## Add the direct controller

Follow [deep-dives Step 4](../deep-dives/4-add-controller.md).

*   Use the `KCC_USE_DIRECT_RECONCILERS` flag
    [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/dev/tasks/run-e2e#L27).

*Tips* The `KCC_USE_DIRECT_RECONCILERS` will override the `tf2crd` and `dcl2crd`
label to force using the direct controller, but it will not affect the releases
which will still use the Terraform/DCL based controllers until the Direct
controller is fully ready. This allows developing the API and controller in
different Config Connector releases without code freeze.

### PR Reviews

*   We require the roundtrip fuzz tests to cover all the fields in `spec` and
    `status.observedState` fields
    [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/f313b00c52f09c4a52a2eb5fe2c15fa4b30a05fd/pkg/controller/direct/discoveryengine/fuzzers.go#L26-L47)
*   We require the MockGCP pass without any change to the `create.yaml`,
    `update.yaml` and `_generated_object_<resource>.golden.yaml` files.

## Switch to the direct controller (optional)

Once the direct controller is fully implemented, you can swith to the direct
controller.

*   Remove the `cnrm.cloud.google.com/dcl2crd: "true"` or
    `cnrm.cloud.google.com/tf2crd: "true"` API tag
    [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92).
    Run `dev/tasks/generate-crds` to update the CRDs, now all the API and
    controller have been migration to the direct approach.

*   Remove the `KCC_USE_DIRECT_RECONCILERS` flag.

### PR Reviews

*   We require the MockGCP unchange and pass.

## Add new API and mapper

The PR shall add back those fields that commented out in Step 3. You need to
update both `types.go` and `mapping.go`, and may need to adjust the direct
controller as well to support those new fields.

### PR Reviews

*   The MockGCP test shall reflect the new fields in `create.yaml`,
    `update.yaml` (if mutable), `_generated_object_<resource>.golden.yaml`, and
    `_http.log`.
*   We require the roundtrip fuzz tests to cover the new fields in `spec` or/and
    `status.observedState` fields
    [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/0bbac86ace6ab2f4051b574f026d5fe47fa05b75/pkg/controller/direct/redis/cluster/roundtrip_test.go#L92)
    (For mapper)
*   We require no /*NOTYET .. */ in the API and Mapper code.
