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

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAIReasoningEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngine) *krm.VertexAIReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIReasoningEngineSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Spec = VertexAIReasoningEngineSpec_ReasoningEngineSpec_FromProto(mapCtx, in.GetSpec())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.ContextSpec = ReasoningEngineContextSpec_FromProto(mapCtx, in.GetContextSpec())
	out.EncryptionSpec = EncryptionSpecV1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}

func VertexAIReasoningEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIReasoningEngineSpec) *pb.ReasoningEngine {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngine{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Spec = VertexAIReasoningEngineSpec_ReasoningEngineSpec_ToProto(mapCtx, in.Spec)
	out.Etag = direct.ValueOf(in.Etag)
	out.ContextSpec = ReasoningEngineContextSpec_ToProto(mapCtx, in.ContextSpec)
	out.EncryptionSpec = EncryptionSpecV1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	return out
}

func VertexAIReasoningEngineSpec_ReasoningEngineSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineSpec) *krm.VertexAIReasoningEngineSpec_ReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIReasoningEngineSpec_ReasoningEngineSpec{}
	out.ServiceAccount = in.ServiceAccount
	out.PackageSpec = ReasoningEngineSpec_PackageSpec_FromProto(mapCtx, in.GetPackageSpec())
	out.DeploymentSpec = ReasoningEngineSpec_DeploymentSpec_FromProto(mapCtx, in.GetDeploymentSpec())
	out.ClassMethods = direct.Slice_FromProto(mapCtx, in.GetClassMethods(), direct.Struct_FromProto)
	out.AgentFramework = direct.LazyPtr(in.GetAgentFramework())
	// source_code_spec is handled via a separate field in KRM
	// TODO: implement mapping for source_code_spec if available in proto
	return out
}

func VertexAIReasoningEngineSpec_ReasoningEngineSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIReasoningEngineSpec_ReasoningEngineSpec) *pb.ReasoningEngineSpec {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineSpec{}
	out.ServiceAccount = in.ServiceAccount
	out.PackageSpec = ReasoningEngineSpec_PackageSpec_ToProto(mapCtx, in.PackageSpec)
	out.DeploymentSpec = ReasoningEngineSpec_DeploymentSpec_ToProto(mapCtx, in.DeploymentSpec)
	out.ClassMethods = direct.Slice_ToProto(mapCtx, in.ClassMethods, direct.Struct_ToProto)
	out.AgentFramework = direct.ValueOf(in.AgentFramework)
	return out
}

func ReasoningEngineSpec_PackageSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineSpec_PackageSpec) *krm.ReasoningEngineSpec_PackageSpec {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineSpec_PackageSpec{}
	out.PickleObjectGCSURI = direct.LazyPtr(in.GetPickleObjectGcsUri())
	out.DependencyFilesGCSURI = direct.LazyPtr(in.GetDependencyFilesGcsUri())
	out.RequirementsGCSURI = direct.LazyPtr(in.GetRequirementsGcsUri())
	out.PythonVersion = direct.LazyPtr(in.GetPythonVersion())
	return out
}

func ReasoningEngineSpec_PackageSpec_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineSpec_PackageSpec) *pb.ReasoningEngineSpec_PackageSpec {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineSpec_PackageSpec{}
	out.PickleObjectGcsUri = direct.ValueOf(in.PickleObjectGCSURI)
	out.DependencyFilesGcsUri = direct.ValueOf(in.DependencyFilesGCSURI)
	out.RequirementsGcsUri = direct.ValueOf(in.RequirementsGCSURI)
	out.PythonVersion = direct.ValueOf(in.PythonVersion)
	return out
}

func ReasoningEngineSpec_DeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineSpec_DeploymentSpec) *krm.ReasoningEngineSpec_DeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineSpec_DeploymentSpec{}
	out.Env = direct.Slice_FromProto(mapCtx, in.GetEnv(), EnvVar_FromProto)
	out.SecretEnv = direct.Slice_FromProto(mapCtx, in.GetSecretEnv(), SecretEnvVar_FromProto)
	out.PSCInterfaceConfig = PSCInterfaceConfig_FromProto(mapCtx, in.GetPscInterfaceConfig())
	out.MinInstances = in.MinInstances
	out.MaxInstances = in.MaxInstances
	out.ResourceLimits = in.ResourceLimits
	out.ContainerConcurrency = in.ContainerConcurrency
	return out
}

func ReasoningEngineSpec_DeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineSpec_DeploymentSpec) *pb.ReasoningEngineSpec_DeploymentSpec {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineSpec_DeploymentSpec{}
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_ToProto)
	out.SecretEnv = direct.Slice_ToProto(mapCtx, in.SecretEnv, SecretEnvVar_ToProto)
	out.PscInterfaceConfig = PSCInterfaceConfig_ToProto(mapCtx, in.PSCInterfaceConfig)
	out.MinInstances = in.MinInstances
	out.MaxInstances = in.MaxInstances
	out.ResourceLimits = in.ResourceLimits
	out.ContainerConcurrency = in.ContainerConcurrency
	return out
}

func EnvVar_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar) *krm.EnvVar {
	if in == nil {
		return nil
	}
	out := &krm.EnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}

func EnvVar_ToProto(mapCtx *direct.MapContext, in *krm.EnvVar) *pb.EnvVar {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}

func SecretEnvVar_FromProto(mapCtx *direct.MapContext, in *pb.SecretEnvVar) *krm.SecretEnvVar {
	if in == nil {
		return nil
	}
	out := &krm.SecretEnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SecretRef = SecretRef_FromProto(mapCtx, in.GetSecretRef())
	return out
}

func SecretEnvVar_ToProto(mapCtx *direct.MapContext, in *krm.SecretEnvVar) *pb.SecretEnvVar {
	if in == nil {
		return nil
	}
	out := &pb.SecretEnvVar{}
	out.Name = direct.ValueOf(in.Name)
	out.SecretRef = SecretRef_ToProto(mapCtx, in.SecretRef)
	return out
}

func SecretRef_FromProto(mapCtx *direct.MapContext, in *pb.SecretRef) *krm.SecretRef {
	if in == nil {
		return nil
	}
	out := &krm.SecretRef{}
	out.Secret = direct.LazyPtr(in.GetSecret())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func SecretRef_ToProto(mapCtx *direct.MapContext, in *krm.SecretRef) *pb.SecretRef {
	if in == nil {
		return nil
	}
	out := &pb.SecretRef{}
	out.Secret = direct.ValueOf(in.Secret)
	out.Version = direct.ValueOf(in.Version)
	return out
}

func PSCInterfaceConfig_FromProto(mapCtx *direct.MapContext, in *pb.PscInterfaceConfig) *krm.PSCInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &krm.PSCInterfaceConfig{}
	out.NetworkAttachment = direct.LazyPtr(in.GetNetworkAttachment())
	out.DNSPeeringConfigs = direct.Slice_FromProto(mapCtx, in.GetDnsPeeringConfigs(), DNSPeeringConfig_FromProto)
	return out
}

func PSCInterfaceConfig_ToProto(mapCtx *direct.MapContext, in *krm.PSCInterfaceConfig) *pb.PscInterfaceConfig {
	if in == nil {
		return nil
	}
	out := &pb.PscInterfaceConfig{}
	out.NetworkAttachment = direct.ValueOf(in.NetworkAttachment)
	out.DnsPeeringConfigs = direct.Slice_ToProto(mapCtx, in.DNSPeeringConfigs, DNSPeeringConfig_ToProto)
	return out
}

func DNSPeeringConfig_FromProto(mapCtx *direct.MapContext, in *pb.DnsPeeringConfig) *krm.DNSPeeringConfig {
	if in == nil {
		return nil
	}
	out := &krm.DNSPeeringConfig{}
	out.Domain = direct.LazyPtr(in.GetDomain())
	out.TargetProject = direct.LazyPtr(in.GetTargetProject())
	out.TargetNetwork = direct.LazyPtr(in.GetTargetNetwork())
	return out
}

func DNSPeeringConfig_ToProto(mapCtx *direct.MapContext, in *krm.DNSPeeringConfig) *pb.DnsPeeringConfig {
	if in == nil {
		return nil
	}
	out := &pb.DnsPeeringConfig{}
	out.Domain = direct.ValueOf(in.Domain)
	out.TargetProject = direct.ValueOf(in.TargetProject)
	out.TargetNetwork = direct.ValueOf(in.TargetNetwork)
	return out
}

func ReasoningEngineContextSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec) *krm.ReasoningEngineContextSpec {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec{}
	out.MemoryBankConfig = ReasoningEngineContextSpec_MemoryBankConfig_FromProto(mapCtx, in.GetMemoryBankConfig())
	return out
}

func ReasoningEngineContextSpec_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec) *pb.ReasoningEngineContextSpec {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec{}
	out.MemoryBankConfig = ReasoningEngineContextSpec_MemoryBankConfig_ToProto(mapCtx, in.MemoryBankConfig)
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec_MemoryBankConfig) *krm.ReasoningEngineContextSpec_MemoryBankConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec_MemoryBankConfig{}
	out.GenerationConfig = ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig_FromProto(mapCtx, in.GetGenerationConfig())
	out.SimilaritySearchConfig = ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig_FromProto(mapCtx, in.GetSimilaritySearchConfig())
	out.TtlConfig = ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_FromProto(mapCtx, in.GetTtlConfig())
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec_MemoryBankConfig) *pb.ReasoningEngineContextSpec_MemoryBankConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec_MemoryBankConfig{}
	out.GenerationConfig = ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig_ToProto(mapCtx, in.GenerationConfig)
	out.SimilaritySearchConfig = ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig_ToProto(mapCtx, in.SimilaritySearchConfig)
	out.TtlConfig = ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_ToProto(mapCtx, in.TtlConfig)
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig) *krm.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig{}
	out.Model = direct.LazyPtr(in.GetModel())
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig) *pb.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec_MemoryBankConfig_GenerationConfig{}
	out.Model = direct.ValueOf(in.Model)
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig) *krm.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig{}
	out.EmbeddingModel = direct.LazyPtr(in.GetEmbeddingModel())
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig) *pb.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec_MemoryBankConfig_SimilaritySearchConfig{}
	out.EmbeddingModel = direct.ValueOf(in.EmbeddingModel)
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig) *krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig{}
	out.DefaultTTL = direct.Duration_FromProto(mapCtx, in.GetDefaultTtl())
	out.GranularTTLConfig = ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig_FromProto(mapCtx, in.GetGranularTtlConfig())
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig) *pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig{}
	out.Ttl = &pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_DefaultTtl{
		DefaultTtl: direct.Duration_ToProto(mapCtx, in.DefaultTTL),
	}
	if in.GranularTTLConfig != nil {
		out.Ttl = &pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig_{
			GranularTtlConfig: ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig_ToProto(mapCtx, in.GranularTTLConfig),
		}
	}
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig) *krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig{}
	// TODO: Add fields once available in proto
	return out
}

func ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig) *pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngineContextSpec_MemoryBankConfig_TtlConfig_GranularTtlConfig{}
	return out
}

func VertexAIReasoningEngineObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReasoningEngine) *krm.VertexAIReasoningEngineObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIReasoningEngineObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func VertexAIReasoningEngineObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIReasoningEngineObservedState) *pb.ReasoningEngine {
	if in == nil {
		return nil
	}
	out := &pb.ReasoningEngine{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
