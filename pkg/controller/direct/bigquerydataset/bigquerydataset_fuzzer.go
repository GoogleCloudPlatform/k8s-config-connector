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

package bigquerydataset

import (
	"time"

	bigquery "cloud.google.com/go/bigquery"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(bigQueryDatasetFuzzer())
}

func bigQueryDatasetFuzzer() fuzztesting.KRMFuzzer_NoProto {
	// Note: We must use the _NoProto form of the KRM fuzzer because
	// bigquery.DatasetMetadata (from cloud.google.com/go/bigquery) is a
	// handwritten Go client library struct, not a protobuf message, and
	// therefore does not implement proto.Message.
	f := fuzztesting.NewKRMTypedFuzzer_NoProto[*bigquery.DatasetMetadata, krm.BigQueryDatasetSpec, krm.BigQueryDatasetStatus](&bigquery.DatasetMetadata{},
		BigQueryDatasetSpec_FromProto, BigQueryDatasetSpec_ToProto,
		BigQueryDatasetStatus_FromProto, BigQueryDatasetStatus_ToProto,
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
	f.StatusField(".FullID")

	f.Unimplemented_NotYetTriaged(".Labels")
	f.Unimplemented_NotYetTriaged(".Tags")
	f.Unimplemented_NotYetTriaged(".ExternalDatasetReference")

	f.FilterSpec = func(in *bigquery.DatasetMetadata) {
		// Zero out status-only fields
		in.CreationTime = time.Time{}
		in.LastModifiedTime = time.Time{}
		in.ETag = ""
		in.FullID = ""

		// Round durations to match the KRM millisecond/hour precision
		in.DefaultPartitionExpiration = (in.DefaultPartitionExpiration / time.Millisecond) * time.Millisecond
		in.DefaultTableExpiration = (in.DefaultTableExpiration / time.Millisecond) * time.Millisecond

		// Restrict MaxTimeTravel to between 48 and 168 hours to avoid formatting issues with scientific notation
		hours := int64(in.MaxTimeTravel.Hours())
		if hours < 48 || hours > 168 {
			hours = (hours % 121)
			if hours < 0 {
				hours = -hours
			}
			hours += 48
		}
		in.MaxTimeTravel = time.Duration(hours) * time.Hour

		// Normalize access entries to respect KRM single-entity union logic and unmapped fields
		for _, access := range in.Access {
			if access == nil {
				continue
			}
			switch access.EntityType {
			case bigquery.DomainEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case bigquery.GroupEmailEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case bigquery.UserEmailEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case bigquery.SpecialGroupEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case bigquery.ViewEntity:
				access.Entity = ""
				access.Routine = nil
				access.Dataset = nil
			case bigquery.IAMMemberEntity:
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			case bigquery.RoutineEntity:
				access.Entity = ""
				access.View = nil
				access.Dataset = nil
			case bigquery.DatasetEntity:
				access.Entity = ""
				access.View = nil
				access.Routine = nil
			default:
				access.EntityType = bigquery.UserEmailEntity
				access.View = nil
				access.Routine = nil
				access.Dataset = nil
			}
			access.Condition = nil
		}
	}

	f.FilterStatus = func(in *bigquery.DatasetMetadata) {
		// Zero out everything except status fields
		in.Name = ""
		in.Description = ""
		in.Location = ""
		in.DefaultTableExpiration = 0
		in.DefaultPartitionExpiration = 0
		in.DefaultCollation = ""
		in.DefaultEncryptionConfig = nil
		in.MaxTimeTravel = 0
		in.StorageBillingModel = ""
		in.IsCaseInsensitive = false
		in.Access = nil
		in.Labels = nil
		in.Tags = nil
		in.ExternalDatasetReference = nil

		// Restrict dates to between year 1970 and 2200 to avoid UnixMillisecond overflow (which happens in year 2262)
		creationMillis := in.CreationTime.UnixMilli()
		if creationMillis < 0 || creationMillis > 7258118400000 { // Year 2200 is 7258118400000
			creationMillis = (creationMillis % 7258118400000)
			if creationMillis < 0 {
				creationMillis = -creationMillis
			}
		}
		in.CreationTime = time.UnixMilli(creationMillis).UTC()

		modifiedMillis := in.LastModifiedTime.UnixMilli()
		if modifiedMillis < 0 || modifiedMillis > 7258118400000 {
			modifiedMillis = (modifiedMillis % 7258118400000)
			if modifiedMillis < 0 {
				modifiedMillis = -modifiedMillis
			}
		}
		in.LastModifiedTime = time.UnixMilli(modifiedMillis).UTC()

		// FullID needs to match the projectID:datasetID form so selflink parses successfully in status.ToProto
		in.FullID = "projectid:datasetid"
	}

	return f
}
