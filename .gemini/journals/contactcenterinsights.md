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

### [2026-06-05] Initial Scaffolding and Identity for CCInsightsIssueModel
- **Context**: Greenfield implementation of CCInsightsIssueModel types, CRD, and IdentityV2 under `contactcenterinsights.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: When `generate-types` first ran, the main `IssueModel` and nested config/observedState sub-types were marked as unreachable and commented out because they were not yet referenced in the spec and status fields of `CCInsightsIssueModelSpec` and `CCInsightsIssueModelObservedState`.
- **Solution**: Explicitly defined the spec fields referencing `IssueModel_InputDataConfig` and observedState fields referencing the nested observedState sub-types inside `issuemodel_types.go`, then re-ran `generate.sh`. The generator successfully detected them as reachable and fully compiled them into `types.generated.go`.
- **Impact**: Provides proper type scaffolding, deepcopy, CRD, and IdentityV2 for CCInsightsIssueModel without unreachable/commented-out required types.

### [2026-06-05] CCInsightsConversation Greenfield Types Scaffolding
- **Context**: Greenfield implementation of CCInsightsConversation (Step 1) as part of Issue #9260.
- **Problem**: CCInsightsConversation is a new service `contactcenterinsights.cnrm.cloud.google.com/v1alpha1`. We needed to build the Google APIs pb files, scaffold the types via `generate-types`, define custom properties under Spec and ObservedState, and set up IdentityV2 and External References.
- **Solution**: 
  - Scaffolded package structure in `apis/contactcenterinsights/v1alpha1/`.
  - Configured `generate.sh` to compile Google API proto definitions from `resources.proto` using `generate-proto.sh`, then ran `generate-types` to build KRM fields and Go models.
  - Implemented `ccinsightsconversation_identity.go` (implementing `identity.IdentityV2` and `identity.Resource`) and `ccinsightsconversation_reference.go` (implementing `refs.Ref`), ensuring standard camelCase mapping (e.g., `{conversation}` to `Conversation` in the struct fields) to prevent initialization panic in `gcpurls.Template`.
  - Added unit testing to verify the URL-parsing behavior of the identity structure.
- **Impact**: Provides a robust foundation for the future controller and mapper implementation of CCInsightsConversation.

### [2026-07-06] CCInsightsPhraseMatcher Controller Implementation
- **Context**: Implementing Phase 2 (Controller and E2E Tests) for CCInsightsPhraseMatcher.
- **Problem**: `CreatePhraseMatcherRequest` in `google.cloud.contactcenterinsights.v1` does not have a separate `phrase_matcher_id` field, which is unusual for GCP APIs that allow user-provided IDs.
- **Solution**: Set the full resource name (including the user-provided ID) on the `PhraseMatcher.Name` field in the `Create` request.
- **Impact**: Other resources in `contactcenterinsights` (like `View`) likely follow this same pattern.

### [2026-07-02] CCInsightsAnalysisRule Greenfield Types Scaffolding
- **Context**: Implementing Greenfield direct types, CRD, and IdentityV2 for `CCInsightsAnalysisRule` resource (mapped to `google.cloud.contactcenterinsights.v1.AnalysisRule`).
- **Problem**: CCInsightsAnalysisRule's `AnnotatorSelector` field has nested reference lists to other resources (`phraseMatchers` and `issueModels`). If left as `[]string`, they would bypass KCC's strict reference guidelines.
- **Solution**: 
  1. Manually defined `AnnotatorSelector` within `ccinsightsanalysisrule_types.go`, overriding the generated version.
  2. Declared `PhraseMatchers []CCInsightsPhraseMatcherRef` and `IssueModels []CCInsightsIssueModelRef` as proper, strong KCC references.
  3. Created `ccinsightsphrasematcher_reference.go` to expose `CCInsightsPhraseMatcherRef` as a canonical reference type.
  4. Implemented `ccinsightsanalysisrule_identity.go` (implementing `identity.IdentityV2` and `identity.Resource`) and `ccinsightsanalysisrule_reference.go` (implementing `refs.Ref`).
  5. Wrote unit tests for identity and reference validation and confirmed they compile and pass perfectly.
- **Impact**: Establishes a completely compliant, strongly-typed direct type foundation for `CCInsightsAnalysisRule` with robust references and URL validation.
