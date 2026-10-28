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
	pb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/beyondcorp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AppConnectionApplicationEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection_ApplicationEndpoint) *krm.AppConnectionApplicationEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectionApplicationEndpoint{}
	out.Host = in.GetHost()
	out.Port = int(in.GetPort())
	return out
}

func AppConnectionApplicationEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectionApplicationEndpoint) *pb.AppConnection_ApplicationEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection_ApplicationEndpoint{}
	out.Host = in.Host
	out.Port = int32(in.Port)
	return out
}

func AppConnectionGateway_FromProto(mapCtx *direct.MapContext, in *pb.AppConnection_Gateway) *krm.AppConnectionGateway {
	if in == nil {
		return nil
	}
	out := &krm.AppConnectionGateway{}
	out.AppGateway = in.GetAppGateway()
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.URI = direct.LazyPtr(in.GetUri())
	out.IngressPort = direct.LazyPtr(int(in.GetIngressPort()))
	return out
}

func AppConnectionGateway_ToProto(mapCtx *direct.MapContext, in *krm.AppConnectionGateway) *pb.AppConnection_Gateway {
	if in == nil {
		return nil
	}
	out := &pb.AppConnection_Gateway{}
	out.AppGateway = in.AppGateway
	out.Type = direct.Enum_ToProto[pb.AppConnection_Gateway_Type](mapCtx, in.Type)
	out.Uri = direct.ValueOf(in.URI)
	out.IngressPort = int32(direct.ValueOf(in.IngressPort))
	return out
}
