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
// proto.message: google.cloud.dataproc.v1.NodeGroup
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataprocNodeGroupFuzzer())
}

func dataprocNodeGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NodeGroup{},
		DataprocNodeGroupSpec_FromProto, DataprocNodeGroupSpec_ToProto,
		DataprocNodeGroupObservedState_FromProto, DataprocNodeGroupObservedState_ToProto,
	)

	f.SpecFields.Insert(".roles")
	f.SpecFields.Insert(".node_group_config")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".node_group_config")

	f.UnimplementedFields.Insert(".name") // special field
	return f
}
