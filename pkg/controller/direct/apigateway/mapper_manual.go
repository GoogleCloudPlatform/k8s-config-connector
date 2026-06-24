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

package apigateway

import (
	"encoding/base64"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krmapigatewayv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIGatewayAPIConfig_File_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_File) *krmapigatewayv1alpha1.APIGatewayAPIConfig_File {
	if in == nil {
		return nil
	}
	out := &krmapigatewayv1alpha1.APIGatewayAPIConfig_File{}
	out.Path = direct.LazyPtr(in.GetPath())

	if in.GetContents() != nil {
		s := base64.StdEncoding.EncodeToString(in.GetContents())
		out.Contents = &s
	}
	return out
}

func APIGatewayAPIConfig_File_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmapigatewayv1alpha1.APIGatewayAPIConfig_File) *pb.ApiConfig_File {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_File{}
	out.Path = direct.ValueOf(in.Path)
	if in.Contents != nil {
		b, err := base64.StdEncoding.DecodeString(*in.Contents)
		if err != nil {
			mapCtx.Errorf("error decoding base64 contents for file %q: %v", direct.ValueOf(in.Path), err)
		} else {
			out.Contents = b
		}
	}
	return out
}
