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
	"time"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/durationpb"
)

func BigtableTableSpec_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.BigtableTableSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigtableTableSpec{}

	if retentionPeriod := in.GetChangeStreamConfig().GetRetentionPeriod(); retentionPeriod != nil {
		out.ChangeStreamRetention = durationFromProto(mapCtx, retentionPeriod)
	}

	if in.GetDeletionProtection() {
		out.DeletionProtection = PtrTo("PROTECTED")
	}

	if in.ColumnFamilies != nil {
		for k := range in.ColumnFamilies {
			columnFamily := krmv1beta1.Table_ColumnFamilies{
				Family: PtrTo(k),
			}
			out.ColumnFamily = append(out.ColumnFamily, columnFamily)
		}
	}
	// out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())

	// out.AutomatedBackupPolicy = Table_AutomatedBackupPolicy_FromProto(mapCtx, in.GetAutomatedBackupPolicy())

	return out
}

func BigtableTableSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigtableTableSpec) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}

	if s := ValueOf(in.ChangeStreamRetention); s != "" && s != "0" {
		retentionPeriod := durationToProto(mapCtx, s)

		out.ChangeStreamConfig = &pb.ChangeStreamConfig{
			RetentionPeriod: retentionPeriod,
		}
	}

	switch ValueOf(in.DeletionProtection) {
	case "PROTECTED":
		out.DeletionProtection = true

	case "UNPROTECTED", "":
		out.DeletionProtection = false

	default:
		mapCtx.Errorf("invalid value for deletionProtection %q", ValueOf(in.DeletionProtection))
	}

	if in.ColumnFamily != nil {
		out.ColumnFamilies = make(map[string]*pb.ColumnFamily)
		for _, columnFamily := range in.ColumnFamily {
			id := ValueOf(columnFamily.Family)
			out.ColumnFamilies[id] = &pb.ColumnFamily{}
		}
	}

	// out.Granularity = direct.Enum_ToProto[pb.Table_TimestampGranularity](mapCtx, in.Granularity)
	// out.ChangeStreamConfig = ChangeStreamConfig_ToProto(mapCtx, in.ChangeStreamConfig)
	// if oneof := Table_AutomatedBackupPolicy_ToProto(mapCtx, in.AutomatedBackupPolicy); oneof != nil {
	// 	out.AutomatedBackupConfig = &pb.Table_AutomatedBackupPolicy_{AutomatedBackupPolicy: oneof}
	// }

	return out
}

func BigtableTableStatus_FromProto(mapCtx *direct.MapContext, in *pb.Table) *krmv1beta1.BigtableTableStatus {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigtableTableStatus{}

	return out
}

func BigtableTableStatus_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigtableTableStatus) *pb.Table {
	if in == nil {
		return nil
	}
	out := &pb.Table{}

	return out
}

func durationToProto(mapCtx *direct.MapContext, in string) *durationpb.Duration {
	if in == "" {
		return nil
	}

	d, err := time.ParseDuration(in)
	if err != nil {
		mapCtx.Errorf("invalid duration %q: %w", in, err)
		return nil
	}
	return durationpb.New(d)
}

func durationFromProto(mapCtx *direct.MapContext, in *durationpb.Duration) *string {
	if in == nil {
		return nil
	}
	s := in.AsDuration().String()
	return &s
}
