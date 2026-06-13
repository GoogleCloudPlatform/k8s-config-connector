# Journal: ComputeSSLCertificate Transition to Direct KRM Types

## Observations

1. **Schema Integrity & Omitempty Tag:**
   - The root field `spec` of the `ComputeSSLCertificate` resource had to be marked as `required` in the OpenAPI v1 CRD schema to remain compatible with the baseline. This was achieved by removing the `omitempty` tag on the `Spec` field of the `ComputeSSLCertificate` struct (using ``json:"spec"`` instead of ``json:"spec,omitempty"``).

2. **Excluding Standard Fields (No ProjectRef/ExternalRef):**
   - Following the strict schema compatibility mandate, we omitted standard direct-controller fields such as `spec.projectRef`, `status.externalRef`, and `status.observedState` from `sslcertificate_types.go` because they were not present in the baseline CRD.

3. **OneOf OpenAPI Constraints & Custom Validation Rules:**
   - The baseline CRD contains strict `oneOf` constraints for `spec.certificate` and `spec.privateKey` to enforce that exactly one of `value` and `valueFrom` is configured.
   - To preserve these constraints in the generated CRD schema, we updated the validation generator script (`scripts/add-validation-to-crds/parse-crds.go`) to include `ComputeSSLCertificate` in the list of resources processed for the signature `value,valueFrom`.

4. **Handcoded Spec & Status Mappers:**
   - The GCP API protobuf defines `Certificate` and `PrivateKey` as simple `string` pointers, but KRM models them as custom secret-key reference structures (`SslcertificateCertificate` and `SslcertificatePrivateKey`).
   - To resolve these type mismatches and prevent Go compilation errors in automatically generated mappers, we wrote custom spec and status mapping functions in `pkg/controller/direct/compute/sslcertificate_mapper.go`.
   - The generator successfully detected these handcoded functions and excluded the conflicting auto-generated assignments from `mapper.generated.go`.
