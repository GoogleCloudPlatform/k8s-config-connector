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

package mockservicedirectory

import (
	"context"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type LookupServiceV1 struct {
	*MockService
	pb.UnimplementedLookupServiceServer
}

func (s *LookupServiceV1) ResolveService(ctx context.Context, req *pb.ResolveServiceRequest) (*pb.ResolveServiceResponse, error) {
	name, err := s.parseServiceName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	service := &pb.Service{}
	if err := s.storage.Get(ctx, fqn, service); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	response := &pb.ResolveServiceResponse{
		Service: proto.Clone(service).(*pb.Service),
	}

	// Fetch endpoints
	prefix := fqn + "/endpoints/"
	if err := s.storage.List(ctx, (&pb.Endpoint{}).ProtoReflect().Descriptor(), storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		endpoint := obj.(*pb.Endpoint)
		response.Service.Endpoints = append(response.Service.Endpoints, endpoint)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
