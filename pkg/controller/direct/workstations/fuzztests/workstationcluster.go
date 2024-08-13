// Copyright 2024 Google LLC
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

package fuzztests

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/workstations"
)

func init() {
	RegisterKRMFuzzer(workstationclusterFuzzer())
}

func workstationclusterFuzzer() KRMFuzzer {
	f := NewKRMTypedFuzzer(&pb.WorkstationCluster{},
		workstations.WorkstationsWorkstationClusterSpec_FromProto, workstations.WorkstationsWorkstationClusterSpec_ToProto,
		workstations.WorkstationsWorkstationClusterObservedState_FromProto, workstations.WorkstationsWorkstationClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")
	f.UnimplementedFields.Insert(".private_cluster_config.cluster_hostname")
	f.UnimplementedFields.Insert(".private_cluster_config.service_attachment_uri")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".private_cluster_config")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".subnetwork")
	f.SpecFields.Insert(".network")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".control_plane_ip")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".uid")

	return f
}
