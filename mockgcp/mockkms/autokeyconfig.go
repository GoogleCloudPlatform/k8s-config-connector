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

// +tool:mockgcp-support
// proto.service: google.cloud.kms.v1.AutokeyAdmin
// proto.message: google.cloud.kms.v1.AutokeyConfig

package mockkms

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
)

type autokeyAdminServer struct {
	*MockService
	pb.UnimplementedAutokeyAdminServer
}

func (r *autokeyAdminServer) GetAutokeyConfig(ctx context.Context, req *pb.GetAutokeyConfigRequest) (*pb.AutokeyConfig, error) {
	name, err := r.parseAutokeyConfigName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AutokeyConfig{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			obj.Name = fqn
			obj.Etag = "mock-etag"
			obj.State = pb.AutokeyConfig_UNINITIALIZED
			r.storage.Create(ctx, fqn, obj)
			return obj, nil
		}
		return nil, err
	}

	return obj, nil
}

func (r *autokeyAdminServer) UpdateAutokeyConfig(ctx context.Context, req *pb.UpdateAutokeyConfigRequest) (*pb.AutokeyConfig, error) {
	reqName := req.GetAutokeyConfig().GetName()
	name, err := r.parseAutokeyConfigName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	isFolder := name.folder != ""

	obj := proto.CloneOf(req.GetAutokeyConfig())
	if obj == nil {
		obj = &pb.AutokeyConfig{}
	}
	obj.Name = fqn
	obj.Etag = "mock-etag"

	// Validation and clearing logic based on resolution mode and type (Folder vs Project)
	if isFolder {
		if obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_DEDICATED_KEY_PROJECT && obj.GetKeyProject() == "" {
			return nil, status.Errorf(codes.InvalidArgument, "'key_project' must be specified when 'key_project_resolution_mode' is 'DEDICATED_KEY_PROJECT'.")
		}
		if obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_RESOURCE_PROJECT && obj.GetKeyProject() != "" {
			return nil, status.Errorf(codes.InvalidArgument, "'key_project' cannot be specified for a resource when 'key_project_resolution_mode' is 'RESOURCE_PROJECT'.")
		}
		if obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_DISABLED && obj.GetKeyProject() != "" {
			return nil, status.Errorf(codes.InvalidArgument, "'key_project' cannot be specified for a resource when 'key_project_resolution_mode' is 'DISABLED'.")
		}
	} else {
		if obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_DEDICATED_KEY_PROJECT {
			return nil, status.Errorf(codes.InvalidArgument, "The 'key_project_resolution_mode' cannot be 'DEDICATED_KEY_PROJECT' for a project resource.")
		}
		if obj.GetKeyProject() != "" {
			return nil, status.Errorf(codes.Unimplemented, "Updating key-project in autokey config for a project is not supported.")
		}
	}

	// Clearing keyProject for RESOURCE_PROJECT or DISABLED modes
	if obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_RESOURCE_PROJECT ||
		obj.GetKeyProjectResolutionMode() == pb.AutokeyConfig_DISABLED {
		obj.KeyProject = ""
	}

	switch obj.GetKeyProjectResolutionMode() {
	case pb.AutokeyConfig_RESOURCE_PROJECT:
		obj.State = pb.AutokeyConfig_ACTIVE
	case pb.AutokeyConfig_DEDICATED_KEY_PROJECT:
		if len(obj.GetKeyProject()) > 0 {
			obj.State = pb.AutokeyConfig_ACTIVE
		} else {
			obj.State = pb.AutokeyConfig_UNINITIALIZED
		}
	case pb.AutokeyConfig_DISABLED:
		obj.State = pb.AutokeyConfig_UNINITIALIZED
	default:
		if len(obj.GetKeyProject()) > 0 {
			obj.State = pb.AutokeyConfig_ACTIVE
		} else {
			obj.State = pb.AutokeyConfig_UNINITIALIZED
		}
	}
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *autokeyAdminServer) ShowEffectiveAutokeyConfig(ctx context.Context, req *pb.ShowEffectiveAutokeyConfigRequest) (*pb.ShowEffectiveAutokeyConfigResponse, error) {
	project := req.Parent
	obj := &pb.ShowEffectiveAutokeyConfigResponse{}
	obj.KeyProject = project

	return obj, nil
}

type autokeyConfigName struct {
	folder  string
	project string
}

func (a *autokeyConfigName) String() string {
	if a.folder != "" {
		return "folders/" + a.folder + "/autokeyConfig"
	}
	return "projects/" + a.project + "/autokeyConfig"
}

// parseAutokeyConfigName parses a string into an AutoKeyConfig name.
// The expected form is `folders/{FOLDER_NUMBER}/autokeyConfig` or `projects/{PROJECT_NUMBER}/autokeyConfig`.
func (r *autokeyAdminServer) parseAutokeyConfigName(name string) (*autokeyConfigName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 3 && tokens[2] == "autokeyConfig" {
		if tokens[0] == "folders" {
			return &autokeyConfigName{folder: tokens[1]}, nil
		}
		if tokens[0] == "projects" {
			return &autokeyConfigName{project: tokens[1]}, nil
		}
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
