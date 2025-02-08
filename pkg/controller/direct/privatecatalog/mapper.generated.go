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
	pb "cloud.google.com/go/privatecatalog/apiv1beta1/privatecatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privatecatalog/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func PrivatecatalogVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.PrivatecatalogVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogVersionObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogVersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.PrivatecatalogVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogVersionSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogVersionSpec) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Version_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.Version {
	if in == nil {
		return nil
	}
	out := &krm.Version{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Version_ToProto(mapCtx *direct.MapContext, in *krm.Version) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Asset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func VersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.VersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Asset = Asset_FromProto(mapCtx, in.GetAsset())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func VersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Asset = Asset_ToProto(mapCtx, in.Asset)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
