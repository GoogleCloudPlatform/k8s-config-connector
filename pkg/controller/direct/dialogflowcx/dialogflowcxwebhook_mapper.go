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

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DialogflowCXWebhookSpec_FromProto(mapCtx *direct.MapContext, in *pb.Webhook) *krm.DialogflowCXWebhookSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowCXWebhookSpec{}
	out.DisplayName = in.GetDisplayName()
	out.GenericWebService = GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.ServiceDirectory = ServiceDirectoryConfig_FromProto(mapCtx, in.GetServiceDirectory())
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}

func DialogflowCXWebhookSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXWebhookSpec) *pb.Webhook {
	if in == nil {
		return nil
	}
	out := &pb.Webhook{}
	out.DisplayName = in.DisplayName
	if oneof := GenericWebService_ToProto(mapCtx, in.GenericWebService); oneof != nil {
		out.Webhook = &pb.Webhook_GenericWebService_{GenericWebService: oneof}
	}
	if oneof := ServiceDirectoryConfig_ToProto(mapCtx, in.ServiceDirectory); oneof != nil {
		out.Webhook = &pb.Webhook_ServiceDirectory{ServiceDirectory: oneof}
	}
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}

func GenericWebService_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_GenericWebService) *krm.GenericWebService {
	if in == nil {
		return nil
	}
	out := &krm.GenericWebService{}
	out.URI = in.GetUri()
	out.RequestHeaders = in.RequestHeaders
	if len(in.AllowedCaCerts) > 0 {
		out.AllowedCACerts = make([]string, len(in.AllowedCaCerts))
		for i, v := range in.AllowedCaCerts {
			out.AllowedCACerts[i] = string(v)
		}
	}
	return out
}

func GenericWebService_ToProto(mapCtx *direct.MapContext, in *krm.GenericWebService) *pb.Webhook_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_GenericWebService{}
	out.Uri = in.URI
	out.RequestHeaders = in.RequestHeaders
	if len(in.AllowedCACerts) > 0 {
		out.AllowedCaCerts = make([][]byte, len(in.AllowedCACerts))
		for i, v := range in.AllowedCACerts {
			out.AllowedCaCerts[i] = []byte(v)
		}
	}
	return out
}

func ServiceDirectoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Webhook_ServiceDirectoryConfig) *krm.ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.ServiceDirectoryConfig{}
	out.GenericWebService = GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.Service = in.GetService()
	return out
}

func ServiceDirectoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.ServiceDirectoryConfig) *pb.Webhook_ServiceDirectoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Webhook_ServiceDirectoryConfig{}
	out.GenericWebService = GenericWebService_ToProto(mapCtx, in.GenericWebService)
	out.Service = in.Service
	return out
}

func DialogflowCXWebhookObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Webhook) *krm.DialogflowCXWebhookObservedState {
	if in == nil {
		return nil
	}
	return &krm.DialogflowCXWebhookObservedState{}
}

func DialogflowCXWebhookObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowCXWebhookObservedState) *pb.Webhook {
	if in == nil {
		return nil
	}
	return &pb.Webhook{}
}
