// Copyright 2025 Google LLC
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
// proto.service: google.devtools.sourcerepo.v1.SourceRepo
// proto.message: google.devtools.sourcerepo.v1.Repo

package mocksourcerepo

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/sourcerepo/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *SourceRepoServer) GetRepo(ctx context.Context, req *pb.GetRepoRequest) (*pb.Repo, error) {
	name, err := s.parseRepoName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Repo{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Repo %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *SourceRepoServer) ListRepos(ctx context.Context, req *pb.ListReposRequest) (*pb.ListReposResponse, error) {
	name, err := s.parseRepoName(req.GetName() + "/repos/dummy")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(name.String(), "dummy")
	response := &pb.ListReposResponse{}

	repoKind := (&pb.Repo{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, repoKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		repo := obj.(*pb.Repo)
		// ListRepos does not set the size.
		repo.Size = 0
		response.Repos = append(response.Repos, repo)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SourceRepoServer) CreateRepo(ctx context.Context, req *pb.CreateRepoRequest) (*pb.Repo, error) {
	name, err := s.parseRepoName(req.GetRepo().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := proto.CloneOf(req.GetRepo())
	obj.Name = fqn
	obj.Url = fmt.Sprintf("https://source.developers.google.com/p/%s/r/%s", name.Project.ID, name.Repo)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SourceRepoServer) DeleteRepo(ctx context.Context, req *pb.DeleteRepoRequest) (*emptypb.Empty, error) {
	name, err := s.parseRepoName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Repo{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *SourceRepoServer) UpdateRepo(ctx context.Context, req *pb.UpdateRepoRequest) (*pb.Repo, error) {
	name, err := s.parseRepoName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Repo{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	if updateMask != nil {
		if err := fields.UpdateByFieldMask(obj, req.GetRepo(), updateMask.GetPaths()); err != nil {
			return nil, err
		}
	} else {
		// Full update if no mask
		obj = req.GetRepo()
		obj.Name = fqn
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SourceRepoServer) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	s.MockService.Lock()
	defer s.MockService.Unlock()

	policy := s.iamPolicies[req.Resource]
	if policy == nil {
		policy = &iampb.Policy{}
		policy.Etag = s.computeEtag(policy)
	}
	return policy, nil
}

func (s *SourceRepoServer) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	s.MockService.Lock()
	defer s.MockService.Unlock()

	s.iamPolicies[req.Resource] = req.Policy
	req.Policy.Etag = s.computeEtag(req.Policy)
	return req.Policy, nil
}

func (s *SourceRepoServer) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	// For mock, just allow all requested permissions
	return &iampb.TestIamPermissionsResponse{
		Permissions: req.Permissions,
	}, nil
}

func (s *SourceRepoServer) computeEtag(policy *iampb.Policy) []byte {
	b, err := proto.Marshal(policy)
	if err != nil {
		panic(fmt.Errorf("error marshaling policy: %w", err))
	}
	hash := md5.Sum(b)
	return hash[:]
}

type repoName struct {
	Project *projects.ProjectData
	Repo    string
}

func (n *repoName) String() string {
	return fmt.Sprintf("projects/%s/repos/%s", n.Project.ID, n.Repo)
}

// parseRepoName parses a string into a repoName.
// The expected form is `projects/{project}/repos/{repo}` where repo can contain slashes.
func (s *MockService) parseRepoName(name string) (*repoName, error) {
	tokens := strings.SplitN(name, "/", 4)
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "repos" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &repoName{
			Project: project,
			Repo:    tokens[3],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
