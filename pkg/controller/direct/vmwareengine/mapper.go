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

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	krmvmwareenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	krmvmwareenginev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func NetworkPolicy_NetworkServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krmvmwareenginev1alpha1.NetworkPolicy_NetworkServiceObservedState {
	return NetworkPolicy_NetworkServiceObservedState_v1alpha1_FromProto(mapCtx, in)
}

func NetworkPolicy_NetworkService_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPolicy_NetworkService) *krmvmwareenginev1alpha1.NetworkPolicy_NetworkService {
	return NetworkPolicy_NetworkService_v1alpha1_FromProto(mapCtx, in)
}

func NetworkPolicy_NetworkService_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.NetworkPolicy_NetworkService) *pb.NetworkPolicy_NetworkService {
	return NetworkPolicy_NetworkService_v1alpha1_ToProto(mapCtx, in)
}

func ExternalAccessRule_IPRange_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.ExternalAccessRule_IPRange) *pb.ExternalAccessRule_IpRange {
	return ExternalAccessRule_IPRange_ToProto(mapCtx, in)
}

func VMwareEngineNetworkPeeringSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkPeering) *krmvmwareenginev1alpha1.VMwareEngineNetworkPeeringSpec {
	return VMwareEngineNetworkPeeringSpec_FromProto(mapCtx, in)
}



func VMwareEngineNetworkPeeringSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.VMwareEngineNetworkPeeringSpec) *pb.NetworkPeering {
	return VMwareEngineNetworkPeeringSpec_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func Nsx_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.Nsx) *pb.Nsx {
	return Nsx_v1alpha1_ToProto(mapCtx, in)
}

func Nsx_FromProto(mapCtx *direct.MapContext, in *pb.Nsx) *krmvmwareenginev1alpha1.Nsx {
	return Nsx_v1alpha1_FromProto(mapCtx, in)
}

func Hcx_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.Hcx) *pb.Hcx {
	return Hcx_v1alpha1_ToProto(mapCtx, in)
}

func Vcenter_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.Vcenter) *pb.Vcenter {
	return Vcenter_v1alpha1_ToProto(mapCtx, in)
}

func NetworkPolicy_NetworkServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.NetworkPolicy_NetworkServiceObservedState) *pb.NetworkPolicy_NetworkService {
	return NetworkPolicy_NetworkServiceObservedState_v1alpha1_ToProto(mapCtx, in)
}

func Hcx_FromProto(mapCtx *direct.MapContext, in *pb.Hcx) *krmvmwareenginev1alpha1.Hcx {
	return Hcx_v1alpha1_FromProto(mapCtx, in)
}

func Vcenter_FromProto(mapCtx *direct.MapContext, in *pb.Vcenter) *krmvmwareenginev1alpha1.Vcenter {
	return Vcenter_v1alpha1_FromProto(mapCtx, in)
}

func NodeTypeConfig_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.NodeTypeConfig) *pb.NodeTypeConfig {
	return NodeTypeConfig_v1alpha1_ToProto(mapCtx, in)
}

func NodeTypeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeTypeConfig) *krmvmwareenginev1alpha1.NodeTypeConfig {
	return NodeTypeConfig_v1alpha1_FromProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func VMwareEngineNetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork) *krmvmwareenginev1alpha1.VMwareEngineNetworkSpec {
	return VMwareEngineNetworkSpec_v1alpha1_FromProto(mapCtx, in)
}

func VMwareEngineExternalAccessRuleSpec_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.VMwareEngineExternalAccessRuleSpec) *pb.ExternalAccessRule {
	return VMwareEngineExternalAccessRuleSpec_v1alpha1_ToProto(mapCtx, in)
}

func VMwareEngineExternalAddressSpec_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1beta1.VMwareEngineExternalAddressSpec) *pb.ExternalAddress {
	return VMwareEngineExternalAddressSpec_v1beta1_ToProto(mapCtx, in)
}

func VMwareEngineExternalAddressSpec_FromProto(mapCtx *direct.MapContext, in *pb.ExternalAddress) *krmvmwareenginev1beta1.VMwareEngineExternalAddressSpec {
	return VMwareEngineExternalAddressSpec_v1beta1_FromProto(mapCtx, in)
}

func VMwareEngineNetworkSpec_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.VMwareEngineNetworkSpec) *pb.VmwareEngineNetwork {
	return VMwareEngineNetworkSpec_v1alpha1_ToProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func VmwareEngineNetwork_VPCNetworkObservedState_ToProto(mapCtx *direct.MapContext, in *krmvmwareenginev1alpha1.VmwareEngineNetwork_VPCNetworkObservedState) *pb.VmwareEngineNetwork_VpcNetwork {
	return VmwareEngineNetwork_VPCNetworkObservedState_v1alpha1_ToProto(mapCtx, in)
}



func VmwareEngineNetwork_VPCNetworkObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VmwareEngineNetwork_VpcNetwork) *krmvmwareenginev1alpha1.VmwareEngineNetwork_VPCNetworkObservedState {
	return VmwareEngineNetwork_VPCNetworkObservedState_v1alpha1_FromProto(mapCtx, in)
}


