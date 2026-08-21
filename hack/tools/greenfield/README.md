# KCC Resource Coverage Analysis Tool

This tool calculates the coverage of Google Cloud Platform (GCP) resources in Kubernetes Config Connector (KCC) relative to the published definitions in the `googleapis/googleapis` repository.

It is used to identify gaps in resource coverage and prioritize the "easiest" next resources to implement using the direct reconciliation model.

## Features

- **SHA-based Analysis**: Compare specific versions of both `googleapis` and `k8s-config-connector`.
- **Capability Detection**: Identifies which GCP resources support standard CRUD operations (Create, Read, Update, Delete) based on gRPC service definitions.
- **Hierarchy Analysis**: Detects "leaf" resources whose parents are Projects, Folders, Organizations, or Locations.
- **Heuristic Matching**: Links GCP proto types to KCC CRDs by accounting for service aliases and naming conventions (e.g., stripping "Compute" prefix from `ComputeInstance`).
- **Prioritization**: Recommends the Top K missing manageable resources to implement, sorted by ease of implementation (Leaf, Next Layer, etc.), displaying their available and missing CRUD operations.

## Usage

```bash
python3 hack/tools/greenfield/calculate_coverage.py <googleapis_sha> <kcc_sha> [k] [--update-gap] [--resources-list]
```

### Parameters

- `<googleapis_sha>`: The git SHA, branch, or tag in the `googleapis/googleapis` repo. The pinned commit for standard gap analysis is [3f9c9d7](https://github.com/googleapis/googleapis/commit/3f9c9d72cb20768ca4ac9f12030faaf43b13c231) (or use `master` for the bleeding edge).
- `<kcc_sha>`: The git SHA in the `k8s-config-connector` repo. Use `LOCAL` to compare against your current working directory.
- `[k]`: (Optional) The number of top manageable resources to list. Defaults to 10.
- `--update-gap`: (Optional) If provided, updates the `gap_analysis.txt` file with the coverage snapshot.
- `--resources-list`: (Optional) If provided, outputs a full structured list of all resources with their layer and CRUD status to `resources_list.json`.

### Example

Compare the pinned googleapis protos against your local KCC state with CRUD analysis and save both reports:
```bash
python3 hack/tools/greenfield/calculate_coverage.py 3f9c9d7 LOCAL 20 --update-gap --resources-list
```

## Metric Definitions

- **Total GCP Resources**: Every unique `google.api.resource` or `google.api.resource_definition` found in the protos.
  - **Manageable**: Resources that have at least one CRUD function (`CREATE`, `READ`, `UPDATE`, or `DELETE`) but are not in KCC.
    - **Adoptable**: Resources that have no `CREATE` function, but have at least one other CRUD function.
    - **Partially Manageable(Creatable)**: Resources that have the `CREATE` function but lack the `DELETE` function.
    - **Fully Manageable**: Resources that have both `CREATE` and `DELETE` functions.
      - **Leaf (Easy)**: Resources that are fully manageable and have flat parentage (Project, Folder, Org, or Location). These are generally the most straightforward to implement as direct resources.
      - **Next Layer**: Resources that are fully manageable and have one dependency.
  - **Unmanageable**: Resources that have no CRUD function.
## Unified Resource Policy

When implementing missing resources identified by this tool, follow the **Unified Direct Pattern**:

1.  **Single logical Kind**: If a GCP resource supports multiple hierarchies (e.g., Global and Regional) but shares the same logical name, do **not** create separate KRM Kinds.
2.  **Extended Existing Kinds**: If a resource identified as "missing" is actually a hierarchical variant of an existing KCC Kind (e.g., Regional Secret Manager Secret), the correct action is to **extend the existing Kind** with a `location` field rather than creating a new one.
3.  **Location-Aware Identity**: Use `gcpurls.Template` to implement identity logic that branches based on the `location` field.

## Implementation Workflow

This tool is integrated into the project's agentic workflows. The chore file `.agents/greenfield-direct-new-resource-types.md` uses this script to automatically identify and create implementation tasks for missing resources.
