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
