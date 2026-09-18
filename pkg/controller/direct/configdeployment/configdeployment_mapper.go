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
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ConfigDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigDeploymentSpec) *pb.Deployment {
	if in == nil {
		return nil
	}
	out := &pb.Deployment{}
	if oneof := TerraformBlueprint_ToProto(mapCtx, in.TerraformBlueprint); oneof != nil {
		out.Blueprint = &pb.Deployment_TerraformBlueprint{TerraformBlueprint: oneof}
	}
	out.Labels = in.Labels
	out.ArtifactsGcsBucket = in.ArtifactsGCSBucket
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = &in.ServiceAccountRef.External
	}
	out.ImportExistingResources = in.ImportExistingResources
	if in.WorkerPoolRef != nil {
		out.WorkerPool = &in.WorkerPoolRef.External
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.QuotaValidation = direct.Enum_ToProto[pb.QuotaValidation](mapCtx, in.QuotaValidation)
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_ToProto(mapCtx, in.ProviderConfig)
	return out
}

func ConfigDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Deployment) *krm.ConfigDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigDeploymentSpec{}
	out.TerraformBlueprint = TerraformBlueprint_FromProto(mapCtx, in.GetTerraformBlueprint())
	out.Labels = in.Labels
	out.ArtifactsGCSBucket = in.ArtifactsGcsBucket
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.ImportExistingResources = in.ImportExistingResources
	if in.GetWorkerPool() != "" {
		out.WorkerPoolRef = &cloudbuildv1beta1.CloudBuildWorkerPoolRef{External: in.GetWorkerPool()}
	}
	out.TfVersionConstraint = in.TfVersionConstraint
	out.QuotaValidation = direct.Enum_FromProto(mapCtx, in.GetQuotaValidation())
	out.Annotations = in.Annotations
	out.ProviderConfig = ProviderConfig_FromProto(mapCtx, in.GetProviderConfig())
	return out
}

func ProviderConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProviderConfig) *pb.ProviderConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProviderConfig{}
	out.SourceType = direct.PtrTo(direct.Enum_ToProto[pb.ProviderConfig_ProviderSource](mapCtx, in.SourceType))
	return out
}

func ProviderConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProviderConfig) *krm.ProviderConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProviderConfig{}
	out.SourceType = direct.Enum_FromProto(mapCtx, in.GetSourceType())
	return out
}
