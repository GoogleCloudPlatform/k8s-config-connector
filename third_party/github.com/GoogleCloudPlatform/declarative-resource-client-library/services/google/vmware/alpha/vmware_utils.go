// Copyright 2023 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package vmware contains methods and types for handling vmware GCP resources.
package alpha

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

const deletePollIntervalSec = 60
const deleteTimeoutMin = 60
const maxPollErrors = 5

// PC deletion is "soft" which means that after the operation completes it still exists.
// There are other resources that need to be removed up as a part of the cleanup (eg. VEN)
// and there is a reference from PC which precludes them from doing so.
// We need to poll for PC existence up until it actually gets deleted before proceeding.
func (op *deletePrivateCloudOperation) do(ctx context.Context, r *PrivateCloud, c *Client) error {
	r, err := c.GetPrivateCloud(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "PrivateCloud not found, returning. Original error: %w", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetPrivateCloud checking for existence. error: %w", err)
		return err
	}

	url, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	url, err = dcl.AddQueryParams(url, map[string]string{"force": "true", "delay_hours": "0"})
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	if _, err := dcl.SendRequest(ctx, c.Config, "DELETE", url, body, c.Config.RetryProvider); err != nil {
		return fmt.Errorf("failed to delete PrivateCloud: %w", err)
	}

	// Loop until the PC is deleted
	deadline := time.Now().Add(deleteTimeoutMin * time.Minute)
	pollErrors := 0
	for time.Now().Before(deadline) {
		r, err = c.GetPrivateCloud(ctx, r)
		if err != nil {
			if dcl.IsNotFound(err) {
				c.Config.Logger.InfoWithContextf(ctx, "PrivateCloud successfully deleted")
				return nil
			}
			pollErrors++
			c.Config.Logger.WarningWithContextf(ctx, "GetPrivateCloud checking for existence. error: %w", err)
			if pollErrors >= maxPollErrors {
				return err
			}
		}
		time.Sleep(deletePollIntervalSec * time.Second)
	}

	return fmt.Errorf("timeout reached while waiting for PrivateCloud to be deleted")
}
