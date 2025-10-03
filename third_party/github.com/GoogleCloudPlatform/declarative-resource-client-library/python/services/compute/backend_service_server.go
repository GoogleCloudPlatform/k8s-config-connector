// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for BackendService.
type BackendServiceServer struct{}

// ProtoToBackendServiceBackendsBalancingModeEnum converts a BackendServiceBackendsBalancingModeEnum enum from its proto representation.
func ProtoToComputeBackendServiceBackendsBalancingModeEnum(e computepb.ComputeBackendServiceBackendsBalancingModeEnum) *compute.BackendServiceBackendsBalancingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceBackendsBalancingModeEnum_name[int32(e)]; ok {
		e := compute.BackendServiceBackendsBalancingModeEnum(n[len("ComputeBackendServiceBackendsBalancingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceProtocolEnum converts a BackendServiceProtocolEnum enum from its proto representation.
func ProtoToComputeBackendServiceProtocolEnum(e computepb.ComputeBackendServiceProtocolEnum) *compute.BackendServiceProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceProtocolEnum_name[int32(e)]; ok {
		e := compute.BackendServiceProtocolEnum(n[len("ComputeBackendServiceProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceSessionAffinityEnum converts a BackendServiceSessionAffinityEnum enum from its proto representation.
func ProtoToComputeBackendServiceSessionAffinityEnum(e computepb.ComputeBackendServiceSessionAffinityEnum) *compute.BackendServiceSessionAffinityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceSessionAffinityEnum_name[int32(e)]; ok {
		e := compute.BackendServiceSessionAffinityEnum(n[len("ComputeBackendServiceSessionAffinityEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceLoadBalancingSchemeEnum converts a BackendServiceLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeBackendServiceLoadBalancingSchemeEnum(e computepb.ComputeBackendServiceLoadBalancingSchemeEnum) *compute.BackendServiceLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := compute.BackendServiceLoadBalancingSchemeEnum(n[len("ComputeBackendServiceLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceCdnPolicyCacheModeEnum converts a BackendServiceCdnPolicyCacheModeEnum enum from its proto representation.
func ProtoToComputeBackendServiceCdnPolicyCacheModeEnum(e computepb.ComputeBackendServiceCdnPolicyCacheModeEnum) *compute.BackendServiceCdnPolicyCacheModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceCdnPolicyCacheModeEnum_name[int32(e)]; ok {
		e := compute.BackendServiceCdnPolicyCacheModeEnum(n[len("ComputeBackendServiceCdnPolicyCacheModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceLocalityLbPolicyEnum converts a BackendServiceLocalityLbPolicyEnum enum from its proto representation.
func ProtoToComputeBackendServiceLocalityLbPolicyEnum(e computepb.ComputeBackendServiceLocalityLbPolicyEnum) *compute.BackendServiceLocalityLbPolicyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeBackendServiceLocalityLbPolicyEnum_name[int32(e)]; ok {
		e := compute.BackendServiceLocalityLbPolicyEnum(n[len("ComputeBackendServiceLocalityLbPolicyEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceBackends converts a BackendServiceBackends resource from its proto representation.
func ProtoToComputeBackendServiceBackends(p *computepb.ComputeBackendServiceBackends) *compute.BackendServiceBackends {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceBackends{
		Description:               dcl.StringOrNil(p.Description),
		Group:                     dcl.StringOrNil(p.Group),
		BalancingMode:             ProtoToComputeBackendServiceBackendsBalancingModeEnum(p.GetBalancingMode()),
		MaxUtilization:            dcl.Float64OrNil(p.MaxUtilization),
		MaxRate:                   dcl.Int64OrNil(p.MaxRate),
		MaxRatePerInstance:        dcl.Float64OrNil(p.MaxRatePerInstance),
		MaxRatePerEndpoint:        dcl.Float64OrNil(p.MaxRatePerEndpoint),
		MaxConnections:            dcl.Int64OrNil(p.MaxConnections),
		MaxConnectionsPerInstance: dcl.Int64OrNil(p.MaxConnectionsPerInstance),
		MaxConnectionsPerEndpoint: dcl.Int64OrNil(p.MaxConnectionsPerEndpoint),
		CapacityScaler:            dcl.Float64OrNil(p.CapacityScaler),
		Failover:                  dcl.Bool(p.Failover),
	}
	return obj
}

// ProtoToBackendServiceFailoverPolicy converts a BackendServiceFailoverPolicy resource from its proto representation.
func ProtoToComputeBackendServiceFailoverPolicy(p *computepb.ComputeBackendServiceFailoverPolicy) *compute.BackendServiceFailoverPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceFailoverPolicy{
		DisableConnectionDrainOnFailover: dcl.Bool(p.DisableConnectionDrainOnFailover),
		DropTrafficIfUnhealthy:           dcl.Bool(p.DropTrafficIfUnhealthy),
		FailoverRatio:                    dcl.Float64OrNil(p.FailoverRatio),
	}
	return obj
}

// ProtoToBackendServiceConnectionDraining converts a BackendServiceConnectionDraining resource from its proto representation.
func ProtoToComputeBackendServiceConnectionDraining(p *computepb.ComputeBackendServiceConnectionDraining) *compute.BackendServiceConnectionDraining {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceConnectionDraining{
		DrainingTimeoutSec: dcl.Int64OrNil(p.DrainingTimeoutSec),
	}
	return obj
}

// ProtoToBackendServiceIap converts a BackendServiceIap resource from its proto representation.
func ProtoToComputeBackendServiceIap(p *computepb.ComputeBackendServiceIap) *compute.BackendServiceIap {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceIap{
		Enabled:                  dcl.Bool(p.Enabled),
		OAuth2ClientId:           dcl.StringOrNil(p.Oauth2ClientId),
		OAuth2ClientSecret:       dcl.StringOrNil(p.Oauth2ClientSecret),
		OAuth2ClientSecretSha256: dcl.StringOrNil(p.Oauth2ClientSecretSha256),
	}
	return obj
}

// ProtoToBackendServiceCdnPolicy converts a BackendServiceCdnPolicy resource from its proto representation.
func ProtoToComputeBackendServiceCdnPolicy(p *computepb.ComputeBackendServiceCdnPolicy) *compute.BackendServiceCdnPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceCdnPolicy{
		CacheKeyPolicy:          ProtoToComputeBackendServiceCdnPolicyCacheKeyPolicy(p.GetCacheKeyPolicy()),
		SignedUrlCacheMaxAgeSec: dcl.Int64OrNil(p.SignedUrlCacheMaxAgeSec),
		RequestCoalescing:       dcl.Bool(p.RequestCoalescing),
		CacheMode:               ProtoToComputeBackendServiceCdnPolicyCacheModeEnum(p.GetCacheMode()),
		DefaultTtl:              dcl.Int64OrNil(p.DefaultTtl),
		MaxTtl:                  dcl.Int64OrNil(p.MaxTtl),
		ClientTtl:               dcl.Int64OrNil(p.ClientTtl),
		NegativeCaching:         dcl.Bool(p.NegativeCaching),
		ServeWhileStale:         dcl.Int64OrNil(p.ServeWhileStale),
	}
	for _, r := range p.GetSignedUrlKeyNames() {
		obj.SignedUrlKeyNames = append(obj.SignedUrlKeyNames, r)
	}
	for _, r := range p.GetNegativeCachingPolicy() {
		obj.NegativeCachingPolicy = append(obj.NegativeCachingPolicy, *ProtoToComputeBackendServiceCdnPolicyNegativeCachingPolicy(r))
	}
	for _, r := range p.GetBypassCacheOnRequestHeaders() {
		obj.BypassCacheOnRequestHeaders = append(obj.BypassCacheOnRequestHeaders, *ProtoToComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders(r))
	}
	return obj
}

// ProtoToBackendServiceCdnPolicyCacheKeyPolicy converts a BackendServiceCdnPolicyCacheKeyPolicy resource from its proto representation.
func ProtoToComputeBackendServiceCdnPolicyCacheKeyPolicy(p *computepb.ComputeBackendServiceCdnPolicyCacheKeyPolicy) *compute.BackendServiceCdnPolicyCacheKeyPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceCdnPolicyCacheKeyPolicy{
		IncludeProtocol:    dcl.Bool(p.IncludeProtocol),
		IncludeHost:        dcl.Bool(p.IncludeHost),
		IncludeQueryString: dcl.Bool(p.IncludeQueryString),
	}
	for _, r := range p.GetQueryStringWhitelist() {
		obj.QueryStringWhitelist = append(obj.QueryStringWhitelist, r)
	}
	for _, r := range p.GetQueryStringBlacklist() {
		obj.QueryStringBlacklist = append(obj.QueryStringBlacklist, r)
	}
	for _, r := range p.GetIncludeHttpHeaders() {
		obj.IncludeHttpHeaders = append(obj.IncludeHttpHeaders, r)
	}
	for _, r := range p.GetIncludeNamedCookies() {
		obj.IncludeNamedCookies = append(obj.IncludeNamedCookies, r)
	}
	return obj
}

// ProtoToBackendServiceCdnPolicyNegativeCachingPolicy converts a BackendServiceCdnPolicyNegativeCachingPolicy resource from its proto representation.
func ProtoToComputeBackendServiceCdnPolicyNegativeCachingPolicy(p *computepb.ComputeBackendServiceCdnPolicyNegativeCachingPolicy) *compute.BackendServiceCdnPolicyNegativeCachingPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceCdnPolicyNegativeCachingPolicy{
		Code: dcl.Int64OrNil(p.Code),
		Ttl:  dcl.Int64OrNil(p.Ttl),
	}
	return obj
}

// ProtoToBackendServiceCdnPolicyBypassCacheOnRequestHeaders converts a BackendServiceCdnPolicyBypassCacheOnRequestHeaders resource from its proto representation.
func ProtoToComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders(p *computepb.ComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders) *compute.BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceCdnPolicyBypassCacheOnRequestHeaders{
		HeaderName: dcl.StringOrNil(p.HeaderName),
	}
	return obj
}

// ProtoToBackendServiceLogConfig converts a BackendServiceLogConfig resource from its proto representation.
func ProtoToComputeBackendServiceLogConfig(p *computepb.ComputeBackendServiceLogConfig) *compute.BackendServiceLogConfig {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceLogConfig{
		Enable:     dcl.Bool(p.Enable),
		SampleRate: dcl.Float64OrNil(p.SampleRate),
	}
	return obj
}

// ProtoToBackendServiceSecuritySettings converts a BackendServiceSecuritySettings resource from its proto representation.
func ProtoToComputeBackendServiceSecuritySettings(p *computepb.ComputeBackendServiceSecuritySettings) *compute.BackendServiceSecuritySettings {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceSecuritySettings{
		ClientTlsPolicy: dcl.StringOrNil(p.ClientTlsPolicy),
	}
	for _, r := range p.GetSubjectAltNames() {
		obj.SubjectAltNames = append(obj.SubjectAltNames, r)
	}
	return obj
}

// ProtoToBackendServiceConsistentHash converts a BackendServiceConsistentHash resource from its proto representation.
func ProtoToComputeBackendServiceConsistentHash(p *computepb.ComputeBackendServiceConsistentHash) *compute.BackendServiceConsistentHash {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceConsistentHash{
		HttpCookie:      ProtoToComputeBackendServiceConsistentHashHttpCookie(p.GetHttpCookie()),
		HttpHeaderName:  dcl.StringOrNil(p.HttpHeaderName),
		MinimumRingSize: dcl.Int64OrNil(p.MinimumRingSize),
	}
	return obj
}

// ProtoToBackendServiceConsistentHashHttpCookie converts a BackendServiceConsistentHashHttpCookie resource from its proto representation.
func ProtoToComputeBackendServiceConsistentHashHttpCookie(p *computepb.ComputeBackendServiceConsistentHashHttpCookie) *compute.BackendServiceConsistentHashHttpCookie {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceConsistentHashHttpCookie{
		Name: dcl.StringOrNil(p.Name),
		Path: dcl.StringOrNil(p.Path),
		Ttl:  ProtoToComputeBackendServiceConsistentHashHttpCookieTtl(p.GetTtl()),
	}
	return obj
}

// ProtoToBackendServiceConsistentHashHttpCookieTtl converts a BackendServiceConsistentHashHttpCookieTtl resource from its proto representation.
func ProtoToComputeBackendServiceConsistentHashHttpCookieTtl(p *computepb.ComputeBackendServiceConsistentHashHttpCookieTtl) *compute.BackendServiceConsistentHashHttpCookieTtl {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceConsistentHashHttpCookieTtl{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceCircuitBreakers converts a BackendServiceCircuitBreakers resource from its proto representation.
func ProtoToComputeBackendServiceCircuitBreakers(p *computepb.ComputeBackendServiceCircuitBreakers) *compute.BackendServiceCircuitBreakers {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceCircuitBreakers{
		MaxRequestsPerConnection: dcl.Int64OrNil(p.MaxRequestsPerConnection),
		MaxConnections:           dcl.Int64OrNil(p.MaxConnections),
		MaxPendingRequests:       dcl.Int64OrNil(p.MaxPendingRequests),
		MaxRequests:              dcl.Int64OrNil(p.MaxRequests),
		MaxRetries:               dcl.Int64OrNil(p.MaxRetries),
	}
	return obj
}

// ProtoToBackendServiceOutlierDetection converts a BackendServiceOutlierDetection resource from its proto representation.
func ProtoToComputeBackendServiceOutlierDetection(p *computepb.ComputeBackendServiceOutlierDetection) *compute.BackendServiceOutlierDetection {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceOutlierDetection{
		ConsecutiveErrors:                  dcl.Int64OrNil(p.ConsecutiveErrors),
		Interval:                           ProtoToComputeBackendServiceOutlierDetectionInterval(p.GetInterval()),
		BaseEjectionTime:                   ProtoToComputeBackendServiceOutlierDetectionBaseEjectionTime(p.GetBaseEjectionTime()),
		MaxEjectionPercent:                 dcl.Int64OrNil(p.MaxEjectionPercent),
		EnforcingConsecutiveErrors:         dcl.Int64OrNil(p.EnforcingConsecutiveErrors),
		EnforcingSuccessRate:               dcl.Int64OrNil(p.EnforcingSuccessRate),
		SuccessRateMinimumHosts:            dcl.Int64OrNil(p.SuccessRateMinimumHosts),
		SuccessRateRequestVolume:           dcl.Int64OrNil(p.SuccessRateRequestVolume),
		SuccessRateStdevFactor:             dcl.Int64OrNil(p.SuccessRateStdevFactor),
		ConsecutiveGatewayFailure:          dcl.Int64OrNil(p.ConsecutiveGatewayFailure),
		EnforcingConsecutiveGatewayFailure: dcl.Int64OrNil(p.EnforcingConsecutiveGatewayFailure),
	}
	return obj
}

// ProtoToBackendServiceOutlierDetectionInterval converts a BackendServiceOutlierDetectionInterval resource from its proto representation.
func ProtoToComputeBackendServiceOutlierDetectionInterval(p *computepb.ComputeBackendServiceOutlierDetectionInterval) *compute.BackendServiceOutlierDetectionInterval {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceOutlierDetectionInterval{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceOutlierDetectionBaseEjectionTime converts a BackendServiceOutlierDetectionBaseEjectionTime resource from its proto representation.
func ProtoToComputeBackendServiceOutlierDetectionBaseEjectionTime(p *computepb.ComputeBackendServiceOutlierDetectionBaseEjectionTime) *compute.BackendServiceOutlierDetectionBaseEjectionTime {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceOutlierDetectionBaseEjectionTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceMaxStreamDuration converts a BackendServiceMaxStreamDuration resource from its proto representation.
func ProtoToComputeBackendServiceMaxStreamDuration(p *computepb.ComputeBackendServiceMaxStreamDuration) *compute.BackendServiceMaxStreamDuration {
	if p == nil {
		return nil
	}
	obj := &compute.BackendServiceMaxStreamDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendService converts a BackendService resource from its proto representation.
func ProtoToBackendService(p *computepb.ComputeBackendService) *compute.BackendService {
	obj := &compute.BackendService{
		Id:                   dcl.Int64OrNil(p.Id),
		Name:                 dcl.StringOrNil(p.Name),
		Description:          dcl.StringOrNil(p.Description),
		SelfLink:             dcl.StringOrNil(p.SelfLink),
		SelfLinkWithId:       dcl.StringOrNil(p.SelfLinkWithId),
		TimeoutSec:           dcl.Int64OrNil(p.TimeoutSec),
		Port:                 dcl.Int64OrNil(p.Port),
		Protocol:             ProtoToComputeBackendServiceProtocolEnum(p.GetProtocol()),
		Fingerprint:          dcl.StringOrNil(p.Fingerprint),
		PortName:             dcl.StringOrNil(p.PortName),
		EnableCdn:            dcl.Bool(p.EnableCdn),
		SessionAffinity:      ProtoToComputeBackendServiceSessionAffinityEnum(p.GetSessionAffinity()),
		AffinityCookieTtlSec: dcl.Int64OrNil(p.AffinityCookieTtlSec),
		Location:             dcl.StringOrNil(p.Location),
		FailoverPolicy:       ProtoToComputeBackendServiceFailoverPolicy(p.GetFailoverPolicy()),
		LoadBalancingScheme:  ProtoToComputeBackendServiceLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		ConnectionDraining:   ProtoToComputeBackendServiceConnectionDraining(p.GetConnectionDraining()),
		Iap:                  ProtoToComputeBackendServiceIap(p.GetIap()),
		CdnPolicy:            ProtoToComputeBackendServiceCdnPolicy(p.GetCdnPolicy()),
		SecurityPolicy:       dcl.StringOrNil(p.SecurityPolicy),
		LogConfig:            ProtoToComputeBackendServiceLogConfig(p.GetLogConfig()),
		SecuritySettings:     ProtoToComputeBackendServiceSecuritySettings(p.GetSecuritySettings()),
		LocalityLbPolicy:     ProtoToComputeBackendServiceLocalityLbPolicyEnum(p.GetLocalityLbPolicy()),
		ConsistentHash:       ProtoToComputeBackendServiceConsistentHash(p.GetConsistentHash()),
		CircuitBreakers:      ProtoToComputeBackendServiceCircuitBreakers(p.GetCircuitBreakers()),
		OutlierDetection:     ProtoToComputeBackendServiceOutlierDetection(p.GetOutlierDetection()),
		Network:              dcl.StringOrNil(p.Network),
		MaxStreamDuration:    ProtoToComputeBackendServiceMaxStreamDuration(p.GetMaxStreamDuration()),
		Project:              dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToComputeBackendServiceBackends(r))
	}
	for _, r := range p.GetHealthChecks() {
		obj.HealthChecks = append(obj.HealthChecks, r)
	}
	for _, r := range p.GetCustomRequestHeaders() {
		obj.CustomRequestHeaders = append(obj.CustomRequestHeaders, r)
	}
	for _, r := range p.GetCustomResponseHeaders() {
		obj.CustomResponseHeaders = append(obj.CustomResponseHeaders, r)
	}
	return obj
}

// BackendServiceBackendsBalancingModeEnumToProto converts a BackendServiceBackendsBalancingModeEnum enum to its proto representation.
func ComputeBackendServiceBackendsBalancingModeEnumToProto(e *compute.BackendServiceBackendsBalancingModeEnum) computepb.ComputeBackendServiceBackendsBalancingModeEnum {
	if e == nil {
		return computepb.ComputeBackendServiceBackendsBalancingModeEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceBackendsBalancingModeEnum_value["BackendServiceBackendsBalancingModeEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceBackendsBalancingModeEnum(v)
	}
	return computepb.ComputeBackendServiceBackendsBalancingModeEnum(0)
}

// BackendServiceProtocolEnumToProto converts a BackendServiceProtocolEnum enum to its proto representation.
func ComputeBackendServiceProtocolEnumToProto(e *compute.BackendServiceProtocolEnum) computepb.ComputeBackendServiceProtocolEnum {
	if e == nil {
		return computepb.ComputeBackendServiceProtocolEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceProtocolEnum_value["BackendServiceProtocolEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceProtocolEnum(v)
	}
	return computepb.ComputeBackendServiceProtocolEnum(0)
}

// BackendServiceSessionAffinityEnumToProto converts a BackendServiceSessionAffinityEnum enum to its proto representation.
func ComputeBackendServiceSessionAffinityEnumToProto(e *compute.BackendServiceSessionAffinityEnum) computepb.ComputeBackendServiceSessionAffinityEnum {
	if e == nil {
		return computepb.ComputeBackendServiceSessionAffinityEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceSessionAffinityEnum_value["BackendServiceSessionAffinityEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceSessionAffinityEnum(v)
	}
	return computepb.ComputeBackendServiceSessionAffinityEnum(0)
}

// BackendServiceLoadBalancingSchemeEnumToProto converts a BackendServiceLoadBalancingSchemeEnum enum to its proto representation.
func ComputeBackendServiceLoadBalancingSchemeEnumToProto(e *compute.BackendServiceLoadBalancingSchemeEnum) computepb.ComputeBackendServiceLoadBalancingSchemeEnum {
	if e == nil {
		return computepb.ComputeBackendServiceLoadBalancingSchemeEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceLoadBalancingSchemeEnum_value["BackendServiceLoadBalancingSchemeEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceLoadBalancingSchemeEnum(v)
	}
	return computepb.ComputeBackendServiceLoadBalancingSchemeEnum(0)
}

// BackendServiceCdnPolicyCacheModeEnumToProto converts a BackendServiceCdnPolicyCacheModeEnum enum to its proto representation.
func ComputeBackendServiceCdnPolicyCacheModeEnumToProto(e *compute.BackendServiceCdnPolicyCacheModeEnum) computepb.ComputeBackendServiceCdnPolicyCacheModeEnum {
	if e == nil {
		return computepb.ComputeBackendServiceCdnPolicyCacheModeEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceCdnPolicyCacheModeEnum_value["BackendServiceCdnPolicyCacheModeEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceCdnPolicyCacheModeEnum(v)
	}
	return computepb.ComputeBackendServiceCdnPolicyCacheModeEnum(0)
}

// BackendServiceLocalityLbPolicyEnumToProto converts a BackendServiceLocalityLbPolicyEnum enum to its proto representation.
func ComputeBackendServiceLocalityLbPolicyEnumToProto(e *compute.BackendServiceLocalityLbPolicyEnum) computepb.ComputeBackendServiceLocalityLbPolicyEnum {
	if e == nil {
		return computepb.ComputeBackendServiceLocalityLbPolicyEnum(0)
	}
	if v, ok := computepb.ComputeBackendServiceLocalityLbPolicyEnum_value["BackendServiceLocalityLbPolicyEnum"+string(*e)]; ok {
		return computepb.ComputeBackendServiceLocalityLbPolicyEnum(v)
	}
	return computepb.ComputeBackendServiceLocalityLbPolicyEnum(0)
}

// BackendServiceBackendsToProto converts a BackendServiceBackends resource to its proto representation.
func ComputeBackendServiceBackendsToProto(o *compute.BackendServiceBackends) *computepb.ComputeBackendServiceBackends {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceBackends{
		Description:               dcl.ValueOrEmptyString(o.Description),
		Group:                     dcl.ValueOrEmptyString(o.Group),
		BalancingMode:             ComputeBackendServiceBackendsBalancingModeEnumToProto(o.BalancingMode),
		MaxUtilization:            dcl.ValueOrEmptyDouble(o.MaxUtilization),
		MaxRate:                   dcl.ValueOrEmptyInt64(o.MaxRate),
		MaxRatePerInstance:        dcl.ValueOrEmptyDouble(o.MaxRatePerInstance),
		MaxRatePerEndpoint:        dcl.ValueOrEmptyDouble(o.MaxRatePerEndpoint),
		MaxConnections:            dcl.ValueOrEmptyInt64(o.MaxConnections),
		MaxConnectionsPerInstance: dcl.ValueOrEmptyInt64(o.MaxConnectionsPerInstance),
		MaxConnectionsPerEndpoint: dcl.ValueOrEmptyInt64(o.MaxConnectionsPerEndpoint),
		CapacityScaler:            dcl.ValueOrEmptyDouble(o.CapacityScaler),
		Failover:                  dcl.ValueOrEmptyBool(o.Failover),
	}
	return p
}

// BackendServiceFailoverPolicyToProto converts a BackendServiceFailoverPolicy resource to its proto representation.
func ComputeBackendServiceFailoverPolicyToProto(o *compute.BackendServiceFailoverPolicy) *computepb.ComputeBackendServiceFailoverPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceFailoverPolicy{
		DisableConnectionDrainOnFailover: dcl.ValueOrEmptyBool(o.DisableConnectionDrainOnFailover),
		DropTrafficIfUnhealthy:           dcl.ValueOrEmptyBool(o.DropTrafficIfUnhealthy),
		FailoverRatio:                    dcl.ValueOrEmptyDouble(o.FailoverRatio),
	}
	return p
}

// BackendServiceConnectionDrainingToProto converts a BackendServiceConnectionDraining resource to its proto representation.
func ComputeBackendServiceConnectionDrainingToProto(o *compute.BackendServiceConnectionDraining) *computepb.ComputeBackendServiceConnectionDraining {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceConnectionDraining{
		DrainingTimeoutSec: dcl.ValueOrEmptyInt64(o.DrainingTimeoutSec),
	}
	return p
}

// BackendServiceIapToProto converts a BackendServiceIap resource to its proto representation.
func ComputeBackendServiceIapToProto(o *compute.BackendServiceIap) *computepb.ComputeBackendServiceIap {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceIap{
		Enabled:                  dcl.ValueOrEmptyBool(o.Enabled),
		Oauth2ClientId:           dcl.ValueOrEmptyString(o.OAuth2ClientId),
		Oauth2ClientSecret:       dcl.ValueOrEmptyString(o.OAuth2ClientSecret),
		Oauth2ClientSecretSha256: dcl.ValueOrEmptyString(o.OAuth2ClientSecretSha256),
	}
	return p
}

// BackendServiceCdnPolicyToProto converts a BackendServiceCdnPolicy resource to its proto representation.
func ComputeBackendServiceCdnPolicyToProto(o *compute.BackendServiceCdnPolicy) *computepb.ComputeBackendServiceCdnPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceCdnPolicy{
		CacheKeyPolicy:          ComputeBackendServiceCdnPolicyCacheKeyPolicyToProto(o.CacheKeyPolicy),
		SignedUrlCacheMaxAgeSec: dcl.ValueOrEmptyInt64(o.SignedUrlCacheMaxAgeSec),
		RequestCoalescing:       dcl.ValueOrEmptyBool(o.RequestCoalescing),
		CacheMode:               ComputeBackendServiceCdnPolicyCacheModeEnumToProto(o.CacheMode),
		DefaultTtl:              dcl.ValueOrEmptyInt64(o.DefaultTtl),
		MaxTtl:                  dcl.ValueOrEmptyInt64(o.MaxTtl),
		ClientTtl:               dcl.ValueOrEmptyInt64(o.ClientTtl),
		NegativeCaching:         dcl.ValueOrEmptyBool(o.NegativeCaching),
		ServeWhileStale:         dcl.ValueOrEmptyInt64(o.ServeWhileStale),
	}
	for _, r := range o.SignedUrlKeyNames {
		p.SignedUrlKeyNames = append(p.SignedUrlKeyNames, r)
	}
	for _, r := range o.NegativeCachingPolicy {
		p.NegativeCachingPolicy = append(p.NegativeCachingPolicy, ComputeBackendServiceCdnPolicyNegativeCachingPolicyToProto(&r))
	}
	for _, r := range o.BypassCacheOnRequestHeaders {
		p.BypassCacheOnRequestHeaders = append(p.BypassCacheOnRequestHeaders, ComputeBackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto(&r))
	}
	return p
}

// BackendServiceCdnPolicyCacheKeyPolicyToProto converts a BackendServiceCdnPolicyCacheKeyPolicy resource to its proto representation.
func ComputeBackendServiceCdnPolicyCacheKeyPolicyToProto(o *compute.BackendServiceCdnPolicyCacheKeyPolicy) *computepb.ComputeBackendServiceCdnPolicyCacheKeyPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceCdnPolicyCacheKeyPolicy{
		IncludeProtocol:    dcl.ValueOrEmptyBool(o.IncludeProtocol),
		IncludeHost:        dcl.ValueOrEmptyBool(o.IncludeHost),
		IncludeQueryString: dcl.ValueOrEmptyBool(o.IncludeQueryString),
	}
	for _, r := range o.QueryStringWhitelist {
		p.QueryStringWhitelist = append(p.QueryStringWhitelist, r)
	}
	for _, r := range o.QueryStringBlacklist {
		p.QueryStringBlacklist = append(p.QueryStringBlacklist, r)
	}
	for _, r := range o.IncludeHttpHeaders {
		p.IncludeHttpHeaders = append(p.IncludeHttpHeaders, r)
	}
	for _, r := range o.IncludeNamedCookies {
		p.IncludeNamedCookies = append(p.IncludeNamedCookies, r)
	}
	return p
}

// BackendServiceCdnPolicyNegativeCachingPolicyToProto converts a BackendServiceCdnPolicyNegativeCachingPolicy resource to its proto representation.
func ComputeBackendServiceCdnPolicyNegativeCachingPolicyToProto(o *compute.BackendServiceCdnPolicyNegativeCachingPolicy) *computepb.ComputeBackendServiceCdnPolicyNegativeCachingPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceCdnPolicyNegativeCachingPolicy{
		Code: dcl.ValueOrEmptyInt64(o.Code),
		Ttl:  dcl.ValueOrEmptyInt64(o.Ttl),
	}
	return p
}

// BackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto converts a BackendServiceCdnPolicyBypassCacheOnRequestHeaders resource to its proto representation.
func ComputeBackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto(o *compute.BackendServiceCdnPolicyBypassCacheOnRequestHeaders) *computepb.ComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders{
		HeaderName: dcl.ValueOrEmptyString(o.HeaderName),
	}
	return p
}

// BackendServiceLogConfigToProto converts a BackendServiceLogConfig resource to its proto representation.
func ComputeBackendServiceLogConfigToProto(o *compute.BackendServiceLogConfig) *computepb.ComputeBackendServiceLogConfig {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceLogConfig{
		Enable:     dcl.ValueOrEmptyBool(o.Enable),
		SampleRate: dcl.ValueOrEmptyDouble(o.SampleRate),
	}
	return p
}

// BackendServiceSecuritySettingsToProto converts a BackendServiceSecuritySettings resource to its proto representation.
func ComputeBackendServiceSecuritySettingsToProto(o *compute.BackendServiceSecuritySettings) *computepb.ComputeBackendServiceSecuritySettings {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceSecuritySettings{
		ClientTlsPolicy: dcl.ValueOrEmptyString(o.ClientTlsPolicy),
	}
	for _, r := range o.SubjectAltNames {
		p.SubjectAltNames = append(p.SubjectAltNames, r)
	}
	return p
}

// BackendServiceConsistentHashToProto converts a BackendServiceConsistentHash resource to its proto representation.
func ComputeBackendServiceConsistentHashToProto(o *compute.BackendServiceConsistentHash) *computepb.ComputeBackendServiceConsistentHash {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceConsistentHash{
		HttpCookie:      ComputeBackendServiceConsistentHashHttpCookieToProto(o.HttpCookie),
		HttpHeaderName:  dcl.ValueOrEmptyString(o.HttpHeaderName),
		MinimumRingSize: dcl.ValueOrEmptyInt64(o.MinimumRingSize),
	}
	return p
}

// BackendServiceConsistentHashHttpCookieToProto converts a BackendServiceConsistentHashHttpCookie resource to its proto representation.
func ComputeBackendServiceConsistentHashHttpCookieToProto(o *compute.BackendServiceConsistentHashHttpCookie) *computepb.ComputeBackendServiceConsistentHashHttpCookie {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceConsistentHashHttpCookie{
		Name: dcl.ValueOrEmptyString(o.Name),
		Path: dcl.ValueOrEmptyString(o.Path),
		Ttl:  ComputeBackendServiceConsistentHashHttpCookieTtlToProto(o.Ttl),
	}
	return p
}

// BackendServiceConsistentHashHttpCookieTtlToProto converts a BackendServiceConsistentHashHttpCookieTtl resource to its proto representation.
func ComputeBackendServiceConsistentHashHttpCookieTtlToProto(o *compute.BackendServiceConsistentHashHttpCookieTtl) *computepb.ComputeBackendServiceConsistentHashHttpCookieTtl {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceConsistentHashHttpCookieTtl{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceCircuitBreakersToProto converts a BackendServiceCircuitBreakers resource to its proto representation.
func ComputeBackendServiceCircuitBreakersToProto(o *compute.BackendServiceCircuitBreakers) *computepb.ComputeBackendServiceCircuitBreakers {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceCircuitBreakers{
		MaxRequestsPerConnection: dcl.ValueOrEmptyInt64(o.MaxRequestsPerConnection),
		MaxConnections:           dcl.ValueOrEmptyInt64(o.MaxConnections),
		MaxPendingRequests:       dcl.ValueOrEmptyInt64(o.MaxPendingRequests),
		MaxRequests:              dcl.ValueOrEmptyInt64(o.MaxRequests),
		MaxRetries:               dcl.ValueOrEmptyInt64(o.MaxRetries),
	}
	return p
}

// BackendServiceOutlierDetectionToProto converts a BackendServiceOutlierDetection resource to its proto representation.
func ComputeBackendServiceOutlierDetectionToProto(o *compute.BackendServiceOutlierDetection) *computepb.ComputeBackendServiceOutlierDetection {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceOutlierDetection{
		ConsecutiveErrors:                  dcl.ValueOrEmptyInt64(o.ConsecutiveErrors),
		Interval:                           ComputeBackendServiceOutlierDetectionIntervalToProto(o.Interval),
		BaseEjectionTime:                   ComputeBackendServiceOutlierDetectionBaseEjectionTimeToProto(o.BaseEjectionTime),
		MaxEjectionPercent:                 dcl.ValueOrEmptyInt64(o.MaxEjectionPercent),
		EnforcingConsecutiveErrors:         dcl.ValueOrEmptyInt64(o.EnforcingConsecutiveErrors),
		EnforcingSuccessRate:               dcl.ValueOrEmptyInt64(o.EnforcingSuccessRate),
		SuccessRateMinimumHosts:            dcl.ValueOrEmptyInt64(o.SuccessRateMinimumHosts),
		SuccessRateRequestVolume:           dcl.ValueOrEmptyInt64(o.SuccessRateRequestVolume),
		SuccessRateStdevFactor:             dcl.ValueOrEmptyInt64(o.SuccessRateStdevFactor),
		ConsecutiveGatewayFailure:          dcl.ValueOrEmptyInt64(o.ConsecutiveGatewayFailure),
		EnforcingConsecutiveGatewayFailure: dcl.ValueOrEmptyInt64(o.EnforcingConsecutiveGatewayFailure),
	}
	return p
}

// BackendServiceOutlierDetectionIntervalToProto converts a BackendServiceOutlierDetectionInterval resource to its proto representation.
func ComputeBackendServiceOutlierDetectionIntervalToProto(o *compute.BackendServiceOutlierDetectionInterval) *computepb.ComputeBackendServiceOutlierDetectionInterval {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceOutlierDetectionInterval{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceOutlierDetectionBaseEjectionTimeToProto converts a BackendServiceOutlierDetectionBaseEjectionTime resource to its proto representation.
func ComputeBackendServiceOutlierDetectionBaseEjectionTimeToProto(o *compute.BackendServiceOutlierDetectionBaseEjectionTime) *computepb.ComputeBackendServiceOutlierDetectionBaseEjectionTime {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceOutlierDetectionBaseEjectionTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceMaxStreamDurationToProto converts a BackendServiceMaxStreamDuration resource to its proto representation.
func ComputeBackendServiceMaxStreamDurationToProto(o *compute.BackendServiceMaxStreamDuration) *computepb.ComputeBackendServiceMaxStreamDuration {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeBackendServiceMaxStreamDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceToProto converts a BackendService resource to its proto representation.
func BackendServiceToProto(resource *compute.BackendService) *computepb.ComputeBackendService {
	p := &computepb.ComputeBackendService{
		Id:                   dcl.ValueOrEmptyInt64(resource.Id),
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Description:          dcl.ValueOrEmptyString(resource.Description),
		SelfLink:             dcl.ValueOrEmptyString(resource.SelfLink),
		SelfLinkWithId:       dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		TimeoutSec:           dcl.ValueOrEmptyInt64(resource.TimeoutSec),
		Port:                 dcl.ValueOrEmptyInt64(resource.Port),
		Protocol:             ComputeBackendServiceProtocolEnumToProto(resource.Protocol),
		Fingerprint:          dcl.ValueOrEmptyString(resource.Fingerprint),
		PortName:             dcl.ValueOrEmptyString(resource.PortName),
		EnableCdn:            dcl.ValueOrEmptyBool(resource.EnableCdn),
		SessionAffinity:      ComputeBackendServiceSessionAffinityEnumToProto(resource.SessionAffinity),
		AffinityCookieTtlSec: dcl.ValueOrEmptyInt64(resource.AffinityCookieTtlSec),
		Location:             dcl.ValueOrEmptyString(resource.Location),
		FailoverPolicy:       ComputeBackendServiceFailoverPolicyToProto(resource.FailoverPolicy),
		LoadBalancingScheme:  ComputeBackendServiceLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme),
		ConnectionDraining:   ComputeBackendServiceConnectionDrainingToProto(resource.ConnectionDraining),
		Iap:                  ComputeBackendServiceIapToProto(resource.Iap),
		CdnPolicy:            ComputeBackendServiceCdnPolicyToProto(resource.CdnPolicy),
		SecurityPolicy:       dcl.ValueOrEmptyString(resource.SecurityPolicy),
		LogConfig:            ComputeBackendServiceLogConfigToProto(resource.LogConfig),
		SecuritySettings:     ComputeBackendServiceSecuritySettingsToProto(resource.SecuritySettings),
		LocalityLbPolicy:     ComputeBackendServiceLocalityLbPolicyEnumToProto(resource.LocalityLbPolicy),
		ConsistentHash:       ComputeBackendServiceConsistentHashToProto(resource.ConsistentHash),
		CircuitBreakers:      ComputeBackendServiceCircuitBreakersToProto(resource.CircuitBreakers),
		OutlierDetection:     ComputeBackendServiceOutlierDetectionToProto(resource.OutlierDetection),
		Network:              dcl.ValueOrEmptyString(resource.Network),
		MaxStreamDuration:    ComputeBackendServiceMaxStreamDurationToProto(resource.MaxStreamDuration),
		Project:              dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Backends {
		p.Backends = append(p.Backends, ComputeBackendServiceBackendsToProto(&r))
	}
	for _, r := range resource.HealthChecks {
		p.HealthChecks = append(p.HealthChecks, r)
	}
	for _, r := range resource.CustomRequestHeaders {
		p.CustomRequestHeaders = append(p.CustomRequestHeaders, r)
	}
	for _, r := range resource.CustomResponseHeaders {
		p.CustomResponseHeaders = append(p.CustomResponseHeaders, r)
	}

	return p
}

// ApplyBackendService handles the gRPC request by passing it to the underlying BackendService Apply() method.
func (s *BackendServiceServer) applyBackendService(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeBackendServiceRequest) (*computepb.ComputeBackendService, error) {
	p := ProtoToBackendService(request.GetResource())
	res, err := c.ApplyBackendService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackendServiceToProto(res)
	return r, nil
}

// ApplyBackendService handles the gRPC request by passing it to the underlying BackendService Apply() method.
func (s *BackendServiceServer) ApplyComputeBackendService(ctx context.Context, request *computepb.ApplyComputeBackendServiceRequest) (*computepb.ComputeBackendService, error) {
	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBackendService(ctx, cl, request)
}

// DeleteBackendService handles the gRPC request by passing it to the underlying BackendService Delete() method.
func (s *BackendServiceServer) DeleteComputeBackendService(ctx context.Context, request *computepb.DeleteComputeBackendServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackendService(ctx, ProtoToBackendService(request.GetResource()))

}

// ListComputeBackendService handles the gRPC request by passing it to the underlying BackendServiceList() method.
func (s *BackendServiceServer) ListComputeBackendService(ctx context.Context, request *computepb.ListComputeBackendServiceRequest) (*computepb.ListComputeBackendServiceResponse, error) {
	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackendService(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeBackendService
	for _, r := range resources.Items {
		rp := BackendServiceToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeBackendServiceResponse{Items: protos}, nil
}

func createConfigBackendService(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
