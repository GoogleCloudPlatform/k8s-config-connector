package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kcc "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	k8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

// SecureSourceManagerInstanceSpec defines the desired state of SecureSourceManagerInstance.
type SecureSourceManagerInstanceSpec struct {

	// The region where the instance will be deployed.
	Location string `json:"location,omitempty"`

	//  Immutable. Customer-managed encryption key name
	KmsKeyRef *kcc.ResourceReference `json:"kmsKeyRef,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef kcc.ResourceReference `json:"projectRef"`

	ResourceID string `json:"resourceID,omitempty"`
}

// SecureSourceManagerInstanceStatus defines the observed state of SecureSourceManagerInstance.
type SecureSourceManagerInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringDashboard's current state. */
	Conditions []k8s.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// Output only. Current state of the instance.
	State string `json:"state,omitempty"`

	// Output only. An optional field providing information about the current
	// instance state.
	StateNote string `json:"stateNote,omitempty"`

	// Output only. A list of hostnames for this instance.
	HostConfig *SecureSourceManagerInstance_HostConfig `json:"hostConfig,omitempty"`
}

type SecureSourceManagerInstance_HostConfig struct {
	// Output only. HTML hostname.
	Html string `json:"html,omitempty"`
	// Output only. API hostname. This is the hostname to use for **Host: Data
	// Plane** endpoints.
	Api string `json:"api,omitempty"`
	// Output only. Git HTTP hostname.
	GitHttp string `json:"gitHttp,omitempty"`
	// Output only. Git SSH hostname.
	GitSsh string `json:"gitSsh,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=securesourcemanagerinstances,scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// SecureSourceManagerInstance is the Schema for the securesourcemanagerinstances API.
type SecureSourceManagerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecureSourceManagerInstanceSpec   `json:"spec,omitempty"`
	Status SecureSourceManagerInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecureSourceManagerInstanceList contains a list of SecureSourceManagerInstance.
type SecureSourceManagerInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerInstance{}, &SecureSourceManagerInstanceList{})
}
