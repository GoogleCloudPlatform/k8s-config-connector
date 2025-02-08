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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Fulfillment_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment) *krm.Fulfillment {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.GenericWebService = Fulfillment_GenericWebService_FromProto(mapCtx, in.GetGenericWebService())
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.Features = direct.Slice_FromProto(mapCtx, in.Features, Fulfillment_Feature_FromProto)
	return out
}
func Fulfillment_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment) *pb.Fulfillment {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if oneof := Fulfillment_GenericWebService_ToProto(mapCtx, in.GenericWebService); oneof != nil {
		out.Fulfillment = &pb.Fulfillment_GenericWebService_{GenericWebService: oneof}
	}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.Features = direct.Slice_ToProto(mapCtx, in.Features, Fulfillment_Feature_ToProto)
	return out
}
func Fulfillment_Feature_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_Feature) *krm.Fulfillment_Feature {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_Feature{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Fulfillment_Feature_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_Feature) *pb.Fulfillment_Feature {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_Feature{}
	out.Type = direct.Enum_ToProto[pb.Fulfillment_Feature_Type](mapCtx, in.Type)
	return out
}
func Fulfillment_GenericWebService_FromProto(mapCtx *direct.MapContext, in *pb.Fulfillment_GenericWebService) *krm.Fulfillment_GenericWebService {
	if in == nil {
		return nil
	}
	out := &krm.Fulfillment_GenericWebService{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Username = direct.LazyPtr(in.GetUsername())
	out.Password = direct.LazyPtr(in.GetPassword())
	out.RequestHeaders = in.RequestHeaders
	out.IsCloudFunction = direct.LazyPtr(in.GetIsCloudFunction())
	return out
}
func Fulfillment_GenericWebService_ToProto(mapCtx *direct.MapContext, in *krm.Fulfillment_GenericWebService) *pb.Fulfillment_GenericWebService {
	if in == nil {
		return nil
	}
	out := &pb.Fulfillment_GenericWebService{}
	out.Uri = direct.ValueOf(in.URI)
	out.Username = direct.ValueOf(in.Username)
	out.Password = direct.ValueOf(in.Password)
	out.RequestHeaders = in.RequestHeaders
	out.IsCloudFunction = direct.ValueOf(in.IsCloudFunction)
	return out
}
