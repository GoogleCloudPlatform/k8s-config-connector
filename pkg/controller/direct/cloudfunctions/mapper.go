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

package cloudfunctions

import (
	pb "cloud.google.com/go/functions/apiv1/functionspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudfunctions/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudFunctionsFunctionSpec_FromProto(mapCtx *direct.MapContext, in *pb.CloudFunction) *krm.CloudFunctionsFunctionSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudFunctionsFunctionSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: SourceArchiveURL
	// (near miss): "SourceArchiveURL" vs "SourceArchiveUrl"
	out.SourceRepository = FunctionSourceRepository_FromProto(mapCtx, in.GetSourceRepository())
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// (near miss): "HTTPSTrigger" vs "HttpsTrigger"
	out.EventTrigger = FunctionEventTrigger_FromProto(mapCtx, in.GetEventTrigger())
	// MISSING: Status
	out.EntryPoint = direct.LazyPtr(in.GetEntryPoint())
	out.Runtime = in.GetRuntime()
	out.Timeout = direct.StringDuration_FromProto(mapCtx, in.GetTimeout())
	if in.GetAvailableMemoryMb() != 0 {
		val := int64(in.GetAvailableMemoryMb())
		out.AvailableMemoryMb = &val
	}
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	out.EnvironmentVariables = in.EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	if in.GetMaxInstances() != 0 {
		val := int64(in.GetMaxInstances())
		out.MaxInstances = &val
	}
	// MISSING: MinInstances
	// MISSING: VPCConnector
	// MISSING: VPCConnectorEgressSettings
	// (near miss): "VPCConnectorEgressSettings" vs "VpcConnectorEgressSettings"
	out.IngressSettings = direct.Enum_FromProto(mapCtx, in.GetIngressSettings())
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}

func CloudFunctionsFunctionSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudFunctionsFunctionSpec) *pb.CloudFunction {
	if in == nil {
		return nil
	}
	out := &pb.CloudFunction{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	// MISSING: SourceArchiveURL
	// (near miss): "SourceArchiveURL" vs "SourceArchiveUrl"
	if oneof := FunctionSourceRepository_ToProto(mapCtx, in.SourceRepository); oneof != nil {
		out.SourceCode = &pb.CloudFunction_SourceRepository{SourceRepository: oneof}
	}
	// MISSING: SourceUploadURL
	// MISSING: HTTPSTrigger
	// (near miss): "HTTPSTrigger" vs "HttpsTrigger"
	if oneof := FunctionEventTrigger_ToProto(mapCtx, in.EventTrigger); oneof != nil {
		out.Trigger = &pb.CloudFunction_EventTrigger{EventTrigger: oneof}
	}
	// MISSING: Status
	out.EntryPoint = direct.ValueOf(in.EntryPoint)
	out.Runtime = in.Runtime
	out.Timeout = direct.StringDuration_ToProto(mapCtx, in.Timeout)
	if in.AvailableMemoryMb != nil {
		out.AvailableMemoryMb = int32(*in.AvailableMemoryMb)
	}
	// MISSING: ServiceAccountEmail
	// MISSING: UpdateTime
	// MISSING: VersionID
	// MISSING: Labels
	out.EnvironmentVariables = in.EnvironmentVariables
	// MISSING: BuildEnvironmentVariables
	// MISSING: Network
	if in.MaxInstances != nil {
		out.MaxInstances = int32(*in.MaxInstances)
	}
	// MISSING: MinInstances
	// MISSING: VPCConnector
	// MISSING: VPCConnectorEgressSettings
	// (near miss): "VPCConnectorEgressSettings" vs "VpcConnectorEgressSettings"
	out.IngressSettings = direct.Enum_ToProto[pb.CloudFunction_IngressSettings](mapCtx, in.IngressSettings)
	// MISSING: KMSKeyName
	// MISSING: BuildWorkerPool
	// MISSING: BuildID
	// MISSING: BuildName
	// MISSING: SecretEnvironmentVariables
	// MISSING: SecretVolumes
	// MISSING: SourceToken
	// MISSING: DockerRepository
	// MISSING: DockerRegistry
	// MISSING: AutomaticUpdatePolicy
	// MISSING: OnDeployUpdatePolicy
	// MISSING: BuildServiceAccount
	return out
}

func FunctionEventTrigger_FromProto(mapCtx *direct.MapContext, in *pb.EventTrigger) *krm.FunctionEventTrigger {
	if in == nil {
		return nil
	}
	out := &krm.FunctionEventTrigger{}
	out.EventType = in.GetEventType()
	if in.GetResource() != "" {
		out.ResourceRef = krm.EventTriggerResourceRef{External: in.GetResource()}
	}
	out.Service = direct.LazyPtr(in.GetService())
	if in.GetFailurePolicy() != nil {
		val := true
		out.FailurePolicy = &val
	}
	return out
}

func FunctionEventTrigger_ToProto(mapCtx *direct.MapContext, in *krm.FunctionEventTrigger) *pb.EventTrigger {
	if in == nil {
		return nil
	}
	out := &pb.EventTrigger{}
	out.EventType = in.EventType
	out.Resource = in.ResourceRef.External
	out.Service = direct.ValueOf(in.Service)
	if in.FailurePolicy != nil && *in.FailurePolicy {
		out.FailurePolicy = &pb.FailurePolicy{
			Action: &pb.FailurePolicy_Retry_{
				Retry: &pb.FailurePolicy_Retry{},
			},
		}
	}
	return out
}
