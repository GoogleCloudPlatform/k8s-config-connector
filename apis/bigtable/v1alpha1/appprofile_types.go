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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BigtableAppProfileGVK = GroupVersion.WithKind("BigtableAppProfile")

// BigtableAppProfileSpec defines the desired state of BigtableAppProfile
// +kcc:proto=google.bigtable.admin.v2.AppProfile
type Parent struct {
	// +required
	InstanceRef *refv1beta1.ResourceRef `json:"instanceRef"`
}

// BigtableAppProfileSpec defines the desired state of BigtableAppProfile
// +kcc:proto=google.bigtable.admin.v2.AppProfile
type BigtableAppProfileSpec struct {
	// +required
	Parent Parent `json:",inline"`
	// The BigtableAppProfile name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Long form description of the use case for this AppProfile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.description
	Description *string `json:"description,omitempty"`

	// Use a multi-cluster routing policy.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.multi_cluster_routing_use_any
	MultiClusterRoutingUseAny *AppProfile_MultiClusterRoutingUseAny `json:"multiClusterRoutingUseAny,omitempty"`

	// Use a single-cluster routing policy.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.single_cluster_routing
	SingleClusterRouting *AppProfile_SingleClusterRouting `json:"singleClusterRouting,omitempty"`

	// This field has been deprecated in favor of `standard_isolation.priority`.
	//  If you set this field, `standard_isolation.priority` will be set instead.
	//
	//  The priority of requests sent using this app profile.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.priority
	Priority *string `json:"priority,omitempty"`

	// The standard options used for isolating this app profile's traffic from
	//  other use cases.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.standard_isolation
	StandardIsolation *AppProfile_StandardIsolation `json:"standardIsolation,omitempty"`

	// Specifies that this app profile is intended for read-only usage via the
	//  Data Boost feature.
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

	// Strongly validated etag for optimistic concurrency control. Preserve the
	//  value returned from `GetAppProfile` when calling `UpdateAppProfile` to
	//  fail the request if there has been a modification in the mean time. The
	//  `update_mask` of the request need not include `etag` for this protection
	//  to apply.
	//  See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and
	//  [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more
	//  details.
	// +kcc:proto:field=google.bigtable.admin.v2.AppProfile.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtableappprofile;gcpbigtableappprofiles
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
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

func init() {
	SchemeBuilder.Register(&BigtableAppProfile{}, &BigtableAppProfileList{})
}
