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

package mocklivestream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type dvrSessionName struct {
	Project    *projects.ProjectData
	Location   string
	Channel    string
	DVRSession string
}

func (n *dvrSessionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/channels/" + n.Channel + "/dvrSessions/" + n.DVRSession
}

func (s *MockService) parseDvrSessionName(name string) (*dvrSessionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "channels" && tokens[6] == "dvrSessions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &dvrSessionName{
			Project:    project,
			Location:   tokens[3],
			Channel:    tokens[5],
			DVRSession: tokens[7],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *LivestreamServer) GetDvrSession(ctx context.Context, req *pb.GetDvrSessionRequest) (*pb.DvrSession, error) {
	name, err := s.parseDvrSessionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DvrSession{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource %q was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *LivestreamServer) CreateDvrSession(ctx context.Context, req *pb.CreateDvrSessionRequest) (*longrunning.Operation, error) {
	dvrSessionID := req.GetDvrSessionId()

	reqName := fmt.Sprintf("%s/dvrSessions/%s", req.GetParent(), dvrSessionID)
	name, err := s.parseDvrSessionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetDvrSession()).(*pb.DvrSession)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *LivestreamServer) UpdateDvrSession(ctx context.Context, req *pb.UpdateDvrSessionRequest) (*longrunning.Operation, error) {
	reqObj := req.GetDvrSession()
	name, err := s.parseDvrSessionName(reqObj.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := &pb.DvrSession{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Update fields specified in updateMask if available, or update obj directly
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		proto.Reset(obj)
		proto.Merge(obj, reqObj)
	} else {
		for _, path := range paths {
			switch path {
			case "dvr_manifests", "dvrManifests":
				obj.DvrManifests = reqObj.GetDvrManifests()
			case "dvr_windows", "dvrWindows":
				obj.DvrWindows = reqObj.GetDvrWindows()
			}
		}
	}
	obj.Name = fqn
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *LivestreamServer) DeleteDvrSession(ctx context.Context, req *pb.DeleteDvrSessionRequest) (*longrunning.Operation, error) {
	name, err := s.parseDvrSessionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deletedObj := &pb.DvrSession{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		EndTime:    timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}

	return s.operations.StartLRO(ctx, fqn, metadata, func() (proto.Message, error) {
		return &emptypb.Empty{}, nil
	})
}

func (s *LivestreamServer) ListDvrSessions(ctx context.Context, req *pb.ListDvrSessionsRequest) (*pb.ListDvrSessionsResponse, error) {
	parent := req.GetParent()
	findPrefix := parent + "/dvrSessions/"

	response := &pb.ListDvrSessionsResponse{}
	dvrSessionKind := (&pb.DvrSession{}).ProtoReflect().Descriptor()
	err := s.storage.List(ctx, dvrSessionKind, storage.ListOptions{
		Prefix: findPrefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.DvrSession)
		response.DvrSessions = append(response.DvrSessions, item)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
