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

### [2026-10-02] LiveStreamChannel direct controller and GCP behavior
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `LiveStreamChannel`.
- **Findings**:
  1. GCP server defaults for `Channel`:
     - `inputConfig.inputSwitchMode` defaults to `FAILOVER_PREFER_PRIMARY` (enum 1).
     - `logConfig.logSeverity` defaults to `OFF` (enum 1).
     - `timecodeConfig.source` defaults to `MEDIA_TIMESTAMP` (enum 1).
     - `manifests`: `maxSegmentCount` defaults to 5, `segmentKeepDuration` defaults to `60s`.
     - `elementaryStreams.videoStream.h264`: `entropyCoder` defaults to `"cabac"`, `gopDuration` defaults to `2s`, `profile` defaults to `"main"`, `vbvSizeBits` defaults to `bitrateBps`, and `vbvFullnessBits` defaults to `0.9 * bitrateBps`.
     - `elementaryStreams.audioStream`: `sampleRateHertz` defaults to `48000`.
  2. Input and Asset references in GCP API responses return canonical resource paths using the numeric GCP Project Number rather than Project ID. Used `projectMapper.ReplaceProjectNumberWithIDInLink` to normalize both `inputAttachments[].input` and `staticOverlays[].asset` in `compareChannel`.
  3. Stream constraints: For `fmp4` containers, each `muxStream` must contain exactly one video or audio elementary stream. Text streams with `cea608` or `cea708` passthrough are embedded in video and cannot be included into `muxStreams`.
  4. In YAML manifests for KRM, mapping keys like `y:` in `position` must be quoted (`"y":`) to prevent YAML 1.1 boolean parsing as `true`.

### [2026-10-03] LiveStreamChannel MockGCP and Alignment verification
- **Context**: Aligning MockGCP logs with RealGCP output for `LiveStreamChannel` (`livestream.cnrm.cloud.google.com/v1alpha1`).
- **Findings & Verification**:
  1. Verified MockGCP implementation for `LiveStreamChannel` in `mockgcp/mocklivestream/channels.go` handling `CreateChannel`, `GetChannel`, `UpdateChannel`, `DeleteChannel`, and `ListChannels`.
  2. Verified that server defaults (`inputConfig`, `logConfig`, `timecodeConfig`, `manifests`, `elementaryStreams` H264 & Audio settings) and project number canonicalization in `inputAttachments` and `staticOverlays` match Real GCP behavior.
  3. Verified both `livestreamchannel-minimal` and `livestreamchannel-maximal` test fixtures pass `hack/compare-mock "fixtures/livestreamchannel"` and `TestGoldenLogAlignment` with 0 diffs.

