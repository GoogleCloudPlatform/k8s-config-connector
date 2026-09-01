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

package mockapihub

import (
	"context"
	"strings"

	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ApiHubServer) GetAttribute(ctx context.Context, req *pb.GetAttributeRequest) (*pb.Attribute, error) {
	name, err := s.parseAttributeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Attribute{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) CreateAttribute(ctx context.Context, req *pb.CreateAttributeRequest) (*pb.Attribute, error) {
	reqName := req.Parent + "/attributes/" + req.AttributeId
	name, err := s.parseAttributeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Attribute).(*pb.Attribute)
	obj.Name = fqn

	now := timestamppb.Now()
	obj.CreateTime = now
	obj.UpdateTime = now
	obj.DefinitionType = pb.Attribute_USER_DEFINED
	obj.Mandatory = false

	if obj.Cardinality == 0 {
		obj.Cardinality = 1
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) UpdateAttribute(ctx context.Context, req *pb.UpdateAttributeRequest) (*pb.Attribute, error) {
	attrName := req.GetAttribute().GetName()

	name, err := s.parseAttributeName(attrName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Attribute{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updateMask := req.GetUpdateMask()
	paths := updateMask.GetPaths()
	if len(paths) == 0 {
		// Treat as all paths if empty
		paths = []string{
			"display_name", "description", "scope", "data_type", "allowed_values", "cardinality",
		}
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetAttribute().GetDisplayName()
		case "description":
			obj.Description = req.GetAttribute().GetDescription()
		case "scope":
			obj.Scope = req.GetAttribute().GetScope()
		case "data_type", "dataType":
			obj.DataType = req.GetAttribute().GetDataType()
		case "allowed_values", "allowedValues":
			obj.AllowedValues = req.GetAttribute().GetAllowedValues()
		case "cardinality":
			obj.Cardinality = req.GetAttribute().GetCardinality()
			if obj.Cardinality == 0 {
				obj.Cardinality = 1
			}
		}
	}

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *ApiHubServer) DeleteAttribute(ctx context.Context, req *pb.DeleteAttributeRequest) (*emptypb.Empty, error) {
	name, err := s.parseAttributeName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Attribute{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ApiHubServer) ListAttributes(ctx context.Context, req *pb.ListAttributesRequest) (*pb.ListAttributesResponse, error) {
	name, err := s.parseAttributeName(req.Parent + "/attributes/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListAttributesResponse{}

	findPrefix := "projects/" + name.Project.ID + "/locations/" + name.Location + "/attributes/"

	attributeKind := (&pb.Attribute{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, attributeKind, storage.ListOptions{}, func(obj proto.Message) error {
		attribute := obj.(*pb.Attribute)
		if strings.HasPrefix(attribute.GetName(), findPrefix) {
			response.Attributes = append(response.Attributes, attribute)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
