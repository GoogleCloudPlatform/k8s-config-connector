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

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Policy) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"KubernetesNamespaceAdmissionRules", "KubernetesServiceAccountAdmissionRules", "IstioServiceIdentityAdmissionRules", "ClusterAdmissionRules"}, r.KubernetesNamespaceAdmissionRules, r.KubernetesServiceAccountAdmissionRules, r.IstioServiceIdentityAdmissionRules, r.ClusterAdmissionRules); err != nil {
		return err
	}
	if err := dcl.Required(r, "defaultAdmissionRule"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultAdmissionRule) {
		if err := r.DefaultAdmissionRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *PolicyAdmissionWhitelistPatterns) validate() error {
	return nil
}
func (r *PolicyClusterAdmissionRules) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *PolicyKubernetesNamespaceAdmissionRules) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *PolicyKubernetesServiceAccountAdmissionRules) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *PolicyIstioServiceIdentityAdmissionRules) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *PolicyDefaultAdmissionRule) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *Policy) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://binaryauthorization.googleapis.com/v1", params)
}

func (r *Policy) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/policy", nr.basePath(), userBasePath, params), nil
}

func (r *Policy) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
	}
	return dcl.URL("projects/{{project}}/policy:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Policy) SetPolicyVerb() string {
	return "POST"
}

func (r *Policy) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
	}
	return dcl.URL("projects/{{project}}/policy:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Policy) IAMPolicyVersion() int {
	return 0
}

// policyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type policyApiOperation interface {
	do(context.Context, *Policy, *Client) error
}

// newUpdatePolicyUpdatePolicyRequest creates a request for an
// Policy resource's UpdatePolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdatePolicyUpdatePolicyRequest(ctx context.Context, f *Policy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandPolicyAdmissionWhitelistPatternsSlice(c, f.AdmissionWhitelistPatterns, res); err != nil {
		return nil, fmt.Errorf("error expanding AdmissionWhitelistPatterns into admissionWhitelistPatterns: %w", err)
	} else if v != nil {
		req["admissionWhitelistPatterns"] = v
	}
	if v, err := expandPolicyClusterAdmissionRulesMap(c, f.ClusterAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding ClusterAdmissionRules into clusterAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clusterAdmissionRules"] = v
	}
	if v, err := expandPolicyKubernetesNamespaceAdmissionRulesMap(c, f.KubernetesNamespaceAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesNamespaceAdmissionRules into kubernetesNamespaceAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["kubernetesNamespaceAdmissionRules"] = v
	}
	if v, err := expandPolicyKubernetesServiceAccountAdmissionRulesMap(c, f.KubernetesServiceAccountAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesServiceAccountAdmissionRules into kubernetesServiceAccountAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["kubernetesServiceAccountAdmissionRules"] = v
	}
	if v, err := expandPolicyIstioServiceIdentityAdmissionRulesMap(c, f.IstioServiceIdentityAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding IstioServiceIdentityAdmissionRules into istioServiceIdentityAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["istioServiceIdentityAdmissionRules"] = v
	}
	if v, err := expandPolicyDefaultAdmissionRule(c, f.DefaultAdmissionRule, res); err != nil {
		return nil, fmt.Errorf("error expanding DefaultAdmissionRule into defaultAdmissionRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["defaultAdmissionRule"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.GlobalPolicyEvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		req["globalPolicyEvaluationMode"] = v
	}
	return req, nil
}

// marshalUpdatePolicyUpdatePolicyRequest converts the update into
// the final JSON request body.
func marshalUpdatePolicyUpdatePolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updatePolicyUpdatePolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updatePolicyUpdatePolicyOperation) do(ctx context.Context, r *Policy, c *Client) error {
	_, err := c.GetPolicy(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdatePolicy")
	if err != nil {
		return err
	}

	req, err := newUpdatePolicyUpdatePolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdatePolicyUpdatePolicyRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createPolicyOperation struct {
	response map[string]interface{}
}

func (op *createPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getPolicyRaw(ctx context.Context, r *Policy) ([]byte, error) {

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

func (c *Client) policyDiffsForRawDesired(ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (initial, desired *Policy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Policy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Policy); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Policy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetPolicy(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Policy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Policy resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Policy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Policy: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Policy: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractPolicyFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizePolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Policy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Policy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizePolicyInitialState(rawInitial, rawDesired *Policy) (*Policy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.KubernetesNamespaceAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.KubernetesNamespaceAdmissionRules = map[string]PolicyKubernetesNamespaceAdmissionRules{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.KubernetesServiceAccountAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.KubernetesServiceAccountAdmissionRules = map[string]PolicyKubernetesServiceAccountAdmissionRules{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.IstioServiceIdentityAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.IstioServiceIdentityAdmissionRules = map[string]PolicyIstioServiceIdentityAdmissionRules{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.ClusterAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules) {
			rawInitial.ClusterAdmissionRules = map[string]PolicyClusterAdmissionRules{}
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

func canonicalizePolicyDesiredState(rawDesired, rawInitial *Policy, opts ...dcl.ApplyOption) (*Policy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultAdmissionRule = canonicalizePolicyDefaultAdmissionRule(rawDesired.DefaultAdmissionRule, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Policy{}
	canonicalDesired.AdmissionWhitelistPatterns = canonicalizePolicyAdmissionWhitelistPatternsSlice(rawDesired.AdmissionWhitelistPatterns, rawInitial.AdmissionWhitelistPatterns, opts...)
	if dcl.IsZeroValue(rawDesired.ClusterAdmissionRules) || (dcl.IsEmptyValueIndirect(rawDesired.ClusterAdmissionRules) && dcl.IsEmptyValueIndirect(rawInitial.ClusterAdmissionRules)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ClusterAdmissionRules = rawInitial.ClusterAdmissionRules
	} else {
		canonicalDesired.ClusterAdmissionRules = rawDesired.ClusterAdmissionRules
	}
	if dcl.IsZeroValue(rawDesired.KubernetesNamespaceAdmissionRules) || (dcl.IsEmptyValueIndirect(rawDesired.KubernetesNamespaceAdmissionRules) && dcl.IsEmptyValueIndirect(rawInitial.KubernetesNamespaceAdmissionRules)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.KubernetesNamespaceAdmissionRules = rawInitial.KubernetesNamespaceAdmissionRules
	} else {
		canonicalDesired.KubernetesNamespaceAdmissionRules = rawDesired.KubernetesNamespaceAdmissionRules
	}
	if dcl.IsZeroValue(rawDesired.KubernetesServiceAccountAdmissionRules) || (dcl.IsEmptyValueIndirect(rawDesired.KubernetesServiceAccountAdmissionRules) && dcl.IsEmptyValueIndirect(rawInitial.KubernetesServiceAccountAdmissionRules)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.KubernetesServiceAccountAdmissionRules = rawInitial.KubernetesServiceAccountAdmissionRules
	} else {
		canonicalDesired.KubernetesServiceAccountAdmissionRules = rawDesired.KubernetesServiceAccountAdmissionRules
	}
	if canonicalizePolicyISIAR(rawDesired.IstioServiceIdentityAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules) {
		canonicalDesired.IstioServiceIdentityAdmissionRules = rawInitial.IstioServiceIdentityAdmissionRules
	} else {
		canonicalDesired.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
	}
	canonicalDesired.DefaultAdmissionRule = canonicalizePolicyDefaultAdmissionRule(rawDesired.DefaultAdmissionRule, rawInitial.DefaultAdmissionRule, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.GlobalPolicyEvaluationMode) || (dcl.IsEmptyValueIndirect(rawDesired.GlobalPolicyEvaluationMode) && dcl.IsEmptyValueIndirect(rawInitial.GlobalPolicyEvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.GlobalPolicyEvaluationMode = rawInitial.GlobalPolicyEvaluationMode
	} else {
		canonicalDesired.GlobalPolicyEvaluationMode = rawDesired.GlobalPolicyEvaluationMode
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	if canonicalDesired.KubernetesNamespaceAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules, rawDesired.ClusterAdmissionRules) {
			canonicalDesired.KubernetesNamespaceAdmissionRules = map[string]PolicyKubernetesNamespaceAdmissionRules{}
		}
	}

	if canonicalDesired.KubernetesServiceAccountAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules, rawDesired.ClusterAdmissionRules) {
			canonicalDesired.KubernetesServiceAccountAdmissionRules = map[string]PolicyKubernetesServiceAccountAdmissionRules{}
		}
	}

	if canonicalDesired.IstioServiceIdentityAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.ClusterAdmissionRules) {
			canonicalDesired.IstioServiceIdentityAdmissionRules = map[string]PolicyIstioServiceIdentityAdmissionRules{}
		}
	}

	if canonicalDesired.ClusterAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules) {
			canonicalDesired.ClusterAdmissionRules = map[string]PolicyClusterAdmissionRules{}
		}
	}

	return canonicalDesired, nil
}

func canonicalizePolicyNewState(c *Client, rawNew, rawDesired *Policy) (*Policy, error) {

	if dcl.IsEmptyValueIndirect(rawNew.AdmissionWhitelistPatterns) && dcl.IsEmptyValueIndirect(rawDesired.AdmissionWhitelistPatterns) {
		rawNew.AdmissionWhitelistPatterns = rawDesired.AdmissionWhitelistPatterns
	} else {
		rawNew.AdmissionWhitelistPatterns = canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c, rawDesired.AdmissionWhitelistPatterns, rawNew.AdmissionWhitelistPatterns)
	}

	if dcl.IsEmptyValueIndirect(rawNew.ClusterAdmissionRules) && dcl.IsEmptyValueIndirect(rawDesired.ClusterAdmissionRules) {
		rawNew.ClusterAdmissionRules = rawDesired.ClusterAdmissionRules
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.KubernetesNamespaceAdmissionRules) && dcl.IsEmptyValueIndirect(rawDesired.KubernetesNamespaceAdmissionRules) {
		rawNew.KubernetesNamespaceAdmissionRules = rawDesired.KubernetesNamespaceAdmissionRules
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.KubernetesServiceAccountAdmissionRules) && dcl.IsEmptyValueIndirect(rawDesired.KubernetesServiceAccountAdmissionRules) {
		rawNew.KubernetesServiceAccountAdmissionRules = rawDesired.KubernetesServiceAccountAdmissionRules
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IstioServiceIdentityAdmissionRules) && dcl.IsEmptyValueIndirect(rawDesired.IstioServiceIdentityAdmissionRules) {
		rawNew.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
	} else {
		if canonicalizePolicyISIAR(rawDesired.IstioServiceIdentityAdmissionRules, rawNew.IstioServiceIdentityAdmissionRules) {
			rawNew.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultAdmissionRule) && dcl.IsEmptyValueIndirect(rawDesired.DefaultAdmissionRule) {
		rawNew.DefaultAdmissionRule = rawDesired.DefaultAdmissionRule
	} else {
		rawNew.DefaultAdmissionRule = canonicalizeNewPolicyDefaultAdmissionRule(c, rawDesired.DefaultAdmissionRule, rawNew.DefaultAdmissionRule)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.GlobalPolicyEvaluationMode) && dcl.IsEmptyValueIndirect(rawDesired.GlobalPolicyEvaluationMode) {
		rawNew.GlobalPolicyEvaluationMode = rawDesired.GlobalPolicyEvaluationMode
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.UpdateTime) && dcl.IsEmptyValueIndirect(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	return rawNew, nil
}

func canonicalizePolicyAdmissionWhitelistPatterns(des, initial *PolicyAdmissionWhitelistPatterns, opts ...dcl.ApplyOption) *PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyAdmissionWhitelistPatterns{}

	if dcl.StringCanonicalize(des.NamePattern, initial.NamePattern) || dcl.IsZeroValue(des.NamePattern) {
		cDes.NamePattern = initial.NamePattern
	} else {
		cDes.NamePattern = des.NamePattern
	}

	return cDes
}

func canonicalizePolicyAdmissionWhitelistPatternsSlice(des, initial []PolicyAdmissionWhitelistPatterns, opts ...dcl.ApplyOption) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyAdmissionWhitelistPatterns, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyAdmissionWhitelistPatterns(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyAdmissionWhitelistPatterns, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyAdmissionWhitelistPatterns(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyAdmissionWhitelistPatterns(c *Client, des, nw *PolicyAdmissionWhitelistPatterns) *PolicyAdmissionWhitelistPatterns {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyAdmissionWhitelistPatterns while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.NamePattern, nw.NamePattern) {
		nw.NamePattern = des.NamePattern
	}

	return nw
}

func canonicalizeNewPolicyAdmissionWhitelistPatternsSet(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyAdmissionWhitelistPatterns
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyAdmissionWhitelistPatternsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyAdmissionWhitelistPatterns(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyAdmissionWhitelistPatterns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyAdmissionWhitelistPatterns(c, &d, &n))
	}

	return items
}

func canonicalizePolicyClusterAdmissionRules(des, initial *PolicyClusterAdmissionRules, opts ...dcl.ApplyOption) *PolicyClusterAdmissionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyClusterAdmissionRules{}

	if dcl.IsZeroValue(des.EvaluationMode) || (dcl.IsEmptyValueIndirect(des.EvaluationMode) && dcl.IsEmptyValueIndirect(initial.EvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) || (dcl.IsEmptyValueIndirect(des.EnforcementMode) && dcl.IsEmptyValueIndirect(initial.EnforcementMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyClusterAdmissionRulesSlice(des, initial []PolicyClusterAdmissionRules, opts ...dcl.ApplyOption) []PolicyClusterAdmissionRules {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyClusterAdmissionRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyClusterAdmissionRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyClusterAdmissionRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyClusterAdmissionRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyClusterAdmissionRules(c *Client, des, nw *PolicyClusterAdmissionRules) *PolicyClusterAdmissionRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyClusterAdmissionRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyClusterAdmissionRulesSet(c *Client, des, nw []PolicyClusterAdmissionRules) []PolicyClusterAdmissionRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyClusterAdmissionRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyClusterAdmissionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyClusterAdmissionRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyClusterAdmissionRulesSlice(c *Client, des, nw []PolicyClusterAdmissionRules) []PolicyClusterAdmissionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyClusterAdmissionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyClusterAdmissionRules(c, &d, &n))
	}

	return items
}

func canonicalizePolicyKubernetesNamespaceAdmissionRules(des, initial *PolicyKubernetesNamespaceAdmissionRules, opts ...dcl.ApplyOption) *PolicyKubernetesNamespaceAdmissionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyKubernetesNamespaceAdmissionRules{}

	if dcl.IsZeroValue(des.EvaluationMode) || (dcl.IsEmptyValueIndirect(des.EvaluationMode) && dcl.IsEmptyValueIndirect(initial.EvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) || (dcl.IsEmptyValueIndirect(des.EnforcementMode) && dcl.IsEmptyValueIndirect(initial.EnforcementMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyKubernetesNamespaceAdmissionRulesSlice(des, initial []PolicyKubernetesNamespaceAdmissionRules, opts ...dcl.ApplyOption) []PolicyKubernetesNamespaceAdmissionRules {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyKubernetesNamespaceAdmissionRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyKubernetesNamespaceAdmissionRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyKubernetesNamespaceAdmissionRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyKubernetesNamespaceAdmissionRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyKubernetesNamespaceAdmissionRules(c *Client, des, nw *PolicyKubernetesNamespaceAdmissionRules) *PolicyKubernetesNamespaceAdmissionRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyKubernetesNamespaceAdmissionRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyKubernetesNamespaceAdmissionRulesSet(c *Client, des, nw []PolicyKubernetesNamespaceAdmissionRules) []PolicyKubernetesNamespaceAdmissionRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyKubernetesNamespaceAdmissionRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyKubernetesNamespaceAdmissionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyKubernetesNamespaceAdmissionRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyKubernetesNamespaceAdmissionRulesSlice(c *Client, des, nw []PolicyKubernetesNamespaceAdmissionRules) []PolicyKubernetesNamespaceAdmissionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyKubernetesNamespaceAdmissionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyKubernetesNamespaceAdmissionRules(c, &d, &n))
	}

	return items
}

func canonicalizePolicyKubernetesServiceAccountAdmissionRules(des, initial *PolicyKubernetesServiceAccountAdmissionRules, opts ...dcl.ApplyOption) *PolicyKubernetesServiceAccountAdmissionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyKubernetesServiceAccountAdmissionRules{}

	if dcl.IsZeroValue(des.EvaluationMode) || (dcl.IsEmptyValueIndirect(des.EvaluationMode) && dcl.IsEmptyValueIndirect(initial.EvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) || (dcl.IsEmptyValueIndirect(des.EnforcementMode) && dcl.IsEmptyValueIndirect(initial.EnforcementMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyKubernetesServiceAccountAdmissionRulesSlice(des, initial []PolicyKubernetesServiceAccountAdmissionRules, opts ...dcl.ApplyOption) []PolicyKubernetesServiceAccountAdmissionRules {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyKubernetesServiceAccountAdmissionRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyKubernetesServiceAccountAdmissionRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyKubernetesServiceAccountAdmissionRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyKubernetesServiceAccountAdmissionRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyKubernetesServiceAccountAdmissionRules(c *Client, des, nw *PolicyKubernetesServiceAccountAdmissionRules) *PolicyKubernetesServiceAccountAdmissionRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyKubernetesServiceAccountAdmissionRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyKubernetesServiceAccountAdmissionRulesSet(c *Client, des, nw []PolicyKubernetesServiceAccountAdmissionRules) []PolicyKubernetesServiceAccountAdmissionRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyKubernetesServiceAccountAdmissionRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyKubernetesServiceAccountAdmissionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyKubernetesServiceAccountAdmissionRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyKubernetesServiceAccountAdmissionRulesSlice(c *Client, des, nw []PolicyKubernetesServiceAccountAdmissionRules) []PolicyKubernetesServiceAccountAdmissionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyKubernetesServiceAccountAdmissionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyKubernetesServiceAccountAdmissionRules(c, &d, &n))
	}

	return items
}

func canonicalizePolicyIstioServiceIdentityAdmissionRules(des, initial *PolicyIstioServiceIdentityAdmissionRules, opts ...dcl.ApplyOption) *PolicyIstioServiceIdentityAdmissionRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyIstioServiceIdentityAdmissionRules{}

	if dcl.IsZeroValue(des.EvaluationMode) || (dcl.IsEmptyValueIndirect(des.EvaluationMode) && dcl.IsEmptyValueIndirect(initial.EvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) || (dcl.IsEmptyValueIndirect(des.EnforcementMode) && dcl.IsEmptyValueIndirect(initial.EnforcementMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyIstioServiceIdentityAdmissionRulesSlice(des, initial []PolicyIstioServiceIdentityAdmissionRules, opts ...dcl.ApplyOption) []PolicyIstioServiceIdentityAdmissionRules {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyIstioServiceIdentityAdmissionRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyIstioServiceIdentityAdmissionRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyIstioServiceIdentityAdmissionRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyIstioServiceIdentityAdmissionRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyIstioServiceIdentityAdmissionRules(c *Client, des, nw *PolicyIstioServiceIdentityAdmissionRules) *PolicyIstioServiceIdentityAdmissionRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyIstioServiceIdentityAdmissionRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyIstioServiceIdentityAdmissionRulesSet(c *Client, des, nw []PolicyIstioServiceIdentityAdmissionRules) []PolicyIstioServiceIdentityAdmissionRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyIstioServiceIdentityAdmissionRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyIstioServiceIdentityAdmissionRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyIstioServiceIdentityAdmissionRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyIstioServiceIdentityAdmissionRulesSlice(c *Client, des, nw []PolicyIstioServiceIdentityAdmissionRules) []PolicyIstioServiceIdentityAdmissionRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyIstioServiceIdentityAdmissionRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyIstioServiceIdentityAdmissionRules(c, &d, &n))
	}

	return items
}

func canonicalizePolicyDefaultAdmissionRule(des, initial *PolicyDefaultAdmissionRule, opts ...dcl.ApplyOption) *PolicyDefaultAdmissionRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyDefaultAdmissionRule{}

	if dcl.IsZeroValue(des.EvaluationMode) || (dcl.IsEmptyValueIndirect(des.EvaluationMode) && dcl.IsEmptyValueIndirect(initial.EvaluationMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) || (dcl.IsEmptyValueIndirect(des.EnforcementMode) && dcl.IsEmptyValueIndirect(initial.EnforcementMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyDefaultAdmissionRuleSlice(des, initial []PolicyDefaultAdmissionRule, opts ...dcl.ApplyOption) []PolicyDefaultAdmissionRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyDefaultAdmissionRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyDefaultAdmissionRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyDefaultAdmissionRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyDefaultAdmissionRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyDefaultAdmissionRule(c *Client, des, nw *PolicyDefaultAdmissionRule) *PolicyDefaultAdmissionRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyDefaultAdmissionRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyDefaultAdmissionRuleSet(c *Client, des, nw []PolicyDefaultAdmissionRule) []PolicyDefaultAdmissionRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []PolicyDefaultAdmissionRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := comparePolicyDefaultAdmissionRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewPolicyDefaultAdmissionRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewPolicyDefaultAdmissionRuleSlice(c *Client, des, nw []PolicyDefaultAdmissionRule) []PolicyDefaultAdmissionRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyDefaultAdmissionRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyDefaultAdmissionRule(c, &d, &n))
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
func diffPolicy(c *Client, desired, actual *Policy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.AdmissionWhitelistPatterns, actual.AdmissionWhitelistPatterns, dcl.DiffInfo{ObjectFunction: comparePolicyAdmissionWhitelistPatternsNewStyle, EmptyObject: EmptyPolicyAdmissionWhitelistPatterns, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("AdmissionWhitelistPatterns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterAdmissionRules, actual.ClusterAdmissionRules, dcl.DiffInfo{ObjectFunction: comparePolicyClusterAdmissionRulesNewStyle, EmptyObject: EmptyPolicyClusterAdmissionRules, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("ClusterAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesNamespaceAdmissionRules, actual.KubernetesNamespaceAdmissionRules, dcl.DiffInfo{ObjectFunction: comparePolicyKubernetesNamespaceAdmissionRulesNewStyle, EmptyObject: EmptyPolicyKubernetesNamespaceAdmissionRules, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("KubernetesNamespaceAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesServiceAccountAdmissionRules, actual.KubernetesServiceAccountAdmissionRules, dcl.DiffInfo{ObjectFunction: comparePolicyKubernetesServiceAccountAdmissionRulesNewStyle, EmptyObject: EmptyPolicyKubernetesServiceAccountAdmissionRules, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("KubernetesServiceAccountAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IstioServiceIdentityAdmissionRules, actual.IstioServiceIdentityAdmissionRules, dcl.DiffInfo{ObjectFunction: comparePolicyIstioServiceIdentityAdmissionRulesNewStyle, EmptyObject: EmptyPolicyIstioServiceIdentityAdmissionRules, CustomDiff: canonicalizePolicyISIAR, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("IstioServiceIdentityAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultAdmissionRule, actual.DefaultAdmissionRule, dcl.DiffInfo{ObjectFunction: comparePolicyDefaultAdmissionRuleNewStyle, EmptyObject: EmptyPolicyDefaultAdmissionRule, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("DefaultAdmissionRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GlobalPolicyEvaluationMode, actual.GlobalPolicyEvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("GlobalPolicyEvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
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
func comparePolicyAdmissionWhitelistPatternsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		desiredNotPointer, ok := d.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns or *PolicyAdmissionWhitelistPatterns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		actualNotPointer, ok := a.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NamePattern, actual.NamePattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("NamePattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyClusterAdmissionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyClusterAdmissionRules)
	if !ok {
		desiredNotPointer, ok := d.(PolicyClusterAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyClusterAdmissionRules or *PolicyClusterAdmissionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyClusterAdmissionRules)
	if !ok {
		actualNotPointer, ok := a.(PolicyClusterAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyClusterAdmissionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyKubernetesNamespaceAdmissionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyKubernetesNamespaceAdmissionRules)
	if !ok {
		desiredNotPointer, ok := d.(PolicyKubernetesNamespaceAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyKubernetesNamespaceAdmissionRules or *PolicyKubernetesNamespaceAdmissionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyKubernetesNamespaceAdmissionRules)
	if !ok {
		actualNotPointer, ok := a.(PolicyKubernetesNamespaceAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyKubernetesNamespaceAdmissionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyKubernetesServiceAccountAdmissionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyKubernetesServiceAccountAdmissionRules)
	if !ok {
		desiredNotPointer, ok := d.(PolicyKubernetesServiceAccountAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyKubernetesServiceAccountAdmissionRules or *PolicyKubernetesServiceAccountAdmissionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyKubernetesServiceAccountAdmissionRules)
	if !ok {
		actualNotPointer, ok := a.(PolicyKubernetesServiceAccountAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyKubernetesServiceAccountAdmissionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyIstioServiceIdentityAdmissionRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyIstioServiceIdentityAdmissionRules)
	if !ok {
		desiredNotPointer, ok := d.(PolicyIstioServiceIdentityAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyIstioServiceIdentityAdmissionRules or *PolicyIstioServiceIdentityAdmissionRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyIstioServiceIdentityAdmissionRules)
	if !ok {
		actualNotPointer, ok := a.(PolicyIstioServiceIdentityAdmissionRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyIstioServiceIdentityAdmissionRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyDefaultAdmissionRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyDefaultAdmissionRule)
	if !ok {
		desiredNotPointer, ok := d.(PolicyDefaultAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyDefaultAdmissionRule or *PolicyDefaultAdmissionRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyDefaultAdmissionRule)
	if !ok {
		actualNotPointer, ok := a.(PolicyDefaultAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyDefaultAdmissionRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
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
func (r *Policy) urlNormalized() *Policy {
	normalized := dcl.Copy(*r).(Policy)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Policy) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdatePolicy" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
		}
		return dcl.URL("projects/{{project}}/policy", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Policy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Policy) marshal(c *Client) ([]byte, error) {
	m, err := expandPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Policy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalPolicy decodes JSON responses into the Policy resource schema.
func unmarshalPolicy(b []byte, c *Client, res *Policy) (*Policy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapPolicy(m, c, res)
}

func unmarshalMapPolicy(m map[string]interface{}, c *Client, res *Policy) (*Policy, error) {

	flattened := flattenPolicy(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandPolicy expands Policy into a JSON request object.
func expandPolicy(c *Client, f *Policy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := expandPolicyAdmissionWhitelistPatternsSlice(c, f.AdmissionWhitelistPatterns, res); err != nil {
		return nil, fmt.Errorf("error expanding AdmissionWhitelistPatterns into admissionWhitelistPatterns: %w", err)
	} else if v != nil {
		m["admissionWhitelistPatterns"] = v
	}
	if v, err := expandPolicyClusterAdmissionRulesMap(c, f.ClusterAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding ClusterAdmissionRules into clusterAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["clusterAdmissionRules"] = v
	}
	if v, err := expandPolicyKubernetesNamespaceAdmissionRulesMap(c, f.KubernetesNamespaceAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesNamespaceAdmissionRules into kubernetesNamespaceAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesNamespaceAdmissionRules"] = v
	}
	if v, err := expandPolicyKubernetesServiceAccountAdmissionRulesMap(c, f.KubernetesServiceAccountAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding KubernetesServiceAccountAdmissionRules into kubernetesServiceAccountAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["kubernetesServiceAccountAdmissionRules"] = v
	}
	if v, err := expandPolicyIstioServiceIdentityAdmissionRulesMap(c, f.IstioServiceIdentityAdmissionRules, res); err != nil {
		return nil, fmt.Errorf("error expanding IstioServiceIdentityAdmissionRules into istioServiceIdentityAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["istioServiceIdentityAdmissionRules"] = v
	}
	if v, err := expandPolicyDefaultAdmissionRule(c, f.DefaultAdmissionRule, res); err != nil {
		return nil, fmt.Errorf("error expanding DefaultAdmissionRule into defaultAdmissionRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultAdmissionRule"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.GlobalPolicyEvaluationMode; dcl.ValueShouldBeSent(v) {
		m["globalPolicyEvaluationMode"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenPolicy flattens Policy from a JSON request object into the
// Policy type.
func flattenPolicy(c *Client, i interface{}, res *Policy) *Policy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Policy{}
	resultRes.AdmissionWhitelistPatterns = flattenPolicyAdmissionWhitelistPatternsSlice(c, m["admissionWhitelistPatterns"], res)
	resultRes.ClusterAdmissionRules = flattenPolicyClusterAdmissionRulesMap(c, m["clusterAdmissionRules"], res)
	resultRes.KubernetesNamespaceAdmissionRules = flattenPolicyKubernetesNamespaceAdmissionRulesMap(c, m["kubernetesNamespaceAdmissionRules"], res)
	resultRes.KubernetesServiceAccountAdmissionRules = flattenPolicyKubernetesServiceAccountAdmissionRulesMap(c, m["kubernetesServiceAccountAdmissionRules"], res)
	resultRes.IstioServiceIdentityAdmissionRules = flattenPolicyIstioServiceIdentityAdmissionRulesMap(c, m["istioServiceIdentityAdmissionRules"], res)
	resultRes.DefaultAdmissionRule = flattenPolicyDefaultAdmissionRule(c, m["defaultAdmissionRule"], res)
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.GlobalPolicyEvaluationMode = flattenPolicyGlobalPolicyEvaluationModeEnum(m["globalPolicyEvaluationMode"])
	resultRes.SelfLink = dcl.FlattenString(m["name"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])

	return resultRes
}

// expandPolicyAdmissionWhitelistPatternsMap expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsMap(c *Client, f map[string]PolicyAdmissionWhitelistPatterns, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyAdmissionWhitelistPatternsSlice expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsSlice(c *Client, f []PolicyAdmissionWhitelistPatterns, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyAdmissionWhitelistPatternsMap flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsMap(c *Client, i interface{}, res *Policy) map[string]PolicyAdmissionWhitelistPatterns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	items := make(map[string]PolicyAdmissionWhitelistPatterns)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyAdmissionWhitelistPatternsSlice flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsSlice(c *Client, i interface{}, res *Policy) []PolicyAdmissionWhitelistPatterns {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	items := make([]PolicyAdmissionWhitelistPatterns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyAdmissionWhitelistPatterns expands an instance of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatterns(c *Client, f *PolicyAdmissionWhitelistPatterns, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NamePattern; !dcl.IsEmptyValueIndirect(v) {
		m["namePattern"] = v
	}

	return m, nil
}

// flattenPolicyAdmissionWhitelistPatterns flattens an instance of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatterns(c *Client, i interface{}, res *Policy) *PolicyAdmissionWhitelistPatterns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyAdmissionWhitelistPatterns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyAdmissionWhitelistPatterns
	}
	r.NamePattern = dcl.FlattenString(m["namePattern"])

	return r
}

// expandPolicyClusterAdmissionRulesMap expands the contents of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRulesMap(c *Client, f map[string]PolicyClusterAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyClusterAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyClusterAdmissionRulesSlice expands the contents of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRulesSlice(c *Client, f []PolicyClusterAdmissionRules, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyClusterAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyClusterAdmissionRulesMap flattens the contents of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesMap(c *Client, i interface{}, res *Policy) map[string]PolicyClusterAdmissionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyClusterAdmissionRules{}
	}

	if len(a) == 0 {
		return map[string]PolicyClusterAdmissionRules{}
	}

	items := make(map[string]PolicyClusterAdmissionRules)
	for k, item := range a {
		items[k] = *flattenPolicyClusterAdmissionRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyClusterAdmissionRulesSlice flattens the contents of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesSlice(c *Client, i interface{}, res *Policy) []PolicyClusterAdmissionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRules{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRules{}
	}

	items := make([]PolicyClusterAdmissionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyClusterAdmissionRules expands an instance of PolicyClusterAdmissionRules into a JSON
// request object.
func expandPolicyClusterAdmissionRules(c *Client, f *PolicyClusterAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyClusterAdmissionRules flattens an instance of PolicyClusterAdmissionRules from a JSON
// response object.
func flattenPolicyClusterAdmissionRules(c *Client, i interface{}, res *Policy) *PolicyClusterAdmissionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyClusterAdmissionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyClusterAdmissionRules
	}
	r.EvaluationMode = flattenPolicyClusterAdmissionRulesEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyClusterAdmissionRulesEnforcementModeEnum(m["enforcementMode"])

	return r
}

// expandPolicyKubernetesNamespaceAdmissionRulesMap expands the contents of PolicyKubernetesNamespaceAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesNamespaceAdmissionRulesMap(c *Client, f map[string]PolicyKubernetesNamespaceAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyKubernetesNamespaceAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyKubernetesNamespaceAdmissionRulesSlice expands the contents of PolicyKubernetesNamespaceAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesNamespaceAdmissionRulesSlice(c *Client, f []PolicyKubernetesNamespaceAdmissionRules, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyKubernetesNamespaceAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyKubernetesNamespaceAdmissionRulesMap flattens the contents of PolicyKubernetesNamespaceAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesNamespaceAdmissionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesNamespaceAdmissionRules{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesNamespaceAdmissionRules{}
	}

	items := make(map[string]PolicyKubernetesNamespaceAdmissionRules)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesNamespaceAdmissionRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyKubernetesNamespaceAdmissionRulesSlice flattens the contents of PolicyKubernetesNamespaceAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesNamespaceAdmissionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesNamespaceAdmissionRules{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesNamespaceAdmissionRules{}
	}

	items := make([]PolicyKubernetesNamespaceAdmissionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesNamespaceAdmissionRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyKubernetesNamespaceAdmissionRules expands an instance of PolicyKubernetesNamespaceAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesNamespaceAdmissionRules(c *Client, f *PolicyKubernetesNamespaceAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyKubernetesNamespaceAdmissionRules flattens an instance of PolicyKubernetesNamespaceAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRules(c *Client, i interface{}, res *Policy) *PolicyKubernetesNamespaceAdmissionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyKubernetesNamespaceAdmissionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyKubernetesNamespaceAdmissionRules
	}
	r.EvaluationMode = flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(m["enforcementMode"])

	return r
}

// expandPolicyKubernetesServiceAccountAdmissionRulesMap expands the contents of PolicyKubernetesServiceAccountAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesServiceAccountAdmissionRulesMap(c *Client, f map[string]PolicyKubernetesServiceAccountAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyKubernetesServiceAccountAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyKubernetesServiceAccountAdmissionRulesSlice expands the contents of PolicyKubernetesServiceAccountAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesServiceAccountAdmissionRulesSlice(c *Client, f []PolicyKubernetesServiceAccountAdmissionRules, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyKubernetesServiceAccountAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesMap flattens the contents of PolicyKubernetesServiceAccountAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesServiceAccountAdmissionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesServiceAccountAdmissionRules{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesServiceAccountAdmissionRules{}
	}

	items := make(map[string]PolicyKubernetesServiceAccountAdmissionRules)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesServiceAccountAdmissionRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesSlice flattens the contents of PolicyKubernetesServiceAccountAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesServiceAccountAdmissionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesServiceAccountAdmissionRules{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesServiceAccountAdmissionRules{}
	}

	items := make([]PolicyKubernetesServiceAccountAdmissionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesServiceAccountAdmissionRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyKubernetesServiceAccountAdmissionRules expands an instance of PolicyKubernetesServiceAccountAdmissionRules into a JSON
// request object.
func expandPolicyKubernetesServiceAccountAdmissionRules(c *Client, f *PolicyKubernetesServiceAccountAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyKubernetesServiceAccountAdmissionRules flattens an instance of PolicyKubernetesServiceAccountAdmissionRules from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRules(c *Client, i interface{}, res *Policy) *PolicyKubernetesServiceAccountAdmissionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyKubernetesServiceAccountAdmissionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyKubernetesServiceAccountAdmissionRules
	}
	r.EvaluationMode = flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(m["enforcementMode"])

	return r
}

// expandPolicyIstioServiceIdentityAdmissionRulesMap expands the contents of PolicyIstioServiceIdentityAdmissionRules into a JSON
// request object.
func expandPolicyIstioServiceIdentityAdmissionRulesMap(c *Client, f map[string]PolicyIstioServiceIdentityAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyIstioServiceIdentityAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyIstioServiceIdentityAdmissionRulesSlice expands the contents of PolicyIstioServiceIdentityAdmissionRules into a JSON
// request object.
func expandPolicyIstioServiceIdentityAdmissionRulesSlice(c *Client, f []PolicyIstioServiceIdentityAdmissionRules, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyIstioServiceIdentityAdmissionRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyIstioServiceIdentityAdmissionRulesMap flattens the contents of PolicyIstioServiceIdentityAdmissionRules from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesMap(c *Client, i interface{}, res *Policy) map[string]PolicyIstioServiceIdentityAdmissionRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyIstioServiceIdentityAdmissionRules{}
	}

	if len(a) == 0 {
		return map[string]PolicyIstioServiceIdentityAdmissionRules{}
	}

	items := make(map[string]PolicyIstioServiceIdentityAdmissionRules)
	for k, item := range a {
		items[k] = *flattenPolicyIstioServiceIdentityAdmissionRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyIstioServiceIdentityAdmissionRulesSlice flattens the contents of PolicyIstioServiceIdentityAdmissionRules from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesSlice(c *Client, i interface{}, res *Policy) []PolicyIstioServiceIdentityAdmissionRules {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyIstioServiceIdentityAdmissionRules{}
	}

	if len(a) == 0 {
		return []PolicyIstioServiceIdentityAdmissionRules{}
	}

	items := make([]PolicyIstioServiceIdentityAdmissionRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyIstioServiceIdentityAdmissionRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyIstioServiceIdentityAdmissionRules expands an instance of PolicyIstioServiceIdentityAdmissionRules into a JSON
// request object.
func expandPolicyIstioServiceIdentityAdmissionRules(c *Client, f *PolicyIstioServiceIdentityAdmissionRules, res *Policy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyIstioServiceIdentityAdmissionRules flattens an instance of PolicyIstioServiceIdentityAdmissionRules from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRules(c *Client, i interface{}, res *Policy) *PolicyIstioServiceIdentityAdmissionRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyIstioServiceIdentityAdmissionRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyIstioServiceIdentityAdmissionRules
	}
	r.EvaluationMode = flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(m["enforcementMode"])

	return r
}

// expandPolicyDefaultAdmissionRuleMap expands the contents of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRuleMap(c *Client, f map[string]PolicyDefaultAdmissionRule, res *Policy) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyDefaultAdmissionRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyDefaultAdmissionRuleSlice expands the contents of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRuleSlice(c *Client, f []PolicyDefaultAdmissionRule, res *Policy) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyDefaultAdmissionRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyDefaultAdmissionRuleMap flattens the contents of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleMap(c *Client, i interface{}, res *Policy) map[string]PolicyDefaultAdmissionRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyDefaultAdmissionRule{}
	}

	if len(a) == 0 {
		return map[string]PolicyDefaultAdmissionRule{}
	}

	items := make(map[string]PolicyDefaultAdmissionRule)
	for k, item := range a {
		items[k] = *flattenPolicyDefaultAdmissionRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleSlice flattens the contents of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleSlice(c *Client, i interface{}, res *Policy) []PolicyDefaultAdmissionRule {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRule{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRule{}
	}

	items := make([]PolicyDefaultAdmissionRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandPolicyDefaultAdmissionRule expands an instance of PolicyDefaultAdmissionRule into a JSON
// request object.
func expandPolicyDefaultAdmissionRule(c *Client, f *PolicyDefaultAdmissionRule, res *Policy) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyDefaultAdmissionRule flattens an instance of PolicyDefaultAdmissionRule from a JSON
// response object.
func flattenPolicyDefaultAdmissionRule(c *Client, i interface{}, res *Policy) *PolicyDefaultAdmissionRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyDefaultAdmissionRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyDefaultAdmissionRule
	}
	r.EvaluationMode = flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(m["enforcementMode"])

	return r
}

// flattenPolicyClusterAdmissionRulesEvaluationModeEnumMap flattens the contents of PolicyClusterAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyClusterAdmissionRulesEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	items := make(map[string]PolicyClusterAdmissionRulesEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyClusterAdmissionRulesEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEvaluationModeEnumSlice flattens the contents of PolicyClusterAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyClusterAdmissionRulesEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRulesEvaluationModeEnum{}
	}

	items := make([]PolicyClusterAdmissionRulesEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRulesEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyClusterAdmissionRulesEvaluationModeEnum with the same value as that string.
func flattenPolicyClusterAdmissionRulesEvaluationModeEnum(i interface{}) *PolicyClusterAdmissionRulesEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyClusterAdmissionRulesEvaluationModeEnumRef(s)
}

// flattenPolicyClusterAdmissionRulesEnforcementModeEnumMap flattens the contents of PolicyClusterAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEnforcementModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyClusterAdmissionRulesEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	items := make(map[string]PolicyClusterAdmissionRulesEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyClusterAdmissionRulesEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEnforcementModeEnumSlice flattens the contents of PolicyClusterAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyClusterAdmissionRulesEnforcementModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyClusterAdmissionRulesEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyClusterAdmissionRulesEnforcementModeEnum{}
	}

	items := make([]PolicyClusterAdmissionRulesEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyClusterAdmissionRulesEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyClusterAdmissionRulesEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyClusterAdmissionRulesEnforcementModeEnum with the same value as that string.
func flattenPolicyClusterAdmissionRulesEnforcementModeEnum(i interface{}) *PolicyClusterAdmissionRulesEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyClusterAdmissionRulesEnforcementModeEnumRef(s)
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumMap flattens the contents of PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum{}
	}

	items := make(map[string]PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumSlice flattens the contents of PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum{}
	}

	items := make([]PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum with the same value as that string.
func flattenPolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum(i interface{}) *PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyKubernetesNamespaceAdmissionRulesEvaluationModeEnumRef(s)
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumMap flattens the contents of PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum{}
	}

	items := make(map[string]PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumSlice flattens the contents of PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum{}
	}

	items := make([]PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum with the same value as that string.
func flattenPolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum(i interface{}) *PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyKubernetesNamespaceAdmissionRulesEnforcementModeEnumRef(s)
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumMap flattens the contents of PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum{}
	}

	items := make(map[string]PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumSlice flattens the contents of PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum{}
	}

	items := make([]PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum with the same value as that string.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum(i interface{}) *PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyKubernetesServiceAccountAdmissionRulesEvaluationModeEnumRef(s)
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumMap flattens the contents of PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum{}
	}

	items := make(map[string]PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumSlice flattens the contents of PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum{}
	}

	items := make([]PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum with the same value as that string.
func flattenPolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum(i interface{}) *PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyKubernetesServiceAccountAdmissionRulesEnforcementModeEnumRef(s)
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumMap flattens the contents of PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum{}
	}

	items := make(map[string]PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumSlice flattens the contents of PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum{}
	}

	items := make([]PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum with the same value as that string.
func flattenPolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum(i interface{}) *PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyIstioServiceIdentityAdmissionRulesEvaluationModeEnumRef(s)
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumMap flattens the contents of PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum{}
	}

	items := make(map[string]PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumSlice flattens the contents of PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum from a JSON
// response object.
func flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum{}
	}

	items := make([]PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum with the same value as that string.
func flattenPolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum(i interface{}) *PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyIstioServiceIdentityAdmissionRulesEnforcementModeEnumRef(s)
}

// flattenPolicyDefaultAdmissionRuleEvaluationModeEnumMap flattens the contents of PolicyDefaultAdmissionRuleEvaluationModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyDefaultAdmissionRuleEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	items := make(map[string]PolicyDefaultAdmissionRuleEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEvaluationModeEnumSlice flattens the contents of PolicyDefaultAdmissionRuleEvaluationModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyDefaultAdmissionRuleEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRuleEvaluationModeEnum{}
	}

	items := make([]PolicyDefaultAdmissionRuleEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyDefaultAdmissionRuleEvaluationModeEnum with the same value as that string.
func flattenPolicyDefaultAdmissionRuleEvaluationModeEnum(i interface{}) *PolicyDefaultAdmissionRuleEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyDefaultAdmissionRuleEvaluationModeEnumRef(s)
}

// flattenPolicyDefaultAdmissionRuleEnforcementModeEnumMap flattens the contents of PolicyDefaultAdmissionRuleEnforcementModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEnforcementModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyDefaultAdmissionRuleEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	items := make(map[string]PolicyDefaultAdmissionRuleEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEnforcementModeEnumSlice flattens the contents of PolicyDefaultAdmissionRuleEnforcementModeEnum from a JSON
// response object.
func flattenPolicyDefaultAdmissionRuleEnforcementModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyDefaultAdmissionRuleEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyDefaultAdmissionRuleEnforcementModeEnum{}
	}

	items := make([]PolicyDefaultAdmissionRuleEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyDefaultAdmissionRuleEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyDefaultAdmissionRuleEnforcementModeEnum with the same value as that string.
func flattenPolicyDefaultAdmissionRuleEnforcementModeEnum(i interface{}) *PolicyDefaultAdmissionRuleEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyDefaultAdmissionRuleEnforcementModeEnumRef(s)
}

// flattenPolicyGlobalPolicyEvaluationModeEnumMap flattens the contents of PolicyGlobalPolicyEvaluationModeEnum from a JSON
// response object.
func flattenPolicyGlobalPolicyEvaluationModeEnumMap(c *Client, i interface{}, res *Policy) map[string]PolicyGlobalPolicyEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyGlobalPolicyEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyGlobalPolicyEvaluationModeEnum{}
	}

	items := make(map[string]PolicyGlobalPolicyEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyGlobalPolicyEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyGlobalPolicyEvaluationModeEnumSlice flattens the contents of PolicyGlobalPolicyEvaluationModeEnum from a JSON
// response object.
func flattenPolicyGlobalPolicyEvaluationModeEnumSlice(c *Client, i interface{}, res *Policy) []PolicyGlobalPolicyEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	items := make([]PolicyGlobalPolicyEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyGlobalPolicyEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyGlobalPolicyEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyGlobalPolicyEvaluationModeEnum with the same value as that string.
func flattenPolicyGlobalPolicyEvaluationModeEnum(i interface{}) *PolicyGlobalPolicyEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return PolicyGlobalPolicyEvaluationModeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Policy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalPolicy(b, c, r)
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
		return true
	}
}

type policyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         policyApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]policyDiff, error) {
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
	var diffs []policyDiff
	// For each operation name, create a policyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := policyDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (policyApiOperation, error) {
	switch opName {

	case "updatePolicyUpdatePolicyOperation":
		return &updatePolicyUpdatePolicyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractPolicyFields(r *Policy) error {
	vDefaultAdmissionRule := r.DefaultAdmissionRule
	if vDefaultAdmissionRule == nil {
		// note: explicitly not the empty object.
		vDefaultAdmissionRule = &PolicyDefaultAdmissionRule{}
	}
	if err := extractPolicyDefaultAdmissionRuleFields(r, vDefaultAdmissionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultAdmissionRule) {
		r.DefaultAdmissionRule = vDefaultAdmissionRule
	}
	return nil
}
func extractPolicyAdmissionWhitelistPatternsFields(r *Policy, o *PolicyAdmissionWhitelistPatterns) error {
	return nil
}
func extractPolicyClusterAdmissionRulesFields(r *Policy, o *PolicyClusterAdmissionRules) error {
	return nil
}
func extractPolicyKubernetesNamespaceAdmissionRulesFields(r *Policy, o *PolicyKubernetesNamespaceAdmissionRules) error {
	return nil
}
func extractPolicyKubernetesServiceAccountAdmissionRulesFields(r *Policy, o *PolicyKubernetesServiceAccountAdmissionRules) error {
	return nil
}
func extractPolicyIstioServiceIdentityAdmissionRulesFields(r *Policy, o *PolicyIstioServiceIdentityAdmissionRules) error {
	return nil
}
func extractPolicyDefaultAdmissionRuleFields(r *Policy, o *PolicyDefaultAdmissionRule) error {
	return nil
}

func postReadExtractPolicyFields(r *Policy) error {
	vDefaultAdmissionRule := r.DefaultAdmissionRule
	if vDefaultAdmissionRule == nil {
		// note: explicitly not the empty object.
		vDefaultAdmissionRule = &PolicyDefaultAdmissionRule{}
	}
	if err := postReadExtractPolicyDefaultAdmissionRuleFields(r, vDefaultAdmissionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultAdmissionRule) {
		r.DefaultAdmissionRule = vDefaultAdmissionRule
	}
	return nil
}
func postReadExtractPolicyAdmissionWhitelistPatternsFields(r *Policy, o *PolicyAdmissionWhitelistPatterns) error {
	return nil
}
func postReadExtractPolicyClusterAdmissionRulesFields(r *Policy, o *PolicyClusterAdmissionRules) error {
	return nil
}
func postReadExtractPolicyKubernetesNamespaceAdmissionRulesFields(r *Policy, o *PolicyKubernetesNamespaceAdmissionRules) error {
	return nil
}
func postReadExtractPolicyKubernetesServiceAccountAdmissionRulesFields(r *Policy, o *PolicyKubernetesServiceAccountAdmissionRules) error {
	return nil
}
func postReadExtractPolicyIstioServiceIdentityAdmissionRulesFields(r *Policy, o *PolicyIstioServiceIdentityAdmissionRules) error {
	return nil
}
func postReadExtractPolicyDefaultAdmissionRuleFields(r *Policy, o *PolicyDefaultAdmissionRule) error {
	return nil
}
