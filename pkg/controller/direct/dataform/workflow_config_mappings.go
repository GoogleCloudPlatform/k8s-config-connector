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

package dataform

import (
	pb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataformRepositoryWorkflowConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowConfig) *krm.DataformRepositoryWorkflowConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataformRepositoryWorkflowConfigSpec{}
	if in.GetReleaseConfig() != "" {
		out.ReleaseConfigRef = refs.DataformRepositoryReleaseConfigRef{
			External: in.GetReleaseConfig(),
		}
	}
	out.InvocationConfig = WorkflowConfigInvocationConfig_FromProto(mapCtx, in.GetInvocationConfig())
	out.CronSchedule = direct.LazyPtr(in.GetCronSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}

func DataformRepositoryWorkflowConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataformRepositoryWorkflowConfigSpec) *pb.WorkflowConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowConfig{}
	out.ReleaseConfig = in.ReleaseConfigRef.External
	out.InvocationConfig = WorkflowConfigInvocationConfig_ToProto(mapCtx, in.InvocationConfig)
	out.CronSchedule = direct.ValueOf(in.CronSchedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}

func WorkflowConfigInvocationConfig_FromProto(mapCtx *direct.MapContext, in *pb.InvocationConfig) *krm.WorkflowConfigInvocationConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowConfigInvocationConfig{}
	if in.GetIncludedTargets() != nil {
		out.IncludedTargets = make([]krm.WorkflowConfigTarget, len(in.GetIncludedTargets()))
		for i, t := range in.GetIncludedTargets() {
			out.IncludedTargets[i] = WorkflowConfigTarget_FromProto(mapCtx, t)
		}
	}
	out.IncludedTags = in.GetIncludedTags()
	out.TransitiveDependenciesIncluded = direct.LazyPtr(in.GetTransitiveDependenciesIncluded())
	out.TransitiveDependentsIncluded = direct.LazyPtr(in.GetTransitiveDependentsIncluded())
	out.FullyRefreshIncrementalTablesEnabled = direct.LazyPtr(in.GetFullyRefreshIncrementalTablesEnabled())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{
			External: in.GetServiceAccount(),
		}
	}
	return out
}

func WorkflowConfigInvocationConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowConfigInvocationConfig) *pb.InvocationConfig {
	if in == nil {
		return nil
	}
	out := &pb.InvocationConfig{}
	if in.IncludedTargets != nil {
		out.IncludedTargets = make([]*pb.Target, len(in.IncludedTargets))
		for i, t := range in.IncludedTargets {
			out.IncludedTargets[i] = WorkflowConfigTarget_ToProto(mapCtx, t)
		}
	}
	out.IncludedTags = in.IncludedTags
	out.TransitiveDependenciesIncluded = direct.ValueOf(in.TransitiveDependenciesIncluded)
	out.TransitiveDependentsIncluded = direct.ValueOf(in.TransitiveDependentsIncluded)
	out.FullyRefreshIncrementalTablesEnabled = direct.ValueOf(in.FullyRefreshIncrementalTablesEnabled)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	return out
}

func WorkflowConfigTarget_FromProto(mapCtx *direct.MapContext, in *pb.Target) krm.WorkflowConfigTarget {
	if in == nil {
		return krm.WorkflowConfigTarget{}
	}
	out := krm.WorkflowConfigTarget{}
	out.Database = direct.LazyPtr(in.GetDatabase())
	out.Schema = direct.LazyPtr(in.GetSchema())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}

func WorkflowConfigTarget_ToProto(mapCtx *direct.MapContext, in krm.WorkflowConfigTarget) *pb.Target {
	out := &pb.Target{}
	out.Database = direct.ValueOf(in.Database)
	out.Schema = direct.ValueOf(in.Schema)
	out.Name = direct.ValueOf(in.Name)
	return out
}
