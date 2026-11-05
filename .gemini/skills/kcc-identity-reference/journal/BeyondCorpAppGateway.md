# BeyondCorpAppGateway Identity and Reference Journal

- The resource kind is `BeyondCorpAppGateway`, located under `apis/beyondcorp/v1alpha1`.
- CAI metadata format: `//beyondcorp.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/appGateways/{{APP_GATEWAY}}`.
- Template format: `gcpurls.Template[BeyondCorpAppGatewayIdentity]("beyondcorp.googleapis.com", "projects/{project}/locations/{location}/appGateways/{appGateway}")`.
- Added `ExternalRef *string` (`json:"externalRef,omitempty"`) to `BeyondCorpAppGatewayStatus` in `apis/beyondcorp/v1alpha1/beyondcorpappgateway_types.go`.
- Regenerated CRDs and identity code via `controllerbuilder generate-identity`:
  - Added `${CONTROLLERBUILDER} generate-identity` to `apis/beyondcorp/generate.sh`.
  - Generated `BeyondCorpAppGatewayIdentity` with `status.externalRef` drift detection cross-check in `GetIdentity`.
  - Generated `BeyondCorpAppGatewayRef` and canonical `refs.Normalize` in `beyondcorpappgateway_reference.generated.go`.
- Added unit tests in `beyondcorpappgateway_identity_test.go` verifying `FromExternal`, `GetIdentity`, and `Normalize`.
