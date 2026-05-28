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

package osconfig

import (
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FixedOrPercent_Fixed_ToProto(mapCtx *direct.MapContext, in *int64) *pb.FixedOrPercent_Fixed {
	if in == nil {
		return nil
	}
	return &pb.FixedOrPercent_Fixed{Fixed: int32(*in)}
}

func FixedOrPercent_Percent_ToProto(mapCtx *direct.MapContext, in *int64) *pb.FixedOrPercent_Percent {
	if in == nil {
		return nil
	}
	return &pb.FixedOrPercent_Percent{Percent: int32(*in)}
}

func FixedOrPercent_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &krm.FixedOrPercent{}
	out.Fixed = direct.LazyPtr(int64(in.GetFixed()))
	out.Percent = direct.LazyPtr(int64(in.GetPercent()))
	return out
}

func FixedOrPercent_ToProto(mapCtx *direct.MapContext, in *krm.FixedOrPercent) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if oneof := FixedOrPercent_Fixed_ToProto(mapCtx, in.Fixed); oneof != nil {
		out.Mode = oneof
	}
	if oneof := FixedOrPercent_Percent_ToProto(mapCtx, in.Percent); oneof != nil {
		out.Mode = oneof
	}
	return out
}
