# KCC Resource Name vs. GCP Resource Name

This snippet explains how Config Connector (KCC) handles resource name and the relationship between Kubernetes identifiers and Google Cloud Platform (GCP) identifiers.

## Conceptual Mapping

In a typical KCC resource, identity is managed through three primary fields:

1.  **`metadata.name`**: The standard Kubernetes resource name. By default, KCC uses this as the name of the resource on GCP.
2.  **`spec.resourceID`**: An optional field that allows users to specify a GCP resource name that differs from `metadata.name`. This is useful if the desired GCP name contains characters invalid in Kubernetes or if the user wants to decouple the two names. For more details on `resourceID`, see [resourceid.md](./resourceid.md).
3.  **`status.externalRef`**: The canonical, fully-qualified GCP resource name (e.g., `projects/my-project/zones/us-central1-a/instances/my-instance`). This is populated by the controller after the resource is successfully reconciled.

## Why the Proto `name` Field is Excluded

When generating KRM types from GCP protocol buffers, the `name` field found in the proto message is deliberately omitted from the KRM `Spec`. There are several reasons for this:

### 1. Kubernetes Idioms
Kubernetes uses `metadata.name` as the primary identifier. Including a redundant `spec.name` field would violate KRM conventions and create ambiguity regarding which field defines the resource's identity.

### 2. Structured vs. Unstructured Identity
GCP proto `name` fields are sometimes a "relative resource name" — full path string like `projects/my-project/zones/us-central1-a/instances/my-instance`. 

KCC decomposes this unstructured string into a declarative, structured `Spec`:
*   **`projectRef`**: Specifies the project.
*   **`location`**: Specifies the region or zone.
*   **`resourceID`**: Specifies the final resource identifier.

Asking users to provide the full path string in a `spec.name` field would be error-prone and less Kubernetes-native than providing structured references.

### 3. Server-Generated and Output-Only Fields
In many GCP APIs, the `name` field is marked as `OUTPUT_ONLY`. This is because the full path is often only known or finalized after the server processes the creation request. 

The KCC type generator automatically excludes `OUTPUT_ONLY` fields from the `Spec` to ensure that only user-configurable fields are present. These values are instead mapped to `status.observedState` or `status.externalRef`.
