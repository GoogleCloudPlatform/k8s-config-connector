# NotebookInstanceV2 Types Journal

## Shortcomings & Solutions

### Type Name Collisions across API Versions / Services

* **Context:** In Config Connector, multiple resources/services under the same Group may share Go package directories (e.g. `apis/notebooks/v1alpha1/`). When generating direct types for a new service version (e.g. notebooks `v2`), helper structs like `VMImage` or `BootDisk` can clash with those of existing services (e.g. notebooks `v1` used by `NotebooksEnvironment`).
* **Observation:** The code generator (`generate-types`) automatically prunes or overwrites `types.generated.go`. If a struct name like `VMImage` is shared, generating v2 will overwrite the v1 definition, altering its schema and causing regressions on existing CRDs/controllers (such as `NotebooksEnvironment`).
* **Solution:** Handwrite/customize these helper types inside the primary `<resource>_types.go` file (`notebookinstancev2_types.go`) and prefix them with the kind's name (e.g. `InstanceVMImage`, `InstanceBootDisk`, `InstanceGCESetup`). The code generator will recognize them as existing types and skip generating them in `types.generated.go`, thus completely isolating the schemas and preventing regressions on existing resources.
