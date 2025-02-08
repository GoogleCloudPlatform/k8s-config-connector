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


// +kcc:proto=google.cloud.iot.v1.DeviceRegistry
type DeviceRegistry struct {
	// The identifier of this device registry. For example, `myRegistry`.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.id
	ID *string `json:"id,omitempty"`

	// The resource path name. For example,
	//  `projects/example-project/locations/us-central1/registries/my-registry`.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.name
	Name *string `json:"name,omitempty"`

	// The configuration for notification of telemetry events received from the
	//  device. All telemetry events that were successfully published by the
	//  device and acknowledged by Cloud IoT Core are guaranteed to be
	//  delivered to Cloud Pub/Sub. If multiple configurations match a message,
	//  only the first matching configuration is used. If you try to publish a
	//  device telemetry event using MQTT without specifying a Cloud Pub/Sub topic
	//  for the device's registry, the connection closes automatically. If you try
	//  to do so using an HTTP connection, an error is returned. Up to 10
	//  configurations may be provided.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.event_notification_configs
	EventNotificationConfigs []EventNotificationConfig `json:"eventNotificationConfigs,omitempty"`

	// The configuration for notification of new states received from the device.
	//  State updates are guaranteed to be stored in the state history, but
	//  notifications to Cloud Pub/Sub are not guaranteed. For example, if
	//  permissions are misconfigured or the specified topic doesn't exist, no
	//  notification will be published but the state will still be stored in Cloud
	//  IoT Core.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.state_notification_config
	StateNotificationConfig *StateNotificationConfig `json:"stateNotificationConfig,omitempty"`

	// The MQTT configuration for this device registry.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.mqtt_config
	MqttConfig *MqttConfig `json:"mqttConfig,omitempty"`

	// The DeviceService (HTTP) configuration for this device registry.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.http_config
	HTTPConfig *HttpConfig `json:"httpConfig,omitempty"`

	// **Beta Feature**
	//
	//  The default logging verbosity for activity from devices in this registry.
	//  The verbosity level can be overridden by Device.log_level.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.log_level
	LogLevel *string `json:"logLevel,omitempty"`

	// The credentials used to verify the device credentials. No more than 10
	//  credentials can be bound to a single registry at a time. The verification
	//  process occurs at the time of device creation or update. If this field is
	//  empty, no verification is performed. Otherwise, the credentials of a newly
	//  created device or added credentials of an updated device should be signed
	//  with one of these registry credentials.
	//
	//  Note, however, that existing devices will never be affected by
	//  modifications to this list of credentials: after a device has been
	//  successfully created in a registry, it should be able to connect even if
	//  its registry credentials are revoked, deleted, or modified.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceRegistry.credentials
	Credentials []RegistryCredential `json:"credentials,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.EventNotificationConfig
type EventNotificationConfig struct {
	// If the subfolder name matches this string exactly, this configuration will
	//  be used. The string must not include the leading '/' character. If empty,
	//  all strings are matched. This field is used only for telemetry events;
	//  subfolders are not supported for state changes.
	// +kcc:proto:field=google.cloud.iot.v1.EventNotificationConfig.subfolder_matches
	SubfolderMatches *string `json:"subfolderMatches,omitempty"`

	// A Cloud Pub/Sub topic name. For example,
	//  `projects/myProject/topics/deviceEvents`.
	// +kcc:proto:field=google.cloud.iot.v1.EventNotificationConfig.pubsub_topic_name
	PubsubTopicName *string `json:"pubsubTopicName,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.HttpConfig
type HttpConfig struct {
	// If enabled, allows devices to use DeviceService via the HTTP protocol.
	//  Otherwise, any requests to DeviceService will fail for this registry.
	// +kcc:proto:field=google.cloud.iot.v1.HttpConfig.http_enabled_state
	HTTPEnabledState *string `json:"httpEnabledState,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.MqttConfig
type MqttConfig struct {
	// If enabled, allows connections using the MQTT protocol. Otherwise, MQTT
	//  connections to this registry will fail.
	// +kcc:proto:field=google.cloud.iot.v1.MqttConfig.mqtt_enabled_state
	MqttEnabledState *string `json:"mqttEnabledState,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.PublicKeyCertificate
type PublicKeyCertificate struct {
	// The certificate format.
	// +kcc:proto:field=google.cloud.iot.v1.PublicKeyCertificate.format
	Format *string `json:"format,omitempty"`

	// The certificate data.
	// +kcc:proto:field=google.cloud.iot.v1.PublicKeyCertificate.certificate
	Certificate *string `json:"certificate,omitempty"`

	// [Output only] The certificate details. Used only for X.509 certificates.
	// +kcc:proto:field=google.cloud.iot.v1.PublicKeyCertificate.x509_details
	X509Details *X509CertificateDetails `json:"x509Details,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.RegistryCredential
type RegistryCredential struct {
	// A public key certificate used to verify the device credentials.
	// +kcc:proto:field=google.cloud.iot.v1.RegistryCredential.public_key_certificate
	PublicKeyCertificate *PublicKeyCertificate `json:"publicKeyCertificate,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.StateNotificationConfig
type StateNotificationConfig struct {
	// A Cloud Pub/Sub topic name. For example,
	//  `projects/myProject/topics/deviceEvents`.
	// +kcc:proto:field=google.cloud.iot.v1.StateNotificationConfig.pubsub_topic_name
	PubsubTopicName *string `json:"pubsubTopicName,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.X509CertificateDetails
type X509CertificateDetails struct {
	// The entity that signed the certificate.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.issuer
	Issuer *string `json:"issuer,omitempty"`

	// The entity the certificate and public key belong to.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.subject
	Subject *string `json:"subject,omitempty"`

	// The time the certificate becomes valid.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.start_time
	StartTime *string `json:"startTime,omitempty"`

	// The time the certificate becomes invalid.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.expiry_time
	ExpiryTime *string `json:"expiryTime,omitempty"`

	// The algorithm used to sign the certificate.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.signature_algorithm
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`

	// The type of public key in the certificate.
	// +kcc:proto:field=google.cloud.iot.v1.X509CertificateDetails.public_key_type
	PublicKeyType *string `json:"publicKeyType,omitempty"`
}
