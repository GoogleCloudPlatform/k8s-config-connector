// Copyright 2024 Google LLC
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

package v1alpha1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigQueryConnectionConnectionGVK = GroupVersion.WithKind("BigQueryConnectionConnection")

type Parent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`
}

// BigQueryConnectionConnectionSpec defines the desired state to connect BigQuery to external resources
// +kcc:proto=google.cloud.bigquery.connection.v1.Connection
type BigQueryConnectionConnectionSpec struct {
	Parent `json:",inline"`

	// The BigQuery ConnectionID. This is a server-generated ID in the UUID format.
	// If not provided, ConfigConnector will create a new Connection and store the UUID in `status.serviceGeneratedID` field.
	// + optional
	ResourceID *string `json:"resourceID,omitempty"`

	// User provided display name for the connection.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// User provided description.
	Description *string `json:"description,omitempty"`

	/* NOTYET
	// Cloud SQL properties.
	CloudSql *CloudSqlProperties `json:"cloudSql,omitempty"`
	*/

	/* NOTYET
	// Amazon Web Services (AWS) properties.
	Aws *AwsProperties `json:"aws,omitempty"`
	*/

	/* NOTYET
	// Azure properties.
	Azure *AzureProperties `json:"azure,omitempty"`
	*/

	/* NOTYET
	// Cloud Spanner properties.
	CloudSpanner *CloudSpannerProperties `json:"cloudSpanner,omitempty"`
	*/

	/* NOTYET
	// Spark properties.
	Spark *SparkProperties `json:"spark,omitempty"`
	*/

	/* NOTYET
	// Optional. Salesforce DataCloud properties. This field is intended for
	//  use only by Salesforce partner projects. This field contains properties
	//  for your Salesforce DataCloud connection.
	SalesforceDataCloud *SalesforceDataCloudProperties `json:"salesforceDataCloud,omitempty"`
	*/

	// Use Cloud Resource properties.
	CloudResourceSpec *CloudResourcePropertiesSpec `json:"cloudResource,omitempty"`
}

// BigQueryConnectionConnectionStatus defines the config connector machine state of BigQueryConnectionConnection
type BigQueryConnectionConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigQueryConnectionConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *BigQueryConnectionConnectionObservedState `json:"observedState,omitempty"`
}

// BigQueryConnectionConnectionSpec defines the desired state of BigQueryConnectionConnection
// +kcc:proto=google.cloud.bigquery.connection.v1.Connection
type BigQueryConnectionConnectionObservedState struct {
	CloudResource *CloudResourcePropertiesStatus `json:"cloudResource,omitempty"`

	// The display name for the connection.
	FriendlyName *string `json:"friendlyName,omitempty"`

	// The description for the connection.
	Description *string `json:"description,omitempty"`

	/*
		// Cloud SQL properties.
		CloudSql *CloudSqlProperties `json:"cloudSql,omitempty"`

		// Amazon Web Services (AWS) properties.
		Aws *AwsProperties `json:"aws,omitempty"`

		// Azure properties.
		Azure *AzureProperties `json:"azure,omitempty"`

		// Cloud Spanner properties.
		CloudSpanner *CloudSpannerProperties `json:"cloudSpanner,omitempty"`

		// Spark properties.
		Spark *SparkProperties `json:"spark,omitempty"`

		// Optional. Salesforce DataCloud properties. This field is intended for
		//  use only by Salesforce partner projects. This field contains properties
		//  for your Salesforce DataCloud connection.
		SalesforceDataCloud *SalesforceDataCloudProperties `json:"salesforceDataCloud,omitempty"`
	*/

	// Output only. True, if credential is configured for this connection.
	HasCredential *bool `json:"hasCredential,omitempty"`
}

type CloudResourcePropertiesSpec struct{}

// +kcc:proto=google.cloud.bigquery.connection.v1.CloudResourceProperties
type CloudResourcePropertiesStatus struct {
	// Output only. The account ID of the service created for the purpose of this
	//  connection.
	//
	//  The service account does not have any permissions associated with it
	//  when it is created. After creation, customers delegate permissions
	//  to the service account. When the connection is used in the context of an
	//  operation in BigQuery, the service account will be used to connect to the
	//  desired resources in GCP.
	//
	//  The account ID is in the form of:
	//    <service-1234>@gcp-sa-bigquery-cloudresource.iam.gserviceaccount.com
	ServiceAccountID *string `json:"serviceAccountID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigqueryconnectionconnection;gcpbigqueryconnectionconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigQueryConnectionConnection is the Schema for the BigQueryConnectionConnection API
// +k8s:openapi-gen=true
type BigQueryConnectionConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigQueryConnectionConnectionSpec   `json:"spec,omitempty"`
	Status BigQueryConnectionConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigQueryConnectionConnectionList contains a list of BigQueryConnectionConnection
type BigQueryConnectionConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigQueryConnectionConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigQueryConnectionConnection{}, &BigQueryConnectionConnectionList{})
}
