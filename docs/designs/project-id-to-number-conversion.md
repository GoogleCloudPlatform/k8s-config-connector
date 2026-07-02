# Project ID to Project Number Conversion

## Problem Statement

Certain GCP services, particularly within the Compute Engine and Vertex AI APIs, return project numbers in resource identifiers, URLs even if the user initially provided project IDs in their configuration; Or reference fields where KCC by default to resolve references within the project ID. 

Config Connector maintains the desired state as specified by the user. When project ID is provided but the GCP API returns a project number, KCC may detect a permanent diff during the reconciliation loop. This results in unnecessary API calls and can prevent the resource from ever reaching a stable "UpToDate" status.

## Design Decision

To eliminate these permanent diffs, KCC normalizes project identifiers to a consistent format. Depending on the specific requirements of the GCP API and the resource type, we employ one of two strategies:

1.  **Normalization to Project Number (Spec-side)**: If the API primarily returns project numbers or requires them for specific fields, we convert the user-provided project ID to its corresponding project number during the reference resolution or normalization phase. This ensures the desired state matches the actual state returned by the API.
2.  **Normalization to Project ID (Status/Export-side)**: If the user expects project IDs in their configuration, but the API returns project numbers (e.g., in resource links), we convert the project numbers back to project IDs during the `Find`, `Export`, or status update phase.

## Implementation Details

KCC utilizes the `ProjectMapper` utility (defined in `apis/common/projects/mapper.go`) to perform these conversions. 

### `ProjectMapper` and `ProjectCache`

The `ProjectMapper` leverages an in-memory, expiring `ProjectCache` to minimize redundant calls to the Cloud Resource Manager API (`resourcemanager.projects.get`). This ensures that the conversion is performant and does not significantly increase reconciliation latency.

### Common Implementation Patterns

*   **Reference Resolution (`Resolve...Refs`)**: For "direct" controllers, normalization is often performed in the `Resolve...Refs` function before the adapter is created. This ensures the adapter receives a "clean" desired state.
*   **Reference Methods**: Reference types (e.g., `ComputeNetworkRef`) can implement a `ConvertToProjectNumber` method to encapsulate the conversion logic for specific fields.
*   **Adapter `Find` Phase**: If the project number is required only for constructing the API request URL, the conversion may be performed within the `Find` method of the adapter.

## Existing Resources Using ID-to-Number Conversion

The following resources currently implement project ID to number conversion to avoid reconciliation diffs:

| Resource | Implementation Location | Logic Description |
| :--- | :--- | :--- |
| **ComputeReservation** | `pkg/controller/direct/compute/reservation_resolverefs.go` | Normalizes `Spec.ShareSettings.ProjectMap` keys and values to project numbers. |
| **ComputeFutureReservation** | `pkg/controller/direct/compute/futurereservation_resolverefs.go` | Normalizes `Spec.ShareSettings.ProjectMap` keys and values to project numbers. |
| **VertexAIMetadataStore** | `pkg/controller/direct/vertexai/vertexaimetadatastore_controller.go` | Converts project ID to number in `Find` to construct the correct resource URL. |
| **TagsTagKey** | `pkg/controller/direct/tags/tagstagkey_controller.go` | Normalizes the `Spec.Parent` field to use project numbers when referencing a project. |
| **ServiceNetworkingPeeredDNSDomain** | `pkg/controller/direct/servicenetworking/servicenetworkingpeereddnsdomain_controller.go` | Normalizes `Spec.NetworkRef` to use project numbers. |
| **CloudBuildWorkerPool** | `pkg/controller/direct/cloudbuild/workerpool_controller.go` | Normalizes `Spec.NetworkConfig.PeeredNetworkRef` to use project numbers. |
| **SecureSourceManagerRepository** | `pkg/controller/direct/securesourcemanager/securesourcemanagerrepository_controller.go` | Normalizes `Spec.InstanceRef` to use project numbers. |

## Guidance for New Resources

When implementing a new resource, check the http log of the fixtures tests to see if the GCP API returns project numbers for fields where users typically provide project IDs. 
If a permanent diff is observed in tests, implement normalization using `ProjectMapper`. 

Prefer performing the conversion during the reference resolution phase to ensure the KRM object remains stable throughout the reconciliation process.
