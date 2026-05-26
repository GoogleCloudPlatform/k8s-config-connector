// Copyright 2026 Google LLC
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

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DeployedIndexAuthConfig_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexAuthConfig) *krm.DeployedIndexAuthConfig {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexAuthConfig{}
	out.AuthProvider = DeployedIndexAuthConfig_AuthProvider_FromProto(mapCtx, in.GetAuthProvider())
	return out
}

func DeployedIndexAuthConfig_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexAuthConfig) *pb.DeployedIndexAuthConfig {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexAuthConfig{}
	out.AuthProvider = DeployedIndexAuthConfig_AuthProvider_ToProto(mapCtx, in.AuthProvider)
	return out
}

func DeployedIndexAuthConfig_AuthProvider_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndexAuthConfig_AuthProvider) *krm.DeployedIndexAuthConfig_AuthProvider {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexAuthConfig_AuthProvider{}
	out.Audiences = in.Audiences
	out.AllowedIssuers = in.AllowedIssuers
	return out
}

func DeployedIndexAuthConfig_AuthProvider_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexAuthConfig_AuthProvider) *pb.DeployedIndexAuthConfig_AuthProvider {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndexAuthConfig_AuthProvider{}
	out.Audiences = in.Audiences
	out.AllowedIssuers = in.AllowedIssuers
	return out
}

func DeployedIndexObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeployedIndex) *krm.DeployedIndexObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployedIndexObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Index = direct.LazyPtr(in.GetIndex())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.PrivateEndpoints = IndexPrivateEndpoints_FromProto(mapCtx, in.GetPrivateEndpoints())
	out.IndexSyncTime = direct.StringTimestamp_FromProto(mapCtx, in.GetIndexSyncTime())
	// MISSING: AutomaticResources
	// MISSING: DedicatedResources
	out.EnableAccessLogging = direct.LazyPtr(in.GetEnableAccessLogging())
	// MISSING: EnableDatapointUpsertLogging
	out.DeployedIndexAuthConfig = DeployedIndexAuthConfig_FromProto(mapCtx, in.GetDeployedIndexAuthConfig())
	out.ReservedIPRanges = in.ReservedIpRanges
	out.DeploymentGroup = direct.LazyPtr(in.GetDeploymentGroup())
	out.PSCAutomationConfigs = direct.Slice_FromProto(mapCtx, in.PscAutomationConfigs, PSCAutomationConfig_FromProto)
	return out
}

func DeployedIndexObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployedIndexObservedState) *pb.DeployedIndex {
	if in == nil {
		return nil
	}
	out := &pb.DeployedIndex{}
	out.Id = direct.ValueOf(in.ID)
	out.Index = direct.ValueOf(in.Index)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.PrivateEndpoints = IndexPrivateEndpoints_ToProto(mapCtx, in.PrivateEndpoints)
	out.IndexSyncTime = direct.StringTimestamp_ToProto(mapCtx, in.IndexSyncTime)
	// MISSING: AutomaticResources
	// MISSING: DedicatedResources
	out.EnableAccessLogging = direct.ValueOf(in.EnableAccessLogging)
	// MISSING: EnableDatapointUpsertLogging
	out.DeployedIndexAuthConfig = DeployedIndexAuthConfig_ToProto(mapCtx, in.DeployedIndexAuthConfig)
	out.ReservedIpRanges = in.ReservedIPRanges
	out.DeploymentGroup = direct.ValueOf(in.DeploymentGroup)
	out.PscAutomationConfigs = direct.Slice_ToProto(mapCtx, in.PSCAutomationConfigs, PSCAutomationConfig_ToProto)
	return out
}

func IndexPrivateEndpoints_FromProto(mapCtx *direct.MapContext, in *pb.IndexPrivateEndpoints) *krm.IndexPrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &krm.IndexPrivateEndpoints{}
	// MISSING: MatchGrpcAddress
	// (near miss): "MatchGrpcAddress" vs "MatchGRPCAddress"
	out.ServiceAttachment = direct.LazyPtr(in.GetServiceAttachment())
	// MISSING: PSCAutomatedEndpoints
	return out
}

func IndexPrivateEndpoints_ToProto(mapCtx *direct.MapContext, in *krm.IndexPrivateEndpoints) *pb.IndexPrivateEndpoints {
	if in == nil {
		return nil
	}
	out := &pb.IndexPrivateEndpoints{}
	// MISSING: MatchGrpcAddress
	// (near miss): "MatchGrpcAddress" vs "MatchGRPCAddress"
	out.ServiceAttachment = direct.ValueOf(in.ServiceAttachment)
	// MISSING: PSCAutomatedEndpoints
	return out
}

func PSCAutomationConfig_FromProto(mapCtx *direct.MapContext, in *pb.PSCAutomationConfig) *krm.PSCAutomationConfig {
	if in == nil {
		return nil
	}
	out := &krm.PSCAutomationConfig{}
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Network = direct.LazyPtr(in.GetNetwork())
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	// MISSING: State
	// MISSING: ErrorMessage
	return out
}

func PSCAutomationConfig_ToProto(mapCtx *direct.MapContext, in *krm.PSCAutomationConfig) *pb.PSCAutomationConfig {
	if in == nil {
		return nil
	}
	out := &pb.PSCAutomationConfig{}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.Network = direct.ValueOf(in.Network)
	// MISSING: IPAddress
	// MISSING: ForwardingRule
	// MISSING: State
	// MISSING: ErrorMessage
	return out
}

func PrivateServiceConnectConfig_FromProto(mapCtx *direct.MapContext, in *pb.PrivateServiceConnectConfig) *krm.PrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &krm.PrivateServiceConnectConfig{}
	out.EnablePrivateServiceConnect = direct.LazyPtr(in.GetEnablePrivateServiceConnect())
	out.ProjectAllowlist = in.ProjectAllowlist
	out.PSCAutomationConfigs = direct.Slice_FromProto(mapCtx, in.PscAutomationConfigs, PSCAutomationConfig_FromProto)
	// MISSING: ServiceAttachment
	return out
}

func PrivateServiceConnectConfig_ToProto(mapCtx *direct.MapContext, in *krm.PrivateServiceConnectConfig) *pb.PrivateServiceConnectConfig {
	if in == nil {
		return nil
	}
	out := &pb.PrivateServiceConnectConfig{}
	out.EnablePrivateServiceConnect = direct.ValueOf(in.EnablePrivateServiceConnect)
	out.ProjectAllowlist = in.ProjectAllowlist
	out.PscAutomationConfigs = direct.Slice_ToProto(mapCtx, in.PSCAutomationConfigs, PSCAutomationConfig_ToProto)
	// MISSING: ServiceAttachment
	return out
}

func VertexAIIndexEndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.VertexAIIndexEndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIIndexEndpointObservedState{}
	// MISSING: Name
	out.DeployedIndexes = direct.Slice_FromProto(mapCtx, in.DeployedIndexes, DeployedIndexObservedState_FromProto)
	// MISSING: Etag
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: EnablePrivateServiceConnect
	out.PublicEndpointDomainName = direct.LazyPtr(in.GetPublicEndpointDomainName())
	// MISSING: SatisfiesPzs
	// (near miss): "SatisfiesPzs" vs "SatisfiesPZS"
	// MISSING: SatisfiesPzi
	// (near miss): "SatisfiesPzi" vs "SatisfiesPZI"
	return out
}

func VertexAIIndexEndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIIndexEndpointObservedState) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	out.DeployedIndexes = direct.Slice_ToProto(mapCtx, in.DeployedIndexes, DeployedIndexObservedState_ToProto)
	// MISSING: Etag
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: EnablePrivateServiceConnect
	out.PublicEndpointDomainName = direct.ValueOf(in.PublicEndpointDomainName)
	// MISSING: SatisfiesPzs
	// (near miss): "SatisfiesPzs" vs "SatisfiesPZS"
	// MISSING: SatisfiesPzi
	// (near miss): "SatisfiesPzi" vs "SatisfiesPZI"
	return out
}

func VertexAIIndexEndpointSpec_FromProto(mapCtx *direct.MapContext, in *pb.IndexEndpoint) *krm.VertexAIIndexEndpointSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIIndexEndpointSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Etag
	out.Labels = in.Labels
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_FromProto(mapCtx, in.GetPrivateServiceConnectConfig())
	out.PublicEndpointEnabled = direct.LazyPtr(in.GetPublicEndpointEnabled())
	out.EncryptionSpec = EncryptionSpec_v1_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}

func VertexAIIndexEndpointSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIIndexEndpointSpec) *pb.IndexEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.IndexEndpoint{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Etag
	out.Labels = in.Labels
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	// MISSING: EnablePrivateServiceConnect
	out.PrivateServiceConnectConfig = PrivateServiceConnectConfig_ToProto(mapCtx, in.PrivateServiceConnectConfig)
	out.PublicEndpointEnabled = direct.ValueOf(in.PublicEndpointEnabled)
	out.EncryptionSpec = EncryptionSpec_v1_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}

func EncryptionSpec_v1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{
		KMSKeyRef: &refsv1beta1.KMSCryptoKeyRef{
			External: in.KmsKeyName,
		},
	}
	return out
}

func EncryptionSpec_v1_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = in.KMSKeyRef.External
	}
	return out
}
