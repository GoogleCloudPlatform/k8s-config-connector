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

// +tool:mockgcp-support
// proto.service: google.cloud.networkservices.v1.DepService
// proto.message: google.cloud.networkservices.v1.LbTrafficExtension

package mocknetworkservices

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NetworkServicesServer) GetLbTrafficExtension(ctx context.Context, req *pb.GetLbTrafficExtensionRequest) (*pb.LbTrafficExtension, error) {
	name, err := s.parseLbTrafficExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LbTrafficExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListLbTrafficExtensions(ctx context.Context, req *pb.ListLbTrafficExtensionsRequest) (*pb.ListLbTrafficExtensionsResponse, error) {
	response := &pb.ListLbTrafficExtensionsResponse{}

	parent, err := s.parseLbTrafficExtensionParent(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/lbTrafficExtensions/"

	findKind := (&pb.LbTrafficExtension{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.LbTrafficExtension)
		response.LbTrafficExtensions = append(response.LbTrafficExtensions, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateLbTrafficExtension(ctx context.Context, req *pb.CreateLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/lbTrafficExtensions/" + req.LbTrafficExtensionId
	name, err := s.parseLbTrafficExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.LbTrafficExtension)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.normalizeLbTrafficExtension(ctx, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) UpdateLbTrafficExtension(ctx context.Context, req *pb.UpdateLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetLbTrafficExtension().GetName()

	name, err := s.parseLbTrafficExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.LbTrafficExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.LbTrafficExtension.CreateTime = obj.CreateTime
		req.LbTrafficExtension.UpdateTime = timestamppb.New(now)
		req.LbTrafficExtension.Name = obj.Name
		obj = req.LbTrafficExtension
	} else {
		// gcloud uses camelCase for some fields in updateMask; handle both.
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetLbTrafficExtension().GetLabels()
			case "description":
				obj.Description = req.GetLbTrafficExtension().GetDescription()
			case "name":
				if req.GetLbTrafficExtension().GetName() != obj.GetName() {
					return nil, status.Errorf(codes.InvalidArgument, "field name is immutable")
				}
			case "forwardingRules", "forwarding_rules":
				obj.ForwardingRules = req.GetLbTrafficExtension().GetForwardingRules()
			case "extensionChains", "extension_chains":
				obj.ExtensionChains = req.GetLbTrafficExtension().GetExtensionChains()
			case "loadBalancingScheme", "load_balancing_scheme":
				if req.GetLbTrafficExtension().GetLoadBalancingScheme() != obj.GetLoadBalancingScheme() {
					return nil, status.Errorf(codes.InvalidArgument, "field load_balancing_scheme is immutable")
				}
			case "metadata":
				obj.Metadata = req.GetLbTrafficExtension().GetMetadata()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.normalizeLbTrafficExtension(ctx, obj); err != nil {
		return nil, err
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NetworkServicesServer) DeleteLbTrafficExtension(ctx context.Context, req *pb.DeleteLbTrafficExtensionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLbTrafficExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.LbTrafficExtension{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	now := time.Now()
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		result := &emptypb.Empty{}
		return result, nil
	})
}

type lbTrafficExtensionParent struct {
	Project  *projects.ProjectData
	Location string
}

func (p *lbTrafficExtensionParent) String() string {
	return "projects/" + p.Project.ID + "/locations/" + p.Location
}

func (s *NetworkServicesServer) parseLbTrafficExtensionParent(parent string) (*lbTrafficExtensionParent, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &lbTrafficExtensionParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

type lbTrafficExtensionName struct {
	Project                *projects.ProjectData
	Location               string
	LbTrafficExtensionName string
}

func (n *lbTrafficExtensionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/lbTrafficExtensions/" + n.LbTrafficExtensionName
}

func (s *NetworkServicesServer) parseLbTrafficExtensionName(name string) (*lbTrafficExtensionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lbTrafficExtensions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &lbTrafficExtensionName{
			Project:                project,
			Location:               tokens[3],
			LbTrafficExtensionName: tokens[5],
		}

		return name, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *NetworkServicesServer) normalizeLbTrafficExtension(ctx context.Context, obj *pb.LbTrafficExtension) error {
	for i, rule := range obj.ForwardingRules {
		newRule, err := s.replaceProjectIDWithNumberInURL(ctx, rule)
		if err != nil {
			return err
		}
		obj.ForwardingRules[i] = newRule
	}
	for _, chain := range obj.ExtensionChains {
		for _, extension := range chain.Extensions {
			if extension.Service != "" {
				newService, err := s.replaceProjectIDWithNumberInURL(ctx, extension.Service)
				if err != nil {
					return err
				}
				extension.Service = newService
			}
		}
	}
	return nil
}
