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

package mockdevicestreaming

import (
	"context"
	"math/rand"
	"strings"
	"time"

	pb "cloud.google.com/go/devicestreaming/apiv1/devicestreamingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type DirectAccessV1 struct {
	*MockService

	pb.UnimplementedDirectAccessServiceServer
}

func randomSessionID() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, 13)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return "session-" + string(b)
}

// Creates a new DeviceSession.
func (s *DirectAccessV1) CreateDeviceSession(ctx context.Context, req *pb.CreateDeviceSessionRequest) (*pb.DeviceSession, error) {
	sessionID := req.DeviceSessionId
	if sessionID == "" {
		sessionID = randomSessionID()
	}

	reqSessionName := req.Parent + "/deviceSessions/" + sessionID

	name, err := s.parseDeviceStreamingSessionName(reqSessionName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.DeviceSession)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.State = pb.DeviceSession_ACTIVE // Always ACTIVE in mock

	if obj.GetTtl() != nil {
		ttl := obj.GetTtl().AsDuration()
		expireTime := obj.CreateTime.AsTime().Add(ttl)
		obj.Expiration = &pb.DeviceSession_ExpireTime{
			ExpireTime: timestamppb.New(expireTime),
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// Gets a DeviceSession.
func (s *DirectAccessV1) GetDeviceSession(ctx context.Context, req *pb.GetDeviceSessionRequest) (*pb.DeviceSession, error) {
	name, err := s.parseDeviceStreamingSessionName(req.Name)
	if err != nil {
		return nil, err
	}

	var session pb.DeviceSession
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &session); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return &session, nil
}

// Updates a DeviceSession.
func (s *DirectAccessV1) UpdateDeviceSession(ctx context.Context, req *pb.UpdateDeviceSessionRequest) (*pb.DeviceSession, error) {
	name, err := s.parseDeviceStreamingSessionName(req.DeviceSession.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.DeviceSession{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(existing)
	updated.Name = name.String()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}
	for _, path := range paths {
		switch path {
		case "ttl", "expire_time", "expireTime":
			updated.Expiration = req.DeviceSession.Expiration
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if updated.GetTtl() != nil {
		ttl := updated.GetTtl().AsDuration()
		var createTime time.Time
		if updated.CreateTime != nil {
			createTime = updated.CreateTime.AsTime()
		} else {
			createTime = timestamppb.Now().AsTime()
		}
		expireTime := createTime.Add(ttl)
		updated.Expiration = &pb.DeviceSession_ExpireTime{
			ExpireTime: timestamppb.New(expireTime),
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

// Cancels (deletes) a DeviceSession.
func (s *DirectAccessV1) CancelDeviceSession(ctx context.Context, req *pb.CancelDeviceSessionRequest) (*emptypb.Empty, error) {
	name, err := s.parseDeviceStreamingSessionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.DeviceSession{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type deviceStreamingSessionName struct {
	Project       *projects.ProjectData
	DeviceSession string
}

func (n *deviceStreamingSessionName) String() string {
	return "projects/" + n.Project.ID + "/deviceSessions/" + n.DeviceSession
}

func (s *MockService) parseDeviceStreamingSessionName(name string) (*deviceStreamingSessionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "deviceSessions" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &deviceStreamingSessionName{
			Project:       project,
			DeviceSession: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name %q is not valid. expected format is projects/<projectID>/deviceSessions/<devicesession>",
			name,
		)
	}
}
