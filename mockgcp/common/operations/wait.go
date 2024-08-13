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

package operations

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"k8s.io/apimachinery/pkg/util/wait"
)

// Wait for the operation to complete.
// If the operation is not done after timeout, an error is returned.
func (o *Operations) Wait(ctx context.Context, opName string, timeout time.Duration) (*longrunningpb.Operation, error) {
	var ret *longrunningpb.Operation
	if err := wait.PollImmediateWithContext(ctx, 100*time.Millisecond, timeout, func(ctx context.Context) (bool, error) {
		op, err := o.GetOperation(ctx, &longrunningpb.GetOperationRequest{
			Name: opName,
		})
		if err != nil {
			return false, fmt.Errorf("getting operation %q: %w", opName, err)
		}
		if op.Done {
			ret = op
			return true, nil
		}
		return false, nil
	}); err != nil {
		return nil, err
	}

	return ret, nil
}
