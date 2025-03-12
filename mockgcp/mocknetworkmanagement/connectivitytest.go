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

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "get",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
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

	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(time.Now()),
		EndTime:      timestamppb.New(time.Now()),
		Target:       name.String(),
		Verb:         "update",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
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
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "delete",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
		return &emptypb.Empty{}, nil
	})
}

type connectivityTestName struct {
	Project         *projects.ProjectData
	ConnectivityTestName string
}

func (n *connectivityTestName) String() string {
	return fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", n.Project.ID, n.ConnectivityTestName)
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
			Project:         project,
			ConnectivityTestName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


</example>

```go
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

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "get",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
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

	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(time.Now()),
		EndTime:      timestamppb.New(time.Now()),
		Target:       name.String(),
		Verb:         "update",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
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
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "delete",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
		return &emptypb.Empty{}, nil
	})
}

type connectivityTestName struct {
	Project         *projects.ProjectData
	ConnectivityTestName string
}

func (n *connectivityTestName) String() string {
	return fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", n.Project.ID, n.ConnectivityTestName)
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
			Project:         project,
			ConnectivityTestName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


</example>

```go
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

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "get",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
		return obj, nil
	})
}
// Updates the configuration of an existing `ConnectivityTest`.
// After you update a test, the reachability analysis is performed as part
// of the long running operation, which completes when the analysis completes.
// The Reachability state in the test resource is updated with the new result.
//
// If the endpoint specifications in `ConnectivityTest` are invalid
// (for example, they contain non-existent resources in the network, or the
// user does not have read permissions to the network configurations of
// listed projects), then the reachability result returns a value of
// `UNKNOWN`.
//
// If the endpoint specifications in `ConnectivityTest` are incomplete, the
// reachability result returns a value of `AMBIGUOUS`. See the documentation
// in `ConnectivityTest` for for more details.
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

	obj.ReachabilityDetails = &pb.ReachabilityDetails{}
	obj.ReachabilityDetails.Result = pb.ReachabilityDetails_REACHABLE
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	lroMetadata := &pb.OperationMetadata{
		CreateTime:   timestamppb.New(time.Now()),
		EndTime:      timestamppb.New(time.Now()),
		Target:       name.String(),
		Verb:         "update",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
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
		CreateTime:   timestamppb.New(now),
		EndTime:      timestamppb.New(now),
		Target:       name.String(),
		Verb:         "delete",
		ApiVersion:   "v1",
		CancelRequested: false,
	}
	return s.operations.StartLRO(ctx, name.String(), lroMetadata, func() (proto.Message, error) {
		// By default, immediately finish the LRO with success.
		// (Tests can override if needed)
		return &emptypb.Empty{}, nil
	})
}

type connectivityTestName struct {
	Project         *projects.ProjectData
	ConnectivityTestName string
}

func (n *connectivityTestName) String() string {
	return fmt.Sprintf("projects/%s/locations/global/connectivityTests/%s", n.Project.ID, n.ConnectivityTestName)
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
			Project:         project,
			ConnectivityTestName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}


</example>

```go
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
// proto.message: google.cloud.networkmanagement.v1.ListConnectivityTestsResponse

package mocknetworkmanagement

import (
	"context"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkmanagement/v1"
	"google.golang.org/protobuf/proto"
)

func (s *reachabilityService) ListConnectivityTests(ctx context.Context, req *pb.ListConnectivityTestsRequest) (*pb.ListConnectivityTestsResponse, error) {
	response := &pb.ListConnectivityTestsResponse{}
	testKind := (&pb.ConnectivityTest{}).ProtoReflect().Descriptor()

	if err := s.storage.List(ctx, testKind, storage.ListOptions{Prefix: req.Parent}, func(msg proto.Message) error {
		test := msg.(*pb.ConnectivityTest)
		response.Resources = append(response.Resources, test)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}


