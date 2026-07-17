// Copyright 2026 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.container.v1.NodePool
// api.group: container.cnrm.cloud.google.com

package container

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(containerNodePoolFuzzer())
}

func ContainerNodePoolStatus_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.ContainerNodePoolStatus {
	return &krm.ContainerNodePoolStatus{}
}

func ContainerNodePoolStatus_ToProto(mapCtx *direct.MapContext, in *krm.ContainerNodePoolStatus) *pb.NodePool {
	return &pb.NodePool{}
}

func containerNodePoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NodePool{},
		ContainerNodePoolSpec_FromProto, ContainerNodePoolSpec_ToProto,
		ContainerNodePoolStatus_FromProto, ContainerNodePoolStatus_ToProto,
	)

	// Identity and parent reference fields
	f.Unimplemented_Identity(".name")

	// Unimplemented top-level fields
	f.Unimplemented_NotYetTriaged(".self_link")
	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".status_message")
	f.Unimplemented_NotYetTriaged(".instance_group_urls")
	f.Unimplemented_NotYetTriaged(".pod_ipv4_cidr_size")
	f.Unimplemented_NotYetTriaged(".best_effort_provisioning")
	f.Unimplemented_NotYetTriaged(".node_drain_config")
	f.Unimplemented_NotYetTriaged(".maintenance_policy")
	f.Unimplemented_NotYetTriaged(".queued_provisioning")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".update_info")
	f.Unimplemented_NotYetTriaged(".conditions")
	f.Unimplemented_NotYetTriaged(".max_pods_constraint")

	// Unimplemented sub-fields in config (NodeConfig)
	f.Unimplemented_NotYetTriaged(".config.containerd_config")
	f.Unimplemented_NotYetTriaged(".config.local_nvme_ssd_block_config")
	f.Unimplemented_NotYetTriaged(".config.secondary_boot_disk_update_strategy")
	f.Unimplemented_NotYetTriaged(".config.taint_config")
	f.Unimplemented_NotYetTriaged(".config.workload_metadata_config")
	f.Unimplemented_NotYetTriaged(".config.max_run_duration")
	f.Unimplemented_NotYetTriaged(".config.local_ssd_encryption_mode")
	f.Unimplemented_NotYetTriaged(".config.effective_cgroup_mode")
	f.Unimplemented_NotYetTriaged(".config.flex_start")
	f.Unimplemented_NotYetTriaged(".config.boot_disk")
	f.Unimplemented_NotYetTriaged(".config.enable_confidential_storage")
	f.Unimplemented_NotYetTriaged(".config.secondary_boot_disks")
	f.Unimplemented_NotYetTriaged(".config.storage_pools")
	f.Unimplemented_NotYetTriaged(".config.logging_config")
	f.Unimplemented_NotYetTriaged(".config.sole_tenant_config")
	f.Unimplemented_NotYetTriaged(".config.resource_manager_tags")
	f.Unimplemented_NotYetTriaged(".config.gpu_sharing_config")
	f.Unimplemented_NotYetTriaged(".config.gpu_driver_byos_config")
	f.Unimplemented_NotYetTriaged(".config.advanced_machine_features.performance_monitoring_unit")
	f.Unimplemented_NotYetTriaged(".config.node_group")
	f.Unimplemented_NotYetTriaged(".config.service_account")
	f.Unimplemented_NotYetTriaged(".config.boot_disk_kms_key")
	f.Unimplemented_NotYetTriaged(".config.gpu_direct_config")
	f.SpecField(".config.kubelet_config")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.allowed_unsafe_sysctls")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.container_log_max_files")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.container_log_max_size")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.crash_loop_back_off")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.eviction_max_pod_grace_period_seconds")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.eviction_minimum_reclaim")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.eviction_soft")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.eviction_soft_grace_period")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.insecure_kubelet_readonly_port_enabled")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.max_parallel_image_pulls")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.memory_manager")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.shutdown_grace_period_critical_pods_seconds")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.shutdown_grace_period_seconds")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.single_process_oom_kill")
	f.Unimplemented_NotYetTriaged(".config.kubelet_config.topology_manager")
	f.Unimplemented_NotYetTriaged(".config.disk_size_gb")
	f.Unimplemented_NotYetTriaged(".config.taints")
	f.Unimplemented_NotYetTriaged(".config.gvnic")
	f.Unimplemented_NotYetTriaged(".config.sandbox_config.type")
	f.Unimplemented_NotYetTriaged(".config.consolidation_delay")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.hugepages")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.swap_config")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.accurate_time_config")
	f.Unimplemented_NotYetTriaged(".config.windows_node_config.os_version")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.node_kernel_module_loading")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.transparent_hugepage_enabled")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.transparent_hugepage_defrag")
	f.Unimplemented_NotYetTriaged(".config.linux_node_config.custom_node_init")
	f.Unimplemented_NotYetTriaged(".config.accelerators")

	// Unimplemented sub-fields in network_config
	f.Unimplemented_NotYetTriaged(".network_config.accelerator_network_profile")
	f.Unimplemented_NotYetTriaged(".network_config.network_tier_config")
	f.Unimplemented_NotYetTriaged(".network_config.pod_cidr_overprovision_config")
	f.Unimplemented_NotYetTriaged(".network_config.pod_ipv4_cidr_block")
	f.Unimplemented_NotYetTriaged(".network_config.pod_ipv4_range_utilization")
	f.Unimplemented_NotYetTriaged(".network_config.network_performance_config")
	f.Unimplemented_NotYetTriaged(".network_config.additional_pod_network_configs[].max_pods_per_node")

	// Unimplemented upgrade settings strategy and management options
	f.Unimplemented_NotYetTriaged(".upgrade_settings.strategy")
	f.Unimplemented_NotYetTriaged(".upgrade_settings.blue_green_settings.autoscaled_rollout_policy")
	f.Unimplemented_NotYetTriaged(".management.upgrade_options")
	f.Unimplemented_NotYetTriaged(".autoscaling.autoprovisioned")
	f.Unimplemented_NotYetTriaged(".autoscaling.enabled")

	// Fuzzer Spec fields
	f.SpecField(".autoscaling")
	f.SpecField(".initial_node_count")
	f.SpecField(".management")
	f.SpecField(".max_pods_constraint")
	f.SpecField(".network_config")
	f.SpecField(".config")
	f.SpecField(".locations")
	f.SpecField(".placement_policy")
	f.SpecField(".upgrade_settings")
	f.SpecField(".version")

	return f
}
