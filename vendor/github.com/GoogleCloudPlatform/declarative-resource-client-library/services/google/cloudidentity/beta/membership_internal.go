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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Membership) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"DisplayName"}, r.DisplayName); err != nil {
		return err
	}
	if err := dcl.Required(r, "preferredMemberKey"); err != nil {
		return err
	}
	if err := dcl.Required(r, "roles"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Group, "Group"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PreferredMemberKey) {
		if err := r.PreferredMemberKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DisplayName) {
		if err := r.DisplayName.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MemberKey) {
		if err := r.MemberKey.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipPreferredMemberKey) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *MembershipRoles) validate() error {
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.ExpiryDetail) {
		if err := r.ExpiryDetail.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RestrictionEvaluations) {
		if err := r.RestrictionEvaluations.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipRolesExpiryDetail) validate() error {
	return nil
}
func (r *MembershipRolesRestrictionEvaluations) validate() error {
	if !dcl.IsEmptyValueIndirect(r.MemberRestrictionEvaluation) {
		if err := r.MemberRestrictionEvaluation.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) validate() error {
	return nil
}
func (r *MembershipDisplayName) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"GivenName"}, r.GivenName); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FamilyName"}, r.FamilyName); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FullName"}, r.FullName); err != nil {
		return err
	}
	return nil
}
func (r *MembershipMemberKey) validate() error {
	return nil
}
func (r *Membership) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudidentity.googleapis.com/v1beta1/", params)
}

func (r *Membership) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
		"name":  dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{group}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Membership) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
	}
	return dcl.URL("groups/{{group}}/memberships", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
	}
	return dcl.URL("groups/{{group}}/memberships/", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"group": dcl.ValueOrEmptyString(nr.Group),
		"name":  dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{group}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

// membershipApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type membershipApiOperation interface {
	do(context.Context, *Membership, *Client) error
}

// newUpdateMembershipUpdateMembershipRequest creates a request for an
// Membership resource's UpdateMembership update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateMembershipUpdateMembershipRequest(ctx context.Context, f *Membership, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := dcl.DeriveField("groups/%s/memberships/%s", f.Name, dcl.SelfLinkToName(f.Group), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v, err := expandMembershipRolesSlice(c, f.Roles, res); err != nil {
		return nil, fmt.Errorf("error expanding Roles into roles: %w", err)
	} else if v != nil {
		req["roles"] = v
	}
	if v, err := expandMembershipMemberKey(c, f.MemberKey, res); err != nil {
		return nil, fmt.Errorf("error expanding MemberKey into memberKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["memberKey"] = v
	}
	return req, nil
}

// marshalUpdateMembershipUpdateMembershipRequest converts the update into
// the final JSON request body.
func marshalUpdateMembershipUpdateMembershipRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateMembershipUpdateMembershipOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listMembershipRaw(ctx context.Context, r *Membership, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != MembershipMaxPage {
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

type listMembershipOperation struct {
	Memberships []map[string]interface{} `json:"memberships"`
	Token       string                   `json:"nextPageToken"`
}

func (c *Client) listMembership(ctx context.Context, r *Membership, pageToken string, pageSize int32) ([]*Membership, string, error) {
	b, err := c.listMembershipRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listMembershipOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Membership
	for _, v := range m.Memberships {
		res, err := unmarshalMapMembership(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Group = r.Group
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllMembership(ctx context.Context, f func(*Membership) bool, resources []*Membership) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteMembership(ctx, res)
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

type deleteMembershipOperation struct{}

func (op *deleteMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	r, err := c.GetMembership(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Membership not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetMembership checking for existence. error: %v", err)
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
		_, err := c.GetMembership(ctx, r)
		if dcl.IsNotFound(err) {
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
type createMembershipOperation struct {
	response map[string]interface{}
}

func (op *createMembershipOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a Membership with the wrong Name.
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

	if _, err := c.GetMembership(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getMembershipRaw(ctx context.Context, r *Membership) ([]byte, error) {

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

func (c *Client) membershipDiffsForRawDesired(ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (initial, desired *Membership, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Membership
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Membership); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Membership, got %T", sh)
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
		desired, err := canonicalizeMembershipDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetMembership(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Membership resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Membership resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Membership resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Membership: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Membership: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractMembershipFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeMembershipInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Membership: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeMembershipDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Membership: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffMembership(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeMembershipInitialState(rawInitial, rawDesired *Membership) (*Membership, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.DisplayName) {
		// Check if anything else is set.
		if dcl.AnySet() {
			rawInitial.DisplayName = EmptyMembershipDisplayName
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeMembershipDesiredState(rawDesired, rawInitial *Membership, opts ...dcl.ApplyOption) (*Membership, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.PreferredMemberKey = canonicalizeMembershipPreferredMemberKey(rawDesired.PreferredMemberKey, nil, opts...)
		rawDesired.DisplayName = canonicalizeMembershipDisplayName(rawDesired.DisplayName, nil, opts...)
		rawDesired.MemberKey = canonicalizeMembershipMemberKey(rawDesired.MemberKey, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Membership{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.PreferredMemberKey = canonicalizeMembershipPreferredMemberKey(rawDesired.PreferredMemberKey, rawInitial.PreferredMemberKey, opts...)
	canonicalDesired.Roles = canonicalizeMembershipRolesSlice(rawDesired.Roles, rawInitial.Roles, opts...)
	canonicalDesired.MemberKey = canonicalizeMembershipMemberKey(rawDesired.MemberKey, rawInitial.MemberKey, opts...)
	if dcl.NameToSelfLink(rawDesired.Group, rawInitial.Group) {
		canonicalDesired.Group = rawInitial.Group
	} else {
		canonicalDesired.Group = rawDesired.Group
	}

	if canonicalDesired.DisplayName != nil {
		// Check if anything else is set.
		if dcl.AnySet() {
			canonicalDesired.DisplayName = EmptyMembershipDisplayName
		}
	}

	return canonicalDesired, nil
}

func canonicalizeMembershipNewState(c *Client, rawNew, rawDesired *Membership) (*Membership, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PreferredMemberKey) && dcl.IsEmptyValueIndirect(rawDesired.PreferredMemberKey) {
		rawNew.PreferredMemberKey = rawDesired.PreferredMemberKey
	} else {
		rawNew.PreferredMemberKey = canonicalizeNewMembershipPreferredMemberKey(c, rawDesired.PreferredMemberKey, rawNew.PreferredMemberKey)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Roles) && dcl.IsEmptyValueIndirect(rawDesired.Roles) {
		rawNew.Roles = rawDesired.Roles
	} else {
		rawNew.Roles = canonicalizeNewMembershipRolesSet(c, rawDesired.Roles, rawNew.Roles)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Type) && dcl.IsEmptyValueIndirect(rawDesired.Type) {
		rawNew.Type = rawDesired.Type
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeliverySetting) && dcl.IsEmptyValueIndirect(rawDesired.DeliverySetting) {
		rawNew.DeliverySetting = rawDesired.DeliverySetting
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		rawNew.DisplayName = canonicalizeNewMembershipDisplayName(c, rawDesired.DisplayName, rawNew.DisplayName)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MemberKey) && dcl.IsEmptyValueIndirect(rawDesired.MemberKey) {
		rawNew.MemberKey = rawDesired.MemberKey
	} else {
		rawNew.MemberKey = canonicalizeNewMembershipMemberKey(c, rawDesired.MemberKey, rawNew.MemberKey)
	}

	rawNew.Group = rawDesired.Group

	return rawNew, nil
}

func canonicalizeMembershipPreferredMemberKey(des, initial *MembershipPreferredMemberKey, opts ...dcl.ApplyOption) *MembershipPreferredMemberKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipPreferredMemberKey{}

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

func canonicalizeMembershipPreferredMemberKeySlice(des, initial []MembershipPreferredMemberKey, opts ...dcl.ApplyOption) []MembershipPreferredMemberKey {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipPreferredMemberKey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipPreferredMemberKey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipPreferredMemberKey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipPreferredMemberKey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipPreferredMemberKey(c *Client, des, nw *MembershipPreferredMemberKey) *MembershipPreferredMemberKey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipPreferredMemberKey while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewMembershipPreferredMemberKeySet(c *Client, des, nw []MembershipPreferredMemberKey) []MembershipPreferredMemberKey {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipPreferredMemberKey
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipPreferredMemberKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipPreferredMemberKey(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipPreferredMemberKeySlice(c *Client, des, nw []MembershipPreferredMemberKey) []MembershipPreferredMemberKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipPreferredMemberKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipPreferredMemberKey(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRoles(des, initial *MembershipRoles, opts ...dcl.ApplyOption) *MembershipRoles {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRoles{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	cDes.ExpiryDetail = canonicalizeMembershipRolesExpiryDetail(des.ExpiryDetail, initial.ExpiryDetail, opts...)
	cDes.RestrictionEvaluations = canonicalizeMembershipRolesRestrictionEvaluations(des.RestrictionEvaluations, initial.RestrictionEvaluations, opts...)

	return cDes
}

func canonicalizeMembershipRolesSlice(des, initial []MembershipRoles, opts ...dcl.ApplyOption) []MembershipRoles {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRoles, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRoles(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRoles, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRoles(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRoles(c *Client, des, nw *MembershipRoles) *MembershipRoles {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRoles while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	nw.ExpiryDetail = canonicalizeNewMembershipRolesExpiryDetail(c, des.ExpiryDetail, nw.ExpiryDetail)
	nw.RestrictionEvaluations = canonicalizeNewMembershipRolesRestrictionEvaluations(c, des.RestrictionEvaluations, nw.RestrictionEvaluations)

	return nw
}

func canonicalizeNewMembershipRolesSet(c *Client, des, nw []MembershipRoles) []MembershipRoles {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipRoles
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipRolesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipRoles(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipRolesSlice(c *Client, des, nw []MembershipRoles) []MembershipRoles {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRoles
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRoles(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesExpiryDetail(des, initial *MembershipRolesExpiryDetail, opts ...dcl.ApplyOption) *MembershipRolesExpiryDetail {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesExpiryDetail{}

	if dcl.IsZeroValue(des.ExpireTime) || (dcl.IsEmptyValueIndirect(des.ExpireTime) && dcl.IsEmptyValueIndirect(initial.ExpireTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExpireTime = initial.ExpireTime
	} else {
		cDes.ExpireTime = des.ExpireTime
	}

	return cDes
}

func canonicalizeMembershipRolesExpiryDetailSlice(des, initial []MembershipRolesExpiryDetail, opts ...dcl.ApplyOption) []MembershipRolesExpiryDetail {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesExpiryDetail, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesExpiryDetail(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesExpiryDetail, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesExpiryDetail(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesExpiryDetail(c *Client, des, nw *MembershipRolesExpiryDetail) *MembershipRolesExpiryDetail {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesExpiryDetail while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipRolesExpiryDetailSet(c *Client, des, nw []MembershipRolesExpiryDetail) []MembershipRolesExpiryDetail {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipRolesExpiryDetail
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipRolesExpiryDetailNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipRolesExpiryDetail(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipRolesExpiryDetailSlice(c *Client, des, nw []MembershipRolesExpiryDetail) []MembershipRolesExpiryDetail {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesExpiryDetail
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesExpiryDetail(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesRestrictionEvaluations(des, initial *MembershipRolesRestrictionEvaluations, opts ...dcl.ApplyOption) *MembershipRolesRestrictionEvaluations {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesRestrictionEvaluations{}

	cDes.MemberRestrictionEvaluation = canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(des.MemberRestrictionEvaluation, initial.MemberRestrictionEvaluation, opts...)

	return cDes
}

func canonicalizeMembershipRolesRestrictionEvaluationsSlice(des, initial []MembershipRolesRestrictionEvaluations, opts ...dcl.ApplyOption) []MembershipRolesRestrictionEvaluations {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesRestrictionEvaluations, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesRestrictionEvaluations(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesRestrictionEvaluations, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesRestrictionEvaluations(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesRestrictionEvaluations(c *Client, des, nw *MembershipRolesRestrictionEvaluations) *MembershipRolesRestrictionEvaluations {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesRestrictionEvaluations while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MemberRestrictionEvaluation = canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, des.MemberRestrictionEvaluation, nw.MemberRestrictionEvaluation)

	return nw
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsSet(c *Client, des, nw []MembershipRolesRestrictionEvaluations) []MembershipRolesRestrictionEvaluations {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipRolesRestrictionEvaluations
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipRolesRestrictionEvaluationsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluations(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsSlice(c *Client, des, nw []MembershipRolesRestrictionEvaluations) []MembershipRolesRestrictionEvaluations {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesRestrictionEvaluations
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluations(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(des, initial *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, opts ...dcl.ApplyOption) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}

	return cDes
}

func canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(des, initial []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, opts ...dcl.ApplyOption) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, des, nw *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSet(c *Client, des, nw []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, des, nw []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipDisplayName(des, initial *MembershipDisplayName, opts ...dcl.ApplyOption) *MembershipDisplayName {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.GivenName != nil || (initial != nil && initial.GivenName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.GivenName = nil
			if initial != nil {
				initial.GivenName = nil
			}
		}
	}

	if des.FamilyName != nil || (initial != nil && initial.FamilyName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.FamilyName = nil
			if initial != nil {
				initial.FamilyName = nil
			}
		}
	}

	if des.FullName != nil || (initial != nil && initial.FullName != nil) {
		// Check if anything else is set.
		if dcl.AnySet() {
			des.FullName = nil
			if initial != nil {
				initial.FullName = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipDisplayName{}

	return cDes
}

func canonicalizeMembershipDisplayNameSlice(des, initial []MembershipDisplayName, opts ...dcl.ApplyOption) []MembershipDisplayName {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipDisplayName, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipDisplayName(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipDisplayName, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipDisplayName(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipDisplayName(c *Client, des, nw *MembershipDisplayName) *MembershipDisplayName {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipDisplayName while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.GivenName, nw.GivenName) {
		nw.GivenName = des.GivenName
	}
	if dcl.StringCanonicalize(des.FamilyName, nw.FamilyName) {
		nw.FamilyName = des.FamilyName
	}
	if dcl.StringCanonicalize(des.FullName, nw.FullName) {
		nw.FullName = des.FullName
	}

	return nw
}

func canonicalizeNewMembershipDisplayNameSet(c *Client, des, nw []MembershipDisplayName) []MembershipDisplayName {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipDisplayName
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipDisplayNameNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipDisplayName(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipDisplayNameSlice(c *Client, des, nw []MembershipDisplayName) []MembershipDisplayName {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipDisplayName
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipDisplayName(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipMemberKey(des, initial *MembershipMemberKey, opts ...dcl.ApplyOption) *MembershipMemberKey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipMemberKey{}

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

func canonicalizeMembershipMemberKeySlice(des, initial []MembershipMemberKey, opts ...dcl.ApplyOption) []MembershipMemberKey {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipMemberKey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipMemberKey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipMemberKey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipMemberKey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipMemberKey(c *Client, des, nw *MembershipMemberKey) *MembershipMemberKey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipMemberKey while comparing non-nil desired to nil actual.  Returning desired object.")
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

func canonicalizeNewMembershipMemberKeySet(c *Client, des, nw []MembershipMemberKey) []MembershipMemberKey {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipMemberKey
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipMemberKeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipMemberKey(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipMemberKeySlice(c *Client, des, nw []MembershipMemberKey) []MembershipMemberKey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipMemberKey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipMemberKey(c, &d, &n))
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
func diffMembership(c *Client, desired, actual *Membership, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PreferredMemberKey, actual.PreferredMemberKey, dcl.DiffInfo{ObjectFunction: compareMembershipPreferredMemberKeyNewStyle, EmptyObject: EmptyMembershipPreferredMemberKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PreferredMemberKey")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Roles, actual.Roles, dcl.DiffInfo{Type: "Set", ObjectFunction: compareMembershipRolesNewStyle, EmptyObject: EmptyMembershipRoles, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Roles")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DeliverySetting, actual.DeliverySetting, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeliverySetting")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMembershipDisplayNameNewStyle, EmptyObject: EmptyMembershipDisplayName, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemberKey, actual.MemberKey, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareMembershipMemberKeyNewStyle, EmptyObject: EmptyMembershipMemberKey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemberKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Group, actual.Group, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Group")); len(ds) != 0 || err != nil {
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
func compareMembershipPreferredMemberKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipPreferredMemberKey)
	if !ok {
		desiredNotPointer, ok := d.(MembershipPreferredMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipPreferredMemberKey or *MembershipPreferredMemberKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipPreferredMemberKey)
	if !ok {
		actualNotPointer, ok := a.(MembershipPreferredMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipPreferredMemberKey", a)
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

func compareMembershipRolesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRoles)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRoles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRoles or *MembershipRoles", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRoles)
	if !ok {
		actualNotPointer, ok := a.(MembershipRoles)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRoles", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExpiryDetail, actual.ExpiryDetail, dcl.DiffInfo{ObjectFunction: compareMembershipRolesExpiryDetailNewStyle, EmptyObject: EmptyMembershipRolesExpiryDetail, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpiryDetail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RestrictionEvaluations, actual.RestrictionEvaluations, dcl.DiffInfo{ObjectFunction: compareMembershipRolesRestrictionEvaluationsNewStyle, EmptyObject: EmptyMembershipRolesRestrictionEvaluations, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RestrictionEvaluations")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesExpiryDetailNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesExpiryDetail)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesExpiryDetail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesExpiryDetail or *MembershipRolesExpiryDetail", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesExpiryDetail)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesExpiryDetail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesExpiryDetail", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ExpireTime, actual.ExpireTime, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExpireTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesRestrictionEvaluationsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesRestrictionEvaluations)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesRestrictionEvaluations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluations or *MembershipRolesRestrictionEvaluations", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesRestrictionEvaluations)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesRestrictionEvaluations)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluations", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MemberRestrictionEvaluation, actual.MemberRestrictionEvaluation, dcl.DiffInfo{ObjectFunction: compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle, EmptyObject: EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemberRestrictionEvaluation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	if !ok {
		desiredNotPointer, ok := d.(MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation or *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	if !ok {
		actualNotPointer, ok := a.(MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipDisplayNameNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipDisplayName)
	if !ok {
		desiredNotPointer, ok := d.(MembershipDisplayName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipDisplayName or *MembershipDisplayName", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipDisplayName)
	if !ok {
		actualNotPointer, ok := a.(MembershipDisplayName)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipDisplayName", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GivenName, actual.GivenName, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GivenName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FamilyName, actual.FamilyName, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FamilyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FullName, actual.FullName, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FullName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipMemberKeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipMemberKey)
	if !ok {
		desiredNotPointer, ok := d.(MembershipMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipMemberKey or *MembershipMemberKey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipMemberKey)
	if !ok {
		actualNotPointer, ok := a.(MembershipMemberKey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipMemberKey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
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
func (r *Membership) urlNormalized() *Membership {
	normalized := dcl.Copy(*r).(Membership)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Group = dcl.SelfLinkToName(r.Group)
	return &normalized
}

func (r *Membership) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateMembership" {
		fields := map[string]interface{}{
			"group": dcl.ValueOrEmptyString(nr.Group),
			"name":  dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("groups/{{group}}/memberships/{{name}}:modifyMembershipRoles", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Membership resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Membership) marshal(c *Client) ([]byte, error) {
	m, err := expandMembership(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Membership: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalMembership decodes JSON responses into the Membership resource schema.
func unmarshalMembership(b []byte, c *Client, res *Membership) (*Membership, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapMembership(m, c, res)
}

func unmarshalMapMembership(m map[string]interface{}, c *Client, res *Membership) (*Membership, error) {

	flattened := flattenMembership(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandMembership expands Membership into a JSON request object.
func expandMembership(c *Client, f *Membership) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("groups/%s/memberships/%s", f.Name, dcl.SelfLinkToName(f.Group), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandMembershipPreferredMemberKey(c, f.PreferredMemberKey, res); err != nil {
		return nil, fmt.Errorf("error expanding PreferredMemberKey into preferredMemberKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["preferredMemberKey"] = v
	}
	if v, err := expandMembershipRolesSlice(c, f.Roles, res); err != nil {
		return nil, fmt.Errorf("error expanding Roles into roles: %w", err)
	} else if v != nil {
		m["roles"] = v
	}
	if v, err := expandMembershipMemberKey(c, f.MemberKey, res); err != nil {
		return nil, fmt.Errorf("error expanding MemberKey into memberKey: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["memberKey"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Group into group: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["group"] = v
	}

	return m, nil
}

// flattenMembership flattens Membership from a JSON request object into the
// Membership type.
func flattenMembership(c *Client, i interface{}, res *Membership) *Membership {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Membership{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.PreferredMemberKey = flattenMembershipPreferredMemberKey(c, m["preferredMemberKey"], res)
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.Roles = flattenMembershipRolesSlice(c, m["roles"], res)
	resultRes.Type = flattenMembershipTypeEnum(m["type"])
	resultRes.DeliverySetting = flattenMembershipDeliverySettingEnum(m["deliverySetting"])
	resultRes.DisplayName = flattenMembershipDisplayName(c, m["displayName"], res)
	resultRes.MemberKey = flattenMembershipMemberKey(c, m["memberKey"], res)
	resultRes.Group = dcl.FlattenString(m["group"])

	return resultRes
}

// expandMembershipPreferredMemberKeyMap expands the contents of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKeyMap(c *Client, f map[string]MembershipPreferredMemberKey, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipPreferredMemberKey(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipPreferredMemberKeySlice expands the contents of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKeySlice(c *Client, f []MembershipPreferredMemberKey, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipPreferredMemberKey(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipPreferredMemberKeyMap flattens the contents of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKeyMap(c *Client, i interface{}, res *Membership) map[string]MembershipPreferredMemberKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipPreferredMemberKey{}
	}

	if len(a) == 0 {
		return map[string]MembershipPreferredMemberKey{}
	}

	items := make(map[string]MembershipPreferredMemberKey)
	for k, item := range a {
		items[k] = *flattenMembershipPreferredMemberKey(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipPreferredMemberKeySlice flattens the contents of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKeySlice(c *Client, i interface{}, res *Membership) []MembershipPreferredMemberKey {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipPreferredMemberKey{}
	}

	if len(a) == 0 {
		return []MembershipPreferredMemberKey{}
	}

	items := make([]MembershipPreferredMemberKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipPreferredMemberKey(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipPreferredMemberKey expands an instance of MembershipPreferredMemberKey into a JSON
// request object.
func expandMembershipPreferredMemberKey(c *Client, f *MembershipPreferredMemberKey, res *Membership) (map[string]interface{}, error) {
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

// flattenMembershipPreferredMemberKey flattens an instance of MembershipPreferredMemberKey from a JSON
// response object.
func flattenMembershipPreferredMemberKey(c *Client, i interface{}, res *Membership) *MembershipPreferredMemberKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipPreferredMemberKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipPreferredMemberKey
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandMembershipRolesMap expands the contents of MembershipRoles into a JSON
// request object.
func expandMembershipRolesMap(c *Client, f map[string]MembershipRoles, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRoles(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesSlice expands the contents of MembershipRoles into a JSON
// request object.
func expandMembershipRolesSlice(c *Client, f []MembershipRoles, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRoles(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesMap flattens the contents of MembershipRoles from a JSON
// response object.
func flattenMembershipRolesMap(c *Client, i interface{}, res *Membership) map[string]MembershipRoles {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRoles{}
	}

	if len(a) == 0 {
		return map[string]MembershipRoles{}
	}

	items := make(map[string]MembershipRoles)
	for k, item := range a {
		items[k] = *flattenMembershipRoles(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipRolesSlice flattens the contents of MembershipRoles from a JSON
// response object.
func flattenMembershipRolesSlice(c *Client, i interface{}, res *Membership) []MembershipRoles {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRoles{}
	}

	if len(a) == 0 {
		return []MembershipRoles{}
	}

	items := make([]MembershipRoles, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRoles(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipRoles expands an instance of MembershipRoles into a JSON
// request object.
func expandMembershipRoles(c *Client, f *MembershipRoles, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := expandMembershipRolesExpiryDetail(c, f.ExpiryDetail, res); err != nil {
		return nil, fmt.Errorf("error expanding ExpiryDetail into expiryDetail: %w", err)
	} else if v != nil {
		m["expiryDetail"] = v
	}
	if v, err := expandMembershipRolesRestrictionEvaluations(c, f.RestrictionEvaluations, res); err != nil {
		return nil, fmt.Errorf("error expanding RestrictionEvaluations into restrictionEvaluations: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["restrictionEvaluations"] = v
	}

	return m, nil
}

// flattenMembershipRoles flattens an instance of MembershipRoles from a JSON
// response object.
func flattenMembershipRoles(c *Client, i interface{}, res *Membership) *MembershipRoles {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRoles{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRoles
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ExpiryDetail = flattenMembershipRolesExpiryDetail(c, m["expiryDetail"], res)
	r.RestrictionEvaluations = flattenMembershipRolesRestrictionEvaluations(c, m["restrictionEvaluations"], res)

	return r
}

// expandMembershipRolesExpiryDetailMap expands the contents of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetailMap(c *Client, f map[string]MembershipRolesExpiryDetail, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesExpiryDetail(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesExpiryDetailSlice expands the contents of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetailSlice(c *Client, f []MembershipRolesExpiryDetail, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesExpiryDetail(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesExpiryDetailMap flattens the contents of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetailMap(c *Client, i interface{}, res *Membership) map[string]MembershipRolesExpiryDetail {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesExpiryDetail{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesExpiryDetail{}
	}

	items := make(map[string]MembershipRolesExpiryDetail)
	for k, item := range a {
		items[k] = *flattenMembershipRolesExpiryDetail(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipRolesExpiryDetailSlice flattens the contents of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetailSlice(c *Client, i interface{}, res *Membership) []MembershipRolesExpiryDetail {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesExpiryDetail{}
	}

	if len(a) == 0 {
		return []MembershipRolesExpiryDetail{}
	}

	items := make([]MembershipRolesExpiryDetail, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesExpiryDetail(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipRolesExpiryDetail expands an instance of MembershipRolesExpiryDetail into a JSON
// request object.
func expandMembershipRolesExpiryDetail(c *Client, f *MembershipRolesExpiryDetail, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ExpireTime; !dcl.IsEmptyValueIndirect(v) {
		m["expireTime"] = v
	}

	return m, nil
}

// flattenMembershipRolesExpiryDetail flattens an instance of MembershipRolesExpiryDetail from a JSON
// response object.
func flattenMembershipRolesExpiryDetail(c *Client, i interface{}, res *Membership) *MembershipRolesExpiryDetail {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesExpiryDetail{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesExpiryDetail
	}
	r.ExpireTime = dcl.FlattenString(m["expireTime"])

	return r
}

// expandMembershipRolesRestrictionEvaluationsMap expands the contents of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMap(c *Client, f map[string]MembershipRolesRestrictionEvaluations, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluations(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesRestrictionEvaluationsSlice expands the contents of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsSlice(c *Client, f []MembershipRolesRestrictionEvaluations, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluations(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesRestrictionEvaluationsMap flattens the contents of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMap(c *Client, i interface{}, res *Membership) map[string]MembershipRolesRestrictionEvaluations {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluations{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluations{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluations)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluations(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsSlice flattens the contents of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsSlice(c *Client, i interface{}, res *Membership) []MembershipRolesRestrictionEvaluations {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluations{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluations{}
	}

	items := make([]MembershipRolesRestrictionEvaluations, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluations(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipRolesRestrictionEvaluations expands an instance of MembershipRolesRestrictionEvaluations into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluations(c *Client, f *MembershipRolesRestrictionEvaluations, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, f.MemberRestrictionEvaluation, res); err != nil {
		return nil, fmt.Errorf("error expanding MemberRestrictionEvaluation into memberRestrictionEvaluation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["memberRestrictionEvaluation"] = v
	}

	return m, nil
}

// flattenMembershipRolesRestrictionEvaluations flattens an instance of MembershipRolesRestrictionEvaluations from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluations(c *Client, i interface{}, res *Membership) *MembershipRolesRestrictionEvaluations {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesRestrictionEvaluations{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesRestrictionEvaluations
	}
	r.MemberRestrictionEvaluation = flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, m["memberRestrictionEvaluation"], res)

	return r
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap expands the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap(c *Client, f map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice expands the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, f []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationMap(c *Client, i interface{}, res *Membership) map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationSlice(c *Client, i interface{}, res *Membership) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation expands an instance of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation into a JSON
// request object.
func expandMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, f *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation flattens an instance of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation(c *Client, i interface{}, res *Membership) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation
	}
	r.State = flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(m["state"])

	return r
}

// expandMembershipDisplayNameMap expands the contents of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayNameMap(c *Client, f map[string]MembershipDisplayName, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipDisplayName(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipDisplayNameSlice expands the contents of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayNameSlice(c *Client, f []MembershipDisplayName, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipDisplayName(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipDisplayNameMap flattens the contents of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayNameMap(c *Client, i interface{}, res *Membership) map[string]MembershipDisplayName {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipDisplayName{}
	}

	if len(a) == 0 {
		return map[string]MembershipDisplayName{}
	}

	items := make(map[string]MembershipDisplayName)
	for k, item := range a {
		items[k] = *flattenMembershipDisplayName(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipDisplayNameSlice flattens the contents of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayNameSlice(c *Client, i interface{}, res *Membership) []MembershipDisplayName {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipDisplayName{}
	}

	if len(a) == 0 {
		return []MembershipDisplayName{}
	}

	items := make([]MembershipDisplayName, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipDisplayName(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipDisplayName expands an instance of MembershipDisplayName into a JSON
// request object.
func expandMembershipDisplayName(c *Client, f *MembershipDisplayName, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipDisplayName flattens an instance of MembershipDisplayName from a JSON
// response object.
func flattenMembershipDisplayName(c *Client, i interface{}, res *Membership) *MembershipDisplayName {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipDisplayName{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipDisplayName
	}
	r.GivenName = dcl.FlattenString(m["givenName"])
	r.FamilyName = dcl.FlattenString(m["familyName"])
	r.FullName = dcl.FlattenString(m["fullName"])

	return r
}

// expandMembershipMemberKeyMap expands the contents of MembershipMemberKey into a JSON
// request object.
func expandMembershipMemberKeyMap(c *Client, f map[string]MembershipMemberKey, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipMemberKey(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipMemberKeySlice expands the contents of MembershipMemberKey into a JSON
// request object.
func expandMembershipMemberKeySlice(c *Client, f []MembershipMemberKey, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipMemberKey(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipMemberKeyMap flattens the contents of MembershipMemberKey from a JSON
// response object.
func flattenMembershipMemberKeyMap(c *Client, i interface{}, res *Membership) map[string]MembershipMemberKey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipMemberKey{}
	}

	if len(a) == 0 {
		return map[string]MembershipMemberKey{}
	}

	items := make(map[string]MembershipMemberKey)
	for k, item := range a {
		items[k] = *flattenMembershipMemberKey(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipMemberKeySlice flattens the contents of MembershipMemberKey from a JSON
// response object.
func flattenMembershipMemberKeySlice(c *Client, i interface{}, res *Membership) []MembershipMemberKey {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipMemberKey{}
	}

	if len(a) == 0 {
		return []MembershipMemberKey{}
	}

	items := make([]MembershipMemberKey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipMemberKey(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipMemberKey expands an instance of MembershipMemberKey into a JSON
// request object.
func expandMembershipMemberKey(c *Client, f *MembershipMemberKey, res *Membership) (map[string]interface{}, error) {
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

// flattenMembershipMemberKey flattens an instance of MembershipMemberKey from a JSON
// response object.
func flattenMembershipMemberKey(c *Client, i interface{}, res *Membership) *MembershipMemberKey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipMemberKey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipMemberKey
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumMap flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumMap(c *Client, i interface{}, res *Membership) map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	items := make(map[string]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum)
	for k, item := range a {
		items[k] = *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumSlice flattens the contents of MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum from a JSON
// response object.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumSlice(c *Client, i interface{}, res *Membership) []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	if len(a) == 0 {
		return []MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum{}
	}

	items := make([]MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum with the same value as that string.
func flattenMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum(i interface{}) *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationStateEnumRef(s)
}

// flattenMembershipTypeEnumMap flattens the contents of MembershipTypeEnum from a JSON
// response object.
func flattenMembershipTypeEnumMap(c *Client, i interface{}, res *Membership) map[string]MembershipTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipTypeEnum{}
	}

	items := make(map[string]MembershipTypeEnum)
	for k, item := range a {
		items[k] = *flattenMembershipTypeEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipTypeEnumSlice flattens the contents of MembershipTypeEnum from a JSON
// response object.
func flattenMembershipTypeEnumSlice(c *Client, i interface{}, res *Membership) []MembershipTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipTypeEnum{}
	}

	if len(a) == 0 {
		return []MembershipTypeEnum{}
	}

	items := make([]MembershipTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipTypeEnum with the same value as that string.
func flattenMembershipTypeEnum(i interface{}) *MembershipTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MembershipTypeEnumRef(s)
}

// flattenMembershipDeliverySettingEnumMap flattens the contents of MembershipDeliverySettingEnum from a JSON
// response object.
func flattenMembershipDeliverySettingEnumMap(c *Client, i interface{}, res *Membership) map[string]MembershipDeliverySettingEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipDeliverySettingEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipDeliverySettingEnum{}
	}

	items := make(map[string]MembershipDeliverySettingEnum)
	for k, item := range a {
		items[k] = *flattenMembershipDeliverySettingEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipDeliverySettingEnumSlice flattens the contents of MembershipDeliverySettingEnum from a JSON
// response object.
func flattenMembershipDeliverySettingEnumSlice(c *Client, i interface{}, res *Membership) []MembershipDeliverySettingEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipDeliverySettingEnum{}
	}

	if len(a) == 0 {
		return []MembershipDeliverySettingEnum{}
	}

	items := make([]MembershipDeliverySettingEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipDeliverySettingEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipDeliverySettingEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipDeliverySettingEnum with the same value as that string.
func flattenMembershipDeliverySettingEnum(i interface{}) *MembershipDeliverySettingEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MembershipDeliverySettingEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Membership) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalMembership(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Group == nil && ncr.Group == nil {
			c.Config.Logger.Info("Both Group fields null - considering equal.")
		} else if nr.Group == nil || ncr.Group == nil {
			c.Config.Logger.Info("Only one Group field is null - considering unequal.")
			return false
		} else if *nr.Group != *ncr.Group {
			return false
		}
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

type membershipDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         membershipApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToMembershipDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]membershipDiff, error) {
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
	var diffs []membershipDiff
	// For each operation name, create a membershipDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := membershipDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToMembershipApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToMembershipApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (membershipApiOperation, error) {
	switch opName {

	case "updateMembershipUpdateMembershipOperation":
		return &updateMembershipUpdateMembershipOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractMembershipFields(r *Membership) error {
	vPreferredMemberKey := r.PreferredMemberKey
	if vPreferredMemberKey == nil {
		// note: explicitly not the empty object.
		vPreferredMemberKey = &MembershipPreferredMemberKey{}
	}
	if err := extractMembershipPreferredMemberKeyFields(r, vPreferredMemberKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPreferredMemberKey) {
		r.PreferredMemberKey = vPreferredMemberKey
	}
	vDisplayName := r.DisplayName
	if vDisplayName == nil {
		// note: explicitly not the empty object.
		vDisplayName = &MembershipDisplayName{}
	}
	if err := extractMembershipDisplayNameFields(r, vDisplayName); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDisplayName) {
		r.DisplayName = vDisplayName
	}
	vMemberKey := r.MemberKey
	if vMemberKey == nil {
		// note: explicitly not the empty object.
		vMemberKey = &MembershipMemberKey{}
	}
	if err := extractMembershipMemberKeyFields(r, vMemberKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMemberKey) {
		r.MemberKey = vMemberKey
	}
	return nil
}
func extractMembershipPreferredMemberKeyFields(r *Membership, o *MembershipPreferredMemberKey) error {
	return nil
}
func extractMembershipRolesFields(r *Membership, o *MembershipRoles) error {
	vExpiryDetail := o.ExpiryDetail
	if vExpiryDetail == nil {
		// note: explicitly not the empty object.
		vExpiryDetail = &MembershipRolesExpiryDetail{}
	}
	if err := extractMembershipRolesExpiryDetailFields(r, vExpiryDetail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExpiryDetail) {
		o.ExpiryDetail = vExpiryDetail
	}
	vRestrictionEvaluations := o.RestrictionEvaluations
	if vRestrictionEvaluations == nil {
		// note: explicitly not the empty object.
		vRestrictionEvaluations = &MembershipRolesRestrictionEvaluations{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsFields(r, vRestrictionEvaluations); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRestrictionEvaluations) {
		o.RestrictionEvaluations = vRestrictionEvaluations
	}
	return nil
}
func extractMembershipRolesExpiryDetailFields(r *Membership, o *MembershipRolesExpiryDetail) error {
	return nil
}
func extractMembershipRolesRestrictionEvaluationsFields(r *Membership, o *MembershipRolesRestrictionEvaluations) error {
	vMemberRestrictionEvaluation := o.MemberRestrictionEvaluation
	if vMemberRestrictionEvaluation == nil {
		// note: explicitly not the empty object.
		vMemberRestrictionEvaluation = &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r, vMemberRestrictionEvaluation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMemberRestrictionEvaluation) {
		o.MemberRestrictionEvaluation = vMemberRestrictionEvaluation
	}
	return nil
}
func extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r *Membership, o *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) error {
	return nil
}
func extractMembershipDisplayNameFields(r *Membership, o *MembershipDisplayName) error {
	return nil
}
func extractMembershipMemberKeyFields(r *Membership, o *MembershipMemberKey) error {
	return nil
}

func postReadExtractMembershipFields(r *Membership) error {
	vPreferredMemberKey := r.PreferredMemberKey
	if vPreferredMemberKey == nil {
		// note: explicitly not the empty object.
		vPreferredMemberKey = &MembershipPreferredMemberKey{}
	}
	if err := postReadExtractMembershipPreferredMemberKeyFields(r, vPreferredMemberKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPreferredMemberKey) {
		r.PreferredMemberKey = vPreferredMemberKey
	}
	vDisplayName := r.DisplayName
	if vDisplayName == nil {
		// note: explicitly not the empty object.
		vDisplayName = &MembershipDisplayName{}
	}
	if err := postReadExtractMembershipDisplayNameFields(r, vDisplayName); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDisplayName) {
		r.DisplayName = vDisplayName
	}
	vMemberKey := r.MemberKey
	if vMemberKey == nil {
		// note: explicitly not the empty object.
		vMemberKey = &MembershipMemberKey{}
	}
	if err := postReadExtractMembershipMemberKeyFields(r, vMemberKey); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMemberKey) {
		r.MemberKey = vMemberKey
	}
	return nil
}
func postReadExtractMembershipPreferredMemberKeyFields(r *Membership, o *MembershipPreferredMemberKey) error {
	return nil
}
func postReadExtractMembershipRolesFields(r *Membership, o *MembershipRoles) error {
	vExpiryDetail := o.ExpiryDetail
	if vExpiryDetail == nil {
		// note: explicitly not the empty object.
		vExpiryDetail = &MembershipRolesExpiryDetail{}
	}
	if err := extractMembershipRolesExpiryDetailFields(r, vExpiryDetail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExpiryDetail) {
		o.ExpiryDetail = vExpiryDetail
	}
	vRestrictionEvaluations := o.RestrictionEvaluations
	if vRestrictionEvaluations == nil {
		// note: explicitly not the empty object.
		vRestrictionEvaluations = &MembershipRolesRestrictionEvaluations{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsFields(r, vRestrictionEvaluations); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRestrictionEvaluations) {
		o.RestrictionEvaluations = vRestrictionEvaluations
	}
	return nil
}
func postReadExtractMembershipRolesExpiryDetailFields(r *Membership, o *MembershipRolesExpiryDetail) error {
	return nil
}
func postReadExtractMembershipRolesRestrictionEvaluationsFields(r *Membership, o *MembershipRolesRestrictionEvaluations) error {
	vMemberRestrictionEvaluation := o.MemberRestrictionEvaluation
	if vMemberRestrictionEvaluation == nil {
		// note: explicitly not the empty object.
		vMemberRestrictionEvaluation = &MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation{}
	}
	if err := extractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r, vMemberRestrictionEvaluation); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMemberRestrictionEvaluation) {
		o.MemberRestrictionEvaluation = vMemberRestrictionEvaluation
	}
	return nil
}
func postReadExtractMembershipRolesRestrictionEvaluationsMemberRestrictionEvaluationFields(r *Membership, o *MembershipRolesRestrictionEvaluationsMemberRestrictionEvaluation) error {
	return nil
}
func postReadExtractMembershipDisplayNameFields(r *Membership, o *MembershipDisplayName) error {
	return nil
}
func postReadExtractMembershipMemberKeyFields(r *Membership, o *MembershipMemberKey) error {
	return nil
}
