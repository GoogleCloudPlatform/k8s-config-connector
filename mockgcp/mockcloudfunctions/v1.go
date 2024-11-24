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
	"fmt"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
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
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		klog.Warningf("update_mask was not provided in request, should be required")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetFunction(), req.UpdateMask.GetPaths()); err != nil {
		return nil, fmt.Errorf("applying updates: %w", err)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
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
		return nil, err
	}

	return s.operations.NewLRO(ctx)
}
