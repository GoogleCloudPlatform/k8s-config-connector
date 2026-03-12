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

package bigtable

import (
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	krmbigtablev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Status_v1alpha1_FromProto(mapCtx *direct.MapContext, in *status.Status) *krmbigtablev1alpha1.Status {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Details = direct.Slice_FromProto(mapCtx, in.GetDetails(), Any_v1alpha1_FromProto)
	return out
}

func Status_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	out.Details = direct.Slice_ToProto(mapCtx, in.Details, Any_v1alpha1_ToProto)
	return out
}

func Any_v1alpha1_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krmbigtablev1alpha1.Any {
	if in == nil {
		return nil
	}
	out := &krmbigtablev1alpha1.Any{}
	out.TypeURL = direct.LazyPtr(in.GetTypeUrl())
	out.Value = in.GetValue()
	return out
}

func Any_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigtablev1alpha1.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	out := &anypb.Any{}
	out.TypeUrl = direct.ValueOf(in.TypeURL)
	out.Value = in.Value
	return out
}

func Type_Struct_Encoding_DelimitedBytes_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.Type_Struct_Encoding_DelimitedBytes) *krm.Type_Struct_Encoding_DelimitedBytes {
	if in == nil {
		return nil
	}
	out := &krm.Type_Struct_Encoding_DelimitedBytes{}
	out.Delimiter = in.GetDelimiter()
	return out
}
func Type_Struct_Encoding_DelimitedBytes_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Type_Struct_Encoding_DelimitedBytes) *pb.Type_Struct_Encoding_DelimitedBytes {
	if in == nil {
		return nil
	}
	out := &pb.Type_Struct_Encoding_DelimitedBytes{}
	out.Delimiter = in.Delimiter
	return out
}
