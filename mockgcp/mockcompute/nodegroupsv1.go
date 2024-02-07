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

package mockcompute

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type NodeGroupsV1 struct {
	*MockService
	pb.UnimplementedNodeGroupsServer
}

func (s *NodeGroupsV1) Get(ctx context.Context, req *pb.GetNodeGroupRequest) (*pb.NodeGroup, error) {
	name, err := s.newNodeGroupName(req.GetProject(), req.GetZone(), req.GetNodeGroup())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeGroup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *NodeGroupsV1) Insert(ctx context.Context, req *pb.InsertNodeGroupRequest) (*pb.Operation, error) {
	name, err := s.newNodeGroupName(req.GetProject(), req.GetZone(), req.GetNodeGroupResource().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	id := s.generateID()

	obj := proto.Clone(req.GetNodeGroupResource()).(*pb.NodeGroup)
	obj.SelfLink = PtrTo("https://compute.googleapis.com/compute/v1/" + name.String())
	obj.CreationTimestamp = PtrTo(s.nowString())
	obj.Id = &id
	obj.Kind = PtrTo("compute#nodegroup")

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating node group: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *NodeGroupsV1) Patch(ctx context.Context, req *pb.PatchNodeGroupRequest) (*pb.Operation, error) {
	name, err := s.newNodeGroupName(req.GetProject(), req.GetZone(), req.GetNodeGroup())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.NodeGroup{}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error updating node group: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *NodeGroupsV1) Delete(ctx context.Context, req *pb.DeleteNodeGroupRequest) (*pb.Operation, error) {
	name, err := s.newNodeGroupName(req.GetProject(), req.GetZone(), req.GetNodeGroup())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeGroup{}
	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return s.newLRO(ctx, name.Project.ID)
}

func (s *NodeGroupsV1) SetNodeTemplate(ctx context.Context, req *pb.SetNodeTemplateNodeGroupRequest) (*pb.Operation, error) {
	name, err := s.newNodeTemplateNodeGroupName(req.GetProject(), req.GetZone(), req.GetNodeGroup(), req.GetNodeGroupsSetNodeTemplateRequestResource().GetNodeTemplate())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.NodeGroup{}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, status.Errorf(codes.Internal, "error set node template for node group: %v", err)
	}

	return s.newLRO(ctx, name.Project.ID)
}

type nodeGroupName struct {
	Project *projects.ProjectData
	Zone    string
	Name    string
}

type nodeTemplateNodeGroupName struct {
	Project          *projects.ProjectData
	Zone             string
	NodeGroupName    string
	NodeTemplateName string
}

func (n *nodeGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/nodeGroups/" + n.Name
}

func (n *nodeTemplateNodeGroupName) String() string {
	return "projects/" + n.Project.ID + "/zones/" + n.Zone + "/nodeGroups/" + n.NodeGroupName + "/setNodeTemplate/" + n.NodeTemplateName
}

// newNodeGroupName builds a normalized nodeGroupName from the constituent parts.
// The expected form is `projects/{project}/zones/{zone}/nodeGroups/{nodeGroup}`.
func (s *MockService) newNodeGroupName(project string, zone string, name string) (*nodeGroupName, error) {
	projectObj, err := s.projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &nodeGroupName{
		Project: projectObj,
		Zone:    zone,
		Name:    name,
	}, nil
}

// newNodeGroupName builds a normalized nodeGroupName from the constituent parts.
// The expected form is `projects/{project}/zones/{zone}/nodeGroups/{nodeGroup}/setNodeTemplate/{nodeTemplate}`.
func (s *MockService) newNodeTemplateNodeGroupName(project string, zone string, nodeGroupName string, nodeTemplateName string) (*nodeTemplateNodeGroupName, error) {
	projectObj, err := s.projects.GetProjectByID(project)
	if err != nil {
		return nil, err
	}

	return &nodeTemplateNodeGroupName{
		Project:          projectObj,
		Zone:             zone,
		NodeGroupName:    nodeGroupName,
		NodeTemplateName: nodeTemplateName,
	}, nil
}
