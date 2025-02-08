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

package config

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/config/v1alpha1"
)
func ConfigResourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Resource) *krm.ConfigResourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigResourceObservedState{}
	// MISSING: Name
	// MISSING: TerraformInfo
	// MISSING: CaiAssets
	// MISSING: Intent
	// MISSING: State
	return out
}
func ConfigResourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigResourceObservedState) *pb.Resource {
	if in == nil {
		return nil
	}
	out := &pb.Resource{}
	// MISSING: Name
	// MISSING: TerraformInfo
	// MISSING: CaiAssets
	// MISSING: Intent
	// MISSING: State
	return out
}
func ConfigResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Resource) *krm.ConfigResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigResourceSpec{}
	// MISSING: Name
	// MISSING: TerraformInfo
	// MISSING: CaiAssets
	// MISSING: Intent
	// MISSING: State
	return out
}
func ConfigResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigResourceSpec) *pb.Resource {
	if in == nil {
		return nil
	}
	out := &pb.Resource{}
	// MISSING: Name
	// MISSING: TerraformInfo
	// MISSING: CaiAssets
	// MISSING: Intent
	// MISSING: State
	return out
}
func ResourceCAIInfo_FromProto(mapCtx *direct.MapContext, in *pb.ResourceCAIInfo) *krm.ResourceCAIInfo {
	if in == nil {
		return nil
	}
	out := &krm.ResourceCAIInfo{}
	out.FullResourceName = direct.LazyPtr(in.GetFullResourceName())
	return out
}
func ResourceCAIInfo_ToProto(mapCtx *direct.MapContext, in *krm.ResourceCAIInfo) *pb.ResourceCAIInfo {
	if in == nil {
		return nil
	}
	out := &pb.ResourceCAIInfo{}
	out.FullResourceName = direct.ValueOf(in.FullResourceName)
	return out
}
func ResourceTerraformInfo_FromProto(mapCtx *direct.MapContext, in *pb.ResourceTerraformInfo) *krm.ResourceTerraformInfo {
	if in == nil {
		return nil
	}
	out := &krm.ResourceTerraformInfo{}
	out.Address = direct.LazyPtr(in.GetAddress())
	out.Type = direct.LazyPtr(in.GetType())
	out.ID = direct.LazyPtr(in.GetId())
	return out
}
func ResourceTerraformInfo_ToProto(mapCtx *direct.MapContext, in *krm.ResourceTerraformInfo) *pb.ResourceTerraformInfo {
	if in == nil {
		return nil
	}
	out := &pb.ResourceTerraformInfo{}
	out.Address = direct.ValueOf(in.Address)
	out.Type = direct.ValueOf(in.Type)
	out.Id = direct.ValueOf(in.ID)
	return out
}
