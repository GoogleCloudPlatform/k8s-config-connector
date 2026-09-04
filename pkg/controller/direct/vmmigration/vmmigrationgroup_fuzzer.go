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
// proto.message: google.cloud.vmmigration.v1.Group
// api.group: vmmigration.cnrm.cloud.google.com

package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vMmigrationGroupFuzzer())
}

func vMmigrationGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Group{},
		VMMigrationGroupSpec_FromProto, VMMigrationGroupSpec_ToProto,
		VMMigrationGroupObservedState_FromProto, VMMigrationGroupObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".migration_target_type")

	f.FilterSpec = func(in *pb.Group) {
		cleanEmptyMessages(in.ProtoReflect())
	}

	f.FilterStatus = func(in *pb.Group) {
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
