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

package discoveryengine

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	out := &krm.Schema{}
	out.StructSchema = StructSchema_FromProto(mapCtx, in.GetStructSchema())
	out.JsonSchema = direct.LazyPtr(in.GetJsonSchema())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	out := &pb.Schema{}
	if oneof := StructSchema_ToProto(mapCtx, in.StructSchema); oneof != nil {
		out.Schema = &pb.Schema_StructSchema{StructSchema: oneof}
	}
	if oneof := Schema_JsonSchema_ToProto(mapCtx, in.JsonSchema); oneof != nil {
		out.Schema = oneof
	}
	out.Name = direct.ValueOf(in.Name)
	return out
}
