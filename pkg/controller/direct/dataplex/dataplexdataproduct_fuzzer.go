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
// proto.message: google.cloud.dataplex.v1.DataProduct
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexDataProductFuzzer())
}

func dataplexDataProductFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataProduct{},
		DataplexDataProductSpec_FromProto, DataplexDataProductSpec_ToProto,
		DataplexDataProductObservedState_FromProto, DataplexDataProductObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".labels")
	f.SpecField(".description")
	f.SpecField(".owner_emails")
	f.SpecField(".access_approval_config")
	f.SpecField(".access_approval_config.approver_emails")
	f.SpecField(".access_groups")
	f.SpecField(".access_groups[].principal.google_group")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".asset_count")
	f.StatusField(".etag")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".icon")

	f.FilterSpec = func(in *pb.DataProduct) {
		for k, v := range in.AccessGroups {
			v.Id = k
			v.DisplayName = k
			v.Description = ""
			if v.Principal != nil {
				v.Principal.ServiceAccount = nil
			}
		}
	}

	return f
}
