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

func AIPlatformNotebookExecutionJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookExecutionJob) *krm.AIPlatformNotebookExecutionJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformNotebookExecutionJobSpec{}
	out.DataformRepositorySource = NotebookExecutionJob_DataformRepositorySource_FromProto(mapCtx, in.GetDataformRepositorySource())
	out.GCSNotebookSource = NotebookExecutionJob_GCSNotebookSource_FromProto(mapCtx, in.GetGcsNotebookSource())
	out.DirectNotebookSource = NotebookExecutionJob_DirectNotebookSource_FromProto(mapCtx, in.GetDirectNotebookSource())
	if in.GetNotebookRuntimeTemplateResourceName() != "" {
		out.NotebookRuntimeTemplateRef = &krm.AIPlatformNotebookRuntimeTemplateRef{
			External: in.GetNotebookRuntimeTemplateResourceName(),
		}
	}
	out.CustomEnvironmentSpec = NotebookExecutionJob_CustomEnvironmentSpecWithRequired_FromProto(mapCtx, in.GetCustomEnvironmentSpec())
	if in.GetGcsOutputUri() != "" {
		out.GCSOutputURI = direct.LazyPtr(in.GetGcsOutputUri())
	}
	if in.GetExecutionUser() != "" {
		out.ExecutionUser = direct.LazyPtr(in.GetExecutionUser())
	}
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.WorkbenchRuntime = NotebookExecutionJob_WorkbenchRuntime_FromProto(mapCtx, in.GetWorkbenchRuntime())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ExecutionTimeout = direct.StringDuration_FromProto(mapCtx, in.GetExecutionTimeout())
	if in.GetScheduleResourceName() != "" {
		out.ScheduleRef = &krm.AIPlatformScheduleRef{
			External: in.GetScheduleResourceName(),
		}
	}
	out.KernelName = direct.LazyPtr(in.GetKernelName())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}

func AIPlatformNotebookExecutionJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformNotebookExecutionJobSpec) *pb.NotebookExecutionJob {
	if in == nil {
		return nil
	}
	out := &pb.NotebookExecutionJob{}
	if oneof := NotebookExecutionJob_DataformRepositorySource_ToProto(mapCtx, in.DataformRepositorySource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_DataformRepositorySource_{DataformRepositorySource: oneof}
	}
	if oneof := NotebookExecutionJob_GCSNotebookSource_ToProto(mapCtx, in.GCSNotebookSource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_GcsNotebookSource_{GcsNotebookSource: oneof}
	}
	if oneof := NotebookExecutionJob_DirectNotebookSource_ToProto(mapCtx, in.DirectNotebookSource); oneof != nil {
		out.NotebookSource = &pb.NotebookExecutionJob_DirectNotebookSource_{DirectNotebookSource: oneof}
	}
	if in.NotebookRuntimeTemplateRef != nil {
		out.EnvironmentSpec = &pb.NotebookExecutionJob_NotebookRuntimeTemplateResourceName{
			NotebookRuntimeTemplateResourceName: in.NotebookRuntimeTemplateRef.External,
		}
	}
	if oneof := NotebookExecutionJob_CustomEnvironmentSpecWithRequired_ToProto(mapCtx, in.CustomEnvironmentSpec); oneof != nil {
		out.EnvironmentSpec = &pb.NotebookExecutionJob_CustomEnvironmentSpec_{CustomEnvironmentSpec: oneof}
	}
	if oneof := AIPlatformNotebookExecutionJobSpec_GcsOutputUri_ToProto(mapCtx, in.GCSOutputURI); oneof != nil {
		out.ExecutionSink = oneof
	}
	if oneof := AIPlatformNotebookExecutionJobSpec_ExecutionUser_ToProto(mapCtx, in.ExecutionUser); oneof != nil {
		out.ExecutionIdentity = oneof
	}
	if in.ServiceAccountRef != nil {
		out.ExecutionIdentity = &pb.NotebookExecutionJob_ServiceAccount{
			ServiceAccount: in.ServiceAccountRef.External,
		}
	}
	if oneof := NotebookExecutionJob_WorkbenchRuntime_ToProto(mapCtx, in.WorkbenchRuntime); oneof != nil {
		out.RuntimeEnvironment = &pb.NotebookExecutionJob_WorkbenchRuntime_{WorkbenchRuntime: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ExecutionTimeout = direct.StringDuration_ToProto(mapCtx, in.ExecutionTimeout)
	if in.ScheduleRef != nil {
		out.ScheduleResourceName = in.ScheduleRef.External
	}
	out.KernelName = direct.ValueOf(in.KernelName)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	return out
}
