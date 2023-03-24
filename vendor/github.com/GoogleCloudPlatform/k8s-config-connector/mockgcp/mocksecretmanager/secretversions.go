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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	secretmanager "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

// Creates a new [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] containing secret data and attaches
// it to an existing [Secret][google.cloud.secretmanager.v1.Secret].
func (s *MockService) AddSecretVersion(ctx context.Context, req *secretmanager.AddSecretVersionRequest) (*secretmanager.SecretVersion, error) {
	secretName, err := s.parseSecretName(req.Parent)
	if err != nil {
		return nil, err
	}

	var secret secretmanager.Secret

	if err := s.storage.Get(ctx, secretName.String(), &secret); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "secret %q not found", req.Parent)

		}
		return nil, status.Errorf(codes.Internal, "error reading secret: %v", err)
	}

	secretVersionKind := (&secretmanager.SecretVersion{}).ProtoReflect().Descriptor()

	ids := make(map[int]bool)
	if err := s.storage.List(ctx, secretVersionKind, storage.ListOptions{
		Prefix: secretName.String() + "/",
	}, func(obj proto.Message) error {
		secretVersion := obj.(*secretmanager.SecretVersion)
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
		req.Payload = &secretmanager.SecretPayload{}
	}
	secretObj.Data = map[string][]byte{
		"data": req.Payload.Data,
	}

	secretVersionObj := &secretmanager.SecretVersion{}
	secretVersionObj.Name = secretVersionName.String()
	secretVersionObj.CreateTime = timestamppb.Now()
	secretVersionObj.State = secretmanager.SecretVersion_ENABLED

	if err := s.kube.Create(ctx, secretObj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating secret data: %v", err)
	}
	klog.Infof("created Secret %v", secretObj.GetNamespace()+"/"+secretObj.GetName())

	if err := s.storage.Create(ctx, secretVersionObj.Name, secretVersionObj); err != nil {
		// TODO: Delete secret data?
		// TODO: Owner ref?
		return nil, status.Errorf(codes.Internal, "error creating secret version: %v", err)
	}
	klog.Infof("created SecretManagerSecretVersion %v", secretVersionObj.Name)

	return secretVersionObj, nil
}

// Lists [SecretVersions][google.cloud.secretmanager.v1.SecretVersion]. This call does not return secret
// data.
func (s *MockService) ListSecretVersions(context.Context, *secretmanager.ListSecretVersionsRequest) (*secretmanager.ListSecretVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "ListSecretVersions not implemented")
}

// Gets metadata for a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// `projects/*/secrets/*/versions/latest` is an alias to the most recently
// created [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
func (s *MockService) GetSecretVersion(ctx context.Context, req *secretmanager.GetSecretVersionRequest) (*secretmanager.SecretVersion, error) {
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

func (s *MockService) getSecretVersion(ctx context.Context, name *secretVersionName) (*secretmanager.SecretVersion, error) {
	secretVersionObj := &secretmanager.SecretVersion{}
	fqn := name.String()

	if err := s.storage.Get(ctx, fqn, secretVersionObj); err != nil {
		// TODO: Delete secret data?
		// TODO: Owner ref?
		return nil, status.Errorf(codes.Internal, "error creating secret version: %v", err)
	}
	return secretVersionObj, nil
}

func (s *MockService) accessSecret(ctx context.Context, secretVersion *secretmanager.SecretVersion) (*corev1.Secret, error) {
	name, err := s.parseSecretVersionName(secretVersion.Name)
	if err != nil {
		return nil, err
	}

	key := name.kubernetesSecretID()
	secretObj := &corev1.Secret{}

	if err := s.kube.Get(ctx, key, secretObj); err != nil {
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
func (s *MockService) AccessSecretVersion(ctx context.Context, req *secretmanager.AccessSecretVersionRequest) (*secretmanager.AccessSecretVersionResponse, error) {
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

	response := &secretmanager.AccessSecretVersionResponse{}
	response.Name = secretVersion.Name
	if data, ok := secretData.Data["data"]; ok {
		response.Payload = &secretmanager.SecretPayload{
			Data: data,
		}
	}

	return response, nil
}

// Disables a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] to
// [DISABLED][google.cloud.secretmanager.v1.SecretVersion.State.DISABLED].
func (s *MockService) DisableSecretVersion(ctx context.Context, req *secretmanager.DisableSecretVersionRequest) (*secretmanager.SecretVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "DisableSecretVersion not implemented")
}

// Enables a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] to
// [ENABLED][google.cloud.secretmanager.v1.SecretVersion.State.ENABLED].
func (s *MockService) EnableSecretVersion(ctx context.Context, req *secretmanager.EnableSecretVersionRequest) (*secretmanager.SecretVersion, error) {
	name, err := s.parseSecretVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	secretVersion, err := s.getSecretVersion(ctx, name)
	if err != nil {
		return nil, err
	}

	secretVersion.State = secretmanager.SecretVersion_ENABLED
	fqn := secretVersion.Name
	if err := s.storage.Update(ctx, fqn, secretVersion); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating secret version: %v", err)
	}

	return secretVersion, nil
}

// Destroys a [SecretVersion][google.cloud.secretmanager.v1.SecretVersion].
//
// Sets the [state][google.cloud.secretmanager.v1.SecretVersion.state] of the [SecretVersion][google.cloud.secretmanager.v1.SecretVersion] to
// [DESTROYED][google.cloud.secretmanager.v1.SecretVersion.State.DESTROYED] and irrevocably destroys the
// secret data.
func (s *MockService) DestroySecretVersion(context.Context, *secretmanager.DestroySecretVersionRequest) (*secretmanager.SecretVersion, error) {
	// Note that the secret version still exists in the list
	return nil, status.Errorf(codes.Unimplemented, "DestroySecretVersion not implemented")
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
		project, err := s.projects.GetProjectByNumber(tokens[1])
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
	secretVersionKind := (&secretmanager.SecretVersion{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, secretVersionKind, storage.ListOptions{
		Prefix: secretVersionName.String() + "/",
	}, func(obj proto.Message) error {
		secretVersion := obj.(*secretmanager.SecretVersion)
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
