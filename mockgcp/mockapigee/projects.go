// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockapigee

import (
	"context"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
)

type projectsServer struct {
	*MockService
	pb.UnimplementedProjectsServerServer
}

func (s *projectsServer) ProvisionOrganizationProject(ctx context.Context, req *pb.ProvisionOrganizationProjectRequest) (*longrunningpb.Operation, error) {
	var name *OrganizationName

	projectID := ""

	// parent := req.GetParent()
	tokens := strings.Split(req.GetName(), "/")
	if len(tokens) == 2 && tokens[0] == "projects" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		projectID = project.ID

		// Name is same as project ID
		name = &OrganizationName{
			ID: project.ID,
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", req.GetName())
	}

	now := time.Now()

	fqn := name.String()

	obj := &pb.GoogleCloudApigeeV1Organization{}
	obj.Name = name.ID
	obj.CreatedAt = now.UnixMilli()
	obj.LastModifiedAt = now.UnixMilli()
	obj.ProjectId = projectID
	obj.State = "ACTIVE"

	obj.BillingType = "EVALUATION"
	obj.SubscriptionType = "TRIAL"

	expiresAt := now.Add(60 * 24 * time.Hour)
	obj.ExpiresAt = expiresAt.UnixMilli()

	obj.CaCertificate = []byte("LS0t...")

	if obj.AddonsConfig != nil {
		if obj.AddonsConfig.MonetizationConfig != nil {
			if !obj.AddonsConfig.MonetizationConfig.GetEnabled() {
				obj.AddonsConfig.MonetizationConfig = nil
			}
		}
	}
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := "organizations/" + name.ID
	opMetadata := &pb.GoogleCloudApigeeV1OperationMetadata{
		OperationType:      "INSERT",
		State:              "IN_PROGRESS",
		TargetResourceName: fqn,
	}
	return s.operations.StartLRO(ctx, prefix, opMetadata, func() (proto.Message, error) {
		opMetadata.State = "FINISHED"
		opMetadata.Progress = &pb.GoogleCloudApigeeV1OperationMetadataProgress{
			Description: "Succeeded",
			PercentDone: 100,
		}
		return &emptypb.Empty{}, nil
	})
}
