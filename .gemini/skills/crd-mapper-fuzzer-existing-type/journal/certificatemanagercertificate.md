# Journal: CertificateManagerCertificate Transition to Direct KRM Types

## Observations

1. **Schema Integrity and Compatibility:**
   - Pre-existing metadata labels such as `cnrm.cloud.google.com/tf2crd=true` and `cnrm.cloud.google.com/stability-level=stable` were explicitly preserved on the `CertificateManagerCertificate` Go struct to maintain correct legacy controller mapping.
   - The root field `spec` was marked with the KubeBuilder validation tag `// +required` to ensure that it aligns with the baseline CRD validation schema.
   - Added standard fields like `projectRef` without a `kind` property by using `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` as recommended by the skill instructions.

2. **Parsing Comments Gotcha:**
   - The GCP API documentation comment for `IssuanceConfigRef` contained `projects/*/locations/*/certificateIssuanceConfigs/*.`. The compiler/generator interpreted the `*/` in the comment as the premature end of a block comment `/* ... */`, leading to syntax/parsing errors because the rest of the comment was treated as actual Go code.
   - This was solved by rewriting the format description in the comment as `projects/{{project}}/locations/{{location}}/certificateIssuanceConfigs/{{name}}`.

3. **Handcoded Reference/Mapper Fields:**
   - In `apis/certificatemanager/v1beta1/certificate_types.go`, we defined custom local references `CertificateDNSAuthorizationRef` and `CertificateIssuanceConfigRef` with standard fields `external`, `name`, and `namespace`.
   - Because of these custom structs and field differences compared to the standard generated `Certificate_SelfManagedCertificate` and `Certificate_ManagedCertificate` types, the automatic mapping generator skipped generating their conversions.
   - We successfully resolved this by implementing hand-coded `ToProto` and `FromProto` mapping functions in a separate Go file `pkg/controller/direct/certificatemanager/certificate_mapper.go`.
