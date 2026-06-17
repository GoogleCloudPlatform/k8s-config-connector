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

package bigquery

import (
	"time"

	pb "cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/bigquerydataset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(bigQueryDatasetFuzzer())
}

func bigQueryDatasetFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.DatasetMetadata{},
		bigquerydataset.BigQueryDatasetSpec_FromProto, bigquerydataset.BigQueryDatasetSpec_ToProto,
		bigquerydataset.BigQueryDatasetStatus_FromProto, bigquerydataset.BigQueryDatasetStatus_ToProto,
	)

	f.SpecField(".Access")
	f.SpecField(".DefaultCollation")
	f.SpecField(".DefaultPartitionExpiration")
	f.SpecField(".DefaultTableExpiration")
	f.SpecField(".DefaultEncryptionConfig")
	f.SpecField(".Description")
	f.SpecField(".Location")
	f.SpecField(".MaxTimeTravel")
	f.SpecField(".IsCaseInsensitive")
	f.SpecField(".Name")
	f.SpecField(".StorageBillingModel")

	f.StatusField(".ETag")
	f.StatusField(".CreationTime")
	f.StatusField(".LastModifiedTime")

	f.IdentityField(".FullID")

	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".ExternalDatasetReference")
	f.Unimplemented_NotYetTriaged(".Tags")

	f.FilterSpec = func(in *pb.DatasetMetadata) {
		// MaxTimeTravel must be between 48 and 168 hours to avoid scientific notation
		hours := int64(in.MaxTimeTravel.Hours())
		if hours < 48 || hours > 168 {
			hours = 168
		}
		in.MaxTimeTravel = time.Duration(hours) * time.Hour

		in.DefaultPartitionExpiration = in.DefaultPartitionExpiration.Round(time.Millisecond)
		in.DefaultTableExpiration = in.DefaultTableExpiration.Round(time.Millisecond)

		// align Access Entry EntityType and fields
		for _, entry := range in.Access {
			entry.Condition = nil // Condition is not supported in KRM
			if entry.EntityType < 1 || entry.EntityType > 8 {
				entry.EntityType = pb.UserEmailEntity // Default/fallback
			}
			switch entry.EntityType {
			case pb.DomainEntity, pb.GroupEmailEntity, pb.UserEmailEntity, pb.SpecialGroupEntity, pb.IAMMemberEntity:
				entry.View = nil
				entry.Routine = nil
				entry.Dataset = nil
			case pb.ViewEntity:
				entry.Entity = ""
				entry.Routine = nil
				entry.Dataset = nil
				if entry.View == nil {
					entry.View = &pb.Table{}
				}
			case pb.RoutineEntity:
				entry.Entity = ""
				entry.View = nil
				entry.Dataset = nil
				if entry.Routine == nil {
					entry.Routine = &pb.Routine{}
				}
			case pb.DatasetEntity:
				entry.Entity = ""
				entry.View = nil
				entry.Routine = nil
				if entry.Dataset == nil {
					entry.Dataset = &pb.DatasetAccessEntry{
						Dataset: &pb.Dataset{},
					}
				}
			}
		}
	}

	f.FilterStatus = func(in *pb.DatasetMetadata) {
		in.CreationTime = in.CreationTime.Truncate(time.Millisecond)
		in.LastModifiedTime = in.LastModifiedTime.Truncate(time.Millisecond)
		if in.CreationTime.UnixMilli() <= 0 {
			in.CreationTime = time.UnixMilli(1)
		}
		if in.LastModifiedTime.UnixMilli() <= 0 {
			in.LastModifiedTime = time.UnixMilli(1)
		}
	}

	return f
}
