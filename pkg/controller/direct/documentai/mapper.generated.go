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

package documentai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ProcessorType_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorType) *krm.ProcessorType {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorType{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.AvailableLocations = direct.Slice_FromProto(mapCtx, in.AvailableLocations, ProcessorType_LocationInfo_FromProto)
	out.AllowCreation = direct.LazyPtr(in.GetAllowCreation())
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	out.SampleDocumentUris = in.SampleDocumentUris
	return out
}
func ProcessorType_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorType) *pb.ProcessorType {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorType{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.Category = direct.ValueOf(in.Category)
	out.AvailableLocations = direct.Slice_ToProto(mapCtx, in.AvailableLocations, ProcessorType_LocationInfo_ToProto)
	out.AllowCreation = direct.ValueOf(in.AllowCreation)
	out.LaunchStage = direct.Enum_ToProto[pb.LaunchStage](mapCtx, in.LaunchStage)
	out.SampleDocumentUris = in.SampleDocumentUris
	return out
}
func ProcessorType_LocationInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorType_LocationInfo) *krm.ProcessorType_LocationInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorType_LocationInfo{}
	out.LocationID = direct.LazyPtr(in.GetLocationId())
	return out
}
func ProcessorType_LocationInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorType_LocationInfo) *pb.ProcessorType_LocationInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorType_LocationInfo{}
	out.LocationId = direct.ValueOf(in.LocationID)
	return out
}
