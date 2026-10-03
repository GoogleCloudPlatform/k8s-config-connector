// +tool:fuzz-gen
// proto.message: google.bigtable.admin.v2.MemoryLayer
// api.group: bigtable.cnrm.cloud.google.com

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(bigtableMemoryLayerFuzzer())
}

func bigtableMemoryLayerFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.MemoryLayer{},
		BigtableMemoryLayerSpec_v1alpha1_FromProto, BigtableMemoryLayerSpec_v1alpha1_ToProto,
	)
	f.UnimplementedFields.Insert(".name", ".memory_config") 
	return f
}
