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

package v1beta1

import (
	billingv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

var ProjectGVK = GroupVersion.WithKind("Project")

// ProjectSpec defines the desired state of Project
// +kcc:spec:proto=google.cloud.resourcemanager.v3.Project
type ProjectSpec struct {
	/* The billing account of the project. */
	BillingAccountRef *billingv1alpha1.BillingAccountRef `json:"billingAccountRef,omitempty"`

	/* The folder that this resource belongs to. Changing this forces the
	   resource to be migrated to the newly specified folder. Only one of
	   folderRef or organizationRef may be specified. */
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.parent
	FolderRef *refsv1beta1.FolderRef `json:"folderRef,omitempty"`

	/* The display name of the project. */
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.display_name
	DisplayName *string `json:"name"`

	/* The organization that this resource belongs to. Changing this
	   forces the resource to be migrated to the newly specified
	   organization. Only one of folderRef or organizationRef may be
	   specified. */
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.parent
	OrganizationRef *OrganizationRef `json:"organizationRef,omitempty"`

	/* Immutable. Optional. The projectId of the resource. Used
	   for creation and acquisition. When unset, the value of `metadata.name`
	   is used as the default. */
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.project_id
	ResourceID *string `json:"resourceID,omitempty"`
}

type OrganizationRef struct {
	// The 'name' field of an organization, when not managed by Config Connector.
	// +optional
	External string `json:"external,omitempty"`

	// The 'name' field of an 'Organization' resource.
	// +optional
	Name string `json:"name,omitempty"`

	// The 'namespace' field of an 'Organization' resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

func (r *OrganizationRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "resourcemanager.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "Organization",
	}
}

func (r *OrganizationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *OrganizationRef) GetExternal() string {
	return r.External
}

func (r *OrganizationRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

// ProjectStatus defines the config connector machine state of Project
// +kcc:observedstate:proto=google.cloud.resourcemanager.v3.Project
type ProjectStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The numeric identifier of the project. */
	// +kcc:proto:field=google.cloud.resourcemanager.v3.Project.name
	Number *string `json:"number,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpproject;gcpprojects
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// Project is the Schema for the Project API
// +k8s:openapi-gen=true
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ProjectList contains a list of Project
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
