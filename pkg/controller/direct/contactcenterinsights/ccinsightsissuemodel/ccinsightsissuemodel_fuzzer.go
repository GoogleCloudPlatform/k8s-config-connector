// Copyright 2026 Google LLC
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
// proto.message: google.cloud.contactcenterinsights.v1.IssueModel
// api.group: contactcenterinsights.cnrm.cloud.google.com

package ccinsightsissuemodel

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(ccinsightsIssueModelFuzzer())
}

func ccinsightsIssueModelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.IssueModel{},
		CCInsightsIssueModelSpec_FromProto, CCInsightsIssueModelSpec_ToProto,
		CCInsightsIssueModelObservedState_FromProto, CCInsightsIssueModelObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".model_type")
	f.SpecField(".language_code")
	f.SpecField(".input_data_config.medium")
	f.SpecField(".input_data_config.filter")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".issue_count")
	f.StatusField(".state")
	f.StatusField(".input_data_config.training_conversations_count")
	f.StatusField(".training_stats")
	f.StatusField(".training_stats.analyzed_conversations_count")
	f.StatusField(".training_stats.unclassified_conversations_count")

	f.Unimplemented_NotYetTriaged(".training_stats.issue_stats")

	f.Unimplemented_Identity(".name")

	return f
}
