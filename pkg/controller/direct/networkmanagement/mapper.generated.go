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
// krm.group: networkmanagement.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkmanagement.v1

package networkmanagement

import (
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AbortInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AbortInfo) *krmv1alpha1.AbortInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.AbortInfoObservedState{}
	out.Cause = direct.Enum_FromProto(mapCtx, in.GetCause())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.ProjectsMissingPermission = in.ProjectsMissingPermission
	return out
}
func AbortInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.AbortInfoObservedState) *pb.AbortInfo {
	if in == nil {
		return nil
	}
	out := &pb.AbortInfo{}
	out.Cause = direct.Enum_ToProto[pb.AbortInfo_Cause](mapCtx, in.Cause)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.ProjectsMissingPermission = in.ProjectsMissingPermission
	return out
}
func AppEngineVersionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AppEngineVersionInfo) *krmv1alpha1.AppEngineVersionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.AppEngineVersionInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Runtime = direct.LazyPtr(in.GetRuntime())
	out.Environment = direct.LazyPtr(in.GetEnvironment())
	return out
}
func AppEngineVersionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.AppEngineVersionInfoObservedState) *pb.AppEngineVersionInfo {
	if in == nil {
		return nil
	}
	out := &pb.AppEngineVersionInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Runtime = direct.ValueOf(in.Runtime)
	out.Environment = direct.ValueOf(in.Environment)
	return out
}
func CloudFunctionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunctionInfo) *krmv1alpha1.CloudFunctionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudFunctionInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	return out
}
func CloudFunctionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudFunctionInfoObservedState) *pb.CloudFunctionInfo {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunctionInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Location = direct.ValueOf(in.Location)
	out.VersionId = direct.ValueOf(in.VersionID)
	return out
}
func CloudRunRevisionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunRevisionInfo) *krmv1alpha1.CloudRunRevisionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudRunRevisionInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ServiceURI = direct.LazyPtr(in.GetServiceUri())
	return out
}
func CloudRunRevisionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudRunRevisionInfoObservedState) *pb.CloudRunRevisionInfo {
	if in == nil {
		return nil
	}
	out := &pb.CloudRunRevisionInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Location = direct.ValueOf(in.Location)
	out.ServiceUri = direct.ValueOf(in.ServiceURI)
	return out
}
func CloudSQLInstanceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudSQLInstanceInfo) *krmv1alpha1.CloudSQLInstanceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.CloudSQLInstanceInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func CloudSQLInstanceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.CloudSQLInstanceInfoObservedState) *pb.CloudSQLInstanceInfo {
	if in == nil {
		return nil
	}
	out := &pb.CloudSQLInstanceInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func DeliverInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeliverInfo) *krmv1alpha1.DeliverInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DeliverInfoObservedState{}
	out.Target = direct.Enum_FromProto(mapCtx, in.GetTarget())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.StorageBucket = direct.LazyPtr(in.GetStorageBucket())
	out.PSCGoogleAPITarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	return out
}
func DeliverInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DeliverInfoObservedState) *pb.DeliverInfo {
	if in == nil {
		return nil
	}
	out := &pb.DeliverInfo{}
	out.Target = direct.Enum_ToProto[pb.DeliverInfo_Target](mapCtx, in.Target)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.StorageBucket = direct.ValueOf(in.StorageBucket)
	out.PscGoogleApiTarget = direct.ValueOf(in.PSCGoogleAPITarget)
	return out
}
func DropInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DropInfo) *krmv1alpha1.DropInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DropInfoObservedState{}
	out.Cause = direct.Enum_FromProto(mapCtx, in.GetCause())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.DestinationIP = direct.LazyPtr(in.GetDestinationIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func DropInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DropInfoObservedState) *pb.DropInfo {
	if in == nil {
		return nil
	}
	out := &pb.DropInfo{}
	out.Cause = direct.Enum_ToProto[pb.DropInfo_Cause](mapCtx, in.Cause)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.SourceIp = direct.ValueOf(in.SourceIP)
	out.DestinationIp = direct.ValueOf(in.DestinationIP)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func EndpointInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EndpointInfo) *krmv1alpha1.EndpointInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EndpointInfoObservedState{}
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.DestinationIP = direct.LazyPtr(in.GetDestinationIp())
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.SourcePort = direct.LazyPtr(in.GetSourcePort())
	out.DestinationPort = direct.LazyPtr(in.GetDestinationPort())
	out.SourceNetworkURI = direct.LazyPtr(in.GetSourceNetworkUri())
	out.DestinationNetworkURI = direct.LazyPtr(in.GetDestinationNetworkUri())
	out.SourceAgentURI = direct.LazyPtr(in.GetSourceAgentUri())
	return out
}
func EndpointInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EndpointInfoObservedState) *pb.EndpointInfo {
	if in == nil {
		return nil
	}
	out := &pb.EndpointInfo{}
	out.SourceIp = direct.ValueOf(in.SourceIP)
	out.DestinationIp = direct.ValueOf(in.DestinationIP)
	out.Protocol = direct.ValueOf(in.Protocol)
	out.SourcePort = direct.ValueOf(in.SourcePort)
	out.DestinationPort = direct.ValueOf(in.DestinationPort)
	out.SourceNetworkUri = direct.ValueOf(in.SourceNetworkURI)
	out.DestinationNetworkUri = direct.ValueOf(in.DestinationNetworkURI)
	out.SourceAgentUri = direct.ValueOf(in.SourceAgentURI)
	return out
}
func Endpoint_AppEngineVersionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_AppEngineVersionEndpoint) *krmv1alpha1.Endpoint_AppEngineVersionEndpoint {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Endpoint_AppEngineVersionEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Endpoint_AppEngineVersionEndpoint_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Endpoint_AppEngineVersionEndpoint) *pb.Endpoint_AppEngineVersionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_AppEngineVersionEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Endpoint_CloudFunctionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_CloudFunctionEndpoint) *krmv1alpha1.Endpoint_CloudFunctionEndpoint {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Endpoint_CloudFunctionEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Endpoint_CloudFunctionEndpoint_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Endpoint_CloudFunctionEndpoint) *pb.Endpoint_CloudFunctionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_CloudFunctionEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Endpoint_CloudRunRevisionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_CloudRunRevisionEndpoint) *krmv1alpha1.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Endpoint_CloudRunRevisionEndpoint{}
	// MISSING: URI
	return out
}
func Endpoint_CloudRunRevisionEndpoint_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Endpoint_CloudRunRevisionEndpoint) *pb.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_CloudRunRevisionEndpoint{}
	// MISSING: URI
	return out
}
func FirewallInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirewallInfo) *krmv1alpha1.FirewallInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.FirewallInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Direction = direct.LazyPtr(in.GetDirection())
	out.Action = direct.LazyPtr(in.GetAction())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.TargetTags = in.TargetTags
	out.TargetServiceAccounts = in.TargetServiceAccounts
	out.Policy = direct.LazyPtr(in.GetPolicy())
	out.PolicyURI = direct.LazyPtr(in.GetPolicyUri())
	out.FirewallRuleType = direct.Enum_FromProto(mapCtx, in.GetFirewallRuleType())
	return out
}
func FirewallInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.FirewallInfoObservedState) *pb.FirewallInfo {
	if in == nil {
		return nil
	}
	out := &pb.FirewallInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Direction = direct.ValueOf(in.Direction)
	out.Action = direct.ValueOf(in.Action)
	out.Priority = direct.ValueOf(in.Priority)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.TargetTags = in.TargetTags
	out.TargetServiceAccounts = in.TargetServiceAccounts
	out.Policy = direct.ValueOf(in.Policy)
	out.PolicyUri = direct.ValueOf(in.PolicyURI)
	out.FirewallRuleType = direct.Enum_ToProto[pb.FirewallInfo_FirewallRuleType](mapCtx, in.FirewallRuleType)
	return out
}
func ForwardInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ForwardInfo) *krmv1alpha1.ForwardInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ForwardInfoObservedState{}
	out.Target = direct.Enum_FromProto(mapCtx, in.GetTarget())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	return out
}
func ForwardInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ForwardInfoObservedState) *pb.ForwardInfo {
	if in == nil {
		return nil
	}
	out := &pb.ForwardInfo{}
	out.Target = direct.Enum_ToProto[pb.ForwardInfo_Target](mapCtx, in.Target)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	return out
}
func ForwardingRuleInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRuleInfo) *krmv1alpha1.ForwardingRuleInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ForwardingRuleInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.MatchedProtocol = direct.LazyPtr(in.GetMatchedProtocol())
	out.MatchedPortRange = direct.LazyPtr(in.GetMatchedPortRange())
	out.VIP = direct.LazyPtr(in.GetVip())
	out.Target = direct.LazyPtr(in.GetTarget())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.Region = direct.LazyPtr(in.GetRegion())
	out.LoadBalancerName = direct.LazyPtr(in.GetLoadBalancerName())
	out.PSCServiceAttachmentURI = direct.LazyPtr(in.GetPscServiceAttachmentUri())
	out.PSCGoogleAPITarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	return out
}
func ForwardingRuleInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ForwardingRuleInfoObservedState) *pb.ForwardingRuleInfo {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRuleInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.MatchedProtocol = direct.ValueOf(in.MatchedProtocol)
	out.MatchedPortRange = direct.ValueOf(in.MatchedPortRange)
	out.Vip = direct.ValueOf(in.VIP)
	out.Target = direct.ValueOf(in.Target)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.Region = direct.ValueOf(in.Region)
	out.LoadBalancerName = direct.ValueOf(in.LoadBalancerName)
	out.PscServiceAttachmentUri = direct.ValueOf(in.PSCServiceAttachmentURI)
	out.PscGoogleApiTarget = direct.ValueOf(in.PSCGoogleAPITarget)
	return out
}
func GKEMasterInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GKEMasterInfo) *krmv1alpha1.GKEMasterInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GKEMasterInfoObservedState{}
	out.ClusterURI = direct.LazyPtr(in.GetClusterUri())
	out.ClusterNetworkURI = direct.LazyPtr(in.GetClusterNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.DNSEndpoint = direct.LazyPtr(in.GetDnsEndpoint())
	return out
}
func GKEMasterInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GKEMasterInfoObservedState) *pb.GKEMasterInfo {
	if in == nil {
		return nil
	}
	out := &pb.GKEMasterInfo{}
	out.ClusterUri = direct.ValueOf(in.ClusterURI)
	out.ClusterNetworkUri = direct.ValueOf(in.ClusterNetworkURI)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.DnsEndpoint = direct.ValueOf(in.DNSEndpoint)
	return out
}
func GoogleServiceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleServiceInfo) *krmv1alpha1.GoogleServiceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.GoogleServiceInfoObservedState{}
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.GoogleServiceType = direct.Enum_FromProto(mapCtx, in.GetGoogleServiceType())
	return out
}
func GoogleServiceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.GoogleServiceInfoObservedState) *pb.GoogleServiceInfo {
	if in == nil {
		return nil
	}
	out := &pb.GoogleServiceInfo{}
	out.SourceIp = direct.ValueOf(in.SourceIP)
	out.GoogleServiceType = direct.Enum_ToProto[pb.GoogleServiceInfo_GoogleServiceType](mapCtx, in.GoogleServiceType)
	return out
}
func InstanceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InstanceInfo) *krmv1alpha1.InstanceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.InstanceInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Interface = direct.LazyPtr(in.GetInterface())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.NetworkTags = in.NetworkTags
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.PSCNetworkAttachmentURI = direct.LazyPtr(in.GetPscNetworkAttachmentUri())
	return out
}
func InstanceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.InstanceInfoObservedState) *pb.InstanceInfo {
	if in == nil {
		return nil
	}
	out := &pb.InstanceInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Interface = direct.ValueOf(in.Interface)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.NetworkTags = in.NetworkTags
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.PscNetworkAttachmentUri = direct.ValueOf(in.PSCNetworkAttachmentURI)
	return out
}
func LatencyDistributionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LatencyDistribution) *krmv1alpha1.LatencyDistributionObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LatencyDistributionObservedState{}
	out.LatencyPercentiles = direct.Slice_FromProto(mapCtx, in.LatencyPercentiles, LatencyPercentileObservedState_FromProto)
	return out
}
func LatencyDistributionObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LatencyDistributionObservedState) *pb.LatencyDistribution {
	if in == nil {
		return nil
	}
	out := &pb.LatencyDistribution{}
	out.LatencyPercentiles = direct.Slice_ToProto(mapCtx, in.LatencyPercentiles, LatencyPercentileObservedState_ToProto)
	return out
}
func LatencyPercentileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LatencyPercentile) *krmv1alpha1.LatencyPercentileObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LatencyPercentileObservedState{}
	out.Percent = direct.LazyPtr(in.GetPercent())
	out.LatencyMicros = direct.LazyPtr(in.GetLatencyMicros())
	return out
}
func LatencyPercentileObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LatencyPercentileObservedState) *pb.LatencyPercentile {
	if in == nil {
		return nil
	}
	out := &pb.LatencyPercentile{}
	out.Percent = direct.ValueOf(in.Percent)
	out.LatencyMicros = direct.ValueOf(in.LatencyMicros)
	return out
}
func LoadBalancerBackendInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerBackendInfo) *krmv1alpha1.LoadBalancerBackendInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LoadBalancerBackendInfoObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.InstanceURI = direct.LazyPtr(in.GetInstanceUri())
	out.BackendServiceURI = direct.LazyPtr(in.GetBackendServiceUri())
	out.InstanceGroupURI = direct.LazyPtr(in.GetInstanceGroupUri())
	out.NetworkEndpointGroupURI = direct.LazyPtr(in.GetNetworkEndpointGroupUri())
	out.BackendBucketURI = direct.LazyPtr(in.GetBackendBucketUri())
	out.PSCServiceAttachmentURI = direct.LazyPtr(in.GetPscServiceAttachmentUri())
	out.PSCGoogleAPITarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	out.HealthCheckURI = direct.LazyPtr(in.GetHealthCheckUri())
	out.HealthCheckFirewallsConfigState = direct.Enum_FromProto(mapCtx, in.GetHealthCheckFirewallsConfigState())
	return out
}
func LoadBalancerBackendInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LoadBalancerBackendInfoObservedState) *pb.LoadBalancerBackendInfo {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancerBackendInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.InstanceUri = direct.ValueOf(in.InstanceURI)
	out.BackendServiceUri = direct.ValueOf(in.BackendServiceURI)
	out.InstanceGroupUri = direct.ValueOf(in.InstanceGroupURI)
	out.NetworkEndpointGroupUri = direct.ValueOf(in.NetworkEndpointGroupURI)
	out.BackendBucketUri = direct.ValueOf(in.BackendBucketURI)
	out.PscServiceAttachmentUri = direct.ValueOf(in.PSCServiceAttachmentURI)
	out.PscGoogleApiTarget = direct.ValueOf(in.PSCGoogleAPITarget)
	out.HealthCheckUri = direct.ValueOf(in.HealthCheckURI)
	out.HealthCheckFirewallsConfigState = direct.Enum_ToProto[pb.LoadBalancerBackendInfo_HealthCheckFirewallsConfigState](mapCtx, in.HealthCheckFirewallsConfigState)
	return out
}
func LoadBalancerBackendObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerBackend) *krmv1alpha1.LoadBalancerBackendObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LoadBalancerBackendObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.HealthCheckFirewallState = direct.Enum_FromProto(mapCtx, in.GetHealthCheckFirewallState())
	out.HealthCheckAllowingFirewallRules = in.HealthCheckAllowingFirewallRules
	out.HealthCheckBlockingFirewallRules = in.HealthCheckBlockingFirewallRules
	return out
}
func LoadBalancerBackendObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LoadBalancerBackendObservedState) *pb.LoadBalancerBackend {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancerBackend{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.HealthCheckFirewallState = direct.Enum_ToProto[pb.LoadBalancerBackend_HealthCheckFirewallState](mapCtx, in.HealthCheckFirewallState)
	out.HealthCheckAllowingFirewallRules = in.HealthCheckAllowingFirewallRules
	out.HealthCheckBlockingFirewallRules = in.HealthCheckBlockingFirewallRules
	return out
}
func LoadBalancerInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerInfo) *krmv1alpha1.LoadBalancerInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LoadBalancerInfoObservedState{}
	out.LoadBalancerType = direct.Enum_FromProto(mapCtx, in.GetLoadBalancerType())
	out.HealthCheckURI = direct.LazyPtr(in.GetHealthCheckUri())
	out.Backends = direct.Slice_FromProto(mapCtx, in.Backends, LoadBalancerBackendObservedState_FromProto)
	out.BackendType = direct.Enum_FromProto(mapCtx, in.GetBackendType())
	out.BackendURI = direct.LazyPtr(in.GetBackendUri())
	return out
}
func LoadBalancerInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LoadBalancerInfoObservedState) *pb.LoadBalancerInfo {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancerInfo{}
	out.LoadBalancerType = direct.Enum_ToProto[pb.LoadBalancerInfo_LoadBalancerType](mapCtx, in.LoadBalancerType)
	out.HealthCheckUri = direct.ValueOf(in.HealthCheckURI)
	out.Backends = direct.Slice_ToProto(mapCtx, in.Backends, LoadBalancerBackendObservedState_ToProto)
	out.BackendType = direct.Enum_ToProto[pb.LoadBalancerInfo_BackendType](mapCtx, in.BackendType)
	out.BackendUri = direct.ValueOf(in.BackendURI)
	return out
}
func NATInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NatInfo) *krmv1alpha1.NATInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NATInfoObservedState{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.OldSourceIP = direct.LazyPtr(in.GetOldSourceIp())
	out.NewSourceIP = direct.LazyPtr(in.GetNewSourceIp())
	out.OldDestinationIP = direct.LazyPtr(in.GetOldDestinationIp())
	out.NewDestinationIP = direct.LazyPtr(in.GetNewDestinationIp())
	out.OldSourcePort = direct.LazyPtr(in.GetOldSourcePort())
	out.NewSourcePort = direct.LazyPtr(in.GetNewSourcePort())
	out.OldDestinationPort = direct.LazyPtr(in.GetOldDestinationPort())
	out.NewDestinationPort = direct.LazyPtr(in.GetNewDestinationPort())
	out.RouterURI = direct.LazyPtr(in.GetRouterUri())
	out.NATGatewayName = direct.LazyPtr(in.GetNatGatewayName())
	return out
}
func NATInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NATInfoObservedState) *pb.NatInfo {
	if in == nil {
		return nil
	}
	out := &pb.NatInfo{}
	out.Type = direct.Enum_ToProto[pb.NatInfo_Type](mapCtx, in.Type)
	out.Protocol = direct.ValueOf(in.Protocol)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.OldSourceIp = direct.ValueOf(in.OldSourceIP)
	out.NewSourceIp = direct.ValueOf(in.NewSourceIP)
	out.OldDestinationIp = direct.ValueOf(in.OldDestinationIP)
	out.NewDestinationIp = direct.ValueOf(in.NewDestinationIP)
	out.OldSourcePort = direct.ValueOf(in.OldSourcePort)
	out.NewSourcePort = direct.ValueOf(in.NewSourcePort)
	out.OldDestinationPort = direct.ValueOf(in.OldDestinationPort)
	out.NewDestinationPort = direct.ValueOf(in.NewDestinationPort)
	out.RouterUri = direct.ValueOf(in.RouterURI)
	out.NatGatewayName = direct.ValueOf(in.NATGatewayName)
	return out
}
func NetworkInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInfo) *krmv1alpha1.NetworkInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NetworkInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.MatchedSubnetURI = direct.LazyPtr(in.GetMatchedSubnetUri())
	out.MatchedIPRange = direct.LazyPtr(in.GetMatchedIpRange())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func NetworkInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkInfoObservedState) *pb.NetworkInfo {
	if in == nil {
		return nil
	}
	out := &pb.NetworkInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.MatchedSubnetUri = direct.ValueOf(in.MatchedSubnetURI)
	out.MatchedIpRange = direct.ValueOf(in.MatchedIPRange)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func NetworkManagementConnectivityTestObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectivityTest) *krmv1alpha1.NetworkManagementConnectivityTestObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NetworkManagementConnectivityTestObservedState{}
	// MISSING: Name
	out.Source = EndpointObservedState_FromProto(mapCtx, in.GetSource())
	out.Destination = EndpointObservedState_FromProto(mapCtx, in.GetDestination())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ReachabilityDetails = ReachabilityDetailsObservedState_FromProto(mapCtx, in.GetReachabilityDetails())
	out.ProbingDetails = ProbingDetailsObservedState_FromProto(mapCtx, in.GetProbingDetails())
	out.ReturnReachabilityDetails = ReachabilityDetailsObservedState_FromProto(mapCtx, in.GetReturnReachabilityDetails())
	return out
}
func NetworkManagementConnectivityTestObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkManagementConnectivityTestObservedState) *pb.ConnectivityTest {
	if in == nil {
		return nil
	}
	out := &pb.ConnectivityTest{}
	// MISSING: Name
	out.Source = EndpointObservedState_ToProto(mapCtx, in.Source)
	out.Destination = EndpointObservedState_ToProto(mapCtx, in.Destination)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ReachabilityDetails = ReachabilityDetailsObservedState_ToProto(mapCtx, in.ReachabilityDetails)
	out.ProbingDetails = ProbingDetailsObservedState_ToProto(mapCtx, in.ProbingDetails)
	out.ReturnReachabilityDetails = ReachabilityDetailsObservedState_ToProto(mapCtx, in.ReturnReachabilityDetails)
	return out
}
func NetworkManagementConnectivityTestSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectivityTest) *krmv1alpha1.NetworkManagementConnectivityTestSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.NetworkManagementConnectivityTestSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Source = Endpoint_FromProto(mapCtx, in.GetSource())
	out.Destination = Endpoint_FromProto(mapCtx, in.GetDestination())
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.RelatedProjects = NetworkManagementConnectivityTestSpec_RelatedProjects_FromProto(mapCtx, in.RelatedProjects)
	out.Labels = in.Labels
	out.RoundTrip = direct.LazyPtr(in.GetRoundTrip())
	out.BypassFirewallChecks = direct.LazyPtr(in.GetBypassFirewallChecks())
	return out
}
func NetworkManagementConnectivityTestSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.NetworkManagementConnectivityTestSpec) *pb.ConnectivityTest {
	if in == nil {
		return nil
	}
	out := &pb.ConnectivityTest{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Source = Endpoint_ToProto(mapCtx, in.Source)
	out.Destination = Endpoint_ToProto(mapCtx, in.Destination)
	out.Protocol = direct.ValueOf(in.Protocol)
	out.RelatedProjects = NetworkManagementConnectivityTestSpec_RelatedProjects_ToProto(mapCtx, in.RelatedProjects)
	out.Labels = in.Labels
	out.RoundTrip = direct.ValueOf(in.RoundTrip)
	out.BypassFirewallChecks = direct.ValueOf(in.BypassFirewallChecks)
	return out
}
func ProbingDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails) *krmv1alpha1.ProbingDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ProbingDetailsObservedState{}
	out.Result = direct.Enum_FromProto(mapCtx, in.GetResult())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	out.Error = StatusObservedState_FromProto(mapCtx, in.GetError())
	out.AbortCause = direct.Enum_FromProto(mapCtx, in.GetAbortCause())
	out.SentProbeCount = direct.LazyPtr(in.GetSentProbeCount())
	out.SuccessfulProbeCount = direct.LazyPtr(in.GetSuccessfulProbeCount())
	out.EndpointInfo = EndpointInfoObservedState_FromProto(mapCtx, in.GetEndpointInfo())
	out.ProbingLatency = LatencyDistributionObservedState_FromProto(mapCtx, in.GetProbingLatency())
	out.DestinationEgressLocation = ProbingDetails_EdgeLocationObservedState_FromProto(mapCtx, in.GetDestinationEgressLocation())
	return out
}
func ProbingDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ProbingDetailsObservedState) *pb.ProbingDetails {
	if in == nil {
		return nil
	}
	out := &pb.ProbingDetails{}
	out.Result = direct.Enum_ToProto[pb.ProbingDetails_ProbingResult](mapCtx, in.Result)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	out.Error = StatusObservedState_ToProto(mapCtx, in.Error)
	out.AbortCause = direct.Enum_ToProto[pb.ProbingDetails_ProbingAbortCause](mapCtx, in.AbortCause)
	out.SentProbeCount = direct.ValueOf(in.SentProbeCount)
	out.SuccessfulProbeCount = direct.ValueOf(in.SuccessfulProbeCount)
	out.EndpointInfo = EndpointInfoObservedState_ToProto(mapCtx, in.EndpointInfo)
	out.ProbingLatency = LatencyDistributionObservedState_ToProto(mapCtx, in.ProbingLatency)
	out.DestinationEgressLocation = ProbingDetails_EdgeLocationObservedState_ToProto(mapCtx, in.DestinationEgressLocation)
	return out
}
func ProbingDetails_EdgeLocationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails_EdgeLocation) *krmv1alpha1.ProbingDetails_EdgeLocationObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ProbingDetails_EdgeLocationObservedState{}
	out.MetropolitanArea = direct.LazyPtr(in.GetMetropolitanArea())
	return out
}
func ProbingDetails_EdgeLocationObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ProbingDetails_EdgeLocationObservedState) *pb.ProbingDetails_EdgeLocation {
	if in == nil {
		return nil
	}
	out := &pb.ProbingDetails_EdgeLocation{}
	out.MetropolitanArea = direct.ValueOf(in.MetropolitanArea)
	return out
}
func ProxyConnectionInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProxyConnectionInfo) *krmv1alpha1.ProxyConnectionInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ProxyConnectionInfoObservedState{}
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.OldSourceIP = direct.LazyPtr(in.GetOldSourceIp())
	out.NewSourceIP = direct.LazyPtr(in.GetNewSourceIp())
	out.OldDestinationIP = direct.LazyPtr(in.GetOldDestinationIp())
	out.NewDestinationIP = direct.LazyPtr(in.GetNewDestinationIp())
	out.OldSourcePort = direct.LazyPtr(in.GetOldSourcePort())
	out.NewSourcePort = direct.LazyPtr(in.GetNewSourcePort())
	out.OldDestinationPort = direct.LazyPtr(in.GetOldDestinationPort())
	out.NewDestinationPort = direct.LazyPtr(in.GetNewDestinationPort())
	out.SubnetURI = direct.LazyPtr(in.GetSubnetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	return out
}
func ProxyConnectionInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ProxyConnectionInfoObservedState) *pb.ProxyConnectionInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProxyConnectionInfo{}
	out.Protocol = direct.ValueOf(in.Protocol)
	out.OldSourceIp = direct.ValueOf(in.OldSourceIP)
	out.NewSourceIp = direct.ValueOf(in.NewSourceIP)
	out.OldDestinationIp = direct.ValueOf(in.OldDestinationIP)
	out.NewDestinationIp = direct.ValueOf(in.NewDestinationIP)
	out.OldSourcePort = direct.ValueOf(in.OldSourcePort)
	out.NewSourcePort = direct.ValueOf(in.NewSourcePort)
	out.OldDestinationPort = direct.ValueOf(in.OldDestinationPort)
	out.NewDestinationPort = direct.ValueOf(in.NewDestinationPort)
	out.SubnetUri = direct.ValueOf(in.SubnetURI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	return out
}
func ReachabilityDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReachabilityDetails) *krmv1alpha1.ReachabilityDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ReachabilityDetailsObservedState{}
	out.Result = direct.Enum_FromProto(mapCtx, in.GetResult())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	out.Error = StatusObservedState_FromProto(mapCtx, in.GetError())
	out.Traces = direct.Slice_FromProto(mapCtx, in.Traces, TraceObservedState_FromProto)
	return out
}
func ReachabilityDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ReachabilityDetailsObservedState) *pb.ReachabilityDetails {
	if in == nil {
		return nil
	}
	out := &pb.ReachabilityDetails{}
	out.Result = direct.Enum_ToProto[pb.ReachabilityDetails_Result](mapCtx, in.Result)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	out.Error = StatusObservedState_ToProto(mapCtx, in.Error)
	out.Traces = direct.Slice_ToProto(mapCtx, in.Traces, TraceObservedState_ToProto)
	return out
}
func RedisClusterInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RedisClusterInfo) *krmv1alpha1.RedisClusterInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RedisClusterInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.DiscoveryEndpointIPAddress = direct.LazyPtr(in.GetDiscoveryEndpointIpAddress())
	out.SecondaryEndpointIPAddress = direct.LazyPtr(in.GetSecondaryEndpointIpAddress())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func RedisClusterInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RedisClusterInfoObservedState) *pb.RedisClusterInfo {
	if in == nil {
		return nil
	}
	out := &pb.RedisClusterInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.DiscoveryEndpointIpAddress = direct.ValueOf(in.DiscoveryEndpointIPAddress)
	out.SecondaryEndpointIpAddress = direct.ValueOf(in.SecondaryEndpointIPAddress)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func RedisInstanceInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RedisInstanceInfo) *krmv1alpha1.RedisInstanceInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RedisInstanceInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.PrimaryEndpointIP = direct.LazyPtr(in.GetPrimaryEndpointIp())
	out.ReadEndpointIP = direct.LazyPtr(in.GetReadEndpointIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func RedisInstanceInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RedisInstanceInfoObservedState) *pb.RedisInstanceInfo {
	if in == nil {
		return nil
	}
	out := &pb.RedisInstanceInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.PrimaryEndpointIp = direct.ValueOf(in.PrimaryEndpointIP)
	out.ReadEndpointIp = direct.ValueOf(in.ReadEndpointIP)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func RouteInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RouteInfo) *krmv1alpha1.RouteInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RouteInfoObservedState{}
	out.RouteType = direct.Enum_FromProto(mapCtx, in.GetRouteType())
	out.NextHopType = direct.Enum_FromProto(mapCtx, in.GetNextHopType())
	out.RouteScope = direct.Enum_FromProto(mapCtx, in.GetRouteScope())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Region = direct.LazyPtr(in.GetRegion())
	out.DestIPRange = direct.LazyPtr(in.GetDestIpRange())
	out.NextHop = direct.LazyPtr(in.GetNextHop())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.Priority = direct.LazyPtr(in.GetPriority())
	out.InstanceTags = in.InstanceTags
	out.SrcIPRange = direct.LazyPtr(in.GetSrcIpRange())
	out.DestPortRanges = in.DestPortRanges
	out.SrcPortRanges = in.SrcPortRanges
	out.Protocols = in.Protocols
	out.NccHubURI = in.NccHubUri
	out.NccSpokeURI = in.NccSpokeUri
	out.AdvertisedRouteSourceRouterURI = in.AdvertisedRouteSourceRouterUri
	out.AdvertisedRouteNextHopURI = in.AdvertisedRouteNextHopUri
	return out
}
func RouteInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RouteInfoObservedState) *pb.RouteInfo {
	if in == nil {
		return nil
	}
	out := &pb.RouteInfo{}
	out.RouteType = direct.Enum_ToProto[pb.RouteInfo_RouteType](mapCtx, in.RouteType)
	out.NextHopType = direct.Enum_ToProto[pb.RouteInfo_NextHopType](mapCtx, in.NextHopType)
	out.RouteScope = direct.Enum_ToProto[pb.RouteInfo_RouteScope](mapCtx, in.RouteScope)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Region = direct.ValueOf(in.Region)
	out.DestIpRange = direct.ValueOf(in.DestIPRange)
	out.NextHop = direct.ValueOf(in.NextHop)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.Priority = direct.ValueOf(in.Priority)
	out.InstanceTags = in.InstanceTags
	out.SrcIpRange = direct.ValueOf(in.SrcIPRange)
	out.DestPortRanges = in.DestPortRanges
	out.SrcPortRanges = in.SrcPortRanges
	out.Protocols = in.Protocols
	out.NccHubUri = in.NccHubURI
	out.NccSpokeUri = in.NccSpokeURI
	out.AdvertisedRouteSourceRouterUri = in.AdvertisedRouteSourceRouterURI
	out.AdvertisedRouteNextHopUri = in.AdvertisedRouteNextHopURI
	return out
}
func ServerlessNegInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ServerlessNegInfo) *krmv1alpha1.ServerlessNegInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ServerlessNegInfoObservedState{}
	out.NegURI = direct.LazyPtr(in.GetNegUri())
	return out
}
func ServerlessNegInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ServerlessNegInfoObservedState) *pb.ServerlessNegInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServerlessNegInfo{}
	out.NegUri = direct.ValueOf(in.NegURI)
	return out
}
func StepObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Step) *krmv1alpha1.StepObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StepObservedState{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CausesDrop = direct.LazyPtr(in.GetCausesDrop())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Instance = InstanceInfoObservedState_FromProto(mapCtx, in.GetInstance())
	out.Firewall = FirewallInfoObservedState_FromProto(mapCtx, in.GetFirewall())
	out.Route = RouteInfoObservedState_FromProto(mapCtx, in.GetRoute())
	out.Endpoint = EndpointInfoObservedState_FromProto(mapCtx, in.GetEndpoint())
	out.GoogleService = GoogleServiceInfoObservedState_FromProto(mapCtx, in.GetGoogleService())
	out.ForwardingRule = ForwardingRuleInfoObservedState_FromProto(mapCtx, in.GetForwardingRule())
	out.VPNGateway = VPNGatewayInfoObservedState_FromProto(mapCtx, in.GetVpnGateway())
	out.VPNTunnel = VPNTunnelInfoObservedState_FromProto(mapCtx, in.GetVpnTunnel())
	out.VPCConnector = VPCConnectorInfoObservedState_FromProto(mapCtx, in.GetVpcConnector())
	out.Deliver = DeliverInfoObservedState_FromProto(mapCtx, in.GetDeliver())
	out.Forward = ForwardInfoObservedState_FromProto(mapCtx, in.GetForward())
	out.Abort = AbortInfoObservedState_FromProto(mapCtx, in.GetAbort())
	out.Drop = DropInfoObservedState_FromProto(mapCtx, in.GetDrop())
	out.LoadBalancer = LoadBalancerInfoObservedState_FromProto(mapCtx, in.GetLoadBalancer())
	out.Network = NetworkInfoObservedState_FromProto(mapCtx, in.GetNetwork())
	out.GKEMaster = GKEMasterInfoObservedState_FromProto(mapCtx, in.GetGkeMaster())
	out.CloudSQLInstance = CloudSQLInstanceInfoObservedState_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.RedisInstance = RedisInstanceInfoObservedState_FromProto(mapCtx, in.GetRedisInstance())
	out.RedisCluster = RedisClusterInfoObservedState_FromProto(mapCtx, in.GetRedisCluster())
	out.CloudFunction = CloudFunctionInfoObservedState_FromProto(mapCtx, in.GetCloudFunction())
	out.AppEngineVersion = AppEngineVersionInfoObservedState_FromProto(mapCtx, in.GetAppEngineVersion())
	out.CloudRunRevision = CloudRunRevisionInfoObservedState_FromProto(mapCtx, in.GetCloudRunRevision())
	out.NAT = NATInfoObservedState_FromProto(mapCtx, in.GetNat())
	out.ProxyConnection = ProxyConnectionInfoObservedState_FromProto(mapCtx, in.GetProxyConnection())
	out.LoadBalancerBackendInfo = LoadBalancerBackendInfoObservedState_FromProto(mapCtx, in.GetLoadBalancerBackendInfo())
	out.StorageBucket = StorageBucketInfoObservedState_FromProto(mapCtx, in.GetStorageBucket())
	out.ServerlessNeg = ServerlessNegInfoObservedState_FromProto(mapCtx, in.GetServerlessNeg())
	return out
}
func StepObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.StepObservedState) *pb.Step {
	if in == nil {
		return nil
	}
	out := &pb.Step{}
	out.Description = direct.ValueOf(in.Description)
	out.State = direct.Enum_ToProto[pb.Step_State](mapCtx, in.State)
	out.CausesDrop = direct.ValueOf(in.CausesDrop)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	if oneof := InstanceInfoObservedState_ToProto(mapCtx, in.Instance); oneof != nil {
		out.StepInfo = &pb.Step_Instance{Instance: oneof}
	}
	if oneof := FirewallInfoObservedState_ToProto(mapCtx, in.Firewall); oneof != nil {
		out.StepInfo = &pb.Step_Firewall{Firewall: oneof}
	}
	if oneof := RouteInfoObservedState_ToProto(mapCtx, in.Route); oneof != nil {
		out.StepInfo = &pb.Step_Route{Route: oneof}
	}
	if oneof := EndpointInfoObservedState_ToProto(mapCtx, in.Endpoint); oneof != nil {
		out.StepInfo = &pb.Step_Endpoint{Endpoint: oneof}
	}
	if oneof := GoogleServiceInfoObservedState_ToProto(mapCtx, in.GoogleService); oneof != nil {
		out.StepInfo = &pb.Step_GoogleService{GoogleService: oneof}
	}
	if oneof := ForwardingRuleInfoObservedState_ToProto(mapCtx, in.ForwardingRule); oneof != nil {
		out.StepInfo = &pb.Step_ForwardingRule{ForwardingRule: oneof}
	}
	if oneof := VPNGatewayInfoObservedState_ToProto(mapCtx, in.VPNGateway); oneof != nil {
		out.StepInfo = &pb.Step_VpnGateway{VpnGateway: oneof}
	}
	if oneof := VPNTunnelInfoObservedState_ToProto(mapCtx, in.VPNTunnel); oneof != nil {
		out.StepInfo = &pb.Step_VpnTunnel{VpnTunnel: oneof}
	}
	if oneof := VPCConnectorInfoObservedState_ToProto(mapCtx, in.VPCConnector); oneof != nil {
		out.StepInfo = &pb.Step_VpcConnector{VpcConnector: oneof}
	}
	if oneof := DeliverInfoObservedState_ToProto(mapCtx, in.Deliver); oneof != nil {
		out.StepInfo = &pb.Step_Deliver{Deliver: oneof}
	}
	if oneof := ForwardInfoObservedState_ToProto(mapCtx, in.Forward); oneof != nil {
		out.StepInfo = &pb.Step_Forward{Forward: oneof}
	}
	if oneof := AbortInfoObservedState_ToProto(mapCtx, in.Abort); oneof != nil {
		out.StepInfo = &pb.Step_Abort{Abort: oneof}
	}
	if oneof := DropInfoObservedState_ToProto(mapCtx, in.Drop); oneof != nil {
		out.StepInfo = &pb.Step_Drop{Drop: oneof}
	}
	if oneof := LoadBalancerInfoObservedState_ToProto(mapCtx, in.LoadBalancer); oneof != nil {
		out.StepInfo = &pb.Step_LoadBalancer{LoadBalancer: oneof}
	}
	if oneof := NetworkInfoObservedState_ToProto(mapCtx, in.Network); oneof != nil {
		out.StepInfo = &pb.Step_Network{Network: oneof}
	}
	if oneof := GKEMasterInfoObservedState_ToProto(mapCtx, in.GKEMaster); oneof != nil {
		out.StepInfo = &pb.Step_GkeMaster{GkeMaster: oneof}
	}
	if oneof := CloudSQLInstanceInfoObservedState_ToProto(mapCtx, in.CloudSQLInstance); oneof != nil {
		out.StepInfo = &pb.Step_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := RedisInstanceInfoObservedState_ToProto(mapCtx, in.RedisInstance); oneof != nil {
		out.StepInfo = &pb.Step_RedisInstance{RedisInstance: oneof}
	}
	if oneof := RedisClusterInfoObservedState_ToProto(mapCtx, in.RedisCluster); oneof != nil {
		out.StepInfo = &pb.Step_RedisCluster{RedisCluster: oneof}
	}
	if oneof := CloudFunctionInfoObservedState_ToProto(mapCtx, in.CloudFunction); oneof != nil {
		out.StepInfo = &pb.Step_CloudFunction{CloudFunction: oneof}
	}
	if oneof := AppEngineVersionInfoObservedState_ToProto(mapCtx, in.AppEngineVersion); oneof != nil {
		out.StepInfo = &pb.Step_AppEngineVersion{AppEngineVersion: oneof}
	}
	if oneof := CloudRunRevisionInfoObservedState_ToProto(mapCtx, in.CloudRunRevision); oneof != nil {
		out.StepInfo = &pb.Step_CloudRunRevision{CloudRunRevision: oneof}
	}
	if oneof := NATInfoObservedState_ToProto(mapCtx, in.NAT); oneof != nil {
		out.StepInfo = &pb.Step_Nat{Nat: oneof}
	}
	if oneof := ProxyConnectionInfoObservedState_ToProto(mapCtx, in.ProxyConnection); oneof != nil {
		out.StepInfo = &pb.Step_ProxyConnection{ProxyConnection: oneof}
	}
	if oneof := LoadBalancerBackendInfoObservedState_ToProto(mapCtx, in.LoadBalancerBackendInfo); oneof != nil {
		out.StepInfo = &pb.Step_LoadBalancerBackendInfo{LoadBalancerBackendInfo: oneof}
	}
	if oneof := StorageBucketInfoObservedState_ToProto(mapCtx, in.StorageBucket); oneof != nil {
		out.StepInfo = &pb.Step_StorageBucket{StorageBucket: oneof}
	}
	if oneof := ServerlessNegInfoObservedState_ToProto(mapCtx, in.ServerlessNeg); oneof != nil {
		out.StepInfo = &pb.Step_ServerlessNeg{ServerlessNeg: oneof}
	}
	return out
}
func StorageBucketInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StorageBucketInfo) *krmv1alpha1.StorageBucketInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.StorageBucketInfoObservedState{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	return out
}
func StorageBucketInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.StorageBucketInfoObservedState) *pb.StorageBucketInfo {
	if in == nil {
		return nil
	}
	out := &pb.StorageBucketInfo{}
	out.Bucket = direct.ValueOf(in.Bucket)
	return out
}
func TraceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trace) *krmv1alpha1.TraceObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.TraceObservedState{}
	out.EndpointInfo = EndpointInfoObservedState_FromProto(mapCtx, in.GetEndpointInfo())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, StepObservedState_FromProto)
	out.ForwardTraceID = direct.LazyPtr(in.GetForwardTraceId())
	return out
}
func TraceObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.TraceObservedState) *pb.Trace {
	if in == nil {
		return nil
	}
	out := &pb.Trace{}
	out.EndpointInfo = EndpointInfoObservedState_ToProto(mapCtx, in.EndpointInfo)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, StepObservedState_ToProto)
	out.ForwardTraceId = direct.ValueOf(in.ForwardTraceID)
	return out
}
func VPCConnectorInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpcConnectorInfo) *krmv1alpha1.VPCConnectorInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VPCConnectorInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func VPCConnectorInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VPCConnectorInfoObservedState) *pb.VpcConnectorInfo {
	if in == nil {
		return nil
	}
	out := &pb.VpcConnectorInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func VPNGatewayInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpnGatewayInfo) *krmv1alpha1.VPNGatewayInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VPNGatewayInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	// MISSING: VPNTunnelURI
	// (near miss): "VPNTunnelURI" vs "VpnTunnelURI"
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func VPNGatewayInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VPNGatewayInfoObservedState) *pb.VpnGatewayInfo {
	if in == nil {
		return nil
	}
	out := &pb.VpnGatewayInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	// MISSING: VPNTunnelURI
	// (near miss): "VPNTunnelURI" vs "VpnTunnelURI"
	out.Region = direct.ValueOf(in.Region)
	return out
}
func VPNTunnelInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.VpnTunnelInfo) *krmv1alpha1.VPNTunnelInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.VPNTunnelInfoObservedState{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.SourceGateway = direct.LazyPtr(in.GetSourceGateway())
	out.RemoteGateway = direct.LazyPtr(in.GetRemoteGateway())
	out.RemoteGatewayIP = direct.LazyPtr(in.GetRemoteGatewayIp())
	out.SourceGatewayIP = direct.LazyPtr(in.GetSourceGatewayIp())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.Region = direct.LazyPtr(in.GetRegion())
	out.RoutingType = direct.Enum_FromProto(mapCtx, in.GetRoutingType())
	return out
}
func VPNTunnelInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.VPNTunnelInfoObservedState) *pb.VpnTunnelInfo {
	if in == nil {
		return nil
	}
	out := &pb.VpnTunnelInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.SourceGateway = direct.ValueOf(in.SourceGateway)
	out.RemoteGateway = direct.ValueOf(in.RemoteGateway)
	out.RemoteGatewayIp = direct.ValueOf(in.RemoteGatewayIP)
	out.SourceGatewayIp = direct.ValueOf(in.SourceGatewayIP)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.Region = direct.ValueOf(in.Region)
	out.RoutingType = direct.Enum_ToProto[pb.VpnTunnelInfo_RoutingType](mapCtx, in.RoutingType)
	return out
}
