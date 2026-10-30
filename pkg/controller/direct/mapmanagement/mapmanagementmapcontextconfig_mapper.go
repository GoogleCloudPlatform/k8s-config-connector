// Copyright 2026 Google LLC
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

package mapmanagement

import (
	mapmanagementpb "cloud.google.com/go/maps/mapmanagement/apiv2/mapmanagementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/mapmanagement/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MapmanagementMapContextConfigObservedState_FromProto(mapCtx *direct.MapContext, in *mapmanagementpb.MapContextConfig) *krm.MapmanagementMapContextConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MapmanagementMapContextConfigObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func MapmanagementMapContextConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MapmanagementMapContextConfigObservedState) *mapmanagementpb.MapContextConfig {
	if in == nil {
		return nil
	}
	out := &mapmanagementpb.MapContextConfig{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}

func MapmanagementMapContextConfigSpec_FromProto(mapCtx *direct.MapContext, in *mapmanagementpb.MapContextConfig) *krm.MapmanagementMapContextConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.MapmanagementMapContextConfigSpec{}
	// MISSING: Name
	if in.GetMapConfig() != "" {
		out.MapConfigRef = &krm.MapManagementMapConfigRef{External: in.GetMapConfig()}
	}
	if in.GetStyleConfig() != "" {
		out.StyleConfigRef = &krm.MapManagementStyleConfigRef{External: in.GetStyleConfig()}
	}
	for _, v := range in.GetDataset() {
		out.DatasetRefs = append(out.DatasetRefs, krm.MapManagementDatasetRef{External: v})
	}
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.MapVariants = direct.EnumSlice_FromProto(mapCtx, in.MapVariants)
	return out
}

func MapmanagementMapContextConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.MapmanagementMapContextConfigSpec) *mapmanagementpb.MapContextConfig {
	if in == nil {
		return nil
	}
	out := &mapmanagementpb.MapContextConfig{}
	// MISSING: Name
	if in.MapConfigRef != nil {
		out.MapConfig = in.MapConfigRef.External
	}
	if in.StyleConfigRef != nil {
		out.StyleConfig = in.StyleConfigRef.External
	}
	for _, v := range in.DatasetRefs {
		out.Dataset = append(out.Dataset, v.External)
	}
	out.Alias = direct.ValueOf(in.Alias)
	out.MapVariants = direct.EnumSlice_ToProto[mapmanagementpb.MapContextConfig_MapVariant](mapCtx, in.MapVariants)
	return out
}
