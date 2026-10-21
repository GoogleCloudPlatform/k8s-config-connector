# ComputeOrganizationSecurityPolicyAssociation Direct Type Transition Journal

## Overview
`ComputeOrganizationSecurityPolicyAssociation` (`compute.cnrm.cloud.google.com/v1alpha1`) maps to the GCP protobuf message `SecurityPolicyAssociation` in `google.cloud.compute.v1beta`.

## Learnings & Observations

- **Proto Package Version (`v1beta`):**
  Unlike many Compute resources that reside in `google.cloud.compute.v1`, `SecurityPolicyAssociation` is defined in `google.cloud.compute.v1beta`.
  - In `apis/compute/generate.sh`, specify the fully-qualified proto name:
    `--resource ComputeOrganizationSecurityPolicyAssociation:google.cloud.compute.v1beta.SecurityPolicyAssociation`.
  - In `generate-mapper`, include the `v1beta` package in the `--service` flag:
    `--service google.cloud.compute.v1,google.cloud.compute.v1beta`.

- **Strict Schema Compatibility:**
  The baseline CRD schema must remain completely unchanged:
  - `spec.attachmentId` (string, required)
  - `spec.policyId` (string, required)
  - `spec.resourceID` (string, optional)
  - `status.displayName` (string, optional)
  - `status.conditions` and `status.observedGeneration`
  - In the Go struct, used `AttachmentID` and `PolicyID` with `json:"attachmentId"` and `json:"policyId"` tags to respect acronym conventions while preserving the exact schema field names.
  - Omitted top-level doc comments on the `Spec` and `Status` struct definitions to avoid introducing descriptions on `.spec` and `.status`, which would trigger `_no_description` diffs in `dev/tasks/diff-crds`.
  - Retained `cnrm.cloud.google.com/tf2crd: "true"` label on the CRD struct.

- **Hand-Coded Mappers:**
  Handcoded conversion functions were implemented in `pkg/controller/direct/compute/computeorganizationsecuritypolicyassociation_mapper.go`:
  - `AttachmentID` and `PolicyID` are non-pointer strings in KRM but `*string` pointers in `computepb.SecurityPolicyAssociation`.
  - `policyId` in KRM maps to `SecurityPolicyId` in the proto message.
  - `resourceID` in KRM maps to `Name` in the proto message.
  `generate-mapper` recognized these existing functions and automatically skipped duplicate generation in `mapper.generated.go`.

- **Verification:**
  - Ran `dev/tasks/diff-crds` which returned completely empty output, proving 100% schema compatibility.
  - Ran unit tests verifying the mapper roundtrips and nil safety.
