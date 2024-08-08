package cluster

// func Cluster_StateInfo_FromProto(ctx *direct.MapContext, in *pb.Cluster_StateInfo) *krm.Cluster_StateInfo {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.Cluster_StateInfo{}
// 	out.UpdateInfo = Cluster_StateInfo_UpdateInfo_FromProto(ctx, in.GetUpdateInfo())
// 	return out
// }
// func Cluster_StateInfo_ToProto(ctx *MapContext, in *krm.Cluster_StateInfo) *pb.Cluster_StateInfo {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster_StateInfo{}
// 	if in.UpdateInfo != nil {
// 		oneof := &pb.Cluster_StateInfo_UpdateInfo_{}
// 		out.Info = oneof
// 		oneof.UpdateInfo = Cluster_StateInfo_UpdateInfo_ToProto(ctx, in.UpdateInfo)
// 	}

// 	return out
// }
// func Cluster_StateInfo_UpdateInfo_FromProto(ctx *MapContext, in *pb.Cluster_StateInfo_UpdateInfo) *krm.Cluster_StateInfo_UpdateInfo {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.Cluster_StateInfo_UpdateInfo{}
// 	out.TargetShardCount = (in.TargetShardCount)
// 	out.TargetReplicaCount = (in.TargetReplicaCount)
// 	return out
// }
// func Cluster_StateInfo_UpdateInfo_ToProto(ctx *MapContext, in *krm.Cluster_StateInfo_UpdateInfo) *pb.Cluster_StateInfo_UpdateInfo {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.Cluster_StateInfo_UpdateInfo{}
// 	out.TargetShardCount = in.TargetShardCount
// 	out.TargetReplicaCount = in.TargetReplicaCount
// 	return out
// }

// func PscConfig_FromProto(ctx *MapContext, in *pb.PscConfig) *krm.PscConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.PscConfig{}
// 	out.Network = LazyPtr(in.Network)
// 	return out
// }
// func PscConfig_ToProto(ctx *MapContext, in *krm.PscConfig) *pb.PscConfig {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.PscConfig{}

// 	if in.Network != nil {
// 		out.Network = ValueOf(in.Network)
// 	}

// 	// if in.ServiceConnectionPolicyRef != nil {
// 	// 	ref := ctx.ResolveServiceConnectionPolicyRef(in.ServiceConnectionPolicyRef)
// 	// 	out.ServiceConnectionPolicy = ValueOf(ref)
// 	// }

// 	return out
// }

// func DiscoveryEndpoint_FromProto(ctx *MapContext, in *pb.DiscoveryEndpoint) *krm.DiscoveryEndpoint {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.DiscoveryEndpoint{}
// 	out.Address = LazyPtr(in.Address)
// 	out.Port = LazyPtr(in.Port)
// 	out.PscConfig = PscConfig_FromProto(ctx, in.PscConfig)
// 	return out
// }
// func DiscoveryEndpoint_ToProto(ctx *MapContext, in *krm.DiscoveryEndpoint) *pb.DiscoveryEndpoint {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.DiscoveryEndpoint{}
// 	out.Address = ValueOf(in.Address)
// 	out.Port = ValueOf(in.Port)
// 	out.PscConfig = PscConfig_ToProto(ctx, in.PscConfig)
// 	return out
// }
// func PscConnection_FromProto(ctx *MapContext, in *pb.PscConnection) *krm.PscConnection {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &krm.PscConnection{}
// 	out.PscConnectionId = LazyPtr(in.PscConnectionId)
// 	out.Address = LazyPtr(in.Address)
// 	out.ForwardingRule = LazyPtr(in.ForwardingRule)
// 	out.ProjectId = LazyPtr(in.ProjectId)
// 	out.Network = LazyPtr(in.Network)
// 	return out
// }
// func PscConnection_ToProto(ctx *MapContext, in *krm.PscConnection) *pb.PscConnection {
// 	if in == nil {
// 		return nil
// 	}
// 	out := &pb.PscConnection{}
// 	out.PscConnectionId = ValueOf(in.PscConnectionId)
// 	out.Address = ValueOf(in.Address)
// 	out.ForwardingRule = ValueOf(in.ForwardingRule)
// 	out.ProjectId = ValueOf(in.ProjectId)
// 	out.Network = ValueOf(in.Network)
// 	return out
// }
