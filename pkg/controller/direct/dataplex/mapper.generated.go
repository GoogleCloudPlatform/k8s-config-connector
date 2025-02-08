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

package dataplex

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Action_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.Action {
	if in == nil {
		return nil
	}
	out := &krm.Action{}
	out.Category = direct.Enum_FromProto(mapCtx, in.GetCategory())
	out.Issue = direct.LazyPtr(in.GetIssue())
	out.DetectTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDetectTime())
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	out.DataLocations = in.DataLocations
	out.InvalidDataFormat = Action_InvalidDataFormat_FromProto(mapCtx, in.GetInvalidDataFormat())
	out.IncompatibleDataSchema = Action_IncompatibleDataSchema_FromProto(mapCtx, in.GetIncompatibleDataSchema())
	out.InvalidDataPartition = Action_InvalidDataPartition_FromProto(mapCtx, in.GetInvalidDataPartition())
	out.MissingData = Action_MissingData_FromProto(mapCtx, in.GetMissingData())
	out.MissingResource = Action_MissingResource_FromProto(mapCtx, in.GetMissingResource())
	out.UnauthorizedResource = Action_UnauthorizedResource_FromProto(mapCtx, in.GetUnauthorizedResource())
	out.FailedSecurityPolicyApply = Action_FailedSecurityPolicyApply_FromProto(mapCtx, in.GetFailedSecurityPolicyApply())
	out.InvalidDataOrganization = Action_InvalidDataOrganization_FromProto(mapCtx, in.GetInvalidDataOrganization())
	return out
}
func Action_ToProto(mapCtx *direct.MapContext, in *krm.Action) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	out.Category = direct.Enum_ToProto[pb.Action_Category](mapCtx, in.Category)
	out.Issue = direct.ValueOf(in.Issue)
	out.DetectTime = direct.StringTimestamp_ToProto(mapCtx, in.DetectTime)
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	out.DataLocations = in.DataLocations
	if oneof := Action_InvalidDataFormat_ToProto(mapCtx, in.InvalidDataFormat); oneof != nil {
		out.Details = &pb.Action_InvalidDataFormat_{InvalidDataFormat: oneof}
	}
	if oneof := Action_IncompatibleDataSchema_ToProto(mapCtx, in.IncompatibleDataSchema); oneof != nil {
		out.Details = &pb.Action_IncompatibleDataSchema_{IncompatibleDataSchema: oneof}
	}
	if oneof := Action_InvalidDataPartition_ToProto(mapCtx, in.InvalidDataPartition); oneof != nil {
		out.Details = &pb.Action_InvalidDataPartition_{InvalidDataPartition: oneof}
	}
	if oneof := Action_MissingData_ToProto(mapCtx, in.MissingData); oneof != nil {
		out.Details = &pb.Action_MissingData_{MissingData: oneof}
	}
	if oneof := Action_MissingResource_ToProto(mapCtx, in.MissingResource); oneof != nil {
		out.Details = &pb.Action_MissingResource_{MissingResource: oneof}
	}
	if oneof := Action_UnauthorizedResource_ToProto(mapCtx, in.UnauthorizedResource); oneof != nil {
		out.Details = &pb.Action_UnauthorizedResource_{UnauthorizedResource: oneof}
	}
	if oneof := Action_FailedSecurityPolicyApply_ToProto(mapCtx, in.FailedSecurityPolicyApply); oneof != nil {
		out.Details = &pb.Action_FailedSecurityPolicyApply_{FailedSecurityPolicyApply: oneof}
	}
	if oneof := Action_InvalidDataOrganization_ToProto(mapCtx, in.InvalidDataOrganization); oneof != nil {
		out.Details = &pb.Action_InvalidDataOrganization_{InvalidDataOrganization: oneof}
	}
	return out
}
func ActionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.ActionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ActionObservedState{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	out.Name = direct.LazyPtr(in.GetName())
	out.Lake = direct.LazyPtr(in.GetLake())
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Asset = direct.LazyPtr(in.GetAsset())
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
func ActionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ActionObservedState) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	out.Name = direct.ValueOf(in.Name)
	out.Lake = direct.ValueOf(in.Lake)
	out.Zone = direct.ValueOf(in.Zone)
	out.Asset = direct.ValueOf(in.Asset)
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
func Action_FailedSecurityPolicyApply_FromProto(mapCtx *direct.MapContext, in *pb.Action_FailedSecurityPolicyApply) *krm.Action_FailedSecurityPolicyApply {
	if in == nil {
		return nil
	}
	out := &krm.Action_FailedSecurityPolicyApply{}
	out.Asset = direct.LazyPtr(in.GetAsset())
	return out
}
func Action_FailedSecurityPolicyApply_ToProto(mapCtx *direct.MapContext, in *krm.Action_FailedSecurityPolicyApply) *pb.Action_FailedSecurityPolicyApply {
	if in == nil {
		return nil
	}
	out := &pb.Action_FailedSecurityPolicyApply{}
	out.Asset = direct.ValueOf(in.Asset)
	return out
}
func Action_IncompatibleDataSchema_FromProto(mapCtx *direct.MapContext, in *pb.Action_IncompatibleDataSchema) *krm.Action_IncompatibleDataSchema {
	if in == nil {
		return nil
	}
	out := &krm.Action_IncompatibleDataSchema{}
	out.Table = direct.LazyPtr(in.GetTable())
	out.ExistingSchema = direct.LazyPtr(in.GetExistingSchema())
	out.NewSchema = direct.LazyPtr(in.GetNewSchema())
	out.SampledDataLocations = in.SampledDataLocations
	out.SchemaChange = direct.Enum_FromProto(mapCtx, in.GetSchemaChange())
	return out
}
func Action_IncompatibleDataSchema_ToProto(mapCtx *direct.MapContext, in *krm.Action_IncompatibleDataSchema) *pb.Action_IncompatibleDataSchema {
	if in == nil {
		return nil
	}
	out := &pb.Action_IncompatibleDataSchema{}
	out.Table = direct.ValueOf(in.Table)
	out.ExistingSchema = direct.ValueOf(in.ExistingSchema)
	out.NewSchema = direct.ValueOf(in.NewSchema)
	out.SampledDataLocations = in.SampledDataLocations
	out.SchemaChange = direct.Enum_ToProto[pb.Action_IncompatibleDataSchema_SchemaChange](mapCtx, in.SchemaChange)
	return out
}
func Action_InvalidDataFormat_FromProto(mapCtx *direct.MapContext, in *pb.Action_InvalidDataFormat) *krm.Action_InvalidDataFormat {
	if in == nil {
		return nil
	}
	out := &krm.Action_InvalidDataFormat{}
	out.SampledDataLocations = in.SampledDataLocations
	out.ExpectedFormat = direct.LazyPtr(in.GetExpectedFormat())
	out.NewFormat = direct.LazyPtr(in.GetNewFormat())
	return out
}
func Action_InvalidDataFormat_ToProto(mapCtx *direct.MapContext, in *krm.Action_InvalidDataFormat) *pb.Action_InvalidDataFormat {
	if in == nil {
		return nil
	}
	out := &pb.Action_InvalidDataFormat{}
	out.SampledDataLocations = in.SampledDataLocations
	out.ExpectedFormat = direct.ValueOf(in.ExpectedFormat)
	out.NewFormat = direct.ValueOf(in.NewFormat)
	return out
}
func Action_InvalidDataOrganization_FromProto(mapCtx *direct.MapContext, in *pb.Action_InvalidDataOrganization) *krm.Action_InvalidDataOrganization {
	if in == nil {
		return nil
	}
	out := &krm.Action_InvalidDataOrganization{}
	return out
}
func Action_InvalidDataOrganization_ToProto(mapCtx *direct.MapContext, in *krm.Action_InvalidDataOrganization) *pb.Action_InvalidDataOrganization {
	if in == nil {
		return nil
	}
	out := &pb.Action_InvalidDataOrganization{}
	return out
}
func Action_InvalidDataPartition_FromProto(mapCtx *direct.MapContext, in *pb.Action_InvalidDataPartition) *krm.Action_InvalidDataPartition {
	if in == nil {
		return nil
	}
	out := &krm.Action_InvalidDataPartition{}
	out.ExpectedStructure = direct.Enum_FromProto(mapCtx, in.GetExpectedStructure())
	return out
}
func Action_InvalidDataPartition_ToProto(mapCtx *direct.MapContext, in *krm.Action_InvalidDataPartition) *pb.Action_InvalidDataPartition {
	if in == nil {
		return nil
	}
	out := &pb.Action_InvalidDataPartition{}
	out.ExpectedStructure = direct.Enum_ToProto[pb.Action_InvalidDataPartition_PartitionStructure](mapCtx, in.ExpectedStructure)
	return out
}
func Action_MissingData_FromProto(mapCtx *direct.MapContext, in *pb.Action_MissingData) *krm.Action_MissingData {
	if in == nil {
		return nil
	}
	out := &krm.Action_MissingData{}
	return out
}
func Action_MissingData_ToProto(mapCtx *direct.MapContext, in *krm.Action_MissingData) *pb.Action_MissingData {
	if in == nil {
		return nil
	}
	out := &pb.Action_MissingData{}
	return out
}
func Action_MissingResource_FromProto(mapCtx *direct.MapContext, in *pb.Action_MissingResource) *krm.Action_MissingResource {
	if in == nil {
		return nil
	}
	out := &krm.Action_MissingResource{}
	return out
}
func Action_MissingResource_ToProto(mapCtx *direct.MapContext, in *krm.Action_MissingResource) *pb.Action_MissingResource {
	if in == nil {
		return nil
	}
	out := &pb.Action_MissingResource{}
	return out
}
func Action_UnauthorizedResource_FromProto(mapCtx *direct.MapContext, in *pb.Action_UnauthorizedResource) *krm.Action_UnauthorizedResource {
	if in == nil {
		return nil
	}
	out := &krm.Action_UnauthorizedResource{}
	return out
}
func Action_UnauthorizedResource_ToProto(mapCtx *direct.MapContext, in *krm.Action_UnauthorizedResource) *pb.Action_UnauthorizedResource {
	if in == nil {
		return nil
	}
	out := &pb.Action_UnauthorizedResource{}
	return out
}
func DataplexActionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.DataplexActionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexActionObservedState{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
func DataplexActionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexActionObservedState) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
func DataplexActionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Action) *krm.DataplexActionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexActionSpec{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
func DataplexActionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexActionSpec) *pb.Action {
	if in == nil {
		return nil
	}
	out := &pb.Action{}
	// MISSING: Category
	// MISSING: Issue
	// MISSING: DetectTime
	// MISSING: Name
	// MISSING: Lake
	// MISSING: Zone
	// MISSING: Asset
	// MISSING: DataLocations
	// MISSING: InvalidDataFormat
	// MISSING: IncompatibleDataSchema
	// MISSING: InvalidDataPartition
	// MISSING: MissingData
	// MISSING: MissingResource
	// MISSING: UnauthorizedResource
	// MISSING: FailedSecurityPolicyApply
	// MISSING: InvalidDataOrganization
	return out
}
