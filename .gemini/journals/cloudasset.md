### [2026-07-19] Scaffolding Greenfield Direct Types for CloudAssetSavedQuery
- **Context**: Implementing Greenfield direct KRM types, CRD, and IdentityV2 for `CloudAssetSavedQuery` (under the `cloudasset` API group) in `apis/cloudasset/v1alpha1/`.
- **Problem**: Cloud Asset API's URL format includes a version prefix `v1/` (`projects/{project}/savedQueries/{savedQuery}`). Standard relative URL parsing with `gcpurls.Template` failed when handling fully-qualified GCP URLs because the presence of the `v1/` prefix shifted the segment matching offset.
- **Solution**: Implemented a dual-template pattern for the `CloudAssetSavedQueryIdentityFormat`:
  1. `CloudAssetSavedQueryIdentityFormat` with the host and `v1/projects/{project}/savedQueries/{savedQuery}`.
  2. `CloudAssetSavedQueryIdentityFormatRelative` with the relative path starting from `projects/{project}/savedQueries/{savedQuery}`.
  In the `FromExternal` parser, we attempt the first (absolute/versioned) format and fall back to the relative format if the absolute one does not match. This ensures robust parsing for both forms.
- **Impact**: Other resources with version prefixes in their GCP URLs can safely employ this dual-template pattern to avoid path-matching panics and successfully parse both fully-qualified and relative resource identifiers.
