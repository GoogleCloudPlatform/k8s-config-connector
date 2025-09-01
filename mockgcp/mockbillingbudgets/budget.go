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
// proto.service: google.cloud.billing.budgets.v1.BudgetService
// proto.message: google.cloud.billing.budgets.v1.Budget

package mockbillingbudgets

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/billing/budgets/apiv1beta1/budgetspb"
)

func (s *BudgetServiceServer) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.Budget, error) {
	budgetID := fmt.Sprintf("%x", time.Now().UnixNano())
	reqName := fmt.Sprintf("%s/budgets/%s", req.GetParent(), budgetID)
	name, err := s.parseBudgetName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetBudget()).(*pb.Budget)
	obj.Name = fqn

	s.populateDefaultsForBudget(obj)

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BudgetServiceServer) populateDefaultsForBudget(obj *pb.Budget) {
	if obj.BudgetFilter == nil {
		obj.BudgetFilter = &pb.Filter{}
	}
	if obj.BudgetFilter.UsagePeriod == nil {
		obj.BudgetFilter.UsagePeriod = &pb.Filter_CalendarPeriod{
			CalendarPeriod: pb.CalendarPeriod_MONTH,
		}
	}
	if obj.BudgetFilter.CreditTypesTreatment == pb.Filter_CREDIT_TYPES_TREATMENT_UNSPECIFIED {
		obj.BudgetFilter.CreditTypesTreatment = pb.Filter_INCLUDE_ALL_CREDITS
	}
	for _, thresholdRule := range obj.ThresholdRules {
		if thresholdRule.SpendBasis == pb.ThresholdRule_BASIS_UNSPECIFIED {
			thresholdRule.SpendBasis = pb.ThresholdRule_CURRENT_SPEND
		}
	}
}

func (s *BudgetServiceServer) GetBudget(ctx context.Context, req *pb.GetBudgetRequest) (*pb.Budget, error) {
	name, err := s.parseBudgetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Budget{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BudgetServiceServer) ListBudgets(ctx context.Context, req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
	name, err := s.parseBudgetName(req.Parent + "/budgets/dummy")
	if err != nil {
		return nil, err
	}

	response := &pb.ListBudgetsResponse{}

	findPrefix := fmt.Sprintf("billingAccounts/%s/budgets/", name.BillingAccount)

	metadataStoreKind := (&pb.Budget{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, metadataStoreKind, storage.ListOptions{}, func(obj proto.Message) error {
		budget := obj.(*pb.Budget)
		if strings.HasPrefix(budget.GetName(), findPrefix) {
			response.Budgets = append(response.Budgets, budget)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *BudgetServiceServer) UpdateBudget(ctx context.Context, req *pb.UpdateBudgetRequest) (*pb.Budget, error) {
	name, err := s.parseBudgetName(req.Budget.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Budget{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		default:
			return nil, fmt.Errorf("unhandled path in mockgcp UpdateBudget: %w", err)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *BudgetServiceServer) DeleteBudget(ctx context.Context, req *pb.DeleteBudgetRequest) (*emptypb.Empty, error) {
	name, err := s.parseBudgetName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Budget{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type budgetName struct {
	BillingAccount string
	Budget         string
}

func (b *budgetName) String() string {
	return fmt.Sprintf("billingAccounts/%s/budgets/%s", b.BillingAccount, b.Budget)
}

func (s *MockService) parseBudgetName(name string) (*budgetName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 4 && tokens[0] == "billingAccounts" && tokens[2] == "budgets" {
		budget := &budgetName{
			BillingAccount: tokens[1],
			Budget:         tokens[3],
		}

		return budget, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
