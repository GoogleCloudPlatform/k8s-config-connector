# How to Create a CRD for an existing terraform resource

As part of moving resources from terraform controllers to direct controllers, we want to create a normal CRD using the controller-runtime framework.

These direct CRDs are defined with go types under `apis/<service>/<version>/<kind>_types.go`.  The service should not include the `cnrm.cloud.google.com` suffix, just the first component (e.g. `artifactregistry` not `artifactregistry.cnrm.cloud.google.com`).

When we initially create the direct CRD, we want to make sure that we keep the same schema as the old Terraform CRD, for compatibility and so we can roll back.  Ideally the generated CRD (under `config/crds`) does not change.
We accept some changes, more details below.  You can easily see changes by running `dev/tasks/diff-crds` after running `generate.sh`.

For a CRD we are migrating from Terraform or DCL, we do not (initially) need a Mapper or Fuzzer; we accept or reject the PR based on the CRD definition.  So send the generated mapper functions in your PR, but don't worry about a fuzzer yet.

## Creating a CRD

We are able to generate a good starting point with some tooling we have created.  Begin by creating or adding to the generate.sh script located in `apis/<service>/<version>/generate.sh`.
Please read `apis/bigtable/v1alpha1/generate.sh` for a good example to follow.  Make sure you have added your resource to `generate-types`.

If you then run the `generate.sh` script, it should generate the shared types `types.generated.go` and a scaffolding for the types that cannot be generated, file for the type, something like `apis/<service>/<version>/<kind>_types.go`.

If you compare the CRD (using `dev/tasks/diff-crds`) you will likely see some changes.  Iterate on the types until they match.  Most of the types will already be defined automatically under `types.generated.go`, those are likely correct so you need to focus on `<Kind>Spec` and `<Kind>Status` in `<kind>_types.go`.

There is often a good starting point for the Spec and Schema types generated under `types.generated.go`

Some changes are acceptable:

* Field descriptions can change, particularly for metadata.
* `status.observedGeneration` will now be an `int64`
* `int32` vs `int64` differences are often acceptable.

Changes to the schema itself, such as a field being added or removed, are not acceptable.
* Watch out for json tag casing (e.g. `objectId` vs `objectID`).
* Do not add new fields (like `projectRef` or extra status fields) that were not present in the Terraform CRD.

We are not yet switching the controller to direct, so we want the same labels; the controller-runtime annotations should look something like this:

```go
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
```

## Creating a Fuzz-Test

For a resource migration, we do not require a fuzz-test initially.

When adding a fuzz-test (later):
Create a fuzz-test under `pkg/controller/direct/<service>/<kind>_fuzzer.go` by using the example `pkg/controller/direct/orgpolicy/policy_fuzzer.go`.

You will also need to make sure the package is imported in `pkg/controller/direct/register/register.go`; that ensures that the fuzz testers will pick it up.

You can then run the fuzz-test with `dev/ci/presubmits/fuzz-roundtrippers`.  We want the test to pass, but likely it will identify some fields that are not correctly mapped in the fuzz-test.
Update the fuzz test until it passes, by adding spec / status / unimplemented fields.  In general when a field is unimplemented, try to describe why it is not being implemented -
categorize it as an identity field / volatile field etc.
