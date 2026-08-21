### [2026-05-20] BigQueryAnalyticsHubListing Identity & Refs
- **Context**: Implementing IdentityV2 and refs.Ref for BigQueryAnalyticsHubListing (v1beta1).
- **Problem**: The resource has a parent resource (DataExchange) that must be part of its identity. The previous implementation used shortened filenames (listing_identity.go) and custom methods (ParseDataExchangeIdentity).
- **Solution**: 
    1. Renamed files to use the full kind name (bigqueryanalyticshublisting_identity.go) for consistency with the canonical ArtifactRegistry example.
    2. Used `ParseExternalToIdentity()` on the parent reference to resolve its identity cleanly in `getIdentityFromBigQueryAnalyticsHubListingSpec`.
    3. Added unit tests for both identity (parsing and string conversion) and reference (GVK validation).
- **Impact**: Better adherence to standards makes the codebase easier to navigate and maintain. Standardized filenames help tools and humans find the right files quickly.
