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

package resourcemanager

import (
	"strings"

	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ProjectSpec_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ProjectSpec {
	if in == nil {
		return nil
	}
	out := &krm.ProjectSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ResourceID = direct.LazyPtr(in.GetProjectId())

	parent := in.GetParent()
	if parent != "" {
		if strings.HasPrefix(parent, "folders/") {
			out.FolderRef = &refsv1beta1.FolderRef{
				External: parent,
			}
		} else if strings.HasPrefix(parent, "organizations/") {
			out.OrganizationRef = &krm.OrganizationRef{
				External: parent,
			}
		} else {
			mapCtx.Errorf("unknown parent format: %q", parent)
		}
	}

	return out
}

func ProjectSpec_ToProto(mapCtx *direct.MapContext, in *krm.ProjectSpec) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ProjectId = direct.ValueOf(in.ResourceID)

	if in.FolderRef != nil {
		out.Parent = in.FolderRef.External
	} else if in.OrganizationRef != nil {
		out.Parent = in.OrganizationRef.External
	}
	return out
}

func ProjectStatus_FromProto(mapCtx *direct.MapContext, in *pb.Project) *krm.ProjectStatus {
	if in == nil {
		return nil
	}
	out := &krm.ProjectStatus{}
	if in.GetName() != "" {
		tokens := strings.Split(in.GetName(), "/")
		if len(tokens) == 2 && tokens[0] == "projects" {
			out.Number = direct.LazyPtr(tokens[1])
		} else {
			mapCtx.Errorf("unexpected project name format: %q", in.GetName())
		}
	}
	return out
}

func ProjectStatus_ToProto(mapCtx *direct.MapContext, in *krm.ProjectStatus) *pb.Project {
	if in == nil {
		return nil
	}
	out := &pb.Project{}
	if in.Number != nil {
		out.Name = "projects/" + *in.Number
	}
	return out
}
