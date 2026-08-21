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
// proto.message: google.cloud.storageinsights.v1.DatasetConfig
// api.group: storageinsights.cnrm.cloud.google.com

package storageinsights

import (
	"strings"

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(StorageInsightsDatasetConfigFuzzer())
}

func StorageInsightsDatasetConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DatasetConfig{},
		StorageInsightsDatasetConfigSpec_FromProto, StorageInsightsDatasetConfigSpec_ToProto,
		StorageInsightsDatasetConfigObservedState_FromProto, StorageInsightsDatasetConfigObservedState_ToProto,
	)

	f.FilterSpec = func(in *pb.DatasetConfig) {
		if path, ok := in.SourceOptions.(*pb.DatasetConfig_CloudStorageObjectPath); ok {
			if !strings.HasPrefix(path.CloudStorageObjectPath, "gs://") {
				path.CloudStorageObjectPath = "gs://" + path.CloudStorageObjectPath
			}
		}
	}

	f.SpecField(".description")
	f.SpecField(".labels")
	f.SpecField(".organization_number")
	f.SpecField(".organization_scope")
	f.SpecField(".source_projects")
	f.SpecField(".source_folders")
	f.SpecField(".cloud_storage_object_path")
	f.SpecField(".include_cloud_storage_locations")
	f.SpecField(".exclude_cloud_storage_locations")
	f.SpecField(".include_cloud_storage_buckets")
	f.SpecField(".exclude_cloud_storage_buckets")
	f.SpecField(".include_newly_created_buckets")
	f.SpecField(".skip_verification_and_ingest")
	f.SpecField(".retention_period_days")
	f.SpecField(".identity")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".uid")
	f.StatusField(".link")
	f.StatusField(".dataset_config_state")

	f.Unimplemented_NotYetTriaged(".status")
	f.Unimplemented_NotYetTriaged(".identity.name")
	f.Unimplemented_Identity(".name")

	return f
}
