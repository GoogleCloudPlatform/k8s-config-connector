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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackendServiceBackend_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Backend) *krm.BackendServiceBackend {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceBackend{}
	out.BalancingMode = in.BalancingMode
	if in.CapacityScaler != nil {
		v := float64(*in.CapacityScaler)
		out.CapacityScaler = &v
	}
	out.Description = in.Description
	out.Failover = in.Failover
	if in.MaxConnections != nil {
		v := int(*in.MaxConnections)
		out.MaxConnections = &v
	}
	if in.MaxConnectionsPerEndpoint != nil {
		v := int(*in.MaxConnectionsPerEndpoint)
		out.MaxConnectionsPerEndpoint = &v
	}
	if in.MaxConnectionsPerInstance != nil {
		v := int(*in.MaxConnectionsPerInstance)
		out.MaxConnectionsPerInstance = &v
	}
	if in.MaxRate != nil {
		v := int(*in.MaxRate)
		out.MaxRate = &v
	}
	if in.MaxRatePerEndpoint != nil {
		v := float64(*in.MaxRatePerEndpoint)
		out.MaxRatePerEndpoint = &v
	}
	if in.MaxRatePerInstance != nil {
		v := float64(*in.MaxRatePerInstance)
		out.MaxRatePerInstance = &v
	}
	if in.MaxUtilization != nil {
		v := float64(*in.MaxUtilization)
		out.MaxUtilization = &v
	}
	return out
}

func BackendServiceBackend_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceBackend) *pb.Backend {
	if in == nil {
		return nil
	}
	out := &pb.Backend{}
	out.BalancingMode = in.BalancingMode
	if in.CapacityScaler != nil {
		v := float32(*in.CapacityScaler)
		out.CapacityScaler = &v
	}
	out.Description = in.Description
	out.Failover = in.Failover
	if in.MaxConnections != nil {
		v := int32(*in.MaxConnections)
		out.MaxConnections = &v
	}
	if in.MaxConnectionsPerEndpoint != nil {
		v := int32(*in.MaxConnectionsPerEndpoint)
		out.MaxConnectionsPerEndpoint = &v
	}
	if in.MaxConnectionsPerInstance != nil {
		v := int32(*in.MaxConnectionsPerInstance)
		out.MaxConnectionsPerInstance = &v
	}
	if in.MaxRate != nil {
		v := int32(*in.MaxRate)
		out.MaxRate = &v
	}
	if in.MaxRatePerEndpoint != nil {
		v := float32(*in.MaxRatePerEndpoint)
		out.MaxRatePerEndpoint = &v
	}
	if in.MaxRatePerInstance != nil {
		v := float32(*in.MaxRatePerInstance)
		out.MaxRatePerInstance = &v
	}
	if in.MaxUtilization != nil {
		v := float32(*in.MaxUtilization)
		out.MaxUtilization = &v
	}
	return out
}

func BackendServiceBaseEjectionTime_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.BackendServiceBaseEjectionTime {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceBaseEjectionTime{}
	if in.Nanos != nil {
		v := int64(*in.Nanos)
		out.Nanos = &v
	}
	if in.Seconds != nil {
		out.Seconds = int64(*in.Seconds)
	}
	return out
}
func BackendServiceBaseEjectionTime_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceBaseEjectionTime) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		v := int32(*in.Nanos)
		out.Nanos = &v
	}
	out.Seconds = &in.Seconds
	return out
}

func BackendServiceConnectTimeout_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.BackendServiceConnectTimeout {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceConnectTimeout{}
	if in.Nanos != nil {
		v := int64(*in.Nanos)
		out.Nanos = &v
	}
	if in.Seconds != nil {
		out.Seconds = int64(*in.Seconds)
	}
	return out
}
func BackendServiceConnectTimeout_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceConnectTimeout) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		v := int32(*in.Nanos)
		out.Nanos = &v
	}
	out.Seconds = &in.Seconds
	return out
}

func BackendServiceTtl_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.BackendServiceTtl {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceTtl{}
	if in.Nanos != nil {
		v := int64(*in.Nanos)
		out.Nanos = &v
	}
	if in.Seconds != nil {
		out.Seconds = int64(*in.Seconds)
	}
	return out
}
func BackendServiceTtl_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceTtl) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		v := int32(*in.Nanos)
		out.Nanos = &v
	}
	out.Seconds = &in.Seconds
	return out
}

func BackendServiceInterval_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.BackendServiceInterval {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceInterval{}
	if in.Nanos != nil {
		v := int64(*in.Nanos)
		out.Nanos = &v
	}
	if in.Seconds != nil {
		out.Seconds = int64(*in.Seconds)
	}
	return out
}
func BackendServiceInterval_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceInterval) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	if in.Nanos != nil {
		v := int32(*in.Nanos)
		out.Nanos = &v
	}
	out.Seconds = &in.Seconds
	return out
}

func BackendServiceBypassCacheOnRequestHeaders_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader) *krm.BackendServiceBypassCacheOnRequestHeaders {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceBypassCacheOnRequestHeaders{}
	if in.HeaderName != nil {
		out.HeaderName = *in.HeaderName
	}
	return out
}
func BackendServiceBypassCacheOnRequestHeaders_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceBypassCacheOnRequestHeaders) *pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader{}
	out.HeaderName = &in.HeaderName
	return out
}

func BackendServiceCircuitBreakers_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CircuitBreakers) *krm.BackendServiceCircuitBreakers {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCircuitBreakers{}
	// out.ConnectTimeout = BackendServiceConnectTimeout_v1beta1_FromProto(mapCtx, in.GetConnectTimeout())
	if in.MaxConnections != nil {
		v := int(*in.MaxConnections)
		out.MaxConnections = &v
	}
	if in.MaxPendingRequests != nil {
		v := int(*in.MaxPendingRequests)
		out.MaxPendingRequests = &v
	}
	if in.MaxRequests != nil {
		v := int(*in.MaxRequests)
		out.MaxRequests = &v
	}
	if in.MaxRequestsPerConnection != nil {
		v := int(*in.MaxRequestsPerConnection)
		out.MaxRequestsPerConnection = &v
	}
	if in.MaxRetries != nil {
		v := int(*in.MaxRetries)
		out.MaxRetries = &v
	}
	return out
}
func BackendServiceCircuitBreakers_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCircuitBreakers) *pb.CircuitBreakers {
	if in == nil {
		return nil
	}
	out := &pb.CircuitBreakers{}
	// out.ConnectTimeout = BackendServiceConnectTimeout_v1beta1_ToProto(mapCtx, in.ConnectTimeout)
	if in.MaxConnections != nil {
		v := int32(*in.MaxConnections)
		out.MaxConnections = &v
	}
	if in.MaxPendingRequests != nil {
		v := int32(*in.MaxPendingRequests)
		out.MaxPendingRequests = &v
	}
	if in.MaxRequests != nil {
		v := int32(*in.MaxRequests)
		out.MaxRequests = &v
	}
	if in.MaxRequestsPerConnection != nil {
		v := int32(*in.MaxRequestsPerConnection)
		out.MaxRequestsPerConnection = &v
	}
	if in.MaxRetries != nil {
		v := int32(*in.MaxRetries)
		out.MaxRetries = &v
	}
	return out
}

func BackendServiceConnectionTrackingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceConnectionTrackingPolicy) *krm.BackendServiceConnectionTrackingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceConnectionTrackingPolicy{}
	out.ConnectionPersistenceOnUnhealthyBackends = in.ConnectionPersistenceOnUnhealthyBackends
	out.EnableStrongAffinity = in.EnableStrongAffinity
	if in.IdleTimeoutSec != nil {
		v := int(*in.IdleTimeoutSec)
		out.IdleTimeoutSec = &v
	}
	out.TrackingMode = in.TrackingMode
	return out
}
func BackendServiceConnectionTrackingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceConnectionTrackingPolicy) *pb.BackendServiceConnectionTrackingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceConnectionTrackingPolicy{}
	out.ConnectionPersistenceOnUnhealthyBackends = in.ConnectionPersistenceOnUnhealthyBackends
	out.EnableStrongAffinity = in.EnableStrongAffinity
	if in.IdleTimeoutSec != nil {
		v := int32(*in.IdleTimeoutSec)
		out.IdleTimeoutSec = &v
	}
	out.TrackingMode = in.TrackingMode
	return out
}

func BackendServiceCustomPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy) *krm.BackendServiceCustomPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCustomPolicy{}
	out.Data = in.Data
	if in.Name != nil {
		out.Name = *in.Name
	}
	return out
}
func BackendServiceCustomPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCustomPolicy) *pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy{}
	out.Data = in.Data
	out.Name = &in.Name
	return out
}

func BackendServicePolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy) *krm.BackendServicePolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServicePolicy{}
	if in.Name != nil {
		out.Name = *in.Name
	}
	return out
}
func BackendServicePolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServicePolicy) *pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy{}
	out.Name = &in.Name
	return out
}

func BackendServiceSubsetting_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Subsetting) *krm.BackendServiceSubsetting {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceSubsetting{}
	if in.Policy != nil {
		out.Policy = *in.Policy
	}
	return out
}
func BackendServiceSubsetting_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceSubsetting) *pb.Subsetting {
	if in == nil {
		return nil
	}
	out := &pb.Subsetting{}
	out.Policy = &in.Policy
	return out
}

func BackendServiceFailoverPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceFailoverPolicy) *krm.BackendServiceFailoverPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceFailoverPolicy{}
	out.DisableConnectionDrainOnFailover = in.DisableConnectionDrainOnFailover
	out.DropTrafficIfUnhealthy = in.DropTrafficIfUnhealthy
	if in.FailoverRatio != nil {
		v := float64(*in.FailoverRatio)
		out.FailoverRatio = &v
	}
	return out
}
func BackendServiceFailoverPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceFailoverPolicy) *pb.BackendServiceFailoverPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceFailoverPolicy{}
	out.DisableConnectionDrainOnFailover = in.DisableConnectionDrainOnFailover
	out.DropTrafficIfUnhealthy = in.DropTrafficIfUnhealthy
	if in.FailoverRatio != nil {
		v := float32(*in.FailoverRatio)
		out.FailoverRatio = &v
	}
	return out
}

func BackendServiceLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLogConfig) *krm.BackendServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLogConfig{}
	out.Enable = in.Enable
	if in.SampleRate != nil {
		v := float64(*in.SampleRate)
		out.SampleRate = &v
	}
	return out
}
func BackendServiceLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLogConfig) *pb.BackendServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLogConfig{}
	out.Enable = in.Enable
	if in.SampleRate != nil {
		v := float32(*in.SampleRate)
		out.SampleRate = &v
	}
	return out
}

func BackendServiceNegativeCachingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicyNegativeCachingPolicy) *krm.BackendServiceNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceNegativeCachingPolicy{}
	if in.Code != nil {
		v := int(*in.Code)
		out.Code = &v
	}
	if in.Ttl != nil {
		v := int(*in.Ttl)
		out.Ttl = &v
	}
	return out
}
func BackendServiceNegativeCachingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceNegativeCachingPolicy) *pb.BackendServiceCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicyNegativeCachingPolicy{}
	if in.Code != nil {
		v := int32(*in.Code)
		out.Code = &v
	}
	if in.Ttl != nil {
		v := int32(*in.Ttl)
		out.Ttl = &v
	}
	return out
}

func BackendServiceOutlierDetection_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.OutlierDetection) *krm.BackendServiceOutlierDetection {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceOutlierDetection{}
	out.BaseEjectionTime = BackendServiceBaseEjectionTime_v1beta1_FromProto(mapCtx, in.GetBaseEjectionTime())
	if in.ConsecutiveErrors != nil {
		v := int(*in.ConsecutiveErrors)
		out.ConsecutiveErrors = &v
	}
	if in.ConsecutiveGatewayFailure != nil {
		v := int(*in.ConsecutiveGatewayFailure)
		out.ConsecutiveGatewayFailure = &v
	}
	if in.EnforcingConsecutiveErrors != nil {
		v := int(*in.EnforcingConsecutiveErrors)
		out.EnforcingConsecutiveErrors = &v
	}
	if in.EnforcingConsecutiveGatewayFailure != nil {
		v := int(*in.EnforcingConsecutiveGatewayFailure)
		out.EnforcingConsecutiveGatewayFailure = &v
	}
	if in.EnforcingSuccessRate != nil {
		v := int(*in.EnforcingSuccessRate)
		out.EnforcingSuccessRate = &v
	}
	out.Interval = BackendServiceInterval_v1beta1_FromProto(mapCtx, in.GetInterval())
	if in.MaxEjectionPercent != nil {
		v := int(*in.MaxEjectionPercent)
		out.MaxEjectionPercent = &v
	}
	if in.SuccessRateMinimumHosts != nil {
		v := int(*in.SuccessRateMinimumHosts)
		out.SuccessRateMinimumHosts = &v
	}
	if in.SuccessRateRequestVolume != nil {
		v := int(*in.SuccessRateRequestVolume)
		out.SuccessRateRequestVolume = &v
	}
	if in.SuccessRateStdevFactor != nil {
		v := int(*in.SuccessRateStdevFactor)
		out.SuccessRateStdevFactor = &v
	}
	return out
}
func BackendServiceOutlierDetection_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceOutlierDetection) *pb.OutlierDetection {
	if in == nil {
		return nil
	}
	out := &pb.OutlierDetection{}
	out.BaseEjectionTime = BackendServiceBaseEjectionTime_v1beta1_ToProto(mapCtx, in.BaseEjectionTime)
	if in.ConsecutiveErrors != nil {
		v := int32(*in.ConsecutiveErrors)
		out.ConsecutiveErrors = &v
	}
	if in.ConsecutiveGatewayFailure != nil {
		v := int32(*in.ConsecutiveGatewayFailure)
		out.ConsecutiveGatewayFailure = &v
	}
	if in.EnforcingConsecutiveErrors != nil {
		v := int32(*in.EnforcingConsecutiveErrors)
		out.EnforcingConsecutiveErrors = &v
	}
	if in.EnforcingConsecutiveGatewayFailure != nil {
		v := int32(*in.EnforcingConsecutiveGatewayFailure)
		out.EnforcingConsecutiveGatewayFailure = &v
	}
	if in.EnforcingSuccessRate != nil {
		v := int32(*in.EnforcingSuccessRate)
		out.EnforcingSuccessRate = &v
	}
	out.Interval = BackendServiceInterval_v1beta1_ToProto(mapCtx, in.Interval)
	if in.MaxEjectionPercent != nil {
		v := int32(*in.MaxEjectionPercent)
		out.MaxEjectionPercent = &v
	}
	if in.SuccessRateMinimumHosts != nil {
		v := int32(*in.SuccessRateMinimumHosts)
		out.SuccessRateMinimumHosts = &v
	}
	if in.SuccessRateRequestVolume != nil {
		v := int32(*in.SuccessRateRequestVolume)
		out.SuccessRateRequestVolume = &v
	}
	if in.SuccessRateStdevFactor != nil {
		v := int32(*in.SuccessRateStdevFactor)
		out.SuccessRateStdevFactor = &v
	}
	return out
}

func BackendServiceSecuritySettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings) *krm.BackendServiceSecuritySettings {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceSecuritySettings{}
	if in.GetClientTlsPolicy() != "" {
		out.ClientTLSPolicyRef = refsv1beta1.NetworkSecurityClientTLSPolicyRef{External: in.GetClientTlsPolicy()}
	}
	out.SubjectAltNames = in.SubjectAltNames
	return out
}
func BackendServiceSecuritySettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceSecuritySettings) *pb.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings{}
	out.ClientTlsPolicy = &in.ClientTLSPolicyRef.External
	out.SubjectAltNames = in.SubjectAltNames
	return out
}

func ComputeBackendServiceSpec_HealthChecks_FromProto(mapCtx *direct.MapContext, in []string) []krm.BackendServiceHealthChecks {
	if len(in) == 0 {
		return nil
	}
	var out []krm.BackendServiceHealthChecks
	for _, v := range in {
		out = append(out, krm.BackendServiceHealthChecks{
			HealthCheckRef: &refsv1beta1.ComputeHealthCheckRef{External: v},
		})
	}
	return out
}
func ComputeBackendServiceSpec_HealthChecks_ToProto(mapCtx *direct.MapContext, in []krm.BackendServiceHealthChecks) []string {
	if len(in) == 0 {
		return nil
	}
	var out []string
	for _, v := range in {
		if v.HealthCheckRef != nil {
			out = append(out, v.HealthCheckRef.External)
		} else if v.HttpHealthCheckRef != nil {
			out = append(out, v.HttpHealthCheckRef.External)
		}
	}
	return out
}

func BackendServiceLocalityLbPolicies_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfig) *krm.BackendServiceLocalityLbPolicies {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLocalityLbPolicies{}
	out.CustomPolicy = BackendServiceCustomPolicy_v1beta1_FromProto(mapCtx, in.GetCustomPolicy())
	out.Policy = BackendServicePolicy_v1beta1_FromProto(mapCtx, in.GetPolicy())
	return out
}
func BackendServiceLocalityLbPolicies_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLocalityLbPolicies) *pb.BackendServiceLocalityLoadBalancingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfig{}
	out.CustomPolicy = BackendServiceCustomPolicy_v1beta1_ToProto(mapCtx, in.CustomPolicy)
	out.Policy = BackendServicePolicy_v1beta1_ToProto(mapCtx, in.Policy)
	return out
}

func ComputeBackendServiceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendService) *krm.ComputeBackendServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendServiceSpec{}

	if in.AffinityCookieTtlSec != nil {
		v := int(*in.AffinityCookieTtlSec)
		out.AffinityCookieTtlSec = &v
	}
	out.Backend = direct.Slice_FromProto(mapCtx, in.Backends, BackendServiceBackend_v1beta1_FromProto)
	out.CdnPolicy = BackendServiceCdnPolicy_v1beta1_FromProto(mapCtx, in.GetCdnPolicy())
	out.EnableCdn = in.EnableCDN
	out.Iap = BackendServiceIap_v1beta1_FromProto(mapCtx, in.GetIap())
	out.ResourceID = in.Name
	out.CircuitBreakers = BackendServiceCircuitBreakers_v1beta1_FromProto(mapCtx, in.GetCircuitBreakers())
	out.CompressionMode = in.CompressionMode
	if in.ConnectionDraining != nil && in.ConnectionDraining.DrainingTimeoutSec != nil {
		v := int(*in.ConnectionDraining.DrainingTimeoutSec)
		out.ConnectionDrainingTimeoutSec = &v
	}
	out.ConnectionTrackingPolicy = BackendServiceConnectionTrackingPolicy_v1beta1_FromProto(mapCtx, in.GetConnectionTrackingPolicy())
	out.ConsistentHash = BackendServiceConsistentHash_v1beta1_FromProto(mapCtx, in.GetConsistentHash())
	out.CustomRequestHeaders = in.CustomRequestHeaders
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	if in.GetEdgeSecurityPolicy() != "" {
		out.EdgeSecurityPolicyRef = &refsv1beta1.ComputeSecurityPolicyRef{External: in.GetEdgeSecurityPolicy()}
	}
	out.FailoverPolicy = BackendServiceFailoverPolicy_v1beta1_FromProto(mapCtx, in.GetFailoverPolicy())
	out.HealthChecks = ComputeBackendServiceSpec_HealthChecks_FromProto(mapCtx, in.HealthChecks)
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.LocalityLbPolicies = direct.Slice_FromProto(mapCtx, in.LocalityLbPolicies, BackendServiceLocalityLbPolicies_v1beta1_FromProto)
	out.LocalityLbPolicy = in.LocalityLbPolicy
	out.LogConfig = BackendServiceLogConfig_v1beta1_FromProto(mapCtx, in.GetLogConfig())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.OutlierDetection = BackendServiceOutlierDetection_v1beta1_FromProto(mapCtx, in.GetOutlierDetection())
	out.PortName = in.PortName
	out.Protocol = in.Protocol
	out.SecurityPolicy = in.SecurityPolicy
	out.SecuritySettings = BackendServiceSecuritySettings_v1beta1_FromProto(mapCtx, in.GetSecuritySettings())
	out.SessionAffinity = in.SessionAffinity
	out.Subsetting = BackendServiceSubsetting_v1beta1_FromProto(mapCtx, in.GetSubsetting())
	if in.TimeoutSec != nil {
		v := int(*in.TimeoutSec)
		out.TimeoutSec = &v
	}
	return out
}

func ComputeBackendServiceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceSpec) *pb.BackendService {
	if in == nil {
		return nil
	}
	out := &pb.BackendService{}

	if in.AffinityCookieTtlSec != nil {
		v := int32(*in.AffinityCookieTtlSec)
		out.AffinityCookieTtlSec = &v
	}
	out.Backends = direct.Slice_ToProto(mapCtx, in.Backend, BackendServiceBackend_v1beta1_ToProto)
	out.CdnPolicy = BackendServiceCdnPolicy_v1beta1_ToProto(mapCtx, in.CdnPolicy)
	out.EnableCDN = in.EnableCdn
	out.Iap = BackendServiceIap_v1beta1_ToProto(mapCtx, in.Iap)
	if in.ResourceID != nil {
		out.Name = in.ResourceID
	}
	out.CircuitBreakers = BackendServiceCircuitBreakers_v1beta1_ToProto(mapCtx, in.CircuitBreakers)
	out.CompressionMode = in.CompressionMode
	if in.ConnectionDrainingTimeoutSec != nil {
		v := int32(*in.ConnectionDrainingTimeoutSec)
		out.ConnectionDraining = &pb.ConnectionDraining{
			DrainingTimeoutSec: &v,
		}
	}
	out.ConnectionTrackingPolicy = BackendServiceConnectionTrackingPolicy_v1beta1_ToProto(mapCtx, in.ConnectionTrackingPolicy)
	out.ConsistentHash = BackendServiceConsistentHash_v1beta1_ToProto(mapCtx, in.ConsistentHash)
	out.CustomRequestHeaders = in.CustomRequestHeaders
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description
	if in.EdgeSecurityPolicyRef != nil {
		out.EdgeSecurityPolicy = &in.EdgeSecurityPolicyRef.External
	}
	out.FailoverPolicy = BackendServiceFailoverPolicy_v1beta1_ToProto(mapCtx, in.FailoverPolicy)
	out.HealthChecks = ComputeBackendServiceSpec_HealthChecks_ToProto(mapCtx, in.HealthChecks)
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.LocalityLbPolicies = direct.Slice_ToProto(mapCtx, in.LocalityLbPolicies, BackendServiceLocalityLbPolicies_v1beta1_ToProto)
	out.LocalityLbPolicy = in.LocalityLbPolicy
	out.LogConfig = BackendServiceLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	out.OutlierDetection = BackendServiceOutlierDetection_v1beta1_ToProto(mapCtx, in.OutlierDetection)
	out.PortName = in.PortName
	out.Protocol = in.Protocol
	out.SecurityPolicy = in.SecurityPolicy
	out.SecuritySettings = BackendServiceSecuritySettings_v1beta1_ToProto(mapCtx, in.SecuritySettings)
	out.SessionAffinity = in.SessionAffinity
	out.Subsetting = BackendServiceSubsetting_v1beta1_ToProto(mapCtx, in.Subsetting)
	if in.TimeoutSec != nil {
		v := int32(*in.TimeoutSec)
		out.TimeoutSec = &v
	}
	return out
}

func BackendServiceCdnPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicy) *krm.BackendServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_FromProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendServiceBypassCacheOnRequestHeaders_v1beta1_FromProto)
	out.CacheKeyPolicy = BackendServiceCacheKeyPolicy_v1beta1_FromProto(mapCtx, in.GetCacheKeyPolicy())
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		v := int(*in.ClientTtl)
		out.ClientTtl = &v
	}
	if in.DefaultTtl != nil {
		v := int(*in.DefaultTtl)
		out.DefaultTtl = &v
	}
	if in.MaxTtl != nil {
		v := int(*in.MaxTtl)
		out.MaxTtl = &v
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		v := int(*in.SignedUrlCacheMaxAgeSec)
		out.SignedUrlCacheMaxAgeSec = &v
	}
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_FromProto(mapCtx, in.NegativeCachingPolicy, BackendServiceNegativeCachingPolicy_v1beta1_FromProto)
	if in.ServeWhileStale != nil {
		v := int(*in.ServeWhileStale)
		out.ServeWhileStale = &v
	}
	return out
}
func BackendServiceCdnPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCdnPolicy) *pb.BackendServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_ToProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendServiceBypassCacheOnRequestHeaders_v1beta1_ToProto)
	out.CacheKeyPolicy = BackendServiceCacheKeyPolicy_v1beta1_ToProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = in.CacheMode
	if in.ClientTtl != nil {
		v := int32(*in.ClientTtl)
		out.ClientTtl = &v
	}
	if in.DefaultTtl != nil {
		v := int32(*in.DefaultTtl)
		out.DefaultTtl = &v
	}
	if in.MaxTtl != nil {
		v := int32(*in.MaxTtl)
		out.MaxTtl = &v
	}
	if in.SignedUrlCacheMaxAgeSec != nil {
		v := int64(*in.SignedUrlCacheMaxAgeSec)
		out.SignedUrlCacheMaxAgeSec = &v
	}
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_ToProto(mapCtx, in.NegativeCachingPolicy, BackendServiceNegativeCachingPolicy_v1beta1_ToProto)
	if in.ServeWhileStale != nil {
		v := int32(*in.ServeWhileStale)
		out.ServeWhileStale = &v
	}
	return out
}
func BackendServiceConsistentHash_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConsistentHashLoadBalancerSettings) *krm.BackendServiceConsistentHash {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceConsistentHash{}
	if in.MinimumRingSize != nil {
		v := int(*in.MinimumRingSize)
		out.MinimumRingSize = &v
	}
	out.HttpCookie = BackendServiceHttpCookie_v1beta1_FromProto(mapCtx, in.GetHttpCookie())
	out.HttpHeaderName = in.HttpHeaderName
	return out
}
func BackendServiceConsistentHash_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceConsistentHash) *pb.ConsistentHashLoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := &pb.ConsistentHashLoadBalancerSettings{}
	if in.MinimumRingSize != nil {
		v := int64(*in.MinimumRingSize)
		out.MinimumRingSize = &v
	}
	out.HttpCookie = BackendServiceHttpCookie_v1beta1_ToProto(mapCtx, in.HttpCookie)
	out.HttpHeaderName = in.HttpHeaderName
	return out
}

func BackendServiceHttpCookie_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConsistentHashLoadBalancerSettingsHttpCookie) *krm.BackendServiceHttpCookie {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceHttpCookie{}
	out.Name = in.Name
	out.Path = in.Path
	if in.GetTtl() != nil {
		out.Ttl = &krm.BackendServiceTtl{}
		if in.GetTtl().Nanos != nil {
			vNanos := int64(*in.GetTtl().Nanos)
			out.Ttl.Nanos = &vNanos
		}
		if in.GetTtl().Seconds != nil {
			out.Ttl.Seconds = *in.GetTtl().Seconds
		}
	}
	return out
}

func BackendServiceHttpCookie_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceHttpCookie) *pb.ConsistentHashLoadBalancerSettingsHttpCookie {
	if in == nil {
		return nil
	}
	out := &pb.ConsistentHashLoadBalancerSettingsHttpCookie{}
	out.Name = in.Name
	out.Path = in.Path
	if in.Ttl != nil {
		out.Ttl = &pb.Duration{}
		if in.Ttl.Nanos != nil {
			vNanos := int32(*in.Ttl.Nanos)
			out.Ttl.Nanos = &vNanos
		}
		vSecs := in.Ttl.Seconds
		out.Ttl.Seconds = &vSecs
	}
	return out
}

func BackendServiceIap_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceIAP) *krm.BackendServiceIap {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceIap{}
	out.Oauth2ClientId = in.Oauth2ClientId
	// Oauth2ClientIdRef is not retrieved from proto, as it's a Kubernetes-only ref field
	out.Oauth2ClientSecretSha256 = in.Oauth2ClientSecretSha256
	return out
}

func BackendServiceIap_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceIap) *pb.BackendServiceIAP {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceIAP{}
	if in.Oauth2ClientIdRef != nil {
		out.Oauth2ClientId = &in.Oauth2ClientIdRef.External
	} else {
		out.Oauth2ClientId = in.Oauth2ClientId
	}
	if in.Oauth2ClientSecret != nil {
		out.Oauth2ClientSecret = in.Oauth2ClientSecret.Value
	}
	out.Oauth2ClientSecretSha256 = in.Oauth2ClientSecretSha256
	return out
}

func ComputeBackendServiceStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendService) *krm.ComputeBackendServiceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendServiceStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		v := int(*in.Id)
		out.GeneratedId = &v
	}
	out.SelfLink = in.SelfLink
	return out
}

func ComputeBackendServiceStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceStatus) *pb.BackendService {
	if in == nil {
		return nil
	}
	out := &pb.BackendService{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.GeneratedId != nil {
		v := uint64(*in.GeneratedId)
		out.Id = &v
	}
	out.SelfLink = in.SelfLink
	return out
}
