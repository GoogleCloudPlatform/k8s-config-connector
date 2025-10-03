// Copyright 2025 Google LLC
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

// +generated:types
// krm.group: gkehub.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.gkehub.v1beta
// resource: GKEHubFeatureMembership:MembershipFeatureSpec

package v1beta1

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.BinauthzConfig
type BinauthzConfig struct {
	// Whether binauthz is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.BinauthzConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ConfigSync
type ConfigSync struct {
	// Optional. Git repo configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.git
	Git *GitConfig `json:"git,omitempty"`

	// Optional. Specifies whether the Config Sync Repo is
	//  in "hierarchical" or "unstructured" mode.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.source_format
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// Optional. Enables the installation of ConfigSync.
	//  If set to true, ConfigSync resources will be created and the other
	//  ConfigSync fields will be applied if exist.
	//  If set to false, all other ConfigSync fields will be ignored, ConfigSync
	//  resources will be deleted.
	//  If omitted, ConfigSync resources will be managed depends on the presence
	//  of the git or oci field.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Optional. Set to true to enable the Config Sync admission webhook to
	//  prevent drifts. If set to `false`, disables the Config Sync admission
	//  webhook and does not prevent drifts.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.prevent_drift
	PreventDrift *bool `json:"preventDrift,omitempty"`

	// Optional. OCI repo configuration for the cluster
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.oci
	Oci *OciConfig `json:"oci,omitempty"`

	// Optional. Set to true to stop syncing configs for a single cluster.
	//  Default to false.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSync.stop_syncing
	StopSyncing *bool `json:"stopSyncing,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.GitConfig
type GitConfig struct {
	// Required. The URL of the Git repository to use as the source of truth.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.sync_repo
	SyncRepo *string `json:"syncRepo,omitempty"`

	// Optional. The branch of the repository to sync from. Default: master.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.sync_branch
	SyncBranch *string `json:"syncBranch,omitempty"`

	// Optional. The path within the Git repository that represents the top level
	//  of the repo to sync. Default: the root directory of the repository.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.policy_dir
	PolicyDir *string `json:"policyDir,omitempty"`

	// Optional. Period in seconds between consecutive syncs. Default: 15.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.sync_wait_secs
	SyncWaitSecs *int64 `json:"syncWaitSecs,omitempty"`

	// Optional. Git revision (tag or hash) to check out. Default HEAD.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.sync_rev
	SyncRev *string `json:"syncRev,omitempty"`

	// Required. Type of secret configured for access to the Git repo.
	//  Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount,
	//  githubapp or none.
	//  The validation of this is case-sensitive.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.secret_type
	SecretType *string `json:"secretType,omitempty"`

	// Optional. URL for the HTTPS proxy to be used when communicating with the
	//  Git repo.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.https_proxy
	HTTPSProxy *string `json:"httpsProxy,omitempty"`

	// Optional. The Google Cloud Service Account Email used for auth when
	//  secret_type is gcpServiceAccount.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GitConfig.gcp_service_account_email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerConfig
type HierarchyControllerConfig struct {
	// Whether Hierarchy Controller is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Whether pod tree labels are enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerConfig.enable_pod_tree_labels
	EnablePodTreeLabels *bool `json:"enablePodTreeLabels,omitempty"`

	// Whether hierarchical resource quota is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerConfig.enable_hierarchical_resource_quota
	EnableHierarchicalResourceQuota *bool `json:"enableHierarchicalResourceQuota,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.OciConfig
type OciConfig struct {
	// Required. The OCI image repository URL for the package to sync from.
	//  e.g. `LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME`.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OciConfig.sync_repo
	SyncRepo *string `json:"syncRepo,omitempty"`

	// Optional. The absolute path of the directory that contains
	//  the local resources.  Default: the root directory of the image.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OciConfig.policy_dir
	PolicyDir *string `json:"policyDir,omitempty"`

	// Optional. Period in seconds between consecutive syncs. Default: 15.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OciConfig.sync_wait_secs
	SyncWaitSecs *int64 `json:"syncWaitSecs,omitempty"`

	// Required. Type of secret configured for access to the OCI repo.
	//  Must be one of gcenode, gcpserviceaccount, k8sserviceaccount or none.
	//  The validation of this is case-sensitive.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OciConfig.secret_type
	SecretType *string `json:"secretType,omitempty"`

	// Optional. The Google Cloud Service Account Email used for auth when
	//  secret_type is gcpServiceAccount.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OciConfig.gcp_service_account_email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyController
type PolicyController struct {
	// Enables the installation of Policy Controller.
	//  If false, the rest of PolicyController fields take no
	//  effect.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Installs the default template library along with Policy Controller.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.template_library_installed
	TemplateLibraryInstalled *bool `json:"templateLibraryInstalled,omitempty"`

	// Sets the interval for Policy Controller Audit Scans (in seconds).
	//  When set to 0, this disables audit functionality altogether.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.audit_interval_seconds
	AuditIntervalSeconds *int64 `json:"auditIntervalSeconds,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks.
	//  Namespaces do not need to currently exist on the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.exemptable_namespaces
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects
	//  other than the object currently being evaluated.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.referential_rules_enabled
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`

	// Logs all denies and dry run failures.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.log_denies_enabled
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`

	// Enable or disable mutation in policy controller.
	//  If true, mutation CRDs, webhook and controller deployment
	//  will be deployed to the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.mutation_enabled
	MutationEnabled *bool `json:"mutationEnabled,omitempty"`

	// Monitoring specifies the configuration of monitoring.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.monitoring
	Monitoring *PolicyControllerMonitoring `json:"monitoring,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMonitoring
type PolicyControllerMonitoring struct {
	// Specifies the list of backends Policy Controller will export to.
	//  An empty list would effectively disable metrics export.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMonitoring.backends
	Backends []string `json:"backends,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.BundleInstallSpec
type BundleInstallSpec struct {
	// The set of namespaces to be exempted from the bundle.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.BundleInstallSpec.exempted_namespaces
	ExemptedNamespaces []string `json:"exemptedNamespaces,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.MonitoringConfig
type MonitoringConfig struct {
	// Specifies the list of backends Policy Controller will export to.
	//  An empty list would effectively disable metrics export.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MonitoringConfig.backends
	Backends []string `json:"backends,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.PolicyContentSpec
type PolicyContentSpec struct {

	// TODO: unsupported map type with key string and value message

	// Configures the installation of the Template Library.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyContentSpec.template_library
	TemplateLibrary *TemplateLibraryConfig `json:"templateLibrary,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.Toleration
type PolicyControllerDeploymentConfig_Toleration struct {
	// Matches a taint key (not necessarily unique).
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.Toleration.key
	Key *string `json:"key,omitempty"`

	// Matches a taint operator.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.Toleration.operator
	Operator *string `json:"operator,omitempty"`

	// Matches a taint value.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.Toleration.value
	Value *string `json:"value,omitempty"`

	// Matches a taint effect.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.Toleration.effect
	Effect *string `json:"effect,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.ResourceList
type ResourceList struct {
	// Memory requirement expressed in Kubernetes resource units.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.ResourceList.memory
	Memory *string `json:"memory,omitempty"`

	// CPU requirement expressed in Kubernetes resource units.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.ResourceList.cpu
	CPU *string `json:"cpu,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.ResourceRequirements
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed for use by
	//  the running container.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.ResourceRequirements.limits
	Limits *ResourceList `json:"limits,omitempty"`

	// Requests describes the amount of compute resources reserved for the
	//  container by the kube-scheduler.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.ResourceRequirements.requests
	Requests *ResourceList `json:"requests,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.TemplateLibraryConfig
type TemplateLibraryConfig struct {
	// Configures the manner in which the template library is installed on the
	//  cluster.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.TemplateLibraryConfig.installation
	Installation *string `json:"installation,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.MembershipSpec
type MembershipSpec struct {
	// Deprecated: use `management` instead
	//  Enables automatic control plane management.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipSpec.control_plane
	ControlPlane *string `json:"controlPlane,omitempty"`

	// Enables automatic Service Mesh management.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipSpec.management
	Management *string `json:"management,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyController
type PolicyControllerObservedState struct {
	// Output only. Last time this membership spec was updated.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyController.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
