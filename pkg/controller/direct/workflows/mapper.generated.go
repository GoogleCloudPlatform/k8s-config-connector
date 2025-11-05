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

// +generated:mapper
// krm.group: workflows.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.workflows.v1

package workflows

import (
	pb "cloud.google.com/go/workflows/apiv1/workflowspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Workflow_StateError_FromProto(mapCtx *direct.MapContext, in *pb.Workflow_StateError) *krm.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &krm.Workflow_StateError{}
	out.Details = direct.LazyPtr(in.GetDetails())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Workflow_StateError_ToProto(mapCtx *direct.MapContext, in *krm.Workflow_StateError) *pb.Workflow_StateError {
	if in == nil {
		return nil
	}
	out := &pb.Workflow_StateError{}
	out.Details = direct.ValueOf(in.Details)
	out.Type = direct.Enum_ToProto[pb.Workflow_StateError_Type](mapCtx, in.Type)
	return out
}
