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

package backupdr

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupdrManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupdrManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupdrManagementServerSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func BackupdrManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupdrManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Type
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	// MISSING: Networks
	// MISSING: Etag
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ManagementServer_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.ManagementServer {
	if in == nil {
		return nil
	}
	out := &krm.ManagementServer{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ManagementServer_ToProto(mapCtx *direct.MapContext, in *krm.ManagementServer) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Type = direct.Enum_ToProto[pb.ManagementServer_InstanceType](mapCtx, in.Type)
	// MISSING: ManagementURI
	// MISSING: WorkforceIdentityBasedManagementURI
	// MISSING: State
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Oauth2ClientID
	// MISSING: WorkforceIdentityBasedOauth2ClientID
	// MISSING: BaProxyURI
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.ManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementServerObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Type
	out.ManagementURI = ManagementURI_FromProto(mapCtx, in.GetManagementUri())
	out.WorkforceIdentityBasedManagementURI = WorkforceIdentityBasedManagementURI_FromProto(mapCtx, in.GetWorkforceIdentityBasedManagementUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Networks
	// MISSING: Etag
	out.Oauth2ClientID = direct.LazyPtr(in.GetOauth2ClientId())
	out.WorkforceIdentityBasedOauth2ClientID = WorkforceIdentityBasedOAuth2ClientID_FromProto(mapCtx, in.GetWorkforceIdentityBasedOauth2ClientId())
	out.BaProxyURI = in.BaProxyUri
	out.SatisfiesPzs = direct.BoolValue_FromProto(mapCtx, in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func ManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Type
	out.ManagementUri = ManagementURI_ToProto(mapCtx, in.ManagementURI)
	out.WorkforceIdentityBasedManagementUri = WorkforceIdentityBasedManagementURI_ToProto(mapCtx, in.WorkforceIdentityBasedManagementURI)
	out.State = direct.Enum_ToProto[pb.ManagementServer_InstanceState](mapCtx, in.State)
	// MISSING: Networks
	// MISSING: Etag
	out.Oauth2ClientId = direct.ValueOf(in.Oauth2ClientID)
	out.WorkforceIdentityBasedOauth2ClientId = WorkforceIdentityBasedOAuth2ClientID_ToProto(mapCtx, in.WorkforceIdentityBasedOauth2ClientID)
	out.BaProxyUri = in.BaProxyURI
	out.SatisfiesPzs = direct.BoolValue_ToProto(mapCtx, in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func ManagementURI_FromProto(mapCtx *direct.MapContext, in *pb.ManagementURI) *krm.ManagementURI {
	if in == nil {
		return nil
	}
	out := &krm.ManagementURI{}
	// MISSING: WebUi
	// MISSING: Api
	return out
}
func ManagementURI_ToProto(mapCtx *direct.MapContext, in *krm.ManagementURI) *pb.ManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.ManagementURI{}
	// MISSING: WebUi
	// MISSING: Api
	return out
}
func ManagementURIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementURI) *krm.ManagementURIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementURIObservedState{}
	out.WebUi = direct.LazyPtr(in.GetWebUi())
	out.Api = direct.LazyPtr(in.GetApi())
	return out
}
func ManagementURIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementURIObservedState) *pb.ManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.ManagementURI{}
	out.WebUi = direct.ValueOf(in.WebUi)
	out.Api = direct.ValueOf(in.Api)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.PeeringMode = direct.Enum_FromProto(mapCtx, in.GetPeeringMode())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.PeeringMode = direct.Enum_ToProto[pb.NetworkConfig_PeeringMode](mapCtx, in.PeeringMode)
	return out
}
func WorkforceIdentityBasedManagementURI_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedManagementURI) *krm.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedManagementURI{}
	// MISSING: FirstPartyManagementURI
	// MISSING: ThirdPartyManagementURI
	return out
}
func WorkforceIdentityBasedManagementURI_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedManagementURI) *pb.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedManagementURI{}
	// MISSING: FirstPartyManagementURI
	// MISSING: ThirdPartyManagementURI
	return out
}
func WorkforceIdentityBasedManagementURIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedManagementURI) *krm.WorkforceIdentityBasedManagementURIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedManagementURIObservedState{}
	out.FirstPartyManagementURI = direct.LazyPtr(in.GetFirstPartyManagementUri())
	out.ThirdPartyManagementURI = direct.LazyPtr(in.GetThirdPartyManagementUri())
	return out
}
func WorkforceIdentityBasedManagementURIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedManagementURIObservedState) *pb.WorkforceIdentityBasedManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedManagementURI{}
	out.FirstPartyManagementUri = direct.ValueOf(in.FirstPartyManagementURI)
	out.ThirdPartyManagementUri = direct.ValueOf(in.ThirdPartyManagementURI)
	return out
}
func WorkforceIdentityBasedOAuth2ClientID_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedOAuth2ClientID) *krm.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedOAuth2ClientID{}
	// MISSING: FirstPartyOauth2ClientID
	// MISSING: ThirdPartyOauth2ClientID
	return out
}
func WorkforceIdentityBasedOAuth2ClientID_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedOAuth2ClientID) *pb.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedOAuth2ClientID{}
	// MISSING: FirstPartyOauth2ClientID
	// MISSING: ThirdPartyOauth2ClientID
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedOAuth2ClientID) *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedOAuth2ClientIDObservedState{}
	out.FirstPartyOauth2ClientID = direct.LazyPtr(in.GetFirstPartyOauth2ClientId())
	out.ThirdPartyOauth2ClientID = direct.LazyPtr(in.GetThirdPartyOauth2ClientId())
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState) *pb.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedOAuth2ClientID{}
	out.FirstPartyOauth2ClientId = direct.ValueOf(in.FirstPartyOauth2ClientID)
	out.ThirdPartyOauth2ClientId = direct.ValueOf(in.ThirdPartyOauth2ClientID)
	return out
}
