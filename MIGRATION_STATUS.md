# Migration Status: ComputeTargetHTTPSProxy (Terraform -> Direct)

**Date:** Wednesday, February 25, 2026
**Status:** PR #1 (API & Mappers) Complete. Ready for PR #2 (MockGCP).

## Summary
The goal is to migrate the `ComputeTargetHTTPSProxy` resource to the KCC Direct Controller while maintaining strict backward compatibility for `v1beta1`.

## Completed Work (PR #1)
- [x] **API Types**: Defined in `apis/compute/v1beta1/targethttpsproxy_types.go`.
    - Implemented with strict parity to existing CRD.
    - Added CEL validation for mutual exclusivity of `sslCertificates` and `certificateManagerCertificates`.
    - Used pointer types for optional fields.
- [x] **Identity**: Defined in `apis/compute/v1beta1/targethttpsproxy_identity.go`.
    - Supports Global (`projects/{{project}}/global/targetHttpsProxies/{{name}}`) and Regional identifiers.
- [x] **References**: Defined in `apis/compute/v1beta1/targethttpsproxy_reference.go`.
    - Handles refs for `ComputeURLMap`, `ComputeSSLCertificate`, `ComputeSSLPolicy`, `CertificateManagerCertificate`, `CertificateManagerCertificateMap`, and `NetworkSecurityServerTLSPolicy`.
- [x] **Manual Mappers**: Implemented in `pkg/controller/direct/compute/targethttpsproxy_mapper.go`.
    - Handles the dual-purpose `ssl_certificates` proto field.
- [x] **Fuzzer**: Registered in `pkg/controller/direct/compute/targethttpsproxy_fuzzer.go`.
- [x] **Code Generation**: Ran `dev/tasks/generate-crds` and `deepcopy-gen`. Packages compile successfully.

## Pending Work (PR #2 & #3)
- [ ] **MockGCP Updates**:
    - Verify if `SetUrlMap`, `SetSslCertificates`, `SetQuicOverride`, and `SetCertificateMap` methods are needed in `mockcompute`.
    - Implement/fix any gaps in the mock storage for TargetHttpsProxies.
- [ ] **Direct Controller**:
    - Implement `pkg/controller/direct/compute/targethttpsproxy_controller.go`.
    - Implement reference resolution using `targethttpsproxy_resolverefs.go`.
- [ ] **Routing & Validation**:
    - Update `pkg/controller/resourceconfig/static_config.go` to route to the direct controller.
    - Add E2E fixtures in `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computetargethttpsproxy`.
    - Record golden logs using `hack/record-gcp` and verify with `hack/compare-mock`.

## Key Constraints
- **Status Fields (Backward Compatibility)**: Fields `creationTimestamp`, `proxyId`, and `selfLink` are maintained directly under `status` to match the existing Terraform-based CRD structure. In a standard Direct Controller, these would typically be placed under `status.observedState`, but they are kept in their legacy positions to avoid breaking changes for existing users.
- **Identity Formats**: Must use Asset Inventory format for `externalRef`.
- **Surgical Changes**: Avoid adding new fields (like `authorizationPolicy`) to maintain minimal diff from Terraform implementation until migration is stable.
- **Labels**: Not supported by the underlying GCP proto; excluded from Spec.

## Technical Decisions & Rationale
- **Acronym Casing (Backward Compatibility)**: We have retained `serverTlsPolicyRef` (instead of `serverTLSPolicyRef`) and `proxyId` (instead of `proxyID`). This matches the existing CRD and avoids breaking changes. These fields have existing exceptions in `tests/apichecks/testdata/exceptions/acronyms.txt`.
- **Dual-Purpose Certificate Mapping**: The GCP proto field `ssl_certificates` is overloaded. In the KRM, we split this into `sslCertificates` (for classic SSL certificates) and `certificateManagerCertificates` (for Certificate Manager certificates). The `targethttpsproxy_mapper.go` contains manual logic to prefix/un-prefix these URLs with `//certificatemanager.googleapis.com` based on their type.
- **Mutual Exclusivity**: Per GCP requirements, `sslCertificates` and `certificateManagerCertificates` are mutually exclusive. We have enforced this via CEL validation in the `ComputeTargetHTTPSProxySpec` struct.
- **API Check Golden Files**: `missingfields.txt` and `missingrefs.txt` were updated via `WRITE_GOLDEN_OUTPUT=1`. The new CRD surfaced that `projectRef` was not explicitly covered in the existing `basic` test fixtures, which triggered the `apichecks` failure.
- **MockGCP Implementation Gaps**: Initial research shows that while `SetUrlMap` and `SetQuicOverride` are implemented in `mockcompute`, `SetSslCertificates` and `SetCertificateMap` are missing and will likely be required for the direct controller's `Update` logic.

## Helpful Commands
- `go build ./apis/compute/v1beta1/...`
- `go build ./pkg/controller/direct/compute/...`
- `./dev/tasks/generate-crds`
