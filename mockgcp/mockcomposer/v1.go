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

package mockcomposer

import (
	"context"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/orchestration/airflow/service/v1"
)

type ComposerV1 struct {
	*MockService
	pb.UnimplementedEnvironmentsServer
}

func (s *ComposerV1) GetEnvironment(ctx context.Context, req *pb.GetEnvironmentRequest) (*pb.Environment, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error reading environment: %v", err)
		}
	}

	return obj, nil
}
func (s *ComposerV1) CreateEnvironment(ctx context.Context, req *pb.CreateEnvironmentRequest) (*longrunning.Operation, error) {
	reqName := req.GetEnvironment().GetName()
	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if req.GetParent()+"/environments/"+name.EnvironmentName != fqn {
		return nil, status.Errorf(codes.InvalidArgument, "name %q does not match parent", reqName)
	}

	obj := proto.Clone(req.Environment).(*pb.Environment)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating environment: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *ComposerV1) UpdateEnvironment(ctx context.Context, req *pb.UpdateEnvironmentRequest) (*longrunning.Operation, error) {
	reqName := req.GetEnvironment().GetName()

	name, err := s.parseEnvironmentName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Environment{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", reqName)
		}
		return nil, status.Errorf(codes.Internal, "error reading environment: %v", err)
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "labels":
			obj.Labels = req.GetEnvironment().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating environment: %v", err)
	}

	return s.operations.NewLRO(ctx)
}

func (s *ComposerV1) DeleteEnvironment(ctx context.Context, req *pb.DeleteEnvironmentRequest) (*longrunning.Operation, error) {
	name, err := s.parseEnvironmentName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	kind := (&pb.Environment{}).ProtoReflect().Descriptor()
	if err := s.storage.Delete(ctx, kind, fqn); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "environment %q not found", name)
		} else {
			return nil, status.Errorf(codes.Internal, "error deleting environment: %v", err)
		}
	}

	return s.operations.NewLRO(ctx)
}
