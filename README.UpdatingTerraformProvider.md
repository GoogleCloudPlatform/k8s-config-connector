# Updating the Google Terraform Provider

1.  Find the latest release TF provider version tag from the
    [TF provider’s GitHub release page](https://github.com/hashicorp/terraform-provider-google-beta/releases).
    This will be of the format `v4.X.Y.` Read through the provider’s release
    notes and make note of any aspects that may cause us difficulty with
    upgrading to this version (especially breaking changes).
2.  Update the TF version cloned by the
    [`third_party/ Makefile`](third_party/Makefile) to this latest TF provider
    release.
3.  Run `make ensure` from the repo’s root. This will clone the version
    provided into `third_party/github.com/hashicorp`, and then add our custom
    patches. If there are any issues with the patches, see "Fixing a TF
    Provider Patch" below.
4.  Run `make vet` to check if the updated code compiles. If it does not, then
    there might either be issues in the patches, or a dependency with a new
    contract.
    1.  If the issues are with a patch, treat this similarly to a patch
        application issue and follow "Fixing a TF Provider Patch" below.
    1.  If the issue is with the Terraform provider breaking its contract in
        some way, there may be larger issues that we will need to work through.
        Investigate and, if non-trivial, bring this up for further discussion.
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
    1.  Fields may not be existent in the underlying GCP resource, and we need
        to turn them into directives: Add the fields into the `directives` list
        in the service mapping.
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
        the `-update` flag to regenerate the golden file. The two tests in
        particular that may require this are:
        1.  `go test -v ./pkg/crd/fielddesc -test.run TestOutputMatches -update`
        1.  `go test -v ./pkg/crd/template -test.run TestSpecAndStatusToYAML
            -update`
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
        instructions [here](README.ChangingTerraform.md) to create a TF provider
        patch to fix the bug locally. Surface the issue separately to the
        terraform team by creating an
        [issue](https://github.com/hashicorp/terraform-provider-google/issues)
        in terraform-provider-google repository if this bug has not been
        reported.
    1.  If this bug requires Terraform domain knowledge to fix, report the issue
        in the terraform-provider-google repo as instructed above, and punt the
        TF provider upgrade. You should bring this topic as a new issue for
        further discussion.

## Fixing a TF Provider Patch

Patch files will fail to apply if there is a git merge conflict. We will need to
manually resolve these conflicts before we will be able to successfully upgrade
the TF provider. The following steps describe a strategy to help resolve
conflicts and update our patch files:

1.  Comment the affected patch files’ `git apply` calls out of
    `third_party/Makefile` and re-run `make ensure` until it completes. This
    will set your working directory to have the unpatched TF code for those
    particular files.
2.  Run `git add third_party/` && `git commit -m "Update TF provider under third_party"`.
    This lets HEAD reflect the unpatched TF, and `git diff` will now show the
    full difference between the unpatched and our desired patched state.
3.  Uncomment the first affected patch application in `third_party/Makefile` and
    run `make ensure` again. Note the line number that fails.
4.  Open the patch file up and note what the change is. Each block starting with
    `@@` denotes a separate patch in the patch file, as well as commonly a
    function name and line number. Note that these could have changed between TF
    versions, but they can serve as a useful guide.
5.  Open the file being patched in `third_party/`, and manually update the file
    to include the desired patch.
6.  Remove that particular patch from the patch file, and re-run the particular
    `git apply` for that patch file. This will attempt the other patches. If
    there are other patch issues that fail, repeat steps 4-5 for the affected
    patch.
7.  Once `git apply` goes through smoothly, run `git diff
    third_party/github.com/[PATH_TO_FILE] > /tmp/patch`. This will save the
    required patches for the file to a temporary file `/tmp/patch`. Now open
    both `/tmp/patch` and `hack/terraform-overrides/[PATCH_FILENAME].patch` side
    by side and replace the patches of the target file in
    `hack/terraform-overrides/[PATCH_FILENAME].patch` with the patches from
    `/tmp/patch`. Note some `[PATCH_FILENAME].patch` contains patches to
    multiple target files, and you should only replace the patches for the file
    you are dealing with.
8.  Run `git add hack/terraform-overrides third_party/ && git commit --amend` in
    order to update your work-in-progress commit to include the changes for that
    patch file.
9.  Repeat steps 3-7 for each affected patch file.
10. Once all patches go through smoothly, re-run `make ensure`. This will do one
    last verification pass, and then ensure our repo is up-to-date with all the
    changes.

## Setting up generate_field_changes tool

On your dev machine, run the following commands:

1.  `sudo apt-get install pip`
2.  `pip3 install gitpython`
3.  `pip3 install deepdiff`
