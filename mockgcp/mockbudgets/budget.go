// Copyright 2022 Google LLC
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

package mockbudgets

import (
	"context"
	"crypto/md5"
	"math/rand"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"k8s.io/klog"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/billing/budgets/v1"
)

// Reference:
// https://cloud.google.com/billing/docs/reference/budget/rest
// https://billingbudgets.googleapis.com/$discovery/rest?version=v1

type BudgetV1 struct {
	*MockService
	pb.UnimplementedBudgetServiceServer
}

func (s *BudgetV1) GetBudget(ctx context.Context, req *pb.GetBudgetRequest) (*pb.Budget, error) {
	name, err := s.parseBudgetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Budget{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "01234567890"

func GenerateBudgetId() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	n := make([]byte, 6)
	for i := range n {
		n[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b) + "-" + string(n)
}

// https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets/create
func (s *BudgetV1) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.Budget, error) {

	fqn := req.Parent + "/budgets/" + GenerateBudgetId()
	budget := proto.Clone(req.GetBudget()).(*pb.Budget)
	budget.Name = fqn
	budget.Etag = string(computeEtag(budget))
	if err := s.storage.Create(ctx, fqn, budget); err != nil {
		return nil, err
	}
	return budget, nil
}

func (s *BudgetV1) PatchBudget(ctx context.Context, req *pb.UpdateBudgetRequest) (*pb.Budget, error) {
	reqName := req.GetBudget().GetName()

	name, err := s.parseBudgetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Budget{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// Required. A list of fields to be updated in this request.
	paths := req.GetUpdateMask().GetPaths()

	// TODO: Some sort of helper for fieldmask?
	for _, path := range paths {
		switch path {
		case "displayName":
			obj.DisplayName = req.GetBudget().GetDisplayName()
		// TODO add more fields ?
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BudgetV1) DeleteBudget(ctx context.Context, req *pb.DeleteBudgetRequest) (*emptypb.Empty, error) {
	name, err := s.parseBudgetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	oldObj := &pb.Budget{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type budgetName struct {
	BillingAccountId string
	BudgetId         string
}

func (n *budgetName) String() string {
	return "billingAccounts/" + n.BillingAccountId + "/budgets/" + n.BudgetId
}

// parseBudgetName parses a string into a budgetName
// The expected form is billingAccounts/{billingAccountId}/budgets/{budgetId}
func (s *MockService) parseBudgetName(name string) (*budgetName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "billingAccounts" && tokens[2] == "budgets" {
		name := &budgetName{
			BillingAccountId: tokens[1],
			BudgetId:         tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func computeEtag(obj proto.Message) []byte {
	// TODO: Do we risk exposing internal fields?  Doesn't matter on a mock, I guess
	b, err := proto.Marshal(obj)
	if err != nil {
		klog.Fatalf("failed to marshal proto object: %v", err)
	}
	hash := md5.Sum(b)
	return hash[:]
}
