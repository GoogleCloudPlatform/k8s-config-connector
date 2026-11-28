// Copyright 2026 High Five Labs
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

// Package tags, as the
// Artifact Registry API documents them. A tag must be in an existing
// repository; its package and version are not checked, since the mock holds
// no packages (nothing is ever pushed to it).

package mockartifactregistry

import (
	"context"
	"strings"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type tagName struct {
	repository *artifactRegistryName
	Package    string
	Tag        string
}

func (n *tagName) String() string {
	return n.repository.String() + "/packages/" + n.Package + "/tags/" + n.Tag
}

func (s *ArtifactRegistryV1) parseTagName(ctx context.Context, name string) (*tagName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 10 || tokens[6] != "packages" || tokens[8] != "tags" {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
	repo, err := s.parseArtifactRegistryName(strings.Join(tokens[:6], "/"))
	if err != nil {
		return nil, err
	}
	if err := s.storage.Get(ctx, repo.String(), &pb.Repository{}); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Repository %q not found.", repo.String())
		}
		return nil, err
	}
	return &tagName{repository: repo, Package: tokens[7], Tag: tokens[9]}, nil
}

func (s *ArtifactRegistryV1) GetTag(ctx context.Context, req *pb.GetTagRequest) (*pb.Tag, error) {
	n, err := s.parseTagName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.Tag{}
	if err := s.storage.Get(ctx, n.String(), obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}
	return obj, nil
}

func (s *ArtifactRegistryV1) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.Tag, error) {
	n, err := s.parseTagName(ctx, req.GetParent()+"/tags/"+req.GetTagId())
	if err != nil {
		return nil, err
	}
	if req.GetTag().GetVersion() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "tag must name a version")
	}
	obj := proto.Clone(req.GetTag()).(*pb.Tag)
	obj.Name = n.String()
	if err := s.storage.Create(ctx, n.String(), obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ArtifactRegistryV1) UpdateTag(ctx context.Context, req *pb.UpdateTagRequest) (*pb.Tag, error) {
	n, err := s.parseTagName(ctx, req.GetTag().GetName())
	if err != nil {
		return nil, err
	}
	obj := &pb.Tag{}
	if err := s.storage.Get(ctx, n.String(), obj); err != nil {
		return nil, err
	}
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "version":
			obj.Version = req.GetTag().GetVersion()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}
	if err := s.storage.Update(ctx, n.String(), obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *ArtifactRegistryV1) DeleteTag(ctx context.Context, req *pb.DeleteTagRequest) (*emptypb.Empty, error) {
	n, err := s.parseTagName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, n.String(), &pb.Tag{}); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
