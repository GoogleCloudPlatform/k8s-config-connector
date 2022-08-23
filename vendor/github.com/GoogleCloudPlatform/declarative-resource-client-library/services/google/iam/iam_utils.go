// Copyright 2022 Google LLC. All Rights Reserved.
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
package iam

import (
	"bytes"
	"context"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// EncodeIAMCreateRequest encodes the create request for an iam resource.
func EncodeIAMCreateRequest(m map[string]interface{}, resourceName, idField string) map[string]interface{} {
	req := make(map[string]interface{})
	// Put base object into object field.
	dcl.PutMapEntry(req, []string{resourceName}, m)
	// Move name field from from nested object to id field.
	dcl.MoveMapEntry(req, []string{resourceName, "name"}, []string{idField})
	// Delete projectID field.
	delete(req[resourceName].(map[string]interface{}), "projectId")
	// Change value of id field to only last part after / and before @.
	idParts := regexp.MustCompile("([^@/]+)[^/]*$").FindStringSubmatch(*req[idField].(*string))
	if len(idParts) < 2 {
		return req
	}
	req[idField] = idParts[1]
	return req
}

// EncodeRoleCreateRequest properly encodes the create request for an iam role.
func EncodeRoleCreateRequest(m map[string]interface{}) map[string]interface{} {
	return EncodeIAMCreateRequest(m, "role", "roleId")
}

// EncodeServiceAccountCreateRequest properly encodes the create request for an iam service account.
func EncodeServiceAccountCreateRequest(m map[string]interface{}) map[string]interface{} {
	return EncodeIAMCreateRequest(m, "serviceAccount", "accountId")
}

// canonicalizeServiceAccountName compares service account names ignoring the part after @.
func canonicalizeServiceAccountName(m, n interface{}) bool {
	mStr, ok := m.(*string)
	if !ok {
		return false
	}
	nStr, ok := n.(*string)
	if !ok {
		return false
	}
	if mStr == nil && nStr == nil {
		return true
	}
	if mStr == nil || nStr == nil {
		return false
	}
	// Compare values before @.
	mVal := strings.Split(*mStr, "@")[0]
	nVal := strings.Split(*nStr, "@")[0]
	return dcl.PartialSelfLinkToSelfLink(&mVal, &nVal)
}

func (c *Client) GetWorkloadIdentityPool(ctx context.Context, r *WorkloadIdentityPool) (*WorkloadIdentityPool, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	b, err := c.getWorkloadIdentityPoolRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalWorkloadIdentityPool(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeWorkloadIdentityPoolNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) getWorkloadIdentityPoolRaw(ctx context.Context, r *WorkloadIdentityPool) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	var resp *dcl.RetryDetails
	// Retry until project is ready.
	ctt, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	err = dcl.Do(ctt, func(ctt context.Context) (*dcl.RetryDetails, error) {
		var err error
		resp, err = dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
		if err != nil {
			if gerr, ok := err.(*googleapi.Error); ok {
				if gerr.Code == 403 && strings.HasPrefix(gerr.Message, "Permission 'iam.workloadIdentityPools.get' denied on resource") {
					return &dcl.RetryDetails{}, dcl.OperationNotDone{}
				}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
