# DeveloperConnect Service Journal

### 2026-06-05 gcpurls.Template snake_case matching requirement (DevConnectInsightsConfig)
- **Context**: Implementing IdentityV2 for `DevConnectInsightsConfig` in the `developerconnect` group, with the pattern `projects/{project}/locations/{location}/insightsConfigs/{insights_config}`.
- **Problem**: The `gcpurls.Template` parses placeholders and maps them to struct fields by converting both to lowercase. If a placeholder contains an underscore (e.g. `{insights_config}`), the lowercase representation `insights_config` is searched in the lowercase field name map of the struct. Using `InsightsConfig` as the field name results in a lowercase representation of `insightsconfig` (without underscore), causing a startup panic: `panic: field "insights_config" not found in struct v1alpha1.DevConnectInsightsConfigIdentity`.
- **Solution**: The Go identity struct field must match the casing of the placeholder. For `{insights_config}`, we must name the Go field `Insights_config string` in the identity struct (its lowercase representation is `insights_config`, which matches the placeholder perfectly).
- **Impact**: Any developer connect or other GCP services with snake_case parameters in their resource identifier URLs must follow this Go field naming convention to avoid template initialization panics.

### 2026-06-05 gcpurls.Template Underscore Mapping in Identity (DevConnectAccountConnector)
- **Context**: Implementing IdentityV2 for `DevConnectAccountConnector`.
- **Problem**: The URL template has a multi-word placeholder with an underscore: `projects/{project}/locations/{location}/accountConnectors/{account_connector}`. When naming the Go struct field `AccountConnector`, `gcpurls.Template` failed to map it and panicked: `panic: field "account_connector" not found in struct v1alpha1.DevConnectAccountConnectorIdentity`.
- **Solution**: The `gcpurls.Template` matches fields case-insensitively, converting both to lower case. Lowercasing `AccountConnector` produces `accountconnector`, which does not match `account_connector`. Changing the Go struct field name to `Account_connector` (which lowercases to `account_connector`) successfully matched the placeholder.
- **Impact**: Any future KCC direct resource implementation utilizing a URL template with underscores in placeholders must use Go struct fields with underscores (e.g., `My_placeholder`) in their Identity struct to avoid panic during reflection-based mapping.
