# PrivateCACertificate KRM Transition Journal

## Learnings & Observations

### 1. Reusing Shared Sub-Types Safely (Pruning Risks)
In package `v1beta1`, `PrivateCACertificate` shares sub-types (such as `ObjectID` and `KeyUsage`) with `PrivateCACAPool` and `PrivateCACertificateAuthority`. 
When defining custom structures locally (like `CertificateAdditionalExtensions`), if you annotate them with `// +kcc:proto` of a shared proto message, `generate-types` thinks you are overriding that proto type globally. This comments out the canonical generated type in `types.generated.go`, breaking compilation for any other resources that depend on it.
**Resolution**: Do not annotate custom/renamed nested structures with proto tags shared across the package. Instead, omit the `// +kcc:proto` annotation on the overridden types and implement targeted handcoded mappings inside `mapper.go`.

### 2. Defining Nested Structs as Pointers
The legacy CRD definitions used required sub-structures as value fields inside nested specs. Moving them to pointers with `// +required` annotations generates the exact same OpenAPI schema, but completely avoids pointer-to-value assignment compilation issues inside generated mappers.

### 3. Preserving required Validation inside Slices
When a sub-type (like `CertificateUnknownExtendedKeyUsages` or `CertificatePolicyIds`) is inside an array/slice, kubebuilder schema generation may omit the required constraint of its inner fields (like `objectIdPath`).
**Resolution**: Explicitly define local types with the `// +kubebuilder:validation:Required` marker on the required fields to force the generator to include them in the element items schema.

### 4. API Version
`PrivateCACertificate` has always been served as a `v1beta1` resource in the upstream master branch. It was never in `v1alpha1`. For this reason, we implement and generate it cleanly in `v1beta1`, ensuring we don't serve any other API versions. No alpha-to-beta promotion was done in this PR.
