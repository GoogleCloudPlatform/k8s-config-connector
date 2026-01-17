# Development Decisions

This document tracks the current status of decisions having to do with developing the KCC project.
If a decision needs to be made or changed, please file an issue for that decision to be made.

| Short Name | Brief Description | Date of Decision | Status | Link to More Details |
| :--- | :--- | :--- | :--- | :--- |
| Labels in Direct Resources | Use Kubernetes `metadata.labels` as the single source of truth for GCP labels. Remove `labels` from `Spec`. | 2025-10-09 | Accepted | [Link](../ai/handle-labels-for-direct-resource.md) |
| Server Generated ID | Use `status.externalRef` for server-generated IDs. For migrated resources, continue writing to `spec.resourceID` for compatibility. | 2025-12-03 | Accepted | [Link](../ai/server-generated-id.md) |
| Graceful Orphan Deletion | Allow deletion of resources when parent dependencies are missing by using `status.externalRef` instead of resolving from `spec`. | 2025-10-03 | Proposed | [Link](../designs/graceful-orphan-deletion.md) |
| Ignore Unspecified Fields | Use annotation `cnrm.cloud.google.com/default-to-gcp-fields` to preserve GCP values for unspecified fields in `spec`. | 2025-09-09 | Proposed | [Link](../designs/ignore-unspecified-fields-direct-controllers.md) |
| Stateful Reconciliation | Use `status.lastModifiedCookie` to store hashes of applied spec and observed GCP state to detect changes and drift efficiently. | 2025-09-30 | Proposed | [Link](../designs/stateful-reconciliation-with-cookie.md) |
