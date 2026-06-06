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

package mockdns

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type changesServer struct {
	*MockService
	pb.UnimplementedChangesServerServer
}

func (s *changesServer) CreateChange(ctx context.Context, req *pb.CreateChangeRequest) (*pb.Change, error) {
	projectData, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}

	zoneName := req.GetManagedZone()
	zoneFqn := fmt.Sprintf("projects/%s/managedZones/%s", projectData.ID, zoneName)
	var zone pb.ManagedZone
	if err := s.storage.Get(ctx, zoneFqn, &zone); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "managedZone %q not found", zoneFqn)
		}
		return nil, err
	}

	change := req.GetChange()
	if change == nil {
		return nil, status.Errorf(codes.InvalidArgument, "change cannot be nil")
	}

	// Apply deletions first
	for _, rrset := range change.GetDeletions() {
		name, err := s.parseResourceRecordSetName(projectData.ID, zoneName, rrset.GetName(), rrset.GetType())
		if err != nil {
			return nil, err
		}
		fqn := name.String()
		if err := s.storage.Delete(ctx, fqn, &pb.ResourceRecordSet{}); err != nil {
			if !apierrors.IsNotFound(err) {
				return nil, err
			}
		}
	}

	// Apply additions
	for _, rrset := range change.GetAdditions() {
		name, err := s.parseResourceRecordSetName(projectData.ID, zoneName, rrset.GetName(), rrset.GetType())
		if err != nil {
			return nil, err
		}
		fqn := name.String()
		obj := proto.Clone(rrset).(*pb.ResourceRecordSet)
		obj.Kind = PtrTo("dns#resourceRecordSet")
		if obj.SignatureRrdatas == nil {
			obj.SignatureRrdatas = []string{}
		}
		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			return nil, err
		}
	}

	changeId := fmt.Sprintf("%d", time.Now().UnixNano())
	change.Id = PtrTo(changeId)
	change.StartTime = PtrTo(time.Now().Format(time.RFC3339))
	change.Status = PtrTo("done")
	change.Kind = PtrTo("dns#change")

	changeFqn := fmt.Sprintf("projects/%s/managedZones/%s/changes/%s", projectData.ID, zoneName, changeId)
	if err := s.storage.Create(ctx, changeFqn, change); err != nil {
		return nil, err
	}

	return change, nil
}

func (s *changesServer) GetChange(ctx context.Context, req *pb.GetChangeRequest) (*pb.Change, error) {
	projectData, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}
	zoneName := req.GetManagedZone()
	changeId := req.GetChangeId()
	changeFqn := fmt.Sprintf("projects/%s/managedZones/%s/changes/%s", projectData.ID, zoneName, changeId)

	obj := &pb.Change{}
	if err := s.storage.Get(ctx, changeFqn, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "change %q not found", changeId)
		}
		return nil, err
	}

	return obj, nil
}

func (s *changesServer) ListChanges(ctx context.Context, req *pb.ListChangesRequest) (*pb.ChangesListResponse, error) {
	projectData, err := s.Projects.GetProjectByIDOrNumber(req.GetProject())
	if err != nil {
		return nil, err
	}
	zoneName := req.GetManagedZone()

	prefix := fmt.Sprintf("projects/%s/managedZones/%s/changes/", projectData.ID, zoneName)
	changeKind := (&pb.Change{}).ProtoReflect().Descriptor()

	items := make([]*pb.Change, 0)
	if err := s.storage.List(ctx, changeKind, storage.ListOptions{Prefix: prefix}, func(item proto.Message) error {
		change := item.(*pb.Change)
		items = append(items, change)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ChangesListResponse{
		Changes: items,
		Kind:    PtrTo("dns#changesListResponse"),
	}, nil
}
