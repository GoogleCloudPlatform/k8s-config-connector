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


// +kcc:proto=google.cloud.iot.v1.Device
type Device struct {
	// The user-defined device identifier. The device ID must be unique
	//  within a device registry.
	// +kcc:proto:field=google.cloud.iot.v1.Device.id
	ID *string `json:"id,omitempty"`

	// The resource path name. For example,
	//  `projects/p1/locations/us-central1/registries/registry0/devices/dev0` or
	//  `projects/p1/locations/us-central1/registries/registry0/devices/{num_id}`.
	//  When `name` is populated as a response from the service, it always ends
	//  in the device numeric ID.
	// +kcc:proto:field=google.cloud.iot.v1.Device.name
	Name *string `json:"name,omitempty"`

	// [Output only] A server-defined unique numeric ID for the device. This is a
	//  more compact way to identify devices, and it is globally unique.
	// +kcc:proto:field=google.cloud.iot.v1.Device.num_id
	NumID *uint64 `json:"numID,omitempty"`

	// The credentials used to authenticate this device. To allow credential
	//  rotation without interruption, multiple device credentials can be bound to
	//  this device. No more than 3 credentials can be bound to a single device at
	//  a time. When new credentials are added to a device, they are verified
	//  against the registry credentials. For details, see the description of the
	//  `DeviceRegistry.credentials` field.
	// +kcc:proto:field=google.cloud.iot.v1.Device.credentials
	Credentials []DeviceCredential `json:"credentials,omitempty"`

	// [Output only] The last time an MQTT `PINGREQ` was received. This field
	//  applies only to devices connecting through MQTT. MQTT clients usually only
	//  send `PINGREQ` messages if the connection is idle, and no other messages
	//  have been sent. Timestamps are periodically collected and written to
	//  storage; they may be stale by a few minutes.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_heartbeat_time
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty"`

	// [Output only] The last time a telemetry event was received. Timestamps are
	//  periodically collected and written to storage; they may be stale by a few
	//  minutes.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_event_time
	LastEventTime *string `json:"lastEventTime,omitempty"`

	// [Output only] The last time a state event was received. Timestamps are
	//  periodically collected and written to storage; they may be stale by a few
	//  minutes.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_state_time
	LastStateTime *string `json:"lastStateTime,omitempty"`

	// [Output only] The last time a cloud-to-device config version acknowledgment
	//  was received from the device. This field is only for configurations
	//  sent through MQTT.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_config_ack_time
	LastConfigAckTime *string `json:"lastConfigAckTime,omitempty"`

	// [Output only] The last time a cloud-to-device config version was sent to
	//  the device.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_config_send_time
	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty"`

	// If a device is blocked, connections or requests from this device will fail.
	//  Can be used to temporarily prevent the device from connecting if, for
	//  example, the sensor is generating bad data and needs maintenance.
	// +kcc:proto:field=google.cloud.iot.v1.Device.blocked
	Blocked *bool `json:"blocked,omitempty"`

	// [Output only] The time the most recent error occurred, such as a failure to
	//  publish to Cloud Pub/Sub. This field is the timestamp of
	//  'last_error_status'.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_error_time
	LastErrorTime *string `json:"lastErrorTime,omitempty"`

	// [Output only] The error message of the most recent error, such as a failure
	//  to publish to Cloud Pub/Sub. 'last_error_time' is the timestamp of this
	//  field. If no errors have occurred, this field has an empty message
	//  and the status code 0 == OK. Otherwise, this field is expected to have a
	//  status code other than OK.
	// +kcc:proto:field=google.cloud.iot.v1.Device.last_error_status
	LastErrorStatus *Status `json:"lastErrorStatus,omitempty"`

	// The most recent device configuration, which is eventually sent from
	//  Cloud IoT Core to the device. If not present on creation, the
	//  configuration will be initialized with an empty payload and version value
	//  of `1`. To update this field after creation, use the
	//  `DeviceManager.ModifyCloudToDeviceConfig` method.
	// +kcc:proto:field=google.cloud.iot.v1.Device.config
	Config *DeviceConfig `json:"config,omitempty"`

	// [Output only] The state most recently received from the device. If no state
	//  has been reported, this field is not present.
	// +kcc:proto:field=google.cloud.iot.v1.Device.state
	State *DeviceState `json:"state,omitempty"`

	// **Beta Feature**
	//
	//  The logging verbosity for device activity. If unspecified,
	//  DeviceRegistry.log_level will be used.
	// +kcc:proto:field=google.cloud.iot.v1.Device.log_level
	LogLevel *string `json:"logLevel,omitempty"`

	// The metadata key-value pairs assigned to the device. This metadata is not
	//  interpreted or indexed by Cloud IoT Core. It can be used to add contextual
	//  information for the device.
	//
	//  Keys must conform to the regular expression [a-zA-Z][a-zA-Z0-9-_.+~%]+ and
	//  be less than 128 bytes in length.
	//
	//  Values are free-form strings. Each value must be less than or equal to 32
	//  KB in size.
	//
	//  The total size of all keys and values must be less than 256 KB, and the
	//  maximum number of key-value pairs is 500.
	// +kcc:proto:field=google.cloud.iot.v1.Device.metadata
	Metadata map[string]string `json:"metadata,omitempty"`

	// Gateway-related configuration and state.
	// +kcc:proto:field=google.cloud.iot.v1.Device.gateway_config
	GatewayConfig *GatewayConfig `json:"gatewayConfig,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.DeviceConfig
type DeviceConfig struct {
	// [Output only] The version of this update. The version number is assigned by
	//  the server, and is always greater than 0 after device creation. The
	//  version must be 0 on the `CreateDevice` request if a `config` is
	//  specified; the response of `CreateDevice` will always have a value of 1.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceConfig.version
	Version *int64 `json:"version,omitempty"`

	// [Output only] The time at which this configuration version was updated in
	//  Cloud IoT Core. This timestamp is set by the server.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceConfig.cloud_update_time
	CloudUpdateTime *string `json:"cloudUpdateTime,omitempty"`

	// [Output only] The time at which Cloud IoT Core received the
	//  acknowledgment from the device, indicating that the device has received
	//  this configuration version. If this field is not present, the device has
	//  not yet acknowledged that it received this version. Note that when
	//  the config was sent to the device, many config versions may have been
	//  available in Cloud IoT Core while the device was disconnected, and on
	//  connection, only the latest version is sent to the device. Some
	//  versions may never be sent to the device, and therefore are never
	//  acknowledged. This timestamp is set by Cloud IoT Core.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceConfig.device_ack_time
	DeviceAckTime *string `json:"deviceAckTime,omitempty"`

	// The device configuration data.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceConfig.binary_data
	BinaryData []byte `json:"binaryData,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.DeviceCredential
type DeviceCredential struct {
	// A public key used to verify the signature of JSON Web Tokens (JWTs).
	//  When adding a new device credential, either via device creation or via
	//  modifications, this public key credential may be required to be signed by
	//  one of the registry level certificates. More specifically, if the
	//  registry contains at least one certificate, any new device credential
	//  must be signed by one of the registry certificates. As a result,
	//  when the registry contains certificates, only X.509 certificates are
	//  accepted as device credentials. However, if the registry does
	//  not contain a certificate, self-signed certificates and public keys will
	//  be accepted. New device credentials must be different from every
	//  registry-level certificate.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceCredential.public_key
	PublicKey *PublicKeyCredential `json:"publicKey,omitempty"`

	// [Optional] The time at which this credential becomes invalid. This
	//  credential will be ignored for new client authentication requests after
	//  this timestamp; however, it will not be automatically deleted.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceCredential.expiration_time
	ExpirationTime *string `json:"expirationTime,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.DeviceState
type DeviceState struct {
	// [Output only] The time at which this state version was updated in Cloud
	//  IoT Core.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceState.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The device state data.
	// +kcc:proto:field=google.cloud.iot.v1.DeviceState.binary_data
	BinaryData []byte `json:"binaryData,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.GatewayConfig
type GatewayConfig struct {
	// Indicates whether the device is a gateway.
	// +kcc:proto:field=google.cloud.iot.v1.GatewayConfig.gateway_type
	GatewayType *string `json:"gatewayType,omitempty"`

	// Indicates how to authorize and/or authenticate devices to access the
	//  gateway.
	// +kcc:proto:field=google.cloud.iot.v1.GatewayConfig.gateway_auth_method
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty"`

	// [Output only] The ID of the gateway the device accessed most recently.
	// +kcc:proto:field=google.cloud.iot.v1.GatewayConfig.last_accessed_gateway_id
	LastAccessedGatewayID *string `json:"lastAccessedGatewayID,omitempty"`

	// [Output only] The most recent time at which the device accessed the gateway
	//  specified in `last_accessed_gateway`.
	// +kcc:proto:field=google.cloud.iot.v1.GatewayConfig.last_accessed_gateway_time
	LastAccessedGatewayTime *string `json:"lastAccessedGatewayTime,omitempty"`
}

// +kcc:proto=google.cloud.iot.v1.PublicKeyCredential
type PublicKeyCredential struct {
	// The format of the key.
	// +kcc:proto:field=google.cloud.iot.v1.PublicKeyCredential.format
	Format *string `json:"format,omitempty"`

	// The key data.
	// +kcc:proto:field=google.cloud.iot.v1.PublicKeyCredential.key
	Key *string `json:"key,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}
