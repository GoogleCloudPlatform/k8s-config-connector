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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme

	GKEHubFeatureMembershipGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "GKEHubFeatureMembership",
	}
)

type FeaturemembershipBinauthz struct {
	/* Whether binauthz is enabled in this cluster. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type FeaturemembershipConfigSync struct {
	// +optional
	Git *FeaturemembershipGit `json:"git,omitempty"`

	// +optional
	MetricsGcpServiceAccountRef *refs.MetricsGcpServiceAccountRef `json:"metricsGcpServiceAccountRef,omitempty"`

	// +optional
	Oci *FeaturemembershipOci `json:"oci,omitempty"`

	/* Set to true to enable the Config Sync admission webhook to prevent drifts. If set to `false`, disables the Config Sync admission webhook and does not prevent drifts. */
	// +optional
	PreventDrift *bool `json:"preventDrift,omitempty"`

	/* Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode. */
	// +optional
	SourceFormat *string `json:"sourceFormat,omitempty"`
}

type FeaturemembershipConfigmanagement struct {
	/* **DEPRECATED** Binauthz configuration for the cluster. This field will be ignored and should not be set. */
	// +optional
	Binauthz *FeaturemembershipBinauthz `json:"binauthz,omitempty"`

	/* Config Sync configuration for the cluster. */
	// +optional
	ConfigSync *FeaturemembershipConfigSync `json:"configSync,omitempty"`

	/* Hierarchy Controller is no longer available. Use https://github.com/kubernetes-sigs/hierarchical-namespaces instead. */
	// +optional
	HierarchyController *FeaturemembershipHierarchyController `json:"hierarchyController,omitempty"`

	/* **DEPRECATED** Configuring Policy Controller through the configmanagement feature is no longer recommended. Use the policycontroller feature instead. */
	// +optional
	PolicyController *FeaturemembershipPolicyController `json:"policyController,omitempty"`

	/* Optional. Version of ACM to install. Defaults to the latest version. */
	// +optional
	Version *string `json:"version,omitempty"`

	/* Optional. Whether to automatically manage the configmanagement Feature. There are 3 accepted values. MANAGEMENT_UNSPECIFIED means that the mamangement mode is unspecified. MANAGEMENT_AUTOMATIC means that Google manages the Feature for the cluster. MANAGEMENT_MANUAL means that users should manage the Feature for the cluster. */
	// +optional
	Management *string `json:"management,omitempty"`
}

type FeaturemembershipGit struct {
	// +optional
	GcpServiceAccountRef *refs.IAMServiceAccountRef `json:"gcpServiceAccountRef,omitempty"`

	/* URL for the HTTPS proxy to be used when communicating with the Git repo. */
	// +optional
	HttpsProxy *string `json:"httpsProxy,omitempty"`

	/* The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository. */
	// +optional
	PolicyDir *string `json:"policyDir,omitempty"`

	/* Type of secret configured for access to the Git repo. Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive. */
	// +optional
	SecretType *string `json:"secretType,omitempty"`

	/* The branch of the repository to sync from. Default: master. */
	// +optional
	SyncBranch *string `json:"syncBranch,omitempty"`

	/* The URL of the Git repository to use as the source of truth. */
	// +optional
	SyncRepo *string `json:"syncRepo,omitempty"`

	/* Git revision (tag or hash) to check out. Default HEAD. */
	// +optional
	SyncRev *string `json:"syncRev,omitempty"`

	/* Period in seconds between consecutive syncs. Default: 15. */
	// +optional
	SyncWaitSecs *string `json:"syncWaitSecs,omitempty"`
}

type FeaturemembershipHierarchyController struct {
	/* Whether hierarchical resource quota is enabled in this cluster. */
	// +optional
	EnableHierarchicalResourceQuota *bool `json:"enableHierarchicalResourceQuota,omitempty"`

	/* Whether pod tree labels are enabled in this cluster. */
	// +optional
	EnablePodTreeLabels *bool `json:"enablePodTreeLabels,omitempty"`

	/* Whether Hierarchy Controller is enabled in this cluster. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
}

type FeaturemembershipMesh struct {
	/* **DEPRECATED** Whether to automatically manage Service Mesh control planes. Possible values: CONTROL_PLANE_MANAGEMENT_UNSPECIFIED, AUTOMATIC, MANUAL */
	// +optional
	ControlPlane *string `json:"controlPlane,omitempty"`

	/* Whether to automatically manage Service Mesh. Possible values: MANAGEMENT_UNSPECIFIED, MANAGEMENT_AUTOMATIC, MANAGEMENT_MANUAL */
	// +optional
	Management *string `json:"management,omitempty"`
}

type FeaturemembershipMonitoring struct {
	/* Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export. */
	// +optional
	Backends []string `json:"backends,omitempty"`
}

type FeaturemembershipOci struct {
	// +optional
	GcpServiceAccountRef *refs.IAMServiceAccountRef `json:"gcpServiceAccountRef,omitempty"`

	/* The absolute path of the directory that contains the local resources. Default: the root directory of the image. */
	// +optional
	PolicyDir *string `json:"policyDir,omitempty"`

	/* Type of secret configured for access to the OCI Image. Must be one of gcenode, gcpserviceaccount or none. The validation of this is case-sensitive. */
	// +optional
	SecretType *string `json:"secretType,omitempty"`

	/* The OCI image repository URL for the package to sync from. e.g. LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME. */
	// +optional
	SyncRepo *string `json:"syncRepo,omitempty"`

	/* Period in seconds(int64 format) between consecutive syncs. Default: 15. */
	// +optional
	SyncWaitSecs *string `json:"syncWaitSecs,omitempty"`
}

type FeaturemembershipPolicyContent struct {
	/* Configures the installation of the Template Library. */
	// +optional
	TemplateLibrary *FeaturemembershipTemplateLibrary `json:"templateLibrary,omitempty"`
}

type FeaturemembershipPolicyController struct {
	/* Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether. */
	// +optional
	AuditIntervalSeconds *string `json:"auditIntervalSeconds,omitempty"`

	/* Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster. */
	// +optional
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	/* Logs all denies and dry run failures. */
	// +optional
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`

	/* Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: ["cloudmonitoring", "prometheus"]. Default: ["cloudmonitoring", "prometheus"] */
	// +optional
	Monitoring *FeaturemembershipMonitoring `json:"monitoring,omitempty"`

	/* Enable or disable mutation in policy controller. If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster. */
	// +optional
	MutationEnabled *bool `json:"mutationEnabled,omitempty"`

	/* Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated. */
	// +optional
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`

	/* Installs the default template library along with Policy Controller. */
	// +optional
	TemplateLibraryInstalled *bool `json:"templateLibraryInstalled,omitempty"`
}

type FeaturemembershipPolicyControllerHubConfig struct {
	/* Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether. */
	// +optional
	AuditIntervalSeconds *int64 `json:"auditIntervalSeconds,omitempty"`

	/* The maximum number of audit violations to be stored in a constraint. If not set, the internal default of 20 will be used. */
	// +optional
	ConstraintViolationLimit *int64 `json:"constraintViolationLimit,omitempty"`

	/* The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster. */
	// +optional
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	/* Configures the mode of the Policy Controller installation. Possible values: INSTALL_SPEC_UNSPECIFIED, INSTALL_SPEC_NOT_INSTALLED, INSTALL_SPEC_ENABLED, INSTALL_SPEC_SUSPENDED, INSTALL_SPEC_DETACHED */
	// +optional
	InstallSpec *string `json:"installSpec,omitempty"`

	/* Logs all denies and dry run failures. */
	// +optional
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`

	/* Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: ["cloudmonitoring", "prometheus"]. Default: ["cloudmonitoring", "prometheus"] */
	// +optional
	Monitoring *FeaturemembershipMonitoring `json:"monitoring,omitempty"`

	/* Enables the ability to mutate resources using Policy Controller. */
	// +optional
	MutationEnabled *bool `json:"mutationEnabled,omitempty"`

	/* Specifies the desired policy content on the cluster. */
	// +optional
	PolicyContent *FeaturemembershipPolicyContent `json:"policyContent,omitempty"`

	/* Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated. */
	// +optional
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`
}

type FeaturemembershipPolicycontroller struct {
	/* Policy Controller configuration for the cluster. */
	PolicyControllerHubConfig FeaturemembershipPolicyControllerHubConfig `json:"policyControllerHubConfig"`

	/* Optional. Version of Policy Controller to install. Defaults to the latest version. */
	// +optional
	Version *string `json:"version,omitempty"`
}

type FeaturemembershipTemplateLibrary struct {
	/* Configures the manner in which the template library is installed on the cluster. Possible values: INSTALLATION_UNSPECIFIED, NOT_INSTALLED, ALL */
	// +optional
	Installation *string `json:"installation,omitempty"`
}

type GKEHubFeatureMembershipSpec struct {
	/* Config Management-specific spec. */
	// +optional
	Configmanagement *FeaturemembershipConfigmanagement `json:"configmanagement,omitempty"`

	/* Immutable. */
	FeatureRef FeatureRef `json:"featureRef"`

	/* Immutable. The location of the feature */
	Location string `json:"location"`

	/* Immutable. The location of the membership */
	// +optional
	MembershipLocation *string `json:"membershipLocation,omitempty"`

	/* Immutable. */
	MembershipRef MembershipRef `json:"membershipRef"`

	/* Manage Mesh Features */
	// +optional
	Mesh *FeaturemembershipMesh `json:"mesh,omitempty"`

	/* Policy Controller-specific spec. */
	// +optional
	Policycontroller *FeaturemembershipPolicycontroller `json:"policycontroller,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef FeatureProjectRef `json:"projectRef"`
}

type GKEHubFeatureMembershipStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubFeatureMembership's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubfeaturemembership;gcpgkehubfeaturememberships
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubFeatureMembership is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubFeatureMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   GKEHubFeatureMembershipSpec   `json:"spec,omitempty"`
	Status GKEHubFeatureMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubFeatureMembershipList contains a list of GKEHubFeatureMembership
type GKEHubFeatureMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubFeatureMembership `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubFeatureMembership{}, &GKEHubFeatureMembershipList{})
}
