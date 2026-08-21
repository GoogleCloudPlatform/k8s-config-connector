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
package alpha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Group) validate() error {

	if err := dcl.Required(r, "groupKey"); err != nil {
		return err
	}
	if err := dcl.Required(r, "parent"); err != nil {
		return err
	}
	if err := dcl.Required(r, "labels"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GroupKey) {
		if err := r.GroupKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DirectMemberCountPerType) {
		if err := r.DirectMemberCountPerType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DynamicGroupMetadata) {
		if err := r.DynamicGroupMetadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GroupGroupKey) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *GroupAdditionalGroupKeys) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *GroupDirectMemberCountPerType) validate() error {
	return nil
}
func (r *GroupDerivedAliases) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *GroupDynamicGroupMetadata) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GroupDynamicGroupMetadataQueries) validate() error {
	return nil
}
func (r *GroupDynamicGroupMetadataStatus) validate() error {
	return nil
}
func (r *GroupPosixGroups) validate() error {
	return nil
}
func (r *Group) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudidentity.googleapis.com/v1beta1/", params)
}

func (r *Group) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Group) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("groups?view=BASIC&parent=customers/{{parent}}", nr.basePath(), userBasePath, params), nil

}

func (r *Group) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"initialGroupConfig": dcl.ValueOrEmptyString(nr.InitialGroupConfig),
	}
	return dcl.URL("groups?initialGroupConfig={{initialGroupConfig}}", nr.basePath(), userBasePath, params), nil

}

func (r *Group) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, params), nil
}

// groupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type groupApiOperation interface {
	do(context.Context, *Group, *Client) error
}

// newUpdateGroupUpdateGroupRequest creates a request for an
// Group resource's UpdateGroup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGroupUpdateGroupRequest(ctx context.Context, f *Group, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("groups/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandGroupDynamicGroupMetadata(c, f.DynamicGroupMetadata, res); err != nil {
		return nil, fmt.Errorf("error expanding DynamicGroupMetadata into dynamicGroupMetadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["dynamicGroupMetadata"] = v
	}
	if v, err := expandGroupPosixGroupsSlice(c, f.PosixGroups, res); err != nil {
		return nil, fmt.Errorf("error expanding PosixGroups into posixGroups: %w", err)
	} else if v != nil {
		req["posixGroups"] = v
	}
	return req, nil
}

// marshalUpdateGroupUpdateGroupRequest converts the update into
// the final JSON request body.
func marshalUpdateGroupUpdateGroupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGroupUpdateGroupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGroupUpdateGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	_, err := c.GetGroup(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateGroup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateGroupUpdateGroupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateGroupUpdateGroupRequest(c, req)
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
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listGroupRaw(ctx context.Context, r *Group, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GroupMaxPage {
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
	return ioutil.ReadAll(resp.Response.Body)
}

type listGroupOperation struct {
	Groups []map[string]interface{} `json:"groups"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listGroup(ctx context.Context, r *Group, pageToken string, pageSize int32) ([]*Group, string, error) {
	b, err := c.listGroupRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Group
	for _, v := range m.Groups {
		res, err := unmarshalMapGroup(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGroup(ctx context.Context, f func(*Group) bool, resources []*Group) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGroup(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteGroupOperation struct{}

func (op *deleteGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	r, err := c.GetGroup(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.InfoWithContextf(ctx, "Group not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetGroup checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetGroup(ctx, r)
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, nil
		}
		if retriesRemaining > 0 {
			retriesRemaining--
			return &dcl.RetryDetails{}, dcl.OperationNotDone{}
		}
		return nil, dcl.NotDeletedError{ExistingResource: r}
	}, c.Config.RetryProvider)
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createGroupOperation struct {
	response map[string]interface{}
}

func (op *createGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	if r.Name != nil {
		// Allowing creation to continue with Name set could result in a Group with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetGroup(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGroupRaw(ctx context.Context, r *Group) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
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

func (c *Client) groupDiffsForRawDesired(ctx context.Context, rawDesired *Group, opts ...dcl.ApplyOption) (initial, desired *Group, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Group
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Group); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Group, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeGroupDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGroup(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Group resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Group resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Group resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Group: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Group: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractGroupFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Group: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Group: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGroupInitialState(rawInitial, rawDesired *Group) (*Group, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGroupDesiredState(rawDesired, rawInitial *Group, opts ...dcl.ApplyOption) (*Group, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.GroupKey = canonicalizeGroupGroupKey(rawDesired.GroupKey, nil, opts...)
		rawDesired.DirectMemberCountPerType = canonicalizeGroupDirectMemberCountPerType(rawDesired.DirectMemberCountPerType, nil, opts...)
		rawDesired.DynamicGroupMetadata = canonicalizeGroupDynamicGroupMetadata(rawDesired.DynamicGroupMetadata, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Group{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.GroupKey = canonicalizeGroupGroupKey(rawDesired.GroupKey, rawInitial.GroupKey, opts...)
	canonicalDesired.AdditionalGroupKeys = canonicalizeGroupAdditionalGroupKeysSlice(rawDesired.AdditionalGroupKeys, rawInitial.AdditionalGroupKeys, opts...)
	if dcl.StringCanonicalize(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.DynamicGroupMetadata = canonicalizeGroupDynamicGroupMetadata(rawDesired.DynamicGroupMetadata, rawInitial.DynamicGroupMetadata, opts...)
	canonicalDesired.PosixGroups = canonicalizeGroupPosixGroupsSlice(rawDesired.PosixGroups, rawInitial.PosixGroups, opts...)
	if dcl.IsZeroValue(rawDesired.InitialGroupConfig) || (dcl.IsEmptyValueIndirect(rawDesired.InitialGroupConfig) && dcl.IsEmptyValueIndirect(rawInitial.InitialGroupConfig)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.InitialGroupConfig = rawInitial.InitialGroupConfig
	} else {
		canonicalDesired.InitialGroupConfig = rawDesired.InitialGroupConfig
	}
	return canonicalDesired, nil
}

func canonicalizeGroupNewState(c *Client, rawNew, rawDesired *Group) (*Group, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.GroupKey) && dcl.IsEmptyValueIndirect(rawDesired.GroupKey) {
		rawNew.GroupKey = rawDesired.GroupKey
	} else {
		rawNew.GroupKey = canonicalizeNewGroupGroupKey(c, rawDesired.GroupKey, rawNew.GroupKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AdditionalGroupKeys) && dcl.IsEmptyValueIndirect(rawDesired.AdditionalGroupKeys) {
		rawNew.AdditionalGroupKeys = rawDesired.AdditionalGroupKeys
	} else {
		rawNew.AdditionalGroupKeys = rawDesired.AdditionalGroupKeys
	}

	if dcl.IsEmptyValueIndirect(rawNew.Parent) && dcl.IsEmptyValueIndirect(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		if dcl.StringCanonicalize(rawDesired.Parent, rawNew.Parent) {
			rawNew.Parent = rawDesired.Parent
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DirectMemberCount) && dcl.IsEmptyValueIndirect(rawDesired.DirectMemberCount) {
		rawNew.DirectMemberCount = rawDesired.DirectMemberCount
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DirectMemberCountPerType) && dcl.IsEmptyValueIndirect(rawDesired.DirectMemberCountPerType) {
		rawNew.DirectMemberCountPerType = rawDesired.DirectMemberCountPerType
	} else {
		rawNew.DirectMemberCountPerType = canonicalizeNewGroupDirectMemberCountPerType(c, rawDesired.DirectMemberCountPerType, rawNew.DirectMemberCountPerType)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DerivedAliases) && dcl.IsEmptyValueIndirect(rawDesired.DerivedAliases) {
		rawNew.DerivedAliases = rawDesired.DerivedAliases
	} else {
		rawNew.DerivedAliases = canonicalizeNewGroupDerivedAliasesSlice(c, rawDesired.DerivedAliases, rawNew.DerivedAliases)
	}

	if dcl.IsEmptyValueIndirect(rawNew.DynamicGroupMetadata) && dcl.IsEmptyValueIndirect(rawDesired.DynamicGroupMetadata) {
		rawNew.DynamicGroupMetadata = rawDesired.DynamicGroupMetadata
	} else {
		rawNew.DynamicGroupMetadata = canonicalizeNewGroupDynamicGroupMetadata(c, rawDesired.DynamicGroupMetadata, rawNew.DynamicGroupMetadata)
	}

	if dcl.IsEmptyValueIndirect(rawNew.PosixGroups) && dcl.IsEmptyValueIndirect(rawDesired.PosixGroups) {
		rawNew.PosixGroups = rawDesired.PosixGroups
	} else {
		rawNew.PosixGroups = canonicalizeNewGroupPosixGroupsSlice(c, rawDesired.PosixGroups, rawNew.PosixGroups)
	}

	rawNew.InitialGroupConfig = rawDesired.InitialGroupConfig

	return rawNew, nil
}

func canonicalizeGroupGroupKey(des, initial *GroupGroupKey, opts ...dcl.ApplyOption) *GroupGroupKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupGroupKey{}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}

	return cDes
}

func canonicalizeGroupGroupKeySlice(des, initial []GroupGroupKey, opts ...dcl.ApplyOption) []GroupGroupKey {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupGroupKey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupGroupKey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupGroupKey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupGroupKey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupGroupKey(c *Client, des, nw *GroupGroupKey) *GroupGroupKey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupGroupKey while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}

	return nw
}

func canonicalizeNewGroupGroupKeySet(c *Client, des, nw []GroupGroupKey) []GroupGroupKey {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupGroupKey
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupGroupKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupGroupKey(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupGroupKeySlice(c *Client, des, nw []GroupGroupKey) []GroupGroupKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupGroupKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupGroupKey(c, &d, &n))
	}

	return items
}

func canonicalizeGroupAdditionalGroupKeys(des, initial *GroupAdditionalGroupKeys, opts ...dcl.ApplyOption) *GroupAdditionalGroupKeys {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupAdditionalGroupKeys{}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}

	return cDes
}

func canonicalizeGroupAdditionalGroupKeysSlice(des, initial []GroupAdditionalGroupKeys, opts ...dcl.ApplyOption) []GroupAdditionalGroupKeys {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupAdditionalGroupKeys, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupAdditionalGroupKeys(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupAdditionalGroupKeys, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupAdditionalGroupKeys(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupAdditionalGroupKeys(c *Client, des, nw *GroupAdditionalGroupKeys) *GroupAdditionalGroupKeys {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupAdditionalGroupKeys while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}

	return nw
}

func canonicalizeNewGroupAdditionalGroupKeysSet(c *Client, des, nw []GroupAdditionalGroupKeys) []GroupAdditionalGroupKeys {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupAdditionalGroupKeys
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupAdditionalGroupKeysNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupAdditionalGroupKeys(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupAdditionalGroupKeysSlice(c *Client, des, nw []GroupAdditionalGroupKeys) []GroupAdditionalGroupKeys {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupAdditionalGroupKeys
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupAdditionalGroupKeys(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDirectMemberCountPerType(des, initial *GroupDirectMemberCountPerType, opts ...dcl.ApplyOption) *GroupDirectMemberCountPerType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDirectMemberCountPerType{}

	return cDes
}

func canonicalizeGroupDirectMemberCountPerTypeSlice(des, initial []GroupDirectMemberCountPerType, opts ...dcl.ApplyOption) []GroupDirectMemberCountPerType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDirectMemberCountPerType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDirectMemberCountPerType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDirectMemberCountPerType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDirectMemberCountPerType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDirectMemberCountPerType(c *Client, des, nw *GroupDirectMemberCountPerType) *GroupDirectMemberCountPerType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDirectMemberCountPerType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewGroupDirectMemberCountPerTypeSet(c *Client, des, nw []GroupDirectMemberCountPerType) []GroupDirectMemberCountPerType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupDirectMemberCountPerType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupDirectMemberCountPerTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupDirectMemberCountPerType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupDirectMemberCountPerTypeSlice(c *Client, des, nw []GroupDirectMemberCountPerType) []GroupDirectMemberCountPerType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDirectMemberCountPerType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDirectMemberCountPerType(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDerivedAliases(des, initial *GroupDerivedAliases, opts ...dcl.ApplyOption) *GroupDerivedAliases {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDerivedAliases{}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}

	return cDes
}

func canonicalizeGroupDerivedAliasesSlice(des, initial []GroupDerivedAliases, opts ...dcl.ApplyOption) []GroupDerivedAliases {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDerivedAliases, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDerivedAliases(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDerivedAliases, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDerivedAliases(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDerivedAliases(c *Client, des, nw *GroupDerivedAliases) *GroupDerivedAliases {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDerivedAliases while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}

	return nw
}

func canonicalizeNewGroupDerivedAliasesSet(c *Client, des, nw []GroupDerivedAliases) []GroupDerivedAliases {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupDerivedAliases
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupDerivedAliasesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupDerivedAliases(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupDerivedAliasesSlice(c *Client, des, nw []GroupDerivedAliases) []GroupDerivedAliases {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDerivedAliases
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDerivedAliases(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadata(des, initial *GroupDynamicGroupMetadata, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadata{}

	cDes.Queries = canonicalizeGroupDynamicGroupMetadataQueriesSlice(des.Queries, initial.Queries, opts...)

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataSlice(des, initial []GroupDynamicGroupMetadata, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadata {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadata, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadata(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadata, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadata(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadata(c *Client, des, nw *GroupDynamicGroupMetadata) *GroupDynamicGroupMetadata {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadata while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Queries = canonicalizeNewGroupDynamicGroupMetadataQueriesSlice(c, des.Queries, nw.Queries)
	nw.Status = canonicalizeNewGroupDynamicGroupMetadataStatus(c, des.Status, nw.Status)

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataSet(c *Client, des, nw []GroupDynamicGroupMetadata) []GroupDynamicGroupMetadata {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupDynamicGroupMetadata
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupDynamicGroupMetadata(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupDynamicGroupMetadataSlice(c *Client, des, nw []GroupDynamicGroupMetadata) []GroupDynamicGroupMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadataQueries(des, initial *GroupDynamicGroupMetadataQueries, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadataQueries {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadataQueries{}

	if dcl.IsZeroValue(des.ResourceType) || (dcl.IsEmptyValueIndirect(des.ResourceType) && dcl.IsEmptyValueIndirect(initial.ResourceType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResourceType = initial.ResourceType
	} else {
		cDes.ResourceType = des.ResourceType
	}
	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		cDes.Query = initial.Query
	} else {
		cDes.Query = des.Query
	}

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataQueriesSlice(des, initial []GroupDynamicGroupMetadataQueries, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadataQueries, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadataQueries(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadataQueries, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadataQueries(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadataQueries(c *Client, des, nw *GroupDynamicGroupMetadataQueries) *GroupDynamicGroupMetadataQueries {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadataQueries while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataQueriesSet(c *Client, des, nw []GroupDynamicGroupMetadataQueries) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupDynamicGroupMetadataQueries
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataQueriesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupDynamicGroupMetadataQueries(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupDynamicGroupMetadataQueriesSlice(c *Client, des, nw []GroupDynamicGroupMetadataQueries) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadataQueries
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadataQueries(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadataStatus(des, initial *GroupDynamicGroupMetadataStatus, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadataStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadataStatus{}

	if dcl.IsZeroValue(des.Status) || (dcl.IsEmptyValueIndirect(des.Status) && dcl.IsEmptyValueIndirect(initial.Status)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Status = initial.Status
	} else {
		cDes.Status = des.Status
	}
	if dcl.IsZeroValue(des.StatusTime) || (dcl.IsEmptyValueIndirect(des.StatusTime) && dcl.IsEmptyValueIndirect(initial.StatusTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.StatusTime = initial.StatusTime
	} else {
		cDes.StatusTime = des.StatusTime
	}

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataStatusSlice(des, initial []GroupDynamicGroupMetadataStatus, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadataStatus {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadataStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadataStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadataStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadataStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadataStatus(c *Client, des, nw *GroupDynamicGroupMetadataStatus) *GroupDynamicGroupMetadataStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadataStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataStatusSet(c *Client, des, nw []GroupDynamicGroupMetadataStatus) []GroupDynamicGroupMetadataStatus {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupDynamicGroupMetadataStatus
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupDynamicGroupMetadataStatus(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupDynamicGroupMetadataStatusSlice(c *Client, des, nw []GroupDynamicGroupMetadataStatus) []GroupDynamicGroupMetadataStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadataStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadataStatus(c, &d, &n))
	}

	return items
}

func canonicalizeGroupPosixGroups(des, initial *GroupPosixGroups, opts ...dcl.ApplyOption) *GroupPosixGroups {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupPosixGroups{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Gid, initial.Gid) || dcl.IsZeroValue(des.Gid) {
		cDes.Gid = initial.Gid
	} else {
		cDes.Gid = des.Gid
	}
	if dcl.StringCanonicalize(des.SystemId, initial.SystemId) || dcl.IsZeroValue(des.SystemId) {
		cDes.SystemId = initial.SystemId
	} else {
		cDes.SystemId = des.SystemId
	}

	return cDes
}

func canonicalizeGroupPosixGroupsSlice(des, initial []GroupPosixGroups, opts ...dcl.ApplyOption) []GroupPosixGroups {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupPosixGroups, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupPosixGroups(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupPosixGroups, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupPosixGroups(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupPosixGroups(c *Client, des, nw *GroupPosixGroups) *GroupPosixGroups {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupPosixGroups while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Gid, nw.Gid) {
		nw.Gid = des.Gid
	}
	if dcl.StringCanonicalize(des.SystemId, nw.SystemId) {
		nw.SystemId = des.SystemId
	}

	return nw
}

func canonicalizeNewGroupPosixGroupsSet(c *Client, des, nw []GroupPosixGroups) []GroupPosixGroups {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []GroupPosixGroups
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareGroupPosixGroupsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewGroupPosixGroups(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewGroupPosixGroupsSlice(c *Client, des, nw []GroupPosixGroups) []GroupPosixGroups {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupPosixGroups
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupPosixGroups(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffGroup(c *Client, desired, actual *Group, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupKey, actual.GroupKey, dcl.DiffInfo{ObjectFunction: compareGroupGroupKeyNewStyle, EmptyObject: EmptyGroupGroupKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalGroupKeys, actual.AdditionalGroupKeys, dcl.DiffInfo{ObjectFunction: compareGroupAdditionalGroupKeysNewStyle, EmptyObject: EmptyGroupAdditionalGroupKeys, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalGroupKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DirectMemberCount, actual.DirectMemberCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DirectMemberCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DirectMemberCountPerType, actual.DirectMemberCountPerType, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareGroupDirectMemberCountPerTypeNewStyle, EmptyObject: EmptyGroupDirectMemberCountPerType, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DirectMemberCountPerType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DerivedAliases, actual.DerivedAliases, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareGroupDerivedAliasesNewStyle, EmptyObject: EmptyGroupDerivedAliases, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DerivedAliases")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DynamicGroupMetadata, actual.DynamicGroupMetadata, dcl.DiffInfo{ObjectFunction: compareGroupDynamicGroupMetadataNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadata, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("DynamicGroupMetadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PosixGroups, actual.PosixGroups, dcl.DiffInfo{ObjectFunction: compareGroupPosixGroupsNewStyle, EmptyObject: EmptyGroupPosixGroups, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("PosixGroups")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitialGroupConfig, actual.InitialGroupConfig, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InitialGroupConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareGroupGroupKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupGroupKey)
	if !ok {
		desiredNotPointer, ok := d.(GroupGroupKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupGroupKey or *GroupGroupKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupGroupKey)
	if !ok {
		actualNotPointer, ok := a.(GroupGroupKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupGroupKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupAdditionalGroupKeysNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupAdditionalGroupKeys)
	if !ok {
		desiredNotPointer, ok := d.(GroupAdditionalGroupKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupAdditionalGroupKeys or *GroupAdditionalGroupKeys", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupAdditionalGroupKeys)
	if !ok {
		actualNotPointer, ok := a.(GroupAdditionalGroupKeys)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupAdditionalGroupKeys", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDirectMemberCountPerTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDirectMemberCountPerType)
	if !ok {
		desiredNotPointer, ok := d.(GroupDirectMemberCountPerType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDirectMemberCountPerType or *GroupDirectMemberCountPerType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDirectMemberCountPerType)
	if !ok {
		actualNotPointer, ok := a.(GroupDirectMemberCountPerType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDirectMemberCountPerType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UserCount, actual.UserCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupCount, actual.GroupCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDerivedAliasesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDerivedAliases)
	if !ok {
		desiredNotPointer, ok := d.(GroupDerivedAliases)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDerivedAliases or *GroupDerivedAliases", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDerivedAliases)
	if !ok {
		actualNotPointer, ok := a.(GroupDerivedAliases)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDerivedAliases", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadata)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadata or *GroupDynamicGroupMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadata)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Queries, actual.Queries, dcl.DiffInfo{ObjectFunction: compareGroupDynamicGroupMetadataQueriesNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadataQueries, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Queries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareGroupDynamicGroupMetadataStatusNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadataStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataQueriesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadataQueries)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadataQueries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataQueries or *GroupDynamicGroupMetadataQueries", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadataQueries)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadataQueries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataQueries", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadataStatus)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadataStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataStatus or *GroupDynamicGroupMetadataStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadataStatus)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadataStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusTime, actual.StatusTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("StatusTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupPosixGroupsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupPosixGroups)
	if !ok {
		desiredNotPointer, ok := d.(GroupPosixGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupPosixGroups or *GroupPosixGroups", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupPosixGroups)
	if !ok {
		actualNotPointer, ok := a.(GroupPosixGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupPosixGroups", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gid, actual.Gid, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Gid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SystemId, actual.SystemId, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SystemId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Group) urlNormalized() *Group {
	normalized := dcl.Copy(*r).(Group)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Parent = dcl.SelfLinkToName(r.Parent)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.InitialGroupConfig = r.InitialGroupConfig
	return &normalized
}

func (r *Group) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateGroup" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Group resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Group) marshal(c *Client) ([]byte, error) {
	m, err := expandGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Group: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGroup decodes JSON responses into the Group resource schema.
func unmarshalGroup(b []byte, c *Client, res *Group) (*Group, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGroup(m, c, res)
}

func unmarshalMapGroup(m map[string]interface{}, c *Client, res *Group) (*Group, error) {

	flattened := flattenGroup(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandGroup expands Group into a JSON request object.
func expandGroup(c *Client, f *Group) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("groups/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandGroupGroupKey(c, f.GroupKey, res); err != nil {
		return nil, fmt.Errorf("error expanding GroupKey into groupKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["groupKey"] = v
	}
	if v, err := expandGroupAdditionalGroupKeysSlice(c, f.AdditionalGroupKeys, res); err != nil {
		return nil, fmt.Errorf("error expanding AdditionalGroupKeys into additionalGroupKeys: %w", err)
	} else if v != nil {
		m["additionalGroupKeys"] = v
	}
	if v := f.Parent; dcl.ValueShouldBeSent(v) {
		m["parent"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandGroupDynamicGroupMetadata(c, f.DynamicGroupMetadata, res); err != nil {
		return nil, fmt.Errorf("error expanding DynamicGroupMetadata into dynamicGroupMetadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dynamicGroupMetadata"] = v
	}
	if v, err := expandGroupPosixGroupsSlice(c, f.PosixGroups, res); err != nil {
		return nil, fmt.Errorf("error expanding PosixGroups into posixGroups: %w", err)
	} else if v != nil {
		m["posixGroups"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding InitialGroupConfig into initialGroupConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["initialGroupConfig"] = v
	}

	return m, nil
}

// flattenGroup flattens Group from a JSON request object into the
// Group type.
func flattenGroup(c *Client, i interface{}, res *Group) *Group {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Group{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.GroupKey = flattenGroupGroupKey(c, m["groupKey"], res)
	resultRes.AdditionalGroupKeys = flattenGroupAdditionalGroupKeysSlice(c, m["additionalGroupKeys"], res)
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.DirectMemberCount = dcl.FlattenInteger(m["directMemberCount"])
	resultRes.DirectMemberCountPerType = flattenGroupDirectMemberCountPerType(c, m["directMemberCountPerType"], res)
	resultRes.DerivedAliases = flattenGroupDerivedAliasesSlice(c, m["derivedAliases"], res)
	resultRes.DynamicGroupMetadata = flattenGroupDynamicGroupMetadata(c, m["dynamicGroupMetadata"], res)
	resultRes.PosixGroups = flattenGroupPosixGroupsSlice(c, m["posixGroups"], res)
	resultRes.InitialGroupConfig = flattenGroupInitialGroupConfigEnum(m["initialGroupConfig"])

	return resultRes
}

// expandGroupGroupKeyMap expands the contents of GroupGroupKey into a JSON
// request object.
func expandGroupGroupKeyMap(c *Client, f map[string]GroupGroupKey, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupGroupKey(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupGroupKeySlice expands the contents of GroupGroupKey into a JSON
// request object.
func expandGroupGroupKeySlice(c *Client, f []GroupGroupKey, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupGroupKey(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupGroupKeyMap flattens the contents of GroupGroupKey from a JSON
// response object.
func flattenGroupGroupKeyMap(c *Client, i interface{}, res *Group) map[string]GroupGroupKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupGroupKey{}
	}

	if len(a) == 0 {
		return map[string]GroupGroupKey{}
	}

	items := make(map[string]GroupGroupKey)
	for k, item := range a {
		items[k] = *flattenGroupGroupKey(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupGroupKeySlice flattens the contents of GroupGroupKey from a JSON
// response object.
func flattenGroupGroupKeySlice(c *Client, i interface{}, res *Group) []GroupGroupKey {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupGroupKey{}
	}

	if len(a) == 0 {
		return []GroupGroupKey{}
	}

	items := make([]GroupGroupKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupGroupKey(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupGroupKey expands an instance of GroupGroupKey into a JSON
// request object.
func expandGroupGroupKey(c *Client, f *GroupGroupKey, res *Group) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}

	return m, nil
}

// flattenGroupGroupKey flattens an instance of GroupGroupKey from a JSON
// response object.
func flattenGroupGroupKey(c *Client, i interface{}, res *Group) *GroupGroupKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupGroupKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupGroupKey
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandGroupAdditionalGroupKeysMap expands the contents of GroupAdditionalGroupKeys into a JSON
// request object.
func expandGroupAdditionalGroupKeysMap(c *Client, f map[string]GroupAdditionalGroupKeys, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupAdditionalGroupKeys(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupAdditionalGroupKeysSlice expands the contents of GroupAdditionalGroupKeys into a JSON
// request object.
func expandGroupAdditionalGroupKeysSlice(c *Client, f []GroupAdditionalGroupKeys, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupAdditionalGroupKeys(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupAdditionalGroupKeysMap flattens the contents of GroupAdditionalGroupKeys from a JSON
// response object.
func flattenGroupAdditionalGroupKeysMap(c *Client, i interface{}, res *Group) map[string]GroupAdditionalGroupKeys {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupAdditionalGroupKeys{}
	}

	if len(a) == 0 {
		return map[string]GroupAdditionalGroupKeys{}
	}

	items := make(map[string]GroupAdditionalGroupKeys)
	for k, item := range a {
		items[k] = *flattenGroupAdditionalGroupKeys(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupAdditionalGroupKeysSlice flattens the contents of GroupAdditionalGroupKeys from a JSON
// response object.
func flattenGroupAdditionalGroupKeysSlice(c *Client, i interface{}, res *Group) []GroupAdditionalGroupKeys {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupAdditionalGroupKeys{}
	}

	if len(a) == 0 {
		return []GroupAdditionalGroupKeys{}
	}

	items := make([]GroupAdditionalGroupKeys, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupAdditionalGroupKeys(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupAdditionalGroupKeys expands an instance of GroupAdditionalGroupKeys into a JSON
// request object.
func expandGroupAdditionalGroupKeys(c *Client, f *GroupAdditionalGroupKeys, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}

	return m, nil
}

// flattenGroupAdditionalGroupKeys flattens an instance of GroupAdditionalGroupKeys from a JSON
// response object.
func flattenGroupAdditionalGroupKeys(c *Client, i interface{}, res *Group) *GroupAdditionalGroupKeys {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupAdditionalGroupKeys{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupAdditionalGroupKeys
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandGroupDirectMemberCountPerTypeMap expands the contents of GroupDirectMemberCountPerType into a JSON
// request object.
func expandGroupDirectMemberCountPerTypeMap(c *Client, f map[string]GroupDirectMemberCountPerType, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDirectMemberCountPerType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDirectMemberCountPerTypeSlice expands the contents of GroupDirectMemberCountPerType into a JSON
// request object.
func expandGroupDirectMemberCountPerTypeSlice(c *Client, f []GroupDirectMemberCountPerType, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDirectMemberCountPerType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDirectMemberCountPerTypeMap flattens the contents of GroupDirectMemberCountPerType from a JSON
// response object.
func flattenGroupDirectMemberCountPerTypeMap(c *Client, i interface{}, res *Group) map[string]GroupDirectMemberCountPerType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDirectMemberCountPerType{}
	}

	if len(a) == 0 {
		return map[string]GroupDirectMemberCountPerType{}
	}

	items := make(map[string]GroupDirectMemberCountPerType)
	for k, item := range a {
		items[k] = *flattenGroupDirectMemberCountPerType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupDirectMemberCountPerTypeSlice flattens the contents of GroupDirectMemberCountPerType from a JSON
// response object.
func flattenGroupDirectMemberCountPerTypeSlice(c *Client, i interface{}, res *Group) []GroupDirectMemberCountPerType {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDirectMemberCountPerType{}
	}

	if len(a) == 0 {
		return []GroupDirectMemberCountPerType{}
	}

	items := make([]GroupDirectMemberCountPerType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDirectMemberCountPerType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupDirectMemberCountPerType expands an instance of GroupDirectMemberCountPerType into a JSON
// request object.
func expandGroupDirectMemberCountPerType(c *Client, f *GroupDirectMemberCountPerType, res *Group) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenGroupDirectMemberCountPerType flattens an instance of GroupDirectMemberCountPerType from a JSON
// response object.
func flattenGroupDirectMemberCountPerType(c *Client, i interface{}, res *Group) *GroupDirectMemberCountPerType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDirectMemberCountPerType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDirectMemberCountPerType
	}
	r.UserCount = dcl.FlattenInteger(m["userCount"])
	r.GroupCount = dcl.FlattenInteger(m["groupCount"])

	return r
}

// expandGroupDerivedAliasesMap expands the contents of GroupDerivedAliases into a JSON
// request object.
func expandGroupDerivedAliasesMap(c *Client, f map[string]GroupDerivedAliases, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDerivedAliases(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDerivedAliasesSlice expands the contents of GroupDerivedAliases into a JSON
// request object.
func expandGroupDerivedAliasesSlice(c *Client, f []GroupDerivedAliases, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDerivedAliases(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDerivedAliasesMap flattens the contents of GroupDerivedAliases from a JSON
// response object.
func flattenGroupDerivedAliasesMap(c *Client, i interface{}, res *Group) map[string]GroupDerivedAliases {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDerivedAliases{}
	}

	if len(a) == 0 {
		return map[string]GroupDerivedAliases{}
	}

	items := make(map[string]GroupDerivedAliases)
	for k, item := range a {
		items[k] = *flattenGroupDerivedAliases(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupDerivedAliasesSlice flattens the contents of GroupDerivedAliases from a JSON
// response object.
func flattenGroupDerivedAliasesSlice(c *Client, i interface{}, res *Group) []GroupDerivedAliases {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDerivedAliases{}
	}

	if len(a) == 0 {
		return []GroupDerivedAliases{}
	}

	items := make([]GroupDerivedAliases, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDerivedAliases(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupDerivedAliases expands an instance of GroupDerivedAliases into a JSON
// request object.
func expandGroupDerivedAliases(c *Client, f *GroupDerivedAliases, res *Group) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}

	return m, nil
}

// flattenGroupDerivedAliases flattens an instance of GroupDerivedAliases from a JSON
// response object.
func flattenGroupDerivedAliases(c *Client, i interface{}, res *Group) *GroupDerivedAliases {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDerivedAliases{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDerivedAliases
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandGroupDynamicGroupMetadataMap expands the contents of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadataMap(c *Client, f map[string]GroupDynamicGroupMetadata, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadata(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataSlice expands the contents of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadataSlice(c *Client, f []GroupDynamicGroupMetadata, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadata(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataMap flattens the contents of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadataMap(c *Client, i interface{}, res *Group) map[string]GroupDynamicGroupMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadata{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadata{}
	}

	items := make(map[string]GroupDynamicGroupMetadata)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadata(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupDynamicGroupMetadataSlice flattens the contents of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadataSlice(c *Client, i interface{}, res *Group) []GroupDynamicGroupMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadata{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadata{}
	}

	items := make([]GroupDynamicGroupMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadata(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupDynamicGroupMetadata expands an instance of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadata(c *Client, f *GroupDynamicGroupMetadata, res *Group) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGroupDynamicGroupMetadataQueriesSlice(c, f.Queries, res); err != nil {
		return nil, fmt.Errorf("error expanding Queries into queries: %w", err)
	} else if v != nil {
		m["queries"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadata flattens an instance of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadata(c *Client, i interface{}, res *Group) *GroupDynamicGroupMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadata
	}
	r.Queries = flattenGroupDynamicGroupMetadataQueriesSlice(c, m["queries"], res)
	r.Status = flattenGroupDynamicGroupMetadataStatus(c, m["status"], res)

	return r
}

// expandGroupDynamicGroupMetadataQueriesMap expands the contents of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueriesMap(c *Client, f map[string]GroupDynamicGroupMetadataQueries, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadataQueries(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataQueriesSlice expands the contents of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueriesSlice(c *Client, f []GroupDynamicGroupMetadataQueries, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadataQueries(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataQueriesMap flattens the contents of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesMap(c *Client, i interface{}, res *Group) map[string]GroupDynamicGroupMetadataQueries {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataQueries{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataQueries{}
	}

	items := make(map[string]GroupDynamicGroupMetadataQueries)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataQueries(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesSlice flattens the contents of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesSlice(c *Client, i interface{}, res *Group) []GroupDynamicGroupMetadataQueries {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataQueries{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataQueries{}
	}

	items := make([]GroupDynamicGroupMetadataQueries, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataQueries(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupDynamicGroupMetadataQueries expands an instance of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueries(c *Client, f *GroupDynamicGroupMetadataQueries, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadataQueries flattens an instance of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueries(c *Client, i interface{}, res *Group) *GroupDynamicGroupMetadataQueries {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadataQueries{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadataQueries
	}
	r.ResourceType = flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(m["resourceType"])
	r.Query = dcl.FlattenString(m["query"])

	return r
}

// expandGroupDynamicGroupMetadataStatusMap expands the contents of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatusMap(c *Client, f map[string]GroupDynamicGroupMetadataStatus, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadataStatus(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataStatusSlice expands the contents of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatusSlice(c *Client, f []GroupDynamicGroupMetadataStatus, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadataStatus(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataStatusMap flattens the contents of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusMap(c *Client, i interface{}, res *Group) map[string]GroupDynamicGroupMetadataStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataStatus{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataStatus{}
	}

	items := make(map[string]GroupDynamicGroupMetadataStatus)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataStatus(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusSlice flattens the contents of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusSlice(c *Client, i interface{}, res *Group) []GroupDynamicGroupMetadataStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataStatus{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataStatus{}
	}

	items := make([]GroupDynamicGroupMetadataStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataStatus(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupDynamicGroupMetadataStatus expands an instance of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatus(c *Client, f *GroupDynamicGroupMetadataStatus, res *Group) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusTime; !dcl.IsEmptyValueIndirect(v) {
		m["statusTime"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadataStatus flattens an instance of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatus(c *Client, i interface{}, res *Group) *GroupDynamicGroupMetadataStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadataStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadataStatus
	}
	r.Status = flattenGroupDynamicGroupMetadataStatusStatusEnum(m["status"])
	r.StatusTime = dcl.FlattenString(m["statusTime"])

	return r
}

// expandGroupPosixGroupsMap expands the contents of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroupsMap(c *Client, f map[string]GroupPosixGroups, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupPosixGroups(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupPosixGroupsSlice expands the contents of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroupsSlice(c *Client, f []GroupPosixGroups, res *Group) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupPosixGroups(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupPosixGroupsMap flattens the contents of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroupsMap(c *Client, i interface{}, res *Group) map[string]GroupPosixGroups {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupPosixGroups{}
	}

	if len(a) == 0 {
		return map[string]GroupPosixGroups{}
	}

	items := make(map[string]GroupPosixGroups)
	for k, item := range a {
		items[k] = *flattenGroupPosixGroups(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenGroupPosixGroupsSlice flattens the contents of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroupsSlice(c *Client, i interface{}, res *Group) []GroupPosixGroups {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupPosixGroups{}
	}

	if len(a) == 0 {
		return []GroupPosixGroups{}
	}

	items := make([]GroupPosixGroups, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupPosixGroups(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandGroupPosixGroups expands an instance of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroups(c *Client, f *GroupPosixGroups, res *Group) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Gid; !dcl.IsEmptyValueIndirect(v) {
		m["gid"] = v
	}
	if v := f.SystemId; !dcl.IsEmptyValueIndirect(v) {
		m["systemId"] = v
	}

	return m, nil
}

// flattenGroupPosixGroups flattens an instance of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroups(c *Client, i interface{}, res *Group) *GroupPosixGroups {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupPosixGroups{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupPosixGroups
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Gid = dcl.FlattenString(m["gid"])
	r.SystemId = dcl.FlattenString(m["systemId"])

	return r
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumMap flattens the contents of GroupDynamicGroupMetadataQueriesResourceTypeEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumMap(c *Client, i interface{}, res *Group) map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	items := make(map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumSlice flattens the contents of GroupDynamicGroupMetadataQueriesResourceTypeEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumSlice(c *Client, i interface{}, res *Group) []GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	items := make([]GroupDynamicGroupMetadataQueriesResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GroupDynamicGroupMetadataQueriesResourceTypeEnum with the same value as that string.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(i interface{}) *GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GroupDynamicGroupMetadataQueriesResourceTypeEnumRef(s)
}

// flattenGroupDynamicGroupMetadataStatusStatusEnumMap flattens the contents of GroupDynamicGroupMetadataStatusStatusEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusStatusEnumMap(c *Client, i interface{}, res *Group) map[string]GroupDynamicGroupMetadataStatusStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	items := make(map[string]GroupDynamicGroupMetadataStatusStatusEnum)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataStatusStatusEnum(item.(interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusStatusEnumSlice flattens the contents of GroupDynamicGroupMetadataStatusStatusEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusStatusEnumSlice(c *Client, i interface{}, res *Group) []GroupDynamicGroupMetadataStatusStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	items := make([]GroupDynamicGroupMetadataStatusStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataStatusStatusEnum(item.(interface{})))
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusStatusEnum asserts that an interface is a string, and returns a
// pointer to a *GroupDynamicGroupMetadataStatusStatusEnum with the same value as that string.
func flattenGroupDynamicGroupMetadataStatusStatusEnum(i interface{}) *GroupDynamicGroupMetadataStatusStatusEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GroupDynamicGroupMetadataStatusStatusEnumRef(s)
}

// flattenGroupInitialGroupConfigEnumMap flattens the contents of GroupInitialGroupConfigEnum from a JSON
// response object.
func flattenGroupInitialGroupConfigEnumMap(c *Client, i interface{}, res *Group) map[string]GroupInitialGroupConfigEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupInitialGroupConfigEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupInitialGroupConfigEnum{}
	}

	items := make(map[string]GroupInitialGroupConfigEnum)
	for k, item := range a {
		items[k] = *flattenGroupInitialGroupConfigEnum(item.(interface{}))
	}

	return items
}

// flattenGroupInitialGroupConfigEnumSlice flattens the contents of GroupInitialGroupConfigEnum from a JSON
// response object.
func flattenGroupInitialGroupConfigEnumSlice(c *Client, i interface{}, res *Group) []GroupInitialGroupConfigEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupInitialGroupConfigEnum{}
	}

	if len(a) == 0 {
		return []GroupInitialGroupConfigEnum{}
	}

	items := make([]GroupInitialGroupConfigEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupInitialGroupConfigEnum(item.(interface{})))
	}

	return items
}

// flattenGroupInitialGroupConfigEnum asserts that an interface is a string, and returns a
// pointer to a *GroupInitialGroupConfigEnum with the same value as that string.
func flattenGroupInitialGroupConfigEnum(i interface{}) *GroupInitialGroupConfigEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return GroupInitialGroupConfigEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Group) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGroup(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type groupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         groupApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToGroupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]groupDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []groupDiff
	// For each operation name, create a groupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := groupDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToGroupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToGroupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (groupApiOperation, error) {
	switch opName {

	case "updateGroupUpdateGroupOperation":
		return &updateGroupUpdateGroupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractGroupFields(r *Group) error {
	vGroupKey := r.GroupKey
	if vGroupKey == nil {
		// note: explicitly not the empty object.
		vGroupKey = &GroupGroupKey{}
	}
	if err := extractGroupGroupKeyFields(r, vGroupKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGroupKey) {
		r.GroupKey = vGroupKey
	}
	vDirectMemberCountPerType := r.DirectMemberCountPerType
	if vDirectMemberCountPerType == nil {
		// note: explicitly not the empty object.
		vDirectMemberCountPerType = &GroupDirectMemberCountPerType{}
	}
	if err := extractGroupDirectMemberCountPerTypeFields(r, vDirectMemberCountPerType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDirectMemberCountPerType) {
		r.DirectMemberCountPerType = vDirectMemberCountPerType
	}
	vDynamicGroupMetadata := r.DynamicGroupMetadata
	if vDynamicGroupMetadata == nil {
		// note: explicitly not the empty object.
		vDynamicGroupMetadata = &GroupDynamicGroupMetadata{}
	}
	if err := extractGroupDynamicGroupMetadataFields(r, vDynamicGroupMetadata); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDynamicGroupMetadata) {
		r.DynamicGroupMetadata = vDynamicGroupMetadata
	}
	return nil
}
func extractGroupGroupKeyFields(r *Group, o *GroupGroupKey) error {
	return nil
}
func extractGroupAdditionalGroupKeysFields(r *Group, o *GroupAdditionalGroupKeys) error {
	return nil
}
func extractGroupDirectMemberCountPerTypeFields(r *Group, o *GroupDirectMemberCountPerType) error {
	return nil
}
func extractGroupDerivedAliasesFields(r *Group, o *GroupDerivedAliases) error {
	return nil
}
func extractGroupDynamicGroupMetadataFields(r *Group, o *GroupDynamicGroupMetadata) error {
	vStatus := o.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &GroupDynamicGroupMetadataStatus{}
	}
	if err := extractGroupDynamicGroupMetadataStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStatus) {
		o.Status = vStatus
	}
	return nil
}
func extractGroupDynamicGroupMetadataQueriesFields(r *Group, o *GroupDynamicGroupMetadataQueries) error {
	return nil
}
func extractGroupDynamicGroupMetadataStatusFields(r *Group, o *GroupDynamicGroupMetadataStatus) error {
	return nil
}
func extractGroupPosixGroupsFields(r *Group, o *GroupPosixGroups) error {
	return nil
}

func postReadExtractGroupFields(r *Group) error {
	vGroupKey := r.GroupKey
	if vGroupKey == nil {
		// note: explicitly not the empty object.
		vGroupKey = &GroupGroupKey{}
	}
	if err := postReadExtractGroupGroupKeyFields(r, vGroupKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGroupKey) {
		r.GroupKey = vGroupKey
	}
	vDirectMemberCountPerType := r.DirectMemberCountPerType
	if vDirectMemberCountPerType == nil {
		// note: explicitly not the empty object.
		vDirectMemberCountPerType = &GroupDirectMemberCountPerType{}
	}
	if err := postReadExtractGroupDirectMemberCountPerTypeFields(r, vDirectMemberCountPerType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDirectMemberCountPerType) {
		r.DirectMemberCountPerType = vDirectMemberCountPerType
	}
	vDynamicGroupMetadata := r.DynamicGroupMetadata
	if vDynamicGroupMetadata == nil {
		// note: explicitly not the empty object.
		vDynamicGroupMetadata = &GroupDynamicGroupMetadata{}
	}
	if err := postReadExtractGroupDynamicGroupMetadataFields(r, vDynamicGroupMetadata); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDynamicGroupMetadata) {
		r.DynamicGroupMetadata = vDynamicGroupMetadata
	}
	return nil
}
func postReadExtractGroupGroupKeyFields(r *Group, o *GroupGroupKey) error {
	return nil
}
func postReadExtractGroupAdditionalGroupKeysFields(r *Group, o *GroupAdditionalGroupKeys) error {
	return nil
}
func postReadExtractGroupDirectMemberCountPerTypeFields(r *Group, o *GroupDirectMemberCountPerType) error {
	return nil
}
func postReadExtractGroupDerivedAliasesFields(r *Group, o *GroupDerivedAliases) error {
	return nil
}
func postReadExtractGroupDynamicGroupMetadataFields(r *Group, o *GroupDynamicGroupMetadata) error {
	vStatus := o.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &GroupDynamicGroupMetadataStatus{}
	}
	if err := extractGroupDynamicGroupMetadataStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStatus) {
		o.Status = vStatus
	}
	return nil
}
func postReadExtractGroupDynamicGroupMetadataQueriesFields(r *Group, o *GroupDynamicGroupMetadataQueries) error {
	return nil
}
func postReadExtractGroupDynamicGroupMetadataStatusFields(r *Group, o *GroupDynamicGroupMetadataStatus) error {
	return nil
}
func postReadExtractGroupPosixGroupsFields(r *Group, o *GroupPosixGroups) error {
	return nil
}
