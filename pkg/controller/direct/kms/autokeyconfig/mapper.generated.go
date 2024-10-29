// Copyright 2024 Google LLC
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

package kmsautokeyconfig

import (
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutokeyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutokeyConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	if in.KeyProject != "" {
		out.KeyProject = &refs.ProjectRef{
			External: in.KeyProject,
		}
	}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func AutokeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutokeyConfig) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	out.Name = direct.ValueOf(in.Name)
	if in.KeyProject != nil {
		out.KeyProject = in.KeyProject.External
	}
	out.State = direct.Enum_ToProto[pb.AutokeyConfig_State](mapCtx, in.State)
	return out
}

func KMSAutokeyConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}

func KMSAutokeyConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigSpec{}
	parent, _ := krm.ParseKMSAutokeyConfigExternal(in.Name)
	out.FolderRef = &refs.FolderRef{
		External: parent.String(),
	}
	if in.GetKeyProject() != "" {
		out.KeyProjectRef = &refs.ProjectRef{
			External: in.GetKeyProject(),
		}
	}
	return out
}

func KMSAutokeyConfig_FromFields(mapCtx *direct.MapContext, id *krm.KMSAutokeyConfigRef, keyProject *refs.Project) *pb.AutokeyConfig {
	out := &pb.AutokeyConfig{}
	parent, err := id.Parent()
	if err != nil {
		return nil
	}
	out.Name = krm.AsKMSAutokeyConfigExternal(parent)
	if keyProject != nil {
		out.KeyProject = "projects/" + keyProject.ProjectID // keyProject expects project of the form `projects/<projectId>` or `projects/<projectNumber>`
	}
	return out
}

func KMSAutokeyConfigStatusObservedState_FromProto(mapCtx *direct.MapContext, updated *kmspb.AutokeyConfig) *krm.KMSAutokeyConfigObservedState {
	if updated == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigObservedState{}
	out.State = direct.Enum_FromProto[pb.AutokeyConfig_State](mapCtx, updated.State)
	return out
}
