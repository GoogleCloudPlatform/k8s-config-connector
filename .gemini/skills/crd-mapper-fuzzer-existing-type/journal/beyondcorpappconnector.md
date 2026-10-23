# Journal: Direct KRM Types Implementation for BeyondCorpAppConnector

## Observations & Learnings

1. **CRD Schema Parity & Versioning**:
   - The baseline CRD `beyondcorpappconnectors.beyondcorp.cnrm.cloud.google.com` is in version `v1alpha1`.
   - Preserved legacy labels on the struct definition:
     - `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"`
     - `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"`
     - `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"`
     - `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"`
   - Marking `Spec` in `BeyondCorpAppConnector` with `// +required` ensures `required: ["spec"]` is generated in the root OpenAPI validation schema.

2. **Top-Level Struct Comments**:
   - The baseline CRD does not contain top-level descriptions for `spec` or `status`.
   - Omitted doc comments on `BeyondCorpAppConnectorSpec` and `BeyondCorpAppConnectorStatus` to avoid triggering `_no_description.spec=true` and `_no_description.status=true` schema diffs.

3. **Reference Types**:
   - The baseline CRD defines `spec.projectRef` with fields `external`, `name`, and `namespace`, without a `kind` field.
   - Used `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` rather than `refsv1beta1.ProjectRef` (which includes `kind`), ensuring the schema matches exactly and allowing `scripts/add-validation-to-crds` to automatically inject the `oneOf` constraint.

4. **Required Nested and Scalar Fields**:
   - `region`: In the baseline CRD, `region` is a required string. Implemented as non-pointer scalar `Region string \`json:"region"\``.
   - `principalInfo`: Defined `AppConnectorPrincipalInfo` and `AppConnectorPrincipalInfoServiceAccount` with `// +kcc:proto` annotations. `email` is required and modeled as non-pointer scalar `Email string \`json:"email"\``. The nested struct `ServiceAccount` uses a pointer with `// +required` to allow clean schema validation and pointer traversal.

5. **generate.sh Multi-Service Configuration**:
   - The beyondcorp API directory hosts multiple services: `google.cloud.beyondcorp.clientconnectorservices.v1` and `google.cloud.beyondcorp.appconnectors.v1`.
   - Consolidated `generate-types` in `apis/beyondcorp/generate.sh` by supplying both services comma-separated in `--service` and declaring both `--resource` entries. This produces a unified `types.generated.go` and eliminates the previous workaround of renaming `types.generated.go` to `clientconnectorservice_types.generated.go`.
