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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ApplyHash_FromProto(mapCtx *direct.MapContext, in *pb.ApplyHash) *krm.ApplyHash {
	if in == nil {
		return nil
	}
	out := &krm.ApplyHash{}
	out.UuidFromBytes = Empty_FromProto(mapCtx, in.GetUuidFromBytes())
	return out
}
func ApplyHash_ToProto(mapCtx *direct.MapContext, in *krm.ApplyHash) *pb.ApplyHash {
	if in == nil {
		return nil
	}
	out := &pb.ApplyHash{}
	if oneof := Empty_ToProto(mapCtx, in.UuidFromBytes); oneof != nil {
		out.HashFunction = &pb.ApplyHash_UuidFromBytes{UuidFromBytes: oneof}
	}
	return out
}
func AssignSpecificValue_FromProto(mapCtx *direct.MapContext, in *pb.AssignSpecificValue) *krm.AssignSpecificValue {
	if in == nil {
		return nil
	}
	out := &krm.AssignSpecificValue{}
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func AssignSpecificValue_ToProto(mapCtx *direct.MapContext, in *krm.AssignSpecificValue) *pb.AssignSpecificValue {
	if in == nil {
		return nil
	}
	out := &pb.AssignSpecificValue{}
	out.Value = direct.ValueOf(in.Value)
	return out
}
func ClouddmsConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ClouddmsConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConversionWorkspaceObservedState{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ClouddmsConversionWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConversionWorkspaceSpec{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConversionWorkspaceSpec) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsMappingRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MappingRule) *krm.ClouddmsMappingRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMappingRuleObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func ClouddmsMappingRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMappingRuleObservedState) *pb.MappingRule {
	if in == nil {
		return nil
	}
	out := &pb.MappingRule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func ClouddmsMappingRuleSpec_FromProto(mapCtx *direct.MapContext, in *pb.MappingRule) *krm.ClouddmsMappingRuleSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMappingRuleSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func ClouddmsMappingRuleSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMappingRuleSpec) *pb.MappingRule {
	if in == nil {
		return nil
	}
	out := &pb.MappingRule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func ClouddmsMigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobSpec) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.ClouddmsPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsPrivateConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.ClouddmsPrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsPrivateConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsPrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ConditionalColumnSetValue_FromProto(mapCtx *direct.MapContext, in *pb.ConditionalColumnSetValue) *krm.ConditionalColumnSetValue {
	if in == nil {
		return nil
	}
	out := &krm.ConditionalColumnSetValue{}
	out.SourceTextFilter = SourceTextFilter_FromProto(mapCtx, in.GetSourceTextFilter())
	out.SourceNumericFilter = SourceNumericFilter_FromProto(mapCtx, in.GetSourceNumericFilter())
	out.ValueTransformation = ValueTransformation_FromProto(mapCtx, in.GetValueTransformation())
	out.CustomFeatures = CustomFeatures_FromProto(mapCtx, in.GetCustomFeatures())
	return out
}
func ConditionalColumnSetValue_ToProto(mapCtx *direct.MapContext, in *krm.ConditionalColumnSetValue) *pb.ConditionalColumnSetValue {
	if in == nil {
		return nil
	}
	out := &pb.ConditionalColumnSetValue{}
	if oneof := SourceTextFilter_ToProto(mapCtx, in.SourceTextFilter); oneof != nil {
		out.SourceFilter = &pb.ConditionalColumnSetValue_SourceTextFilter{SourceTextFilter: oneof}
	}
	if oneof := SourceNumericFilter_ToProto(mapCtx, in.SourceNumericFilter); oneof != nil {
		out.SourceFilter = &pb.ConditionalColumnSetValue_SourceNumericFilter{SourceNumericFilter: oneof}
	}
	out.ValueTransformation = ValueTransformation_ToProto(mapCtx, in.ValueTransformation)
	out.CustomFeatures = CustomFeatures_ToProto(mapCtx, in.CustomFeatures)
	return out
}
func ConvertRowIdToColumn_FromProto(mapCtx *direct.MapContext, in *pb.ConvertRowIdToColumn) *krm.ConvertRowIdToColumn {
	if in == nil {
		return nil
	}
	out := &krm.ConvertRowIdToColumn{}
	out.OnlyIfNoPrimaryKey = direct.LazyPtr(in.GetOnlyIfNoPrimaryKey())
	return out
}
func ConvertRowIdToColumn_ToProto(mapCtx *direct.MapContext, in *krm.ConvertRowIdToColumn) *pb.ConvertRowIdToColumn {
	if in == nil {
		return nil
	}
	out := &pb.ConvertRowIdToColumn{}
	out.OnlyIfNoPrimaryKey = direct.ValueOf(in.OnlyIfNoPrimaryKey)
	return out
}
func DoubleComparisonFilter_FromProto(mapCtx *direct.MapContext, in *pb.DoubleComparisonFilter) *krm.DoubleComparisonFilter {
	if in == nil {
		return nil
	}
	out := &krm.DoubleComparisonFilter{}
	out.ValueComparison = direct.Enum_FromProto(mapCtx, in.GetValueComparison())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func DoubleComparisonFilter_ToProto(mapCtx *direct.MapContext, in *krm.DoubleComparisonFilter) *pb.DoubleComparisonFilter {
	if in == nil {
		return nil
	}
	out := &pb.DoubleComparisonFilter{}
	out.ValueComparison = direct.Enum_ToProto[pb.ValueComparison](mapCtx, in.ValueComparison)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func EntityMove_FromProto(mapCtx *direct.MapContext, in *pb.EntityMove) *krm.EntityMove {
	if in == nil {
		return nil
	}
	out := &krm.EntityMove{}
	out.NewSchema = direct.LazyPtr(in.GetNewSchema())
	return out
}
func EntityMove_ToProto(mapCtx *direct.MapContext, in *krm.EntityMove) *pb.EntityMove {
	if in == nil {
		return nil
	}
	out := &pb.EntityMove{}
	out.NewSchema = direct.ValueOf(in.NewSchema)
	return out
}
func FilterTableColumns_FromProto(mapCtx *direct.MapContext, in *pb.FilterTableColumns) *krm.FilterTableColumns {
	if in == nil {
		return nil
	}
	out := &krm.FilterTableColumns{}
	out.IncludeColumns = in.IncludeColumns
	out.ExcludeColumns = in.ExcludeColumns
	return out
}
func FilterTableColumns_ToProto(mapCtx *direct.MapContext, in *krm.FilterTableColumns) *pb.FilterTableColumns {
	if in == nil {
		return nil
	}
	out := &pb.FilterTableColumns{}
	out.IncludeColumns = in.IncludeColumns
	out.ExcludeColumns = in.ExcludeColumns
	return out
}
func IntComparisonFilter_FromProto(mapCtx *direct.MapContext, in *pb.IntComparisonFilter) *krm.IntComparisonFilter {
	if in == nil {
		return nil
	}
	out := &krm.IntComparisonFilter{}
	out.ValueComparison = direct.Enum_FromProto(mapCtx, in.GetValueComparison())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func IntComparisonFilter_ToProto(mapCtx *direct.MapContext, in *krm.IntComparisonFilter) *pb.IntComparisonFilter {
	if in == nil {
		return nil
	}
	out := &pb.IntComparisonFilter{}
	out.ValueComparison = direct.Enum_ToProto[pb.ValueComparison](mapCtx, in.ValueComparison)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func MappingRule_FromProto(mapCtx *direct.MapContext, in *pb.MappingRule) *krm.MappingRule {
	if in == nil {
		return nil
	}
	out := &krm.MappingRule{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RuleScope = direct.Enum_FromProto(mapCtx, in.GetRuleScope())
	out.Filter = MappingRuleFilter_FromProto(mapCtx, in.GetFilter())
	out.RuleOrder = direct.LazyPtr(in.GetRuleOrder())
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	out.SingleEntityRename = SingleEntityRename_FromProto(mapCtx, in.GetSingleEntityRename())
	out.MultiEntityRename = MultiEntityRename_FromProto(mapCtx, in.GetMultiEntityRename())
	out.EntityMove = EntityMove_FromProto(mapCtx, in.GetEntityMove())
	out.SingleColumnChange = SingleColumnChange_FromProto(mapCtx, in.GetSingleColumnChange())
	out.MultiColumnDataTypeChange = MultiColumnDatatypeChange_FromProto(mapCtx, in.GetMultiColumnDataTypeChange())
	out.ConditionalColumnSetValue = ConditionalColumnSetValue_FromProto(mapCtx, in.GetConditionalColumnSetValue())
	out.ConvertRowidColumn = ConvertRowIdToColumn_FromProto(mapCtx, in.GetConvertRowidColumn())
	out.SetTablePrimaryKey = SetTablePrimaryKey_FromProto(mapCtx, in.GetSetTablePrimaryKey())
	out.SinglePackageChange = SinglePackageChange_FromProto(mapCtx, in.GetSinglePackageChange())
	out.SourceSqlChange = SourceSqlChange_FromProto(mapCtx, in.GetSourceSqlChange())
	out.FilterTableColumns = FilterTableColumns_FromProto(mapCtx, in.GetFilterTableColumns())
	return out
}
func MappingRule_ToProto(mapCtx *direct.MapContext, in *krm.MappingRule) *pb.MappingRule {
	if in == nil {
		return nil
	}
	out := &pb.MappingRule{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.State = direct.Enum_ToProto[pb.MappingRule_State](mapCtx, in.State)
	out.RuleScope = direct.Enum_ToProto[pb.DatabaseEntityType](mapCtx, in.RuleScope)
	out.Filter = MappingRuleFilter_ToProto(mapCtx, in.Filter)
	out.RuleOrder = direct.ValueOf(in.RuleOrder)
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	if oneof := SingleEntityRename_ToProto(mapCtx, in.SingleEntityRename); oneof != nil {
		out.Details = &pb.MappingRule_SingleEntityRename{SingleEntityRename: oneof}
	}
	if oneof := MultiEntityRename_ToProto(mapCtx, in.MultiEntityRename); oneof != nil {
		out.Details = &pb.MappingRule_MultiEntityRename{MultiEntityRename: oneof}
	}
	if oneof := EntityMove_ToProto(mapCtx, in.EntityMove); oneof != nil {
		out.Details = &pb.MappingRule_EntityMove{EntityMove: oneof}
	}
	if oneof := SingleColumnChange_ToProto(mapCtx, in.SingleColumnChange); oneof != nil {
		out.Details = &pb.MappingRule_SingleColumnChange{SingleColumnChange: oneof}
	}
	if oneof := MultiColumnDatatypeChange_ToProto(mapCtx, in.MultiColumnDataTypeChange); oneof != nil {
		out.Details = &pb.MappingRule_MultiColumnDataTypeChange{MultiColumnDataTypeChange: oneof}
	}
	if oneof := ConditionalColumnSetValue_ToProto(mapCtx, in.ConditionalColumnSetValue); oneof != nil {
		out.Details = &pb.MappingRule_ConditionalColumnSetValue{ConditionalColumnSetValue: oneof}
	}
	if oneof := ConvertRowIdToColumn_ToProto(mapCtx, in.ConvertRowidColumn); oneof != nil {
		out.Details = &pb.MappingRule_ConvertRowidColumn{ConvertRowidColumn: oneof}
	}
	if oneof := SetTablePrimaryKey_ToProto(mapCtx, in.SetTablePrimaryKey); oneof != nil {
		out.Details = &pb.MappingRule_SetTablePrimaryKey{SetTablePrimaryKey: oneof}
	}
	if oneof := SinglePackageChange_ToProto(mapCtx, in.SinglePackageChange); oneof != nil {
		out.Details = &pb.MappingRule_SinglePackageChange{SinglePackageChange: oneof}
	}
	if oneof := SourceSqlChange_ToProto(mapCtx, in.SourceSqlChange); oneof != nil {
		out.Details = &pb.MappingRule_SourceSqlChange{SourceSqlChange: oneof}
	}
	if oneof := FilterTableColumns_ToProto(mapCtx, in.FilterTableColumns); oneof != nil {
		out.Details = &pb.MappingRule_FilterTableColumns{FilterTableColumns: oneof}
	}
	return out
}
func MappingRuleFilter_FromProto(mapCtx *direct.MapContext, in *pb.MappingRuleFilter) *krm.MappingRuleFilter {
	if in == nil {
		return nil
	}
	out := &krm.MappingRuleFilter{}
	out.ParentEntity = direct.LazyPtr(in.GetParentEntity())
	out.EntityNamePrefix = direct.LazyPtr(in.GetEntityNamePrefix())
	out.EntityNameSuffix = direct.LazyPtr(in.GetEntityNameSuffix())
	out.EntityNameContains = direct.LazyPtr(in.GetEntityNameContains())
	out.Entities = in.Entities
	return out
}
func MappingRuleFilter_ToProto(mapCtx *direct.MapContext, in *krm.MappingRuleFilter) *pb.MappingRuleFilter {
	if in == nil {
		return nil
	}
	out := &pb.MappingRuleFilter{}
	out.ParentEntity = direct.ValueOf(in.ParentEntity)
	out.EntityNamePrefix = direct.ValueOf(in.EntityNamePrefix)
	out.EntityNameSuffix = direct.ValueOf(in.EntityNameSuffix)
	out.EntityNameContains = direct.ValueOf(in.EntityNameContains)
	out.Entities = in.Entities
	return out
}
func MappingRuleObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MappingRule) *krm.MappingRuleObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MappingRuleObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func MappingRuleObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MappingRuleObservedState) *pb.MappingRule {
	if in == nil {
		return nil
	}
	out := &pb.MappingRule{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: RuleScope
	// MISSING: Filter
	// MISSING: RuleOrder
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	// MISSING: SingleEntityRename
	// MISSING: MultiEntityRename
	// MISSING: EntityMove
	// MISSING: SingleColumnChange
	// MISSING: MultiColumnDataTypeChange
	// MISSING: ConditionalColumnSetValue
	// MISSING: ConvertRowidColumn
	// MISSING: SetTablePrimaryKey
	// MISSING: SinglePackageChange
	// MISSING: SourceSqlChange
	// MISSING: FilterTableColumns
	return out
}
func MultiColumnDatatypeChange_FromProto(mapCtx *direct.MapContext, in *pb.MultiColumnDatatypeChange) *krm.MultiColumnDatatypeChange {
	if in == nil {
		return nil
	}
	out := &krm.MultiColumnDatatypeChange{}
	out.SourceDataTypeFilter = direct.LazyPtr(in.GetSourceDataTypeFilter())
	out.SourceTextFilter = SourceTextFilter_FromProto(mapCtx, in.GetSourceTextFilter())
	out.SourceNumericFilter = SourceNumericFilter_FromProto(mapCtx, in.GetSourceNumericFilter())
	out.NewDataType = direct.LazyPtr(in.GetNewDataType())
	out.OverrideLength = direct.LazyPtr(in.GetOverrideLength())
	out.OverrideScale = direct.LazyPtr(in.GetOverrideScale())
	out.OverridePrecision = direct.LazyPtr(in.GetOverridePrecision())
	out.OverrideFractionalSecondsPrecision = direct.LazyPtr(in.GetOverrideFractionalSecondsPrecision())
	out.CustomFeatures = CustomFeatures_FromProto(mapCtx, in.GetCustomFeatures())
	return out
}
func MultiColumnDatatypeChange_ToProto(mapCtx *direct.MapContext, in *krm.MultiColumnDatatypeChange) *pb.MultiColumnDatatypeChange {
	if in == nil {
		return nil
	}
	out := &pb.MultiColumnDatatypeChange{}
	out.SourceDataTypeFilter = direct.ValueOf(in.SourceDataTypeFilter)
	if oneof := SourceTextFilter_ToProto(mapCtx, in.SourceTextFilter); oneof != nil {
		out.SourceFilter = &pb.MultiColumnDatatypeChange_SourceTextFilter{SourceTextFilter: oneof}
	}
	if oneof := SourceNumericFilter_ToProto(mapCtx, in.SourceNumericFilter); oneof != nil {
		out.SourceFilter = &pb.MultiColumnDatatypeChange_SourceNumericFilter{SourceNumericFilter: oneof}
	}
	out.NewDataType = direct.ValueOf(in.NewDataType)
	out.OverrideLength = direct.ValueOf(in.OverrideLength)
	out.OverrideScale = direct.ValueOf(in.OverrideScale)
	out.OverridePrecision = direct.ValueOf(in.OverridePrecision)
	out.OverrideFractionalSecondsPrecision = direct.ValueOf(in.OverrideFractionalSecondsPrecision)
	out.CustomFeatures = CustomFeatures_ToProto(mapCtx, in.CustomFeatures)
	return out
}
func MultiEntityRename_FromProto(mapCtx *direct.MapContext, in *pb.MultiEntityRename) *krm.MultiEntityRename {
	if in == nil {
		return nil
	}
	out := &krm.MultiEntityRename{}
	out.NewNamePattern = direct.LazyPtr(in.GetNewNamePattern())
	out.SourceNameTransformation = direct.Enum_FromProto(mapCtx, in.GetSourceNameTransformation())
	return out
}
func MultiEntityRename_ToProto(mapCtx *direct.MapContext, in *krm.MultiEntityRename) *pb.MultiEntityRename {
	if in == nil {
		return nil
	}
	out := &pb.MultiEntityRename{}
	out.NewNamePattern = direct.ValueOf(in.NewNamePattern)
	out.SourceNameTransformation = direct.Enum_ToProto[pb.EntityNameTransformation](mapCtx, in.SourceNameTransformation)
	return out
}
func RoundToScale_FromProto(mapCtx *direct.MapContext, in *pb.RoundToScale) *krm.RoundToScale {
	if in == nil {
		return nil
	}
	out := &krm.RoundToScale{}
	out.Scale = direct.LazyPtr(in.GetScale())
	return out
}
func RoundToScale_ToProto(mapCtx *direct.MapContext, in *krm.RoundToScale) *pb.RoundToScale {
	if in == nil {
		return nil
	}
	out := &pb.RoundToScale{}
	out.Scale = direct.ValueOf(in.Scale)
	return out
}
func SetTablePrimaryKey_FromProto(mapCtx *direct.MapContext, in *pb.SetTablePrimaryKey) *krm.SetTablePrimaryKey {
	if in == nil {
		return nil
	}
	out := &krm.SetTablePrimaryKey{}
	out.PrimaryKeyColumns = in.PrimaryKeyColumns
	out.PrimaryKey = direct.LazyPtr(in.GetPrimaryKey())
	return out
}
func SetTablePrimaryKey_ToProto(mapCtx *direct.MapContext, in *krm.SetTablePrimaryKey) *pb.SetTablePrimaryKey {
	if in == nil {
		return nil
	}
	out := &pb.SetTablePrimaryKey{}
	out.PrimaryKeyColumns = in.PrimaryKeyColumns
	out.PrimaryKey = direct.ValueOf(in.PrimaryKey)
	return out
}
func SingleColumnChange_FromProto(mapCtx *direct.MapContext, in *pb.SingleColumnChange) *krm.SingleColumnChange {
	if in == nil {
		return nil
	}
	out := &krm.SingleColumnChange{}
	out.DataType = direct.LazyPtr(in.GetDataType())
	out.Charset = direct.LazyPtr(in.GetCharset())
	out.Collation = direct.LazyPtr(in.GetCollation())
	out.Length = direct.LazyPtr(in.GetLength())
	out.Precision = direct.LazyPtr(in.GetPrecision())
	out.Scale = direct.LazyPtr(in.GetScale())
	out.FractionalSecondsPrecision = direct.LazyPtr(in.GetFractionalSecondsPrecision())
	out.Array = direct.LazyPtr(in.GetArray())
	out.ArrayLength = direct.LazyPtr(in.GetArrayLength())
	out.Nullable = direct.LazyPtr(in.GetNullable())
	out.AutoGenerated = direct.LazyPtr(in.GetAutoGenerated())
	out.Udt = direct.LazyPtr(in.GetUdt())
	out.CustomFeatures = CustomFeatures_FromProto(mapCtx, in.GetCustomFeatures())
	out.SetValues = in.SetValues
	out.Comment = direct.LazyPtr(in.GetComment())
	return out
}
func SingleColumnChange_ToProto(mapCtx *direct.MapContext, in *krm.SingleColumnChange) *pb.SingleColumnChange {
	if in == nil {
		return nil
	}
	out := &pb.SingleColumnChange{}
	out.DataType = direct.ValueOf(in.DataType)
	out.Charset = direct.ValueOf(in.Charset)
	out.Collation = direct.ValueOf(in.Collation)
	out.Length = direct.ValueOf(in.Length)
	out.Precision = direct.ValueOf(in.Precision)
	out.Scale = direct.ValueOf(in.Scale)
	out.FractionalSecondsPrecision = direct.ValueOf(in.FractionalSecondsPrecision)
	out.Array = direct.ValueOf(in.Array)
	out.ArrayLength = direct.ValueOf(in.ArrayLength)
	out.Nullable = direct.ValueOf(in.Nullable)
	out.AutoGenerated = direct.ValueOf(in.AutoGenerated)
	out.Udt = direct.ValueOf(in.Udt)
	out.CustomFeatures = CustomFeatures_ToProto(mapCtx, in.CustomFeatures)
	out.SetValues = in.SetValues
	out.Comment = direct.ValueOf(in.Comment)
	return out
}
func SingleEntityRename_FromProto(mapCtx *direct.MapContext, in *pb.SingleEntityRename) *krm.SingleEntityRename {
	if in == nil {
		return nil
	}
	out := &krm.SingleEntityRename{}
	out.NewName = direct.LazyPtr(in.GetNewName())
	return out
}
func SingleEntityRename_ToProto(mapCtx *direct.MapContext, in *krm.SingleEntityRename) *pb.SingleEntityRename {
	if in == nil {
		return nil
	}
	out := &pb.SingleEntityRename{}
	out.NewName = direct.ValueOf(in.NewName)
	return out
}
func SinglePackageChange_FromProto(mapCtx *direct.MapContext, in *pb.SinglePackageChange) *krm.SinglePackageChange {
	if in == nil {
		return nil
	}
	out := &krm.SinglePackageChange{}
	out.PackageDescription = direct.LazyPtr(in.GetPackageDescription())
	out.PackageBody = direct.LazyPtr(in.GetPackageBody())
	return out
}
func SinglePackageChange_ToProto(mapCtx *direct.MapContext, in *krm.SinglePackageChange) *pb.SinglePackageChange {
	if in == nil {
		return nil
	}
	out := &pb.SinglePackageChange{}
	out.PackageDescription = direct.ValueOf(in.PackageDescription)
	out.PackageBody = direct.ValueOf(in.PackageBody)
	return out
}
func SourceNumericFilter_FromProto(mapCtx *direct.MapContext, in *pb.SourceNumericFilter) *krm.SourceNumericFilter {
	if in == nil {
		return nil
	}
	out := &krm.SourceNumericFilter{}
	out.SourceMinScaleFilter = direct.LazyPtr(in.GetSourceMinScaleFilter())
	out.SourceMaxScaleFilter = direct.LazyPtr(in.GetSourceMaxScaleFilter())
	out.SourceMinPrecisionFilter = direct.LazyPtr(in.GetSourceMinPrecisionFilter())
	out.SourceMaxPrecisionFilter = direct.LazyPtr(in.GetSourceMaxPrecisionFilter())
	out.NumericFilterOption = direct.Enum_FromProto(mapCtx, in.GetNumericFilterOption())
	return out
}
func SourceNumericFilter_ToProto(mapCtx *direct.MapContext, in *krm.SourceNumericFilter) *pb.SourceNumericFilter {
	if in == nil {
		return nil
	}
	out := &pb.SourceNumericFilter{}
	out.SourceMinScaleFilter = direct.ValueOf(in.SourceMinScaleFilter)
	out.SourceMaxScaleFilter = direct.ValueOf(in.SourceMaxScaleFilter)
	out.SourceMinPrecisionFilter = direct.ValueOf(in.SourceMinPrecisionFilter)
	out.SourceMaxPrecisionFilter = direct.ValueOf(in.SourceMaxPrecisionFilter)
	out.NumericFilterOption = direct.Enum_ToProto[pb.NumericFilterOption](mapCtx, in.NumericFilterOption)
	return out
}
func SourceSqlChange_FromProto(mapCtx *direct.MapContext, in *pb.SourceSqlChange) *krm.SourceSqlChange {
	if in == nil {
		return nil
	}
	out := &krm.SourceSqlChange{}
	out.SqlCode = direct.LazyPtr(in.GetSqlCode())
	return out
}
func SourceSqlChange_ToProto(mapCtx *direct.MapContext, in *krm.SourceSqlChange) *pb.SourceSqlChange {
	if in == nil {
		return nil
	}
	out := &pb.SourceSqlChange{}
	out.SqlCode = direct.ValueOf(in.SqlCode)
	return out
}
func SourceTextFilter_FromProto(mapCtx *direct.MapContext, in *pb.SourceTextFilter) *krm.SourceTextFilter {
	if in == nil {
		return nil
	}
	out := &krm.SourceTextFilter{}
	out.SourceMinLengthFilter = direct.LazyPtr(in.GetSourceMinLengthFilter())
	out.SourceMaxLengthFilter = direct.LazyPtr(in.GetSourceMaxLengthFilter())
	return out
}
func SourceTextFilter_ToProto(mapCtx *direct.MapContext, in *krm.SourceTextFilter) *pb.SourceTextFilter {
	if in == nil {
		return nil
	}
	out := &pb.SourceTextFilter{}
	out.SourceMinLengthFilter = direct.ValueOf(in.SourceMinLengthFilter)
	out.SourceMaxLengthFilter = direct.ValueOf(in.SourceMaxLengthFilter)
	return out
}
func ValueListFilter_FromProto(mapCtx *direct.MapContext, in *pb.ValueListFilter) *krm.ValueListFilter {
	if in == nil {
		return nil
	}
	out := &krm.ValueListFilter{}
	out.ValuePresentList = direct.Enum_FromProto(mapCtx, in.GetValuePresentList())
	out.Values = in.Values
	out.IgnoreCase = direct.LazyPtr(in.GetIgnoreCase())
	return out
}
func ValueListFilter_ToProto(mapCtx *direct.MapContext, in *krm.ValueListFilter) *pb.ValueListFilter {
	if in == nil {
		return nil
	}
	out := &pb.ValueListFilter{}
	out.ValuePresentList = direct.Enum_ToProto[pb.ValuePresentInList](mapCtx, in.ValuePresentList)
	out.Values = in.Values
	out.IgnoreCase = direct.ValueOf(in.IgnoreCase)
	return out
}
func ValueTransformation_FromProto(mapCtx *direct.MapContext, in *pb.ValueTransformation) *krm.ValueTransformation {
	if in == nil {
		return nil
	}
	out := &krm.ValueTransformation{}
	out.IsNull = Empty_FromProto(mapCtx, in.GetIsNull())
	out.ValueList = ValueListFilter_FromProto(mapCtx, in.GetValueList())
	out.IntComparison = IntComparisonFilter_FromProto(mapCtx, in.GetIntComparison())
	out.DoubleComparison = DoubleComparisonFilter_FromProto(mapCtx, in.GetDoubleComparison())
	out.AssignNull = Empty_FromProto(mapCtx, in.GetAssignNull())
	out.AssignSpecificValue = AssignSpecificValue_FromProto(mapCtx, in.GetAssignSpecificValue())
	out.AssignMinValue = Empty_FromProto(mapCtx, in.GetAssignMinValue())
	out.AssignMaxValue = Empty_FromProto(mapCtx, in.GetAssignMaxValue())
	out.RoundScale = RoundToScale_FromProto(mapCtx, in.GetRoundScale())
	out.ApplyHash = ApplyHash_FromProto(mapCtx, in.GetApplyHash())
	return out
}
func ValueTransformation_ToProto(mapCtx *direct.MapContext, in *krm.ValueTransformation) *pb.ValueTransformation {
	if in == nil {
		return nil
	}
	out := &pb.ValueTransformation{}
	if oneof := Empty_ToProto(mapCtx, in.IsNull); oneof != nil {
		out.Filter = &pb.ValueTransformation_IsNull{IsNull: oneof}
	}
	if oneof := ValueListFilter_ToProto(mapCtx, in.ValueList); oneof != nil {
		out.Filter = &pb.ValueTransformation_ValueList{ValueList: oneof}
	}
	if oneof := IntComparisonFilter_ToProto(mapCtx, in.IntComparison); oneof != nil {
		out.Filter = &pb.ValueTransformation_IntComparison{IntComparison: oneof}
	}
	if oneof := DoubleComparisonFilter_ToProto(mapCtx, in.DoubleComparison); oneof != nil {
		out.Filter = &pb.ValueTransformation_DoubleComparison{DoubleComparison: oneof}
	}
	if oneof := Empty_ToProto(mapCtx, in.AssignNull); oneof != nil {
		out.Action = &pb.ValueTransformation_AssignNull{AssignNull: oneof}
	}
	if oneof := AssignSpecificValue_ToProto(mapCtx, in.AssignSpecificValue); oneof != nil {
		out.Action = &pb.ValueTransformation_AssignSpecificValue{AssignSpecificValue: oneof}
	}
	if oneof := Empty_ToProto(mapCtx, in.AssignMinValue); oneof != nil {
		out.Action = &pb.ValueTransformation_AssignMinValue{AssignMinValue: oneof}
	}
	if oneof := Empty_ToProto(mapCtx, in.AssignMaxValue); oneof != nil {
		out.Action = &pb.ValueTransformation_AssignMaxValue{AssignMaxValue: oneof}
	}
	if oneof := RoundToScale_ToProto(mapCtx, in.RoundScale); oneof != nil {
		out.Action = &pb.ValueTransformation_RoundScale{RoundScale: oneof}
	}
	if oneof := ApplyHash_ToProto(mapCtx, in.ApplyHash); oneof != nil {
		out.Action = &pb.ValueTransformation_ApplyHash{ApplyHash: oneof}
	}
	return out
}
