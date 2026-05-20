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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeHealthCheckGVK = GroupVersion.WithKind("ComputeHealthCheck")

// ComputeHealthCheckSpec defines the desired state of ComputeHealthCheck
// +kcc:spec:proto=google.cloud.compute.v1.HealthCheck
type ComputeHealthCheckSpec struct {
	// Location represents the geographical location of the ComputeHealthCheck. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// How often (in seconds) to send a health check. The default value is 5 seconds.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.check_interval_sec
	CheckIntervalSec *int64 `json:"checkIntervalSec,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.grpc_health_check
	GRPCHealthCheck *HealthCheckGRPCHealthCheck `json:"grpcHealthCheck,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.healthy_threshold
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.http2_health_check
	HTTP2HealthCheck *HealthCheckHTTP2HealthCheck `json:"http2HealthCheck,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.http_health_check
	HTTPHealthCheck *HealthCheckHTTPHealthCheck `json:"httpHealthCheck,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.https_health_check
	HTTPSHealthCheck *HealthCheckHTTPSHealthCheck `json:"httpsHealthCheck,omitempty"`

	// Configure logging on this health check.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.log_config
	LogConfig *HealthCheckLogConfig `json:"logConfig,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.ssl_health_check
	SSLHealthCheck *HealthCheckSSLHealthCheck `json:"sslHealthCheck,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.tcp_health_check
	TCPHealthCheck *HealthCheckTCPHealthCheck `json:"tcpHealthCheck,omitempty"`

	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.timeout_sec
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheck.unhealthy_threshold
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.GRPCHealthCheck
type HealthCheckGRPCHealthCheck struct {
	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	//   - Empty serviceName means the overall status of all services at the backend.
	//   - Non-empty serviceName means the health of that gRPC service, as defined by the owner of the service.
	// The grpcServiceName can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.GRPCHealthCheck.grpc_service_name
	GRPCServiceName *string `json:"grpcServiceName,omitempty"`

	// The port number for the health check request.
	// Must be specified if portName and portSpecification are not set
	// or if port_specification is USE_FIXED_PORT. Valid values are 1 through 65535.
	// +kcc:proto:field=google.cloud.compute.v1.GRPCHealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.GRPCHealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, gRPC health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.GRPCHealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HTTP2HealthCheck
type HealthCheckHTTP2HealthCheck struct {
	// The value of the host header in the HTTP2 health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.host
	Host *string `json:"host,omitempty"`

	// The TCP port number for the HTTP2 health check request.
	// The default value is 443.
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, HTTP2 health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The request path of the HTTP2 health check request.
	// The default value is /.
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.request_path
	RequestPath *string `json:"requestPath,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.HTTP2HealthCheck.response
	Response *string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HTTPHealthCheck
type HealthCheckHTTPHealthCheck struct {
	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.host
	Host *string `json:"host,omitempty"`

	// The TCP port number for the HTTP health check request.
	// The default value is 80.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, HTTP health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The request path of the HTTP health check request.
	// The default value is /.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.request_path
	RequestPath *string `json:"requestPath,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPHealthCheck.response
	Response *string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HTTPSHealthCheck
type HealthCheckHTTPSHealthCheck struct {
	// The value of the host header in the HTTPS health check request.
	// If left empty (default value), the public IP on behalf of which this health
	// check is performed will be used.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.host
	Host *string `json:"host,omitempty"`

	// The TCP port number for the HTTPS health check request.
	// The default value is 443.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, HTTPS health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The request path of the HTTPS health check request.
	// The default value is /.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.request_path
	RequestPath *string `json:"requestPath,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.HTTPSHealthCheck.response
	Response *string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HealthCheckLogConfig
type HealthCheckLogConfig struct {
	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	// +kcc:proto:field=google.cloud.compute.v1.HealthCheckLogConfig.enable
	Enable *bool `json:"enable,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SSLHealthCheck
type HealthCheckSSLHealthCheck struct {
	// The TCP port number for the SSL health check request.
	// The default value is 443.
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, SSL health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The application data to send once the SSL connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.request
	Request *string `json:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.SSLHealthCheck.response
	Response *string `json:"response,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.TCPHealthCheck
type HealthCheckTCPHealthCheck struct {
	// The TCP port number for the TCP health check request.
	// The default value is 443.
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.port
	Port *int64 `json:"port,omitempty"`

	// Port name as defined in InstanceGroup#NamedPort#name. If both port and
	// port_name are defined, port takes precedence.
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.port_name
	PortName *string `json:"portName,omitempty"`

	// Specifies how port is selected for health checking, can be one of the
	// following values:
	//   * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	// If not specified, TCP health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"].
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.port_specification
	PortSpecification *string `json:"portSpecification,omitempty"`

	// Specifies the type of proxy header to append before sending data to the
	// backend. Default value: "NONE" Possible values: ["NONE", "PROXY_V1"].
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// The application data to send once the TCP connection has been
	// established (default value is empty). If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.request
	Request *string `json:"request,omitempty"`

	// The bytes to match against the beginning of the response data. If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	// +kcc:proto:field=google.cloud.compute.v1.TCPHealthCheck.response
	Response *string `json:"response,omitempty"`
}

// ComputeHealthCheckStatus defines the config connector machine state of ComputeHealthCheck
type ComputeHealthCheckStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeHealthCheck's current state. */
	Conditions []commonv1alpha1.Condition `json:"conditions,omitempty"`

	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// +optional
	ObservedState *ComputeHealthCheckObservedState `json:"observedState,omitempty"`
}

// ComputeHealthCheckObservedState is the state of the ComputeHealthCheck resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.HealthCheck
type ComputeHealthCheckObservedState struct {
	/* The type of the health check. One of HTTP, HTTPS, TCP, or SSL. */
	// +optional
	Type *string `json:"type,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputehealthcheck;gcpcomputehealthchecks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeHealthCheck is the Schema for the ComputeHealthCheck API
// +k8s:openapi-gen=true
type ComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status ComputeHealthCheckStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeHealthCheckList contains a list of ComputeHealthCheck
type ComputeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeHealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeHealthCheck{}, &ComputeHealthCheckList{})
}
