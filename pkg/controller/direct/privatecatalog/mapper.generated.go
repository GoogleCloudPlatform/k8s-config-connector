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

package privatecatalog

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/privatecatalog/apiv1beta1/privatecatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privatecatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AssetReference_FromProto(mapCtx *direct.MapContext, in *pb.AssetReference) *krm.AssetReference {
	if in == nil {
		return nil
	}
	out := &krm.AssetReference{}
	// MISSING: ID
	// MISSING: Description
	// MISSING: Inputs
	// MISSING: ValidationStatus
	// MISSING: ValidationOperation
	// MISSING: Asset
	// MISSING: GcsPath
	// MISSING: GitSource
	// MISSING: GcsSource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func AssetReference_ToProto(mapCtx *direct.MapContext, in *krm.AssetReference) *pb.AssetReference {
	if in == nil {
		return nil
	}
	out := &pb.AssetReference{}
	// MISSING: ID
	// MISSING: Description
	// MISSING: Inputs
	// MISSING: ValidationStatus
	// MISSING: ValidationOperation
	// MISSING: Asset
	// MISSING: GcsPath
	// MISSING: GitSource
	// MISSING: GcsSource
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Version = direct.ValueOf(in.Version)
	return out
}
func AssetReferenceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AssetReference) *krm.AssetReferenceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetReferenceObservedState{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Inputs = Inputs_FromProto(mapCtx, in.GetInputs())
	out.ValidationStatus = direct.Enum_FromProto(mapCtx, in.GetValidationStatus())
	out.ValidationOperation = Operation_FromProto(mapCtx, in.GetValidationOperation())
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.GcsPath = direct.LazyPtr(in.GetGcsPath())
	out.GitSource = GitSource_FromProto(mapCtx, in.GetGitSource())
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Version
	return out
}
func AssetReferenceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetReferenceObservedState) *pb.AssetReference {
	if in == nil {
		return nil
	}
	out := &pb.AssetReference{}
	out.Id = direct.ValueOf(in.ID)
	out.Description = direct.ValueOf(in.Description)
	out.Inputs = Inputs_ToProto(mapCtx, in.Inputs)
	out.ValidationStatus = direct.Enum_ToProto[pb.AssetReference_AssetValidationState](mapCtx, in.ValidationStatus)
	out.ValidationOperation = Operation_ToProto(mapCtx, in.ValidationOperation)
	if oneof := AssetReferenceObservedState_Asset_ToProto(mapCtx, in.Asset); oneof != nil {
		out.Source = oneof
	}
	if oneof := AssetReferenceObservedState_GcsPath_ToProto(mapCtx, in.GcsPath); oneof != nil {
		out.Source = oneof
	}
	if oneof := GitSource_ToProto(mapCtx, in.GitSource); oneof != nil {
		out.Source = &pb.AssetReference_GitSource{GitSource: oneof}
	}
	out.GcsSource = GcsSource_ToProto(mapCtx, in.GcsSource)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Version
	return out
}
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	// MISSING: GcsPath
	// MISSING: Generation
	// MISSING: UpdateTime
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	// MISSING: GcsPath
	// MISSING: Generation
	// MISSING: UpdateTime
	return out
}
func GcsSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GcsSourceObservedState{}
	out.GcsPath = direct.LazyPtr(in.GetGcsPath())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func GcsSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GcsSourceObservedState) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.GcsPath = direct.ValueOf(in.GcsPath)
	out.Generation = direct.ValueOf(in.Generation)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func GitSource_FromProto(mapCtx *direct.MapContext, in *pb.GitSource) *krm.GitSource {
	if in == nil {
		return nil
	}
	out := &krm.GitSource{}
	out.Repo = direct.LazyPtr(in.GetRepo())
	out.Dir = direct.LazyPtr(in.GetDir())
	out.Commit = direct.LazyPtr(in.GetCommit())
	out.Branch = direct.LazyPtr(in.GetBranch())
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}
func GitSource_ToProto(mapCtx *direct.MapContext, in *krm.GitSource) *pb.GitSource {
	if in == nil {
		return nil
	}
	out := &pb.GitSource{}
	out.Repo = direct.ValueOf(in.Repo)
	out.Dir = direct.ValueOf(in.Dir)
	if oneof := GitSource_Commit_ToProto(mapCtx, in.Commit); oneof != nil {
		out.Ref = oneof
	}
	if oneof := GitSource_Branch_ToProto(mapCtx, in.Branch); oneof != nil {
		out.Ref = oneof
	}
	if oneof := GitSource_Tag_ToProto(mapCtx, in.Tag); oneof != nil {
		out.Ref = oneof
	}
	return out
}
func Inputs_FromProto(mapCtx *direct.MapContext, in *pb.Inputs) *krm.Inputs {
	if in == nil {
		return nil
	}
	out := &krm.Inputs{}
	// MISSING: Parameters
	return out
}
func Inputs_ToProto(mapCtx *direct.MapContext, in *krm.Inputs) *pb.Inputs {
	if in == nil {
		return nil
	}
	out := &pb.Inputs{}
	// MISSING: Parameters
	return out
}
func InputsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Inputs) *krm.InputsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InputsObservedState{}
	out.Parameters = Parameters_FromProto(mapCtx, in.GetParameters())
	return out
}
func InputsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InputsObservedState) *pb.Inputs {
	if in == nil {
		return nil
	}
	out := &pb.Inputs{}
	out.Parameters = Parameters_ToProto(mapCtx, in.Parameters)
	return out
}
func PrivatecatalogProductObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.PrivatecatalogProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogProductObservedState{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogProductObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogProductObservedState) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogProductSpec_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.PrivatecatalogProductSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogProductSpec{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogProductSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogProductSpec) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Product_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.Product {
	if in == nil {
		return nil
	}
	out := &krm.Product{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Product_ToProto(mapCtx *direct.MapContext, in *krm.Product) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	// MISSING: Name
	// MISSING: AssetType
	// MISSING: DisplayMetadata
	// MISSING: IconURI
	// MISSING: AssetReferences
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func ProductObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.ProductObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProductObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AssetType = direct.LazyPtr(in.GetAssetType())
	out.DisplayMetadata = DisplayMetadata_FromProto(mapCtx, in.GetDisplayMetadata())
	out.IconURI = direct.LazyPtr(in.GetIconUri())
	out.AssetReferences = direct.Slice_FromProto(mapCtx, in.AssetReferences, AssetReference_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ProductObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProductObservedState) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	out.Name = direct.ValueOf(in.Name)
	out.AssetType = direct.ValueOf(in.AssetType)
	out.DisplayMetadata = DisplayMetadata_ToProto(mapCtx, in.DisplayMetadata)
	out.IconUri = direct.ValueOf(in.IconURI)
	out.AssetReferences = direct.Slice_ToProto(mapCtx, in.AssetReferences, AssetReference_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
