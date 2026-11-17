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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func Model_FromProto(mapCtx *direct.MapContext, in *pb.Model) *krm.Model {
	if in == nil {
		return nil
	}
	out := &krm.Model{}
	out.Name = direct.LazyPtr(in.GetName())
	out.VersionAliases = in.VersionAliases
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.VersionDescription = direct.LazyPtr(in.GetVersionDescription())
	out.DefaultCheckpointID = direct.LazyPtr(in.GetDefaultCheckpointId())
	out.PredictSchemata = PredictSchemata_FromProto(mapCtx, in.GetPredictSchemata())
	out.MetadataSchemaURI = direct.LazyPtr(in.GetMetadataSchemaUri())
	out.Metadata = Value_FromProto(mapCtx, in.GetMetadata())
	out.PipelineJob = direct.LazyPtr(in.GetPipelineJob())
	out.ContainerSpec = ModelContainerSpecWithRequired_FromProto(mapCtx, in.GetContainerSpec())
	out.ArtifactURI = direct.LazyPtr(in.GetArtifactUri())
	out.ExplanationSpec = explanationSpecWithRequired_v1_FromProto(mapCtx, in.GetExplanationSpec())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.DataStats = Model_DataStats_FromProto(mapCtx, in.GetDataStats())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	out.BaseModelSource = Model_BaseModelSourceWithRequired_FromProto(mapCtx, in.GetBaseModelSource())
	return out
}

func Model_ToProto(mapCtx *direct.MapContext, in *krm.Model) *pb.Model {
	if in == nil {
		return nil
	}
	out := &pb.Model{}
	out.Name = direct.ValueOf(in.Name)
	out.VersionAliases = in.VersionAliases
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.VersionDescription = direct.ValueOf(in.VersionDescription)
	out.DefaultCheckpointId = direct.ValueOf(in.DefaultCheckpointID)
	out.PredictSchemata = PredictSchemata_ToProto(mapCtx, in.PredictSchemata)
	out.MetadataSchemaUri = direct.ValueOf(in.MetadataSchemaURI)
	out.Metadata = Value_ToProto(mapCtx, in.Metadata)
	out.PipelineJob = direct.ValueOf(in.PipelineJob)
	out.ContainerSpec = ModelContainerSpecWithRequired_ToProto(mapCtx, in.ContainerSpec)
	out.ArtifactUri = direct.ValueOf(in.ArtifactURI)
	out.ExplanationSpec = explanationSpecWithRequired_v1_ToProto(mapCtx, in.ExplanationSpec)
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.DataStats = Model_DataStats_ToProto(mapCtx, in.DataStats)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	out.BaseModelSource = Model_BaseModelSourceWithRequired_ToProto(mapCtx, in.BaseModelSource)
	return out
}
