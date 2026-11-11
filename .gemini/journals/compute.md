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
  5. **Custom Mappers**: Implemented `ComputeVMExtensionPolicySpec_v1alpha1_FromProto` and `ComputeVMExtensionPolicySpec_v1alpha1_ToProto` in `pkg/controller/direct/compute/computevmextensionpolicy_mapper.go` so `generate-mapper` avoids attempting to map the moved output-only fields to Spec.
  6. **Testing**: Added unit tests in `computevmextensionpolicy_identity_test.go` covering `FromExternal`, `ParentString`, `GetIdentity` (including status drift detection), and `ComputeVMExtensionPolicyRef.ValidateExternal`/`Normalize`. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` for the removed Spec fields.

