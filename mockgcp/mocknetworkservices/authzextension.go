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
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.AuthzExtension

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

func (s *NetworkServicesServer) GetAuthzExtension(ctx context.Context, req *pb.GetAuthzExtensionRequest) (*pb.AuthzExtension, error) {
	name, err := s.parseAuthzExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AuthzExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NetworkServicesServer) ListAuthzExtensions(ctx context.Context, req *pb.ListAuthzExtensionsRequest) (*pb.ListAuthzExtensionsResponse, error) {
	response := &pb.ListAuthzExtensionsResponse{}

	parent, err := s.parseAuthzExtensionParent(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/authzExtensions/"

	findKind := (&pb.AuthzExtension{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.AuthzExtension)
		response.AuthzExtensions = append(response.AuthzExtensions, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateAuthzExtension(ctx context.Context, req *pb.CreateAuthzExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/authzExtensions/" + req.AuthzExtensionId
	name, err := s.parseAuthzExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.AuthzExtension)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.normalizeAuthzExtension(ctx, obj); err != nil {
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

func (s *NetworkServicesServer) UpdateAuthzExtension(ctx context.Context, req *pb.UpdateAuthzExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetAuthzExtension().GetName()

	name, err := s.parseAuthzExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.AuthzExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.AuthzExtension.CreateTime = obj.CreateTime
		req.AuthzExtension.UpdateTime = timestamppb.New(now)
		req.AuthzExtension.Name = obj.Name
		obj = req.AuthzExtension
	} else {
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetAuthzExtension().GetLabels()
			case "description":
				obj.Description = req.GetAuthzExtension().GetDescription()
			case "name":
				if req.GetAuthzExtension().GetName() != obj.GetName() {
					return nil, status.Errorf(codes.InvalidArgument, "field name is immutable")
				}
			case "loadBalancingScheme", "load_balancing_scheme":
				if req.GetAuthzExtension().GetLoadBalancingScheme() != obj.GetLoadBalancingScheme() {
					return nil, status.Errorf(codes.InvalidArgument, "field load_balancing_scheme is immutable")
				}
			case "authority":
				obj.Authority = req.GetAuthzExtension().GetAuthority()
			case "service":
				obj.Service = req.GetAuthzExtension().GetService()
			case "timeout":
				obj.Timeout = req.GetAuthzExtension().GetTimeout()
			case "failOpen", "fail_open":
				obj.FailOpen = req.GetAuthzExtension().GetFailOpen()
			case "metadata":
				obj.Metadata = req.GetAuthzExtension().GetMetadata()
			case "forwardHeaders", "forward_headers":
				obj.ForwardHeaders = req.GetAuthzExtension().GetForwardHeaders()
			case "wireFormat", "wire_format":
				obj.WireFormat = req.GetAuthzExtension().GetWireFormat()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.normalizeAuthzExtension(ctx, obj); err != nil {
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

func (s *NetworkServicesServer) DeleteAuthzExtension(ctx context.Context, req *pb.DeleteAuthzExtensionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAuthzExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AuthzExtension{}
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

type authzExtensionParent struct {
	Project  *projects.ProjectData
	Location string
}

func (p *authzExtensionParent) String() string {
	return "projects/" + p.Project.ID + "/locations/" + p.Location
}

func (s *NetworkServicesServer) parseAuthzExtensionParent(parent string) (*authzExtensionParent, error) {
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

		return &authzExtensionParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
}

type authzExtensionName struct {
	Project            *projects.ProjectData
	Location           string
	AuthzExtensionName string
}

func (n *authzExtensionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/authzExtensions/" + n.AuthzExtensionName
}

func (s *NetworkServicesServer) parseAuthzExtensionName(name string) (*authzExtensionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "authzExtensions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &authzExtensionName{
			Project:            project,
			Location:           tokens[3],
			AuthzExtensionName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *NetworkServicesServer) normalizeAuthzExtension(ctx context.Context, obj *pb.AuthzExtension) error {
	newService, err := s.replaceProjectIDWithNumberInURL(ctx, obj.Service)
	if err != nil {
		return err
	}
	obj.Service = newService
	return nil
}
