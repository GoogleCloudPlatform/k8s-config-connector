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

### [2026-09-29] CCInsightsAnalysisRule Greenfield Types Scaffolding
- **Context**: Greenfield implementation of CCInsightsAnalysisRule (Step 1: types, CRD, and IdentityV2) under `contactcenterinsights.cnrm.cloud.google.com/v1alpha1`.
- **Solution**:
  - Added `CCInsightsAnalysisRule:AnalysisRule` to `apis/contactcenterinsights/generate.sh`.
  - Defined `CCInsightsAnalysisRuleSpec` and `CCInsightsAnalysisRuleObservedState` in `apis/contactcenterinsights/v1alpha1/ccinsightsanalysisrule_types.go`, using pointer fields for all primitives (including `Location *string`).
  - Implemented `IdentityV2` in `ccinsightsanalysisrule_identity.go` with template `projects/{project}/locations/{location}/analysisRules/{analysisRule}` and added URL format exception to `pkg/gcpurls/registry_test.go`.
  - Implemented `CCInsightsAnalysisRuleRef` in `ccinsightsanalysisrule_reference.go` delegating `Normalize` to `refs.Normalize`.
  - Added comprehensive identity unit tests with `cmp.Diff`.
- **Impact**: Provides valid direct KRM types, CRD, and IdentityV2 for CCInsightsAnalysisRule.

### [2026-09-30] CCInsightsAnalysisRule Direct Controller and E2E Fixtures
- **Context**: Implementing Greenfield direct controller, fuzzer, and recording/verifying E2E fixtures for `CCInsightsAnalysisRule`.
- **Problem**:
  1. `CreateAnalysisRuleRequest` takes parent and `AnalysisRule` proto with server-generated ID returned on `created.Name`.
  2. GCP API requires `display_name` to be non-empty upon creation, and if `active: true`, at least one annotator must be enabled in `annotatorSelector`.
  3. Server-assigned numeric ID required normalization in `tests/e2e/normalize.go`, `tests/e2e/replacements.go`, and `pkg/cais/caistesting/testing.go`.
- **Solution**:
  1. Implemented isolated controller under `pkg/controller/direct/contactcenterinsights/ccinsightsanalysisrule/` using official GAPIC REST client (`contactcenterinsights.NewRESTClient`).
  2. Defaulted `DisplayName` to resource name in `AdapterForObject` if unspecified, and populated default in `compare`.
  3. Implemented round-trip fuzzer registered via `fuzztesting.RegisterKRMFuzzer`.
  4. Added `CCInsightsAnalysisRule` to `tests/e2e/normalize.go` (`${analysisRuleId}`), `replacements.go`, and `caistesting/testing.go`.
  5. Recorded minimal and maximal fixtures against real GCP (`RECORD_AUDIT_PROBE=1 ./hack/record-gcp`), and verified re-reconciliation and audit log generation.
- **Impact**: Fully functional direct controller and E2E test coverage for `CCInsightsAnalysisRule`.

### [2026-10-04] CCInsightsAnalysisRule MockGCP and Alignment verification
- **Context**: Aligning MockGCP logs with RealGCP output for `CCInsightsAnalysisRule` (`contactcenterinsights.cnrm.cloud.google.com/v1alpha1`).
- **Findings & Verification**:
  1. Verified MockGCP implementation for `CCInsightsAnalysisRule` in `mockgcp/mockcontactcenterinsights/ccinsightsanalysisrule.go` handling `CreateAnalysisRule`, `GetAnalysisRule`, `UpdateAnalysisRule`, `DeleteAnalysisRule`, and `ListAnalysisRules`.
  2. Verified 404 error message formatting matching real GCP (`"No AnalysisRule found for project: \`%d\` and AnalysisRule Id: \`%s\`."`).
  3. Verified ID generation, timestamp updates, field mask updates, and URL normalizations for `${analysisRuleId}` in `mockgcp/mockcontactcenterinsights/normalize.go`.
  4. Verified both `ccinsightsanalysisrule-minimal` and `ccinsightsanalysisrule-maximal` test fixtures pass `hack/compare-mock "fixtures/ccinsightsanalysisrule"` and `TestGoldenLogAlignment` with 0 diffs.

### [2026-11-14] CCInsightsQAScorecardRevision Greenfield Prepopulate (Step 1a)
- **Context**: Mechanically prepopulating direct KRM types, IdentityV2, reference, CRD, and `judgement_queue.yaml` for `CCInsightsQAScorecardRevision` (`google.cloud.contactcenterinsights.v1.QaScorecardRevision`).
- **Implementation**:
  1. Configured `apis/contactcenterinsights/generate.sh` with `--prepopulate-resource` for `CCInsightsQAScorecardRevision:QaScorecardRevision` in `generate-types` and `generate-identity`.
  2. Generated `apis/contactcenterinsights/v1alpha1/ccinsightsqascorecardrevision_types.go`, identity, reference, deepcopy, mapper, and CRD.
  3. Added `CCInsightsQAScorecardRevision` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  4. Added untriaged findings for `CCInsightsQAScorecardRevision` into `apis/contactcenterinsights/judgement_queue.yaml`.
  5. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` via golden test recording.
- **Verification**: Verified `apis/contactcenterinsights/...`, `pkg/controller/direct/contactcenterinsights/...`, and `tests/apichecks/...` all pass with zero errors.

### [2026-11-14] CCInsightsQAScorecardRevision Greenfield Types Judgement Triage (Step 1b)
- **Context**: Triaging `judgement_queue.yaml` entries and finalizing direct KRM types and identity for `CCInsightsQAScorecardRevision` (`google.cloud.contactcenterinsights.v1.QaScorecardRevision`).
- **Parent & Identity Resolution**:
  - `CCInsightsQAScorecardRevision` is a child resource under `qaScorecards/{qa_scorecard}`.
  - Replaced `projectRef` and `location` with `spec.parentRef *CCInsightsQAScorecardRef`, marked `// +required` and immutable via `+kubebuilder:validation:XValidation:rule="self == oldSelf"`.
  - Re-generated identity via `generate-identity`, which resolves `ParentRef` and builds identity from the parent's external reference and `resourceID`.
- **Nested Snapshot Struct Customization**:
  - `snapshot` represents a `QaScorecard`. Moved `QaScorecard` and `QaScorecardObservedState` into `ccinsightsqascorecardrevision_types.go` with `+kcc:proto=` annotations.
  - Removed `name` (resource identifier) from `spec.snapshot` as it is output-only/server-set from the parent scorecard, and preserved it under `status.observedState.snapshot.name`.
- **Queue Triage & Verification**:
  - Resolved all 5 `judgement_queue.yaml` entries (`untriaged-bulk-generation`, `identity-root-unknown`, `location-or-parent-ref`, `parent-ref-guessed`, and `possible-reference-by-description`).
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` via golden recording.
  - Verified `generate.sh` produces zero diff, and all tests in `apis/contactcenterinsights/...`, `pkg/controller/direct/contactcenterinsights/...`, and `tests/apichecks/...` pass cleanly.

### [2026-11-14] CCInsightsQAScorecardRevision Greenfield Direct Controller & E2E Fixtures (Step 2)
- **Context**: Implementing Phase 2 (direct controller, fuzzer, and E2E fixtures recorded against RealGCP) for `CCInsightsQAScorecardRevision`.
- **GCP API Nuances & Architecture**:
  1. **Scorecard Lifecycle & Revision Mutability**: `QaScorecardRevision` is completely immutable in GCP (there is no `UpdateQaScorecardRevision` RPC). When a `QaScorecard` is created, GCP automatically creates an initial revision in state `EDITABLE`.
  2. **Precondition for Creating Revisions**: A scorecard can only have one `EDITABLE` revision at a time. Attempting to call `CreateQaScorecardRevision` while another revision is in state `EDITABLE` returns `400 FAILED_PRECONDITION: Precondition check failed.`. To create a new revision, any in-flight `EDITABLE` revision on the parent scorecard must first be tuned (`TuneQaScorecardRevision`) until `State == READY` and deployed (`DeployQaScorecardRevision`). The controller's `Create` method proactively checks for and prepares (`prepareParentScorecard`) any in-flight initial `EDITABLE` revision.
  3. **ID Length Limit (Discrepancy with Proto)**: While the protobuf documentation for `CreateQaScorecardRevisionRequest` states `qaScorecardRevisionId` allows 4-64 characters matching `^[a-z0-9-]{4,64}$`, real GCP rejects IDs longer than 32 characters with `400 INVALID_ARGUMENT: Request contains an invalid argument.`. The fixture metadata names use `rev-minimal-${uniqueId}` and `rev-maximal-${uniqueId}` (<= 32 characters).
  4. **Required Snapshot Field**: GCP requires `snapshot.displayName` to be populated when creating a revision; omitting it returns `400 INVALID_ARGUMENT: Request contains an invalid argument.`.
  5. **Protobuf Validation (`protovalidate.go`)**:
     - Custom verbs (e.g. `:tuneQaScorecardRevision`, `:deploy`) were previously stripped in `expandHTTPPathTemplate`, causing different RPCs sharing the same path template to collide. `matchHTTPRule` and `httpBindingRule` now preserve and match `customVerb`.
     - Real GCP returns `source` on `QaScorecard` snapshots which is absent in the pinned `.build/googleapis.pb` descriptor; `QaScorecardRevision` and `ListQaScorecardRevisionsResponse` were added to `discardUnknown = true`.
- **Verification**:
  - `FOCUS=CCInsightsQAScorecardRevision go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers` passes 200-seed round-trip.
  - Both `ccinsightsqascorecardrevision-minimal` and `ccinsightsqascorecardrevision-maximal` passed all 7/7 automated validation checks in `unified_test.go` (`hack/record-gcp`) against real GCP (`cnrm-barni-4`) with genuine `_http.log` and `_audit_probe.log` recorded.







