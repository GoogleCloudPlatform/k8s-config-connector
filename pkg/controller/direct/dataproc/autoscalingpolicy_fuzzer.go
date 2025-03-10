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
// proto.message: google.cloud.dataproc.v1.AutoscalingPolicy
// api.group: dataproc.cnrm.cloud.google.com

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(autoscalingPolicyFuzzer())
}

func autoscalingPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AutoscalingPolicy{},
		AutoscalingPolicySpec_FromProto, AutoscalingPolicySpec_ToProto,
		AutoscalingPolicyObservedState_FromProto, AutoscalingPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".basic_algorithm")
	f.SpecFields.Insert(".worker_config")
	f.SpecFields.Insert(".secondary_worker_config")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".name")
	f.SpecFields.Insert("." + util.AutoscalingPolicyIdFieldPath)

	return f
}


