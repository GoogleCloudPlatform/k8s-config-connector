// Copyright 2026 Google LLC
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

package mockfinancialservices

import (
	"context"
	"testing"

	pb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestFinancialServicesBacktestResultLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	locationParent := "projects/test-project/locations/us-central1"
	instanceID := "test-instance"
	instanceName := locationParent + "/instances/" + instanceID
	backtestResultID := "test-backtest"
	backtestResultName := instanceName + "/backtestResults/" + backtestResultID

	// 1. Try to create BacktestResult when parent instance does not exist -> NotFound
	createReq := &pb.CreateBacktestResultRequest{
		Parent:           instanceName,
		BacktestResultId: backtestResultID,
		BacktestResult: &pb.BacktestResult{
			Dataset: instanceName + "/datasets/test-dataset",
			Model:   instanceName + "/models/test-model",
			EndTime: timestamppb.Now(),
			PerformanceTarget: &pb.BacktestResult_PerformanceTarget{
				PartyInvestigationsPerPeriodHint: 100,
			},
		},
	}
	_, err := server.CreateBacktestResult(ctx, createReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when parent instance missing, got: %v", err)
	}

	// 2. Create parent instance
	_, err = server.CreateInstance(ctx, &pb.CreateInstanceRequest{
		Parent:     locationParent,
		InstanceId: instanceID,
		Instance: &pb.Instance{
			KmsKey: "projects/test-project/locations/us-central1/keyRings/test-kr/cryptoKeys/test-key",
		},
	})
	if err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}

	// 3. Get before creation -> NotFound
	_, err = server.GetBacktestResult(ctx, &pb.GetBacktestResultRequest{Name: backtestResultName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent backtest result, got: %v", err)
	}

	// 4. Create BacktestResult with existing parent -> Success
	op, err := server.CreateBacktestResult(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateBacktestResult failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateBacktestResult")
	}

	// 5. Get after creation
	br, err := server.GetBacktestResult(ctx, &pb.GetBacktestResultRequest{Name: backtestResultName})
	if err != nil {
		t.Fatalf("GetBacktestResult failed: %v", err)
	}
	if br.GetName() != backtestResultName {
		t.Errorf("expected name %q, got %q", backtestResultName, br.GetName())
	}
	if br.GetState() != pb.BacktestResult_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", br.GetState())
	}
	if br.GetPerformanceTarget().GetPartyInvestigationsPerPeriodHint() != 100 {
		t.Errorf("expected PartyInvestigationsPerPeriodHint=100, got %v", br.GetPerformanceTarget().GetPartyInvestigationsPerPeriodHint())
	}

	// 6. List BacktestResults
	listResp, err := server.ListBacktestResults(ctx, &pb.ListBacktestResultsRequest{Parent: instanceName})
	if err != nil {
		t.Fatalf("ListBacktestResults failed: %v", err)
	}
	if len(listResp.GetBacktestResults()) != 1 {
		t.Fatalf("expected 1 backtest result, got %d", len(listResp.GetBacktestResults()))
	}

	// 7. Update BacktestResult
	updateReq := &pb.UpdateBacktestResultRequest{
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"backtest_periods", "performance_target.party_investigations_per_period_hint"},
		},
		BacktestResult: &pb.BacktestResult{
			Name:            backtestResultName,
			BacktestPeriods: 6,
			PerformanceTarget: &pb.BacktestResult_PerformanceTarget{
				PartyInvestigationsPerPeriodHint: 200,
			},
		},
	}
	_, err = server.UpdateBacktestResult(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateBacktestResult failed: %v", err)
	}

	// 8. Verify update
	br, err = server.GetBacktestResult(ctx, &pb.GetBacktestResultRequest{Name: backtestResultName})
	if err != nil {
		t.Fatalf("GetBacktestResult after update failed: %v", err)
	}
	if br.GetBacktestPeriods() != 6 {
		t.Errorf("expected BacktestPeriods=6, got %v", br.GetBacktestPeriods())
	}
	if br.GetPerformanceTarget().GetPartyInvestigationsPerPeriodHint() != 200 {
		t.Errorf("expected PartyInvestigationsPerPeriodHint=200, got %v", br.GetPerformanceTarget().GetPartyInvestigationsPerPeriodHint())
	}

	// 9. Delete BacktestResult
	_, err = server.DeleteBacktestResult(ctx, &pb.DeleteBacktestResultRequest{Name: backtestResultName})
	if err != nil {
		t.Fatalf("DeleteBacktestResult failed: %v", err)
	}

	// 10. Get after delete -> NotFound
	_, err = server.GetBacktestResult(ctx, &pb.GetBacktestResultRequest{Name: backtestResultName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}
