# Overview

This README provides guidance on modifying the Config Connector's [local copy](third_party/github.com/hashicorp/terraform-provider-google-beta) of the [Terraform Provider Google (TPG) Beta](https://github.com/hashicorp/terraform-provider-google-beta) library, maintained as a [Git Subtree](https://www.atlassian.com/git/tutorials/git-subtree).

Approximately 75% of Config Connector resources leverage Terraform. To verify if a Config Connector resource leverages Terraform, search for the presence of `cnrm.cloud.google.com/tf2crd: "true"` within its CRD schema. Config Connector has historically implemented numerous custom modifications (patches) on top of the standard TPG Beta library. These customizations are accessible through the Git Subtree's [commit history](https://github.com/GoogleCloudPlatform/k8s-config-connector/commits/master/third_party/github.com/hashicorp/terraform-provider-google-beta). Alternatively, if you're solely interested in tracking modifications for a specific file, you can use the following command and focus on commits made after  `Remove legacy third_party dir`:

```
$ git log --follow -- third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/spanner/resource_spanner_database.go

commit a6783bf0abc6b4dcd948532bffb938fb82f60620
Author: Shuxian Cai <shuxiancai@google.com>
Date:   Wed Nov 29 00:08:40 2023 +0000

    deletion_protection_default.patch
    
    The provider is introducing breaking changes to database resources to
    prevent accidental stateful resource deletion in Terraform. KCC
    customers rely on deletion behavior being consistent with its current
    behavior, and the accidental stateful resource deletion scenario
    isn't as much of an issue in KCC, so we patch this new `true` default
    out.

commit d68d4edf28f095627696a0eb956d60adde167a0d
Author: Shuxian Cai <shuxiancai@google.com>
Date:   Wed Nov 29 00:04:21 2023 +0000

    Remove legacy third_party dir
...

```

The current Git Subtree incorporates [v4.84.0](https://github.com/hashicorp/terraform-provider-google-beta/tree/v4.84.0) of TPG Beta with 60+ Config Connector-specific customizations. We are currently evaluating the optimal future strategy for managing Terraform as a dependency, including potential migration to the v5.* major version of TPG Beta and addressing associated breaking changes.

Until a long-term strategy is established, modifying this local Git Subtree may be necessary under specific scenarios, including:

* Bug Fixes: Addressing identified bugs within the TPG Beta library. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/commit/d60a04e87165d3610d37601d0c1964e28af49d5f)
* Missing Fields: Implementing fields recently introduced in TPG v5 or the underlying GCP API but absent in the current subtree. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/commit/1978b77d6cc9a579c2d5ff6bad1d75ba619d23de)
* Missing Resources: Adding resources recently introduced in TPG v5 but not present in the current subtree. [Example](https://github.com/GoogleCloudPlatform/k8s-config-connector/commit/f54d7dff2047496fa874a63bbc798d0a2b97f7e6)


# Make Your Change in the Git Subtree

1.  Implement Your Change:
    * Make your desired modifications directly within the [third_party/github.com/hashicorp/terraform-provider-google-beta](third_party/github.com/hashicorp/terraform-provider-google-beta) directory.
    * Run `make vet` to check if the updated code compiles.

1.  Commit Your Change:
    Craft a clear commit message. Include a concise title and detailed description outlining the changes and their purpose. Refer to the provided examples for guidance.

1.  Generate Files (Optional):
    If adding new fields to an existing resource, run `make ready-pr` to update all the generated files. Then add the generated files as a separate commit.
 
1.  Test Your change:
    * [Run tests](README.NewResourceFromTerraform.md#run-tests) associated with the modified resource/CRD.
    * [Run sample tests](README.NewResourceFromTerraform.md#rundisable-sample-tests) associated with the modified resource/CRD.

1.  Submit a Pull Request:
    * While your Pull Request can contain other commits, make sure you have grouped all changes within the aforementioned [directory](third_party/github.com/hashicorp/terraform-provider-google-beta) into a single, dedicated commit.

1.  New Resource from Terraform (Optional):
    For new Terraform resources, follow the steps outlined in [README.NewResourceFromTerraform.md](README.NewResourceFromTerraform.md) to create another PR and introduce a corresponding Config Connector resource.
