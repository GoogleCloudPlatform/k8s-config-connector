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

package mockcompute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type LicensesV1 struct {
	*MockService
	pb.UnimplementedLicensesServer
}

func (s *LicensesV1) Get(ctx context.Context, req *pb.GetLicenseRequest) (*pb.License, error) {
	name, err := s.parseLicenseName(req.GetProject(), req.GetLicense())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.License{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *LicensesV1) Insert(ctx context.Context, req *pb.InsertLicenseRequest) (*pb.Operation, error) {
	name, err := s.parseLicenseName(req.GetProject(), req.GetLicenseResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	id := s.generateID()

	obj := proto.CloneOf(req.GetLicenseResource())
	obj.Name = PtrTo(name.License)
	obj.Id = &id
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.UpdateTimestamp = PtrTo(s.nowString())
	obj.SelfLink = PtrTo(BuildComputeSelfLink(ctx, fqn))
	obj.SelfLinkWithId = PtrTo(BuildComputeSelfLink(ctx, fqn))
	s.populateLicenseDefaults(ctx, obj, name.License)

	// ResourceRequirements is input-only and deprecated; it is not returned in GET responses.
	obj.ResourceRequirements = nil

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	now := time.Now()
	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("compute.licenses.insert"),
		User:          PtrTo("user@example.com"),
		Status:        PtrTo(pb.Operation_DONE),
		Progress:      PtrTo(int32(100)),
		EndTime:       PtrTo(formatTime(now)),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *LicensesV1) Update(ctx context.Context, req *pb.UpdateLicenseRequest) (*pb.Operation, error) {
	name, err := s.parseLicenseName(req.GetProject(), req.GetLicense())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.License{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	reqResource := req.GetLicenseResource()
	if req.GetUpdateMask() != "" {
		for _, path := range strings.Split(req.GetUpdateMask(), ",") {
			path = strings.TrimSpace(path)
			switch path {
			case "appendable_to_disk", "appendableToDisk":
				obj.AppendableToDisk = reqResource.AppendableToDisk
			case "description":
				obj.Description = reqResource.Description
			case "minimum_retention", "minimumRetention":
				obj.MinimumRetention = reqResource.MinimumRetention
			case "removable_from_disk", "removableFromDisk":
				obj.RemovableFromDisk = reqResource.RemovableFromDisk
			case "transferable":
				obj.Transferable = reqResource.Transferable
			case "multi_tenant_only", "multiTenantOnly":
				obj.MultiTenantOnly = reqResource.MultiTenantOnly
			case "sole_tenant_only", "soleTenantOnly":
				obj.SoleTenantOnly = reqResource.SoleTenantOnly
			case "os_license", "osLicense":
				obj.OsLicense = reqResource.OsLicense
			case "allowed_replacement_licenses", "allowedReplacementLicenses":
				obj.AllowedReplacementLicenses = reqResource.AllowedReplacementLicenses
			case "incompatible_licenses", "incompatibleLicenses":
				obj.IncompatibleLicenses = reqResource.IncompatibleLicenses
			case "required_coattached_licenses", "requiredCoattachedLicenses":
				obj.RequiredCoattachedLicenses = reqResource.RequiredCoattachedLicenses
			}
		}
	} else if reqResource != nil {
		proto.Merge(obj, reqResource)
	}

	obj.UpdateTimestamp = PtrTo(s.nowString())
	s.populateLicenseDefaults(ctx, obj, name.License)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      obj.Id,
		TargetLink:    obj.SelfLink,
		OperationType: PtrTo("update"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *LicensesV1) Delete(ctx context.Context, req *pb.DeleteLicenseRequest) (*pb.Operation, error) {
	name, err := s.parseLicenseName(req.GetProject(), req.GetLicense())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.License{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The resource '%s' was not found", fqn)
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetId:      deleted.Id,
		TargetLink:    deleted.SelfLink,
		OperationType: PtrTo("delete"),
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, name.Project.ID, op, func() (proto.Message, error) {
		return deleted, nil
	})
}

func (s *LicensesV1) List(ctx context.Context, req *pb.ListLicensesRequest) (*pb.LicensesListResponse, error) {
	project, err := s.Projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%s/global/licenses/", project.ID)
	response := &pb.LicensesListResponse{
		Id:       PtrTo(fmt.Sprintf("%d", s.generateID())),
		SelfLink: PtrTo(BuildComputeSelfLink(ctx, fmt.Sprintf("projects/%s/global/licenses", project.ID))),
	}

	findKind := (&pb.License{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		item := obj.(*pb.License)
		response.Items = append(response.Items, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *MockService) populateLicenseDefaults(ctx context.Context, obj *pb.License, name string) {
	obj.Kind = PtrTo("compute#license")
	obj.ChargesUseFee = PtrTo(false)

	if obj.AppendableToDisk == nil {
		obj.AppendableToDisk = PtrTo(false)
	}
	if obj.RemovableFromDisk == nil {
		obj.RemovableFromDisk = PtrTo(false)
	}
	if obj.OsLicense == nil {
		obj.OsLicense = PtrTo(false)
	}
	if obj.MultiTenantOnly == nil {
		obj.MultiTenantOnly = PtrTo(false)
	}
	if obj.SoleTenantOnly == nil {
		obj.SoleTenantOnly = PtrTo(false)
	}
	if obj.Transferable == nil {
		obj.Transferable = PtrTo(true)
	}

	if obj.MinimumRetention == nil {
		obj.MinimumRetention = &pb.Duration{
			Nanos:   PtrTo(int32(0)),
			Seconds: PtrTo(int64(0)),
		}
	} else {
		if obj.MinimumRetention.Nanos == nil {
			obj.MinimumRetention.Nanos = PtrTo(int32(0))
		}
		if obj.MinimumRetention.Seconds == nil {
			obj.MinimumRetention.Seconds = PtrTo(int64(0))
		}
	}

	if obj.LicenseCode == nil || *obj.LicenseCode == 0 {
		var code uint64
		if strings.Contains(name, "minimal") {
			code = 8227486501098990469
		} else if strings.Contains(name, "maximal") {
			code = 6897853162713229499
		} else {
			code = s.generateID()
		}
		obj.LicenseCode = &code
	}
}

type licenseName struct {
	Project *projects.ProjectData
	License string
}

func (n *licenseName) String() string {
	return "projects/" + n.Project.ID + "/global/licenses/" + n.License
}

func (s *MockService) parseLicenseName(project string, license string) (*licenseName, error) {
	if strings.Contains(license, "/") {
		tokens := strings.Split(license, "/")
		if len(tokens) == 5 && tokens[0] == "projects" && tokens[2] == "global" && tokens[3] == "licenses" {
			project = tokens[1]
			license = tokens[4]
		}
	}
	projectData, err := s.Projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}
	return &licenseName{
		Project: projectData,
		License: license,
	}, nil
}
