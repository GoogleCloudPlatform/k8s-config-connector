# Design Decision: ProjectRef in Migrating Resources

## Context
As Config Connector (KCC) migrates resources from legacy controllers (Terraform/DCL) to the "Direct" controller approach, we must maintain strict backward compatibility for `v1beta1` resources.

A common pattern for **new** Direct resources is to include a `projectRef` field in the `Spec` to define the parent GCP project. However, adding this field to **existing** resources during migration presents several challenges.

## Problem Statement
Adding `spec.projectRef` to a resource currently managed by a legacy controller (while keeping the same API version) introduces the following risks:

1. **Ignored Intent:** Legacy controllers (TF/DCL) are unaware of the `projectRef` field. If a user specifies it, the controller will ignore it and continue using the legacy resolution logic (namespace name or `cnrm.cloud.google.com/project-id` annotation), leading to "silent failures" where the resource is deployed to the wrong project.
2. **Controller Divergence:** If a user forces the Direct controller (via annotation), the resource might suddenly respect `projectRef` and move to a different project. This creates inconsistent behavior between controllers for the same API version.
3. **Schema Pollution:** Adding fields to a stable version (`v1beta1`) that only work under specific internal conditions is confusing for users and breaks the "Source of Truth" principle for the CRD schema.

## Decision
For all resources migrating from Terraform or DCL to the Direct controller within an existing API version:

1. **Do Not Add ProjectRef:** The `projectRef` field MUST NOT be added to the Go types or the CRD schema for existing versions.
2. **Use Annotation Resolution:** The Direct controller implementation MUST use `refsv1beta1.ResolveProjectFromAnnotation` to determine the parent project. This ensures parity with the legacy controllers.
3. **Support Location/Region Fields:** While `projectRef` is avoided, `location` or `region` fields can be included if they were already part of the schema or are required for the Direct controller's identity resolution, provided they follow existing patterns.

## Future Roadmap
The `projectRef` field should only be introduced after the TF Controller is no longer supported for the resource.

Once the direct controller is the sole engine for the resource, we could theoretically introduce projectRef as an optional, additive field.
We should either use the direct controller or enhance the existing webhooks to check `projectRef first`, and fall back to the `project-id` annotation if `projectRef` isn't provided.
However, we'd still have to support the annotation indefinitely to avoid breaking existing users.
