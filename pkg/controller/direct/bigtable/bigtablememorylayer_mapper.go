package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableMemoryLayerSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableMemoryLayerSpec) *pb.MemoryLayer {
	if in == nil {
		return nil
	}
	out := &pb.MemoryLayer{}
	// Setting an empty MemoryConfig formally enables the MemoryLayer in GCP
	out.MemoryConfig = &pb.MemoryLayer_MemoryConfig{}
	return out
}

func BigtableMemoryLayerSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MemoryLayer) *krm.BigtableMemoryLayerSpec {
	if in == nil {
		return nil
	}
	// KRM Spec holds no other fields; ResourceID/ClusterRef are managed by the Adapter logic
	return &krm.BigtableMemoryLayerSpec{}
}
