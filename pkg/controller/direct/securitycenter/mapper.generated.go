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

package securitycenter

import (
	pb "cloud.google.com/go/securitycenter/apiv2/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ValuedResource_FromProto(mapCtx *direct.MapContext, in *pb.ValuedResource) *krm.ValuedResource {
	if in == nil {
		return nil
	}
	out := &krm.ValuedResource{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Resource = direct.LazyPtr(in.GetResource())
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ResourceValue = direct.Enum_FromProto(mapCtx, in.GetResourceValue())
	out.ExposedScore = direct.LazyPtr(in.GetExposedScore())
	out.ResourceValueConfigsUsed = direct.Slice_FromProto(mapCtx, in.ResourceValueConfigsUsed, ResourceValueConfigMetadata_FromProto)
	return out
}
func ValuedResource_ToProto(mapCtx *direct.MapContext, in *krm.ValuedResource) *pb.ValuedResource {
	if in == nil {
		return nil
	}
	out := &pb.ValuedResource{}
	out.Name = direct.ValueOf(in.Name)
	out.Resource = direct.ValueOf(in.Resource)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ResourceValue = direct.Enum_ToProto[pb.ValuedResource_ResourceValue](mapCtx, in.ResourceValue)
	out.ExposedScore = direct.ValueOf(in.ExposedScore)
	out.ResourceValueConfigsUsed = direct.Slice_ToProto(mapCtx, in.ResourceValueConfigsUsed, ResourceValueConfigMetadata_ToProto)
	return out
}
