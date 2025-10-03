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
// Package gkehub includes all of the code for gkehub.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

// getMembershipSpecs returns a map of membership specs taken from the get response of the feature membership's feature object.
func getMembershipSpecs(ctx context.Context, r *FeatureMembership, c *Client) (map[string]any, error) {
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
	b, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	membershipSpecs, ok := m["membershipSpecs"].(map[string]any)
	if !ok {
		return map[string]any{}, nil
	}
	return membershipSpecs, nil
}

// Return the value if it exists, default otherwise
func valueOrDefaultString(val *string, def string) string {
	if dcl.ValueOrEmptyString(val) == "" {
		return def
	}
	return dcl.ValueOrEmptyString(val)
}

// Return the full key for a given FeatureMembership's entry in the membershipSpecs field.
func membershipSpecKey(r *FeatureMembership) string {
	params := map[string]any{
		"project":    dcl.ValueOrEmptyString(r.Project),
		"location":   valueOrDefaultString(r.MembershipLocation, "global"),
		"membership": dcl.ValueOrEmptyString(r.Membership),
	}

	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/memberships/{{membership}}", params)
}

// Find and return the key and value in membershipSpecs matching the given membership.
func findMembershipSpec(membership string, membershipLocation string, membershipSpecs map[string]any) (string, map[string]any, error) {
	for key, value := range membershipSpecs {
		if strings.HasSuffix(key, fmt.Sprintf("%s/memberships/%s", membershipLocation, membership)) {
			spec, ok := value.(map[string]any)
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

func sendFeatureUpdate(ctx context.Context, req map[string]any, c *Client, u string) error {
	c.Config.Logger.Infof("Created update: %#v", req)
	body, err := json.Marshal(req)
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
	req := map[string]any{
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
	_, spec, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), valueOrDefaultString(nr.MembershipLocation, "global"), membershipSpecs)
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
		m, ok := spec.(map[string]any)
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
	key, _, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), valueOrDefaultString(nr.MembershipLocation, "global"), membershipSpecs)
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
	req := map[string]any{
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
	key, _, err := findMembershipSpec(dcl.ValueOrEmptyString(nr.Membership), valueOrDefaultString(nr.MembershipLocation, "global"), membershipSpecs)
	if err != nil {
		return err
	}
	membershipSpecs[key] = map[string]any{}
	req := map[string]any{
		"membershipSpecs": membershipSpecs,
	}
	return sendFeatureUpdate(ctx, req, c, u)
}

// CompareFeatureMembershipConfigmanagementHierarchyControllerNewStyle exists only for unit-testing the diff library.
func CompareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a any, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	return compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a, fn)
}

// This function behaves the same way as the generated diff function, except that it explicitly
// checks for emptiness as well.
func emptyHNCSameAsAllFalse(d, a any) bool {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			fmt.Printf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController or *FeatureMembershipConfigmanagementHierarchyController\n", d)
			return false
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			fmt.Printf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController\n", a)
			return false
		}
		actual = &actualNotPointer
	}

	if actual == nil && desired == nil {
		return true
	}
	if actual == nil || desired == nil {
		return false
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, dcl.FieldName{FieldName: "Enabled"}); len(ds) != 0 || err != nil {
		if err != nil {
			fmt.Print(err)
			return false
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnablePodTreeLabels, actual.EnablePodTreeLabels, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, dcl.FieldName{FieldName: "EnablePodTreeLabels"}); len(ds) != 0 || err != nil {
		if err != nil {
			fmt.Print(err)
			return false
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableHierarchicalResourceQuota, actual.EnableHierarchicalResourceQuota, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, dcl.FieldName{FieldName: "EnableHierarchicalResourceQuota"}); len(ds) != 0 || err != nil {
		if err != nil {
			fmt.Print(err)
			return false
		}
		diffs = append(diffs, ds...)
	}

	if len(diffs) > 0 {
		return false
	}

	if desired.Empty() != actual.Empty() {
		return false
	}
	return true
}

func flattenHierarchyControllerConfig(c *Client, i any, v *FeatureMembership) *FeatureMembershipConfigmanagementHierarchyController {
	m, ok := i.(map[string]any)
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementHierarchyController{}

	// Compared to the generated code, we removed the part where we skip flattening the API response
	// if the return value is empty (i.e. HNC = {}). This is because the Hub API returns the same
	// empty object for both {} (empty config) and {fieldA: false, fieldB: false, fieldC: false}. We
	// always flatten the response into the latter form i.e. explicitly stating false values, so that
	// it fits more easily into the declarative pattern and avoids a permadiff bug.
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.EnablePodTreeLabels = dcl.FlattenBool(m["enablePodTreeLabels"])
	r.EnableHierarchicalResourceQuota = dcl.FlattenBool(m["enableHierarchicalResourceQuota"])

	return r
}

func expandHierarchyControllerConfig(c *Client, f *FeatureMembershipConfigmanagementHierarchyController, res *FeatureMembership) (map[string]any, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]any)
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.EnablePodTreeLabels; !dcl.IsEmptyValueIndirect(v) {
		m["enablePodTreeLabels"] = v
	}
	if v := f.EnableHierarchicalResourceQuota; !dcl.IsEmptyValueIndirect(v) {
		m["enableHierarchicalResourceQuota"] = v
	}

	return m, nil
}
