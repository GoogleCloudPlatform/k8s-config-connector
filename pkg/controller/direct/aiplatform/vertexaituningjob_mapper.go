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

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAITuningJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.TuningJob) *krm.VertexAITuningJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAITuningJobSpec{}
	out.BaseModel = direct.LazyPtr(in.GetBaseModel())
	out.SupervisedTuningSpec = SupervisedTuningSpec_FromProto(mapCtx, in.GetSupervisedTuningSpec())
	out.TunedModelDisplayName = direct.LazyPtr(in.GetTunedModelDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	return out
}

func VertexAITuningJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAITuningJobSpec) *pb.TuningJob {
	if in == nil {
		return nil
	}
	out := &pb.TuningJob{}
	if oneof := VertexAITuningJobSpec_BaseModel_ToProto(mapCtx, in.BaseModel); oneof != nil {
		out.SourceModel = oneof
	}
	if oneof := SupervisedTuningSpec_ToProto(mapCtx, in.SupervisedTuningSpec); oneof != nil {
		out.TuningSpec = &pb.TuningJob_SupervisedTuningSpec{SupervisedTuningSpec: oneof}
	}
	out.TunedModelDisplayName = direct.ValueOf(in.TunedModelDisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	return out
}
