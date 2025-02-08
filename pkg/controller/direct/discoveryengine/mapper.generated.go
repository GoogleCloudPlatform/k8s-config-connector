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
func CustomTuningModel_FromProto(mapCtx *direct.MapContext, in *pb.CustomTuningModel) *krm.CustomTuningModel {
	if in == nil {
		return nil
	}
	out := &krm.CustomTuningModel{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ModelVersion = direct.LazyPtr(in.GetModelVersion())
	out.ModelState = direct.Enum_FromProto(mapCtx, in.GetModelState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.TrainingStartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetTrainingStartTime())
	// MISSING: Metrics
	out.ErrorMessage = direct.LazyPtr(in.GetErrorMessage())
	return out
}
func CustomTuningModel_ToProto(mapCtx *direct.MapContext, in *krm.CustomTuningModel) *pb.CustomTuningModel {
	if in == nil {
		return nil
	}
	out := &pb.CustomTuningModel{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ModelVersion = direct.ValueOf(in.ModelVersion)
	out.ModelState = direct.Enum_ToProto[pb.CustomTuningModel_ModelState](mapCtx, in.ModelState)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.TrainingStartTime = direct.StringTimestamp_ToProto(mapCtx, in.TrainingStartTime)
	// MISSING: Metrics
	out.ErrorMessage = direct.ValueOf(in.ErrorMessage)
	return out
}
