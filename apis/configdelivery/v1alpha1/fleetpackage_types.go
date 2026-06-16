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

package v1alpha1

import (
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConfigDeliveryFleetPackageGVK = GroupVersion.WithKind("ConfigDeliveryFleetPackage")

// ConfigDeliveryFleetPackageSpec defines the desired state of ConfigDeliveryFleetPackage
// +kcc:spec:proto=google.cloud.configdelivery.v1.FleetPackage
type ConfigDeliveryFleetPackageSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The ConfigDeliveryFleetPackage name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Information specifying the source of kubernetes configuration to deploy.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.resource_bundle_selector
	ResourceBundleSelector *FleetPackage_ResourceBundleSelector `json:"resourceBundleSelector,omitempty"`

	// Optional. Configuration to select target clusters to deploy kubernetes configuration to.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.target
	Target *FleetPackage_Target `json:"target,omitempty"`

	// Optional. The strategy to use to deploy kubernetes configuration to clusters.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.rollout_strategy
	RolloutStrategy *RolloutStrategy `json:"rolloutStrategy,omitempty"`

	// Required. Information specifying how to map a `ResourceBundle` variant to a target cluster.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.variant_selector
	VariantSelector *FleetPackage_VariantSelector `json:"variantSelector,omitempty"`

	// Optional. Information around how to handle kubernetes resources at the target clusters when the `FleetPackage` is deleted.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.deletion_propagation_policy
	DeletionPropagationPolicy *string `json:"deletionPropagationPolicy,omitempty"`

	// Optional. The desired state of the fleet package.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.state
	State *string `json:"state,omitempty"`
}

// ConfigDeliveryFleetPackageStatus defines the config connector machine state of ConfigDeliveryFleetPackage
type ConfigDeliveryFleetPackageStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConfigDeliveryFleetPackage resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *ConfigDeliveryFleetPackageObservedState `json:"observedState,omitempty"`
}

// ConfigDeliveryFleetPackageObservedState is the state of the ConfigDeliveryFleetPackage resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.configdelivery.v1.FleetPackage
type ConfigDeliveryFleetPackageObservedState struct {
	// Output only. Time at which the `FleetPackage` was created.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `FleetPackage` was updated.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Information containing the rollout status of the `FleetPackage` across all the target clusters.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.info
	Info *FleetPackageInfoObservedState `json:"info,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconfigdeliveryfleetpackage;gcpconfigdeliveryfleetpackages
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConfigDeliveryFleetPackage is the Schema for the ConfigDeliveryFleetPackage API
// +k8s:openapi-gen=true
type ConfigDeliveryFleetPackage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConfigDeliveryFleetPackageSpec   `json:"spec,omitempty"`
	Status ConfigDeliveryFleetPackageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConfigDeliveryFleetPackageList contains a list of ConfigDeliveryFleetPackage
type ConfigDeliveryFleetPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeliveryFleetPackage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigDeliveryFleetPackage{}, &ConfigDeliveryFleetPackageList{})
}

// +kcc:proto=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository
type FleetPackage_CloudBuildRepository struct {
	// Optional. variants_pattern is a glob pattern that will be used to find
	//  variants in the repository. Examples: `variants/*.yaml`, `us-*`
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository.variants_pattern
	VariantsPattern *string `json:"variantsPattern,omitempty"`

	// Required. Name of the cloud build repository.
	//  Format is projects/{p}/locations/{l}/connections/{c}/repositories/{r}.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository.name
	RepositoryRef *cloudbuildv1beta1.RepositoryRef `json:"repositoryRef,omitempty"`

	// Optional. path to the directory or file within the repository that
	//  contains the kubernetes configuration. If unspecified, path is assumed to
	//  the top level root directory of the repository.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository.path
	Path *string `json:"path,omitempty"`

	// Required. git tag of the underlying git repository.
	//  The git tag must be in the semantic version format `vX.Y.Z`.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository.tag
	Tag *string `json:"tag,omitempty"`

	// Required. Google service account to use in CloudBuild triggers to fetch
	//  and store kubernetes configuration.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.CloudBuildRepository.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

// +kcc:proto=google.cloud.configdelivery.v1.FleetPackage.ResourceBundleTag
type FleetPackage_ResourceBundleTag struct {
	// Required. Name of the `ResourceBundle`.
	//  Format is projects/{p}/locations/{l}/resourceBundles/{r}.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.ResourceBundleTag.name
	ResourceBundleRef *ConfigDeliveryResourceBundleRef `json:"resourceBundleRef,omitempty"`

	// Required. Tag refers to a version of the release in a `ResourceBundle`.
	//  This is a Git tag in the semantic version format `vX.Y.Z`.
	// +kcc:proto:field=google.cloud.configdelivery.v1.FleetPackage.ResourceBundleTag.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.configdelivery.v1.Fleet
type Fleet struct {
	// Required. The host project for the GKE fleet. Format is
	//  `projects/{project}`.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Fleet.project
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// Optional. selector allows targeting a subset of fleet members using their
	//  labels.
	// +kcc:proto:field=google.cloud.configdelivery.v1.Fleet.selector
	Selector *Fleet_LabelSelector `json:"selector,omitempty"`
}
