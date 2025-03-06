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
// proto.message: google.cloud.tpu.v1.Node
// api.group: tpu.cnrm.cloud.google.com

package tpu

import (
	pb "cloud.google.com/go/tpu/apiv1/tpupb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(TPUNodeFuzzer())
}

func TPUNodeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Node{},
		TPUNodeSpec_FromProto, TPUNodeSpec_ToProto,
		TPUNodeStatus_FromProto, TPUNodeStatus_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".accelerator_type")
	f.SpecFields.Insert(".tensorflow_version")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".cidr_block")
	f.SpecFields.Insert(".scheduling_config")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".use_service_networking")

	f.StatusFields.Insert(".ip_address")
	f.StatusFields.Insert(".port")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".health_description")
	f.StatusFields.Insert(".service_account")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".network_endpoints")
	f.StatusFields.Insert(".health")
	f.StatusFields.Insert(".api_version")
	f.StatusFields.Insert(".symptoms")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
