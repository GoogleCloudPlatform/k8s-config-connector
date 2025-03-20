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
// proto.service: google.cloud.modelarmor.v1.ModelArmor
// proto.message: google.cloud.modelarmor.v1.Template

package mockmodelarmor

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/modelarmor/v1"
)

func (s *ModelArmorV1) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Template{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ModelArmorV1) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.Template, error) {
	reqName := req.Parent + "/templates/" + req.TemplateId
	name, err := s.parseTemplateName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Template).(*pb.Template)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	s.populateDefaultsForTemplate(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ModelArmorV1) populateDefaultsForTemplate(obj *pb.Template) {

}

func (s *ModelArmorV1) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.Template, error) {
	name, err := s.parseTemplateName(req.GetTemplate().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Template{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ModelArmorV1) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*emptypb.Empty, error) {
	name, err := s.parseTemplateName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Template{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type templateName struct {
	Project      string
	Location     string
	TemplateName string
}

func (n *templateName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/templates/%s", n.Project, n.Location, n.TemplateName)
}

// parseTemplateName parses a string into a templateName.
// The expected form is `projects/*/locations/*/templates/*`.
func (s *MockService) parseTemplateName(name string) (*templateName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "templates" {
		name := &templateName{
			Project:      tokens[1],
			Location:     tokens[3],
			TemplateName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
