# ComputeOrganizationSecurityPolicyRule Direct Type Transition Journal

## Overview
`ComputeOrganizationSecurityPolicyRule` (`compute.cnrm.cloud.google.com/v1alpha1`) maps to the GCP protobuf message `SecurityPolicyRule` in `google.cloud.compute.v1`.

## Learnings & Observations

- **Proto Package Version (`google.cloud.compute.v1`):**
  `SecurityPolicyRule` is defined in `google.cloud.compute.v1` (`cloud.google.com/go/compute/apiv1/computepb`), consistent with other Compute resources such as `ComputeOrganizationSecurityPolicy` and `ComputeOrganizationSecurityPolicyAssociation`.
  - In `apis/compute/generate.sh`, specify:
    `--resource ComputeOrganizationSecurityPolicyRule:SecurityPolicyRule`.
  - In `generate-mapper`, use `--service google.cloud.compute.v1`.

- **Strict Schema Compatibility:**
  The baseline CRD schema must remain completely unchanged:
  - `spec.action` (string, required)
  - `spec.description` (*string, optional)
  - `spec.direction` (*string, optional)
  - `spec.enableLogging` (*bool, optional)
  - `spec.match` (required object: `OrganizationSecurityPolicyRuleMatch`)
    - `match.config` (required object: `OrganizationSecurityPolicyRuleMatchConfig`)
      - `config.destIpRanges` ([]string, optional)
      - `config.layer4Config` (required array: `[]OrganizationSecurityPolicyRuleMatchConfigLayer4Config`)
        - `layer4Config[].ipProtocol` (string, required)
        - `layer4Config[].ports` ([]string, optional)
      - `config.srcIpRanges` ([]string, optional)
    - `match.description` (*string, optional)
    - `match.versionedExpr` (*string, optional)
  - `spec.policyId` (string, required)
  - `spec.preview` (*bool, optional)
  - `spec.resourceID` (*string, optional)
  - `spec.targetResources` ([]string, optional)
  - `spec.targetServiceAccounts` ([]string, optional)
  - `status.conditions` and `status.observedGeneration`
  - In the Go struct, used `PolicyID` with `json:"policyId"` tag to respect acronym conventions while preserving the exact schema field name.
  - Omitted top-level doc comments on the `Spec` and `Status` struct definitions to avoid introducing descriptions on `.spec` and `.status`, which would trigger `_no_description` diffs in `dev/tasks/diff-crds`.
  - Retained `cnrm.cloud.google.com/tf2crd: "true"` label on the CRD struct.

- **Hand-Coded Mappers:**
  Handcoded conversion functions were added to `pkg/controller/direct/compute/mappers.go` alongside other compute mappers:
  - `Action` is a non-pointer string in KRM but `*string` pointer in `computepb.SecurityPolicyRule`.
  - `ResourceID` in KRM represents the priority of the rule (maps to `Priority *int32` in the proto).
  - `Match` has a custom nested structure (`OrganizationSecurityPolicyRuleMatch`), and several fields (`Direction`, `EnableLogging`, `TargetResources`, `TargetServiceAccounts`) exist in GCP organization security policy rules in the REST API but are not present on `pb.SecurityPolicyRule` in `cloud.google.com/go/compute/apiv1/computepb`.
  `generate-mapper` recognized these existing functions in `mappers.go` and automatically skipped duplicate generation in `mapper.generated.go`.

- **Verification:**
  - Ran `dev/tasks/diff-crds` which returned completely empty output, proving 100% schema compatibility.
  - Ran unit tests verifying the mapper roundtrips and nil safety.
