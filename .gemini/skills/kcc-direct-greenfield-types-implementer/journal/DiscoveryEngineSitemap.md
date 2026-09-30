# DiscoveryEngineSitemap Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Schema Design**:
   - The GCP API for `Sitemap` (`google.cloud.discoveryengine.v1.Sitemap`) defines `uri`, `name` (output only), and `create_time` (output only).
   - In KRM `DiscoveryEngineSitemapSpec`, `uri` is mapped from proto `uri` (inside oneof `feed`), `Location`, `ProjectRef`, and `DataStoreRef` define the parent resource hierarchy, and `ResourceID` is optional for referencing pre-existing sitemaps.
   - Output-only field `name` maps to `status.externalRef` and `create_time` maps to `status.observedState.createTime`.

2. **Server-Generated ID**:
   - The sitemap ID is server-generated in GCP (`CreateSitemapRequest` takes `parent` and `sitemap` without a custom ID parameter).
   - `DiscoveryEngineSitemapIdentity` implements `identity.ServerGeneratedIdentity`.
   - `GetIdentity` defaults `specIdentity.Sitemap` from `status.externalRef` when `spec.resourceID` is unset, and verifies drift if `spec.resourceID` is specified.

3. **Identity Format**:
   - Identity format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}/siteSearchEngine/sitemaps/{sitemap}`.
   - `ParentString()` returns `projects/{project}/locations/{location}/collections/{collection}/dataStores/{dataStore}/siteSearchEngine`.
   - Registered URL template pattern exception added in `pkg/gcpurls/registry_test.go`.

4. **Reference Implementation**:
   - `DiscoveryEngineSitemapRef` implements `refs.Ref` and delegates normalization strictly to `refs.Normalize`.
