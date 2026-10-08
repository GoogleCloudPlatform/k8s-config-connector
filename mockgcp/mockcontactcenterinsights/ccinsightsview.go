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

package mockcontactcenterinsights

import (
	"context"
	"hash/fnv"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
)

func generateViewID(name string) string {
	h := fnv.New64a()
	h.Write([]byte(name))
	return strconv.FormatUint(h.Sum64(), 10)
}

func (s *ContactCenterInsightsServer) GetView(ctx context.Context, req *pb.GetViewRequest) (*pb.View, error) {
	name, err := s.parseViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.View{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No view found for project: `%d` and view Id: `%s`.", name.Project.Number, name.View)
		}
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) CreateView(ctx context.Context, req *pb.CreateViewRequest) (*pb.View, error) {
	reqName := req.GetView().GetName()
	name, err := s.parseViewName(reqName)
	if err != nil {
		return nil, err
	}

	// Generate a stable numeric view ID based on the requested view name to mimic real GCP behavior
	viewID := generateViewID(name.View)
	name.View = viewID

	fqn := name.String()

	obj := proto.Clone(req.View).(*pb.View)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) UpdateView(ctx context.Context, req *pb.UpdateViewRequest) (*pb.View, error) {
	reqName := req.GetView().GetName()
	name, err := s.parseViewName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.View{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No view found for project: `%d` and view Id: `%s`.", name.Project.Number, name.View)
		}
		return nil, err
	}

	if req.UpdateMask != nil && len(req.UpdateMask.Paths) > 0 {
		if err := fields.UpdateByFieldMask(obj, req.View, req.UpdateMask.Paths); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update CCInsightsView fields: %v", err)
		}
	} else {
		obj = proto.Clone(req.View).(*pb.View)
	}

	obj.Name = fqn

	now := timestamppb.Now()
	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ContactCenterInsightsServer) DeleteView(ctx context.Context, req *pb.DeleteViewRequest) (*emptypb.Empty, error) {
	name, err := s.parseViewName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.View{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "No view found for project: `%d` and view Id: `%s`.", name.Project.Number, name.View)
		}
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
