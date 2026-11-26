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

package designcenter

import (
	pb "cloud.google.com/go/designcenter/apiv1/designcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/designcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DeploymentMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentMetadata) *krm.DeploymentMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentMetadataObservedState{}
	out.Revision = direct.LazyPtr(in.GetRevision())
	out.ComponentOutputParameters = direct.Slice_FromProto(mapCtx, in.GetComponentOutputParameters(), ComponentOutputParametersObservedState_FromProto)
	out.Error = DeploymentErrorObservedState_FromProto(mapCtx, in.GetError())
	out.Build = direct.LazyPtr(in.GetBuild())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.RetryAttempts = direct.LazyPtr(in.GetRetryAttempts())
	out.ApplicationOutputParameters = ApplicationOutputParametersObservedState_FromProto(mapCtx, in.GetApplicationOutputParameters())
	return out
}

func DeploymentMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentMetadataObservedState) *pb.DeploymentMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentMetadata{}
	out.Revision = direct.ValueOf(in.Revision)
	out.ComponentOutputParameters = direct.Slice_ToProto(mapCtx, in.ComponentOutputParameters, ComponentOutputParametersObservedState_ToProto)
	out.Error = DeploymentErrorObservedState_ToProto(mapCtx, in.Error)
	out.Build = direct.ValueOf(in.Build)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.RetryAttempts = direct.ValueOf(in.RetryAttempts)
	out.ApplicationOutputParameters = ApplicationOutputParametersObservedState_ToProto(mapCtx, in.ApplicationOutputParameters)
	return out
}

func DeploymentGroupMetadataObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentGroupMetadata) *krm.DeploymentGroupMetadataObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentGroupMetadataObservedState{}
	out.DeploymentGroup = direct.LazyPtr(in.GetDeploymentGroup())
	if in.GetDeploymentMetadataMap() != nil {
		out.DeploymentMetadataMap = make(map[string]krm.DeploymentMetadata, len(in.GetDeploymentMetadataMap()))
		for k, v := range in.GetDeploymentMetadataMap() {
			if m := DeploymentMetadata_FromProto(mapCtx, v); m != nil {
				out.DeploymentMetadataMap[k] = *m
			}
		}
	}
	return out
}

func DeploymentGroupMetadataObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentGroupMetadataObservedState) *pb.DeploymentGroupMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentGroupMetadata{}
	out.DeploymentGroup = direct.ValueOf(in.DeploymentGroup)
	if in.DeploymentMetadataMap != nil {
		out.DeploymentMetadataMap = make(map[string]*pb.DeploymentMetadata, len(in.DeploymentMetadataMap))
		for k, v := range in.DeploymentMetadataMap {
			copyV := v
			if m := DeploymentMetadata_ToProto(mapCtx, &copyV); m != nil {
				out.DeploymentMetadataMap[k] = m
			}
		}
	}
	return out
}

func DeploymentMetadata_FromProto(mapCtx *direct.MapContext, in *pb.DeploymentMetadata) *krm.DeploymentMetadata {
	if in == nil {
		return nil
	}
	out := &krm.DeploymentMetadata{}
	out.Revision = direct.LazyPtr(in.GetRevision())
	out.ComponentOutputParameters = direct.Slice_FromProto(mapCtx, in.GetComponentOutputParameters(), ComponentOutputParametersObservedState_FromProto)
	out.Error = DeploymentErrorObservedState_FromProto(mapCtx, in.GetError())
	out.Build = direct.LazyPtr(in.GetBuild())
	out.WorkerPool = direct.LazyPtr(in.GetWorkerPool())
	out.RetryAttempts = direct.LazyPtr(in.GetRetryAttempts())
	out.ApplicationOutputParameters = ApplicationOutputParametersObservedState_FromProto(mapCtx, in.GetApplicationOutputParameters())
	return out
}

func DeploymentMetadata_ToProto(mapCtx *direct.MapContext, in *krm.DeploymentMetadata) *pb.DeploymentMetadata {
	if in == nil {
		return nil
	}
	out := &pb.DeploymentMetadata{}
	out.Revision = direct.ValueOf(in.Revision)
	out.ComponentOutputParameters = direct.Slice_ToProto(mapCtx, in.ComponentOutputParameters, ComponentOutputParametersObservedState_ToProto)
	out.Error = DeploymentErrorObservedState_ToProto(mapCtx, in.Error)
	out.Build = direct.ValueOf(in.Build)
	out.WorkerPool = direct.ValueOf(in.WorkerPool)
	out.RetryAttempts = direct.ValueOf(in.RetryAttempts)
	out.ApplicationOutputParameters = ApplicationOutputParametersObservedState_ToProto(mapCtx, in.ApplicationOutputParameters)
	return out
}
