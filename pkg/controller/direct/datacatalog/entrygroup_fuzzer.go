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
// proto.message: google.cloud.datacatalog.v1.EntryGroup
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogEntryGroupFuzzer())
}

func dataCatalogEntryGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EntryGroup{},
		DataCatalogEntryGroupSpec_FromProto, DataCatalogEntryGroupSpec_ToProto,
		DataCatalogEntryGroupObservedState_FromProto, DataCatalogEntryGroupObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".transferred_to_dataplex")

	f.StatusFields.Insert(".data_catalog_timestamps")

	f.UnimplementedFields.Insert(".name")                                // special field
	f.UnimplementedFields.Insert(".data_catalog_timestamps.expire_time") // this is moved to observed state
	return f
}
