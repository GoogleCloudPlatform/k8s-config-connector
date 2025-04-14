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
// proto.message: google.cloud.discoveryengine.v1.TargetSite

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataStoreTargetSiteFuzzer())
}

func dataStoreTargetSiteFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetSite{},
		DataStoreTargetSiteSpec_FromProto, DataStoreTargetSiteSpec_ToProto,
		TargetSiteObservedState_FromProto, TargetSiteObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".exact_match")
	f.SpecFields.Insert(".provided_uri_pattern")
	f.SpecFields.Insert(".type")

	f.StatusFields.Insert(".generated_uri_pattern")
	f.StatusFields.Insert(".root_domain_uri")
	f.StatusFields.Insert(".site_verification_info")
	f.StatusFields.Insert(".indexing_status")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".failure_reason")

	return f
}
