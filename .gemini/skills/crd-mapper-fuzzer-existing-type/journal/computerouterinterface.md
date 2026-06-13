# Journal: ComputeRouterInterface Transition to Direct KRM Types

## Context
Transitioned the `ComputeRouterInterface` resource to use direct Go KRM types under `apis/compute/v1beta1/routerinterface_types.go`, while maintaining strict schema compatibility with the baseline CRD.

## Learnings & Patterns Identified

### 1. Acronym Field Matching
When a KRM reference field matches a Proto field but uses different capitalization for acronyms (e.g., `PrivateIpAddressRef` in KRM vs `private_ip_address` in Proto / `PrivateIPAddress` in the Go protobuf struct), renaming the Go struct field to match the acronym exactly (e.g., `PrivateIPAddressRef`) allows the generator to automatically map the field properly. Since the JSON tag `json:"privateIpAddressRef"` remains unchanged, the CRD schema stays identical.

### 2. Standard Field Mapping via Json Tag Alignment
The KRM-level `resourceID` field maps to `name` in the Proto message. By naming the Go struct field `Name` but preserving the json tag `json:"resourceID,omitempty"`, we let `generate-mapper` fully automate the `Name` field mapping to and from the proto. The CRD schema is unaffected since the JSON property name remains `resourceID`.

### 3. Fuzzer Setup for Empty Status Types
When a resource does not have any status fields mapping to proto, a custom status FromProto and ToProto mapping can be implemented locally in the fuzzer file:
```go
func statusFromProto(mapCtx *direct.MapContext, in *pb.RouterInterface) *krm.ComputeRouterInterfaceStatus {
	return &krm.ComputeRouterInterfaceStatus{}
}

func statusToProto(mapCtx *direct.MapContext, in *krm.ComputeRouterInterfaceStatus) *pb.RouterInterface {
	if in == nil {
		return nil
	}
	return &pb.RouterInterface{}
}
```
Returning a non-nil empty struct `&pb.RouterInterface{}` in `statusToProto` prevents the fuzzing library from flagging the round-trip as invalid.
