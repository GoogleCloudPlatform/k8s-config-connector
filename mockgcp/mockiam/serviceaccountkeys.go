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

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/admin/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *IAMServer) GetServiceAccountKey(ctx context.Context, req *pb.GetServiceAccountKeyRequest) (*pb.ServiceAccountKey, error) {
	name, err := s.parseServiceAccountKeyName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	obj := &pb.ServiceAccountKey{}
	fqn := name.String()
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Unknown service account key")
		}
		return nil, err
	}

	// Don't return key after initial creation
	ret := ProtoClone(obj)
	ret.PrivateKeyData = nil
	ret.PrivateKeyType = pb.ServiceAccountPrivateKeyType_TYPE_UNSPECIFIED

	return ret, nil
}

func (s *IAMServer) ListServiceAccountKeys(ctx context.Context, req *pb.ListServiceAccountKeysRequest) (*pb.ListServiceAccountKeysResponse, error) {
	parent, err := s.parseServiceAccountName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	if len(req.KeyTypes) != 0 {
		return nil, fmt.Errorf("ListServiceAccountKeys key_types not implemented")
	}

	prefix := parent.String() + "/keys/"

	var keys []*pb.ServiceAccountKey
	keyKind := (&pb.ServiceAccountKey{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, keyKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		key, ok := obj.(*pb.ServiceAccountKey)
		if !ok {
			return status.Errorf(codes.Internal, "unexpected resource type: %T", obj)
		}

		// Don't return key after initial creation
		ret := ProtoClone(key)
		ret.PrivateKeyData = nil
		ret.PrivateKeyType = pb.ServiceAccountPrivateKeyType_TYPE_UNSPECIFIED

		keys = append(keys, ret)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListServiceAccountKeysResponse{
		Keys: keys,
	}, nil
}

func (s *IAMServer) CreateServiceAccountKey(ctx context.Context, req *pb.CreateServiceAccountKeyRequest) (*pb.ServiceAccountKey, error) {
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

	key := &pb.ServiceAccountKey{}
	key.KeyAlgorithm = pb.ServiceAccountKeyAlgorithm_KEY_ALG_RSA_2048
	key.KeyOrigin = pb.ServiceAccountKeyOrigin_GOOGLE_PROVIDED
	key.KeyType = pb.ListServiceAccountKeysRequest_USER_MANAGED
	key.Name = name.String()
	key.PrivateKeyData = []byte("This should really be a key, but instead we just use a placeholder value.  I don't think it matters for our tests.")
	key.PrivateKeyType = pb.ServiceAccountPrivateKeyType_TYPE_GOOGLE_CREDENTIALS_FILE
	key.ValidBeforeTime = timestamppb.New(time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC))
	key.ValidAfterTime = timestamppb.New(now.Truncate(time.Second))

	fqn := name.String()
	if err := s.storage.Create(ctx, fqn, key); err != nil {
		return nil, err
	}

	return key, nil
}

func (s *IAMServer) DeleteServiceAccountKey(ctx context.Context, req *pb.DeleteServiceAccountKeyRequest) (*empty.Empty, error) {
	name, err := s.parseServiceAccountKeyName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	deletedObj := &pb.ServiceAccountKey{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
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
