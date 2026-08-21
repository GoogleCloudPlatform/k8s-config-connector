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

package mockdns

import (
	"context"
	"fmt"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/dns/v1"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type dnsOperations struct {
	storage storage.Storage
	pb.UnimplementedManagedZoneOperationsServerServer
}

func newDNSOperationsService(storage storage.Storage) *dnsOperations {
	return &dnsOperations{
		storage: storage,
	}
}

func (s *dnsOperations) GetManagedZoneOperation(ctx context.Context, req *pb.GetManagedZoneOperationRequest) (*pb.Operation, error) {
	fqn := s.managedZoneOperationFQN(req.GetProject(), req.GetManagedZone(), req.GetOperation())
	lro, err := s.getOperation(ctx, fqn)
	if err != nil {
		return nil, err
	}

	return lro, nil
}

func (s *dnsOperations) managedZoneOperationFQN(project, managedZone, operation string) string {
	return fmt.Sprintf("projects/%s/managedZones/%s/operations/%s", project, managedZone, operation)
}

func (s *dnsOperations) startLRO0(ctx context.Context, op *pb.Operation, fqn string, callback func() (proto.Message, error)) (*pb.Operation, error) {
	log := klog.FromContext(ctx)

	now := time.Now()

	if op == nil {
		op = &pb.Operation{}
	}

	op.StartTime = PtrTo(formatTime(now))

	if op.Status == nil {
		op.Status = PtrTo("running")
	}

	op.Kind = PtrTo("dns#operation")

	log.Info("storing operation", "fqn", fqn)
	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, err
	}

	go func() {
		result, err := callback()
		finished := &pb.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		finished.Status = PtrTo("done")

		if err != nil {
			klog.Warningf("TODO: more fully handle LRO error %v", err)
		} else {
			// The LRO result does not appear to be returned in the operation
			klog.V(4).Infof("LRO result: %+v", result)
		}
		if err := s.storage.Update(ctx, fqn, finished); err != nil {
			klog.Warningf("error updating LRO: %v", err)
			return
		}
	}()

	return op, nil
}

func (s *dnsOperations) StartLRO(ctx context.Context, lroPrefix string, op *pb.Operation, callback func() (proto.Message, error)) (*pb.Operation, error) {
	now := time.Now()
	nanos := now.UnixNano()

	op.Id = PtrTo(fmt.Sprintf("%d", nanos))

	fqn := lroPrefix + "operations/" + *op.Id

	return s.startLRO0(ctx, op, fqn, callback)
}

// Gets the latest state of a long-running operation.  Clients can use this
// method to poll the operation result at intervals as recommended by the API
// service.
func (s *dnsOperations) getOperation(ctx context.Context, fqn string) (*pb.Operation, error) {
	op := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, op); err != nil {
		return nil, err
	}

	return op, nil
}

func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
