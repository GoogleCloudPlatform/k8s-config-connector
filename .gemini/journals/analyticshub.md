### [2026-10-29] Greenfield Step 1b Judgement Triage for BigQueryAnalyticsHubQueryTemplate
- **Context**: Triaged judgement queue and finalized direct KRM types and identity for BigQueryAnalyticsHubQueryTemplate (issue #14666).
- **Problem**: QueryTemplate is a child resource of DataExchange with URL pattern `projects/{project}/locations/{location}/dataExchanges/{data_exchange}/queryTemplates/{query_template}`. The initial scaffolding had open entries for `identity-root-unknown`, guessed `dataExchangeRef`, and `location-or-parent-ref`. Additionally, `CreateQueryTemplateRequest` has `query_template_id` while `.name` was marked `OUTPUT_ONLY`.
- **Solution**: Replaced `ProjectRef`, `Location`, and guessed `DataExchangeRef` on `BigQueryAnalyticsHubQueryTemplateSpec` with a single immutable, required `ParentRef *BigQueryAnalyticsHubDataExchangeRef`. Confirmed `query_template_id` is client-specified so `BigQueryAnalyticsHubQueryTemplateIdentity` implements `identity.IdentityV2`. Regenerated types, identity, CRD, and verified fuzzer and controller tests.
- **Impact**: Clean single-parent hierarchy conforming to KCC direct conventions and ready for Step 2 controller reconciliation testing.

### [2026-10-30] Greenfield Step 2 Controller & Fixtures for BigQueryAnalyticsHubQueryTemplate
- **Context**: Greenfield direct controller, fuzzer, and real-GCP E2E fixtures for `BigQueryAnalyticsHubQueryTemplate` (#14699).
- **Problem**: During real-GCP recording:
  1. GCP rejected QueryTemplate creation under standard DataExchanges with "Query templates are only supported for data clean rooms." DataExchanges need `sharingEnvironmentConfig.dcrExchangeConfig` enabled to act as clean rooms.
  2. `routine` is mandatory on create, and GCP requires the routine name in `definitionBody` to match the QueryTemplate resource ID (e.g. `<resource_id>() AS (...)`).
  3. GCP requires `primaryContact` to equal the authenticated caller's email (`overseer-kcc-tester@${projectId}.iam.gserviceaccount.com`).
  4. GCP rejects updating `displayName`, `documentation`, and `proposer`. Furthermore, GCP drops `proposer` from read responses because it is deprecated in favor of `primaryContact`, causing `CompareGreenfieldSpec` to report a diff and try to update it.
- **Solution**:
  1. Exposed `sharingEnvironmentConfig.dcrExchangeConfig` on `BigQueryAnalyticsHubDataExchange` so test fixtures can create clean room exchanges.
  2. Crafted minimal and maximal fixtures with valid TVF routine definition bodies matching the template ID and authenticated primary contact.
  3. Added a `normalize` hook in `CompareGreenfieldSpec` in `QueryTemplateAdapter.Update` to blank out `Proposer`, avoiding false drift on the deprecated field.
- **Impact**: Clean real-GCP reconciliation and re-reconciliation without diffs or update failures. Future tests and controllers under Analytics Hub clean rooms can reference these requirements.
