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

// +generated:mapper
// krm.group: datacatalog.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krmv1beta1.DataCatalogTaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DataCatalogTaxonomyObservedState{}
	// MISSING: Name
	out.PolicyTagCount = direct.LazyPtr(in.GetPolicyTagCount())
	out.TaxonomyTimestamps = SystemTimestampsv1beta1_FromProto(mapCtx, in.TaxonomyTimestamps)
	out.Service = Taxonomy_Service_FromProto(mapCtx, in.GetService())
	return out
}
func DataCatalogTaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DataCatalogTaxonomyObservedState) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	out.PolicyTagCount = direct.ValueOf(in.PolicyTagCount)
	out.TaxonomyTimestamps = SystemTimestampsv1beta1_ToProto(mapCtx, in.TaxonomyTimestamps)
	out.Service = Taxonomy_Service_ToProto(mapCtx, in.Service)
	return out
}
func DataCatalogTaxonomySpec_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krmv1beta1.DataCatalogTaxonomySpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.DataCatalogTaxonomySpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ActivatedPolicyTypes = direct.EnumSlice_FromProto(mapCtx, in.ActivatedPolicyTypes)
	return out
}
func DataCatalogTaxonomySpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.DataCatalogTaxonomySpec) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ActivatedPolicyTypes = direct.EnumSlice_ToProto[pb.Taxonomy_PolicyType](mapCtx, in.ActivatedPolicyTypes)
	return out
}
func Taxonomy_Service_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy_Service) *krmv1beta1.Taxonomy_Service {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.Taxonomy_Service{}
	out.Name = direct.Enum_FromProto(mapCtx, in.GetName())
	out.Identity = direct.LazyPtr(in.GetIdentity())
	return out
}
func Taxonomy_Service_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.Taxonomy_Service) *pb.Taxonomy_Service {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy_Service{}
	out.Name = direct.Enum_ToProto[pb.ManagingSystem](mapCtx, in.Name)
	out.Identity = direct.ValueOf(in.Identity)
	return out
}

func SystemTimestampsv1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SystemTimestamps) *krmv1beta1.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	return out
}
func SystemTimestampsv1beta1_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.SystemTimestamps) *pb.SystemTimestamps {
	if in == nil {
		return nil
	}
	out := &pb.SystemTimestamps{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	return out
}
