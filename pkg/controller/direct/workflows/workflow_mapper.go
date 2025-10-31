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

package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkflowsWorkflow_StateError_FromProto(mapCtx *direct.MapContext, in *pb.Workflow_StateError) *krm.WorkflowsWorkflow_StateError {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowsWorkflow_StateError{
		Details: direct.LazyPtr(in.GetDetails()),
		Type:    direct.LazyPtr(in.GetType().String()),
	}
	return out
}

func WorkflowsWorkflow_StateError_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowsWorkflow_StateError) *pb.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &pb.Workflow_StateError{
		Details: direct.ValueOf(in.Details),
		Type:    direct.Enum_ToProto[pb.Workflow_StateError_Type](mapCtx, in.Type),
	}
	return out
}
