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
// proto.message: google.cloud.talent.v4.Company
// api.group: cloudtalentsolution.cnrm.cloud.google.com

package cloudtalentsolution

import (
	pb "cloud.google.com/go/talent/apiv4/talentpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudTalentSolutionCompanyFuzzer())
}

func cloudTalentSolutionCompanyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Company{},
		CloudTalentSolutionCompanySpec_FromProto, CloudTalentSolutionCompanySpec_ToProto,
		CloudTalentSolutionCompanyObservedState_FromProto, CloudTalentSolutionCompanyObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".external_id")
	f.SpecField(".size")
	f.SpecField(".headquarters_address")
	f.SpecField(".hiring_agency")
	f.SpecField(".eeo_text")
	f.SpecField(".website_uri")
	f.SpecField(".career_site_uri")
	f.SpecField(".image_uri")
	f.SpecField(".keyword_searchable_job_custom_attributes")

	f.StatusField(".derived_info")
	f.StatusField(".suspended")

	f.Unimplemented_Identity(".name")

	f.FilterSpec = func(in *pb.Company) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.Company) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	return f
}

func cleanEmptyMessages(m protoreflect.Message) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsList() || fd.IsMap() {
				return true
			}
			sub := v.Message()
			cleanEmptyMessages(sub)
			// check if sub has any populated fields now
			hasFields := false
			sub.Range(func(fd2 protoreflect.FieldDescriptor, v2 protoreflect.Value) bool {
				hasFields = true
				return false
			})
			if !hasFields {
				m.Clear(fd)
			}
		}
		return true
	})
}
