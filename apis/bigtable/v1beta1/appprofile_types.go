// Copyright 2025 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableAppProfileGVK = GroupVersion.WithKind("BigtableAppProfile")

type BigtableAppProfileParent struct {
	// +required
	InstanceRef *InstanceRef `json:"instanceRef,omitempty"`
}

// BigtableAppProfileSpec defines the desired state of BigtableAppProfile
// +kcc:spec:proto=google.bigtable.admin.v2.AppProfile
type BigtableAppProfileSpec struct {
	// The BigtableAppProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	BigtableAppProfileParent `json:",inline"`

	// Long form description of the use case for this AppProfile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.description
	Description *string `json:"description,omitempty"`

	// The set of clusters to route to, if using multi cluster routing. The order is ignored; clusters will be tried
	// in order of distance. If left empty, all clusters are eligible.
	MultiClusterRoutingClusterIds []string `json:"multiClusterRoutingClusterIds,omitempty"`

	// Use a multi-cluster routing policy.
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty"`

	// Use a single-cluster routing policy.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.single_cluster_routing
	SingleClusterRouting *AppProfile_SingleClusterRouting `json:"singleClusterRouting,omitempty"`

	// The standard options used for isolating this app profile's traffic from
	//  other use cases.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.standard_isolation
	StandardIsolation *AppProfile_StandardIsolation `json:"standardIsolation,omitempty"`

	// Specifies that this app profile is intended for read-only usage via the
	//  Data Boost feature. Please opt-in to this feature by setting the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.data_boost_isolation_read_only
	DataBoostIsolationReadOnly *AppProfile_DataBoostIsolationReadOnly `json:"dataBoostIsolationReadOnly,omitempty"`
}

// BigtableAppProfileStatus defines the config connector machine state of BigtableAppProfile
type BigtableAppProfileStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the BigtableAppProfile resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// The unique name of the app profile. Values are of the form
	//  `projects/{project}/instances/{instance}/appProfiles/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.name
	Name *string `json:"name,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *BigtableAppProfileObservedState `json:"observedState,omitempty"`
}

// BigtableAppProfileObservedState is the state of the BigtableAppProfile resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.bigtable.admin.v2.AppProfile
type BigtableAppProfileObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableappprofile;gcpbigtableappprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableAppProfile is the Schema for the BigtableAppProfile API
// +k8s:openapi-gen=true
type BigtableAppProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   BigtableAppProfileSpec   `json:"spec,omitempty"`
	Status BigtableAppProfileStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BigtableAppProfileList contains a list of BigtableAppProfile
type BigtableAppProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableAppProfile `json:"items"`
}

// We have to manually override AppProfile_SingleClusterRouting because of clusterId vs clusterID
// +kcc:proto=google.bigtable.admin.v2.AppProfile.SingleClusterRouting
type AppProfile_SingleClusterRouting struct {
	// The cluster to which read/write requests should be routed.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.SingleClusterRouting.cluster_id
	ClusterID *string `json:"clusterId,omitempty"`

	// Whether or not `CheckAndMutateRow` and `ReadModifyWriteRow` requests are
	//  allowed by this app profile. It is unsafe to send these requests to
	//  the same table/row/column in multiple clusters.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.SingleClusterRouting.allow_transactional_writes
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BigtableAppProfile{}, &BigtableAppProfileList{})
}
