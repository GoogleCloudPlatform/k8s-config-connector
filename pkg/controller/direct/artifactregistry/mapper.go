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

package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CleanupPolicyCondition_TagState_ToProto(mapCtx *direct.MapContext, in *string) *pb.CleanupPolicyCondition_TagState {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.CleanupPolicyCondition_TagState](mapCtx, in)
	return &val
}

func ArtifactRegistryRepositoryRef_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositoryRef {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositoryRef{}
	out.External = in.GetName()
	return out
}

func ArtifactRegistryRepositoryRef_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositoryRef) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Name = in.External
	return out
}
