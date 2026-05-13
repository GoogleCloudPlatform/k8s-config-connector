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

package direct

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func WaitOperation(ctx context.Context, opClient longrunning.OperationsClient, op *longrunning.Operation, result proto.Message) error {
	for {
		if op.Done {
			if op.GetError() != nil {
				return fmt.Errorf("operation failed: %v", op.GetError())
			}
			if result != nil && op.GetResponse() != nil {
				if err := anypb.UnmarshalTo(op.GetResponse(), result, proto.UnmarshalOptions{}); err != nil {
					return fmt.Errorf("unmarshaling operation response: %w", err)
				}
			}
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
		}

		var err error
		op, err = opClient.GetOperation(ctx, &longrunning.GetOperationRequest{Name: op.Name})
		if err != nil {
			return fmt.Errorf("getting operation: %w", err)
		}
	}
}
