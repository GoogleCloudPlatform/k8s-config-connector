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

package mockapigee

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

type apiproductName struct {
	Organization string
	Apiproduct   string
}

func (n *apiproductName) String() string {
	return fmt.Sprintf("organizations/%v/apiproducts/%v", n.Organization, n.Apiproduct)
}

func (s *MockService) parseApiproductName(name string) (*apiproductName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 || tokens[0] != "organizations" || tokens[2] != "apiproducts" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid apiproduct name %q: must be of the form organizations/org/apiproducts/apiproduct", name)
	}

	return &apiproductName{
		Organization: tokens[1],
		Apiproduct:   tokens[3],
	}, nil
}

type apiproductsServer struct {
	*MockService
	pb.UnimplementedOrganizationsApiproductsServerServer
}

func (s *apiproductsServer) GetOrganizationsApiproduct(ctx context.Context, req *pb.GetOrganizationsApiproductRequest) (*pb.GoogleCloudApigeeV1ApiProduct, error) {
	name, err := s.parseApiproductName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1ApiProduct{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "resource %s not found", fqn)
		}
		return nil, err
	}
	return obj, nil
}

func (s *apiproductsServer) CreateOrganizationsApiproduct(ctx context.Context, req *pb.CreateOrganizationsApiproductRequest) (*pb.GoogleCloudApigeeV1ApiProduct, error) {
	reqName := req.Parent + "/apiproducts/" + req.OrganizationsApiproduct.Name
	name, err := s.parseApiproductName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.OrganizationsApiproduct).(*pb.GoogleCloudApigeeV1ApiProduct)
	obj.Name = name.Apiproduct

	// Set timestamps
	now := int64(1420070400000) // Jan 1, 2015 - stable dummy timestamp
	obj.CreatedAt = now
	obj.LastModifiedAt = now

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *apiproductsServer) UpdateOrganizationsApiproduct(ctx context.Context, req *pb.UpdateOrganizationsApiproductRequest) (*pb.GoogleCloudApigeeV1ApiProduct, error) {
	name, err := s.parseApiproductName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.GoogleCloudApigeeV1ApiProduct{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updated := proto.Clone(req.GetOrganizationsApiproduct()).(*pb.GoogleCloudApigeeV1ApiProduct)
	updated.Name = obj.Name
	updated.CreatedAt = obj.CreatedAt
	updated.LastModifiedAt = int64(1420070401000) // Slightly later stable timestamp

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *apiproductsServer) DeleteOrganizationsApiproduct(ctx context.Context, req *pb.DeleteOrganizationsApiproductRequest) (*pb.GoogleCloudApigeeV1ApiProduct, error) {
	name, err := s.parseApiproductName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.GoogleCloudApigeeV1ApiProduct{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return oldObj, nil
}
