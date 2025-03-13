package apis

type APIArgs struct {
	Group           string
	Version         string
	Kind            string
	ProtoResource   string
	PackageProtoTag string
	KindProtoTag    string

	// ProtoMessageName is the last component of the proto message name, e.g. for google.cloud.v1.Foo, it will be "Foo"
	ProtoMessageName string
	// ProtoMessageFullName is the fully qualified proto message name, e.g. google.cloud.v1.Foo
	ProtoMessageFullName string
}

const TypesTemplate = `
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

package {{ .Version }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var {{ .Kind }}GVK = GroupVersion.WithKind("{{ .Kind }}")

// {{ .Kind }}Spec defines the desired state of {{ .Kind }}
{{- if .KindProtoTag }}
// +kcc:proto={{ .KindProtoTag }}
{{- end }}
type {{ .Kind }}Spec struct {
	// The {{ .Kind }} name. If not given, the metadata.name will be used.
	ResourceID *string ` + "`" + `json:"resourceID,omitempty"` + "`" + `
}

// {{ .Kind }}Status defines the config connector machine state of {{ .Kind }}
type {{ .Kind }}Status struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition ` + "`" + `json:"conditions,omitempty"` + "`" + ` 

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 ` + "`" + `json:"observedGeneration,omitempty"` + "`" + `

	// A unique specifier for the {{ .Kind }} resource in GCP.
	ExternalRef *string ` + "`" + `json:"externalRef,omitempty"` + "`" + `

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *{{ .Kind }}ObservedState ` + "`" + `json:"observedState,omitempty"` + "`" + `
}

// {{ .Kind }}ObservedState is the state of the {{ .Kind }} resource as most recently observed in GCP.
{{- if .KindProtoTag }}
// +kcc:proto={{ .KindProtoTag }}
{{- end }}
type {{ .Kind }}ObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralization below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcp{{ .Kind | ToLower }};gcp{{ .Kind | ToLower }}s
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// {{ .Kind }} is the Schema for the {{ .Kind }} API
// +k8s:openapi-gen=true
type {{ .Kind }} struct {
	metav1.TypeMeta   ` + "`" + `json:",inline"` + "`" + `
	metav1.ObjectMeta ` + "`" + `json:"metadata,omitempty"` + "`" + `

	// +required
	Spec   {{ .Kind }}Spec   ` + "`" + `json:"spec,omitempty"` + "`" + `
	Status {{ .Kind }}Status ` + "`" + `json:"status,omitempty"` + "`" + `
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// {{ .Kind }}List contains a list of {{ .Kind }}
type {{ .Kind }}List struct {
	metav1.TypeMeta ` + "`" + `json:",inline"` + "`" + `
	metav1.ListMeta ` + "`" + `json:"metadata,omitempty"` + "`" + `
	Items           []{{ .Kind }} ` + "`" + `json:"items"` + "`" + `
}

func init() {
	SchemeBuilder.Register(&{{ .Kind }}{}, &{{ .Kind }}List{})
}
`
