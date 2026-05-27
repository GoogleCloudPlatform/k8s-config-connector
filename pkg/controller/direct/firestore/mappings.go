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

package firestore

import (
	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IndexFields_Order_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_Order_ {
	if in == nil {
		return nil
	}

	v := direct.Enum_ToProto[pb.Index_IndexField_Order](mapCtx, in)
	out := &pb.Index_IndexField_Order_{Order: v}
	return out
}

func IndexFields_ArrayConfig_ToProto(mapCtx *direct.MapContext, in *string) *pb.Index_IndexField_ArrayConfig_ {
	if in == nil {
		return nil
	}

	v := direct.Enum_ToProto[pb.Index_IndexField_ArrayConfig](mapCtx, in)
	out := &pb.Index_IndexField_ArrayConfig_{ArrayConfig: v}
	return out
}

func Field_TTLConfig_Spec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Field_TTLConfig_Spec) *pb.Field_TtlConfig {
	if in == nil {
		return nil
	}

	enabled := direct.ValueOf(in.Enabled)
	if enabled {
		return &pb.Field_TtlConfig{}
	} else {
		// This is an unusual API: the absence of the TTLConfig indicates that TTL is disabled.
		return nil
	}
}

func Field_TTLConfig_Spec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Field_TtlConfig) *krm.Field_TTLConfig_Spec {
	if in == nil {
		return nil
	}

	// The presence of the TTLConfig indicates that TTL is enabled.
	return &krm.Field_TTLConfig_Spec{
		Enabled: direct.PtrTo(true),
	}
}

func Index_IndexField_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Index_IndexField) *krm.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &krm.Index_IndexField{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Order = direct.Enum_FromProto(mapCtx, in.GetOrder())
	out.ArrayConfig = direct.Enum_FromProto(mapCtx, in.GetArrayConfig())
	out.VectorConfig = Index_IndexField_VectorConfig_v1alpha1_FromProto(mapCtx, in.GetVectorConfig())
	return out
}

func Index_IndexField_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField) *pb.Index_IndexField {
	if in == nil {
		return nil
	}
	out := &pb.Index_IndexField{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	if oneof := IndexFields_Order_ToProto(mapCtx, in.Order); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := IndexFields_ArrayConfig_ToProto(mapCtx, in.ArrayConfig); oneof != nil {
		out.ValueMode = oneof
	}
	if oneof := Index_IndexField_VectorConfig_v1alpha1_ToProto(mapCtx, in.VectorConfig); oneof != nil {
		out.ValueMode = &pb.Index_IndexField_VectorConfig_{VectorConfig: oneof}
	}
	return out
}

// Stubs to prevent controllerbuilder from generating them and causing compile errors

func Index_IndexField_SearchConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in any) *krm.Index_IndexField_SearchConfig {
	return nil
}
func Index_IndexField_SearchConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_SearchConfig) any {
	return nil
}
func Index_IndexField_SearchConfig_SearchGeoSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in any) *krm.Index_IndexField_SearchConfig_SearchGeoSpec {
	return nil
}
func Index_IndexField_SearchConfig_SearchGeoSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_SearchConfig_SearchGeoSpec) any {
	return nil
}
func Index_IndexField_SearchConfig_SearchTextIndexSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in any) *krm.Index_IndexField_SearchConfig_SearchTextIndexSpec {
	return nil
}
func Index_IndexField_SearchConfig_SearchTextIndexSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_SearchConfig_SearchTextIndexSpec) any {
	return nil
}
func Index_IndexField_SearchConfig_SearchTextSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in any) *krm.Index_IndexField_SearchConfig_SearchTextSpec {
	return nil
}
func Index_IndexField_SearchConfig_SearchTextSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Index_IndexField_SearchConfig_SearchTextSpec) any {
	return nil
}
