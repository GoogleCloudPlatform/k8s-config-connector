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

package bigtable

import (
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Status_v1alpha1_FromProto(mapCtx *direct.MapContext, in *status.Status) *common.Status {
	if in == nil {
		return nil
	}
	out := &common.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func Status_v1alpha1_ToProto(mapCtx *direct.MapContext, in *common.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

func Any_v1alpha1_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *common.DeprecatedAny {
	if in == nil {
		return nil
	}
	out := &common.DeprecatedAny{}
	out.TypeURL = direct.LazyPtr(in.GetTypeUrl())
	out.Value = in.GetValue()
	return out
}

func Any_v1alpha1_ToProto(mapCtx *direct.MapContext, in *common.DeprecatedAny) *anypb.Any {
	if in == nil {
		return nil
	}
	out := &anypb.Any{}
	out.TypeUrl = direct.ValueOf(in.TypeURL)
	out.Value = in.Value
	return out
}
