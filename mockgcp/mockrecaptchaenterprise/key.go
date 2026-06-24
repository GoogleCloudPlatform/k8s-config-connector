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

// +tool:mockgcp-support
// proto.service: google.cloud.recaptchaenterprise.v1.RecaptchaEnterpriseService
// proto.message: google.cloud.recaptchaenterprise.v1.Key

package mockrecaptchaenterprise

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/uuid"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
)

func (s *recaptchaEnterpriseService) GetKey(ctx context.Context, req *pb.GetKeyRequest) (*pb.Key, error) {
	name, err := s.parseKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Key{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	return obj, nil
}

func (s *recaptchaEnterpriseService) CreateKey(ctx context.Context, req *pb.CreateKeyRequest) (*pb.Key, error) {
	keyID := uuid.New().String()

	reqName := fmt.Sprintf("%s/keys/%s", req.GetParent(), keyID)
	name, err := s.parseKeyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetKey()).(*pb.Key)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	return obj, nil
}

func (s *recaptchaEnterpriseService) UpdateKey(ctx context.Context, req *pb.UpdateKeyRequest) (*pb.Key, error) {
	reqName := req.GetKey().GetName()

	name, err := s.parseKeyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Key{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.DisplayName = req.GetKey().GetDisplayName()
		obj.PlatformSettings = req.GetKey().GetPlatformSettings()
		obj.Labels = req.GetKey().GetLabels()
		obj.TestingOptions = req.GetKey().GetTestingOptions()
		obj.WafSettings = req.GetKey().GetWafSettings()
	} else {
		for _, path := range paths {
			if path == "display_name" || path == "displayName" {
				obj.DisplayName = req.GetKey().GetDisplayName()
			} else if path == "web_settings" || path == "webSettings" || strings.HasPrefix(path, "web_settings.") || strings.HasPrefix(path, "webSettings.") {
				obj.PlatformSettings = req.GetKey().GetPlatformSettings()
			} else if path == "android_settings" || path == "androidSettings" || strings.HasPrefix(path, "android_settings.") || strings.HasPrefix(path, "androidSettings.") {
				obj.PlatformSettings = req.GetKey().GetPlatformSettings()
			} else if path == "ios_settings" || path == "iosSettings" || strings.HasPrefix(path, "ios_settings.") || strings.HasPrefix(path, "iosSettings.") {
				obj.PlatformSettings = req.GetKey().GetPlatformSettings()
			} else if path == "express_settings" || path == "expressSettings" || strings.HasPrefix(path, "express_settings.") || strings.HasPrefix(path, "expressSettings.") {
				obj.PlatformSettings = req.GetKey().GetPlatformSettings()
			} else if path == "labels" {
				obj.Labels = req.GetKey().GetLabels()
			} else if path == "testing_options" || path == "testingOptions" || strings.HasPrefix(path, "testing_options.") || strings.HasPrefix(path, "testingOptions.") {
				obj.TestingOptions = req.GetKey().GetTestingOptions()
			} else if path == "waf_settings" || path == "wafSettings" || strings.HasPrefix(path, "waf_settings.") || strings.HasPrefix(path, "wafSettings.") {
				obj.WafSettings = req.GetKey().GetWafSettings()
			} else {
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.ReplaceAll(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
	return obj, nil
}

func (s *recaptchaEnterpriseService) DeleteKey(ctx context.Context, req *pb.DeleteKeyRequest) (*emptypb.Empty, error) {
	name, err := s.parseKeyName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.Key{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type keyName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *keyName) String() string {
	return fmt.Sprintf("projects/%s/keys/%s", n.Project.ID, n.Name)
}

// parseKeyName parses a string into a Key name.
// The expected form is `projects/*/keys/*`.
func (s *MockService) parseKeyName(name string) (*keyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "keys" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &keyName{
			Project: project,
			Name:    tokens[3],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *recaptchaEnterpriseService) ListKeys(ctx context.Context, req *pb.ListKeysRequest) (*pb.ListKeysResponse, error) {
	response := &pb.ListKeysResponse{}
	kind := (&pb.Key{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: req.Parent}, func(msg proto.Message) error {
		keyObj, ok := msg.(*pb.Key)
		if ok {
			name, err := s.parseKeyName(keyObj.Name)
			if err == nil {
				keyCopy := proto.Clone(keyObj).(*pb.Key)
				keyCopy.Name = strings.ReplaceAll(keyCopy.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))
				response.Keys = append(response.Keys, keyCopy)
			}
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error listing Key in %q: %v", req.Parent, err)
	}
	return response, nil
}
