package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/structpb"
)

// Override but should be unreachable.
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func StructSchema_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func Schema_JsonSchema_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) *string {
	mapCtx.NotImplemented()
	return nil
}

// Override but should be unreachable
// Would not be needed if we did a reachability analysis in our proto generation.
func Schema_JsonSchema_ToProto(mapCtx *direct.MapContext, in *string) *pb.Schema_StructSchema {
	mapCtx.NotImplemented()
	return nil
}
