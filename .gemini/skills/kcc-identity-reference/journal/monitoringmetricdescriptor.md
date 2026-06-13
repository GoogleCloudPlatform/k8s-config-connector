# MonitoringMetricDescriptor Identity and Reference Journal

## Observations

1. **Missing from CAI Names list**:
   `MonitoringMetricDescriptor` is not listed in `docs/ai/metadata/cloudassetinventory_names.jsonl` (e.g. not handled by CAIS). We successfully registered an exception for it in `pkg/gcpurls/registry_test.go` under `ignoredTemplates`.

2. **Metric ID Containing Slashes**:
   The `type` field (which maps to `{metricid}` / resource ID) often contains slashes (e.g. `custom.googleapis.com/invoice/paid/amount`). Standard `gcpurls.Template` split logic splits segments by `/` and checks for segment count equality, which fails when fields have nested slashes. To handle this cleanly and safely without breaking existing standard template parsing, we implemented a robust custom parser in `FromExternal` that manually extracts the `Project` and `MetricID` segments when the reference is structured as `projects/{project}/metricDescriptors/{metricid}`.
