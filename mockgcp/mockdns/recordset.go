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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type resourceRecordSetsService struct {
	*MockService
	pb.UnimplementedResourceRecordSetsServerServer
}

func (s *resourceRecordSetsService) GetResourceRecordSet(ctx context.Context, req *pb.GetResourceRecordSetRequest) (*pb.ResourceRecordSet, error) {
	name, err := s.parseResourceRecordSetName(req.GetProject(), req.GetManagedZone(), req.GetName(), req.GetType())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ResourceRecordSet{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "resourceRecordSet %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *resourceRecordSetsService) CreateResourceRecordSet(ctx context.Context, req *pb.CreateResourceRecordSetRequest) (*pb.ResourceRecordSet, error) {
	name, err := s.parseResourceRecordSetName(req.GetProject(), req.GetManagedZone(), req.GetResourceRecordSet().GetName(), req.GetResourceRecordSet().GetType())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.ResourceRecordSet).(*pb.ResourceRecordSet)
	obj.Kind = PtrTo("dns#resourceRecordSet")

	if obj.SignatureRrdatas == nil {
		obj.SignatureRrdatas = []string{}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *resourceRecordSetsService) PatchResourceRecordSet(ctx context.Context, req *pb.PatchResourceRecordSetRequest) (*pb.ResourceRecordSet, error) {
	name, err := s.parseResourceRecordSetName(req.GetProject(), req.GetManagedZone(), req.GetName(), req.GetType())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	var existing pb.ResourceRecordSet
	if err := s.storage.Get(ctx, fqn, &existing); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "resourceRecordSet %q not found", fqn)
		}
		return nil, err
	}

	updated := proto.Clone(&existing).(*pb.ResourceRecordSet)
	if req.GetResourceRecordSet().GetTtl() != 0 {
		updated.Ttl = req.GetResourceRecordSet().Ttl
	}
	if len(req.GetResourceRecordSet().GetRrdatas()) > 0 {
		updated.Rrdatas = req.GetResourceRecordSet().GetRrdatas()
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s *resourceRecordSetsService) DeleteResourceRecordSet(ctx context.Context, req *pb.DeleteResourceRecordSetRequest) (*pb.ResourceRecordSetsDeleteResponse, error) {
	name, err := s.parseResourceRecordSetName(req.GetProject(), req.GetManagedZone(), req.GetName(), req.GetType())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.ResourceRecordSet{}); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "resourceRecordSet %q not found", fqn)
		}
		return nil, err
	}

	resp := &pb.ResourceRecordSetsDeleteResponse{}

	return resp, nil
}

func (s *resourceRecordSetsService) ListResourceRecordSets(ctx context.Context, req *pb.ListResourceRecordSetsRequest) (*pb.ResourceRecordSetsListResponse, error) {
	name, err := s.parseResourceRecordSetName(req.GetProject(), req.GetManagedZone(), "dummyname", "dummytype")
	if err != nil {
		return nil, err
	}

	if req.PageToken != nil {
		return nil, status.Errorf(codes.Unimplemented, "pagination not implemented in mock")
	}

	prefix := strings.TrimSuffix(name.String(), "/rrsets/dummyname/dummytype") + "/rrsets/"

	recordSetKind := (&pb.ResourceRecordSet{}).ProtoReflect().Descriptor()

	items := make([]*pb.ResourceRecordSet, 0)
	if err := s.storage.List(ctx, recordSetKind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
		recordSet := item.(*pb.ResourceRecordSet)
		if req.Name != nil && ValueOf(recordSet.Name) != req.GetName() {
			return nil
		}
		if req.Type != nil && ValueOf(recordSet.Type) != req.GetType() {
			return nil
		}
		items = append(items, recordSet)
		return nil
	}); err != nil {
		return nil, err
	}

	response := &pb.ResourceRecordSetsListResponse{
		Rrsets: items,
		Kind:   PtrTo("dns#resourceRecordSetsListResponse"),
	}

	if req.MaxResults != nil {
		maxResults := req.GetMaxResults()
		if len(items) > int(maxResults) {
			response.Rrsets = items[:maxResults]
			response.NextPageToken = PtrTo("mock-pagination-token")
		}
	}

	return response, nil
}

type resourceRecordSetName struct {
	Project     *projects.ProjectData
	ManagedZone string
	Name        string
	Type        string
}

func (n *resourceRecordSetName) String() string {
	return "projects/" + n.Project.ID + "/managedZones/" + n.ManagedZone + "/rrsets/" + n.Name + "/" + n.Type
}

func (s *MockService) parseResourceRecordSetName(project, managedZone, name, recordType string) (*resourceRecordSetName, error) {
	projectData, err := s.projects.GetProjectByIDOrNumber(project)
	if err != nil {
		return nil, err
	}
	return &resourceRecordSetName{
		Project:     projectData,
		ManagedZone: managedZone,
		Name:        name,
		Type:        recordType,
	}, nil
}
