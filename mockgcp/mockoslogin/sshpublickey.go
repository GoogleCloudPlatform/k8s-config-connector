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

package mockoslogin

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	pb "cloud.google.com/go/oslogin/apiv1/osloginpb"
	commonpb "cloud.google.com/go/oslogin/common/commonpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type osLoginServer struct {
	*MockService
	pb.UnimplementedOsLoginServiceServer
}

func computeFingerprint(keyStr string) string {
	cleaned := strings.TrimSpace(keyStr)
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(cleaned))
	if err == nil {
		h := sha256.Sum256(pubKey.Marshal())
		return hex.EncodeToString(h[:])
	}
	h := sha256.Sum256([]byte(cleaned))
	return hex.EncodeToString(h[:])
}

func (s *osLoginServer) CreateSshPublicKey(ctx context.Context, req *pb.CreateSshPublicKeyRequest) (*commonpb.SshPublicKey, error) {
	if req.SshPublicKey == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ssh_public_key is required")
	}
	obj := proto.Clone(req.SshPublicKey).(*commonpb.SshPublicKey)
	if obj.Fingerprint == "" {
		obj.Fingerprint = computeFingerprint(obj.Key)
	}
	obj.Name = fmt.Sprintf("%s/sshPublicKeys/%s", req.Parent, obj.Fingerprint)

	if err := s.storage.Create(ctx, obj.Name, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *osLoginServer) ImportSshPublicKey(ctx context.Context, req *pb.ImportSshPublicKeyRequest) (*pb.ImportSshPublicKeyResponse, error) {
	if req.SshPublicKey == nil {
		return nil, status.Errorf(codes.InvalidArgument, "ssh_public_key is required")
	}
	obj := proto.Clone(req.SshPublicKey).(*commonpb.SshPublicKey)
	if obj.Fingerprint == "" {
		obj.Fingerprint = computeFingerprint(obj.Key)
	}
	obj.Name = fmt.Sprintf("%s/sshPublicKeys/%s", req.Parent, obj.Fingerprint)

	existing := &commonpb.SshPublicKey{}
	if err := s.storage.Get(ctx, obj.Name, existing); err == nil {
		if err := s.storage.Update(ctx, obj.Name, obj); err != nil {
			return nil, err
		}
	} else {
		if err := s.storage.Create(ctx, obj.Name, obj); err != nil {
			return nil, err
		}
	}

	sshPublicKeys := make(map[string]*commonpb.SshPublicKey)
	kind := (&commonpb.SshPublicKey{}).ProtoReflect().Descriptor()
	prefix := req.Parent + "/sshPublicKeys/"
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(o proto.Message) error {
		k := o.(*commonpb.SshPublicKey)
		if strings.HasPrefix(k.Name, prefix) {
			sshPublicKeys[k.Fingerprint] = k
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ImportSshPublicKeyResponse{
		LoginProfile: &pb.LoginProfile{
			Name:          req.Parent,
			SshPublicKeys: sshPublicKeys,
		},
	}, nil
}

func (s *osLoginServer) GetSshPublicKey(ctx context.Context, req *pb.GetSshPublicKeyRequest) (*commonpb.SshPublicKey, error) {
	obj := &commonpb.SshPublicKey{}
	if err := s.storage.Get(ctx, req.Name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SSH public key %q not found", req.Name)
		}
		return nil, err
	}
	return obj, nil
}

func (s *osLoginServer) UpdateSshPublicKey(ctx context.Context, req *pb.UpdateSshPublicKeyRequest) (*commonpb.SshPublicKey, error) {
	obj := &commonpb.SshPublicKey{}
	if err := s.storage.Get(ctx, req.Name, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SSH public key %q not found", req.Name)
		}
		return nil, err
	}

	if req.SshPublicKey != nil {
		if req.UpdateMask == nil || len(req.UpdateMask.Paths) == 0 {
			if req.SshPublicKey.ExpirationTimeUsec != 0 {
				obj.ExpirationTimeUsec = req.SshPublicKey.ExpirationTimeUsec
			}
			if req.SshPublicKey.Key != "" {
				obj.Key = req.SshPublicKey.Key
			}
		} else {
			for _, path := range req.UpdateMask.Paths {
				switch path {
				case "expiration_time_usec", "expirationTimeUsec":
					obj.ExpirationTimeUsec = req.SshPublicKey.ExpirationTimeUsec
				case "key":
					obj.Key = req.SshPublicKey.Key
				}
			}
		}
	}

	if err := s.storage.Update(ctx, req.Name, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *osLoginServer) DeleteSshPublicKey(ctx context.Context, req *pb.DeleteSshPublicKeyRequest) (*emptypb.Empty, error) {
	oldObj := &commonpb.SshPublicKey{}
	if err := s.storage.Delete(ctx, req.Name, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "SSH public key %q not found", req.Name)
		}
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *osLoginServer) GetLoginProfile(ctx context.Context, req *pb.GetLoginProfileRequest) (*pb.LoginProfile, error) {
	sshPublicKeys := make(map[string]*commonpb.SshPublicKey)
	kind := (&commonpb.SshPublicKey{}).ProtoReflect().Descriptor()
	prefix := req.Name + "/sshPublicKeys/"
	if err := s.storage.List(ctx, kind, storage.ListOptions{}, func(o proto.Message) error {
		k := o.(*commonpb.SshPublicKey)
		if strings.HasPrefix(k.Name, prefix) {
			sshPublicKeys[k.Fingerprint] = k
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.LoginProfile{
		Name:          req.Name,
		SshPublicKeys: sshPublicKeys,
	}, nil
}

func (s *osLoginServer) DeletePosixAccount(ctx context.Context, req *pb.DeletePosixAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
