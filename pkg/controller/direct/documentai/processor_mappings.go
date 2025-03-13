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
	pb "cloud.google.com/go/documentai/apiv1/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DocumentAIProcessorSpec_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.DocumentAIProcessorSpec {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorSpec{}
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// NOTYET
	// out.DefaultProcessorVersion = direct.LazyPtr(in.GetDefaultProcessorVersion())
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}
func DocumentAIProcessorSpec_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorSpec) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// NOTYET
	// out.DefaultProcessorVersion = direct.ValueOf(in.DefaultProcessorVersion)
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}
func DocumentAIProcessorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.DocumentAIProcessorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentAIProcessorObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DefaultProcessorVersion = direct.LazyPtr(in.GetDefaultProcessorVersion())
	out.ProcessorVersionAliases = direct.Slice_FromProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_FromProto)
	out.ProcessEndpoint = direct.LazyPtr(in.GetProcessEndpoint())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func DocumentAIProcessorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentAIProcessorObservedState) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Processor_State](mapCtx, in.State)
	out.DefaultProcessorVersion = direct.ValueOf(in.DefaultProcessorVersion)
	out.ProcessorVersionAliases = direct.Slice_ToProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_ToProto)
	out.ProcessEndpoint = direct.ValueOf(in.ProcessEndpoint)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
