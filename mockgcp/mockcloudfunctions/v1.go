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

package mockcloudfunctions

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/functions/v1"
)

type CloudFunctionsV1 struct {
	*MockService
	pb.UnimplementedCloudFunctionsServiceServer
}

func (s *CloudFunctionsV1) GetFunction(ctx context.Context, req *pb.GetFunctionRequest) (*pb.CloudFunction, error) {
	name, err := s.parseFunctionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CloudFunction{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "function %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading function: %v", err)
		}
	}

	return obj, nil
}
func (s *CloudFunctionsV1) CreateFunction(ctx context.Context, req *pb.CreateFunctionRequest) (*longrunning.Operation, error) {
	reqName := /*req.Location + "/functions/" +*/ req.GetFunction().GetName()
	name, err := s.parseFunctionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Function).(*pb.CloudFunction)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating function: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *CloudFunctionsV1) DeleteFunction(ctx context.Context, req *pb.DeleteFunctionRequest) (*longrunning.Operation, error) {
	name, err := s.parseFunctionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	kind := (&pb.CloudFunction{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "function %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting function: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
