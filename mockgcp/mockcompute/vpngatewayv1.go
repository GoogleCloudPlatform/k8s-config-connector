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

type VPNGatewaysV1 struct {
	*MockService
	pb.UnimplementedVpnGatewaysServer
}

func (s *VPNGatewaysV1) Get(ctx context.Context, req *pb.GetVpnGatewayRequest) (*pb.VpnGateway, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/vpnGateways/" + req.GetVpnGateway()
	name, err := s.parseRegionalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.VpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *VPNGatewaysV1) Insert(ctx context.Context, req *pb.InsertVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/vpnGateways/" + req.GetVpnGatewayResource().GetName()
	name, err := s.parseRegionalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetVpnGatewayResource()).(*pb.VpnGateway)
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#vpnGateway")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *VPNGatewaysV1) Delete(ctx context.Context, req *pb.DeleteVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/vpnGateways/" + req.GetVpnGateway()
	name, err := s.parseRegionalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.VpnGateway{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *VPNGatewaysV1) SetLabels(ctx context.Context, req *pb.SetLabelsVpnGatewayRequest) (*pb.Operation, error) {
	reqName := "projects/" + req.GetProject() + "/regions/" + req.GetRegion() + "/vpnGateways/" + req.GetResource()
	name, err := s.parseRegionalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.VpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Labels = req.GetRegionSetLabelsRequestResource().GetLabels()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

type regionalVpnGatewayName struct {
	Project *projects.ProjectData
	Region  string
	Name    string
}

func (n *regionalVpnGatewayName) String() string {
	return "projects/" + n.Project.ID + "/regions/" + n.Region + "/vpnGateways/" + n.Name
}

// parseRegionalVpnGatewayName parses a string into a regionalVpnGatewayName.
// The expected form is `projects/*/regions/*/vpnGateways/*`.
func (s *MockService) parseRegionalVpnGatewayName(name string) (*regionalVpnGatewayName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "vpnGateways" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &regionalVpnGatewayName{
			Project: project,
			Region:  tokens[3],
			Name:    tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
