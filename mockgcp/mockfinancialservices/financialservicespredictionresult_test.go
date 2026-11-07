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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreatePredictionResultParentNotFound(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	parent := "projects/test-project/locations/us-central1/instances/nonexistent-instance"
	req := &pb.CreatePredictionResultRequest{
		Parent:             parent,
		PredictionResultId: "pred-1",
		PredictionResult:   &pb.PredictionResult{},
	}

	_, err := server.CreatePredictionResult(ctx, req)
	if err == nil {
		t.Fatalf("expected error when parent instance does not exist, got nil")
	}
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected codes.NotFound, got %v (%v)", status.Code(err), err)
	}
}
