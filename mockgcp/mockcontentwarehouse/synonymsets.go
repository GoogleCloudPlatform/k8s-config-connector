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

package mockcontentwarehouse

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/contentwarehouse/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type SynonymSetName struct {
	Project  *projects.ProjectName
	Location string
	Context  string
}

func (n *SynonymSetName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/synonymSets/%s", n.Project.OriginalValue, n.Location, n.Context)
}

func (s *SynonymSetV1) parseSynonymSetName(name string) (*SynonymSetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "synonymSets" {
		projectName, err := projects.ParseProjectIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		return &SynonymSetName{
			Project:  projectName,
			Location: tokens[3],
			Context:  tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *SynonymSetV1) canonicalizeName(name string) (string, error) {
	parsed, err := s.parseSynonymSetName(name)
	if err != nil {
		return "", err
	}
	proj, err := s.Projects.GetProjectByIDOrNumber(parsed.Project.OriginalValue)
	if err == nil {
		// Canonical format uses project number
		return fmt.Sprintf("projects/%d/locations/%s/synonymSets/%s", proj.Number, parsed.Location, parsed.Context), nil
	}
	// Fallback to parsed if not found
	return parsed.String(), nil
}

func (s *SynonymSetV1) CreateSynonymSet(ctx context.Context, req *pb.CreateSynonymSetRequest) (*pb.SynonymSet, error) {
	parent := req.GetParent() // projects/{project_number}/locations/{location}
	synonymSet := req.GetSynonymSet()

	if synonymSet == nil {
		return nil, status.Errorf(codes.InvalidArgument, "synonym_set is required")
	}

	if synonymSet.GetContext() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "synonym_set.context is required")
	}

	// Construct fully qualified name
	fqn := fmt.Sprintf("%s/synonymSets/%s", parent, synonymSet.GetContext())
	canonicalFQN, err := s.canonicalizeName(fqn)
	if err != nil {
		return nil, err
	}

	// Check if already exists
	existing := &pb.SynonymSet{}
	if err := s.storage.Get(ctx, canonicalFQN, existing); err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "synonym set %q already exists", canonicalFQN)
	}

	synonymSet.Name = canonicalFQN

	if err := s.storage.Create(ctx, canonicalFQN, synonymSet); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating synonym set: %v", err)
	}

	return synonymSet, nil
}

func (s *SynonymSetV1) GetSynonymSet(ctx context.Context, req *pb.GetSynonymSetRequest) (*pb.SynonymSet, error) {
	canonicalFQN, err := s.canonicalizeName(req.GetName())
	if err != nil {
		return nil, err
	}

	obj := &pb.SynonymSet{}
	if err := s.storage.Get(ctx, canonicalFQN, obj); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	return obj, nil
}

func (s *SynonymSetV1) UpdateSynonymSet(ctx context.Context, req *pb.UpdateSynonymSetRequest) (*pb.SynonymSet, error) {
	canonicalFQN, err := s.canonicalizeName(req.GetName())
	if err != nil {
		return nil, err
	}

	synonymSet := req.GetSynonymSet()
	if synonymSet == nil {
		return nil, status.Errorf(codes.InvalidArgument, "synonym_set is required")
	}

	// Check if already exists (Update throws NOT_FOUND if it is not found)
	existing := &pb.SynonymSet{}
	if err := s.storage.Get(ctx, canonicalFQN, existing); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	synonymSet.Name = canonicalFQN
	if synonymSet.GetContext() == "" {
		synonymSet.Context = existing.GetContext()
	}

	if err := s.storage.Update(ctx, canonicalFQN, synonymSet); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating synonym set: %v", err)
	}

	return synonymSet, nil
}

func (s *SynonymSetV1) DeleteSynonymSet(ctx context.Context, req *pb.DeleteSynonymSetRequest) (*emptypb.Empty, error) {
	canonicalFQN, err := s.canonicalizeName(req.GetName())
	if err != nil {
		return nil, err
	}

	deletedObject := &pb.SynonymSet{}
	if err := s.storage.Delete(ctx, canonicalFQN, deletedObject); err != nil {
		return nil, s.mapNotFoundError(err, canonicalFQN)
	}

	return &emptypb.Empty{}, nil
}

func (s *SynonymSetV1) mapNotFoundError(err error, name string) error {
	if status.Code(err) == codes.NotFound {
		formattedName := name
		if strings.HasPrefix(name, "projects/") {
			formattedName = "Projects/" + strings.TrimPrefix(name, "projects/")
		}
		return status.Errorf(codes.NotFound, "SynonymSet %s not found.", formattedName)
	}
	return err
}

func (s *SynonymSetV1) ListSynonymSets(ctx context.Context, req *pb.ListSynonymSetsRequest) (*pb.ListSynonymSetsResponse, error) {
	parent := req.GetParent()
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be specified")
	}

	var synonymSets []*pb.SynonymSet
	if err := s.storage.List(ctx, (&pb.SynonymSet{}).ProtoReflect().Descriptor(), storage.ListOptions{
		Prefix: parent + "/synonymSets/",
	}, func(obj proto.Message) error {
		synonymSet := obj.(*pb.SynonymSet)
		synonymSets = append(synonymSets, synonymSet)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListSynonymSetsResponse{
		SynonymSets: synonymSets,
	}, nil
}
