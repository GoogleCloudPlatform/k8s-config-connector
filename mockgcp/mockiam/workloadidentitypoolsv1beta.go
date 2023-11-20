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

package mockiam

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/iam/v1beta"
)

type WorkloadIdentityPoolsV1Beta struct {
	*MockService
	pb.UnimplementedWorkloadIdentityPoolsServer
}

func (s *WorkloadIdentityPoolsV1Beta) GetWorkloadIdentityPoolProvider(ctx context.Context, req *pb.GetWorkloadIdentityPoolProviderRequest) (*pb.WorkloadIdentityPoolProvider, error) {
	name, err := s.parseWorkloadIdentityPoolProviderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkloadIdentityPoolProvider{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPoolProvider %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading workloadIdentityPoolProvider: %v", err)
		}
	}

	return obj, nil
}
func (s *WorkloadIdentityPoolsV1Beta) CreateWorkloadIdentityPoolProvider(ctx context.Context, req *pb.CreateWorkloadIdentityPoolProviderRequest) (*longrunning.Operation, error) {
	reqName := req.GetWorkloadIdentityPoolProvider().GetName()
	name, err := s.parseWorkloadIdentityPoolProviderName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.WorkloadIdentityPoolProvider).(*pb.WorkloadIdentityPoolProvider)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating workloadIdentityPoolProvider: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *WorkloadIdentityPoolsV1Beta) UpdateWorkloadIdentityPoolProvider(ctx context.Context, req *pb.UpdateWorkloadIdentityPoolProviderRequest) (*longrunning.Operation, error) {
	reqName := req.GetWorkloadIdentityPoolProvider().GetName()

	name, err := s.parseWorkloadIdentityPoolProviderName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WorkloadIdentityPoolProvider{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPoolProvider %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading workloadIdentityPoolProvider: %v", err)
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
			obj.Description = req.GetWorkloadIdentityPoolProvider().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating workloadIdentityPoolProvider: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *WorkloadIdentityPoolsV1Beta) DeleteWorkloadIdentityPoolProvider(ctx context.Context, req *pb.DeleteWorkloadIdentityPoolProviderRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkloadIdentityPoolProviderName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.WorkloadIdentityPoolProvider{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPoolProvider %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting workloadIdentityPoolProvider: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}

func (s *WorkloadIdentityPoolsV1Beta) GetWorkloadIdentityPool(ctx context.Context, req *pb.GetWorkloadIdentityPoolRequest) (*pb.WorkloadIdentityPool, error) {
	name, err := s.parseWorkloadIdentityPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading workloadIdentityPool: %v", err)
		}
	}

	return obj, nil
}

func (s *WorkloadIdentityPoolsV1Beta) CreateWorkloadIdentityPool(ctx context.Context, req *pb.CreateWorkloadIdentityPoolRequest) (*longrunning.Operation, error) {
	reqName := req.GetWorkloadIdentityPool().GetName()
	name, err := s.parseWorkloadIdentityPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.WorkloadIdentityPool).(*pb.WorkloadIdentityPool)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating workloadIdentityPool: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *WorkloadIdentityPoolsV1Beta) UpdateWorkloadIdentityPool(ctx context.Context, req *pb.UpdateWorkloadIdentityPoolRequest) (*longrunning.Operation, error) {
	reqName := req.GetWorkloadIdentityPool().GetName()

	name, err := s.parseWorkloadIdentityPoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPool %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading workloadIdentityPool: %v", err)
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
			obj.Description = req.GetWorkloadIdentityPool().GetDescription()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating workloadIdentityPool: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *WorkloadIdentityPoolsV1Beta) DeleteWorkloadIdentityPool(ctx context.Context, req *pb.DeleteWorkloadIdentityPoolRequest) (*longrunning.Operation, error) {
	name, err := s.parseWorkloadIdentityPoolName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.WorkloadIdentityPool{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "workloadIdentityPool %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting workloadIdentityPool: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
