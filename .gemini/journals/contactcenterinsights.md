# Contact Center Insights Service Journal

### [2026-06-05] CCInsightsView Greenfield Types Scaffolding
- **Context**: Implementing Greenfield direct types, CRD, and IdentityV2 for `CCInsightsView` resource (mapped to `google.cloud.contactcenterinsights.v1.View`).
- **Problem**: Greenfield types need to strictly follow the KRM layout, specifically:
  1. Scalar primitive fields (including `Location`) must be pointer types, e.g., `*string`, regardless of optionality/required status.
  2. The custom `generate-types` script generates a template containing general fields, but resource-specific spec fields (like `displayName` and `value`) must be manually declared.
  3. The `View` resource in Google APIs is simple but has a standard project/location/views format, which maps perfectly to standard URL templates.
- **Solution**:
  1. Manually enhanced the `view_types.go` file to include `DisplayName *string` and `Value *string` in `CCInsightsViewSpec`.
  2. Defined `Location *string` as a pointer.
  3. Implemented `CCInsightsViewIdentity` with format `projects/{project}/locations/{location}/views/{view}` and the respective `IdentityV2` interface methods, including cross-checking `Status.ExternalRef`.
  4. Implemented `CCInsightsViewRef` and registered it for resource referencing.
  5. Wrote comprehensive unit tests to validate identity and reference logic.
- **Impact**: Establishes a standard, fully validated direct type foundation for Contact Center Insights (CCInsightsView), passing all presubmits and client-go code generation cleanly.

### [2026-06-05] CCInsightsPhraseMatcher Greenfield Types and Identity
- **Context**: Greenfield types, CRD, and IdentityV2 implementation for `CCInsightsPhraseMatcher` (GCP `PhraseMatcher` under `contactcenterinsights.googleapis.com`).
- **Problem**: The proto-to-KRM generator comments out nested structs as "unreachable types" initially because `phrasematcher_types.go` is stubbed out. Also, we had to ensure scalar types are represented as pointers to satisfy KCC's strict pointer rule.
- **Solution**: We populated the nested structs (such as `PhraseMatchRuleGroup` and `PhraseMatchRule`) explicitly in `phrasematcher_types.go` and defined `Location` in `CCInsightsPhraseMatcherSpec` directly as `*string`. Running `generate.sh` then resolved all unreachable types cleanly.
- **Impact**: All future developers working on `contactcenterinsights` can leverage the scaffolded types, CRDs, and identity structure, knowing that the compiler and validators are fully satisfied with the pointer fields.
