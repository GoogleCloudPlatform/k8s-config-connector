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

package kms

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/kms/inventory/apiv1/inventorypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func KmsProtectedResourcesSummaryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResourcesSummary) *krm.KmsProtectedResourcesSummaryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KmsProtectedResourcesSummaryObservedState{}
	// MISSING: Name
	// MISSING: ResourceCount
	// MISSING: ProjectCount
	// MISSING: ResourceTypes
	// MISSING: CloudProducts
	// MISSING: Locations
	return out
}
func KmsProtectedResourcesSummaryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KmsProtectedResourcesSummaryObservedState) *pb.ProtectedResourcesSummary {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResourcesSummary{}
	// MISSING: Name
	// MISSING: ResourceCount
	// MISSING: ProjectCount
	// MISSING: ResourceTypes
	// MISSING: CloudProducts
	// MISSING: Locations
	return out
}
func KmsProtectedResourcesSummarySpec_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResourcesSummary) *krm.KmsProtectedResourcesSummarySpec {
	if in == nil {
		return nil
	}
	out := &krm.KmsProtectedResourcesSummarySpec{}
	// MISSING: Name
	// MISSING: ResourceCount
	// MISSING: ProjectCount
	// MISSING: ResourceTypes
	// MISSING: CloudProducts
	// MISSING: Locations
	return out
}
func KmsProtectedResourcesSummarySpec_ToProto(mapCtx *direct.MapContext, in *krm.KmsProtectedResourcesSummarySpec) *pb.ProtectedResourcesSummary {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResourcesSummary{}
	// MISSING: Name
	// MISSING: ResourceCount
	// MISSING: ProjectCount
	// MISSING: ResourceTypes
	// MISSING: CloudProducts
	// MISSING: Locations
	return out
}
func ProtectedResourcesSummary_FromProto(mapCtx *direct.MapContext, in *pb.ProtectedResourcesSummary) *krm.ProtectedResourcesSummary {
	if in == nil {
		return nil
	}
	out := &krm.ProtectedResourcesSummary{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ResourceCount = direct.LazyPtr(in.GetResourceCount())
	out.ProjectCount = direct.LazyPtr(in.GetProjectCount())
	out.ResourceTypes = in.ResourceTypes
	out.CloudProducts = in.CloudProducts
	out.Locations = in.Locations
	return out
}
func ProtectedResourcesSummary_ToProto(mapCtx *direct.MapContext, in *krm.ProtectedResourcesSummary) *pb.ProtectedResourcesSummary {
	if in == nil {
		return nil
	}
	out := &pb.ProtectedResourcesSummary{}
	out.Name = direct.ValueOf(in.Name)
	out.ResourceCount = direct.ValueOf(in.ResourceCount)
	out.ProjectCount = direct.ValueOf(in.ProjectCount)
	out.ResourceTypes = in.ResourceTypes
	out.CloudProducts = in.CloudProducts
	out.Locations = in.Locations
	return out
}
