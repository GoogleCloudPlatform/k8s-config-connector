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

package cloudidentity

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	api "google.golang.org/api/cloudidentity/v1beta1"
)

func WaitForCloudIdentityOp(ctx context.Context, op *api.Operation) error {
	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		if op.Done {
			if op.Error != nil {
				return true, fmt.Errorf("operation %q completed with error: %v", op.Name, op.Error)
			} else {
				return true, nil
			}
		}
		return false, nil
	})
}
