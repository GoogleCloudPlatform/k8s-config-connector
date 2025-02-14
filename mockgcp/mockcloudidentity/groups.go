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

package mockcloudidentity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/klog/v2"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/apps/cloudidentity/groups/v1beta1"
)

type groupsServer struct {
	*MockService
	pb.UnimplementedGroupsServerServer
}

func (s *groupsServer) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.Group, error) {
	name, err := s.parseGroupName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.PermissionDenied, "Error(2017): Permission denied for group resource '%s' (or it may not exist).", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *groupsServer) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("groups/%x", time.Now().UnixMilli())

	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := timestamppb.Now()

	obj := proto.Clone(req.Group).(*pb.Group)
	obj.Name = PtrTo(fmt.Sprintf("groups/%s", name.Name))
	obj.CreateTime = now
	obj.UpdateTime = now

	obj.AdditionalGroupKeys = append(obj.AdditionalGroupKeys, obj.GroupKey)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Terraform client is _very_ fussy about format here
	// (to the point of being arguably broken)
	// return s.operations.DoneLRO(ctx, "", lro, obj)

	// additionalGroupKeys are not populated in the LRO
	retObj := proto.Clone(obj).(*pb.Group)
	retObj.AdditionalGroupKeys = nil
	go func() {
		time.Sleep(time.Second)
		if err := s.addAdditionalGroupKeys(context.Background(), name); err != nil {
			klog.Fatalf("error adding additionalGroupKeys: %v", err)
		}
	}()
	return buildLRO(retObj)
}

func (s *groupsServer) addAdditionalGroupKeys(ctx context.Context, name *groupName) error {
	fqn := name.String()

	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return err
	}

	addAdditionalGroup(obj, obj.GroupKey)

	// Add a test-google-a.com additionalGroupKey which is auto-added by the service.
	for _, groupKey := range obj.AdditionalGroupKeys {
		id := groupKey.GetId()
		if strings.HasSuffix(id, ".test-google-a.com") {
			continue
		}
		newGroup := &pb.EntityKey{
			Id: proto.String(fmt.Sprintf("%s.test-google-a.com", id)),
		}
		addAdditionalGroup(obj, newGroup)
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return err
	}

	return nil
}

// addAdditionalGroup adds a group to additionalGroups, unless it is already found.
func addAdditionalGroup(obj *pb.Group, entityKey *pb.EntityKey) {
	for _, existing := range obj.AdditionalGroupKeys {
		if proto.Equal(existing, entityKey) {
			return
		}
	}
	obj.AdditionalGroupKeys = append(obj.AdditionalGroupKeys, entityKey)
}

func (s *groupsServer) PatchGroup(ctx context.Context, req *pb.PatchGroupRequest) (*longrunning.Operation, error) {
	reqName := req.GetName()

	name, err := s.parseGroupName(reqName)
	if err != nil {
		return nil, err
	}

	now := timestamppb.Now()

	fqn := name.String()
	obj := &pb.Group{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range strings.Split(req.GetUpdateMask(), ",") {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetGroup().DisplayName
		case "description":
			obj.Description = req.GetGroup().Description
		case "labels":
			obj.Labels = req.GetGroup().Labels
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = now

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// additionalGroupKeys and groupKey are not populated in the LRO response
	retObj := proto.Clone(obj).(*pb.Group)
	retObj.AdditionalGroupKeys = nil
	retObj.GroupKey = &pb.EntityKey{}
	return buildLRO(retObj)
}

func (s *groupsServer) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*longrunning.Operation, error) {
	name, err := s.parseGroupName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Group{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	// Returns a non-standard LRO
	lro := &longrunning.Operation{}
	lro.Done = true
	return lro, nil
}

type groupName struct {
	Name string
}

func (n *groupName) String() string {
	return fmt.Sprintf("groups/%s", n.Name)
}

func (s *MockService) parseGroupName(name string) (*groupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "groups" {
		name := &groupName{
			Name: tokens[1],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
