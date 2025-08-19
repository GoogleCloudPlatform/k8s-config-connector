// Copyright 2023 Google LLC
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

package mockbilling

import (
	"context"

	"cloud.google.com/go/iam/apiv1/iampb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/billing/v1"
)

type BillingV1 struct {
	*MockService
	pb.UnimplementedCloudBillingServer
}

func (s *BillingV1) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	response := &iampb.TestIamPermissionsResponse{}
	// HACK: assume all permissions
	for _, permission := range req.Permissions {
		response.Permissions = append(response.Permissions, permission)
	}
	return response, nil
}

func (s *BillingV1) GetProjectBillingInfo(ctx context.Context, req *pb.GetProjectBillingInfoRequest) (*pb.ProjectBillingInfo, error) {
	projectName, err := projects.ParseProjectName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := projectName.String() + "/billingInfo"

	obj := &pb.ProjectBillingInfo{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			// Expected if billing info has not yet been set
			obj.Name = fqn
			obj.BillingEnabled = false
			obj.ProjectId = projectName.ProjectID
		} else {
			return nil, err
		}
	}

	return obj, nil
}

func (s *BillingV1) UpdateProjectBillingInfo(ctx context.Context, req *pb.UpdateProjectBillingInfoRequest) (*pb.ProjectBillingInfo, error) {
	projectName, err := projects.ParseProjectName(req.GetName())
	if err != nil {
		return nil, err
	}

	billingAccountName, err := s.parseBillingAccountName(req.GetProjectBillingInfo().GetBillingAccountName())
	if err != nil {
		return nil, err
	}

	fqn := projectName.String() + "/billingInfo"

	obj := proto.Clone(req.GetProjectBillingInfo()).(*pb.ProjectBillingInfo)
	obj.BillingAccountName = billingAccountName.String()
	obj.Name = fqn
	obj.BillingEnabled = true
	obj.ProjectId = projectName.ProjectID

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BillingV1) ListBillingAccounts(ctx context.Context, req *pb.ListBillingAccountsRequest) (*pb.ListBillingAccountsResponse, error) {
	// For now, return a dummy billing account.
	// In a more complete mock, this would retrieve from storage.
	dummyAccount := &pb.BillingAccount{
		Name:         "billingAccounts/000000-123456-000000",
		Open:         true,
		DisplayName:  "Mock Billing Account",
		CurrencyCode: "USD",
		Parent:       "organizations/12345678",
	}

	response := &pb.ListBillingAccountsResponse{
		BillingAccounts: []*pb.BillingAccount{dummyAccount},
	}
	return response, nil
}
