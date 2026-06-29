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

package mockosconfig

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type OSConfigServer struct {
	*MockService
	pb.UnimplementedOsConfigServiceServer
}

func (s *OSConfigServer) GetGuestPolicy(ctx context.Context, req *pb.GetGuestPolicyRequest) (*pb.GuestPolicy, error) {
	name, err := s.parseGuestPolicyName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.GuestPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "GuestPolicy %s not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *OSConfigServer) ListGuestPolicies(ctx context.Context, req *pb.ListGuestPoliciesRequest) (*pb.ListGuestPoliciesResponse, error) {
	var resources []*pb.GuestPolicy
	kind := (&pb.GuestPolicy{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		gp := obj.(*pb.GuestPolicy)
		// Match parent/project prefix
		if strings.HasPrefix(gp.Name, req.Parent+"/") {
			resources = append(resources, gp)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.ListGuestPoliciesResponse{
		GuestPolicies: resources,
	}, nil
}

func (s *OSConfigServer) CreateGuestPolicy(ctx context.Context, req *pb.CreateGuestPolicyRequest) (*pb.GuestPolicy, error) {
	name, err := s.parseGuestPolicyName(req.Parent + "/guestPolicies/" + req.GuestPolicyId)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.Clone(req.GuestPolicy).(*pb.GuestPolicy)
	obj.Name = fqn

	t := timestamppb.New(time.Now())
	obj.CreateTime = t
	obj.UpdateTime = t

	// Generate etag
	obj.Etag = s.computeEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *OSConfigServer) UpdateGuestPolicy(ctx context.Context, req *pb.UpdateGuestPolicyRequest) (*pb.GuestPolicy, error) {
	name, err := s.parseGuestPolicyName(req.GuestPolicy.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.GuestPolicy{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.GuestPolicy).(*pb.GuestPolicy)
	updated.Name = fqn
	updated.CreateTime = obj.CreateTime
	updated.UpdateTime = timestamppb.New(time.Now())
	updated.Etag = s.computeEtag(updated)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *OSConfigServer) DeleteGuestPolicy(ctx context.Context, req *pb.DeleteGuestPolicyRequest) (*emptypb.Empty, error) {
	name, err := s.parseGuestPolicyName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	oldObj := &pb.GuestPolicy{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OSConfigServer) computeEtag(obj *pb.GuestPolicy) string {
	h := md5.New()
	h.Write([]byte(obj.Name + obj.Description))
	return hex.EncodeToString(h.Sum(nil))[:16]
}

type guestPolicyName struct {
	Project     *projects.ProjectData
	GuestPolicy string
}

func (n *guestPolicyName) String() string {
	return fmt.Sprintf("projects/%s/guestPolicies/%s", n.Project.ID, n.GuestPolicy)
}

func (s *MockService) parseGuestPolicyName(name string) (*guestPolicyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "guestPolicies" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &guestPolicyName{
			Project:     project,
			GuestPolicy: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid guest policy name %q", name)
}
