// Copyright 2024 Google LLC
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

package mockdns

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type managedZonesService struct {
	*MockService
	pb.UnimplementedManagedZonesServerServer
}

func (s *managedZonesService) GetManagedZone(ctx context.Context, req *pb.GetManagedZoneRequest) (*pb.ManagedZone, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ManagedZone{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "managedZone %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *managedZonesService) CreateManagedZone(ctx context.Context, req *pb.CreateManagedZoneRequest) (*pb.ManagedZone, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/" + req.GetManagedZone().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ManagedZone).(*pb.ManagedZone)

	obj.CreationTime = PtrTo(time.Now().UTC().Format(time.RFC3339))

	if obj.CloudLoggingConfig == nil {
		obj.CloudLoggingConfig = &pb.ManagedZoneCloudLoggingConfig{}
	}
	obj.CloudLoggingConfig.Kind = PtrTo("dns#managedZoneCloudLoggingConfig")
	obj.Id = PtrTo[uint64](1234567890)
	obj.Kind = PtrTo("dns#managedZone")
	obj.NameServers = []string{
		"ns-cloud-c1.googledomains.com.",
		"ns-cloud-c2.googledomains.com.",
		"ns-cloud-c3.googledomains.com.",
		"ns-cloud-c4.googledomains.com.",
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *managedZonesService) UpdateManagedZone(ctx context.Context, req *pb.UpdateManagedZoneRequest) (*pb.Operation, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/" + req.GetManagedZone().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.ManagedZone

	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "managedZone %q not found", name)
		}
		return nil, err
	}

	updated := proto.Clone(req.ManagedZone).(*pb.ManagedZone)

	// These fields are output only and cannot be changed.
	updated.CreationTime = existing.CreationTime
	updated.CloudLoggingConfig = existing.CloudLoggingConfig
	updated.Id = existing.Id
	updated.Kind = existing.Kind
	updated.NameServers = existing.NameServers

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Status: PtrTo("done"),
		Type:   PtrTo("UPDATE"),
		User:   PtrTo(currentUser(ctx)),
		ZoneContext: &pb.OperationManagedZoneContext{
			NewValue: updated,
			OldValue: &existing,
		},
	}
	lroPrefix := "projects/" + req.GetProject() + "/managedZones/" + req.GetManagedZone().GetName() + "/"
	return s.operations.StartLRO(ctx, lroPrefix, op, func() (proto.Message, error) {
		return updated, nil
	})
}

func (s *managedZonesService) PatchManagedZone(ctx context.Context, req *pb.PatchManagedZoneRequest) (*pb.Operation, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/" + req.GetManagedZone().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	var existing pb.ManagedZone

	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "managedZone %q not found", name)
		}
		return nil, err
	}

	updated := proto.Clone(&existing).(*pb.ManagedZone)
	if req.GetManagedZone().Description != nil {
		updated.Description = req.GetManagedZone().Description
	}
	if req.GetManagedZone().DnsName != nil {
		updated.DnsName = req.GetManagedZone().DnsName
	}
	if req.GetManagedZone().Visibility != nil {
		updated.Visibility = req.GetManagedZone().Visibility
	}
	if req.GetManagedZone().PrivateVisibilityConfig != nil {
		updated.PrivateVisibilityConfig = req.GetManagedZone().PrivateVisibilityConfig
	}
	if req.GetManagedZone().CloudLoggingConfig != nil {
		updated.CloudLoggingConfig = req.GetManagedZone().CloudLoggingConfig
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		Status: PtrTo("done"),
		Type:   PtrTo("UPDATE"),
		User:   PtrTo(currentUser(ctx)),
		ZoneContext: &pb.OperationManagedZoneContext{
			NewValue: updated,
			OldValue: &existing,
		},
	}

	lroPrefix := "projects/" + req.GetProject() + "/managedZones/" + req.GetManagedZone().GetName() + "/"
	return s.operations.StartLRO(ctx, lroPrefix, op, func() (proto.Message, error) {
		return updated, nil
	})
}

func (s *managedZonesService) DeleteManagedZone(ctx context.Context, req *pb.DeleteManagedZoneRequest) (*emptypb.Empty, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/" + req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	var existing pb.ManagedZone

	if err := s.storage.Delete(ctx, fqn, &existing); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *managedZonesService) ListManagedZones(ctx context.Context, req *pb.ListManagedZonesRequest) (*pb.ManagedZonesListResponse, error) {
	name, err := s.parseManagedZoneName("projects/" + req.GetProject() + "/managedZones/dummy")
	if err != nil {
		return nil, err
	}

	if req.DnsName != nil || req.PageToken != nil {
		return nil, status.Errorf(codes.Unimplemented, "filtering and pagination are not implemented in mock")
	}

	prefix := strings.TrimSuffix(name.String(), "dummy")

	zoneKind := (&pb.ManagedZone{}).ProtoReflect().Descriptor()

	items := make([]*pb.ManagedZone, 0)
	if err := s.storage.List(ctx, zoneKind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
		zone := item.(*pb.ManagedZone)
		items = append(items, zone)
		return nil
	}); err != nil {
		return nil, err
	}

	response := &pb.ManagedZonesListResponse{
		ManagedZones: items,
		Kind:         PtrTo("dns#managedZonesListResponse"),
	}

	if req.MaxResults != nil {
		maxResults := req.GetMaxResults()
		if len(items) > int(maxResults) {
			response.ManagedZones = items[:maxResults]
			response.NextPageToken = PtrTo("mock-pagination-token")
		}
	}

	return response, nil
}

type managedZoneName struct {
	Project *projects.ProjectData
	Name    string
}

func (n *managedZoneName) String() string {
	return "projects/" + n.Project.ID + "/managedZones/" + n.Name
}

// parseManagedZoneName parses a string into a managedZoneName.
// The expected form is `projects/<projectID>/managedZones/<name>`.
func (s *MockService) parseManagedZoneName(name string) (*managedZoneName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "managedZones" {
		project, err := s.projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &managedZoneName{
			Project: project,
			Name:    tokens[3],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func currentUser(ctx context.Context) string {
	return "user@example.com"
}
