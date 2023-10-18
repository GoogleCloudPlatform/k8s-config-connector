# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The third_party directory provides the patches version of the Terraform provider, that is used by
# Config Conector.
# This directory and the patched provider MUST be checked into version control, as the cnrm code is
# used as a in tools (including but not limited to) anthos-cli.
.PHONY: all
all: clone-terraform-google-provider-beta apply-patches

.PHONY: clone-terraform-google-provider-beta
clone-terraform-google-provider-beta:
	rm -rf github.com/hashicorp/terraform-provider-google-beta
	mkdir -p github.com/hashicorp
	cd github.com/hashicorp && git clone https://github.com/hashicorp/terraform-provider-google-beta --branch v4.84.0 --depth 1
	rm -rf github.com/hashicorp/terraform-provider-google-beta/.git

.PHONY: apply-patches
apply-patches:
	# Fix to the google_container_cluster resource:
	# (1)The TF provider added a change that waits for GKE clusters to be ready
	# before the import completes. As we use import before the resource actually
	# exists, this breaks our normal flow.
	# (2)Add support for importing ContainerCluster from CAI resource name which uses the legacy 'zones/{zone}' instead
	# of 'locations/{location}'
	# (3)Fix faulty return value on `expandIPAllocationPolicy`
	# TODO(b/208277276): Upstream change to TF
	# (4)Preserve basic authentication configuration option that has been removed in tf4.x upgrade.
	# See b/206133327 for context.
	# (5)tpu_config field is causing constant diff during ContainerCluster resource reconciliation.
	# Ignore the field for now to fix build break and unblock KCC releases.
	# See b/239812293 for context
	# (6)The default value of `cluster_autoscaling.auto_provisioning_defaults.disk_type`
	# causes an unexpected diff because when the cluster enables autopilot, the
	# computed result of `cluster_autoscaling.auto_provisioning_defaults` has
	# `disk_type` field unset.
	# TODO(b/258864199): Remove this patch if disk_type field no longer causes unexpected diffs.
	git apply ../hack/terraform-overrides/container_cluster.patch
	# Fix to the google_container_node_pool resource:
	# (1)Add support for importing ContainerNodePool from CAI resource name which uses the legacy 'zones/{zone}' instead
	# of 'locations/{location}'
	# (2)Do not require a GET to wait for the node pool to be ready during import.
	# (3)Add a way to disable autoscaling by setting min_node_count and max_node_count to 0
	# This change will not be upstreamed to TF (they already allow users to
	# disable autoscaling by unsetting the field) and was implemented because
	# Kubernetes does not allow objects to be set from non-null to null.
	git apply ../hack/terraform-overrides/container_nodepool.patch
	# Terraform client uses a list with MaxItems==1 for nested fields and lists with size limits
	# We assume that a list with MaxItems 1 is a field. The following patches enforce this behaviour on the client's behalf.
	git apply ../hack/terraform-overrides/compute_resource_policy.patch
	# Fix to the google_project resource:
	# (1) When you call GET on a project that hasn't been created yet, the
	# underlying API always returns 403, which the provider treats as a real
	# error. Instead, treat this the same as 404 NotFound for projects.
	# (2) add support for importing Projects by their ProjectNumber. Asset
	# inventory uses the project number ID format to refer to projects so this
	# is needed to for support for going from asset to KCC resource
	git apply ../hack/terraform-overrides/project.patch
	# Fix to the google_folder resource:
	# (1) Folders should be treated as gone if the state is DELETE_REQUESTED.
	# This is consistent with how projects are treated in the provider.
	# (2) Set the 'name' field during the TF import. This is required by the
	# config-connector CLI, specifically to retain the folder ID when
	# converting the folder's TF state to a KCC resource skeleton.
	# (3) Replace 'parent' field with 'parent_org_id' and 'parent_folder_id' so
	# that we can convert these fields to resource references. This is
	# necessary since KCC can't map a TF field to the appropriate KRM field
	# when there are multiple possible KRM fields a TF field can be mapped
	# back to.
	# (4) On create, check to see if a folder with the given display_name in
	# the given parent exists before trying to create a new folder. If one
	# exists, read that one instead. This allows users to acquire folders
	# by specifying the folder's display_name and parent in KCC.
	git apply ../hack/terraform-overrides/folder.patch
	# Our logic requires a 'project' field on ServiceNetworkingConnection.
	# ServiceNetworkingConnection doesn't have one, so we add it.
	git apply ../hack/terraform-overrides/service_networking_connection_project.patch
	# Fixes to the google_dataflow_job resource:
	# 1) Change the TF read to return more fields so that they can be used
	# as part of the resource's live state when generating diffs.
	# 2) Remove the CustomizeDiff that treats any diffs for batch jobs as
	# ForceNew since it triggers a ForceNew even for diffs that have been
	# suppressed by DiffSuppressFunc functions.
	# 3) Make format of 'subnetwork' field consistent with other TF
	# reference fields
	# 4) Don't remove resource on read just because the job is in a
	# terminal state. Instead, wait for the user to explicitly remove it.
	# 5) Allow for acquisition of active jobs by name.
	# 6) Ensure field additional_experiments is readable.
	git apply ../hack/terraform-overrides/dataflow_job.patch
	# Don't remove resource on read just because the job is in a terminal
	# state. Instead, wait for the user to explicitly remove it.
	git apply ../hack/terraform-overrides/dataflow_flex_template_job.patch
	# Make the same changes in
	# https://github.com/hashicorp/terraform-provider-google-beta/pull/6357
	# to update google_dataflow_flex_template_job to send SdkPipeline parameters
	# via environment block fields, and to fix the issue that the returned
    # maxWorkers was always 0.
	git apply ../hack/terraform-overrides/dataflow_flex_template_job_environment_fields.patch
	# Adds support for TF imports for google_compute_instance_from_template
	git apply ../hack/terraform-overrides/compute_instance_from_template_import.patch
	# The `StorageBucketAccessControl.get` call returns 400 when its entity
	# refers to a service account that no longer exists. As the API
	# automatically prunes access controls for deleted users, this should
	# be treated the same as a 404.
	# Tracking issue: https://github.com/terraform-providers/terraform-provider-google/issues/5852
	git apply ../hack/terraform-overrides/storage_bucket_access_control.patch
	# Suppress a diff between topic ID vs topic relative resource name
	git apply ../hack/terraform-overrides/source_repo_repository.patch
	# 1) Add support to google_project_service's import for one platform resource
	# name
	# 2) make Read return nil when resource/service is not enabled
	git apply ../hack/terraform-overrides/project_service.patch
	# Change the default value for create_disposition to CREATE_NEVER so
	# that referencing a non-existent BigQuery table does not cause
	# BigQuery to create the table by default.
	git apply ../hack/terraform-overrides/bigquery_job_create_disposition_defaults.patch
	# Treat backends as computed so they are merged by the diff.
	git apply ../hack/terraform-overrides/backend_service_computed_backend.patch
	# Fixes to google_bigtable_instance:
	# 1) Set "instance_type" as part of the TF read, otherwise a diff would
	# be generated on every reconcile, causing an update to be triggered
	# even when there is no drift.
	# 2) Render "deletion_protection" useless so that it can be removed
	# later. Its reason for existing is considered incompatible with KCC.
	# It was also causing diffs to be generated on every reconciliation,
	# thereby triggering constant updates.
	# 3) Set "zone" field as required since there is no default value from KCC
	git apply ../hack/terraform-overrides/bigtable_instance.patch
	# Enables service accounts to be imported by unique ID (in addition to
	# email)
	git apply ../hack/terraform-overrides/service_account.patch
	# Makes 'service' immutable since that is what uniquely identifies an
	# IAM audit config. Letting this field be mutable will lead to a UX
	# that is inconsistent with other resources whose identifiers cannot be
	# changed.
	git apply ../hack/terraform-overrides/iam_audit_config.patch
	# Update the default description from "Managed by Terraform" to "Managed by
	# Config Connector" in dns_managed_zone and dns_policy.
	git apply ../hack/terraform-overrides/dns_resources_default_description.patch
	# The newly-introduced confidential_instance_config field on
	# compute_instance_template is incorrectly marked as Computed, causing
	# a persistent ForceNew diff.
	git apply ../hack/terraform-overrides/instance_template_confidential_instance_config.patch
	# When a schema is not specified, the TF provider panics, as during KRM2TF
	# conversion after presetting immutable fields for read, empty maps are
	# converted to nil, which fail a map type assertion.
	# TODO(b/168805790): Preset empty immutable fields to their zero value.
	git apply ../hack/terraform-overrides/bigquery_table_schema_fix.patch
	# TODO(b/170059118): upstream the following schema fixes
	# `private_interconnect_info` field should be an object rather than a list
	git apply ../hack/terraform-overrides/compute_interconnect_attachment.patch
	# `server_ca_cert` field should be an object rather than a list
	# `replication_type` field is deprecated, remove the defaulting to avoid unexpected diffs.
	git apply ../hack/terraform-overrides/sql_database_instance.patch
	# Re-mark `settings` as a required field. Remove patch once we support
	# the `clone` field (b/179734967)
	git apply ../hack/terraform-overrides/sql_database_instance_settings.patch
	# `project_team` field should be an object rather than a list
	git apply ../hack/terraform-overrides/storage_default_object_access_control.patch
	# 1) Add combined_resources.go so that we can create custom TF resources
	# that combined multiple TF resources.
	# 2) Set unique_writer_identity to false for Org/Folder/BillingAccount logging sink
	# on read, this avoids constant diff being generated in the combined KCC resource.
	git apply ../hack/terraform-overrides/combined_resources.patch
	# Add a combined logging log sink resource to the provider, in order to
	# avoid needing to implement more generic resource merging in KCC.
	git apply ../hack/terraform-overrides/combined_log_sink.patch
	# Add a combined IAM role resource to the provider, in order to avoid
	# needing to implement more generic resource merging in KCC.
	git apply ../hack/terraform-overrides/combined_iam_custom_role.patch
	# The provider is introducing breaking changes to database resources to
	# prevent accidental stateful resource deletion in Terraform. KCC
	# customers rely on deletion behavior being consistent with its current
	# behavior, and the accidental stateful resource deletion scenario
	# isn't as much of an issue in KCC, so we patch this new `true` default
	# out.
	git apply ../hack/terraform-overrides/deletion_protection_default.patch
	# Return an empty instance state on the TF read if the policy is
	# "<never>" (i.e. an empty policy). This fixes an issue with
	# BigtableGCPolicy where it could not create an underlying policy at
	# all even if it was non-existent. Note that this is a partial fix, and
	# a fuller fix for BigtableGCPolicy should follow (b/173321896).
	git apply ../hack/terraform-overrides/bigtable_gc_policy_read_empty_on_empty_policy.patch
	# Rename the 'instance' field to 'instance_name' to be consistent with google_bigtable_table;
	# otherwise iam controllers cannot heuristically construct an object with all required information (instance in this case) to invoke bigtable_table_iam.
	# Remove this patch when https://github.com/hashicorp/terraform-provider-google/issues/7939 is resolved.
	git apply ../hack/terraform-overrides/iam_bigtable_table.patch
	# Remove MaxItems==1 for spec.MaxAge so that the field is still treated as a list instead of an object
	git apply ../hack/terraform-overrides/bigtable_gc_policy.patch
	# Set the default value of the virtual field `deletion_protection` in google_bigquery_table
	# to false. We determined to ignore this field in order to avoid breaking changes,
	# but meanwhile we need to disable the deletion protection to enable `delete`.
	git apply ../hack/terraform-overrides/bigquery_table_deletion_protection_default.patch
	# Keep auth_string as optional and non-sensitive to avoid breaking schema changes.
	# This field should be output-only and be placed into status. We have a tracking bug (b/182289489) to
	# fix it eventually.
	git apply ../hack/terraform-overrides/redis_instance.patch
	# Replace 'parent' with an 'org_id' field that doesn't require an
	# "organizations/" prefix. This is done to remove the need for a
	# valueTemplate for the organization-id container annotation which we
	# needed previously to add the "organizations/" prefix to the
	# annotation value. This is done so that we can more easily convert
	# organization-id annotations to external references by just taking the
	# annotation value and putting it into `external` without having to
	# prefix it with "organizations/"
	git apply ../hack/terraform-overrides/access_context_manager_access_policy_parent.patch
	# Add a combined org policy resource to the provider, in order to avoid
	# needing to implement more generic resource merging in KCC.
	git apply ../hack/terraform-overrides/combined_org_policy.patch
	# Modifying the description of compute_disk.size to not include "Terraform",
	# which strips the text of the description. In addition updating the description
	# to describe KCC behavior.
	git apply ../hack/terraform-overrides/compute_disk_description.patch
	# Make iap subfield optional so ComputeBackendService can be acquired with IAP enabled
	# (b/180415885)
	git apply ../hack/terraform-overrides/compute_backend_service.patch
	# When GET is called on a recently deleted CloudIdentityGroup, a 403 is
	# returned. This handles 403 errors as NotFound for CloudIdentityGroups.
	git apply ../hack/terraform-overrides/cloud_identity_group.patch
	# Suppress iam etag diff. This is needed to avoid constant diff generated
	# because terraform will always set etag as empty in the old state.
	# TODO(b/188205161): Upstream the change
	git apply ../hack/terraform-overrides/iam_etag_diff_suppress.patch
	# Remove the project Get request before enabling the given service; this Get request causes
	# a chicken-and-egg problem when enabling `cloudresourcemanager.googleapis.com` API.
	git apply ../hack/terraform-overrides/project_service_remove_project_read.patch
	# Remove the logic in the read function that keeps retrying the GET
	# until it stops returning 404. This is problematic for KCC since KCC
	# does a read before it creates the resource, which means the GET will
	# always return a 404, and will continue to be retried until the
	# timeout.
	git apply ../hack/terraform-overrides/storage_bucket_read.patch
	# Preserve the node_medadata configuration option that has been removed in tf4.x upgrade.
	# See b/206133327 for context.
	git apply ../hack/terraform-overrides/container_cluster_node_metadata_tf4.x_upgrade.patch
	# Preserve the 'self_link' field that has been removed in tf4.x upgrade.
	# See b/206136846 for context.
	git apply ../hack/terraform-overrides/kms_crypto_key_self_link_tf4.x_upgrade.patch
	# Preserve the 'self_link' field that has been removed in tf4.x upgrade.
	# See b/206136846 for context.
	git apply ../hack/terraform-overrides/kms_key_ring_self_link_tf4.x_upgrade.patch
	# Revert the change in tf4.x upgrade that makes one of the source parameters required in ingress firewalls
	# since it's a breaking change.
	git apply ../hack/terraform-overrides/compute_firewall.patch
	# The newly-introduced `share_settings` field is an optional computed list. When an optional
	# computed list field is omitted, it causes a persistent RequiresNew diff.
	# Removing `Computed: true` in the schema provides a workaround for the unexpected diff result.
	git apply ../hack/terraform-overrides/compute_reservation.patch
	# Ignore routing_policy due to complexity introduced with top layer rrdatas/rrdatasRefs
	# See b/233267548 for context
	git apply ../hack/terraform-overrides/dns_ignore_routing_policy.patch
	# Change field "layer_7_ddos_defense_config" to "layer7_ddos_defense_config"
	# See b/220930832 for context
	git apply ../hack/terraform-overrides/compute_security_policy_fieldname.patch
	# Change ip_version field description for compute_global_address to include `This field
	# can only be specified for a global address.`
	# See b/238334650 for context
	git apply ../hack/terraform-overrides/compute_global_address_ip_version_description.patch
	# Allow specifying custom HTTP clients to support mock testing
	# See b/238356047 for context
	git apply ../hack/terraform-overrides/mock-http-clients.patch
	# Change TF read to return non-empty resource when state is DESTROYED.
	# This was the behavior before the upgrade to TF v4.41.0. This is done
	# since some users might have been depending on KCC to update
	# status.destroyTime when the resource is DESTROYED.
	# TODO(b/254533796): Re-evaluate if this patch is worth keeping.
	git apply ../hack/terraform-overrides/secret_manager_secret_version_destroyed.patch
	# Avoid breaking change from TF with MaxItem = 1, which will update type from array to object in CRD
	# See b/283317413 for context
	git apply ../hack/terraform-overrides/compute_forwarding_rule.patch
	# Fix the incorrect `Computed: true` schema for the immutable, unreadable,
	# and non-computed `params` field in the
	# `google_compute_instance_from_template` resource.
	git apply ../hack/terraform-overrides/compute_instance_from_template_params_field.patch
	# Correct the field type of terminal_condition field. b/296456403.
	git apply ../hack/terraform-overrides/runservice_terminalcondition.patch
	# Mark certificate id as readonly
	git apply ../hack/terraform-overrides/compute_managed_ssl_certificate.patch
