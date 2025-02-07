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

package v1alpha1


// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap
type CertificateMap struct {
	// A user-defined name of the Certificate Map. Certificate Map names must be
	//  unique globally and match pattern
	//  `projects/*/locations/*/certificateMaps/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.name
	Name *string `json:"name,omitempty"`

	// One or more paragraphs of text description of a certificate map.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.description
	Description *string `json:"description,omitempty"`

	// Set of labels associated with a Certificate Map.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget
type CertificateMap_GclbTarget struct {
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig
type CertificateMap_GclbTarget_IpConfig struct {
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap
type CertificateMapObservedState struct {
	// Output only. The creation timestamp of a Certificate Map.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp of a Certificate Map.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of GCLB targets that use this Certificate Map.
	//  A Target Proxy is only present on this list if it's attached to a
	//  Forwarding Rule.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.gclb_targets
	GclbTargets []CertificateMap_GclbTarget `json:"gclbTargets,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget
type CertificateMap_GclbTargetObservedState struct {
	// Output only. This field returns the resource name in the following
	//  format:
	//  `//compute.googleapis.com/projects/*/global/targetHttpsProxies/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.target_https_proxy
	TargetHTTPSProxy *string `json:"targetHTTPSProxy,omitempty"`

	// Output only. This field returns the resource name in the following
	//  format:
	//  `//compute.googleapis.com/projects/*/global/targetSslProxies/*`.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.target_ssl_proxy
	TargetSslProxy *string `json:"targetSslProxy,omitempty"`

	// Output only. IP configurations for this Target Proxy where the
	//  Certificate Map is serving.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.ip_configs
	IPConfigs []CertificateMap_GclbTarget_IpConfig `json:"ipConfigs,omitempty"`
}

// +kcc:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig
type CertificateMap_GclbTarget_IpConfigObservedState struct {
	// Output only. An external IP address.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. Ports.
	// +kcc:proto:field=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig.ports
	Ports []uint32 `json:"ports,omitempty"`
}
