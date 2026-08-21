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

package mockdialogflow

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
)

func (s *sipTrunksServer) GetSipTrunk(ctx context.Context, req *pb.GetSipTrunkRequest) (*pb.SipTrunk, error) {
	name, err := s.parseSipTrunkName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SipTrunk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sipTrunksServer) CreateSipTrunk(ctx context.Context, req *pb.CreateSipTrunkRequest) (*pb.SipTrunk, error) {
	reqName := req.GetSipTrunk().GetName()
	name, err := s.parseSipTrunkName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.CloneOf(req.GetSipTrunk())
	obj.Name = fqn

	// Populate simulated connection
	obj.Connections = []*pb.Connection{
		{
			ConnectionId: "connection-1",
			State:        pb.Connection_CONNECTED,
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sipTrunksServer) UpdateSipTrunk(ctx context.Context, req *pb.UpdateSipTrunkRequest) (*pb.SipTrunk, error) {
	reqName := req.GetSipTrunk().GetName()
	name, err := s.parseSipTrunkName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SipTrunk{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Apply update mask
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.ExpectedHostname = req.GetSipTrunk().GetExpectedHostname()
		obj.DisplayName = req.GetSipTrunk().GetDisplayName()
	} else {
		for _, path := range paths {
			switch path {
			case "expected_hostname", "expectedHostname":
				obj.ExpectedHostname = req.GetSipTrunk().GetExpectedHostname()
			case "display_name", "displayName":
				obj.DisplayName = req.GetSipTrunk().GetDisplayName()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid/supported", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sipTrunksServer) DeleteSipTrunk(ctx context.Context, req *pb.DeleteSipTrunkRequest) (*emptypb.Empty, error) {
	name, err := s.parseSipTrunkName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.SipTrunk{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type sipTrunkName struct {
	Project  *projects.ProjectData
	Location string
	SipTrunk string
}

func (n *sipTrunkName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/sipTrunks/%s", n.Project.ID, n.Location, n.SipTrunk)
}

func (s *MockService) parseSipTrunkName(name string) (*sipTrunkName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sipTrunks" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &sipTrunkName{
			Project:  project,
			Location: tokens[3],
			SipTrunk: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/sipTrunks/{siptrunk}", name)
}
