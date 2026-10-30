# ComputeManagedSSLCertificate Migration Journal

## Overview
Migrated `ComputeManagedSSLCertificate` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Client & API:**
   - Used Google Compute Engine REST GAPIC client `compute.NewSslCertificatesRESTClient`.
   - Added `newSslCertificatesClient` to `pkg/controller/direct/compute/client.go`.
   - Uses `Insert`, `Get`, `Delete` methods on `SslCertificatesClient`.

2. **Resource Immutability:**
   - In GCP, SSL certificates are immutable and cannot be updated once created.
   - `Update` detects diffs using `compareComputeManagedSSLCertificate` (`mappers.OnlySpecFields` + `common.DiffForTopLevelFields`). If a diff exists, it reports the diff via `structuredreporting.ReportDiff` and returns an error explaining that the resource is immutable.

3. **Status & Export:**
   - Status maps `observedState` (`certificateID`, `creationTimestamp`, `selfLink`) using `ComputeManagedSSLCertificateObservedState_v1beta1_FromProto`.
   - `Export` maps `ComputeManagedSSLCertificateSpec_v1beta1_FromProto` and populates `spec.projectRef.external`.

4. **Testing & Golden Files:**
   - Ran against real GCP with `hack/record-gcp` and verified with `hack/compare-mock`.
   - Re-reconciliation test passes cleanly with 0 diffs.
   - Golden comparison files (`_http_old_controller.log`, `_http_old_controller_mock.log`, `_final_object_old_controller.golden.yaml`, `_final_object.diff`) generated and aligned.
