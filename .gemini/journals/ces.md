### [2026-06-03] Handling Greenfield Protos Added After Pinned Commit
- **Context**: Implementing `CESApp` types (Google Cloud CES v1beta) under `ces.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The `google/cloud/ces` proto files did not exist in the pinned googleapis SHA `731d7f2ab6e4e2ea15030c95039e2cb66174d4fb`. Therefore, compiling `googleapis.pb` with the standard script left out the CES package, causing `generate-types` to fail.
- **Solution**: Modified `dev/tools/controllerbuilder/generate-proto.sh` to checkout `google/cloud/ces/v1beta` from `origin/master` after resetting to the pinned SHA, and added the path explicitly to the protoc compile command. This allowed us to successfully compile the `googleapis.pb` descriptor containing the ces service definitions.
- **Impact**: Any subsequent developer or automated pipeline running `./apis/ces/v1alpha1/generate.sh` can successfully regenerate types without missing proto errors.
