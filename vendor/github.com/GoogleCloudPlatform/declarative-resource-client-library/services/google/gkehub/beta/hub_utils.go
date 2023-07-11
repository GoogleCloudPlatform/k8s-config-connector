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
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func expandHubReferenceLink(_ *Client, val *string, _ *Membership) (interface{}, error) {
	if val == nil {
		return nil, nil
	}

	v := *val

	if strings.HasPrefix(v, "https:") {
		return strings.Replace(strings.Replace(strings.Replace(*val, "https:", "", 1), "v1/", "", 1), "v1beta1/", "", 1), nil
	} else if strings.HasPrefix(v, "//container.googleapis.com") {
		return v, nil
	}
	return "//container.googleapis.com/" + v, nil
}

func flattenHubReferenceLink(_ *Client, config interface{}, _ *Membership) *string {
	v, ok := config.(string)
	if !ok {
		return nil
	}

	v = strings.Replace(v, "//container.googleapis.com/", "", 1)

	return &v
}

// Feature has custom url methods because it uses v1beta endpoints instead of v1beta1.
func (r *Feature) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", "https://gkehub.googleapis.com/v1beta/", userBasePath, params), nil
}

func (r *Feature) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features", "https://gkehub.googleapis.com/v1beta/", userBasePath, params), nil

}

func (r *Feature) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":     dcl.ValueOrEmptyString(nr.Name),
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features?featureId={{name}}", "https://gkehub.googleapis.com/v1beta/", userBasePath, params), nil

}

func (r *Feature) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":     dcl.ValueOrEmptyString(nr.Name),
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{name}}", "https://gkehub.googleapis.com/v1beta/", userBasePath, params), nil
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
	err = o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

// getMembershipSpecs returns a map of membership specs taken from the get response of the feature membership's feature object.
func getMembershipSpecs(ctx context.Context, r *FeatureMembership, c *Client) (map[string]interface{}, error) {
	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	u = strings.Replace(u, "v1beta1", "v1beta", 1)
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	membershipSpecs, ok := m["membershipSpecs"].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}, nil
	}
	return membershipSpecs, nil
}

// Return the full key for a given FeatureMembership's entry in the membershipSpecs field.
func membershipSpecKey(r *FeatureMembership) string {
	params := map[string]interface{}{
		"project":    dcl.ValueOrEmptyString(r.Project),
		"location":   dcl.ValueOrEmptyString(r.Location),
		"membership": dcl.ValueOrEmptyString(r.Membership),
	}

	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/memberships/{{membership}}", params)
}

// Find and return the key and value in membershipSpecs matching the given membership.
func findMembershipSpec(membership string, membershipSpecs map[string]interface{}) (string, map[string]interface{}, error) {
	for key, value := range membershipSpecs {
		if strings.HasSuffix(key, membership) {
			spec, ok := value.(map[string]interface{})
			if !ok {
				return "", nil, errors.New("membership spec was not of map type")
			}
			return key, spec, nil
		}
	}
	return "", nil, &googleapi.Error{
		Code:    404,
		Message: "feature membership not found in feature membership specs",
	}
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
	err = o.Wait(ctx, c.Config, "https://gkehub.googleapis.com/v1beta/", "GET")

	if err != nil {
		return err
	}

	return nil
}

func (op *createFeatureMembershipOperation) do(ctx context.Context, r *FeatureMembership, c *Client) error {
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}
	u = strings.Replace(u, "v1beta1", "v1beta", 1)

	nr := r.urlNormalized()
	membershipSpecs, err := getMembershipSpecs(ctx, nr, c)
	if err != nil {
		return err
	}
	m, err := expandFeatureMembership(c, nr)
	if err != nil {
		return err
	}
	if err := dcl.PutMapEntry(membershipSpecs, []string{membershipSpecKey(nr)}, m); err != nil {
		return err
	}
	req := map[string]interface{}{
		"membershipSpecs": membershipSpecs,
	}
	return sendFeatureUpdate(ctx, req, c, u)
}

// GetFeatureMembership returns a feature membership object retrieved from the membershipSpecs field of a feature.
func (c *Client) GetFeatureMembership(ctx context.Context, r *FeatureMembership) (*FeatureMembership, error) {
	nr := r.urlNormalized()
	membershipSpecs, err := getMembershipSpecs(ctx, nr, c)
	if err != nil {
		return nil, err
	}
	_, spec, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), membershipSpecs)
	if err != nil {
		return nil, err
	}
	result, err := unmarshalMapFeatureMembership(spec, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Feature = r.Feature
	result.Membership = r.Membership

	c.Config.Logger.Infof("Retrieved raw result state: %v", result)
	c.Config.Logger.Infof("Canonicalizing with specified state: %v", r)
	result, err = canonicalizeFeatureMembershipNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	c.Config.Logger.Infof("Created result state: %v", result)

	return result, nil
}

// HasNext always returns false because a feature membership list never has a next page.
func (l *FeatureMembershipList) HasNext() bool {
	return false
}

// Next returns nil because it will never be called.
func (l *FeatureMembershipList) Next(_ context.Context, _ *Client) error {
	return nil
}

// ListFeatureMembership returns a list of feature memberships retrieved from the membershipSpecs field of a feature.
func (c *Client) ListFeatureMembership(ctx context.Context, project, location, feature string) (*FeatureMembershipList, error) {
	r := &FeatureMembership{
		Project:  &project,
		Location: &location,
		Feature:  &feature,
	}
	membershipSpecs, err := getMembershipSpecs(ctx, r, c)
	if err != nil {
		return nil, err
	}
	var list *FeatureMembershipList
	for key, spec := range membershipSpecs {
		m, ok := spec.(map[string]interface{})
		if !ok {
			return nil, errors.New("membership spec was not of map type")
		}
		ri, err := unmarshalMapFeatureMembership(m, c, r)
		if err != nil {
			return nil, err
		}
		ri.Project = r.Project
		ri.Location = r.Location
		ri.Feature = r.Feature
		ri.Membership = dcl.SelfLinkToName(&key)
		list.Items = append(list.Items, ri)
	}
	return list, nil
}

func (op *updateFeatureMembershipUpdateFeatureMembershipOperation) do(ctx context.Context, r *FeatureMembership, c *Client) error {
	nr := r.urlNormalized()
	u, err := r.updateURL(c.Config.BasePath, "UpdateFeatureMembership")
	if err != nil {
		return err
	}
	u = strings.Replace(u, "v1beta1", "v1beta", 1)

	membershipSpecs, err := getMembershipSpecs(ctx, r, c)
	if err != nil {
		return err
	}
	key, _, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), membershipSpecs)
	if err != nil {
		return err
	}
	m, err := expandFeatureMembership(c, r)
	if err != nil {
		return err
	}
	if err := dcl.PutMapEntry(membershipSpecs, []string{key}, m); err != nil {
		return err
	}
	req := map[string]interface{}{
		"membershipSpecs": membershipSpecs,
	}
	return sendFeatureUpdate(ctx, req, c, u)
}

func (op *deleteFeatureMembershipOperation) do(ctx context.Context, r *FeatureMembership, c *Client) error {
	nr := r.urlNormalized()
	u, err := nr.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}
	u = strings.Replace(u, "v1beta1", "v1beta", 1)

	membershipSpecs, err := getMembershipSpecs(ctx, nr, c)
	if err != nil {
		return err
	}
	key, _, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), membershipSpecs)
	if err != nil {
		return err
	}
	membershipSpecs[key] = map[string]interface{}{}
	req := map[string]interface{}{
		"membershipSpecs": membershipSpecs,
	}
	return sendFeatureUpdate(ctx, req, c, u)
}

// CompareFeatureMembershipConfigmanagementHierarchyControllerNewStyle exists only for unit-testing the diff library.
func CompareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	return compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a, fn)
}
