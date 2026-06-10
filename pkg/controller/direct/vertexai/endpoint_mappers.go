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

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	rpcpb "google.golang.org/genproto/googleapis/rpc/status"
	typepb "google.golang.org/genproto/googleapis/type/money"
	"google.golang.org/protobuf/types/known/structpb"
)

func Value_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *structpb.Value {
	return nil
}

func Value_v1alpha1_ToProto(mapCtx *direct.MapContext, in *structpb.Value) *structpb.Value {
	return nil
}

func ListValue_v1alpha1_FromProto(mapCtx *direct.MapContext, in *structpb.ListValue) *structpb.ListValue {
	return nil
}

func ListValue_v1alpha1_ToProto(mapCtx *direct.MapContext, in *structpb.ListValue) *structpb.ListValue {
	return nil
}

func ExplanationMetadata_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata) *pb.ExplanationMetadata {
	return nil
}

func ExplanationMetadata_v1alpha1_ToProto(mapCtx *direct.MapContext, in *pb.ExplanationMetadata) *pb.ExplanationMetadata {
	return nil
}

func Presets_Query_ToProto(mapCtx *direct.MapContext, in *string) *pb.Presets_Query {
	return nil
}

func PrivateServiceConnectConfig_ProjectAllowlist_FromProto(mapCtx *direct.MapContext, in []string) []refsv1beta1.ProjectRef {
	var out []refsv1beta1.ProjectRef
	for _, v := range in {
		out = append(out, refsv1beta1.ProjectRef{External: v})
	}
	return out
}

func PrivateServiceConnectConfig_ProjectAllowlist_ToProto(mapCtx *direct.MapContext, in []refsv1beta1.ProjectRef) []string {
	var out []string
	for _, v := range in {
		out = append(out, v.External)
	}
	return out
}

func Money_v1alpha1_FromProto(mapCtx *direct.MapContext, in *typepb.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Money) *typepb.Money {
	if in == nil {
		return nil
	}
	out := &typepb.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func Status_v1alpha1_FromProto(mapCtx *direct.MapContext, in *rpcpb.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func Status_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.Status) *rpcpb.Status {
	if in == nil {
		return nil
	}
	out := &rpcpb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}
