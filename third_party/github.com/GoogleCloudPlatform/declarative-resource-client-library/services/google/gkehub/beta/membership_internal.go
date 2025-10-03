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

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Endpoint) {
		if err := r.Endpoint.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.State) {
		if err := r.State.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Authority) {
		if err := r.Authority.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpoint) validate() error {
	if !dcl.IsEmptyValueIndirect(r.GkeCluster) {
		if err := r.GkeCluster.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubernetesMetadata) {
		if err := r.KubernetesMetadata.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.KubernetesResource) {
		if err := r.KubernetesResource.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpointGkeCluster) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesMetadata) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResource) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ResourceOptions) {
		if err := r.ResourceOptions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *MembershipEndpointKubernetesResourceMembershipResources) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResourceConnectResources) validate() error {
	return nil
}
func (r *MembershipEndpointKubernetesResourceResourceOptions) validate() error {
	return nil
}
func (r *MembershipState) validate() error {
	return nil
}
func (r *MembershipAuthority) validate() error {
	return nil
}
func (r *Membership) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://gkehub.googleapis.com/v1beta1/", params)
}

func (r *Membership) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Membership) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships?membershipId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Membership) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Membership) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Membership) SetPolicyVerb() string {
	return ""
}

func (r *Membership) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{}
	return dcl.URL("", nr.basePath(), userBasePath, fields)
}

func (r *Membership) IAMPolicyVersion() int {
	return 3
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

	if v, err := expandMembershipEndpoint(c, f.Endpoint, res); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["endpoint"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.ExternalId; !dcl.IsEmptyValueIndirect(v) {
		req["externalId"] = v
	}
	if v, err := expandMembershipAuthority(c, f.Authority, res); err != nil {
		return nil, fmt.Errorf("error expanding Authority into authority: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["authority"] = v
	}
	if v := f.InfrastructureType; !dcl.IsEmptyValueIndirect(v) {
		req["infrastructureType"] = v
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

func (op *updateMembershipUpdateMembershipOperation) do(ctx context.Context, r *Membership, c *Client) error {
	_, err := c.GetMembership(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateMembership")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateMembershipUpdateMembershipRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateMembershipUpdateMembershipRequest(c, req)
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
	Resources []map[string]interface{} `json:"resources"`
	Token     string                   `json:"nextPageToken"`
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
	for _, v := range m.Resources {
		res, err := unmarshalMapMembership(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
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
		rawDesired.Endpoint = canonicalizeMembershipEndpoint(rawDesired.Endpoint, nil, opts...)
		rawDesired.State = canonicalizeMembershipState(rawDesired.State, nil, opts...)
		rawDesired.Authority = canonicalizeMembershipAuthority(rawDesired.Authority, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Membership{}
	canonicalDesired.Endpoint = canonicalizeMembershipEndpoint(rawDesired.Endpoint, rawInitial.Endpoint, opts...)
	if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.ExternalId, rawInitial.ExternalId) {
		canonicalDesired.ExternalId = rawInitial.ExternalId
	} else {
		canonicalDesired.ExternalId = rawDesired.ExternalId
	}
	canonicalDesired.Authority = canonicalizeMembershipAuthority(rawDesired.Authority, rawInitial.Authority, opts...)
	if dcl.IsZeroValue(rawDesired.InfrastructureType) || (dcl.IsEmptyValueIndirect(rawDesired.InfrastructureType) && dcl.IsEmptyValueIndirect(rawInitial.InfrastructureType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.InfrastructureType = rawInitial.InfrastructureType
	} else {
		canonicalDesired.InfrastructureType = rawDesired.InfrastructureType
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeMembershipNewState(c *Client, rawNew, rawDesired *Membership) (*Membership, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Endpoint) && dcl.IsEmptyValueIndirect(rawDesired.Endpoint) {
		rawNew.Endpoint = rawDesired.Endpoint
	} else {
		rawNew.Endpoint = canonicalizeNewMembershipEndpoint(c, rawDesired.Endpoint, rawNew.Endpoint)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
		rawNew.State = canonicalizeNewMembershipState(c, rawDesired.State, rawNew.State)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DeleteTime) && dcl.IsEmptyValueIndirect(rawDesired.DeleteTime) {
		rawNew.DeleteTime = rawDesired.DeleteTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ExternalId) && dcl.IsEmptyValueIndirect(rawDesired.ExternalId) {
		rawNew.ExternalId = rawDesired.ExternalId
	} else {
		if dcl.StringCanonicalize(rawDesired.ExternalId, rawNew.ExternalId) {
			rawNew.ExternalId = rawDesired.ExternalId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastConnectionTime) && dcl.IsEmptyValueIndirect(rawDesired.LastConnectionTime) {
		rawNew.LastConnectionTime = rawDesired.LastConnectionTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.UniqueId) && dcl.IsEmptyValueIndirect(rawDesired.UniqueId) {
		rawNew.UniqueId = rawDesired.UniqueId
	} else {
		if dcl.StringCanonicalize(rawDesired.UniqueId, rawNew.UniqueId) {
			rawNew.UniqueId = rawDesired.UniqueId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Authority) && dcl.IsEmptyValueIndirect(rawDesired.Authority) {
		rawNew.Authority = rawDesired.Authority
	} else {
		rawNew.Authority = canonicalizeNewMembershipAuthority(c, rawDesired.Authority, rawNew.Authority)
	}

	if dcl.IsEmptyValueIndirect(rawNew.InfrastructureType) && dcl.IsEmptyValueIndirect(rawDesired.InfrastructureType) {
		rawNew.InfrastructureType = rawDesired.InfrastructureType
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeMembershipEndpoint(des, initial *MembershipEndpoint, opts ...dcl.ApplyOption) *MembershipEndpoint {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpoint{}

	cDes.GkeCluster = canonicalizeMembershipEndpointGkeCluster(des.GkeCluster, initial.GkeCluster, opts...)
	cDes.KubernetesResource = canonicalizeMembershipEndpointKubernetesResource(des.KubernetesResource, initial.KubernetesResource, opts...)

	return cDes
}

func canonicalizeMembershipEndpointSlice(des, initial []MembershipEndpoint, opts ...dcl.ApplyOption) []MembershipEndpoint {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpoint, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpoint(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpoint, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpoint(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpoint(c *Client, des, nw *MembershipEndpoint) *MembershipEndpoint {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpoint while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.GkeCluster = canonicalizeNewMembershipEndpointGkeCluster(c, des.GkeCluster, nw.GkeCluster)
	nw.KubernetesMetadata = canonicalizeNewMembershipEndpointKubernetesMetadata(c, des.KubernetesMetadata, nw.KubernetesMetadata)
	nw.KubernetesResource = canonicalizeNewMembershipEndpointKubernetesResource(c, des.KubernetesResource, nw.KubernetesResource)

	return nw
}

func canonicalizeNewMembershipEndpointSet(c *Client, des, nw []MembershipEndpoint) []MembershipEndpoint {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpoint
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpoint(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointSlice(c *Client, des, nw []MembershipEndpoint) []MembershipEndpoint {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpoint
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpoint(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointGkeCluster(des, initial *MembershipEndpointGkeCluster, opts ...dcl.ApplyOption) *MembershipEndpointGkeCluster {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointGkeCluster{}

	if dcl.IsZeroValue(des.ResourceLink) || (dcl.IsEmptyValueIndirect(des.ResourceLink) && dcl.IsEmptyValueIndirect(initial.ResourceLink)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResourceLink = initial.ResourceLink
	} else {
		cDes.ResourceLink = des.ResourceLink
	}

	return cDes
}

func canonicalizeMembershipEndpointGkeClusterSlice(des, initial []MembershipEndpointGkeCluster, opts ...dcl.ApplyOption) []MembershipEndpointGkeCluster {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointGkeCluster, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointGkeCluster(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointGkeCluster, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointGkeCluster(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointGkeCluster(c *Client, des, nw *MembershipEndpointGkeCluster) *MembershipEndpointGkeCluster {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointGkeCluster while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipEndpointGkeClusterSet(c *Client, des, nw []MembershipEndpointGkeCluster) []MembershipEndpointGkeCluster {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointGkeCluster
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointGkeClusterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointGkeCluster(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointGkeClusterSlice(c *Client, des, nw []MembershipEndpointGkeCluster) []MembershipEndpointGkeCluster {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointGkeCluster
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointGkeCluster(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesMetadata(des, initial *MembershipEndpointKubernetesMetadata, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointKubernetesMetadata{}

	return cDes
}

func canonicalizeMembershipEndpointKubernetesMetadataSlice(des, initial []MembershipEndpointKubernetesMetadata, opts ...dcl.ApplyOption) []MembershipEndpointKubernetesMetadata {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointKubernetesMetadata, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointKubernetesMetadata(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointKubernetesMetadata, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointKubernetesMetadata(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointKubernetesMetadata(c *Client, des, nw *MembershipEndpointKubernetesMetadata) *MembershipEndpointKubernetesMetadata {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointKubernetesMetadata while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.KubernetesApiServerVersion, nw.KubernetesApiServerVersion) {
		nw.KubernetesApiServerVersion = des.KubernetesApiServerVersion
	}
	if dcl.StringCanonicalize(des.NodeProviderId, nw.NodeProviderId) {
		nw.NodeProviderId = des.NodeProviderId
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesMetadataSet(c *Client, des, nw []MembershipEndpointKubernetesMetadata) []MembershipEndpointKubernetesMetadata {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointKubernetesMetadata
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointKubernetesMetadata(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointKubernetesMetadataSlice(c *Client, des, nw []MembershipEndpointKubernetesMetadata) []MembershipEndpointKubernetesMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResource(des, initial *MembershipEndpointKubernetesResource, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResource {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointKubernetesResource{}

	if dcl.StringCanonicalize(des.MembershipCrManifest, initial.MembershipCrManifest) || dcl.IsZeroValue(des.MembershipCrManifest) {
		cDes.MembershipCrManifest = initial.MembershipCrManifest
	} else {
		cDes.MembershipCrManifest = des.MembershipCrManifest
	}
	cDes.ResourceOptions = canonicalizeMembershipEndpointKubernetesResourceResourceOptions(des.ResourceOptions, initial.ResourceOptions, opts...)

	return cDes
}

func canonicalizeMembershipEndpointKubernetesResourceSlice(des, initial []MembershipEndpointKubernetesResource, opts ...dcl.ApplyOption) []MembershipEndpointKubernetesResource {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointKubernetesResource, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointKubernetesResource(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointKubernetesResource, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointKubernetesResource(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointKubernetesResource(c *Client, des, nw *MembershipEndpointKubernetesResource) *MembershipEndpointKubernetesResource {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointKubernetesResource while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MembershipCrManifest = des.MembershipCrManifest
	nw.MembershipResources = canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSlice(c, des.MembershipResources, nw.MembershipResources)
	nw.ConnectResources = canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSlice(c, des.ConnectResources, nw.ConnectResources)
	nw.ResourceOptions = canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c, des.ResourceOptions, nw.ResourceOptions)

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceSet(c *Client, des, nw []MembershipEndpointKubernetesResource) []MembershipEndpointKubernetesResource {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointKubernetesResource
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointKubernetesResource(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointKubernetesResourceSlice(c *Client, des, nw []MembershipEndpointKubernetesResource) []MembershipEndpointKubernetesResource {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResource
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResource(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceMembershipResources(des, initial *MembershipEndpointKubernetesResourceMembershipResources, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointKubernetesResourceMembershipResources{}

	if dcl.StringCanonicalize(des.Manifest, initial.Manifest) || dcl.IsZeroValue(des.Manifest) {
		cDes.Manifest = initial.Manifest
	} else {
		cDes.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, initial.ClusterScoped) || dcl.IsZeroValue(des.ClusterScoped) {
		cDes.ClusterScoped = initial.ClusterScoped
	} else {
		cDes.ClusterScoped = des.ClusterScoped
	}

	return cDes
}

func canonicalizeMembershipEndpointKubernetesResourceMembershipResourcesSlice(des, initial []MembershipEndpointKubernetesResourceMembershipResources, opts ...dcl.ApplyOption) []MembershipEndpointKubernetesResourceMembershipResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointKubernetesResourceMembershipResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointKubernetesResourceMembershipResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointKubernetesResourceMembershipResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointKubernetesResourceMembershipResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResources(c *Client, des, nw *MembershipEndpointKubernetesResourceMembershipResources) *MembershipEndpointKubernetesResourceMembershipResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointKubernetesResourceMembershipResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Manifest, nw.Manifest) {
		nw.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, nw.ClusterScoped) {
		nw.ClusterScoped = des.ClusterScoped
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSet(c *Client, des, nw []MembershipEndpointKubernetesResourceMembershipResources) []MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointKubernetesResourceMembershipResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceMembershipResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceMembershipResources) []MembershipEndpointKubernetesResourceMembershipResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceMembershipResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceMembershipResources(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceConnectResources(des, initial *MembershipEndpointKubernetesResourceConnectResources, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointKubernetesResourceConnectResources{}

	if dcl.StringCanonicalize(des.Manifest, initial.Manifest) || dcl.IsZeroValue(des.Manifest) {
		cDes.Manifest = initial.Manifest
	} else {
		cDes.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, initial.ClusterScoped) || dcl.IsZeroValue(des.ClusterScoped) {
		cDes.ClusterScoped = initial.ClusterScoped
	} else {
		cDes.ClusterScoped = des.ClusterScoped
	}

	return cDes
}

func canonicalizeMembershipEndpointKubernetesResourceConnectResourcesSlice(des, initial []MembershipEndpointKubernetesResourceConnectResources, opts ...dcl.ApplyOption) []MembershipEndpointKubernetesResourceConnectResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointKubernetesResourceConnectResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointKubernetesResourceConnectResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointKubernetesResourceConnectResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointKubernetesResourceConnectResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResources(c *Client, des, nw *MembershipEndpointKubernetesResourceConnectResources) *MembershipEndpointKubernetesResourceConnectResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointKubernetesResourceConnectResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Manifest, nw.Manifest) {
		nw.Manifest = des.Manifest
	}
	if dcl.BoolCanonicalize(des.ClusterScoped, nw.ClusterScoped) {
		nw.ClusterScoped = des.ClusterScoped
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSet(c *Client, des, nw []MembershipEndpointKubernetesResourceConnectResources) []MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointKubernetesResourceConnectResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceConnectResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceConnectResources) []MembershipEndpointKubernetesResourceConnectResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceConnectResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceConnectResources(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipEndpointKubernetesResourceResourceOptions(des, initial *MembershipEndpointKubernetesResourceResourceOptions, opts ...dcl.ApplyOption) *MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipEndpointKubernetesResourceResourceOptions{}

	if dcl.StringCanonicalize(des.ConnectVersion, initial.ConnectVersion) || dcl.IsZeroValue(des.ConnectVersion) {
		cDes.ConnectVersion = initial.ConnectVersion
	} else {
		cDes.ConnectVersion = des.ConnectVersion
	}
	if dcl.BoolCanonicalize(des.V1Beta1Crd, initial.V1Beta1Crd) || dcl.IsZeroValue(des.V1Beta1Crd) {
		cDes.V1Beta1Crd = initial.V1Beta1Crd
	} else {
		cDes.V1Beta1Crd = des.V1Beta1Crd
	}

	return cDes
}

func canonicalizeMembershipEndpointKubernetesResourceResourceOptionsSlice(des, initial []MembershipEndpointKubernetesResourceResourceOptions, opts ...dcl.ApplyOption) []MembershipEndpointKubernetesResourceResourceOptions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipEndpointKubernetesResourceResourceOptions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipEndpointKubernetesResourceResourceOptions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipEndpointKubernetesResourceResourceOptions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipEndpointKubernetesResourceResourceOptions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c *Client, des, nw *MembershipEndpointKubernetesResourceResourceOptions) *MembershipEndpointKubernetesResourceResourceOptions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipEndpointKubernetesResourceResourceOptions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ConnectVersion, nw.ConnectVersion) {
		nw.ConnectVersion = des.ConnectVersion
	}
	if dcl.BoolCanonicalize(des.V1Beta1Crd, nw.V1Beta1Crd) {
		nw.V1Beta1Crd = des.V1Beta1Crd
	}

	return nw
}

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptionsSet(c *Client, des, nw []MembershipEndpointKubernetesResourceResourceOptions) []MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipEndpointKubernetesResourceResourceOptions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, des, nw []MembershipEndpointKubernetesResourceResourceOptions) []MembershipEndpointKubernetesResourceResourceOptions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipEndpointKubernetesResourceResourceOptions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipEndpointKubernetesResourceResourceOptions(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipState(des, initial *MembershipState, opts ...dcl.ApplyOption) *MembershipState {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipState{}

	return cDes
}

func canonicalizeMembershipStateSlice(des, initial []MembershipState, opts ...dcl.ApplyOption) []MembershipState {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipState, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipState(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipState, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipState(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipState(c *Client, des, nw *MembershipState) *MembershipState {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipState while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewMembershipStateSet(c *Client, des, nw []MembershipState) []MembershipState {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipState
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipStateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipState(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipStateSlice(c *Client, des, nw []MembershipState) []MembershipState {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipState
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipState(c, &d, &n))
	}

	return items
}

func canonicalizeMembershipAuthority(des, initial *MembershipAuthority, opts ...dcl.ApplyOption) *MembershipAuthority {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &MembershipAuthority{}

	if dcl.StringCanonicalize(des.Issuer, initial.Issuer) || dcl.IsZeroValue(des.Issuer) {
		cDes.Issuer = initial.Issuer
	} else {
		cDes.Issuer = des.Issuer
	}

	return cDes
}

func canonicalizeMembershipAuthoritySlice(des, initial []MembershipAuthority, opts ...dcl.ApplyOption) []MembershipAuthority {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]MembershipAuthority, 0, len(des))
		for _, d := range des {
			cd := canonicalizeMembershipAuthority(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]MembershipAuthority, 0, len(des))
	for i, d := range des {
		cd := canonicalizeMembershipAuthority(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewMembershipAuthority(c *Client, des, nw *MembershipAuthority) *MembershipAuthority {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for MembershipAuthority while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Issuer, nw.Issuer) {
		nw.Issuer = des.Issuer
	}
	if dcl.StringCanonicalize(des.WorkloadIdentityPool, nw.WorkloadIdentityPool) {
		nw.WorkloadIdentityPool = des.WorkloadIdentityPool
	}
	if dcl.StringCanonicalize(des.IdentityProvider, nw.IdentityProvider) {
		nw.IdentityProvider = des.IdentityProvider
	}

	return nw
}

func canonicalizeNewMembershipAuthoritySet(c *Client, des, nw []MembershipAuthority) []MembershipAuthority {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []MembershipAuthority
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareMembershipAuthorityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewMembershipAuthority(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewMembershipAuthoritySlice(c *Client, des, nw []MembershipAuthority) []MembershipAuthority {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []MembershipAuthority
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewMembershipAuthority(c, &d, &n))
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
	if ds, err := dcl.Diff(desired.Endpoint, actual.Endpoint, dcl.DiffInfo{ObjectFunction: compareMembershipEndpointNewStyle, EmptyObject: EmptyMembershipEndpoint, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Endpoint")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMembershipStateNewStyle, EmptyObject: EmptyMembershipState, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.DeleteTime, actual.DeleteTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DeleteTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExternalId, actual.ExternalId, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("ExternalId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastConnectionTime, actual.LastConnectionTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastConnectionTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UniqueId, actual.UniqueId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UniqueId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Authority, actual.Authority, dcl.DiffInfo{ObjectFunction: compareMembershipAuthorityNewStyle, EmptyObject: EmptyMembershipAuthority, OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Authority")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InfrastructureType, actual.InfrastructureType, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("InfrastructureType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Location, actual.Location, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Location")); len(ds) != 0 || err != nil {
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
func compareMembershipEndpointNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpoint)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpoint or *MembershipEndpoint", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpoint)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpoint)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpoint", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.GkeCluster, actual.GkeCluster, dcl.DiffInfo{ObjectFunction: compareMembershipEndpointGkeClusterNewStyle, EmptyObject: EmptyMembershipEndpointGkeCluster, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GkeCluster")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesMetadata, actual.KubernetesMetadata, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesMetadataNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesMetadata, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesMetadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesResource, actual.KubernetesResource, dcl.DiffInfo{ObjectFunction: compareMembershipEndpointKubernetesResourceNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResource, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesResource")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointGkeClusterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointGkeCluster)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointGkeCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointGkeCluster or *MembershipEndpointGkeCluster", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointGkeCluster)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointGkeCluster)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointGkeCluster", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceLink, actual.ResourceLink, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceLink")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesMetadata)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesMetadata or *MembershipEndpointKubernetesMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesMetadata)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KubernetesApiServerVersion, actual.KubernetesApiServerVersion, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KubernetesApiServerVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeProviderId, actual.NodeProviderId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeProviderId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NodeCount, actual.NodeCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NodeCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VcpuCount, actual.VcpuCount, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VcpuCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemoryMb, actual.MemoryMb, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemoryMb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResource)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResource or *MembershipEndpointKubernetesResource", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResource)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResource)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResource", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MembershipCrManifest, actual.MembershipCrManifest, dcl.DiffInfo{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MembershipCrManifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MembershipResources, actual.MembershipResources, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceMembershipResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MembershipResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectResources, actual.ConnectResources, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceConnectResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectResources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceOptions, actual.ResourceOptions, dcl.DiffInfo{ObjectFunction: compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle, EmptyObject: EmptyMembershipEndpointKubernetesResourceResourceOptions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceMembershipResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceMembershipResources)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceMembershipResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceMembershipResources or *MembershipEndpointKubernetesResourceMembershipResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceMembershipResources)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceMembershipResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceMembershipResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Manifest, actual.Manifest, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Manifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterScoped, actual.ClusterScoped, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterScoped")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceConnectResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceConnectResources)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceConnectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceConnectResources or *MembershipEndpointKubernetesResourceConnectResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceConnectResources)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceConnectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceConnectResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Manifest, actual.Manifest, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Manifest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterScoped, actual.ClusterScoped, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ClusterScoped")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipEndpointKubernetesResourceResourceOptionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipEndpointKubernetesResourceResourceOptions)
	if !ok {
		desiredNotPointer, ok := d.(MembershipEndpointKubernetesResourceResourceOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceResourceOptions or *MembershipEndpointKubernetesResourceResourceOptions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipEndpointKubernetesResourceResourceOptions)
	if !ok {
		actualNotPointer, ok := a.(MembershipEndpointKubernetesResourceResourceOptions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipEndpointKubernetesResourceResourceOptions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConnectVersion, actual.ConnectVersion, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.V1Beta1Crd, actual.V1Beta1Crd, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("V1beta1Crd")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipStateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipState)
	if !ok {
		desiredNotPointer, ok := d.(MembershipState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipState or *MembershipState", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipState)
	if !ok {
		actualNotPointer, ok := a.(MembershipState)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipState", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Code, actual.Code, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Code")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareMembershipAuthorityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*MembershipAuthority)
	if !ok {
		desiredNotPointer, ok := d.(MembershipAuthority)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipAuthority or *MembershipAuthority", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*MembershipAuthority)
	if !ok {
		actualNotPointer, ok := a.(MembershipAuthority)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a MembershipAuthority", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Issuer, actual.Issuer, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateMembershipUpdateMembershipOperation")}, fn.AddNest("Issuer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WorkloadIdentityPool, actual.WorkloadIdentityPool, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WorkloadIdentityPool")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IdentityProvider, actual.IdentityProvider, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IdentityProvider")); len(ds) != 0 || err != nil {
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
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.ExternalId = dcl.SelfLinkToName(r.ExternalId)
	normalized.UniqueId = dcl.SelfLinkToName(r.UniqueId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Membership) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateMembership" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/memberships/{{name}}", nr.basePath(), userBasePath, fields), nil

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
	if v, err := expandMembershipEndpoint(c, f.Endpoint, res); err != nil {
		return nil, fmt.Errorf("error expanding Endpoint into endpoint: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endpoint"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/locations/%s/memberships/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.ExternalId; dcl.ValueShouldBeSent(v) {
		m["externalId"] = v
	}
	if v, err := expandMembershipAuthority(c, f.Authority, res); err != nil {
		return nil, fmt.Errorf("error expanding Authority into authority: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["authority"] = v
	}
	if v := f.InfrastructureType; dcl.ValueShouldBeSent(v) {
		m["infrastructureType"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
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
	resultRes.Endpoint = flattenMembershipEndpoint(c, m["endpoint"], res)
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenMembershipState(c, m["state"], res)
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.DeleteTime = dcl.FlattenString(m["deleteTime"])
	resultRes.ExternalId = dcl.FlattenString(m["externalId"])
	resultRes.LastConnectionTime = dcl.FlattenString(m["lastConnectionTime"])
	resultRes.UniqueId = dcl.FlattenString(m["uniqueId"])
	resultRes.Authority = flattenMembershipAuthority(c, m["authority"], res)
	resultRes.InfrastructureType = flattenMembershipInfrastructureTypeEnum(m["infrastructureType"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandMembershipEndpointMap expands the contents of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpointMap(c *Client, f map[string]MembershipEndpoint, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpoint(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointSlice expands the contents of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpointSlice(c *Client, f []MembershipEndpoint, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpoint(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointMap flattens the contents of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpointMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpoint {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpoint{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpoint{}
	}

	items := make(map[string]MembershipEndpoint)
	for k, item := range a {
		items[k] = *flattenMembershipEndpoint(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointSlice flattens the contents of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpointSlice(c *Client, i interface{}, res *Membership) []MembershipEndpoint {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpoint{}
	}

	if len(a) == 0 {
		return []MembershipEndpoint{}
	}

	items := make([]MembershipEndpoint, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpoint(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpoint expands an instance of MembershipEndpoint into a JSON
// request object.
func expandMembershipEndpoint(c *Client, f *MembershipEndpoint, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandMembershipEndpointGkeCluster(c, f.GkeCluster, res); err != nil {
		return nil, fmt.Errorf("error expanding GkeCluster into gkeCluster: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["gkeCluster"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResource(c, f.KubernetesResource, res); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesResource into kubernetesResource: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesResource"] = v
	}

	return m, nil
}

// flattenMembershipEndpoint flattens an instance of MembershipEndpoint from a JSON
// response object.
func flattenMembershipEndpoint(c *Client, i interface{}, res *Membership) *MembershipEndpoint {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpoint{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpoint
	}
	r.GkeCluster = flattenMembershipEndpointGkeCluster(c, m["gkeCluster"], res)
	r.KubernetesMetadata = flattenMembershipEndpointKubernetesMetadata(c, m["kubernetesMetadata"], res)
	r.KubernetesResource = flattenMembershipEndpointKubernetesResource(c, m["kubernetesResource"], res)

	return r
}

// expandMembershipEndpointGkeClusterMap expands the contents of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeClusterMap(c *Client, f map[string]MembershipEndpointGkeCluster, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointGkeCluster(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointGkeClusterSlice expands the contents of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeClusterSlice(c *Client, f []MembershipEndpointGkeCluster, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointGkeCluster(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointGkeClusterMap flattens the contents of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeClusterMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointGkeCluster {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointGkeCluster{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointGkeCluster{}
	}

	items := make(map[string]MembershipEndpointGkeCluster)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointGkeCluster(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointGkeClusterSlice flattens the contents of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeClusterSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointGkeCluster {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointGkeCluster{}
	}

	if len(a) == 0 {
		return []MembershipEndpointGkeCluster{}
	}

	items := make([]MembershipEndpointGkeCluster, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointGkeCluster(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointGkeCluster expands an instance of MembershipEndpointGkeCluster into a JSON
// request object.
func expandMembershipEndpointGkeCluster(c *Client, f *MembershipEndpointGkeCluster, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandHubReferenceLink(c, f.ResourceLink, res); err != nil {
		return nil, fmt.Errorf("error expanding ResourceLink into resourceLink: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceLink"] = v
	}

	return m, nil
}

// flattenMembershipEndpointGkeCluster flattens an instance of MembershipEndpointGkeCluster from a JSON
// response object.
func flattenMembershipEndpointGkeCluster(c *Client, i interface{}, res *Membership) *MembershipEndpointGkeCluster {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointGkeCluster{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointGkeCluster
	}
	r.ResourceLink = flattenHubReferenceLink(c, m["resourceLink"], res)

	return r
}

// expandMembershipEndpointKubernetesMetadataMap expands the contents of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadataMap(c *Client, f map[string]MembershipEndpointKubernetesMetadata, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesMetadata(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesMetadataSlice expands the contents of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadataSlice(c *Client, f []MembershipEndpointKubernetesMetadata, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesMetadata(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesMetadataMap flattens the contents of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadataMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointKubernetesMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesMetadata{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesMetadata{}
	}

	items := make(map[string]MembershipEndpointKubernetesMetadata)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesMetadata(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointKubernetesMetadataSlice flattens the contents of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadataSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointKubernetesMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesMetadata{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesMetadata{}
	}

	items := make([]MembershipEndpointKubernetesMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesMetadata(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointKubernetesMetadata expands an instance of MembershipEndpointKubernetesMetadata into a JSON
// request object.
func expandMembershipEndpointKubernetesMetadata(c *Client, f *MembershipEndpointKubernetesMetadata, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipEndpointKubernetesMetadata flattens an instance of MembershipEndpointKubernetesMetadata from a JSON
// response object.
func flattenMembershipEndpointKubernetesMetadata(c *Client, i interface{}, res *Membership) *MembershipEndpointKubernetesMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesMetadata
	}
	r.KubernetesApiServerVersion = dcl.FlattenString(m["kubernetesApiServerVersion"])
	r.NodeProviderId = dcl.FlattenString(m["nodeProviderId"])
	r.NodeCount = dcl.FlattenInteger(m["nodeCount"])
	r.VcpuCount = dcl.FlattenInteger(m["vcpuCount"])
	r.MemoryMb = dcl.FlattenInteger(m["memoryMb"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// expandMembershipEndpointKubernetesResourceMap expands the contents of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMap(c *Client, f map[string]MembershipEndpointKubernetesResource, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResource(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceSlice expands the contents of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceSlice(c *Client, f []MembershipEndpointKubernetesResource, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResource(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceMap flattens the contents of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointKubernetesResource {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResource{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResource{}
	}

	items := make(map[string]MembershipEndpointKubernetesResource)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResource(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceSlice flattens the contents of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointKubernetesResource {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResource{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResource{}
	}

	items := make([]MembershipEndpointKubernetesResource, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResource(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointKubernetesResource expands an instance of MembershipEndpointKubernetesResource into a JSON
// request object.
func expandMembershipEndpointKubernetesResource(c *Client, f *MembershipEndpointKubernetesResource, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MembershipCrManifest; !dcl.IsEmptyValueIndirect(v) {
		m["membershipCrManifest"] = v
	}
	if v, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, f.ResourceOptions, res); err != nil {
		return nil, fmt.Errorf("error expanding ResourceOptions into resourceOptions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resourceOptions"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResource flattens an instance of MembershipEndpointKubernetesResource from a JSON
// response object.
func flattenMembershipEndpointKubernetesResource(c *Client, i interface{}, res *Membership) *MembershipEndpointKubernetesResource {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResource{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResource
	}
	r.MembershipCrManifest = dcl.FlattenSecretValue(m["membershipCrManifest"])
	r.MembershipResources = flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice(c, m["membershipResources"], res)
	r.ConnectResources = flattenMembershipEndpointKubernetesResourceConnectResourcesSlice(c, m["connectResources"], res)
	r.ResourceOptions = flattenMembershipEndpointKubernetesResourceResourceOptions(c, m["resourceOptions"], res)

	return r
}

// expandMembershipEndpointKubernetesResourceMembershipResourcesMap expands the contents of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResourcesMap(c *Client, f map[string]MembershipEndpointKubernetesResourceMembershipResources, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceMembershipResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceMembershipResourcesSlice expands the contents of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, f []MembershipEndpointKubernetesResourceMembershipResources, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceMembershipResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceMembershipResourcesMap flattens the contents of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResourcesMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointKubernetesResourceMembershipResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceMembershipResources{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceMembershipResources{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceMembershipResources)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceMembershipResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice flattens the contents of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResourcesSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointKubernetesResourceMembershipResources {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceMembershipResources{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceMembershipResources{}
	}

	items := make([]MembershipEndpointKubernetesResourceMembershipResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceMembershipResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceMembershipResources expands an instance of MembershipEndpointKubernetesResourceMembershipResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceMembershipResources(c *Client, f *MembershipEndpointKubernetesResourceMembershipResources, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Manifest; !dcl.IsEmptyValueIndirect(v) {
		m["manifest"] = v
	}
	if v := f.ClusterScoped; !dcl.IsEmptyValueIndirect(v) {
		m["clusterScoped"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceMembershipResources flattens an instance of MembershipEndpointKubernetesResourceMembershipResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceMembershipResources(c *Client, i interface{}, res *Membership) *MembershipEndpointKubernetesResourceMembershipResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceMembershipResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceMembershipResources
	}
	r.Manifest = dcl.FlattenString(m["manifest"])
	r.ClusterScoped = dcl.FlattenBool(m["clusterScoped"])

	return r
}

// expandMembershipEndpointKubernetesResourceConnectResourcesMap expands the contents of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResourcesMap(c *Client, f map[string]MembershipEndpointKubernetesResourceConnectResources, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceConnectResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceConnectResourcesSlice expands the contents of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, f []MembershipEndpointKubernetesResourceConnectResources, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceConnectResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceConnectResourcesMap flattens the contents of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResourcesMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointKubernetesResourceConnectResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceConnectResources{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceConnectResources{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceConnectResources)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceConnectResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceConnectResourcesSlice flattens the contents of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResourcesSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointKubernetesResourceConnectResources {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceConnectResources{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceConnectResources{}
	}

	items := make([]MembershipEndpointKubernetesResourceConnectResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceConnectResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceConnectResources expands an instance of MembershipEndpointKubernetesResourceConnectResources into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceConnectResources(c *Client, f *MembershipEndpointKubernetesResourceConnectResources, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Manifest; !dcl.IsEmptyValueIndirect(v) {
		m["manifest"] = v
	}
	if v := f.ClusterScoped; !dcl.IsEmptyValueIndirect(v) {
		m["clusterScoped"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceConnectResources flattens an instance of MembershipEndpointKubernetesResourceConnectResources from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceConnectResources(c *Client, i interface{}, res *Membership) *MembershipEndpointKubernetesResourceConnectResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceConnectResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceConnectResources
	}
	r.Manifest = dcl.FlattenString(m["manifest"])
	r.ClusterScoped = dcl.FlattenBool(m["clusterScoped"])

	return r
}

// expandMembershipEndpointKubernetesResourceResourceOptionsMap expands the contents of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptionsMap(c *Client, f map[string]MembershipEndpointKubernetesResourceResourceOptions, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipEndpointKubernetesResourceResourceOptionsSlice expands the contents of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, f []MembershipEndpointKubernetesResourceResourceOptions, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipEndpointKubernetesResourceResourceOptions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipEndpointKubernetesResourceResourceOptionsMap flattens the contents of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptionsMap(c *Client, i interface{}, res *Membership) map[string]MembershipEndpointKubernetesResourceResourceOptions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipEndpointKubernetesResourceResourceOptions{}
	}

	if len(a) == 0 {
		return map[string]MembershipEndpointKubernetesResourceResourceOptions{}
	}

	items := make(map[string]MembershipEndpointKubernetesResourceResourceOptions)
	for k, item := range a {
		items[k] = *flattenMembershipEndpointKubernetesResourceResourceOptions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipEndpointKubernetesResourceResourceOptionsSlice flattens the contents of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptionsSlice(c *Client, i interface{}, res *Membership) []MembershipEndpointKubernetesResourceResourceOptions {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipEndpointKubernetesResourceResourceOptions{}
	}

	if len(a) == 0 {
		return []MembershipEndpointKubernetesResourceResourceOptions{}
	}

	items := make([]MembershipEndpointKubernetesResourceResourceOptions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipEndpointKubernetesResourceResourceOptions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipEndpointKubernetesResourceResourceOptions expands an instance of MembershipEndpointKubernetesResourceResourceOptions into a JSON
// request object.
func expandMembershipEndpointKubernetesResourceResourceOptions(c *Client, f *MembershipEndpointKubernetesResourceResourceOptions, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ConnectVersion; !dcl.IsEmptyValueIndirect(v) {
		m["connectVersion"] = v
	}
	if v := f.V1Beta1Crd; !dcl.IsEmptyValueIndirect(v) {
		m["v1beta1Crd"] = v
	}

	return m, nil
}

// flattenMembershipEndpointKubernetesResourceResourceOptions flattens an instance of MembershipEndpointKubernetesResourceResourceOptions from a JSON
// response object.
func flattenMembershipEndpointKubernetesResourceResourceOptions(c *Client, i interface{}, res *Membership) *MembershipEndpointKubernetesResourceResourceOptions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipEndpointKubernetesResourceResourceOptions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipEndpointKubernetesResourceResourceOptions
	}
	r.ConnectVersion = dcl.FlattenString(m["connectVersion"])
	r.V1Beta1Crd = dcl.FlattenBool(m["v1beta1Crd"])

	return r
}

// expandMembershipStateMap expands the contents of MembershipState into a JSON
// request object.
func expandMembershipStateMap(c *Client, f map[string]MembershipState, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipState(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipStateSlice expands the contents of MembershipState into a JSON
// request object.
func expandMembershipStateSlice(c *Client, f []MembershipState, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipState(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipStateMap flattens the contents of MembershipState from a JSON
// response object.
func flattenMembershipStateMap(c *Client, i interface{}, res *Membership) map[string]MembershipState {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipState{}
	}

	if len(a) == 0 {
		return map[string]MembershipState{}
	}

	items := make(map[string]MembershipState)
	for k, item := range a {
		items[k] = *flattenMembershipState(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipStateSlice flattens the contents of MembershipState from a JSON
// response object.
func flattenMembershipStateSlice(c *Client, i interface{}, res *Membership) []MembershipState {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipState{}
	}

	if len(a) == 0 {
		return []MembershipState{}
	}

	items := make([]MembershipState, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipState(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipState expands an instance of MembershipState into a JSON
// request object.
func expandMembershipState(c *Client, f *MembershipState, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenMembershipState flattens an instance of MembershipState from a JSON
// response object.
func flattenMembershipState(c *Client, i interface{}, res *Membership) *MembershipState {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipState{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipState
	}
	r.Code = flattenMembershipStateCodeEnum(m["code"])

	return r
}

// expandMembershipAuthorityMap expands the contents of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthorityMap(c *Client, f map[string]MembershipAuthority, res *Membership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandMembershipAuthority(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandMembershipAuthoritySlice expands the contents of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthoritySlice(c *Client, f []MembershipAuthority, res *Membership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandMembershipAuthority(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenMembershipAuthorityMap flattens the contents of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthorityMap(c *Client, i interface{}, res *Membership) map[string]MembershipAuthority {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipAuthority{}
	}

	if len(a) == 0 {
		return map[string]MembershipAuthority{}
	}

	items := make(map[string]MembershipAuthority)
	for k, item := range a {
		items[k] = *flattenMembershipAuthority(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenMembershipAuthoritySlice flattens the contents of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthoritySlice(c *Client, i interface{}, res *Membership) []MembershipAuthority {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipAuthority{}
	}

	if len(a) == 0 {
		return []MembershipAuthority{}
	}

	items := make([]MembershipAuthority, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipAuthority(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandMembershipAuthority expands an instance of MembershipAuthority into a JSON
// request object.
func expandMembershipAuthority(c *Client, f *MembershipAuthority, res *Membership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Issuer; !dcl.IsEmptyValueIndirect(v) {
		m["issuer"] = v
	}

	return m, nil
}

// flattenMembershipAuthority flattens an instance of MembershipAuthority from a JSON
// response object.
func flattenMembershipAuthority(c *Client, i interface{}, res *Membership) *MembershipAuthority {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &MembershipAuthority{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyMembershipAuthority
	}
	r.Issuer = dcl.FlattenString(m["issuer"])
	r.WorkloadIdentityPool = dcl.FlattenString(m["workloadIdentityPool"])
	r.IdentityProvider = dcl.FlattenString(m["identityProvider"])

	return r
}

// flattenMembershipStateCodeEnumMap flattens the contents of MembershipStateCodeEnum from a JSON
// response object.
func flattenMembershipStateCodeEnumMap(c *Client, i interface{}, res *Membership) map[string]MembershipStateCodeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipStateCodeEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipStateCodeEnum{}
	}

	items := make(map[string]MembershipStateCodeEnum)
	for k, item := range a {
		items[k] = *flattenMembershipStateCodeEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipStateCodeEnumSlice flattens the contents of MembershipStateCodeEnum from a JSON
// response object.
func flattenMembershipStateCodeEnumSlice(c *Client, i interface{}, res *Membership) []MembershipStateCodeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipStateCodeEnum{}
	}

	if len(a) == 0 {
		return []MembershipStateCodeEnum{}
	}

	items := make([]MembershipStateCodeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipStateCodeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipStateCodeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipStateCodeEnum with the same value as that string.
func flattenMembershipStateCodeEnum(i interface{}) *MembershipStateCodeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MembershipStateCodeEnumRef(s)
}

// flattenMembershipInfrastructureTypeEnumMap flattens the contents of MembershipInfrastructureTypeEnum from a JSON
// response object.
func flattenMembershipInfrastructureTypeEnumMap(c *Client, i interface{}, res *Membership) map[string]MembershipInfrastructureTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]MembershipInfrastructureTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]MembershipInfrastructureTypeEnum{}
	}

	items := make(map[string]MembershipInfrastructureTypeEnum)
	for k, item := range a {
		items[k] = *flattenMembershipInfrastructureTypeEnum(item.(interface{}))
	}

	return items
}

// flattenMembershipInfrastructureTypeEnumSlice flattens the contents of MembershipInfrastructureTypeEnum from a JSON
// response object.
func flattenMembershipInfrastructureTypeEnumSlice(c *Client, i interface{}, res *Membership) []MembershipInfrastructureTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []MembershipInfrastructureTypeEnum{}
	}

	if len(a) == 0 {
		return []MembershipInfrastructureTypeEnum{}
	}

	items := make([]MembershipInfrastructureTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenMembershipInfrastructureTypeEnum(item.(interface{})))
	}

	return items
}

// flattenMembershipInfrastructureTypeEnum asserts that an interface is a string, and returns a
// pointer to a *MembershipInfrastructureTypeEnum with the same value as that string.
func flattenMembershipInfrastructureTypeEnum(i interface{}) *MembershipInfrastructureTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return MembershipInfrastructureTypeEnumRef(s)
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

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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
	vEndpoint := r.Endpoint
	if vEndpoint == nil {
		// note: explicitly not the empty object.
		vEndpoint = &MembershipEndpoint{}
	}
	if err := extractMembershipEndpointFields(r, vEndpoint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndpoint) {
		r.Endpoint = vEndpoint
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &MembershipState{}
	}
	if err := extractMembershipStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	vAuthority := r.Authority
	if vAuthority == nil {
		// note: explicitly not the empty object.
		vAuthority = &MembershipAuthority{}
	}
	if err := extractMembershipAuthorityFields(r, vAuthority); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAuthority) {
		r.Authority = vAuthority
	}
	return nil
}
func extractMembershipEndpointFields(r *Membership, o *MembershipEndpoint) error {
	vGkeCluster := o.GkeCluster
	if vGkeCluster == nil {
		// note: explicitly not the empty object.
		vGkeCluster = &MembershipEndpointGkeCluster{}
	}
	if err := extractMembershipEndpointGkeClusterFields(r, vGkeCluster); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGkeCluster) {
		o.GkeCluster = vGkeCluster
	}
	vKubernetesMetadata := o.KubernetesMetadata
	if vKubernetesMetadata == nil {
		// note: explicitly not the empty object.
		vKubernetesMetadata = &MembershipEndpointKubernetesMetadata{}
	}
	if err := extractMembershipEndpointKubernetesMetadataFields(r, vKubernetesMetadata); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetesMetadata) {
		o.KubernetesMetadata = vKubernetesMetadata
	}
	vKubernetesResource := o.KubernetesResource
	if vKubernetesResource == nil {
		// note: explicitly not the empty object.
		vKubernetesResource = &MembershipEndpointKubernetesResource{}
	}
	if err := extractMembershipEndpointKubernetesResourceFields(r, vKubernetesResource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetesResource) {
		o.KubernetesResource = vKubernetesResource
	}
	return nil
}
func extractMembershipEndpointGkeClusterFields(r *Membership, o *MembershipEndpointGkeCluster) error {
	return nil
}
func extractMembershipEndpointKubernetesMetadataFields(r *Membership, o *MembershipEndpointKubernetesMetadata) error {
	return nil
}
func extractMembershipEndpointKubernetesResourceFields(r *Membership, o *MembershipEndpointKubernetesResource) error {
	vResourceOptions := o.ResourceOptions
	if vResourceOptions == nil {
		// note: explicitly not the empty object.
		vResourceOptions = &MembershipEndpointKubernetesResourceResourceOptions{}
	}
	if err := extractMembershipEndpointKubernetesResourceResourceOptionsFields(r, vResourceOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceOptions) {
		o.ResourceOptions = vResourceOptions
	}
	return nil
}
func extractMembershipEndpointKubernetesResourceMembershipResourcesFields(r *Membership, o *MembershipEndpointKubernetesResourceMembershipResources) error {
	return nil
}
func extractMembershipEndpointKubernetesResourceConnectResourcesFields(r *Membership, o *MembershipEndpointKubernetesResourceConnectResources) error {
	return nil
}
func extractMembershipEndpointKubernetesResourceResourceOptionsFields(r *Membership, o *MembershipEndpointKubernetesResourceResourceOptions) error {
	return nil
}
func extractMembershipStateFields(r *Membership, o *MembershipState) error {
	return nil
}
func extractMembershipAuthorityFields(r *Membership, o *MembershipAuthority) error {
	return nil
}

func postReadExtractMembershipFields(r *Membership) error {
	vEndpoint := r.Endpoint
	if vEndpoint == nil {
		// note: explicitly not the empty object.
		vEndpoint = &MembershipEndpoint{}
	}
	if err := postReadExtractMembershipEndpointFields(r, vEndpoint); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndpoint) {
		r.Endpoint = vEndpoint
	}
	vState := r.State
	if vState == nil {
		// note: explicitly not the empty object.
		vState = &MembershipState{}
	}
	if err := postReadExtractMembershipStateFields(r, vState); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vState) {
		r.State = vState
	}
	vAuthority := r.Authority
	if vAuthority == nil {
		// note: explicitly not the empty object.
		vAuthority = &MembershipAuthority{}
	}
	if err := postReadExtractMembershipAuthorityFields(r, vAuthority); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAuthority) {
		r.Authority = vAuthority
	}
	return nil
}
func postReadExtractMembershipEndpointFields(r *Membership, o *MembershipEndpoint) error {
	vGkeCluster := o.GkeCluster
	if vGkeCluster == nil {
		// note: explicitly not the empty object.
		vGkeCluster = &MembershipEndpointGkeCluster{}
	}
	if err := extractMembershipEndpointGkeClusterFields(r, vGkeCluster); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGkeCluster) {
		o.GkeCluster = vGkeCluster
	}
	vKubernetesMetadata := o.KubernetesMetadata
	if vKubernetesMetadata == nil {
		// note: explicitly not the empty object.
		vKubernetesMetadata = &MembershipEndpointKubernetesMetadata{}
	}
	if err := extractMembershipEndpointKubernetesMetadataFields(r, vKubernetesMetadata); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetesMetadata) {
		o.KubernetesMetadata = vKubernetesMetadata
	}
	vKubernetesResource := o.KubernetesResource
	if vKubernetesResource == nil {
		// note: explicitly not the empty object.
		vKubernetesResource = &MembershipEndpointKubernetesResource{}
	}
	if err := extractMembershipEndpointKubernetesResourceFields(r, vKubernetesResource); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKubernetesResource) {
		o.KubernetesResource = vKubernetesResource
	}
	return nil
}
func postReadExtractMembershipEndpointGkeClusterFields(r *Membership, o *MembershipEndpointGkeCluster) error {
	return nil
}
func postReadExtractMembershipEndpointKubernetesMetadataFields(r *Membership, o *MembershipEndpointKubernetesMetadata) error {
	return nil
}
func postReadExtractMembershipEndpointKubernetesResourceFields(r *Membership, o *MembershipEndpointKubernetesResource) error {
	vResourceOptions := o.ResourceOptions
	if vResourceOptions == nil {
		// note: explicitly not the empty object.
		vResourceOptions = &MembershipEndpointKubernetesResourceResourceOptions{}
	}
	if err := extractMembershipEndpointKubernetesResourceResourceOptionsFields(r, vResourceOptions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResourceOptions) {
		o.ResourceOptions = vResourceOptions
	}
	return nil
}
func postReadExtractMembershipEndpointKubernetesResourceMembershipResourcesFields(r *Membership, o *MembershipEndpointKubernetesResourceMembershipResources) error {
	return nil
}
func postReadExtractMembershipEndpointKubernetesResourceConnectResourcesFields(r *Membership, o *MembershipEndpointKubernetesResourceConnectResources) error {
	return nil
}
func postReadExtractMembershipEndpointKubernetesResourceResourceOptionsFields(r *Membership, o *MembershipEndpointKubernetesResourceResourceOptions) error {
	return nil
}
func postReadExtractMembershipStateFields(r *Membership, o *MembershipState) error {
	return nil
}
func postReadExtractMembershipAuthorityFields(r *Membership, o *MembershipAuthority) error {
	return nil
}
