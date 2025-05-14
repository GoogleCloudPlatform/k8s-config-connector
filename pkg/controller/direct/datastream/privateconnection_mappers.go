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

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DatastreamPrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.DatastreamPrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamPrivateConnectionSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.VPCPeeringConfig = VpcPeeringConfig_FromProto(mapCtx, in.GetVpcPeeringConfig())
	return out
}
func DatastreamPrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamPrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.VpcPeeringConfig = VpcPeeringConfig_ToProto(mapCtx, in.VPCPeeringConfig)
	return out
}
func Error_FromProto(mapCtx *direct.MapContext, in *pb.Error) *krm.Error {
	if in == nil {
		return nil
	}
	out := &krm.Error{}
	out.Reason = direct.LazyPtr(in.GetReason())
	out.ErrorUUID = direct.LazyPtr(in.GetErrorUuid())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.ErrorTime = direct.StringTimestamp_FromProto(mapCtx, in.GetErrorTime())
	out.Details = in.Details
	return out
}
func Error_ToProto(mapCtx *direct.MapContext, in *krm.Error) *pb.Error {
	if in == nil {
		return nil
	}
	out := &pb.Error{}
	out.Reason = direct.ValueOf(in.Reason)
	out.ErrorUuid = direct.ValueOf(in.ErrorUUID)
	out.Message = direct.ValueOf(in.Message)
	out.ErrorTime = direct.StringTimestamp_ToProto(mapCtx, in.ErrorTime)
	out.Details = in.Details
	return out
}
func VpcPeeringConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConfig) *krm.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcPeeringConfig{}
	if in.GetVpc() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetVpc()}
	}
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	return out
}
func VpcPeeringConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcPeeringConfig) *pb.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConfig{}
	if in.NetworkRef != nil {
		out.Vpc = in.NetworkRef.External
	}
	out.Subnet = direct.ValueOf(in.Subnet)
	return out
}
