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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for BackendService.
type BackendServiceServer struct{}

// ProtoToBackendServiceBackendsBalancingModeEnum converts a BackendServiceBackendsBalancingModeEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceBackendsBalancingModeEnum(e betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum) *beta.BackendServiceBackendsBalancingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum_name[int32(e)]; ok {
		e := beta.BackendServiceBackendsBalancingModeEnum(n[len("ComputeBetaBackendServiceBackendsBalancingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceProtocolEnum converts a BackendServiceProtocolEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceProtocolEnum(e betapb.ComputeBetaBackendServiceProtocolEnum) *beta.BackendServiceProtocolEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceProtocolEnum_name[int32(e)]; ok {
		e := beta.BackendServiceProtocolEnum(n[len("ComputeBetaBackendServiceProtocolEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceSessionAffinityEnum converts a BackendServiceSessionAffinityEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceSessionAffinityEnum(e betapb.ComputeBetaBackendServiceSessionAffinityEnum) *beta.BackendServiceSessionAffinityEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceSessionAffinityEnum_name[int32(e)]; ok {
		e := beta.BackendServiceSessionAffinityEnum(n[len("ComputeBetaBackendServiceSessionAffinityEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceLoadBalancingSchemeEnum converts a BackendServiceLoadBalancingSchemeEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceLoadBalancingSchemeEnum(e betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum) *beta.BackendServiceLoadBalancingSchemeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum_name[int32(e)]; ok {
		e := beta.BackendServiceLoadBalancingSchemeEnum(n[len("ComputeBetaBackendServiceLoadBalancingSchemeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceCdnPolicyCacheModeEnum converts a BackendServiceCdnPolicyCacheModeEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceCdnPolicyCacheModeEnum(e betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum) *beta.BackendServiceCdnPolicyCacheModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum_name[int32(e)]; ok {
		e := beta.BackendServiceCdnPolicyCacheModeEnum(n[len("ComputeBetaBackendServiceCdnPolicyCacheModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceLocalityLbPolicyEnum converts a BackendServiceLocalityLbPolicyEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceLocalityLbPolicyEnum(e betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum) *beta.BackendServiceLocalityLbPolicyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum_name[int32(e)]; ok {
		e := beta.BackendServiceLocalityLbPolicyEnum(n[len("ComputeBetaBackendServiceLocalityLbPolicyEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceSubsettingPolicyEnum converts a BackendServiceSubsettingPolicyEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceSubsettingPolicyEnum(e betapb.ComputeBetaBackendServiceSubsettingPolicyEnum) *beta.BackendServiceSubsettingPolicyEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceSubsettingPolicyEnum_name[int32(e)]; ok {
		e := beta.BackendServiceSubsettingPolicyEnum(n[len("ComputeBetaBackendServiceSubsettingPolicyEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceConnectionTrackingPolicyTrackingModeEnum converts a BackendServiceConnectionTrackingPolicyTrackingModeEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum(e betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum) *beta.BackendServiceConnectionTrackingPolicyTrackingModeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum_name[int32(e)]; ok {
		e := beta.BackendServiceConnectionTrackingPolicyTrackingModeEnum(n[len("ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum converts a BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum enum from its proto representation.
func ProtoToComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(e betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum) *beta.BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum_name[int32(e)]; ok {
		e := beta.BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(n[len("ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackendServiceBackends converts a BackendServiceBackends resource from its proto representation.
func ProtoToComputeBetaBackendServiceBackends(p *betapb.ComputeBetaBackendServiceBackends) *beta.BackendServiceBackends {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceBackends{
		Description:               dcl.StringOrNil(p.Description),
		Group:                     dcl.StringOrNil(p.Group),
		BalancingMode:             ProtoToComputeBetaBackendServiceBackendsBalancingModeEnum(p.GetBalancingMode()),
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
func ProtoToComputeBetaBackendServiceFailoverPolicy(p *betapb.ComputeBetaBackendServiceFailoverPolicy) *beta.BackendServiceFailoverPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceFailoverPolicy{
		DisableConnectionDrainOnFailover: dcl.Bool(p.DisableConnectionDrainOnFailover),
		DropTrafficIfUnhealthy:           dcl.Bool(p.DropTrafficIfUnhealthy),
		FailoverRatio:                    dcl.Float64OrNil(p.FailoverRatio),
	}
	return obj
}

// ProtoToBackendServiceConnectionDraining converts a BackendServiceConnectionDraining resource from its proto representation.
func ProtoToComputeBetaBackendServiceConnectionDraining(p *betapb.ComputeBetaBackendServiceConnectionDraining) *beta.BackendServiceConnectionDraining {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceConnectionDraining{
		DrainingTimeoutSec: dcl.Int64OrNil(p.DrainingTimeoutSec),
	}
	return obj
}

// ProtoToBackendServiceIap converts a BackendServiceIap resource from its proto representation.
func ProtoToComputeBetaBackendServiceIap(p *betapb.ComputeBetaBackendServiceIap) *beta.BackendServiceIap {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceIap{
		Enabled:                  dcl.Bool(p.Enabled),
		OAuth2ClientId:           dcl.StringOrNil(p.Oauth2ClientId),
		OAuth2ClientSecret:       dcl.StringOrNil(p.Oauth2ClientSecret),
		OAuth2ClientSecretSha256: dcl.StringOrNil(p.Oauth2ClientSecretSha256),
	}
	return obj
}

// ProtoToBackendServiceCdnPolicy converts a BackendServiceCdnPolicy resource from its proto representation.
func ProtoToComputeBetaBackendServiceCdnPolicy(p *betapb.ComputeBetaBackendServiceCdnPolicy) *beta.BackendServiceCdnPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCdnPolicy{
		CacheKeyPolicy:          ProtoToComputeBetaBackendServiceCdnPolicyCacheKeyPolicy(p.GetCacheKeyPolicy()),
		SignedUrlCacheMaxAgeSec: dcl.Int64OrNil(p.SignedUrlCacheMaxAgeSec),
		RequestCoalescing:       dcl.Bool(p.RequestCoalescing),
		CacheMode:               ProtoToComputeBetaBackendServiceCdnPolicyCacheModeEnum(p.GetCacheMode()),
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
		obj.NegativeCachingPolicy = append(obj.NegativeCachingPolicy, *ProtoToComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy(r))
	}
	for _, r := range p.GetBypassCacheOnRequestHeaders() {
		obj.BypassCacheOnRequestHeaders = append(obj.BypassCacheOnRequestHeaders, *ProtoToComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders(r))
	}
	return obj
}

// ProtoToBackendServiceCdnPolicyCacheKeyPolicy converts a BackendServiceCdnPolicyCacheKeyPolicy resource from its proto representation.
func ProtoToComputeBetaBackendServiceCdnPolicyCacheKeyPolicy(p *betapb.ComputeBetaBackendServiceCdnPolicyCacheKeyPolicy) *beta.BackendServiceCdnPolicyCacheKeyPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCdnPolicyCacheKeyPolicy{
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
func ProtoToComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy(p *betapb.ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy) *beta.BackendServiceCdnPolicyNegativeCachingPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCdnPolicyNegativeCachingPolicy{
		Code: dcl.Int64OrNil(p.Code),
		Ttl:  dcl.Int64OrNil(p.Ttl),
	}
	return obj
}

// ProtoToBackendServiceCdnPolicyBypassCacheOnRequestHeaders converts a BackendServiceCdnPolicyBypassCacheOnRequestHeaders resource from its proto representation.
func ProtoToComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders(p *betapb.ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders) *beta.BackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCdnPolicyBypassCacheOnRequestHeaders{
		HeaderName: dcl.StringOrNil(p.HeaderName),
	}
	return obj
}

// ProtoToBackendServiceLogConfig converts a BackendServiceLogConfig resource from its proto representation.
func ProtoToComputeBetaBackendServiceLogConfig(p *betapb.ComputeBetaBackendServiceLogConfig) *beta.BackendServiceLogConfig {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceLogConfig{
		Enable:     dcl.Bool(p.Enable),
		SampleRate: dcl.Float64OrNil(p.SampleRate),
	}
	return obj
}

// ProtoToBackendServiceSecuritySettings converts a BackendServiceSecuritySettings resource from its proto representation.
func ProtoToComputeBetaBackendServiceSecuritySettings(p *betapb.ComputeBetaBackendServiceSecuritySettings) *beta.BackendServiceSecuritySettings {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceSecuritySettings{
		ClientTlsPolicy: dcl.StringOrNil(p.ClientTlsPolicy),
		Authentication:  dcl.StringOrNil(p.Authentication),
	}
	for _, r := range p.GetSubjectAltNames() {
		obj.SubjectAltNames = append(obj.SubjectAltNames, r)
	}
	return obj
}

// ProtoToBackendServiceConsistentHash converts a BackendServiceConsistentHash resource from its proto representation.
func ProtoToComputeBetaBackendServiceConsistentHash(p *betapb.ComputeBetaBackendServiceConsistentHash) *beta.BackendServiceConsistentHash {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceConsistentHash{
		HttpCookie:      ProtoToComputeBetaBackendServiceConsistentHashHttpCookie(p.GetHttpCookie()),
		HttpHeaderName:  dcl.StringOrNil(p.HttpHeaderName),
		MinimumRingSize: dcl.Int64OrNil(p.MinimumRingSize),
	}
	return obj
}

// ProtoToBackendServiceConsistentHashHttpCookie converts a BackendServiceConsistentHashHttpCookie resource from its proto representation.
func ProtoToComputeBetaBackendServiceConsistentHashHttpCookie(p *betapb.ComputeBetaBackendServiceConsistentHashHttpCookie) *beta.BackendServiceConsistentHashHttpCookie {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceConsistentHashHttpCookie{
		Name: dcl.StringOrNil(p.Name),
		Path: dcl.StringOrNil(p.Path),
		Ttl:  ProtoToComputeBetaBackendServiceConsistentHashHttpCookieTtl(p.GetTtl()),
	}
	return obj
}

// ProtoToBackendServiceConsistentHashHttpCookieTtl converts a BackendServiceConsistentHashHttpCookieTtl resource from its proto representation.
func ProtoToComputeBetaBackendServiceConsistentHashHttpCookieTtl(p *betapb.ComputeBetaBackendServiceConsistentHashHttpCookieTtl) *beta.BackendServiceConsistentHashHttpCookieTtl {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceConsistentHashHttpCookieTtl{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceCircuitBreakers converts a BackendServiceCircuitBreakers resource from its proto representation.
func ProtoToComputeBetaBackendServiceCircuitBreakers(p *betapb.ComputeBetaBackendServiceCircuitBreakers) *beta.BackendServiceCircuitBreakers {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCircuitBreakers{
		ConnectTimeout:           ProtoToComputeBetaBackendServiceCircuitBreakersConnectTimeout(p.GetConnectTimeout()),
		MaxRequestsPerConnection: dcl.Int64OrNil(p.MaxRequestsPerConnection),
		MaxConnections:           dcl.Int64OrNil(p.MaxConnections),
		MaxPendingRequests:       dcl.Int64OrNil(p.MaxPendingRequests),
		MaxRequests:              dcl.Int64OrNil(p.MaxRequests),
		MaxRetries:               dcl.Int64OrNil(p.MaxRetries),
	}
	return obj
}

// ProtoToBackendServiceCircuitBreakersConnectTimeout converts a BackendServiceCircuitBreakersConnectTimeout resource from its proto representation.
func ProtoToComputeBetaBackendServiceCircuitBreakersConnectTimeout(p *betapb.ComputeBetaBackendServiceCircuitBreakersConnectTimeout) *beta.BackendServiceCircuitBreakersConnectTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceCircuitBreakersConnectTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceOutlierDetection converts a BackendServiceOutlierDetection resource from its proto representation.
func ProtoToComputeBetaBackendServiceOutlierDetection(p *betapb.ComputeBetaBackendServiceOutlierDetection) *beta.BackendServiceOutlierDetection {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceOutlierDetection{
		ConsecutiveErrors:                  dcl.Int64OrNil(p.ConsecutiveErrors),
		Interval:                           ProtoToComputeBetaBackendServiceOutlierDetectionInterval(p.GetInterval()),
		BaseEjectionTime:                   ProtoToComputeBetaBackendServiceOutlierDetectionBaseEjectionTime(p.GetBaseEjectionTime()),
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
func ProtoToComputeBetaBackendServiceOutlierDetectionInterval(p *betapb.ComputeBetaBackendServiceOutlierDetectionInterval) *beta.BackendServiceOutlierDetectionInterval {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceOutlierDetectionInterval{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceOutlierDetectionBaseEjectionTime converts a BackendServiceOutlierDetectionBaseEjectionTime resource from its proto representation.
func ProtoToComputeBetaBackendServiceOutlierDetectionBaseEjectionTime(p *betapb.ComputeBetaBackendServiceOutlierDetectionBaseEjectionTime) *beta.BackendServiceOutlierDetectionBaseEjectionTime {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceOutlierDetectionBaseEjectionTime{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendServiceSubsetting converts a BackendServiceSubsetting resource from its proto representation.
func ProtoToComputeBetaBackendServiceSubsetting(p *betapb.ComputeBetaBackendServiceSubsetting) *beta.BackendServiceSubsetting {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceSubsetting{
		Policy: ProtoToComputeBetaBackendServiceSubsettingPolicyEnum(p.GetPolicy()),
	}
	return obj
}

// ProtoToBackendServiceConnectionTrackingPolicy converts a BackendServiceConnectionTrackingPolicy resource from its proto representation.
func ProtoToComputeBetaBackendServiceConnectionTrackingPolicy(p *betapb.ComputeBetaBackendServiceConnectionTrackingPolicy) *beta.BackendServiceConnectionTrackingPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceConnectionTrackingPolicy{
		TrackingMode:                             ProtoToComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum(p.GetTrackingMode()),
		ConnectionPersistenceOnUnhealthyBackends: ProtoToComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(p.GetConnectionPersistenceOnUnhealthyBackends()),
		IdleTimeoutSec:                           dcl.Int64OrNil(p.IdleTimeoutSec),
	}
	return obj
}

// ProtoToBackendServiceMaxStreamDuration converts a BackendServiceMaxStreamDuration resource from its proto representation.
func ProtoToComputeBetaBackendServiceMaxStreamDuration(p *betapb.ComputeBetaBackendServiceMaxStreamDuration) *beta.BackendServiceMaxStreamDuration {
	if p == nil {
		return nil
	}
	obj := &beta.BackendServiceMaxStreamDuration{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToBackendService converts a BackendService resource from its proto representation.
func ProtoToBackendService(p *betapb.ComputeBetaBackendService) *beta.BackendService {
	obj := &beta.BackendService{
		Name:                     dcl.StringOrNil(p.Name),
		Description:              dcl.StringOrNil(p.Description),
		SelfLink:                 dcl.StringOrNil(p.SelfLink),
		SelfLinkWithId:           dcl.StringOrNil(p.SelfLinkWithId),
		TimeoutSec:               dcl.Int64OrNil(p.TimeoutSec),
		Port:                     dcl.Int64OrNil(p.Port),
		Protocol:                 ProtoToComputeBetaBackendServiceProtocolEnum(p.GetProtocol()),
		Fingerprint:              dcl.StringOrNil(p.Fingerprint),
		PortName:                 dcl.StringOrNil(p.PortName),
		EnableCdn:                dcl.Bool(p.EnableCdn),
		SessionAffinity:          ProtoToComputeBetaBackendServiceSessionAffinityEnum(p.GetSessionAffinity()),
		AffinityCookieTtlSec:     dcl.Int64OrNil(p.AffinityCookieTtlSec),
		Location:                 dcl.StringOrNil(p.Location),
		FailoverPolicy:           ProtoToComputeBetaBackendServiceFailoverPolicy(p.GetFailoverPolicy()),
		LoadBalancingScheme:      ProtoToComputeBetaBackendServiceLoadBalancingSchemeEnum(p.GetLoadBalancingScheme()),
		ConnectionDraining:       ProtoToComputeBetaBackendServiceConnectionDraining(p.GetConnectionDraining()),
		Iap:                      ProtoToComputeBetaBackendServiceIap(p.GetIap()),
		CdnPolicy:                ProtoToComputeBetaBackendServiceCdnPolicy(p.GetCdnPolicy()),
		SecurityPolicy:           dcl.StringOrNil(p.SecurityPolicy),
		LogConfig:                ProtoToComputeBetaBackendServiceLogConfig(p.GetLogConfig()),
		SecuritySettings:         ProtoToComputeBetaBackendServiceSecuritySettings(p.GetSecuritySettings()),
		LocalityLbPolicy:         ProtoToComputeBetaBackendServiceLocalityLbPolicyEnum(p.GetLocalityLbPolicy()),
		ConsistentHash:           ProtoToComputeBetaBackendServiceConsistentHash(p.GetConsistentHash()),
		CircuitBreakers:          ProtoToComputeBetaBackendServiceCircuitBreakers(p.GetCircuitBreakers()),
		OutlierDetection:         ProtoToComputeBetaBackendServiceOutlierDetection(p.GetOutlierDetection()),
		Network:                  dcl.StringOrNil(p.Network),
		Subsetting:               ProtoToComputeBetaBackendServiceSubsetting(p.GetSubsetting()),
		ConnectionTrackingPolicy: ProtoToComputeBetaBackendServiceConnectionTrackingPolicy(p.GetConnectionTrackingPolicy()),
		MaxStreamDuration:        ProtoToComputeBetaBackendServiceMaxStreamDuration(p.GetMaxStreamDuration()),
		Project:                  dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToComputeBetaBackendServiceBackends(r))
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
func ComputeBetaBackendServiceBackendsBalancingModeEnumToProto(e *beta.BackendServiceBackendsBalancingModeEnum) betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum_value["BackendServiceBackendsBalancingModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum(v)
	}
	return betapb.ComputeBetaBackendServiceBackendsBalancingModeEnum(0)
}

// BackendServiceProtocolEnumToProto converts a BackendServiceProtocolEnum enum to its proto representation.
func ComputeBetaBackendServiceProtocolEnumToProto(e *beta.BackendServiceProtocolEnum) betapb.ComputeBetaBackendServiceProtocolEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceProtocolEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceProtocolEnum_value["BackendServiceProtocolEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceProtocolEnum(v)
	}
	return betapb.ComputeBetaBackendServiceProtocolEnum(0)
}

// BackendServiceSessionAffinityEnumToProto converts a BackendServiceSessionAffinityEnum enum to its proto representation.
func ComputeBetaBackendServiceSessionAffinityEnumToProto(e *beta.BackendServiceSessionAffinityEnum) betapb.ComputeBetaBackendServiceSessionAffinityEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceSessionAffinityEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceSessionAffinityEnum_value["BackendServiceSessionAffinityEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceSessionAffinityEnum(v)
	}
	return betapb.ComputeBetaBackendServiceSessionAffinityEnum(0)
}

// BackendServiceLoadBalancingSchemeEnumToProto converts a BackendServiceLoadBalancingSchemeEnum enum to its proto representation.
func ComputeBetaBackendServiceLoadBalancingSchemeEnumToProto(e *beta.BackendServiceLoadBalancingSchemeEnum) betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum_value["BackendServiceLoadBalancingSchemeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum(v)
	}
	return betapb.ComputeBetaBackendServiceLoadBalancingSchemeEnum(0)
}

// BackendServiceCdnPolicyCacheModeEnumToProto converts a BackendServiceCdnPolicyCacheModeEnum enum to its proto representation.
func ComputeBetaBackendServiceCdnPolicyCacheModeEnumToProto(e *beta.BackendServiceCdnPolicyCacheModeEnum) betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum_value["BackendServiceCdnPolicyCacheModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum(v)
	}
	return betapb.ComputeBetaBackendServiceCdnPolicyCacheModeEnum(0)
}

// BackendServiceLocalityLbPolicyEnumToProto converts a BackendServiceLocalityLbPolicyEnum enum to its proto representation.
func ComputeBetaBackendServiceLocalityLbPolicyEnumToProto(e *beta.BackendServiceLocalityLbPolicyEnum) betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum_value["BackendServiceLocalityLbPolicyEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum(v)
	}
	return betapb.ComputeBetaBackendServiceLocalityLbPolicyEnum(0)
}

// BackendServiceSubsettingPolicyEnumToProto converts a BackendServiceSubsettingPolicyEnum enum to its proto representation.
func ComputeBetaBackendServiceSubsettingPolicyEnumToProto(e *beta.BackendServiceSubsettingPolicyEnum) betapb.ComputeBetaBackendServiceSubsettingPolicyEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceSubsettingPolicyEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceSubsettingPolicyEnum_value["BackendServiceSubsettingPolicyEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceSubsettingPolicyEnum(v)
	}
	return betapb.ComputeBetaBackendServiceSubsettingPolicyEnum(0)
}

// BackendServiceConnectionTrackingPolicyTrackingModeEnumToProto converts a BackendServiceConnectionTrackingPolicyTrackingModeEnum enum to its proto representation.
func ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnumToProto(e *beta.BackendServiceConnectionTrackingPolicyTrackingModeEnum) betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum_value["BackendServiceConnectionTrackingPolicyTrackingModeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum(v)
	}
	return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnum(0)
}

// BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnumToProto converts a BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum enum to its proto representation.
func ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnumToProto(e *beta.BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum) betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum {
	if e == nil {
		return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(0)
	}
	if v, ok := betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum_value["BackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum"+string(*e)]; ok {
		return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(v)
	}
	return betapb.ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnum(0)
}

// BackendServiceBackendsToProto converts a BackendServiceBackends resource to its proto representation.
func ComputeBetaBackendServiceBackendsToProto(o *beta.BackendServiceBackends) *betapb.ComputeBetaBackendServiceBackends {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceBackends{
		Description:               dcl.ValueOrEmptyString(o.Description),
		Group:                     dcl.ValueOrEmptyString(o.Group),
		BalancingMode:             ComputeBetaBackendServiceBackendsBalancingModeEnumToProto(o.BalancingMode),
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
func ComputeBetaBackendServiceFailoverPolicyToProto(o *beta.BackendServiceFailoverPolicy) *betapb.ComputeBetaBackendServiceFailoverPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceFailoverPolicy{
		DisableConnectionDrainOnFailover: dcl.ValueOrEmptyBool(o.DisableConnectionDrainOnFailover),
		DropTrafficIfUnhealthy:           dcl.ValueOrEmptyBool(o.DropTrafficIfUnhealthy),
		FailoverRatio:                    dcl.ValueOrEmptyDouble(o.FailoverRatio),
	}
	return p
}

// BackendServiceConnectionDrainingToProto converts a BackendServiceConnectionDraining resource to its proto representation.
func ComputeBetaBackendServiceConnectionDrainingToProto(o *beta.BackendServiceConnectionDraining) *betapb.ComputeBetaBackendServiceConnectionDraining {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceConnectionDraining{
		DrainingTimeoutSec: dcl.ValueOrEmptyInt64(o.DrainingTimeoutSec),
	}
	return p
}

// BackendServiceIapToProto converts a BackendServiceIap resource to its proto representation.
func ComputeBetaBackendServiceIapToProto(o *beta.BackendServiceIap) *betapb.ComputeBetaBackendServiceIap {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceIap{
		Enabled:                  dcl.ValueOrEmptyBool(o.Enabled),
		Oauth2ClientId:           dcl.ValueOrEmptyString(o.OAuth2ClientId),
		Oauth2ClientSecret:       dcl.ValueOrEmptyString(o.OAuth2ClientSecret),
		Oauth2ClientSecretSha256: dcl.ValueOrEmptyString(o.OAuth2ClientSecretSha256),
	}
	return p
}

// BackendServiceCdnPolicyToProto converts a BackendServiceCdnPolicy resource to its proto representation.
func ComputeBetaBackendServiceCdnPolicyToProto(o *beta.BackendServiceCdnPolicy) *betapb.ComputeBetaBackendServiceCdnPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCdnPolicy{
		CacheKeyPolicy:          ComputeBetaBackendServiceCdnPolicyCacheKeyPolicyToProto(o.CacheKeyPolicy),
		SignedUrlCacheMaxAgeSec: dcl.ValueOrEmptyInt64(o.SignedUrlCacheMaxAgeSec),
		RequestCoalescing:       dcl.ValueOrEmptyBool(o.RequestCoalescing),
		CacheMode:               ComputeBetaBackendServiceCdnPolicyCacheModeEnumToProto(o.CacheMode),
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
		p.NegativeCachingPolicy = append(p.NegativeCachingPolicy, ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicyToProto(&r))
	}
	for _, r := range o.BypassCacheOnRequestHeaders {
		p.BypassCacheOnRequestHeaders = append(p.BypassCacheOnRequestHeaders, ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto(&r))
	}
	return p
}

// BackendServiceCdnPolicyCacheKeyPolicyToProto converts a BackendServiceCdnPolicyCacheKeyPolicy resource to its proto representation.
func ComputeBetaBackendServiceCdnPolicyCacheKeyPolicyToProto(o *beta.BackendServiceCdnPolicyCacheKeyPolicy) *betapb.ComputeBetaBackendServiceCdnPolicyCacheKeyPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCdnPolicyCacheKeyPolicy{
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
func ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicyToProto(o *beta.BackendServiceCdnPolicyNegativeCachingPolicy) *betapb.ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCdnPolicyNegativeCachingPolicy{
		Code: dcl.ValueOrEmptyInt64(o.Code),
		Ttl:  dcl.ValueOrEmptyInt64(o.Ttl),
	}
	return p
}

// BackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto converts a BackendServiceCdnPolicyBypassCacheOnRequestHeaders resource to its proto representation.
func ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeadersToProto(o *beta.BackendServiceCdnPolicyBypassCacheOnRequestHeaders) *betapb.ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCdnPolicyBypassCacheOnRequestHeaders{
		HeaderName: dcl.ValueOrEmptyString(o.HeaderName),
	}
	return p
}

// BackendServiceLogConfigToProto converts a BackendServiceLogConfig resource to its proto representation.
func ComputeBetaBackendServiceLogConfigToProto(o *beta.BackendServiceLogConfig) *betapb.ComputeBetaBackendServiceLogConfig {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceLogConfig{
		Enable:     dcl.ValueOrEmptyBool(o.Enable),
		SampleRate: dcl.ValueOrEmptyDouble(o.SampleRate),
	}
	return p
}

// BackendServiceSecuritySettingsToProto converts a BackendServiceSecuritySettings resource to its proto representation.
func ComputeBetaBackendServiceSecuritySettingsToProto(o *beta.BackendServiceSecuritySettings) *betapb.ComputeBetaBackendServiceSecuritySettings {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceSecuritySettings{
		ClientTlsPolicy: dcl.ValueOrEmptyString(o.ClientTlsPolicy),
		Authentication:  dcl.ValueOrEmptyString(o.Authentication),
	}
	for _, r := range o.SubjectAltNames {
		p.SubjectAltNames = append(p.SubjectAltNames, r)
	}
	return p
}

// BackendServiceConsistentHashToProto converts a BackendServiceConsistentHash resource to its proto representation.
func ComputeBetaBackendServiceConsistentHashToProto(o *beta.BackendServiceConsistentHash) *betapb.ComputeBetaBackendServiceConsistentHash {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceConsistentHash{
		HttpCookie:      ComputeBetaBackendServiceConsistentHashHttpCookieToProto(o.HttpCookie),
		HttpHeaderName:  dcl.ValueOrEmptyString(o.HttpHeaderName),
		MinimumRingSize: dcl.ValueOrEmptyInt64(o.MinimumRingSize),
	}
	return p
}

// BackendServiceConsistentHashHttpCookieToProto converts a BackendServiceConsistentHashHttpCookie resource to its proto representation.
func ComputeBetaBackendServiceConsistentHashHttpCookieToProto(o *beta.BackendServiceConsistentHashHttpCookie) *betapb.ComputeBetaBackendServiceConsistentHashHttpCookie {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceConsistentHashHttpCookie{
		Name: dcl.ValueOrEmptyString(o.Name),
		Path: dcl.ValueOrEmptyString(o.Path),
		Ttl:  ComputeBetaBackendServiceConsistentHashHttpCookieTtlToProto(o.Ttl),
	}
	return p
}

// BackendServiceConsistentHashHttpCookieTtlToProto converts a BackendServiceConsistentHashHttpCookieTtl resource to its proto representation.
func ComputeBetaBackendServiceConsistentHashHttpCookieTtlToProto(o *beta.BackendServiceConsistentHashHttpCookieTtl) *betapb.ComputeBetaBackendServiceConsistentHashHttpCookieTtl {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceConsistentHashHttpCookieTtl{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceCircuitBreakersToProto converts a BackendServiceCircuitBreakers resource to its proto representation.
func ComputeBetaBackendServiceCircuitBreakersToProto(o *beta.BackendServiceCircuitBreakers) *betapb.ComputeBetaBackendServiceCircuitBreakers {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCircuitBreakers{
		ConnectTimeout:           ComputeBetaBackendServiceCircuitBreakersConnectTimeoutToProto(o.ConnectTimeout),
		MaxRequestsPerConnection: dcl.ValueOrEmptyInt64(o.MaxRequestsPerConnection),
		MaxConnections:           dcl.ValueOrEmptyInt64(o.MaxConnections),
		MaxPendingRequests:       dcl.ValueOrEmptyInt64(o.MaxPendingRequests),
		MaxRequests:              dcl.ValueOrEmptyInt64(o.MaxRequests),
		MaxRetries:               dcl.ValueOrEmptyInt64(o.MaxRetries),
	}
	return p
}

// BackendServiceCircuitBreakersConnectTimeoutToProto converts a BackendServiceCircuitBreakersConnectTimeout resource to its proto representation.
func ComputeBetaBackendServiceCircuitBreakersConnectTimeoutToProto(o *beta.BackendServiceCircuitBreakersConnectTimeout) *betapb.ComputeBetaBackendServiceCircuitBreakersConnectTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceCircuitBreakersConnectTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceOutlierDetectionToProto converts a BackendServiceOutlierDetection resource to its proto representation.
func ComputeBetaBackendServiceOutlierDetectionToProto(o *beta.BackendServiceOutlierDetection) *betapb.ComputeBetaBackendServiceOutlierDetection {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceOutlierDetection{
		ConsecutiveErrors:                  dcl.ValueOrEmptyInt64(o.ConsecutiveErrors),
		Interval:                           ComputeBetaBackendServiceOutlierDetectionIntervalToProto(o.Interval),
		BaseEjectionTime:                   ComputeBetaBackendServiceOutlierDetectionBaseEjectionTimeToProto(o.BaseEjectionTime),
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
func ComputeBetaBackendServiceOutlierDetectionIntervalToProto(o *beta.BackendServiceOutlierDetectionInterval) *betapb.ComputeBetaBackendServiceOutlierDetectionInterval {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceOutlierDetectionInterval{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceOutlierDetectionBaseEjectionTimeToProto converts a BackendServiceOutlierDetectionBaseEjectionTime resource to its proto representation.
func ComputeBetaBackendServiceOutlierDetectionBaseEjectionTimeToProto(o *beta.BackendServiceOutlierDetectionBaseEjectionTime) *betapb.ComputeBetaBackendServiceOutlierDetectionBaseEjectionTime {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceOutlierDetectionBaseEjectionTime{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceSubsettingToProto converts a BackendServiceSubsetting resource to its proto representation.
func ComputeBetaBackendServiceSubsettingToProto(o *beta.BackendServiceSubsetting) *betapb.ComputeBetaBackendServiceSubsetting {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceSubsetting{
		Policy: ComputeBetaBackendServiceSubsettingPolicyEnumToProto(o.Policy),
	}
	return p
}

// BackendServiceConnectionTrackingPolicyToProto converts a BackendServiceConnectionTrackingPolicy resource to its proto representation.
func ComputeBetaBackendServiceConnectionTrackingPolicyToProto(o *beta.BackendServiceConnectionTrackingPolicy) *betapb.ComputeBetaBackendServiceConnectionTrackingPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceConnectionTrackingPolicy{
		TrackingMode:                             ComputeBetaBackendServiceConnectionTrackingPolicyTrackingModeEnumToProto(o.TrackingMode),
		ConnectionPersistenceOnUnhealthyBackends: ComputeBetaBackendServiceConnectionTrackingPolicyConnectionPersistenceOnUnhealthyBackendsEnumToProto(o.ConnectionPersistenceOnUnhealthyBackends),
		IdleTimeoutSec:                           dcl.ValueOrEmptyInt64(o.IdleTimeoutSec),
	}
	return p
}

// BackendServiceMaxStreamDurationToProto converts a BackendServiceMaxStreamDuration resource to its proto representation.
func ComputeBetaBackendServiceMaxStreamDurationToProto(o *beta.BackendServiceMaxStreamDuration) *betapb.ComputeBetaBackendServiceMaxStreamDuration {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaBackendServiceMaxStreamDuration{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// BackendServiceToProto converts a BackendService resource to its proto representation.
func BackendServiceToProto(resource *beta.BackendService) *betapb.ComputeBetaBackendService {
	p := &betapb.ComputeBetaBackendService{
		Name:                     dcl.ValueOrEmptyString(resource.Name),
		Description:              dcl.ValueOrEmptyString(resource.Description),
		SelfLink:                 dcl.ValueOrEmptyString(resource.SelfLink),
		SelfLinkWithId:           dcl.ValueOrEmptyString(resource.SelfLinkWithId),
		TimeoutSec:               dcl.ValueOrEmptyInt64(resource.TimeoutSec),
		Port:                     dcl.ValueOrEmptyInt64(resource.Port),
		Protocol:                 ComputeBetaBackendServiceProtocolEnumToProto(resource.Protocol),
		Fingerprint:              dcl.ValueOrEmptyString(resource.Fingerprint),
		PortName:                 dcl.ValueOrEmptyString(resource.PortName),
		EnableCdn:                dcl.ValueOrEmptyBool(resource.EnableCdn),
		SessionAffinity:          ComputeBetaBackendServiceSessionAffinityEnumToProto(resource.SessionAffinity),
		AffinityCookieTtlSec:     dcl.ValueOrEmptyInt64(resource.AffinityCookieTtlSec),
		Location:                 dcl.ValueOrEmptyString(resource.Location),
		FailoverPolicy:           ComputeBetaBackendServiceFailoverPolicyToProto(resource.FailoverPolicy),
		LoadBalancingScheme:      ComputeBetaBackendServiceLoadBalancingSchemeEnumToProto(resource.LoadBalancingScheme),
		ConnectionDraining:       ComputeBetaBackendServiceConnectionDrainingToProto(resource.ConnectionDraining),
		Iap:                      ComputeBetaBackendServiceIapToProto(resource.Iap),
		CdnPolicy:                ComputeBetaBackendServiceCdnPolicyToProto(resource.CdnPolicy),
		SecurityPolicy:           dcl.ValueOrEmptyString(resource.SecurityPolicy),
		LogConfig:                ComputeBetaBackendServiceLogConfigToProto(resource.LogConfig),
		SecuritySettings:         ComputeBetaBackendServiceSecuritySettingsToProto(resource.SecuritySettings),
		LocalityLbPolicy:         ComputeBetaBackendServiceLocalityLbPolicyEnumToProto(resource.LocalityLbPolicy),
		ConsistentHash:           ComputeBetaBackendServiceConsistentHashToProto(resource.ConsistentHash),
		CircuitBreakers:          ComputeBetaBackendServiceCircuitBreakersToProto(resource.CircuitBreakers),
		OutlierDetection:         ComputeBetaBackendServiceOutlierDetectionToProto(resource.OutlierDetection),
		Network:                  dcl.ValueOrEmptyString(resource.Network),
		Subsetting:               ComputeBetaBackendServiceSubsettingToProto(resource.Subsetting),
		ConnectionTrackingPolicy: ComputeBetaBackendServiceConnectionTrackingPolicyToProto(resource.ConnectionTrackingPolicy),
		MaxStreamDuration:        ComputeBetaBackendServiceMaxStreamDurationToProto(resource.MaxStreamDuration),
		Project:                  dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Backends {
		p.Backends = append(p.Backends, ComputeBetaBackendServiceBackendsToProto(&r))
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
func (s *BackendServiceServer) applyBackendService(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaBackendServiceRequest) (*betapb.ComputeBetaBackendService, error) {
	p := ProtoToBackendService(request.GetResource())
	res, err := c.ApplyBackendService(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackendServiceToProto(res)
	return r, nil
}

// ApplyBackendService handles the gRPC request by passing it to the underlying BackendService Apply() method.
func (s *BackendServiceServer) ApplyComputeBetaBackendService(ctx context.Context, request *betapb.ApplyComputeBetaBackendServiceRequest) (*betapb.ComputeBetaBackendService, error) {
	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBackendService(ctx, cl, request)
}

// DeleteBackendService handles the gRPC request by passing it to the underlying BackendService Delete() method.
func (s *BackendServiceServer) DeleteComputeBetaBackendService(ctx context.Context, request *betapb.DeleteComputeBetaBackendServiceRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackendService(ctx, ProtoToBackendService(request.GetResource()))

}

// ListComputeBetaBackendService handles the gRPC request by passing it to the underlying BackendServiceList() method.
func (s *BackendServiceServer) ListComputeBetaBackendService(ctx context.Context, request *betapb.ListComputeBetaBackendServiceRequest) (*betapb.ListComputeBetaBackendServiceResponse, error) {
	cl, err := createConfigBackendService(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackendService(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaBackendService
	for _, r := range resources.Items {
		rp := BackendServiceToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaBackendServiceResponse{Items: protos}, nil
}

func createConfigBackendService(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
