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
	"reflect"
	"sort"

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
				TargetTypes: append([]string{}, access.Dataset.TargetTypes...),
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

func foundDiffDatasetAccessEntry(a1, a2 []*bigquery.AccessEntry) bool {
	if len(a1) != len(a2) {
		return true
	}
	sortAccessEntries(a1)
	sortAccessEntries(a2)
	for i := range a1 {
		if a1[i].EntityType != bigquery.RoutineEntity && a1[i].EntityType != bigquery.ViewEntity && a1[i].EntityType != bigquery.DatasetEntity {
			if !reflect.DeepEqual(a1[i], a2[i]) {
				return true
			}
			continue
		}
		if reflect.DeepEqual(a1[i].View, a2[i].View) {
			return true
		}
		if reflect.DeepEqual(a1[i].Routine, a2[i].Routine) {
			return true
		}
		if reflect.DeepEqual(a1[i].Dataset, a2[i].Dataset) {
			return true
		}
	}
	return false
}

func sortAccessEntries(entries []*bigquery.AccessEntry) {
	if entries == nil {
		return
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Role != entries[j].Role {
			return entries[i].Role < entries[j].Role
		}
		if entries[i].EntityType != entries[j].EntityType {
			return entries[i].EntityType < entries[j].EntityType
		}
		return entries[i].Entity < entries[j].Entity
	})
}
