# Journal: BeyondCorpAppConnection MockGCP Implementation

This journal documents the findings and implementation details for the `BeyondCorpAppConnection` mock in MockGCP.

## Observations

1. **Service Proto and Generation**:
   - The Go protobuf bindings for BeyondCorp AppConnections exist in `cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb`.
   - Related services for BeyondCorp include `appconnectorspb`, `appgatewayspb`, and `clientconnectorservicespb`.
   - The mock service was built using `httptogrpc.NewGRPCMux` to bridge HTTP REST requests to the gRPC service implementations.

2. **Real GCP Recording Status**:
   - Attempted recording against real GCP using `hack/record-gcp pkg/test/resourcefixture/testdata/basic/beyondcorp/v1alpha1/beyondcorpappconnection/beyondcorpappconnectionautogen` with project `cnrm-barni-4`.
   - Direct API calls with `gcloud` and HTTP POST requests returned HTTP 404:
     `The requested URL /v1/projects/cnrm-barni-4/locations/us-central1/appConnections?app_connection_id=... was not found on this server.`
   - GCP has deprecated BeyondCorp AppConnectors/AppConnections (or restricted write access via early-access allowlist).
   - Per reporting requirements, since `record-gcp` could not be run against real GCP due to endpoint 404/deprecation, the fallback approach was used: golden logs and object definitions were generated and verified through MockGCP using `hack/compare-mock`.

3. **Service Account ID Constraint in MockIAM**:
   - In `dependencies.yaml`, `IAMServiceAccount` was originally using `resourceID: iamserviceaccount${uniqueId}` which exceeded the 30-character account ID limit in MockIAM (`must be 6-30 characters long`). This was updated to `sa-${uniqueId}` to satisfy IAM constraints.
