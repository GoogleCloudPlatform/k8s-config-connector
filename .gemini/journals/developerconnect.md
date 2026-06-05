# DeveloperConnect Service Journal

### 2026-06-05 gcpurls.Template snake_case matching requirement
- **Context**: Implementing IdentityV2 for `DevConnectInsightsConfig` in the `developerconnect` group, with the pattern `projects/{project}/locations/{location}/insightsConfigs/{insights_config}`.
- **Problem**: The `gcpurls.Template` parses placeholders and maps them to struct fields by converting both to lowercase. If a placeholder contains an underscore (e.g. `{insights_config}`), the lowercase representation `insights_config` is searched in the lowercase field name map of the struct. Using `InsightsConfig` as the field name results in a lowercase representation of `insightsconfig` (without underscore), causing a startup panic: `panic: field "insights_config" not found in struct v1alpha1.DevConnectInsightsConfigIdentity`.
- **Solution**: The Go identity struct field must match the casing of the placeholder. For `{insights_config}`, we must name the Go field `Insights_config string` in the identity struct (its lowercase representation is `insights_config`, which matches the placeholder perfectly).
- **Impact**: Any developer connect or other GCP services with snake_case parameters in their resource identifier URLs must follow this Go field naming convention to avoid template initialization panics.
