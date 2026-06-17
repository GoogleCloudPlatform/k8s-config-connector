### [2026-06-03] Implement direct types for DataformFolder
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DataformFolder` (v1alpha1).
- **Problem**: The proto package `google.cloud.dataform.v1beta1` has no `Folder` message, causing `generate-types` to fail.
- **Solution**: We created a custom protobuf definition for `Folder` in a new `overlay.proto` file and appended it to the local `dataform.proto` under `.build/third_party/googleapis` before compiling the protocol buffer descriptor. This allowed KCC's code generator to parse the custom proto and auto-scaffold all the KRM types, deepcopy logic, and mappers.
- **Impact**: Enables smooth scaffolding for any GCP resources where the corresponding proto files do not exist or are outdated in the pinned `googleapis` submodule, by defining them locally and updating the compiled descriptor.
