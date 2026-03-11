// Copyright 2024 Google LLC
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

package networkconnectivity

import (
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"

	krmnetworkconnectivityv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetworkConnectivityServiceConnectionPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicySpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.Network != "" {
		out.Network = &krmcomputev1beta1.ComputeNetworkRef{External: in.Network}
	}
	out.PscConfig = PscConfig_v1alpha1_FromProto(mapCtx, in.GetPscConfig())
	out.ServiceClass = direct.LazyPtr(in.GetServiceClass())
	return out
}

func NetworkConnectivityServiceConnectionPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicySpec) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.Description = direct.ValueOf(in.Description)
	if in.Network != nil {
		out.Network = in.Network.External
	}
	out.PscConfig = PscConfig_v1alpha1_ToProto(mapCtx, in.PscConfig)
	out.ServiceClass = direct.ValueOf(in.ServiceClass)
	return out
}

func NetworkConnectivityServiceConnectionPolicyObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ServiceConnectionPolicy) *krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicyObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Infrastructure = direct.LazyPtr(in.GetInfrastructure())
	out.PscConnections = direct.Slice_FromProto(mapCtx, in.GetPscConnections(), pscConnection_v1alpha1_FromProto)
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func NetworkConnectivityServiceConnectionPolicyObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.NetworkConnectivityServiceConnectionPolicyObservedState) *pb.ServiceConnectionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.ServiceConnectionPolicy{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.Infrastructure = direct.ValueOf(in.Infrastructure)
	out.PscConnections = direct.Slice_ToProto(mapCtx, in.PscConnections, pscConnection_v1alpha1_ToProto)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func pscConnection_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PscConnection) *krmnetworkconnectivityv1alpha1.PscConnection {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.PscConnection{}
	out.ConsumerAddress = direct.LazyPtr(in.GetConsumerAddress())
	out.ConsumerForwardingRule = direct.LazyPtr(in.GetConsumerForwardingRule())
	out.ConsumerTargetProject = direct.LazyPtr(in.GetConsumerTargetProject())
	out.Error = GoogleRpcStatus_v1alpha1_FromProto(mapCtx, in.GetError())
	out.ErrorInfo = GoogleRpcErrorInfo_v1alpha1_FromProto(mapCtx, in.GetErrorInfo())
	out.ErrorType = direct.LazyPtr(in.GetErrorType())
	out.GceOperation = direct.LazyPtr(in.GetGceOperation())
	out.ProducerInstanceID = direct.LazyPtr(in.GetProducerInstanceId())
	out.PscConnectionID = direct.LazyPtr(in.GetPscConnectionId())
	out.SelectedSubnetwork = direct.LazyPtr(in.GetSelectedSubnetwork())
	out.State = direct.LazyPtr(in.GetState())
	return out
}

func pscConnection_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.PscConnection) *pb.PscConnection {
	if in == nil {
		return nil
	}
	out := &pb.PscConnection{}
	out.ConsumerAddress = direct.ValueOf(in.ConsumerAddress)
	out.ConsumerForwardingRule = direct.ValueOf(in.ConsumerForwardingRule)
	out.ConsumerTargetProject = direct.ValueOf(in.ConsumerTargetProject)
	out.Error = GoogleRpcStatus_v1alpha1_ToProto(mapCtx, in.Error)
	out.ErrorInfo = GoogleRpcErrorInfo_v1alpha1_ToProto(mapCtx, in.ErrorInfo)
	out.ErrorType = direct.ValueOf(in.ErrorType)
	out.GceOperation = direct.ValueOf(in.GceOperation)
	out.ProducerInstanceId = direct.ValueOf(in.ProducerInstanceID)
	out.PscConnectionId = direct.ValueOf(in.PscConnectionID)
	out.SelectedSubnetwork = direct.ValueOf(in.SelectedSubnetwork)
	out.State = direct.ValueOf(in.State)
	return out
}

func Migration_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Migration) *krm.Migration {
	if in == nil {
		return nil
	}
	out := &krm.Migration{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Target = direct.LazyPtr(in.GetTarget())
	return out
}
func Migration_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Migration) *pb.Migration {
	if in == nil {
		return nil
	}
	out := &pb.Migration{}
	out.Source = direct.ValueOf(in.Source)
	out.Target = direct.ValueOf(in.Target)
	return out
}
func NetworkConnectivityInternalRangeObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.NetworkConnectivityInternalRangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityInternalRangeObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Users = in.Users
	return out
}
func NetworkConnectivityInternalRangeObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityInternalRangeObservedState) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Users = in.Users
	return out
}
func NetworkConnectivityInternalRangeSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.InternalRange) *krm.NetworkConnectivityInternalRangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConnectivityInternalRangeSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.IPCIDRRange = direct.LazyPtr(in.GetIpCidrRange())
	out.Labels = in.Labels
	out.Migration = Migration_v1beta1_FromProto(mapCtx, in.GetMigration())
	if in.GetNetwork() != "" {
		out.NetworkRef = &krmcomputev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.Overlaps = in.Overlaps
	out.Peering = direct.LazyPtr(in.GetPeering())
	out.PrefixLength = direct.LazyPtr(in.GetPrefixLength())
	out.TargetCIDRRange = in.TargetCidrRange
	out.Usage = direct.LazyPtr(in.GetUsage())
	return out
}
func NetworkConnectivityInternalRangeSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConnectivityInternalRangeSpec) *pb.InternalRange {
	if in == nil {
		return nil
	}
	out := &pb.InternalRange{}
	out.Description = direct.ValueOf(in.Description)
	out.IpCidrRange = direct.ValueOf(in.IPCIDRRange)
	out.Labels = in.Labels
	out.Migration = Migration_v1beta1_ToProto(mapCtx, in.Migration)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.Overlaps = in.Overlaps
	out.Peering = direct.ValueOf(in.Peering)
	out.PrefixLength = direct.ValueOf(in.PrefixLength)
	out.TargetCidrRange = in.TargetCIDRRange
	out.Usage = direct.ValueOf(in.Usage)
	return out
}

func PscConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PscConfig) *krmnetworkconnectivityv1alpha1.PscConfig {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.PscConfig{}
	out.Limit = direct.LazyPtr(in.GetLimit())
	out.ProducerInstanceLocation = direct.LazyPtr(in.GetProducerInstanceLocation())
	out.Subnetworks = PscConfig_Subnetworks_FromProto(mapCtx, in.GetSubnetworks())
	return out
}
func PscConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.PscConfig) *pb.PscConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscConfig{}
	out.Limit = direct.ValueOf(in.Limit)
	out.ProducerInstanceLocation = direct.ValueOf(in.ProducerInstanceLocation)
	out.Subnetworks = PscConfig_Subnetworks_ToProto(mapCtx, in.Subnetworks)
	return out
}

func PscConfig_Subnetworks_FromProto(mapCtx *direct.MapContext, in []string) []refs.ComputeSubnetworkRef {
	if in == nil {
		return nil
	}
	var out []refs.ComputeSubnetworkRef
	for _, s := range in {
		out = append(out, refs.ComputeSubnetworkRef{External: s})
	}
	return out
}
func PscConfig_Subnetworks_ToProto(mapCtx *direct.MapContext, in []refs.ComputeSubnetworkRef) []string {
	if in == nil {
		return nil
	}
	var out []string
	for _, ref := range in {
		out = append(out, ref.External)
	}
	return out
}

func Policy_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krmnetworkconnectivityv1alpha1.Policy {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.Policy{}
	out.AuditConfigs = direct.Slice_FromProto(mapCtx, in.AuditConfigs, AuditConfig_v1alpha1_FromProto)
	out.Bindings = direct.Slice_FromProto(mapCtx, in.Bindings, Binding_v1alpha1_FromProto)
	out.Etag = in.GetEtag()
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func Policy_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.AuditConfigs = direct.Slice_ToProto(mapCtx, in.AuditConfigs, AuditConfig_v1alpha1_ToProto)
	out.Bindings = direct.Slice_ToProto(mapCtx, in.Bindings, Binding_v1alpha1_ToProto)
	out.Etag = in.Etag
	out.Version = direct.ValueOf(in.Version)
	return out
}

func AuditConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuditConfig) *krmnetworkconnectivityv1alpha1.AuditConfig {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.AuditConfig{}
	out.AuditLogConfigs = direct.Slice_FromProto(mapCtx, in.AuditLogConfigs, AuditLogConfig_v1alpha1_FromProto)
	out.Service = direct.LazyPtr(in.GetService())
	return out
}

func AuditConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.AuditConfig) *pb.AuditConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuditConfig{}
	out.AuditLogConfigs = direct.Slice_ToProto(mapCtx, in.AuditLogConfigs, AuditLogConfig_v1alpha1_ToProto)
	out.Service = direct.ValueOf(in.Service)
	return out
}

func AuditLogConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.AuditLogConfig) *krmnetworkconnectivityv1alpha1.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.AuditLogConfig{}
	out.ExemptedMembers = in.ExemptedMembers
	out.LogType = direct.LazyPtr(in.GetLogType())
	return out
}

func AuditLogConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.AuditLogConfig) *pb.AuditLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.AuditLogConfig{}
	out.ExemptedMembers = in.ExemptedMembers
	out.LogType = direct.ValueOf(in.LogType)
	return out
}

func Binding_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Binding) *krmnetworkconnectivityv1alpha1.Binding {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.Binding{}
	out.Condition = Expr_v1alpha1_FromProto(mapCtx, in.GetCondition())
	out.Members = in.Members
	out.Role = direct.LazyPtr(in.GetRole())
	return out
}

func Binding_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.Binding) *pb.Binding {
	if in == nil {
		return nil
	}
	out := &pb.Binding{}
	out.Condition = Expr_v1alpha1_ToProto(mapCtx, in.Condition)
	out.Members = in.Members
	out.Role = direct.ValueOf(in.Role)
	return out
}

func Expr_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Expr) *krmnetworkconnectivityv1alpha1.Expr {
	if in == nil {
		return nil
	}
	out := &krmnetworkconnectivityv1alpha1.Expr{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Expression = direct.LazyPtr(in.GetExpression())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.Title = direct.LazyPtr(in.GetTitle())
	return out
}

func Expr_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmnetworkconnectivityv1alpha1.Expr) *pb.Expr {
	if in == nil {
		return nil
	}
	out := &pb.Expr{}
	out.Description = direct.ValueOf(in.Description)
	out.Expression = direct.ValueOf(in.Expression)
	out.Location = direct.ValueOf(in.Location)
	out.Title = direct.ValueOf(in.Title)
	return out
}
