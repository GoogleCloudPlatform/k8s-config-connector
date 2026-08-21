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
// Package apigee contains methods and types for handling apigee GCP resources.
package apigee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// createProperties adds an update to apply Properties after create.
func createProperties(inOps []organizationApiOperation) ([]organizationApiOperation, error) {
	for _, op := range inOps {
		if _, ok := op.(*createOrganizationOperation); ok {
			return append(inOps, &updateOrganizationUpdateOrganizationOperation{FieldDiffs: []*dcl.FieldDiff{{FieldName: "properties"}}}), nil
		}
	}
	return inOps, nil
}

// listOrganizationRaw is located here because other functions required by the list endpoint are custom.
func (c *Client) listOrganizationRaw(ctx context.Context, r *Organization, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != OrganizationMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return io.ReadAll(resp.Response.Body)
}

type listOrganizationOperation struct {
	Organizations []map[string]interface{} `json:"organizations"`
	Token         string                   `json:"nextPageToken"`
}

// listOrganization is a custom method which handles the different format that apigeeOrganization's list method returns.
func (c *Client) listOrganization(ctx context.Context, r *Organization, pageToken string, pageSize int32) ([]*Organization, string, error) {
	b, err := c.listOrganizationRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listOrganizationOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Organization
	for _, v := range m.Organizations {
		name, ok := v["organization"].(string)
		if !ok {
			return l, "", fmt.Errorf("name field was %T, not string", v["apigeeOrganization"])
		}
		projectIDs, ok := v["projectIds"].([]interface{})
		if !ok {
			return l, "", fmt.Errorf("project ids field was %T, not slice", v["projectIds"])
		}
		projectID, ok := projectIDs[0].(string)
		if !ok {
			return l, "", fmt.Errorf("project id was %T, not slice", projectIDs[0])
		}
		l = append(l, &Organization{
			Name:      dcl.String(name),
			ProjectId: dcl.String(projectID),
		})
	}

	return l, m.Token, nil
}

// HasNext returns whether the list has a next page.
func (l *OrganizationList) HasNext() bool {
	return l.nextToken != ""
}

// Next advances the list to its next page.
func (l *OrganizationList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listOrganization(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

// ListOrganization returns a list of apigee organizations which the client has permission to access.
func (c *Client) ListOrganization(ctx context.Context) (*OrganizationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	return c.ListOrganizationWithMaxResults(ctx, OrganizationMaxPage)

}

// ListOrganizationWithMaxResults returns a list of apigee organizations with the given page size.
func (c *Client) ListOrganizationWithMaxResults(ctx context.Context, pageSize int32) (*OrganizationList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	r := &Organization{}
	items, token, err := c.listOrganization(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &OrganizationList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.
func (op *updateOrganizationUpdateOrganizationOperation) do(ctx context.Context, r *Organization, c *Client) error {
	_, err := c.GetOrganization(ctx, r.urlNormalized())
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateOrganization")
	if err != nil {
		return err
	}

	req, err := newUpdateOrganizationUpdateOrganizationRequest(ctx, r, c)
	if err != nil {
		return err
	}
	dcl.PutMapEntry(req, []string{"name"}, r.Name)

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateOrganizationUpdateOrganizationRequest(c, req)
	if err != nil {
		return err
	}
	if _, err := dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider); err != nil {
		return err
	}

	return nil
}

func (c *Client) listEnvironmentRaw(ctx context.Context, r *Environment, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != EnvironmentMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return io.ReadAll(resp.Response.Body)
}

// listEnvironment is needed because the ListEnvironments method returns a list of environment names rather than objects.
func (c *Client) listEnvironment(ctx context.Context, r *Environment, pageToken string, pageSize int32) ([]*Environment, string, error) {
	b, err := c.listEnvironmentRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	names := make([]string, 0)
	if err := json.Unmarshal(b, &names); err != nil {
		return nil, "", err
	}

	var l []*Environment
	for _, name := range names {
		res := Environment{
			Name:               &name,
			ApigeeOrganization: r.ApigeeOrganization,
		}
		l = append(l, &res)
	}

	return l, "", nil
}

// HasNext returns an empty token.
func (l *EnvironmentList) HasNext() bool {
	return l.nextToken != ""
}

// Next returns an error because environment list has no pages.
func (l *EnvironmentList) Next(ctx context.Context, c *Client) error {
	return fmt.Errorf("no next page")
}

// ListEnvironment returns an EnvironmentList containing all Environment resources in the given organization.
func (c *Client) ListEnvironment(ctx context.Context, apigeeOrganization string) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	return c.ListEnvironmentWithMaxResults(ctx, apigeeOrganization, EnvironmentMaxPage)

}

// ListEnvironmentWithMaxResults returns an EnvironmentList containing all Environment resources in the given organization.
func (c *Client) ListEnvironmentWithMaxResults(ctx context.Context, apigeeOrganization string, pageSize int32) (*EnvironmentList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0))
	defer cancel()

	r := &Environment{ApigeeOrganization: &apigeeOrganization}
	items, token, err := c.listEnvironment(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &EnvironmentList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}
