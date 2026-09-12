// Copyright 2024 Google LLC
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
// api.group: bigquery.cnrm.cloud.google.com

package bigquerydataset

import (
	"time"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	pb "cloud.google.com/go/bigquery"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(bigqueryDatasetFuzzer())
}

func bigqueryDatasetFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&pb.DatasetMetadata{},
		BigQueryDatasetSpec_FromProto, BigQueryDatasetSpec_ToProto,
		BigQueryDatasetStatus_FromProto, BigQueryDatasetStatus_ToProto,
	)

	f.FilterStatus = func(in *pb.DatasetMetadata) {
		in.CreationTime = time.UnixMilli(1600000000000)
		in.LastModifiedTime = time.UnixMilli(1600000000000)
		in.FullID = "project:dataset"
	}
	f.CmpOptions = []cmp.Option{
		cmpopts.IgnoreUnexported(pb.Table{}, pb.Routine{}, pb.Dataset{}),
	}
	f.FilterSpec = func(in *pb.DatasetMetadata) {
		// BigQueryDatasetSpec rounds MaxTimeTravel to hours and converts to string via fmt.Sprintf("%v").
		// Constrain MaxTimeTravel to valid BigQuery range (48 to 168 hours) to avoid parsing issues
		in.MaxTimeTravel = 168 * time.Hour

		// BigQueryDatasetSpec rounds expirations to milliseconds
		in.DefaultPartitionExpiration = in.DefaultPartitionExpiration.Truncate(time.Millisecond)
		in.DefaultTableExpiration = in.DefaultTableExpiration.Truncate(time.Millisecond)

		// Fix up Access
		for _, access := range in.Access {
			if access == nil {
				continue
			}

			// Valid entity types:
			// DomainEntity, GroupEmailEntity, UserEmailEntity, SpecialGroupEntity, ViewEntity, IAMMemberEntity, RoutineEntity, DatasetEntity
			types := []pb.EntityType{
				pb.DomainEntity, pb.GroupEmailEntity, pb.UserEmailEntity, pb.SpecialGroupEntity,
				pb.ViewEntity, pb.IAMMemberEntity, pb.RoutineEntity, pb.DatasetEntity,
			}
			access.EntityType = types[uint64(access.EntityType)%uint64(len(types))]

			// Clear out non-matching fields
			switch access.EntityType {
			case pb.DomainEntity, pb.GroupEmailEntity, pb.UserEmailEntity, pb.SpecialGroupEntity, pb.IAMMemberEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case pb.ViewEntity:
				access.Entity = ""
				access.Routine = nil
				access.Dataset = nil
			case pb.RoutineEntity:
				access.Entity = ""
				access.View = nil
				access.Dataset = nil
			case pb.DatasetEntity:
				access.Entity = ""
				access.View = nil
				access.Routine = nil
			}
		}
	}

	f.SpecField(".Name")
	f.SpecField(".Description")
	f.SpecField(".Location")
	f.SpecField(".DefaultTableExpiration")
	f.SpecField(".Access")
	f.SpecField(".Access.Role")
	f.SpecField(".Access.EntityType")
	f.SpecField(".Access.Entity")
	f.SpecField(".Access.View")
	f.SpecField(".Access.View.ProjectID")
	f.SpecField(".Access.View.DatasetID")
	f.SpecField(".Access.View.TableID")
	f.SpecField(".Access.Routine")
	f.SpecField(".Access.Routine.ProjectID")
	f.SpecField(".Access.Routine.DatasetID")
	f.SpecField(".Access.Routine.RoutineID")
	f.SpecField(".Access.Dataset")
	f.SpecField(".Access.Dataset.Dataset")
	f.SpecField(".Access.Dataset.Dataset.ProjectID")
	f.SpecField(".Access.Dataset.Dataset.DatasetID")
	f.SpecField(".Access.Dataset.TargetTypes")
	f.SpecField(".DefaultEncryptionConfig")
	f.SpecField(".DefaultEncryptionConfig.KMSKeyName")
	f.SpecField(".DefaultPartitionExpiration")
	f.SpecField(".DefaultCollation")
	f.SpecField(".MaxTimeTravel")
	f.SpecField(".StorageBillingModel")
	f.SpecField(".IsCaseInsensitive")

	f.StatusField(".ETag")
	f.StatusField(".CreationTime")
	f.StatusField(".LastModifiedTime")
	f.StatusField(".FullID")

	f.Unimplemented_NotYetTriaged(".Condition")
	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Tags")
	f.Unimplemented_NotYetTriaged(".ExternalDatasetReference")
	f.Unimplemented_NotYetTriaged(".View.c")
	f.Unimplemented_NotYetTriaged(".Routine.c")
	f.Unimplemented_NotYetTriaged(".Dataset.Dataset.c")

	return f
}
