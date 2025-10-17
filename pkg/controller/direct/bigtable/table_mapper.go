// Copyright 2025 Google LLC
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

package bigtable

import (
	"strconv"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigtableTableSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.BigtableTableSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigtableTableSpec{}

	// Note: Bigtable proto 1.38 -> 1.40 changed the ColumnFamily from a slice to a map
	if in.GetColumnFamilies() != nil {
		out.ColumnFamily = []*krm.TableColumnFamily{}
		for _, v := range in.GetColumnFamilies() {
			if v == nil {
				continue
			}
			if cf := TableColumnFamily_v1beta1_FromProto(mapCtx, v); cf != nil {
				out.ColumnFamily = append(out.ColumnFamily, cf)
			}
		}
	}

	// Note: Bigtable proto 1.38 -> 1.40 changed the ChangeStreamRetention from a single field to a struct
	if changeStreamConfig := ChangeStreamConfig_v1beta1_FromProto(mapCtx, in.GetChangeStreamConfig()); changeStreamConfig != nil {
		out.ChangeStreamRetention = changeStreamConfig.RetentionPeriod
	}

	// Note: Bigtable proto 1.38 -> 1.40 changed the DeletionProtection type from string to bool; we handle the conversion.
	s := strconv.FormatBool(in.GetDeletionProtection())
	out.DeletionProtection = &s

	// MISSING: Granularity
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func BigtableTableSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableTableSpec) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}

	// Note: Bigtable proto 1.38 -> 1.40 changed the ColumnFamily from a slice to a map
	if in.ColumnFamily != nil {
		out.ColumnFamilies = map[string]*pb.ColumnFamily{}
		for _, v := range in.ColumnFamily {
			if v == nil {
				continue
			}
			out.ColumnFamilies[v.FamilyID] = TableColumnFamily_v1beta1_ToProto(mapCtx, v)
		}
	}

	// Note: Bigtable proto 1.38 -> 1.40 changed the ChangeStreamRetention from a single field to a struct
	if in.ChangeStreamRetention != nil {
		out.ChangeStreamConfig = &pb.ChangeStreamConfig{
			RetentionPeriod: direct.Duration_ToProto(mapCtx, in.ChangeStreamRetention),
		}
	}

	// Note: Bigtable proto 1.38 -> 1.40 changed the DeletionProtection type from string to bool; we handle the conversion.
	out.DeletionProtection, _ = strconv.ParseBool(direct.ValueOf(in.DeletionProtection))

	// MISSING: Granularity
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}

func BigtableTableObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krm.BigtableTableObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigtableTableObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// TODO: map type string message for field ClusterStates
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	if restoreInfo := RestoreInfo_v1beta1_FromProto(mapCtx, in.GetRestoreInfo()); restoreInfo != nil {
		out.RestoreInfo = restoreInfo
	}

	if in.GetClusterStates() != nil {
		out.ClusterStates = map[string]krm.Table_ClusterState{}
		for k, v := range in.GetClusterStates() {
			if v == nil {
				continue
			}
			if cs := Table_ClusterState_v1beta1_FromProto(mapCtx, v); cs != nil {
				out.ClusterStates[k] = *cs
			}
		}
	}
	// MISSING: ChangeStreamConfig
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
func BigtableTableObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.BigtableTableObservedState) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}
	out.Name = direct.ValueOf(in.Name)

	if in.ClusterStates != nil {
		out.ClusterStates = map[string]*pb.Table_ClusterState{}
		for k, v := range in.ClusterStates {
			out.ClusterStates[k] = Table_ClusterState_v1beta1_ToProto(mapCtx, &v)
		}
	}
	// MISSING: ColumnFamilies
	// MISSING: Granularity
	if in.RestoreInfo != nil {
		out.RestoreInfo = RestoreInfo_v1beta1_ToProto(mapCtx, in.RestoreInfo)
	}
	// MISSING: ChangeStreamConfig
	// MISSING: AutomatedBackupPolicy
	// MISSING: RowKeySchema
	return out
}
