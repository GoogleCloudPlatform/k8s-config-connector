# BeyondCorpAppGateway Direct Controller Journal

## Implementation Details

- **Controller Architecture**: Implemented `Adapter` under `pkg/controller/direct/beyondcorp/beyondcorpappgateway/beyondcorpappgateway_controller.go` using the official `cloud.google.com/go/beyondcorp/appgateways/apiv1` GAPIC client.
- **Client Protocol**: Utilized `m.config.GRPCClientOptions()` with `gcp.NewClient` because GAPIC does not provide `NewRESTClient` for this API.
- **Immutability Handling**: BeyondCorp AppGateways cannot be updated in GCP. The controller checks for diffs using `common.CompareBrownfieldSpecAndLabels` and surfaces an error on updates while passing steady-state re-reconciliation without sending mutating requests.
- **LRO Handling**: `CreateAppGateway` and `DeleteAppGateway` return operations that are waited on with `op.Wait(ctx)`. Re-reads the latest state using `GetAppGateway` after creation LRO completion before updating status.
- **Export Support**: Implemented `Export` and `AdapterForURL` parsing canonical external references (`projects/{project}/locations/{location}/appGateways/{appGateway}`).
- **Testing against MockGCP**: Verified using `hack/compare-mock` against `pkg/test/resourcefixture/testdata/basic/beyondcorp/v1alpha1/beyondcorpappgateway/beyondcorpappgatewayautogen`. Both direct and fallback terraform reconcilers passed creation, steady-state re-reconciliation, and deletion.
