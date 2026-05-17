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

package mockiam

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IAMServer) GetServiceAccountKey(ctx context.Context, req *adminpb.GetServiceAccountKeyRequest) (*adminpb.ServiceAccountKey, error) {
	name, err := s.parseServiceAccountKeyName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	obj := &adminpb.ServiceAccountKey{}
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Unknown service account key")
		}
		return nil, err
	}

	// Don't return key after initial creation
	ret := proto.Clone(obj).(*adminpb.ServiceAccountKey)
	ret.PrivateKeyData = nil
	ret.PrivateKeyType = adminpb.ServiceAccountPrivateKeyType_TYPE_UNSPECIFIED

	return ret, nil
}

func (s *IAMServer) ListServiceAccountKeys(ctx context.Context, req *adminpb.ListServiceAccountKeysRequest) (*adminpb.ListServiceAccountKeysResponse, error) {
	parent, err := s.parseServiceAccountName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	if len(req.KeyTypes) != 0 {
		return nil, fmt.Errorf("ListServiceAccountKeys key_types not implemented")
	}

	prefix := parent.String() + "/keys/"

	var keys []*adminpb.ServiceAccountKey
	keyKind := (&adminpb.ServiceAccountKey{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, keyKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		key, ok := obj.(*adminpb.ServiceAccountKey)
		if !ok {
			return status.Errorf(codes.Internal, "unexpected resource type: %T", obj)
		}

		// Don't return key after initial creation
		ret := proto.Clone(key).(*adminpb.ServiceAccountKey)
		ret.PrivateKeyData = nil
		ret.PrivateKeyType = adminpb.ServiceAccountPrivateKeyType_TYPE_UNSPECIFIED

		keys = append(keys, ret)
		return nil
	}); err != nil {
		return nil, err
	}

	return &adminpb.ListServiceAccountKeysResponse{
		Keys: keys,
	}, nil
}

func (s *IAMServer) CreateServiceAccountKey(ctx context.Context, req *adminpb.CreateServiceAccountKeyRequest) (*adminpb.ServiceAccountKey, error) {
	parent, err := s.parseServiceAccountName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	name := &serviceAccountKeyName{
		serviceAccountName: *parent,
	}

	now := time.Now()

	// TODO: Something more real
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", now.UnixNano())))
	name.Key = hex.EncodeToString(hash[:])

	key := &adminpb.ServiceAccountKey{}
	key.KeyAlgorithm = adminpb.ServiceAccountKeyAlgorithm_KEY_ALG_RSA_2048
	key.KeyOrigin = adminpb.ServiceAccountKeyOrigin_GOOGLE_PROVIDED
	key.KeyType = adminpb.ListServiceAccountKeysRequest_USER_MANAGED
	key.Name = name.String()
	key.PrivateKeyData = []byte("This should really be a key, but instead we just use a placeholder value.  I don't think it matters for our tests.")
	key.PrivateKeyType = adminpb.ServiceAccountPrivateKeyType_TYPE_GOOGLE_CREDENTIALS_FILE
	key.ValidBeforeTime = timestamppb.New(time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC))
	key.ValidAfterTime = timestamppb.New(now.Truncate(time.Second))

	fqn := name.String()
	if err := s.storage.Create(ctx, fqn, key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *IAMServer) DeleteServiceAccountKey(ctx context.Context, req *adminpb.DeleteServiceAccountKeyRequest) (*emptypb.Empty, error) {
	name, err := s.parseServiceAccountKeyName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deletedObj := &adminpb.ServiceAccountKey{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type serviceAccountKeyName struct {
	serviceAccountName
	Key string
}

func (n *serviceAccountKeyName) String() string {
	return n.serviceAccountName.String() + "/keys/" + n.Key
}

func (s *MockService) parseServiceAccountKeyName(ctx context.Context, name string) (*serviceAccountKeyName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[4] == "keys" {
		parentName := strings.Join(tokens[0:4], "/")
		parent, err := s.parseServiceAccountName(ctx, parentName)
		if err != nil {
			return nil, err
		}

		return &serviceAccountKeyName{
			serviceAccountName: *parent,
			Key:                tokens[5],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
