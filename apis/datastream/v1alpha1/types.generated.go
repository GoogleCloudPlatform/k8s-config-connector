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

// +generated:types
// krm.group: datastream.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.datastream.v1
// resource: DatastreamPrivateConnection:PrivateConnection
// resource: DatastreamConnectionProfile:ConnectionProfile
// resource: DatastreamRoute:Route

package v1alpha1

// +kcc:proto=google.cloud.datastream.v1.BigQueryProfile
type BigQueryProfile struct {
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfig struct {
	// Input only. PEM-encoded private key associated with the Client Certificate.
	//  If this field is used then the 'client_certificate' and the
	//  'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key
	ClientKey *string `json:"clientKey,omitempty"`

	// Input only. PEM-encoded certificate that will be used by the replica to
	//  authenticate against the source database server. If this field is used
	//  then the 'client_key' and the 'ca_certificate' fields are mandatory.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate
	ClientCertificate *string `json:"clientCertificate,omitempty"`

	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfig struct {
	// Input only. PEM-encoded certificate of the CA that signed the source
	//  database server's certificate.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate
	CACertificate *string `json:"caCertificate,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Route
type Route struct {

	// Labels.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Display name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. Destination address for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_address
	DestinationAddress *string `json:"destinationAddress,omitempty"`

	// Destination port for connection
	// +kcc:proto:field=google.cloud.datastream.v1.Route.destination_port
	DestinationPort *int32 `json:"destinationPort,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.StaticServiceIpConnectivity
type StaticServiceIPConnectivity struct {
}

// +kcc:proto=google.cloud.datastream.v1.MysqlSslConfig
type MysqlSSLConfigObservedState struct {
	// Output only. Indicates whether the client_key field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_key_set
	ClientKeySet *bool `json:"clientKeySet,omitempty"`

	// Output only. Indicates whether the client_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.client_certificate_set
	ClientCertificateSet *bool `json:"clientCertificateSet,omitempty"`

	// Output only. Indicates whether the ca_certificate field is set.
	// +kcc:proto:field=google.cloud.datastream.v1.MysqlSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.OracleSslConfig
type OracleSSLConfigObservedState struct {
	// Output only. Indicates whether the ca_certificate field has been set for
	//  this Connection-Profile.
	// +kcc:proto:field=google.cloud.datastream.v1.OracleSslConfig.ca_certificate_set
	CACertificateSet *bool `json:"caCertificateSet,omitempty"`
}

// +kcc:proto=google.cloud.datastream.v1.Route
type RouteObservedState struct {
	// Output only. The resource's name.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.name
	Name *string `json:"name,omitempty"`

	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update time of the resource.
	// +kcc:proto:field=google.cloud.datastream.v1.Route.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
