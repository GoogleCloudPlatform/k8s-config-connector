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
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/secretmanager/v1"
)

type SecretsV1 struct {
	*MockService

	pb.UnimplementedSecretManagerServiceServer
}

// Creates a new [Secret][google.cloud.secretmanager.v1.Secret] containing no [SecretVersions][google.cloud.secretmanager.v1.SecretVersion].
func (s *SecretsV1) CreateSecret(ctx context.Context, req *pb.CreateSecretRequest) (*pb.Secret, error) {
	secretID := req.SecretId
	if secretID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "SecretId is required")
	}

	parent, err := projects.ParseProjectName(req.Parent)
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProject(parent)
	if err != nil {
		return nil, err
	}

	name := secretName{
		Project:    project,
		SecretName: secretID,
	}
	fqn := name.String()

	obj := proto.Clone(req.Secret).(*pb.Secret)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	if obj.Replication == nil {
		return nil, fmt.Errorf("Secret.replication must be specified.")
	}
	obj.Etag = computeEtag(obj)
	if err := s.populateDefaultsForSecret(ctx, obj); err != nil {
		return nil, err
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *SecretsV1) populateDefaultsForSecret(ctx context.Context, obj *pb.Secret) error {
	for _, version := range obj.VersionAliases {
		versionName := obj.Name + "/versions/" + strconv.FormatInt(version, 10)
		_, err := s.GetSecretVersion(ctx, &pb.GetSecretVersionRequest{Name: versionName})
		if err != nil {
			return fmt.Errorf("Aliases cannot be assigned to versions that don't exist")
		}
	}
	// TTL and ExpireTime are OneOf, but the GCP service always converts TTL to expireTime before storing the object.
	if obj.GetTtl() != nil {
		expirateTime := timestamppb.Now().AsTime().Add(obj.GetTtl().AsDuration())
		obj.Expiration = &pb.Secret_ExpireTime{
			ExpireTime: timestamppb.New(expirateTime),
		}
	}
	return nil
}

// Gets metadata for a given [Secret][google.cloud.secretmanager.v1.Secret].
func (s *SecretsV1) GetSecret(ctx context.Context, req *pb.GetSecretRequest) (*pb.Secret, error) {
	name, err := s.parseSecretName(req.Name)
	if err != nil {
		return nil, err
	}

	var secret pb.Secret
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, &secret); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Secret [%s] not found.", fqn)
		}
		return nil, err
	}

	return &secret, nil
}

// ProtoClone is a type-safe wrapper around proto.Clone
func ProtoClone[T proto.Message](t T) T {
	return proto.Clone(t).(T)
}

// Update metadata for a given [Secret][google.cloud.secretmanager.v1.Secret].
func (s *SecretsV1) UpdateSecret(ctx context.Context, req *pb.UpdateSecretRequest) (*pb.Secret, error) {
	name, err := s.parseSecretName(req.Secret.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	existing := &pb.Secret{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)
	updated.Name = name.String()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}
	for _, path := range paths {
		switch path {
		case "topics":
			updated.Topics = req.Secret.GetTopics()
		case "customerManagedEncryption":
			updated.CustomerManagedEncryption = req.Secret.GetCustomerManagedEncryption()
		case "rotation":
			updated.Rotation = req.Secret.GetRotation()
			if len(req.Secret.GetTopics()) == 0 {
				return nil, fmt.Errorf("There must be at least one topic configured when a Rotation policy is set.")
			}
		case "rotation.rotationPeriod":
			updated.Rotation.RotationPeriod = req.Secret.GetRotation().RotationPeriod
		case "annotations":
			updated.Annotations = req.Secret.GetAnnotations()
		case "labels":
			updated.Labels = req.Secret.GetLabels()
		case "versionAliases":
			updated.VersionAliases = req.Secret.GetVersionAliases()
		case "expireTime":
			updated.Expiration = &pb.Secret_ExpireTime{
				ExpireTime: req.Secret.GetExpireTime(),
			}
		case "ttl":
			updated.Expiration = &pb.Secret_Ttl{
				Ttl: req.Secret.GetTtl(),
			}
		case "expiration":
			updated.Expiration = req.Secret.GetExpiration()
		case "rotation.nextRotationTime":
			updated.Rotation.NextRotationTime = req.Secret.Rotation.NextRotationTime
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.populateDefaultsForSecret(ctx, updated); err != nil {
		return nil, err
	}
	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

// Deletes a [Secret][google.cloud.secretmanager.v1.Secret].
func (s *SecretsV1) DeleteSecret(ctx context.Context, req *pb.DeleteSecretRequest) (*emptypb.Empty, error) {
	name, err := s.parseSecretName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Secret{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	// TODO: Delete secret versions?

	return &emptypb.Empty{}, nil
}

type secretName struct {
	Project    *projects.ProjectData
	SecretName string
}

func (n *secretName) String() string {
	return "projects/" + strconv.FormatInt(n.Project.Number, 10) + "/secrets/" + n.SecretName
}

// parseSecretName parses a string into a secretName.
// The expected form is projects/<projectID>/secrets/<secretName>
func (s *MockService) parseSecretName(name string) (*secretName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "secrets" {
		projectName, err := projects.ParseProjectName("projects/" + tokens[1])
		if err != nil {
			return nil, err
		}

		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &secretName{
			Project:    project,
			SecretName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
