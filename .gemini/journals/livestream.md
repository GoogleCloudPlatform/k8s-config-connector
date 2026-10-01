### [2026-07-02] LiveStreamAsset types generation and proto package correction
- **Context**: Implementing KRM types, CRD, and IdentityV2 for `LiveStreamAsset` under `livestream.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: 
  1. The issue instructions suggested the service name as `google.cloud.livestream.v1`, but the actual Google API proto package for livestream in googleapis is `google.cloud.video.livestream.v1`.
  2. Running the scaffolding generator initially with empty spec fields caused the nested structures `Asset_VideoAsset` and `Asset_ImageAsset` to be pruned as unreachable (commented out) in `types.generated.go`.
- **Solution**:
  1. Corrected the service name in `generate.sh` to `google.cloud.video.livestream.v1`.
  2. Mapped `Labels`, `Video`, `Image`, and `Crc32c` fields into `LiveStreamAssetSpec` in `livestreamasset_types.go` first, and then ran `./apis/livestream/v1alpha1/generate.sh` again. This allowed the generator to see those nested types as reachable, generating them as active, clean Go structs in `types.generated.go`.
- **Learnings**: Always check `proto-list-final.yaml` or googleapis repository to verify the actual proto package name. To ensure nested structs are not commented out in `types.generated.go`, reference them in `_types.go` Spec/Status first, then run/re-run the generator.

### [2026-10-01] LiveStreamInput controller implementation and GCP behavior
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `LiveStreamInput`.
- **Findings**:
  1. `LiveStreamInput` labels cannot be updated after creation in GCP LiveStream API (`googleapi: Error 400: The request was invalid: labels can't be updated`). Spec labels should remain constant between creation and updates, while `preprocessingConfig` and `securityRules` are mutable.
  2. `LiveStreamInput` generates a dynamic `uri` stream key and IP address upon creation. Added normalization rules in `mockgcp/mocklivestream/normalize.go` so that the stream key and URI in responses are normalized for golden file comparisons.

