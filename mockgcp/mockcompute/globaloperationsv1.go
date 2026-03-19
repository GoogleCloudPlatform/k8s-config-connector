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

package mockcompute

import (
	"context"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type GlobalOperationsV1 struct {
	*MockService
	pb.UnimplementedGlobalOperationsServer
}

func (s *GlobalOperationsV1) Get(ctx context.Context, req *pb.GetGlobalOperationRequest) (*pb.Operation, error) {
	projectID := lastComponent(req.GetProject())
	operation := lastComponent(req.GetOperation())
	fqn := s.globalOperationFQN(projectID, operation)
	lro, err := s.getOperation(ctx, fqn)
	if err != nil {
		return nil, err
	}

	return lro, nil
}

func (s *GlobalOperationsV1) Wait(ctx context.Context, req *pb.WaitGlobalOperationRequest) (*pb.Operation, error) {
	projectID := lastComponent(req.GetProject())
	operation := lastComponent(req.GetOperation())
	fqn := s.globalOperationFQN(projectID, operation)

	deadline := 2 * time.Minute
	timeoutAt := time.Now().Add(deadline)
	for {
		lro, err := s.getOperation(ctx, fqn)
		if err != nil {
			return nil, err
		}
		switch ValueOf(lro.Status) {
		case pb.Operation_DONE:
			return lro, nil
		}

		if time.Now().After(timeoutAt) {
			return lro, nil
		}

		time.Sleep(2 * time.Second)
	}
}

func (s *GlobalOperationsV1) List(ctx context.Context, req *pb.ListGlobalOperationsRequest) (*pb.OperationList, error) {
	// TODO: Implement List
	return &pb.OperationList{}, nil
}

func (s *GlobalOperationsV1) AggregatedList(ctx context.Context, req *pb.AggregatedListGlobalOperationsRequest) (*pb.OperationAggregatedList, error) {
	// TODO: Implement AggregatedList
	return &pb.OperationAggregatedList{}, nil
}
