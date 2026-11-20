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

package mockmemorystore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type aclPolicyName struct {
	Project  *projects.ProjectData
	Location string
	Name     string
}

func (n *aclPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/aclPolicies/" + n.Name
}

func (r *instanceServer) parseAclPolicyName(name string) (*aclPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "aclPolicies" {
		project, err := r.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &aclPolicyName{
			Project:  project,
			Location: tokens[3],
			Name:     tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (r *instanceServer) GetAclPolicy(ctx context.Context, req *pb.GetAclPolicyRequest) (*pb.AclPolicy, error) {
	name, err := r.parseAclPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AclPolicy{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return proto.CloneOf(obj), nil
}

func (r *instanceServer) CreateAclPolicy(ctx context.Context, req *pb.CreateAclPolicyRequest) (*pb.AclPolicy, error) {
	reqName := fmt.Sprintf("%s/aclPolicies/%s", req.GetParent(), req.GetAclPolicyId())
	name, err := r.parseAclPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.CloneOf(req.GetAclPolicy())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.AclPolicy_ACTIVE

	if err := r.populateDefaultsForAclPolicy(name, obj); err != nil {
		return nil, err
	}

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	retObj := proto.CloneOf(obj)
	// In real GCP, createTime and updateTime are omitted in the CreateAclPolicy response.
	retObj.CreateTime = nil
	retObj.UpdateTime = nil
	return retObj, nil
}

func (r *instanceServer) populateDefaultsForAclPolicy(name *aclPolicyName, obj *pb.AclPolicy) error {
	if obj.State == pb.AclPolicy_STATE_UNSPECIFIED {
		obj.State = pb.AclPolicy_ACTIVE
	}
	return nil
}

func (r *instanceServer) UpdateAclPolicy(ctx context.Context, req *pb.UpdateAclPolicyRequest) (*longrunning.Operation, error) {
	reqName := req.GetAclPolicy().GetName()
	name, err := r.parseAclPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.AclPolicy{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask is required")
	}

	for _, path := range paths {
		switch path {
		case "rules":
			obj.Rules = req.GetAclPolicy().GetRules()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mockgcp", path)
		}
	}

	if err := r.populateDefaultsForAclPolicy(name, obj); err != nil {
		return nil, err
	}

	obj.UpdateTime = timestamppb.New(now)
	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		retObj := proto.CloneOf(obj)
		retObj.State = pb.AclPolicy_ACTIVE
		if err := r.storage.Update(ctx, fqn, retObj); err != nil {
			return nil, err
		}
		return retObj, nil
	})
}

func (r *instanceServer) DeleteAclPolicy(ctx context.Context, req *pb.DeleteAclPolicyRequest) (*longrunning.Operation, error) {
	name, err := r.parseAclPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.AclPolicy{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     fqn,
		Verb:       "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		deletedObj := &pb.AclPolicy{}
		if err := r.storage.Delete(ctx, fqn, deletedObj); err != nil {
			return nil, err
		}
		return &emptypb.Empty{}, nil
	})
}

func (r *instanceServer) ListAclPolicies(ctx context.Context, req *pb.ListAclPoliciesRequest) (*pb.ListAclPoliciesResponse, error) {
	tokens := strings.Split(req.GetParent(), "/")
	if len(tokens) != 4 || tokens[0] != "projects" || tokens[2] != "locations" {
		return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", req.GetParent())
	}
	project, err := r.Projects.GetProjectByIDOrNumber(tokens[1])
	if err != nil {
		return nil, err
	}

	location := tokens[3]
	prefix := fmt.Sprintf("projects/%s/locations/%s/aclPolicies/", project.ID, location)

	response := &pb.ListAclPoliciesResponse{}
	findKind := (&pb.AclPolicy{}).ProtoReflect().Descriptor()
	if err := r.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		item := obj.(*pb.AclPolicy)
		response.AclPolicies = append(response.AclPolicies, item)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}
