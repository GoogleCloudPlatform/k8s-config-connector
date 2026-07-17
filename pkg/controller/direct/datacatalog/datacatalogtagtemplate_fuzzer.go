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
// proto.message: google.cloud.datacatalog.v1.TagTemplate
// api.group: datacatalog.cnrm.cloud.google.com

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataCatalogTagTemplateFuzzer())
}

func dataCatalogTagTemplateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TagTemplate{},
		DataCatalogTagTemplateSpec_FromProto, DataCatalogTagTemplateSpec_ToProto,
		DataCatalogTagTemplateObservedState_FromProto, DataCatalogTagTemplateObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".is_publicly_readable")

	f.StatusFields.Insert(".dataplex_transfer_status")

	f.UnimplementedFields.Insert(".name")   // special field
	f.UnimplementedFields.Insert(".fields") // .fields.type is a oneof + enum combination. fuzzer sets both enums

	return f
}
