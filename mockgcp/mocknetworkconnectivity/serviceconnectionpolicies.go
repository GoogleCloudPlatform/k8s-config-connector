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

package mocknetworkconnectivity

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
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
)

type serviceConnectionPolicies struct {
	*MockService
	pb.UnimplementedProjectsLocationsServiceConnectionPoliciesServerServer
}

func (r *serviceConnectionPolicies) GetProjectsLocationsServiceConnectionPolicy(ctx context.Context, req *pb.GetProjectsLocationsServiceConnectionPolicyRequest) (*pb.ServiceConnectionPolicy, error) {
	name, err := r.parseServiceConnectionPolicyName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ServiceConnectionPolicy{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (r *serviceConnectionPolicies) CreateProjectsLocationsServiceConnectionPolicy(ctx context.Context, req *pb.CreateProjectsLocationsServiceConnectionPolicyRequest) (*longrunning.Operation, error) {
	reqName := fmt.Sprintf("%s/serviceConnectionPolicies/%s", req.GetParent(), req.GetServiceConnectionPolicyId())
	name, err := r.parseServiceConnectionPolicyName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetProjectsLocationsServiceConnectionPolicy()).(*pb.ServiceConnectionPolicy)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	r.populateDefaultsForServiceConnectionPolicy(name, obj)

	obj.Etag = computeEtag(obj)

	if err := r.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()

		if err := r.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return redactedForLRO(obj), nil
	})
}

// redactedForLRO returns a version of the ServiceConnectionPolicy with many fields not set,
// which is what the LRO returns
func redactedForLRO(obj *pb.ServiceConnectionPolicy) *pb.ServiceConnectionPolicy {
	retObj := proto.Clone(obj).(*pb.ServiceConnectionPolicy)
	retObj.Description = ""
	retObj.Infrastructure = ""
	retObj.PscConfig = nil
	retObj.ServiceClass = ""
	retObj.Network = ""

	return retObj
}

func (r *serviceConnectionPolicies) populateDefaultsForServiceConnectionPolicy(name *serviceConnectionPolicyName, obj *pb.ServiceConnectionPolicy) {
	if obj.Infrastructure == "" {
		obj.Infrastructure = "PSC"
	}
}

func (r *serviceConnectionPolicies) PatchProjectsLocationsServiceConnectionPolicy(ctx context.Context, req *pb.PatchProjectsLocationsServiceConnectionPolicyRequest) (*longrunning.Operation, error) {
	log := klog.FromContext(ctx)

	reqName := req.GetName()

	name, err := r.parseServiceConnectionPolicyName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := &pb.ServiceConnectionPolicy{}
	if err := r.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UpdateTime = timestamppb.New(now)

	if req.GetUpdateMask() != "" {
		paths := strings.Split(req.GetUpdateMask(), ",")

		patch := req.GetProjectsLocationsServiceConnectionPolicy()
		// TODO: Some sort of helper for fieldmask?
		for _, path := range paths {
			switch path {
			case "psc_config":
				obj.PscConfig = patch.GetPscConfig()

			default:
				log.Info("unsupported update_mask", "req", req)
				return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not supported by mock", path)
			}
		}
	} else {
		// If update_mask is not specified, all fields are overwritten
		patch := req.GetProjectsLocationsServiceConnectionPolicy()
		obj.PscConfig = patch.GetPscConfig()
	}

	obj.Etag = computeEtag(obj)

	if err := r.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "update",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return redactedForLRO(obj), nil
	})
}

func (r *serviceConnectionPolicies) DeleteProjectsLocationsServiceConnectionPolicy(ctx context.Context, req *pb.DeleteProjectsLocationsServiceConnectionPolicyRequest) (*longrunning.Operation, error) {
	name, err := r.parseServiceConnectionPolicyName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	oldObj := &pb.ServiceConnectionPolicy{}
	if err := r.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		RequestedCancellation: false,
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "delete",
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return r.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type serviceConnectionPolicyName struct {
	Project                     *projects.ProjectData
	Location                    string
	ServiceConnectionPolicyName string
}

func (n *serviceConnectionPolicyName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/serviceConnectionPolicies/" + n.ServiceConnectionPolicyName
}

// parseServiceConnectionPolicyName parses a string into an serviceConnectionPolicyName.
// The expected form is `projects/*/locations/*/serviceConnectionPolicies/*`.
func (r *serviceConnectionPolicies) parseServiceConnectionPolicyName(name string) (*serviceConnectionPolicyName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "serviceConnectionPolicies" {
		project, err := r.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &serviceConnectionPolicyName{
			Project:                     project,
			Location:                    tokens[3],
			ServiceConnectionPolicyName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
