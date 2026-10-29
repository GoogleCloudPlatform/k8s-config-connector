# ComputeManagedSSLCertificate Journal

## Context
Implementing direct KRM types and updating `apis/compute/generate.sh` for `ComputeManagedSSLCertificate` in `apis/compute/v1beta1/`.

## Challenges & Solutions

### 1. Multi-resource Proto Mapping to `google.cloud.compute.v1.SslCertificate`
- **Context**: In GCP Compute API, both self-managed certificates (`ComputeSSLCertificate`) and managed certificates (`ComputeManagedSSLCertificate`) map to the same protobuf message `google.cloud.compute.v1.SslCertificate`.
- **Solution**: Added `--resource ComputeManagedSSLCertificate:SslCertificate` to `apis/compute/generate.sh` under the `v1beta1` section alongside `ComputeSSLCertificate:SslCertificate`.

### 2. Strict Schema Compatibility for `spec.managed`
- **Problem**: The proto message `SslCertificateManagedSslCertificate` contains output-only fields `domain_status` and `status`. Generating default types would expose these output-only fields under `spec.managed`, altering the CRD schema.
- **Solution**: Defined `ComputeManagedSSLCertificateManaged` in `computemanagedsslcertificate_types.go` containing only `Domains []string` with `// +required`, maintaining exact parity with the baseline CRD schema.

### 3. Parent Reference Without `kind` Field
- **Problem**: Baseline CRD's `spec.projectRef` lacks a `kind` field. Using `apis/refs/v1beta1.ProjectRef` causes `dev/tasks/diff-crds` to flag `+ spec.projectRef.kind=string`.
- **Solution**: Imported `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` which lacks the `kind` field.

### 4. Preserving OpenAPI Integer Without Format
- **Problem**: In the baseline CRD, `observedState.certificateID` and `status.observedGeneration` are defined as plain `type: integer` without `format: int32` or `format: int64`. Using `*int64` would cause `controller-gen` to emit `format: int64`.
- **Solution**: Typed both fields as Go `*int` so `controller-gen` generates unformatted `type: integer`.

### 5. Handcoded Mappers in `computemanagedsslcertificate_mapper.go`
- **Problem**: The proto `SslCertificate` uses `*uint64` for `id`, while KRM uses `*int` for `certificateID`. In addition, `ResourceID` in KRM maps to `Name` in protobuf.
- **Solution**: Handcoded `ComputeManagedSSLCertificateObservedState_v1beta1_{From,To}Proto` and `ComputeManagedSSLCertificateSpec_v1beta1_{From,To}Proto` in `pkg/controller/direct/compute/computemanagedsslcertificate_mapper.go`. `generate-mapper` automatically detects these non-generated functions and skips generating conflicting duplicates in `mapper.generated.go`.

### 6. Storage Version and Multi-version CRD
- **Context**: The baseline CRD was generated via Terraform service mappings with `storageVersion: v1alpha1` and `v1alpha1ToV1beta1: true`.
- **Solution**: Promoted `v1beta1` to storage version using `// +kubebuilder:storageversion`, while preserving `v1alpha1` via `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"`.
