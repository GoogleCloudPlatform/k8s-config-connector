// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.tpu.v2.Node
// crd.kind: TPUVirtualMachine

package tpu

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/tpu/v2"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(virtualMachineFuzzer())
}

func virtualMachineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Node{},
		TPUVirtualMachineSpec_FromProto, TPUVirtualMachineSpec_ToProto,
		TPUVirtualMachineObservedState_FromProto, TPUVirtualMachineObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".accelerator_type")
	f.SpecFields.Insert(".runtime_version")
	f.SpecFields.Insert(".network_config")
	f.SpecFields.Insert(".network_configs")
	f.SpecFields.Insert(".cidr_block")
	f.SpecFields.Insert(".service_account")
	f.SpecFields.Insert(".scheduling_config")
	f.SpecFields.Insert(".metadata")
	f.SpecFields.Insert(".tags")
	f.SpecFields.Insert(".data_disks")
	f.SpecFields.Insert(".shielded_instance_config")
	f.SpecFields.Insert(".accelerator_config")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".health")
	f.StatusFields.Insert(".network_endpoints")
	f.StatusFields.Insert(".health_description")
	f.StatusFields.Insert(".symptoms")
	f.StatusFields.Insert(".queued_resource")
	f.StatusFields.Insert(".multislice_node")

	// network_endpoints are tricky, they mix spec and status
	f.SpecFields.Insert(".network_endpoints")
	f.StatusFields.Insert(".network_endpoints.access_config")

	// System / naming fields
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".id")

	// Labels
	f.UnimplementedFields.Insert(".labels")

	// Fields that are not good KRM/gitops API design
	f.UnimplementedFields.Insert(".api_version")

	// Volatile fields we don't want to implement
	f.UnimplementedFields.Insert(".create_time")

	return f
}
