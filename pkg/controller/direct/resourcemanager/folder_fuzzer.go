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
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzFolder())
}

func fuzzFolder() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Folder{},
		FolderSpec_FromProto, FolderSpec_ToProto,
		FolderStatus_FromProto, FolderStatus_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".parent")

	f.StatusField(".name")
	f.StatusField(".state")

	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".delete_time")
	f.Unimplemented_NotYetTriaged(".etag")

	f.FilterSpec = func(in *pb.Folder) {
		if in.Parent != "" {
			// Ensure parent starts with folders/ or organizations/
			if in.Parent[0]%2 == 0 {
				in.Parent = "folders/" + sanitizeID(in.Parent)
			} else {
				in.Parent = "organizations/" + sanitizeID(in.Parent)
			}
		}
	}

	f.FilterStatus = func(in *pb.Folder) {
		if in.Name != "" {
			in.Name = "folders/" + sanitizeID(in.Name)
		}
	}

	return f
}
