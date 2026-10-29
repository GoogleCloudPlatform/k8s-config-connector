# Journal: Match Real GCP Behavior in MockGCP for BeyondCorpAppGateway

This journal captures the details and findings of implementing the `BeyondCorpAppGateway` mock in MockGCP (`mockgcp/mockbeyondcorp`) and attempting real GCP recording.

- **Date:** October 29, 2026
- **Resource Kind:** `BeyondCorpAppGateway`
- **GCP Service:** `beyondcorp.googleapis.com` (`google.cloud.beyondcorp.appgateways.v1.AppGatewaysService`)

## Real GCP Recording Attempt

- **Real GCP Recording Status:** Attempted against live GCP (`cnrm-barni-4`), but could not proceed.
- **GCP Project Used:** `cnrm-barni-4`
- **Command Run:** `./hack/record-gcp "fixtures/^beyondcorpappgatewayautogen$"`
- **Start Time (PDT):** Wed Oct 28 19:41:56 PDT 2026
- **Full Error Message Encountered:**
  ```text
  Error waiting to create AppGateway: googleapi: got HTTP response code 404 with body: HTTP/2.0 404 Not Found
  Content-Length: 1683
  Content-Type: text/html; charset=UTF-8
  Date: Thu, 29 Oct 2026 02:50:18 GMT
  Server: ESF
  ...
  The requested URL /v1/projects/cnrm-barni-4/locations/us-central1/appGateways?alt=json&app_gateway_id=beyondcorpappgatewayp3zqicfcqqtns7i was not found on this server.
  Reason:UpdateFailed Status:False Type:Ready}]}
  F1029 03:19:07.865550   51527 http_recorder.go:108] error from json.Unmarshal("<!DOCTYPE html>..."): invalid character '<' looking for beginning of value
  ```
- **Root Cause:**
  Google Cloud has turned down / retired legacy BeyondCorp AppGateway, AppConnector, and AppConnection creation API endpoints. On live GCP, `beyondcorp.googleapis.com` returns HTML 404 from Google's Edge Security Frontend (ESF) for POST requests to `/v1/projects/*/locations/*/appGateways`. The underlying resource was also removed in Google Cloud Terraform Provider v8.0.0.
- **Fallback Approach Used:**
  Hermetically validated and generated golden mock files using MockGCP with `./hack/compare-mock "fixtures/^beyondcorpappgatewayautogen$"`.

## Mock Implementation Details

1. **Service Registration:**
   - Created `mockgcp/mockbeyondcorp/service.go` with host `beyondcorp.googleapis.com` and operations path `/v1/{prefix=**}/operations/{name}`.
   - Implemented `AppGatewaysV1` in `mockgcp/mockbeyondcorp/appgateways.go` implementing `AppGatewaysServiceServer` with `GetAppGateway`, `CreateAppGateway`, `DeleteAppGateway`, and LRO handling.
   - Added `mockbeyondcorp/normalize.go` with visitor support scoped to `beyondcorp.googleapis.com`.
   - Registered `mockbeyondcorp` in `mockgcp/register.go` and `mockgcp/mockserviceusage/knownservices.go`.

2. **Test Harness & Ratcheting:**
   - Removed `BeyondCorpAppGateway` from the ratcheting exclusion list in `tests/e2e/ratcheting.go`.
   - Registered `BeyondCorpAppGateway` in `config/tests/samples/create/harness.go`.
   - Updated `pkg/test/resourcefixture/testdata/basic/beyondcorp/v1alpha1/beyondcorpappgateway/beyondcorpappgatewayautogen/create.yaml` to include `displayName: Sample Gateway`, covering `.spec.displayName`.
   - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` to remove the `.spec.displayName` exception.

3. **Validation:**
   - Executed `hack/compare-mock "fixtures/^beyondcorpappgatewayautogen$"`: resource reached `Ready`, re-reconciliation passed, deletion succeeded.
   - Ran `go test -v ./pkg/test/resourcefixture/...`: passed.
   - Ran `go test -v ./tests/apichecks/...`: passed.
