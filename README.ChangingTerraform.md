# Overview

This README goes over how to make changes to our
[Terraform library](third_party/github.com/hashicorp/terraform-provider-google-beta).

There are two approaches:

1.  [Make your change upstream](#make-your-change-upstream)
1.  [Make your change locally](#make-your-change-locally) first, and then
    upstream the change.

Generally speaking:

*   If your change is cheap or non-urgent (e.g. changes to field descriptions,
    minor logic changes), prefer making the change upstream.
*   Otherwise, make the change locally first.

Either way, it is recommended that you consult someone from the
[Terraform Provider Google](https://github.com/hashicorp/terraform-provider-google)
project about the change you intend to make first to get
an early signal about the likelihood of your change being accepted upstream.

# Make Your Change Upstream

The upstream
[Terraform repository](https://github.com/hashicorp/terraform-provider-google-beta)
is generated using
[Magic Modules](https://github.com/googleCloudPlatform/magic-modules). This
means that to actually "upstream" your change, you need to make changes to Magic
Modules.

Read through the Magic Modules README.md to learn how to get started.

You can stop reading now if your change is expected to be accepted by
Terraform. Once the change is accepted, it will be adopted to Config Connector
in the next [Terraform update](README.UpdatingTerraformProvider.md).


# Make Your Change Locally

1.  Make your desired change locally in
    [third_party/github.com/hashicorp/terraform-provider-google-beta](third_party/github.com/hashicorp/terraform-provider-google-beta).

1.  Create a patch file in [terraform-overrides](hack/terraform-overrides):

    ```bash
    git diff master -- third_party/github.com > hack/terraform-overrides/my-changes.patch
    ```

    This patch file will ensure that your changes are preserved whenever we
    update our Terraform library.

    Note that if you make any more modifications to your change, then the patch
    file needs to be updated.

1.  Create an issue to track the upstreaming of your change.

1.  Edit the [third_party/Makefile](third_party/Makefile). Add your patch to the
    list of patches applied in the `apply-patches:` section. Add a link to your
    tracking issue and comment on the purpose of the patch.

1.  Update the copy of the Terraform library.

    ```bash
    make ensure
    ```

1.  (Optional) Commit all your changes and submit for review. To keep the change
    as discrete, reviewable, and reappliable unit, the commit and review should
    *ONLY* contain your Terraform changes and no other changes to the project.

1.  (Optional) [Make your change upstream](#make-your-change-upstream)

1.  (Optional) Once you have a Pull Request, update the tracking issue with a
    link to the PR.

1.  (Optional) Once the change has been submitted upstream, pull in a new
    version of the Terraform library through [Terraform update](README.UpdatingTerraformProvider.md)
    and remove your patch from [terraform-overrides](hack/terraform-overrides)
    and the [third_party/Makefile](third_party/Makefile).
