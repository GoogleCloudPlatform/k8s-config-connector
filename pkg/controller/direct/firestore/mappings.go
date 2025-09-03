package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IndexFields_Order_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_Order_ {
	if in == nil {
		return nil
	}

	v := direct.Enum_ToProto[pb.Index_IndexField_Order](mapCtx, in)
	out := &pb.Index_IndexField_Order_{Order: v}
	return out
}

func IndexFields_ArrayConfig_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_ArrayConfig_ {
	if in == nil {
		return nil
	}

	v := direct.Enum_ToProto[pb.Index_IndexField_ArrayConfig](mapCtx, in)
	out := &pb.Index_IndexField_ArrayConfig_{ArrayConfig: v}
	return out
}
