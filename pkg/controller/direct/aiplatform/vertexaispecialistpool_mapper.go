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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/genproto/googleapis/type/money"
)

func Money_FromProto(mapCtx *direct.MapContext, in *money.Money) *krm.Money {
	if in == nil {
		return nil
	}
	out := &krm.Money{}
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.Units = direct.LazyPtr(in.GetUnits())
	out.Nanos = direct.LazyPtr(in.GetNanos())
	return out
}

func Money_ToProto(mapCtx *direct.MapContext, in *krm.Money) *money.Money {
	if in == nil {
		return nil
	}
	out := &money.Money{}
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.Units = direct.ValueOf(in.Units)
	out.Nanos = direct.ValueOf(in.Nanos)
	return out
}

func SpecialistPool_Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	return out
}

func SpecialistPool_Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	return out
}

// VertexAIDataLabelingJob mappings are skipped since it's managed by another controller package.
func VertexAIDataLabelingJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobObservedState {
	return nil
}

func VertexAIDataLabelingJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobObservedState) *pb.DataLabelingJob {
	return nil
}

func VertexAIDataLabelingJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataLabelingJob) *krm.VertexAIDataLabelingJobSpec {
	return nil
}

func VertexAIDataLabelingJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIDataLabelingJobSpec) *pb.DataLabelingJob {
	return nil
}
