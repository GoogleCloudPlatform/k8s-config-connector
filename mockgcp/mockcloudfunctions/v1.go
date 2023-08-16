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
	"k8s.io/klog/v2"

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

func (s *CloudFunctionsV1) UpdateFunction(ctx context.Context, req *pb.UpdateFunctionRequest) (*longrunning.Operation, error) {
	reqName := req.GetFunction().GetName()

	name, err := s.parseFunctionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.CloudFunction{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "cloudFunction %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading cloudFunction: %v", err)
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetFunction().GetDescription()
		case "labels":
			obj.Labels = req.GetFunction().GetLabels()
		case "timeout":
			obj.Timeout = req.GetFunction().GetTimeout()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating cloudFunction: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *CloudFunctionsV1) DeleteFunction(ctx context.Context, req *pb.DeleteFunctionRequest) (*longrunning.Operation, error) {
	name, err := s.parseFunctionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.CloudFunction{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "function %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting function: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
