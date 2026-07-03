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
// proto.service: google.cloud.dataproc.v1.SessionController
// proto.message: google.cloud.dataproc.v1.Session

package mockdataproc

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sessionControllerServer struct {
	*MockService
	pb.UnimplementedSessionControllerServer
}

func (s *sessionControllerServer) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.Session, error) {
	name, err := s.parseSessionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Session{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *sessionControllerServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSessionName(req.Parent + "/sessions/" + req.SessionId)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.Session).(*pb.Session)
	obj.Name = fqn
	obj.Uuid = "00000000-0000-0000-0000-000000000001"
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.Session_ACTIVE
	obj.StateTime = timestamppb.New(now)
	obj.Creator = "test-user@google.com"
	obj.StateHistory = []*pb.Session_SessionStateHistory{
		{
			State:          pb.Session_ACTIVE,
			StateStartTime: timestamppb.New(now),
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_SESSION_OPERATION_TYPE_UNSPECIFIED,
		Description:   "Create Session",
		SessionUuid:   obj.Uuid,
		CreateTime:    timestamppb.New(now),
		DoneTime:      timestamppb.New(now),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sessionControllerServer) DeleteSession(ctx context.Context, req *pb.DeleteSessionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSessionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Session{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_SESSION_OPERATION_TYPE_UNSPECIFIED,
		Description:   "Delete Session",
		CreateTime:    timestamppb.New(time.Now()),
		DoneTime:      timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *sessionControllerServer) TerminateSession(ctx context.Context, req *pb.TerminateSessionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseSessionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Session{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.Session_TERMINATED
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_SESSION_OPERATION_TYPE_UNSPECIFIED,
		Description:   "Terminate Session",
		CreateTime:    timestamppb.New(time.Now()),
		DoneTime:      timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sessionControllerServer) ListSessions(ctx context.Context, req *pb.ListSessionsRequest) (*pb.ListSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSessions not implemented")
}

type sessionName struct {
	Project  string
	Location string
	Session  string
}

func (n *sessionName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/sessions/%s", n.Project, n.Location, n.Session)
}

func (s *MockService) parseSessionName(name string) (*sessionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "sessions" {
		return &sessionName{
			Project:  tokens[1],
			Location: tokens[3],
			Session:  tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
