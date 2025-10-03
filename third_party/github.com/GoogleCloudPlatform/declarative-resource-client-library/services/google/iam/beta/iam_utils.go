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
package beta

import (
	"bytes"
	"context"
	"fmt"
	"io"
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
	b, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// normalizeServiceAccountName converts name to short name and removes domain from the tail.
// Example: Example: projects/xyz/serviceAccounts/test-id-xcad-1665079476@xyz.iam.gserviceaccount.com becomes test-id-xcad-1665079476
func normalizeServiceAccountName(name *string) *string {
	newName := dcl.SelfLinkToName(name)
	*newName = strings.Split(*newName, "@")[0]
	return newName
}

func (r *ServiceAccount) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	nr.Name = normalizeServiceAccountName(nr.Name)
	params := map[string]any{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", nr.basePath(), userBasePath, params), nil
}

func (r *ServiceAccount) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	nr.Name = normalizeServiceAccountName(nr.Name)
	params := map[string]any{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts", nr.basePath(), userBasePath, params), nil
}

func (r *ServiceAccount) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	nr.Name = normalizeServiceAccountName(nr.Name)
	params := map[string]any{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com", nr.basePath(), userBasePath, params), nil
}

// SetPolicyURL constructs url for setting IAM Policy.
func (r *ServiceAccount) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]any{
		"project": *nr.Project,
		"name":    *nr.Name,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *ServiceAccount) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]any{
		"project": *nr.Project,
		"name":    *nr.Name,
	}
	return dcl.URL("projects/{{project}}/serviceAccounts/{{name}}@{{project}}.iam.gserviceaccount.com:getIamPolicy", nr.basePath(), userBasePath, fields)
}

// We are using CUSTOM_URL_METHOD trait on GetIAMPolicy and SetIAMPolicy which requires us to define these custom methods for IAM.

// SetPolicyVerb sets the verb for SetPolicy.
func (r *ServiceAccount) SetPolicyVerb() string {
	return "POST"
}

// IAMPolicyVersion defines version for IAMPolicy.
func (r *ServiceAccount) IAMPolicyVersion() int {
	return 3
}

// GetPolicy gets the IAM policy.
func (r *ServiceAccount) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	body.WriteString(fmt.Sprintf(`{"options":{"requestedPolicyVersion": %d}}`, r.IAMPolicyVersion()))
	return u, "POST", body, nil
}
