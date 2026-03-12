// Copyright 2025 Google LLC
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
// proto.message: google.cloud.networkservices.v1.LbRouteExtension

package mocknetworkservices

import (
	"context"
	"fmt"
	"strconv"
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

func (s *NetworkServicesServer) GetLbRouteExtension(ctx context.Context, req *pb.GetLbRouteExtensionRequest) (*pb.LbRouteExtension, error) {
	name, err := s.parseLbRouteExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.LbRouteExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}
func (s *NetworkServicesServer) ListLbRouteExtensions(ctx context.Context, req *pb.ListLbRouteExtensionsRequest) (*pb.ListLbRouteExtensionsResponse, error) {
	response := &pb.ListLbRouteExtensionsResponse{}

	parent, err := s.parseLbRouteExtensionParent(req.Parent)
	if err != nil {
		return nil, err
	}
	prefix := parent.String() + "/lbRouteExtensions/"

	findKind := (&pb.LbRouteExtension{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		item := obj.(*pb.LbRouteExtension)
		response.LbRouteExtensions = append(response.LbRouteExtensions, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *NetworkServicesServer) CreateLbRouteExtension(ctx context.Context, req *pb.CreateLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/lbRouteExtensions/" + req.LbRouteExtensionId
	name, err := s.parseLbRouteExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.LbRouteExtension).(*pb.LbRouteExtension)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.normalizeLbRouteExtension(ctx, obj); err != nil {
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

func (s *NetworkServicesServer) UpdateLbRouteExtension(ctx context.Context, req *pb.UpdateLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetLbRouteExtension().GetName()

	name, err := s.parseLbRouteExtensionName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.LbRouteExtension{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		req.LbRouteExtension.CreateTime = obj.CreateTime
		req.LbRouteExtension.UpdateTime = timestamppb.New(now)
		req.LbRouteExtension.Name = obj.Name
		obj = req.LbRouteExtension
	} else {
		// gcloud uses camelCase for some fields in updateMask; handle both.
		for _, path := range paths {
			switch path {
			case "labels":
				obj.Labels = req.GetLbRouteExtension().GetLabels()
			case "description":
				obj.Description = req.GetLbRouteExtension().GetDescription()
			case "name":
				if req.GetLbRouteExtension().GetName() != obj.GetName() {
					return nil, status.Errorf(codes.InvalidArgument, "field name is immutable")
				}
			case "forwardingRules", "forwarding_rules":
				obj.ForwardingRules = req.GetLbRouteExtension().GetForwardingRules()
			case "extensionChains", "extension_chains":
				obj.ExtensionChains = req.GetLbRouteExtension().GetExtensionChains()
			case "loadBalancingScheme", "load_balancing_scheme":
				if req.GetLbRouteExtension().GetLoadBalancingScheme() != obj.GetLoadBalancingScheme() {
					return nil, status.Errorf(codes.InvalidArgument, "field load_balancing_scheme is immutable")
				}
			case "metadata":
				obj.Metadata = req.GetLbRouteExtension().GetMetadata()
			default:
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
			}
		}
		obj.UpdateTime = timestamppb.New(now)
	}

	if err := s.normalizeLbRouteExtension(ctx, obj); err != nil {
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

func (s *NetworkServicesServer) DeleteLbRouteExtension(ctx context.Context, req *pb.DeleteLbRouteExtensionRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseLbRouteExtensionName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.LbRouteExtension{}
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

type lbRouteExtensionParent struct {
	Project  *projects.ProjectData
	Location string
}

func (p *lbRouteExtensionParent) String() string {
	return "projects/" + p.Project.ID + "/locations/" + p.Location
}

func (s *NetworkServicesServer) parseLbRouteExtensionParent(parent string) (*lbRouteExtensionParent, error) {
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

		return &lbRouteExtensionParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
	}
}

type lbRouteExtensionName struct {
	Project              *projects.ProjectData
	Location             string
	LbRouteExtensionName string
}

func (n *lbRouteExtensionName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/lbRouteExtensions/" + n.LbRouteExtensionName
}

func (s *NetworkServicesServer) parseLbRouteExtensionName(name string) (*lbRouteExtensionName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "lbRouteExtensions" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &lbRouteExtensionName{
			Project:              project,
			Location:             tokens[3],
			LbRouteExtensionName: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *NetworkServicesServer) normalizeLbRouteExtension(ctx context.Context, obj *pb.LbRouteExtension) error {
	for i, rule := range obj.ForwardingRules {
		newRule, err := s.replaceProjectIDWithNumberInURL(ctx, rule)
		if err != nil {
			return err
		}
		obj.ForwardingRules[i] = newRule
	}
	for _, chain := range obj.ExtensionChains {
		for _, extension := range chain.Extensions {
			newService, err := s.replaceProjectIDWithNumberInURL(ctx, extension.Service)
			if err != nil {
				return err
			}
			extension.Service = newService
		}
	}
	return nil
}

func (s *NetworkServicesServer) replaceProjectIDWithNumberInURL(ctx context.Context, url string) (string, error) {
	if !strings.HasPrefix(url, "https://www.googleapis.com/compute/v1/projects/") &&
		!strings.HasPrefix(url, "https://compute.googleapis.com/compute/v1/projects/") {
		return url, nil
	}

	// Format is https://[hostname]/compute/v1/projects/[projectID]/...
	tokens := strings.Split(url, "/")
	if len(tokens) < 7 {
		return url, nil
	}

	projectIDOrNumber := tokens[6]
	project, err := s.Projects.GetProjectByIDOrNumber(projectIDOrNumber)
	if err != nil {
		return url, nil // Should we return error or just return original URL? Returning original for now.
	}

	tokens[6] = strconv.FormatInt(project.Number, 10)
	return strings.Join(tokens, "/"), nil
}
