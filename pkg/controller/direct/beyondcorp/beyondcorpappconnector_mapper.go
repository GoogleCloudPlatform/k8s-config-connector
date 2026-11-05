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

package beyondcorp

import (
	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BeyondCorpAppConnectorSpec_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector) *krm.BeyondCorpAppConnectorSpec {
	if in == nil {
		return nil
	}
	out := &krm.BeyondCorpAppConnectorSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.PrincipalInfo = AppConnectorPrincipalInfo_FromProto(mapCtx, in.GetPrincipalInfo())
	return out
}

func BeyondCorpAppConnectorSpec_ToProto(mapCtx *direct.MapContext, in *krm.BeyondCorpAppConnectorSpec) *pb.AppConnector {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.PrincipalInfo = AppConnectorPrincipalInfo_ToProto(mapCtx, in.PrincipalInfo)
	return out
}

func AppConnectorPrincipalInfo_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector_PrincipalInfo) *krm.AppConnectorPrincipalInfo {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectorPrincipalInfo{}
	out.ServiceAccount = AppConnectorPrincipalInfoServiceAccount_FromProto(mapCtx, in.GetServiceAccount())
	return out
}

func AppConnectorPrincipalInfo_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectorPrincipalInfo) *pb.AppConnector_PrincipalInfo {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector_PrincipalInfo{}
	if oneof := AppConnectorPrincipalInfoServiceAccount_ToProto(mapCtx, in.ServiceAccount); oneof != nil {
		out.Type = &pb.AppConnector_PrincipalInfo_ServiceAccount_{ServiceAccount: oneof}
	}
	return out
}

func AppConnectorPrincipalInfoServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.AppConnector_PrincipalInfo_ServiceAccount) *krm.AppConnectorPrincipalInfoServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectorPrincipalInfoServiceAccount{}
	out.Email = in.GetEmail()
	return out
}

func AppConnectorPrincipalInfoServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectorPrincipalInfoServiceAccount) *pb.AppConnector_PrincipalInfo_ServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.AppConnector_PrincipalInfo_ServiceAccount{}
	out.Email = in.Email
	return out
}
