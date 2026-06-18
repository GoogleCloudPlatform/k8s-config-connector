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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexDataAttributeBindingSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataAttributeBinding) *v1alpha1.DataplexDataAttributeBindingSpec {
	if in == nil {
		return nil
	}
	out := &v1alpha1.DataplexDataAttributeBindingSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels

	if in.GetResource() != "" {
		out.ResourceRef = &v1alpha1.EntityRef{External: in.GetResource()}
	}

	if v := in.GetAttributes(); len(v) != 0 {
		for i := range v {
			out.AttributeRefs = append(out.AttributeRefs, v1alpha1.DataAttributeRef{External: v[i]})
		}
	}

	out.Paths = direct.Slice_FromProto(mapCtx, in.Paths, DataAttributeBindingPath_FromProto)
	return out
}

func DataplexDataAttributeBindingSpec_ToProto(mapCtx *direct.MapContext, in *v1alpha1.DataplexDataAttributeBindingSpec) *pb.DataAttributeBinding {
	if in == nil {
		return nil
	}
	out := &pb.DataAttributeBinding{}
	out.Description = direct.ValueOf(in.Description)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels

	if in.ResourceRef != nil {
		out.ResourceReference = &pb.DataAttributeBinding_Resource{
			Resource: in.ResourceRef.External,
		}
	}

	if v := in.AttributeRefs; len(v) != 0 {
		for i := range v {
			out.Attributes = append(out.Attributes, v[i].External)
		}
	}

	out.Paths = direct.Slice_ToProto(mapCtx, in.Paths, DataAttributeBindingPath_ToProto)
	return out
}
