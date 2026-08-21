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

package resourcemanager

import (
	"strings"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzProject())
}

func fuzzProject() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Project{},
		ProjectSpec_FromProto, ProjectSpec_ToProto,
		ProjectStatus_FromProto, ProjectStatus_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".parent")
	f.SpecField(".project_id")

	f.StatusField(".name")

	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_NotYetTriaged(".etag")
	f.Unimplemented_NotYetTriaged(".labels")

	f.FilterSpec = func(in *pb.Project) {
		if in.Parent != "" {
			// Ensure parent starts with folders/ or organizations/
			if in.Parent[0]%2 == 0 {
				in.Parent = "folders/" + sanitizeID(in.Parent)
			} else {
				in.Parent = "organizations/" + sanitizeID(in.Parent)
			}
		}
	}

	f.FilterStatus = func(in *pb.Project) {
		if in.Name != "" {
			in.Name = "projects/" + sanitizeID(in.Name)
		}
	}

	return f
}

func sanitizeID(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			sb.WriteRune(r)
		}
	}
	if sb.Len() == 0 {
		return "12345"
	}
	return sb.String()
}
