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
	"math/rand"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cloud.google.com/go/billing/budgets/apiv1beta1/budgetspb"
	moneypb "google.golang.org/genproto/googleapis/type/money"
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

	// We randomize the order of projects to simulate non-determinism
	// since the order is not guaranteed by the API.
	reorderSlice(obj.GetBudgetFilter().GetProjects())

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
	if req.Budget == nil {
		return nil, status.Errorf(codes.InvalidArgument, "budget is required")
	}

	name, err := s.parseBudgetName(req.Budget.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Budget{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if len(req.GetUpdateMask().GetPaths()) == 0 {
		return nil, fmt.Errorf("empty update mask in mockgcp UpdateBudget")
	}

	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "budgetFilter.creditTypes":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			obj.BudgetFilter.CreditTypes = req.Budget.GetBudgetFilter().GetCreditTypes()
		case "budgetFilter.customPeriod":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			if customPeriod := req.Budget.GetBudgetFilter().GetCustomPeriod(); customPeriod != nil {
				obj.BudgetFilter.UsagePeriod = &pb.Filter_CustomPeriod{
					CustomPeriod: customPeriod,
				}
			} else {
				obj.BudgetFilter.UsagePeriod = nil
			}

		case "budgetFilter.calendarPeriod":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			calendarPeriod := req.Budget.GetBudgetFilter().GetCalendarPeriod()
			obj.BudgetFilter.UsagePeriod = &pb.Filter_CalendarPeriod{
				CalendarPeriod: calendarPeriod,
			}

		case "budgetFilter.creditTypesTreatment":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			obj.BudgetFilter.CreditTypesTreatment = req.Budget.GetBudgetFilter().GetCreditTypesTreatment()
		case "budgetFilter.labels":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			obj.BudgetFilter.Labels = req.Budget.GetBudgetFilter().GetLabels()
		case "budgetFilter.projects":
			if obj.BudgetFilter == nil {
				obj.BudgetFilter = &pb.Filter{}
			}
			obj.BudgetFilter.Projects = req.Budget.GetBudgetFilter().GetProjects()

		case "allUpdatesRule.disableDefaultIamRecipients":
			if obj.AllUpdatesRule == nil {
				obj.AllUpdatesRule = &pb.AllUpdatesRule{}
			}
			obj.AllUpdatesRule.DisableDefaultIamRecipients = req.Budget.GetAllUpdatesRule().GetDisableDefaultIamRecipients()
		case "allUpdatesRule.monitoringNotificationChannels":
			if obj.AllUpdatesRule == nil {
				obj.AllUpdatesRule = &pb.AllUpdatesRule{}
			}
			obj.AllUpdatesRule.MonitoringNotificationChannels = req.Budget.GetAllUpdatesRule().GetMonitoringNotificationChannels()
		case "allUpdatesRule.pubsubTopic":
			if obj.AllUpdatesRule == nil {
				obj.AllUpdatesRule = &pb.AllUpdatesRule{}
			}
			obj.AllUpdatesRule.PubsubTopic = req.Budget.GetAllUpdatesRule().GetPubsubTopic()
		case "amount.specifiedAmount.nanos":
			if obj.Amount == nil {
				obj.Amount = &pb.BudgetAmount{}
			}
			if obj.Amount.GetSpecifiedAmount() == nil {
				obj.Amount.BudgetAmount = &pb.BudgetAmount_SpecifiedAmount{
					SpecifiedAmount: &moneypb.Money{},
				}
			}
			obj.Amount.GetSpecifiedAmount().Nanos = req.Budget.GetAmount().GetSpecifiedAmount().GetNanos()
		case "amount.specifiedAmount.units":
			if obj.Amount == nil {
				obj.Amount = &pb.BudgetAmount{}
			}
			if obj.Amount.GetSpecifiedAmount() == nil {
				obj.Amount.BudgetAmount = &pb.BudgetAmount_SpecifiedAmount{
					SpecifiedAmount: &moneypb.Money{},
				}
			}
			obj.Amount.GetSpecifiedAmount().Units = req.Budget.GetAmount().GetSpecifiedAmount().GetUnits()
		case "displayName":
			common.MustCopyField(req.Budget, obj, path)
		case "thresholdRules":
			obj.ThresholdRules = req.Budget.GetThresholdRules()
		default:
			return nil, fmt.Errorf("unhandled path %q in mockgcp UpdateBudget", path)
		}
	}

	// We randomize the order of projects to simulate non-determinism
	// since the order is not guaranteed by the API.
	reorderSlice(obj.GetBudgetFilter().GetProjects())

	obj.Etag = fields.ComputeWeakEtag(obj)

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

// reorderSlice predictably reorders the given slice in place.
// So that our tests produce consistent output,
// we want the same input slice to always produce the same output output.
// But the output order does not always match the input order, so we test
// controllers against assumptions that the order is preserved.
func reorderSlice[T any](slice []T) {
	// We order using a pseudo-random shuffle with a seed derived from the length of the slice.
	// It's an unusual choice, but it gives us a deterministic yet non-trivial reordering.
	// We don't want to depend on the values, because the values often include ${uniqueID} which changes between test runs.
	n := len(slice)
	seed := int64(n * 2654435761) // Knuth's multiplicative hash
	random := rand.New(rand.NewSource(seed))
	random.Shuffle(n, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
