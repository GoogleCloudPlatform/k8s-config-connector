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
// proto.message: google.bigtable.admin.v2.LogicalView

package mockanalytics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Note: we use the "real" proto (not mockgcp), because the client uses GRPC.
	pb "cloud.google.com/go/analytics/admin/apiv1alpha/adminpb"
)

func (s *analyticsAdminServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	name, err := s.parseAccountName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Account{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "%v not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *analyticsAdminServer) ListAccounts(ctx context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	response := &pb.ListAccountsResponse{}
	findKind := (&pb.Account{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{}, func(obj proto.Message) error {
		account := obj.(*pb.Account)
		response.Accounts = append(response.Accounts, account)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *analyticsAdminServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*emptypb.Empty, error) {
	name, err := s.parseAccountName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Account{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *analyticsAdminServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.Account, error) {
	name, err := s.parseAccountName(req.Account.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	existing := &pb.Account{}
	if err := s.storage.Get(ctx, fqn, existing); err != nil {
		return nil, err
	}

	updated := ProtoClone(existing)

	// Required. The set of fields to update.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "display_name":
			updated.DisplayName = req.GetAccount().GetDisplayName()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *analyticsAdminServer) ProvisionAccountTicket(ctx context.Context, req *pb.ProvisionAccountTicketRequest) (*pb.ProvisionAccountTicketResponse, error) {
	// Service-generated resource ID.
	fqn := "123456"

	obj := ProtoClone(req.Account)
	obj.Name = fmt.Sprintf("accounts/%s", fqn)
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &pb.ProvisionAccountTicketResponse{AccountTicketId: "ASDFGHJKL123456"}, nil
}

//// Returns all accounts accessible by the caller.
////
//// Note that these accounts might not currently have GA properties.
//// Soft-deleted (ie: "trashed") accounts are excluded by default.
//// Returns an empty list if no relevant accounts are found.
//rpc ListAccounts(ListAccountsRequest) returns (ListAccountsResponse) {
//option (google.api.http) = {
//get: "/v1beta/accounts"
//};
//}

type accountName struct {
	accountID string
}

func (n *accountName) String() string {
	return fmt.Sprintf("accounts/%s", n.accountID)
}

// parseLogicalViewName parses a string into a logicalViewName.
// The expected form is `projects/*/instances/*/logicalViews/*`.
func (s *analyticsAdminServer) parseAccountName(name string) (*accountName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 2 && tokens[0] == "accounts" {

		name := &accountName{
			accountID: tokens[1],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func ProtoClone[T proto.Message](obj T) T {
	return proto.Clone(obj).(T)
}
