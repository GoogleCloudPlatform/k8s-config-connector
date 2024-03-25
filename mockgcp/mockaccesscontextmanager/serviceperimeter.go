// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockaccesscontextmanager

import (
	"context"

	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/identity/accesscontextmanager/v1"
	"google.golang.org/genproto/googleapis/longrunning"
)

func (s *AccessContextManagerV1) GetServicePerimeter(ctx context.Context, req *pb.GetServicePerimeterRequest) (*pb.ServicePerimeter, error) {
	name, err := s.parseServicePerimeterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServicePerimeter{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *AccessContextManagerV1) CreateServicePerimeter(ctx context.Context, req *pb.CreateServicePerimeterRequest) (*longrunning.Operation, error) {
	reqName := req.GetServicePerimeter().GetName()
	name, err := s.parseServicePerimeterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ServicePerimeter).(*pb.ServicePerimeter)
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) UpdateServicePerimeter(ctx context.Context, req *pb.UpdateServicePerimeterRequest) (*longrunning.Operation, error) {
	reqName := req.GetServicePerimeter().GetName()

	name, err := s.parseServicePerimeterName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.ServicePerimeter{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}

func (s *AccessContextManagerV1) DeleteServicePerimeter(ctx context.Context, req *pb.DeleteServicePerimeterRequest) (*longrunning.Operation, error) {
	name, err := s.parseServicePerimeterName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.ServicePerimeter{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
