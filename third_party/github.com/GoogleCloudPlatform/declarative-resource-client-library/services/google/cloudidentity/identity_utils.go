// Copyright 2024 Google LLC. All Rights Reserved.
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
// Package cloudidentity contains support for DCL Cloud Identity.
package cloudidentity

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// do creates a update request for groups.
func (op *updateMembershipUpdateMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	var currentMembership *Membership
	if sh := dcl.FetchStateHint(op.ApplyOptions); sh != nil {
		if r, ok := sh.(*Membership); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Membership, got %T", sh)
		} else {
			currentMembership = r
		}
	}
	if currentMembership == nil {
		currentMembership = r
	}

	currentMembership, err := c.GetMembership(ctx, currentMembership)
	if err != nil {
		return err
	}

	add := []interface{}{}
	for _, newRole := range r.Roles {
		found := false
		for _, currentRole := range currentMembership.Roles {
			if *currentRole.Name == *newRole.Name {
				found = true
			}
		}

		if !found {
			add = append(add, newRole)
		}
	}

	delete := []string{}
	for _, currentRole := range currentMembership.Roles {
		found := false
		for _, newRole := range r.Roles {
			if *currentRole.Name == *newRole.Name {
				found = true
			}
		}

		if !found {
			delete = append(delete, *currentRole.Name)
		}
	}

	chg := map[string]interface{}{
		"addRoles": []interface{}{
			add,
		},
		"removeRoles": delete,
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateMembership")
	if err != nil {
		return err
	}

	req, err := json.Marshal(chg)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}
