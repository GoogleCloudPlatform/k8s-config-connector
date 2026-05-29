# BigtableAppProfile Identity and Reference Journal

- The resource kind is `BigtableAppProfile`, but the existing files were named `appprofile_identity.go` and `appprofile_reference.go`. I renamed them to `bigtableappprofile_identity.go` and `bigtableappprofile_reference.go` to follow the "full lowercase Kind" convention.
- The `BigtableAppProfile` resource depends on a `BigtableInstance`. The identity resolution needs to resolve the `InstanceRef` to get the instance ID and project ID.
- The existing `InstanceRef` uses an older pattern, but it still has `NormalizedExternal` which is useful for resolving the reference.
- I used `bigtableadmin.googleapis.com` as the host in `gcpurls.Template`, which matched the existing implementation and the resource type in CAI, although the CAI `nameFormats` used `bigtable.googleapis.com`. An exception was already present in `pkg/gcpurls/registry_test.go`.
- I had to update the direct controller `pkg/controller/direct/bigtable/appprofile_controller.go` to use the new identity struct and fields.
- `runtime.DefaultUnstructuredConverter.FromUnstructured` was used to handle conversion from `unstructured.Unstructured` in `getIdentityFromBigtableAppProfileSpec`.
