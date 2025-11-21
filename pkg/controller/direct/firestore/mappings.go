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
