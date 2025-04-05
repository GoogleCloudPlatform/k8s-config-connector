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
// krm.version: v1alpha1
// proto.service: google.cloud.datacatalog.v1

package datacatalog

import (
	pb "cloud.google.com/go/datacatalog/apiv1/datacatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datacatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataCatalogTaxonomyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.DataCatalogTaxonomyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTaxonomyObservedState{}
	// MISSING: Name
	out.PolicyTagCount = direct.LazyPtr(in.GetPolicyTagCount())
	out.TaxonomyTimestamps = SystemTimestamps_FromProto(mapCtx, in.GetTaxonomyTimestamps())
	out.Service = Taxonomy_Service_FromProto(mapCtx, in.GetService())
	return out
}
func DataCatalogTaxonomyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTaxonomyObservedState) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	out.PolicyTagCount = direct.ValueOf(in.PolicyTagCount)
	out.TaxonomyTimestamps = SystemTimestamps_ToProto(mapCtx, in.TaxonomyTimestamps)
	out.Service = Taxonomy_Service_ToProto(mapCtx, in.Service)
	return out
}
func DataCatalogTaxonomySpec_FromProto(mapCtx *direct.MapContext, in *pb.Taxonomy) *krm.DataCatalogTaxonomySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataCatalogTaxonomySpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ActivatedPolicyTypes = direct.EnumSlice_FromProto(mapCtx, in.ActivatedPolicyTypes)
	return out
}
func DataCatalogTaxonomySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataCatalogTaxonomySpec) *pb.Taxonomy {
	if in == nil {
		return nil
	}
	out := &pb.Taxonomy{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ActivatedPolicyTypes = direct.EnumSlice_ToProto[pb.Taxonomy_PolicyType](mapCtx, in.ActivatedPolicyTypes)
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
