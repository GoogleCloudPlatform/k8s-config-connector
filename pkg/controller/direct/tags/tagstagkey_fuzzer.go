// Copyright 2025 Google LLC
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
// proto.message: google.cloud.resourcemanager.v3.TagKey
// api.group: tags.cnrm.cloud.google.com

package tags

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(tagsTagKeyFuzzer())
}

func tagsTagKeyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TagKey{},
		TagsTagKeySpec_FromProto, TagsTagKeySpec_ToProto,
		TagsTagKeyStatus_FromProto, TagsTagKeyStatus_ToProto,
	)

	f.SpecField(".parent")
	f.SpecField(".short_name")
	f.SpecField(".description")
	f.SpecField(".purpose")
	f.SpecField(".purpose_data")

	f.StatusField(".name")
	f.StatusField(".namespaced_name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Etag()

	return f
}
