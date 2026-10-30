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

### [2026-10-29] DataformReleaseConfig oneof mapper and external-only compilation result
- **Context**: Triaging judgement queue and finalizing direct KRM types and identity for `DataformReleaseConfig`.
- **Problem**: `NotebookRuntimeOptions.gcs_output_bucket` is part of a proto `oneof execution_sink`. Converting it to `GCSOutputBucketRef *storagev1beta1.StorageBucketRef` caused `generate-mapper` to emit direct field assignment `out.GcsOutputBucket = ...`, which fails to compile because protobuf oneof variants require setting the interface wrapper struct `out.ExecutionSink = &pb.NotebookRuntimeOptions_GcsOutputBucket{...}`.
- **Solution**: Implemented handwritten mappers for `NotebookRuntimeOptions` in `pkg/controller/direct/dataform/dataformreleaseconfig_mapper.go`. `generate-mapper` detects existing mapping functions in the package and automatically skips generating duplicates in `mapper.generated.go`. Also defined an external-only reference and IdentityV2 for `DataformCompilationResultRef` (`dataform.googleapis.com/CompilationResult`) since KCC does not have a CRD for compilation results.
- **Impact**: Resolves mapper generation compilation issues for oneof reference fields and provides external-only reference support for DataformCompilationResult.

### [2026-10-30] DataformReleaseConfig scheduled release and immutable field constraints
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DataformReleaseConfig` (Issue #14701).
- **Problem**: Setting `cronSchedule` in E2E fixtures caused real GCP to fail with `googleapi: Error 400: Automatic release is not supported in first-party repositories that enabled strictActAsChecks`. Furthermore, attempting to modify `codeCompilationConfig` in `update.yaml` returned `googleapi: Error 400: com.google.apps.framework.request.BadRequestException: Request update_mask contains immutable fields: [code_compilation_config]`.
- **Solution**: Repositories without third-party git remote credentials (first-party / Google-hosted repositories) do not support automatic scheduled compilation via `cronSchedule`/`timeZone`. Omit `cronSchedule` and `timeZone` in test fixtures targeting basic repositories. Treat `codeCompilationConfig` as immutable during updates, only mutating fields like `gitCommitish` and `disabled`.
- **Impact**: Prevents invalid GCP API requests during E2E tests and clarifies which fields can be updated on `DataformReleaseConfig`.

