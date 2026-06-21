### 2026-06-03 Handling Missing GCP Proto via mockgcp/apis For Greenfield Scaffolding
- **Context**: Scaffolding direct KRM types and IdentityV2 for `AppOptimizeReport` under `appoptimize.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: The proto definition for `google.cloud.appoptimize.v1beta.Report` was missing from the pinned `googleapis` repository commit (`731d7f2ab6`), causing `generate-types` to fail.
- **Solution**: Created a custom, mock proto file at `mockgcp/apis/google/cloud/appoptimize/v1beta/report.proto`, registered it in `dev/tools/controllerbuilder/generate-proto.sh`, and cleared the cached protobuf descriptor (`.build/*.pb`) to force recompilation.
- **Impact**: Provides a repeatable strategy for scaffolding direct types for new resources whose Google APIs are not yet published or accessible in the public googleapis repo.
