These scripts are a work in progress.

The goal is to create tooling that will (one day) be able to generate the 3 PRs that make up building a resource.

For now, the goal is to capture the current process, and get commits in place that will need to be manually edited,
and then engineers will `git commit --fixup` or `git absorb`

# Step 1: mockgcp

In this step, we get mockgcp support for a resource.  We assume that gcloud supports it.

There should be a few steps:

1) A test script that invokes gcloud
2) The captured "golden" output from running the test
3) The implementation of the mock (which may include service.go if this is the first resource for a service)
4) The captured "golden" output from running the mock
5) Iterate the mock/normalization and rerun the test until the golden output matches


This corresponds to the following commits:

1. create test script for for ${GCLOUD_COMMAND}
1. capture golden output for ${RUN_TEST}
1. mockgcp support for ${SERVICE} ${RESOURCE}

Later changes should be squashed into this structure.


# Step 2: CRD / Mapper / Fuzzer

Note: this is _much_ easier if you are doing a new or alpha resource.
Matching the terraform schema can be painful, because terraform does not follow the kubernetes/declarative model.

Steps:

1) We write a generator script, and run it to generate the types other than Spec & Status
2) We generate the Spec and Status, and commit it (we'll have to merge it manually, currently)
3) Generate a fuzzer, and commit it
4) Iterate the fuzzer and type definitions until the fuzzer passes (i.e. until we round-trip)

This corresponds to the following commits:

1. generated scripts for ${CRD_KIND}
1. run script to generate types
1. Update spec and status and any other types (e.g. for refs)
1. dev/tasks/generate-crds to generate types and api machinery
1. generate fuzzer

