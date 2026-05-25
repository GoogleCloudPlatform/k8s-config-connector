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

package customjob

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CustomJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomJobSpec) *krm.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CustomJobSpec{}

	if in.GetPersistentResourceId() != "" {
		out.PersistentResourceRef = &krm.VertexAIPersistentResourceRef{External: in.GetPersistentResourceId()}
	}

	out.CustomJobWorkerPoolSpecs = direct.Slice_FromProto(mapCtx, in.WorkerPoolSpecs, CustomJobWorkerPoolSpec_FromProto)
	out.CustomJobScheduling = CustomJobScheduling_FromProto(mapCtx, in.Scheduling)

	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	if in.GetNetwork() != "" {
		out.NetworkRef = &computev1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	out.ReservedIPRanges = in.ReservedIpRanges
	out.BaseOutputDirectory = CustomJobGcsDestination_FromProto(mapCtx, in.GetBaseOutputDirectory())
	out.ProtectedArtifactLocationID = direct.LazyPtr(in.GetProtectedArtifactLocationId())
	if in.GetTensorboard() != "" {
		out.TensorboardRef = &krm.VertexAITensorboardRef{External: in.GetTensorboard()}
	}
	out.EnableWebAccess = direct.LazyPtr(in.GetEnableWebAccess())
	out.EnableDashboardAccess = direct.LazyPtr(in.GetEnableDashboardAccess())

	return out
}

func CustomJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CustomJobSpec) *pb.CustomJobSpec {
	if in == nil {
		return nil
	}
	out := &pb.CustomJobSpec{}

	if in.PersistentResourceRef != nil {
		out.PersistentResourceId = in.PersistentResourceRef.External
	}

	out.WorkerPoolSpecs = direct.Slice_ToProto(mapCtx, in.CustomJobWorkerPoolSpecs, CustomJobWorkerPoolSpec_ToProto)
	out.Scheduling = CustomJobScheduling_ToProto(mapCtx, in.CustomJobScheduling)

	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	out.ReservedIpRanges = in.ReservedIPRanges
	out.BaseOutputDirectory = CustomJobGcsDestination_ToProto(mapCtx, in.BaseOutputDirectory)
	out.ProtectedArtifactLocationId = direct.ValueOf(in.ProtectedArtifactLocationID)
	if in.TensorboardRef != nil {
		out.Tensorboard = in.TensorboardRef.External
	}
	out.EnableWebAccess = direct.ValueOf(in.EnableWebAccess)
	out.EnableDashboardAccess = direct.ValueOf(in.EnableDashboardAccess)

	return out
}
