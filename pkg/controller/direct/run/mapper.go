package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Volume_FromProto(mapCtx *direct.MapContext, in *pb.Volume) *krm.Volume {
	if in == nil {
		return nil
	}
	out := &krm.Volume{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Secret = SecretVolumeSource_FromProto(mapCtx, in.GetSecret())
	out.CloudSQLInstance = CloudSQLInstance_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.EmptyDir = EmptyDirVolumeSource_FromProto(mapCtx, in.GetEmptyDir())
	out.NFS = NFSVolumeSource_FromProto(mapCtx, in.GetNfs())
	out.GCS = GCSVolumeSource_FromProto(mapCtx, in.GetGcs())
	return out
}

func Volume_ToProto(mapCtx *direct.MapContext, in *krm.Volume) *pb.Volume {
	if in == nil {
		return nil
	}
	out := &pb.Volume{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := SecretVolumeSource_ToProto(mapCtx, in.Secret); oneof != nil {
		out.VolumeType = &pb.Volume_Secret{Secret: oneof}
	}
	if oneof := CloudSQLInstance_ToProto(mapCtx, in.CloudSQLInstance); oneof != nil {
		out.VolumeType = &pb.Volume_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := EmptyDirVolumeSource_ToProto(mapCtx, in.EmptyDir); oneof != nil {
		out.VolumeType = &pb.Volume_EmptyDir{EmptyDir: oneof}
	}
	if oneof := NFSVolumeSource_ToProto(mapCtx, in.NFS); oneof != nil {
		out.VolumeType = &pb.Volume_Nfs{Nfs: oneof}
	}
	if oneof := GCSVolumeSource_ToProto(mapCtx, in.GCS); oneof != nil {
		out.VolumeType = &pb.Volume_Gcs{Gcs: oneof}
	}
	return out
}
