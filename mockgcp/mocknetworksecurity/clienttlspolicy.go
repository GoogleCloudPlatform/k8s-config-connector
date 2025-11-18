// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:mockgcp-support
// proto.service: google.cloud.networksecurity.v1beta1.NetworkSecurity
// proto.message: google.cloud.networksecurity.v1beta1.ClientTlsPolicy

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityServer) CreateClientTlsPolicy(ctx context.Context, req *pb.CreateClientTlsPolicyRequest) (*longrunning.Operation, error) {
	name := req.Parent + "/clientTlsPolicies/" + req.ClientTlsPolicyId

	fqn := name

	now := time.Now()

	obj := ProtoClone(req.ClientTlsPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1beta1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := ProtoClone(obj)
		return result, nil
	})
}

func (s *NetworkSecurityServer) GetClientTlsPolicy(ctx context.Context, req *pb.GetClientTlsPolicyRequest) (*pb.ClientTlsPolicy, error) {
	name, err := s.parseClientTlsPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ClientTlsPolicy{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *NetworkSecurityServer) ListClientTlsPolicies(ctx context.Context, req *pb.ListClientTlsPoliciesRequest) (*pb.ListClientTlsPoliciesResponse, error) {
	prefixName, err := s.parseClientTlsPolicyName(req.Parent + "/clientTlsPolicies/dummy")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(prefixName.String(), "dummy")

	response := &pb.ListClientTlsPoliciesResponse{}

	filterKind := (&pb.ClientTlsPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, filterKind, storage.ListOptions{Prefix: prefix}, func(obj protoreflect.ProtoMessage) error {
		clientTlsPolicy := obj.(*pb.ClientTlsPolicy)
		response.ClientTlsPolicies = append(response.ClientTlsPolicies, clientTlsPolicy)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *NetworkSecurityServer) UpdateClientTlsPolicy(ctx context.Context, req *pb.UpdateClientTlsPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseClientTlsPolicyName(req.GetClientTlsPolicy().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.ClientTlsPolicy{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := ProtoClone(obj)
	updated.UpdateTime = timestamppb.New(time.Now())

	updatePaths := req.GetUpdateMask().GetPaths()
	if len(updatePaths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "no update paths specified")
	}
	for _, path := range updatePaths {
		switch req.GetUpdateMask().GetPaths()[0] {
		case "sni":
			updated.Sni = req.GetClientTlsPolicy().GetSni()
		case "client_certificate":
			updated.ClientCertificate = req.GetClientTlsPolicy().GetClientCertificate()
		case "server_validation_ca":
			updated.ServerValidationCa = req.GetClientTlsPolicy().GetServerValidationCa()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
		}
	}

	if err := s.storage.Update(ctx, name.String(), updated); err != nil {
		return nil, err
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1beta1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := ProtoClone(updated)
		return result, nil
	})
}

func (s *NetworkSecurityServer) DeleteClientTlsPolicy(ctx context.Context, req *pb.DeleteClientTlsPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseClientTlsPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pb.ClientTlsPolicy{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1beta1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type clientTlsPolicyName struct {
	Project           *projects.ProjectData
	Location          string
	ClientTlsPolicyID string
}

func (n *clientTlsPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/clientTlsPolicies/" + n.ClientTlsPolicyID
}

func (s *NetworkSecurityServer) parseClientTlsPolicyName(name string) (*clientTlsPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "clientTlsPolicies" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &clientTlsPolicyName{
			Project:           project,
			Location:          tokens[3],
			ClientTlsPolicyID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
