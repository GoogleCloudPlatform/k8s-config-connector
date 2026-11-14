# Compute Service Journals

### [2026-06-20] Direct Migration of ComputeSecurityPolicy
- **Context**: Implementing the direct controller and test fixtures for `ComputeSecurityPolicy` (issue #10602).
- **Problem**: `ComputeSecurityPolicy` has distinct clients for global vs. regional resources, namely `SecurityPoliciesClient` and `RegionSecurityPoliciesClient`. Both are needed to fully support both global and regional KCC resource configurations without separate controllers.
- **Solution**: Evaluated the resource identity to dynamically detect if `Region` is specified. If `Region` is non-empty, we instantiate and invoke regional client methods (e.g., `InsertRegionSecurityPolicyRequest`, `PatchRegionSecurityPolicyRequest`, `DeleteRegionSecurityPolicyRequest`, and regional `Get`), otherwise we default to the global client methods. Copied GCP fingerprint metadata back to desired model in `Update` to avoid any modification conflicts.
- **Impact**: Highly modular, unified adapter pattern that completely avoids duplicating reconcilers for global and regional variants of Compute API resources.

### [2026-06-23] Direct Migration of ComputeAutoscaler
- **Context**: Implementing the direct controller and test fixtures for `ComputeAutoscaler` (issue #10727).
- **Problem**: `ComputeAutoscaler` has a target reference (`TargetRef`) referring to `ComputeInstanceGroupManager`, which is a legacy DCL-reconciled resource. Since MockGCP does not support mocking the DCL-managed `InstanceGroupManagers` API, any `ComputeAutoscaler` E2E tests containing this dependency will be skipped on MockGCP.
- **Solution**: We implemented the direct controller for `ComputeAutoscaler` and registered `k8s.ReconcilerTypeDirect` in `static_config.go` under `SupportedControllers`. To support resolving the referenced resource, we wrote a helper `resolveComputeAutoscalerRefs` to dynamically fetch the status fields (`status.externalRef` or `status.selfLink`) of the referenced `ComputeInstanceGroupManager`. We verified mappers and fuzzer correctness by running targeted fuzz roundtrip tests with `FOCUS=ComputeAutoscaler`.
- **Impact**: Dynamic testing ensures the direct controller is automatically vetted, and any mappers or fuzzers are fully verified, even if real API dependencies prevent executing the full E2E test against MockGCP.

### [2026-06-28] Direct Migration of ComputeTargetHTTPSProxy and MockGCP Fixes
- **Context**: Implementing the direct controller and test fixtures for `ComputeTargetHTTPSProxy` (issue #10925).
- **Problem**: Ran into two MockGCP bugs:
  1. `CertificateManagerCertificates` references in the direct controller are mapped with `//certificatemanager.googleapis.com/` prefix, which was not handled by mock compute `globalTargetHttpsProxies` / `regionalTargetHttpsProxies` (they only trimmed `https://certificatemanager.googleapis.com/v1/`).
  2. `sslCertificates` references parsed by `parseGlobalSslCertificateName` and `parseRegionalSslCertificateName` assumed a relative URL with exactly 5 or 6 tokens, and failed with `sslCertName <nil> is not valid` when absolute compute URLs were passed by the direct controller.
- **Solution**:
  1. Updated the mock `globaltargethttpsproxiesv1.go` and `regionaltargethttpsproxiesv1.go` to support both `https://` and `//` prefixes when parsing certificate manager certificate strings.
  2. Added trimming of standard absolute compute URL prefixes to `parseGlobalSslCertificateName` and `parseRegionalSslCertificateName` inside MockGCP compute SSL certificate services, ensuring relative and absolute URL formats are parsed seamlessly.
- **Impact**: Ensures complete compatibility and seamless alignment between direct controllers and MockGCP/legacy controllers for SSL certificates and Certificate Manager integrations.

### [2026-06-29] Direct Migration of ComputeNodeTemplate
- **Context**: Implementing the direct controller and test fixtures for `ComputeNodeTemplate` (issue #10951).
- **Problem**: `ComputeNodeTemplate` is completely immutable in GCP (cannot be updated after creation). In direct reconciliation, if a resource cannot be updated, we must still detect diffs and return an error.
- **Solution**: Implemented the direct controller under `pkg/controller/direct/compute/` and supported client options using `newNodeTemplatesClient` in `client.go`. In the adapter's `Update` method, we perform standard `compareComputeNodeTemplate` comparison on spec fields. If a diff is found, we surface the diff back to the user via `structuredreporting.ReportDiff` and return a descriptive error stating that the resource is immutable.
- **Impact**: Ensures strict correctness and immutability compliance while surfacing exact diffs and errors dynamically on the resource status. Tests passed successfully against MockGCP.

### [2026-06-29] Direct Migration of ComputeInstanceGroupManager
- **Context**: Implementing the direct controller and test fixtures for `ComputeInstanceGroupManager` (issue #10958).
- **Problem**: `ComputeInstanceGroupManager` can be zonal or regional, meaning both `InstanceGroupManagersClient` and `RegionInstanceGroupManagersClient` are required. Furthermore, GCP returns full URLs for references (e.g., `InstanceTemplate`, `HealthCheck`, zones in `DistributionPolicy`), whereas users may specify relative names, causing false diffs.
- **Solution**:
  1. Added zonal and regional REST client instantiations under `pkg/controller/direct/compute/client.go`.
  2. Implemented `ComputeInstanceGroupManagerAdapter` routing requests to zonal/regional clients depending on the location of the resource.
  3. Pre-normalized all URL fields (e.g., `InstanceTemplate`, `HealthCheck`, `TargetPools`, and zone names) to relative paths using `refs.TrimComputeURIPrefix` and `lastComponent` before diff comparison.
  4. Duplicated `regionalcomputeinstancegroupmanager` and `zonalcomputeinstancegroupmanager` to `-direct` variants with direct-reconciler annotations, successfully recorded golden files and HTTP traffic against MockGCP, and passed the full presubmit suites.
- **Impact**: Clean package-isolated direct controller supporting both zonal and regional instance group managers while completely avoiding any drift loops or false diffs.

### [2026-10-15] Direct Migration of ComputeProjectMetadata
- **Context**: Implementing direct controller and test fixtures for `ComputeProjectMetadata` (issue #13992).
- **Implementation**:
  1. Added `newProjectsClient` to `gcpClient` under `pkg/controller/direct/compute/client.go` using GAPIC REST client `compute.NewProjectsRESTClient`.
  2. Implemented `ComputeProjectMetadataAdapter` under `pkg/controller/direct/compute/computeprojectmetadata_controller.go`, supporting `Find`, `Create`, `Update`, `Delete`, and `Export`.
  3. Registered dual controllers in `pkg/controller/resourceconfig/static_config.go` with `SupportedControllers: [Direct, Terraform]` keeping default as `Terraform` for brownfield backward compatibility.
  4. Executed `hack/compare-mock` and verified E2E test scenarios for direct reconciliation, legacy fallback, and `TestMigrationToDirect`.
- **Impact**: Successfully enables direct reconciliation for `ComputeProjectMetadata` with 0 diffs during takeover and steady-state re-reconciliation.

### [2026-10-30] Direct Migration of ComputeManagedSSLCertificate
- **Context**: Implementing direct controller and test fixtures for `ComputeManagedSSLCertificate` (issue #14697).
- **Problem**: In GCP, SSL certificates are immutable and cannot be updated after creation. The legacy controller had server defaults (`type: "MANAGED"`, `name`) and status fields that needed clean handling during brownfield takeover without drift loops or diff mismatches.
- **Solution**:
  1. Added `newSslCertificatesClient` in `pkg/controller/direct/compute/client.go` using GAPIC REST client (`compute.NewSslCertificatesRESTClient`).
  2. Implemented `ComputeManagedSSLCertificateAdapter` supporting `Find`, `Create`, `Update`, `Delete`, and `Export`.
  3. In `Update`, used `common.CompareBrownfieldSpec` to automatically adopt server-defaulted fields (`type: "MANAGED"`, `name`) via its `MergeUnsetFields` step, reporting diffs via `structuredreporting.ReportDiff` and returning an error explaining immutability if diffs exist.
  4. Registered dual controllers in `pkg/controller/resourceconfig/static_config.go` with `SupportedControllers: [Direct, Terraform]` while keeping `DefaultController` as legacy `Terraform`.
  5. Validated against real GCP and MockGCP with recorded test fixtures and golden comparisons.
- **Impact**: Provides a robust, idiomatic direct controller for `ComputeManagedSSLCertificate` with zero-write re-reconciliation and clean immutability enforcement.

### [2026-11-10] Greenfield: Prepopulate types, identity, and judgement queue for ComputeCrossSiteNetwork
- **Context**: Greenfield implementation of `ComputeCrossSiteNetwork` (issue #15355).
- **Proto & Service**: `CrossSiteNetwork` is a GA resource under Google Cloud Compute Engine v1 (and present in `cloud.google.com/go/compute/apiv1/computepb`), but was not present in the pinned googleapis SHA in `.build/third_party/googleapis`. Following the existing pattern in `dev/tools/controllerbuilder/generate-proto.sh` (e.g. `SecurityPolicyAssociation`), we patched `google/cloud/compute/v1/compute.proto` to include `CrossSiteNetwork` message and service definitions, keeping the pinned commit intact.
- **Prepopulation**: Ran `generate-types --prepopulate-resource` and `generate-identity` with pattern `projects/{project}/global/crossSiteNetworks/{crossSiteNetwork}`.
- **Release Bundle**: Added `ComputeCrossSiteNetwork` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Validation**: Added entry to `alpha-missingfields.txt` via golden test runner. All tests in `apis/compute/...`, `tests/apichecks/...`, and `dev/tasks/build-release-bundle` passed cleanly.

### [2026-11-11] Greenfield Prepopulate of ComputeInstantSnapshot
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for Greenfield resource `ComputeInstantSnapshot` (issue #15354).
- **Problem**: When configuring `apis/compute/generate.sh`, passing `--prepopulate-resource` to the existing multi-resource `generate-types` command enabled `--emit-message-maps`, causing `AutoscalingPolicy` to be un-commented in `apis/compute/v1alpha1/types.generated.go`. This triggered a compilation conflict in `mapper.generated.go` because a handwritten `FixedOrPercent_v1alpha1_FromProto` had a mismatched return type.
- **Solution**:
  1. Kept the existing multi-resource `generate-types` invocation untouched without `--prepopulate-resource`.
  2. Added a separate, dedicated `generate-types` invocation with `--prepopulate-resource` and `generate-identity` strictly for `ComputeInstantSnapshot`.
  3. Added `--prepopulate-resource` to `generate-mapper` (alongside `--multiversion`), which correctly emits inline message map conversions without compiler collisions.
  4. Added `ComputeInstantSnapshot` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  5. Added non-idiomatic proto comment exception to `tests/apichecks/testdata/exceptions/comments.txt` and updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Complies with Greenfield Step 1 review standards by enabling `--prepopulate-resource` on `generate-mapper` while safely isolating `--prepopulate-resource` in `generate-types` for newly scaffolded resources in shared services.

### [2026-11-11] Greenfield Judgement Queue Triage for ComputeInstantSnapshot
- **Context**: Triaging `judgement_queue.yaml` and finalizing direct KRM types, identity, and references for `ComputeInstantSnapshot` (issue #15378).
- **Findings**:
  1. `ComputeInstantSnapshot` can be created either for a zonal disk or a regional disk. In GCP CAIS, it has two name formats: `//compute.googleapis.com/projects/{project}/zones/{zone}/instantSnapshots/{name}` and `//compute.googleapis.com/projects/{project}/regions/{region}/instantSnapshots/{name}`.
  2. Because `controllerbuilder generate-identity` only generates a single template pattern and cannot express dual zonal and regional templates for a resource, `ComputeInstantSnapshotIdentity` was promoted to handwritten `computeinstantsnapshot_identity.go`, mirroring `ComputeDisk` (`ZonalComputeInstantSnapshotIdentityFormat` and `RegionalComputeInstantSnapshotIdentityFormat`).
  3. `SourceDisk` is a reference to a `ComputeDisk` and was converted to `SourceDiskRef *computev1beta1.ComputeDiskRef` with `json:"sourceDiskRef,omitempty"` and `// +kcc:proto:field=google.cloud.compute.v1.InstantSnapshot.source_disk`.
  4. Top-level `Labels` and output-only fields (`architecture`, `diskSizeGB`, `labelFingerprint`, `region`, `resourceStatus`, `satisfiesPzi`, `satisfiesPzs`, `sourceDiskID`, `status`, `zone`) were moved to `ComputeInstantSnapshotObservedState` or omitted from Spec.
  5. In `apis/compute/generate.sh`, removed the isolated prepopulation block and added `ComputeInstantSnapshot:InstantSnapshot` to the main `generate-types` invocation without `--prepopulate-resource`.
- **Impact**: Cleanly resolved all 24 entries in `apis/compute/judgement_queue.yaml`. All unit tests, CAIS golden tests, and API checks pass.

### [2026-11-11] Greenfield Controller, Fuzzer, and E2E Fixtures for ComputeInstantSnapshot
- **Context**: Implementing the direct controller, fuzzer, and real-GCP recorded E2E fixtures for `ComputeInstantSnapshot` (issue #15386).
- **Findings**:
  1. `InstantSnapshots` in Compute Engine are regional or zonal and require separate clients: `compute.InstantSnapshotsClient` and `compute.RegionInstantSnapshotsClient`. Added `newInstantSnapshotsClient` and `newRegionInstantSnapshotsClient` in `pkg/controller/direct/compute/client.go`.
  2. Instant snapshots in GCP are immutable (no update or patch method exists on the client). In `Update`, `common.CompareGreenfieldSpec` checks for changes and returns a clear error if modified.
  3. When `SourceDisk` is returned from GCP `Get`, it has the full URL format (`https://www.googleapis.com/...`), while `SourceDiskRef` on KRM objects or resolved by `NormalizeGreenfieldReferences` uses the relative format (`projects/...`). A normalization hook was supplied to `CompareGreenfieldSpec` to strip the compute URL prefix, preventing false diffs during re-reconciliation.
  4. Fuzzer round-trip testing (`FOCUS=ComputeInstantSnapshot`) required marking unmapped proto fields (`.params`, `.source_instant_snapshot_group`, `.source_instant_snapshot_group_id`) as `f.Unimplemented_NotYetTriaged`.
  5. Minimal and maximal fixtures were created and recorded against real GCP with `RECORD_AUDIT_PROBE=1 ./hack/record-gcp`. Both minimal and maximal fixtures passed all 7 automated validation checks.
- **Impact**: Clean direct controller and fuzzer implementation for `ComputeInstantSnapshot` verified end-to-end against real GCP.

### [2026-11-11] MockGCP Implementation and Alignment for ComputeInstantSnapshot
- **Context**: Implementing MockGCP service handlers and aligning mock golden logs with authentic real GCP traffic for `ComputeInstantSnapshot` (issue #15421).
- **Implementation**:
  1. Implemented `InstantSnapshotsV1` (zonal) and `RegionInstantSnapshotsV1` (regional) in `mockgcp/mockcompute/instantsnapshotsv1.go`, handling Get, Insert, Delete, and SetLabels.
  2. Registered both gRPC servers and HTTP mux handlers in `mockgcp/mockcompute/service.go`.
  3. Added `ComputeInstantSnapshot` to supported GroupKinds for MockGCP testing in `config/tests/samples/create/harness.go`.
  4. Updated empty labels fingerprint in `mockgcp/mockcompute/utils.go` to return standard Compute Engine constant `42WmSpB8rSM=` when no labels are specified.
  5. Populated `SatisfiesPzi: true` on instant snapshots to match real GCP behavior observed in KRM status objects.
  6. Automatically populated source disk metadata (`SourceDiskId`, `DiskSizeGb`) from mock storage when available, defaulting to standard 10GB pd size.
  7. Ran `hack/compare-mock` and verified 0 diff in golden objects, generated `_http_mock.log` files, and verified clean pass of `TestGoldenLogAlignment` and `./dev/ci/presubmits/tests-e2e-fixtures-suite compute`.
- **Impact**: Provides hermetic MockGCP simulation for both zonal and regional instant snapshots with zero diff against authentic real GCP golden logs.

### [2026-11-11] Greenfield Prepopulate of ComputeHealthCheckService
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for Greenfield resource `ComputeHealthCheckService` (issue #15400).
- **Implementation**:
  1. Configured dedicated `generate-types` and `generate-identity` commands for `ComputeHealthCheckService:HealthCheckService` in `apis/compute/generate.sh` using `--prepopulate-resource`.
  2. Executed `./apis/compute/generate.sh`, which automatically generated `computehealthcheckservice_types.go`, `computehealthcheckservice_identity.generated.go`, `computehealthcheckservice_reference.generated.go`, `computehealthcheckservice_identity_generated_test.go`, and CRD manifests.
  3. Added `ComputeHealthCheckService` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Added `computehealthcheckservices.compute.cnrm.cloud.google.com` entries to `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Prepopulation completed mechanically with all queue entries left open for Step 1b triage. Idempotency verified and all tests pass.

### [2026-11-11] Greenfield Judgement Queue Triage for ComputeHealthCheckService
- **Context**: Triaging `apis/compute/judgement_queue.yaml` and finalizing direct KRM types, identity, and references for `ComputeHealthCheckService` (issue #15418).
- **Findings & Resolutions**:
  1. `ComputeHealthCheckService` is a regional resource under `projects/{project}/regions/{region}/healthCheckServices/{healthcheckservice}`. In `computehealthcheckservice_types.go`, `spec.location` was made a required string (`Location string`) representing the region name, matching standard Compute conventions.
  2. `spec.region` was confirmed as output-only per proto comments ("It is not settable as a field in the request body") and moved to `ComputeHealthCheckServiceObservedState.Region`.
  3. Verified documentation link and updated `+kcc:source:resource-docs` to `https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthCheckServices`, removing guess annotations.
  4. Server-set fields (`creationTimestamp`, `id`, `kind`, `selfLink`) were confirmed output-only and retained in `ComputeHealthCheckServiceObservedState` with guess markers removed.
  5. Deterministic identity and reference generation via `controllerbuilder generate-identity` were retained in `apis/compute/generate.sh`, while `ComputeHealthCheckService:HealthCheckService` was moved into the primary `generate-types` invocation.
  6. Modeled `spec.healthChecks` and `spec.networkEndpointGroups` as typed KCC references (`HealthCheckRefs []*computev1beta1.ComputeHealthCheckRef` and `NetworkEndpointGroupRefs []*computev1beta1.ComputeNetworkEndpointGroupRef`) referencing existing KCC resources `ComputeHealthCheck` and `ComputeNetworkEndpointGroup`.
  7. Resolved all 9 entries in `apis/compute/judgement_queue.yaml` for `ComputeHealthCheckService`.
- **Impact**: Cleanly resolved the entire judgement queue for `ComputeHealthCheckService`. CRDs and mapper bindings updated, and all unit tests and API checks pass.

### [2026-11-11] Greenfield Prepopulate of ComputeVMExtensionPolicy
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, and `judgement_queue.yaml` for `ComputeVMExtensionPolicy` under `apis/compute/v1alpha1/` (issue #15353).
- **Proto & Service Details**: `VmExtensionPolicy` and `ZoneVmExtensionPolicies` service exist in `cloud.google.com/go/compute/apiv1/computepb` and newer googleapis revisions, but are missing from the pinned googleapis commit (1765b559c4). Following the established pattern for Compute (e.g. `SecurityPolicyAssociation`), patched `dev/tools/controllerbuilder/generate-proto.sh` to inject `VmExtensionPolicy` and `ZoneVmExtensionPolicies` into `google/cloud/compute/v1/compute.proto`.
- **Implementation**:
  1. Configured `apis/compute/generate.sh` with `generate-types` (including `ComputeVMExtensionPolicy`), `generate-identity`, and `--prepopulate-resource` on `generate-mapper`, and added manual mapper helpers in `pkg/controller/direct/compute/computevmextensionpolicy_mapper.go` for `ExtensionPolicies` map conversion.
  2. Registered `ComputeVMExtensionPolicy` in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  3. Generated types, CRD, and `apis/compute/judgement_queue.yaml` with open triage findings.
  4. Verified conformance tests and updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-11] Greenfield Judgement Queue Triage for ComputeVMExtensionPolicy
- **Context**: Triaging `judgement_queue.yaml` and finalizing direct KRM types, identity, and references for `ComputeVMExtensionPolicy` (issue #15399).
- **Findings & Resolutions**:
  1. **Pattern & Service Hierarchy**: `ComputeVMExtensionPolicy` is managed via the `ZoneVmExtensionPolicies` REST API under `projects/{project}/zones/{zone}/vmExtensionPolicies/{vmextensionpolicy}`. Locked in this pattern in `apis/compute/generate.sh` via `--pattern ComputeVMExtensionPolicy=projects/{project}/zones/{zone}/vmExtensionPolicies/{vmextensionpolicy}` and resolved `identity-multi-pattern`.
  2. **Location**: Retained `spec.location` as a required string representing the zone name (`location-parent-unknown`), matching Compute conventions (e.g. `ComputeInstantSnapshot`).
  3. **Documentation**: Updated `+kcc:source:resource-docs=` to `https://cloud.google.com/compute/docs/reference/rest/v1/zoneVmExtensionPolicies` and removed `// +kcc:guess=source-link reason=verify-resource-docs-link`, resolving `verify-resource-docs-link`.
  4. **Output-Only Fields**: `globalResourceLink`, `managedByGlobal`, `state`, and `updateTimestamp` were identified as output-only per proto comments and moved from `ComputeVMExtensionPolicySpec` to `ComputeVMExtensionPolicyObservedState`. Removed resolved `// +kcc:guess=placement` comments.
  5. Custom Mappers: Implemented `ComputeVMExtensionPolicySpec_v1alpha1_FromProto` and `ComputeVMExtensionPolicySpec_v1alpha1_ToProto` in `pkg/controller/direct/compute/computevmextensionpolicy_mapper.go` so `generate-mapper` avoids attempting to map the moved output-only fields to Spec.
  6. Testing: Added unit tests in `computevmextensionpolicy_identity_test.go` covering `FromExternal`, `ParentString`, `GetIdentity` (including status drift detection), and `ComputeVMExtensionPolicyRef.ValidateExternal`/`Normalize`. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for the removed Spec fields.

### [2026-11-13] Greenfield Prepopulate of ComputeGlobalVMExtensionPolicy
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, and `judgement_queue.yaml` for `ComputeGlobalVMExtensionPolicy` under `apis/compute/v1alpha1/` (issue #15534).
- **Proto & Service Details**: `GlobalVmExtensionPolicy` and `GlobalVmExtensionPolicies` service exist in the Compute v1 Discovery document (`projects/{project}/global/vmExtensionPolicies/{globalVmExtensionPolicy}`), but were missing from the pinned googleapis commit (1765b559c4). Following the established pattern for Compute (e.g. `SecurityPolicyAssociation`, `VmExtensionPolicy`), patched `dev/tools/controllerbuilder/generate-proto.sh` to inject `GlobalVmExtensionPolicy` messages and `GlobalVmExtensionPolicies` service into `google/cloud/compute/v1/compute.proto`.
- **Implementation**:
  1. Added `GlobalVmExtensionPolicy` definitions to `dev/tools/controllerbuilder/generate-proto.sh`.
  2. Registered `ComputeGlobalVMExtensionPolicy` in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  3. Configured `apis/compute/generate.sh` with dedicated `generate-types --prepopulate-resource` and `generate-identity` for `ComputeGlobalVMExtensionPolicy:GlobalVmExtensionPolicy`.
  4. Added `--skip-proto` for `GlobalVmExtensionPolicy` messages to `generate-mapper` in `apis/compute/generate.sh` to prevent undefined symbol compilation errors against the currently pinned `cloud.google.com/go/compute` client library version (`v1.63.0`) which has not yet exposed Go SDK types for global VM extension policies.
  5. Scaffolded `apis/compute/v1alpha1/computeglobalvmextensionpolicy_types.go`, generated identity, reference, deepcopy, CRD, and `judgement_queue.yaml` with open triage findings.
  6. Verified compilation of `apis/compute/...` and `pkg/controller/direct/compute/...`, conformance tests, and updated golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.

### [2026-11-13] Greenfield Judgement Queue Triage for ComputeGlobalVMExtensionPolicy
- **Context**: Triaged `judgement_queue.yaml` and finalized direct KRM types, identity, and references for `ComputeGlobalVMExtensionPolicy` (issue #15558).
- **Findings & Resolutions**:
  1. **Pattern & Service Hierarchy**: `ComputeGlobalVMExtensionPolicy` is managed via the `GlobalVmExtensionPolicies` REST API under `projects/{project}/global/vmExtensionPolicies/{globalvmextensionpolicy}`. Locked in this pattern in `apis/compute/generate.sh` via `--pattern ComputeGlobalVMExtensionPolicy=projects/{project}/global/vmExtensionPolicies/{globalvmextensionpolicy}` and resolved `identity-multi-pattern`.
  2. **Location**: Removed `spec.location` because `ComputeGlobalVMExtensionPolicy` is global and has no location/region/zone segment (`location-parent-unknown`). Made `spec.projectRef` required with `// +required`.
  3. **Documentation**: Updated `+kcc:source:resource-docs=` to `https://cloud.google.com/compute/docs/reference/rest/v1/globalVmExtensionPolicies` and removed `// +kcc:guess=source-link reason=verify-resource-docs-link`, resolving `verify-resource-docs-link`.
  4. **Output-Only Fields & RolloutOperation Structs**:
     - `scopedResourceStatus` and `updateTimestamp` were identified as output-only per proto comments and moved from `ComputeGlobalVMExtensionPolicySpec` to `ComputeGlobalVMExtensionPolicyObservedState`.
     - `GlobalVmExtensionPolicyRolloutOperation` was split into desired state (`GlobalVMExtensionPolicyRolloutOperation` containing only user-configurable `rolloutInput`) and observed state (`GlobalVMExtensionPolicyRolloutOperationObservedState` containing output-only `rolloutStatus`).
     - Moved `GlobalVMExtensionPolicyRolloutOperation*` structs to `computeglobalvmextensionpolicy_types.go` to preserve `map[string]...` in `LocationRolloutStatus` without requiring repo-wide `--emit-message-maps`.
  5. **Triage Queue**: Resolved all 27 entries for `ComputeGlobalVMExtensionPolicy` in `apis/compute/judgement_queue.yaml`.
  6. **Testing**: Added unit tests in `computeglobalvmextensionpolicy_identity_test.go` covering `FromExternal`, `ParentString`, `GetIdentity` (including status drift detection), and `ComputeGlobalVMExtensionPolicyRef.ValidateExternal`/`Normalize`. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for the removed Spec fields.




### [2026-11-13] Greenfield Prepopulate of ComputeNotificationEndpoint
- **Context**: Prepopulating direct KRM types, identity, reference, and `judgement_queue.yaml` for Greenfield resource `ComputeNotificationEndpoint` under `apis/compute/v1alpha1/` (issue #15537).
- **Implementation**:
  1. Configured dedicated `generate-types` and `generate-identity` commands for `ComputeNotificationEndpoint:NotificationEndpoint` in `apis/compute/generate.sh` using `--prepopulate-resource` and `--include-skipped-output`.
  2. Executed `./apis/compute/generate.sh`, which automatically generated `computenotificationendpoint_types.go`, `computenotificationendpoint_identity.generated.go`, `computenotificationendpoint_reference.generated.go`, `computenotificationendpoint_identity_generated_test.go`, and CRD manifests.
  3. Added `ComputeNotificationEndpoint` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with entries for `computenotificationendpoints.compute.cnrm.cloud.google.com`.
  5. Kept all entries in `apis/compute/judgement_queue.yaml` with `status: open` for Step 1b triage.
- **Impact**: Prepopulation completed mechanically with all queue entries left open for Step 1b triage. Generation is idempotent and all unit tests and API checks pass cleanly.

### [2026-11-13] Greenfield Judgement Queue Triage for ComputeNotificationEndpoint
- **Context**: Triaging `apis/compute/judgement_queue.yaml` and finalizing direct KRM types, identity, and references for `ComputeNotificationEndpoint` (issue #15557).
- **Findings & Resolutions**:
  1. **Pattern & Service Hierarchy**: `ComputeNotificationEndpoint` is a regional resource under `projects/{project}/regions/{region}/notificationEndpoints/{notificationendpoint}`, managed via `RegionNotificationEndpoints`.
  2. **Location & Parent**: Modeled parentage using `spec.projectRef` (`*refsv1beta1.ProjectRef`) and `spec.location` (`Location string`) as required fields, matching established Compute conventions.
  3. **Documentation**: Updated `+kcc:source:resource-docs=` to `https://cloud.google.com/compute/docs/reference/rest/v1/regionNotificationEndpoints` and removed `// +kcc:guess=source-link reason=verify-resource-docs-link`, resolving `verify-resource-docs-link`.
  4. **Output-Only Fields**: `region` is confirmed output-only per proto comments ("It is not settable as a field in the request body") and moved from `ComputeNotificationEndpointSpec` to `ComputeNotificationEndpointObservedState`. Server-set fields (`creationTimestamp`, `id`, `kind`, `selfLink`) were confirmed output-only in `ObservedState`, and all guess placement annotations were removed.
  5. **Script & Mapper Alignment**: In `apis/compute/generate.sh`, removed the isolated prepopulation block and added `ComputeNotificationEndpoint:NotificationEndpoint` to the primary `generate-types` invocation without `--prepopulate-resource`. Mapper bindings regenerated cleanly.
  6. **Testing**: Added comprehensive unit tests in `computenotificationendpoint_identity_test.go` covering `FromExternal`, `ParentString`, `GetIdentity` (including status drift detection), and `ComputeNotificationEndpointRef.ValidateExternal`/`Normalize`. Re-recorded golden exceptions in `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Cleanly resolved all 9 entries in `apis/compute/judgement_queue.yaml`. All unit tests, API checks, and linters pass cleanly.

### [2026-11-12] Greenfield: Prepopulate direct KRM types, identity, and judgement queue for ComputeHealthAggregationPolicy
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `ComputeHealthAggregationPolicy` (issue #15466).
- **Proto & Service**:
  - `HealthAggregationPolicy` is defined in Google Cloud Compute API v1 (`google/cloud/compute/v1/compute.proto`), and managed via the `RegionHealthAggregationPolicies` service (`projects/{project}/regions/{region}/healthAggregationPolicies/{healthAggregationPolicy}`).
  - Pinned `apis/git.versions` (`1765b559c4`) predates the inclusion of `HealthAggregationPolicy` in `compute.proto` (which was merged upstream in revision 20261026). Similar to `SecurityPolicyAssociation`, patched `dev/tools/controllerbuilder/generate-proto.sh` to include `HealthAggregationPolicy` and `RegionHealthAggregationPolicies` during proto compilation.
- **Scaffolding**:
  - Ran `generate-types` with `--prepopulate-resource` for `ComputeHealthAggregationPolicy:HealthAggregationPolicy` and `generate-identity`.
  - Added `ComputeHealthAggregationPolicy` to `apis/compute/generate.sh` and added to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Regenerated CRDs and updated test exception for alpha missing fields (`tests/apichecks/testdata/exceptions/alpha-missingfields.txt`).
- **Verification**:
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-12] Greenfield: Triage judgement queue and finalize direct KRM types and identity for ComputeHealthAggregationPolicy
- **Context**: Triaged and finalized direct KRM types, identity, and references for `ComputeHealthAggregationPolicy` (issue #15482) following Greenfield Step 1b (`judgement/SKILL.md`).
- **Judgement Queue Resolutions & Learnings**:
  1. `untriaged-bulk-generation`: Accepted after validating fields, conventions, references, and omissions.
  2. `verify-resource-docs-link`: Corrected resource documentation URL to `https://cloud.google.com/compute/docs/reference/rest/v1/regionHealthAggregationPolicies`.
  3. `location-parent-unknown`: Verified `HealthAggregationPolicy` is a regional resource located under `projects/{project}/regions/{region}/healthAggregationPolicies/{healthAggregationPolicy}`. Kept `spec.location` as a required string representing the region, matching Compute service conventions (`ComputeNetworkAttachment`, `ComputeHealthCheckService`).
  4. `location-or-parent-ref` for `spec.region`: The proto field `region` is an output-only URL string populated by GCP server. Removed `spec.region` and moved to `status.observedState.region`.
  5. `server-set-field-placed` for `creationTimestamp`, `id`, `kind`, `selfLink`, and `selfLinkWithId`: Confirmed server-generated output-only fields; verified their placement under `status.observedState`.
  6. Incorporated `ComputeHealthAggregationPolicy:HealthAggregationPolicy` into the primary `generate-types` and `generate-identity` commands in `apis/compute/generate.sh`.
  7. Regenerated types, CRDs, mappers, deepcopy code, and updated `alpha-missingfields.txt`.
- **Verification**:
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-13] Greenfield Controller, Fuzzer, and E2E Fixtures for ComputeVMExtensionPolicy
- **Context**: Implementing direct controller, fuzzer, and real-GCP recorded E2E fixtures for `ComputeVMExtensionPolicy` (issue #15517).
- **Findings & Implementation**:
  1. **Client**: Utilized GAPIC REST client `gcp.NewZoneVmExtensionPoliciesRESTClient(ctx, opts...)` matching the `ZoneVmExtensionPolicies` service in GCP Compute Engine API v1.
  2. **Controller Logic**: Implemented `VmExtensionPolicyAdapter` satisfying all 5 canonical Greenfield controller standards:
     - REST client options configured via `m.config.RESTClientOptions()`.
     - `common.NormalizeGreenfieldReferences(ctx, reader, obj, nil)` called first in `AdapterForObject`.
     - `ComputeVMExtensionPolicyIdentity` (`obj.GetIdentity`) used to obtain `project`, `zone`, and `vmextensionpolicy` name.
     - `common.CompareGreenfieldSpec` with `structuredreporting.ReportDiff` for change detection in `Update`.
     - `setStatus` called on all reconciliation paths, including no-op updates.
  3. **RPC Lifecycle**: Handled `Insert` (POST), `Get` (GET), `Update` (PATCH), and `Delete` (DELETE) operations, awaiting the resulting Compute `ZoneOperations` LROs via `op.Wait(ctx)`.
  4. **Fuzzer**: Autogenerated `computevmextensionpolicy_fuzzer.generated.go` via `controllerbuilder generate-fuzzer --prepopulate-resource`. Verified round-trip fidelity with `FOCUS=ComputeVMExtensionPolicy go test ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`.
  5. **E2E Golden Testing**: Authored minimal and maximal test fixtures with valid versions (`2.57.0` and `2.58.0`) and configurations for `ops-agent`. Successfully executed and recorded golden HTTP traffic against real GCP (`RECORD_AUDIT_PROBE=1 ./hack/record-gcp`), verifying V1-V7 validation checks, zero-diff re-reconciliation, and clean deletion.

### [2026-11-13] MockGCP Implementation and Alignment for ComputeVMExtensionPolicy
- **Context**: Implementing MockGCP service handlers and aligning mock golden logs with authentic real GCP traffic for `ComputeVMExtensionPolicy` (issue #15556).
- **Findings & Implementation**:
  1. **Mock Service**: Implemented `ZoneVmExtensionPoliciesV1` in `mockgcp/mockcompute/computevmextensionpolicies.go`, handling `Get`, `Insert`, `Update`, and `Delete` requests using the GAPIC computepb types (`cloud.google.com/go/compute/apiv1/computepb`), backing them with mock storage and zonal LROs (`startZonalLRO`).
  2. **HTTP Routing**: Registered routes (`GET`, `POST`, `PATCH`, `DELETE`) for `/compute/v1/projects/{project}/zones/{zone}/vmExtensionPolicies` in `mockgcp/mockcompute/service.go` using `runtime.ForwardResponseMessage` to ensure correct HTTP response header and payload formatting.
  3. **Harness Support**: Added `ComputeVMExtensionPolicy` to the mock supported kinds in `config/tests/samples/create/harness.go`.
  4. **Normalization**: Configured normalization in `mockgcp/mockcompute/normalize.go` and `tests/e2e/normalize.go` for `updateTimestamp` (`mockgcpregistry.PlaceholderTimestamp` / `mockgcpregistry.PlaceholderTime`) matching server timestamps.
  5. **Validation**: Ran `hack/compare-mock` for minimal and maximal fixtures; both passed with zero diff in golden objects and mock HTTP logs (`_http_mock.log`), and verified `./dev/ci/presubmits/tests-e2e-fixtures-suite compute` succeeded.

### [2026-11-13] Greenfield Prepopulate of ComputeInterconnectGroup
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, and `judgement_queue.yaml` for Greenfield resource `ComputeInterconnectGroup` under `apis/compute/v1alpha1/` (issue #15613).
- **Proto & Service Details**: `InterconnectGroup` is defined in Google Cloud Compute API v1 (`google/cloud/compute/v1/compute.proto`), and managed via the `InterconnectGroups` service (`projects/{project}/global/interconnectGroups/{interconnectgroup}`).
- **Implementation**:
  1. Configured `apis/compute/generate.sh` with `generate-types` (adding `ComputeInterconnectGroup:InterconnectGroup` with `--prepopulate-resource`) and `generate-identity`.
  2. Executed `./apis/compute/generate.sh`, automatically generating `computeinterconnectgroup_types.go`, `computeinterconnectgroup_identity.generated.go`, `computeinterconnectgroup_reference.generated.go`, `computeinterconnectgroup_identity_generated_test.go`, and CRD manifests.
  3. Added `ComputeInterconnectGroup` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Added `computeinterconnectgroups.compute.cnrm.cloud.google.com` entries to `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
  5. Recorded review items in `apis/compute/judgement_queue.yaml` with `status: open` for Step 1b triage.
- **Verification**:
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-14] Greenfield: Triage judgement queue and finalize direct KRM types and identity for ComputeInterconnectGroup
- **Context**: Triaged and finalized direct KRM types, identity, and references for `ComputeInterconnectGroup` (issue #15621) following Greenfield Step 1b (`judgement/SKILL.md`).
- **Judgement Queue Resolutions & Learnings**:
  1. `untriaged-bulk-generation`: Resolved after triaging all fields, omissions, references, and KRM conventions.
  2. `verify-resource-docs-link`: Corrected resource documentation link to `https://cloud.google.com/compute/docs/reference/rest/v1/interconnectGroups`.
  3. `location-parent-unknown`: Verified `InterconnectGroup` is a global resource parented by project (`projects/{project}/global/interconnectGroups/{interconnectgroup}`). Removed `spec.location` and kept `spec.projectRef` (`// +required`).
  4. `possible-reference-by-sibling` on `.spec.interconnects.KEY.interconnect`: Converted to `InterconnectRef *ComputeInterconnectRef` in `InterconnectGroupInterconnect` (moved to `computeinterconnectgroup_types.go`). `generate-mapper` generates conversion to/from proto string field.
  5. `message-map-derived` on `.spec.interconnects`: Accepted; accurately models proto `map<string, InterconnectGroupInterconnect>`.
  6. `well-known-output-only-pattern-in-comment` on `configured` and `physicalStructure`: Proto comments mark `InterconnectGroupConfigured` and `InterconnectGroupPhysicalStructure` as output-only. Moved `Configured *InterconnectGroupConfigured` and `PhysicalStructure *InterconnectGroupPhysicalStructure` from `ComputeInterconnectGroupSpec` to `ComputeInterconnectGroupObservedState`. Marked corresponding `possible-reference-by-sibling` and `location-or-parent-ref` on their nested fields as `not-applicable` because status fields do not use KRM references.
  7. `server-set-field-placed` on `creationTimestamp`, `etag`, `id`, `kind`, and `selfLink`: Confirmed server-generated output-only fields; verified placement in `ComputeInterconnectGroupObservedState` and removed `+kcc:guess` comments.
- **Verification**:
  - Re-ran `./apis/compute/generate.sh` and verified CRD generation, mappers, deepcopy code.
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-14] Greenfield Step 1a: Prepopulate ComputeStoragePool
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, CRD, and `judgement_queue.yaml` for `ComputeStoragePool` (issue #15640).
- **Findings & Implementation**:
  1. **Proto & Pattern**: Compute StoragePool message is in `google.cloud.compute.v1` (`StoragePool`). Due to compute proto RPC HTTP bindings lacking `{name=...}` assignment patterns, `--pattern ComputeStoragePool=projects/{project}/zones/{zone}/storagePools/{storagePool}` was provided to `generate-identity` in `apis/compute/generate.sh` to resolve pattern ambiguity.
  2. **Generation**: Added `--prepopulate-resource` block for `ComputeStoragePool:StoragePool` and `generate-identity` in `apis/compute/generate.sh`. Ran script to generate KRM types (`computestoragepool_types.go`), IdentityV2 (`computestoragepool_identity.generated.go`), reference (`computestoragepool_reference.generated.go`), CRD (`apiextensions.k8s.io_v1_customresourcedefinition_computestoragepools.compute.cnrm.cloud.google.com.yaml`), and mapper stubs.
  3. **Release Bundle & Apichecks**: Added `ComputeStoragePool` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` via `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...`.
  4. **Validation**: Verified `apis/compute/generate.sh` idempotency, `TestComputeStoragePoolIdentity_Conformance`, apichecks, cais tests, and full package builds with zero errors.

### [2026-11-14] Greenfield Step 1b: Triage judgement queue and finalize direct KRM types and identity for ComputeStoragePool
- **Context**: Triaged and finalized direct KRM types, identity, and references for `ComputeStoragePool` (issue #15652) following Greenfield Step 1b (`judgement/SKILL.md`).
- **Judgement Queue Resolutions & Learnings**:
  1. `untriaged-bulk-generation`: Resolved after triaging all spec and status fields, omissions, docs link, output-only fields, and KRM conventions.
  2. `verify-resource-docs-link`: Verified and updated `+kcc:source:resource-docs` link to `https://cloud.google.com/compute/docs/reference/rest/v1/storagePools`.
  3. `location-parent-unknown`: Verified `StoragePool` is a zonal resource parented by project and zone (`projects/{project}/zones/{zone}/storagePools/{storagePool}`). Kept `spec.projectRef` (`*refsv1beta1.ProjectRef` with `// +required`) and `spec.location` (`string` with `// +required`, representing the zone name) matching Compute conventions.
  4. `location-or-parent-ref` and `well-known-output-only-pattern-in-comment` on `zone`: In the proto, `StoragePool.zone` is explicitly output-only (`[Output Only] URL of the zone where the storage pool resides...`). Removed `zone` from Spec and moved to `ComputeStoragePoolObservedState.zone`; zone location on creation is specified via `spec.location`.
  5. `possible-status-field-in-spec` and `well-known-output-only-pattern-in-comment` on `state`, `status`, and `resource_status`: Proto comments mark `state`, `status`, and `resource_status` (and all nested fields in `StoragePoolResourceStatus`) as output-only. Moved them from `ComputeStoragePoolSpec` to `ComputeStoragePoolObservedState`.
  6. Top-level `labels` and `label_fingerprint`: Removed top-level `labels` from `ComputeStoragePoolSpec` (reconciled via `metadata.labels`); moved `label_fingerprint` to `ComputeStoragePoolObservedState`.
  7. `server-set-field-placed` on `creationTimestamp`, `id`, `kind`, `selfLink`, and `selfLinkWithID`: Confirmed output-only server-set fields and removed guess markers.
  8. Implemented unit tests in `computestoragepool_identity_test.go` covering identity parsing, parent string formatting, spec identity resolution, and drift detection.
- **Verification**:
  - Re-ran `./apis/compute/generate.sh` and verified CRD generation, mappers, deepcopy code.
  - Re-recorded `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-14] Greenfield Prepopulate of ComputeInterconnectAttachmentGroup
- **Context**: Prepopulating direct KRM types, IdentityV2, reference, and `judgement_queue.yaml` for Greenfield resource `ComputeInterconnectAttachmentGroup` under `apis/compute/v1alpha1/` (issue #15656).
- **Proto & Service Details**: `InterconnectAttachmentGroup` is defined in Google Cloud Compute API v1 (`google/cloud/compute/v1/compute.proto`), and managed via the `InterconnectAttachmentGroups` service (`projects/{project}/global/interconnectAttachmentGroups/{interconnectattachmentgroup}`).
- **Implementation**:
  1. Added allowlist URL exception to `pkg/gcpurls/registry_test.go` for `//compute.googleapis.com/projects/{}/global/interconnectAttachmentGroups/{}`.
  2. Added `ComputeInterconnectAttachmentGroup:InterconnectAttachmentGroup` to `generate-types` and `generate-identity` in `apis/compute/generate.sh`.
  3. Added `ComputeInterconnectAttachmentGroup` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Executed `./apis/compute/generate.sh`, automatically generating `computeinterconnectattachmentgroup_types.go`, `computeinterconnectattachmentgroup_identity.generated.go`, `computeinterconnectattachmentgroup_reference.generated.go`, `computeinterconnectattachmentgroup_identity_generated_test.go`, and CRD manifests.
  5. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` with entries for `computeinterconnectattachmentgroups.compute.cnrm.cloud.google.com`.
  6. Kept all entries in `apis/compute/judgement_queue.yaml` with `status: open` for Step 1b triage.
- **Verification**:
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.

### [2026-11-14] Greenfield: Triage judgement queue and finalize direct KRM types and identity for ComputeInterconnectAttachmentGroup
- **Context**: Triaged and finalized direct KRM types, identity, and references for `ComputeInterconnectAttachmentGroup` (issue #15665) following Greenfield Step 1b (`judgement/SKILL.md`).
- **Judgement Queue Resolutions & Learnings**:
  1. `untriaged-bulk-generation`: Resolved after triaging all fields, omissions, references, and KRM conventions.
  2. `verify-resource-docs-link`: Corrected resource documentation link to `https://cloud.google.com/compute/docs/reference/rest/v1/interconnectAttachmentGroups`.
  3. `location-parent-unknown`: Verified `InterconnectAttachmentGroup` is a global resource parented by project (`projects/{project}/global/interconnectAttachmentGroups/{interconnectattachmentgroup}`). Removed `spec.location` and kept `spec.projectRef` (`// +required`).
  4. `possible-reference-by-sibling` on `.spec.interconnectGroup`: Converted to `InterconnectGroupRef *ComputeInterconnectGroupRef` in `ComputeInterconnectAttachmentGroupSpec`. `generate-mapper` automatically maps to/from proto string field.
  5. `message-map-derived` on `.spec.attachments`: Accepted; accurately models proto `map<string, InterconnectAttachmentGroupAttachment>`. Moved `InterconnectAttachmentGroupAttachment` to `computeinterconnectattachmentgroup_types.go` and converted `attachment` to `AttachmentRef *computev1beta1.ComputeInterconnectAttachmentRef`. `generate-mapper` automatically generates bidirectional conversion between `*computev1beta1.ComputeInterconnectAttachmentRef` and proto string field `attachment`.
  6. `well-known-output-only-pattern-in-comment` on `configured` and `logicalStructure`: Proto comments mark `InterconnectAttachmentGroupConfigured` and `InterconnectAttachmentGroupLogicalStructure` as output-only. Moved `Configured *InterconnectAttachmentGroupConfigured` and `LogicalStructure *InterconnectAttachmentGroupLogicalStructure` from `ComputeInterconnectAttachmentGroupSpec` to `ComputeInterconnectAttachmentGroupObservedState`. Marked corresponding `location-or-parent-ref` on nested fields as `not-applicable` because status fields do not use KRM references.
  7. `server-set-field-placed` on `creationTimestamp`, `etag`, `id`, `kind`, and `selfLink`: Confirmed server-generated output-only fields; verified placement in `ComputeInterconnectAttachmentGroupObservedState` and removed `+kcc:guess` comments.
- **Verification**:
  - Re-ran `./apis/compute/generate.sh` and verified CRD generation, mappers, deepcopy code.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
  - All unit tests (`apis/compute/...`), conformance identity tests, and API checks (`tests/apichecks/...`) pass cleanly.
