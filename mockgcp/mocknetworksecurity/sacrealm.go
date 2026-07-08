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

	pbv1 "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SSERealmServer implements the SSERealmServiceServer gRPC interface,
// providing the mock implementation and alignment verification for NetworkSecuritySACRealm under Phase 3.
type SSERealmServer struct {
	*MockService
	pbv1.UnimplementedSSERealmServiceServer
}

func (s *SSERealmServer) CreateSACRealm(ctx context.Context, req *pbv1.CreateSACRealmRequest) (*longrunning.Operation, error) {
	project, location, err := s.parseParent(req.Parent)
	if err != nil {
		return nil, err
	}

	name := req.Parent + "/sacRealms/" + req.SacRealmId
	fqn := name

	obj := proto.CloneOf(req.SacRealm)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pbv1.SACRealm_PENDING_PARTNER_ATTACHMENT
	obj.PairingKey = &pbv1.SACRealm_PairingKey{
		Key:        "dummy-pairing-key",
		ExpireTime: timestamppb.New(time.Now().Add(7 * 24 * time.Hour)),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name,
		Verb:                  "create",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", project.ID, location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := proto.CloneOf(obj)
		return result, nil
	})
}

func (s *SSERealmServer) GetSACRealm(ctx context.Context, req *pbv1.GetSACRealmRequest) (*pbv1.SACRealm, error) {
	name, err := s.parseSACRealmName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pbv1.SACRealm{}
	obj.Name = fqn
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *SSERealmServer) DeleteSACRealm(ctx context.Context, req *pbv1.DeleteSACRealmRequest) (*longrunning.Operation, error) {
	name, err := s.parseSACRealmName(req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Delete(ctx, name.String(), &pbv1.SACRealm{}); err != nil {
		return nil, err
	}
	now := time.Now()
	lroMetadata := &pbv1.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (protoreflect.ProtoMessage, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type sacRealmName struct {
	Project    *projects.ProjectData
	Location   string
	SACRealmID string
}

func (n *sacRealmName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/sacRealms/" + n.SACRealmID
}

func (s *SSERealmServer) parseSACRealmName(name string) (*sacRealmName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sacRealms" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, err
		}
		name := &sacRealmName{
			Project:    project,
			Location:   tokens[3],
			SACRealmID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SSERealmServer) parseParent(parent string) (*projects.ProjectData, string, error) {
	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProject(&projects.ProjectName{ProjectID: tokens[1]})
		if err != nil {
			return nil, "", err
		}
		return project, tokens[3], nil
	}
	return nil, "", status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}
