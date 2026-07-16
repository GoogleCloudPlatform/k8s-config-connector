# MapManagement Service Journal

### 2026-07-02 MapManagement Protobuf Paths and ControllerBuilder Configuration
- **Context**: Implementing the Greenfield types for `MapManagementMapConfig` (`MapConfig` GCP resource).
- **Problem**: The GCP API and protobuf files for `mapmanagement` are located under the `google/maps/mapmanagement/v2beta` path in `googleapis`, instead of the standard `google/cloud/` prefix. Consequently, the default `protoc` command inside `dev/tools/controllerbuilder/generate-proto.sh` did not capture or compile the `v2beta` protos, causing code generation to fail with `failed to find the proto message google.maps.mapmanagement.v2beta.MapConfig: proto: not found`.
- **Solution**: 
  1. Updated the service name in the `generate-types` command argument to `google.maps.mapmanagement.v2beta`.
  2. Modified `dev/tools/controllerbuilder/generate-proto.sh` to explicitly compile the paths `${THIRD_PARTY}/googleapis/google/maps/mapmanagement/*/*.proto`.
  3. Cleared out cached `.pb` files under `.build/` to force regeneration.
- **Impact**: Allows `controllerbuilder` to correctly locate and process Google Maps API family protobufs. Subsequent map management resources will be able to generate types seamlessly.
