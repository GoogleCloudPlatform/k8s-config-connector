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
// proto.message: google.cloud.vmwareengine.v1.PrivateCloud
// api.group: vmwareengine.cnrm.cloud.google.com

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vmwareEnginePrivateCloudFuzzer())
}

func vmwareEnginePrivateCloudFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PrivateCloud{},
		VMwareEnginePrivateCloudSpec_FromProto, VMwareEnginePrivateCloudSpec_ToProto,
		VMwareEnginePrivateCloudObservedState_FromProto, VMwareEnginePrivateCloudObservedState_ToProto,
	)

	f.SpecFields.Insert(".network_config")
	f.SpecFields.Insert(".management_cluster")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".type")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".expire_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".network_config")
	f.StatusFields.Insert(".hcx")
	f.StatusFields.Insert(".nsx")
	f.StatusFields.Insert(".vcenter")
	f.StatusFields.Insert(".uid")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
