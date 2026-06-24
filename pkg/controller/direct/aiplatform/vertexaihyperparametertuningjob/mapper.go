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

package vertexaihyperparametertuningjob

import (
	"strconv"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	vertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func CustomJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomJobSpec) *krm.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobSpec{}
	out.PersistentResourceID = direct.LazyPtr(in.GetPersistentResourceId())
	out.WorkerPoolSpecs = direct.Slice_FromProto(mapCtx, in.WorkerPoolSpecs, WorkerPoolSpec_FromProto)
	out.Scheduling = Scheduling_FromProto(mapCtx, in.GetScheduling())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.ReservedIPRanges = in.ReservedIpRanges
	out.PSCInterfaceConfig = PSCInterfaceConfig_FromProto(mapCtx, in.GetPscInterfaceConfig())
	out.BaseOutputDirectory = GCSDestination_FromProto(mapCtx, in.GetBaseOutputDirectory())
	out.ProtectedArtifactLocationID = direct.LazyPtr(in.GetProtectedArtifactLocationId())
	if in.GetTensorboard() != "" {
		out.TensorboardRef = &vertexaiv1alpha1.VertexAITensorboardRef{External: in.GetTensorboard()}
	}
	out.EnableWebAccess = direct.LazyPtr(in.GetEnableWebAccess())
	out.EnableDashboardAccess = direct.LazyPtr(in.GetEnableDashboardAccess())
	out.Experiment = direct.LazyPtr(in.GetExperiment())
	out.ExperimentRun = direct.LazyPtr(in.GetExperimentRun())
	if len(in.Models) > 0 {
		out.ModelRefs = make([]krm.AIPlatformModelRef, len(in.Models))
		for i, m := range in.Models {
			out.ModelRefs[i] = krm.AIPlatformModelRef{External: m}
		}
	}
	return out
}

func CustomJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobSpec) *pb.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &pb.CustomJobSpec{}
	out.PersistentResourceId = direct.ValueOf(in.PersistentResourceID)
	out.WorkerPoolSpecs = direct.Slice_ToProto(mapCtx, in.WorkerPoolSpecs, WorkerPoolSpec_ToProto)
	out.Scheduling = Scheduling_ToProto(mapCtx, in.Scheduling)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.ReservedIpRanges = in.ReservedIPRanges
	out.PscInterfaceConfig = PSCInterfaceConfig_ToProto(mapCtx, in.PSCInterfaceConfig)
	out.BaseOutputDirectory = GCSDestination_ToProto(mapCtx, in.BaseOutputDirectory)
	out.ProtectedArtifactLocationId = direct.ValueOf(in.ProtectedArtifactLocationID)
	if in.TensorboardRef != nil {
		out.Tensorboard = in.TensorboardRef.External
	}
	out.EnableWebAccess = direct.ValueOf(in.EnableWebAccess)
	out.EnableDashboardAccess = direct.ValueOf(in.EnableDashboardAccess)
	out.Experiment = direct.ValueOf(in.Experiment)
	out.ExperimentRun = direct.ValueOf(in.ExperimentRun)
	if len(in.ModelRefs) > 0 {
		out.Models = make([]string, len(in.ModelRefs))
		for i, ref := range in.ModelRefs {
			out.Models[i] = ref.External
		}
	}
	return out
}

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	out := &krm.Int32Value{}
	val := in.GetValue()
	out.Value = &val
	return out
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil {
		return nil
	}
	out := &wrapperspb.Int32Value{}
	if in.Value != nil {
		out.Value = *in.Value
	}
	return out
}

func Status_FromProto(mapCtx *direct.MapContext, in *status.Status) *common.Status {
	if in == nil {
		return nil
	}
	out := &common.Status{}
	code := in.GetCode()
	out.Code = &code
	msg := in.GetMessage()
	out.Message = &msg
	return out
}

func Status_ToProto(mapCtx *direct.MapContext, in *common.Status) *status.Status {
	if in == nil {
		return nil
	}
	out := &status.Status{}
	if in.Code != nil {
		out.Code = *in.Code
	}
	if in.Message != nil {
		out.Message = *in.Message
	}
	return out
}

func Value_ToProto(mapCtx *direct.MapContext, in *krm.Value) *structpb.Value {
	if in == nil {
		return nil
	}
	out := &structpb.Value{}
	if in.BoolValue != nil {
		out.Kind = &structpb.Value_BoolValue{
			BoolValue: direct.ValueOf(in.BoolValue),
		}
	}
	if in.NullValue != nil {
		value, err := strconv.Atoi(direct.ValueOf(in.NullValue))
		if err != nil {
			mapCtx.Errorf("error converting value %s from string to int", direct.ValueOf(in.NullValue))
		}
		out.Kind = &structpb.Value_NullValue{
			NullValue: structpb.NullValue(value),
		}
	}
	if in.NumberValue != nil {
		out.Kind = &structpb.Value_NumberValue{
			NumberValue: direct.ValueOf(in.NumberValue),
		}
	}
	if in.StringValue != nil {
		out.Kind = &structpb.Value_StringValue{
			StringValue: direct.ValueOf(in.StringValue),
		}
	}
	if in.StructValue != nil {
		out.Kind = &structpb.Value_StructValue{
			StructValue: StructValue_ToProto(mapCtx, in.StructValue),
		}
	}
	return out
}

func Value_FromProto(mapCtx *direct.MapContext, in *structpb.Value) *krm.Value {
	if in == nil {
		return nil
	}
	out := &krm.Value{}
	switch in.GetKind().(type) {
	case *structpb.Value_StringValue:
		value := in.GetStringValue()
		out.StringValue = &value
	case *structpb.Value_NumberValue:
		value := in.GetNumberValue()
		out.NumberValue = &value
	case *structpb.Value_NullValue:
		value := in.GetNullValue().String()
		out.NullValue = &value
	case *structpb.Value_BoolValue:
		value := in.GetBoolValue()
		out.BoolValue = &value
	case *structpb.Value_StructValue:
		out.StructValue = StructValue_FromProto(mapCtx, in.GetStructValue())
	}
	return out
}

func StructValue_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for key, val := range in.Fields {
		out[key] = val.GetStringValue()
	}
	return out
}

func StructValue_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{}
	if len(in) > 0 {
		out.Fields = make(map[string]*structpb.Value)
	}
	for key, val := range in {
		value := &structpb.Value_StringValue{
			StringValue: val,
		}
		out.Fields[key] = &structpb.Value{
			Kind: value,
		}
	}
	return out
}
