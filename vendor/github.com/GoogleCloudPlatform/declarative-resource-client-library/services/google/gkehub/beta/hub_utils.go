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
// Package gkehub includes all of the code for gkehub.
package beta

import (
	"bytes"
	"context"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// Feature has custom url methods because it uses v1beta endpoints instead of v1beta1.
func (r *Feature) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", "https://gkehub.googleapis.com/v1/", userBasePath, params), nil
}

func (r *Feature) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features", "https://gkehub.googleapis.com/v1/", userBasePath, params), nil

}

func (r *Feature) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":     dcl.ValueOrEmptyString(nr.Name),
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features?featureId={{name}}", "https://gkehub.googleapis.com/v1/", userBasePath, params), nil

}

func (r *Feature) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":     dcl.ValueOrEmptyString(nr.Name),
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", "https://gkehub.googleapis.com/v1/", userBasePath, params), nil
}

func (op *updateFeatureUpdateFeatureOperation) do(ctx context.Context, r *Feature, c *Client) error {
	_, err := c.GetFeature(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateFeature")
	if err != nil {
		return err
	}
	u = strings.Replace(u, "v1beta1", "v1beta", 1)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": "labels,spec"})
	if err != nil {
		return err
	}

	req, err := newUpdateFeatureUpdateFeatureRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFeatureUpdateFeatureRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func sendFeatureUpdate(ctx context.Context, req map[string]interface{}, c *Client, u string) error {
	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := marshalUpdateFeatureUpdateFeatureRequest(c, req)
	if err != nil {
		return err
	}
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": "membershipSpecs"})
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1/", "GET")

	if err != nil {
		return err
	}

	return nil
}
