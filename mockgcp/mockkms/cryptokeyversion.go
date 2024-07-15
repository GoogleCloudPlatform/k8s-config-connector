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

// +mockgcp-support
// apiVersion: kms.cnrm.cloud.google.com/v1beta1
// kind: KMSCryptoKeyVersion
// service: google.cloud.kms.v1.KeyManagementService
// resource: CryptoKeyVersion

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

	findPrefix := parentName.String()

	var matchingObjects []*pb.CryptoKeyVersion
	endpointKind := (&pb.CryptoKeyVersion{}).ProtoReflect().Descriptor()
	if err := r.storage.List(ctx, endpointKind, storage.ListOptions{}, func(obj proto.Message) error {
		cryptoKeyVersion := obj.(*pb.CryptoKeyVersion)
		if strings.HasPrefix(cryptoKeyVersion.Name, findPrefix) {
			matchingObjects = append(matchingObjects, cryptoKeyVersion)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListCryptoKeyVersionsResponse{
		CryptoKeyVersions: matchingObjects,
		NextPageToken:     "",
	}, nil
}

func (r *kmsServer) CreateCryptoKeyVersion(ctx context.Context, req *pb.CreateCryptoKeyVersionRequest) (*pb.CryptoKeyVersion, error) {
	parentName, err := r.parseCryptoKeyName(req.GetParent())
	if err != nil {
		return nil, err
	}

	id := strconv.FormatInt(time.Now().UnixNano(), 10)

	// The server-generated crypto key version name is the concatenation of the parent crypto key's
	// resource name and '/cryptoKeyVersions/' followed by a unique, server-assigned identifier.
	// For example, if the parent crypto key's resource name is
	// 'projects/1/locations/us-central1/keyRings/my-key-ring/cryptoKeys/my-crypto-key',
	// then the server-generated crypto key version name might be
	// 'projects/1/locations/us-central1/keyRings/my-key-ring/cryptoKeys/my-crypto-key/cryptoKeyVersions/123'
	name := &CryptoKeyVersionName{
		CryptoKeyName: parentName,
		Name:          id,
	}
	fqn := name.String()

	obj := proto.Clone(req.GetCryptoKeyVersion()).(*pb.CryptoKeyVersion)
	obj.Name = fqn

	r.populateDefaultsForCryptoKeyVersion(name, obj)

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

func (r *kmsServer) populateDefaultsForCryptoKeyVersion(name *CryptoKeyVersionName, obj *pb.CryptoKeyVersion) {

}

type CryptoKeyVersionName struct {
	*CryptoKeyName
	Name string
}

func (n *CryptoKeyVersionName) String() string {
	return fmt.Sprintf("%s/cryptoKeyVersions/%d", n.CryptoKeyName.String(), n.Name)
}

// parseCryptoKeyVersionName parses a string into a CryptoKeyVersionName.
// The expected form is `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
func (r *kmsServer) parseCryptoKeyVersionName(name string) (*CryptoKeyVersionName, error) {
	parts := strings.Split(name, "/")
	if len(parts) != 8 {
		return nil, status.Errorf(codes.InvalidArgument, "CryptoKeyVersion name must be in the form of projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*, got %v", name)
	}

	cryptoKeyName, err := r.parseCryptoKeyName(strings.Join(parts[0:6], "/"))
	if err != nil {
		return nil, err
	}

	id := parts[7]
	// TODO:  validate id is numeric

	return &CryptoKeyVersionName{
		CryptoKeyName: cryptoKeyName,
		Name:          id,
	}, nil
}
