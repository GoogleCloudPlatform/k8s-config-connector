// Copyright 2025 Google LLC
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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzCloudDMSMappingRule())
}

func fuzzCloudDMSMappingRule() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MappingRule{},
		CloudDMSMappingRuleSpec_FromProto, CloudDMSMappingRuleSpec_ToProto,
		CloudDMSMappingRuleObservedState_FromProto, CloudDMSMappingRuleObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".rule_scope")
	f.SpecField(".filter")

	f.StatusField(".state")
	f.StatusField(".revision_id")
	f.StatusField(".revision_create_time")

	f.Unimplemented_Identity(".name")
	return f
}
