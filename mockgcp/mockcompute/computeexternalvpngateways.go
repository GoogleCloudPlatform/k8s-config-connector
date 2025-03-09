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
// proto.service: google.cloud.compute.v1.ExternalVpnGateways
// proto.message: google.cloud.compute.v1.ExternalVpnGateway

package mockcompute

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type externalVPNGateways struct {
	*MockService
	pb.UnimplementedExternalVpnGatewaysServer
}

func (s *externalVPNGateways) Get(ctx context.Context, req *pb.GetExternalVpnGatewayRequest) (*pb.ExternalVpnGateway, error) {
	reqName := fmt.Sprintf("projects/%s/global/externalVpnGateways/%s", req.GetProject(), req.GetExternalVpnGateway())
	name, err := s.parseExternalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ExternalVpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ExternalVpnGateway %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *externalVPNGateways) Insert(ctx context.Context, req *pb.InsertExternalVpnGatewayRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/externalVpnGateways/%s", req.GetProject(), req.GetExternalVpnGatewayResource().GetName())
	name, err := s.parseExternalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GetExternalVpnGatewayResource()).(*pb.ExternalVpnGateway)
	obj.Id = proto.Uint64(s.generateID())
	obj.SelfLink = PtrTo(buildComputeSelfLink(ctx, fqn))
	obj.Kind = PtrTo("compute#externalVpnGateway")
	obj.CreationTimestamp = PtrTo(s.nowString())

	obj.LabelFingerprint = PtrTo(labelsFingerprint(obj.GetLabels()))

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.externalVpnGateways.insert"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *externalVPNGateways) Delete(ctx context.Context, req *pb.DeleteExternalVpnGatewayRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/externalVpnGateways/%s", req.GetProject(), req.GetExternalVpnGateway())
	name, err := s.parseExternalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deleted := &pb.ExternalVpnGateway{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.externalVpnGateways.delete"),
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *externalVPNGateways) SetLabels(ctx context.Context, req *pb.SetLabelsExternalVpnGatewayRequest) (*pb.Operation, error) {
	reqName := fmt.Sprintf("projects/%s/global/externalVpnGateways/%s", req.GetProject(), req.GetResource())
	name, err := s.parseExternalVpnGatewayName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ExternalVpnGateway{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.GetLabelFingerprint() != req.GetGlobalSetLabelsRequestResource().GetLabelFingerprint() {
		return nil, status.Errorf(codes.FailedPrecondition, "LabelFingerprint mismatch")
	}
	obj.Labels = req.GetGlobalSetLabelsRequestResource().GetLabels()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("update"),
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		User:          PtrTo("user@example.com"),
	}
	return s.computeOperations.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

type externalVpnGatewayName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *externalVpnGatewayName) String() string {
	return "projects/" + n.Project.ID + "/global/externalVpnGateways/" + n.Name
}

// parseExternalVpnGatewayName parses a string into a externalVpnGatewayName.
// The expected form is `locations/global/firewallPolicies/*`.
func (s *MockService) parseExternalVpnGatewayName(name string) (*externalVpnGatewayName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "externalVpnGateways" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &externalVpnGatewayName{
			Project: project,
			Name:    tokens[4],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
