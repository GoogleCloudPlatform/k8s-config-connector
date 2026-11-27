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

package designcenter

import (
	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DesignCenterConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Connection) *krm.DesignCenterConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DesignCenterConnectionSpec{}
	if in.GetDestinationComponentUri() != "" {
		out.DestinationComponentRef = &krm.DesignCenterComponentRef{External: in.GetDestinationComponentUri()}
	}
	out.SourceComponentParameters = direct.Slice_FromProto(mapCtx, in.SourceComponentParameters, Parameter_FromProto)
	out.DestinationComponentParameters = direct.Slice_FromProto(mapCtx, in.DestinationComponentParameters, Parameter_FromProto)
	return out
}

func DesignCenterConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DesignCenterConnectionSpec) *pb.Connection {
	if in == nil {
		return nil
	}
	out := &pb.Connection{}
	if in.DestinationComponentRef != nil {
		out.DestinationComponentUri = in.DestinationComponentRef.External
	}
	out.SourceComponentParameters = direct.Slice_ToProto(mapCtx, in.SourceComponentParameters, Parameter_ToProto)
	out.DestinationComponentParameters = direct.Slice_ToProto(mapCtx, in.DestinationComponentParameters, Parameter_ToProto)
	return out
}
