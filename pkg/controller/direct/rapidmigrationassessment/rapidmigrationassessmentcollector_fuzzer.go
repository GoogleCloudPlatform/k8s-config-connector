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
// proto.message: google.cloud.rapidmigrationassessment.v1.Collector
// api.group: rapidmigrationassessment.cnrm.cloud.google.com

package rapidmigrationassessment

import (
	pb "cloud.google.com/go/rapidmigrationassessment/apiv1/rapidmigrationassessmentpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(rapidMigrationAssessmentCollectorFuzzer())
}

func rapidMigrationAssessmentCollectorFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Collector{},
		RapidMigrationAssessmentCollectorSpec_FromProto, RapidMigrationAssessmentCollectorSpec_ToProto,
		RapidMigrationAssessmentCollectorObservedState_FromProto, RapidMigrationAssessmentCollectorObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".service_account")
	f.SpecField(".expected_asset_count")
	f.SpecField(".collection_days")
	f.SpecField(".eula_uri")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".bucket")
	f.StatusField(".state")
	f.StatusField(".client_version")
	f.StatusField(".guest_os_scan")
	f.StatusField(".vsphere_scan")

	f.Unimplemented_Identity(".name")

	f.FilterSpec = func(in *pb.Collector) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.Collector) {
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
