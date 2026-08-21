// Copyright 2026 Google LLC
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
// proto.message: google.cloud.dataplex.v1.MetadataJob
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexMetadataJobFuzzer())
}

func dataplexMetadataJobFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.MetadataJob{},
		DataplexMetadataJobSpec_FromProto, DataplexMetadataJobSpec_ToProto,
		DataplexMetadataJobObservedState_FromProto, DataplexMetadataJobObservedState_ToProto,
	)

	f.SpecField(".type")
	f.SpecField(".import_spec")
	f.SpecField(".export_spec")

	f.StatusField(".uid")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".import_result")
	f.StatusField(".export_result")
	f.StatusField(".status")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	// Unmapped fields in ImportJobResult
	f.Unimplemented_NotYetTriaged(".import_result.deleted_entry_links")
	f.Unimplemented_NotYetTriaged(".import_result.unchanged_entry_links")
	f.Unimplemented_NotYetTriaged(".import_result.created_entry_links")
	f.Unimplemented_NotYetTriaged(".import_result.recreated_entry_links")

	// Unmapped fields in ImportJobScope
	f.Unimplemented_NotYetTriaged(".import_spec.scope.glossaries")
	f.Unimplemented_NotYetTriaged(".import_spec.scope.entry_link_types")
	f.Unimplemented_NotYetTriaged(".import_spec.scope.referenced_entry_scopes")

	// Unmapped fields in ExportJobScope
	f.Unimplemented_NotYetTriaged(".export_spec.scope.projects")
	f.Unimplemented_NotYetTriaged(".export_spec.scope.entry_types")
	f.Unimplemented_NotYetTriaged(".export_spec.scope.aspect_types")

	return f
}
