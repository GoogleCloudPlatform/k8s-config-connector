# AIPlatformCachedContent Journal

## Resource
- Kind: `AIPlatformCachedContent`
- Proto Message: `google.cloud.aiplatform.v1.CachedContent`
- Service: `aiplatform`

## Notes & Findings
- **Proto Source Path Pinning**: `apis/aiplatform/generate.sh` pins `--proto-source-path ${REPO_ROOT}/.build/googleapis-e7513c828b8db536fa12fc5493d541f323aedf90.pb` for existing resources to preserve CRD backward compatibility against `master`.
- **Identity & Reference**: Generated `aiplatformcachedcontent_identity.generated.go` and `aiplatformcachedcontent_reference.generated.go` with pattern `projects/{project}/locations/{location}/cachedContents/{cachedContent}`.
- **Empty Oneof Variant Struct**: `spec.tools[].googleMaps` is an empty `oneof` variant struct (`message GoogleMaps {}`). Added `"aiplatformcachedcontents.aiplatform.cnrm.cloud.google.com": true` to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` with comment `// GCP known pattern: Empty oneof option (spec.tools[].googleMaps)`.
- **Mapper Helpers**: `LatLng` and `CachedContent_UsageMetadata` mapper helper functions (`*_FromProto` / `*_ToProto`) were added to `pkg/controller/direct/aiplatform/aiplatformcachedcontent_mapper.go`. Also `generate-mapper` service list was updated with `google.type`.
