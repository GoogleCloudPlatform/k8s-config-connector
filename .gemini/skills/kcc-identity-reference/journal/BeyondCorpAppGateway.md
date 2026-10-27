# BeyondCorpAppGateway Identity and Reference Journal

- The resource kind is `BeyondCorpAppGateway`, located under `apis/beyondcorp/v1alpha1`.
- CAI metadata format: `//beyondcorp.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/appGateways/{{APP_GATEWAY}}`.
- Template format: `gcpurls.Template[BeyondCorpAppGatewayIdentity]("beyondcorp.googleapis.com", "projects/{project}/locations/{location}/appGateways/{appGateway}")`.
- The resource is currently managed by the Terraform (legacy) reconciler and does not have `status.externalRef` or `status.name` in `BeyondCorpAppGatewayStatus`.
- In adherence to the skill instructions:
  - We did not alter the CRD schema.
  - In `GetIdentity`, the parsed `specIdentity` is returned without status cross-checks since no status identity field is present in the schema.
  - `Normalize` in `BeyondCorpAppGatewayRef` uses `refs.NormalizeWithFallback` (legacy fallback pattern) with a check that `status.conditions` has `Ready == True`. If not ready, it returns `""` to bubble up `k8s.NewReferenceNotReadyError`.
- Deterministic generation via `controllerbuilder generate-identity`:
  - Added `${CONTROLLERBUILDER} generate-identity` to `apis/beyondcorp/generate.sh`.
  - Handwritten `Normalize` in `beyondcorpappgateway_reference.go` is respected by `controllerbuilder`, which automatically omitted emitting the default `Normalize` in `beyondcorpappgateway_reference.generated.go`.
- Registered `apis/beyondcorp/v1alpha1` imports in `pkg/gcpurls/registry_test.go` and `pkg/cli/powertools/cais/cmd_test.go`.
- Updated the fixture golden output `pkg/test/resourcefixture/testdata/basic/beyondcorp/v1alpha1/beyondcorpappgateway/beyondcorpappgatewayautogen/_identities.yaml` with the resolved CAIS URL.
