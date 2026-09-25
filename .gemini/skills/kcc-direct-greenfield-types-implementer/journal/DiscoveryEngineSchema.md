# DiscoveryEngineSchema Greenfield Types Implementation Notes

- Resource Kind: `DiscoveryEngineSchema`
- GCP Proto: `google.cloud.discoveryengine.v1.Schema`
- URL Format: `projects/{project}/locations/{location}/dataStores/{datastore}/schemas/{schema}`
- Identity & Reference implemented following `IdentityV2` and `refs.Ref` using `gcpurls.Template`.
- Added URL template exception to `pkg/gcpurls/registry_test.go` as DiscoveryEngineSchema is not present in CAI metadata.
- Registered in `pkg/controller/resourceconfig/static_config.go` with `ReconcilerTypeDirect`.
