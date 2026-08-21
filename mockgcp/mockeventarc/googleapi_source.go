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
// proto.service: google.cloud.eventarc.v1.Eventarc
// proto.message: google.cloud.eventarc.v1.GoogleApiSource

package mockeventarc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *EventarcV1) GetGoogleApiSource(ctx context.Context, req *pb.GetGoogleApiSourceRequest) (*pb.GoogleApiSource, error) {
	name, err := s.parseGoogleApiSourceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleApiSource{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *EventarcV1) CreateGoogleApiSource(ctx context.Context, req *pb.CreateGoogleApiSourceRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/googleApiSources/%s", req.GetParent(), req.GetGoogleApiSourceId())
	name, err := s.parseGoogleApiSourceName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()
	obj := proto.CloneOf(req.GetGoogleApiSource())
	obj.Name = fqn
	obj.Uid = name.GoogleApiSource
	obj.Etag = "mock-etag-value"
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Return an LRO that doesnt finish immediately
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lro, err := s.operations.NewLRO(ctx)
	if err != nil {
		return nil, err
	}
	lro.Done = false
	lro.Metadata, err = anypb.New(lroMetadata)
	if err != nil {
		return nil, err
	}
	// Use the fully qualified type name to ensure compatibility with the expected output.
	lro.Metadata.TypeUrl = "type.googleapis.com/google.cloud.eventarc.v1.OperationMetadata"
	lro.Name = fmt.Sprintf("projects/%s/locations/%s/operations/%s", name.Project.ID, name.Location, strings.Split(lro.Name, "/")[len(strings.Split(lro.Name, "/"))-1])

	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *EventarcV1) UpdateGoogleApiSource(ctx context.Context, req *pb.UpdateGoogleApiSourceRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetGoogleApiSource().GetName()
	name, err := s.parseGoogleApiSourceName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.GoogleApiSource{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetGoogleApiSource().GetDisplayName()
		case "labels":
			obj.Labels = req.GetGoogleApiSource().GetLabels()
		case "annotations":
			obj.Annotations = req.GetGoogleApiSource().GetAnnotations()
		case "destination":
			obj.Destination = req.GetGoogleApiSource().GetDestination()
		case "crypto_key_name", "cryptoKeyName":
			obj.CryptoKeyName = req.GetGoogleApiSource().GetCryptoKeyName()
		case "logging_config", "loggingConfig":
			obj.LoggingConfig = req.GetGoogleApiSource().GetLoggingConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported for update", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroRet := proto.CloneOf(obj)
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(time.Now()),
		Target:     fqn,
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return lroRet, nil
	})
}

func (s *EventarcV1) DeleteGoogleApiSource(ctx context.Context, req *pb.DeleteGoogleApiSourceRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseGoogleApiSourceName(req.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	deletedObj := &pb.GoogleApiSource{}

	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     fqn,
		Verb:       "delete",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return deletedObj, nil
	})
}

type googleApiSourceName struct {
	Project         *projects.ProjectData
	Location        string
	GoogleApiSource string
}

func (n *googleApiSourceName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/googleApiSources/%s", n.Project.ID, n.Location, n.GoogleApiSource)
}

// parseGoogleApiSourceName parses a string into a googleApiSourceName.
// The expected form is `projects/*/locations/*/googleApiSources/*`.
func (s *MockService) parseGoogleApiSourceName(name string) (*googleApiSourceName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "googleApiSources" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &googleApiSourceName{
			Project:         project,
			Location:        tokens[3],
			GoogleApiSource: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
