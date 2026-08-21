# Note
This document provides an overview of how this project previously updated the Google Terraform Provider (TPG). It serves as a resource for individuals interested in understanding past update procedures and gaining insights into potential future strategies.

We are actively evaluating the optimal long-term strategy for managing Terraform as a dependency. In the interim, for those requiring modifications to the Terraform provider code, please refer to the [README.ChangingTerraform.md](README.ChangingTerraform.md) document for detailed instructions.


# Updating the Google Terraform Provider

1.  Find the latest release TF provider version tag from the
    [TF provider’s GitHub release page](https://github.com/hashicorp/terraform-provider-google-beta/releases).
    This will be of the format `v5.X.Y.` Read through the provider’s release
    notes and make note of any aspects that may cause us difficulty with
    upgrading to this version (especially breaking changes).
2.  Run `git subtree pull --prefix third_party/github.com/hashicorp/terraform-provider-google-beta https://github.com/hashicorp/terraform-provider-google-beta.git v5.X.Y --squash` from the repo’s root. This will update the provider code to the specified version.
3.  Git may prompt you on merge CONFLICT, carefully review and resolve the CONFLICT to complete the git merge.
4.  Run `make vet` to check if the updated code compiles. 
5.  Run `make generate manifests resource-docs generate-go-client` to update all
    the generated files.
6.  Run `python3
    scripts/generate-field-diffs-tf-upgrade/generate_field_changes.py` to
    generate diff summary of our CRDs. Check "Setting up generate_field_changes
    tool" section below to set up if this is your first time using it. (Double
    check with `git diff config/crds/resources` and look through any changes).
    There may be special situations (as described below) that fields added by
    the TF provider need to be configured manually. If so, modify the resources’
    service mappings accordingly and rerun the command from step 3.
    1.  Fields may be implicit resource references that we need to KRM-ify:
        1.  Modify the resources’ service mappings to configure the references.
        1.  (Optional) Consider adding the reference fields in testdata and
            samples if it is an important ask from users or belongs to a
            core user case.
    1.  Fields may not exist in the underlying GCP resource, and we need
        to turn them into directives: Add the fields into the `directives` list
        in the service mapping. `directives` fields do not exist in resource CRD.
    1.  Fields may not exist in the underlying GCP resource, but are returned
        by Terraform on read, we need to add them into the `mutableButUnreadableFields`
        list in the service mapping. The fields are unreadable but the spec value of
        the fields can be modified.
    1.  Fields can't be supported or will result in the suboptimal UX (e.g.
        multi-kind reference in the legacy style) due to lack of feature
        support, and we need to ignore the fields right now: Add the fields into
        the `ignoredFields` list in the service mapping, file a GitHub issue 
        to support this field, and add a TODO with the issue ID right above the
        entry in the ignoredFields list.
7.  Rerun `python3
    scripts/generate-field-diffs-tf-upgrade/generate_field_changes.py` and save 
    the output change summary so it can be used in step 11.
8.  Run `make test` to run the unit tests.
    1.  IAM policy support may have been added to existing resources. If so, add
        IAM support to the resource’s service mapping. If there is an issue with
        IAM support and it can’t be added, make a tracking GitHub issue and
        block it temporarily in the [service mapping test code](config/tests/servicemapping/servicemapping_test.go).
    1.  If any new directive fields are added, you might run into issues with
        ResourceSkeleton’s `TestNewFromAsset` and `TestNewFromURI` tests. Any
        lingering fields in the expected vs. actual resource skeleton will
        indicate which fields are suspect. If it makes sense for the field to
        become a directive, add it to the resource’s corresponding
        servicemapping.
    1.  Golden files may need to be updated. If so, run the relevant test with
        the `WRITE_GOLDEN_OUTPUT` env-var to regenerate the golden file. The two tests in
        particular that may require this are:
        1.  `WRITE_GOLDEN_OUTPUT=1 go test -v ./pkg/crd/fielddesc -test.run TestOutputMatches`
        1.  `WRITE_GOLDEN_OUTPUT=1 go test -v ./pkg/crd/template -test.run TestSpecAndStatusToYAML`
9.  Run `make ready-pr` to ensure the PR is ready to be sent out.
10. Run `git commit -m "Update TF provider to [version]"` (replacing
    `[version]` with the desired TF version.
11. Run `git commit --amend` and paste the change summary in the commit message.
12. Push the commits to the topic branch of your forked repo.
    Send a PR for review.
13. There are chances that presubmit tests in GitHub Action might fail due to the
    TF provider bugs in the latest version. Perform a judgment call on the next
    step following the below recommendations.
    1.  If you understand the bug well and want to unblock the TF provider
        upgrade quickly to delivery some user requested features, follow the
        [instructions](README.ChangingTerraform.md) to create a TF provider
        patch to fix the bug locally. Surface the issue separately to the
        terraform team by creating an
        [issue](https://github.com/hashicorp/terraform-provider-google/issues)
        in terraform-provider-google repository if this bug has not been
        reported.
    1.  If this bug requires Terraform domain knowledge to fix, report the issue
        in the terraform-provider-google repo as instructed above, and punt the
        TF provider upgrade. You should bring this topic as a new issue for
        further discussion.

## Setting up generate_field_changes tool

On your dev machine, run the following commands:

1.  `sudo apt-get install pip`
2.  `pip3 install gitpython`
3.  `pip3 install deepdiff`
