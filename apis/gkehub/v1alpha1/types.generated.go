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


// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.BinauthzConfig
type BinauthzConfig struct {
	// Whether binauthz is enabled in this cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.BinauthzConfig.enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.BinauthzState
type BinauthzState struct {
	// The state of the binauthz webhook.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.BinauthzState.webhook
	Webhook *string `json:"webhook,omitempty"`

	// The version of binauthz that is installed.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.BinauthzState.version
	Version *BinauthzVersion `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.BinauthzVersion
type BinauthzVersion struct {
	// The version of the binauthz webhook.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.BinauthzVersion.webhook_version
	WebhookVersion *string `json:"webhookVersion,omitempty"`
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

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState
type ConfigSyncDeploymentState struct {
	// Deployment state of the importer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.importer
	Importer *string `json:"importer,omitempty"`

	// Deployment state of the syncer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.syncer
	Syncer *string `json:"syncer,omitempty"`

	// Deployment state of the git-sync pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.git_sync
	GitSync *string `json:"gitSync,omitempty"`

	// Deployment state of the monitor pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.monitor
	Monitor *string `json:"monitor,omitempty"`

	// Deployment state of reconciler-manager pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.reconciler_manager
	ReconcilerManager *string `json:"reconcilerManager,omitempty"`

	// Deployment state of root-reconciler
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.root_reconciler
	RootReconciler *string `json:"rootReconciler,omitempty"`

	// Deployment state of admission-webhook
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.admission_webhook
	AdmissionWebhook *string `json:"admissionWebhook,omitempty"`

	// Deployment state of resource-group-controller-manager
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.resource_group_controller_manager
	ResourceGroupControllerManager *string `json:"resourceGroupControllerManager,omitempty"`

	// Deployment state of otel-collector
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncDeploymentState.otel_collector
	OtelCollector *string `json:"otelCollector,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncError
type ConfigSyncError struct {
	// A string representing the user facing error message
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncState
type ConfigSyncState struct {
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion
type ConfigSyncVersion struct {
	// Version of the deployed importer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.importer
	Importer *string `json:"importer,omitempty"`

	// Version of the deployed syncer pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.syncer
	Syncer *string `json:"syncer,omitempty"`

	// Version of the deployed git-sync pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.git_sync
	GitSync *string `json:"gitSync,omitempty"`

	// Version of the deployed monitor pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.monitor
	Monitor *string `json:"monitor,omitempty"`

	// Version of the deployed reconciler-manager pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.reconciler_manager
	ReconcilerManager *string `json:"reconcilerManager,omitempty"`

	// Version of the deployed reconciler container in root-reconciler pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.root_reconciler
	RootReconciler *string `json:"rootReconciler,omitempty"`

	// Version of the deployed admission-webhook pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.admission_webhook
	AdmissionWebhook *string `json:"admissionWebhook,omitempty"`

	// Version of the deployed resource-group-controller-manager pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.resource_group_controller_manager
	ResourceGroupControllerManager *string `json:"resourceGroupControllerManager,omitempty"`

	// Version of the deployed otel-collector pod
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ConfigSyncVersion.otel_collector
	OtelCollector *string `json:"otelCollector,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.ErrorResource
type ErrorResource struct {
	// Path in the git repo of the erroneous config
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ErrorResource.source_path
	SourcePath *string `json:"sourcePath,omitempty"`

	// Metadata name of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ErrorResource.resource_name
	ResourceName *string `json:"resourceName,omitempty"`

	// Namespace of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ErrorResource.resource_namespace
	ResourceNamespace *string `json:"resourceNamespace,omitempty"`

	// Group/version/kind of the resource that is causing an error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.ErrorResource.resource_gvk
	ResourceGvk *GroupVersionKind `json:"resourceGvk,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.GatekeeperDeploymentState
type GatekeeperDeploymentState struct {
	// Status of gatekeeper-controller-manager pod.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GatekeeperDeploymentState.gatekeeper_controller_manager_state
	GatekeeperControllerManagerState *string `json:"gatekeeperControllerManagerState,omitempty"`

	// Status of gatekeeper-audit deployment.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GatekeeperDeploymentState.gatekeeper_audit
	GatekeeperAudit *string `json:"gatekeeperAudit,omitempty"`

	// Status of the pod serving the mutation webhook.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GatekeeperDeploymentState.gatekeeper_mutation
	GatekeeperMutation *string `json:"gatekeeperMutation,omitempty"`
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

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.GroupVersionKind
type GroupVersionKind struct {
	// Kubernetes Group
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GroupVersionKind.group
	Group *string `json:"group,omitempty"`

	// Kubernetes Version
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GroupVersionKind.version
	Version *string `json:"version,omitempty"`

	// Kubernetes Kind
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.GroupVersionKind.kind
	Kind *string `json:"kind,omitempty"`
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

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerDeploymentState
type HierarchyControllerDeploymentState struct {
	// The deployment state for open source HNC (e.g. v0.7.0-hc.0)
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerDeploymentState.hnc
	Hnc *string `json:"hnc,omitempty"`

	// The deployment state for Hierarchy Controller extension (e.g. v0.7.0-hc.1)
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerDeploymentState.extension
	Extension *string `json:"extension,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerState
type HierarchyControllerState struct {
	// The version for Hierarchy Controller
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerState.version
	Version *HierarchyControllerVersion `json:"version,omitempty"`

	// The deployment state for Hierarchy Controller
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerState.state
	State *HierarchyControllerDeploymentState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerVersion
type HierarchyControllerVersion struct {
	// Version for open source HNC
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerVersion.hnc
	Hnc *string `json:"hnc,omitempty"`

	// Version for Hierarchy Controller extension
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.HierarchyControllerVersion.extension
	Extension *string `json:"extension,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.InstallError
type InstallError struct {
	// A string representing the user facing error message
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.InstallError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec
type MembershipSpec struct {
	// Optional. Config Sync configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.config_sync
	ConfigSync *ConfigSync `json:"configSync,omitempty"`

	// Optional. Policy Controller configuration for the cluster.
	//  Deprecated: Configuring Policy Controller through the configmanagement
	//  feature is no longer recommended. Use the policycontroller feature instead.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.policy_controller
	PolicyController *PolicyController `json:"policyController,omitempty"`

	// Optional. Binauthz conifguration for the cluster. Deprecated: This field
	//  will be ignored and should not be set.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.binauthz
	Binauthz *BinauthzConfig `json:"binauthz,omitempty"`

	// Optional. Hierarchy Controller configuration for the cluster.
	//  Deprecated: Configuring Hierarchy Controller through the configmanagement
	//  feature is no longer recommended. Use
	//  https://github.com/kubernetes-sigs/hierarchical-namespaces instead.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.hierarchy_controller
	HierarchyController *HierarchyControllerConfig `json:"hierarchyController,omitempty"`

	// Optional. Version of ACM installed.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.version
	Version *string `json:"version,omitempty"`

	// Optional. The user-specified cluster name used by Config Sync
	//  cluster-name-selector annotation or ClusterSelector, for applying configs
	//  to only a subset of clusters. Omit this field if the cluster's fleet
	//  membership name is used by Config Sync cluster-name-selector annotation or
	//  ClusterSelector. Set this field if a name different from the cluster's
	//  fleet membership name is used by Config Sync cluster-name-selector
	//  annotation or ClusterSelector.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Optional. Enables automatic Feature management.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.MembershipSpec.management
	Management *string `json:"management,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.MembershipState
type MembershipState struct {
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

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.OperatorState
type OperatorState struct {
	// The semenatic version number of the operator
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OperatorState.version
	Version *string `json:"version,omitempty"`

	// The state of the Operator's deployment
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OperatorState.deployment_state
	DeploymentState *string `json:"deploymentState,omitempty"`

	// Install errors.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.OperatorState.errors
	Errors []InstallError `json:"errors,omitempty"`
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

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMigration
type PolicyControllerMigration struct {
	// Stage of the migration.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMigration.stage
	Stage *string `json:"stage,omitempty"`

	// Last time this membership spec was copied to PoCo feature.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMigration.copy_time
	CopyTime *string `json:"copyTime,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMonitoring
type PolicyControllerMonitoring struct {
	// Specifies the list of backends Policy Controller will export to.
	//  An empty list would effectively disable metrics export.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerMonitoring.backends
	Backends []string `json:"backends,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerState
type PolicyControllerState struct {
	// The version of Gatekeeper Policy Controller deployed.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerState.version
	Version *PolicyControllerVersion `json:"version,omitempty"`

	// The state about the policy controller installation.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerState.deployment_state
	DeploymentState *GatekeeperDeploymentState `json:"deploymentState,omitempty"`

	// Record state of ACM -> PoCo Hub migration for this feature.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerState.migration
	Migration *PolicyControllerMigration `json:"migration,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerVersion
type PolicyControllerVersion struct {
	// The gatekeeper image tag that is composed of ACM version, git tag, build
	//  number.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.PolicyControllerVersion.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.SyncError
type SyncError struct {
	// An ACM defined error code
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncError.code
	Code *string `json:"code,omitempty"`

	// A description of the error
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncError.error_message
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// A list of config(s) associated with the error, if any
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncError.error_resources
	ErrorResources []ErrorResource `json:"errorResources,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.configmanagement.v1beta.SyncState
type SyncState struct {
	// Token indicating the state of the repo.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.source_token
	SourceToken *string `json:"sourceToken,omitempty"`

	// Token indicating the state of the importer.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.import_token
	ImportToken *string `json:"importToken,omitempty"`

	// Token indicating the state of the syncer.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.sync_token
	SyncToken *string `json:"syncToken,omitempty"`

	// Deprecated: use last_sync_time instead.
	//  Timestamp of when ACM last successfully synced the repo
	//  The time format is specified in https://golang.org/pkg/time/#Time.String
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.last_sync
	LastSync *string `json:"lastSync,omitempty"`

	// Timestamp type of when ACM last successfully synced the repo
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.last_sync_time
	LastSyncTime *string `json:"lastSyncTime,omitempty"`

	// Sync status code
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.code
	Code *string `json:"code,omitempty"`

	// A list of errors resulting from problematic configs.
	//  This list will be truncated after 100 errors, although it is
	//  unlikely for that many errors to simultaneously exist.
	// +kcc:proto:field=google.cloud.gkehub.configmanagement.v1beta.SyncState.errors
	Errors []SyncError `json:"errors,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.metering.v1beta.MembershipState
type MembershipState struct {
	// The time stamp of the most recent measurement of the number of vCPUs
	//  in the cluster.
	// +kcc:proto:field=google.cloud.gkehub.metering.v1beta.MembershipState.last_measurement_time
	LastMeasurementTime *string `json:"lastMeasurementTime,omitempty"`

	// The vCPUs capacity in the cluster according to the most recent
	//  measurement (1/1000 precision).
	// +kcc:proto:field=google.cloud.gkehub.metering.v1beta.MembershipState.precise_last_measured_cluster_vcpu_capacity
	PreciseLastMeasuredClusterVcpuCapacity *float32 `json:"preciseLastMeasuredClusterVcpuCapacity,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.multiclusteringress.v1beta.FeatureSpec
type FeatureSpec struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD.
	//  Example: `projects/foo-proj/locations/global/memberships/bar`
	// +kcc:proto:field=google.cloud.gkehub.multiclusteringress.v1beta.FeatureSpec.config_membership
	ConfigMembership *string `json:"configMembership,omitempty"`

	// Customer's billing structure
	// +kcc:proto:field=google.cloud.gkehub.multiclusteringress.v1beta.FeatureSpec.billing
	Billing *string `json:"billing,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.BundleInstallSpec
type BundleInstallSpec struct {
	// The set of namespaces to be exempted from the bundle.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.BundleInstallSpec.exempted_namespaces
	ExemptedNamespaces []string `json:"exemptedNamespaces,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.HubConfig
type HubConfig struct {
	// The install_spec represents the intended state specified by the
	//  latest request that mutated install_spec in the feature spec,
	//  not the lifecycle state of the
	//  feature observed by the Hub feature controller
	//  that is reported in the feature state.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.install_spec
	InstallSpec *string `json:"installSpec,omitempty"`

	// Sets the interval for Policy Controller Audit Scans (in seconds).
	//  When set to 0, this disables audit functionality altogether.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.audit_interval_seconds
	AuditIntervalSeconds *int64 `json:"auditIntervalSeconds,omitempty"`

	// The set of namespaces that are excluded from Policy Controller checks.
	//  Namespaces do not need to currently exist on the cluster.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.exemptable_namespaces
	ExemptableNamespaces []string `json:"exemptableNamespaces,omitempty"`

	// Enables the ability to use Constraint Templates that reference to objects
	//  other than the object currently being evaluated.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.referential_rules_enabled
	ReferentialRulesEnabled *bool `json:"referentialRulesEnabled,omitempty"`

	// Logs all denies and dry run failures.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.log_denies_enabled
	LogDeniesEnabled *bool `json:"logDeniesEnabled,omitempty"`

	// Enables the ability to mutate resources using Policy Controller.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.mutation_enabled
	MutationEnabled *bool `json:"mutationEnabled,omitempty"`

	// Monitoring specifies the configuration of monitoring.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.monitoring
	Monitoring *MonitoringConfig `json:"monitoring,omitempty"`

	// Specifies the desired policy content on the cluster
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.policy_content
	PolicyContent *PolicyContentSpec `json:"policyContent,omitempty"`

	// The maximum number of audit violations to be stored in a constraint.
	//  If not set, the internal default (currently 20) will be used.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.HubConfig.constraint_violation_limit
	ConstraintViolationLimit *int64 `json:"constraintViolationLimit,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.MembershipSpec
type MembershipSpec struct {
	// Policy Controller configuration for the cluster.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MembershipSpec.policy_controller_hub_config
	PolicyControllerHubConfig *HubConfig `json:"policyControllerHubConfig,omitempty"`

	// Version of Policy Controller installed.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MembershipSpec.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.MembershipState
type MembershipState struct {

	// TODO: unsupported map type with key string and value message


	// The overall Policy Controller lifecycle state observed by the Hub Feature
	//  controller.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MembershipState.state
	State *string `json:"state,omitempty"`

	// The overall content state observed by the Hub Feature controller.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MembershipState.policy_content_state
	PolicyContentState *PolicyContentState `json:"policyContentState,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.MonitoringConfig
type MonitoringConfig struct {
	// Specifies the list of backends Policy Controller will export to.
	//  An empty list would effectively disable metrics export.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.MonitoringConfig.backends
	Backends []string `json:"backends,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.OnClusterState
type OnClusterState struct {
	// The lifecycle state of this component.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.OnClusterState.state
	State *string `json:"state,omitempty"`

	// Surface potential errors or information logs.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.OnClusterState.details
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.PolicyContentSpec
type PolicyContentSpec struct {

	// TODO: unsupported map type with key string and value message


	// Configures the installation of the Template Library.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyContentSpec.template_library
	TemplateLibrary *TemplateLibraryConfig `json:"templateLibrary,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.PolicyContentState
type PolicyContentState struct {
	// The state of the template library
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyContentState.template_library_state
	TemplateLibraryState *OnClusterState `json:"templateLibraryState,omitempty"`

	// TODO: unsupported map type with key string and value message


	// The state of the referential data sync configuration.  This could
	//  represent the state of either the syncSet object(s) or the config
	//  object, depending on the version of PoCo configured by the user.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyContentState.referential_sync_config_state
	ReferentialSyncConfigState *OnClusterState `json:"referentialSyncConfigState,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig
type PolicyControllerDeploymentConfig struct {
	// Pod replica count.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.replica_count
	ReplicaCount *int64 `json:"replicaCount,omitempty"`

	// Container resource requirements.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.container_resources
	ContainerResources *ResourceRequirements `json:"containerResources,omitempty"`

	// Pod anti-affinity enablement. Deprecated: use `pod_affinity` instead.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.pod_anti_affinity
	PodAntiAffinity *bool `json:"podAntiAffinity,omitempty"`

	// Pod tolerations of node taints.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.pod_tolerations
	PodTolerations []PolicyControllerDeploymentConfig_Toleration `json:"podTolerations,omitempty"`

	// Pod affinity configuration.
	// +kcc:proto:field=google.cloud.gkehub.policycontroller.v1beta.PolicyControllerDeploymentConfig.pod_affinity
	PodAffinity *string `json:"podAffinity,omitempty"`
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
	Cpu *string `json:"cpu,omitempty"`
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

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.MembershipState
type MembershipState struct {
}

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.MembershipState.Condition
type MembershipState_Condition struct {
	// Unique identifier of the condition which describes the condition
	//  recognizable to the user.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.Condition.code
	Code *string `json:"code,omitempty"`

	// Links contains actionable information.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.Condition.documentation_link
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// A short summary about the issue.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.Condition.details
	Details *string `json:"details,omitempty"`

	// Severity level of the condition.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.Condition.severity
	Severity *string `json:"severity,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.MembershipState.ControlPlaneManagement
type MembershipState_ControlPlaneManagement struct {
	// Explanation of state.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.ControlPlaneManagement.details
	Details []StatusDetails `json:"details,omitempty"`

	// LifecycleState of control plane management.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.ControlPlaneManagement.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.MembershipState.DataPlaneManagement
type MembershipState_DataPlaneManagement struct {
	// Lifecycle status of data plane management.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.DataPlaneManagement.state
	State *string `json:"state,omitempty"`

	// Explanation of the status.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.MembershipState.DataPlaneManagement.details
	Details []StatusDetails `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.servicemesh.v1beta.StatusDetails
type StatusDetails struct {
	// A machine-readable code that further describes a broad status.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.StatusDetails.code
	Code *string `json:"code,omitempty"`

	// Human-readable explanation of code.
	// +kcc:proto:field=google.cloud.gkehub.servicemesh.v1beta.StatusDetails.details
	Details *string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.CommonFeatureSpec
type CommonFeatureSpec struct {
	// Multicluster Ingress-specific spec.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.CommonFeatureSpec.multiclusteringress
	Multiclusteringress *FeatureSpec `json:"multiclusteringress,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.CommonFeatureState
type CommonFeatureState struct {
}

// +kcc:proto=google.cloud.gkehub.v1beta.Feature
type Feature struct {

	// Labels for this Feature.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Hub-wide Feature configuration. If this Feature does not support
	//  any Hub-wide configuration, this field may be unused.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.spec
	Spec *CommonFeatureSpec `json:"spec,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.gkehub.v1beta.FeatureResourceState
type FeatureResourceState struct {
	// The current state of the Feature resource in the Hub API.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.FeatureResourceState.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.FeatureState
type FeatureState struct {
	// The high-level, machine-readable status of this Feature.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.FeatureState.code
	Code *string `json:"code,omitempty"`

	// A human-readable description of the current status.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.FeatureState.description
	Description *string `json:"description,omitempty"`

	// The time this status and any related Feature-specific details were updated.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.FeatureState.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.MembershipFeatureSpec
type MembershipFeatureSpec struct {
	// Config Management-specific spec.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureSpec.configmanagement
	Configmanagement *MembershipSpec `json:"configmanagement,omitempty"`

	// Anthos Service Mesh-specific spec
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureSpec.mesh
	Mesh *MembershipSpec `json:"mesh,omitempty"`

	// Policy Controller spec.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureSpec.policycontroller
	Policycontroller *MembershipSpec `json:"policycontroller,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.MembershipFeatureState
type MembershipFeatureState struct {
	// Service Mesh-specific state.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureState.servicemesh
	Servicemesh *MembershipState `json:"servicemesh,omitempty"`

	// Metering-specific state.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureState.metering
	Metering *MembershipState `json:"metering,omitempty"`

	// Config Management-specific state.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureState.configmanagement
	Configmanagement *MembershipState `json:"configmanagement,omitempty"`

	// Policycontroller-specific state.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureState.policycontroller
	Policycontroller *MembershipState `json:"policycontroller,omitempty"`

	// The high-level state of this Feature for a single membership.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.MembershipFeatureState.state
	State *FeatureState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.CommonFeatureState
type CommonFeatureStateObservedState struct {
	// Output only. The "running state" of the Feature in this Hub.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.CommonFeatureState.state
	State *FeatureState `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.gkehub.v1beta.Feature
type FeatureObservedState struct {
	// Output only. The full, unique name of this Feature resource in the format
	//  `projects/*/locations/*/features/*`.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.name
	Name *string `json:"name,omitempty"`

	// Output only. State of the Feature resource itself.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.resource_state
	ResourceState *FeatureResourceState `json:"resourceState,omitempty"`

	// Output only. The Hub-wide Feature state.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.state
	State *CommonFeatureState `json:"state,omitempty"`

	// TODO: unsupported map type with key string and value message


	// Output only. When the Feature resource was created.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. When the Feature resource was last updated.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. When the Feature resource was deleted.
	// +kcc:proto:field=google.cloud.gkehub.v1beta.Feature.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`
}
