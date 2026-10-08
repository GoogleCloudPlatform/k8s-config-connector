# CCInsightsAnalysisRule Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Stability Label**:
   - Added `cnrm.cloud.google.com/stability-level: alpha` as a metadata label to the CRD schema.

2. **API & Types Scaffolding**:
   - Mapped `CCInsightsAnalysisRule:AnalysisRule` from `google.cloud.contactcenterinsights.v1`.
   - Included all scalar primitives as pointers (`DisplayName *string`, `ConversationFilter *string`, `AnalysisPercentage *float64`, `Active *bool`).
   - Kept `Location *string` as required pointer field in `CCInsightsAnalysisRuleSpec`.
   - Leveraged existing nested structs (`AnnotatorSelector`) in `types.generated.go`.

3. **Identity & Reference**:
   - Implemented identity under `apis/contactcenterinsights/v1alpha1/ccinsightsanalysisrule_identity.go`.
   - The GCP resource URL format is `projects/{project}/locations/{location}/analysisRules/{analysisRule}`.
   - Added registration exception in `pkg/gcpurls/registry_test.go` for CAI alignment.
   - Implemented standard `CCInsightsAnalysisRuleRef` in `apis/contactcenterinsights/v1alpha1/ccinsightsanalysisrule_reference.go` delegating `Normalize` to `refs.Normalize`.
