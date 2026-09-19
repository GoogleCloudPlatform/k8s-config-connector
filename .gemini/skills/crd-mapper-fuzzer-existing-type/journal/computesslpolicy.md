# ComputeSSLPolicy Direct KRM Migration Journal

## Learnings & Observations

- **Omission of ProjectRef & Location**:
  In the baseline CRD, `ComputeSSLPolicy` does not contain `projectRef` or `location` fields in its `Spec`. It is a global resource whose parent project/namespace is historically handled implicitly (e.g., inherited). For strict schema-compatibility, we must completely omit these standard direct-controller fields from the spec.

- **Status Field Constraints**:
  To ensure 100% schema-compatibility with the baseline CRD, standard direct-controller fields like `observedState` and `externalRef` under `Status` must be excluded. The final generated CRD from our Go types produced an absolutely empty diff when compared against the baseline CRD.

- **Duplicate GVK declarations**:
  The existing reference file `sslpolicy_reference.go` had already declared `ComputeSSLPolicyGVK`. Redeclaring it in `sslpolicy_types.go` led to a build failure during `go vet`. We resolved this by removing the duplicate GVK declaration from `sslpolicy_types.go`.

- **Acronym Convention Alignment and Automated Mapping**:
  By renaming `MinTlsVersion` to `MinTLSVersion` (while preserving `json:"minTlsVersion,omitempty"` to prevent any schema changes), we aligned the Go field name with the acronym naming convention and resolved the "near miss" error in the generator. This allowed the mapping generator (`generate-mapper`) to automatically and fully map all spec and status fields of `ComputeSSLPolicy`. Consequently, the handwritten `computesslpolicy_mapper.go` file was completely deleted, maximizing code reuse and relying 100% on automatically generated mappers.
