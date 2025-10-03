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
// proto.service: google.bigtable.admin.v2.BigtableInstanceAdmin
// proto.message: google.bigtable.admin.v2.AppProfile

package mockbigtable

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/apimachinery/pkg/util/sets"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
)

func (s *instanceAdminServer) GetAppProfile(ctx context.Context, req *pb.GetAppProfileRequest) (*pb.AppProfile, error) {
	name, err := s.parseAppProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.AppProfile{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%v not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) CreateAppProfile(ctx context.Context, req *pb.CreateAppProfileRequest) (*pb.AppProfile, error) {
	reqName := req.Parent + "/appProfiles/" + req.AppProfileId
	name, err := s.parseAppProfileName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.AppProfile)
	obj.Name = fqn
	if obj.Isolation == nil {
		obj.Isolation = &pb.AppProfile_StandardIsolation_{
			StandardIsolation: &pb.AppProfile_StandardIsolation{
				Priority: pb.AppProfile_PRIORITY_HIGH,
			},
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *instanceAdminServer) UpdateAppProfile(ctx context.Context, req *pb.UpdateAppProfileRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseAppProfileName(req.GetAppProfile().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.AppProfile{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	isAsync := false

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			updated.Description = req.GetAppProfile().GetDescription()
		case "multiClusterRoutingUseAny", "multi_cluster_routing_use_any":
			updated.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{
				MultiClusterRoutingUseAny: req.GetAppProfile().GetMultiClusterRoutingUseAny(),
			}
		case "multiClusterRoutingUseAny.clusterIds", "multi_cluster_routing_use_any.cluster_ids":
			updated.RoutingPolicy = &pb.AppProfile_MultiClusterRoutingUseAny_{
				MultiClusterRoutingUseAny: req.GetAppProfile().GetMultiClusterRoutingUseAny(),
			}
		case "singleClusterRouting", "single_cluster_routing":
			updated.RoutingPolicy = &pb.AppProfile_SingleClusterRouting_{
				SingleClusterRouting: req.GetAppProfile().GetSingleClusterRouting(),
			}
		case "singleClusterRouting.clusterId", "single_cluster_routing.cluster_id":
			isAsync = true
			if updated.GetSingleClusterRouting() == nil {
				updated.RoutingPolicy = &pb.AppProfile_SingleClusterRouting_{
					SingleClusterRouting: &pb.AppProfile_SingleClusterRouting{},
				}
			}
			updated.GetSingleClusterRouting().ClusterId = req.GetAppProfile().GetSingleClusterRouting().GetClusterId()
		case "standardIsolation", "standard_isolation":
			updated.Isolation = &pb.AppProfile_StandardIsolation_{
				StandardIsolation: req.GetAppProfile().GetStandardIsolation(),
			}
		case "dataBoostIsolationReadOnly", "data_boost_isolation_read_only", "dataBoostIsolationReadOnly.computeBillingOwner":
			updated.Isolation = &pb.AppProfile_DataBoostIsolationReadOnly_{
				DataBoostIsolationReadOnly: req.GetAppProfile().GetDataBoostIsolationReadOnly(),
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	metadata := &pb.UpdateAppProfileMetadata{}
	zone := "us-central1-a" // TODO
	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), zone)

	lroRet := ProtoClone(updated)
	updatePaths := sets.New(req.GetUpdateMask().GetPaths()...)
	// Only return in LRO whatever has actually been updated/changed.
	if !updatePaths.Has("standard_isolation") && !updatePaths.Has("standardIsolation") && !updatePaths.Has("dataBoostIsolationReadOnly") && !updatePaths.Has("data_boost_isolation_read_only") {
		lroRet.Isolation = nil
	}
	if !updatePaths.Has("description") {
		lroRet.Description = ""
	}
	if !updatePaths.Has("single_cluster_routing") && !updatePaths.Has("singleClusterRouting") && !updatePaths.Has("single_cluster_routing.cluster_id") && !updatePaths.Has("singleClusterRouting.clusterId") && !updatePaths.Has("multi_cluster_routing_use_any") && !updatePaths.Has("multiClusterRoutingUseAny") {
		lroRet.RoutingPolicy = nil
	}

	if isAsync {
		return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
			return lroRet, nil
		})
	} else {
		return s.operations.DoneLRO(ctx, prefix, metadata, lroRet)
	}
}

func (s *instanceAdminServer) ListAppProfiles(ctx context.Context, req *pb.ListAppProfilesRequest) (*pb.ListAppProfilesResponse, error) {
	instanceName, err := s.parseInstanceName(req.GetParent())
	if err != nil {
		return nil, err
	}

	appProfile, err := s.listAppProfilesForInstance(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	response := &pb.ListAppProfilesResponse{}
	response.AppProfiles = appProfile

	return response, nil
}

func (s *instanceAdminServer) listAppProfilesForInstance(ctx context.Context, instanceName *instanceName) ([]*pb.AppProfile, error) {
	if instanceName.InstanceName == "-" {
		return nil, fmt.Errorf("mock does not implement ListAppProfiles for wildcard instances")
	}

	var response []*pb.AppProfile

	findKind := (&pb.AppProfile{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{
		Prefix: instanceName.String() + "/appProfiles/",
	}, func(obj proto.Message) error {
		appProfile := obj.(*pb.AppProfile)
		response = append(response, appProfile)
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Name < response[j].Name
	})

	return response, nil
}

func (s *instanceAdminServer) DeleteAppProfile(ctx context.Context, req *pb.DeleteAppProfileRequest) (*emptypb.Empty, error) {
	name, err := s.parseAppProfileName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.AppProfile{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type appProfileName struct {
	instanceName
	AppProfile string
}

func (n *appProfileName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/appProfiles/%s", n.Project.ID, n.InstanceName, n.AppProfile)
}

// parseAppProfileName parses a string into a appProfileName.
// The expected form is `projects/*/instances/*/appProfiles/*`.
func (s *instanceAdminServer) parseAppProfileName(name string) (*appProfileName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "appProfiles" {
		instanceName, err := s.parseInstanceName(strings.Join(tokens[0:4], "/"))
		if err != nil {
			return nil, err
		}

		name := &appProfileName{
			instanceName: *instanceName,
			AppProfile:   tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
