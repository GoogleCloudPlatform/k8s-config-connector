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
	"hash/crc32"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/secretmanager/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

// Creates a new [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] containing secret data and attaches
// it to an existing [Secret][google.cloud.secretmanager.v1.Secret].
func (s *SecretsV1) AddSecretVersion(ctx context.Context, req *pb.AddSecretVersionRequest) (*pb.SecretVersion, error) {
	secretName, err := s.parseSecretName(req.Parent)
	if err != nil {
		return nil, err
	}

	var secret pb.Secret

	if err := s.storage.Get(ctx, secretName.String(), &secret); err != nil {
		return nil, err
	}

	secretVersionKind := (&pb.SecretVersion{}).ProtoReflect().Descriptor()

	ids := make(map[int]bool)
	if err := s.storage.List(ctx, secretVersionKind, storage.ListOptions{
		Prefix: secretName.String() + "/",
	}, func(obj proto.Message) error {
		secretVersion := obj.(*pb.SecretVersion)
		id, err := strconv.Atoi(lastComponent(secretVersion.Name))
		if err != nil {
			return status.Errorf(codes.Internal, "error parsing Name %q: %v", secretVersion.Name, err)
		}
		ids[id] = true
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error reading secret versions: %v", err)
	}

	var version int
	for i := 1; ; i++ {
		if !ids[i] {
			version = i
			break
		}
	}

	secretVersionName := secretVersionName{
		Project:    secretName.Project,
		SecretName: secretName.SecretName,
		Version:    strconv.Itoa(version),
	}

	secretObj := &corev1.Secret{}
	secretKey := secretVersionName.kubernetesSecretID()
	secretObj.Name = secretKey.Name
	secretObj.Namespace = secretKey.Namespace
	if req.Payload == nil {
		req.Payload = &pb.SecretPayload{}
	}
	secretObj.Data = map[string][]byte{
		"data": req.Payload.Data,
	}

	secretVersionObj := &pb.SecretVersion{}
	secretVersionObj.Name = secretVersionName.String()
	secretVersionObj.CreateTime = timestamppb.Now()
	secretVersionObj.State = pb.SecretVersion_ENABLED
	secretVersionObj.Etag = computeEtag(secretVersionObj)

	// TODO: Copy from secret
	if secretVersionObj.ReplicationStatus == nil {
		secretVersionObj.ReplicationStatus = &pb.ReplicationStatus{}
	}
	if secretVersionObj.ReplicationStatus.ReplicationStatus == nil {
		secretVersionObj.ReplicationStatus.ReplicationStatus = &pb.ReplicationStatus_Automatic{
			Automatic: &pb.ReplicationStatus_AutomaticStatus{},
		}
	}

	// Ensure namespace exists
	// (Would be good to clean this up / align with project creation)
	{
		ns := &corev1.Namespace{}
		ns.SetName(secretKey.Namespace)
		if err := s.KubeClient.Create(ctx, ns); err != nil {
			if apierrors.IsAlreadyExists(err) {
				// somewhat expected
			} else {
				return nil, status.Errorf(codes.Internal, "error creating namespace: %v", err)
			}
		}
	}

	if err := s.KubeClient.Create(ctx, secretObj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating secret data: %v", err)
	}
	klog.Infof("created Secret %v", secretObj.GetNamespace()+"/"+secretObj.GetName())

	if err := s.storage.Create(ctx, secretVersionObj.Name, secretVersionObj); err != nil {
		// TODO: Delete secret data?
		// TODO: Owner ref?
		return nil, err
	}
	klog.Infof("created SecretManagerSecretVersion %v", secretVersionObj.Name)

	return secretVersionObj, nil
}

// Gets metadata for a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// `projects/*/secrets/*/versions/latest` is an alias to the most recently
// created [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
func (s *SecretsV1) GetSecretVersion(ctx context.Context, req *pb.GetSecretVersionRequest) (*pb.SecretVersion, error) {
	rawName, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	name, err := s.resolveLatestVersion(ctx, rawName)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	return secretVersion, nil
}

func (s *MockService) getSecretVersion(ctx context.Context, name *secretVersionName) (*pb.SecretVersion, error) {
	secretVersionObj := &pb.SecretVersion{}
	fqn := name.String()

	if err := s.storage.Get(ctx, fqn, secretVersionObj); err != nil {
		// TODO: Delete secret data?
		// TODO: Owner ref?
		return nil, err
	}
	return secretVersionObj, nil
}

func (s *MockService) accessSecret(ctx context.Context, secretVersion *pb.SecretVersion) (*corev1.Secret, error) {
	name, err := s.parseSecretVersionName(secretVersion.Name)
	if err != nil {
		return nil, err
	}

	key := name.kubernetesSecretID()
	secretObj := &corev1.Secret{}

	if err := s.KubeClient.Get(ctx, key, secretObj); err != nil {
		if apierrors.IsNotFound(err) {
			klog.Infof("did not find secret with id %v", key)
			return nil, status.Errorf(codes.NotFound, "secret version %q not found", name)
		}
		return nil, err
	}

	return secretObj, nil
}

// Accesses a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion]. This call returns the secret data.
//
// `projects/*/secrets/*/versions/latest` is an alias to the most recently
// created [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
func (s *SecretsV1) AccessSecretVersion(ctx context.Context, req *pb.AccessSecretVersionRequest) (*pb.AccessSecretVersionResponse, error) {
	rawName, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	name, err := s.resolveLatestVersion(ctx, rawName)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	secretData, err := s.accessSecret(ctx, secretVersion)
	if err != nil {
		return nil, err
	}

	response := &pb.AccessSecretVersionResponse{}
	response.Name = secretVersion.Name
	if data, ok := secretData.Data["data"]; ok {
		response.Payload = &pb.SecretPayload{
			Data: data,
		}

		crcTable := crc32.MakeTable(crc32.Castagnoli)
		dataCrc32c := int64(crc32.Checksum(response.Payload.Data, crcTable))
		response.Payload.DataCrc32C = &dataCrc32c
	}

	return response, nil
}

// Enables a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] to
// [ENABLED][google.cloud.secretmanager.v1.SecretVersion.State.ENABLED].
func (s *SecretsV1) EnableSecretVersion(ctx context.Context, req *pb.EnableSecretVersionRequest) (*pb.SecretVersion, error) {
	name, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	secretVersion.State = pb.SecretVersion_ENABLED
	fqn := secretVersion.Name
	if err := s.storage.Update(ctx, fqn, secretVersion); err != nil {
		return nil, err
	}

	return secretVersion, nil
}

func (s *SecretsV1) DisableSecretVersion(ctx context.Context, req *pb.DisableSecretVersionRequest) (*pb.SecretVersion, error) {
	name, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	secretVersion.State = pb.SecretVersion_DISABLED
	fqn := secretVersion.Name
	if err := s.storage.Update(ctx, fqn, secretVersion); err != nil {
		return nil, err
	}

	return secretVersion, nil
}

// Destroys a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] to
// [DESTROYED][google.cloud.secretmanager.v1.SecretVersion.State.DESTROYED] and irrevocably destroys the
// secret data.
func (s *SecretsV1) DestroySecretVersion(ctx context.Context, req *pb.DestroySecretVersionRequest) (*pb.SecretVersion, error) {
	// Note that the secret version still exists in the list
	name, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	now := timestamppb.Now()
	// TODO: Delete the kube secret

	secretVersion.State = pb.SecretVersion_DESTROYED
	secretVersion.DestroyTime = now

	fqn := secretVersion.Name
	if err := s.storage.Update(ctx, fqn, secretVersion); err != nil {
		return nil, err
	}

	return secretVersion, nil
}

func lastComponent(s string) string {
	i := strings.LastIndex(s, "/")
	return s[i+1:]
}

type secretVersionName struct {
	Project    *projects.ProjectData
	SecretName string
	Version    string
}

func (n *secretVersionName) String() string {
	return "projects/" + strconv.FormatInt(n.Project.Number, 10) + "/secrets/" + n.SecretName + "/versions/" + n.Version
}

func (s *MockService) parseSecretVersionName(name string) (*secretVersionName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "secrets" && tokens[4] == "versions" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &secretVersionName{
			Project:    project,
			SecretName: tokens[3],
			Version:    tokens[5],
		}

		if name.Version == "latest" {
			// OK
		} else {
			_, err := strconv.Atoi(name.Version)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid (invalid version)", name)
			}
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *secretVersionName) kubernetesSecretID() types.NamespacedName {
	key := types.NamespacedName{
		Name:      fmt.Sprintf("secretmanager-%s-%s", s.SecretName, s.Version),
		Namespace: strconv.FormatInt(s.Project.Number, 10),
	}
	return key
}

func (s *MockService) resolveLatestVersion(ctx context.Context, secretVersionName *secretVersionName) (*secretVersionName, error) {
	if secretVersionName.Version != "latest" {
		return secretVersionName, nil
	}

	maxVersion := 0
	secretVersionKind := (&pb.SecretVersion{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, secretVersionKind, storage.ListOptions{
		Prefix: secretVersionName.String() + "/",
	}, func(obj proto.Message) error {
		secretVersion := obj.(*pb.SecretVersion)
		v, err := strconv.Atoi(lastComponent(secretVersion.Name))
		if err != nil {
			return status.Errorf(codes.Internal, "error parsing Name %q: %v", secretVersion, err)
		}
		if v > maxVersion {
			maxVersion = v
		}
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error reading secret versions: %v", err)
	}

	if maxVersion == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "version %q not known", secretVersionName)
	}

	ret := *secretVersionName
	ret.Version = strconv.Itoa(maxVersion)
	return &ret, nil
}
