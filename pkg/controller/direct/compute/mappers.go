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
	"strconv"
	"strings"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto converts a v1.SecurityPolicy proto to a v1alpha1.ComputeOrganizationSecurityPolicySpec.
// We hand-code this function because KRM displayName and parent map to proto ShortName and Parent.
func ComputeOrganizationSecurityPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec{}
	out.Description = in.Description
	out.DisplayName = in.GetShortName()
	out.Parent = in.GetParent()
	out.Type = in.Type
	return out
}

// ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto converts a v1alpha1.ComputeOrganizationSecurityPolicySpec to a v1.SecurityPolicy proto.
// We hand-code this function because KRM displayName and parent map to proto ShortName and Parent.
func ComputeOrganizationSecurityPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeOrganizationSecurityPolicySpec) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.Description = in.Description
	if in.DisplayName != "" {
		out.ShortName = &in.DisplayName
	}
	if in.Parent != "" {
		out.Parent = &in.Parent
	}
	out.Type = in.Type
	return out
}

// ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto converts a v1.SecurityPolicy proto to a v1alpha1.ComputeOrganizationSecurityPolicyStatus.
// We hand-code this function because policyId maps to Id (which requires string to uint64 conversion).
func ComputeOrganizationSecurityPolicyStatus_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.SecurityPolicy) *krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus{}
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		idStr := strconv.FormatUint(*in.Id, 10)
		out.PolicyId = &idStr
	}
	return out
}

// ComputeOrganizationSecurityPolicyStatus_v1alpha1_ToProto converts a v1alpha1.ComputeOrganizationSecurityPolicyStatus to a v1.SecurityPolicy proto.
// We hand-code this function because policyId maps to Id (which requires string to uint64 conversion).
func ComputeOrganizationSecurityPolicyStatus_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeOrganizationSecurityPolicyStatus) *pb.SecurityPolicy {
	if in == nil {
		return nil
	}
	out := &pb.SecurityPolicy{}
	out.Fingerprint = in.Fingerprint
	if in.PolicyId != nil {
		idVal, err := strconv.ParseUint(*in.PolicyId, 10, 64)
		if err != nil {
			mapCtx.Errorf("parsing policyId %q: %v", *in.PolicyId, err)
		} else {
			out.Id = &idVal
		}
	}
	return out
}

// ComputeExternalVPNGatewayInterface_v1beta1_FromProto maps a pb.ExternalVpnGatewayInterface to a krm.ComputeExternalVPNGatewayInterface.
// It is handcoded here because of type mismatches: KRM ID is *int64, while Proto ID is *uint32.
func ComputeExternalVPNGatewayInterface_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ExternalVpnGatewayInterface) *krm.ComputeExternalVPNGatewayInterface {
	if in == nil {
		return nil
	}
	out := &krm.ComputeExternalVPNGatewayInterface{}
	if in.Id != nil {
		idVal := int64(*in.Id)
		out.ID = &idVal
	}
	out.IPAddress = in.IpAddress
	return out
}

// ComputeExternalVPNGatewayInterface_v1beta1_ToProto maps a krm.ComputeExternalVPNGatewayInterface to a pb.ExternalVpnGatewayInterface.
// It is handcoded here because of type mismatches: KRM ID is *int64, while Proto ID is *uint32.
func ComputeExternalVPNGatewayInterface_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeExternalVPNGatewayInterface) *pb.ExternalVpnGatewayInterface {
	if in == nil {
		return nil
	}
	out := &pb.ExternalVpnGatewayInterface{}
	if in.ID != nil {
		idVal := uint32(*in.ID)
		out.Id = &idVal
	}
	out.IpAddress = in.IPAddress
	return out
}

// ComputeExternalVPNGatewayStatus_v1beta1_FromProto maps a pb.ExternalVpnGateway to a krm.ComputeExternalVPNGatewayStatus.
// It is handcoded here to organize status mappings cleanly within mappers.go.
func ComputeExternalVPNGatewayStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ExternalVpnGateway) *krm.ComputeExternalVPNGatewayStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeExternalVPNGatewayStatus{}
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	return out
}

// ComputeExternalVPNGatewayStatus_v1beta1_ToProto maps a krm.ComputeExternalVPNGatewayStatus to a pb.ExternalVpnGateway.
// It is handcoded here to organize status mappings cleanly within mappers.go.
func ComputeExternalVPNGatewayStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeExternalVPNGatewayStatus) *pb.ExternalVpnGateway {
	if in == nil {
		return nil
	}
	out := &pb.ExternalVpnGateway{}
	out.LabelFingerprint = in.LabelFingerprint
	out.SelfLink = in.SelfLink
	return out
}

// ComputeBackendServiceSpec_v1beta1_ToProto maps ComputeBackendServiceSpec to the protobuf BackendService.
// Handcoded to handle custom reference structures, flat mapping for connection draining timeout, and regional vs global locations.
func ComputeBackendServiceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceSpec) *pb.BackendService {
	if in == nil {
		return nil
	}
	out := &pb.BackendService{}

	out.AffinityCookieTtlSec = in.AffinityCookieTtlSec
	out.Backends = ComputeBackendServiceSpec_Backends_v1beta1_ToProto(mapCtx, in.Backend)
	out.CdnPolicy = BackendServiceCdnPolicy_v1beta1_ToProto(mapCtx, in.CDNPolicy)
	out.CircuitBreakers = CircuitBreakers_v1beta1_ToProto(mapCtx, in.CircuitBreakers)
	out.CompressionMode = in.CompressionMode

	if in.ConnectionDrainingTimeoutSec != nil {
		out.ConnectionDraining = &pb.ConnectionDraining{
			DrainingTimeoutSec: in.ConnectionDrainingTimeoutSec,
		}
	}

	out.ConnectionTrackingPolicy = BackendServiceConnectionTrackingPolicy_v1beta1_ToProto(mapCtx, in.ConnectionTrackingPolicy)
	out.ConsistentHash = ConsistentHashLoadBalancerSettings_v1beta1_ToProto(mapCtx, in.ConsistentHash)
	out.CustomRequestHeaders = in.CustomRequestHeaders
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description

	if in.EdgeSecurityPolicyRef != nil {
		out.EdgeSecurityPolicy = &in.EdgeSecurityPolicyRef.External
	}

	out.EnableCDN = in.EnableCdn
	out.FailoverPolicy = BackendServiceFailoverPolicy_v1beta1_ToProto(mapCtx, in.FailoverPolicy)
	out.HealthChecks = ComputeBackendServiceSpec_HealthChecks_v1beta1_ToProto(mapCtx, in.HealthChecks)
	out.Iap = BackendServiceIap_v1beta1_ToProto(mapCtx, in.Iap)
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.LocalityLbPolicies = direct.Slice_ToProto(mapCtx, in.LocalityLbPolicies, BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_ToProto)
	out.LocalityLbPolicy = in.LocalityLbPolicy

	if in.Location != "" && in.Location != "global" {
		out.Region = &in.Location
	}

	out.LogConfig = BackendServiceLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)

	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}

	out.OutlierDetection = OutlierDetection_v1beta1_ToProto(mapCtx, in.OutlierDetection)
	out.PortName = in.PortName
	out.Protocol = in.Protocol
	out.Name = in.ResourceID

	if in.SecurityPolicyRef != nil {
		out.SecurityPolicy = &in.SecurityPolicyRef.External
	} else if in.SecurityPolicy != nil {
		out.SecurityPolicy = in.SecurityPolicy
	}

	out.SecuritySettings = SecuritySettings_v1beta1_ToProto(mapCtx, in.SecuritySettings)
	out.SessionAffinity = in.SessionAffinity
	out.Subsetting = Subsetting_v1beta1_ToProto(mapCtx, in.Subsetting)
	out.TimeoutSec = in.TimeoutSec

	return out
}

// ComputeBackendServiceSpec_v1beta1_FromProto maps protobuf BackendService to ComputeBackendServiceSpec.
// Handcoded to handle custom reference structures, flat mapping for connection draining timeout, and regional vs global locations.
func ComputeBackendServiceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendService) *krm.ComputeBackendServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendServiceSpec{}

	out.AffinityCookieTtlSec = in.AffinityCookieTtlSec
	out.Backend = ComputeBackendServiceSpec_Backends_v1beta1_FromProto(mapCtx, in.Backends)
	out.CDNPolicy = BackendServiceCdnPolicy_v1beta1_FromProto(mapCtx, in.CdnPolicy)
	out.CircuitBreakers = CircuitBreakers_v1beta1_FromProto(mapCtx, in.CircuitBreakers)
	out.CompressionMode = in.CompressionMode

	if in.ConnectionDraining != nil {
		out.ConnectionDrainingTimeoutSec = in.ConnectionDraining.DrainingTimeoutSec
	}

	out.ConnectionTrackingPolicy = BackendServiceConnectionTrackingPolicy_v1beta1_FromProto(mapCtx, in.ConnectionTrackingPolicy)
	out.ConsistentHash = ConsistentHashLoadBalancerSettings_v1beta1_FromProto(mapCtx, in.ConsistentHash)
	out.CustomRequestHeaders = in.CustomRequestHeaders
	out.CustomResponseHeaders = in.CustomResponseHeaders
	out.Description = in.Description

	if in.EdgeSecurityPolicy != nil && *in.EdgeSecurityPolicy != "" {
		out.EdgeSecurityPolicyRef = &krm.ComputeSecurityPolicyRef{External: *in.EdgeSecurityPolicy}
	}

	out.EnableCdn = in.EnableCDN
	out.FailoverPolicy = BackendServiceFailoverPolicy_v1beta1_FromProto(mapCtx, in.FailoverPolicy)
	out.HealthChecks = ComputeBackendServiceSpec_HealthChecks_v1beta1_FromProto(mapCtx, in.HealthChecks)
	out.Iap = BackendServiceIap_v1beta1_FromProto(mapCtx, in.Iap)
	out.LoadBalancingScheme = in.LoadBalancingScheme
	out.LocalityLbPolicies = direct.Slice_FromProto(mapCtx, in.LocalityLbPolicies, BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_FromProto)
	out.LocalityLbPolicy = in.LocalityLbPolicy

	if in.Region != nil && *in.Region != "" {
		out.Location = lastComponent(*in.Region)
	} else {
		out.Location = "global"
	}

	out.LogConfig = BackendServiceLogConfig_v1beta1_FromProto(mapCtx, in.LogConfig)

	if in.Network != nil && *in.Network != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: *in.Network}
	}

	out.OutlierDetection = OutlierDetection_v1beta1_FromProto(mapCtx, in.OutlierDetection)
	out.PortName = in.PortName
	out.Protocol = in.Protocol
	out.ResourceID = in.Name

	if in.SecurityPolicy != nil && *in.SecurityPolicy != "" {
		out.SecurityPolicy = in.SecurityPolicy
		out.SecurityPolicyRef = &krm.ComputeSecurityPolicyRef{External: *in.SecurityPolicy}
	}

	out.SecuritySettings = SecuritySettings_v1beta1_FromProto(mapCtx, in.SecuritySettings)
	out.SessionAffinity = in.SessionAffinity
	out.Subsetting = Subsetting_v1beta1_FromProto(mapCtx, in.Subsetting)
	out.TimeoutSec = in.TimeoutSec

	return out
}

// ComputeBackendServiceStatus_v1beta1_ToProto maps ComputeBackendServiceStatus to protobuf BackendService.
// Handcoded to preserve custom fields like GeneratedId, CreationTimestamp, Fingerprint, and SelfLink.
func ComputeBackendServiceStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceStatus) *pb.BackendService {
	if in == nil {
		return nil
	}
	out := &pb.BackendService{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.GeneratedId != nil {
		val := uint64(*in.GeneratedId)
		out.Id = &val
	}
	out.SelfLink = in.SelfLink
	return out
}

// ComputeBackendServiceStatus_v1beta1_FromProto maps protobuf BackendService to ComputeBackendServiceStatus.
// Handcoded to preserve custom fields like GeneratedId, CreationTimestamp, Fingerprint, and SelfLink.
func ComputeBackendServiceStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendService) *krm.ComputeBackendServiceStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeBackendServiceStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Fingerprint = in.Fingerprint
	if in.Id != nil {
		val := int64(*in.Id)
		out.GeneratedId = &val
	}
	out.SelfLink = in.SelfLink
	return out
}

// ComputeBackendServiceSpec_Backends_v1beta1_ToProto maps slice of Backend KRM objects to pb.Backend slice.
func ComputeBackendServiceSpec_Backends_v1beta1_ToProto(mapCtx *direct.MapContext, in []krm.Backend) []*pb.Backend {
	if in == nil {
		return nil
	}
	out := make([]*pb.Backend, len(in))
	for i, b := range in {
		out[i] = Backend_v1beta1_ToProto(mapCtx, &b)
	}
	return out
}

// ComputeBackendServiceSpec_Backends_v1beta1_FromProto maps slice of pb.Backend to Backend KRM objects slice.
func ComputeBackendServiceSpec_Backends_v1beta1_FromProto(mapCtx *direct.MapContext, in []*pb.Backend) []krm.Backend {
	if in == nil {
		return nil
	}
	out := make([]krm.Backend, len(in))
	for i, b := range in {
		out[i] = *Backend_v1beta1_FromProto(mapCtx, b)
	}
	return out
}

// Backend_v1beta1_ToProto maps a Backend KRM struct to pb.Backend.
// Handcoded to map custom nested group references.
func Backend_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Backend) *pb.Backend {
	if in == nil {
		return nil
	}
	out := &pb.Backend{}
	out.BalancingMode = in.BalancingMode
	out.CapacityScaler = in.CapacityScaler
	out.Description = in.Description
	out.Failover = in.Failover
	out.Group = BackendGroup_v1beta1_ToProto(mapCtx, in.Group)
	out.MaxConnections = in.MaxConnections
	out.MaxConnectionsPerEndpoint = in.MaxConnectionsPerEndpoint
	out.MaxConnectionsPerInstance = in.MaxConnectionsPerInstance
	out.MaxRate = in.MaxRate
	out.MaxRatePerEndpoint = in.MaxRatePerEndpoint
	out.MaxRatePerInstance = in.MaxRatePerInstance
	out.MaxUtilization = in.MaxUtilization
	return out
}

// Backend_v1beta1_FromProto maps pb.Backend to a Backend KRM struct.
// Handcoded to map custom nested group references.
func Backend_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Backend) *krm.Backend {
	if in == nil {
		return nil
	}
	out := &krm.Backend{}
	out.BalancingMode = in.BalancingMode
	out.CapacityScaler = in.CapacityScaler
	out.Description = in.Description
	out.Failover = in.Failover
	out.Group = BackendGroup_v1beta1_FromProto(mapCtx, in.Group)
	out.MaxConnections = in.MaxConnections
	out.MaxConnectionsPerEndpoint = in.MaxConnectionsPerEndpoint
	out.MaxConnectionsPerInstance = in.MaxConnectionsPerInstance
	out.MaxRate = in.MaxRate
	out.MaxRatePerEndpoint = in.MaxRatePerEndpoint
	out.MaxRatePerInstance = in.MaxRatePerInstance
	out.MaxUtilization = in.MaxUtilization
	return out
}

// BackendGroup_v1beta1_ToProto converts KRM BackendGroup to proto group selfLink string.
func BackendGroup_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendGroup) *string {
	if in == nil {
		return nil
	}
	if in.InstanceGroupRef != nil && in.InstanceGroupRef.External != "" {
		return &in.InstanceGroupRef.External
	}
	if in.NetworkEndpointGroupRef != nil && in.NetworkEndpointGroupRef.External != "" {
		return &in.NetworkEndpointGroupRef.External
	}
	return nil
}

// BackendGroup_v1beta1_FromProto converts proto group selfLink string to KRM BackendGroup.
func BackendGroup_v1beta1_FromProto(mapCtx *direct.MapContext, in *string) *krm.BackendGroup {
	if in == nil || *in == "" {
		return nil
	}
	val := *in
	if strings.Contains(val, "/instanceGroups/") {
		return &krm.BackendGroup{
			InstanceGroupRef: &krm.ComputeInstanceGroupRef{External: val},
		}
	}
	if strings.Contains(val, "/networkEndpointGroups/") {
		return &krm.BackendGroup{
			NetworkEndpointGroupRef: &krm.ComputeNetworkEndpointGroupRef{External: val},
		}
	}
	return &krm.BackendGroup{
		InstanceGroupRef: &krm.ComputeInstanceGroupRef{External: val},
	}
}

// ComputeBackendServiceSpec_HealthChecks_v1beta1_ToProto converts KRM HealthChecks slice to proto health check selfLink slice.
func ComputeBackendServiceSpec_HealthChecks_v1beta1_ToProto(mapCtx *direct.MapContext, in []krm.BackendserviceHealthChecks) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, h := range in {
		if h.HealthCheckRef != nil {
			out[i] = h.HealthCheckRef.External
		} else if h.HttpHealthCheckRef != nil {
			out[i] = h.HttpHealthCheckRef.External
		}
	}
	return out
}

// ComputeBackendServiceSpec_HealthChecks_v1beta1_FromProto converts proto health check selfLink slice to KRM HealthChecks slice.
func ComputeBackendServiceSpec_HealthChecks_v1beta1_FromProto(mapCtx *direct.MapContext, in []string) []krm.BackendserviceHealthChecks {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	outKRM := make([]krm.BackendserviceHealthChecks, len(in))
	for i, url := range out {
		if strings.Contains(url, "/httpHealthChecks/") {
			outKRM[i] = krm.BackendserviceHealthChecks{
				HttpHealthCheckRef: &krm.ComputeHTTPHealthCheckRef{External: url},
			}
		} else {
			outKRM[i] = krm.BackendserviceHealthChecks{
				HealthCheckRef: &krm.ComputeHealthCheckRef{External: url},
			}
		}
	}
	return outKRM
}

// BackendServiceDuration_v1beta1_ToProto maps BackendServiceDuration to pb.Duration.
func BackendServiceDuration_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceDuration) *pb.Duration {
	if in == nil {
		return nil
	}
	out := &pb.Duration{}
	out.Seconds = &in.Seconds
	out.Nanos = in.Nanos
	return out
}

// BackendServiceDuration_v1beta1_FromProto maps pb.Duration to BackendServiceDuration.
func BackendServiceDuration_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Duration) *krm.BackendServiceDuration {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceDuration{}
	if in.Seconds != nil {
		out.Seconds = *in.Seconds
	}
	out.Nanos = in.Nanos
	return out
}

// CircuitBreakers_v1beta1_ToProto maps CircuitBreakers to pb.CircuitBreakers.
func CircuitBreakers_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CircuitBreakers) *pb.CircuitBreakers {
	if in == nil {
		return nil
	}
	out := &pb.CircuitBreakers{}
	out.MaxConnections = in.MaxConnections
	out.MaxPendingRequests = in.MaxPendingRequests
	out.MaxRequests = in.MaxRequests
	out.MaxRequestsPerConnection = in.MaxRequestsPerConnection
	out.MaxRetries = in.MaxRetries
	return out
}

// CircuitBreakers_v1beta1_FromProto maps pb.CircuitBreakers to CircuitBreakers.
func CircuitBreakers_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CircuitBreakers) *krm.CircuitBreakers {
	if in == nil {
		return nil
	}
	out := &krm.CircuitBreakers{}
	out.MaxConnections = in.MaxConnections
	out.MaxPendingRequests = in.MaxPendingRequests
	out.MaxRequests = in.MaxRequests
	out.MaxRequestsPerConnection = in.MaxRequestsPerConnection
	out.MaxRetries = in.MaxRetries
	return out
}

// BackendServiceHttpCookie_v1beta1_ToProto maps BackendServiceHttpCookie to pb.ConsistentHashLoadBalancerSettingsHttpCookie.
func BackendServiceHttpCookie_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceHttpCookie) *pb.ConsistentHashLoadBalancerSettingsHttpCookie {
	if in == nil {
		return nil
	}
	out := &pb.ConsistentHashLoadBalancerSettingsHttpCookie{}
	out.Name = in.Name
	out.Path = in.Path
	out.Ttl = BackendServiceDuration_v1beta1_ToProto(mapCtx, in.Ttl)
	return out
}

// BackendServiceHttpCookie_v1beta1_FromProto maps pb.ConsistentHashLoadBalancerSettingsHttpCookie to BackendServiceHttpCookie.
func BackendServiceHttpCookie_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConsistentHashLoadBalancerSettingsHttpCookie) *krm.BackendServiceHttpCookie {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceHttpCookie{}
	out.Name = in.Name
	out.Path = in.Path
	out.Ttl = BackendServiceDuration_v1beta1_FromProto(mapCtx, in.Ttl)
	return out
}

// ConsistentHashLoadBalancerSettings_v1beta1_ToProto maps ConsistentHashLoadBalancerSettings to pb.ConsistentHashLoadBalancerSettings.
func ConsistentHashLoadBalancerSettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ConsistentHashLoadBalancerSettings) *pb.ConsistentHashLoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := &pb.ConsistentHashLoadBalancerSettings{}
	out.HttpCookie = BackendServiceHttpCookie_v1beta1_ToProto(mapCtx, in.HttpCookie)
	out.HttpHeaderName = in.HttpHeaderName
	out.MinimumRingSize = in.MinimumRingSize
	return out
}

// ConsistentHashLoadBalancerSettings_v1beta1_FromProto maps pb.ConsistentHashLoadBalancerSettings to ConsistentHashLoadBalancerSettings.
func ConsistentHashLoadBalancerSettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ConsistentHashLoadBalancerSettings) *krm.ConsistentHashLoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := &krm.ConsistentHashLoadBalancerSettings{}
	out.HttpCookie = BackendServiceHttpCookie_v1beta1_FromProto(mapCtx, in.HttpCookie)
	out.HttpHeaderName = in.HttpHeaderName
	out.MinimumRingSize = in.MinimumRingSize
	return out
}

// OutlierDetection_v1beta1_ToProto maps OutlierDetection to pb.OutlierDetection.
func OutlierDetection_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.OutlierDetection) *pb.OutlierDetection {
	if in == nil {
		return nil
	}
	out := &pb.OutlierDetection{}
	out.BaseEjectionTime = BackendServiceDuration_v1beta1_ToProto(mapCtx, in.BaseEjectionTime)
	out.ConsecutiveErrors = in.ConsecutiveErrors
	out.ConsecutiveGatewayFailure = in.ConsecutiveGatewayFailure
	out.EnforcingConsecutiveErrors = in.EnforcingConsecutiveErrors
	out.EnforcingConsecutiveGatewayFailure = in.EnforcingConsecutiveGatewayFailure
	out.EnforcingSuccessRate = in.EnforcingSuccessRate
	out.Interval = BackendServiceDuration_v1beta1_ToProto(mapCtx, in.Interval)
	out.MaxEjectionPercent = in.MaxEjectionPercent
	out.SuccessRateMinimumHosts = in.SuccessRateMinimumHosts
	out.SuccessRateRequestVolume = in.SuccessRateRequestVolume
	out.SuccessRateStdevFactor = in.SuccessRateStdevFactor
	return out
}

// OutlierDetection_v1beta1_FromProto maps pb.OutlierDetection to OutlierDetection.
func OutlierDetection_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.OutlierDetection) *krm.OutlierDetection {
	if in == nil {
		return nil
	}
	out := &krm.OutlierDetection{}
	out.BaseEjectionTime = BackendServiceDuration_v1beta1_FromProto(mapCtx, in.BaseEjectionTime)
	out.ConsecutiveErrors = in.ConsecutiveErrors
	out.ConsecutiveGatewayFailure = in.ConsecutiveGatewayFailure
	out.EnforcingConsecutiveErrors = in.EnforcingConsecutiveErrors
	out.EnforcingConsecutiveGatewayFailure = in.EnforcingConsecutiveGatewayFailure
	out.EnforcingSuccessRate = in.EnforcingSuccessRate
	out.Interval = BackendServiceDuration_v1beta1_FromProto(mapCtx, in.Interval)
	out.MaxEjectionPercent = in.MaxEjectionPercent
	out.SuccessRateMinimumHosts = in.SuccessRateMinimumHosts
	out.SuccessRateRequestVolume = in.SuccessRateRequestVolume
	out.SuccessRateStdevFactor = in.SuccessRateStdevFactor
	return out
}

// CacheKeyPolicy_v1beta1_ToProto maps CacheKeyPolicy to pb.CacheKeyPolicy.
func CacheKeyPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.CacheKeyPolicy) *pb.CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CacheKeyPolicy{}
	out.IncludeHost = in.IncludeHost
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.IncludeNamedCookies = in.IncludeNamedCookies
	out.IncludeProtocol = in.IncludeProtocol
	out.IncludeQueryString = in.IncludeQueryString
	out.QueryStringBlacklist = in.QueryStringBlacklist
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

// CacheKeyPolicy_v1beta1_FromProto maps pb.CacheKeyPolicy to CacheKeyPolicy.
func CacheKeyPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.CacheKeyPolicy) *krm.CacheKeyPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CacheKeyPolicy{}
	out.IncludeHost = in.IncludeHost
	out.IncludeHttpHeaders = in.IncludeHttpHeaders
	out.IncludeNamedCookies = in.IncludeNamedCookies
	out.IncludeProtocol = in.IncludeProtocol
	out.IncludeQueryString = in.IncludeQueryString
	out.QueryStringBlacklist = in.QueryStringBlacklist
	out.QueryStringWhitelist = in.QueryStringWhitelist
	return out
}

// BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_ToProto maps BackendServiceCdnPolicyBypassCacheOnRequestHeader to pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader.
func BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCdnPolicyBypassCacheOnRequestHeader) *pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader{}
	out.HeaderName = &in.HeaderName
	return out
}

// BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_FromProto maps pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader to BackendServiceCdnPolicyBypassCacheOnRequestHeader.
func BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicyBypassCacheOnRequestHeader) *krm.BackendServiceCdnPolicyBypassCacheOnRequestHeader {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCdnPolicyBypassCacheOnRequestHeader{}
	if in.HeaderName != nil {
		out.HeaderName = *in.HeaderName
	}
	return out
}

// BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_ToProto maps BackendServiceCdnPolicyNegativeCachingPolicy to pb.BackendServiceCdnPolicyNegativeCachingPolicy.
func BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCdnPolicyNegativeCachingPolicy) *pb.BackendServiceCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicyNegativeCachingPolicy{}
	out.Code = in.Code
	out.Ttl = in.Ttl
	return out
}

// BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_FromProto maps pb.BackendServiceCdnPolicyNegativeCachingPolicy to BackendServiceCdnPolicyNegativeCachingPolicy.
func BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicyNegativeCachingPolicy) *krm.BackendServiceCdnPolicyNegativeCachingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCdnPolicyNegativeCachingPolicy{}
	out.Code = in.Code
	out.Ttl = in.Ttl
	return out
}

// BackendServiceCdnPolicy_v1beta1_ToProto maps BackendServiceCdnPolicy to pb.BackendServiceCdnPolicy.
func BackendServiceCdnPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceCdnPolicy) *pb.BackendServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_ToProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_ToProto)
	out.CacheKeyPolicy = CacheKeyPolicy_v1beta1_ToProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = in.CacheMode
	out.ClientTtl = in.ClientTtl
	out.DefaultTtl = in.DefaultTtl
	out.MaxTtl = in.MaxTtl
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_ToProto(mapCtx, in.NegativeCachingPolicy, BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_ToProto)
	out.ServeWhileStale = in.ServeWhileStale
	out.SignedUrlCacheMaxAgeSec = in.SignedUrlCacheMaxAgeSec
	return out
}

// BackendServiceCdnPolicy_v1beta1_FromProto maps pb.BackendServiceCdnPolicy to BackendServiceCdnPolicy.
func BackendServiceCdnPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceCdnPolicy) *krm.BackendServiceCdnPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceCdnPolicy{}
	out.BypassCacheOnRequestHeaders = direct.Slice_FromProto(mapCtx, in.BypassCacheOnRequestHeaders, BackendServiceCdnPolicyBypassCacheOnRequestHeader_v1beta1_FromProto)
	out.CacheKeyPolicy = CacheKeyPolicy_v1beta1_FromProto(mapCtx, in.CacheKeyPolicy)
	out.CacheMode = in.CacheMode
	out.ClientTtl = in.ClientTtl
	out.DefaultTtl = in.DefaultTtl
	out.MaxTtl = in.MaxTtl
	out.NegativeCaching = in.NegativeCaching
	out.NegativeCachingPolicy = direct.Slice_FromProto(mapCtx, in.NegativeCachingPolicy, BackendServiceCdnPolicyNegativeCachingPolicy_v1beta1_FromProto)
	out.ServeWhileStale = in.ServeWhileStale
	out.SignedUrlCacheMaxAgeSec = in.SignedUrlCacheMaxAgeSec
	return out
}

// BackendServiceIap_v1beta1_ToProto maps BackendServiceIap to pb.BackendServiceIAP.
func BackendServiceIap_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceIap) *pb.BackendServiceIAP {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceIAP{}
	enabled := true
	out.Enabled = &enabled

	if in.Oauth2ClientIdRef != nil && in.Oauth2ClientIdRef.External != "" {
		out.Oauth2ClientId = &in.Oauth2ClientIdRef.External
	} else if in.Oauth2ClientId != nil {
		out.Oauth2ClientId = in.Oauth2ClientId
	}

	if in.Oauth2ClientSecret != nil && in.Oauth2ClientSecret.Value != nil {
		out.Oauth2ClientSecret = in.Oauth2ClientSecret.Value
	}

	out.Oauth2ClientSecretSha256 = in.Oauth2ClientSecretSha256

	return out
}

// BackendServiceIap_v1beta1_FromProto maps pb.BackendServiceIAP to BackendServiceIap.
func BackendServiceIap_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceIAP) *krm.BackendServiceIap {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceIap{}
	if in.Oauth2ClientId != nil && *in.Oauth2ClientId != "" {
		out.Oauth2ClientId = in.Oauth2ClientId
		out.Oauth2ClientIdRef = &krm.BackendServiceOauth2ClientIdRef{External: *in.Oauth2ClientId}
	}
	if in.Oauth2ClientSecret != nil && *in.Oauth2ClientSecret != "" {
		out.Oauth2ClientSecret = &secret.Legacy{
			Value: in.Oauth2ClientSecret,
		}
	}
	out.Oauth2ClientSecretSha256 = in.Oauth2ClientSecretSha256
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_ToProto maps BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy to pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy.
func BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy) *pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy{}
	out.Data = in.Data
	out.Name = &in.Name
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_FromProto maps pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy to BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy.
func BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy) *krm.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy{}
	out.Data = in.Data
	if in.Name != nil {
		out.Name = *in.Name
	}
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_ToProto maps BackendServiceLocalityLoadBalancingPolicyConfigPolicy to pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy.
func BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLocalityLoadBalancingPolicyConfigPolicy) *pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy{}
	out.Name = &in.Name
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_FromProto maps pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy to BackendServiceLocalityLoadBalancingPolicyConfigPolicy.
func BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfigPolicy) *krm.BackendServiceLocalityLoadBalancingPolicyConfigPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLocalityLoadBalancingPolicyConfigPolicy{}
	if in.Name != nil {
		out.Name = *in.Name
	}
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_ToProto maps BackendServiceLocalityLoadBalancingPolicyConfig to pb.BackendServiceLocalityLoadBalancingPolicyConfig.
func BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLocalityLoadBalancingPolicyConfig) *pb.BackendServiceLocalityLoadBalancingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLocalityLoadBalancingPolicyConfig{}
	out.CustomPolicy = BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_ToProto(mapCtx, in.CustomPolicy)
	out.Policy = BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_ToProto(mapCtx, in.Policy)
	return out
}

// BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_FromProto maps BackendServiceLocalityLoadBalancingPolicyConfig to BackendServiceLocalityLoadBalancingPolicyConfig.
func BackendServiceLocalityLoadBalancingPolicyConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLocalityLoadBalancingPolicyConfig) *krm.BackendServiceLocalityLoadBalancingPolicyConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLocalityLoadBalancingPolicyConfig{}
	out.CustomPolicy = BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicy_v1beta1_FromProto(mapCtx, in.CustomPolicy)
	out.Policy = BackendServiceLocalityLoadBalancingPolicyConfigPolicy_v1beta1_FromProto(mapCtx, in.Policy)
	return out
}

// BackendServiceLogConfig_v1beta1_ToProto maps BackendServiceLogConfig to pb.BackendServiceLogConfig.
func BackendServiceLogConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BackendServiceLogConfig) *pb.BackendServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &pb.BackendServiceLogConfig{}
	out.Enable = in.Enable
	if in.SampleRate != nil {
		val := float32(*in.SampleRate)
		out.SampleRate = &val
	}
	return out
}

// BackendServiceLogConfig_v1beta1_FromProto maps pb.BackendServiceLogConfig to BackendServiceLogConfig.
func BackendServiceLogConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackendServiceLogConfig) *krm.BackendServiceLogConfig {
	if in == nil {
		return nil
	}
	out := &krm.BackendServiceLogConfig{}
	out.Enable = in.Enable
	if in.SampleRate != nil {
		val := float64(*in.SampleRate)
		out.SampleRate = &val
	}
	return out
}

// SecuritySettings_v1beta1_ToProto maps SecuritySettings to pb.SecuritySettings.
func SecuritySettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.SecuritySettings) *pb.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &pb.SecuritySettings{}
	if in.ClientTLSPolicyRef != nil {
		out.ClientTlsPolicy = &in.ClientTLSPolicyRef.External
	}
	if in.SubjectAltNames != nil {
		out.SubjectAltNames = in.SubjectAltNames
	}
	return out
}

// SecuritySettings_v1beta1_FromProto maps pb.SecuritySettings to SecuritySettings.
func SecuritySettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SecuritySettings) *krm.SecuritySettings {
	if in == nil {
		return nil
	}
	out := &krm.SecuritySettings{}
	if in.ClientTlsPolicy != nil && *in.ClientTlsPolicy != "" {
		out.ClientTLSPolicyRef = &krm.ComputeSecuritySettingsClientTLSPolicyRef{External: *in.ClientTlsPolicy}
	}
	if in.SubjectAltNames != nil {
		out.SubjectAltNames = in.SubjectAltNames
	}
	return out
}

// Subsetting_v1beta1_ToProto maps Subsetting to pb.Subsetting.
func Subsetting_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Subsetting) *pb.Subsetting {
	if in == nil {
		return nil
	}
	out := &pb.Subsetting{}
	out.Policy = &in.Policy
	return out
}

// Subsetting_v1beta1_FromProto maps pb.Subsetting to Subsetting.
func Subsetting_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Subsetting) *krm.Subsetting {
	if in == nil {
		return nil
	}
	out := &krm.Subsetting{}
	if in.Policy != nil {
		out.Policy = *in.Policy
	}
	return out
}
