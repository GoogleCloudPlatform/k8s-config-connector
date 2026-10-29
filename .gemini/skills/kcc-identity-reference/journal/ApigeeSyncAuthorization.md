# ApigeeSyncAuthorization Identity and Reference Journal

## Observations & Learnings

- **MockGCP Protos & Hand-Written Identity/Ref**: Apigee APIs do not have standard `cloud.google.com/go/apigee/apiv1/apigeepb` protobuf packages in the Google Cloud Go SDK; instead MockGCP defines `mockgcp.cloud.apigee.v1`. Because `controllerbuilder generate-identity` rejects MockGCP protobuf packages, the identity, reference, and unit test files were hand-written following canonical patterns.
- **Singleton Subresource Pattern**: `ApigeeSyncAuthorization` represents a singleton subresource under an Apigee organization with the URL format `organizations/{organization}/syncAuthorization`. The template placeholder `{organization}` maps to `Organization string` in `ApigeeSyncAuthorizationIdentity`.
- **Identity Resolution from Spec**: In the brownfield CRD, `ApigeeSyncAuthorization` does not declare a separate `organizationRef`; instead, the organization is passed via `spec.resourceID` (defaulting to `metadata.name`). Prefixes such as `organizations/` and `projects/` are trimmed during identity parsing.
- **No Schema Changes / Legacy Fallback**: `ApigeeSyncAuthorizationStatus` does not contain `externalRef` or `name`. Per skill rules, the CRD schema was left untouched. In `ApigeeSyncAuthorizationRef.Normalize()`, `refs.NormalizeWithFallback` was used to ensure the referenced resource's `Ready` condition is `True` before building the external URL from spec.
- **Cloud Asset Inventory Exception**: `//apigee.googleapis.com/organizations/{}/syncAuthorization` is not tracked by Cloud Asset Inventory (`cloudassetinventory_names.jsonl`). An entry was added to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` so `TestRegisteredTemplatesMatchCAI` continues to pass.
