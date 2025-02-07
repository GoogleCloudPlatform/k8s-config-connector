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

package cloudcontrolspartner

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudcontrolspartner/apiv1beta/cloudcontrolspartnerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudcontrolspartner/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Partner_FromProto(mapCtx *direct.MapContext, in *pb.Partner) *krm.Partner {
	if in == nil {
		return nil
	}
	out := &krm.Partner{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Skus = direct.Slice_FromProto(mapCtx, in.Skus, Sku_FromProto)
	out.EkmSolutions = direct.Slice_FromProto(mapCtx, in.EkmSolutions, EkmMetadata_FromProto)
	out.OperatedCloudRegions = in.OperatedCloudRegions
	out.PartnerProjectID = direct.LazyPtr(in.GetPartnerProjectId())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Partner_ToProto(mapCtx *direct.MapContext, in *krm.Partner) *pb.Partner {
	if in == nil {
		return nil
	}
	out := &pb.Partner{}
	out.Name = direct.ValueOf(in.Name)
	out.Skus = direct.Slice_ToProto(mapCtx, in.Skus, Sku_ToProto)
	out.EkmSolutions = direct.Slice_ToProto(mapCtx, in.EkmSolutions, EkmMetadata_ToProto)
	out.OperatedCloudRegions = in.OperatedCloudRegions
	out.PartnerProjectId = direct.ValueOf(in.PartnerProjectID)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PartnerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Partner) *krm.PartnerObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PartnerObservedState{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PartnerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PartnerObservedState) *pb.Partner {
	if in == nil {
		return nil
	}
	out := &pb.Partner{}
	// MISSING: Name
	// MISSING: Skus
	// MISSING: EkmSolutions
	// MISSING: OperatedCloudRegions
	// MISSING: PartnerProjectID
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Sku_FromProto(mapCtx *direct.MapContext, in *pb.Sku) *krm.Sku {
	if in == nil {
		return nil
	}
	out := &krm.Sku{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Sku_ToProto(mapCtx *direct.MapContext, in *krm.Sku) *pb.Sku {
	if in == nil {
		return nil
	}
	out := &pb.Sku{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
