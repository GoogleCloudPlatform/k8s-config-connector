// Copyright 2024 Google LLC
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
// krm.apiVersion: kms.cnrm.cloud.google.com/v1beta1
// krm.kind: KMSCryptoKeyVersion
// proto.service: google.cloud.kms.v1.KeyManagementService
// proto.resource: CryptoKeyVersion

package mockkms

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/kms/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (r *kmsServer) GetCryptoKeyVersion(ctx context.Context, req *pb.GetCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {
	name, err := r.parseCryptoKeyVersionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CryptoKeyVersion{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "CryptoKeyVersion %s not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) ListCryptoKeyVersions(ctx context.Context, req *pb.ListCryptoKeyVersionsRequest) (*pb.ListCryptoKeyVersionsResponse, error) {
	parentName, err := r.parseCryptoKeyName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentFQN := parentName.String()

	response, err := r.listCryptoKeyVersions(ctx, parentFQN)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *kmsServer) listCryptoKeyVersions(ctx context.Context, parentFQN string) (*pb.ListCryptoKeyVersionsResponse, error) {
	namePrefix := parentFQN + "/cryptoKeyVersions/"

	response := &pb.ListCryptoKeyVersionsResponse{}

	// Network must not have any subnets depending on it
	cryptoKeyVersionKind := (&pb.CryptoKeyVersion{}).ProtoReflect().Descriptor()
	if err := r.storage.List(ctx, cryptoKeyVersionKind, storage.ListOptions{}, func(obj proto.Message) error {
		cryptoKeyVersion := obj.(*pb.CryptoKeyVersion)
		if strings.HasPrefix(cryptoKeyVersion.GetName(), namePrefix) {
			response.CryptoKeyVersions = append(response.CryptoKeyVersions, cryptoKeyVersion)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	response.TotalSize = int32(len(response.CryptoKeyVersions))

	return response, nil
}

func (r *kmsServer) CreateCryptoKeyVersion(ctx context.Context, req *pb.CreateCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {

	versions, err := r.listCryptoKeyVersions(ctx, req.GetParent())
	if err != nil {
		return nil, err
	}

	var maxVersion int64
	for _, version := range versions.CryptoKeyVersions {
		last := lastComponent(version.GetName())
		n, err := strconv.ParseInt(last, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid key version name %q", version.GetName())
		}
		if maxVersion < n {
			maxVersion = n
		}
	}

	nextVersion := maxVersion + 1

	reqName := fmt.Sprintf("%s/cryptoKeyVersions/%d", req.GetParent(), nextVersion)
	name, err := r.parseCryptoKeyVersionName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	var obj *pb.CryptoKeyVersion
	obj = proto.Clone(req.GetCryptoKeyVersion()).(*pb.CryptoKeyVersion)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.GenerateTime = timestamppb.New(now)
	obj.State = pb.CryptoKeyVersion_ENABLED
	obj.Algorithm = req.CryptoKeyVersion.GetAlgorithm()
	obj.ProtectionLevel = req.CryptoKeyVersion.GetProtectionLevel()

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) UpdateCryptoKeyVersion(ctx context.Context, req *pb.UpdateCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {
	name, err := r.parseCryptoKeyVersionName(req.GetCryptoKeyVersion().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.CryptoKeyVersion{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	proto.Merge(obj, req.GetCryptoKeyVersion())
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (r *kmsServer) DestroyCryptoKeyVersion(ctx context.Context, req *pb.DestroyCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {
	name, err := r.parseCryptoKeyVersionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.CryptoKeyVersion{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	var parent *pb.CryptoKey
	{
		get := &pb.GetCryptoKeyRequest{
			Name: name.CryptoKeyName.String(),
		}
		cryptoKey, err := r.GetCryptoKey(ctx, get)
		if err != nil {
			return nil, err
		}
		parent = cryptoKey
	}

	destroyScheuledDuration := parent.GetDestroyScheduledDuration().AsDuration()

	obj.State = pb.CryptoKeyVersion_DESTROY_SCHEDULED
	obj.DestroyTime = timestamppb.New(now.Add(destroyScheuledDuration))

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *kmsServer) RestoreCryptoKeyVersion(ctx context.Context, req *pb.RestoreCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {
	name, err := r.parseCryptoKeyVersionName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.CryptoKeyVersion{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	// TODO:  set appropriate state and fields
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type CryptoKeyVersionName struct {
	CryptoKeyName
	CryptoKeyVersionID string
}

func (n *CryptoKeyVersionName) String() string {
	return n.CryptoKeyName.String() + "/cryptoKeyVersions/" + n.CryptoKeyVersionID
}

// parseCryptoKeyVersionName parses a string into a CryptoKeyVersionName.
// The expected form is `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
func (r *kmsServer) parseCryptoKeyVersionName(name string) (*CryptoKeyVersionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[8] == "cryptoKeyVersions" {
		cryptoKeyName, err := r.parseCryptoKeyName(strings.Join(tokens[0:8], "/"))
		if err != nil {
			return nil, err
		}

		// TODO:  validate id is numeric
		id := tokens[9]
		name := &CryptoKeyVersionName{
			CryptoKeyName:      *cryptoKeyName,
			CryptoKeyVersionID: id,
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
