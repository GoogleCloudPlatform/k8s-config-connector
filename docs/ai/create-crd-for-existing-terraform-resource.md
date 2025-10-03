# How to Create a CRD (and Mapper and Fuzzer) for an existing terraform resource

As part of moving resources from terraform controllers to direct controllers, we want to create a normal CRD using the controller-runtime framework.

These direct CRDs are defined with go types under apis/<service>/v1beta1/<kind>_types.go

When we initially create the direct CRD, we want to make sure that we keep the same schema as the old Terraform CRD, for compatability and so we can roll back.  Ideally the generated CRD (under `config/crds`) does not change.
We accept some changes, more details below.

As with all CRD creation, we want to create reviewable PR that has tests.  So we want to send a PR with a CRD, Mapper and Fuzz-test.

## Creating a CRD

We are able to generate a good starting point with some tooling we have created.  Begin by creating or adding to the generate.sh script located in ``apis/<service>/v1beta1/generate.sh`.
Please read `apis/bigtable/v1alpha1/generate.sh` for a good example.  Make sure you have added your resource to `generate-types`.

If you then run the `generate.sh` script, it should generate the shared types `types.generated.go` and some starting points for the type itself.  It should also generate a mapper.

The script also generates two files if they don't exist: `<type>_reference.go` and `<type>_identity.go`; we can't test those yet so if they were added then just comment them out for now.

If you look at the CRD under `config/crds` (you can use `git diff config/crds`) then you will likely see some changes.  Iterate on the types until they match.  Most of the types will already be defined automatically under `types.generated.go`, those are likely correct so you need to focus on `<Kind>Spec` and `<Kind>Status`.

There is often a good starting point for the Spec and Schema types generated under `types.generated.go`

Some changes are acceptable:

* Field descriptions can change, particularly for metadata.
* status.observedGeneration will now be an `int64`

Changes to the schema itself, such as a field being added or removed, are not acceptable.

## Creating a Fuzz-Test

Create a fuzz-test under `pkg/controller/direct/<service>/<kind>_fuzzer.go` by using the example `pkg/controller/direct/orgpolicy/policy_fuzzer.go`.

You will also need to make sure the package is imported in `pkg/controller/direct/register/register.go`; that ensures that the fuzz testers will pick it up.

You can then run the fuzz-test with `dev/ci/presubmits/fuzz-roundtrippers`.  We want the test to pass, but likely it will identify some fields that are not correctly mapped in the fuzz-test.
Update the fuzz test until it passes, by adding spec / status / unimplemented fields.  In general when a field is unimplemented, try to describe why it is not being implemented -
categorize it as an identity field / volatile field etc.