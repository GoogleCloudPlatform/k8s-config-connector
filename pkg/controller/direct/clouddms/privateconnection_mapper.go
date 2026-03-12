// Copyright 2025 Google LLC
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
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	status "google.golang.org/genproto/googleapis/rpc/status"
)

func CloudDMSPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.CloudDMSPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSPrivateConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = CloudDMSPrivateConnectionStatus_FromProto(mapCtx, in.Error)
	return out
}
func CloudDMSPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = CloudDMSPrivateConnectionStatus_ToProto(mapCtx, in.Error)
	return out
}
func CloudDMSPrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.CloudDMSPrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSPrivateConnectionSpec{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.VpcPeeringConfig = VpcPeeringConfig_FromProto(mapCtx, in.GetVpcPeeringConfig())
	return out
}
func CloudDMSPrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSPrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := VpcPeeringConfig_ToProto(mapCtx, in.VpcPeeringConfig); oneof != nil {
		out.Connectivity = &pb.PrivateConnection_VpcPeeringConfig{VpcPeeringConfig: oneof}
	}
	return out
}
func VpcPeeringConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConfig) *krm.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcPeeringConfig{}
	if in.GetVpcName() != "" {
		out.VpcNameRef = &computev1beta1.ComputeNetworkRef{External: in.GetVpcName()}
	}
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	return out
}
func VpcPeeringConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcPeeringConfig) *pb.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConfig{}
	if in.VpcNameRef != nil {
		out.VpcName = in.VpcNameRef.External
	}
	out.Subnet = direct.ValueOf(in.Subnet)
	return out
}

func CloudDMSPrivateConnectionStatus_ToProto(mapCtx *direct.MapContext, in *krm.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

func CloudDMSPrivateConnectionStatus_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.Code)
	out.Message = direct.LazyPtr(in.Message)
	return out
}
