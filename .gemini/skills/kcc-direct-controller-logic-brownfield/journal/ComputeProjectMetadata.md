# ComputeProjectMetadata Direct Controller Logic Journal

## Overview
Implemented the direct controller logic for `ComputeProjectMetadata` (brownfield resource migrating from Terraform).

## Observations & Implementation Insights

### 1. Project-Level Common Instance Metadata Semantics
- `ComputeProjectMetadata` represents the project-level common instance metadata on Google Compute Engine (`Projects.Get` and `Projects.SetCommonInstanceMetadata`).
- In GCP, the metadata subresource exists as part of the project object.
- `Find` queries the project via `ProjectsClient.Get`. If found, `actual` is populated from `project.GetCommonInstanceMetadata()`.
- `Update` and `Create` both delegate to `save`, which compares desired items against actual items, sets the latest `Fingerprint`, and calls `SetCommonInstanceMetadata`.
- `Delete` clears common instance metadata by sending an empty metadata object with the current fingerprint, perfectly mirroring Terraform legacy controller behavior.

### 2. Behavioral Parity & Legacy Provider Quirk
- In the legacy Terraform provider, `expandComputeMetadata` had a minor slice preallocation bug where `make([]*compute.MetadataItems, len(m))` resulted in initial `null` elements in the `items` array sent over HTTP.
- The direct GAPIC REST controller sends clean, valid JSON items without spurious `null` values.
- Direct controller also eliminates an unnecessary trailing `GET` call during resource deletion.
- In `_migration_diffs.json`, migration from legacy Terraform controller to direct controller produces 0 diffs during takeover and steady-state re-reconciliation.
