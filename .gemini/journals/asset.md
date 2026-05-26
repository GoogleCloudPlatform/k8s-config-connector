
### [2026-05-26] CloudAssetSavedQuery Identity Implementation
- **Context**: Implementing `CloudAssetSavedQuery` KRM types and IdentityV2 (PR step 1).
- **Problem**: `SavedQuery` uses multiple parent hierarchies (projects, folders, organizations). Furthermore, `generate-types` conflicts with existing `AssetSavedQuery` types in `apis/asset/v1beta1`.
- **Solution**: Implemented `IdentityV2` using multiple `gcpurls.Template` templates in `FromExternal` and `String` methods, switching based on populated `Parent` fields. Renamed types in `apis/cloudasset/v1alpha1/savedquery_types.go` from `IAMPolicyAnalysisQuery` to `CloudAssetIAMPolicyAnalysisQuery` to prevent mapper redeclaration errors in `pkg/controller/direct/asset/mapper.generated.go`.
- **Impact**: When implementing `cloudasset` resources alongside `asset` resources, care must be taken to namespace struct names inside the `_types.go` files to prevent global mapper conflicts within the same Go package.
