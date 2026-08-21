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
// proto.message: google.cloud.edgecontainer.v1.Machine
// api.group: edgecontainer.cnrm.cloud.google.com

package edgecontainer

import (
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(edgeContainerMachineFuzzer())
}

func edgeContainerMachineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Machine{},
		EdgeContainerMachineSpec_FromProto, EdgeContainerMachineSpec_ToProto,
		EdgeContainerMachineObservedState_FromProto, EdgeContainerMachineObservedState_ToProto,
	)

	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".zone")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".version")
	f.StatusFields.Insert(".disabled")
	f.StatusFields.Insert(".hosted_node")

	f.UnimplementedFields.Insert(".name")

	return f
}
