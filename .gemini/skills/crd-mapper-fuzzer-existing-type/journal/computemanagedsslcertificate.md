# ComputeManagedSSLCertificate KRM Direct Transition Journal

## Learnings & Observations

### 1. Capitalization of ID vs Id
When a protobuf field `id` maps to KRM field `certificateID` (which represents the baseline KRM field name), the generator detects the name difference. If we try to rename the Go struct field to `Id` while keeping `json:"certificateID,omitempty"`, the generator flags `Id` vs `ID` as a "(near miss)" because Go protobuf's naming convention for the `id` field generates `Id` in `pb.SslCertificate` but uses `ID` in other places, or vice-versa. 
Renaming the KRM Go struct field to `ID` (all caps) with `json:"certificateID,omitempty"` successfully maps to the proto field `Id`, which resolves the near miss and fully automates the mapper.

### 2. Multi-Version Storage version setup
For resources with both `v1alpha1` and `v1beta1` versions served, `v1alpha1` serves as the `storage: true` version in the baseline CRD. 
We MUST add the `// +kubebuilder:storageversion` annotation on the Kind struct of the `v1alpha1` version (`apis/compute/v1alpha1/sslcertificate_types.go`), and leave it off the `v1beta1` version. Doing this preserves the original storage version layout of the CRD perfectly and compiles successfully with `controller-gen`.

### 3. Type Mismatch Conversion (`*uint64` vs `*int64`)
The proto field `id` has type `*uint64`, whereas the KRM `certificateID` uses `*int64` (to represent `integer` in OpenAPI schema). Since they are differing pointer types, direct assignment fails compile-checking.
We successfully resolved this type mismatch by hand-coding the observed state conversions in `pkg/controller/direct/compute/mappers.go`. The generator automatically detected the handcoded functions and skipped generating them in `mapper.generated.go`.
