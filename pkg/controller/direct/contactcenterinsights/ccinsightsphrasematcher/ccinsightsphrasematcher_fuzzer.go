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
// proto.message: google.cloud.contactcenterinsights.v1.PhraseMatcher
// api.group: contactcenterinsights.cnrm.cloud.google.com

package ccinsightsphrasematcher

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(ccinsightsPhraseMatcherFuzzer())
}

func ccinsightsPhraseMatcherFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PhraseMatcher{},
		CCInsightsPhraseMatcherSpec_FromProto, CCInsightsPhraseMatcherSpec_ToProto,
		CCInsightsPhraseMatcherObservedState_FromProto, CCInsightsPhraseMatcherObservedState_ToProto,
	)

	f.SpecField(".version_tag")
	f.SpecField(".display_name")
	f.SpecField(".type")
	f.SpecField(".active")
	f.SpecField(".phrase_match_rule_groups")
	f.SpecField(".role_match")

	f.StatusField(".revision_id")
	f.StatusField(".revision_create_time")
	f.StatusField(".activation_update_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")

	return f
}
