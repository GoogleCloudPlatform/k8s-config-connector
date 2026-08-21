# ComputeSSLPolicy Direct KRM Migration Journal

## Learnings & Observations

- **Omission of ProjectRef & Location**:
  In the baseline CRD, `ComputeSSLPolicy` does not contain `projectRef` or `location` fields in its `Spec`. It is a global resource whose parent project/namespace is historically handled implicitly (e.g., inherited). For strict schema-compatibility, we must completely omit these standard direct-controller fields from the spec.

- **Status Field Constraints**:
  To ensure 100% schema-compatibility with the baseline CRD, standard direct-controller fields like `observedState` and `externalRef` under `Status` must be excluded. The final generated CRD from our Go types produced an absolutely empty diff when compared against the baseline CRD.

- **Duplicate GVK declarations**:
  The existing reference file `sslpolicy_reference.go` had already declared `ComputeSSLPolicyGVK`. Redeclaring it in `sslpolicy_types.go` led to a build failure during `go vet`. We resolved this by removing the duplicate GVK declaration from `sslpolicy_types.go`.

- **Automatic Skipper of Mappers**:
  By specifying `ComputeSSLPolicySpec_v1beta1_FromProto`, `ComputeSSLPolicySpec_v1beta1_ToProto`, `ComputeSSLPolicyStatus_v1beta1_FromProto`, and `ComputeSSLPolicyStatus_v1beta1_ToProto` in a handcoded `sslpolicy_mapper.go` file under `pkg/controller/direct/compute/`, the code generator automatically skipped generating conflicting versions inside `mapper.generated.go`.
