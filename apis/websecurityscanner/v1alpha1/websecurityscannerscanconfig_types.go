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

// API sources for WebSecurityScannerScanConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/websecurityscanner/v1/scan_config.proto
// +kcc:source:service-docs=https://cloud.google.com/security-command-center/docs/concepts-web-security-scanner-overview/
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/security-command-center/docs/concepts-web-security-scanner-overview/reference/rest

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var WebSecurityScannerScanConfigGVK = GroupVersion.WithKind("WebSecurityScannerScanConfig")

// WebSecurityScannerScanConfigSpec defines the desired state of WebSecurityScannerScanConfig
// +kcc:spec:proto=google.cloud.websecurityscanner.v1.ScanConfig
type WebSecurityScannerScanConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The WebSecurityScannerScanConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The user provided display name of the ScanConfig.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The maximum QPS during scanning. A valid value ranges from 5 to 20
	//  inclusively. If the field is unspecified or its value is set 0, server will
	//  default to 15. Other values outside of [5, 20] range will be rejected with
	//  INVALID_ARGUMENT error.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.max_qps
	MaxQps *int32 `json:"maxQps,omitempty"`

	// Required. The starting URLs from which the scanner finds site pages.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.starting_urls
	StartingURLs []string `json:"startingURLs,omitempty"`

	// The authentication configuration. If specified, service will use the
	//  authentication configuration during scanning.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.authentication
	Authentication *ScanConfig_Authentication `json:"authentication,omitempty"`

	// The user agent used during scanning.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.user_agent
	UserAgent *string `json:"userAgent,omitempty"`

	// The excluded URL patterns as described in
	//  https://cloud.google.com/security-command-center/docs/how-to-use-web-security-scanner#excluding_urls
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.blacklist_patterns
	BlacklistPatterns []string `json:"blacklistPatterns,omitempty"`

	// The schedule of the ScanConfig.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.schedule
	Schedule *ScanConfig_Schedule `json:"schedule,omitempty"`

	// Controls export of scan configurations and results to Security
	//  Command Center.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.export_to_security_command_center
	ExportToSecurityCommandCenter *string `json:"exportToSecurityCommandCenter,omitempty"`

	// The risk level selected for the scan
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.risk_level
	RiskLevel *string `json:"riskLevel,omitempty"`

	// Whether the scan config is managed by Web Security Scanner, output
	//  only.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.managed_scan
	ManagedScan *bool `json:"managedScan,omitempty"`

	// Whether the scan configuration has enabled static IP address scan feature.
	//  If enabled, the scanner will access applications from static IP addresses.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.static_ip_scan
	StaticIPScan *bool `json:"staticIPScan,omitempty"`

	// Whether to keep scanning even if most requests return HTTP error codes.
	// +kcc:proto:field=google.cloud.websecurityscanner.v1.ScanConfig.ignore_http_status_errors
	IgnoreHTTPStatusErrors *bool `json:"ignoreHTTPStatusErrors,omitempty"`
}

// WebSecurityScannerScanConfigStatus defines the config connector machine state of WebSecurityScannerScanConfig
type WebSecurityScannerScanConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the WebSecurityScannerScanConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *WebSecurityScannerScanConfigObservedState `json:"observedState,omitempty"`
}

// WebSecurityScannerScanConfigObservedState is the state of the WebSecurityScannerScanConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.websecurityscanner.v1.ScanConfig
type WebSecurityScannerScanConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpwebsecurityscannerscanconfig;gcpwebsecurityscannerscanconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WebSecurityScannerScanConfig is the Schema for the WebSecurityScannerScanConfig API
// +k8s:openapi-gen=true
type WebSecurityScannerScanConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   WebSecurityScannerScanConfigSpec   `json:"spec,omitempty"`
	Status WebSecurityScannerScanConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// WebSecurityScannerScanConfigList contains a list of WebSecurityScannerScanConfig
type WebSecurityScannerScanConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebSecurityScannerScanConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WebSecurityScannerScanConfig{}, &WebSecurityScannerScanConfigList{})
}
