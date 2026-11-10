// Copyright 2026 High Five Labs
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

// projects.setDefaultNetworkTier, after setCommonInstanceMetadata in
// projectsv1.go. The tier must be PREMIUM or STANDARD; the project's
// defaultNetworkTier (PREMIUM in a new project) is replaced.

package mockcompute

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

func (s *ProjectsV1) SetDefaultNetworkTier(ctx context.Context, req *pb.SetDefaultNetworkTierProjectRequest) (*pb.Operation, error) {
	tier := req.GetProjectsSetDefaultNetworkTierRequestResource().GetNetworkTier()
	if tier != "PREMIUM" && tier != "STANDARD" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid value for field 'networkTier': '%s'.", tier)
	}
	obj, err := s.Get(ctx, &pb.GetProjectRequest{Project: req.GetProject()})
	if err != nil {
		return nil, err
	}
	obj.DefaultNetworkTier = PtrTo(tier)
	fqn := "projects/" + obj.GetName()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: PtrTo("compute.projects.setDefaultNetworkTier"),
		TargetLink:    PtrTo(fmt.Sprintf("https://compute.googleapis.com/compute/v1/projects/%s", obj.GetName())),
		TargetId:      obj.Id,
		User:          PtrTo("user@example.com"),
	}
	return s.startGlobalLRO(ctx, obj.GetName(), op, func() (proto.Message, error) {
		return obj, nil
	})
}
