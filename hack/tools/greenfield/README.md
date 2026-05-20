# KCC Resource Coverage Analysis Tool

This tool calculates the coverage of Google Cloud Platform (GCP) resources in Kubernetes Config Connector (KCC) relative to the published definitions in the `googleapis/googleapis` repository.

It is used to identify gaps in resource coverage and prioritize the "easiest" next resources to implement using the direct reconciliation model.

## Features

- **SHA-based Analysis**: Compare specific versions of both `googleapis` and `k8s-config-connector`.
- **Capability Detection**: Identifies which GCP resources support `Create` and `Delete` operations based on gRPC service definitions.
- **Hierarchy Analysis**: Detects "leaf" resources whose parents are Projects, Folders, Organizations, or Locations.
- **Heuristic Matching**: Links GCP proto types to KCC CRDs by accounting for service aliases and naming conventions (e.g., stripping "Compute" prefix from `ComputeInstance`).
- **Prioritization**: Recommends the "Next K" resources to implement based on ease of management.

## Usage

```bash
python3 hack/tools/greenfield/calculate_coverage.py <googleapis_sha> <kcc_sha> [k]
```

### Parameters

- `<googleapis_sha>`: The git SHA, branch, or tag in the `googleapis/googleapis` repo (e.g., `master`).
- `<kcc_sha>`: The git SHA in the `k8s-config-connector` repo. Use `LOCAL` to compare against your current working directory.
- `[k]`: (Optional) The number of "easiest" resources to list. Defaults to 10.

### Example

Compare the latest protos against your local KCC state:
```bash
python3 hack/tools/greenfield/calculate_coverage.py master LOCAL 20
```

## Metric Definitions

- **Total GCP Resources**: Every unique `google.api.resource` or `google.api.resource_definition` found in the protos.
- **Manageable Missing**: Resources that have a `Create` or `Upsert` RPC but are not in KCC.
- **Full Lifecycle Missing**: Resources that have both a creation RPC and a termination RPC (`Delete`, `Finish`, `Abort`, etc.).
- **Leaf (Easy) Missing**: Resources that support the full lifecycle and have a flat parentage (Project, Folder, Org, or Location). These are generally the most straightforward to implement as direct resources.

## Unified Resource Policy

When implementing missing resources identified by this tool, follow the **Unified Direct Pattern**:

1.  **Single logical Kind**: If a GCP resource supports multiple hierarchies (e.g., Global and Regional) but shares the same logical name, do **not** create separate KRM Kinds.
2.  **Extended Existing Kinds**: If a resource identified as "missing" is actually a hierarchical variant of an existing KCC Kind (e.g., Regional Secret Manager Secret), the correct action is to **extend the existing Kind** with a `location` field rather than creating a new one.
3.  **Location-Aware Identity**: Use `gcpurls.Template` to implement identity logic that branches based on the `location` field.

## Implementation Workflow

This tool is integrated into the project's agentic workflows. The chore file `.agents/greenfield-direct-new-resource-types.md` uses this script to automatically identify and create implementation tasks for missing resources.
