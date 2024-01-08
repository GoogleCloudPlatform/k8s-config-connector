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
	"k8s.io/klog/v2"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigConnectorContextSpec defines the desired state of ConfigConnectorContext
type ConfigConnectorContextSpec struct {
	// The Google Service Account to be used by Config Connector to
	// authenticate with Google Cloud APIs in the associated namespace.
	GoogleServiceAccount string `json:"googleServiceAccount"`

	// Specifies which project to use for preconditions, quota, and billing for
	// requests made to Google Cloud APIs for resources in the associated
	// namespace. Must be one of 'SERVICE_ACCOUNT_PROJECT',
	// 'RESOURCE_PROJECT', or 'BILLING_PROJECT. Defaults to 'SERVICE_ACCOUNT_PROJECT'. If set to
	// 'SERVICE_ACCOUNT_PROJECT', uses the project that the Google Service
	// Account belongs to. If set to 'RESOURCE_PROJECT', uses the project that
	// the resource belongs to. If set to 'BILLING_PROJECT', uses the project specified by spec.billingProject.
	// +kubebuilder:validation:Enum=SERVICE_ACCOUNT_PROJECT;RESOURCE_PROJECT;BILLING_PROJECT
	RequestProjectPolicy string `json:"requestProjectPolicy,omitempty"`

	// Specifies the project to use for preconditions, quota and billing.
	// Should only be used when requestProjectPolicy is set to BILLING_PROJECT.
	BillingProject string `json:"billingProject,omitempty"`
}

// ConfigConnectorContextStatus defines the observed state of ConfigConnectorContext
type ConfigConnectorContextStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Healthy",type=string,JSONPath=".status.healthy", description="When 'true' the most recent reconcile of the ConfigConnectorContext object succeeded"

// ConfigConnectorContext is the Schema for the ConfigConnectorContexts API
type ConfigConnectorContext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigConnectorContextSpec   `json:"spec"`
	Status ConfigConnectorContextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConnectorContextList contains a list of ConfigConnectorContext
type ConfigConnectorContextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConnectorContext `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigConnectorContext{}, &ConfigConnectorContextList{})
}

var _ addonv1alpha1.CommonObject = &ConfigConnectorContext{}

func (c *ConfigConnectorContext) ComponentName() string {
	// This should not be called, but is needed to satisfy the CommonObject interface.
	// (We only interact with the status fields)
	klog.Fatalf("ComponentName should not be called	on ConfigConnectorContext")
	return ""
}

func (c *ConfigConnectorContext) CommonSpec() addonv1alpha1.CommonSpec {
	// This should not be called, but is needed to satisfy the CommonObject interface.
	// (We only interact with the status fields)
	klog.Fatalf("CommonSpec should not be called on ConfigConnectorContext")
	return addonv1alpha1.CommonSpec{}
}

func (c *ConfigConnectorContext) GetCommonStatus() addonv1alpha1.CommonStatus {
	return c.Status.CommonStatus
}

func (c *ConfigConnectorContext) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func (c *ConfigConnectorContext) GetRequestProjectPolicy() string {
	if c.Spec.RequestProjectPolicy == "" {
		return k8s.ServiceAccountProjectPolicy
	}
	return c.Spec.RequestProjectPolicy
}
