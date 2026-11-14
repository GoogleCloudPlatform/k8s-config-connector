### [2026-07-15] Normalization of modern direct resources
- **Observation**: StorageInsightsDatasetConfig represents a modern direct resource with `status.externalRef` support.
- **Rule**: Standardized its reference type `Normalize` implementation in `storageinsightsdatasetconfig_reference.go` to delegate directly to `refs.Normalize` instead of using the legacy fallback wrapper `refs.NormalizeWithFallback`. This guarantees that dependency validation behaves correctly and avoids premature identity resolution, keeping in line with the direct Greenfield guidelines.
