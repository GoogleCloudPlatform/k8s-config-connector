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

package mocksaasservicemgmt

import (
	"context"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SaasDeploymentsServer struct {
	*MockService
	pb.UnimplementedSaasDeploymentsServer
}

func (s *SaasDeploymentsServer) GetRelease(ctx context.Context, req *pb.GetReleaseRequest) (*pb.Release, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) CreateRelease(ctx context.Context, req *pb.CreateReleaseRequest) (*pb.Release, error) {
	reqName := req.Parent + "/releases/" + req.ReleaseId
	name, err := s.parseReleaseName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Release).(*pb.Release)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.Uid = uuid.NewString()
	obj.Etag = uuid.NewString()

	// Fill observed state fields if required by KRM Spec / GCP alignment
	if obj.Blueprint != nil {
		// Mock engine and version on the blueprint observed state
		obj.Blueprint.Engine = "mock-engine"
		obj.Blueprint.Version = "mock-version"
	}

	// Just some dummy variables to simulate mock behavior for observedState
	obj.InputVariables = []*pb.UnitVariable{
		{
			Variable: "my-var",
			Type:     pb.UnitVariable_STRING,
			Value:    "mock-input-val",
		},
	}
	obj.OutputVariables = []*pb.UnitVariable{
		{
			Variable: "my-output-var",
			Type:     pb.UnitVariable_STRING,
			Value:    "mock-output-val",
		},
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) UpdateRelease(ctx context.Context, req *pb.UpdateReleaseRequest) (*pb.Release, error) {
	reqRelease := req.GetRelease()
	if reqRelease == nil {
		return nil, status.Errorf(codes.InvalidArgument, "release is required")
	}

	name, err := s.parseReleaseName(reqRelease.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		// Default to all paths if empty
		paths = []string{
			"blueprint",
			"release_requirements",
			"input_variable_defaults",
			"labels",
			"annotations",
		}
	}

	for _, path := range paths {
		switch path {
		case "blueprint":
			obj.Blueprint = reqRelease.GetBlueprint()
			if obj.Blueprint != nil {
				obj.Blueprint.Engine = "mock-engine"
				obj.Blueprint.Version = "mock-version"
			}
		case "release_requirements", "releaseRequirements":
			obj.ReleaseRequirements = reqRelease.GetReleaseRequirements()
		case "input_variable_defaults", "inputVariableDefaults":
			obj.InputVariableDefaults = reqRelease.GetInputVariableDefaults()
		case "labels":
			obj.Labels = reqRelease.GetLabels()
		case "annotations":
			obj.Annotations = reqRelease.GetAnnotations()
		}
	}

	obj.UpdateTime = timestamppb.Now()
	obj.Etag = uuid.NewString()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SaasDeploymentsServer) DeleteRelease(ctx context.Context, req *pb.DeleteReleaseRequest) (*emptypb.Empty, error) {
	name, err := s.parseReleaseName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Release{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
