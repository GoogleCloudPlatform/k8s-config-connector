### [2026-07-02] Correct service name mapping for TranscoderJob
- **Context**: Implementing `TranscoderJob` types, CRD, and IdentityV2 under `apis/transcoder/v1alpha1`.
- **Problem**: The issue specified the service as `google.cloud.transcoder.v1`, which did not match the actual package definition `google.cloud.video.transcoder.v1` in the Google APIs proto files. This caused the generator to fail to find the `Job` proto message with `failed to find the proto message google.cloud.transcoder.v1.Job: proto: not found`.
- **Solution**: Changed the `--service` parameter in `generate.sh` to `google.cloud.video.transcoder.v1`.
- **Impact**: The type generator successfully generated all KCC Go structures and CRDs, ensuring subsequent mapper and controller generation steps will use the correct GCP proto import paths.
