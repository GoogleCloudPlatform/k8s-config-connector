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

// +tool:fuzz-gen
// proto.message: google.cloud.dataplex.v1.Environment
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexEnvironmentFuzzer())
}

func dataplexEnvironmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Environment{},
		DataplexEnvironmentSpec_FromProto, DataplexEnvironmentSpec_ToProto,
		DataplexEnvironmentObservedState_FromProto, DataplexEnvironmentObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".infrastructure_spec")
	f.SpecFields.Insert(".session_spec")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".session_status")
	f.StatusFields.Insert(".endpoints")

	f.UnimplementedFields.Insert(".name")

	return f
}


