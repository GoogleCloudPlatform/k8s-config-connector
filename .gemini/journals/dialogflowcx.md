# DialogflowCX Journal

## DialogflowCXAgent Direct Controller Implementation

- **Service-Generated Resource ID**: `DialogflowCXAgent` uses server-assigned IDs (UUIDs). Reconciler logic follows the service-generated ID pattern:
  - `Find()` pre-checks if identity is specified (`!a.id.HasIdentitySpecified()`) and returns `false, nil` without calling API when empty.
  - `Create()` does not pass `resource.Name` in request body, letting GCP assign the UUID.
  - On create success, `parsedID.FromExternal(created.GetName())` sets the server-assigned UUID in `a.id.Agent`.
  - `status.externalRef` and `status.name` are populated in status update.
- **Regional Endpoint Construction**: Dialogflow CX API uses endpoint format `{location}-dialogflow.googleapis.com` (e.g. `global-dialogflow.googleapis.com`). Endpoints must be formatted without `:443` port suffix to cleanly align with GAPIC REST client HTTP request URLs.
- **Reconciliation & CompareGreenfieldSpec**: Using `common.CompareGreenfieldSpec` with `DialogflowCXAgentSpec_FromProto` and `DialogflowCXAgentSpec_ToProto` accurately identifies spec field diffs and produces field mask for `UpdateAgentRequest`.
