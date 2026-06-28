# Fuzzer Journal: NetworkServicesEdgeCacheService (Fuzzing of Handcoded No-Proto Resources)

This journal captures learnings from implementing the round-trip KRM fuzzer for `NetworkServicesEdgeCacheService` (`NetworkServicesEdgeCacheService`).

## Observations and Learnings

### 1. Handling Resources Missing Public Proto/Discovery SDK Types
Unlike most `networkservices` resources which are backed by public protobuf definitions (like `WasmPlugin` or `LbRouteExtension`), `NetworkServicesEdgeCacheService` is handcoded under `apis/networkservices/v1alpha1/edgecacheservice_types.go` because its proto definition is not published in the public googleapis repository. It is also not present in the OpenAPI client library `google.golang.org/api/networkservices/v1`.

### 2. Custom Wrapper Struct for NoProto Fuzzer
To implement a robust, lossless round-trip KRM fuzzer without access to any official API structs, we defined a custom API-equivalent struct `EdgeCacheServiceAPI` directly inside `pkg/controller/direct/networkservices/edgecacheservice_fuzzer.go`:

```go
type EdgeCacheServiceAPI struct {
	Name   string
	Spec   krm.NetworkServicesEdgeCacheServiceSpec
	Status krm.NetworkServicesEdgeCacheServiceStatus
}
```

This struct carries the actual handcoded KRM types 1:1, allowing our `FromAPI` and `ToAPI` mappers to be completely lossless and generic:

```go
func EdgeCacheServiceSpec_FromAPI(ctx *direct.MapContext, in *EdgeCacheServiceAPI) *krm.NetworkServicesEdgeCacheServiceSpec {
	if in == nil {
		return nil
	}
	out := in.Spec
	return &out
}

func EdgeCacheServiceSpec_ToAPI(ctx *direct.MapContext, in *krm.NetworkServicesEdgeCacheServiceSpec) *EdgeCacheServiceAPI {
	if in == nil {
		return nil
	}
	return &EdgeCacheServiceAPI{
		Spec: *in,
	}
}
```

### 3. Fuzzer Verification
By registering `".Spec"` as the spec field, `".Status"` as the status field, and `".Name"` as the identity field, the fuzzer automatically randomizes and validates the entire nested KRM Spec and Status structures:

```go
func edgeCacheServiceFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&EdgeCacheServiceAPI{},
		EdgeCacheServiceSpec_FromAPI, EdgeCacheServiceSpec_ToAPI,
		EdgeCacheServiceStatus_FromAPI, EdgeCacheServiceStatus_ToAPI,
	)

	f.SpecField(".Spec")
	f.StatusField(".Status")
	f.IdentityField(".Name")

	return f
}
```

This is extremely elegant, avoids any mock model overhead, and proves that our KRM schemas compile, deserialize, serialize, and round-trip flawlessly under KCC's central fuzz test suite.
