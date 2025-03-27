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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var DataprocWorkflowTemplateGVK = GroupVersion.WithKind("DataprocWorkflowTemplate")

// DataprocWorkflowTemplateSpec defines the desired state of DataprocWorkflowTemplate
// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplate
type DataprocWorkflowTemplateSpec struct {
	// The project that this resource belongs to.
	ProjectRef refsv1beta1.Reference `json:"projectRef"`

	// The location for the resource
	Location string `json:"location"`
	// The DataprocWorkflowTemplate name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.id
	ID *string `json:"id,omitempty"`

	// Optional. Used to perform a consistent read-modify-write.
	//
	//  This field should be left blank for a `CreateWorkflowTemplate` request. It
	//  is required for an `UpdateWorkflowTemplate` request, and must match the
	//  current server version. A typical update template flow would fetch the
	//  current template with a `GetWorkflowTemplate` request, which will return
	//  the current template with the `version` field filled in with the
	//  current server version. The user updates other fields in the template,
	//  then returns it as part of the `UpdateWorkflowTemplate` request.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.version
	Version *int32 `json:"version,omitempty"`

	// Optional. The labels to associate with this template. These labels
	//  will be propagated to all jobs and clusters created by the workflow
	//  instance.
	//
	//  Label **keys** must contain 1 to 63 characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	//  Label **values** may be empty, but, if present, must contain 1 to 63
	//  characters, and must conform to
	//  [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	//
	//  No more than 32 labels can be associated with a template.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. WorkflowTemplate scheduling information.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.placement
	Placement WorkflowTemplatePlacement `json:"placement"`

	// Required. The Directed Acyclic Graph of Jobs to submit.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.jobs
	Jobs []OrderedJob `json:"jobs,omitempty"`

	// Optional. Template parameters whose values are substituted into the
	//  template. Values for parameters must be provided when the template is
	//  instantiated.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.parameters
	Parameters []TemplateParameter `json:"parameters,omitempty"`

	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see
	//  [JSON representation of
	//  duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
	//  The timeout duration must be from 10 minutes ("600s") to 24 hours
	//  ("86400s"). The timer begins when the first job is submitted. If the
	//  workflow is running at the end of the timeout period, any remaining jobs
	//  are cancelled, the workflow is ended, and if the workflow was running on a
	//  [managed
	//  cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	//  the cluster is deleted.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.dag_timeout
	DagTimeout *string `json:"dagTimeout,omitempty"`

	// Optional. Encryption settings for encrypting workflow template job
	//  arguments.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.encryption_config
	EncryptionConfig *WorkflowTemplate_EncryptionConfig `json:"encryptionConfig,omitempty"`
}

type WorkflowTemplatePlacement struct {
	// A cluster that is managed by the workflow.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplatePlacement.managed_cluster
	ManagedCluster *ManagedCluster `json:"managedCluster,omitempty"`

	// Optional. A selector that chooses target cluster for jobs based
	//  on metadata.
	//
	//  The selector is evaluated at the time each job is submitted.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplatePlacement.cluster_selector
	ClusterSelector *ClusterSelector `json:"clusterSelector,omitempty"`
}

type ManagedCluster struct {
	// Required. The cluster name prefix. A unique cluster name will be formed by
	//  appending a random suffix.
	//
	//  The name must contain only lower-case letters (a-z), numbers (0-9),
	//  and hyphens (-). Must begin with a letter. Cannot begin or end with
	//  hyphen. Must consist of between 2 and 35 characters.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Required. The cluster configuration.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.config
	Config *ClusterConfig `json:"config,omitempty"`

	// Optional. The labels to associate with this cluster.
	//
	//  Label keys must be between 1 and 63 characters long, and must conform to
	//  the following PCRE regular expression:
	//  [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	//  Label values must be between 1 and 63 characters long, and must conform to
	//  the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	//  No more than 32 labels can be associated with a given cluster.
	// +kcc:proto:field=google.cloud.dataproc.v1.ManagedCluster.labels
	Labels map[string]string `json:"labels,omitempty"`
}

type ClusterSelector struct {
	// Optional. The zone where workflow process executes. This parameter does not
	//  affect the selection of the cluster.
	//
	//  If unspecified, the zone of the first cluster matching the selector
	//  is used.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterSelector.zone
	Zone *string `json:"zone,omitempty"`

	// Required. The cluster labels. Cluster must have all labels
	//  to match.
	// +kcc:proto:field=google.cloud.dataproc.v1.ClusterSelector.cluster_labels
	ClusterLabels map[string]string `json:"clusterLabels,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplate_EncryptionConfig) DeepCopyInto(out *WorkflowTemplate_EncryptionConfig) {
	*out = *in
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	return
}

// DataprocWorkflowTemplateStatus defines the config connector machine state of DataprocWorkflowTemplate
type DataprocWorkflowTemplateStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocWorkflowTemplate resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataprocWorkflowTemplateObservedState `json:"observedState,omitempty"`
}

// DataprocWorkflowTemplateObservedState is the state of the DataprocWorkflowTemplate resource as most recently observed in GCP.
// +kcc:proto=google.cloud.dataproc.v1.WorkflowTemplate
type DataprocWorkflowTemplateObservedState struct {
	// Output only. The resource name of the workflow template, as described
	//  in https://cloud.google.com/apis/design/resource_names.
	//
	//  * For `projects.regions.workflowTemplates`, the resource name of the
	//    template has the following format:
	//    `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}`
	//
	//  * For `projects.locations.workflowTemplates`, the resource name of the
	//    template has the following format:
	//    `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.name
	Name *string `json:"name,omitempty"`

	// Output only. The time template was created.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time template was last updated.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. WorkflowTemplate scheduling information.
	// +kcc:proto:field=google.cloud.dataproc.v1.WorkflowTemplate.placement
	Placement *WorkflowTemplatePlacementObservedState `json:"placement,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocworkflowtemplate;gcpdataprocworkflowtemplates
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocWorkflowTemplate is the Schema for the DataprocWorkflowTemplate API
// +k8s:openapi-gen=true
type DataprocWorkflowTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocWorkflowTemplateSpec   `json:"spec,omitempty"`
	Status DataprocWorkflowTemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocWorkflowTemplateList contains a list of DataprocWorkflowTemplate
type DataprocWorkflowTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocWorkflowTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocWorkflowTemplate{}, &DataprocWorkflowTemplateList{})
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataprocWorkflowTemplate) DeepCopyInto(out *DataprocWorkflowTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataprocWorkflowTemplate.
func (in *DataprocWorkflowTemplate) DeepCopy() *DataprocWorkflowTemplate {
	if in == nil {
		return nil
	}
	out := new(DataprocWorkflowTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataprocWorkflowTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataprocWorkflowTemplateList) DeepCopyInto(out *DataprocWorkflowTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataprocWorkflowTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataprocWorkflowTemplateList.
func (in *DataprocWorkflowTemplateList) DeepCopy() *DataprocWorkflowTemplateList {
	if in == nil {
		return nil
	}
	out := new(DataprocWorkflowTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataprocWorkflowTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataprocWorkflowTemplateSpec) DeepCopyInto(out *DataprocWorkflowTemplateSpec) {
	*out = *in
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(int32)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]OrderedJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]TemplateParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DagTimeout != nil {
		in, out := &in.DagTimeout, &out.DagTimeout
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfig != nil {
		in, out := &in.EncryptionConfig, &out.EncryptionConfig
		*out = new(WorkflowTemplate_EncryptionConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataprocWorkflowTemplateSpec.
func (in *DataprocWorkflowTemplateSpec) DeepCopy() *DataprocWorkflowTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(DataprocWorkflowTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataprocWorkflowTemplateStatus) DeepCopyInto(out *DataprocWorkflowTemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(DataprocWorkflowTemplateObservedState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataprocWorkflowTemplateStatus.
func (in *DataprocWorkflowTemplateStatus) DeepCopy() *DataprocWorkflowTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(DataprocWorkflowTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataprocWorkflowTemplateObservedState) DeepCopyInto(out *DataprocWorkflowTemplateObservedState) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	in.Placement.DeepCopyInto(&out.Placement)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataprocWorkflowTemplateObservedState.
func (in *DataprocWorkflowTemplateObservedState) DeepCopy() *DataprocWorkflowTemplateObservedState {
	if in == nil {
		return nil
	}
	out := new(DataprocWorkflowTemplateObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTemplatePlacement) DeepCopyInto(out *WorkflowTemplatePlacement) {
	*out = *in
	if in.ManagedCluster != nil {
		in, out := &in.ManagedCluster, &out.ManagedCluster
		*out = new(ManagedCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(ClusterSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}
