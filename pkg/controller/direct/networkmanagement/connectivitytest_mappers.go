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
	pb "cloud.google.com/go/networkmanagement/apiv1/networkmanagementpb"
	compute "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	container "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkmanagement/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	run "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func EndpointObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.EndpointObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EndpointObservedState{}
	// MISSING: IPAddress
	// MISSING: Port
	// MISSING: Instance
	// MISSING: ForwardingRule
	if in.ForwardingRuleTarget != nil {
		out.ForwardingRuleTarget = direct.ZeroBasedEnum_FromProto(mapCtx, in.GetForwardingRuleTarget())
	}
	out.LoadBalancerID = in.LoadBalancerId
	if in.LoadBalancerType != nil {
		out.LoadBalancerType = direct.ZeroBasedEnum_FromProto(mapCtx, in.GetLoadBalancerType())
	}
	// MISSING: GKEMasterCluster
	// MISSING: FQDN
	// MISSING: CloudSQLInstance
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
	if in.ForwardingRuleTarget != nil {
		out.ForwardingRuleTarget = direct.PtrTo(direct.Enum_ToProto[pb.Endpoint_ForwardingRuleTarget](mapCtx, in.ForwardingRuleTarget))
	}
	out.LoadBalancerId = in.LoadBalancerID
	if in.LoadBalancerType != nil {
		out.LoadBalancerType = direct.PtrTo(direct.Enum_ToProto[pb.LoadBalancerType](mapCtx, in.LoadBalancerType))
	}
	// MISSING: GKEMasterCluster
	// MISSING: FQDN
	// MISSING: CloudSQLInstance
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
func StatusObservedState_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.StatusObservedState {
	if in == nil {
		return nil
	}

	out := &krm.StatusObservedState{
		Code:    direct.LazyPtr(in.Code),
		Message: direct.LazyPtr(in.Message),
	}
	if len(in.Details) == 0 {
		return out
	}
	detailsOut := make([]krm.Any, 0)
	for _, d := range in.Details {
		if d == nil {
			continue
		}
		dOut := krm.Any{
			TypeURL: direct.LazyPtr(d.TypeUrl),
			Value:   direct.ByteSliceToStringPtr(mapCtx, d.Value),
		}
		detailsOut = append(detailsOut, dOut)
	}
	if len(detailsOut) > 0 {
		out.Details = detailsOut
	}
	return out
}
func StatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StatusObservedState) *status.Status {
	if in == nil {
		return nil
	}

	out := &status.Status{
		Code:    direct.ValueOf(in.Code),
		Message: direct.ValueOf(in.Message),
	}
	if len(in.Details) == 0 {
		return out
	}
	detailsOut := make([]*anypb.Any, 0)
	for _, d := range in.Details {
		dOut := &anypb.Any{
			TypeUrl: direct.ValueOf(d.TypeURL),
			Value:   direct.StringPtrToByteSlice(mapCtx, d.Value),
		}
		detailsOut = append(detailsOut, dOut)
	}
	if len(detailsOut) > 0 {
		out.Details = detailsOut
	}
	return out
}
func NetworkManagementConnectivityTestSpec_RelatedProjects_FromProto(mapCtx *direct.MapContext, in []string) []refs.ProjectRef {
	if len(in) == 0 {
		return nil
	}
	out := make([]refs.ProjectRef, 0)
	for _, p := range in {
		if p == "" {
			continue
		}
		projectRef := refs.ProjectRef{External: p}
		out = append(out, projectRef)
	}
	return out
}
func NetworkManagementConnectivityTestSpec_RelatedProjects_ToProto(mapCtx *direct.MapContext, in []refs.ProjectRef) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, 0)
	for _, p := range in {
		if p.External == "" {
			continue
		}
		out = append(out, p.External)
	}
	return out
}

func Endpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint) *krm.Endpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint{}
	out.IPAddress = direct.LazyPtr(in.GetIpAddress())
	out.Port = direct.LazyPtr(in.GetPort())
	if in.GetInstance() != "" {
		out.ComputeInstanceRef = &compute.InstanceRef{External: in.GetInstance()}
	}
	out.ComputeForwardingRuleRef = direct.LazyPtr(in.GetForwardingRule())
	// MISSING: ForwardingRuleTarget
	// MISSING: LoadBalancerID
	// MISSING: LoadBalancerType
	if in.GetGkeMasterCluster() != "" {
		out.ContainerClusterRef = &container.ContainerClusterRef{External: in.GetGkeMasterCluster()}
	}
	out.FQDN = direct.LazyPtr(in.GetFqdn())
	if in.GetCloudSqlInstance() != "" {
		out.SQLInstanceRef = &sqlv1beta1.SQLInstanceRef{External: in.GetCloudSqlInstance()}
	}
	out.RedisInstance = direct.LazyPtr(in.GetRedisInstance())
	out.RedisCluster = direct.LazyPtr(in.GetRedisCluster())
	out.CloudFunction = Endpoint_CloudFunctionEndpoint_FromProto(mapCtx, in.GetCloudFunction())
	out.AppEngineVersion = Endpoint_AppEngineVersionEndpoint_FromProto(mapCtx, in.GetAppEngineVersion())
	out.CloudRunRevision = Endpoint_CloudRunRevisionEndpoint_FromProto(mapCtx, in.GetCloudRunRevision())
	if in.GetNetwork() != "" {
		out.ComputeNetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.NetworkType = direct.Enum_FromProto(mapCtx, in.GetNetworkType())
	if in.GetProjectId() != "" {
		out.ProjectRef = &refs.ProjectRef{External: in.GetProjectId()}
	}
	return out
}
func Endpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint) *pb.Endpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint{}
	out.IpAddress = direct.ValueOf(in.IPAddress)
	out.Port = direct.ValueOf(in.Port)
	if in.ComputeInstanceRef != nil {
		out.Instance = in.ComputeInstanceRef.External
	}
	out.ForwardingRule = direct.ValueOf(in.ComputeForwardingRuleRef)
	// MISSING: ForwardingRuleTarget
	// MISSING: LoadBalancerID
	// MISSING: LoadBalancerType
	if in.ContainerClusterRef != nil {
		out.GkeMasterCluster = in.ContainerClusterRef.External
	}
	out.Fqdn = direct.ValueOf(in.FQDN)
	if in.SQLInstanceRef != nil {
		out.CloudSqlInstance = in.SQLInstanceRef.External
	}
	out.RedisInstance = direct.ValueOf(in.RedisInstance)
	out.RedisCluster = direct.ValueOf(in.RedisCluster)
	out.CloudFunction = Endpoint_CloudFunctionEndpoint_ToProto(mapCtx, in.CloudFunction)
	out.AppEngineVersion = Endpoint_AppEngineVersionEndpoint_ToProto(mapCtx, in.AppEngineVersion)
	out.CloudRunRevision = Endpoint_CloudRunRevisionEndpoint_ToProto(mapCtx, in.CloudRunRevision)
	if in.ComputeNetworkRef != nil {
		out.Network = in.ComputeNetworkRef.External
	}
	out.NetworkType = direct.Enum_ToProto[pb.Endpoint_NetworkType](mapCtx, in.NetworkType)
	if in.ProjectRef != nil {
		out.ProjectId = in.ProjectRef.External
	}
	return out
}

func Endpoint_CloudRunRevisionEndpoint_FromProto(mapCtx *direct.MapContext, in *pb.Endpoint_CloudRunRevisionEndpoint) *krm.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &krm.Endpoint_CloudRunRevisionEndpoint{}
	if in.Uri != "" {
		out.RunRevisionRef = &run.RevisionRef{External: in.GetUri()}
	}
	return out
}
func Endpoint_CloudRunRevisionEndpoint_ToProto(mapCtx *direct.MapContext, in *krm.Endpoint_CloudRunRevisionEndpoint) *pb.Endpoint_CloudRunRevisionEndpoint {
	if in == nil {
		return nil
	}
	out := &pb.Endpoint_CloudRunRevisionEndpoint{}
	if in.RunRevisionRef != nil {
		out.Uri = in.RunRevisionRef.External
	}
	return out
}
