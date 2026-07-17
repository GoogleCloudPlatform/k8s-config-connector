# DataprocCluster Migration Journal

## Overview
This journal tracks the implementation details and learnings from migrating the `DataprocCluster` resource to direct KRM types (`v1beta1`).

## Learnings & Observations

### 1. Struct and Type Name Collisions
In GCP Dataproc Protos, there is a top-level `ClusterConfig` message. However, inside virtual cluster / GKE node pool configuration, GKE's node pool config also contains a nested structure named `ClusterConfig` (representing GKE's GkeNodeConfig details in proto).
In the original legacy generator, this type name collision resulted in the GKE-specific node config overwriting the top-level `ClusterConfig` definition, causing the top-level `spec.config` to lose all its real fields (such as `masterConfig`, `workerConfig`, etc.) in the generated client-go types.
To resolve this:
- We renamed the GKE node configuration Go struct to `ClusterGkeNodeConfig`.
- We updated GKE node pool's referencing `Config` field to point to `ClusterGkeNodeConfig` instead of `ClusterConfig`.
- We declared the top-level `ClusterConfig` struct with its actual correct fields (e.g., `masterConfig`, `workerConfig`, `softwareConfig`, references, etc.).
This successfully preserved the full CRD schema properties of both fields while resolving the type collision cleanly.

### 2. Date-Time Format Overrides
For timestamp-based fields (e.g., `idleStartTime`, `autoDeleteTime`, `stateStartTime`), the `controller-gen` tool generates standard string schemas. To maintain strict schema compatibility with the original CRD schemas (which explicitly specify `format: date-time`), we utilized `// +kubebuilder:validation:Format="date-time"` annotations on the fields.
