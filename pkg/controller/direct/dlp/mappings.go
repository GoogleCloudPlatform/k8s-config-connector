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

package dlp

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dlp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Empty_FromProto(mapCtx *direct.MapContext, in *emptypb.Empty) *krm.Empty {
	if in == nil {
		return nil
	}
	return &krm.Empty{}
}

func Empty_ToProto(mapCtx *direct.MapContext, in *krm.Empty) *emptypb.Empty {
	if in == nil {
		return nil
	}
	return &emptypb.Empty{}
}

func string_FromProto(mapCtx *direct.MapContext, in *timestamppb.Timestamp) *string {
	return direct.StringTimestamp_FromProto(mapCtx, in)
}

func string_ToProto(mapCtx *direct.MapContext, in *string) *timestamppb.Timestamp {
	return direct.StringTimestamp_ToProto(mapCtx, in)
}
