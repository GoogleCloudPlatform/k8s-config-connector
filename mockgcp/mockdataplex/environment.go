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

// +tool:mockgcp-support
// proto.service: google.cloud.dataplex.v1.DataplexService
// proto.message: google.cloud.dataplex.v1.Environment

package mockdataplex

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

	"github.com/google/uuid"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
)

func (s *DataplexV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DataplexV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/environments/%s", req.GetParent(), req.GetEnvironmentId())
	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetEnvironment()).(*pb.Environment)
	obj.Name = fqn
	obj.Uid = uuid.NewString()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.State_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "create",
		ApiVersion: "v1",
	}
	// todo: verify prefix
	return s.operations.StartLRO(ctx, req.Parent, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.GetEnvironment().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// Only care about fields we support for update.
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetEnvironment().GetDescription()
		case "display_name":
			obj.DisplayName = req.GetEnvironment().GetDisplayName()
		case "labels":
			obj.Labels = req.GetEnvironment().GetLabels()
		case "infrastructure_spec":
			obj.InfrastructureSpec = req.GetEnvironment().GetInfrastructureSpec()
		case "session_spec":
			obj.SessionSpec = req.GetEnvironment().GetSessionSpec()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
		ApiVersion: "v1",
	}
	// todo: verify prefix
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DataplexV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Environment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	metadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
		ApiVersion: "v1",
	}
	// todo: verify prefix
	return s.operations.StartLRO(ctx, "", metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(now)
		return &emptypb.Empty{}, nil
	})
}

type environmentName struct {
	Lake          string
	EnvironmentID string
}

func (n *environmentName) String() string {
	return fmt.Sprintf("%s/environments/%s", n.Lake, n.EnvironmentID)
}

// parseEnvironmentName parses a string into an environmentName.
// The expected form is `projects/*/locations/*/lakes/*/environments/*`.
func (s *MockService) parseEnvironmentName(name string) (*environmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lakes" && tokens[6] == "environments" {
		name := &environmentName{
			Lake:          strings.Join(tokens[:len(tokens)-2], "/"),
			EnvironmentID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
