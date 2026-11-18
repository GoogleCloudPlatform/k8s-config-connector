# DialogflowCX Journal

## DialogflowCXAgent Direct Controller Implementation

- **Service-Generated Resource ID**: `DialogflowCXAgent` uses server-assigned IDs (UUIDs). Reconciler logic follows the service-generated ID pattern:
  - `Find()` pre-checks if identity is specified (`!a.id.HasIdentitySpecified()`) and returns `false, nil` without calling API when empty.
  - `Create()` does not pass `resource.Name` in request body, letting GCP assign the UUID.
  - On create success, `parsedID.FromExternal(created.GetName())` sets the server-assigned UUID in `a.id.Agent`.
  - `status.externalRef` and `status.name` are populated in status update.
- **Regional Endpoint Construction**: Dialogflow CX API uses endpoint format `{location}-dialogflow.googleapis.com` (e.g. `global-dialogflow.googleapis.com`). Endpoints must be formatted without `:443` port suffix to cleanly align with GAPIC REST client HTTP request URLs.
- **Reconciliation & CompareGreenfieldSpec**: Using `common.CompareGreenfieldSpec` with `DialogflowCXAgentSpec_FromProto` and `DialogflowCXAgentSpec_ToProto` accurately identifies spec field diffs and produces field mask for `UpdateAgentRequest`.

## DialogflowCXWebhook Direct Controller Implementation

- **Service-Generated Resource ID**: `DialogflowCXWebhook` uses server-assigned IDs (UUIDs). Follows service-generated ID rules:
  - `Find()` pre-checks `!a.id.HasIdentitySpecified()` and returns `false, nil` without calling API when empty.
  - `Create()` clears `resource.Name = ""` before sending `CreateWebhook` request so GCP assigns the UUID.
  - Sets `spec.resourceID` via `createOp.SetSpecResourceID` upon creation to maintain backward compatibility with brownfield Terraform controller.
  - Populates both `status.name` and `status.externalRef` on status updates.
- **Reconciliation & CompareBrownfieldSpec**: Used `common.CompareBrownfieldSpec` with `DialogflowCXWebhookSpec_FromProto` and `DialogflowCXWebhookSpec_ToProto` to compare spec fields and generate the FieldMask for `UpdateWebhookRequest`, ensuring 0-write steady-state re-reconciliation.
- **Client & Endpoint**: GAPIC REST client (`cx.NewWebhooksRESTClient`) with `{location}-dialogflow.googleapis.com` endpoint without port suffix cleanly aligns with MockGCP and real GCP recording.

