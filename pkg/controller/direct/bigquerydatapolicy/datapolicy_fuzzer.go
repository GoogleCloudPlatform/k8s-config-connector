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
// proto.message: google.cloud.bigquery.datapolicies.v1beta1.DataPolicy
// api.group: bigquerydatapolicies.cnrm.cloud.google.com

package bigquerydatapolicy

import (
	pb "cloud.google.com/go/bigquery/datapolicies/apiv1beta1/datapoliciespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigQueryDataPolicyFuzzer())
}

func bigQueryDataPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataPolicy{},
		BigQueryDataPolicySpec_FromProto, BigQueryDataPolicySpec_ToProto,
		BigQueryDataPolicyObservedState_FromProto, BigQueryDataPolicyObservedState_ToProto,
	)

	f.SpecFields.Insert(".data_policy_type")
	f.SpecFields.Insert(".data_policy_id") // KRM resourceID
	f.SpecFields.Insert(".policy_tag")
	f.SpecFields.Insert(".data_masking_policy.predefined_expression")

	f.StatusFields.Insert(".name")

	f.UnimplementedFields.Insert(".name") // Special field for KRM resource name

	return f
}
