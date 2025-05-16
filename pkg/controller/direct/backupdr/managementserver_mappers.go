// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ManagementURIObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ManagementURI) *krm.ManagementURIObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagementURIObservedState{}
	out.WebUI = direct.LazyPtr(in.GetWebUi())
	out.API = direct.LazyPtr(in.GetApi())
	return out
}
func ManagementURIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagementURIObservedState) *pb.ManagementURI {
	if in == nil {
		return nil
	}
	out := &pb.ManagementURI{}
	out.WebUi = direct.ValueOf(in.WebUI)
	out.Api = direct.ValueOf(in.API)
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkforceIdentityBasedOAuth2ClientID) *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkforceIdentityBasedOAuth2ClientIDObservedState{}
	out.FirstPartyOAuth2ClientID = direct.LazyPtr(in.GetFirstPartyOauth2ClientId())
	out.ThirdPartyOAuth2ClientID = direct.LazyPtr(in.GetThirdPartyOauth2ClientId())
	return out
}
func WorkforceIdentityBasedOAuth2ClientIDObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkforceIdentityBasedOAuth2ClientIDObservedState) *pb.WorkforceIdentityBasedOAuth2ClientID {
	if in == nil {
		return nil
	}
	out := &pb.WorkforceIdentityBasedOAuth2ClientID{}
	out.FirstPartyOauth2ClientId = direct.ValueOf(in.FirstPartyOAuth2ClientID)
	out.ThirdPartyOauth2ClientId = direct.ValueOf(in.ThirdPartyOAuth2ClientID)
	return out
}
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
	out.OAuth2ClientID = direct.LazyPtr(in.GetOauth2ClientId())
	out.WorkforceIdentityBasedOAuth2ClientID = WorkforceIdentityBasedOAuth2ClientIDObservedState_FromProto(mapCtx, in.GetWorkforceIdentityBasedOauth2ClientId())
	out.BAProxyURIs = in.BaProxyUri
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
	out.Oauth2ClientId = direct.ValueOf(in.OAuth2ClientID)
	out.WorkforceIdentityBasedOauth2ClientId = WorkforceIdentityBasedOAuth2ClientIDObservedState_ToProto(mapCtx, in.WorkforceIdentityBasedOAuth2ClientID)
	out.BaProxyUri = in.BAProxyURIs
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.PeeringMode = direct.Enum_FromProto(mapCtx, in.GetPeeringMode())
	return out
}
