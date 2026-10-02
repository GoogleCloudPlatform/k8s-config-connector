# SaaSServiceMgmtUnit Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Kind and Acronym Casing**:
   - Resource kind is `SaaSServiceMgmtUnit`, matching `services_acronyms.json` and `cp_resources_list.json`.
   - The GCP API service is `google.cloud.saasplatform.saasservicemgmt.v1beta1` with CRD group `saasservicemgmt.cnrm.cloud.google.com/v1alpha1`.

2. **Schema & Field Mapping**:
   - `Location` is a required pointer `*string` marked with kubebuilder immutable rule (`self == oldSelf`).
   - `ProjectRef` is a required pointer `*refsv1beta1.ProjectRef`.
   - Immutable spec fields: `UnitKind`, `Tenant`, `ManagementMode`.
   - Child types `Unit_MaintenanceSettings`, `UnitCondition`, `UnitDependencyObservedState`, `UnitVariable` were mapped to proto types cleanly.
   - Standard omissions `Name`, `Labels`, `Annotations` are managed via KRM metadata and identity.
   - `Release`, `OngoingOperations`, `PendingOperations`, `ScheduledOperations`, `Dependents`, `Dependencies`, `InputVariables`, `OutputVariables`, `State`, `Conditions`, `SystemManagedState`, `SystemCleanupAt`, `Uid`, `Etag`, `CreateTime`, `UpdateTime` are correctly placed in `ObservedState` as output-only fields.

3. **Identity & Reference**:
   - Resource URL template is `projects/{project}/locations/{location}/units/{unit}` for host `saasservicemgmt.googleapis.com`.
   - Implemented `SaaSServiceMgmtUnitIdentity` implementing `identity.IdentityV2` and `identity.Resource`.
   - Implemented `SaaSServiceMgmtUnitRef` implementing `refs.Ref` using `refs.Normalize`.
   - Added comprehensive unit tests in `saasservicemgmtunit_identity_test.go` using `cmp.Diff`.
