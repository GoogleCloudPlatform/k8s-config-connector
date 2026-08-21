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
	"google.golang.org/protobuf/types/known/durationpb"
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
			State:          pb.Session_CREATING,
			StateStartTime: timestamppb.New(now),
		},
	}
	obj.RuntimeInfo = &pb.RuntimeInfo{}

	if obj.Labels == nil {
		obj.Labels = make(map[string]string)
	}
	for k, v := range map[string]string{
		"goog-dataproc-drz-resource-uuid": "session-" + obj.Uuid,
		"goog-dataproc-location":          name.Location,
		"goog-dataproc-session-id":        name.Session,
		"goog-dataproc-session-uuid":      obj.Uuid,
	} {
		if _, ok := obj.Labels[k]; !ok {
			obj.Labels[k] = v
		}
	}
	if obj.EnvironmentConfig == nil {
		obj.EnvironmentConfig = &pb.EnvironmentConfig{}
	}
	if obj.EnvironmentConfig.ExecutionConfig == nil {
		obj.EnvironmentConfig.ExecutionConfig = &pb.ExecutionConfig{}
	}
	if obj.EnvironmentConfig.ExecutionConfig.IdleTtl == nil {
		obj.EnvironmentConfig.ExecutionConfig.IdleTtl = &durationpb.Duration{Seconds: 3600}
	}
	if obj.EnvironmentConfig.ExecutionConfig.Ttl == nil {
		obj.EnvironmentConfig.ExecutionConfig.Ttl = &durationpb.Duration{Seconds: 86400}
	}
	if obj.EnvironmentConfig.ExecutionConfig.ServiceAccount == "" {
		project, _ := s.Projects.GetProjectByID(name.Project)
		if project != nil {
			obj.EnvironmentConfig.ExecutionConfig.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", project.Number)
		} else {
			obj.EnvironmentConfig.ExecutionConfig.ServiceAccount = fmt.Sprintf("%s-compute@developer.gserviceaccount.com", name.Project)
		}
	}
	if obj.EnvironmentConfig.PeripheralsConfig == nil {
		obj.EnvironmentConfig.PeripheralsConfig = &pb.PeripheralsConfig{
			SparkHistoryServerConfig: &pb.SparkHistoryServerConfig{},
		}
	}
	if obj.RuntimeConfig == nil {
		obj.RuntimeConfig = &pb.RuntimeConfig{}
	}
	if obj.RuntimeConfig.Version == "" {
		obj.RuntimeConfig.Version = "2.2.82"
	}
	if obj.RuntimeConfig.Properties == nil {
		obj.RuntimeConfig.Properties = map[string]string{
			"dataproc:dataproc.tier":                                "premium",
			"spark:spark.dataproc.engine":                           "default",
			"spark:spark.dataproc.lightningEngine.runtime":          "default",
			"spark:spark.dataproc.scaling.version":                  "2",
			"spark:spark.driver.cores":                              "4",
			"spark:spark.driver.memory":                             "9600m",
			"spark:spark.dynamicAllocation.executorAllocationRatio": "0.3",
			"spark:spark.executor.cores":                            "4",
			"spark:spark.executor.instances":                        "2",
			"spark:spark.executor.memory":                           "9600m",
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_CREATE,
		Description:   "Create session",
		SessionUuid:   obj.Uuid,
		CreateTime:    timestamppb.New(now),
		DoneTime:      timestamppb.New(now),
		Labels: map[string]string{
			"cnrm-test":                       "true",
			"goog-dataproc-drz-resource-uuid": "session-" + obj.Uuid,
			"goog-dataproc-location":          name.Location,
			"goog-dataproc-session-id":        name.Session,
			"goog-dataproc-session-uuid":      obj.Uuid,
			"managed-by-cnrm":                 "true",
		},
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

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_DELETE,
		Description:   "Delete session",
		SessionUuid:   deleted.Uuid,
		CreateTime:    timestamppb.New(time.Now()),
		DoneTime:      timestamppb.New(time.Now()),
		Labels: map[string]string{
			"cnrm-test":                       "true",
			"goog-dataproc-drz-resource-uuid": "session-" + deleted.Uuid,
			"goog-dataproc-location":          name.Location,
			"goog-dataproc-session-id":        name.Session,
			"goog-dataproc-session-uuid":      deleted.Uuid,
			"managed-by-cnrm":                 "true",
		},
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

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project, name.Location)
	lroMetadata := &pb.SessionOperationMetadata{
		Session:       fqn,
		OperationType: pb.SessionOperationMetadata_TERMINATE,
		Description:   "Terminate session",
		SessionUuid:   obj.Uuid,
		CreateTime:    timestamppb.New(time.Now()),
		DoneTime:      timestamppb.New(time.Now()),
		Labels: map[string]string{
			"cnrm-test":                       "true",
			"goog-dataproc-drz-resource-uuid": "session-" + obj.Uuid,
			"goog-dataproc-location":          name.Location,
			"goog-dataproc-session-id":        name.Session,
			"goog-dataproc-session-uuid":      obj.Uuid,
			"managed-by-cnrm":                 "true",
		},
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
