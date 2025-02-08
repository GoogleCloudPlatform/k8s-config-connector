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

package networkmanagement

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkmanagement/apiv1beta1/networkmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
)
func AbortInfo_FromProto(mapCtx *direct.MapContext, in *pb.AbortInfo) *krm.AbortInfo {
	if in == nil {
		return nil
	}
	out := &krm.AbortInfo{}
	out.Cause = direct.Enum_FromProto(mapCtx, in.GetCause())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.ProjectsMissingPermission = in.ProjectsMissingPermission
	return out
}
func AbortInfo_ToProto(mapCtx *direct.MapContext, in *krm.AbortInfo) *pb.AbortInfo {
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
func AppEngineVersionInfo_FromProto(mapCtx *direct.MapContext, in *pb.AppEngineVersionInfo) *krm.AppEngineVersionInfo {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineVersionInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Runtime = direct.LazyPtr(in.GetRuntime())
	out.Environment = direct.LazyPtr(in.GetEnvironment())
	return out
}
func AppEngineVersionInfo_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineVersionInfo) *pb.AppEngineVersionInfo {
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
func CloudFunctionInfo_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunctionInfo) *krm.CloudFunctionInfo {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunctionInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.VersionID = direct.LazyPtr(in.GetVersionId())
	return out
}
func CloudFunctionInfo_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunctionInfo) *pb.CloudFunctionInfo {
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
func CloudRunRevisionInfo_FromProto(mapCtx *direct.MapContext, in *pb.CloudRunRevisionInfo) *krm.CloudRunRevisionInfo {
	if in == nil {
		return nil
	}
	out := &krm.CloudRunRevisionInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	out.ServiceURI = direct.LazyPtr(in.GetServiceUri())
	return out
}
func CloudRunRevisionInfo_ToProto(mapCtx *direct.MapContext, in *krm.CloudRunRevisionInfo) *pb.CloudRunRevisionInfo {
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
func CloudSQLInstanceInfo_FromProto(mapCtx *direct.MapContext, in *pb.CloudSQLInstanceInfo) *krm.CloudSQLInstanceInfo {
	if in == nil {
		return nil
	}
	out := &krm.CloudSQLInstanceInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func CloudSQLInstanceInfo_ToProto(mapCtx *direct.MapContext, in *krm.CloudSQLInstanceInfo) *pb.CloudSQLInstanceInfo {
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
func ConnectivityTest_FromProto(mapCtx *direct.MapContext, in *pb.ConnectivityTest) *krm.ConnectivityTest {
	if in == nil {
		return nil
	}
	out := &krm.ConnectivityTest{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Source = Endpoint_FromProto(mapCtx, in.GetSource())
	out.Destination = Endpoint_FromProto(mapCtx, in.GetDestination())
	out.Protocol = direct.LazyPtr(in.GetProtocol())
	out.RelatedProjects = in.RelatedProjects
	// MISSING: DisplayName
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReachabilityDetails
	// MISSING: ProbingDetails
	out.RoundTrip = direct.LazyPtr(in.GetRoundTrip())
	// MISSING: ReturnReachabilityDetails
	out.BypassFirewallChecks = direct.LazyPtr(in.GetBypassFirewallChecks())
	return out
}
func ConnectivityTest_ToProto(mapCtx *direct.MapContext, in *krm.ConnectivityTest) *pb.ConnectivityTest {
	if in == nil {
		return nil
	}
	out := &pb.ConnectivityTest{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Source = Endpoint_ToProto(mapCtx, in.Source)
	out.Destination = Endpoint_ToProto(mapCtx, in.Destination)
	out.Protocol = direct.ValueOf(in.Protocol)
	out.RelatedProjects = in.RelatedProjects
	// MISSING: DisplayName
	out.Labels = in.Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: ReachabilityDetails
	// MISSING: ProbingDetails
	out.RoundTrip = direct.ValueOf(in.RoundTrip)
	// MISSING: ReturnReachabilityDetails
	out.BypassFirewallChecks = direct.ValueOf(in.BypassFirewallChecks)
	return out
}
func ConnectivityTestObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectivityTest) *krm.ConnectivityTestObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConnectivityTestObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.Source = EndpointObservedState_FromProto(mapCtx, in.GetSource())
	// MISSING: Destination
	// MISSING: Protocol
	// MISSING: RelatedProjects
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ReachabilityDetails = ReachabilityDetails_FromProto(mapCtx, in.GetReachabilityDetails())
	out.ProbingDetails = ProbingDetails_FromProto(mapCtx, in.GetProbingDetails())
	// MISSING: RoundTrip
	out.ReturnReachabilityDetails = ReachabilityDetails_FromProto(mapCtx, in.GetReturnReachabilityDetails())
	// MISSING: BypassFirewallChecks
	return out
}
func ConnectivityTestObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConnectivityTestObservedState) *pb.ConnectivityTest {
	if in == nil {
		return nil
	}
	out := &pb.ConnectivityTest{}
	// MISSING: Name
	// MISSING: Description
	out.Source = EndpointObservedState_ToProto(mapCtx, in.Source)
	// MISSING: Destination
	// MISSING: Protocol
	// MISSING: RelatedProjects
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ReachabilityDetails = ReachabilityDetails_ToProto(mapCtx, in.ReachabilityDetails)
	out.ProbingDetails = ProbingDetails_ToProto(mapCtx, in.ProbingDetails)
	// MISSING: RoundTrip
	out.ReturnReachabilityDetails = ReachabilityDetails_ToProto(mapCtx, in.ReturnReachabilityDetails)
	// MISSING: BypassFirewallChecks
	return out
}
func DeliverInfo_FromProto(mapCtx *direct.MapContext, in *pb.DeliverInfo) *krm.DeliverInfo {
	if in == nil {
		return nil
	}
	out := &krm.DeliverInfo{}
	out.Target = direct.Enum_FromProto(mapCtx, in.GetTarget())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.StorageBucket = direct.LazyPtr(in.GetStorageBucket())
	out.PscGoogleApiTarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	return out
}
func DeliverInfo_ToProto(mapCtx *direct.MapContext, in *krm.DeliverInfo) *pb.DeliverInfo {
	if in == nil {
		return nil
	}
	out := &pb.DeliverInfo{}
	out.Target = direct.Enum_ToProto[pb.DeliverInfo_Target](mapCtx, in.Target)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.StorageBucket = direct.ValueOf(in.StorageBucket)
	out.PscGoogleApiTarget = direct.ValueOf(in.PscGoogleApiTarget)
	return out
}
func DropInfo_FromProto(mapCtx *direct.MapContext, in *pb.DropInfo) *krm.DropInfo {
	if in == nil {
		return nil
	}
	out := &krm.DropInfo{}
	out.Cause = direct.Enum_FromProto(mapCtx, in.GetCause())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.DestinationIP = direct.LazyPtr(in.GetDestinationIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func DropInfo_ToProto(mapCtx *direct.MapContext, in *krm.DropInfo) *pb.DropInfo {
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
func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.ForwardingRule = direct.LazyPtr(in.GetForwardingRule())
	// MISSING: ForwardingRuleTarget
	// MISSING: LoadBalancerID
	// MISSING: LoadBalancerType
	out.GkeMasterCluster = direct.LazyPtr(in.GetGkeMasterCluster())
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
	out.CloudSqlInstance = direct.LazyPtr(in.GetCloudSqlInstance())
	out.RedisInstance = direct.LazyPtr(in.GetRedisInstance())
	out.RedisCluster = direct.LazyPtr(in.GetRedisCluster())
	out.CloudFunction = Endpoint_CloudFunctionEndpoint_FromProto(mapCtx, in.GetCloudFunction())
	out.AppEngineVersion = Endpoint_AppEngineVersionEndpoint_FromProto(mapCtx, in.GetAppEngineVersion())
	out.CloudRunRevision = Endpoint_CloudRunRevisionEndpoint_FromProto(mapCtx, in.GetCloudRunRevision())
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.NetworkType = direct.Enum_FromProto(mapCtx, in.GetNetworkType())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Port = direct.ValueOf(in.Port)
	out.Instance = direct.ValueOf(in.Instance)
	out.ForwardingRule = direct.ValueOf(in.ForwardingRule)
	// MISSING: ForwardingRuleTarget
	// MISSING: LoadBalancerID
	// MISSING: LoadBalancerType
	out.GkeMasterCluster = direct.ValueOf(in.GkeMasterCluster)
	out.Fqdn = direct.ValueOf(in.Fqdn)
	out.CloudSqlInstance = direct.ValueOf(in.CloudSqlInstance)
	out.RedisInstance = direct.ValueOf(in.RedisInstance)
	out.RedisCluster = direct.ValueOf(in.RedisCluster)
	out.CloudFunction = Endpoint_CloudFunctionEndpoint_ToProto(mapCtx, in.CloudFunction)
	out.AppEngineVersion = Endpoint_AppEngineVersionEndpoint_ToProto(mapCtx, in.AppEngineVersion)
	out.CloudRunRevision = Endpoint_CloudRunRevisionEndpoint_ToProto(mapCtx, in.CloudRunRevision)
	out.Network = direct.ValueOf(in.Network)
	out.NetworkType = direct.Enum_ToProto[pb.Endpoint_NetworkType](mapCtx, in.NetworkType)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	return out
}
func EndpointInfo_FromProto(mapCtx *direct.MapContext, in *pb.EndpointInfo) *krm.EndpointInfo {
	if in == nil {
		return nil
	}
	out := &krm.EndpointInfo{}
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
func EndpointInfo_ToProto(mapCtx *direct.MapContext, in *krm.EndpointInfo) *pb.EndpointInfo {
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
func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	// MISSING: IPAddress
	// MISSING: Port
	// MISSING: Instance
	// MISSING: ForwardingRule
	out.ForwardingRuleTarget = direct.Enum_FromProto(mapCtx, in.GetForwardingRuleTarget())
	out.LoadBalancerID = in.LoadBalancerId
	out.LoadBalancerType = direct.Enum_FromProto(mapCtx, in.GetLoadBalancerType())
	// MISSING: GkeMasterCluster
	// MISSING: Fqdn
	// MISSING: CloudSqlInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Network
	// MISSING: NetworkType
	// MISSING: ProjectID
	return out
}
func EndpointObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EndpointObservedState) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	// MISSING: IPAddress
	// MISSING: Port
	// MISSING: Instance
	// MISSING: ForwardingRule
	if oneof := EndpointObservedState_ForwardingRuleTarget_ToProto(mapCtx, in.ForwardingRuleTarget); oneof != nil {
		out.ForwardingRuleTarget = oneof
	}
	out.LoadBalancerId = in.LoadBalancerID
	if oneof := EndpointObservedState_LoadBalancerType_ToProto(mapCtx, in.LoadBalancerType); oneof != nil {
		out.LoadBalancerType = oneof
	}
	// MISSING: GkeMasterCluster
	// MISSING: Fqdn
	// MISSING: CloudSqlInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Network
	// MISSING: NetworkType
	// MISSING: ProjectID
	return out
}
func Endpoint_AppEngineVersionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_AppEngineVersionEndpoint) *krm.Endpoint_AppEngineVersionEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint_AppEngineVersionEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Endpoint_AppEngineVersionEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_AppEngineVersionEndpoint) *pb.Endpoint_AppEngineVersionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_AppEngineVersionEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Endpoint_CloudFunctionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_CloudFunctionEndpoint) *krm.Endpoint_CloudFunctionEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint_CloudFunctionEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Endpoint_CloudFunctionEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_CloudFunctionEndpoint) *pb.Endpoint_CloudFunctionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_CloudFunctionEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func Endpoint_CloudRunRevisionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_CloudRunRevisionEndpoint) *krm.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint_CloudRunRevisionEndpoint{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func Endpoint_CloudRunRevisionEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_CloudRunRevisionEndpoint) *pb.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_CloudRunRevisionEndpoint{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func FirewallInfo_FromProto(mapCtx *direct.MapContext, in *pb.FirewallInfo) *krm.FirewallInfo {
	if in == nil {
		return nil
	}
	out := &krm.FirewallInfo{}
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
func FirewallInfo_ToProto(mapCtx *direct.MapContext, in *krm.FirewallInfo) *pb.FirewallInfo {
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
func ForwardInfo_FromProto(mapCtx *direct.MapContext, in *pb.ForwardInfo) *krm.ForwardInfo {
	if in == nil {
		return nil
	}
	out := &krm.ForwardInfo{}
	out.Target = direct.Enum_FromProto(mapCtx, in.GetTarget())
	out.ResourceURI = direct.LazyPtr(in.GetResourceUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	return out
}
func ForwardInfo_ToProto(mapCtx *direct.MapContext, in *krm.ForwardInfo) *pb.ForwardInfo {
	if in == nil {
		return nil
	}
	out := &pb.ForwardInfo{}
	out.Target = direct.Enum_ToProto[pb.ForwardInfo_Target](mapCtx, in.Target)
	out.ResourceUri = direct.ValueOf(in.ResourceURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	return out
}
func ForwardingRuleInfo_FromProto(mapCtx *direct.MapContext, in *pb.ForwardingRuleInfo) *krm.ForwardingRuleInfo {
	if in == nil {
		return nil
	}
	out := &krm.ForwardingRuleInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.MatchedProtocol = direct.LazyPtr(in.GetMatchedProtocol())
	out.MatchedPortRange = direct.LazyPtr(in.GetMatchedPortRange())
	out.Vip = direct.LazyPtr(in.GetVip())
	out.Target = direct.LazyPtr(in.GetTarget())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.Region = direct.LazyPtr(in.GetRegion())
	out.LoadBalancerName = direct.LazyPtr(in.GetLoadBalancerName())
	out.PscServiceAttachmentURI = direct.LazyPtr(in.GetPscServiceAttachmentUri())
	out.PscGoogleApiTarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	return out
}
func ForwardingRuleInfo_ToProto(mapCtx *direct.MapContext, in *krm.ForwardingRuleInfo) *pb.ForwardingRuleInfo {
	if in == nil {
		return nil
	}
	out := &pb.ForwardingRuleInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.MatchedProtocol = direct.ValueOf(in.MatchedProtocol)
	out.MatchedPortRange = direct.ValueOf(in.MatchedPortRange)
	out.Vip = direct.ValueOf(in.Vip)
	out.Target = direct.ValueOf(in.Target)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.Region = direct.ValueOf(in.Region)
	out.LoadBalancerName = direct.ValueOf(in.LoadBalancerName)
	out.PscServiceAttachmentUri = direct.ValueOf(in.PscServiceAttachmentURI)
	out.PscGoogleApiTarget = direct.ValueOf(in.PscGoogleApiTarget)
	return out
}
func GKEMasterInfo_FromProto(mapCtx *direct.MapContext, in *pb.GKEMasterInfo) *krm.GKEMasterInfo {
	if in == nil {
		return nil
	}
	out := &krm.GKEMasterInfo{}
	out.ClusterURI = direct.LazyPtr(in.GetClusterUri())
	out.ClusterNetworkURI = direct.LazyPtr(in.GetClusterNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.DnsEndpoint = direct.LazyPtr(in.GetDnsEndpoint())
	return out
}
func GKEMasterInfo_ToProto(mapCtx *direct.MapContext, in *krm.GKEMasterInfo) *pb.GKEMasterInfo {
	if in == nil {
		return nil
	}
	out := &pb.GKEMasterInfo{}
	out.ClusterUri = direct.ValueOf(in.ClusterURI)
	out.ClusterNetworkUri = direct.ValueOf(in.ClusterNetworkURI)
	out.InternalIp = direct.ValueOf(in.InternalIP)
	out.ExternalIp = direct.ValueOf(in.ExternalIP)
	out.DnsEndpoint = direct.ValueOf(in.DnsEndpoint)
	return out
}
func GoogleServiceInfo_FromProto(mapCtx *direct.MapContext, in *pb.GoogleServiceInfo) *krm.GoogleServiceInfo {
	if in == nil {
		return nil
	}
	out := &krm.GoogleServiceInfo{}
	out.SourceIP = direct.LazyPtr(in.GetSourceIp())
	out.GoogleServiceType = direct.Enum_FromProto(mapCtx, in.GetGoogleServiceType())
	return out
}
func GoogleServiceInfo_ToProto(mapCtx *direct.MapContext, in *krm.GoogleServiceInfo) *pb.GoogleServiceInfo {
	if in == nil {
		return nil
	}
	out := &pb.GoogleServiceInfo{}
	out.SourceIp = direct.ValueOf(in.SourceIP)
	out.GoogleServiceType = direct.Enum_ToProto[pb.GoogleServiceInfo_GoogleServiceType](mapCtx, in.GoogleServiceType)
	return out
}
func InstanceInfo_FromProto(mapCtx *direct.MapContext, in *pb.InstanceInfo) *krm.InstanceInfo {
	if in == nil {
		return nil
	}
	out := &krm.InstanceInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Interface = direct.LazyPtr(in.GetInterface())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.InternalIP = direct.LazyPtr(in.GetInternalIp())
	out.ExternalIP = direct.LazyPtr(in.GetExternalIp())
	out.NetworkTags = in.NetworkTags
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.PscNetworkAttachmentURI = direct.LazyPtr(in.GetPscNetworkAttachmentUri())
	return out
}
func InstanceInfo_ToProto(mapCtx *direct.MapContext, in *krm.InstanceInfo) *pb.InstanceInfo {
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
	out.PscNetworkAttachmentUri = direct.ValueOf(in.PscNetworkAttachmentURI)
	return out
}
func LatencyDistribution_FromProto(mapCtx *direct.MapContext, in *pb.LatencyDistribution) *krm.LatencyDistribution {
	if in == nil {
		return nil
	}
	out := &krm.LatencyDistribution{}
	out.LatencyPercentiles = direct.Slice_FromProto(mapCtx, in.LatencyPercentiles, LatencyPercentile_FromProto)
	return out
}
func LatencyDistribution_ToProto(mapCtx *direct.MapContext, in *krm.LatencyDistribution) *pb.LatencyDistribution {
	if in == nil {
		return nil
	}
	out := &pb.LatencyDistribution{}
	out.LatencyPercentiles = direct.Slice_ToProto(mapCtx, in.LatencyPercentiles, LatencyPercentile_ToProto)
	return out
}
func LatencyPercentile_FromProto(mapCtx *direct.MapContext, in *pb.LatencyPercentile) *krm.LatencyPercentile {
	if in == nil {
		return nil
	}
	out := &krm.LatencyPercentile{}
	out.Percent = direct.LazyPtr(in.GetPercent())
	out.LatencyMicros = direct.LazyPtr(in.GetLatencyMicros())
	return out
}
func LatencyPercentile_ToProto(mapCtx *direct.MapContext, in *krm.LatencyPercentile) *pb.LatencyPercentile {
	if in == nil {
		return nil
	}
	out := &pb.LatencyPercentile{}
	out.Percent = direct.ValueOf(in.Percent)
	out.LatencyMicros = direct.ValueOf(in.LatencyMicros)
	return out
}
func LoadBalancerBackend_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerBackend) *krm.LoadBalancerBackend {
	if in == nil {
		return nil
	}
	out := &krm.LoadBalancerBackend{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.HealthCheckFirewallState = direct.Enum_FromProto(mapCtx, in.GetHealthCheckFirewallState())
	out.HealthCheckAllowingFirewallRules = in.HealthCheckAllowingFirewallRules
	out.HealthCheckBlockingFirewallRules = in.HealthCheckBlockingFirewallRules
	return out
}
func LoadBalancerBackend_ToProto(mapCtx *direct.MapContext, in *krm.LoadBalancerBackend) *pb.LoadBalancerBackend {
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
func LoadBalancerBackendInfo_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerBackendInfo) *krm.LoadBalancerBackendInfo {
	if in == nil {
		return nil
	}
	out := &krm.LoadBalancerBackendInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.InstanceURI = direct.LazyPtr(in.GetInstanceUri())
	out.BackendServiceURI = direct.LazyPtr(in.GetBackendServiceUri())
	out.InstanceGroupURI = direct.LazyPtr(in.GetInstanceGroupUri())
	out.NetworkEndpointGroupURI = direct.LazyPtr(in.GetNetworkEndpointGroupUri())
	out.BackendBucketURI = direct.LazyPtr(in.GetBackendBucketUri())
	out.PscServiceAttachmentURI = direct.LazyPtr(in.GetPscServiceAttachmentUri())
	out.PscGoogleApiTarget = direct.LazyPtr(in.GetPscGoogleApiTarget())
	out.HealthCheckURI = direct.LazyPtr(in.GetHealthCheckUri())
	// MISSING: HealthCheckFirewallsConfigState
	return out
}
func LoadBalancerBackendInfo_ToProto(mapCtx *direct.MapContext, in *krm.LoadBalancerBackendInfo) *pb.LoadBalancerBackendInfo {
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
	out.PscServiceAttachmentUri = direct.ValueOf(in.PscServiceAttachmentURI)
	out.PscGoogleApiTarget = direct.ValueOf(in.PscGoogleApiTarget)
	out.HealthCheckUri = direct.ValueOf(in.HealthCheckURI)
	// MISSING: HealthCheckFirewallsConfigState
	return out
}
func LoadBalancerBackendInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerBackendInfo) *krm.LoadBalancerBackendInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoadBalancerBackendInfoObservedState{}
	// MISSING: Name
	// MISSING: InstanceURI
	// MISSING: BackendServiceURI
	// MISSING: InstanceGroupURI
	// MISSING: NetworkEndpointGroupURI
	// MISSING: BackendBucketURI
	// MISSING: PscServiceAttachmentURI
	// MISSING: PscGoogleApiTarget
	// MISSING: HealthCheckURI
	out.HealthCheckFirewallsConfigState = direct.Enum_FromProto(mapCtx, in.GetHealthCheckFirewallsConfigState())
	return out
}
func LoadBalancerBackendInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoadBalancerBackendInfoObservedState) *pb.LoadBalancerBackendInfo {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancerBackendInfo{}
	// MISSING: Name
	// MISSING: InstanceURI
	// MISSING: BackendServiceURI
	// MISSING: InstanceGroupURI
	// MISSING: NetworkEndpointGroupURI
	// MISSING: BackendBucketURI
	// MISSING: PscServiceAttachmentURI
	// MISSING: PscGoogleApiTarget
	// MISSING: HealthCheckURI
	out.HealthCheckFirewallsConfigState = direct.Enum_ToProto[pb.LoadBalancerBackendInfo_HealthCheckFirewallsConfigState](mapCtx, in.HealthCheckFirewallsConfigState)
	return out
}
func LoadBalancerInfo_FromProto(mapCtx *direct.MapContext, in *pb.LoadBalancerInfo) *krm.LoadBalancerInfo {
	if in == nil {
		return nil
	}
	out := &krm.LoadBalancerInfo{}
	out.LoadBalancerType = direct.Enum_FromProto(mapCtx, in.GetLoadBalancerType())
	out.HealthCheckURI = direct.LazyPtr(in.GetHealthCheckUri())
	out.Backends = direct.Slice_FromProto(mapCtx, in.Backends, LoadBalancerBackend_FromProto)
	out.BackendType = direct.Enum_FromProto(mapCtx, in.GetBackendType())
	out.BackendURI = direct.LazyPtr(in.GetBackendUri())
	return out
}
func LoadBalancerInfo_ToProto(mapCtx *direct.MapContext, in *krm.LoadBalancerInfo) *pb.LoadBalancerInfo {
	if in == nil {
		return nil
	}
	out := &pb.LoadBalancerInfo{}
	out.LoadBalancerType = direct.Enum_ToProto[pb.LoadBalancerInfo_LoadBalancerType](mapCtx, in.LoadBalancerType)
	out.HealthCheckUri = direct.ValueOf(in.HealthCheckURI)
	out.Backends = direct.Slice_ToProto(mapCtx, in.Backends, LoadBalancerBackend_ToProto)
	out.BackendType = direct.Enum_ToProto[pb.LoadBalancerInfo_BackendType](mapCtx, in.BackendType)
	out.BackendUri = direct.ValueOf(in.BackendURI)
	return out
}
func NatInfo_FromProto(mapCtx *direct.MapContext, in *pb.NatInfo) *krm.NatInfo {
	if in == nil {
		return nil
	}
	out := &krm.NatInfo{}
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
	out.NatGatewayName = direct.LazyPtr(in.GetNatGatewayName())
	return out
}
func NatInfo_ToProto(mapCtx *direct.MapContext, in *krm.NatInfo) *pb.NatInfo {
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
	out.NatGatewayName = direct.ValueOf(in.NatGatewayName)
	return out
}
func NetworkInfo_FromProto(mapCtx *direct.MapContext, in *pb.NetworkInfo) *krm.NetworkInfo {
	if in == nil {
		return nil
	}
	out := &krm.NetworkInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.MatchedSubnetURI = direct.LazyPtr(in.GetMatchedSubnetUri())
	out.MatchedIPRange = direct.LazyPtr(in.GetMatchedIpRange())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func NetworkInfo_ToProto(mapCtx *direct.MapContext, in *krm.NetworkInfo) *pb.NetworkInfo {
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
func ProbingDetails_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails) *krm.ProbingDetails {
	if in == nil {
		return nil
	}
	out := &krm.ProbingDetails{}
	out.Result = direct.Enum_FromProto(mapCtx, in.GetResult())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.AbortCause = direct.Enum_FromProto(mapCtx, in.GetAbortCause())
	out.SentProbeCount = direct.LazyPtr(in.GetSentProbeCount())
	out.SuccessfulProbeCount = direct.LazyPtr(in.GetSuccessfulProbeCount())
	out.EndpointInfo = EndpointInfo_FromProto(mapCtx, in.GetEndpointInfo())
	out.ProbingLatency = LatencyDistribution_FromProto(mapCtx, in.GetProbingLatency())
	out.DestinationEgressLocation = ProbingDetails_EdgeLocation_FromProto(mapCtx, in.GetDestinationEgressLocation())
	return out
}
func ProbingDetails_ToProto(mapCtx *direct.MapContext, in *krm.ProbingDetails) *pb.ProbingDetails {
	if in == nil {
		return nil
	}
	out := &pb.ProbingDetails{}
	out.Result = direct.Enum_ToProto[pb.ProbingDetails_ProbingResult](mapCtx, in.Result)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.AbortCause = direct.Enum_ToProto[pb.ProbingDetails_ProbingAbortCause](mapCtx, in.AbortCause)
	out.SentProbeCount = direct.ValueOf(in.SentProbeCount)
	out.SuccessfulProbeCount = direct.ValueOf(in.SuccessfulProbeCount)
	out.EndpointInfo = EndpointInfo_ToProto(mapCtx, in.EndpointInfo)
	out.ProbingLatency = LatencyDistribution_ToProto(mapCtx, in.ProbingLatency)
	out.DestinationEgressLocation = ProbingDetails_EdgeLocation_ToProto(mapCtx, in.DestinationEgressLocation)
	return out
}
func ProbingDetails_EdgeLocation_FromProto(mapCtx *direct.MapContext, in *pb.ProbingDetails_EdgeLocation) *krm.ProbingDetails_EdgeLocation {
	if in == nil {
		return nil
	}
	out := &krm.ProbingDetails_EdgeLocation{}
	out.MetropolitanArea = direct.LazyPtr(in.GetMetropolitanArea())
	return out
}
func ProbingDetails_EdgeLocation_ToProto(mapCtx *direct.MapContext, in *krm.ProbingDetails_EdgeLocation) *pb.ProbingDetails_EdgeLocation {
	if in == nil {
		return nil
	}
	out := &pb.ProbingDetails_EdgeLocation{}
	out.MetropolitanArea = direct.ValueOf(in.MetropolitanArea)
	return out
}
func ProxyConnectionInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProxyConnectionInfo) *krm.ProxyConnectionInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProxyConnectionInfo{}
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
func ProxyConnectionInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProxyConnectionInfo) *pb.ProxyConnectionInfo {
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
func ReachabilityDetails_FromProto(mapCtx *direct.MapContext, in *pb.ReachabilityDetails) *krm.ReachabilityDetails {
	if in == nil {
		return nil
	}
	out := &krm.ReachabilityDetails{}
	out.Result = direct.Enum_FromProto(mapCtx, in.GetResult())
	out.VerifyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetVerifyTime())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	out.Traces = direct.Slice_FromProto(mapCtx, in.Traces, Trace_FromProto)
	return out
}
func ReachabilityDetails_ToProto(mapCtx *direct.MapContext, in *krm.ReachabilityDetails) *pb.ReachabilityDetails {
	if in == nil {
		return nil
	}
	out := &pb.ReachabilityDetails{}
	out.Result = direct.Enum_ToProto[pb.ReachabilityDetails_Result](mapCtx, in.Result)
	out.VerifyTime = direct.StringTimestamp_ToProto(mapCtx, in.VerifyTime)
	out.Error = Status_ToProto(mapCtx, in.Error)
	out.Traces = direct.Slice_ToProto(mapCtx, in.Traces, Trace_ToProto)
	return out
}
func ReachabilityDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReachabilityDetails) *krm.ReachabilityDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReachabilityDetailsObservedState{}
	// MISSING: Result
	// MISSING: VerifyTime
	// MISSING: Error
	out.Traces = direct.Slice_FromProto(mapCtx, in.Traces, TraceObservedState_FromProto)
	return out
}
func ReachabilityDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReachabilityDetailsObservedState) *pb.ReachabilityDetails {
	if in == nil {
		return nil
	}
	out := &pb.ReachabilityDetails{}
	// MISSING: Result
	// MISSING: VerifyTime
	// MISSING: Error
	out.Traces = direct.Slice_ToProto(mapCtx, in.Traces, TraceObservedState_ToProto)
	return out
}
func RedisClusterInfo_FromProto(mapCtx *direct.MapContext, in *pb.RedisClusterInfo) *krm.RedisClusterInfo {
	if in == nil {
		return nil
	}
	out := &krm.RedisClusterInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.DiscoveryEndpointIPAddress = direct.LazyPtr(in.GetDiscoveryEndpointIpAddress())
	out.SecondaryEndpointIPAddress = direct.LazyPtr(in.GetSecondaryEndpointIpAddress())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func RedisClusterInfo_ToProto(mapCtx *direct.MapContext, in *krm.RedisClusterInfo) *pb.RedisClusterInfo {
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
func RedisInstanceInfo_FromProto(mapCtx *direct.MapContext, in *pb.RedisInstanceInfo) *krm.RedisInstanceInfo {
	if in == nil {
		return nil
	}
	out := &krm.RedisInstanceInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.PrimaryEndpointIP = direct.LazyPtr(in.GetPrimaryEndpointIp())
	out.ReadEndpointIP = direct.LazyPtr(in.GetReadEndpointIp())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func RedisInstanceInfo_ToProto(mapCtx *direct.MapContext, in *krm.RedisInstanceInfo) *pb.RedisInstanceInfo {
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
func RouteInfo_FromProto(mapCtx *direct.MapContext, in *pb.RouteInfo) *krm.RouteInfo {
	if in == nil {
		return nil
	}
	out := &krm.RouteInfo{}
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
func RouteInfo_ToProto(mapCtx *direct.MapContext, in *krm.RouteInfo) *pb.RouteInfo {
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
func ServerlessNegInfo_FromProto(mapCtx *direct.MapContext, in *pb.ServerlessNegInfo) *krm.ServerlessNegInfo {
	if in == nil {
		return nil
	}
	out := &krm.ServerlessNegInfo{}
	out.NegURI = direct.LazyPtr(in.GetNegUri())
	return out
}
func ServerlessNegInfo_ToProto(mapCtx *direct.MapContext, in *krm.ServerlessNegInfo) *pb.ServerlessNegInfo {
	if in == nil {
		return nil
	}
	out := &pb.ServerlessNegInfo{}
	out.NegUri = direct.ValueOf(in.NegURI)
	return out
}
func Step_FromProto(mapCtx *direct.MapContext, in *pb.Step) *krm.Step {
	if in == nil {
		return nil
	}
	out := &krm.Step{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CausesDrop = direct.LazyPtr(in.GetCausesDrop())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.Instance = InstanceInfo_FromProto(mapCtx, in.GetInstance())
	out.Firewall = FirewallInfo_FromProto(mapCtx, in.GetFirewall())
	out.Route = RouteInfo_FromProto(mapCtx, in.GetRoute())
	out.Endpoint = EndpointInfo_FromProto(mapCtx, in.GetEndpoint())
	out.GoogleService = GoogleServiceInfo_FromProto(mapCtx, in.GetGoogleService())
	out.ForwardingRule = ForwardingRuleInfo_FromProto(mapCtx, in.GetForwardingRule())
	out.VpnGateway = VpnGatewayInfo_FromProto(mapCtx, in.GetVpnGateway())
	out.VpnTunnel = VpnTunnelInfo_FromProto(mapCtx, in.GetVpnTunnel())
	out.VpcConnector = VpcConnectorInfo_FromProto(mapCtx, in.GetVpcConnector())
	out.Deliver = DeliverInfo_FromProto(mapCtx, in.GetDeliver())
	out.Forward = ForwardInfo_FromProto(mapCtx, in.GetForward())
	out.Abort = AbortInfo_FromProto(mapCtx, in.GetAbort())
	out.Drop = DropInfo_FromProto(mapCtx, in.GetDrop())
	out.LoadBalancer = LoadBalancerInfo_FromProto(mapCtx, in.GetLoadBalancer())
	out.Network = NetworkInfo_FromProto(mapCtx, in.GetNetwork())
	out.GkeMaster = GKEMasterInfo_FromProto(mapCtx, in.GetGkeMaster())
	out.CloudSqlInstance = CloudSQLInstanceInfo_FromProto(mapCtx, in.GetCloudSqlInstance())
	out.RedisInstance = RedisInstanceInfo_FromProto(mapCtx, in.GetRedisInstance())
	out.RedisCluster = RedisClusterInfo_FromProto(mapCtx, in.GetRedisCluster())
	out.CloudFunction = CloudFunctionInfo_FromProto(mapCtx, in.GetCloudFunction())
	out.AppEngineVersion = AppEngineVersionInfo_FromProto(mapCtx, in.GetAppEngineVersion())
	out.CloudRunRevision = CloudRunRevisionInfo_FromProto(mapCtx, in.GetCloudRunRevision())
	out.Nat = NatInfo_FromProto(mapCtx, in.GetNat())
	out.ProxyConnection = ProxyConnectionInfo_FromProto(mapCtx, in.GetProxyConnection())
	out.LoadBalancerBackendInfo = LoadBalancerBackendInfo_FromProto(mapCtx, in.GetLoadBalancerBackendInfo())
	out.StorageBucket = StorageBucketInfo_FromProto(mapCtx, in.GetStorageBucket())
	out.ServerlessNeg = ServerlessNegInfo_FromProto(mapCtx, in.GetServerlessNeg())
	return out
}
func Step_ToProto(mapCtx *direct.MapContext, in *krm.Step) *pb.Step {
	if in == nil {
		return nil
	}
	out := &pb.Step{}
	out.Description = direct.ValueOf(in.Description)
	out.State = direct.Enum_ToProto[pb.Step_State](mapCtx, in.State)
	out.CausesDrop = direct.ValueOf(in.CausesDrop)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	if oneof := InstanceInfo_ToProto(mapCtx, in.Instance); oneof != nil {
		out.StepInfo = &pb.Step_Instance{Instance: oneof}
	}
	if oneof := FirewallInfo_ToProto(mapCtx, in.Firewall); oneof != nil {
		out.StepInfo = &pb.Step_Firewall{Firewall: oneof}
	}
	if oneof := RouteInfo_ToProto(mapCtx, in.Route); oneof != nil {
		out.StepInfo = &pb.Step_Route{Route: oneof}
	}
	if oneof := EndpointInfo_ToProto(mapCtx, in.Endpoint); oneof != nil {
		out.StepInfo = &pb.Step_Endpoint{Endpoint: oneof}
	}
	if oneof := GoogleServiceInfo_ToProto(mapCtx, in.GoogleService); oneof != nil {
		out.StepInfo = &pb.Step_GoogleService{GoogleService: oneof}
	}
	if oneof := ForwardingRuleInfo_ToProto(mapCtx, in.ForwardingRule); oneof != nil {
		out.StepInfo = &pb.Step_ForwardingRule{ForwardingRule: oneof}
	}
	if oneof := VpnGatewayInfo_ToProto(mapCtx, in.VpnGateway); oneof != nil {
		out.StepInfo = &pb.Step_VpnGateway{VpnGateway: oneof}
	}
	if oneof := VpnTunnelInfo_ToProto(mapCtx, in.VpnTunnel); oneof != nil {
		out.StepInfo = &pb.Step_VpnTunnel{VpnTunnel: oneof}
	}
	if oneof := VpcConnectorInfo_ToProto(mapCtx, in.VpcConnector); oneof != nil {
		out.StepInfo = &pb.Step_VpcConnector{VpcConnector: oneof}
	}
	if oneof := DeliverInfo_ToProto(mapCtx, in.Deliver); oneof != nil {
		out.StepInfo = &pb.Step_Deliver{Deliver: oneof}
	}
	if oneof := ForwardInfo_ToProto(mapCtx, in.Forward); oneof != nil {
		out.StepInfo = &pb.Step_Forward{Forward: oneof}
	}
	if oneof := AbortInfo_ToProto(mapCtx, in.Abort); oneof != nil {
		out.StepInfo = &pb.Step_Abort{Abort: oneof}
	}
	if oneof := DropInfo_ToProto(mapCtx, in.Drop); oneof != nil {
		out.StepInfo = &pb.Step_Drop{Drop: oneof}
	}
	if oneof := LoadBalancerInfo_ToProto(mapCtx, in.LoadBalancer); oneof != nil {
		out.StepInfo = &pb.Step_LoadBalancer{LoadBalancer: oneof}
	}
	if oneof := NetworkInfo_ToProto(mapCtx, in.Network); oneof != nil {
		out.StepInfo = &pb.Step_Network{Network: oneof}
	}
	if oneof := GKEMasterInfo_ToProto(mapCtx, in.GkeMaster); oneof != nil {
		out.StepInfo = &pb.Step_GkeMaster{GkeMaster: oneof}
	}
	if oneof := CloudSQLInstanceInfo_ToProto(mapCtx, in.CloudSqlInstance); oneof != nil {
		out.StepInfo = &pb.Step_CloudSqlInstance{CloudSqlInstance: oneof}
	}
	if oneof := RedisInstanceInfo_ToProto(mapCtx, in.RedisInstance); oneof != nil {
		out.StepInfo = &pb.Step_RedisInstance{RedisInstance: oneof}
	}
	if oneof := RedisClusterInfo_ToProto(mapCtx, in.RedisCluster); oneof != nil {
		out.StepInfo = &pb.Step_RedisCluster{RedisCluster: oneof}
	}
	if oneof := CloudFunctionInfo_ToProto(mapCtx, in.CloudFunction); oneof != nil {
		out.StepInfo = &pb.Step_CloudFunction{CloudFunction: oneof}
	}
	if oneof := AppEngineVersionInfo_ToProto(mapCtx, in.AppEngineVersion); oneof != nil {
		out.StepInfo = &pb.Step_AppEngineVersion{AppEngineVersion: oneof}
	}
	if oneof := CloudRunRevisionInfo_ToProto(mapCtx, in.CloudRunRevision); oneof != nil {
		out.StepInfo = &pb.Step_CloudRunRevision{CloudRunRevision: oneof}
	}
	if oneof := NatInfo_ToProto(mapCtx, in.Nat); oneof != nil {
		out.StepInfo = &pb.Step_Nat{Nat: oneof}
	}
	if oneof := ProxyConnectionInfo_ToProto(mapCtx, in.ProxyConnection); oneof != nil {
		out.StepInfo = &pb.Step_ProxyConnection{ProxyConnection: oneof}
	}
	if oneof := LoadBalancerBackendInfo_ToProto(mapCtx, in.LoadBalancerBackendInfo); oneof != nil {
		out.StepInfo = &pb.Step_LoadBalancerBackendInfo{LoadBalancerBackendInfo: oneof}
	}
	if oneof := StorageBucketInfo_ToProto(mapCtx, in.StorageBucket); oneof != nil {
		out.StepInfo = &pb.Step_StorageBucket{StorageBucket: oneof}
	}
	if oneof := ServerlessNegInfo_ToProto(mapCtx, in.ServerlessNeg); oneof != nil {
		out.StepInfo = &pb.Step_ServerlessNeg{ServerlessNeg: oneof}
	}
	return out
}
func StepObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Step) *krm.StepObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StepObservedState{}
	// MISSING: Description
	// MISSING: State
	// MISSING: CausesDrop
	// MISSING: ProjectID
	// MISSING: Instance
	// MISSING: Firewall
	// MISSING: Route
	// MISSING: Endpoint
	// MISSING: GoogleService
	// MISSING: ForwardingRule
	// MISSING: VpnGateway
	// MISSING: VpnTunnel
	// MISSING: VpcConnector
	// MISSING: Deliver
	// MISSING: Forward
	// MISSING: Abort
	// MISSING: Drop
	// MISSING: LoadBalancer
	// MISSING: Network
	// MISSING: GkeMaster
	// MISSING: CloudSqlInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Nat
	// MISSING: ProxyConnection
	out.LoadBalancerBackendInfo = LoadBalancerBackendInfoObservedState_FromProto(mapCtx, in.GetLoadBalancerBackendInfo())
	// MISSING: StorageBucket
	// MISSING: ServerlessNeg
	return out
}
func StepObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StepObservedState) *pb.Step {
	if in == nil {
		return nil
	}
	out := &pb.Step{}
	// MISSING: Description
	// MISSING: State
	// MISSING: CausesDrop
	// MISSING: ProjectID
	// MISSING: Instance
	// MISSING: Firewall
	// MISSING: Route
	// MISSING: Endpoint
	// MISSING: GoogleService
	// MISSING: ForwardingRule
	// MISSING: VpnGateway
	// MISSING: VpnTunnel
	// MISSING: VpcConnector
	// MISSING: Deliver
	// MISSING: Forward
	// MISSING: Abort
	// MISSING: Drop
	// MISSING: LoadBalancer
	// MISSING: Network
	// MISSING: GkeMaster
	// MISSING: CloudSqlInstance
	// MISSING: RedisInstance
	// MISSING: RedisCluster
	// MISSING: CloudFunction
	// MISSING: AppEngineVersion
	// MISSING: CloudRunRevision
	// MISSING: Nat
	// MISSING: ProxyConnection
	if oneof := LoadBalancerBackendInfoObservedState_ToProto(mapCtx, in.LoadBalancerBackendInfo); oneof != nil {
		out.StepInfo = &pb.Step_LoadBalancerBackendInfo{LoadBalancerBackendInfo: oneof}
	}
	// MISSING: StorageBucket
	// MISSING: ServerlessNeg
	return out
}
func StorageBucketInfo_FromProto(mapCtx *direct.MapContext, in *pb.StorageBucketInfo) *krm.StorageBucketInfo {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketInfo{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	return out
}
func StorageBucketInfo_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketInfo) *pb.StorageBucketInfo {
	if in == nil {
		return nil
	}
	out := &pb.StorageBucketInfo{}
	out.Bucket = direct.ValueOf(in.Bucket)
	return out
}
func Trace_FromProto(mapCtx *direct.MapContext, in *pb.Trace) *krm.Trace {
	if in == nil {
		return nil
	}
	out := &krm.Trace{}
	out.EndpointInfo = EndpointInfo_FromProto(mapCtx, in.GetEndpointInfo())
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, Step_FromProto)
	out.ForwardTraceID = direct.LazyPtr(in.GetForwardTraceId())
	return out
}
func Trace_ToProto(mapCtx *direct.MapContext, in *krm.Trace) *pb.Trace {
	if in == nil {
		return nil
	}
	out := &pb.Trace{}
	out.EndpointInfo = EndpointInfo_ToProto(mapCtx, in.EndpointInfo)
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, Step_ToProto)
	out.ForwardTraceId = direct.ValueOf(in.ForwardTraceID)
	return out
}
func TraceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trace) *krm.TraceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TraceObservedState{}
	// MISSING: EndpointInfo
	out.Steps = direct.Slice_FromProto(mapCtx, in.Steps, StepObservedState_FromProto)
	// MISSING: ForwardTraceID
	return out
}
func TraceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TraceObservedState) *pb.Trace {
	if in == nil {
		return nil
	}
	out := &pb.Trace{}
	// MISSING: EndpointInfo
	out.Steps = direct.Slice_ToProto(mapCtx, in.Steps, StepObservedState_ToProto)
	// MISSING: ForwardTraceID
	return out
}
func VpcConnectorInfo_FromProto(mapCtx *direct.MapContext, in *pb.VpcConnectorInfo) *krm.VpcConnectorInfo {
	if in == nil {
		return nil
	}
	out := &krm.VpcConnectorInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func VpcConnectorInfo_ToProto(mapCtx *direct.MapContext, in *krm.VpcConnectorInfo) *pb.VpcConnectorInfo {
	if in == nil {
		return nil
	}
	out := &pb.VpcConnectorInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.Location = direct.ValueOf(in.Location)
	return out
}
func VpnGatewayInfo_FromProto(mapCtx *direct.MapContext, in *pb.VpnGatewayInfo) *krm.VpnGatewayInfo {
	if in == nil {
		return nil
	}
	out := &krm.VpnGatewayInfo{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.URI = direct.LazyPtr(in.GetUri())
	out.NetworkURI = direct.LazyPtr(in.GetNetworkUri())
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.VpnTunnelURI = direct.LazyPtr(in.GetVpnTunnelUri())
	out.Region = direct.LazyPtr(in.GetRegion())
	return out
}
func VpnGatewayInfo_ToProto(mapCtx *direct.MapContext, in *krm.VpnGatewayInfo) *pb.VpnGatewayInfo {
	if in == nil {
		return nil
	}
	out := &pb.VpnGatewayInfo{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Uri = direct.ValueOf(in.URI)
	out.NetworkUri = direct.ValueOf(in.NetworkURI)
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.VpnTunnelUri = direct.ValueOf(in.VpnTunnelURI)
	out.Region = direct.ValueOf(in.Region)
	return out
}
func VpnTunnelInfo_FromProto(mapCtx *direct.MapContext, in *pb.VpnTunnelInfo) *krm.VpnTunnelInfo {
	if in == nil {
		return nil
	}
	out := &krm.VpnTunnelInfo{}
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
func VpnTunnelInfo_ToProto(mapCtx *direct.MapContext, in *krm.VpnTunnelInfo) *pb.VpnTunnelInfo {
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
