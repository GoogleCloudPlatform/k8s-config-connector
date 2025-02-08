// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memorystore

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/memorystore/apiv1beta/memorystorepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CertificateAuthority_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority) *krm.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority{}
	out.ManagedServerCa = CertificateAuthority_ManagedCertificateAuthority_FromProto(mapCtx, in.GetManagedServerCa())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func CertificateAuthority_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority) *pb.CertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority{}
	if oneof := CertificateAuthority_ManagedCertificateAuthority_ToProto(mapCtx, in.ManagedServerCa); oneof != nil {
		out.ServerCa = &pb.CertificateAuthority_ManagedServerCa{ManagedServerCa: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_ManagedCertificateAuthority) *krm.CertificateAuthority_ManagedCertificateAuthority {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ManagedCertificateAuthority{}
	out.CaCerts = direct.Slice_FromProto(mapCtx, in.CaCerts, CertificateAuthority_ManagedCertificateAuthority_CertChain_FromProto)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ManagedCertificateAuthority) *pb.CertificateAuthority_ManagedCertificateAuthority {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_ManagedCertificateAuthority{}
	out.CaCerts = direct.Slice_ToProto(mapCtx, in.CaCerts, CertificateAuthority_ManagedCertificateAuthority_CertChain_ToProto)
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_CertChain_FromProto(mapCtx *direct.MapContext, in *pb.CertificateAuthority_ManagedCertificateAuthority_CertChain) *krm.CertificateAuthority_ManagedCertificateAuthority_CertChain {
	if in == nil {
		return nil
	}
	out := &krm.CertificateAuthority_ManagedCertificateAuthority_CertChain{}
	out.Certificates = in.Certificates
	return out
}
func CertificateAuthority_ManagedCertificateAuthority_CertChain_ToProto(mapCtx *direct.MapContext, in *krm.CertificateAuthority_ManagedCertificateAuthority_CertChain) *pb.CertificateAuthority_ManagedCertificateAuthority_CertChain {
	if in == nil {
		return nil
	}
	out := &pb.CertificateAuthority_ManagedCertificateAuthority_CertChain{}
	out.Certificates = in.Certificates
	return out
}
func MemorystoreInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.MemorystoreInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.MemorystoreInstanceSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: StateInfo
	// MISSING: Uid
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_FromProto(mapCtx, in.GetAuthorizationMode())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	out.ShardCount = direct.LazyPtr(in.GetShardCount())
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_FromProto(mapCtx, in.GetNodeType())
	out.PersistenceConfig = PersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.EngineVersion = direct.LazyPtr(in.GetEngineVersion())
	out.EngineConfigs = in.EngineConfigs
	// MISSING: NodeConfig
	out.ZoneDistributionConfig = ZoneDistributionConfig_FromProto(mapCtx, in.GetZoneDistributionConfig())
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PscAutoConnections
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
func MemorystoreInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.MemorystoreInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: StateInfo
	// MISSING: Uid
	out.ReplicaCount = in.ReplicaCount
	out.AuthorizationMode = direct.Enum_ToProto[pb.Instance_AuthorizationMode](mapCtx, in.AuthorizationMode)
	out.TransitEncryptionMode = direct.Enum_ToProto[pb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	out.ShardCount = direct.ValueOf(in.ShardCount)
	// MISSING: DiscoveryEndpoints
	out.NodeType = direct.Enum_ToProto[pb.Instance_NodeType](mapCtx, in.NodeType)
	out.PersistenceConfig = PersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.EngineVersion = direct.ValueOf(in.EngineVersion)
	out.EngineConfigs = in.EngineConfigs
	// MISSING: NodeConfig
	out.ZoneDistributionConfig = ZoneDistributionConfig_ToProto(mapCtx, in.ZoneDistributionConfig)
	out.DeletionProtectionEnabled = in.DeletionProtectionEnabled
	// MISSING: PscAutoConnections
	// MISSING: Endpoints
	// MISSING: Mode
	return out
}
