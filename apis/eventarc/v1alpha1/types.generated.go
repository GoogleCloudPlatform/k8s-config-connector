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


// +kcc:proto=google.cloud.eventarc.v1.LoggingConfig
type LoggingConfig struct {
	// Optional. The minimum severity of logs that will be sent to
	//  Stackdriver/Platform Telemetry. Logs at severitiy ≥ this value will be
	//  sent, unless it is NONE.
	// +kcc:proto:field=google.cloud.eventarc.v1.LoggingConfig.log_severity
	LogSeverity *string `json:"logSeverity,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline
type Pipeline struct {
	// Identifier. The resource name of the Pipeline. Must be unique within the
	//  location of the project and must be in
	//  `projects/{project}/locations/{location}/pipelines/{pipeline}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.name
	Name *string `json:"name,omitempty"`

	// Optional. User labels attached to the Pipeline that can be used to group
	//  resources. An object containing a list of "key": value pairs. Example: {
	//  "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User-defined annotations. See
	//  https://google.aip.dev/128#annotations.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Display name of resource.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. List of destinations to which messages will be forwarded.
	//  Currently, exactly one destination is supported per Pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.destinations
	Destinations []Pipeline_Destination `json:"destinations,omitempty"`

	// Optional. List of mediation operations to be performed on the message.
	//  Currently, only one Transformation operation is allowed in each Pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.mediations
	Mediations []Pipeline_Mediation `json:"mediations,omitempty"`

	// Optional. Resource name of a KMS crypto key (managed by the user) used to
	//  encrypt/decrypt the event data. If not set, an internal Google-owned key
	//  will be used to encrypt messages. It must match the pattern
	//  "projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}".
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.crypto_key_name
	CryptoKeyName *string `json:"cryptoKeyName,omitempty"`

	// Optional. The payload format expected for the messages received by the
	//  Pipeline. If input_payload_format is set then any messages not matching
	//  this format will be treated as persistent errors. If input_payload_format
	//  is not set, then the message data will be treated as an opaque binary and
	//  no output format can be set on the Pipeline through the
	//  Pipeline.Destination.output_payload_format field. Any Mediations on the
	//  Pipeline that involve access to the data field will fail as persistent
	//  errors.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.input_payload_format
	InputPayloadFormat *Pipeline_MessagePayloadFormat `json:"inputPayloadFormat,omitempty"`

	// Optional. Config to control Platform Logging for Pipelines.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.logging_config
	LoggingConfig *LoggingConfig `json:"loggingConfig,omitempty"`

	// Optional. The retry policy to use in the pipeline.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.retry_policy
	RetryPolicy *Pipeline_RetryPolicy `json:"retryPolicy,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on create requests to ensure that the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination
type Pipeline_Destination struct {
	// Optional. Network config is used to configure how Pipeline resolves and
	//  connects to a destination.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.network_config
	NetworkConfig *Pipeline_Destination_NetworkConfig `json:"networkConfig,omitempty"`

	// Optional. An HTTP endpoint destination described by an URI.
	//  If a DNS FQDN is provided as the endpoint, Pipeline will create a
	//  peering zone to the consumer VPC and forward DNS requests to the VPC
	//  specified by network config to resolve the service endpoint. See:
	//  https://cloud.google.com/dns/docs/zones/zones-overview#peering_zones
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.http_endpoint
	HTTPEndpoint *Pipeline_Destination_HttpEndpoint `json:"httpEndpoint,omitempty"`

	// Optional. The resource name of the Workflow whose Executions are
	//  triggered by the events. The Workflow resource should be deployed in
	//  the same project as the Pipeline. Format:
	//  `projects/{project}/locations/{location}/workflows/{workflow}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.workflow
	Workflow *string `json:"workflow,omitempty"`

	// Optional. The resource name of the Message Bus to which events should
	//  be published. The Message Bus resource should exist in the same project
	//  as the Pipeline. Format:
	//  `projects/{project}/locations/{location}/messageBuses/{message_bus}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.message_bus
	MessageBus *string `json:"messageBus,omitempty"`

	// Optional. The resource name of the Pub/Sub topic to which events should
	//  be published. Format:
	//  `projects/{project}/locations/{location}/topics/{topic}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.topic
	Topic *string `json:"topic,omitempty"`

	// Optional. An authentication config used to authenticate message requests,
	//  such that destinations can verify the source. For example, this can be
	//  used with private GCP destinations that require GCP credentials to access
	//  like Cloud Run. This field is optional and should be set only by users
	//  interested in authenticated push
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.authentication_config
	AuthenticationConfig *Pipeline_Destination_AuthenticationConfig `json:"authenticationConfig,omitempty"`

	// Optional. The message format before it is delivered to the destination.
	//  If not set, the message will be delivered in the format it was originally
	//  delivered to the Pipeline. This field can only be set if
	//  Pipeline.input_payload_format is also set.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.output_payload_format
	OutputPayloadFormat *Pipeline_MessagePayloadFormat `json:"outputPayloadFormat,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig
type Pipeline_Destination_AuthenticationConfig struct {
	// Optional. This authenticate method will apply Google OIDC tokens
	//  signed by a GCP service account to the requests.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.google_oidc
	GoogleOidc *Pipeline_Destination_AuthenticationConfig_OidcToken `json:"googleOidc,omitempty"`

	// Optional. If specified, an [OAuth
	//  token](https://developers.google.com/identity/protocols/OAuth2) will
	//  be generated and attached as an `Authorization` header in the HTTP
	//  request.
	//
	//  This type of authorization should generally only be used when calling
	//  Google APIs hosted on *.googleapis.com.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.oauth_token
	OauthToken *Pipeline_Destination_AuthenticationConfig_OAuthToken `json:"oauthToken,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken
type Pipeline_Destination_AuthenticationConfig_OAuthToken struct {
	// Required. Service account email used to generate the [OAuth
	//  token](https://developers.google.com/identity/protocols/OAuth2).
	//  The principal who calls this API must have
	//  iam.serviceAccounts.actAs permission in the service account. See
	//  https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common
	//  for more information. Eventarc service agents must have
	//  roles/roles/iam.serviceAccountTokenCreator role to allow Pipeline
	//  to create OAuth2 tokens for authenticated requests.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. OAuth scope to be used for generating OAuth access token.
	//  If not specified, "https://www.googleapis.com/auth/cloud-platform"
	//  will be used.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OAuthToken.scope
	Scope *string `json:"scope,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken
type Pipeline_Destination_AuthenticationConfig_OidcToken struct {
	// Required. Service account email used to generate the OIDC Token.
	//  The principal who calls this API must have
	//  iam.serviceAccounts.actAs permission in the service account. See
	//  https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common
	//  for more information. Eventarc service agents must have
	//  roles/roles/iam.serviceAccountTokenCreator role to allow the
	//  Pipeline to create OpenID tokens for authenticated requests.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Audience to be used to generate the OIDC Token. The
	//  audience claim identifies the recipient that the JWT is intended for.
	//  If unspecified, the destination URI will be used.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.AuthenticationConfig.OidcToken.audience
	Audience *string `json:"audience,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.HttpEndpoint
type Pipeline_Destination_HttpEndpoint struct {
	// Required. The URI of the HTTP enpdoint.
	//
	//  The value must be a RFC2396 URI string.
	//  Examples: `https://svc.us-central1.p.local:8080/route`.
	//  Only the HTTPS protocol is supported.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.HttpEndpoint.uri
	URI *string `json:"uri,omitempty"`

	// Optional. The CEL expression used to modify how the destination-bound
	//  HTTP request is constructed.
	//
	//  If a binding expression is not specified here, the message
	//  is treated as a CloudEvent and is mapped to the HTTP request according
	//  to the CloudEvent HTTP Protocol Binding Binary Content Mode. In this
	//  representation, all fields except the `data` and `datacontenttype`
	//  field on the message are mapped to HTTP request headers with a prefix
	//  of `ce-`.
	//
	//  To construct the HTTP request payload and the value of the content-type
	//  HTTP header, the payload format is defined as follows:
	//  1) Use the output_payload_format_type on the Pipeline.Destination if it
	//     is set, else:
	//  2) Use the input_payload_format_type on the Pipeline if it is set,
	//     else:
	//  3) Treat the payload as opaque binary data.
	//
	//  The `data` field of the message is converted to the payload format or
	//  left as-is for case 3) and then attached as the payload of the HTTP
	//  request. The `content-type` header on the HTTP request is set to the
	//  payload format type or left empty for case 3). However, if a mediation
	//  has updated the `datacontenttype` field on the message so that it is
	//  not the same as the payload format type but it is still a prefix of the
	//  payload format type, then the `content-type` header on the HTTP request
	//  is set to this `datacontenttype` value. For example, if the
	//  `datacontenttype` is "application/json" and the payload format type is
	//  "application/json; charset=utf-8", then the `content-type` header on
	//  the HTTP request is set to "application/json; charset=utf-8".
	//
	//  If a non-empty binding expression is specified then this expression is
	//  used to modify the default CloudEvent HTTP Protocol Binding Binary
	//  Content representation.
	//  The result of the CEL expression must be a map of key/value pairs
	//  which is used as follows:
	//  - If a map named `headers` exists on the result of the expression,
	//  then its key/value pairs are directly mapped to the HTTP request
	//  headers. The headers values are constructed from the corresponding
	//  value type’s canonical representation. If the `headers` field doesn’t
	//  exist then the resulting HTTP request will be the headers of the
	//  CloudEvent HTTP Binding Binary Content Mode representation of the final
	//  message. Note: If the specified binding expression, has updated the
	//  `datacontenttype` field on the message so that it is not the same as
	//  the payload format type but it is still a prefix of the payload format
	//  type, then the `content-type` header in the `headers` map is set to
	//  this `datacontenttype` value.
	//  - If a field named `body` exists on the result of the expression then
	//  its value is directly mapped to the body of the request. If the value
	//  of the `body` field is of type bytes or string then it is used for
	//  the HTTP request body as-is, with no conversion. If the body field is
	//  of any other type then it is converted to a JSON string. If the body
	//  field does not exist then the resulting payload of the HTTP request
	//  will be data value of the CloudEvent HTTP Binding Binary Content Mode
	//  representation of the final message as described earlier.
	//  - Any other fields in the resulting expression will be ignored.
	//
	//  The CEL expression may access the incoming CloudEvent message in its
	//  definition, as follows:
	//  - The `data` field of the incoming CloudEvent message can be accessed
	//  using the `message.data` value. Subfields of `message.data` may also be
	//  accessed if an input_payload_format has been specified on the Pipeline.
	//  - Each attribute of the incoming CloudEvent message can be accessed
	//  using the `message.<key>` value, where <key> is replaced with the
	//  name of the attribute.
	//  - Existing headers can be accessed in the CEL expression using the
	//  `headers` variable. The `headers` variable defines a map of key/value
	//  pairs corresponding to the HTTP headers of the CloudEvent HTTP Binding
	//  Binary Content Mode representation of the final message as described
	//  earlier. For example, the following CEL expression can be used to
	//  construct an HTTP request by adding an additional header to the HTTP
	//  headers of the CloudEvent HTTP Binding Binary Content Mode
	//  representation of the final message and by overwriting the body of the
	//  request:
	//
	//  ```
	//  {
	//    "headers": headers.merge({"new-header-key": "new-header-value"}),
	//    "body": "new-body"
	//  }
	//  ```
	//
	//  Additionally, the following CEL extension functions are provided for
	//  use in this CEL expression:
	//  - toBase64Url:
	//    map.toBase64Url() -> string
	//      - Converts a CelValue to a base64url encoded string
	//  - toJsonString: map.toJsonString() -> string
	//      - Converts a CelValue to a JSON string
	//  - merge:
	//    map1.merge(map2) -> map3
	//      - Merges the passed CEL map with the existing CEL map the
	//      function is applied to.
	//      - If the same key exists in both maps, if the key's value is type
	//      map both maps are merged else the value from the passed map is
	//      used.
	//  - denormalize:
	//    map.denormalize() -> map
	//      - Denormalizes a CEL map such that every value of type map or key
	//      in the map is expanded to return a single level map.
	//      - The resulting keys are "." separated indices of the map keys.
	//      - For example:
	//        {
	//          "a": 1,
	//          "b": {
	//            "c": 2,
	//            "d": 3
	//          }
	//          "e": [4, 5]
	//        }
	//        .denormalize()
	//        -> {
	//          "a": 1,
	//          "b.c": 2,
	//          "b.d": 3,
	//          "e.0": 4,
	//          "e.1": 5
	//        }
	//  - setField:
	//    map.setField(key, value) -> message
	//      - Sets the field of the message with the given key to the
	//      given value.
	//      - If the field is not present it will be added.
	//      - If the field is present it will be overwritten.
	//      - The key can be a dot separated path to set a field in a nested
	//      message.
	//      - Key must be of type string.
	//      - Value may be any valid type.
	//  - removeFields:
	//    map.removeFields([key1, key2, ...]) -> message
	//      - Removes the fields of the map with the given keys.
	//      - The keys can be a dot separated path to remove a field in a
	//      nested message.
	//      - If a key is not found it will be ignored.
	//      - Keys must be of type string.
	//  - toMap:
	//    [map1, map2, ...].toMap() -> map
	//      - Converts a CEL list of CEL maps to a single CEL map
	//  - toDestinationPayloadFormat():
	//    message.data.toDestinationPayloadFormat() -> string or bytes
	//      - Converts the message data to the destination payload format
	//      specified in Pipeline.Destination.output_payload_format
	//      - This function is meant to be applied to the message.data field.
	//      - If the destination payload format is not set, the function will
	//      return the message data unchanged.
	//  - toCloudEventJsonWithPayloadFormat:
	//    message.toCloudEventJsonWithPayloadFormat() -> map
	//      - Converts a message to the corresponding structure of JSON
	//      format for CloudEvents
	//      - This function applies toDestinationPayloadFormat() to the
	//      message data. It also sets the corresponding datacontenttype of
	//      the CloudEvent, as indicated by
	//      Pipeline.Destination.output_payload_format. If no
	//      output_payload_format is set it will use the existing
	//      datacontenttype on the CloudEvent if present, else leave
	//      datacontenttype absent.
	//      - This function expects that the content of the message will
	//      adhere to the standard CloudEvent format. If it doesn’t then this
	//      function will fail.
	//      - The result is a CEL map that corresponds to the JSON
	//      representation of the CloudEvent. To convert that data to a JSON
	//      string it can be chained with the toJsonString function.
	//
	//  The Pipeline expects that the message it receives adheres to the
	//  standard CloudEvent format. If it doesn’t then the outgoing message
	//  request may fail with a persistent error.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.HttpEndpoint.message_binding_template
	MessageBindingTemplate *string `json:"messageBindingTemplate,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Destination.NetworkConfig
type Pipeline_Destination_NetworkConfig struct {
	// Required. Name of the NetworkAttachment that allows access to the
	//  consumer VPC. Format:
	//  `projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Destination.NetworkConfig.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Mediation
type Pipeline_Mediation struct {
	// Optional. How the Pipeline is to transform messages
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Mediation.transformation
	Transformation *Pipeline_Mediation_Transformation `json:"transformation,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.Mediation.Transformation
type Pipeline_Mediation_Transformation struct {
	// Optional. The CEL expression template to apply to transform messages.
	//  The following CEL extension functions are provided for
	//  use in this CEL expression:
	//  - merge:
	//    map1.merge(map2) -> map3
	//      - Merges the passed CEL map with the existing CEL map the
	//      function is applied to.
	//      - If the same key exists in both maps, if the key's value is type
	//      map both maps are merged else the value from the passed map is
	//      used.
	//  - denormalize:
	//    map.denormalize() -> map
	//      - Denormalizes a CEL map such that every value of type map or key
	//      in the map is expanded to return a single level map.
	//      - The resulting keys are "." separated indices of the map keys.
	//      - For example:
	//        {
	//          "a": 1,
	//          "b": {
	//            "c": 2,
	//            "d": 3
	//          }
	//          "e": [4, 5]
	//        }
	//        .denormalize()
	//        -> {
	//          "a": 1,
	//          "b.c": 2,
	//          "b.d": 3,
	//          "e.0": 4,
	//          "e.1": 5
	//        }
	//  - setField:
	//    map.setField(key, value) -> message
	//      - Sets the field of the message with the given key to the
	//      given value.
	//      - If the field is not present it will be added.
	//      - If the field is present it will be overwritten.
	//      - The key can be a dot separated path to set a field in a nested
	//      message.
	//      - Key must be of type string.
	//      - Value may be any valid type.
	//  - removeFields:
	//    map.removeFields([key1, key2, ...]) -> message
	//      - Removes the fields of the map with the given keys.
	//      - The keys can be a dot separated path to remove a field in a
	//      nested message.
	//      - If a key is not found it will be ignored.
	//      - Keys must be of type string.
	//  - toMap:
	//    [map1, map2, ...].toMap() -> map
	//      - Converts a CEL list of CEL maps to a single CEL map
	//  - toDestinationPayloadFormat():
	//    message.data.toDestinationPayloadFormat() -> string or bytes
	//      - Converts the message data to the destination payload format
	//      specified in Pipeline.Destination.output_payload_format
	//      - This function is meant to be applied to the message.data field.
	//      - If the destination payload format is not set, the function will
	//      return the message data unchanged.
	//  - toCloudEventJsonWithPayloadFormat:
	//    message.toCloudEventJsonWithPayloadFormat() -> map
	//      - Converts a message to the corresponding structure of JSON
	//      format for CloudEvents
	//      - This function applies toDestinationPayloadFormat() to the
	//      message data. It also sets the corresponding datacontenttype of
	//      the CloudEvent, as indicated by
	//      Pipeline.Destination.output_payload_format. If no
	//      output_payload_format is set it will use the existing
	//      datacontenttype on the CloudEvent if present, else leave
	//      datacontenttype absent.
	//      - This function expects that the content of the message will
	//      adhere to the standard CloudEvent format. If it doesn’t then this
	//      function will fail.
	//      - The result is a CEL map that corresponds to the JSON
	//      representation of the CloudEvent. To convert that data to a JSON
	//      string it can be chained with the toJsonString function.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.Mediation.Transformation.transformation_template
	TransformationTemplate *string `json:"transformationTemplate,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat
type Pipeline_MessagePayloadFormat struct {
	// Optional. Protobuf format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.protobuf
	Protobuf *Pipeline_MessagePayloadFormat_ProtobufFormat `json:"protobuf,omitempty"`

	// Optional. Avro format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.avro
	Avro *Pipeline_MessagePayloadFormat_AvroFormat `json:"avro,omitempty"`

	// Optional. JSON format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.json
	Json *Pipeline_MessagePayloadFormat_JsonFormat `json:"json,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.AvroFormat
type Pipeline_MessagePayloadFormat_AvroFormat struct {
	// Optional. The entire schema definition is stored in this field.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.AvroFormat.schema_definition
	SchemaDefinition *string `json:"schemaDefinition,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.JsonFormat
type Pipeline_MessagePayloadFormat_JsonFormat struct {
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.ProtobufFormat
type Pipeline_MessagePayloadFormat_ProtobufFormat struct {
	// Optional. The entire schema definition is stored in this field.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.MessagePayloadFormat.ProtobufFormat.schema_definition
	SchemaDefinition *string `json:"schemaDefinition,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline.RetryPolicy
type Pipeline_RetryPolicy struct {
	// Optional. The maximum number of delivery attempts for any message. The
	//  value must be between 1 and 100. The default value for this field is 5.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.RetryPolicy.max_attempts
	MaxAttempts *int32 `json:"maxAttempts,omitempty"`

	// Optional. The minimum amount of seconds to wait between retry attempts.
	//  The value must be between 1 and 600. The default value for this field
	//  is 5.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.RetryPolicy.min_retry_delay
	MinRetryDelay *string `json:"minRetryDelay,omitempty"`

	// Optional. The maximum amount of seconds to wait between retry attempts.
	//  The value must be between 1 and 600. The default value for this field
	//  is 60.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.RetryPolicy.max_retry_delay
	MaxRetryDelay *string `json:"maxRetryDelay,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pipeline
type PipelineObservedState struct {
	// Output only. The creation time.
	//  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	//  to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
	//  "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	//  A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	//  to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
	//  "2014-10-02T15:01:23.045123456Z".
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Server-assigned unique identifier for the Pipeline. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pipeline.uid
	Uid *string `json:"uid,omitempty"`
}
