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

package mocknetworksecurity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *NetworkSecurityServer) CreateServerTlsPolicy(ctx context.Context, req *pb.CreateServerTlsPolicyRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/serverTlsPolicies/" + req.ServerTlsPolicyId
	name, err := s.parseServerTlsPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ServerTlsPolicy).(*pb.ServerTlsPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1beta1",
	}
	return s.operations.StartLRO(ctx, req.Parent, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.Clone(obj)
		return result, nil
	})
}

func (s *NetworkSecurityServer) GetServerTlsPolicy(ctx context.Context, req *pb.GetServerTlsPolicyRequest) (*pb.ServerTlsPolicy, error) {
	name, err := s.parseServerTlsPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServerTlsPolicy{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *NetworkSecurityServer) UpdateServerTlsPolicy(ctx context.Context, req *pb.UpdateServerTlsPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseServerTlsPolicyName(req.GetServerTlsPolicy().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.ServerTlsPolicy{}
	if err := s.storage.Get(ctx, name.String(), obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(obj).(*pb.ServerTlsPolicy)
	updated.UpdateTime = timestamppb.New(time.Now())

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		updated = proto.Clone(req.GetServerTlsPolicy()).(*pb.ServerTlsPolicy)
		updated.CreateTime = obj.CreateTime
		updated.UpdateTime = timestamppb.New(time.Now())
		updated.Name = obj.Name
	} else {
		for _, path := range paths {
			switch path {
			case "description":
				updated.Description = req.GetServerTlsPolicy().GetDescription()
			case "labels":
				updated.Labels = req.GetServerTlsPolicy().GetLabels()
			case "allowOpen":
				updated.AllowOpen = req.GetServerTlsPolicy().GetAllowOpen()
			case "serverCertificate":
				updated.ServerCertificate = req.GetServerTlsPolicy().GetServerCertificate()
			case "mtlsPolicy":
				updated.MtlsPolicy = req.GetServerTlsPolicy().GetMtlsPolicy()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q is not yet handled in mock", path)
			}
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
		result := proto.Clone(updated)
		return result, nil
	})
}

func (s *NetworkSecurityServer) DeleteServerTlsPolicy(ctx context.Context, req *pb.DeleteServerTlsPolicyRequest) (*longrunning.Operation, error) {
	name, err := s.parseServerTlsPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pb.ServerTlsPolicy{}); err != nil {
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

type serverTlsPolicyName struct {
	Project           *projects.ProjectData
	Location          string
	ServerTlsPolicyID string
}

func (n *serverTlsPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/serverTlsPolicies/" + n.ServerTlsPolicyID
}

func (s *NetworkSecurityServer) parseServerTlsPolicyName(name string) (*serverTlsPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serverTlsPolicies" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &serverTlsPolicyName{
			Project:           project,
			Location:          tokens[3],
			ServerTlsPolicyID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
