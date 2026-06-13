# Journal: ComputeSharedVPCHostProject KRM Type Migration

## Observations
- `ComputeSharedVPCHostProject` represents a Google Cloud Compute project being registered as a Shared VPC Host Project.
- The resource's original CRD is extremely minimal:
  - It contains an empty `spec` structure (no fields).
  - Its `status` structure only contains `conditions` and `observedGeneration`.
- The corresponding Google Cloud Compute API proto message is `Project`.
- Because the baseline CRD does not have any fields under `spec`, we handcoded `project_types.go` as an empty spec resource, keeping only the standard `conditions` and `observedGeneration` fields in the Status, avoiding adding fields like `projectRef` or `externalRef` to retain strict schema compatibility.

## Implementation Details
1. Configured `apis/compute/v1beta1/generate.sh` to include `--resource ComputeSharedVPCHostProject:Project`.
2. Created a handcoded KRM type file `project_types.go` under `apis/compute/v1beta1/` using the 2026 copyright header.
3. Included the `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"` annotation on the KRM struct to ensure that the legacy Terraform reconciler continues to register and manage the resource until a direct controller is implemented.
4. Ran `apis/compute/v1beta1/generate.sh` to generate the mappers and CRDs.
5. Ran `make manifests` to build all final CRDs and SupportedGVK metadata.
6. Ran `make generate-go-client` to update Go CRD clients.
7. Verified strict schema compatibility with `dev/tasks/diff-crds`.
8. Updated the resource report via the CRD generation hooks (`docs/reports/crd_report.csv` and `docs/reports/crd_report.md`).
