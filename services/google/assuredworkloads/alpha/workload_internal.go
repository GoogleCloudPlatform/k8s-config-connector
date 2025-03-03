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

func (r *Workload) validate() error {

	if err := dcl.Required(r, "displayName"); err != nil {
		return err
	}
	if err := dcl.Required(r, "complianceRegime"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Organization, "Organization"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.KmsSettings) {
		if err := r.KmsSettings.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SaaEnrollmentResponse) {
		if err := r.SaaEnrollmentResponse.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ComplianceStatus) {
		if err := r.ComplianceStatus.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PartnerPermissions) {
		if err := r.PartnerPermissions.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EkmProvisioningResponse) {
		if err := r.EkmProvisioningResponse.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *WorkloadResources) validate() error {
	return nil
}
func (r *WorkloadKmsSettings) validate() error {
	if err := dcl.Required(r, "nextRotationTime"); err != nil {
		return err
	}
	if err := dcl.Required(r, "rotationPeriod"); err != nil {
		return err
	}
	return nil
}
func (r *WorkloadResourceSettings) validate() error {
	return nil
}
func (r *WorkloadSaaEnrollmentResponse) validate() error {
	return nil
}
func (r *WorkloadComplianceStatus) validate() error {
	return nil
}
func (r *WorkloadPartnerPermissions) validate() error {
	return nil
}
func (r *WorkloadEkmProvisioningResponse) validate() error {
	return nil
}
func (r *Workload) basePath() string {
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(r.Location),
	}
	return dcl.Nprintf("https://{{location}}-assuredworkloads.googleapis.com/v1beta1/", params)
}

func (r *Workload) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(nr.Organization),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{organization}}/locations/{{location}}/workloads/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Workload) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(nr.Organization),
		"location":     dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("organizations/{{organization}}/locations/{{location}}/workloads", nr.basePath(), userBasePath, params), nil

}

func (r *Workload) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(nr.Organization),
		"location":     dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("organizations/{{organization}}/locations/{{location}}/workloads", nr.basePath(), userBasePath, params), nil

}

func (r *Workload) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"organization": dcl.ValueOrEmptyString(nr.Organization),
		"location":     dcl.ValueOrEmptyString(nr.Location),
		"name":         dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("organizations/{{organization}}/locations/{{location}}/workloads/{{name}}", nr.basePath(), userBasePath, params), nil
}

// workloadApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type workloadApiOperation interface {
	do(context.Context, *Workload, *Client) error
}

// newUpdateWorkloadUpdateWorkloadRequest creates a request for an
// Workload resource's UpdateWorkload update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWorkloadUpdateWorkloadRequest(ctx context.Context, f *Workload, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	b, err := c.getWorkloadRaw(ctx, f)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	rawEtag, err := dcl.GetMapEntry(
		m,
		[]string{"etag"},
	)
	if err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "Failed to fetch from JSON Path: %v", err)
	} else {
		req["etag"] = rawEtag.(string)
	}
	return req, nil
}

// marshalUpdateWorkloadUpdateWorkloadRequest converts the update into
// the final JSON request body.
func marshalUpdateWorkloadUpdateWorkloadRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWorkloadUpdateWorkloadOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWorkloadUpdateWorkloadOperation) do(ctx context.Context, r *Workload, c *Client) error {
	_, err := c.GetWorkload(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWorkload")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMaskWithPrefix(op.FieldDiffs, "Workload")
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWorkloadUpdateWorkloadRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWorkloadUpdateWorkloadRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listWorkloadRaw(ctx context.Context, r *Workload, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WorkloadMaxPage {
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

type listWorkloadOperation struct {
	Workloads []map[string]interface{} `json:"workloads"`
	Token     string                   `json:"nextPageToken"`
}

func (c *Client) listWorkload(ctx context.Context, r *Workload, pageToken string, pageSize int32) ([]*Workload, string, error) {
	b, err := c.listWorkloadRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWorkloadOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Workload
	for _, v := range m.Workloads {
		res, err := unmarshalMapWorkload(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Organization = r.Organization
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWorkload(ctx context.Context, f func(*Workload) bool, resources []*Workload) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWorkload(ctx, res)
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

type deleteWorkloadOperation struct{}

func (op *deleteWorkloadOperation) do(ctx context.Context, r *Workload, c *Client) error {
	r, err := c.GetWorkload(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Workload not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWorkload checking for existence. error: %v", err)
		return err
	}

	err = r.deleteResources(ctx, c)
	if err != nil {
		return err
	}
	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	_, err = dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return fmt.Errorf("failed to delete Workload: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetWorkload(ctx, r)
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
type createWorkloadOperation struct {
	response map[string]interface{}
}

func (op *createWorkloadOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWorkloadOperation) do(ctx context.Context, r *Workload, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a Workload with the wrong Name.
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

	if _, err := c.GetWorkload(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWorkloadRaw(ctx context.Context, r *Workload) ([]byte, error) {

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

func (c *Client) workloadDiffsForRawDesired(ctx context.Context, rawDesired *Workload, opts ...dcl.ApplyOption) (initial, desired *Workload, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Workload
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Workload); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Workload, got %T", sh)
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
		desired, err := canonicalizeWorkloadDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWorkload(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Workload resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Workload resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Workload resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWorkloadDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Workload: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Workload: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractWorkloadFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWorkloadInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Workload: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWorkloadDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Workload: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWorkload(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWorkloadInitialState(rawInitial, rawDesired *Workload) (*Workload, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWorkloadDesiredState(rawDesired, rawInitial *Workload, opts ...dcl.ApplyOption) (*Workload, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.KmsSettings = canonicalizeWorkloadKmsSettings(rawDesired.KmsSettings, nil, opts...)
		rawDesired.SaaEnrollmentResponse = canonicalizeWorkloadSaaEnrollmentResponse(rawDesired.SaaEnrollmentResponse, nil, opts...)
		rawDesired.ComplianceStatus = canonicalizeWorkloadComplianceStatus(rawDesired.ComplianceStatus, nil, opts...)
		rawDesired.PartnerPermissions = canonicalizeWorkloadPartnerPermissions(rawDesired.PartnerPermissions, nil, opts...)
		rawDesired.EkmProvisioningResponse = canonicalizeWorkloadEkmProvisioningResponse(rawDesired.EkmProvisioningResponse, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Workload{}
	if dcl.IsZeroValue(rawDesired.Name) || (dcl.IsEmptyValueIndirect(rawDesired.Name) && dcl.IsEmptyValueIndirect(rawInitial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.ComplianceRegime) || (dcl.IsEmptyValueIndirect(rawDesired.ComplianceRegime) && dcl.IsEmptyValueIndirect(rawInitial.ComplianceRegime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ComplianceRegime = rawInitial.ComplianceRegime
	} else {
		canonicalDesired.ComplianceRegime = rawDesired.ComplianceRegime
	}
	if dcl.StringCanonicalize(rawDesired.BillingAccount, rawInitial.BillingAccount) {
		canonicalDesired.BillingAccount = rawInitial.BillingAccount
	} else {
		canonicalDesired.BillingAccount = rawDesired.BillingAccount
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.StringCanonicalize(rawDesired.ProvisionedResourcesParent, rawInitial.ProvisionedResourcesParent) {
		canonicalDesired.ProvisionedResourcesParent = rawInitial.ProvisionedResourcesParent
	} else {
		canonicalDesired.ProvisionedResourcesParent = rawDesired.ProvisionedResourcesParent
	}
	canonicalDesired.KmsSettings = canonicalizeWorkloadKmsSettings(rawDesired.KmsSettings, rawInitial.KmsSettings, opts...)
	canonicalDesired.ResourceSettings = canonicalizeWorkloadResourceSettingsSlice(rawDesired.ResourceSettings, rawInitial.ResourceSettings, opts...)
	if dcl.BoolCanonicalize(rawDesired.EnableSovereignControls, rawInitial.EnableSovereignControls) {
		canonicalDesired.EnableSovereignControls = rawInitial.EnableSovereignControls
	} else {
		canonicalDesired.EnableSovereignControls = rawDesired.EnableSovereignControls
	}
	if dcl.IsZeroValue(rawDesired.Partner) || (dcl.IsEmptyValueIndirect(rawDesired.Partner) && dcl.IsEmptyValueIndirect(rawInitial.Partner)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Partner = rawInitial.Partner
	} else {
		canonicalDesired.Partner = rawDesired.Partner
	}
	canonicalDesired.PartnerPermissions = canonicalizeWorkloadPartnerPermissions(rawDesired.PartnerPermissions, rawInitial.PartnerPermissions, opts...)
	if dcl.BoolCanonicalize(rawDesired.ViolationNotificationsEnabled, rawInitial.ViolationNotificationsEnabled) {
		canonicalDesired.ViolationNotificationsEnabled = rawInitial.ViolationNotificationsEnabled
	} else {
		canonicalDesired.ViolationNotificationsEnabled = rawDesired.ViolationNotificationsEnabled
	}
	if dcl.NameToSelfLink(rawDesired.Organization, rawInitial.Organization) {
		canonicalDesired.Organization = rawInitial.Organization
	} else {
		canonicalDesired.Organization = rawDesired.Organization
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeWorkloadNewState(c *Client, rawNew, rawDesired *Workload) (*Workload, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Resources) && dcl.IsEmptyValueIndirect(rawDesired.Resources) {
		rawNew.Resources = rawDesired.Resources
	} else {
		rawNew.Resources = canonicalizeNewWorkloadResourcesSlice(c, rawDesired.Resources, rawNew.Resources)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ComplianceRegime) && dcl.IsEmptyValueIndirect(rawDesired.ComplianceRegime) {
		rawNew.ComplianceRegime = rawDesired.ComplianceRegime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.BillingAccount) && dcl.IsEmptyValueIndirect(rawDesired.BillingAccount) {
		rawNew.BillingAccount = rawDesired.BillingAccount
	} else {
		rawNew.BillingAccount = rawDesired.BillingAccount
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	rawNew.ProvisionedResourcesParent = rawDesired.ProvisionedResourcesParent

	rawNew.KmsSettings = rawDesired.KmsSettings

	rawNew.ResourceSettings = rawDesired.ResourceSettings

	if dcl.IsEmptyValueIndirect(rawNew.KajEnrollmentState) && dcl.IsEmptyValueIndirect(rawDesired.KajEnrollmentState) {
		rawNew.KajEnrollmentState = rawDesired.KajEnrollmentState
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.EnableSovereignControls) && dcl.IsEmptyValueIndirect(rawDesired.EnableSovereignControls) {
		rawNew.EnableSovereignControls = rawDesired.EnableSovereignControls
	} else {
		if dcl.BoolCanonicalize(rawDesired.EnableSovereignControls, rawNew.EnableSovereignControls) {
			rawNew.EnableSovereignControls = rawDesired.EnableSovereignControls
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SaaEnrollmentResponse) && dcl.IsEmptyValueIndirect(rawDesired.SaaEnrollmentResponse) {
		rawNew.SaaEnrollmentResponse = rawDesired.SaaEnrollmentResponse
	} else {
		rawNew.SaaEnrollmentResponse = canonicalizeNewWorkloadSaaEnrollmentResponse(c, rawDesired.SaaEnrollmentResponse, rawNew.SaaEnrollmentResponse)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ComplianceStatus) && dcl.IsEmptyValueIndirect(rawDesired.ComplianceStatus) {
		rawNew.ComplianceStatus = rawDesired.ComplianceStatus
	} else {
		rawNew.ComplianceStatus = canonicalizeNewWorkloadComplianceStatus(c, rawDesired.ComplianceStatus, rawNew.ComplianceStatus)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CompliantButDisallowedServices) && dcl.IsEmptyValueIndirect(rawDesired.CompliantButDisallowedServices) {
		rawNew.CompliantButDisallowedServices = rawDesired.CompliantButDisallowedServices
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.CompliantButDisallowedServices, rawNew.CompliantButDisallowedServices) {
			rawNew.CompliantButDisallowedServices = rawDesired.CompliantButDisallowedServices
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Partner) && dcl.IsEmptyValueIndirect(rawDesired.Partner) {
		rawNew.Partner = rawDesired.Partner
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.PartnerPermissions) && dcl.IsEmptyValueIndirect(rawDesired.PartnerPermissions) {
		rawNew.PartnerPermissions = rawDesired.PartnerPermissions
	} else {
		rawNew.PartnerPermissions = canonicalizeNewWorkloadPartnerPermissions(c, rawDesired.PartnerPermissions, rawNew.PartnerPermissions)
	}

	if dcl.IsEmptyValueIndirect(rawNew.EkmProvisioningResponse) && dcl.IsEmptyValueIndirect(rawDesired.EkmProvisioningResponse) {
		rawNew.EkmProvisioningResponse = rawDesired.EkmProvisioningResponse
	} else {
		rawNew.EkmProvisioningResponse = canonicalizeNewWorkloadEkmProvisioningResponse(c, rawDesired.EkmProvisioningResponse, rawNew.EkmProvisioningResponse)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ViolationNotificationsEnabled) && dcl.IsEmptyValueIndirect(rawDesired.ViolationNotificationsEnabled) {
		rawNew.ViolationNotificationsEnabled = rawDesired.ViolationNotificationsEnabled
	} else {
		if dcl.BoolCanonicalize(rawDesired.ViolationNotificationsEnabled, rawNew.ViolationNotificationsEnabled) {
			rawNew.ViolationNotificationsEnabled = rawDesired.ViolationNotificationsEnabled
		}
	}

	rawNew.Organization = rawDesired.Organization

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeWorkloadResources(des, initial *WorkloadResources, opts ...dcl.ApplyOption) *WorkloadResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadResources{}

	if dcl.IsZeroValue(des.ResourceId) || (dcl.IsEmptyValueIndirect(des.ResourceId) && dcl.IsEmptyValueIndirect(initial.ResourceId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResourceId = initial.ResourceId
	} else {
		cDes.ResourceId = des.ResourceId
	}
	if dcl.IsZeroValue(des.ResourceType) || (dcl.IsEmptyValueIndirect(des.ResourceType) && dcl.IsEmptyValueIndirect(initial.ResourceType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResourceType = initial.ResourceType
	} else {
		cDes.ResourceType = des.ResourceType
	}

	return cDes
}

func canonicalizeWorkloadResourcesSlice(des, initial []WorkloadResources, opts ...dcl.ApplyOption) []WorkloadResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadResources(c *Client, des, nw *WorkloadResources) *WorkloadResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewWorkloadResourcesSet(c *Client, des, nw []WorkloadResources) []WorkloadResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadResourcesSlice(c *Client, des, nw []WorkloadResources) []WorkloadResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadResources(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadKmsSettings(des, initial *WorkloadKmsSettings, opts ...dcl.ApplyOption) *WorkloadKmsSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadKmsSettings{}

	if dcl.IsZeroValue(des.NextRotationTime) || (dcl.IsEmptyValueIndirect(des.NextRotationTime) && dcl.IsEmptyValueIndirect(initial.NextRotationTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NextRotationTime = initial.NextRotationTime
	} else {
		cDes.NextRotationTime = des.NextRotationTime
	}
	if dcl.StringCanonicalize(des.RotationPeriod, initial.RotationPeriod) || dcl.IsZeroValue(des.RotationPeriod) {
		cDes.RotationPeriod = initial.RotationPeriod
	} else {
		cDes.RotationPeriod = des.RotationPeriod
	}

	return cDes
}

func canonicalizeWorkloadKmsSettingsSlice(des, initial []WorkloadKmsSettings, opts ...dcl.ApplyOption) []WorkloadKmsSettings {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadKmsSettings, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadKmsSettings(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadKmsSettings, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadKmsSettings(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadKmsSettings(c *Client, des, nw *WorkloadKmsSettings) *WorkloadKmsSettings {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadKmsSettings while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.RotationPeriod, nw.RotationPeriod) {
		nw.RotationPeriod = des.RotationPeriod
	}

	return nw
}

func canonicalizeNewWorkloadKmsSettingsSet(c *Client, des, nw []WorkloadKmsSettings) []WorkloadKmsSettings {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadKmsSettings
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadKmsSettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadKmsSettings(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadKmsSettingsSlice(c *Client, des, nw []WorkloadKmsSettings) []WorkloadKmsSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadKmsSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadKmsSettings(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadResourceSettings(des, initial *WorkloadResourceSettings, opts ...dcl.ApplyOption) *WorkloadResourceSettings {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadResourceSettings{}

	if dcl.StringCanonicalize(des.ResourceId, initial.ResourceId) || dcl.IsZeroValue(des.ResourceId) {
		cDes.ResourceId = initial.ResourceId
	} else {
		cDes.ResourceId = des.ResourceId
	}
	if dcl.IsZeroValue(des.ResourceType) || (dcl.IsEmptyValueIndirect(des.ResourceType) && dcl.IsEmptyValueIndirect(initial.ResourceType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ResourceType = initial.ResourceType
	} else {
		cDes.ResourceType = des.ResourceType
	}
	if dcl.StringCanonicalize(des.DisplayName, initial.DisplayName) || dcl.IsZeroValue(des.DisplayName) {
		cDes.DisplayName = initial.DisplayName
	} else {
		cDes.DisplayName = des.DisplayName
	}

	return cDes
}

func canonicalizeWorkloadResourceSettingsSlice(des, initial []WorkloadResourceSettings, opts ...dcl.ApplyOption) []WorkloadResourceSettings {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadResourceSettings, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadResourceSettings(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadResourceSettings, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadResourceSettings(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadResourceSettings(c *Client, des, nw *WorkloadResourceSettings) *WorkloadResourceSettings {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadResourceSettings while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ResourceId, nw.ResourceId) {
		nw.ResourceId = des.ResourceId
	}
	if dcl.StringCanonicalize(des.DisplayName, nw.DisplayName) {
		nw.DisplayName = des.DisplayName
	}

	return nw
}

func canonicalizeNewWorkloadResourceSettingsSet(c *Client, des, nw []WorkloadResourceSettings) []WorkloadResourceSettings {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadResourceSettings
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadResourceSettingsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadResourceSettings(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadResourceSettingsSlice(c *Client, des, nw []WorkloadResourceSettings) []WorkloadResourceSettings {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadResourceSettings
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadResourceSettings(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadSaaEnrollmentResponse(des, initial *WorkloadSaaEnrollmentResponse, opts ...dcl.ApplyOption) *WorkloadSaaEnrollmentResponse {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadSaaEnrollmentResponse{}

	if dcl.IsZeroValue(des.SetupErrors) || (dcl.IsEmptyValueIndirect(des.SetupErrors) && dcl.IsEmptyValueIndirect(initial.SetupErrors)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SetupErrors = initial.SetupErrors
	} else {
		cDes.SetupErrors = des.SetupErrors
	}
	if dcl.IsZeroValue(des.SetupStatus) || (dcl.IsEmptyValueIndirect(des.SetupStatus) && dcl.IsEmptyValueIndirect(initial.SetupStatus)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SetupStatus = initial.SetupStatus
	} else {
		cDes.SetupStatus = des.SetupStatus
	}

	return cDes
}

func canonicalizeWorkloadSaaEnrollmentResponseSlice(des, initial []WorkloadSaaEnrollmentResponse, opts ...dcl.ApplyOption) []WorkloadSaaEnrollmentResponse {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadSaaEnrollmentResponse, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadSaaEnrollmentResponse(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadSaaEnrollmentResponse, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadSaaEnrollmentResponse(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadSaaEnrollmentResponse(c *Client, des, nw *WorkloadSaaEnrollmentResponse) *WorkloadSaaEnrollmentResponse {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadSaaEnrollmentResponse while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewWorkloadSaaEnrollmentResponseSet(c *Client, des, nw []WorkloadSaaEnrollmentResponse) []WorkloadSaaEnrollmentResponse {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadSaaEnrollmentResponse
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadSaaEnrollmentResponseNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadSaaEnrollmentResponse(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadSaaEnrollmentResponseSlice(c *Client, des, nw []WorkloadSaaEnrollmentResponse) []WorkloadSaaEnrollmentResponse {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadSaaEnrollmentResponse
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadSaaEnrollmentResponse(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadComplianceStatus(des, initial *WorkloadComplianceStatus, opts ...dcl.ApplyOption) *WorkloadComplianceStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadComplianceStatus{}

	if dcl.IsZeroValue(des.ActiveViolationCount) || (dcl.IsEmptyValueIndirect(des.ActiveViolationCount) && dcl.IsEmptyValueIndirect(initial.ActiveViolationCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ActiveViolationCount = initial.ActiveViolationCount
	} else {
		cDes.ActiveViolationCount = des.ActiveViolationCount
	}
	if dcl.IsZeroValue(des.AcknowledgedViolationCount) || (dcl.IsEmptyValueIndirect(des.AcknowledgedViolationCount) && dcl.IsEmptyValueIndirect(initial.AcknowledgedViolationCount)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AcknowledgedViolationCount = initial.AcknowledgedViolationCount
	} else {
		cDes.AcknowledgedViolationCount = des.AcknowledgedViolationCount
	}

	return cDes
}

func canonicalizeWorkloadComplianceStatusSlice(des, initial []WorkloadComplianceStatus, opts ...dcl.ApplyOption) []WorkloadComplianceStatus {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadComplianceStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadComplianceStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadComplianceStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadComplianceStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadComplianceStatus(c *Client, des, nw *WorkloadComplianceStatus) *WorkloadComplianceStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadComplianceStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewWorkloadComplianceStatusSet(c *Client, des, nw []WorkloadComplianceStatus) []WorkloadComplianceStatus {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadComplianceStatus
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadComplianceStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadComplianceStatus(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadComplianceStatusSlice(c *Client, des, nw []WorkloadComplianceStatus) []WorkloadComplianceStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadComplianceStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadComplianceStatus(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadPartnerPermissions(des, initial *WorkloadPartnerPermissions, opts ...dcl.ApplyOption) *WorkloadPartnerPermissions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadPartnerPermissions{}

	if dcl.BoolCanonicalize(des.DataLogsViewer, initial.DataLogsViewer) || dcl.IsZeroValue(des.DataLogsViewer) {
		cDes.DataLogsViewer = initial.DataLogsViewer
	} else {
		cDes.DataLogsViewer = des.DataLogsViewer
	}
	if dcl.BoolCanonicalize(des.ServiceAccessApprover, initial.ServiceAccessApprover) || dcl.IsZeroValue(des.ServiceAccessApprover) {
		cDes.ServiceAccessApprover = initial.ServiceAccessApprover
	} else {
		cDes.ServiceAccessApprover = des.ServiceAccessApprover
	}
	if dcl.BoolCanonicalize(des.AssuredWorkloadsMonitoring, initial.AssuredWorkloadsMonitoring) || dcl.IsZeroValue(des.AssuredWorkloadsMonitoring) {
		cDes.AssuredWorkloadsMonitoring = initial.AssuredWorkloadsMonitoring
	} else {
		cDes.AssuredWorkloadsMonitoring = des.AssuredWorkloadsMonitoring
	}

	return cDes
}

func canonicalizeWorkloadPartnerPermissionsSlice(des, initial []WorkloadPartnerPermissions, opts ...dcl.ApplyOption) []WorkloadPartnerPermissions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadPartnerPermissions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadPartnerPermissions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadPartnerPermissions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadPartnerPermissions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadPartnerPermissions(c *Client, des, nw *WorkloadPartnerPermissions) *WorkloadPartnerPermissions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadPartnerPermissions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.DataLogsViewer, nw.DataLogsViewer) {
		nw.DataLogsViewer = des.DataLogsViewer
	}
	if dcl.BoolCanonicalize(des.ServiceAccessApprover, nw.ServiceAccessApprover) {
		nw.ServiceAccessApprover = des.ServiceAccessApprover
	}
	if dcl.BoolCanonicalize(des.AssuredWorkloadsMonitoring, nw.AssuredWorkloadsMonitoring) {
		nw.AssuredWorkloadsMonitoring = des.AssuredWorkloadsMonitoring
	}

	return nw
}

func canonicalizeNewWorkloadPartnerPermissionsSet(c *Client, des, nw []WorkloadPartnerPermissions) []WorkloadPartnerPermissions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadPartnerPermissions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadPartnerPermissionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadPartnerPermissions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadPartnerPermissionsSlice(c *Client, des, nw []WorkloadPartnerPermissions) []WorkloadPartnerPermissions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadPartnerPermissions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadPartnerPermissions(c, &d, &n))
	}

	return items
}

func canonicalizeWorkloadEkmProvisioningResponse(des, initial *WorkloadEkmProvisioningResponse, opts ...dcl.ApplyOption) *WorkloadEkmProvisioningResponse {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &WorkloadEkmProvisioningResponse{}

	if dcl.IsZeroValue(des.EkmProvisioningState) || (dcl.IsEmptyValueIndirect(des.EkmProvisioningState) && dcl.IsEmptyValueIndirect(initial.EkmProvisioningState)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EkmProvisioningState = initial.EkmProvisioningState
	} else {
		cDes.EkmProvisioningState = des.EkmProvisioningState
	}
	if dcl.IsZeroValue(des.EkmProvisioningErrorDomain) || (dcl.IsEmptyValueIndirect(des.EkmProvisioningErrorDomain) && dcl.IsEmptyValueIndirect(initial.EkmProvisioningErrorDomain)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EkmProvisioningErrorDomain = initial.EkmProvisioningErrorDomain
	} else {
		cDes.EkmProvisioningErrorDomain = des.EkmProvisioningErrorDomain
	}
	if dcl.IsZeroValue(des.EkmProvisioningErrorMapping) || (dcl.IsEmptyValueIndirect(des.EkmProvisioningErrorMapping) && dcl.IsEmptyValueIndirect(initial.EkmProvisioningErrorMapping)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EkmProvisioningErrorMapping = initial.EkmProvisioningErrorMapping
	} else {
		cDes.EkmProvisioningErrorMapping = des.EkmProvisioningErrorMapping
	}

	return cDes
}

func canonicalizeWorkloadEkmProvisioningResponseSlice(des, initial []WorkloadEkmProvisioningResponse, opts ...dcl.ApplyOption) []WorkloadEkmProvisioningResponse {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]WorkloadEkmProvisioningResponse, 0, len(des))
		for _, d := range des {
			cd := canonicalizeWorkloadEkmProvisioningResponse(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]WorkloadEkmProvisioningResponse, 0, len(des))
	for i, d := range des {
		cd := canonicalizeWorkloadEkmProvisioningResponse(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewWorkloadEkmProvisioningResponse(c *Client, des, nw *WorkloadEkmProvisioningResponse) *WorkloadEkmProvisioningResponse {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for WorkloadEkmProvisioningResponse while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewWorkloadEkmProvisioningResponseSet(c *Client, des, nw []WorkloadEkmProvisioningResponse) []WorkloadEkmProvisioningResponse {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []WorkloadEkmProvisioningResponse
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareWorkloadEkmProvisioningResponseNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewWorkloadEkmProvisioningResponse(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewWorkloadEkmProvisioningResponseSlice(c *Client, des, nw []WorkloadEkmProvisioningResponse) []WorkloadEkmProvisioningResponse {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []WorkloadEkmProvisioningResponse
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewWorkloadEkmProvisioningResponse(c, &d, &n))
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
func diffWorkload(c *Client, desired, actual *Workload, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadUpdateWorkloadOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareWorkloadResourcesNewStyle, EmptyObject: EmptyWorkloadResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ComplianceRegime, actual.ComplianceRegime, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ComplianceRegime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.BillingAccount, actual.BillingAccount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BillingAccount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWorkloadUpdateWorkloadOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProvisionedResourcesParent, actual.ProvisionedResourcesParent, dcl.DiffInfo{Ignore: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProvisionedResourcesParent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KmsSettings, actual.KmsSettings, dcl.DiffInfo{Ignore: true, ObjectFunction: compareWorkloadKmsSettingsNewStyle, EmptyObject: EmptyWorkloadKmsSettings, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceSettings, actual.ResourceSettings, dcl.DiffInfo{Ignore: true, ObjectFunction: compareWorkloadResourceSettingsNewStyle, EmptyObject: EmptyWorkloadResourceSettings, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceSettings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KajEnrollmentState, actual.KajEnrollmentState, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KajEnrollmentState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableSovereignControls, actual.EnableSovereignControls, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableSovereignControls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SaaEnrollmentResponse, actual.SaaEnrollmentResponse, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareWorkloadSaaEnrollmentResponseNewStyle, EmptyObject: EmptyWorkloadSaaEnrollmentResponse, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SaaEnrollmentResponse")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ComplianceStatus, actual.ComplianceStatus, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareWorkloadComplianceStatusNewStyle, EmptyObject: EmptyWorkloadComplianceStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ComplianceStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CompliantButDisallowedServices, actual.CompliantButDisallowedServices, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CompliantButDisallowedServices")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Partner, actual.Partner, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Partner")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PartnerPermissions, actual.PartnerPermissions, dcl.DiffInfo{ObjectFunction: compareWorkloadPartnerPermissionsNewStyle, EmptyObject: EmptyWorkloadPartnerPermissions, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PartnerPermissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EkmProvisioningResponse, actual.EkmProvisioningResponse, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareWorkloadEkmProvisioningResponseNewStyle, EmptyObject: EmptyWorkloadEkmProvisioningResponse, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EkmProvisioningResponse")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ViolationNotificationsEnabled, actual.ViolationNotificationsEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ViolationNotificationsEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Organization, actual.Organization, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Organization")); len(ds) != 0 || err != nil {
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
func compareWorkloadResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadResources)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadResources or *WorkloadResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadResources)
	if !ok {
		actualNotPointer, ok := a.(WorkloadResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceId, actual.ResourceId, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadKmsSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadKmsSettings)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadKmsSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadKmsSettings or *WorkloadKmsSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadKmsSettings)
	if !ok {
		actualNotPointer, ok := a.(WorkloadKmsSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadKmsSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NextRotationTime, actual.NextRotationTime, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NextRotationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RotationPeriod, actual.RotationPeriod, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RotationPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadResourceSettingsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadResourceSettings)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadResourceSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadResourceSettings or *WorkloadResourceSettings", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadResourceSettings)
	if !ok {
		actualNotPointer, ok := a.(WorkloadResourceSettings)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadResourceSettings", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceId, actual.ResourceId, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadSaaEnrollmentResponseNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadSaaEnrollmentResponse)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadSaaEnrollmentResponse)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadSaaEnrollmentResponse or *WorkloadSaaEnrollmentResponse", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadSaaEnrollmentResponse)
	if !ok {
		actualNotPointer, ok := a.(WorkloadSaaEnrollmentResponse)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadSaaEnrollmentResponse", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SetupErrors, actual.SetupErrors, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SetupErrors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SetupStatus, actual.SetupStatus, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SetupStatus")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadComplianceStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadComplianceStatus)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadComplianceStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadComplianceStatus or *WorkloadComplianceStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadComplianceStatus)
	if !ok {
		actualNotPointer, ok := a.(WorkloadComplianceStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadComplianceStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ActiveViolationCount, actual.ActiveViolationCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ActiveViolationCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AcknowledgedViolationCount, actual.AcknowledgedViolationCount, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AcknowledgedViolationCount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadPartnerPermissionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadPartnerPermissions)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadPartnerPermissions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadPartnerPermissions or *WorkloadPartnerPermissions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadPartnerPermissions)
	if !ok {
		actualNotPointer, ok := a.(WorkloadPartnerPermissions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadPartnerPermissions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DataLogsViewer, actual.DataLogsViewer, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DataLogsViewer")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ServiceAccessApprover, actual.ServiceAccessApprover, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ServiceAccessApprover")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AssuredWorkloadsMonitoring, actual.AssuredWorkloadsMonitoring, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AssuredWorkloadsMonitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareWorkloadEkmProvisioningResponseNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*WorkloadEkmProvisioningResponse)
	if !ok {
		desiredNotPointer, ok := d.(WorkloadEkmProvisioningResponse)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadEkmProvisioningResponse or *WorkloadEkmProvisioningResponse", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*WorkloadEkmProvisioningResponse)
	if !ok {
		actualNotPointer, ok := a.(WorkloadEkmProvisioningResponse)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a WorkloadEkmProvisioningResponse", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EkmProvisioningState, actual.EkmProvisioningState, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EkmProvisioningState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EkmProvisioningErrorDomain, actual.EkmProvisioningErrorDomain, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EkmProvisioningErrorDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EkmProvisioningErrorMapping, actual.EkmProvisioningErrorMapping, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EkmProvisioningErrorMapping")); len(ds) != 0 || err != nil {
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
func (r *Workload) urlNormalized() *Workload {
	normalized := dcl.Copy(*r).(Workload)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.BillingAccount = dcl.SelfLinkToName(r.BillingAccount)
	normalized.ProvisionedResourcesParent = dcl.SelfLinkToName(r.ProvisionedResourcesParent)
	normalized.Organization = dcl.SelfLinkToName(r.Organization)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Workload) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWorkload" {
		fields := map[string]interface{}{
			"organization": dcl.ValueOrEmptyString(nr.Organization),
			"location":     dcl.ValueOrEmptyString(nr.Location),
			"name":         dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("organizations/{{organization}}/locations/{{location}}/workloads/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Workload resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Workload) marshal(c *Client) ([]byte, error) {
	m, err := expandWorkload(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Workload: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWorkload decodes JSON responses into the Workload resource schema.
func unmarshalWorkload(b []byte, c *Client, res *Workload) (*Workload, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWorkload(m, c, res)
}

func unmarshalMapWorkload(m map[string]interface{}, c *Client, res *Workload) (*Workload, error) {

	flattened := flattenWorkload(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandWorkload expands Workload into a JSON request object.
func expandWorkload(c *Client, f *Workload) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("organizations/%s/locations/%s/workloads/%s", f.Name, dcl.SelfLinkToName(f.Organization), dcl.SelfLinkToName(f.Location), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.ComplianceRegime; dcl.ValueShouldBeSent(v) {
		m["complianceRegime"] = v
	}
	if v := f.BillingAccount; dcl.ValueShouldBeSent(v) {
		m["billingAccount"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.ProvisionedResourcesParent; dcl.ValueShouldBeSent(v) {
		m["provisionedResourcesParent"] = v
	}
	if v, err := expandWorkloadKmsSettings(c, f.KmsSettings, res); err != nil {
		return nil, fmt.Errorf("error expanding KmsSettings into kmsSettings: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kmsSettings"] = v
	}
	if v, err := expandWorkloadResourceSettingsSlice(c, f.ResourceSettings, res); err != nil {
		return nil, fmt.Errorf("error expanding ResourceSettings into resourceSettings: %w", err)
	} else if v != nil {
		m["resourceSettings"] = v
	}
	if v := f.EnableSovereignControls; dcl.ValueShouldBeSent(v) {
		m["enableSovereignControls"] = v
	}
	if v := f.Partner; dcl.ValueShouldBeSent(v) {
		m["partner"] = v
	}
	if v, err := expandWorkloadPartnerPermissions(c, f.PartnerPermissions, res); err != nil {
		return nil, fmt.Errorf("error expanding PartnerPermissions into partnerPermissions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["partnerPermissions"] = v
	}
	if v := f.ViolationNotificationsEnabled; dcl.ValueShouldBeSent(v) {
		m["violationNotificationsEnabled"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Organization into organization: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["organization"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenWorkload flattens Workload from a JSON request object into the
// Workload type.
func flattenWorkload(c *Client, i interface{}, res *Workload) *Workload {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Workload{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Resources = flattenWorkloadResourcesSlice(c, m["resources"], res)
	resultRes.ComplianceRegime = flattenWorkloadComplianceRegimeEnum(m["complianceRegime"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.BillingAccount = dcl.FlattenString(m["billingAccount"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.ProvisionedResourcesParent = dcl.FlattenSecretValue(m["provisionedResourcesParent"])
	resultRes.KmsSettings = flattenWorkloadKmsSettings(c, m["kmsSettings"], res)
	resultRes.ResourceSettings = flattenWorkloadResourceSettingsSlice(c, m["resourceSettings"], res)
	resultRes.KajEnrollmentState = flattenWorkloadKajEnrollmentStateEnum(m["kajEnrollmentState"])
	resultRes.EnableSovereignControls = dcl.FlattenBool(m["enableSovereignControls"])
	resultRes.SaaEnrollmentResponse = flattenWorkloadSaaEnrollmentResponse(c, m["saaEnrollmentResponse"], res)
	resultRes.ComplianceStatus = flattenWorkloadComplianceStatus(c, m["complianceStatus"], res)
	resultRes.CompliantButDisallowedServices = dcl.FlattenStringSlice(m["compliantButDisallowedServices"])
	resultRes.Partner = flattenWorkloadPartnerEnum(m["partner"])
	resultRes.PartnerPermissions = flattenWorkloadPartnerPermissions(c, m["partnerPermissions"], res)
	resultRes.EkmProvisioningResponse = flattenWorkloadEkmProvisioningResponse(c, m["ekmProvisioningResponse"], res)
	resultRes.ViolationNotificationsEnabled = dcl.FlattenBool(m["violationNotificationsEnabled"])
	resultRes.Organization = dcl.FlattenString(m["organization"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandWorkloadResourcesMap expands the contents of WorkloadResources into a JSON
// request object.
func expandWorkloadResourcesMap(c *Client, f map[string]WorkloadResources, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadResourcesSlice expands the contents of WorkloadResources into a JSON
// request object.
func expandWorkloadResourcesSlice(c *Client, f []WorkloadResources, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadResourcesMap flattens the contents of WorkloadResources from a JSON
// response object.
func flattenWorkloadResourcesMap(c *Client, i interface{}, res *Workload) map[string]WorkloadResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadResources{}
	}

	if len(a) == 0 {
		return map[string]WorkloadResources{}
	}

	items := make(map[string]WorkloadResources)
	for k, item := range a {
		items[k] = *flattenWorkloadResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadResourcesSlice flattens the contents of WorkloadResources from a JSON
// response object.
func flattenWorkloadResourcesSlice(c *Client, i interface{}, res *Workload) []WorkloadResources {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadResources{}
	}

	if len(a) == 0 {
		return []WorkloadResources{}
	}

	items := make([]WorkloadResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadResources expands an instance of WorkloadResources into a JSON
// request object.
func expandWorkloadResources(c *Client, f *WorkloadResources, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceId"] = v
	}
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}

	return m, nil
}

// flattenWorkloadResources flattens an instance of WorkloadResources from a JSON
// response object.
func flattenWorkloadResources(c *Client, i interface{}, res *Workload) *WorkloadResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadResources
	}
	r.ResourceId = dcl.FlattenInteger(m["resourceId"])
	r.ResourceType = flattenWorkloadResourcesResourceTypeEnum(m["resourceType"])

	return r
}

// expandWorkloadKmsSettingsMap expands the contents of WorkloadKmsSettings into a JSON
// request object.
func expandWorkloadKmsSettingsMap(c *Client, f map[string]WorkloadKmsSettings, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadKmsSettings(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadKmsSettingsSlice expands the contents of WorkloadKmsSettings into a JSON
// request object.
func expandWorkloadKmsSettingsSlice(c *Client, f []WorkloadKmsSettings, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadKmsSettings(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadKmsSettingsMap flattens the contents of WorkloadKmsSettings from a JSON
// response object.
func flattenWorkloadKmsSettingsMap(c *Client, i interface{}, res *Workload) map[string]WorkloadKmsSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadKmsSettings{}
	}

	if len(a) == 0 {
		return map[string]WorkloadKmsSettings{}
	}

	items := make(map[string]WorkloadKmsSettings)
	for k, item := range a {
		items[k] = *flattenWorkloadKmsSettings(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadKmsSettingsSlice flattens the contents of WorkloadKmsSettings from a JSON
// response object.
func flattenWorkloadKmsSettingsSlice(c *Client, i interface{}, res *Workload) []WorkloadKmsSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadKmsSettings{}
	}

	if len(a) == 0 {
		return []WorkloadKmsSettings{}
	}

	items := make([]WorkloadKmsSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadKmsSettings(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadKmsSettings expands an instance of WorkloadKmsSettings into a JSON
// request object.
func expandWorkloadKmsSettings(c *Client, f *WorkloadKmsSettings, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NextRotationTime; !dcl.IsEmptyValueIndirect(v) {
		m["nextRotationTime"] = v
	}
	if v := f.RotationPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["rotationPeriod"] = v
	}

	return m, nil
}

// flattenWorkloadKmsSettings flattens an instance of WorkloadKmsSettings from a JSON
// response object.
func flattenWorkloadKmsSettings(c *Client, i interface{}, res *Workload) *WorkloadKmsSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadKmsSettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadKmsSettings
	}
	r.NextRotationTime = dcl.FlattenString(m["nextRotationTime"])
	r.RotationPeriod = dcl.FlattenString(m["rotationPeriod"])

	return r
}

// expandWorkloadResourceSettingsMap expands the contents of WorkloadResourceSettings into a JSON
// request object.
func expandWorkloadResourceSettingsMap(c *Client, f map[string]WorkloadResourceSettings, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadResourceSettings(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadResourceSettingsSlice expands the contents of WorkloadResourceSettings into a JSON
// request object.
func expandWorkloadResourceSettingsSlice(c *Client, f []WorkloadResourceSettings, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadResourceSettings(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadResourceSettingsMap flattens the contents of WorkloadResourceSettings from a JSON
// response object.
func flattenWorkloadResourceSettingsMap(c *Client, i interface{}, res *Workload) map[string]WorkloadResourceSettings {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadResourceSettings{}
	}

	if len(a) == 0 {
		return map[string]WorkloadResourceSettings{}
	}

	items := make(map[string]WorkloadResourceSettings)
	for k, item := range a {
		items[k] = *flattenWorkloadResourceSettings(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadResourceSettingsSlice flattens the contents of WorkloadResourceSettings from a JSON
// response object.
func flattenWorkloadResourceSettingsSlice(c *Client, i interface{}, res *Workload) []WorkloadResourceSettings {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadResourceSettings{}
	}

	if len(a) == 0 {
		return []WorkloadResourceSettings{}
	}

	items := make([]WorkloadResourceSettings, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadResourceSettings(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadResourceSettings expands an instance of WorkloadResourceSettings into a JSON
// request object.
func expandWorkloadResourceSettings(c *Client, f *WorkloadResourceSettings, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceId; !dcl.IsEmptyValueIndirect(v) {
		m["resourceId"] = v
	}
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["displayName"] = v
	}

	return m, nil
}

// flattenWorkloadResourceSettings flattens an instance of WorkloadResourceSettings from a JSON
// response object.
func flattenWorkloadResourceSettings(c *Client, i interface{}, res *Workload) *WorkloadResourceSettings {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadResourceSettings{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadResourceSettings
	}
	r.ResourceId = dcl.FlattenString(m["resourceId"])
	r.ResourceType = flattenWorkloadResourceSettingsResourceTypeEnum(m["resourceType"])
	r.DisplayName = dcl.FlattenString(m["displayName"])

	return r
}

// expandWorkloadSaaEnrollmentResponseMap expands the contents of WorkloadSaaEnrollmentResponse into a JSON
// request object.
func expandWorkloadSaaEnrollmentResponseMap(c *Client, f map[string]WorkloadSaaEnrollmentResponse, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadSaaEnrollmentResponse(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadSaaEnrollmentResponseSlice expands the contents of WorkloadSaaEnrollmentResponse into a JSON
// request object.
func expandWorkloadSaaEnrollmentResponseSlice(c *Client, f []WorkloadSaaEnrollmentResponse, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadSaaEnrollmentResponse(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadSaaEnrollmentResponseMap flattens the contents of WorkloadSaaEnrollmentResponse from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseMap(c *Client, i interface{}, res *Workload) map[string]WorkloadSaaEnrollmentResponse {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadSaaEnrollmentResponse{}
	}

	if len(a) == 0 {
		return map[string]WorkloadSaaEnrollmentResponse{}
	}

	items := make(map[string]WorkloadSaaEnrollmentResponse)
	for k, item := range a {
		items[k] = *flattenWorkloadSaaEnrollmentResponse(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadSaaEnrollmentResponseSlice flattens the contents of WorkloadSaaEnrollmentResponse from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseSlice(c *Client, i interface{}, res *Workload) []WorkloadSaaEnrollmentResponse {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadSaaEnrollmentResponse{}
	}

	if len(a) == 0 {
		return []WorkloadSaaEnrollmentResponse{}
	}

	items := make([]WorkloadSaaEnrollmentResponse, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadSaaEnrollmentResponse(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadSaaEnrollmentResponse expands an instance of WorkloadSaaEnrollmentResponse into a JSON
// request object.
func expandWorkloadSaaEnrollmentResponse(c *Client, f *WorkloadSaaEnrollmentResponse, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SetupErrors; v != nil {
		m["setupErrors"] = v
	}
	if v := f.SetupStatus; !dcl.IsEmptyValueIndirect(v) {
		m["setupStatus"] = v
	}

	return m, nil
}

// flattenWorkloadSaaEnrollmentResponse flattens an instance of WorkloadSaaEnrollmentResponse from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponse(c *Client, i interface{}, res *Workload) *WorkloadSaaEnrollmentResponse {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadSaaEnrollmentResponse{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadSaaEnrollmentResponse
	}
	r.SetupErrors = flattenWorkloadSaaEnrollmentResponseSetupErrorsEnumSlice(c, m["setupErrors"], res)
	r.SetupStatus = flattenWorkloadSaaEnrollmentResponseSetupStatusEnum(m["setupStatus"])

	return r
}

// expandWorkloadComplianceStatusMap expands the contents of WorkloadComplianceStatus into a JSON
// request object.
func expandWorkloadComplianceStatusMap(c *Client, f map[string]WorkloadComplianceStatus, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadComplianceStatus(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadComplianceStatusSlice expands the contents of WorkloadComplianceStatus into a JSON
// request object.
func expandWorkloadComplianceStatusSlice(c *Client, f []WorkloadComplianceStatus, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadComplianceStatus(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadComplianceStatusMap flattens the contents of WorkloadComplianceStatus from a JSON
// response object.
func flattenWorkloadComplianceStatusMap(c *Client, i interface{}, res *Workload) map[string]WorkloadComplianceStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadComplianceStatus{}
	}

	if len(a) == 0 {
		return map[string]WorkloadComplianceStatus{}
	}

	items := make(map[string]WorkloadComplianceStatus)
	for k, item := range a {
		items[k] = *flattenWorkloadComplianceStatus(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadComplianceStatusSlice flattens the contents of WorkloadComplianceStatus from a JSON
// response object.
func flattenWorkloadComplianceStatusSlice(c *Client, i interface{}, res *Workload) []WorkloadComplianceStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadComplianceStatus{}
	}

	if len(a) == 0 {
		return []WorkloadComplianceStatus{}
	}

	items := make([]WorkloadComplianceStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadComplianceStatus(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadComplianceStatus expands an instance of WorkloadComplianceStatus into a JSON
// request object.
func expandWorkloadComplianceStatus(c *Client, f *WorkloadComplianceStatus, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ActiveViolationCount; v != nil {
		m["activeViolationCount"] = v
	}
	if v := f.AcknowledgedViolationCount; v != nil {
		m["acknowledgedViolationCount"] = v
	}

	return m, nil
}

// flattenWorkloadComplianceStatus flattens an instance of WorkloadComplianceStatus from a JSON
// response object.
func flattenWorkloadComplianceStatus(c *Client, i interface{}, res *Workload) *WorkloadComplianceStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadComplianceStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadComplianceStatus
	}
	r.ActiveViolationCount = dcl.FlattenIntSlice(m["activeViolationCount"])
	r.AcknowledgedViolationCount = dcl.FlattenIntSlice(m["acknowledgedViolationCount"])

	return r
}

// expandWorkloadPartnerPermissionsMap expands the contents of WorkloadPartnerPermissions into a JSON
// request object.
func expandWorkloadPartnerPermissionsMap(c *Client, f map[string]WorkloadPartnerPermissions, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadPartnerPermissions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadPartnerPermissionsSlice expands the contents of WorkloadPartnerPermissions into a JSON
// request object.
func expandWorkloadPartnerPermissionsSlice(c *Client, f []WorkloadPartnerPermissions, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadPartnerPermissions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadPartnerPermissionsMap flattens the contents of WorkloadPartnerPermissions from a JSON
// response object.
func flattenWorkloadPartnerPermissionsMap(c *Client, i interface{}, res *Workload) map[string]WorkloadPartnerPermissions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadPartnerPermissions{}
	}

	if len(a) == 0 {
		return map[string]WorkloadPartnerPermissions{}
	}

	items := make(map[string]WorkloadPartnerPermissions)
	for k, item := range a {
		items[k] = *flattenWorkloadPartnerPermissions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadPartnerPermissionsSlice flattens the contents of WorkloadPartnerPermissions from a JSON
// response object.
func flattenWorkloadPartnerPermissionsSlice(c *Client, i interface{}, res *Workload) []WorkloadPartnerPermissions {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadPartnerPermissions{}
	}

	if len(a) == 0 {
		return []WorkloadPartnerPermissions{}
	}

	items := make([]WorkloadPartnerPermissions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadPartnerPermissions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadPartnerPermissions expands an instance of WorkloadPartnerPermissions into a JSON
// request object.
func expandWorkloadPartnerPermissions(c *Client, f *WorkloadPartnerPermissions, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DataLogsViewer; !dcl.IsEmptyValueIndirect(v) {
		m["dataLogsViewer"] = v
	}
	if v := f.ServiceAccessApprover; !dcl.IsEmptyValueIndirect(v) {
		m["serviceAccessApprover"] = v
	}
	if v := f.AssuredWorkloadsMonitoring; !dcl.IsEmptyValueIndirect(v) {
		m["assuredWorkloadsMonitoring"] = v
	}

	return m, nil
}

// flattenWorkloadPartnerPermissions flattens an instance of WorkloadPartnerPermissions from a JSON
// response object.
func flattenWorkloadPartnerPermissions(c *Client, i interface{}, res *Workload) *WorkloadPartnerPermissions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadPartnerPermissions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadPartnerPermissions
	}
	r.DataLogsViewer = dcl.FlattenBool(m["dataLogsViewer"])
	r.ServiceAccessApprover = dcl.FlattenBool(m["serviceAccessApprover"])
	r.AssuredWorkloadsMonitoring = dcl.FlattenBool(m["assuredWorkloadsMonitoring"])

	return r
}

// expandWorkloadEkmProvisioningResponseMap expands the contents of WorkloadEkmProvisioningResponse into a JSON
// request object.
func expandWorkloadEkmProvisioningResponseMap(c *Client, f map[string]WorkloadEkmProvisioningResponse, res *Workload) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandWorkloadEkmProvisioningResponse(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandWorkloadEkmProvisioningResponseSlice expands the contents of WorkloadEkmProvisioningResponse into a JSON
// request object.
func expandWorkloadEkmProvisioningResponseSlice(c *Client, f []WorkloadEkmProvisioningResponse, res *Workload) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandWorkloadEkmProvisioningResponse(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenWorkloadEkmProvisioningResponseMap flattens the contents of WorkloadEkmProvisioningResponse from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseMap(c *Client, i interface{}, res *Workload) map[string]WorkloadEkmProvisioningResponse {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadEkmProvisioningResponse{}
	}

	if len(a) == 0 {
		return map[string]WorkloadEkmProvisioningResponse{}
	}

	items := make(map[string]WorkloadEkmProvisioningResponse)
	for k, item := range a {
		items[k] = *flattenWorkloadEkmProvisioningResponse(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseSlice flattens the contents of WorkloadEkmProvisioningResponse from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseSlice(c *Client, i interface{}, res *Workload) []WorkloadEkmProvisioningResponse {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadEkmProvisioningResponse{}
	}

	if len(a) == 0 {
		return []WorkloadEkmProvisioningResponse{}
	}

	items := make([]WorkloadEkmProvisioningResponse, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadEkmProvisioningResponse(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandWorkloadEkmProvisioningResponse expands an instance of WorkloadEkmProvisioningResponse into a JSON
// request object.
func expandWorkloadEkmProvisioningResponse(c *Client, f *WorkloadEkmProvisioningResponse, res *Workload) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EkmProvisioningState; !dcl.IsEmptyValueIndirect(v) {
		m["ekmProvisioningState"] = v
	}
	if v := f.EkmProvisioningErrorDomain; !dcl.IsEmptyValueIndirect(v) {
		m["ekmProvisioningErrorDomain"] = v
	}
	if v := f.EkmProvisioningErrorMapping; !dcl.IsEmptyValueIndirect(v) {
		m["ekmProvisioningErrorMapping"] = v
	}

	return m, nil
}

// flattenWorkloadEkmProvisioningResponse flattens an instance of WorkloadEkmProvisioningResponse from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponse(c *Client, i interface{}, res *Workload) *WorkloadEkmProvisioningResponse {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &WorkloadEkmProvisioningResponse{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyWorkloadEkmProvisioningResponse
	}
	r.EkmProvisioningState = flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(m["ekmProvisioningState"])
	r.EkmProvisioningErrorDomain = flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(m["ekmProvisioningErrorDomain"])
	r.EkmProvisioningErrorMapping = flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(m["ekmProvisioningErrorMapping"])

	return r
}

// flattenWorkloadResourcesResourceTypeEnumMap flattens the contents of WorkloadResourcesResourceTypeEnum from a JSON
// response object.
func flattenWorkloadResourcesResourceTypeEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadResourcesResourceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadResourcesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadResourcesResourceTypeEnum{}
	}

	items := make(map[string]WorkloadResourcesResourceTypeEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadResourcesResourceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadResourcesResourceTypeEnumSlice flattens the contents of WorkloadResourcesResourceTypeEnum from a JSON
// response object.
func flattenWorkloadResourcesResourceTypeEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadResourcesResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadResourcesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []WorkloadResourcesResourceTypeEnum{}
	}

	items := make([]WorkloadResourcesResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadResourcesResourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadResourcesResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadResourcesResourceTypeEnum with the same value as that string.
func flattenWorkloadResourcesResourceTypeEnum(i interface{}) *WorkloadResourcesResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadResourcesResourceTypeEnumRef(s)
}

// flattenWorkloadComplianceRegimeEnumMap flattens the contents of WorkloadComplianceRegimeEnum from a JSON
// response object.
func flattenWorkloadComplianceRegimeEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadComplianceRegimeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadComplianceRegimeEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadComplianceRegimeEnum{}
	}

	items := make(map[string]WorkloadComplianceRegimeEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadComplianceRegimeEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadComplianceRegimeEnumSlice flattens the contents of WorkloadComplianceRegimeEnum from a JSON
// response object.
func flattenWorkloadComplianceRegimeEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadComplianceRegimeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadComplianceRegimeEnum{}
	}

	if len(a) == 0 {
		return []WorkloadComplianceRegimeEnum{}
	}

	items := make([]WorkloadComplianceRegimeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadComplianceRegimeEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadComplianceRegimeEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadComplianceRegimeEnum with the same value as that string.
func flattenWorkloadComplianceRegimeEnum(i interface{}) *WorkloadComplianceRegimeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadComplianceRegimeEnumRef(s)
}

// flattenWorkloadResourceSettingsResourceTypeEnumMap flattens the contents of WorkloadResourceSettingsResourceTypeEnum from a JSON
// response object.
func flattenWorkloadResourceSettingsResourceTypeEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadResourceSettingsResourceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadResourceSettingsResourceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadResourceSettingsResourceTypeEnum{}
	}

	items := make(map[string]WorkloadResourceSettingsResourceTypeEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadResourceSettingsResourceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadResourceSettingsResourceTypeEnumSlice flattens the contents of WorkloadResourceSettingsResourceTypeEnum from a JSON
// response object.
func flattenWorkloadResourceSettingsResourceTypeEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadResourceSettingsResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadResourceSettingsResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []WorkloadResourceSettingsResourceTypeEnum{}
	}

	items := make([]WorkloadResourceSettingsResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadResourceSettingsResourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadResourceSettingsResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadResourceSettingsResourceTypeEnum with the same value as that string.
func flattenWorkloadResourceSettingsResourceTypeEnum(i interface{}) *WorkloadResourceSettingsResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadResourceSettingsResourceTypeEnumRef(s)
}

// flattenWorkloadKajEnrollmentStateEnumMap flattens the contents of WorkloadKajEnrollmentStateEnum from a JSON
// response object.
func flattenWorkloadKajEnrollmentStateEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadKajEnrollmentStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadKajEnrollmentStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadKajEnrollmentStateEnum{}
	}

	items := make(map[string]WorkloadKajEnrollmentStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadKajEnrollmentStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadKajEnrollmentStateEnumSlice flattens the contents of WorkloadKajEnrollmentStateEnum from a JSON
// response object.
func flattenWorkloadKajEnrollmentStateEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadKajEnrollmentStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadKajEnrollmentStateEnum{}
	}

	if len(a) == 0 {
		return []WorkloadKajEnrollmentStateEnum{}
	}

	items := make([]WorkloadKajEnrollmentStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadKajEnrollmentStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadKajEnrollmentStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadKajEnrollmentStateEnum with the same value as that string.
func flattenWorkloadKajEnrollmentStateEnum(i interface{}) *WorkloadKajEnrollmentStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadKajEnrollmentStateEnumRef(s)
}

// flattenWorkloadSaaEnrollmentResponseSetupErrorsEnumMap flattens the contents of WorkloadSaaEnrollmentResponseSetupErrorsEnum from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseSetupErrorsEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadSaaEnrollmentResponseSetupErrorsEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadSaaEnrollmentResponseSetupErrorsEnum{}
	}

	items := make(map[string]WorkloadSaaEnrollmentResponseSetupErrorsEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadSaaEnrollmentResponseSetupErrorsEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadSaaEnrollmentResponseSetupErrorsEnumSlice flattens the contents of WorkloadSaaEnrollmentResponseSetupErrorsEnum from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseSetupErrorsEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadSaaEnrollmentResponseSetupErrorsEnum{}
	}

	if len(a) == 0 {
		return []WorkloadSaaEnrollmentResponseSetupErrorsEnum{}
	}

	items := make([]WorkloadSaaEnrollmentResponseSetupErrorsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadSaaEnrollmentResponseSetupErrorsEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadSaaEnrollmentResponseSetupErrorsEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadSaaEnrollmentResponseSetupErrorsEnum with the same value as that string.
func flattenWorkloadSaaEnrollmentResponseSetupErrorsEnum(i interface{}) *WorkloadSaaEnrollmentResponseSetupErrorsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadSaaEnrollmentResponseSetupErrorsEnumRef(s)
}

// flattenWorkloadSaaEnrollmentResponseSetupStatusEnumMap flattens the contents of WorkloadSaaEnrollmentResponseSetupStatusEnum from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseSetupStatusEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadSaaEnrollmentResponseSetupStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadSaaEnrollmentResponseSetupStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadSaaEnrollmentResponseSetupStatusEnum{}
	}

	items := make(map[string]WorkloadSaaEnrollmentResponseSetupStatusEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadSaaEnrollmentResponseSetupStatusEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadSaaEnrollmentResponseSetupStatusEnumSlice flattens the contents of WorkloadSaaEnrollmentResponseSetupStatusEnum from a JSON
// response object.
func flattenWorkloadSaaEnrollmentResponseSetupStatusEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadSaaEnrollmentResponseSetupStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadSaaEnrollmentResponseSetupStatusEnum{}
	}

	if len(a) == 0 {
		return []WorkloadSaaEnrollmentResponseSetupStatusEnum{}
	}

	items := make([]WorkloadSaaEnrollmentResponseSetupStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadSaaEnrollmentResponseSetupStatusEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadSaaEnrollmentResponseSetupStatusEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadSaaEnrollmentResponseSetupStatusEnum with the same value as that string.
func flattenWorkloadSaaEnrollmentResponseSetupStatusEnum(i interface{}) *WorkloadSaaEnrollmentResponseSetupStatusEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadSaaEnrollmentResponseSetupStatusEnumRef(s)
}

// flattenWorkloadPartnerEnumMap flattens the contents of WorkloadPartnerEnum from a JSON
// response object.
func flattenWorkloadPartnerEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadPartnerEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadPartnerEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadPartnerEnum{}
	}

	items := make(map[string]WorkloadPartnerEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadPartnerEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadPartnerEnumSlice flattens the contents of WorkloadPartnerEnum from a JSON
// response object.
func flattenWorkloadPartnerEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadPartnerEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadPartnerEnum{}
	}

	if len(a) == 0 {
		return []WorkloadPartnerEnum{}
	}

	items := make([]WorkloadPartnerEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadPartnerEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadPartnerEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadPartnerEnum with the same value as that string.
func flattenWorkloadPartnerEnum(i interface{}) *WorkloadPartnerEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadPartnerEnumRef(s)
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnumMap flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningStateEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningStateEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningStateEnum{}
	}

	items := make(map[string]WorkloadEkmProvisioningResponseEkmProvisioningStateEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnumSlice flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningStateEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadEkmProvisioningResponseEkmProvisioningStateEnum{}
	}

	if len(a) == 0 {
		return []WorkloadEkmProvisioningResponseEkmProvisioningStateEnum{}
	}

	items := make([]WorkloadEkmProvisioningResponseEkmProvisioningStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadEkmProvisioningResponseEkmProvisioningStateEnum with the same value as that string.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningStateEnum(i interface{}) *WorkloadEkmProvisioningResponseEkmProvisioningStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadEkmProvisioningResponseEkmProvisioningStateEnumRef(s)
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumMap flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum{}
	}

	items := make(map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumSlice flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum{}
	}

	if len(a) == 0 {
		return []WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum{}
	}

	items := make([]WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum with the same value as that string.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum(i interface{}) *WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadEkmProvisioningResponseEkmProvisioningErrorDomainEnumRef(s)
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumMap flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumMap(c *Client, i interface{}, res *Workload) map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum{}
	}

	if len(a) == 0 {
		return map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum{}
	}

	items := make(map[string]WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum)
	for k, item := range a {
		items[k] = *flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(item.(interface{}))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumSlice flattens the contents of WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum from a JSON
// response object.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumSlice(c *Client, i interface{}, res *Workload) []WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum{}
	}

	if len(a) == 0 {
		return []WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum{}
	}

	items := make([]WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(item.(interface{})))
	}

	return items
}

// flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum asserts that an interface is a string, and returns a
// pointer to a *WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum with the same value as that string.
func flattenWorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum(i interface{}) *WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return WorkloadEkmProvisioningResponseEkmProvisioningErrorMappingEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Workload) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWorkload(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Organization == nil && ncr.Organization == nil {
			c.Config.Logger.Info("Both Organization fields null - considering equal.")
		} else if nr.Organization == nil || ncr.Organization == nil {
			c.Config.Logger.Info("Only one Organization field is null - considering unequal.")
			return false
		} else if *nr.Organization != *ncr.Organization {
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

type workloadDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         workloadApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToWorkloadDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]workloadDiff, error) {
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
	var diffs []workloadDiff
	// For each operation name, create a workloadDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := workloadDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWorkloadApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWorkloadApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (workloadApiOperation, error) {
	switch opName {

	case "updateWorkloadUpdateWorkloadOperation":
		return &updateWorkloadUpdateWorkloadOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWorkloadFields(r *Workload) error {
	vKmsSettings := r.KmsSettings
	if vKmsSettings == nil {
		// note: explicitly not the empty object.
		vKmsSettings = &WorkloadKmsSettings{}
	}
	if err := extractWorkloadKmsSettingsFields(r, vKmsSettings); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKmsSettings) {
		r.KmsSettings = vKmsSettings
	}
	vSaaEnrollmentResponse := r.SaaEnrollmentResponse
	if vSaaEnrollmentResponse == nil {
		// note: explicitly not the empty object.
		vSaaEnrollmentResponse = &WorkloadSaaEnrollmentResponse{}
	}
	if err := extractWorkloadSaaEnrollmentResponseFields(r, vSaaEnrollmentResponse); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaaEnrollmentResponse) {
		r.SaaEnrollmentResponse = vSaaEnrollmentResponse
	}
	vComplianceStatus := r.ComplianceStatus
	if vComplianceStatus == nil {
		// note: explicitly not the empty object.
		vComplianceStatus = &WorkloadComplianceStatus{}
	}
	if err := extractWorkloadComplianceStatusFields(r, vComplianceStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vComplianceStatus) {
		r.ComplianceStatus = vComplianceStatus
	}
	vPartnerPermissions := r.PartnerPermissions
	if vPartnerPermissions == nil {
		// note: explicitly not the empty object.
		vPartnerPermissions = &WorkloadPartnerPermissions{}
	}
	if err := extractWorkloadPartnerPermissionsFields(r, vPartnerPermissions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPartnerPermissions) {
		r.PartnerPermissions = vPartnerPermissions
	}
	vEkmProvisioningResponse := r.EkmProvisioningResponse
	if vEkmProvisioningResponse == nil {
		// note: explicitly not the empty object.
		vEkmProvisioningResponse = &WorkloadEkmProvisioningResponse{}
	}
	if err := extractWorkloadEkmProvisioningResponseFields(r, vEkmProvisioningResponse); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEkmProvisioningResponse) {
		r.EkmProvisioningResponse = vEkmProvisioningResponse
	}
	return nil
}
func extractWorkloadResourcesFields(r *Workload, o *WorkloadResources) error {
	return nil
}
func extractWorkloadKmsSettingsFields(r *Workload, o *WorkloadKmsSettings) error {
	return nil
}
func extractWorkloadResourceSettingsFields(r *Workload, o *WorkloadResourceSettings) error {
	return nil
}
func extractWorkloadSaaEnrollmentResponseFields(r *Workload, o *WorkloadSaaEnrollmentResponse) error {
	return nil
}
func extractWorkloadComplianceStatusFields(r *Workload, o *WorkloadComplianceStatus) error {
	return nil
}
func extractWorkloadPartnerPermissionsFields(r *Workload, o *WorkloadPartnerPermissions) error {
	return nil
}
func extractWorkloadEkmProvisioningResponseFields(r *Workload, o *WorkloadEkmProvisioningResponse) error {
	return nil
}

func postReadExtractWorkloadFields(r *Workload) error {
	vKmsSettings := r.KmsSettings
	if vKmsSettings == nil {
		// note: explicitly not the empty object.
		vKmsSettings = &WorkloadKmsSettings{}
	}
	if err := postReadExtractWorkloadKmsSettingsFields(r, vKmsSettings); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vKmsSettings) {
		r.KmsSettings = vKmsSettings
	}
	vSaaEnrollmentResponse := r.SaaEnrollmentResponse
	if vSaaEnrollmentResponse == nil {
		// note: explicitly not the empty object.
		vSaaEnrollmentResponse = &WorkloadSaaEnrollmentResponse{}
	}
	if err := postReadExtractWorkloadSaaEnrollmentResponseFields(r, vSaaEnrollmentResponse); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSaaEnrollmentResponse) {
		r.SaaEnrollmentResponse = vSaaEnrollmentResponse
	}
	vComplianceStatus := r.ComplianceStatus
	if vComplianceStatus == nil {
		// note: explicitly not the empty object.
		vComplianceStatus = &WorkloadComplianceStatus{}
	}
	if err := postReadExtractWorkloadComplianceStatusFields(r, vComplianceStatus); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vComplianceStatus) {
		r.ComplianceStatus = vComplianceStatus
	}
	vPartnerPermissions := r.PartnerPermissions
	if vPartnerPermissions == nil {
		// note: explicitly not the empty object.
		vPartnerPermissions = &WorkloadPartnerPermissions{}
	}
	if err := postReadExtractWorkloadPartnerPermissionsFields(r, vPartnerPermissions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPartnerPermissions) {
		r.PartnerPermissions = vPartnerPermissions
	}
	vEkmProvisioningResponse := r.EkmProvisioningResponse
	if vEkmProvisioningResponse == nil {
		// note: explicitly not the empty object.
		vEkmProvisioningResponse = &WorkloadEkmProvisioningResponse{}
	}
	if err := postReadExtractWorkloadEkmProvisioningResponseFields(r, vEkmProvisioningResponse); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEkmProvisioningResponse) {
		r.EkmProvisioningResponse = vEkmProvisioningResponse
	}
	return nil
}
func postReadExtractWorkloadResourcesFields(r *Workload, o *WorkloadResources) error {
	return nil
}
func postReadExtractWorkloadKmsSettingsFields(r *Workload, o *WorkloadKmsSettings) error {
	return nil
}
func postReadExtractWorkloadResourceSettingsFields(r *Workload, o *WorkloadResourceSettings) error {
	return nil
}
func postReadExtractWorkloadSaaEnrollmentResponseFields(r *Workload, o *WorkloadSaaEnrollmentResponse) error {
	return nil
}
func postReadExtractWorkloadComplianceStatusFields(r *Workload, o *WorkloadComplianceStatus) error {
	return nil
}
func postReadExtractWorkloadPartnerPermissionsFields(r *Workload, o *WorkloadPartnerPermissions) error {
	return nil
}
func postReadExtractWorkloadEkmProvisioningResponseFields(r *Workload, o *WorkloadEkmProvisioningResponse) error {
	return nil
}
