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

package bigquerydataset

import (
	"encoding/json"
	"fmt"

	bigquery "cloud.google.com/go/bigquery"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func convertProtoToAPI(u protoreflect.ProtoMessage, v any) error {
	if u == nil {
		return nil
	}

	j, err := protojson.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}

	if err := json.Unmarshal(j, v); err != nil {
		return fmt.Errorf("converting json to cloud API type: %w", err)
	}
	return nil
}

func convertAPIToProto[V protoreflect.ProtoMessage](u any, pV *V) error {
	if u == nil {
		return nil
	}

	j, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("converting proto to json: %w", err)
	}

	var v V
	if err := json.Unmarshal(j, &v); err != nil {
		return fmt.Errorf("converting json to proto type: %w", err)
	}
	*pV = v
	return nil
}
func cloneBigQueryDatasetMetadate(in *bigquery.DatasetMetadata) *bigquery.DatasetMetadata {
	if in == nil {
		return nil
	}
	out := &bigquery.DatasetMetadata{}
	acccessList := []*bigquery.AccessEntry{}
	for _, access := range in.Access {
		curAccess := &bigquery.AccessEntry{
			Role:       access.Role,
			EntityType: access.EntityType,
			Entity:     access.Entity,
			Condition:  access.Condition,
		}
		if access.View != nil {
			curAccess.View = &bigquery.Table{
				ProjectID: access.View.ProjectID,
				DatasetID: access.View.DatasetID,
				TableID:   access.View.TableID,
			}
		}
		if access.Routine != nil {
			curAccess.Routine = &bigquery.Routine{
				ProjectID: access.Routine.ProjectID,
				DatasetID: access.Routine.DatasetID,
				RoutineID: access.Routine.RoutineID,
			}
		}
		if access.Dataset != nil {
			curAccess.Dataset = &bigquery.DatasetAccessEntry{
				Dataset: &bigquery.Dataset{
					ProjectID: access.Dataset.Dataset.ProjectID,
					DatasetID: access.Dataset.Dataset.DatasetID,
				},
			}
		}
		acccessList = append(acccessList, curAccess)
	}
	out.Access = acccessList
	if in.DefaultEncryptionConfig != nil {
		out.DefaultEncryptionConfig = &bigquery.EncryptionConfig{
			KMSKeyName: in.DefaultEncryptionConfig.KMSKeyName,
		}
	}
	out.Location = in.Location
	// if the value to explicitly set to empty in the update request, we set the value.
	// Otherwise, we drop the value.
	if in.DefaultCollation != "" {
		out.DefaultCollation = in.DefaultCollation
	}
	if in.DefaultPartitionExpiration != 0 {
		out.DefaultPartitionExpiration = in.DefaultPartitionExpiration
	}
	if in.DefaultTableExpiration != 0 {
		out.DefaultTableExpiration = in.DefaultTableExpiration
	}
	if in.Description != "" {
		out.Description = in.Description
	}
	if in.MaxTimeTravel != 0 {
		out.MaxTimeTravel = in.MaxTimeTravel
	}
	out.IsCaseInsensitive = in.IsCaseInsensitive
	if in.Name != "" {
		out.Name = in.Name
	}
	if in.StorageBillingModel != "" {
		out.StorageBillingModel = in.StorageBillingModel
	}
	out.CreationTime = in.CreationTime
	out.LastModifiedTime = in.LastModifiedTime
	out.ETag = in.ETag
	out.FullID = in.FullID
	return out
}
