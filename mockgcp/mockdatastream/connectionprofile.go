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

//go:build mockgcp
// +build mockgcp

// Package mockdatastream is a mock for the datastream service.
//
//	proto.service: google.cloud.datastream.v1.Datastream
//
// proto.message: google.cloud.datastream.v1.ConnectionProfile
package mockdatastream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/datastream/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *DatastreamV1) GetConnectionProfile(ctx context.Context, req *pb.GetConnectionProfileRequest) (*pb.ConnectionProfile, error) {
	name, err := s.parseConnectionProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *DatastreamV1) CreateConnectionProfile(ctx context.Context, req *pb.CreateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/connectionProfiles/%s", req.GetParent(), req.GetConnectionProfileId())
	name, err := s.parseConnectionProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetConnectionProfile()).(*pb.ConnectionProfile)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) UpdateConnectionProfile(ctx context.Context, req *pb.UpdateConnectionProfileRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetConnectionProfile().GetName()
	name, err := s.parseConnectionProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ConnectionProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// We clone before, so we can check after, in case the update_mask is empty
	original := proto.Clone(obj).(*pb.ConnectionProfile)

	// Apply the update mask to the object, using a library function.
	// Note: FieldMask.IsValid doesn't work yet.
	updateMask := req.GetUpdateMask()
	if updateMask == nil || len(updateMask.Paths) == 0 {
		updateMask = &fieldmaskpb.FieldMask{}
		for _, p := range allPaths(original) {
			updateMask.Paths = append(updateMask.Paths, p)
		}
	}

	if !updateMask.IsValid(obj) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid fieldmask")
	}

	// Actually do the update
	if err := updateMask.Merge(req.GetConnectionProfile(), obj); err != nil {
		return nil, fmt.Errorf("error applying fieldmask: %w", err)
	}

	// Some fields are "output only", so we preserve their values.
	preserveOutputOnlyFields(original, obj)
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *DatastreamV1) DeleteConnectionProfile(ctx context.Context, req *pb.DeleteConnectionProfileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConnectionProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.ConnectionProfile{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return nil, nil // Operation is 'done' immediately.
	})
}

type connectionProfileName struct {
	Project             *projects.ProjectData
	Location            string
	ConnectionProfileID string
}

func (n *connectionProfileName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/connectionProfiles/" + n.ConnectionProfileID
}

// parseConnectionProfileName parses a string into a connectionProfileName.
// The expected form is `projects/*/locations/*/connectionProfiles/*`.
func (s *MockService) parseConnectionProfileName(name string) (*connectionProfileName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectionProfiles" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &connectionProfileName{
			Project:             project,
			Location:            tokens[3],
			ConnectionProfileID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

// preserveOutputOnlyFields copies known output-only fields from src to dest.
// This function is hand-maintained.
func preserveOutputOnlyFields(src *pb.ConnectionProfile, dest *pb.ConnectionProfile) {
	dest.Name = src.Name // Field 1
}

// allPaths returns all the paths for a resource.
// This function is hand-maintained; we can replace with protoreflect to
// avoid manual maintenance, but adds a dependency.
func allPaths(obj *pb.ConnectionProfile) []string {
	return []string{
		"display_name",
		"oracle_profile",
		"gcs_profile",
		"mysql_profile",
		"bigquery_profile",
		"postgresql_profile",
		"static_service_ip_connectivity",
		"forward_ssh_connectivity",
		"private_connectivity",
		"labels",
		"sql_server_profile",
	}
}
