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

package operations

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "google.golang.org/genproto/googleapis/longrunning"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

type Operations struct {
	storage storage.Storage

	pb.UnimplementedOperationsServer
}

func NewOperationsService(storage storage.Storage) *Operations {
	return &Operations{
		storage: storage,
	}
}

func (s *Operations) NewLRO(ctx context.Context) (*pb.Operation, error) {
	now := time.Now()
	millis := now.UnixMilli()
	id := string(uuid.NewUUID())

	op := &pb.Operation{}

	op.Name = fmt.Sprintf("operations/operation-%d-%s", millis, id)
	op.Done = true

	fqn := op.Name

	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating LRO: %v", err)
	}
	return op, nil
}

func (s *Operations) StartLRO(ctx context.Context, metadata proto.Message, callback func() (proto.Message, error)) (*pb.Operation, error) {
	now := time.Now()
	millis := now.UnixMilli()
	id := uuid.NewUUID()

	op := &pb.Operation{}

	op.Name = fmt.Sprintf("operations/operation-%d-%s", millis, id)
	op.Done = false

	if metadata != nil {
		metadataAny, err := anypb.New(metadata)
		if err != nil {
			return nil, fmt.Errorf("error building anypb for metadata: %w", err)
		}
		rewriteTypes(metadataAny)

		op.Metadata = metadataAny
	}
	fqn := op.Name

	if err := s.storage.Create(ctx, fqn, op); err != nil {
		return nil, status.Errorf(codes.Internal, "error creating LRO: %v", err)
	}

	go func() {
		result, err := callback()
		finished := &pb.Operation{}
		if err2 := s.storage.Get(ctx, fqn, finished); err2 != nil {
			klog.Warningf("error getting LRO: %v", err2)
			return
		}

		finished.Done = true
		if err != nil {
			finished.Result = &pb.Operation_Error{
				Error: &rpcstatus.Status{
					Message: fmt.Sprintf("error processing operation: %v", err),
				},
			}
		} else {
			resultAny, err := anypb.New(result)
			if err != nil {
				klog.Warningf("error building anypb for result: %v", err)
				finished.Result = &pb.Operation_Response{}
			} else {
				rewriteTypes(resultAny)

				finished.Result = &pb.Operation_Response{
					Response: resultAny,
				}
			}
		}
		if err := s.storage.Update(ctx, fqn, finished); err != nil {
			klog.Warningf("error updating LRO: %v", err)
			return
		}
	}()

	return op, nil
}

func rewriteTypes(any *anypb.Any) {
	// Fix our mockgcp hack
	if strings.HasPrefix(any.TypeUrl, "type.googleapis.com/mockgcp.") {
		any.TypeUrl = "type.googleapis.com/google." + strings.TrimPrefix(any.TypeUrl, "type.googleapis.com/mockgcp.")
	}
}

// Gets the latest state of a long-running operation.  Clients can use this
// method to poll the operation result at intervals as recommended by the API
// service.
func (s *Operations) GetOperation(ctx context.Context, req *pb.GetOperationRequest) (*pb.Operation, error) {
	fqn := req.GetName()

	op := &pb.Operation{}
	if err := s.storage.Get(ctx, fqn, op); err != nil {
		return nil, err
	}

	return op, nil
}
