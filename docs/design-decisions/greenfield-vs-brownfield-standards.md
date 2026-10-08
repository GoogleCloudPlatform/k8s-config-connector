# Greenfield vs. Brownfield Standards in Direct Controllers

## Context

This document captures design decisions for **brownfield resources** (resources migrating from legacy Terraform or DCL controllers to Direct controllers) that differ from the standards established for **greenfield resources** (brand-new resources implemented directly with the Google Cloud SDK).

When migrating brownfield resources to direct controllers, strict backward compatibility must be maintained for existing user configurations, manifests, and deployment patterns.

## 1. Resolving Project ID

### Brownfield Behavior

In brownfield resources (i.e. resources supported before we introduce the direct controller), project ID resolution follows a 4-tier hierarchy:

1. **Tier 1: `spec.projectRef`**: Explicit project reference in the resource spec (if supported).
2. **Tier 2:  Resource-level Annotation**: `cnrm.cloud.google.com/project-id` on `metadata.annotations`.
3. **Tier 3:  Namespace-level Annotation**: `cnrm.cloud.google.com/project-id` on the Kubernetes `Namespace` object.
4. **Tier 4: Namespace Name**: The Kubernetes namespace name (assuming 1:1 namespace-to-GCP-project mapping).

#### Implementation Details

The [`container-annotation-handler` mutating admission webhook](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/v1.156.0/pkg/webhook/container_annotation_handler.go) defaults container annotations from the `Namespace` object (or the namespace name) on resource creation so that `cnrm.cloud.google.com/project-id` is always populated on the brownfield resource's metadata before reconciliation.

In direct controllers, the [`refs.ResolveProjectID` utility function](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/v1.156.0/apis/refs/v1beta1/projectref.go#L278) resolves the project ID by first checking `spec.projectRef`, and if unset, falling back to the `cnrm.cloud.google.com/project-id` annotation on the resource object.

### Special Handling in Direct Controllers

Use the `refs.ResolveProjectID` utility function when `spec.projectRef` is not supported in the CRD schema.

### Differences Comparing with Greenfield

- **Implementation**:
  - **Brownfield resource with `spec.projectRef`**: **No difference** comparing with greenfield. Both resolve the project identity directly from the `spec.projectRef` field.
  - **Brownfield resource without `spec.projectRef` and requires a parent Project**: **Different** comparing with greenfield. Direct controllers must call `refs.ResolveProjectID` to resolve the project ID from the resource's `cnrm.cloud.google.com/project-id` annotation, whereas greenfield resources always require explicit parent reference fields in `spec`.
- **Behavior**: Brownfield resources can resolve the project ID from a namespace-level annotation or the Kubernetes namespace name, whereas greenfield resources cannot.
  - **How it works**: The mutating admission webhook (`pkg/webhook/container_annotation_handler.go`) resolves and defaults the annotation based on whether the resource has a `ServiceMapping` (Terraform) or `ServiceMetadata` (DCL) (migrated brownfield resources still retain their `ServiceMapping`/`ServiceMetadata`).

> [!WARNING]
> **Future `ServiceMapping` / `ServiceMetadata` Deletion**: Because the admission webhook relies on the presence of `ServiceMapping`/`ServiceMetadata` to intercept resources, deleting these legacy mapping files in the future will cause silent breaking changes for brownfield resources unless a direct-native container metadata registry is introduced first. It's tracked at https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/12817.
