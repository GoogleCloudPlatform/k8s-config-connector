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

package vision

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/vision/apiv1p3beta1/visionpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vision/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Product_FromProto(mapCtx *direct.MapContext, in *pb.Product) *krm.Product {
	if in == nil {
		return nil
	}
	out := &krm.Product{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ProductCategory = direct.LazyPtr(in.GetProductCategory())
	out.ProductLabels = direct.Slice_FromProto(mapCtx, in.ProductLabels, Product_KeyValue_FromProto)
	return out
}
func Product_ToProto(mapCtx *direct.MapContext, in *krm.Product) *pb.Product {
	if in == nil {
		return nil
	}
	out := &pb.Product{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.ProductCategory = direct.ValueOf(in.ProductCategory)
	out.ProductLabels = direct.Slice_ToProto(mapCtx, in.ProductLabels, Product_KeyValue_ToProto)
	return out
}
func Product_KeyValue_FromProto(mapCtx *direct.MapContext, in *pb.Product_KeyValue) *krm.Product_KeyValue {
	if in == nil {
		return nil
	}
	out := &krm.Product_KeyValue{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func Product_KeyValue_ToProto(mapCtx *direct.MapContext, in *krm.Product_KeyValue) *pb.Product_KeyValue {
	if in == nil {
		return nil
	}
	out := &pb.Product_KeyValue{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
