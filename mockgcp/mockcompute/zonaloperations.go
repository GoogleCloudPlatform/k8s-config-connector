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

package mockcompute

import (
	"context"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/compute/v1"
)

type ZoneOperationsV1 struct {
	*MockService
	pb.UnimplementedZoneOperationsServer
}

func (s *ZoneOperationsV1) Get(ctx context.Context, req *pb.GetZoneOperationRequest) (*pb.Operation, error) {
	fqn := s.zonalOperationFQN(req.Project, req.Zone, req.Operation)

	op := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, op); err != nil {
		return nil, err
	}

	// If operation is still RUNNING, complete it when polled
	if op.Status != nil && *op.Status == pb.Operation_RUNNING {
		op.Status = PtrTo(pb.Operation_DONE)
		op.Progress = PtrTo(int32(100))
		op.EndTime = PtrTo(s.nowString())

		// Update the operation in storage
		if err := s.storage.Update(ctx, fqn, op); err != nil {
			return nil, err
		}
	}

	return op, nil
}

func (s *ZoneOperationsV1) Wait(ctx context.Context, req *pb.WaitZoneOperationRequest) (*pb.Operation, error) {
	fqn := s.zonalOperationFQN(req.Project, req.Zone, req.Operation)

	deadline := 2 * time.Minute
	timeoutAt := time.Now().Add(deadline)
	for {
		op := &pb.Operation{}
		if err := s.storage.Get(ctx, fqn, op); err != nil {
			return nil, err
		}
		switch ValueOf(op.Status) {
		case pb.Operation_DONE:
			return op, nil
		}

		if time.Now().After(timeoutAt) {
			return op, nil
		}

		time.Sleep(2 * time.Second)
	}
}
