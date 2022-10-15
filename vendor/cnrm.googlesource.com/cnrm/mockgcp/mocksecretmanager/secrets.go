// Copyright 2022 Google LLC
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

package mocksecretmanager

import (
	"context"
	"strconv"
	"strings"

	secretmanager "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// Lists [Secrets][google.cloud.secretmanager.v1.Secret].
func (s *MockService) ListSecrets(context.Context, *secretmanager.ListSecretsRequest) (*secretmanager.ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

// Creates a new [Secret][google.cloud.secretmanager.v1.Secret] containing no [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
func (s *MockService) CreateSecret(ctx context.Context, req *secretmanager.CreateSecretRequest) (*secretmanager.Secret, error) {
	secretID := req.SecretId
	if secretID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SecretId is required")
	}

	parent, err := parseProjectName(req.Parent)
	if err != nil {
		return nil, err
	}

	project := s.projects.getProject(parent.Project)
	if project == nil {
		return nil, status.Errorf(codes.InvalidArgument, "project %q not known", parent)
	}

	name := secretName{
		ProjectNumber: project.Number,
		SecretName:    secretID,
	}
	fqn := name.String()

	obj := &secretmanager.Secret{}
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating secret: %v", err)
	}

	return obj, nil
}

// Gets metadata for a given [Secret][google.cloud.secretmanager.v1.Secret].
func (s *MockService) GetSecret(ctx context.Context, req *secretmanager.GetSecretRequest) (*secretmanager.Secret, error) {
	name, err := s.parseSecretName(req.Name)
	if err != nil {
		return nil, err
	}

	project := s.projects.getProjectByNumber(name.ProjectNumber)
	if project == nil {
		return nil, status.Errorf(codes.InvalidArgument, "project %q not known", name.ProjectNumber)
	}

	var secret secretmanager.Secret
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &secret); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "secret %q not found", req.Name)

		}
		return nil, status.Errorf(codes.Internal, "error reading secret: %v", err)
	}

	return &secret, nil
}

// Updates metadata of an existing [Secret][google.cloud.secretmanager.v1.Secret].
func (s *MockService) UpdateSecret(context.Context, *secretmanager.UpdateSecretRequest) (*secretmanager.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

// Deletes a [Secret][google.cloud.secretmanager.v1.Secret].
func (s *MockService) DeleteSecret(context.Context, *secretmanager.DeleteSecretRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

type secretName struct {
	ProjectNumber int64
	SecretName    string
}

func (n *secretName) String() string {
	return "projects/" + strconv.FormatInt(n.ProjectNumber, 10) + "/secrets/" + n.SecretName
}

// parseSecretName parses a string into a secretName.
// The expected form is projects/<projectNumber>/secrets/<secretName>
func (s *MockService) parseSecretName(name string) (*secretName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "secrets" {
		project := s.projects.getProject(tokens[1])
		if project == nil {
			return nil, status.Errorf(codes.InvalidArgument, "project %q not known", tokens[1])
		}

		name := &secretName{
			ProjectNumber: project.Number,
			SecretName:    tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
