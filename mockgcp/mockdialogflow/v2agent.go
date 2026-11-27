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

package mockdialogflow

import (
	"context"
	"strings"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type agentsServer struct {
	*MockService
	pb.UnimplementedAgentsServer
}

func (s *agentsServer) populateDefaults(obj *pb.Agent) {
	if obj.ApiVersion == pb.Agent_API_VERSION_UNSPECIFIED {
		obj.ApiVersion = pb.Agent_API_VERSION_V2
	}
	if obj.Tier == pb.Agent_TIER_UNSPECIFIED {
		obj.Tier = pb.Agent_TIER_STANDARD
	}
}

func (s *agentsServer) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.Agent, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	fqn := parent + "/agent"
	obj := &pb.Agent{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			projectID := strings.TrimPrefix(parent, "projects/")
			if idx := strings.Index(projectID, "/"); idx != -1 {
				projectID = projectID[:idx]
			}
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: No DesignTimeAgent found for project '%s'.", projectID)
		}
		return nil, err
	}

	return obj, nil
}

func (s *agentsServer) SetAgent(ctx context.Context, req *pb.SetAgentRequest) (*pb.Agent, error) {
	agent := req.GetAgent()
	if agent == nil {
		return nil, status.Errorf(codes.InvalidArgument, "agent is required")
	}
	parent := agent.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	fqn := parent + "/agent"
	existing := &pb.Agent{}
	err := s.storage.Get(ctx, fqn, existing)
	if err != nil && status.Code(err) != codes.NotFound {
		return nil, err
	}

	updated := proto.Clone(agent).(*pb.Agent)
	s.populateDefaults(updated)

	if err != nil {
		if err := s.storage.Create(ctx, fqn, updated); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Update(ctx, fqn, updated); err != nil {
			return nil, err
		}
	}

	return updated, nil
}

func (s *agentsServer) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*emptypb.Empty, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent is required")
	}

	fqn := parent + "/agent"
	old := &pb.Agent{}
	if err := s.storage.Delete(ctx, fqn, old); err != nil {
		if status.Code(err) == codes.NotFound {
			projectID := strings.TrimPrefix(parent, "projects/")
			if idx := strings.Index(projectID, "/"); idx != -1 {
				projectID = projectID[:idx]
			}
			return nil, status.Errorf(codes.NotFound, "com.google.apps.framework.request.NotFoundException: No DesignTimeAgent found for project '%s'.", projectID)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
