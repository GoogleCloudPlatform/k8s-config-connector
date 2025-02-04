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

package datastream

import (
	pb "cloud.google.com/go/datastream/apiv1/datastreampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datastream/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DatastreamStreamSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatastreamStreamSpec) *pb.Stream {
	if in == nil {
		return nil
	}
	out := &pb.Stream{}
	out.Labels = in.Labels
	out.SourceConfig = SourceConfig_ToProto(mapCtx, in.SourceConfig)
	out.DestinationConfig = DestinationConfig_ToProto(mapCtx, in.DestinationConfig)
	if in.CustomerManagedEncryptionKeyRef != nil {
		out.CustomerManagedEncryptionKey = direct.LazyPtr(in.CustomerManagedEncryptionKeyRef.External)
	}
	if oneof := Stream_BackfillAllStrategy_ToProto(mapCtx, in.BackfillAll); oneof != nil {
		out.BackfillStrategy = &pb.Stream_BackfillAll{BackfillAll: oneof}
	}
	if oneof := Stream_BackfillNoneStrategy_ToProto(mapCtx, in.BackfillNone); oneof != nil {
		out.BackfillStrategy = &pb.Stream_BackfillNone{BackfillNone: oneof}
	}
	return out
}

func SourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.SourceConfigSpec) *pb.SourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.SourceConfig{}
	if in.SourceConnectionProfileRef != nil {
		out.SourceConnectionProfile = in.SourceConnectionProfileRef.External
	}
	if oneof := OracleSourceConfig_ToProto(mapCtx, in.OracleSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_OracleSourceConfig{OracleSourceConfig: oneof}
	}
	if oneof := MysqlSourceConfig_ToProto(mapCtx, in.MysqlSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_MysqlSourceConfig{MysqlSourceConfig: oneof}
	}
	if oneof := PostgresqlSourceConfig_ToProto(mapCtx, in.PostgresqlSourceConfig); oneof != nil {
		out.SourceStreamConfig = &pb.SourceConfig_PostgresqlSourceConfig{PostgresqlSourceConfig: oneof}
	}

	return out
}

func BigQueryDestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryDestinationConfig) *krm.BigQueryDestinationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDestinationConfigSpec{}
	if oneof := BigQueryDestinationConfig_SingleTargetDataset_FromProto(mapCtx, in.GetSingleTargetDataset()); oneof != nil {
		out.SingleTargetDataset = oneof
	}
	if oneof := BigQueryDestinationConfig_SourceHierarchyDatasets_FromProto(mapCtx, in.GetSourceHierarchyDatasets()); oneof != nil {
		out.SourceHierarchyDatasets = oneof
	}
	out.DataFreshness = direct.StringDuration_FromProto(mapCtx, in.GetDataFreshness())
	return out
}
func BigQueryDestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDestinationConfigSpec) *pb.BigQueryDestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryDestinationConfig{}
	if oneof := BigQueryDestinationConfig_SingleTargetDataset_ToProto(mapCtx, in.SingleTargetDataset); oneof != nil {
		out.DatasetConfig = &pb.BigQueryDestinationConfig_SingleTargetDataset_{SingleTargetDataset: oneof}
	}
	if oneof := BigQueryDestinationConfig_SourceHierarchyDatasets_ToProto(mapCtx, in.SourceHierarchyDatasets); oneof != nil {
		out.DatasetConfig = &pb.BigQueryDestinationConfig_SourceHierarchyDatasets_{SourceHierarchyDatasets: oneof}
	}
	out.DataFreshness = direct.StringDuration_ToProto(mapCtx, in.DataFreshness)
	return out
}
func SourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.SourceConfig) *krm.SourceConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SourceConfigSpec{}
	if in.GetSourceConnectionProfile() != "" {
		out.SourceConnectionProfileRef = &krm.ConnectionProfileRef{
			External: in.GetSourceConnectionProfile(),
		}
	}
	direct.LazyPtr(in.GetSourceConnectionProfile())
	if oneof := OracleSourceConfig_FromProto(mapCtx, in.GetOracleSourceConfig()); oneof != nil {
		out.OracleSourceConfig = oneof
	}
	if oneof := MysqlSourceConfig_FromProto(mapCtx, in.GetMysqlSourceConfig()); oneof != nil {
		out.MysqlSourceConfig = oneof
	}
	if oneof := PostgresqlSourceConfig_FromProto(mapCtx, in.GetPostgresqlSourceConfig()); oneof != nil {
		out.PostgresqlSourceConfig = oneof
	}
	return out
}

func DatastreamStreamSpec_FromProto(mapCtx *direct.MapContext, in *pb.Stream) *krm.DatastreamStreamSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatastreamStreamSpec{}
	out.Labels = in.Labels
	out.SourceConfig = SourceConfig_FromProto(mapCtx, in.GetSourceConfig())
	out.DestinationConfig = DestinationConfig_FromProto(mapCtx, in.GetDestinationConfig())
	out.BackfillAll = Stream_BackfillAllStrategy_FromProto(mapCtx, in.GetBackfillAll())
	out.BackfillNone = Stream_BackfillNoneStrategy_FromProto(mapCtx, in.GetBackfillNone())
	// MISSING: LastRecoveryTime
	return out
}

func MysqlSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MysqlSourceConfig) *krm.MysqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.MysqlSourceConfig{}
	if oneof := MysqlRdbms_FromProto(mapCtx, in.GetIncludeObjects()); oneof != nil {
		out.IncludeObjects = oneof
	}
	if oneof := MysqlRdbms_FromProto(mapCtx, in.GetExcludeObjects()); oneof != nil {
		out.ExcludeObjects = oneof
	}
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	return out
}
func MysqlSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.MysqlSourceConfig) *pb.MysqlSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.MysqlSourceConfig{}
	out.IncludeObjects = MysqlRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = MysqlRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	return out
}

func OracleSourceConfig_FromProto(mapCtx *direct.MapContext, in *pb.OracleSourceConfig) *krm.OracleSourceConfig {
	if in == nil {
		return nil
	}
	out := &krm.OracleSourceConfig{}
	if oneof := OracleRdbms_FromProto(mapCtx, in.GetIncludeObjects()); oneof != nil {
		out.IncludeObjects = oneof
	}
	if oneof := OracleRdbms_FromProto(mapCtx, in.GetExcludeObjects()); oneof != nil {
		out.ExcludeObjects = oneof
	}
	out.MaxConcurrentCdcTasks = direct.LazyPtr(in.GetMaxConcurrentCdcTasks())
	out.MaxConcurrentBackfillTasks = direct.LazyPtr(in.GetMaxConcurrentBackfillTasks())
	if oneof := OracleSourceConfig_DropLargeObjects_FromProto(mapCtx, in.GetDropLargeObjects()); oneof != nil {
		out.DropLargeObjects = oneof
	}
	if oneof := OracleSourceConfig_StreamLargeObjects_FromProto(mapCtx, in.GetStreamLargeObjects()); oneof != nil {
		out.StreamLargeObjects = oneof
	}
	return out
}

func OracleSourceConfig_ToProto(mapCtx *direct.MapContext, in *krm.OracleSourceConfig) *pb.OracleSourceConfig {
	if in == nil {
		return nil
	}
	out := &pb.OracleSourceConfig{}
	out.IncludeObjects = OracleRdbms_ToProto(mapCtx, in.IncludeObjects)
	out.ExcludeObjects = OracleRdbms_ToProto(mapCtx, in.ExcludeObjects)
	out.MaxConcurrentCdcTasks = direct.ValueOf(in.MaxConcurrentCdcTasks)
	out.MaxConcurrentBackfillTasks = direct.ValueOf(in.MaxConcurrentBackfillTasks)
	if oneof := OracleSourceConfig_DropLargeObjects_ToProto(mapCtx, in.DropLargeObjects); oneof != nil {
		out.LargeObjectsHandling = &pb.OracleSourceConfig_DropLargeObjects_{DropLargeObjects: oneof}
	}
	if oneof := OracleSourceConfig_StreamLargeObjects_ToProto(mapCtx, in.StreamLargeObjects); oneof != nil {
		out.LargeObjectsHandling = &pb.OracleSourceConfig_StreamLargeObjects_{StreamLargeObjects: oneof}
	}
	return out
}

func Stream_BackfillAllStrategy_ToProto(mapCtx *direct.MapContext, in *krm.Stream_BackfillAllStrategy) *pb.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &pb.Stream_BackfillAllStrategy{}
	if oneof := OracleRdbms_ToProto(mapCtx, in.OracleExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_OracleExcludedObjects{OracleExcludedObjects: oneof}
	}
	if oneof := MysqlRdbms_ToProto(mapCtx, in.MysqlExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_MysqlExcludedObjects{MysqlExcludedObjects: oneof}
	}
	if oneof := PostgresqlRdbms_ToProto(mapCtx, in.PostgresqlExcludedObjects); oneof != nil {
		out.ExcludedObjects = &pb.Stream_BackfillAllStrategy_PostgresqlExcludedObjects{PostgresqlExcludedObjects: oneof}
	}
	return out
}

func Stream_BackfillAllStrategy_FromProto(mapCtx *direct.MapContext, in *pb.Stream_BackfillAllStrategy) *krm.Stream_BackfillAllStrategy {
	if in == nil {
		return nil
	}
	out := &krm.Stream_BackfillAllStrategy{}
	if oneof := OracleRdbms_FromProto(mapCtx, in.GetOracleExcludedObjects()); oneof != nil {
		out.OracleExcludedObjects = oneof
	}
	if oneof := MysqlRdbms_FromProto(mapCtx, in.GetMysqlExcludedObjects()); oneof != nil {
		out.MysqlExcludedObjects = oneof
	}
	if oneof := PostgresqlRdbms_FromProto(mapCtx, in.GetPostgresqlExcludedObjects()); oneof != nil {
		out.PostgresqlExcludedObjects = oneof
	}
	return out
}

func DestinationConfig_FromProto(mapCtx *direct.MapContext, in *pb.DestinationConfig) *krm.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &krm.DestinationConfig{}
	if in.GetDestinationConnectionProfile() != "" {
		out.DestinationConnectionProfileRef = &krm.ConnectionProfileRef{
			External: in.GetDestinationConnectionProfile(),
		}
	}
	out.GcsDestinationConfig = GcsDestinationConfig_FromProto(mapCtx, in.GetGcsDestinationConfig())
	out.BigqueryDestinationConfig = BigQueryDestinationConfig_FromProto(mapCtx, in.GetBigqueryDestinationConfig())
	return out
}

func DestinationConfig_ToProto(mapCtx *direct.MapContext, in *krm.DestinationConfig) *pb.DestinationConfig {
	if in == nil {
		return nil
	}
	out := &pb.DestinationConfig{}
	if in.DestinationConnectionProfileRef != nil {
		out.DestinationConnectionProfile = in.DestinationConnectionProfileRef.External
	}
	if oneof := GcsDestinationConfig_ToProto(mapCtx, in.GcsDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_GcsDestinationConfig{GcsDestinationConfig: oneof}
	}
	if oneof := BigQueryDestinationConfig_ToProto(mapCtx, in.BigqueryDestinationConfig); oneof != nil {
		out.DestinationStreamConfig = &pb.DestinationConfig_BigqueryDestinationConfig{BigqueryDestinationConfig: oneof}
	}
	return out
}
