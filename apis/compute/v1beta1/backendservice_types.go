// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackendServiceBackend struct {
	// +optional
	BalancingMode *string `json:"balancingMode,omitempty"`

	// +optional
	CapacityScaler *float64 `json:"capacityScaler,omitempty"`

	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	Failover *bool `json:"failover,omitempty"`

	Group BackendServiceGroup `json:"group"`

	// +optional
	MaxConnections *int64 `json:"maxConnections,omitempty"`

	// +optional
	MaxConnectionsPerEndpoint *int64 `json:"maxConnectionsPerEndpoint,omitempty"`

	// +optional
	MaxConnectionsPerInstance *int64 `json:"maxConnectionsPerInstance,omitempty"`

	// +optional
	MaxRate *int64 `json:"maxRate,omitempty"`

	// +optional
	MaxRatePerEndpoint *float64 `json:"maxRatePerEndpoint,omitempty"`

	// +optional
	MaxRatePerInstance *float64 `json:"maxRatePerInstance,omitempty"`

	// +optional
	MaxUtilization *float64 `json:"maxUtilization,omitempty"`
}

type BackendServiceGroup struct {
	// +optional
	InstanceGroupRef *v1alpha1.ResourceRef `json:"instanceGroupRef,omitempty"`

	// +optional
	NetworkEndpointGroupRef *v1alpha1.ResourceRef `json:"networkEndpointGroupRef,omitempty"`
}

type BackendServiceBaseEjectionTime struct {
	// +optional
	Nanos *int64 `json:"nanos,omitempty"`

	Seconds int64 `json:"seconds"`
}

type BackendServiceBypassCacheOnRequestHeaders struct {
	HeaderName string `json:"headerName"`
}

type BackendServiceCacheKeyPolicy struct {
	// +optional
	IncludeHost *bool `json:"includeHost,omitempty"`

	// +optional
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty"`

	// +optional
	IncludeNamedCookies []string `json:"includeNamedCookies,omitempty"`

	// +optional
	IncludeProtocol *bool `json:"includeProtocol,omitempty"`

	// +optional
	IncludeQueryString *bool `json:"includeQueryString,omitempty"`

	// +optional
	QueryStringBlacklist []string `json:"queryStringBlacklist,omitempty"`

	// +optional
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty"`
}

type BackendServiceCdnPolicy struct {
	// +optional
	BypassCacheOnRequestHeaders []BackendServiceBypassCacheOnRequestHeaders `json:"bypassCacheOnRequestHeaders,omitempty"`

	// +optional
	CacheKeyPolicy *BackendServiceCacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	// +optional
	CacheMode *string `json:"cacheMode,omitempty"`

	// +optional
	ClientTtl *int64 `json:"clientTtl,omitempty"`

	// +optional
	DefaultTtl *int64 `json:"defaultTtl,omitempty"`

	// +optional
	MaxTtl *int64 `json:"maxTtl,omitempty"`

	// +optional
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	// +optional
	NegativeCachingPolicy []BackendServiceNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`

	// +optional
	ServeWhileStale *int64 `json:"serveWhileStale,omitempty"`

	// +optional
	SignedUrlCacheMaxAgeSec *int64 `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

type BackendServiceConnectTimeout struct {
	// +optional
	Nanos *int64 `json:"nanos,omitempty"`

	Seconds int64 `json:"seconds"`
}

type BackendServiceCircuitBreakers struct {
	// +optional
	ConnectTimeout *BackendServiceConnectTimeout `json:"connectTimeout,omitempty"`

	// +optional
	MaxConnections *int64 `json:"maxConnections,omitempty"`

	// +optional
	MaxPendingRequests *int64 `json:"maxPendingRequests,omitempty"`

	// +optional
	MaxRequests *int64 `json:"maxRequests,omitempty"`

	// +optional
	MaxRequestsPerConnection *int64 `json:"maxRequestsPerConnection,omitempty"`

	// +optional
	MaxRetries *int64 `json:"maxRetries,omitempty"`
}

type BackendServiceConnectionTrackingPolicy struct {
	// +optional
	ConnectionPersistenceOnUnhealthyBackends *string `json:"connectionPersistenceOnUnhealthyBackends,omitempty"`

	// +optional
	EnableStrongAffinity *bool `json:"enableStrongAffinity,omitempty"`

	// +optional
	IdleTimeoutSec *int64 `json:"idleTimeoutSec,omitempty"`

	// +optional
	TrackingMode *string `json:"trackingMode,omitempty"`
}

type BackendServiceConsistentHash struct {
	// +optional
	HttpCookie *BackendServiceHttpCookie `json:"httpCookie,omitempty"`

	// +optional
	HttpHeaderName *string `json:"httpHeaderName,omitempty"`

	// +optional
	MinimumRingSize *int64 `json:"minimumRingSize,omitempty"`
}

type BackendServiceHttpCookie struct {
	// +optional
	Name *string `json:"name,omitempty"`

	// +optional
	Path *string `json:"path,omitempty"`

	// +optional
	Ttl *BackendServiceTtl `json:"ttl,omitempty"`
}

type BackendServiceTtl struct {
	// +optional
	Nanos *int64 `json:"nanos,omitempty"`

	Seconds int64 `json:"seconds"`
}

type BackendServiceFailoverPolicy struct {
	// +optional
	DisableConnectionDrainOnFailover *bool `json:"disableConnectionDrainOnFailover,omitempty"`

	// +optional
	DropTrafficIfUnhealthy *bool `json:"dropTrafficIfUnhealthy,omitempty"`

	// +optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty"`
}

type BackendServiceHealthChecks struct {
	// +optional
	HealthCheckRef *v1alpha1.ResourceRef `json:"healthCheckRef,omitempty"`

	// +optional
	HttpHealthCheckRef *v1alpha1.ResourceRef `json:"httpHealthCheckRef,omitempty"`
}

type BackendServiceOauth2ClientSecret struct {
	// +optional
	Value *string `json:"value,omitempty"`

	// +optional
	ValueFrom *BackendServiceValueFrom `json:"valueFrom,omitempty"`
}

type BackendServiceValueFrom struct {
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type BackendServiceIap struct {
	// +optional
	Oauth2ClientId *string `json:"oauth2ClientId,omitempty"`

	// +optional
	Oauth2ClientIdRef *v1alpha1.ResourceRef `json:"oauth2ClientIdRef,omitempty"`

	// +optional
	Oauth2ClientSecret *BackendServiceOauth2ClientSecret `json:"oauth2ClientSecret,omitempty"`

	// +optional
	Oauth2ClientSecretSha256 *string `json:"oauth2ClientSecretSha256,omitempty"`
}

type BackendServiceCustomPolicy struct {
	// +optional
	Data *string `json:"data,omitempty"`

	Name string `json:"name"`
}

type BackendServicePolicy struct {
	Name string `json:"name"`
}

type BackendServiceLocalityLbPolicies struct {
	// +optional
	CustomPolicy *BackendServiceCustomPolicy `json:"customPolicy,omitempty"`

	// +optional
	Policy *BackendServicePolicy `json:"policy,omitempty"`
}

type BackendServiceLogConfig struct {
	// +optional
	Enable *bool `json:"enable,omitempty"`

	// +optional
	SampleRate *float64 `json:"sampleRate,omitempty"`
}

type BackendServiceNegativeCachingPolicy struct {
	// +optional
	Code *int64 `json:"code,omitempty"`

	// +optional
	Ttl *int64 `json:"ttl,omitempty"`
}

type BackendServiceInterval struct {
	// +optional
	Nanos *int64 `json:"nanos,omitempty"`

	Seconds int64 `json:"seconds"`
}

type BackendServiceOutlierDetection struct {
	// +optional
	BaseEjectionTime *BackendServiceBaseEjectionTime `json:"baseEjectionTime,omitempty"`

	// +optional
	ConsecutiveErrors *int64 `json:"consecutiveErrors,omitempty"`

	// +optional
	ConsecutiveGatewayFailure *int64 `json:"consecutiveGatewayFailure,omitempty"`

	// +optional
	EnforcingConsecutiveErrors *int64 `json:"enforcingConsecutiveErrors,omitempty"`

	// +optional
	EnforcingConsecutiveGatewayFailure *int64 `json:"enforcingConsecutiveGatewayFailure,omitempty"`

	// +optional
	EnforcingSuccessRate *int64 `json:"enforcingSuccessRate,omitempty"`

	// +optional
	Interval *BackendServiceInterval `json:"interval,omitempty"`

	// +optional
	MaxEjectionPercent *int64 `json:"maxEjectionPercent,omitempty"`

	// +optional
	SuccessRateMinimumHosts *int64 `json:"successRateMinimumHosts,omitempty"`

	// +optional
	SuccessRateRequestVolume *int64 `json:"successRateRequestVolume,omitempty"`

	// +optional
	SuccessRateStdevFactor *int64 `json:"successRateStdevFactor,omitempty"`
}

type BackendServiceSecuritySettings struct {
	ClientTLSPolicyRef v1alpha1.ResourceRef `json:"clientTLSPolicyRef"`

	SubjectAltNames []string `json:"subjectAltNames"`
}

type BackendServiceSubsetting struct {
	Policy string `json:"policy"`
}

// ComputeBackendServiceSpec defines the desired state of ComputeBackendService
type ComputeBackendServiceSpec struct {
	// +optional
	AffinityCookieTtlSec *int64 `json:"affinityCookieTtlSec,omitempty"`

	// +optional
	Backend []BackendServiceBackend `json:"backend,omitempty"`

	// +optional
	CdnPolicy *BackendServiceCdnPolicy `json:"cdnPolicy,omitempty"`

	// +optional
	CircuitBreakers *BackendServiceCircuitBreakers `json:"circuitBreakers,omitempty"`

	// +optional
	CompressionMode *string `json:"compressionMode,omitempty"`

	// +optional
	ConnectionDrainingTimeoutSec *int64 `json:"connectionDrainingTimeoutSec,omitempty"`

	// +optional
	ConnectionTrackingPolicy *BackendServiceConnectionTrackingPolicy `json:"connectionTrackingPolicy,omitempty"`

	// +optional
	ConsistentHash *BackendServiceConsistentHash `json:"consistentHash,omitempty"`

	// +optional
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty"`

	// +optional
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`

	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	EdgeSecurityPolicyRef *v1alpha1.ResourceRef `json:"edgeSecurityPolicyRef,omitempty"`

	// +optional
	EnableCdn *bool `json:"enableCdn,omitempty"`

	// +optional
	FailoverPolicy *BackendServiceFailoverPolicy `json:"failoverPolicy,omitempty"`

	// +optional
	HealthChecks []BackendServiceHealthChecks `json:"healthChecks,omitempty"`

	// +optional
	Iap *BackendServiceIap `json:"iap,omitempty"`

	// +optional
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// +optional
	LocalityLbPolicies []BackendServiceLocalityLbPolicies `json:"localityLbPolicies,omitempty"`

	// +optional
	LocalityLbPolicy *string `json:"localityLbPolicy,omitempty"`

	Location string `json:"location"`

	// +optional
	LogConfig *BackendServiceLogConfig `json:"logConfig,omitempty"`

	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	// +optional
	OutlierDetection *BackendServiceOutlierDetection `json:"outlierDetection,omitempty"`

	// +optional
	PortName *string `json:"portName,omitempty"`

	// +optional
	Protocol *string `json:"protocol,omitempty"`

	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	SecurityPolicy *string `json:"securityPolicy,omitempty"`

	// +optional
	SecurityPolicyRef *v1alpha1.ResourceRef `json:"securityPolicyRef,omitempty"`

	// +optional
	SecuritySettings *BackendServiceSecuritySettings `json:"securitySettings,omitempty"`

	// +optional
	SessionAffinity *string `json:"sessionAffinity,omitempty"`

	// +optional
	Subsetting *BackendServiceSubsetting `json:"subsetting,omitempty"`

	// +optional
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
}

// ComputeBackendServiceStatus defines the config connector machine state of ComputeBackendService
type ComputeBackendServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []commonv1alpha1.Condition `json:"conditions,omitempty"`

	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// +optional
	Fingerprint *string `json:"fingerprint,omitempty"`

	// +optional
	GeneratedId *int64 `json:"generatedId,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendservice;gcpcomputebackendservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendService is the Schema for the ComputeBackendService API
// +k8s:openapi-gen=true
type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status ComputeBackendServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeBackendServiceList contains a list of ComputeBackendService
type ComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendService{}, &ComputeBackendServiceList{})
}
