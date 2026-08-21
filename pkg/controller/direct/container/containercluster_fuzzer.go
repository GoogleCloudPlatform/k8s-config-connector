// Copyright 2026 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.container.v1.Cluster
// api.group: container.cnrm.cloud.google.com

package container

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(containerClusterFuzzer())
}

func containerClusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Cluster{},
		ContainerClusterSpec_FromProto, ContainerClusterSpec_ToProto,
		ClusterObservedState_FromProto, ClusterObservedState_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")

	// Spec fields that are successfully mapped and round-tripped
	f.SpecField(".description")
	f.SpecField(".initial_node_count")
	f.SpecField(".node_config")
	f.SpecField(".node_config.machine_type")
	f.SpecField(".node_config.oauth_scopes")
	f.SpecField(".node_config.service_account")
	f.SpecField(".node_config.metadata")
	f.SpecField(".node_config.image_type")
	f.SpecField(".node_config.labels")
	f.SpecField(".node_config.local_ssd_count")
	f.SpecField(".node_config.tags")
	f.SpecField(".node_config.preemptible")
	f.SpecField(".node_config.disk_type")
	f.SpecField(".node_config.min_cpu_platform")
	f.SpecField(".node_config.sandbox_config")
	f.SpecField(".node_config.node_group")
	f.SpecField(".node_config.reservation_affinity")
	f.SpecField(".node_config.shielded_instance_config")
	f.SpecField(".node_config.linux_node_config")
	f.SpecField(".node_config.kubelet_config")
	f.SpecField(".node_config.gcfs_config")
	f.SpecField(".node_config.spot")
	f.SpecField(".node_config.confidential_nodes")
	f.SpecField(".node_config.fast_socket")
	f.SpecField(".node_config.resource_labels")
	f.SpecField(".node_config.ephemeral_storage_local_ssd_config")
	f.SpecField(".master_auth")
	f.SpecField(".logging_service")
	f.SpecField(".monitoring_service")
	f.SpecField(".network")
	f.SpecField(".cluster_ipv4_cidr")
	f.SpecField(".addons_config")
	f.SpecField(".subnetwork")
	f.SpecField(".enable_kubernetes_alpha")
	f.SpecField(".network_policy")
	f.SpecField(".ip_allocation_policy")
	f.SpecField(".master_authorized_networks_config")
	f.SpecField(".maintenance_policy")
	f.SpecField(".binary_authorization")
	f.SpecField(".resource_usage_export_config")
	f.SpecField(".authenticator_groups_config")
	f.SpecField(".private_cluster_config")
	f.SpecField(".database_encryption")
	f.SpecField(".vertical_pod_autoscaling")
	f.SpecField(".release_channel")
	f.SpecField(".workload_identity_config")
	f.SpecField(".mesh_certificates")
	f.SpecField(".cost_management_config")
	f.SpecField(".notification_config")
	f.SpecField(".confidential_nodes")
	f.SpecField(".identity_service_config")
	f.SpecField(".location")
	f.SpecField(".enable_tpu")
	f.SpecField(".node_pool_defaults")
	f.SpecField(".logging_config")
	f.SpecField(".monitoring_config")
	f.SpecField(".node_pool_auto_config")
	f.SpecField(".security_posture_config")
	f.SpecField(".control_plane_endpoints_config")
	f.SpecField(".enable_k8s_beta_apis")
	f.SpecField(".autopilot")
	f.SpecField(".legacy_abac")
	f.SpecField(".shielded_nodes")
	f.SpecField(".default_max_pods_constraint")
	f.SpecField(".network_config")

	// Status fields that are successfully mapped and round-tripped
	f.StatusField(".master_auth")
	f.StatusField(".private_cluster_config")
	f.StatusField(".control_plane_endpoints_config")

	// Unmapped / Unimplemented Spec fields
	f.Unimplemented_NotYetTriaged(".addons_config.gce_persistent_disk_csi_driver_config")
	f.Unimplemented_NotYetTriaged(".addons_config.gcp_filestore_csi_driver_config")
	f.Unimplemented_NotYetTriaged(".addons_config.gcs_fuse_csi_driver_config")
	f.Unimplemented_NotYetTriaged(".addons_config.high_scale_checkpointing_config")
	f.Unimplemented_NotYetTriaged(".addons_config.kubernetes_dashboard")
	f.Unimplemented_NotYetTriaged(".addons_config.lustre_csi_driver_config")
	f.Unimplemented_NotYetTriaged(".addons_config.node_readiness_config")
	f.Unimplemented_NotYetTriaged(".addons_config.parallelstore_csi_driver_config")
	f.Unimplemented_NotYetTriaged(".addons_config.pod_snapshot_config")
	f.Unimplemented_NotYetTriaged(".addons_config.ray_operator_config")
	f.Unimplemented_NotYetTriaged(".addons_config.slice_controller_config")
	f.Unimplemented_NotYetTriaged(".addons_config.slurm_operator_config")
	f.Unimplemented_NotYetTriaged(".addons_config.stateful_ha_config")
	f.Unimplemented_NotYetTriaged(".alpha_cluster_feature_gates")
	f.Unimplemented_NotYetTriaged(".anonymous_authentication_config")
	f.Unimplemented_NotYetTriaged(".authenticator_groups_config.enabled")
	f.Unimplemented_NotYetTriaged(".autopilot.privileged_admission_config")
	f.Unimplemented_NotYetTriaged(".autopilot.workload_policy_config")
	f.Unimplemented_NotYetTriaged(".autopilot.cluster_policy_config")
	f.Unimplemented_NotYetTriaged(".autoscaling")
	f.Unimplemented_NotYetTriaged(".compliance_posture_config")
	f.Unimplemented_NotYetTriaged(".conditions")
	f.Unimplemented_NotYetTriaged(".confidential_nodes.confidential_instance_type")
	f.Unimplemented_NotYetTriaged(".control_plane_egress")
	f.Unimplemented_NotYetTriaged(".control_plane_endpoints_config.dns_endpoint_config")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".current_master_version")
	f.Unimplemented_NotYetTriaged(".current_node_count")
	f.Unimplemented_NotYetTriaged(".current_node_version")
	f.Unimplemented_NotYetTriaged(".database_encryption.current_state")
	f.Unimplemented_NotYetTriaged(".database_encryption.decryption_keys")
	f.Unimplemented_NotYetTriaged(".database_encryption.last_operation_errors")
	f.Unimplemented_NotYetTriaged(".default_max_pods_constraint.max_pods_per_node")
	f.Unimplemented_NotYetTriaged(".enable_k8s_beta_apis.enabled_apis")
	f.Unimplemented_NotYetTriaged(".endpoint")
	f.Unimplemented_NotYetTriaged(".enterprise_config")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".expire_time")
	f.Unimplemented_NotYetTriaged(".fleet")
	f.Unimplemented_NotYetTriaged(".gke_auto_upgrade_config")
	f.Unimplemented_NotYetTriaged(".id")
	f.Unimplemented_NotYetTriaged(".initial_cluster_version")
	f.Unimplemented_NotYetTriaged(".instance_group_urls")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.additional_ip_ranges_configs[].pod_ipv4_range_names")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.additional_ip_ranges_configs[].status")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.additional_pod_ranges_config.pod_range_info")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.additional_pod_ranges_config.pod_range_names")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.auto_ipam_config")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.cluster_ipv4_cidr")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.create_subnetwork")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.default_pod_ipv4_range_utilization")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.ipv6_access_type")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.network_tier_config")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.node_ipv4_cidr")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.node_ipv4_cidr_block")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.pod_cidr_overprovision_config.disable")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.services_ipv4_cidr")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.services_ipv6_cidr_block")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.subnet_ipv6_cidr_block")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.subnetwork_name")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.tpu_ipv4_cidr_block")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.use_ip_aliases")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy.use_routes")
	f.Unimplemented_NotYetTriaged(".label_fingerprint")
	f.Unimplemented_NotYetTriaged(".legacy_abac.enabled")
	f.Unimplemented_NotYetTriaged(".locations")
	f.Unimplemented_NotYetTriaged(".logging_config.component_config")
	f.Unimplemented_NotYetTriaged(".maintenance_policy.disruption_budget")
	f.Unimplemented_NotYetTriaged(".maintenance_policy.resource_version")
	f.Unimplemented_NotYetTriaged(".maintenance_policy.window")
	f.Unimplemented_NotYetTriaged(".managed_machine_learning_diagnostics_config")
	f.Unimplemented_NotYetTriaged(".managed_opentelemetry_config")
	f.Unimplemented_NotYetTriaged(".master_auth.client_certificate_config")
	f.Unimplemented_NotYetTriaged(".master_authorized_networks_config.cidr_blocks")
	f.Unimplemented_NotYetTriaged(".master_authorized_networks_config.enabled")
	f.Unimplemented_NotYetTriaged(".master_authorized_networks_config.gcp_public_cidrs_access_enabled")
	f.Unimplemented_NotYetTriaged(".master_authorized_networks_config.private_endpoint_enforcement_enabled")
	f.Unimplemented_NotYetTriaged(".monitoring_config.component_config")
	f.Unimplemented_NotYetTriaged(".monitoring_config.managed_prometheus_config")
	f.Unimplemented_NotYetTriaged(".network_config.default_enable_private_nodes")
	f.Unimplemented_NotYetTriaged(".network_config.default_snat_status")
	f.Unimplemented_NotYetTriaged(".network_config.disable_l4_lb_firewall_reconciliation")
	f.Unimplemented_NotYetTriaged(".network_config.dns_config")
	f.Unimplemented_NotYetTriaged(".network_config.gateway_api_config")
	f.Unimplemented_NotYetTriaged(".network_config.network")
	f.Unimplemented_NotYetTriaged(".network_config.network_performance_config")
	f.Unimplemented_NotYetTriaged(".network_config.private_ipv6_google_access")
	f.Unimplemented_NotYetTriaged(".network_config.service_external_ips_config")
	f.Unimplemented_NotYetTriaged(".network_config.subnetwork")
	f.Unimplemented_NotYetTriaged(".node_config.accelerators")
	f.Unimplemented_NotYetTriaged(".node_config.advanced_machine_features")
	f.Unimplemented_NotYetTriaged(".node_config.boot_disk")
	f.Unimplemented_NotYetTriaged(".node_config.boot_disk_kms_key")
	f.Unimplemented_NotYetTriaged(".node_config.consolidation_delay")
	f.Unimplemented_NotYetTriaged(".node_config.disk_size_gb")
	f.Unimplemented_NotYetTriaged(".node_config.effective_cgroup_mode")
	f.Unimplemented_NotYetTriaged(".node_config.enable_confidential_storage")
	f.Unimplemented_NotYetTriaged(".node_config.flex_start")
	f.Unimplemented_NotYetTriaged(".node_config.gpu_direct_config")
	f.Unimplemented_NotYetTriaged(".node_config.gvnic")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.allowed_unsafe_sysctls")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.container_log_max_files")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.container_log_max_size")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.crash_loop_back_off")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.eviction_max_pod_grace_period_seconds")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.eviction_minimum_reclaim")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.eviction_soft")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.eviction_soft_grace_period")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.insecure_kubelet_readonly_port_enabled")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.max_parallel_image_pulls")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.memory_manager")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.shutdown_grace_period_critical_pods_seconds")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.shutdown_grace_period_seconds")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.single_process_oom_kill")
	f.Unimplemented_NotYetTriaged(".node_config.kubelet_config.topology_manager")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.accurate_time_config")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.custom_node_init")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.hugepages")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.node_kernel_module_loading")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.swap_config")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.sysctls")
	f.Unimplemented_NotYetTriaged(".node_config.linux_node_config.transparent_hugepage_defrag")
	f.Unimplemented_NotYetTriaged(".node_config.local_nvme_ssd_block_config")
	f.Unimplemented_NotYetTriaged(".node_config.local_ssd_encryption_mode")
	f.Unimplemented_NotYetTriaged(".node_config.logging_config")
	f.Unimplemented_NotYetTriaged(".node_config.max_run_duration")
	f.Unimplemented_NotYetTriaged(".node_config.reservation_affinity.values")
	f.Unimplemented_NotYetTriaged(".node_config.resource_manager_tags")
	f.Unimplemented_NotYetTriaged(".node_config.sandbox_config.type")
	f.Unimplemented_NotYetTriaged(".node_config.secondary_boot_disk_update_strategy")
	f.Unimplemented_NotYetTriaged(".node_config.secondary_boot_disks")
	f.Unimplemented_NotYetTriaged(".node_config.sole_tenant_config")
	f.Unimplemented_NotYetTriaged(".node_config.storage_pools")
	f.Unimplemented_NotYetTriaged(".node_config.taint_config")
	f.Unimplemented_NotYetTriaged(".node_config.taints")
	f.Unimplemented_NotYetTriaged(".node_config.windows_node_config")
	f.Unimplemented_NotYetTriaged(".node_config.workload_metadata_config")
	f.Unimplemented_NotYetTriaged(".node_creation_config")
	f.Unimplemented_NotYetTriaged(".node_ipv4_cidr_size")
	f.Unimplemented_NotYetTriaged(".node_pool_auto_config.linux_node_config")
	f.Unimplemented_NotYetTriaged(".node_pool_auto_config.network_tags.tags")
	f.Unimplemented_NotYetTriaged(".node_pool_auto_config.node_kubelet_config")
	f.Unimplemented_NotYetTriaged(".node_pool_auto_config.resource_manager_tags")
	f.Unimplemented_NotYetTriaged(".node_pool_defaults.node_config_defaults.containerd_config")
	f.Unimplemented_NotYetTriaged(".node_pool_defaults.node_config_defaults.logging_config")
	f.Unimplemented_NotYetTriaged(".node_pool_defaults.node_config_defaults.node_kubelet_config")
	f.Unimplemented_NotYetTriaged(".node_pools")
	f.Unimplemented_NotYetTriaged(".notification_config.pubsub.filter")
	f.Unimplemented_NotYetTriaged(".pod_autoscaling")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.master_global_access_config")
	f.Unimplemented_NotYetTriaged(".rbac_binding_config")
	f.Unimplemented_NotYetTriaged(".resource_labels")
	f.Unimplemented_NotYetTriaged(".resource_usage_export_config.bigquery_destination")
	f.Unimplemented_NotYetTriaged(".resource_usage_export_config.consumption_metering_config")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".schedule_upgrade_config")
	f.Unimplemented_NotYetTriaged(".secret_manager_config")
	f.Unimplemented_NotYetTriaged(".secret_sync_config")
	f.Unimplemented_NotYetTriaged(".security_posture_config.vulnerability_mode")
	f.Unimplemented_NotYetTriaged(".services_ipv4_cidr")
	f.Unimplemented_NotYetTriaged(".shielded_nodes")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".status_message")
	f.Unimplemented_NotYetTriaged(".tpu_ipv4_cidr_block")
	f.Unimplemented_NotYetTriaged(".user_managed_keys_config")
	f.Unimplemented_NotYetTriaged(".zone")

	// Unmapped / Unimplemented Status fields (specifically)
	f.Unimplemented_NotYetTriaged(".addons_config")
	f.Unimplemented_NotYetTriaged(".authenticator_groups_config")
	f.Unimplemented_NotYetTriaged(".binary_authorization")
	f.Unimplemented_NotYetTriaged(".cluster_ipv4_cidr")
	f.Unimplemented_NotYetTriaged(".confidential_nodes")
	f.Unimplemented_NotYetTriaged(".control_plane_endpoints_config.dns_endpoint_config.allow_external_traffic")
	f.Unimplemented_NotYetTriaged(".control_plane_endpoints_config.dns_endpoint_config.enable_k8s_certs_via_dns")
	f.Unimplemented_NotYetTriaged(".control_plane_endpoints_config.dns_endpoint_config.enable_k8s_tokens_via_dns")
	f.Unimplemented_NotYetTriaged(".control_plane_endpoints_config.ip_endpoints_config")
	f.Unimplemented_NotYetTriaged(".cost_management_config")
	f.Unimplemented_NotYetTriaged(".database_encryption")
	f.Unimplemented_NotYetTriaged(".description")
	f.Unimplemented_NotYetTriaged(".enable_k8s_beta_apis")
	f.Unimplemented_NotYetTriaged(".enable_kubernetes_alpha")
	f.Unimplemented_NotYetTriaged(".enable_tpu")
	f.Unimplemented_NotYetTriaged(".identity_service_config")
	f.Unimplemented_NotYetTriaged(".initial_node_count")
	f.Unimplemented_NotYetTriaged(".ip_allocation_policy")
	f.Unimplemented_NotYetTriaged(".legacy_abac")
	f.Unimplemented_NotYetTriaged(".location")
	f.Unimplemented_NotYetTriaged(".logging_config")
	f.Unimplemented_NotYetTriaged(".logging_service")
	f.Unimplemented_NotYetTriaged(".maintenance_policy")
	f.Unimplemented_NotYetTriaged(".master_auth.client_key")
	f.Unimplemented_NotYetTriaged(".master_auth.password")
	f.Unimplemented_NotYetTriaged(".master_auth.username")
	f.Unimplemented_NotYetTriaged(".master_authorized_networks_config")
	f.Unimplemented_NotYetTriaged(".mesh_certificates")
	f.Unimplemented_NotYetTriaged(".monitoring_config")
	f.Unimplemented_NotYetTriaged(".monitoring_service")
	f.Unimplemented_NotYetTriaged(".network")
	f.Unimplemented_NotYetTriaged(".network_config")
	f.Unimplemented_NotYetTriaged(".network_policy")
	f.Unimplemented_NotYetTriaged(".node_config")
	f.Unimplemented_NotYetTriaged(".node_pool_auto_config")
	f.Unimplemented_NotYetTriaged(".node_pool_defaults")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.enable_private_endpoint")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.enable_private_nodes")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.master_ipv4_cidr_block")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.peering_name")
	f.Unimplemented_NotYetTriaged(".private_cluster_config.private_endpoint_subnetwork")
	f.Unimplemented_NotYetTriaged(".release_channel")
	f.Unimplemented_NotYetTriaged(".resource_usage_export_config")
	f.Unimplemented_NotYetTriaged(".security_posture_config")
	f.Unimplemented_NotYetTriaged(".subnetwork")
	f.Unimplemented_NotYetTriaged(".vertical_pod_autoscaling")
	f.Unimplemented_NotYetTriaged(".workload_identity_config")

	f.FilterSpec = func(in *pb.Cluster) {
		in.SelfLink = ""
		in.SatisfiesPzi = nil
		in.SatisfiesPzs = nil
		if in.Autopilot != nil {
			in.Autopilot.ClusterPolicyConfig = nil
			if !in.Autopilot.Enabled {
				in.Autopilot = nil
			}
		}
		if in.DefaultMaxPodsConstraint != nil {
			if in.DefaultMaxPodsConstraint.MaxPodsPerNode == 0 {
				in.DefaultMaxPodsConstraint = nil
			}
		}
	}
	f.FilterStatus = func(in *pb.Cluster) {
		in.SelfLink = ""
		in.SatisfiesPzi = nil
		in.SatisfiesPzs = nil
		if in.Autopilot != nil {
			in.Autopilot = nil
		}
		if in.DefaultMaxPodsConstraint != nil {
			in.DefaultMaxPodsConstraint = nil
		}
	}

	return f
}
