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
// Package billingbudgets defines functions and types for handling billingbudgets GCP resources.
package billingbudgets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Create and update have custom methods because the request object must be encoded differently depending on whether the endpoint is beta or v1.
func (op *createBudgetOperation) do(ctx context.Context, r *Budget, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	m, err := expandBudget(c, r)
	if err != nil {
		return err
	}
	if strings.HasPrefix(u, "https://billingbudgets.googleapis.com/v1beta1/") {
		n := make(map[string]interface{})
		n["budget"] = m
		m = n
	}
	req, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		details, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
		if err != nil {
			// Occasionally the server will return a 503 with message "The service is currently unavailable."
			return details, err
		}

		o, err := dcl.ResponseBodyAsJSON(details)
		if err != nil {
			return nil, fmt.Errorf("error decoding response body into JSON: %w", err)
		}
		op.response = o
		return details, nil
	}, c.Config.RetryProvider)

	if err != nil {
		return err
	}

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string")
	}
	r.Name = &name

	// Poll for the Budget resource to be created. Budget resources are eventually consistent but do not support operations
	// so we must repeatedly poll to check for their creation.
	start := time.Now()
	err = dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		u, err := r.getURL(c.Config.BasePath)
		if err != nil {
			return nil, err
		}
		getResp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, nil)
		if err != nil {
			// If the error is a transient server error (e.g., 500) or not found (i.e., the resource has not yet been created),
			// continue retrying until the transient error is resolved, the resource is created, or we time out.
			if dcl.IsRetryableRequestError(c.Config, err, true, start) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		getResp.Response.Body.Close()
		return getResp, nil
	}, c.Config.RetryProvider)

	if err != nil {
		return err
	}

	if _, err := c.GetBudget(ctx, r); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (op *updateBudgetUpdateBudgetOperation) do(ctx context.Context, r *Budget, c *Client) error {
	_, err := c.GetBudget(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateBudget")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	m, err := expandBudget(c, r)
	if err != nil {
		return err
	}
	if strings.HasPrefix(u, "https://billingbudgets.googleapis.com/v1beta1/") {
		n := make(map[string]interface{})
		n["budget"] = m
		m = n
	}
	req, err := json.Marshal(m)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", m)
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Returns a list of project ids given a list of project numbers, both in the form projects/<value>.
func flattenBudgetFilterProjects(c *Client, projects interface{}, _ *Budget) []string {
	ctx := context.Background()
	projectInterfaces, ok := projects.([]interface{})
	if !ok {
		return nil
	}
	projectNumbers := make([]string, len(projectInterfaces))
	for i, projectInterface := range projectInterfaces {
		projectNumber, ok := projectInterface.(string)
		if !ok {
			return nil
		}
		projectNumbers[i] = projectNumber
	}
	projectIDs := make([]string, len(projectNumbers))
	for i, projectNumber := range projectNumbers {
		number := strings.TrimPrefix(projectNumber, "projects/")
		resp, err := dcl.SendRequest(ctx, c.Config, "GET", fmt.Sprintf("https://cloudresourcemanager.googleapis.com/v1/projects/%s", number), &bytes.Buffer{}, c.Config.RetryProvider)
		if err != nil {
			return nil
		}
		defer resp.Response.Body.Close()
		b, err := io.ReadAll(resp.Response.Body)
		if err != nil {
			return nil
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil
		}
		id, ok := m["projectId"].(string)
		if !ok {
			return nil
		}
		projectIDs[i] = fmt.Sprintf("projects/%s", id)
	}
	return projectIDs
}

// Returns the labels field as a map with slices as values.
func expandBudgetFilterLabels(_ *Client, labels map[string]BudgetBudgetFilterLabels, _ *Budget) (interface{}, error) {
	m := make(map[string][]string, len(labels))
	for k, v := range labels {
		m[k] = v.Values
	}
	return m, nil
}

func flattenBudgetFilterLabels(_ *Client, labels interface{}, _ *Budget) map[string]BudgetBudgetFilterLabels {
	labelsInterfaceMap, ok := labels.(map[string]interface{})
	if !ok {
		return nil
	}
	labelsStringsMap := make(map[string][]string, len(labelsInterfaceMap))
	for key, labelsInterfaceValue := range labelsInterfaceMap {
		labelsInterfaces, ok := labelsInterfaceValue.([]interface{})
		if !ok {
			return nil
		}
		labelsStrings := make([]string, len(labelsInterfaces))
		for i, labelsInterface := range labelsInterfaces {
			labelsString, ok := labelsInterface.(string)
			if !ok {
				return nil
			}
			labelsStrings[i] = labelsString
		}
		labelsStringsMap[key] = labelsStrings
	}
	m := make(map[string]BudgetBudgetFilterLabels, len(labelsStringsMap))
	for k, v := range labelsStringsMap {
		m[k] = BudgetBudgetFilterLabels{Values: v}
	}
	return m
}
