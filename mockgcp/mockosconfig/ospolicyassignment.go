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

package mockosconfig

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/osconfig/v1"
	"github.com/google/uuid"
)

func (s *OSConfigZonalService) GetOSPolicyAssignment(ctx context.Context, req *pb.GetOSPolicyAssignmentRequest) (*pb.OSPolicyAssignment, error) {
	name, err := s.parseOSPolicyAssignmentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.OSPolicyAssignment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "osPolicyAssignment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading osPolicyAssignment: %v", err)
		}
	}

	return obj, nil
}

func (s *OSConfigZonalService) CreateOSPolicyAssignment(ctx context.Context, req *pb.CreateOSPolicyAssignmentRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/osPolicyAssignments/" + req.OsPolicyAssignmentId
	name, err := s.parseOSPolicyAssignmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.OsPolicyAssignment).(*pb.OSPolicyAssignment)
	obj.Name = fqn
	obj.Etag = computeEtag(obj)

	obj.RevisionCreateTime = timestamppb.New(now)
	obj.RevisionId = uuid.NewString()
	obj.RolloutState = pb.OSPolicyAssignment_IN_PROGRESS
	obj.Uid = uuid.NewString()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating osPolicyAssignment: %v", err)
	}

	return s.operations.StartLRO(ctx, func() (proto.Message, error) {
		obj.Baseline = true
		obj.RolloutState = pb.OSPolicyAssignment_SUCCEEDED

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *OSConfigZonalService) UpdateOSPolicyAssignment(ctx context.Context, req *pb.UpdateOSPolicyAssignmentRequest) (*longrunning.Operation, error) {
	name, err := s.parseOSPolicyAssignmentName(req.GetOsPolicyAssignment().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.OSPolicyAssignment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "osPolicyAssignment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading osPolicyAssignment: %v", err)
		}
	}

	updateMask := req.GetUpdateMask()
	for _, path := range updateMask.Paths {
		switch path {
		case "instanceFilter":
			obj.InstanceFilter = req.GetOsPolicyAssignment().GetInstanceFilter()
		case "rollout":
			obj.Rollout = req.GetOsPolicyAssignment().GetRollout()
		case "osPolicies":
			obj.OsPolicies = req.GetOsPolicyAssignment().GetOsPolicies()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field path %q is not supported", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating osPolicyAssignment: %v", err)
	}

	return s.operations.StartLRO(ctx, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *OSConfigZonalService) DeleteOSPolicyAssignment(ctx context.Context, req *pb.DeleteOSPolicyAssignmentRequest) (*longrunning.Operation, error) {
	name, err := s.parseOSPolicyAssignmentName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.OSPolicyAssignment{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "osPolicyAssignment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting osPolicyAssignment: %v", err)
		}
	}

	return s.operations.StartLRO(ctx, func() (proto.Message, error) {
		return deleted, nil
	})
}

type osPolicyAssignmentName struct {
	Project                *projects.ProjectData
	Location               string
	OSPolicyAssignmentName string
}

func (n *osPolicyAssignmentName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/osPolicyAssignments/" + n.OSPolicyAssignmentName
}

// parseOSPolicyAssignmentName parses a string into a osPolicyAssignmentName.
// The expected form is projects/{project}/locations/{location}/osPolicyAssignments/{os_policy_assignment}@{revisionId}
func (s *MockService) parseOSPolicyAssignmentName(name string) (*osPolicyAssignmentName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "osPolicyAssignments" {
		project, err := s.projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &osPolicyAssignmentName{
			Project:                project,
			Location:               tokens[3],
			OSPolicyAssignmentName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func computeEtag(obj *pb.OSPolicyAssignment) string {
	b, err := proto.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("converting to proto: %v", err))
	}
	hash := md5.Sum(b)
	return base64.StdEncoding.EncodeToString(hash[:])
}
