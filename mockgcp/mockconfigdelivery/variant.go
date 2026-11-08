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

package mockconfigdelivery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type variantName struct {
	Project        *projects.ProjectData
	Location       string
	ResourceBundle string
	Release        string
	Variant        string
}

func (n *variantName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/resourceBundles/" + n.ResourceBundle + "/releases/" + n.Release + "/variants/" + n.Variant
}

func (s *MockService) parseVariantName(name string) (*variantName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 10 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "resourceBundles" && tokens[6] == "releases" && tokens[8] == "variants" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		return &variantName{
			Project:        project,
			Location:       tokens[3],
			ResourceBundle: tokens[5],
			Release:        tokens[7],
			Variant:        tokens[9],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *ConfigDeliveryServer) GetVariant(ctx context.Context, req *pb.GetVariantRequest) (*pb.Variant, error) {
	name, err := s.parseVariantName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Variant{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	out := proto.Clone(obj).(*pb.Variant)
	// resources is input-only and not returned on GET
	out.Resources = nil

	return out, nil
}

func (s *ConfigDeliveryServer) ListVariants(ctx context.Context, req *pb.ListVariantsRequest) (*pb.ListVariantsResponse, error) {
	parentName, err := s.parseReleaseName(req.GetParent())
	if err != nil {
		return nil, err
	}

	parentFqn := parentName.String()
	parentObj := &pb.Release{}
	if err := s.storage.Get(ctx, parentFqn, parentObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentFqn)
		}
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s/resourceBundles/%s/releases/%s/variants/", parentName.Project.ID, parentName.Location, parentName.ResourceBundle, parentName.Release)

	response := &pb.ListVariantsResponse{}
	kind := (&pb.Variant{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := proto.Clone(obj).(*pb.Variant)
		// resources is input-only and not returned in responses
		item.Resources = nil
		response.Variants = append(response.Variants, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *ConfigDeliveryServer) CreateVariant(ctx context.Context, req *pb.CreateVariantRequest) (*longrunning.Operation, error) {
	parentName, err := s.parseReleaseName(req.Parent)
	if err != nil {
		return nil, err
	}

	parentFqn := parentName.String()
	parentObj := &pb.Release{}
	if err := s.storage.Get(ctx, parentFqn, parentObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", parentFqn)
		}
		return nil, err
	}

	reqName := req.Parent + "/variants/" + req.VariantId
	name, err := s.parseVariantName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := proto.Clone(req.Variant).(*pb.Variant)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Errorf(codes.AlreadyExists, "Resource '%s' already exists", fqn)
		}
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.Variant)
		lroResponse.Resources = nil
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) UpdateVariant(ctx context.Context, req *pb.UpdateVariantRequest) (*longrunning.Operation, error) {
	reqName := req.GetVariant().GetName()

	name, err := s.parseVariantName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	fqn := name.String()

	obj := &pb.Variant{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		obj.Labels = req.GetVariant().GetLabels()
		obj.Resources = req.GetVariant().GetResources()
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetVariant().GetLabels()
			case "resources":
				obj.Resources = req.GetVariant().GetResources()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
	}

	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		lroResponse := proto.Clone(obj).(*pb.Variant)
		lroResponse.Resources = nil
		return lroResponse, nil
	})
}

func (s *ConfigDeliveryServer) DeleteVariant(ctx context.Context, req *pb.DeleteVariantRequest) (*longrunning.Operation, error) {
	name, err := s.parseVariantName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	oldObj := &pb.Variant{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}
