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
// proto.message: google.cloud.orchestration.airflow.service.v1.Environment
// api.group: composer.cnrm.cloud.google.com

package composer

import (
	pb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(composerEnvironmentFuzzer())
}

func composerEnvironmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Environment{},
		ComposerEnvironmentSpec_FromProto, ComposerEnvironmentSpec_ToProto,
		ComposerEnvironmentObservedState_FromProto, ComposerEnvironmentObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")          // special field
	f.UnimplementedFields.Insert(".satisfies_pzs") // field for future use
	f.UnimplementedFields.Insert(".satisfies_pzi") // field for future use

	f.SpecFields.Insert(".config")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".storage_config")

	f.StatusFields.Insert(".config")
	f.StatusFields.Insert(".uuid")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")

	return f
}
