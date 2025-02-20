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

package mocksecuresourcemanager

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/securesourcemanager/v1"
)

func (s *secureSourceManagerServer) GetInstance(ctx context.Context, req *pb.GetInstanceRequest) (*pb.Instance, error) {
	name, err := s.parseInstanceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Instance{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *secureSourceManagerServer) CreateInstance(ctx context.Context, req *pb.CreateInstanceRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/instances/" + req.InstanceId
	name, err := s.parseInstanceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Instance).(*pb.Instance)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// TODO: State should be Creating at first, ACTIVE once done
	obj.State = pb.Instance_ACTIVE

	if req.GetInstance().GetKmsKey() != "" {
		obj.KmsKey = req.GetInstance().GetKmsKey()
	}

	if req.GetInstance().GetPrivateConfig() != nil {
		obj.PrivateConfig.IsPrivate = req.GetInstance().GetPrivateConfig().GetIsPrivate()
		obj.PrivateConfig.CaPool = req.GetInstance().GetPrivateConfig().GetCaPool()
		obj.PrivateConfig.HttpServiceAttachment = fmt.Sprintf("projects/tp-project/regions/%s/serviceAttachments/httpAttachment", name.Location)
		obj.PrivateConfig.SshServiceAttachment = fmt.Sprintf("projects/tp-project/regions/%s/serviceAttachments/sshAttachment", name.Location)
	}

	// TODO: Only fill in when ACTIVE
	prefix := fmt.Sprintf("%s-%d", name.InstanceID, name.Project.Number)
	domain := "." + name.Location + ".sourcemanager.dev"
	obj.HostConfig = &pb.Instance_HostConfig{
		Html:    prefix + domain,
		Api:     prefix + "-api" + domain,
		GitHttp: prefix + "-git" + domain,
		GitSsh:  prefix + "-ssh" + domain,
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Verify name
	//  Error 400: resource name is invalid; instance id must match regex '[a-z]([a-z0-9-]{0,38}[a-z0-9])?'\nerror details: name = PreconditionFailure type = INVALID_RESOURCE_NAME subj = projects/<number>/locations/us-central1/instances/securesourcemanagerinstance-xna5xqnqeidom6wiwxgq desc = resource name is invalid; instance id must match regex '[a-z]([a-z0-9-]{0,38}[a-z0-9])?

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *secureSourceManagerServer) DeleteInstance(ctx context.Context, req *pb.DeleteInstanceRequest) (*longrunning.Operation, error) {
	name, err := s.parseInstanceName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.Instance{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// TODO: State should be Deleting at first

	op := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, op, func() (proto.Message, error) {
		op.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type InstanceName struct {
	Project    *projects.ProjectData
	Location   string
	InstanceID string
}

func (n *InstanceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.Location, n.InstanceID)
}

// func (n *InstanceName) Target() string {
// 	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", n.Project.ID, n.Location, n.InstanceID)
// }

// parseInstanceName parses a string into a InstanceName.
// The expected form is projects/*/locations/*/instances/*
func (s *MockService) parseInstanceName(name string) (*InstanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &InstanceName{
			Project:    project,
			Location:   tokens[3],
			InstanceID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
