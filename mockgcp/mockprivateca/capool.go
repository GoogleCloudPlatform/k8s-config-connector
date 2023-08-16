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

package mockprivateca

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
)

type PrivateCAV1 struct {
	*MockService
	pb.UnimplementedCertificateAuthorityServiceServer
}

func (s *PrivateCAV1) GetCaPool(ctx context.Context, req *pb.GetCaPoolRequest) (*pb.CaPool, error) {
	name, err := s.parseCAPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CaPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "caPool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading caPool: %v", err)
		}
	}

	return obj, nil
}

func (s *PrivateCAV1) CreateCaPool(ctx context.Context, req *pb.CreateCaPoolRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/caPools/" + req.CaPoolId
	name, err := s.parseCAPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.CaPool).(*pb.CaPool)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating caPool: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *PrivateCAV1) UpdateCaPool(ctx context.Context, req *pb.UpdateCaPoolRequest) (*longrunning.Operation, error) {
	reqName := req.GetCaPool().GetName()

	name, err := s.parseCAPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CaPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "caPool %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading caPool: %v", err)
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "issuancePolicy":
			obj.IssuancePolicy = req.GetCaPool().GetIssuancePolicy()
		case "publishingOptions":
			obj.PublishingOptions = req.GetCaPool().GetPublishingOptions()
		case "labels":
			obj.Labels = req.GetCaPool().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating caPool: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *PrivateCAV1) DeleteCaPool(ctx context.Context, req *pb.DeleteCaPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseCAPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.CaPool{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "caPool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting caPool: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
