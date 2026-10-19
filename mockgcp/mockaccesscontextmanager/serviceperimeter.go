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
// proto.service: google.identity.accesscontextmanager.v1.AccessContextManager
// proto.message: google.identity.accesscontextmanager.v1.ServicePerimeter

package mockaccesscontextmanager

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/accesscontextmanager/apiv1/accesscontextmanagerpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *AccessContextManagerV1) GetServicePerimeter(ctx context.Context, req *pb.GetServicePerimeterRequest) (*pb.ServicePerimeter, error) {
	name, err := s.parseServicePerimeterName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.ServicePerimeter{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "ServicePerimeter %q not found.", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *AccessContextManagerV1) ListServicePerimeters(ctx context.Context, req *pb.ListServicePerimetersRequest) (*pb.ListServicePerimetersResponse, error) {
	name, err := s.parseServicePerimeterName(req.GetParent() + "/servicePerimeters/dummy")
	if err != nil {
		return nil, err
	}
	prefix := strings.TrimSuffix(name.String(), "dummy")

	response := &pb.ListServicePerimetersResponse{}
	kind := (&pb.ServicePerimeter{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(p proto.Message) error {
		servicePerimeter := p.(*pb.ServicePerimeter)
		response.ServicePerimeters = append(response.ServicePerimeters, servicePerimeter)
		return nil
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "Error listing ServicePerimeters: %v", err)
	}

	return response, nil
}

func (s *AccessContextManagerV1) CreateServicePerimeter(ctx context.Context, req *pb.CreateServicePerimeterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServicePerimeterName(req.GetServicePerimeter().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.CloneOf(req.GetServicePerimeter())
	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/create/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, obj)
}

func (s *AccessContextManagerV1) UpdateServicePerimeter(ctx context.Context, req *pb.UpdateServicePerimeterRequest) (*longrunningpb.Operation, error) {
	if req.GetServicePerimeter() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "service_perimeter is required for update")
	}
	name, err := s.parseServicePerimeterName(req.GetServicePerimeter().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.ServicePerimeter{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := proto.CloneOf(existing)
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "title":
			updated.Title = req.GetServicePerimeter().GetTitle()
		case "description":
			updated.Description = req.GetServicePerimeter().GetDescription()
		case "status":
			updated.Status = req.GetServicePerimeter().GetStatus()
		case "spec":
			updated.Spec = req.GetServicePerimeter().GetSpec()
		case "use_explicit_dry_run_spec", "useExplicitDryRunSpec":
			updated.UseExplicitDryRunSpec = req.GetServicePerimeter().GetUseExplicitDryRunSpec()
		case "status.resources":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.Resources = req.GetServicePerimeter().GetStatus().GetResources()
		case "status.access_levels", "status.accessLevels":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.AccessLevels = req.GetServicePerimeter().GetStatus().GetAccessLevels()
		case "status.restricted_services", "status.restrictedServices":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.RestrictedServices = req.GetServicePerimeter().GetStatus().GetRestrictedServices()
		case "status.vpc_accessible_services", "status.vpcAccessibleServices":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.VpcAccessibleServices = req.GetServicePerimeter().GetStatus().GetVpcAccessibleServices()
		case "status.ingress_policies", "status.ingressPolicies":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.IngressPolicies = req.GetServicePerimeter().GetStatus().GetIngressPolicies()
		case "status.egress_policies", "status.egressPolicies":
			if updated.Status == nil {
				updated.Status = &pb.ServicePerimeterConfig{}
			}
			updated.Status.EgressPolicies = req.GetServicePerimeter().GetStatus().GetEgressPolicies()
		case "spec.resources":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.Resources = req.GetServicePerimeter().GetSpec().GetResources()
		case "spec.access_levels", "spec.accessLevels":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.AccessLevels = req.GetServicePerimeter().GetSpec().GetAccessLevels()
		case "spec.restricted_services", "spec.restrictedServices":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.RestrictedServices = req.GetServicePerimeter().GetSpec().GetRestrictedServices()
		case "spec.vpc_accessible_services", "spec.vpcAccessibleServices":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.VpcAccessibleServices = req.GetServicePerimeter().GetSpec().GetVpcAccessibleServices()
		case "spec.ingress_policies", "spec.ingressPolicies":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.IngressPolicies = req.GetServicePerimeter().GetSpec().GetIngressPolicies()
		case "spec.egress_policies", "spec.egressPolicies":
			if updated.Spec == nil {
				updated.Spec = &pb.ServicePerimeterConfig{}
			}
			updated.Spec.EgressPolicies = req.GetServicePerimeter().GetSpec().GetEgressPolicies()
		default:
			return nil, fmt.Errorf("UpdateServicePerimeter: unsupported update_mask path (in mock) %q", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/update/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, updated)
}

func (s *AccessContextManagerV1) DeleteServicePerimeter(ctx context.Context, req *pb.DeleteServicePerimeterRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseServicePerimeterName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.ServicePerimeter{}); err != nil {
		if status.Code(err) == codes.NotFound {
			// Deleting a non-existent object should succeed idempotently.
			metadata := &pb.AccessContextManagerOperationMetadata{}
			return s.operations.DoneLRO(ctx, fqn, metadata, &emptypb.Empty{})
		}
		return nil, err
	}

	metadata := &pb.AccessContextManagerOperationMetadata{}
	lroPrefix := "operations/" + fqn + "/delete/{{operationID}}"
	return s.operations.DoneLRO(ctx, lroPrefix, metadata, &emptypb.Empty{})
}

type servicePerimeterName struct {
	AccessPolicy     string
	ServicePerimeter string
}

func (n *servicePerimeterName) String() string {
	return "accessPolicies/" + n.AccessPolicy + "/servicePerimeters/" + n.ServicePerimeter
}

// parseServicePerimeterName parses a string into a servicePerimeterName.
// The expected form is accessPolicies/{access_policy}/servicePerimeters/{service_perimeter}
func (s *AccessContextManagerV1) parseServicePerimeterName(name string) (*servicePerimeterName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "accessPolicies" && tokens[2] == "servicePerimeters" {
		return &servicePerimeterName{AccessPolicy: tokens[1], ServicePerimeter: tokens[3]}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
