# LiveStreamChannel Controller Logic Journal

## Observations & Patterns

### 1. Server Defaults & Normalization
- `InputConfig.inputSwitchMode`: Defaults to `FAILOVER_PREFER_PRIMARY` (enum value 1).
- `LogConfig.logSeverity`: Defaults to `OFF` (enum value 1).
- `TimecodeConfig.source`: Defaults to `MEDIA_TIMESTAMP` (enum value 1).
- `Manifests`: `maxSegmentCount` defaults to 5, `segmentKeepDuration` defaults to `60s`.
- `ElementaryStreams[].videoStream.h264`: `entropyCoder` defaults to `"cabac"`, `gopDuration` defaults to `2s`, `profile` defaults to `"main"`, `vbvSizeBits` defaults to `bitrateBps`, and `vbvFullnessBits` defaults to `0.9 * bitrateBps`.
- `ElementaryStreams[].audioStream`: `sampleRateHertz` defaults to `48000`.

### 2. Project Number vs Project ID
- GCP LiveStream API canonicalizes input references (`channel.input_attachments[].input`) and asset references (`channel.static_overlays[].asset`) from project ID to numeric GCP project number in GET responses.
- Used `projectMapper.ReplaceProjectNumberWithIDInLink` to normalize both desired and actual links in `compareChannel` to prevent false diffs during re-reconciliation.

### 3. Stream Constraints
- For `fmp4` container, each `muxStream` must contain exactly one video or audio elementary stream.
- Text streams with `cea608` or `cea708` passthrough cannot be included in `muxStreams` (they are carried inside the video stream).

### 4. YAML 1.1 Quoting
- Mapping key `y:` in `position` must be quoted (`"y":`) to prevent YAML parsers from evaluating it as boolean `true`.
