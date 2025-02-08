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

package datacatalog

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
)
func DatacatalogTaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.DatacatalogTaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogTaxonomyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	// MISSING: ActivatedPolicyTypes
	// MISSING: Service
	return out
}
func DatacatalogTaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogTaxonomyObservedState) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	// MISSING: ActivatedPolicyTypes
	// MISSING: Service
	return out
}
func DatacatalogTaxonomySpec_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.DatacatalogTaxonomySpec {
	if in == nil {
		return nil
	}
	out := &krm.DatacatalogTaxonomySpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	// MISSING: ActivatedPolicyTypes
	// MISSING: Service
	return out
}
func DatacatalogTaxonomySpec_ToProto(mapCtx *direct.MapContext, in *krm.DatacatalogTaxonomySpec) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	// MISSING: ActivatedPolicyTypes
	// MISSING: Service
	return out
}
func SystemTimestamps_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ExpireTime
	return out
}
func SystemTimestamps_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ExpireTime
	return out
}
func SystemTimestampsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krm.SystemTimestampsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SystemTimestampsObservedState{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func SystemTimestampsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SystemTimestampsObservedState) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
func Taxonomy_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.Taxonomy {
	if in == nil {
		return nil
	}
	out := &krm.Taxonomy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	out.ActivatedPolicyTypes = direct.EnumSlice_FromProto(mapCtx, in.ActivatedPolicyTypes)
	// MISSING: Service
	return out
}
func Taxonomy_ToProto(mapCtx *direct.MapContext, in *krm.Taxonomy) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: PolicyTagCount
	// MISSING: TaxonomyTimestamps
	out.ActivatedPolicyTypes = direct.EnumSlice_ToProto[pb.Taxonomy_PolicyType](mapCtx, in.ActivatedPolicyTypes)
	// MISSING: Service
	return out
}
func TaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.TaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TaxonomyObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.PolicyTagCount = direct.LazyPtr(in.GetPolicyTagCount())
	out.TaxonomyTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetTaxonomyTimestamps())
	// MISSING: ActivatedPolicyTypes
	out.Service = Taxonomy_Service_FromProto(mapCtx, in.GetService())
	return out
}
func TaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TaxonomyObservedState) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.PolicyTagCount = direct.ValueOf(in.PolicyTagCount)
	out.TaxonomyTimestamps = SystemTimestamps_ToProto(mapCtx, in.TaxonomyTimestamps)
	// MISSING: ActivatedPolicyTypes
	out.Service = Taxonomy_Service_ToProto(mapCtx, in.Service)
	return out
}
func Taxonomy_Service_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy_Service) *krm.Taxonomy_Service {
	if in == nil {
		return nil
	}
	out := &krm.Taxonomy_Service{}
	out.Name = direct.Enum_FromProto(mapCtx, in.GetName())
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func Taxonomy_Service_ToProto(mapCtx *direct.MapContext, in *krm.Taxonomy_Service) *pb.Taxonomy_Service {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy_Service{}
	out.Name = direct.Enum_ToProto[pb.ManagingSystem](mapCtx, in.Name)
	out.Identity = direct.ValueOf(in.Identity)
	return out
}
