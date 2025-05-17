// Copyright 2022 Google LLC
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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// ConfigConnectorSpec defines the desired state of ConfigConnector
type ConfigConnectorSpec struct {
	addonv1alpha1.CommonSpec `json:"-"`
	// The Google Service Account to be used by Config Connector to authenticate with Google Cloud APIs. This field is used only when running in cluster mode with Workload Identity enabled.
	// See Google Kubernetes Engine (GKE) workload-identity (https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) for details. This field cannot be specified together with `credentialSecretName`.
	// For namespaced mode, use `googleServiceAccount` in ConfigConnectorContext CRD to specify the Google Service Account to be used to authenticate with Google Cloud APIs per namespace.
	GoogleServiceAccount string `json:"googleServiceAccount,omitempty"`

	// The Kubernetes secret that contains the Google Service Account Key's credentials to be used by ConfigConnector to authenticate with Google Cloud APIs. This field is used only when in cluster mode.
	// It's recommended to use `googleServiceAccount` when running ConfigConnector in Google Kubernetes Engine (GKE) clusters with Workload Identity enabled.
	// This field cannot be specified together with `googleServiceAccount`.
	CredentialSecretName string `json:"credentialSecretName,omitempty"`

	// The mode that Config Connector will run in. This can be either 'cluster' or 'namespaced'. The default is 'namespaced'.
	// Cluster mode uses a single Google Service Account to create and manage resources, even if you are using Config Connector to manage multiple Projects.
	// You must specify either `credentialSecretName` or `googleServiceAccount` when in cluster mode, but not both.
	// Namespaced mode allows you to use different Google service accounts for different Projects.
	// When in namespaced mode, you must create a ConfigConnectorContext object per namespace that you want to enable Config Connector in, and each must set `googleServiceAccount` to specify the Google Service Account to be used to authenticate with Google Cloud APIs for the namespace.
	//+kubebuilder:validation:Enum=cluster;namespaced
	Mode string `json:"mode,omitempty"`

	// The actuation mode of Config Connector controls how resources are actuated onto the cloud provider.
	// This can be either 'Reconciling' or 'Paused'.
	// In 'Paused', k8s resources are still reconciled with the api server but not actuated onto the cloud provider.
	// If Config Connector is running in 'namespaced' mode, then the value in ConfigConnectorContext (CCC) takes precedence.
	// If CCC doesn't define a value but ConfigConnector (CC) does, we defer to that value. Otherwise,
	// the default is 'Reconciling' where resources get actuated.
	//+kubebuilder:validation:Enum=Reconciling;Paused
	//+kubebuilder:validation:Optional
	Actuation ActuationMode `json:"actuationMode,omitempty"`

	// StateIntoSpec is the user override of the default value for the
	// 'cnrm.cloud.google.com/state-into-spec' annotation if the annotation is
	// unset for a resource.
	// If the field is set in both the ConfigConnector object and the
	// ConfigConnectorContext object is in the namespaced mode, then the value
	// in the ConfigConnectorContext object will be used.
	// 'Absent' means that unspecified fields in the resource spec stay
	// unspecified after successful reconciliation.
	// 'Merge' means that unspecified fields in the resource spec are populated
	// after a successful reconciliation if those unspecified fields are
	// computed/defaulted by the API. It is only applicable to resources
	// supporting the 'Merge' option.
	//+kubebuilder:validation:Enum=Absent;Merge
	//+kubebuilder:validation:Optional
	StateIntoSpec *StateIntoSpecValue `json:"stateIntoSpec,omitempty"`

	// Experiments are a list of strings that allow opting-in to previews/experiments.
	// This functionality is generally not supported and not expected to work going forwards,
	// the intention is to allow new functionality to be delivered to non-production environments quickly,
	// proving the functionality before it is baked into the supported API.
	Experiments []string `json:"experiments,omitempty"`
}

// ConfigConnectorStatus defines the observed state of ConfigConnector
type ConfigConnectorStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=configconnectors,scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Healthy",type=string,JSONPath=".status.healthy", description="When 'true' the most recent reconcile of the ConfigConnector object succeeded"

// ConfigConnector is the Schema for the configconnectors API
type ConfigConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigConnectorSpec   `json:"spec,omitempty"`
	Status ConfigConnectorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConnectorList contains a list of ConfigConnector
type ConfigConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigConnector{}, &ConfigConnectorList{})
}

func (c *ConfigConnector) ComponentName() string {
	return "configconnector"
}

func (c *ConfigConnector) CommonSpec() addonv1alpha1.CommonSpec {
	return c.Spec.CommonSpec
}

func (c *ConfigConnector) GetCommonStatus() addonv1alpha1.CommonStatus {
	return c.Status.CommonStatus
}

func (c *ConfigConnector) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func (c *ConfigConnector) GetMode() string {
	if c.Spec.Mode == "" {
		return k8s.NamespacedMode
	}
	return c.Spec.Mode
}
