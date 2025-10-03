// Copyright 2022 Google LLC
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

package mockcompute

import (
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type TargetVpnGatewaysV1 struct {
	*MockService
	pb.UnimplementedTargetVpnGatewaysServer
}

func (s *TargetVpnGatewaysV1) Get(ctx context.Context, req *pb.GetTargetVpnGatewayRequest) (*pb.TargetVpnGateway, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetVpnGateways/" + req.GetTargetVpnGateway()
	name, err := s.parseRegionalTargetVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetVpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *TargetVpnGatewaysV1) Insert(ctx context.Context, req *pb.InsertTargetVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetVpnGateways/" + req.GetTargetVpnGatewayResource().GetName()
	name, err := s.parseRegionalTargetVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetTargetVpnGatewayResource()).(*pb.TargetVpnGateway)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#targetVpnGateway")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *TargetVpnGatewaysV1) Delete(ctx context.Context, req *pb.DeleteTargetVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetVpnGateways/" + req.GetTargetVpnGateway()
	name, err := s.parseRegionalTargetVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.TargetVpnGateway{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *TargetVpnGatewaysV1) SetLabels(ctx context.Context, req *pb.SetLabelsTargetVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/targetVpnGateways/" + req.GetResource()
	name, err := s.parseRegionalTargetVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TargetVpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalTargetVpnGatewayName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalTargetVpnGatewayName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/targetVpnGateways/" + n.Name
}

// parseRegionalTargetVpnGatewayName parses a string into a regionalTargetVpnGatewayName.
// The expected form is `projects/*/regions/*/targetVpnGateways/*`.
func (s *MockService) parseRegionalTargetVpnGatewayName(name string) (*regionalTargetVpnGatewayName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "targetVpnGateways" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalTargetVpnGatewayName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
