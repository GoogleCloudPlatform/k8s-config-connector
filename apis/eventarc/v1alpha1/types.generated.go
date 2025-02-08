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


// +kcc:proto=google.cloud.eventarc.v1.CloudRun
type CloudRun struct {
	// Required. The name of the Cloud Run service being addressed. See
	//  https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services.
	//
	//  Only services located in the same project as the trigger object
	//  can be addressed.
	// +kcc:proto:field=google.cloud.eventarc.v1.CloudRun.service
	Service *string `json:"service,omitempty"`

	// Optional. The relative path on the Cloud Run service the events should be
	//  sent to.
	//
	//  The value must conform to the definition of a URI path segment (section 3.3
	//  of RFC2396). Examples: "/route", "route", "route/subroute".
	// +kcc:proto:field=google.cloud.eventarc.v1.CloudRun.path
	Path *string `json:"path,omitempty"`

	// Required. The region the Cloud Run service is deployed in.
	// +kcc:proto:field=google.cloud.eventarc.v1.CloudRun.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Destination
type Destination struct {
	// Cloud Run fully-managed resource that receives the events. The resource
	//  should be in the same project as the trigger.
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.cloud_run
	CloudRun *CloudRun `json:"cloudRun,omitempty"`

	// The Cloud Function resource name. Cloud Functions V1 and V2 are
	//  supported.
	//  Format: `projects/{project}/locations/{location}/functions/{function}`
	//
	//  This is a read-only field. Creating Cloud Functions V1/V2 triggers is
	//  only supported via the Cloud Functions product. An error will be returned
	//  if the user sets this value.
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.cloud_function
	CloudFunction *string `json:"cloudFunction,omitempty"`

	// A GKE service capable of receiving events. The service should be running
	//  in the same project as the trigger.
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.gke
	Gke *GKE `json:"gke,omitempty"`

	// The resource name of the Workflow whose Executions are triggered by
	//  the events. The Workflow resource should be deployed in the same project
	//  as the trigger.
	//  Format: `projects/{project}/locations/{location}/workflows/{workflow}`
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.workflow
	Workflow *string `json:"workflow,omitempty"`

	// An HTTP endpoint destination described by an URI.
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.http_endpoint
	HTTPEndpoint *HttpEndpoint `json:"httpEndpoint,omitempty"`

	// Optional. Network config is used to configure how Eventarc resolves and
	//  connect to a destination.
	//  This should only be used with HttpEndpoint destination type.
	// +kcc:proto:field=google.cloud.eventarc.v1.Destination.network_config
	NetworkConfig *NetworkConfig `json:"networkConfig,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.EventFilter
type EventFilter struct {
	// Required. The name of a CloudEvents attribute. Currently, only a subset of
	//  attributes are supported for filtering. You can [retrieve a specific
	//  provider's supported event
	//  types](/eventarc/docs/list-providers#describe-provider).
	//
	//  All triggers MUST provide a filter for the 'type' attribute.
	// +kcc:proto:field=google.cloud.eventarc.v1.EventFilter.attribute
	Attribute *string `json:"attribute,omitempty"`

	// Required. The value for the attribute.
	// +kcc:proto:field=google.cloud.eventarc.v1.EventFilter.value
	Value *string `json:"value,omitempty"`

	// Optional. The operator used for matching the events with the value of the
	//  filter. If not specified, only events that have an exact key-value pair
	//  specified in the filter are matched. The allowed values are `path_pattern`
	//  and `match-path-pattern`. `path_pattern` is only allowed for GCFv1
	//  triggers.
	// +kcc:proto:field=google.cloud.eventarc.v1.EventFilter.operator
	Operator *string `json:"operator,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.GKE
type GKE struct {
	// Required. The name of the cluster the GKE service is running in. The
	//  cluster must be running in the same project as the trigger being created.
	// +kcc:proto:field=google.cloud.eventarc.v1.GKE.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Required. The name of the Google Compute Engine in which the cluster
	//  resides, which can either be compute zone (for example, us-central1-a) for
	//  the zonal clusters or region (for example, us-central1) for regional
	//  clusters.
	// +kcc:proto:field=google.cloud.eventarc.v1.GKE.location
	Location *string `json:"location,omitempty"`

	// Required. The namespace the GKE service is running in.
	// +kcc:proto:field=google.cloud.eventarc.v1.GKE.namespace
	Namespace *string `json:"namespace,omitempty"`

	// Required. Name of the GKE service.
	// +kcc:proto:field=google.cloud.eventarc.v1.GKE.service
	Service *string `json:"service,omitempty"`

	// Optional. The relative path on the GKE service the events should be sent
	//  to.
	//
	//  The value must conform to the definition of a URI path segment (section 3.3
	//  of RFC2396). Examples: "/route", "route", "route/subroute".
	// +kcc:proto:field=google.cloud.eventarc.v1.GKE.path
	Path *string `json:"path,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.HttpEndpoint
type HttpEndpoint struct {
	// Required. The URI of the HTTP enpdoint.
	//
	//  The value must be a RFC2396 URI string.
	//  Examples: `http://10.10.10.8:80/route`,
	//  `http://svc.us-central1.p.local:8080/`.
	//  Only HTTP and HTTPS protocols are supported. The host can be either a
	//  static IP addressable from the VPC specified by the network config, or
	//  an internal DNS hostname of the service resolvable via Cloud DNS.
	// +kcc:proto:field=google.cloud.eventarc.v1.HttpEndpoint.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.NetworkConfig
type NetworkConfig struct {
	// Required. Name of the NetworkAttachment that allows access to the
	//  customer's VPC. Format:
	//  `projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}`
	// +kcc:proto:field=google.cloud.eventarc.v1.NetworkConfig.network_attachment
	NetworkAttachment *string `json:"networkAttachment,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pubsub
type Pubsub struct {
	// Optional. The name of the Pub/Sub topic created and managed by Eventarc as
	//  a transport for the event delivery. Format:
	//  `projects/{PROJECT_ID}/topics/{TOPIC_NAME}`.
	//
	//  You can set an existing topic for triggers of the type
	//  `google.cloud.pubsub.topic.v1.messagePublished`. The topic you provide
	//  here is not deleted by Eventarc at trigger deletion.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pubsub.topic
	Topic *string `json:"topic,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.StateCondition
type StateCondition struct {
	// The canonical code of the condition.
	// +kcc:proto:field=google.cloud.eventarc.v1.StateCondition.code
	Code *string `json:"code,omitempty"`

	// Human-readable message.
	// +kcc:proto:field=google.cloud.eventarc.v1.StateCondition.message
	Message *string `json:"message,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Transport
type Transport struct {
	// The Pub/Sub topic and subscription used by Eventarc as a transport
	//  intermediary.
	// +kcc:proto:field=google.cloud.eventarc.v1.Transport.pubsub
	Pubsub *Pubsub `json:"pubsub,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Trigger
type Trigger struct {
	// Required. The resource name of the trigger. Must be unique within the
	//  location of the project and must be in
	//  `projects/{project}/locations/{location}/triggers/{trigger}` format.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.name
	Name *string `json:"name,omitempty"`

	// Required. Unordered list. The list of filters that applies to event
	//  attributes. Only events that match all the provided filters are sent to the
	//  destination.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.event_filters
	EventFilters []EventFilter `json:"eventFilters,omitempty"`

	// Optional. The IAM service account email associated with the trigger. The
	//  service account represents the identity of the trigger.
	//
	//  The `iam.serviceAccounts.actAs` permission must be granted on the service
	//  account to allow a principal to impersonate the service account. For more
	//  information, see the
	//  [Roles and permissions](/eventarc/docs/all-roles-permissions) page specific
	//  to the trigger destination.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Required. Destination specifies where the events should be sent to.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.destination
	Destination *Destination `json:"destination,omitempty"`

	// Optional. To deliver messages, Eventarc might use other Google Cloud
	//  products as a transport intermediary. This field contains a reference to
	//  that transport intermediary. This information can be used for debugging
	//  purposes.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.transport
	Transport *Transport `json:"transport,omitempty"`

	// Optional. User labels attached to the triggers that can be used to group
	//  resources.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The name of the channel associated with the trigger in
	//  `projects/{project}/locations/{location}/channels/{channel}` format.
	//  You must provide a channel to receive events from Eventarc SaaS partners.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.channel
	Channel *string `json:"channel,omitempty"`

	// Optional. EventDataContentType specifies the type of payload in MIME
	//  format that is expected from the CloudEvent data field. This is set to
	//  `application/json` if the value is not defined.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.event_data_content_type
	EventDataContentType *string `json:"eventDataContentType,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Pubsub
type PubsubObservedState struct {
	// Output only. The name of the Pub/Sub subscription created and managed by
	//  Eventarc as a transport for the event delivery. Format:
	//  `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`.
	// +kcc:proto:field=google.cloud.eventarc.v1.Pubsub.subscription
	Subscription *string `json:"subscription,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Transport
type TransportObservedState struct {
	// The Pub/Sub topic and subscription used by Eventarc as a transport
	//  intermediary.
	// +kcc:proto:field=google.cloud.eventarc.v1.Transport.pubsub
	Pubsub *PubsubObservedState `json:"pubsub,omitempty"`
}

// +kcc:proto=google.cloud.eventarc.v1.Trigger
type TriggerObservedState struct {
	// Output only. Server-assigned unique identifier for the trigger. The value
	//  is a UUID4 string and guaranteed to remain unchanged until the resource is
	//  deleted.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. To deliver messages, Eventarc might use other Google Cloud
	//  products as a transport intermediary. This field contains a reference to
	//  that transport intermediary. This information can be used for debugging
	//  purposes.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.transport
	Transport *TransportObservedState `json:"transport,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. Whether or not this Trigger satisfies the requirements of
	//  physical zone separation
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	//  other fields, and might be sent only on create requests to ensure that the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.eventarc.v1.Trigger.etag
	Etag *string `json:"etag,omitempty"`
}
