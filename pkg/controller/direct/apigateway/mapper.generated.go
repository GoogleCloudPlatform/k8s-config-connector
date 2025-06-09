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

// +generated:mapper
// proto.service: google.cloud.apigateway.v1
// krm.group: apigateway.cnrm.cloud.google.com
// krm.version: v1beta1

package apigateway

import (
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIGatewayAPIObservedState_ToProto(mapCtx *direct.MapContext, in *krm.APIGatewayAPIObservedState) *pb.Api {
	if in == nil {
		return nil
	}
	out := &pb.Api{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Api_State](mapCtx, in.State)
	return out
}
