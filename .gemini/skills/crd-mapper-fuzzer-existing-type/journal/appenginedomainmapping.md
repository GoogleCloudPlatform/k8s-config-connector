# AppEngineDomainMapping Direct Type Implementation Journal

## Observations & Implementation Steps

1. **Service Proto Inclusion**:
   - The proto for App Engine Admin API is located under `google/appengine/v1/` (`google/appengine/v1/domain_mapping.proto`).
   - `google/appengine/*/*.proto` is already included in `PROTO_FILES` in `dev/tools/controllerbuilder/generate-proto.sh`.

2. **Schema Compatibility & Acronym Field Mappings**:
   - Baseline CRD fields under `spec.sslSettings`:
     - `certificateId` (lowercase 'd')
     - `pendingManagedCertificateId` (lowercase 'd')
     - `sslManagementType` (required enum)
   - To achieve 100% strict schema compatibility:
     - Defined Go struct fields using uppercase acronyms matching proto conventions: `CertificateID *string` with tag `json:"certificateId,omitempty"` and `PendingManagedCertificateID *string` with tag `json:"pendingManagedCertificateId,omitempty"`.
     - Marked `SSLManagementType string` with `// +required` and tag `json:"sslManagementType"` as required non-pointer string to adhere to KCC required scalar conventions. Handcoded `DomainmappingSslSettings_v1alpha1_FromProto` and `ToProto` in `pkg/controller/direct/appengine/mappers.go` using `direct.ValueOf(direct.Enum_FromProto(...))` and `direct.Enum_ToProto[...](mapCtx, &in.SSLManagementType)`.

3. **CRD Root Schema & Descriptions**:
   - The baseline CRD does not mark `spec` as required at the root level; removing `// +required` from the `Spec` field on `AppEngineDomainMapping` avoids `_required.spec=true` diffs.
   - Omitted doc comments on `AppEngineDomainMappingSpec` and `AppEngineDomainMappingStatus` to preserve `_no_description` parity with the baseline CRD.

4. **Identity Resolution**:
   - Proto resource path: `apps/{app}/domainMappings/{domain_mapping}`.
   - For App Engine, `{app}` represents the project ID.
   - Implemented `getIdentityFromAppEngineDomainMappingSpec` in `apis/appengine/v1alpha1/appenginedomainmapping_identity.go` resolving `refs.ResolveProjectID` and `refs.GetResourceID`.

5. **Status Mapping**:
   - The baseline CRD status exposes `name` and `resourceRecords`.
   - Adding `// +kcc:status:proto=google.appengine.v1.DomainMapping` on `AppEngineDomainMappingStatus` allowed `generate-mapper` to automatically generate `AppEngineDomainMappingStatus_v1alpha1_FromProto` and `DomainmappingResourceRecords_v1alpha1_FromProto`/`ToProto`.
   - Neither `observedState` nor `externalRef` was present in the baseline CRD status, and neither was added.

6. **Validation & Verification**:
   - `dev/tasks/diff-crds` passes with zero diffs.
   - Unit tests in `pkg/controller/direct/appengine/mapper_test.go` and identity conformance tests in `apis/appengine/v1alpha1/` pass cleanly.
