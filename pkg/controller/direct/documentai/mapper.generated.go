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
func Processor_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.Processor {
	if in == nil {
		return nil
	}
	out := &krm.Processor{}
	// MISSING: Name
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	out.DefaultProcessorVersion = direct.LazyPtr(in.GetDefaultProcessorVersion())
	// MISSING: ProcessorVersionAliases
	// MISSING: ProcessEndpoint
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Processor_ToProto(mapCtx *direct.MapContext, in *krm.Processor) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	// MISSING: Name
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	out.DefaultProcessorVersion = direct.ValueOf(in.DefaultProcessorVersion)
	// MISSING: ProcessorVersionAliases
	// MISSING: ProcessEndpoint
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ProcessorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Processor) *krm.ProcessorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Type
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: DefaultProcessorVersion
	out.ProcessorVersionAliases = direct.Slice_FromProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_FromProto)
	out.ProcessEndpoint = direct.LazyPtr(in.GetProcessEndpoint())
	// MISSING: CreateTime
	// MISSING: KMSKeyName
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func ProcessorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorObservedState) *pb.Processor {
	if in == nil {
		return nil
	}
	out := &pb.Processor{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Type
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.Processor_State](mapCtx, in.State)
	// MISSING: DefaultProcessorVersion
	out.ProcessorVersionAliases = direct.Slice_ToProto(mapCtx, in.ProcessorVersionAliases, ProcessorVersionAlias_ToProto)
	out.ProcessEndpoint = direct.ValueOf(in.ProcessEndpoint)
	// MISSING: CreateTime
	// MISSING: KMSKeyName
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func ProcessorVersionAlias_FromProto(mapCtx *direct.MapContext, in *pb.ProcessorVersionAlias) *krm.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &krm.ProcessorVersionAlias{}
	out.Alias = direct.LazyPtr(in.GetAlias())
	out.ProcessorVersion = direct.LazyPtr(in.GetProcessorVersion())
	return out
}
func ProcessorVersionAlias_ToProto(mapCtx *direct.MapContext, in *krm.ProcessorVersionAlias) *pb.ProcessorVersionAlias {
	if in == nil {
		return nil
	}
	out := &pb.ProcessorVersionAlias{}
	out.Alias = direct.ValueOf(in.Alias)
	out.ProcessorVersion = direct.ValueOf(in.ProcessorVersion)
	return out
}
