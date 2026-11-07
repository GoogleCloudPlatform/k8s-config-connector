# SecuritycenterSecurityHealthAnalyticsCustomModule Journal

## Summary

Triage `judgement_queue.yaml` and finalize direct KRM types, identity, and reference implementation for `SecuritycenterSecurityHealthAnalyticsCustomModule`.

## Judgement Queue Triage

The following queue entries were triaged and resolved:

1. `reason: untriaged-bulk-generation`
   - Verified that `SecuritycenterSecurityHealthAnalyticsCustomModuleSpec` maps to proto `google.cloud.securitycenter.v1.SecurityHealthAnalyticsCustomModule`.
   - Updated `Spec` to support multi-parent hierarchy (`projectRef`, `folderRef`, `organizationRef`) with CEL validation ensuring at most one parent is specified.
   - Updated `spec.resourceID` comment to document that the ID is server-generated and acquire-only.
   - Status: `resolved` / `resolution: edited`.

2. `reason: identity-multi-pattern`
   - The proto defines 3 resource patterns:
     - `organizations/{organization}/securityHealthAnalyticsSettings/customModules/{custom_module}`
     - `folders/{folder}/securityHealthAnalyticsSettings/customModules/{custom_module}`
     - `projects/{project}/securityHealthAnalyticsSettings/customModules/{custom_module}`
   - Extended `Spec` with optional `projectRef`, `folderRef`, and `organizationRef` fields with CEL validation (`(has(self.projectRef) ? 1 : 0) + (has(self.folderRef) ? 1 : 0) + (has(self.organizationRef) ? 1 : 0) <= 1`).
   - Regenerated identity via `controllerbuilder generate-identity`, which generates multi-parent identity supporting all three scopes (projects, folders, organizations).
   - Status: `resolved` / `resolution: edited`.

3. `reason: identity-server-generated-id`
   - The GCP API `CreateSecurityHealthAnalyticsCustomModuleRequest` takes `parent` and `security_health_analytics_custom_module`. In GCP, the resource name is ignored on create and reset with a server-generated numeric ID.
   - Clarified in `spec.resourceID` doc comment that the ID is server-generated and acquire-only (not defaulting to `metadata.name`).
   - Regenerated `identity.ServerGeneratedIdentity` implementation and verification tests.
   - Status: `resolved` / `resolution: edited`.

## Verification

- Ran `apis/securitycenter/generate.sh`.
- Ran unit tests for `apis/securitycenter/...`.
- Ran `tests/apichecks/...` validating no missing refs, well-formed judgement queue, and CRD backward compatibility.
- Ran `go test ./apis/securitycenter/...`.
