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

package dialogflow

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3beta1/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
)
func SessionEntityType_FromProto(mapCtx *direct.MapContext, in *pb.SessionEntityType) *krm.SessionEntityType {
	if in == nil {
		return nil
	}
	out := &krm.SessionEntityType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.EntityOverrideMode = direct.Enum_FromProto(mapCtx, in.GetEntityOverrideMode())
	out.Entities = direct.Slice_FromProto(mapCtx, in.Entities, EntityType_Entity_FromProto)
	return out
}
func SessionEntityType_ToProto(mapCtx *direct.MapContext, in *krm.SessionEntityType) *pb.SessionEntityType {
	if in == nil {
		return nil
	}
	out := &pb.SessionEntityType{}
	out.Name = direct.ValueOf(in.Name)
	out.EntityOverrideMode = direct.Enum_ToProto[pb.SessionEntityType_EntityOverrideMode](mapCtx, in.EntityOverrideMode)
	out.Entities = direct.Slice_ToProto(mapCtx, in.Entities, EntityType_Entity_ToProto)
	return out
}
