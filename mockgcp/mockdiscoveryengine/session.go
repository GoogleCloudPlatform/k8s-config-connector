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

// +tool:mockgcp-support
// proto.service: google.cloud.discoveryengine.v1.ConversationalSearchService
// proto.message: google.cloud.discoveryengine.v1.Session

package mockdiscoveryengine

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

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type conversationalSearchService struct {
	*MockService
	pb.UnimplementedConversationalSearchServiceServer
}

func (s *conversationalSearchService) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
	name, err := s.parseSessionName(req.GetSession().GetName())
	if err != nil {
		return nil, err
	}
	now := time.Now()

	// Real GCP ignores the client-specified session ID and assigns a numeric session ID.
	sessionID := fmt.Sprintf("%d", now.UnixNano())
	name.Session = sessionID

	fqn := name.String()
	obj := proto.Clone(req.GetSession()).(*pb.Session)
	obj.Name = fqn
	obj.State = pb.Session_IN_PROGRESS
	obj.StartTime = timestamppb.New(now)
	obj.EndTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *conversationalSearchService) DeleteSession(ctx context.Context, req *pb.DeleteSessionRequest) (*emptypb.Empty, error) {
	name, err := s.parseSessionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Session{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *conversationalSearchService) UpdateSession(ctx context.Context, req *pb.UpdateSessionRequest) (*pb.Session, error) {
	name, err := s.parseSessionName(req.GetSession().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Session{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Session %q not found", name)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) > 0 {
		for _, path := range paths {
			switch path {
			case "state":
				obj.State = req.GetSession().GetState()
			case "display_name", "displayName":
				obj.DisplayName = req.GetSession().GetDisplayName()
			}
		}
	} else {
		// Simple merge of updated fields
		proto.Merge(obj, req.GetSession())
	}
	obj.Name = fqn

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *conversationalSearchService) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.Session, error) {
	name, err := s.parseSessionName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Session{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Session %v not found.", name)
		}
		return nil, err
	}
	return obj, nil
}

type sessionName struct {
	Project    *projects.ProjectData
	Location   string
	Collection string
	DataStore  string
	Session    string
}

func (n *sessionName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/collections/%s/dataStores/%s/sessions/%s", n.Project.Number, n.Location, n.Collection, n.DataStore, n.Session)
}

func (s *MockService) parseSessionName(name string) (*sessionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "collections" && tokens[6] == "dataStores" && tokens[8] == "sessions" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		return &sessionName{
			Project:    project,
			Location:   tokens[3],
			Collection: tokens[5],
			DataStore:  tokens[7],
			Session:    tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid session name: %q", name)
}
