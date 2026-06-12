# PrivateCACertificateTemplate Journal

## Observations
- `PrivateCACertificateTemplate` was transitioned from a DCL-based resource to a direct KCC resource.
- The baseline CRD and GCP PrivateCA protos contain shared nested models with `PrivateCACAPool` (e.g., `X509Parameters`, `CertificateIdentityConstraints`, `CertificateExtensionConstraints`, `ObjectID`, `X509Extension`).
- Sharing types between resources requires careful validation coordination (for example, marking nested fields like `ObjectIDPath` in `ObjectID`, and `ObjectID`/`Value` in `X509Extension` as required by removing `,omitempty` and adding `// +required` annotations).
- Defining `CertificateIdentityConstraints` in `certificatetemplate_types.go` instead of relying entirely on auto-generated types in `types.generated.go` allowed us to safely specify custom validation tags (like `// +required`) without they being overwritten by code generators.
- The GCP protobuf for `ObjectId` defines `object_id_path` as a repeated list of `int32`, which differs from DCL's Python schema representing it as `int64`. The direct KRM type uses `int32` which is correct per canonical GCP API specifications.
