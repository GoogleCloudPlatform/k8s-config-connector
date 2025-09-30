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
// proto.message: google.cloud.resourcemanager.v3.TagBinding
// api.group: tags.cnrm.cloud.google.com
// crd.kind: LocationTagBinding

package tags

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(tagsLocationTagBindingFuzzer())
}

func tagsLocationTagBindingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TagBinding{},
		TagsLocationTagBindingSpec_FromProto, TagsLocationTagBindingSpec_ToProto,
		TagsLocationTagBindingObservedState_FromProto, TagsLocationTagBindingObservedState_ToProto,
	)

	f.SpecFields.Insert(".parent")
	f.SpecFields.Insert(".tag_value")

	f.UnimplementedFields.Insert(".name")                      // Output only
	f.UnimplementedFields.Insert(".tag_value_namespaced_name") // Cannot be set together with tag_value

	return f
}
