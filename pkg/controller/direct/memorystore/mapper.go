package memorystore

import (
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PscAutoConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PscAutoConnectionSpec) *krm.PscAutoConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PscAutoConnectionSpec{}
	out.NetworkRef = direct.LazyPtr(in.GetNetworkRef())
	out.ProjectRef = direct.LazyPtr(in.GetProjectRef())
	return out
}
func PscAutoConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PscAutoConnectionSpec) *pb.PscAutoConnectionSpec {
	if in == nil {
		return nil
	}
	out := &pb.PscAutoConnectionSpec{}
	out.NetworkRef = direct.ValueOf(in.NetworkRef)
	out.ProjectRef = direct.ValueOf(in.ProjectRef)
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceSpec{}
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.PscAutoConnectionsSpec = PscAutoConnectionSpec_FromProto(mapCtx, in.PscAutoConnectionSpec)
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	out.PscAutoConnectionSpec = PscAutoConnectionSpec_ToProto(mapCtx, in.PscAutoConnectionsSpec)
	return out
}
