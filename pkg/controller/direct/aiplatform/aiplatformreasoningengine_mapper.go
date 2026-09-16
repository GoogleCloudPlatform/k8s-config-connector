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
	apiextensionsv1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform/apiextensionsv1"
)

func AIPlatformReasoningEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngine) *krm.AIPlatformReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.AIPlatformReasoningEngineSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Spec = ReasoningEngineSpec_FromProto(mapCtx, in.GetSpec())
	// MISSING: Etag
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.Labels = in.GetLabels()
	return out
}

func AIPlatformReasoningEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.AIPlatformReasoningEngineSpec) *pb.ReasoningEngine {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngine{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Spec = ReasoningEngineSpec_ToProto(mapCtx, in.Spec)
	// MISSING: Etag
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.Labels = in.Labels
	return out
}

func ReasoningEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineSpec) *krm.ReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineSpec{}
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.PackageSpec = ReasoningEngineSpec_PackageSpec_FromProto(mapCtx, in.GetPackageSpec())
	out.DeploymentSpec = ReasoningEngineSpec_DeploymentSpec_FromProto(mapCtx, in.GetDeploymentSpec())
	out.ClassMethods = direct.Slice_FromProto(mapCtx, in.ClassMethods, apiextensionsv1.JSON_FromProto)
	out.AgentFramework = direct.LazyPtr(in.GetAgentFramework())
	return out
}

func ReasoningEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineSpec) *pb.ReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineSpec{}
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = direct.LazyPtr(in.ServiceAccountRef.External)
	}
	out.PackageSpec = ReasoningEngineSpec_PackageSpec_ToProto(mapCtx, in.PackageSpec)
	out.DeploymentSpec = ReasoningEngineSpec_DeploymentSpec_ToProto(mapCtx, in.DeploymentSpec)
	out.ClassMethods = direct.Slice_ToProto(mapCtx, in.ClassMethods, apiextensionsv1.JSON_ToProto)
	out.AgentFramework = direct.ValueOf(in.AgentFramework)
	return out
}

func SecretRef_FromProto(mapCtx *direct.MapContext, in *pb.SecretRef) *krm.SecretRef {
	if in == nil {
		return nil
	}
	out := &krm.SecretRef{}
	if in.GetSecret() != "" {
		out.SecretRef = &refsv1beta1.SecretManagerSecretRef{External: in.GetSecret()}
	}
	if in.GetVersion() != "" {
		out.VersionRef = &refsv1beta1.SecretManagerSecretVersionRef{External: in.GetVersion()}
	}
	return out
}

func SecretRef_ToProto(mapCtx *direct.MapContext, in *krm.SecretRef) *pb.SecretRef {
	if in == nil {
		return nil
	}
	out := &pb.SecretRef{}
	if in.SecretRef != nil {
		out.Secret = in.SecretRef.External
	}
	if in.VersionRef != nil {
		out.Version = in.VersionRef.External
	}
	return out
}
