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
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRManagementServerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupDRManagementServerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRManagementServerObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ManagementURI = ManagementURIObservedState_FromProto(mapCtx, in.GetManagementUri())
	out.WorkforceIdentityBasedManagementURI = WorkforceIdentityBasedManagementURIObservedState_FromProto(mapCtx, in.GetWorkforceIdentityBasedManagementUri())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.OAUTH2ClientID = direct.LazyPtr(in.GetOauth2ClientId())
	out.WorkforceIdentityBasedOAUTH2ClientID = WorkforceIdentityBasedOAuth2ClientIDObservedState_FromProto(mapCtx, in.GetWorkforceIdentityBasedOauth2ClientId())
	out.BaProxyURI = in.BaProxyUri
	out.SatisfiesPzs = direct.BoolValue_FromProto(mapCtx, in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func BackupDRManagementServerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRManagementServerObservedState) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ManagementUri = ManagementURIObservedState_ToProto(mapCtx, in.ManagementURI)
	out.WorkforceIdentityBasedManagementUri = WorkforceIdentityBasedManagementURIObservedState_ToProto(mapCtx, in.WorkforceIdentityBasedManagementURI)
	out.State = direct.Enum_ToProto[pb.ManagementServer_InstanceState](mapCtx, in.State)
	out.Oauth2ClientId = direct.ValueOf(in.OAUTH2ClientID)
	out.WorkforceIdentityBasedOauth2ClientId = WorkforceIdentityBasedOAuth2ClientIDObservedState_ToProto(mapCtx, in.WorkforceIdentityBasedOAUTH2ClientID)
	out.BaProxyUri = in.BaProxyURI
	out.SatisfiesPzs = direct.BoolValue_ToProto(mapCtx, in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func BackupDRManagementServerSpec_FromProto(mapCtx *direct.MapContext, in *pb.ManagementServer) *krm.BackupDRManagementServerSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRManagementServerSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Networks = direct.Slice_FromProto(mapCtx, in.Networks, NetworkConfig_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func BackupDRManagementServerSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRManagementServerSpec) *pb.ManagementServer {
	if in == nil {
		return nil
	}
	out := &pb.ManagementServer{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Type = direct.Enum_ToProto[pb.ManagementServer_InstanceType](mapCtx, in.Type)
	out.Networks = direct.Slice_ToProto(mapCtx, in.Networks, NetworkConfig_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func ManagementURI_FromProto(mapCtx *direct.MapContext, in *pb.ManagementURI) *krm.ManagementURI {
	if in == nil {
		return nil
	}
	out := &krm.ManagementURI{}
	// MISSING: WebUi
	// MISSING: API
	return out
}
func ManagementURI_ToProto(mapCtx *direct.MapContext, in *krm.ManagementURI) *pb.ManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.ManagementURI{}
	// MISSING: WebUi
	// MISSING: API
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
	// MISSING: FirstPartyOAUTH2ClientID
	// MISSING: ThirdPartyOAUTH2ClientID
	return out
}
func WorkforceIdentityBasedOAuth2ClientID_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedOAuth2ClientID) *pb.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedOAuth2ClientID{}
	// MISSING: FirstPartyOAUTH2ClientID
	// MISSING: ThirdPartyOAUTH2ClientID
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedOAuth2ClientID) *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedOAuth2ClientIDObservedState{}
	out.FirstPartyOAUTH2ClientID = direct.LazyPtr(in.GetFirstPartyOauth2ClientId())
	out.ThirdPartyOAUTH2ClientID = direct.LazyPtr(in.GetThirdPartyOauth2ClientId())
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState) *pb.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedOAuth2ClientID{}
	out.FirstPartyOauth2ClientId = direct.ValueOf(in.FirstPartyOAUTH2ClientID)
	out.ThirdPartyOauth2ClientId = direct.ValueOf(in.ThirdPartyOAUTH2ClientID)
	return out
}
