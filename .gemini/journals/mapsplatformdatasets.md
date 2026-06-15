### 2026-06-15 Maps Platform Datasets Proto Path Discovery
- **Context**: Implementing direct types for mapsplatformdatasets.cnrm.cloud.google.com/v1alpha1.
- **Problem**: The proto file was located at `google/maps/mapsplatformdatasets/v1/dataset.proto` in googleapis, which was not matched by KCC's default proto generator configuration in `dev/tools/controllerbuilder/generate-proto.sh`. The default only included `google/*/*.proto` and `google/cloud/*/*/*.proto`.
- **Solution**: Added `${THIRD_PARTY}/googleapis/google/maps/*/*/*.proto` to the `protoc` inputs inside `dev/tools/controllerbuilder/generate-proto.sh` to ensure nested maps-related services are compiled, and ran generating types with correct `--service google.maps.mapsplatformdatasets.v1`.
- **Impact**: Future agents working on Google Maps platform resources can directly compile and build their types without encountering proto message not found errors.
