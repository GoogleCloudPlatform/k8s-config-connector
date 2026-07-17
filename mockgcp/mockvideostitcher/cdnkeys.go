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

package mockvideostitcher

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type VideoStitcherV1 struct {
	*MockService
	pb.UnimplementedVideoStitcherServiceServer
}

func (s *VideoStitcherV1) GetCdnKey(ctx context.Context, req *pb.GetCdnKeyRequest) (*pb.CdnKey, error) {
	name, err := s.parseCdnKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CdnKey{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	// Clone to avoid mutating stored object
	ret := proto.CloneOf(obj)
	// Strip input-only/secret fields
	if ret.GetGoogleCdnKey() != nil {
		ret.GetGoogleCdnKey().PrivateKey = nil
	}
	if ret.GetMediaCdnKey() != nil {
		ret.GetMediaCdnKey().PrivateKey = nil
	}
	if ret.GetAkamaiCdnKey() != nil {
		ret.GetAkamaiCdnKey().TokenKey = nil
	}

	return ret, nil
}

func (s *VideoStitcherV1) CreateCdnKey(ctx context.Context, req *pb.CreateCdnKeyRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/cdnKeys/" + req.CdnKeyId
	name, err := s.parseCdnKeyName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.CloneOf(req.CdnKey)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Verb:       "create",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		// Clone to avoid returning internal state
		ret := proto.CloneOf(obj)
		if ret.GetGoogleCdnKey() != nil {
			ret.GetGoogleCdnKey().PrivateKey = nil
		}
		if ret.GetMediaCdnKey() != nil {
			ret.GetMediaCdnKey().PrivateKey = nil
		}
		if ret.GetAkamaiCdnKey() != nil {
			ret.GetAkamaiCdnKey().TokenKey = nil
		}
		return ret, nil
	})
}

func (s *VideoStitcherV1) UpdateCdnKey(ctx context.Context, req *pb.UpdateCdnKeyRequest) (*longrunning.Operation, error) {
	reqName := req.GetCdnKey().GetName()

	name, err := s.parseCdnKeyName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := &pb.CdnKey{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	for _, path := range paths {
		switch path {
		case "hostname":
			obj.Hostname = req.GetCdnKey().GetHostname()
		case "google_cdn_key", "googleCdnKey", "cdn_key_config":
			obj.CdnKeyConfig = req.GetCdnKey().GetCdnKeyConfig()
		case "media_cdn_key", "mediaCdnKey":
			obj.CdnKeyConfig = req.GetCdnKey().GetCdnKeyConfig()
		case "akamai_cdn_key", "akamaiCdnKey":
			obj.CdnKeyConfig = req.GetCdnKey().GetCdnKeyConfig()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Verb:       "update",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		// Clone to avoid returning internal state
		ret := proto.CloneOf(obj)
		if ret.GetGoogleCdnKey() != nil {
			ret.GetGoogleCdnKey().PrivateKey = nil
		}
		if ret.GetMediaCdnKey() != nil {
			ret.GetMediaCdnKey().PrivateKey = nil
		}
		if ret.GetAkamaiCdnKey() != nil {
			ret.GetAkamaiCdnKey().TokenKey = nil
		}
		return ret, nil
	})
}

func (s *VideoStitcherV1) DeleteCdnKey(ctx context.Context, req *pb.DeleteCdnKeyRequest) (*longrunning.Operation, error) {
	name, err := s.parseCdnKeyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.CdnKey{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Verb:       "delete",
		Target:     fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *VideoStitcherV1) ListCdnKeys(ctx context.Context, req *pb.ListCdnKeysRequest) (*pb.ListCdnKeysResponse, error) {
	objs := []*pb.CdnKey{}
	kind := (&pb.CdnKey{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(obj proto.Message) error {
		objs = append(objs, obj.(*pb.CdnKey))
		return nil
	}); err != nil {
		return nil, err
	}

	// Filter by parent (projects/*/locations/*)
	// Parent form: projects/<projectID>/locations/<region>
	tokens := strings.Split(req.Parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectID := tokens[1]
		location := tokens[3]
		filtered := []*pb.CdnKey{}
		for _, obj := range objs {
			name, err := s.parseCdnKeyName(obj.Name)
			if err == nil && name.Project.ID == projectID && name.Location == location {
				// Clone to avoid mutating stored object
				ret := proto.CloneOf(obj)
				// Strip input-only/secret fields
				if ret.GetGoogleCdnKey() != nil {
					ret.GetGoogleCdnKey().PrivateKey = nil
				}
				if ret.GetMediaCdnKey() != nil {
					ret.GetMediaCdnKey().PrivateKey = nil
				}
				if ret.GetAkamaiCdnKey() != nil {
					ret.GetAkamaiCdnKey().TokenKey = nil
				}
				filtered = append(filtered, ret)
			}
		}
		objs = filtered
	}

	return &pb.ListCdnKeysResponse{
		CdnKeys: objs,
	}, nil
}

type cdnKeyName struct {
	Project    *projects.ProjectData
	Location   string
	CdnKeyName string
}

func (n *cdnKeyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/cdnKeys/" + n.CdnKeyName
}

// parseCdnKeyName parses a string into a cdnKeyName.
// The expected form is projects/<projectID>/locations/<region>/cdnKeys/<cdnKeyName>
func (s *MockService) parseCdnKeyName(name string) (*cdnKeyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "cdnKeys" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &cdnKeyName{
			Project:    project,
			Location:   tokens[3],
			CdnKeyName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
