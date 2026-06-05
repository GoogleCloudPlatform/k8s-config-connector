### [2026-06-03] Implement direct types for DataformFolder
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `DataformFolder` (v1alpha1).
- **Problem**: The proto package `google.cloud.dataform.v1beta1` has no `Folder` message, causing `generate-types` to fail.
- **Solution**: We created a custom protobuf definition for `Folder` in a new `overlay.proto` file and appended it to the local `dataform.proto` under `.build/third_party/googleapis` before compiling the protocol buffer descriptor. This allowed KCC's code generator to parse the custom proto and auto-scaffold all the KRM types, deepcopy logic, and mappers.
- **Impact**: Enables smooth scaffolding for any GCP resources where the corresponding proto files do not exist or are outdated in the pinned `googleapis` submodule, by defining them locally and updating the compiled descriptor.

### [2026-06-05] DataformTeamFolder Implementation and Google APIs SHA dependency
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for the `DataformTeamFolder` direct resource (v1alpha1).
- **Problem**: The pinned Google APIs SHA in `apis/git.versions` (`1765b559c42386788ff0c6412491277b4791107a`) did not contain the proto definitions for the newly added `TeamFolder` resource in `google.cloud.dataform.v1.dataform.proto`, causing `generate-types` to fail.
- **Solution**: Updated `apis/git.versions` to point to a newer `googleapis` master commit `29214fa04adfdf39cab5cfc0f08b27b23fcfae1a` (Tue Mar 24, 2026) which includes the `folders and teamFolders related changes` in the v1 dataform proto. Then, successfully ran `generate.sh` to generate the types and deepcopy code.
- **Impact**: Future agents working on other Dataform features or folders can build upon this newer APIs SHA without issues, and any future step implementations (controller, mappers, fuzzer) will find the KRM types fully matching the proto structures.
