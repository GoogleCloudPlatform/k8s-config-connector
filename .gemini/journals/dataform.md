### [2026-06-03] Implement direct types for DataformFolder
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DataformFolder` (v1alpha1).
- **Problem**: The proto package `google.cloud.dataform.v1beta1` has no `Folder` message, causing `generate-types` to fail.
- **Solution**: We created a custom protobuf definition for `Folder` in a new `overlay.proto` file and appended it to the local `dataform.proto` under `.build/third_party/googleapis` before compiling the protocol buffer descriptor. This allowed KCC's code generator to parse the custom proto and auto-scaffold all the KRM types, deepcopy logic, and mappers.
- **Impact**: Enables smooth scaffolding for any GCP resources where the corresponding proto files do not exist or are outdated in the pinned `googleapis` submodule, by defining them locally and updating the compiled descriptor.

### [2026-07-06] Implement direct types for DataformTeamFolder
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DataformTeamFolder` (v1alpha1).
- **Problem**: Finding the correct proto file definition for `TeamFolder` since the Google APIs SHA pinned in older resource implementations (`cdc919ff596e263f2cc55a9780d2f74633da1ced`) didn't contain `TeamFolder` or `teamFolders` messages in `v1beta1`, but it is fully defined in the `google.cloud.dataform.v1` package of more recent Google APIs master branches.
- **Solution**: We targeted the `google.cloud.dataform.v1` proto service directly inside `apis/dataform/v1alpha1/generate.sh` and scaffolded the KRM types using the direct stable v1 proto fields, bypassing the old beta1 package limitation for `TeamFolder`.
- **Impact**: Enables clean scaffolding and future direct-controller implementation for DataformTeamFolder under the direct approach.
