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

package cloudidentity

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	api "google.golang.org/api/cloudidentity/v1beta1"
)

func WaitForCloudIdentityOp(ctx context.Context, op *api.Operation) error {
	if op == nil {
		return nil
	}
	if op.Done {
		if op.Error != nil {
			return fmt.Errorf("operation %q completed with error: %v", op.Name, op.Error)
		}
		return nil
	}
	return fmt.Errorf("operation %q is not done, and no HTTP client is available for polling", op.Name)
}

func WaitForCloudIdentityOpWithClient(ctx context.Context, httpClient *http.Client, op *api.Operation) error {
	if op == nil {
		return nil
	}
	if op.Done {
		if op.Error != nil {
			return fmt.Errorf("operation %q completed with error: %v", op.Name, op.Error)
		}
		return nil
	}

	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		req, err := http.NewRequestWithContext(ctx, "GET", "https://cloudidentity.googleapis.com/v1beta1/"+op.Name, nil)
		if err != nil {
			return false, fmt.Errorf("creating operation GET request: %w", err)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return false, fmt.Errorf("getting operation status: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return false, fmt.Errorf("getting operation status returned status %s", resp.Status)
		}

		var current api.Operation
		if err := json.NewDecoder(resp.Body).Decode(&current); err != nil {
			return false, fmt.Errorf("decoding operation status response: %w", err)
		}

		if current.Done {
			if current.Error != nil {
				return true, fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			}
			return true, nil
		}
		return false, nil
	})
}
