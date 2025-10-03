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

// +tool:mockgcp-support
// proto.service: google.cloud.networkmanagement.v1.ReachabilityService
// proto.message: google.cloud.networkmanagement.v1.ConnectivityTest

package mocknetworkmanagement

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkmanagement/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

type reachabilityService struct {
	*MockService
	pb.UnimplementedReachabilityServiceServer
}

func (s *reachabilityService) GetConnectivityTest(ctx context.Context, req *pb.GetConnectivityTestRequest) (*pb.ConnectivityTest, error) {
	name, err := s.parseConnectivityTestName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.ConnectivityTest{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}

		return nil, err
	}
	// TODO: Remove these lines, or add default values, this is just a workaround for the test
	if obj.GetReachabilityDetails() == nil {
		obj.ReachabilityDetails = &pb.ReachabilityDetails{
			Result: pb.ReachabilityDetails_REACHABLE,
		}
	}
	return obj, nil
}

func (s *reachabilityService) CreateConnectivityTest(ctx context.Context, req *pb.CreateConnectivityTestRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/connectivityTests/" + req.TestId
	name, err := s.parseConnectivityTestName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.Resource).(*pb.ConnectivityTest)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.ReachabilityDetails = &pb.ReachabilityDetails{
		Result: pb.ReachabilityDetails_UNDETERMINED,
		Traces: []*pb.Trace{
			{
				EndpointInfo: &pb.EndpointInfo{
					DestinationIp:    "10.0.0.2",
					DestinationPort:  80,
					Protocol:         "TCP",
					SourceIp:         "10.0.0.1",
					SourceNetworkUri: "projects/${projectId}/global/networks/default",
					SourcePort:       55896,
				},
				ForwardTraceId: 1,
				Steps: []*pb.Step{
					{
						Description: "Initial state: packet originating from a VPC network default.",
						ProjectId:   "${projectId}",
						State:       pb.Step_START_FROM_PRIVATE_NETWORK,
						StepInfo: &pb.Step_Network{
							Network: &pb.NetworkInfo{
								DisplayName:      "default",
								MatchedIpRange:   "10.0.0.0/20",
								MatchedSubnetUri: "projects/${projectId}/regions/us-east4/subnetworks/default",
								Region:           "us-east4",
								Uri:              "projects/${projectId}/global/networks/default",
							},
						},
					},
					{
						Description: "Config checking state: verify EGRESS firewall rule.",
						StepInfo: &pb.Step_Firewall{
							Firewall: &pb.FirewallInfo{
								Action:           "ALLOW",
								Direction:        "EGRESS",
								DisplayName:      "default-allow-egress",
								FirewallRuleType: pb.FirewallInfo_IMPLIED_VPC_FIREWALL_RULE,
								NetworkUri:       "projects/${projectId}/global/networks/default",
								Priority:         65535,
							},
						},
						ProjectId: "${projectId}",
						State:     pb.Step_APPLY_EGRESS_FIREWALL_RULE,
					},
					{
						Description: "Config checking state: verify route.",
						ProjectId:   "${projectId}",
						StepInfo: &pb.Step_Route{
							Route: &pb.RouteInfo{
								DestIpRange: "10.0.0.0/20",
								DisplayName: "default-route-f17b1d9b115a2c1e",
								NetworkUri:  "projects/${projectId}/global/networks/default",
								NextHopType: pb.RouteInfo_NEXT_HOP_NETWORK,
								RouteType:   pb.RouteInfo_SUBNET,
								Uri:         "projects/${projectId}/global/routes/default-route-f17b1d9b115a2c1e",
							},
						},
						State: pb.Step_APPLY_ROUTE,
					},
					{
						StepInfo: &pb.Step_Abort{
							Abort: &pb.AbortInfo{
								Cause:       pb.AbortInfo_UNKNOWN_IP,
								IpAddress:   "10.0.0.2",
								ResourceUri: "projects/${projectId}/global/networks/default",
							},
						},
						Description: "Final state: analysis is aborted due to no endpoints with destination IP address are found in network.",
						ProjectId:   "${projectId}",
						State:       pb.Step_ABORT,
					},
				},
			},
		},
		VerifyTime: timestamppb.New(now),
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		Target:          name.String(),
		Verb:            "create",
		ApiVersion:      "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.Parent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)
		return obj, nil
	})
}

func (s *reachabilityService) UpdateConnectivityTest(ctx context.Context, req *pb.UpdateConnectivityTestRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetResource().GetName()
	name, err := s.parseConnectivityTestName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.ConnectivityTest{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetResource().GetDescription()
		case "labels":
			obj.Labels = req.GetResource().GetLabels()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.Protocol = "TCP"

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(time.Now()),
		Target:          name.String(),
		Verb:            "update",
		ApiVersion:      "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.Parent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *reachabilityService) DeleteConnectivityTest(ctx context.Context, req *pb.DeleteConnectivityTestRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseConnectivityTestName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.ConnectivityTest{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:      timestamppb.New(now),
		Target:          name.String(),
		Verb:            "delete",
		ApiVersion:      "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.Parent(), lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type connectivityTestName struct {
	Project              *projects.ProjectData
	ConnectivityTestName string
}

func (n *connectivityTestName) String() string {
	return fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", n.Project.ID, n.ConnectivityTestName)
}

func (n *connectivityTestName) Parent() string {
	return fmt.Sprintf("projects/%s/locations/global", n.Project.ID)
}

// parseNetworkName parses a string into a networkName.
// The expected form is `projects/*/locations/global/connectivityTests/*`.
func (s *MockService) parseConnectivityTestName(name string) (*connectivityTestName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "connectivityTests" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &connectivityTestName{
			Project:              project,
			ConnectivityTestName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
