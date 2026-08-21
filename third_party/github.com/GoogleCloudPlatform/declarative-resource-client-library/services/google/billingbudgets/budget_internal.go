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
package billingbudgets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Budget) validate() error {

	if err := dcl.Required(r, "amount"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.BillingAccount, "BillingAccount"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.BudgetFilter) {
		if err := r.BudgetFilter.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Amount) {
		if err := r.Amount.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.AllUpdatesRule) {
		if err := r.AllUpdatesRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BudgetBudgetFilter) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"CalendarPeriod", "CustomPeriod"}, r.CalendarPeriod, r.CustomPeriod); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.CustomPeriod) {
		if err := r.CustomPeriod.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BudgetBudgetFilterLabels) validate() error {
	return nil
}
func (r *BudgetBudgetFilterCustomPeriod) validate() error {
	if err := dcl.Required(r, "startDate"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.StartDate) {
		if err := r.StartDate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.EndDate) {
		if err := r.EndDate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BudgetBudgetFilterCustomPeriodStartDate) validate() error {
	return nil
}
func (r *BudgetBudgetFilterCustomPeriodEndDate) validate() error {
	return nil
}
func (r *BudgetAmount) validate() error {
	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"SpecifiedAmount", "LastPeriodAmount"}, r.SpecifiedAmount, r.LastPeriodAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SpecifiedAmount) {
		if err := r.SpecifiedAmount.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LastPeriodAmount) {
		if err := r.LastPeriodAmount.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BudgetAmountSpecifiedAmount) validate() error {
	return nil
}
func (r *BudgetAmountLastPeriodAmount) validate() error {
	return nil
}
func (r *BudgetThresholdRules) validate() error {
	if err := dcl.Required(r, "thresholdPercent"); err != nil {
		return err
	}
	return nil
}
func (r *BudgetAllUpdatesRule) validate() error {
	return nil
}
func (r *Budget) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://billingbudgets.googleapis.com/v1/", params)
}

func (r *Budget) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"billingAccount": dcl.ValueOrEmptyString(nr.BillingAccount),
		"name":           dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("billingAccounts/{{billingAccount}}/budgets/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Budget) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"billingAccount": dcl.ValueOrEmptyString(nr.BillingAccount),
	}
	return dcl.URL("billingAccounts/{{billingAccount}}/budgets", nr.basePath(), userBasePath, params), nil

}

func (r *Budget) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"billingAccount": dcl.ValueOrEmptyString(nr.BillingAccount),
	}
	return dcl.URL("billingAccounts/{{billingAccount}}/budgets", nr.basePath(), userBasePath, params), nil

}

func (r *Budget) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"billingAccount": dcl.ValueOrEmptyString(nr.BillingAccount),
		"name":           dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("billingAccounts/{{billingAccount}}/budgets/{{name}}", nr.basePath(), userBasePath, params), nil
}

// budgetApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type budgetApiOperation interface {
	do(context.Context, *Budget, *Client) error
}

// newUpdateBudgetUpdateBudgetRequest creates a request for an
// Budget resource's UpdateBudget update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBudgetUpdateBudgetRequest(ctx context.Context, f *Budget, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v, err := expandBudgetBudgetFilter(c, f.BudgetFilter, res); err != nil {
		return nil, fmt.Errorf("error expanding BudgetFilter into budgetFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["budgetFilter"] = v
	}
	if v, err := expandBudgetAmount(c, f.Amount, res); err != nil {
		return nil, fmt.Errorf("error expanding Amount into amount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["amount"] = v
	}
	if v, err := expandBudgetThresholdRulesSlice(c, f.ThresholdRules, res); err != nil {
		return nil, fmt.Errorf("error expanding ThresholdRules into thresholdRules: %w", err)
	} else if v != nil {
		req["thresholdRules"] = v
	}
	if v, err := expandBudgetAllUpdatesRule(c, f.AllUpdatesRule, res); err != nil {
		return nil, fmt.Errorf("error expanding AllUpdatesRule into notificationsRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["notificationsRule"] = v
	}
	b, err := c.getBudgetRaw(ctx, f)
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

// marshalUpdateBudgetUpdateBudgetRequest converts the update into
// the final JSON request body.
func marshalUpdateBudgetUpdateBudgetRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBudgetUpdateBudgetOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listBudgetRaw(ctx context.Context, r *Budget, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BudgetMaxPage {
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

type listBudgetOperation struct {
	Budgets []map[string]interface{} `json:"budgets"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listBudget(ctx context.Context, r *Budget, pageToken string, pageSize int32) ([]*Budget, string, error) {
	b, err := c.listBudgetRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBudgetOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Budget
	for _, v := range m.Budgets {
		res, err := unmarshalMapBudget(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.BillingAccount = r.BillingAccount
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBudget(ctx context.Context, f func(*Budget) bool, resources []*Budget) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBudget(ctx, res)
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

type deleteBudgetOperation struct{}

func (op *deleteBudgetOperation) do(ctx context.Context, r *Budget, c *Client) error {
	r, err := c.GetBudget(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Budget not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetBudget checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Budget: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetBudget(ctx, r)
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
type createBudgetOperation struct {
	response map[string]interface{}
}

func (op *createBudgetOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getBudgetRaw(ctx context.Context, r *Budget) ([]byte, error) {

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

func (c *Client) budgetDiffsForRawDesired(ctx context.Context, rawDesired *Budget, opts ...dcl.ApplyOption) (initial, desired *Budget, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Budget
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Budget); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Budget, got %T", sh)
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
		desired, err := canonicalizeBudgetDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBudget(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Budget resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Budget resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Budget resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBudgetDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Budget: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Budget: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractBudgetFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBudgetInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Budget: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBudgetDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Budget: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBudget(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBudgetInitialState(rawInitial, rawDesired *Budget) (*Budget, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBudgetDesiredState(rawDesired, rawInitial *Budget, opts ...dcl.ApplyOption) (*Budget, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.BudgetFilter = canonicalizeBudgetBudgetFilter(rawDesired.BudgetFilter, nil, opts...)
		rawDesired.Amount = canonicalizeBudgetAmount(rawDesired.Amount, nil, opts...)
		rawDesired.AllUpdatesRule = canonicalizeBudgetAllUpdatesRule(rawDesired.AllUpdatesRule, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Budget{}
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
	canonicalDesired.BudgetFilter = canonicalizeBudgetBudgetFilter(rawDesired.BudgetFilter, rawInitial.BudgetFilter, opts...)
	canonicalDesired.Amount = canonicalizeBudgetAmount(rawDesired.Amount, rawInitial.Amount, opts...)
	canonicalDesired.ThresholdRules = canonicalizeBudgetThresholdRulesSlice(rawDesired.ThresholdRules, rawInitial.ThresholdRules, opts...)
	canonicalDesired.AllUpdatesRule = canonicalizeBudgetAllUpdatesRule(rawDesired.AllUpdatesRule, rawInitial.AllUpdatesRule, opts...)
	if dcl.NameToSelfLink(rawDesired.BillingAccount, rawInitial.BillingAccount) {
		canonicalDesired.BillingAccount = rawInitial.BillingAccount
	} else {
		canonicalDesired.BillingAccount = rawDesired.BillingAccount
	}
	return canonicalDesired, nil
}

func canonicalizeBudgetNewState(c *Client, rawNew, rawDesired *Budget) (*Budget, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.BudgetFilter) && dcl.IsEmptyValueIndirect(rawDesired.BudgetFilter) {
		rawNew.BudgetFilter = rawDesired.BudgetFilter
	} else {
		rawNew.BudgetFilter = canonicalizeNewBudgetBudgetFilter(c, rawDesired.BudgetFilter, rawNew.BudgetFilter)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Amount) && dcl.IsEmptyValueIndirect(rawDesired.Amount) {
		rawNew.Amount = rawDesired.Amount
	} else {
		rawNew.Amount = canonicalizeNewBudgetAmount(c, rawDesired.Amount, rawNew.Amount)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ThresholdRules) && dcl.IsEmptyValueIndirect(rawDesired.ThresholdRules) {
		rawNew.ThresholdRules = rawDesired.ThresholdRules
	} else {
		rawNew.ThresholdRules = canonicalizeNewBudgetThresholdRulesSlice(c, rawDesired.ThresholdRules, rawNew.ThresholdRules)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AllUpdatesRule) && dcl.IsEmptyValueIndirect(rawDesired.AllUpdatesRule) {
		rawNew.AllUpdatesRule = rawDesired.AllUpdatesRule
	} else {
		rawNew.AllUpdatesRule = canonicalizeNewBudgetAllUpdatesRule(c, rawDesired.AllUpdatesRule, rawNew.AllUpdatesRule)
	}

	rawNew.BillingAccount = rawDesired.BillingAccount

	return rawNew, nil
}

func canonicalizeBudgetBudgetFilter(des, initial *BudgetBudgetFilter, opts ...dcl.ApplyOption) *BudgetBudgetFilter {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.CalendarPeriod != nil || (initial != nil && initial.CalendarPeriod != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CustomPeriod) {
			des.CalendarPeriod = nil
			if initial != nil {
				initial.CalendarPeriod = nil
			}
		}
	}

	if des.CustomPeriod != nil || (initial != nil && initial.CustomPeriod != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CalendarPeriod) {
			des.CustomPeriod = nil
			if initial != nil {
				initial.CustomPeriod = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetBudgetFilter{}

	if dcl.StringArrayCanonicalize(des.Projects, initial.Projects) {
		cDes.Projects = initial.Projects
	} else {
		cDes.Projects = des.Projects
	}
	if dcl.StringArrayCanonicalize(des.CreditTypes, initial.CreditTypes) {
		cDes.CreditTypes = initial.CreditTypes
	} else {
		cDes.CreditTypes = des.CreditTypes
	}
	if dcl.IsZeroValue(des.CreditTypesTreatment) || (dcl.IsEmptyValueIndirect(des.CreditTypesTreatment) && dcl.IsEmptyValueIndirect(initial.CreditTypesTreatment)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CreditTypesTreatment = initial.CreditTypesTreatment
	} else {
		cDes.CreditTypesTreatment = des.CreditTypesTreatment
	}
	if dcl.StringArrayCanonicalize(des.Services, initial.Services) {
		cDes.Services = initial.Services
	} else {
		cDes.Services = des.Services
	}
	if dcl.StringArrayCanonicalize(des.Subaccounts, initial.Subaccounts) {
		cDes.Subaccounts = initial.Subaccounts
	} else {
		cDes.Subaccounts = des.Subaccounts
	}
	if dcl.IsZeroValue(des.Labels) || (dcl.IsEmptyValueIndirect(des.Labels) && dcl.IsEmptyValueIndirect(initial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Labels = initial.Labels
	} else {
		cDes.Labels = des.Labels
	}
	if dcl.IsZeroValue(des.CalendarPeriod) || (dcl.IsEmptyValueIndirect(des.CalendarPeriod) && dcl.IsEmptyValueIndirect(initial.CalendarPeriod)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CalendarPeriod = initial.CalendarPeriod
	} else {
		cDes.CalendarPeriod = des.CalendarPeriod
	}
	cDes.CustomPeriod = canonicalizeBudgetBudgetFilterCustomPeriod(des.CustomPeriod, initial.CustomPeriod, opts...)

	return cDes
}

func canonicalizeBudgetBudgetFilterSlice(des, initial []BudgetBudgetFilter, opts ...dcl.ApplyOption) []BudgetBudgetFilter {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetBudgetFilter, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetBudgetFilter(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetBudgetFilter, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetBudgetFilter(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetBudgetFilter(c *Client, des, nw *BudgetBudgetFilter) *BudgetBudgetFilter {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetBudgetFilter while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Projects, nw.Projects) {
		nw.Projects = des.Projects
	}
	if dcl.StringArrayCanonicalize(des.CreditTypes, nw.CreditTypes) {
		nw.CreditTypes = des.CreditTypes
	}
	if dcl.StringArrayCanonicalize(des.Services, nw.Services) {
		nw.Services = des.Services
	}
	if dcl.StringArrayCanonicalize(des.Subaccounts, nw.Subaccounts) {
		nw.Subaccounts = des.Subaccounts
	}
	nw.CustomPeriod = canonicalizeNewBudgetBudgetFilterCustomPeriod(c, des.CustomPeriod, nw.CustomPeriod)

	return nw
}

func canonicalizeNewBudgetBudgetFilterSet(c *Client, des, nw []BudgetBudgetFilter) []BudgetBudgetFilter {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetBudgetFilter
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetBudgetFilterNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetBudgetFilter(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetBudgetFilterSlice(c *Client, des, nw []BudgetBudgetFilter) []BudgetBudgetFilter {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetBudgetFilter
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetBudgetFilter(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetBudgetFilterLabels(des, initial *BudgetBudgetFilterLabels, opts ...dcl.ApplyOption) *BudgetBudgetFilterLabels {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetBudgetFilterLabels{}

	if dcl.StringArrayCanonicalize(des.Values, initial.Values) {
		cDes.Values = initial.Values
	} else {
		cDes.Values = des.Values
	}

	return cDes
}

func canonicalizeBudgetBudgetFilterLabelsSlice(des, initial []BudgetBudgetFilterLabels, opts ...dcl.ApplyOption) []BudgetBudgetFilterLabels {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetBudgetFilterLabels, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetBudgetFilterLabels(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetBudgetFilterLabels, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetBudgetFilterLabels(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetBudgetFilterLabels(c *Client, des, nw *BudgetBudgetFilterLabels) *BudgetBudgetFilterLabels {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetBudgetFilterLabels while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Values, nw.Values) {
		nw.Values = des.Values
	}

	return nw
}

func canonicalizeNewBudgetBudgetFilterLabelsSet(c *Client, des, nw []BudgetBudgetFilterLabels) []BudgetBudgetFilterLabels {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetBudgetFilterLabels
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetBudgetFilterLabelsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetBudgetFilterLabels(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetBudgetFilterLabelsSlice(c *Client, des, nw []BudgetBudgetFilterLabels) []BudgetBudgetFilterLabels {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetBudgetFilterLabels
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetBudgetFilterLabels(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetBudgetFilterCustomPeriod(des, initial *BudgetBudgetFilterCustomPeriod, opts ...dcl.ApplyOption) *BudgetBudgetFilterCustomPeriod {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetBudgetFilterCustomPeriod{}

	cDes.StartDate = canonicalizeBudgetBudgetFilterCustomPeriodStartDate(des.StartDate, initial.StartDate, opts...)
	cDes.EndDate = canonicalizeBudgetBudgetFilterCustomPeriodEndDate(des.EndDate, initial.EndDate, opts...)

	return cDes
}

func canonicalizeBudgetBudgetFilterCustomPeriodSlice(des, initial []BudgetBudgetFilterCustomPeriod, opts ...dcl.ApplyOption) []BudgetBudgetFilterCustomPeriod {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetBudgetFilterCustomPeriod, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetBudgetFilterCustomPeriod(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetBudgetFilterCustomPeriod, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetBudgetFilterCustomPeriod(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetBudgetFilterCustomPeriod(c *Client, des, nw *BudgetBudgetFilterCustomPeriod) *BudgetBudgetFilterCustomPeriod {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetBudgetFilterCustomPeriod while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.StartDate = canonicalizeNewBudgetBudgetFilterCustomPeriodStartDate(c, des.StartDate, nw.StartDate)
	nw.EndDate = canonicalizeNewBudgetBudgetFilterCustomPeriodEndDate(c, des.EndDate, nw.EndDate)

	return nw
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodSet(c *Client, des, nw []BudgetBudgetFilterCustomPeriod) []BudgetBudgetFilterCustomPeriod {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetBudgetFilterCustomPeriod
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetBudgetFilterCustomPeriodNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriod(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodSlice(c *Client, des, nw []BudgetBudgetFilterCustomPeriod) []BudgetBudgetFilterCustomPeriod {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetBudgetFilterCustomPeriod
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriod(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetBudgetFilterCustomPeriodStartDate(des, initial *BudgetBudgetFilterCustomPeriodStartDate, opts ...dcl.ApplyOption) *BudgetBudgetFilterCustomPeriodStartDate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetBudgetFilterCustomPeriodStartDate{}

	if dcl.IsZeroValue(des.Year) || (dcl.IsEmptyValueIndirect(des.Year) && dcl.IsEmptyValueIndirect(initial.Year)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Year = initial.Year
	} else {
		cDes.Year = des.Year
	}
	if dcl.IsZeroValue(des.Month) || (dcl.IsEmptyValueIndirect(des.Month) && dcl.IsEmptyValueIndirect(initial.Month)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Month = initial.Month
	} else {
		cDes.Month = des.Month
	}
	if dcl.IsZeroValue(des.Day) || (dcl.IsEmptyValueIndirect(des.Day) && dcl.IsEmptyValueIndirect(initial.Day)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Day = initial.Day
	} else {
		cDes.Day = des.Day
	}

	return cDes
}

func canonicalizeBudgetBudgetFilterCustomPeriodStartDateSlice(des, initial []BudgetBudgetFilterCustomPeriodStartDate, opts ...dcl.ApplyOption) []BudgetBudgetFilterCustomPeriodStartDate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetBudgetFilterCustomPeriodStartDate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetBudgetFilterCustomPeriodStartDate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetBudgetFilterCustomPeriodStartDate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetBudgetFilterCustomPeriodStartDate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetBudgetFilterCustomPeriodStartDate(c *Client, des, nw *BudgetBudgetFilterCustomPeriodStartDate) *BudgetBudgetFilterCustomPeriodStartDate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetBudgetFilterCustomPeriodStartDate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodStartDateSet(c *Client, des, nw []BudgetBudgetFilterCustomPeriodStartDate) []BudgetBudgetFilterCustomPeriodStartDate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetBudgetFilterCustomPeriodStartDate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetBudgetFilterCustomPeriodStartDateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriodStartDate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodStartDateSlice(c *Client, des, nw []BudgetBudgetFilterCustomPeriodStartDate) []BudgetBudgetFilterCustomPeriodStartDate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetBudgetFilterCustomPeriodStartDate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriodStartDate(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetBudgetFilterCustomPeriodEndDate(des, initial *BudgetBudgetFilterCustomPeriodEndDate, opts ...dcl.ApplyOption) *BudgetBudgetFilterCustomPeriodEndDate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetBudgetFilterCustomPeriodEndDate{}

	if dcl.IsZeroValue(des.Year) || (dcl.IsEmptyValueIndirect(des.Year) && dcl.IsEmptyValueIndirect(initial.Year)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Year = initial.Year
	} else {
		cDes.Year = des.Year
	}
	if dcl.IsZeroValue(des.Month) || (dcl.IsEmptyValueIndirect(des.Month) && dcl.IsEmptyValueIndirect(initial.Month)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Month = initial.Month
	} else {
		cDes.Month = des.Month
	}
	if dcl.IsZeroValue(des.Day) || (dcl.IsEmptyValueIndirect(des.Day) && dcl.IsEmptyValueIndirect(initial.Day)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Day = initial.Day
	} else {
		cDes.Day = des.Day
	}

	return cDes
}

func canonicalizeBudgetBudgetFilterCustomPeriodEndDateSlice(des, initial []BudgetBudgetFilterCustomPeriodEndDate, opts ...dcl.ApplyOption) []BudgetBudgetFilterCustomPeriodEndDate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetBudgetFilterCustomPeriodEndDate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetBudgetFilterCustomPeriodEndDate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetBudgetFilterCustomPeriodEndDate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetBudgetFilterCustomPeriodEndDate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetBudgetFilterCustomPeriodEndDate(c *Client, des, nw *BudgetBudgetFilterCustomPeriodEndDate) *BudgetBudgetFilterCustomPeriodEndDate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetBudgetFilterCustomPeriodEndDate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodEndDateSet(c *Client, des, nw []BudgetBudgetFilterCustomPeriodEndDate) []BudgetBudgetFilterCustomPeriodEndDate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetBudgetFilterCustomPeriodEndDate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetBudgetFilterCustomPeriodEndDateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriodEndDate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetBudgetFilterCustomPeriodEndDateSlice(c *Client, des, nw []BudgetBudgetFilterCustomPeriodEndDate) []BudgetBudgetFilterCustomPeriodEndDate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetBudgetFilterCustomPeriodEndDate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetBudgetFilterCustomPeriodEndDate(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetAmount(des, initial *BudgetAmount, opts ...dcl.ApplyOption) *BudgetAmount {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.SpecifiedAmount != nil || (initial != nil && initial.SpecifiedAmount != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.LastPeriodAmount) {
			des.SpecifiedAmount = nil
			if initial != nil {
				initial.SpecifiedAmount = nil
			}
		}
	}

	if des.LastPeriodAmount != nil || (initial != nil && initial.LastPeriodAmount != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.SpecifiedAmount) {
			des.LastPeriodAmount = nil
			if initial != nil {
				initial.LastPeriodAmount = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetAmount{}

	cDes.SpecifiedAmount = canonicalizeBudgetAmountSpecifiedAmount(des.SpecifiedAmount, initial.SpecifiedAmount, opts...)
	cDes.LastPeriodAmount = canonicalizeBudgetAmountLastPeriodAmount(des.LastPeriodAmount, initial.LastPeriodAmount, opts...)

	return cDes
}

func canonicalizeBudgetAmountSlice(des, initial []BudgetAmount, opts ...dcl.ApplyOption) []BudgetAmount {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetAmount, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetAmount(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetAmount, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetAmount(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetAmount(c *Client, des, nw *BudgetAmount) *BudgetAmount {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetAmount while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SpecifiedAmount = canonicalizeNewBudgetAmountSpecifiedAmount(c, des.SpecifiedAmount, nw.SpecifiedAmount)
	nw.LastPeriodAmount = canonicalizeNewBudgetAmountLastPeriodAmount(c, des.LastPeriodAmount, nw.LastPeriodAmount)

	return nw
}

func canonicalizeNewBudgetAmountSet(c *Client, des, nw []BudgetAmount) []BudgetAmount {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetAmount
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetAmountNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetAmount(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetAmountSlice(c *Client, des, nw []BudgetAmount) []BudgetAmount {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetAmount
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetAmount(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetAmountSpecifiedAmount(des, initial *BudgetAmountSpecifiedAmount, opts ...dcl.ApplyOption) *BudgetAmountSpecifiedAmount {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetAmountSpecifiedAmount{}

	if dcl.StringCanonicalize(des.CurrencyCode, initial.CurrencyCode) || dcl.IsZeroValue(des.CurrencyCode) {
		cDes.CurrencyCode = initial.CurrencyCode
	} else {
		cDes.CurrencyCode = des.CurrencyCode
	}
	if dcl.IsZeroValue(des.Units) || (dcl.IsEmptyValueIndirect(des.Units) && dcl.IsEmptyValueIndirect(initial.Units)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Units = initial.Units
	} else {
		cDes.Units = des.Units
	}
	if dcl.IsZeroValue(des.Nanos) || (dcl.IsEmptyValueIndirect(des.Nanos) && dcl.IsEmptyValueIndirect(initial.Nanos)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Nanos = initial.Nanos
	} else {
		cDes.Nanos = des.Nanos
	}

	return cDes
}

func canonicalizeBudgetAmountSpecifiedAmountSlice(des, initial []BudgetAmountSpecifiedAmount, opts ...dcl.ApplyOption) []BudgetAmountSpecifiedAmount {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetAmountSpecifiedAmount, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetAmountSpecifiedAmount(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetAmountSpecifiedAmount, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetAmountSpecifiedAmount(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetAmountSpecifiedAmount(c *Client, des, nw *BudgetAmountSpecifiedAmount) *BudgetAmountSpecifiedAmount {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetAmountSpecifiedAmount while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.CurrencyCode, nw.CurrencyCode) {
		nw.CurrencyCode = des.CurrencyCode
	}

	return nw
}

func canonicalizeNewBudgetAmountSpecifiedAmountSet(c *Client, des, nw []BudgetAmountSpecifiedAmount) []BudgetAmountSpecifiedAmount {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetAmountSpecifiedAmount
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetAmountSpecifiedAmountNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetAmountSpecifiedAmount(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetAmountSpecifiedAmountSlice(c *Client, des, nw []BudgetAmountSpecifiedAmount) []BudgetAmountSpecifiedAmount {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetAmountSpecifiedAmount
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetAmountSpecifiedAmount(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetAmountLastPeriodAmount(des, initial *BudgetAmountLastPeriodAmount, opts ...dcl.ApplyOption) *BudgetAmountLastPeriodAmount {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &BudgetAmountLastPeriodAmount{}

	return cDes
}

func canonicalizeBudgetAmountLastPeriodAmountSlice(des, initial []BudgetAmountLastPeriodAmount, opts ...dcl.ApplyOption) []BudgetAmountLastPeriodAmount {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetAmountLastPeriodAmount, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetAmountLastPeriodAmount(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetAmountLastPeriodAmount, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetAmountLastPeriodAmount(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetAmountLastPeriodAmount(c *Client, des, nw *BudgetAmountLastPeriodAmount) *BudgetAmountLastPeriodAmount {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetAmountLastPeriodAmount while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewBudgetAmountLastPeriodAmountSet(c *Client, des, nw []BudgetAmountLastPeriodAmount) []BudgetAmountLastPeriodAmount {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetAmountLastPeriodAmount
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetAmountLastPeriodAmountNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetAmountLastPeriodAmount(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetAmountLastPeriodAmountSlice(c *Client, des, nw []BudgetAmountLastPeriodAmount) []BudgetAmountLastPeriodAmount {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetAmountLastPeriodAmount
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetAmountLastPeriodAmount(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetThresholdRules(des, initial *BudgetThresholdRules, opts ...dcl.ApplyOption) *BudgetThresholdRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetThresholdRules{}

	if dcl.IsZeroValue(des.ThresholdPercent) || (dcl.IsEmptyValueIndirect(des.ThresholdPercent) && dcl.IsEmptyValueIndirect(initial.ThresholdPercent)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ThresholdPercent = initial.ThresholdPercent
	} else {
		cDes.ThresholdPercent = des.ThresholdPercent
	}
	if dcl.IsZeroValue(des.SpendBasis) || (dcl.IsEmptyValueIndirect(des.SpendBasis) && dcl.IsEmptyValueIndirect(initial.SpendBasis)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SpendBasis = initial.SpendBasis
	} else {
		cDes.SpendBasis = des.SpendBasis
	}

	return cDes
}

func canonicalizeBudgetThresholdRulesSlice(des, initial []BudgetThresholdRules, opts ...dcl.ApplyOption) []BudgetThresholdRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetThresholdRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetThresholdRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetThresholdRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetThresholdRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetThresholdRules(c *Client, des, nw *BudgetThresholdRules) *BudgetThresholdRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetThresholdRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewBudgetThresholdRulesSet(c *Client, des, nw []BudgetThresholdRules) []BudgetThresholdRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetThresholdRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetThresholdRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetThresholdRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetThresholdRulesSlice(c *Client, des, nw []BudgetThresholdRules) []BudgetThresholdRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetThresholdRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetThresholdRules(c, &d, &n))
	}

	return items
}

func canonicalizeBudgetAllUpdatesRule(des, initial *BudgetAllUpdatesRule, opts ...dcl.ApplyOption) *BudgetAllUpdatesRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BudgetAllUpdatesRule{}

	if dcl.IsZeroValue(des.PubsubTopic) || (dcl.IsEmptyValueIndirect(des.PubsubTopic) && dcl.IsEmptyValueIndirect(initial.PubsubTopic)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.PubsubTopic = initial.PubsubTopic
	} else {
		cDes.PubsubTopic = des.PubsubTopic
	}
	if dcl.StringCanonicalize(des.SchemaVersion, initial.SchemaVersion) || dcl.IsZeroValue(des.SchemaVersion) {
		cDes.SchemaVersion = initial.SchemaVersion
	} else {
		cDes.SchemaVersion = des.SchemaVersion
	}
	if dcl.StringArrayCanonicalize(des.MonitoringNotificationChannels, initial.MonitoringNotificationChannels) {
		cDes.MonitoringNotificationChannels = initial.MonitoringNotificationChannels
	} else {
		cDes.MonitoringNotificationChannels = des.MonitoringNotificationChannels
	}
	if dcl.BoolCanonicalize(des.DisableDefaultIamRecipients, initial.DisableDefaultIamRecipients) || dcl.IsZeroValue(des.DisableDefaultIamRecipients) {
		cDes.DisableDefaultIamRecipients = initial.DisableDefaultIamRecipients
	} else {
		cDes.DisableDefaultIamRecipients = des.DisableDefaultIamRecipients
	}

	return cDes
}

func canonicalizeBudgetAllUpdatesRuleSlice(des, initial []BudgetAllUpdatesRule, opts ...dcl.ApplyOption) []BudgetAllUpdatesRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BudgetAllUpdatesRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBudgetAllUpdatesRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BudgetAllUpdatesRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBudgetAllUpdatesRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBudgetAllUpdatesRule(c *Client, des, nw *BudgetAllUpdatesRule) *BudgetAllUpdatesRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BudgetAllUpdatesRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SchemaVersion, nw.SchemaVersion) {
		nw.SchemaVersion = des.SchemaVersion
	}
	if dcl.StringArrayCanonicalize(des.MonitoringNotificationChannels, nw.MonitoringNotificationChannels) {
		nw.MonitoringNotificationChannels = des.MonitoringNotificationChannels
	}
	if dcl.BoolCanonicalize(des.DisableDefaultIamRecipients, nw.DisableDefaultIamRecipients) {
		nw.DisableDefaultIamRecipients = des.DisableDefaultIamRecipients
	}

	return nw
}

func canonicalizeNewBudgetAllUpdatesRuleSet(c *Client, des, nw []BudgetAllUpdatesRule) []BudgetAllUpdatesRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BudgetAllUpdatesRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBudgetAllUpdatesRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBudgetAllUpdatesRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBudgetAllUpdatesRuleSlice(c *Client, des, nw []BudgetAllUpdatesRule) []BudgetAllUpdatesRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BudgetAllUpdatesRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBudgetAllUpdatesRule(c, &d, &n))
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
func diffBudget(c *Client, desired, actual *Budget, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BudgetFilter, actual.BudgetFilter, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareBudgetBudgetFilterNewStyle, EmptyObject: EmptyBudgetBudgetFilter, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BudgetFilter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Amount, actual.Amount, dcl.DiffInfo{ObjectFunction: compareBudgetAmountNewStyle, EmptyObject: EmptyBudgetAmount, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Amount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ThresholdRules, actual.ThresholdRules, dcl.DiffInfo{ObjectFunction: compareBudgetThresholdRulesNewStyle, EmptyObject: EmptyBudgetThresholdRules, OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("ThresholdRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllUpdatesRule, actual.AllUpdatesRule, dcl.DiffInfo{ObjectFunction: compareBudgetAllUpdatesRuleNewStyle, EmptyObject: EmptyBudgetAllUpdatesRule, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("NotificationsRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BillingAccount, actual.BillingAccount, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("BillingAccount")); len(ds) != 0 || err != nil {
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
func compareBudgetBudgetFilterNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetBudgetFilter)
	if !ok {
		desiredNotPointer, ok := d.(BudgetBudgetFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilter or *BudgetBudgetFilter", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetBudgetFilter)
	if !ok {
		actualNotPointer, ok := a.(BudgetBudgetFilter)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilter", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Projects, actual.Projects, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Projects")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreditTypes, actual.CreditTypes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("CreditTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreditTypesTreatment, actual.CreditTypesTreatment, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("CreditTypesTreatment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Services, actual.Services, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Services")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subaccounts, actual.Subaccounts, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Subaccounts")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{ObjectFunction: compareBudgetBudgetFilterLabelsNewStyle, EmptyObject: EmptyBudgetBudgetFilterLabels, OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CalendarPeriod, actual.CalendarPeriod, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("CalendarPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomPeriod, actual.CustomPeriod, dcl.DiffInfo{MergeNestedDiffs: true, ObjectFunction: compareBudgetBudgetFilterCustomPeriodNewStyle, EmptyObject: EmptyBudgetBudgetFilterCustomPeriod, OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("CustomPeriod")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetBudgetFilterLabelsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetBudgetFilterLabels)
	if !ok {
		desiredNotPointer, ok := d.(BudgetBudgetFilterLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterLabels or *BudgetBudgetFilterLabels", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetBudgetFilterLabels)
	if !ok {
		actualNotPointer, ok := a.(BudgetBudgetFilterLabels)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterLabels", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Values, actual.Values, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Values")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetBudgetFilterCustomPeriodNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetBudgetFilterCustomPeriod)
	if !ok {
		desiredNotPointer, ok := d.(BudgetBudgetFilterCustomPeriod)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriod or *BudgetBudgetFilterCustomPeriod", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetBudgetFilterCustomPeriod)
	if !ok {
		actualNotPointer, ok := a.(BudgetBudgetFilterCustomPeriod)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriod", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StartDate, actual.StartDate, dcl.DiffInfo{ObjectFunction: compareBudgetBudgetFilterCustomPeriodStartDateNewStyle, EmptyObject: EmptyBudgetBudgetFilterCustomPeriodStartDate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StartDate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EndDate, actual.EndDate, dcl.DiffInfo{ObjectFunction: compareBudgetBudgetFilterCustomPeriodEndDateNewStyle, EmptyObject: EmptyBudgetBudgetFilterCustomPeriodEndDate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EndDate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetBudgetFilterCustomPeriodStartDateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetBudgetFilterCustomPeriodStartDate)
	if !ok {
		desiredNotPointer, ok := d.(BudgetBudgetFilterCustomPeriodStartDate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriodStartDate or *BudgetBudgetFilterCustomPeriodStartDate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetBudgetFilterCustomPeriodStartDate)
	if !ok {
		actualNotPointer, ok := a.(BudgetBudgetFilterCustomPeriodStartDate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriodStartDate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Year, actual.Year, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Year")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Month, actual.Month, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Month")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Day, actual.Day, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Day")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetBudgetFilterCustomPeriodEndDateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetBudgetFilterCustomPeriodEndDate)
	if !ok {
		desiredNotPointer, ok := d.(BudgetBudgetFilterCustomPeriodEndDate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriodEndDate or *BudgetBudgetFilterCustomPeriodEndDate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetBudgetFilterCustomPeriodEndDate)
	if !ok {
		actualNotPointer, ok := a.(BudgetBudgetFilterCustomPeriodEndDate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetBudgetFilterCustomPeriodEndDate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Year, actual.Year, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Year")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Month, actual.Month, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Month")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Day, actual.Day, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Day")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetAmountNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetAmount)
	if !ok {
		desiredNotPointer, ok := d.(BudgetAmount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAmount or *BudgetAmount", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetAmount)
	if !ok {
		actualNotPointer, ok := a.(BudgetAmount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAmount", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SpecifiedAmount, actual.SpecifiedAmount, dcl.DiffInfo{ObjectFunction: compareBudgetAmountSpecifiedAmountNewStyle, EmptyObject: EmptyBudgetAmountSpecifiedAmount, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SpecifiedAmount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastPeriodAmount, actual.LastPeriodAmount, dcl.DiffInfo{ObjectFunction: compareBudgetAmountLastPeriodAmountNewStyle, EmptyObject: EmptyBudgetAmountLastPeriodAmount, OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("LastPeriodAmount")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetAmountSpecifiedAmountNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetAmountSpecifiedAmount)
	if !ok {
		desiredNotPointer, ok := d.(BudgetAmountSpecifiedAmount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAmountSpecifiedAmount or *BudgetAmountSpecifiedAmount", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetAmountSpecifiedAmount)
	if !ok {
		actualNotPointer, ok := a.(BudgetAmountSpecifiedAmount)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAmountSpecifiedAmount", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CurrencyCode, actual.CurrencyCode, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CurrencyCode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Units, actual.Units, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Units")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Nanos, actual.Nanos, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("Nanos")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetAmountLastPeriodAmountNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareBudgetThresholdRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetThresholdRules)
	if !ok {
		desiredNotPointer, ok := d.(BudgetThresholdRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetThresholdRules or *BudgetThresholdRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetThresholdRules)
	if !ok {
		actualNotPointer, ok := a.(BudgetThresholdRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetThresholdRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ThresholdPercent, actual.ThresholdPercent, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("ThresholdPercent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SpendBasis, actual.SpendBasis, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("SpendBasis")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBudgetAllUpdatesRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BudgetAllUpdatesRule)
	if !ok {
		desiredNotPointer, ok := d.(BudgetAllUpdatesRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAllUpdatesRule or *BudgetAllUpdatesRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BudgetAllUpdatesRule)
	if !ok {
		actualNotPointer, ok := a.(BudgetAllUpdatesRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BudgetAllUpdatesRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.PubsubTopic, actual.PubsubTopic, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("PubsubTopic")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SchemaVersion, actual.SchemaVersion, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("SchemaVersion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MonitoringNotificationChannels, actual.MonitoringNotificationChannels, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("MonitoringNotificationChannels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisableDefaultIamRecipients, actual.DisableDefaultIamRecipients, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBudgetUpdateBudgetOperation")}, fn.AddNest("DisableDefaultIamRecipients")); len(ds) != 0 || err != nil {
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
func (r *Budget) urlNormalized() *Budget {
	normalized := dcl.Copy(*r).(Budget)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.BillingAccount = dcl.SelfLinkToName(r.BillingAccount)
	return &normalized
}

func (r *Budget) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateBudget" {
		fields := map[string]interface{}{
			"billingAccount": dcl.ValueOrEmptyString(nr.BillingAccount),
			"name":           dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("billingAccounts/{{billingAccount}}/budgets/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Budget resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Budget) marshal(c *Client) ([]byte, error) {
	m, err := expandBudget(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Budget: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBudget decodes JSON responses into the Budget resource schema.
func unmarshalBudget(b []byte, c *Client, res *Budget) (*Budget, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBudget(m, c, res)
}

func unmarshalMapBudget(m map[string]interface{}, c *Client, res *Budget) (*Budget, error) {

	flattened := flattenBudget(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandBudget expands Budget into a JSON request object.
func expandBudget(c *Client, f *Budget) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("billingAccounts/%s/budgets/%s", f.Name, dcl.SelfLinkToName(f.BillingAccount), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v, err := expandBudgetBudgetFilter(c, f.BudgetFilter, res); err != nil {
		return nil, fmt.Errorf("error expanding BudgetFilter into budgetFilter: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["budgetFilter"] = v
	}
	if v, err := expandBudgetAmount(c, f.Amount, res); err != nil {
		return nil, fmt.Errorf("error expanding Amount into amount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["amount"] = v
	}
	if v, err := expandBudgetThresholdRulesSlice(c, f.ThresholdRules, res); err != nil {
		return nil, fmt.Errorf("error expanding ThresholdRules into thresholdRules: %w", err)
	} else if v != nil {
		m["thresholdRules"] = v
	}
	if v, err := expandBudgetAllUpdatesRule(c, f.AllUpdatesRule, res); err != nil {
		return nil, fmt.Errorf("error expanding AllUpdatesRule into notificationsRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["notificationsRule"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding BillingAccount into billingAccount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["billingAccount"] = v
	}

	return m, nil
}

// flattenBudget flattens Budget from a JSON request object into the
// Budget type.
func flattenBudget(c *Client, i interface{}, res *Budget) *Budget {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Budget{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.BudgetFilter = flattenBudgetBudgetFilter(c, m["budgetFilter"], res)
	resultRes.Amount = flattenBudgetAmount(c, m["amount"], res)
	resultRes.ThresholdRules = flattenBudgetThresholdRulesSlice(c, m["thresholdRules"], res)
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.AllUpdatesRule = flattenBudgetAllUpdatesRule(c, m["notificationsRule"], res)
	resultRes.BillingAccount = dcl.FlattenString(m["billingAccount"])

	return resultRes
}

// expandBudgetBudgetFilterMap expands the contents of BudgetBudgetFilter into a JSON
// request object.
func expandBudgetBudgetFilterMap(c *Client, f map[string]BudgetBudgetFilter, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetBudgetFilter(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetBudgetFilterSlice expands the contents of BudgetBudgetFilter into a JSON
// request object.
func expandBudgetBudgetFilterSlice(c *Client, f []BudgetBudgetFilter, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetBudgetFilter(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetBudgetFilterMap flattens the contents of BudgetBudgetFilter from a JSON
// response object.
func flattenBudgetBudgetFilterMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilter {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilter{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilter{}
	}

	items := make(map[string]BudgetBudgetFilter)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilter(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetBudgetFilterSlice flattens the contents of BudgetBudgetFilter from a JSON
// response object.
func flattenBudgetBudgetFilterSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilter {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilter{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilter{}
	}

	items := make([]BudgetBudgetFilter, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilter(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetBudgetFilter expands an instance of BudgetBudgetFilter into a JSON
// request object.
func expandBudgetBudgetFilter(c *Client, f *BudgetBudgetFilter, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Projects; v != nil {
		m["projects"] = v
	}
	if v := f.CreditTypes; v != nil {
		m["creditTypes"] = v
	}
	if v := f.CreditTypesTreatment; !dcl.IsEmptyValueIndirect(v) {
		m["creditTypesTreatment"] = v
	}
	if v := f.Services; v != nil {
		m["services"] = v
	}
	if v := f.Subaccounts; v != nil {
		m["subaccounts"] = v
	}
	if v, err := expandBudgetFilterLabels(c, f.Labels, res); err != nil {
		return nil, fmt.Errorf("error expanding Labels into labels: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["labels"] = v
	}
	if v := f.CalendarPeriod; !dcl.IsEmptyValueIndirect(v) {
		m["calendarPeriod"] = v
	}
	if v, err := expandBudgetBudgetFilterCustomPeriod(c, f.CustomPeriod, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomPeriod into customPeriod: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["customPeriod"] = v
	}

	return m, nil
}

// flattenBudgetBudgetFilter flattens an instance of BudgetBudgetFilter from a JSON
// response object.
func flattenBudgetBudgetFilter(c *Client, i interface{}, res *Budget) *BudgetBudgetFilter {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetBudgetFilter{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetBudgetFilter
	}
	r.Projects = flattenBudgetFilterProjects(c, m["projects"], res)
	r.CreditTypes = dcl.FlattenStringSlice(m["creditTypes"])
	r.CreditTypesTreatment = flattenBudgetBudgetFilterCreditTypesTreatmentEnum(m["creditTypesTreatment"])
	r.Services = dcl.FlattenStringSlice(m["services"])
	r.Subaccounts = dcl.FlattenStringSlice(m["subaccounts"])
	r.Labels = flattenBudgetFilterLabels(c, m["labels"], res)
	r.CalendarPeriod = flattenBudgetBudgetFilterCalendarPeriodEnum(m["calendarPeriod"])
	r.CustomPeriod = flattenBudgetBudgetFilterCustomPeriod(c, m["customPeriod"], res)

	return r
}

// expandBudgetBudgetFilterLabelsMap expands the contents of BudgetBudgetFilterLabels into a JSON
// request object.
func expandBudgetBudgetFilterLabelsMap(c *Client, f map[string]BudgetBudgetFilterLabels, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetBudgetFilterLabels(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetBudgetFilterLabelsSlice expands the contents of BudgetBudgetFilterLabels into a JSON
// request object.
func expandBudgetBudgetFilterLabelsSlice(c *Client, f []BudgetBudgetFilterLabels, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetBudgetFilterLabels(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetBudgetFilterLabelsMap flattens the contents of BudgetBudgetFilterLabels from a JSON
// response object.
func flattenBudgetBudgetFilterLabelsMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterLabels {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterLabels{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterLabels{}
	}

	items := make(map[string]BudgetBudgetFilterLabels)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterLabels(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetBudgetFilterLabelsSlice flattens the contents of BudgetBudgetFilterLabels from a JSON
// response object.
func flattenBudgetBudgetFilterLabelsSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterLabels {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterLabels{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterLabels{}
	}

	items := make([]BudgetBudgetFilterLabels, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterLabels(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetBudgetFilterLabels expands an instance of BudgetBudgetFilterLabels into a JSON
// request object.
func expandBudgetBudgetFilterLabels(c *Client, f *BudgetBudgetFilterLabels, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Values; v != nil {
		m["values"] = v
	}

	return m, nil
}

// flattenBudgetBudgetFilterLabels flattens an instance of BudgetBudgetFilterLabels from a JSON
// response object.
func flattenBudgetBudgetFilterLabels(c *Client, i interface{}, res *Budget) *BudgetBudgetFilterLabels {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetBudgetFilterLabels{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetBudgetFilterLabels
	}
	r.Values = dcl.FlattenStringSlice(m["values"])

	return r
}

// expandBudgetBudgetFilterCustomPeriodMap expands the contents of BudgetBudgetFilterCustomPeriod into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodMap(c *Client, f map[string]BudgetBudgetFilterCustomPeriod, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriod(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetBudgetFilterCustomPeriodSlice expands the contents of BudgetBudgetFilterCustomPeriod into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodSlice(c *Client, f []BudgetBudgetFilterCustomPeriod, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriod(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetBudgetFilterCustomPeriodMap flattens the contents of BudgetBudgetFilterCustomPeriod from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterCustomPeriod {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterCustomPeriod{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterCustomPeriod{}
	}

	items := make(map[string]BudgetBudgetFilterCustomPeriod)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterCustomPeriod(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetBudgetFilterCustomPeriodSlice flattens the contents of BudgetBudgetFilterCustomPeriod from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterCustomPeriod {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterCustomPeriod{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterCustomPeriod{}
	}

	items := make([]BudgetBudgetFilterCustomPeriod, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterCustomPeriod(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetBudgetFilterCustomPeriod expands an instance of BudgetBudgetFilterCustomPeriod into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriod(c *Client, f *BudgetBudgetFilterCustomPeriod, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBudgetBudgetFilterCustomPeriodStartDate(c, f.StartDate, res); err != nil {
		return nil, fmt.Errorf("error expanding StartDate into startDate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["startDate"] = v
	}
	if v, err := expandBudgetBudgetFilterCustomPeriodEndDate(c, f.EndDate, res); err != nil {
		return nil, fmt.Errorf("error expanding EndDate into endDate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["endDate"] = v
	}

	return m, nil
}

// flattenBudgetBudgetFilterCustomPeriod flattens an instance of BudgetBudgetFilterCustomPeriod from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriod(c *Client, i interface{}, res *Budget) *BudgetBudgetFilterCustomPeriod {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetBudgetFilterCustomPeriod{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetBudgetFilterCustomPeriod
	}
	r.StartDate = flattenBudgetBudgetFilterCustomPeriodStartDate(c, m["startDate"], res)
	r.EndDate = flattenBudgetBudgetFilterCustomPeriodEndDate(c, m["endDate"], res)

	return r
}

// expandBudgetBudgetFilterCustomPeriodStartDateMap expands the contents of BudgetBudgetFilterCustomPeriodStartDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodStartDateMap(c *Client, f map[string]BudgetBudgetFilterCustomPeriodStartDate, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriodStartDate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetBudgetFilterCustomPeriodStartDateSlice expands the contents of BudgetBudgetFilterCustomPeriodStartDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodStartDateSlice(c *Client, f []BudgetBudgetFilterCustomPeriodStartDate, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriodStartDate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetBudgetFilterCustomPeriodStartDateMap flattens the contents of BudgetBudgetFilterCustomPeriodStartDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodStartDateMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterCustomPeriodStartDate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterCustomPeriodStartDate{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterCustomPeriodStartDate{}
	}

	items := make(map[string]BudgetBudgetFilterCustomPeriodStartDate)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterCustomPeriodStartDate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetBudgetFilterCustomPeriodStartDateSlice flattens the contents of BudgetBudgetFilterCustomPeriodStartDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodStartDateSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterCustomPeriodStartDate {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterCustomPeriodStartDate{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterCustomPeriodStartDate{}
	}

	items := make([]BudgetBudgetFilterCustomPeriodStartDate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterCustomPeriodStartDate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetBudgetFilterCustomPeriodStartDate expands an instance of BudgetBudgetFilterCustomPeriodStartDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodStartDate(c *Client, f *BudgetBudgetFilterCustomPeriodStartDate, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Year; !dcl.IsEmptyValueIndirect(v) {
		m["year"] = v
	}
	if v := f.Month; !dcl.IsEmptyValueIndirect(v) {
		m["month"] = v
	}
	if v := f.Day; !dcl.IsEmptyValueIndirect(v) {
		m["day"] = v
	}

	return m, nil
}

// flattenBudgetBudgetFilterCustomPeriodStartDate flattens an instance of BudgetBudgetFilterCustomPeriodStartDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodStartDate(c *Client, i interface{}, res *Budget) *BudgetBudgetFilterCustomPeriodStartDate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetBudgetFilterCustomPeriodStartDate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetBudgetFilterCustomPeriodStartDate
	}
	r.Year = dcl.FlattenInteger(m["year"])
	r.Month = dcl.FlattenInteger(m["month"])
	r.Day = dcl.FlattenInteger(m["day"])

	return r
}

// expandBudgetBudgetFilterCustomPeriodEndDateMap expands the contents of BudgetBudgetFilterCustomPeriodEndDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodEndDateMap(c *Client, f map[string]BudgetBudgetFilterCustomPeriodEndDate, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriodEndDate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetBudgetFilterCustomPeriodEndDateSlice expands the contents of BudgetBudgetFilterCustomPeriodEndDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodEndDateSlice(c *Client, f []BudgetBudgetFilterCustomPeriodEndDate, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetBudgetFilterCustomPeriodEndDate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetBudgetFilterCustomPeriodEndDateMap flattens the contents of BudgetBudgetFilterCustomPeriodEndDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodEndDateMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterCustomPeriodEndDate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterCustomPeriodEndDate{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterCustomPeriodEndDate{}
	}

	items := make(map[string]BudgetBudgetFilterCustomPeriodEndDate)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterCustomPeriodEndDate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetBudgetFilterCustomPeriodEndDateSlice flattens the contents of BudgetBudgetFilterCustomPeriodEndDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodEndDateSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterCustomPeriodEndDate {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterCustomPeriodEndDate{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterCustomPeriodEndDate{}
	}

	items := make([]BudgetBudgetFilterCustomPeriodEndDate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterCustomPeriodEndDate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetBudgetFilterCustomPeriodEndDate expands an instance of BudgetBudgetFilterCustomPeriodEndDate into a JSON
// request object.
func expandBudgetBudgetFilterCustomPeriodEndDate(c *Client, f *BudgetBudgetFilterCustomPeriodEndDate, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Year; !dcl.IsEmptyValueIndirect(v) {
		m["year"] = v
	}
	if v := f.Month; !dcl.IsEmptyValueIndirect(v) {
		m["month"] = v
	}
	if v := f.Day; !dcl.IsEmptyValueIndirect(v) {
		m["day"] = v
	}

	return m, nil
}

// flattenBudgetBudgetFilterCustomPeriodEndDate flattens an instance of BudgetBudgetFilterCustomPeriodEndDate from a JSON
// response object.
func flattenBudgetBudgetFilterCustomPeriodEndDate(c *Client, i interface{}, res *Budget) *BudgetBudgetFilterCustomPeriodEndDate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetBudgetFilterCustomPeriodEndDate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetBudgetFilterCustomPeriodEndDate
	}
	r.Year = dcl.FlattenInteger(m["year"])
	r.Month = dcl.FlattenInteger(m["month"])
	r.Day = dcl.FlattenInteger(m["day"])

	return r
}

// expandBudgetAmountMap expands the contents of BudgetAmount into a JSON
// request object.
func expandBudgetAmountMap(c *Client, f map[string]BudgetAmount, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetAmount(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetAmountSlice expands the contents of BudgetAmount into a JSON
// request object.
func expandBudgetAmountSlice(c *Client, f []BudgetAmount, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetAmount(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetAmountMap flattens the contents of BudgetAmount from a JSON
// response object.
func flattenBudgetAmountMap(c *Client, i interface{}, res *Budget) map[string]BudgetAmount {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetAmount{}
	}

	if len(a) == 0 {
		return map[string]BudgetAmount{}
	}

	items := make(map[string]BudgetAmount)
	for k, item := range a {
		items[k] = *flattenBudgetAmount(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetAmountSlice flattens the contents of BudgetAmount from a JSON
// response object.
func flattenBudgetAmountSlice(c *Client, i interface{}, res *Budget) []BudgetAmount {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetAmount{}
	}

	if len(a) == 0 {
		return []BudgetAmount{}
	}

	items := make([]BudgetAmount, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetAmount(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetAmount expands an instance of BudgetAmount into a JSON
// request object.
func expandBudgetAmount(c *Client, f *BudgetAmount, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBudgetAmountSpecifiedAmount(c, f.SpecifiedAmount, res); err != nil {
		return nil, fmt.Errorf("error expanding SpecifiedAmount into specifiedAmount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["specifiedAmount"] = v
	}
	if v, err := expandBudgetAmountLastPeriodAmount(c, f.LastPeriodAmount, res); err != nil {
		return nil, fmt.Errorf("error expanding LastPeriodAmount into lastPeriodAmount: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["lastPeriodAmount"] = v
	}

	return m, nil
}

// flattenBudgetAmount flattens an instance of BudgetAmount from a JSON
// response object.
func flattenBudgetAmount(c *Client, i interface{}, res *Budget) *BudgetAmount {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetAmount{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetAmount
	}
	r.SpecifiedAmount = flattenBudgetAmountSpecifiedAmount(c, m["specifiedAmount"], res)
	r.LastPeriodAmount = flattenBudgetAmountLastPeriodAmount(c, m["lastPeriodAmount"], res)

	return r
}

// expandBudgetAmountSpecifiedAmountMap expands the contents of BudgetAmountSpecifiedAmount into a JSON
// request object.
func expandBudgetAmountSpecifiedAmountMap(c *Client, f map[string]BudgetAmountSpecifiedAmount, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetAmountSpecifiedAmount(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetAmountSpecifiedAmountSlice expands the contents of BudgetAmountSpecifiedAmount into a JSON
// request object.
func expandBudgetAmountSpecifiedAmountSlice(c *Client, f []BudgetAmountSpecifiedAmount, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetAmountSpecifiedAmount(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetAmountSpecifiedAmountMap flattens the contents of BudgetAmountSpecifiedAmount from a JSON
// response object.
func flattenBudgetAmountSpecifiedAmountMap(c *Client, i interface{}, res *Budget) map[string]BudgetAmountSpecifiedAmount {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetAmountSpecifiedAmount{}
	}

	if len(a) == 0 {
		return map[string]BudgetAmountSpecifiedAmount{}
	}

	items := make(map[string]BudgetAmountSpecifiedAmount)
	for k, item := range a {
		items[k] = *flattenBudgetAmountSpecifiedAmount(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetAmountSpecifiedAmountSlice flattens the contents of BudgetAmountSpecifiedAmount from a JSON
// response object.
func flattenBudgetAmountSpecifiedAmountSlice(c *Client, i interface{}, res *Budget) []BudgetAmountSpecifiedAmount {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetAmountSpecifiedAmount{}
	}

	if len(a) == 0 {
		return []BudgetAmountSpecifiedAmount{}
	}

	items := make([]BudgetAmountSpecifiedAmount, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetAmountSpecifiedAmount(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetAmountSpecifiedAmount expands an instance of BudgetAmountSpecifiedAmount into a JSON
// request object.
func expandBudgetAmountSpecifiedAmount(c *Client, f *BudgetAmountSpecifiedAmount, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.CurrencyCode; !dcl.IsEmptyValueIndirect(v) {
		m["currencyCode"] = v
	}
	if v := f.Units; v != nil {
		m["units"] = v
	}
	if v := f.Nanos; v != nil {
		m["nanos"] = v
	}

	return m, nil
}

// flattenBudgetAmountSpecifiedAmount flattens an instance of BudgetAmountSpecifiedAmount from a JSON
// response object.
func flattenBudgetAmountSpecifiedAmount(c *Client, i interface{}, res *Budget) *BudgetAmountSpecifiedAmount {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetAmountSpecifiedAmount{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetAmountSpecifiedAmount
	}
	r.CurrencyCode = dcl.FlattenString(m["currencyCode"])
	r.Units = dcl.FlattenInteger(m["units"])
	r.Nanos = dcl.FlattenInteger(m["nanos"])

	return r
}

// expandBudgetAmountLastPeriodAmountMap expands the contents of BudgetAmountLastPeriodAmount into a JSON
// request object.
func expandBudgetAmountLastPeriodAmountMap(c *Client, f map[string]BudgetAmountLastPeriodAmount, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetAmountLastPeriodAmount(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetAmountLastPeriodAmountSlice expands the contents of BudgetAmountLastPeriodAmount into a JSON
// request object.
func expandBudgetAmountLastPeriodAmountSlice(c *Client, f []BudgetAmountLastPeriodAmount, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetAmountLastPeriodAmount(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetAmountLastPeriodAmountMap flattens the contents of BudgetAmountLastPeriodAmount from a JSON
// response object.
func flattenBudgetAmountLastPeriodAmountMap(c *Client, i interface{}, res *Budget) map[string]BudgetAmountLastPeriodAmount {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetAmountLastPeriodAmount{}
	}

	if len(a) == 0 {
		return map[string]BudgetAmountLastPeriodAmount{}
	}

	items := make(map[string]BudgetAmountLastPeriodAmount)
	for k, item := range a {
		items[k] = *flattenBudgetAmountLastPeriodAmount(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetAmountLastPeriodAmountSlice flattens the contents of BudgetAmountLastPeriodAmount from a JSON
// response object.
func flattenBudgetAmountLastPeriodAmountSlice(c *Client, i interface{}, res *Budget) []BudgetAmountLastPeriodAmount {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetAmountLastPeriodAmount{}
	}

	if len(a) == 0 {
		return []BudgetAmountLastPeriodAmount{}
	}

	items := make([]BudgetAmountLastPeriodAmount, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetAmountLastPeriodAmount(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetAmountLastPeriodAmount expands an instance of BudgetAmountLastPeriodAmount into a JSON
// request object.
func expandBudgetAmountLastPeriodAmount(c *Client, f *BudgetAmountLastPeriodAmount, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenBudgetAmountLastPeriodAmount flattens an instance of BudgetAmountLastPeriodAmount from a JSON
// response object.
func flattenBudgetAmountLastPeriodAmount(c *Client, i interface{}, res *Budget) *BudgetAmountLastPeriodAmount {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetAmountLastPeriodAmount{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetAmountLastPeriodAmount
	}

	return r
}

// expandBudgetThresholdRulesMap expands the contents of BudgetThresholdRules into a JSON
// request object.
func expandBudgetThresholdRulesMap(c *Client, f map[string]BudgetThresholdRules, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetThresholdRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetThresholdRulesSlice expands the contents of BudgetThresholdRules into a JSON
// request object.
func expandBudgetThresholdRulesSlice(c *Client, f []BudgetThresholdRules, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetThresholdRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetThresholdRulesMap flattens the contents of BudgetThresholdRules from a JSON
// response object.
func flattenBudgetThresholdRulesMap(c *Client, i interface{}, res *Budget) map[string]BudgetThresholdRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetThresholdRules{}
	}

	if len(a) == 0 {
		return map[string]BudgetThresholdRules{}
	}

	items := make(map[string]BudgetThresholdRules)
	for k, item := range a {
		items[k] = *flattenBudgetThresholdRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetThresholdRulesSlice flattens the contents of BudgetThresholdRules from a JSON
// response object.
func flattenBudgetThresholdRulesSlice(c *Client, i interface{}, res *Budget) []BudgetThresholdRules {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetThresholdRules{}
	}

	if len(a) == 0 {
		return []BudgetThresholdRules{}
	}

	items := make([]BudgetThresholdRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetThresholdRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetThresholdRules expands an instance of BudgetThresholdRules into a JSON
// request object.
func expandBudgetThresholdRules(c *Client, f *BudgetThresholdRules, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ThresholdPercent; !dcl.IsEmptyValueIndirect(v) {
		m["thresholdPercent"] = v
	}
	if v := f.SpendBasis; !dcl.IsEmptyValueIndirect(v) {
		m["spendBasis"] = v
	}

	return m, nil
}

// flattenBudgetThresholdRules flattens an instance of BudgetThresholdRules from a JSON
// response object.
func flattenBudgetThresholdRules(c *Client, i interface{}, res *Budget) *BudgetThresholdRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetThresholdRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetThresholdRules
	}
	r.ThresholdPercent = dcl.FlattenDouble(m["thresholdPercent"])
	r.SpendBasis = flattenBudgetThresholdRulesSpendBasisEnum(m["spendBasis"])

	return r
}

// expandBudgetAllUpdatesRuleMap expands the contents of BudgetAllUpdatesRule into a JSON
// request object.
func expandBudgetAllUpdatesRuleMap(c *Client, f map[string]BudgetAllUpdatesRule, res *Budget) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBudgetAllUpdatesRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBudgetAllUpdatesRuleSlice expands the contents of BudgetAllUpdatesRule into a JSON
// request object.
func expandBudgetAllUpdatesRuleSlice(c *Client, f []BudgetAllUpdatesRule, res *Budget) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBudgetAllUpdatesRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBudgetAllUpdatesRuleMap flattens the contents of BudgetAllUpdatesRule from a JSON
// response object.
func flattenBudgetAllUpdatesRuleMap(c *Client, i interface{}, res *Budget) map[string]BudgetAllUpdatesRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetAllUpdatesRule{}
	}

	if len(a) == 0 {
		return map[string]BudgetAllUpdatesRule{}
	}

	items := make(map[string]BudgetAllUpdatesRule)
	for k, item := range a {
		items[k] = *flattenBudgetAllUpdatesRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBudgetAllUpdatesRuleSlice flattens the contents of BudgetAllUpdatesRule from a JSON
// response object.
func flattenBudgetAllUpdatesRuleSlice(c *Client, i interface{}, res *Budget) []BudgetAllUpdatesRule {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetAllUpdatesRule{}
	}

	if len(a) == 0 {
		return []BudgetAllUpdatesRule{}
	}

	items := make([]BudgetAllUpdatesRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetAllUpdatesRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBudgetAllUpdatesRule expands an instance of BudgetAllUpdatesRule into a JSON
// request object.
func expandBudgetAllUpdatesRule(c *Client, f *BudgetAllUpdatesRule, res *Budget) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.PubsubTopic; !dcl.IsEmptyValueIndirect(v) {
		m["pubsubTopic"] = v
	}
	if v := f.SchemaVersion; !dcl.IsEmptyValueIndirect(v) {
		m["schemaVersion"] = v
	}
	if v := f.MonitoringNotificationChannels; v != nil {
		m["monitoringNotificationChannels"] = v
	}
	if v := f.DisableDefaultIamRecipients; !dcl.IsEmptyValueIndirect(v) {
		m["disableDefaultIamRecipients"] = v
	}

	return m, nil
}

// flattenBudgetAllUpdatesRule flattens an instance of BudgetAllUpdatesRule from a JSON
// response object.
func flattenBudgetAllUpdatesRule(c *Client, i interface{}, res *Budget) *BudgetAllUpdatesRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BudgetAllUpdatesRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBudgetAllUpdatesRule
	}
	r.PubsubTopic = dcl.FlattenString(m["pubsubTopic"])
	r.SchemaVersion = dcl.FlattenString(m["schemaVersion"])
	r.MonitoringNotificationChannels = dcl.FlattenStringSlice(m["monitoringNotificationChannels"])
	r.DisableDefaultIamRecipients = dcl.FlattenBool(m["disableDefaultIamRecipients"])

	return r
}

// flattenBudgetBudgetFilterCreditTypesTreatmentEnumMap flattens the contents of BudgetBudgetFilterCreditTypesTreatmentEnum from a JSON
// response object.
func flattenBudgetBudgetFilterCreditTypesTreatmentEnumMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterCreditTypesTreatmentEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterCreditTypesTreatmentEnum{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterCreditTypesTreatmentEnum{}
	}

	items := make(map[string]BudgetBudgetFilterCreditTypesTreatmentEnum)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterCreditTypesTreatmentEnum(item.(interface{}))
	}

	return items
}

// flattenBudgetBudgetFilterCreditTypesTreatmentEnumSlice flattens the contents of BudgetBudgetFilterCreditTypesTreatmentEnum from a JSON
// response object.
func flattenBudgetBudgetFilterCreditTypesTreatmentEnumSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterCreditTypesTreatmentEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterCreditTypesTreatmentEnum{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterCreditTypesTreatmentEnum{}
	}

	items := make([]BudgetBudgetFilterCreditTypesTreatmentEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterCreditTypesTreatmentEnum(item.(interface{})))
	}

	return items
}

// flattenBudgetBudgetFilterCreditTypesTreatmentEnum asserts that an interface is a string, and returns a
// pointer to a *BudgetBudgetFilterCreditTypesTreatmentEnum with the same value as that string.
func flattenBudgetBudgetFilterCreditTypesTreatmentEnum(i interface{}) *BudgetBudgetFilterCreditTypesTreatmentEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BudgetBudgetFilterCreditTypesTreatmentEnumRef(s)
}

// flattenBudgetBudgetFilterCalendarPeriodEnumMap flattens the contents of BudgetBudgetFilterCalendarPeriodEnum from a JSON
// response object.
func flattenBudgetBudgetFilterCalendarPeriodEnumMap(c *Client, i interface{}, res *Budget) map[string]BudgetBudgetFilterCalendarPeriodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetBudgetFilterCalendarPeriodEnum{}
	}

	if len(a) == 0 {
		return map[string]BudgetBudgetFilterCalendarPeriodEnum{}
	}

	items := make(map[string]BudgetBudgetFilterCalendarPeriodEnum)
	for k, item := range a {
		items[k] = *flattenBudgetBudgetFilterCalendarPeriodEnum(item.(interface{}))
	}

	return items
}

// flattenBudgetBudgetFilterCalendarPeriodEnumSlice flattens the contents of BudgetBudgetFilterCalendarPeriodEnum from a JSON
// response object.
func flattenBudgetBudgetFilterCalendarPeriodEnumSlice(c *Client, i interface{}, res *Budget) []BudgetBudgetFilterCalendarPeriodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetBudgetFilterCalendarPeriodEnum{}
	}

	if len(a) == 0 {
		return []BudgetBudgetFilterCalendarPeriodEnum{}
	}

	items := make([]BudgetBudgetFilterCalendarPeriodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetBudgetFilterCalendarPeriodEnum(item.(interface{})))
	}

	return items
}

// flattenBudgetBudgetFilterCalendarPeriodEnum asserts that an interface is a string, and returns a
// pointer to a *BudgetBudgetFilterCalendarPeriodEnum with the same value as that string.
func flattenBudgetBudgetFilterCalendarPeriodEnum(i interface{}) *BudgetBudgetFilterCalendarPeriodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BudgetBudgetFilterCalendarPeriodEnumRef(s)
}

// flattenBudgetThresholdRulesSpendBasisEnumMap flattens the contents of BudgetThresholdRulesSpendBasisEnum from a JSON
// response object.
func flattenBudgetThresholdRulesSpendBasisEnumMap(c *Client, i interface{}, res *Budget) map[string]BudgetThresholdRulesSpendBasisEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BudgetThresholdRulesSpendBasisEnum{}
	}

	if len(a) == 0 {
		return map[string]BudgetThresholdRulesSpendBasisEnum{}
	}

	items := make(map[string]BudgetThresholdRulesSpendBasisEnum)
	for k, item := range a {
		items[k] = *flattenBudgetThresholdRulesSpendBasisEnum(item.(interface{}))
	}

	return items
}

// flattenBudgetThresholdRulesSpendBasisEnumSlice flattens the contents of BudgetThresholdRulesSpendBasisEnum from a JSON
// response object.
func flattenBudgetThresholdRulesSpendBasisEnumSlice(c *Client, i interface{}, res *Budget) []BudgetThresholdRulesSpendBasisEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BudgetThresholdRulesSpendBasisEnum{}
	}

	if len(a) == 0 {
		return []BudgetThresholdRulesSpendBasisEnum{}
	}

	items := make([]BudgetThresholdRulesSpendBasisEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBudgetThresholdRulesSpendBasisEnum(item.(interface{})))
	}

	return items
}

// flattenBudgetThresholdRulesSpendBasisEnum asserts that an interface is a string, and returns a
// pointer to a *BudgetThresholdRulesSpendBasisEnum with the same value as that string.
func flattenBudgetThresholdRulesSpendBasisEnum(i interface{}) *BudgetThresholdRulesSpendBasisEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BudgetThresholdRulesSpendBasisEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Budget) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBudget(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.BillingAccount == nil && ncr.BillingAccount == nil {
			c.Config.Logger.Info("Both BillingAccount fields null - considering equal.")
		} else if nr.BillingAccount == nil || ncr.BillingAccount == nil {
			c.Config.Logger.Info("Only one BillingAccount field is null - considering unequal.")
			return false
		} else if *nr.BillingAccount != *ncr.BillingAccount {
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

type budgetDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         budgetApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToBudgetDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]budgetDiff, error) {
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
	var diffs []budgetDiff
	// For each operation name, create a budgetDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := budgetDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBudgetApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBudgetApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (budgetApiOperation, error) {
	switch opName {

	case "updateBudgetUpdateBudgetOperation":
		return &updateBudgetUpdateBudgetOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractBudgetFields(r *Budget) error {
	vBudgetFilter := r.BudgetFilter
	if vBudgetFilter == nil {
		// note: explicitly not the empty object.
		vBudgetFilter = &BudgetBudgetFilter{}
	}
	if err := extractBudgetBudgetFilterFields(r, vBudgetFilter); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBudgetFilter) {
		r.BudgetFilter = vBudgetFilter
	}
	vAmount := r.Amount
	if vAmount == nil {
		// note: explicitly not the empty object.
		vAmount = &BudgetAmount{}
	}
	if err := extractBudgetAmountFields(r, vAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAmount) {
		r.Amount = vAmount
	}
	vAllUpdatesRule := r.AllUpdatesRule
	if vAllUpdatesRule == nil {
		// note: explicitly not the empty object.
		vAllUpdatesRule = &BudgetAllUpdatesRule{}
	}
	if err := extractBudgetAllUpdatesRuleFields(r, vAllUpdatesRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAllUpdatesRule) {
		r.AllUpdatesRule = vAllUpdatesRule
	}
	return nil
}
func extractBudgetBudgetFilterFields(r *Budget, o *BudgetBudgetFilter) error {
	vCustomPeriod := o.CustomPeriod
	if vCustomPeriod == nil {
		// note: explicitly not the empty object.
		vCustomPeriod = &BudgetBudgetFilterCustomPeriod{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodFields(r, vCustomPeriod); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomPeriod) {
		o.CustomPeriod = vCustomPeriod
	}
	return nil
}
func extractBudgetBudgetFilterLabelsFields(r *Budget, o *BudgetBudgetFilterLabels) error {
	return nil
}
func extractBudgetBudgetFilterCustomPeriodFields(r *Budget, o *BudgetBudgetFilterCustomPeriod) error {
	vStartDate := o.StartDate
	if vStartDate == nil {
		// note: explicitly not the empty object.
		vStartDate = &BudgetBudgetFilterCustomPeriodStartDate{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodStartDateFields(r, vStartDate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStartDate) {
		o.StartDate = vStartDate
	}
	vEndDate := o.EndDate
	if vEndDate == nil {
		// note: explicitly not the empty object.
		vEndDate = &BudgetBudgetFilterCustomPeriodEndDate{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodEndDateFields(r, vEndDate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndDate) {
		o.EndDate = vEndDate
	}
	return nil
}
func extractBudgetBudgetFilterCustomPeriodStartDateFields(r *Budget, o *BudgetBudgetFilterCustomPeriodStartDate) error {
	return nil
}
func extractBudgetBudgetFilterCustomPeriodEndDateFields(r *Budget, o *BudgetBudgetFilterCustomPeriodEndDate) error {
	return nil
}
func extractBudgetAmountFields(r *Budget, o *BudgetAmount) error {
	vSpecifiedAmount := o.SpecifiedAmount
	if vSpecifiedAmount == nil {
		// note: explicitly not the empty object.
		vSpecifiedAmount = &BudgetAmountSpecifiedAmount{}
	}
	if err := extractBudgetAmountSpecifiedAmountFields(r, vSpecifiedAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpecifiedAmount) {
		o.SpecifiedAmount = vSpecifiedAmount
	}
	vLastPeriodAmount := o.LastPeriodAmount
	if vLastPeriodAmount == nil {
		// note: explicitly not the empty object.
		vLastPeriodAmount = &BudgetAmountLastPeriodAmount{}
	}
	if err := extractBudgetAmountLastPeriodAmountFields(r, vLastPeriodAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLastPeriodAmount) {
		o.LastPeriodAmount = vLastPeriodAmount
	}
	return nil
}
func extractBudgetAmountSpecifiedAmountFields(r *Budget, o *BudgetAmountSpecifiedAmount) error {
	return nil
}
func extractBudgetAmountLastPeriodAmountFields(r *Budget, o *BudgetAmountLastPeriodAmount) error {
	return nil
}
func extractBudgetThresholdRulesFields(r *Budget, o *BudgetThresholdRules) error {
	return nil
}
func extractBudgetAllUpdatesRuleFields(r *Budget, o *BudgetAllUpdatesRule) error {
	return nil
}

func postReadExtractBudgetFields(r *Budget) error {
	vBudgetFilter := r.BudgetFilter
	if vBudgetFilter == nil {
		// note: explicitly not the empty object.
		vBudgetFilter = &BudgetBudgetFilter{}
	}
	if err := postReadExtractBudgetBudgetFilterFields(r, vBudgetFilter); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBudgetFilter) {
		r.BudgetFilter = vBudgetFilter
	}
	vAmount := r.Amount
	if vAmount == nil {
		// note: explicitly not the empty object.
		vAmount = &BudgetAmount{}
	}
	if err := postReadExtractBudgetAmountFields(r, vAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAmount) {
		r.Amount = vAmount
	}
	vAllUpdatesRule := r.AllUpdatesRule
	if vAllUpdatesRule == nil {
		// note: explicitly not the empty object.
		vAllUpdatesRule = &BudgetAllUpdatesRule{}
	}
	if err := postReadExtractBudgetAllUpdatesRuleFields(r, vAllUpdatesRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAllUpdatesRule) {
		r.AllUpdatesRule = vAllUpdatesRule
	}
	return nil
}
func postReadExtractBudgetBudgetFilterFields(r *Budget, o *BudgetBudgetFilter) error {
	vCustomPeriod := o.CustomPeriod
	if vCustomPeriod == nil {
		// note: explicitly not the empty object.
		vCustomPeriod = &BudgetBudgetFilterCustomPeriod{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodFields(r, vCustomPeriod); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCustomPeriod) {
		o.CustomPeriod = vCustomPeriod
	}
	return nil
}
func postReadExtractBudgetBudgetFilterLabelsFields(r *Budget, o *BudgetBudgetFilterLabels) error {
	return nil
}
func postReadExtractBudgetBudgetFilterCustomPeriodFields(r *Budget, o *BudgetBudgetFilterCustomPeriod) error {
	vStartDate := o.StartDate
	if vStartDate == nil {
		// note: explicitly not the empty object.
		vStartDate = &BudgetBudgetFilterCustomPeriodStartDate{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodStartDateFields(r, vStartDate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStartDate) {
		o.StartDate = vStartDate
	}
	vEndDate := o.EndDate
	if vEndDate == nil {
		// note: explicitly not the empty object.
		vEndDate = &BudgetBudgetFilterCustomPeriodEndDate{}
	}
	if err := extractBudgetBudgetFilterCustomPeriodEndDateFields(r, vEndDate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEndDate) {
		o.EndDate = vEndDate
	}
	return nil
}
func postReadExtractBudgetBudgetFilterCustomPeriodStartDateFields(r *Budget, o *BudgetBudgetFilterCustomPeriodStartDate) error {
	return nil
}
func postReadExtractBudgetBudgetFilterCustomPeriodEndDateFields(r *Budget, o *BudgetBudgetFilterCustomPeriodEndDate) error {
	return nil
}
func postReadExtractBudgetAmountFields(r *Budget, o *BudgetAmount) error {
	vSpecifiedAmount := o.SpecifiedAmount
	if vSpecifiedAmount == nil {
		// note: explicitly not the empty object.
		vSpecifiedAmount = &BudgetAmountSpecifiedAmount{}
	}
	if err := extractBudgetAmountSpecifiedAmountFields(r, vSpecifiedAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSpecifiedAmount) {
		o.SpecifiedAmount = vSpecifiedAmount
	}
	vLastPeriodAmount := o.LastPeriodAmount
	if vLastPeriodAmount == nil {
		// note: explicitly not the empty object.
		vLastPeriodAmount = &BudgetAmountLastPeriodAmount{}
	}
	if err := extractBudgetAmountLastPeriodAmountFields(r, vLastPeriodAmount); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLastPeriodAmount) {
		o.LastPeriodAmount = vLastPeriodAmount
	}
	return nil
}
func postReadExtractBudgetAmountSpecifiedAmountFields(r *Budget, o *BudgetAmountSpecifiedAmount) error {
	return nil
}
func postReadExtractBudgetAmountLastPeriodAmountFields(r *Budget, o *BudgetAmountLastPeriodAmount) error {
	return nil
}
func postReadExtractBudgetThresholdRulesFields(r *Budget, o *BudgetThresholdRules) error {
	return nil
}
func postReadExtractBudgetAllUpdatesRuleFields(r *Budget, o *BudgetAllUpdatesRule) error {
	return nil
}
