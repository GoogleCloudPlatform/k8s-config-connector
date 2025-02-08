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

package v1alpha1


// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ConfigSync
type ConfigSync struct {
	// Git repo configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.git
	Git *GitConfig `json:"git,omitempty"`

	// Specifies whether the Config Sync Repo is
	//  in "hierarchical" or "unstructured" mode.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.source_format
	SourceFormat *string `json:"sourceFormat,omitempty"`

	// Enables the installation of ConfigSync.
	//  If set to true, ConfigSync resources will be created and the other
	//  ConfigSync fields will be applied if exist.
	//  If set to false, all other ConfigSync fields will be ignored, ConfigSync
	//  resources will be deleted.
	//  If omitted, ConfigSync resources will be managed depends on the presence
	//  of the git or oci field.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Set to true to enable the Config Sync admission webhook to prevent drifts.
	//  If set to `false`, disables the Config Sync admission webhook and does not
	//  prevent drifts.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.prevent_drift
	PreventDrift *bool `json:"preventDrift,omitempty"`

	// OCI repo configuration for the cluster
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.oci
	Oci *OciConfig `json:"oci,omitempty"`

	// The Email of the Google Cloud Service Account (GSA) used for exporting
	//  Config Sync metrics to Cloud Monitoring when Workload Identity is enabled.
	//  The GSA should have the Monitoring Metric Writer
	//  (roles/monitoring.metricWriter) IAM role.
	//  The Kubernetes ServiceAccount `default` in the namespace
	//  `config-management-monitoring` should be bound to the GSA.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSync.metrics_gcp_service_account_email
	MetricsGcpServiceAccountEmail *string `json:"metricsGcpServiceAccountEmail,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState
type ConfigSyncDeploymentState struct {
	// Deployment state of the importer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.importer
	Importer *string `json:"importer,omitempty"`

	// Deployment state of the syncer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.syncer
	Syncer *string `json:"syncer,omitempty"`

	// Deployment state of the git-sync pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.git_sync
	GitSync *string `json:"gitSync,omitempty"`

	// Deployment state of the monitor pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.monitor
	Monitor *string `json:"monitor,omitempty"`

	// Deployment state of reconciler-manager pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.reconciler_manager
	ReconcilerManager *string `json:"reconcilerManager,omitempty"`

	// Deployment state of root-reconciler
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.root_reconciler
	RootReconciler *string `json:"rootReconciler,omitempty"`

	// Deployment state of admission-webhook
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncDeploymentState.admission_webhook
	AdmissionWebhook *string `json:"admissionWebhook,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ConfigSyncError
type ConfigSyncError struct {
	// A string representing the user facing error message
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ConfigSyncState
type ConfigSyncState struct {
	// The version of ConfigSync deployed
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.version
	Version *ConfigSyncVersion `json:"version,omitempty"`

	// Information about the deployment of ConfigSync, including the version
	//  of the various Pods deployed
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.deployment_state
	DeploymentState *ConfigSyncDeploymentState `json:"deploymentState,omitempty"`

	// The state of ConfigSync's process to sync configs to a cluster
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.sync_state
	SyncState *SyncState `json:"syncState,omitempty"`

	// Errors pertaining to the installation of Config Sync.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.errors
	Errors []ConfigSyncError `json:"errors,omitempty"`

	// The state of the RootSync CRD
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.rootsync_crd
	RootsyncCrd *string `json:"rootsyncCrd,omitempty"`

	// The state of the Reposync CRD
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.reposync_crd
	ReposyncCrd *string `json:"reposyncCrd,omitempty"`

	// The state of CS
	//  This field summarizes the other fields in this message.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncState.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion
type ConfigSyncVersion struct {
	// Version of the deployed importer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.importer
	Importer *string `json:"importer,omitempty"`

	// Version of the deployed syncer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.syncer
	Syncer *string `json:"syncer,omitempty"`

	// Version of the deployed git-sync pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.git_sync
	GitSync *string `json:"gitSync,omitempty"`

	// Version of the deployed monitor pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.monitor
	Monitor *string `json:"monitor,omitempty"`

	// Version of the deployed reconciler-manager pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.reconciler_manager
	ReconcilerManager *string `json:"reconcilerManager,omitempty"`

	// Version of the deployed reconciler container in root-reconciler pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.root_reconciler
	RootReconciler *string `json:"rootReconciler,omitempty"`

	// Version of the deployed admission_webhook pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ConfigSyncVersion.admission_webhook
	AdmissionWebhook *string `json:"admissionWebhook,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.ErrorResource
type ErrorResource struct {
	// Path in the git repo of the erroneous config
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ErrorResource.source_path
	SourcePath *string `json:"sourcePath,omitempty"`

	// Metadata name of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ErrorResource.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// Namespace of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ErrorResource.resource_namespace
	ResourceNamespace *string `json:"resourceNamespace,omitempty"`

	// Group/version/kind of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.ErrorResource.resource_gvk
	ResourceGvk *GroupVersionKind `json:"resourceGvk,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.GatekeeperDeploymentState
type GatekeeperDeploymentState struct {
	// Status of gatekeeper-controller-manager pod.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GatekeeperDeploymentState.gatekeeper_controller_manager_state
	GatekeeperControllerManagerState *string `json:"gatekeeperControllerManagerState,omitempty"`

	// Status of gatekeeper-audit deployment.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GatekeeperDeploymentState.gatekeeper_audit
	GatekeeperAudit *string `json:"gatekeeperAudit,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.GitConfig
type GitConfig struct {
	// The URL of the Git repository to use as the source of truth.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.sync_repo
	SyncRepo *string `json:"syncRepo,omitempty"`

	// The branch of the repository to sync from. Default: master.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.sync_branch
	SyncBranch *string `json:"syncBranch,omitempty"`

	// The path within the Git repository that represents the top level of the
	//  repo to sync. Default: the root directory of the repository.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.policy_dir
	PolicyDir *string `json:"policyDir,omitempty"`

	// Period in seconds between consecutive syncs. Default: 15.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.sync_wait_secs
	SyncWaitSecs *int64 `json:"syncWaitSecs,omitempty"`

	// Git revision (tag or hash) to check out. Default HEAD.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.sync_rev
	SyncRev *string `json:"syncRev,omitempty"`

	// Type of secret configured for access to the Git repo. Must be one of ssh,
	//  cookiefile, gcenode, token, gcpserviceaccount or none. The
	//  validation of this is case-sensitive. Required.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.secret_type
	SecretType *string `json:"secretType,omitempty"`

	// URL for the HTTPS proxy to be used when communicating with the Git repo.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.https_proxy
	HTTPSProxy *string `json:"httpsProxy,omitempty"`

	// The Google Cloud Service Account Email used for auth when secret_type is
	//  gcpServiceAccount.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GitConfig.gcp_service_account_email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.GroupVersionKind
type GroupVersionKind struct {
	// Kubernetes Group
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GroupVersionKind.group
	Group *string `json:"group,omitempty"`

	// Kubernetes Version
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GroupVersionKind.version
	Version *string `json:"version,omitempty"`

	// Kubernetes Kind
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.GroupVersionKind.kind
	Kind *string `json:"kind,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.HierarchyControllerConfig
type HierarchyControllerConfig struct {
	// Whether Hierarchy Controller is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Whether pod tree labels are enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerConfig.enable_pod_tree_labels
	EnablePodTreeLabels *bool `json:"enablePodTreeLabels,omitempty"`

	// Whether hierarchical resource quota is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerConfig.enable_hierarchical_resource_quota
	EnableHierarchicalResourceQuota *bool `json:"enableHierarchicalResourceQuota,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.HierarchyControllerDeploymentState
type HierarchyControllerDeploymentState struct {
	// The deployment state for open source HNC (e.g. v0.7.0-hc.0)
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerDeploymentState.hnc
	Hnc *string `json:"hnc,omitempty"`

	// The deployment state for Hierarchy Controller extension (e.g. v0.7.0-hc.1)
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerDeploymentState.extension
	Extension *string `json:"extension,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.HierarchyControllerState
type HierarchyControllerState struct {
	// The version for Hierarchy Controller
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerState.version
	Version *HierarchyControllerVersion `json:"version,omitempty"`

	// The deployment state for Hierarchy Controller
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerState.state
	State *HierarchyControllerDeploymentState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.HierarchyControllerVersion
type HierarchyControllerVersion struct {
	// Version for open source HNC
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerVersion.hnc
	Hnc *string `json:"hnc,omitempty"`

	// Version for Hierarchy Controller extension
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.HierarchyControllerVersion.extension
	Extension *string `json:"extension,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.InstallError
type InstallError struct {
	// A string representing the user facing error message
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.InstallError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.MembershipSpec
type MembershipSpec struct {
	// Config Sync configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.config_sync
	ConfigSync *ConfigSync `json:"configSync,omitempty"`

	// Policy Controller configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.policy_controller
	PolicyController *PolicyController `json:"policyController,omitempty"`

	// Hierarchy Controller configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.hierarchy_controller
	HierarchyController *HierarchyControllerConfig `json:"hierarchyController,omitempty"`

	// Version of ACM installed.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.version
	Version *string `json:"version,omitempty"`

	// The user-specified cluster name used by Config Sync cluster-name-selector
	//  annotation or ClusterSelector, for applying configs to only a subset
	//  of clusters.
	//  Omit this field if the cluster's fleet membership name is used by Config
	//  Sync cluster-name-selector annotation or ClusterSelector.
	//  Set this field if a name different from the cluster's fleet membership name
	//  is used by Config Sync cluster-name-selector annotation or ClusterSelector.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Enables automatic Feature management.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipSpec.management
	Management *string `json:"management,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.MembershipState
type MembershipState struct {
	// This field is set to the `cluster_name` field of the Membership Spec if it
	//  is not empty. Otherwise, it is set to the cluster's fleet membership name.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.cluster_name
	ClusterName *string `json:"clusterName,omitempty"`

	// Membership configuration in the cluster. This represents the actual state
	//  in the cluster, while the MembershipSpec in the FeatureSpec represents
	//  the intended state
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.membership_spec
	MembershipSpec *MembershipSpec `json:"membershipSpec,omitempty"`

	// Current install status of ACM's Operator
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.operator_state
	OperatorState *OperatorState `json:"operatorState,omitempty"`

	// Current sync status
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.config_sync_state
	ConfigSyncState *ConfigSyncState `json:"configSyncState,omitempty"`

	// PolicyController status
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.policy_controller_state
	PolicyControllerState *PolicyControllerState `json:"policyControllerState,omitempty"`

	// Hierarchy Controller status
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.MembershipState.hierarchy_controller_state
	HierarchyControllerState *HierarchyControllerState `json:"hierarchyControllerState,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.OciConfig
type OciConfig struct {
	// The OCI image repository URL for the package to sync from.
	//  e.g. `LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME`.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OciConfig.sync_repo
	SyncRepo *string `json:"syncRepo,omitempty"`

	// The absolute path of the directory that contains
	//  the local resources.  Default: the root directory of the image.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OciConfig.policy_dir
	PolicyDir *string `json:"policyDir,omitempty"`

	// Period in seconds between consecutive syncs. Default: 15.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OciConfig.sync_wait_secs
	SyncWaitSecs *int64 `json:"syncWaitSecs,omitempty"`

	// Type of secret configured for access to the Git repo.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OciConfig.secret_type
	SecretType *string `json:"secretType,omitempty"`

	// The Google Cloud Service Account Email used for auth when secret_type is
	//  gcpServiceAccount.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OciConfig.gcp_service_account_email
	GcpServiceAccountEmail *string `json:"gcpServiceAccountEmail,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.OperatorState
type OperatorState struct {
	// The semenatic version number of the operator
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OperatorState.version
	Version *string `json:"version,omitempty"`

	// The state of the Operator's deployment
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OperatorState.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Install errors.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.OperatorState.errors
	Errors []InstallError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.PolicyController
type PolicyController struct {
	// Enables the installation of Policy Controller.
	//  If false, the rest of PolicyController fields take no
	//  effect.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Installs the default template library along with Policy Controller.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.template_library_installed
	TemplateLibraryInstalled *bool `json:"templateLibraryInstalled,omitempty"`

	// Sets the interval for Policy Controller Audit Scans (in seconds).
	//  When set to 0, this disables audit functionality altogether.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.audit_interval_seconds
	AuditIntervalSeconds *int64 `json:"auditIntervalSeconds,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks.
	//  Namespaces do not need to currently exist on the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.exemptable_namespaces
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects
	//  other than the object currently being evaluated.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.referential_rules_enabled
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`

	// Logs all denies and dry run failures.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyController.log_denies_enabled
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.PolicyControllerState
type PolicyControllerState struct {
	// The version of Gatekeeper Policy Controller deployed.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyControllerState.version
	Version *PolicyControllerVersion `json:"version,omitempty"`

	// The state about the policy controller installation.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyControllerState.deployment_state
	DeploymentState *GatekeeperDeploymentState `json:"deploymentState,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.PolicyControllerVersion
type PolicyControllerVersion struct {
	// The gatekeeper image tag that is composed of ACM version, git tag, build
	//  number.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.PolicyControllerVersion.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.SyncError
type SyncError struct {
	// An ACM defined error code
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncError.code
	Code *string `json:"code,omitempty"`

	// A description of the error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// A list of config(s) associated with the error, if any
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncError.error_resources
	ErrorResources []ErrorResource `json:"errorResources,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1.SyncState
type SyncState struct {
	// Token indicating the state of the repo.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.source_token
	SourceToken *string `json:"sourceToken,omitempty"`

	// Token indicating the state of the importer.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.import_token
	ImportToken *string `json:"importToken,omitempty"`

	// Token indicating the state of the syncer.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.sync_token
	SyncToken *string `json:"syncToken,omitempty"`

	// Deprecated: use last_sync_time instead.
	//  Timestamp of when ACM last successfully synced the repo
	//  The time format is specified in https://golang.org/pkg/time/#Time.String
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.last_sync
	LastSync *string `json:"lastSync,omitempty"`

	// Timestamp type of when ACM last successfully synced the repo
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.last_sync_time
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Sync status code
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.code
	Code *string `json:"code,omitempty"`

	// A list of errors resulting from problematic configs.
	//  This list will be truncated after 100 errors, although it is
	//  unlikely for that many errors to simultaneously exist.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1.SyncState.errors
	Errors []SyncError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.multiclusteringress.v1.FeatureSpec
type FeatureSpec struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD.
	//  Example: `projects/foo-proj/locations/global/memberships/bar`
	// +kcc:proto:field=google.cloud.gkehub.multiclusteringress.v1.FeatureSpec.config_membership
	ConfigMembership *string `json:"configMembership,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.CommonFeatureSpec
type CommonFeatureSpec struct {
	// Multicluster Ingress-specific spec.
	// +kcc:proto:field=google.cloud.gkehub.v1.CommonFeatureSpec.multiclusteringress
	Multiclusteringress *FeatureSpec `json:"multiclusteringress,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.CommonFeatureState
type CommonFeatureState struct {
}

// +kcc:proto=google.cloud.gkehub.v1.Feature
type Feature struct {

	// GCP labels for this Feature.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Hub-wide Feature configuration. If this Feature does not support any
	//  Hub-wide configuration, this field may be unused.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.spec
	Spec *CommonFeatureSpec `json:"spec,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.gkehub.v1.FeatureResourceState
type FeatureResourceState struct {
	// The current state of the Feature resource in the Hub API.
	// +kcc:proto:field=google.cloud.gkehub.v1.FeatureResourceState.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.FeatureState
type FeatureState struct {
	// The high-level, machine-readable status of this Feature.
	// +kcc:proto:field=google.cloud.gkehub.v1.FeatureState.code
	Code *string `json:"code,omitempty"`

	// A human-readable description of the current status.
	// +kcc:proto:field=google.cloud.gkehub.v1.FeatureState.description
	Description *string `json:"description,omitempty"`

	// The time this status and any related Feature-specific details were updated.
	// +kcc:proto:field=google.cloud.gkehub.v1.FeatureState.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.MembershipFeatureSpec
type MembershipFeatureSpec struct {
	// Config Management-specific spec.
	// +kcc:proto:field=google.cloud.gkehub.v1.MembershipFeatureSpec.configmanagement
	Configmanagement *MembershipSpec `json:"configmanagement,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.MembershipFeatureState
type MembershipFeatureState struct {
	// Config Management-specific state.
	// +kcc:proto:field=google.cloud.gkehub.v1.MembershipFeatureState.configmanagement
	Configmanagement *MembershipState `json:"configmanagement,omitempty"`

	// The high-level state of this Feature for a single membership.
	// +kcc:proto:field=google.cloud.gkehub.v1.MembershipFeatureState.state
	State *FeatureState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.CommonFeatureState
type CommonFeatureStateObservedState struct {
	// Output only. The "running state" of the Feature in this Hub.
	// +kcc:proto:field=google.cloud.gkehub.v1.CommonFeatureState.state
	State *FeatureState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1.Feature
type FeatureObservedState struct {
	// Output only. The full, unique name of this Feature resource in the format
	//  `projects/*/locations/*/features/*`.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the Feature resource itself.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.resource_state
	ResourceState *FeatureResourceState `json:"resourceState,omitempty"`

	// Output only. The Hub-wide Feature state.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.state
	State *CommonFeatureState `json:"state,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. When the Feature resource was created.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the Feature resource was last updated.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. When the Feature resource was deleted.
	// +kcc:proto:field=google.cloud.gkehub.v1.Feature.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`
}
