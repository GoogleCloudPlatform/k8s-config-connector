# ComputeRouterNAT Direct Migration Journal

During the direct migration validation for `ComputeRouterNAT`, we encountered and resolved takeover diff issues during direct takeover (Phase 3) and direct re-reconciliation (Phase 4).

## Observations & Root Causes

### 1. Unspecified Optional Fields Diff
- **Issue:** Optional fields in the KRM spec that are omitted (such as `type` or `endpointTypes`) are parsed as `nil` in the desired Proto representation. However, GCP returns default values (like `PUBLIC` or `["ENDPOINT_TYPE_VM"]`) in `maskedActual`. The top-level comparison `DiffForTopLevelFields` flagged these as differences, resulting in unexpected `PATCH` requests during Phase 3/4 that stripped away those unspecified optional fields.
- **Resolution:** Leveraged `common.MergeUnsetFields` to recursively merge any unset/nil fields from the actual GCP state (`maskedActual`) into the desired state (`clonedDesired`) before defaults are populated and compared.

### 2. Absolute vs. Relative URL Mismatches and List Ordering Non-Determinism
- **Issue:** Reference fields to other compute resources (such as `NatIps`, `DrainNatIps`, `SourceNatActiveRanges`, and `SourceNatDrainRanges` inside rule actions) returned fully-qualified absolute URLs (e.g. `https://www.googleapis.com/...`) from GCP, but were resolved to shorter relative paths (e.g. `projects/...`) in the desired Proto representation. In addition, the list of IPs returned by GCP did not guarantee deterministic ordering, triggering spurious diffs.
- **Resolution:** Added a `normalizeRouterNat` helper function in `computerouternat_controller.go` to normalize all compute reference links (stripping absolute URL host prefixes) and sort list/slice fields alphabetically. This ensured that different URL styles and element orderings do not cause false-positive diffs during direct takeover.
