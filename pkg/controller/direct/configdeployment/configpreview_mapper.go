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

package configdeployment

import (
	pb "cloud.google.com/go/config/apiv1/configpb"
	krmcloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ConfigPreviewSpec_FromProto(mapCtx *direct.MapContext, in *pb.Preview) *krm.ConfigPreviewSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigPreviewSpec{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	if in.GetDeployment() != "" {
		out.DeploymentRef = &krm.ConfigDeploymentRef{External: in.GetDeployment()}
	}
	out.PreviewMode = direct.Enum_FromProto(mapCtx, in.GetPreviewMode())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	if in.GetArtifactsGcsBucket() != "" {
		out.ArtifactsLocation = ArtifactsLocation_FromProto(mapCtx, in.GetArtifactsGcsBucket())
	}
	if in.GetWorkerPool() != "" {
		out.WorkerPoolRef = &krmcloudbuildv1beta1.CloudBuildWorkerPoolRef{External: in.GetWorkerPool()}
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_FromProto(mapCtx, in.GetProviderConfig())
	return out
}

func ConfigPreviewSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigPreviewSpec) *pb.Preview {
	if in == nil {
		return nil
	}
	out := &pb.Preview{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Preview_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	if in.DeploymentRef != nil {
		out.Deployment = in.DeploymentRef.External
	}
	out.PreviewMode = direct.Enum_ToProto[pb.Preview_PreviewMode](mapCtx, in.PreviewMode)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	if in.ArtifactsLocation != nil {
		out.ArtifactsGcsBucket = ArtifactsLocation_ToProto(mapCtx, in.ArtifactsLocation)
	}
	if in.WorkerPoolRef != nil {
		out.WorkerPool = direct.LazyPtr(in.WorkerPoolRef.External)
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_ToProto(mapCtx, in.ProviderConfig)
	return out
}
