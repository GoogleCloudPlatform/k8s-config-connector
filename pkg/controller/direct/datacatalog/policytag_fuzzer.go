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
// proto.message: google.cloud.datacatalog.v1beta1.PolicyTag
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1beta1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogPolicyTagFuzzer())
}

func dataCatalogPolicyTagFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.PolicyTag{},
		DataCatalogPolicyTagSpec_FromProto, DataCatalogPolicyTagSpec_ToProto,
		DataCatalogPolicyTagObservedState_FromProto, DataCatalogPolicyTagObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".parent_policy_tag")

	f.UnimplementedFields.Insert(".name") // special field
	f.UnimplementedFields.Insert(".child_policy_tags")

	return f
}
