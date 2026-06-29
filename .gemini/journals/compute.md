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
