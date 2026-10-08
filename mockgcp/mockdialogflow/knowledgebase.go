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
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type knowledgeBasesServer struct {
	*MockService
	pb.UnimplementedKnowledgeBasesServer
}

type knowledgeBaseName struct {
	Project       string
	Location      string
	KnowledgeBase string
	Agent         bool
}

func (n *knowledgeBaseName) String() string {
	if n.Agent {
		return fmt.Sprintf("projects/%s/agent/knowledgeBases/%s", n.Project, n.KnowledgeBase)
	}
	if n.Location != "" {
		return fmt.Sprintf("projects/%s/locations/%s/knowledgeBases/%s", n.Project, n.Location, n.KnowledgeBase)
	}
	return fmt.Sprintf("projects/%s/knowledgeBases/%s", n.Project, n.KnowledgeBase)
}

func (s *MockService) parseKnowledgeBaseName(name string) (*knowledgeBaseName, error) {
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "knowledgeBases" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project:       project.ID,
			Location:      tokens[3],
			KnowledgeBase: tokens[5],
		}, nil
	}
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "knowledgeBases" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project:       project.ID,
			KnowledgeBase: tokens[3],
		}, nil
	}
	if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "agent" && tokens[3] == "knowledgeBases" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project:       project.ID,
			Agent:         true,
			KnowledgeBase: tokens[4],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid, expected format projects/{project}/locations/{location}/knowledgeBases/{knowledge_base}", name)
}

func (s *MockService) parseKnowledgeBaseParent(parent string) (*knowledgeBaseName, error) {
	if parent == "" {
		return nil, status.Errorf(codes.InvalidArgument, "parent must be provided")
	}

	tokens := strings.Split(parent, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project:  project.ID,
			Location: tokens[3],
		}, nil
	}
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project: project.ID,
		}, nil
	}
	if len(tokens) == 3 && tokens[0] == "projects" && tokens[2] == "agent" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &knowledgeBaseName{
			Project: project.ID,
			Agent:   true,
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

func (s *knowledgeBasesServer) GetKnowledgeBase(ctx context.Context, req *pb.GetKnowledgeBaseRequest) (*pb.KnowledgeBase, error) {
	name, err := s.parseKnowledgeBaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.KnowledgeBase{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			if name.Location != "" {
				return nil, status.Errorf(codes.NotFound, "Knowledge Base with ID %s not found in projects/%s/locations/%s.", name.KnowledgeBase, name.Project, name.Location)
			}
			if name.Agent {
				return nil, status.Errorf(codes.NotFound, "Knowledge Base with ID %s not found in projects/%s/agent.", name.KnowledgeBase, name.Project)
			}
			return nil, status.Errorf(codes.NotFound, "Knowledge Base with ID %s not found in projects/%s.", name.KnowledgeBase, name.Project)
		}
		return nil, err
	}

	return obj, nil
}

func (s *knowledgeBasesServer) CreateKnowledgeBase(ctx context.Context, req *pb.CreateKnowledgeBaseRequest) (*pb.KnowledgeBase, error) {
	parent, err := s.parseKnowledgeBaseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	kbID := ""
	if req.GetKnowledgeBase().GetName() != "" {
		name, err := s.parseKnowledgeBaseName(req.GetKnowledgeBase().GetName())
		if err == nil && name.KnowledgeBase != "" {
			kbID = name.KnowledgeBase
		}
	}
	if kbID == "" {
		n, err := rand.Int(rand.Reader, big.NewInt(1<<62))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
		kbID = base64.RawURLEncoding.EncodeToString([]byte(n.String()))
	}

	parent.KnowledgeBase = kbID
	fqn := parent.String()

	obj := proto.Clone(req.GetKnowledgeBase()).(*pb.KnowledgeBase)
	obj.Name = fqn
	if obj.GetLanguageCode() == "" {
		obj.LanguageCode = "en-US"
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *knowledgeBasesServer) UpdateKnowledgeBase(ctx context.Context, req *pb.UpdateKnowledgeBaseRequest) (*pb.KnowledgeBase, error) {
	name, err := s.parseKnowledgeBaseName(req.GetKnowledgeBase().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.KnowledgeBase{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	obj := req.GetKnowledgeBase()
	mask := req.GetUpdateMask()
	if mask == nil || len(mask.GetPaths()) == 0 {
		existing.DisplayName = obj.GetDisplayName()
		if obj.GetLanguageCode() != "" {
			existing.LanguageCode = obj.GetLanguageCode()
		}
	} else {
		for _, path := range mask.GetPaths() {
			switch path {
			case "display_name", "displayName":
				existing.DisplayName = obj.GetDisplayName()
			case "language_code", "languageCode":
				existing.LanguageCode = obj.GetLanguageCode()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "field %q cannot be updated", path)
			}
		}
	}

	if err := s.storage.Update(ctx, fqn, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *knowledgeBasesServer) DeleteKnowledgeBase(ctx context.Context, req *pb.DeleteKnowledgeBaseRequest) (*emptypb.Empty, error) {
	name, err := s.parseKnowledgeBaseName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deletedObj := &pb.KnowledgeBase{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *knowledgeBasesServer) ListKnowledgeBases(ctx context.Context, req *pb.ListKnowledgeBasesRequest) (*pb.ListKnowledgeBasesResponse, error) {
	parent, err := s.parseKnowledgeBaseParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/knowledgeBases/"
	var results []*pb.KnowledgeBase
	findOpts := storage.ListOptions{
		Prefix: prefix,
	}
	if err := s.storage.List(ctx, (&pb.KnowledgeBase{}).ProtoReflect().Descriptor(), findOpts, func(obj proto.Message) error {
		results = append(results, proto.Clone(obj).(*pb.KnowledgeBase))
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListKnowledgeBasesResponse{
		KnowledgeBases: results,
	}, nil
}
