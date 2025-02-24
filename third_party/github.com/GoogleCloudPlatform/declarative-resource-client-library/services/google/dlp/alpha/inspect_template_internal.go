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
)

func (r *InspectTemplate) validate() error {

	if err := dcl.RequiredParameter(r.Parent, "Parent"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.InspectConfig) {
		if err := r.InspectConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Limits) {
		if err := r.Limits.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigInfoTypes) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigLimits) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) validate() error {
	if !dcl.IsEmptyValueIndirect(r.InfoType) {
		if err := r.InfoType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypes) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Dictionary", "Regex", "SurrogateType", "StoredType"}, r.Dictionary, r.Regex, r.SurrogateType, r.StoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.InfoType) {
		if err := r.InfoType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Dictionary) {
		if err := r.Dictionary.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Regex) {
		if err := r.Regex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SurrogateType) {
		if err := r.SurrogateType.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.StoredType) {
		if err := r.StoredType.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesInfoType) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesDictionary) validate() error {
	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"WordList", "CloudStoragePath"}, r.WordList, r.CloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.WordList) {
		if err := r.WordList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStoragePath) {
		if err := r.CloudStoragePath.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesRegex) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigCustomInfoTypesStoredType) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSet) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetInfoTypes) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRules) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"HotwordRule", "ExclusionRule"}, r.HotwordRule, r.ExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.HotwordRule) {
		if err := r.HotwordRule.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExclusionRule) {
		if err := r.ExclusionRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HotwordRegex) {
		if err := r.HotwordRegex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Proximity) {
		if err := r.Proximity.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.LikelihoodAdjustment) {
		if err := r.LikelihoodAdjustment.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"FixedLikelihood", "RelativeLikelihood"}, r.FixedLikelihood, r.RelativeLikelihood); err != nil {
		return err
	}
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRule) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"Dictionary", "Regex", "ExcludeInfoTypes"}, r.Dictionary, r.Regex, r.ExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Dictionary) {
		if err := r.Dictionary.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Regex) {
		if err := r.Regex.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ExcludeInfoTypes) {
		if err := r.ExcludeInfoTypes.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) validate() error {
	if err := dcl.ValidateExactlyOneOfFieldsSet([]string{"WordList", "CloudStoragePath"}, r.WordList, r.CloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.WordList) {
		if err := r.WordList.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStoragePath) {
		if err := r.CloudStoragePath.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) validate() error {
	return nil
}
func (r *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) validate() error {
	return nil
}
func (r *InspectTemplate) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://dlp.googleapis.com/v2/", params)
}

func (r *InspectTemplate) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *InspectTemplate) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/inspectTemplates", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/inspectTemplates", nr.basePath(), userBasePath, params), nil

}

func (r *InspectTemplate) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/inspectTemplates", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/inspectTemplates", nr.basePath(), userBasePath, params), nil

}

func (r *InspectTemplate) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, params), nil
}

// inspectTemplateApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type inspectTemplateApiOperation interface {
	do(context.Context, *InspectTemplate, *Client) error
}

// newUpdateInspectTemplateUpdateInspectTemplateRequest creates a request for an
// InspectTemplate resource's UpdateInspectTemplate update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateInspectTemplateUpdateInspectTemplateRequest(ctx context.Context, f *InspectTemplate, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandInspectTemplateInspectConfig(c, f.InspectConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding InspectConfig into inspectConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["inspectConfig"] = v
	}
	return req, nil
}

// marshalUpdateInspectTemplateUpdateInspectTemplateRequest converts the update into
// the final JSON request body.
func marshalUpdateInspectTemplateUpdateInspectTemplateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateInspectTemplateUpdateInspectTemplateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listInspectTemplateRaw(ctx context.Context, r *InspectTemplate, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != InspectTemplateMaxPage {
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

type listInspectTemplateOperation struct {
	InspectTemplates []map[string]interface{} `json:"inspectTemplates"`
	Token            string                   `json:"nextPageToken"`
}

func (c *Client) listInspectTemplate(ctx context.Context, r *InspectTemplate, pageToken string, pageSize int32) ([]*InspectTemplate, string, error) {
	b, err := c.listInspectTemplateRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listInspectTemplateOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*InspectTemplate
	for _, v := range m.InspectTemplates {
		res, err := unmarshalMapInspectTemplate(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllInspectTemplate(ctx context.Context, f func(*InspectTemplate) bool, resources []*InspectTemplate) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteInspectTemplate(ctx, res)
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

type deleteInspectTemplateOperation struct{}

func (op *deleteInspectTemplateOperation) do(ctx context.Context, r *InspectTemplate, c *Client) error {
	r, err := c.GetInspectTemplate(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "InspectTemplate not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetInspectTemplate checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete InspectTemplate: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetInspectTemplate(ctx, r)
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
type createInspectTemplateOperation struct {
	response map[string]interface{}
}

func (op *createInspectTemplateOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createInspectTemplateOperation) do(ctx context.Context, r *InspectTemplate, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a InspectTemplate with the wrong Name.
		return fmt.Errorf("server-generated parameter Name was specified by user as %v, should be unspecified", dcl.ValueOrEmptyString(r.Name))
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	// Include Name in URL substitution for initial GET request.
	m := op.response
	r.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))

	if _, err := c.GetInspectTemplate(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getInspectTemplateRaw(ctx context.Context, r *InspectTemplate) ([]byte, error) {

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

func (c *Client) inspectTemplateDiffsForRawDesired(ctx context.Context, rawDesired *InspectTemplate, opts ...dcl.ApplyOption) (initial, desired *InspectTemplate, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *InspectTemplate
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*InspectTemplate); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected InspectTemplate, got %T", sh)
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
		desired, err := canonicalizeInspectTemplateDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetInspectTemplate(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a InspectTemplate resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve InspectTemplate resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that InspectTemplate resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeInspectTemplateDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for InspectTemplate: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for InspectTemplate: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractInspectTemplateFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeInspectTemplateInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for InspectTemplate: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeInspectTemplateDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for InspectTemplate: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffInspectTemplate(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeInspectTemplateInitialState(rawInitial, rawDesired *InspectTemplate) (*InspectTemplate, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeInspectTemplateDesiredState(rawDesired, rawInitial *InspectTemplate, opts ...dcl.ApplyOption) (*InspectTemplate, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.InspectConfig = canonicalizeInspectTemplateInspectConfig(rawDesired.InspectConfig, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &InspectTemplate{}
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
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	canonicalDesired.InspectConfig = canonicalizeInspectTemplateInspectConfig(rawDesired.InspectConfig, rawInitial.InspectConfig, opts...)
	if dcl.NameToSelfLink(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeInspectTemplateNewState(c *Client, rawNew, rawDesired *InspectTemplate) (*InspectTemplate, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.InspectConfig) && dcl.IsEmptyValueIndirect(rawDesired.InspectConfig) {
		rawNew.InspectConfig = rawDesired.InspectConfig
	} else {
		rawNew.InspectConfig = canonicalizeNewInspectTemplateInspectConfig(c, rawDesired.InspectConfig, rawNew.InspectConfig)
	}

	if dcl.IsEmptyValueIndirect(rawNew.LocationId) && dcl.IsEmptyValueIndirect(rawDesired.LocationId) {
		rawNew.LocationId = rawDesired.LocationId
	} else {
		if dcl.StringCanonicalize(rawDesired.LocationId, rawNew.LocationId) {
			rawNew.LocationId = rawDesired.LocationId
		}
	}

	rawNew.Parent = rawDesired.Parent

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeInspectTemplateInspectConfig(des, initial *InspectTemplateInspectConfig, opts ...dcl.ApplyOption) *InspectTemplateInspectConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfig{}

	cDes.InfoTypes = canonicalizeInspectTemplateInspectConfigInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)
	if dcl.IsZeroValue(des.MinLikelihood) || (dcl.IsEmptyValueIndirect(des.MinLikelihood) && dcl.IsEmptyValueIndirect(initial.MinLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MinLikelihood = initial.MinLikelihood
	} else {
		cDes.MinLikelihood = des.MinLikelihood
	}
	cDes.Limits = canonicalizeInspectTemplateInspectConfigLimits(des.Limits, initial.Limits, opts...)
	if dcl.BoolCanonicalize(des.IncludeQuote, initial.IncludeQuote) || dcl.IsZeroValue(des.IncludeQuote) {
		cDes.IncludeQuote = initial.IncludeQuote
	} else {
		cDes.IncludeQuote = des.IncludeQuote
	}
	if dcl.BoolCanonicalize(des.ExcludeInfoTypes, initial.ExcludeInfoTypes) || dcl.IsZeroValue(des.ExcludeInfoTypes) {
		cDes.ExcludeInfoTypes = initial.ExcludeInfoTypes
	} else {
		cDes.ExcludeInfoTypes = des.ExcludeInfoTypes
	}
	cDes.CustomInfoTypes = canonicalizeInspectTemplateInspectConfigCustomInfoTypesSlice(des.CustomInfoTypes, initial.CustomInfoTypes, opts...)
	if dcl.IsZeroValue(des.ContentOptions) || (dcl.IsEmptyValueIndirect(des.ContentOptions) && dcl.IsEmptyValueIndirect(initial.ContentOptions)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ContentOptions = initial.ContentOptions
	} else {
		cDes.ContentOptions = des.ContentOptions
	}
	cDes.RuleSet = canonicalizeInspectTemplateInspectConfigRuleSetSlice(des.RuleSet, initial.RuleSet, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigSlice(des, initial []InspectTemplateInspectConfig, opts ...dcl.ApplyOption) []InspectTemplateInspectConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfig(c *Client, des, nw *InspectTemplateInspectConfig) *InspectTemplateInspectConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewInspectTemplateInspectConfigInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)
	nw.Limits = canonicalizeNewInspectTemplateInspectConfigLimits(c, des.Limits, nw.Limits)
	if dcl.BoolCanonicalize(des.IncludeQuote, nw.IncludeQuote) {
		nw.IncludeQuote = des.IncludeQuote
	}
	if dcl.BoolCanonicalize(des.ExcludeInfoTypes, nw.ExcludeInfoTypes) {
		nw.ExcludeInfoTypes = des.ExcludeInfoTypes
	}
	nw.CustomInfoTypes = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSlice(c, des.CustomInfoTypes, nw.CustomInfoTypes)
	nw.RuleSet = canonicalizeNewInspectTemplateInspectConfigRuleSetSlice(c, des.RuleSet, nw.RuleSet)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigSet(c *Client, des, nw []InspectTemplateInspectConfig) []InspectTemplateInspectConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigSlice(c *Client, des, nw []InspectTemplateInspectConfig) []InspectTemplateInspectConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfig(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigInfoTypes(des, initial *InspectTemplateInspectConfigInfoTypes, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigInfoTypesSlice(des, initial []InspectTemplateInspectConfigInfoTypes, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigInfoTypes(c *Client, des, nw *InspectTemplateInspectConfigInfoTypes) *InspectTemplateInspectConfigInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigInfoTypesSet(c *Client, des, nw []InspectTemplateInspectConfigInfoTypes) []InspectTemplateInspectConfigInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigInfoTypesSlice(c *Client, des, nw []InspectTemplateInspectConfigInfoTypes) []InspectTemplateInspectConfigInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigLimits(des, initial *InspectTemplateInspectConfigLimits, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigLimits {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigLimits{}

	if dcl.IsZeroValue(des.MaxFindingsPerItem) || (dcl.IsEmptyValueIndirect(des.MaxFindingsPerItem) && dcl.IsEmptyValueIndirect(initial.MaxFindingsPerItem)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindingsPerItem = initial.MaxFindingsPerItem
	} else {
		cDes.MaxFindingsPerItem = des.MaxFindingsPerItem
	}
	if dcl.IsZeroValue(des.MaxFindingsPerRequest) || (dcl.IsEmptyValueIndirect(des.MaxFindingsPerRequest) && dcl.IsEmptyValueIndirect(initial.MaxFindingsPerRequest)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindingsPerRequest = initial.MaxFindingsPerRequest
	} else {
		cDes.MaxFindingsPerRequest = des.MaxFindingsPerRequest
	}
	cDes.MaxFindingsPerInfoType = canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(des.MaxFindingsPerInfoType, initial.MaxFindingsPerInfoType, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigLimitsSlice(des, initial []InspectTemplateInspectConfigLimits, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigLimits {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigLimits, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigLimits(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigLimits, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigLimits(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigLimits(c *Client, des, nw *InspectTemplateInspectConfigLimits) *InspectTemplateInspectConfigLimits {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigLimits while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.MaxFindingsPerInfoType = canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, des.MaxFindingsPerInfoType, nw.MaxFindingsPerInfoType)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigLimitsSet(c *Client, des, nw []InspectTemplateInspectConfigLimits) []InspectTemplateInspectConfigLimits {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigLimits
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigLimitsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimits(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigLimitsSlice(c *Client, des, nw []InspectTemplateInspectConfigLimits) []InspectTemplateInspectConfigLimits {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigLimits
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimits(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(des, initial *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}

	cDes.InfoType = canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(des.InfoType, initial.InfoType, opts...)
	if dcl.IsZeroValue(des.MaxFindings) || (dcl.IsEmptyValueIndirect(des.MaxFindings) && dcl.IsEmptyValueIndirect(initial.MaxFindings)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxFindings = initial.MaxFindings
	} else {
		cDes.MaxFindings = des.MaxFindings
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(des, initial []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c *Client, des, nw *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoType = canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, des.InfoType, nw.InfoType)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSet(c *Client, des, nw []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, des, nw []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(des, initial *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(des, initial []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, des, nw *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSet(c *Client, des, nw []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, des, nw []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypes(des, initial *InspectTemplateInspectConfigCustomInfoTypes, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Dictionary != nil || (initial != nil && initial.Dictionary != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Regex, des.SurrogateType, des.StoredType) {
			des.Dictionary = nil
			if initial != nil {
				initial.Dictionary = nil
			}
		}
	}

	if des.Regex != nil || (initial != nil && initial.Regex != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.SurrogateType, des.StoredType) {
			des.Regex = nil
			if initial != nil {
				initial.Regex = nil
			}
		}
	}

	if des.SurrogateType != nil || (initial != nil && initial.SurrogateType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex, des.StoredType) {
			des.SurrogateType = nil
			if initial != nil {
				initial.SurrogateType = nil
			}
		}
	}

	if des.StoredType != nil || (initial != nil && initial.StoredType != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex, des.SurrogateType) {
			des.StoredType = nil
			if initial != nil {
				initial.StoredType = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypes{}

	cDes.InfoType = canonicalizeInspectTemplateInspectConfigCustomInfoTypesInfoType(des.InfoType, initial.InfoType, opts...)
	if dcl.IsZeroValue(des.Likelihood) || (dcl.IsEmptyValueIndirect(des.Likelihood) && dcl.IsEmptyValueIndirect(initial.Likelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Likelihood = initial.Likelihood
	} else {
		cDes.Likelihood = des.Likelihood
	}
	cDes.Dictionary = canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionary(des.Dictionary, initial.Dictionary, opts...)
	cDes.Regex = canonicalizeInspectTemplateInspectConfigCustomInfoTypesRegex(des.Regex, initial.Regex, opts...)
	cDes.SurrogateType = canonicalizeInspectTemplateInspectConfigCustomInfoTypesSurrogateType(des.SurrogateType, initial.SurrogateType, opts...)
	cDes.StoredType = canonicalizeInspectTemplateInspectConfigCustomInfoTypesStoredType(des.StoredType, initial.StoredType, opts...)
	if dcl.IsZeroValue(des.ExclusionType) || (dcl.IsEmptyValueIndirect(des.ExclusionType) && dcl.IsEmptyValueIndirect(initial.ExclusionType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ExclusionType = initial.ExclusionType
	} else {
		cDes.ExclusionType = des.ExclusionType
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypes, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypes(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypes) *InspectTemplateInspectConfigCustomInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoType = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoType(c, des.InfoType, nw.InfoType)
	nw.Dictionary = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionary(c, des.Dictionary, nw.Dictionary)
	nw.Regex = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegex(c, des.Regex, nw.Regex)
	nw.SurrogateType = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, des.SurrogateType, nw.SurrogateType)
	nw.StoredType = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredType(c, des.StoredType, nw.StoredType)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypes) []InspectTemplateInspectConfigCustomInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypes) []InspectTemplateInspectConfigCustomInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesInfoType(des, initial *InspectTemplateInspectConfigCustomInfoTypesInfoType, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesInfoType{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesInfoType, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesInfoType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesInfoType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesInfoType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesInfoType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoType(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesInfoType) *InspectTemplateInspectConfigCustomInfoTypesInfoType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesInfoType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoTypeSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesInfoType) []InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesInfoType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesInfoTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesInfoType) []InspectTemplateInspectConfigCustomInfoTypesInfoType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesInfoType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesInfoType(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionary(des, initial *InspectTemplateInspectConfigCustomInfoTypesDictionary, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.WordList != nil || (initial != nil && initial.WordList != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStoragePath) {
			des.WordList = nil
			if initial != nil {
				initial.WordList = nil
			}
		}
	}

	if des.CloudStoragePath != nil || (initial != nil && initial.CloudStoragePath != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.WordList) {
			des.CloudStoragePath = nil
			if initial != nil {
				initial.CloudStoragePath = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesDictionary{}

	cDes.WordList = canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(des.WordList, initial.WordList, opts...)
	cDes.CloudStoragePath = canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(des.CloudStoragePath, initial.CloudStoragePath, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionarySlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesDictionary, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionary(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesDictionary) *InspectTemplateInspectConfigCustomInfoTypesDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.WordList = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, des.WordList, nw.WordList)
	nw.CloudStoragePath = canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, des.CloudStoragePath, nw.CloudStoragePath)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionarySet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionary) []InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionarySlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionary) []InspectTemplateInspectConfigCustomInfoTypesDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(des, initial *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}

	if dcl.StringArrayCanonicalize(des.Words, initial.Words) {
		cDes.Words = initial.Words
	} else {
		cDes.Words = des.Words
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Words, nw.Words) {
		nw.Words = des.Words
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(des, initial *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesRegex(des, initial *InspectTemplateInspectConfigCustomInfoTypesRegex, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesRegexSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesRegex, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegex(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesRegex) *InspectTemplateInspectConfigCustomInfoTypesRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegexSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesRegex) []InspectTemplateInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegexSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesRegex) []InspectTemplateInspectConfigCustomInfoTypesRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesRegex(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesSurrogateType(des, initial *InspectTemplateInspectConfigCustomInfoTypesSurrogateType, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}
	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesSurrogateType, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesSurrogateType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesSurrogateType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesSurrogateType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesSurrogateType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) *InspectTemplateInspectConfigCustomInfoTypesSurrogateType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesSurrogateType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesSurrogateType) []InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesSurrogateType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesSurrogateType) []InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesSurrogateType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesStoredType(des, initial *InspectTemplateInspectConfigCustomInfoTypesStoredType, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigCustomInfoTypesStoredType{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice(des, initial []InspectTemplateInspectConfigCustomInfoTypesStoredType, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigCustomInfoTypesStoredType, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesStoredType(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesStoredType, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigCustomInfoTypesStoredType(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredType(c *Client, des, nw *InspectTemplateInspectConfigCustomInfoTypesStoredType) *InspectTemplateInspectConfigCustomInfoTypesStoredType {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigCustomInfoTypesStoredType while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredTypeSet(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesStoredType) []InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigCustomInfoTypesStoredType
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigCustomInfoTypesStoredTypeNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredType(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, des, nw []InspectTemplateInspectConfigCustomInfoTypesStoredType) []InspectTemplateInspectConfigCustomInfoTypesStoredType {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigCustomInfoTypesStoredType
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigCustomInfoTypesStoredType(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSet(des, initial *InspectTemplateInspectConfigRuleSet, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSet{}

	cDes.InfoTypes = canonicalizeInspectTemplateInspectConfigRuleSetInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)
	cDes.Rules = canonicalizeInspectTemplateInspectConfigRuleSetRulesSlice(des.Rules, initial.Rules, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetSlice(des, initial []InspectTemplateInspectConfigRuleSet, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSet {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSet(c *Client, des, nw *InspectTemplateInspectConfigRuleSet) *InspectTemplateInspectConfigRuleSet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)
	nw.Rules = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesSlice(c, des.Rules, nw.Rules)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSet) []InspectTemplateInspectConfigRuleSet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSet) []InspectTemplateInspectConfigRuleSet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSet(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetInfoTypes(des, initial *InspectTemplateInspectConfigRuleSetInfoTypes, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetInfoTypesSlice(des, initial []InspectTemplateInspectConfigRuleSetInfoTypes, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypes(c *Client, des, nw *InspectTemplateInspectConfigRuleSetInfoTypes) *InspectTemplateInspectConfigRuleSetInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypesSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetInfoTypes) []InspectTemplateInspectConfigRuleSetInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypesSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetInfoTypes) []InspectTemplateInspectConfigRuleSetInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRules(des, initial *InspectTemplateInspectConfigRuleSetRules, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRules {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.HotwordRule != nil || (initial != nil && initial.HotwordRule != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.ExclusionRule) {
			des.HotwordRule = nil
			if initial != nil {
				initial.HotwordRule = nil
			}
		}
	}

	if des.ExclusionRule != nil || (initial != nil && initial.ExclusionRule != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.HotwordRule) {
			des.ExclusionRule = nil
			if initial != nil {
				initial.ExclusionRule = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRules{}

	cDes.HotwordRule = canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRule(des.HotwordRule, initial.HotwordRule, opts...)
	cDes.ExclusionRule = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRule(des.ExclusionRule, initial.ExclusionRule, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesSlice(des, initial []InspectTemplateInspectConfigRuleSetRules, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRules {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRules, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRules(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRules, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRules(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRules(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRules) *InspectTemplateInspectConfigRuleSetRules {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRules while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRule = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, des.HotwordRule, nw.HotwordRule)
	nw.ExclusionRule = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, des.ExclusionRule, nw.ExclusionRule)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRules) []InspectTemplateInspectConfigRuleSetRules {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRules
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRules(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRules) []InspectTemplateInspectConfigRuleSetRules {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRules
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRules(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRule(des, initial *InspectTemplateInspectConfigRuleSetRulesHotwordRule, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesHotwordRule{}

	cDes.HotwordRegex = canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(des.HotwordRegex, initial.HotwordRegex, opts...)
	cDes.Proximity = canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(des.Proximity, initial.Proximity, opts...)
	cDes.LikelihoodAdjustment = canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(des.LikelihoodAdjustment, initial.LikelihoodAdjustment, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesHotwordRule, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRule(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesHotwordRule) *InspectTemplateInspectConfigRuleSetRulesHotwordRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesHotwordRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.HotwordRegex = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, des.HotwordRegex, nw.HotwordRegex)
	nw.Proximity = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, des.Proximity, nw.Proximity)
	nw.LikelihoodAdjustment = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, des.LikelihoodAdjustment, nw.LikelihoodAdjustment)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRule) []InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRule) []InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(des, initial *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(des, initial *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}

	if dcl.IsZeroValue(des.WindowBefore) || (dcl.IsEmptyValueIndirect(des.WindowBefore) && dcl.IsEmptyValueIndirect(initial.WindowBefore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowBefore = initial.WindowBefore
	} else {
		cDes.WindowBefore = des.WindowBefore
	}
	if dcl.IsZeroValue(des.WindowAfter) || (dcl.IsEmptyValueIndirect(des.WindowAfter) && dcl.IsEmptyValueIndirect(initial.WindowAfter)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WindowAfter = initial.WindowAfter
	} else {
		cDes.WindowAfter = des.WindowAfter
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice(des, initial []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(des, initial *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.FixedLikelihood != nil || (initial != nil && initial.FixedLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.RelativeLikelihood) {
			des.FixedLikelihood = nil
			if initial != nil {
				initial.FixedLikelihood = nil
			}
		}
	}

	if des.RelativeLikelihood != nil || (initial != nil && initial.RelativeLikelihood != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.FixedLikelihood) {
			des.RelativeLikelihood = nil
			if initial != nil {
				initial.RelativeLikelihood = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsZeroValue(des.FixedLikelihood) || (dcl.IsEmptyValueIndirect(des.FixedLikelihood) && dcl.IsEmptyValueIndirect(initial.FixedLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FixedLikelihood = initial.FixedLikelihood
	} else {
		cDes.FixedLikelihood = des.FixedLikelihood
	}
	if dcl.IsZeroValue(des.RelativeLikelihood) || (dcl.IsEmptyValueIndirect(des.RelativeLikelihood) && dcl.IsEmptyValueIndirect(initial.RelativeLikelihood)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RelativeLikelihood = initial.RelativeLikelihood
	} else {
		cDes.RelativeLikelihood = des.RelativeLikelihood
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRule(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRule, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.Dictionary != nil || (initial != nil && initial.Dictionary != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Regex, des.ExcludeInfoTypes) {
			des.Dictionary = nil
			if initial != nil {
				initial.Dictionary = nil
			}
		}
	}

	if des.Regex != nil || (initial != nil && initial.Regex != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.ExcludeInfoTypes) {
			des.Regex = nil
			if initial != nil {
				initial.Regex = nil
			}
		}
	}

	if des.ExcludeInfoTypes != nil || (initial != nil && initial.ExcludeInfoTypes != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.Dictionary, des.Regex) {
			des.ExcludeInfoTypes = nil
			if initial != nil {
				initial.ExcludeInfoTypes = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRule{}

	cDes.Dictionary = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(des.Dictionary, initial.Dictionary, opts...)
	cDes.Regex = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(des.Regex, initial.Regex, opts...)
	cDes.ExcludeInfoTypes = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(des.ExcludeInfoTypes, initial.ExcludeInfoTypes, opts...)
	if dcl.IsZeroValue(des.MatchingType) || (dcl.IsEmptyValueIndirect(des.MatchingType) && dcl.IsEmptyValueIndirect(initial.MatchingType)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MatchingType = initial.MatchingType
	} else {
		cDes.MatchingType = des.MatchingType
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRule, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRule(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRule) *InspectTemplateInspectConfigRuleSetRulesExclusionRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Dictionary = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, des.Dictionary, nw.Dictionary)
	nw.Regex = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, des.Regex, nw.Regex)
	nw.ExcludeInfoTypes = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, des.ExcludeInfoTypes, nw.ExcludeInfoTypes)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRule) []InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRule) []InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.WordList != nil || (initial != nil && initial.WordList != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStoragePath) {
			des.WordList = nil
			if initial != nil {
				initial.WordList = nil
			}
		}
	}

	if des.CloudStoragePath != nil || (initial != nil && initial.CloudStoragePath != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.WordList) {
			des.CloudStoragePath = nil
			if initial != nil {
				initial.CloudStoragePath = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}

	cDes.WordList = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(des.WordList, initial.WordList, opts...)
	cDes.CloudStoragePath = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(des.CloudStoragePath, initial.CloudStoragePath, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.WordList = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, des.WordList, nw.WordList)
	nw.CloudStoragePath = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, des.CloudStoragePath, nw.CloudStoragePath)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}

	if dcl.StringArrayCanonicalize(des.Words, initial.Words) {
		cDes.Words = initial.Words
	} else {
		cDes.Words = des.Words
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Words, nw.Words) {
		nw.Words = des.Words
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}

	if dcl.StringCanonicalize(des.Pattern, initial.Pattern) || dcl.IsZeroValue(des.Pattern) {
		cDes.Pattern = initial.Pattern
	} else {
		cDes.Pattern = des.Pattern
	}
	if dcl.IsZeroValue(des.GroupIndexes) || (dcl.IsEmptyValueIndirect(des.GroupIndexes) && dcl.IsEmptyValueIndirect(initial.GroupIndexes)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GroupIndexes = initial.GroupIndexes
	} else {
		cDes.GroupIndexes = des.GroupIndexes
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}

	cDes.InfoTypes = canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(des.InfoTypes, initial.InfoTypes, opts...)

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.InfoTypes = canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, des.InfoTypes, nw.InfoTypes)

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &d, &n))
	}

	return items
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(des, initial *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, opts ...dcl.ApplyOption) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(des, initial []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, opts ...dcl.ApplyOption) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(des))
		for _, d := range des {
			cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(des))
	for i, d := range des {
		cd := canonicalizeInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, des, nw *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSet(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, des, nw []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &d, &n))
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
func diffInspectTemplate(c *Client, desired, actual *InspectTemplate, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.InspectConfig, actual.InspectConfig, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigNewStyle, EmptyObject: EmptyInspectTemplateInspectConfig, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InspectConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocationId, actual.LocationId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocationId")); len(ds) != 0 || err != nil {
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
func compareInspectTemplateInspectConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfig)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfig or *InspectTemplateInspectConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfig)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigInfoTypesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigInfoTypes, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinLikelihood, actual.MinLikelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MinLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Limits, actual.Limits, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigLimitsNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigLimits, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Limits")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IncludeQuote, actual.IncludeQuote, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("IncludeQuote")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludeInfoTypes, actual.ExcludeInfoTypes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("ExcludeInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomInfoTypes, actual.CustomInfoTypes, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypes, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("CustomInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ContentOptions, actual.ContentOptions, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("ContentOptions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RuleSet, actual.RuleSet, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSet, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("RuleSet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigInfoTypes or *InspectTemplateInspectConfigInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigLimitsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigLimits)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimits or *InspectTemplateInspectConfigLimits", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigLimits)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigLimits)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimits", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerItem, actual.MaxFindingsPerItem, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MaxFindingsPerItem")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerRequest, actual.MaxFindingsPerRequest, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MaxFindingsPerRequest")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindingsPerInfoType, actual.MaxFindingsPerInfoType, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MaxFindingsPerInfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType or *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoType, actual.InfoType, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxFindings, actual.MaxFindings, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MaxFindings")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType or *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypes or *InspectTemplateInspectConfigCustomInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoType, actual.InfoType, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesInfoTypeNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesInfoType, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InfoType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Likelihood, actual.Likelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Likelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dictionary, actual.Dictionary, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesDictionaryNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesDictionary, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Dictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Regex, actual.Regex, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesRegexNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesRegex, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Regex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SurrogateType, actual.SurrogateType, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesSurrogateType, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("SurrogateType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StoredType, actual.StoredType, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesStoredTypeNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesStoredType, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("StoredType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExclusionType, actual.ExclusionType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("ExclusionType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesInfoTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesInfoType)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesInfoType or *InspectTemplateInspectConfigCustomInfoTypesInfoType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesInfoType)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesInfoType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesInfoType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesDictionary)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionary or *InspectTemplateInspectConfigCustomInfoTypesDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesDictionary)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WordList, actual.WordList, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("WordList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStoragePath, actual.CloudStoragePath, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("CloudStoragePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList or *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Words, actual.Words, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Words")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath or *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesRegex)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesRegex or *InspectTemplateInspectConfigCustomInfoTypesRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesRegex)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	return diffs, nil
}

func compareInspectTemplateInspectConfigCustomInfoTypesStoredTypeNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigCustomInfoTypesStoredType)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigCustomInfoTypesStoredType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesStoredType or *InspectTemplateInspectConfigCustomInfoTypesStoredType", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigCustomInfoTypesStoredType)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigCustomInfoTypesStoredType)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigCustomInfoTypesStoredType", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSet)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSet or *InspectTemplateInspectConfigRuleSet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSet)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetInfoTypesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetInfoTypes, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rules, actual.Rules, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRules, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Rules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetInfoTypes or *InspectTemplateInspectConfigRuleSetInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRules)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRules or *InspectTemplateInspectConfigRuleSetRules", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRules)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRules)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRules", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRule, actual.HotwordRule, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRule, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("HotwordRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExclusionRule, actual.ExclusionRule, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRule, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("ExclusionRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesHotwordRule)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRule or *InspectTemplateInspectConfigRuleSetRulesHotwordRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesHotwordRule)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesHotwordRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HotwordRegex, actual.HotwordRegex, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("HotwordRegex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Proximity, actual.Proximity, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Proximity")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LikelihoodAdjustment, actual.LikelihoodAdjustment, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("LikelihoodAdjustment")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex or *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity or *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WindowBefore, actual.WindowBefore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("WindowBefore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WindowAfter, actual.WindowAfter, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("WindowAfter")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment or *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FixedLikelihood, actual.FixedLikelihood, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("FixedLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RelativeLikelihood, actual.RelativeLikelihood, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("RelativeLikelihood")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRule)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRule or *InspectTemplateInspectConfigRuleSetRulesExclusionRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRule)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Dictionary, actual.Dictionary, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Dictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Regex, actual.Regex, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Regex")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExcludeInfoTypes, actual.ExcludeInfoTypes, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("ExcludeInfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MatchingType, actual.MatchingType, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("MatchingType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WordList, actual.WordList, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("WordList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStoragePath, actual.CloudStoragePath, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("CloudStoragePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Words, actual.Words, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Words")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InfoTypes, actual.InfoTypes, dcl.DiffInfo{ObjectFunction: compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle, EmptyObject: EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("InfoTypes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	if !ok {
		desiredNotPointer, ok := d.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes or *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	if !ok {
		actualNotPointer, ok := a.(InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateInspectTemplateUpdateInspectTemplateOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
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
func (r *InspectTemplate) urlNormalized() *InspectTemplate {
	normalized := dcl.Copy(*r).(InspectTemplate)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.LocationId = dcl.SelfLinkToName(r.LocationId)
	normalized.Parent = r.Parent
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *InspectTemplate) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateInspectTemplate" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"parent":   dcl.ValueOrEmptyString(nr.Parent),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("{{parent}}/locations/{{location}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, fields), nil
		}

		return dcl.URL("{{parent}}/inspectTemplates/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the InspectTemplate resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *InspectTemplate) marshal(c *Client) ([]byte, error) {
	m, err := expandInspectTemplate(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling InspectTemplate: %w", err)
	}
	m = encodeInspectTemplateCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalInspectTemplate decodes JSON responses into the InspectTemplate resource schema.
func unmarshalInspectTemplate(b []byte, c *Client, res *InspectTemplate) (*InspectTemplate, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapInspectTemplate(m, c, res)
}

func unmarshalMapInspectTemplate(m map[string]interface{}, c *Client, res *InspectTemplate) (*InspectTemplate, error) {

	flattened := flattenInspectTemplate(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandInspectTemplate expands InspectTemplate into a JSON request object.
func expandInspectTemplate(c *Client, f *InspectTemplate) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v, err := expandInspectTemplateInspectConfig(c, f.InspectConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding InspectConfig into inspectConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["inspectConfig"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Parent into parent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["parent"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenInspectTemplate flattens InspectTemplate from a JSON request object into the
// InspectTemplate type.
func flattenInspectTemplate(c *Client, i interface{}, res *InspectTemplate) *InspectTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &InspectTemplate{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.UpdateTime = dcl.FlattenString(m["updateTime"])
	resultRes.InspectConfig = flattenInspectTemplateInspectConfig(c, m["inspectConfig"], res)
	resultRes.LocationId = dcl.FlattenString(m["locationId"])
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandInspectTemplateInspectConfigMap expands the contents of InspectTemplateInspectConfig into a JSON
// request object.
func expandInspectTemplateInspectConfigMap(c *Client, f map[string]InspectTemplateInspectConfig, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigSlice expands the contents of InspectTemplateInspectConfig into a JSON
// request object.
func expandInspectTemplateInspectConfigSlice(c *Client, f []InspectTemplateInspectConfig, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigMap flattens the contents of InspectTemplateInspectConfig from a JSON
// response object.
func flattenInspectTemplateInspectConfigMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfig{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfig{}
	}

	items := make(map[string]InspectTemplateInspectConfig)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigSlice flattens the contents of InspectTemplateInspectConfig from a JSON
// response object.
func flattenInspectTemplateInspectConfigSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfig{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfig{}
	}

	items := make([]InspectTemplateInspectConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfig expands an instance of InspectTemplateInspectConfig into a JSON
// request object.
func expandInspectTemplateInspectConfig(c *Client, f *InspectTemplateInspectConfig, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}
	if v := f.MinLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["minLikelihood"] = v
	}
	if v, err := expandInspectTemplateInspectConfigLimits(c, f.Limits, res); err != nil {
		return nil, fmt.Errorf("error expanding Limits into limits: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["limits"] = v
	}
	if v := f.IncludeQuote; !dcl.IsEmptyValueIndirect(v) {
		m["includeQuote"] = v
	}
	if v := f.ExcludeInfoTypes; !dcl.IsEmptyValueIndirect(v) {
		m["excludeInfoTypes"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesSlice(c, f.CustomInfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding CustomInfoTypes into customInfoTypes: %w", err)
	} else if v != nil {
		m["customInfoTypes"] = v
	}
	if v := f.ContentOptions; v != nil {
		m["contentOptions"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetSlice(c, f.RuleSet, res); err != nil {
		return nil, fmt.Errorf("error expanding RuleSet into ruleSet: %w", err)
	} else if v != nil {
		m["ruleSet"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfig flattens an instance of InspectTemplateInspectConfig from a JSON
// response object.
func flattenInspectTemplateInspectConfig(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfig
	}
	r.InfoTypes = flattenInspectTemplateInspectConfigInfoTypesSlice(c, m["infoTypes"], res)
	r.MinLikelihood = flattenInspectTemplateInspectConfigMinLikelihoodEnum(m["minLikelihood"])
	r.Limits = flattenInspectTemplateInspectConfigLimits(c, m["limits"], res)
	r.IncludeQuote = dcl.FlattenBool(m["includeQuote"])
	r.ExcludeInfoTypes = dcl.FlattenBool(m["excludeInfoTypes"])
	r.CustomInfoTypes = flattenInspectTemplateInspectConfigCustomInfoTypesSlice(c, m["customInfoTypes"], res)
	r.ContentOptions = flattenInspectTemplateInspectConfigContentOptionsEnumSlice(c, m["contentOptions"], res)
	r.RuleSet = flattenInspectTemplateInspectConfigRuleSetSlice(c, m["ruleSet"], res)

	return r
}

// expandInspectTemplateInspectConfigInfoTypesMap expands the contents of InspectTemplateInspectConfigInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigInfoTypesMap(c *Client, f map[string]InspectTemplateInspectConfigInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigInfoTypesSlice expands the contents of InspectTemplateInspectConfigInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigInfoTypesSlice(c *Client, f []InspectTemplateInspectConfigInfoTypes, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigInfoTypesMap flattens the contents of InspectTemplateInspectConfigInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigInfoTypesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigInfoTypes{}
	}

	items := make(map[string]InspectTemplateInspectConfigInfoTypes)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigInfoTypesSlice flattens the contents of InspectTemplateInspectConfigInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigInfoTypesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigInfoTypes{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigInfoTypes{}
	}

	items := make([]InspectTemplateInspectConfigInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigInfoTypes expands an instance of InspectTemplateInspectConfigInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigInfoTypes(c *Client, f *InspectTemplateInspectConfigInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigInfoTypes flattens an instance of InspectTemplateInspectConfigInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigInfoTypes(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandInspectTemplateInspectConfigLimitsMap expands the contents of InspectTemplateInspectConfigLimits into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMap(c *Client, f map[string]InspectTemplateInspectConfigLimits, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigLimits(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigLimitsSlice expands the contents of InspectTemplateInspectConfigLimits into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsSlice(c *Client, f []InspectTemplateInspectConfigLimits, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigLimits(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigLimitsMap flattens the contents of InspectTemplateInspectConfigLimits from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigLimits {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigLimits{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigLimits{}
	}

	items := make(map[string]InspectTemplateInspectConfigLimits)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigLimits(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigLimitsSlice flattens the contents of InspectTemplateInspectConfigLimits from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigLimits {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigLimits{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigLimits{}
	}

	items := make([]InspectTemplateInspectConfigLimits, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigLimits(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigLimits expands an instance of InspectTemplateInspectConfigLimits into a JSON
// request object.
func expandInspectTemplateInspectConfigLimits(c *Client, f *InspectTemplateInspectConfigLimits, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxFindingsPerItem; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindingsPerItem"] = v
	}
	if v := f.MaxFindingsPerRequest; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindingsPerRequest"] = v
	}
	if v, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, f.MaxFindingsPerInfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding MaxFindingsPerInfoType into maxFindingsPerInfoType: %w", err)
	} else if v != nil {
		m["maxFindingsPerInfoType"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigLimits flattens an instance of InspectTemplateInspectConfigLimits from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimits(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigLimits {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigLimits{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigLimits
	}
	r.MaxFindingsPerItem = dcl.FlattenInteger(m["maxFindingsPerItem"])
	r.MaxFindingsPerRequest = dcl.FlattenInteger(m["maxFindingsPerRequest"])
	r.MaxFindingsPerInfoType = flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c, m["maxFindingsPerInfoType"], res)

	return r
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMap expands the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMap(c *Client, f map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice expands the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, f []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMap flattens the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	items := make(map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice flattens the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}
	}

	items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType expands an instance of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c *Client, f *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, f.InfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoType into infoType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["infoType"] = v
	}
	if v := f.MaxFindings; !dcl.IsEmptyValueIndirect(v) {
		m["maxFindings"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType flattens an instance of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType
	}
	r.InfoType = flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, m["infoType"], res)
	r.MaxFindings = dcl.FlattenInteger(m["maxFindings"])

	return r
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap expands the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap(c *Client, f map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice expands the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, f []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap flattens the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	items := make(map[string]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice flattens the contents of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}

	items := make([]InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType expands an instance of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, f *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType flattens an instance of InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesMap expands the contents of InspectTemplateInspectConfigCustomInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypes, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypes{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypes)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypes{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypes{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypes expands an instance of InspectTemplateInspectConfigCustomInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypes(c *Client, f *InspectTemplateInspectConfigCustomInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesInfoType(c, f.InfoType, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoType into infoType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["infoType"] = v
	}
	if v := f.Likelihood; !dcl.IsEmptyValueIndirect(v) {
		m["likelihood"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dictionary"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, f.SurrogateType, res); err != nil {
		return nil, fmt.Errorf("error expanding SurrogateType into surrogateType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["surrogateType"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesStoredType(c, f.StoredType, res); err != nil {
		return nil, fmt.Errorf("error expanding StoredType into storedType: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["storedType"] = v
	}
	if v := f.ExclusionType; !dcl.IsEmptyValueIndirect(v) {
		m["exclusionType"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypes flattens an instance of InspectTemplateInspectConfigCustomInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypes(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypes
	}
	r.InfoType = flattenInspectTemplateInspectConfigCustomInfoTypesInfoType(c, m["infoType"], res)
	r.Likelihood = flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(m["likelihood"])
	r.Dictionary = flattenInspectTemplateInspectConfigCustomInfoTypesDictionary(c, m["dictionary"], res)
	r.Regex = flattenInspectTemplateInspectConfigCustomInfoTypesRegex(c, m["regex"], res)
	r.SurrogateType = flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, m["surrogateType"], res)
	r.StoredType = flattenInspectTemplateInspectConfigCustomInfoTypesStoredType(c, m["storedType"], res)
	r.ExclusionType = flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(m["exclusionType"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesInfoTypeMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesInfoTypeMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesInfoType, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesInfoType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesInfoTypeMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesInfoTypeMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesInfoType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesInfoType)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesInfoType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesInfoTypeSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesInfoType {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesInfoType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesInfoType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesInfoType expands an instance of InspectTemplateInspectConfigCustomInfoTypesInfoType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesInfoType(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesInfoType, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesInfoType flattens an instance of InspectTemplateInspectConfigCustomInfoTypesInfoType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesInfoType(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesInfoType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesInfoType
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesDictionary, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionarySlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionarySlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesDictionary, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesDictionary)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionarySlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionarySlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionary expands an instance of InspectTemplateInspectConfigCustomInfoTypesDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionary(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesDictionary, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, f.WordList, res); err != nil {
		return nil, fmt.Errorf("error expanding WordList into wordList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["wordList"] = v
	}
	if v, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, f.CloudStoragePath, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStoragePath into cloudStoragePath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStoragePath"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionary flattens an instance of InspectTemplateInspectConfigCustomInfoTypesDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionary(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesDictionary
	}
	r.WordList = flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, m["wordList"], res)
	r.CloudStoragePath = flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, m["cloudStoragePath"], res)

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList expands an instance of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Words; v != nil {
		m["words"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList flattens an instance of InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList
	}
	r.Words = dcl.FlattenStringSlice(m["words"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath expands an instance of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath flattens an instance of InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesRegexMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesRegexMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesRegexSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesRegexSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesRegex, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesRegexMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesRegexMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesRegex)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesRegexSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesRegexSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesRegex expands an instance of InspectTemplateInspectConfigCustomInfoTypesRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesRegex(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesRegex flattens an instance of InspectTemplateInspectConfigCustomInfoTypesRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesRegex(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesSurrogateType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesSurrogateType, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesSurrogateType)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesSurrogateType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesSurrogateType expands an instance of InspectTemplateInspectConfigCustomInfoTypesSurrogateType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesSurrogateType, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateType flattens an instance of InspectTemplateInspectConfigCustomInfoTypesSurrogateType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesSurrogateType(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	_, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesSurrogateType
	}

	return r
}

// expandInspectTemplateInspectConfigCustomInfoTypesStoredTypeMap expands the contents of InspectTemplateInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesStoredTypeMap(c *Client, f map[string]InspectTemplateInspectConfigCustomInfoTypesStoredType, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesStoredType(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice expands the contents of InspectTemplateInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, f []InspectTemplateInspectConfigCustomInfoTypesStoredType, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigCustomInfoTypesStoredType(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesStoredTypeMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesStoredTypeMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesStoredType {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesStoredType)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesStoredType(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesStoredTypeSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesStoredType {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesStoredType, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesStoredType(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigCustomInfoTypesStoredType expands an instance of InspectTemplateInspectConfigCustomInfoTypesStoredType into a JSON
// request object.
func expandInspectTemplateInspectConfigCustomInfoTypesStoredType(c *Client, f *InspectTemplateInspectConfigCustomInfoTypesStoredType, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigCustomInfoTypesStoredType flattens an instance of InspectTemplateInspectConfigCustomInfoTypesStoredType from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesStoredType(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigCustomInfoTypesStoredType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigCustomInfoTypesStoredType{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigCustomInfoTypesStoredType
	}
	r.Name = dcl.FlattenString(m["name"])
	r.CreateTime = dcl.FlattenString(m["createTime"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetMap expands the contents of InspectTemplateInspectConfigRuleSet into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSet, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetSlice expands the contents of InspectTemplateInspectConfigRuleSet into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetSlice(c *Client, f []InspectTemplateInspectConfigRuleSet, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetMap flattens the contents of InspectTemplateInspectConfigRuleSet from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSet{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSet{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSet)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetSlice flattens the contents of InspectTemplateInspectConfigRuleSet from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSet {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSet{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSet{}
	}

	items := make([]InspectTemplateInspectConfigRuleSet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSet expands an instance of InspectTemplateInspectConfigRuleSet into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSet(c *Client, f *InspectTemplateInspectConfigRuleSet, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesSlice(c, f.Rules, res); err != nil {
		return nil, fmt.Errorf("error expanding Rules into rules: %w", err)
	} else if v != nil {
		m["rules"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSet flattens an instance of InspectTemplateInspectConfigRuleSet from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSet(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSet
	}
	r.InfoTypes = flattenInspectTemplateInspectConfigRuleSetInfoTypesSlice(c, m["infoTypes"], res)
	r.Rules = flattenInspectTemplateInspectConfigRuleSetRulesSlice(c, m["rules"], res)

	return r
}

// expandInspectTemplateInspectConfigRuleSetInfoTypesMap expands the contents of InspectTemplateInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetInfoTypesMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetInfoTypesSlice expands the contents of InspectTemplateInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetInfoTypesSlice(c *Client, f []InspectTemplateInspectConfigRuleSetInfoTypes, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetInfoTypesMap flattens the contents of InspectTemplateInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetInfoTypesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetInfoTypes{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetInfoTypes)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetInfoTypesSlice flattens the contents of InspectTemplateInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetInfoTypesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetInfoTypes{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetInfoTypes{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetInfoTypes expands an instance of InspectTemplateInspectConfigRuleSetInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetInfoTypes(c *Client, f *InspectTemplateInspectConfigRuleSetInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetInfoTypes flattens an instance of InspectTemplateInspectConfigRuleSetInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetInfoTypes(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesMap expands the contents of InspectTemplateInspectConfigRuleSetRules into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRules, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRules(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesSlice expands the contents of InspectTemplateInspectConfigRuleSetRules into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRules, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRules(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesMap flattens the contents of InspectTemplateInspectConfigRuleSetRules from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRules {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRules{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRules{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRules)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRules(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesSlice flattens the contents of InspectTemplateInspectConfigRuleSetRules from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRules {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRules{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRules{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRules, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRules(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRules expands an instance of InspectTemplateInspectConfigRuleSetRules into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRules(c *Client, f *InspectTemplateInspectConfigRuleSetRules, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, f.HotwordRule, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRule into hotwordRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRule"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, f.ExclusionRule, res); err != nil {
		return nil, fmt.Errorf("error expanding ExclusionRule into exclusionRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["exclusionRule"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRules flattens an instance of InspectTemplateInspectConfigRuleSetRules from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRules(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRules {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRules{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRules
	}
	r.HotwordRule = flattenInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, m["hotwordRule"], res)
	r.ExclusionRule = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, m["exclusionRule"], res)

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleMap expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRule, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesHotwordRule, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRule)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRule expands an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRule(c *Client, f *InspectTemplateInspectConfigRuleSetRulesHotwordRule, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, f.HotwordRegex, res); err != nil {
		return nil, fmt.Errorf("error expanding HotwordRegex into hotwordRegex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["hotwordRegex"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, f.Proximity, res); err != nil {
		return nil, fmt.Errorf("error expanding Proximity into proximity: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["proximity"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, f.LikelihoodAdjustment, res); err != nil {
		return nil, fmt.Errorf("error expanding LikelihoodAdjustment into likelihoodAdjustment: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["likelihoodAdjustment"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRule flattens an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRule(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesHotwordRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesHotwordRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRule
	}
	r.HotwordRegex = flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, m["hotwordRegex"], res)
	r.Proximity = flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, m["proximity"], res)
	r.LikelihoodAdjustment = flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, m["likelihoodAdjustment"], res)

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex expands an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, f *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex flattens an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityMap expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximitySlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity expands an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, f *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.WindowBefore; !dcl.IsEmptyValueIndirect(v) {
		m["windowBefore"] = v
	}
	if v := f.WindowAfter; !dcl.IsEmptyValueIndirect(v) {
		m["windowAfter"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity flattens an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity
	}
	r.WindowBefore = dcl.FlattenInteger(m["windowBefore"])
	r.WindowAfter = dcl.FlattenInteger(m["windowAfter"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment expands an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, f *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FixedLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["fixedLikelihood"] = v
	}
	if v := f.RelativeLikelihood; !dcl.IsEmptyValueIndirect(v) {
		m["relativeLikelihood"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment flattens an instance of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment
	}
	r.FixedLikelihood = flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(m["fixedLikelihood"])
	r.RelativeLikelihood = dcl.FlattenInteger(m["relativeLikelihood"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRule, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRule, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRule)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRule expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRule into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRule(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRule, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dictionary"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, f.ExcludeInfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding ExcludeInfoTypes into excludeInfoTypes: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["excludeInfoTypes"] = v
	}
	if v := f.MatchingType; !dcl.IsEmptyValueIndirect(v) {
		m["matchingType"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRule flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRule from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRule(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRule
	}
	r.Dictionary = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, m["dictionary"], res)
	r.Regex = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, m["regex"], res)
	r.ExcludeInfoTypes = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, m["excludeInfoTypes"], res)
	r.MatchingType = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(m["matchingType"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionarySlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, f.WordList, res); err != nil {
		return nil, fmt.Errorf("error expanding WordList into wordList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["wordList"] = v
	}
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, f.CloudStoragePath, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStoragePath into cloudStoragePath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStoragePath"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary
	}
	r.WordList = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, m["wordList"], res)
	r.CloudStoragePath = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, m["cloudStoragePath"], res)

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Words; v != nil {
		m["words"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList
	}
	r.Words = dcl.FlattenStringSlice(m["words"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Pattern; !dcl.IsEmptyValueIndirect(v) {
		m["pattern"] = v
	}
	if v := f.GroupIndexes; v != nil {
		m["groupIndexes"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, f.InfoTypes, res); err != nil {
		return nil, fmt.Errorf("error expanding InfoTypes into infoTypes: %w", err)
	} else if v != nil {
		m["infoTypes"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes
	}
	r.InfoTypes = flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c, m["infoTypes"], res)

	return r
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap(c *Client, f map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice expands the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, f []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *InspectTemplate) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes expands an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes into a JSON
// request object.
func expandInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, f *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes, res *InspectTemplate) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes flattens an instance of InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(c *Client, i interface{}, res *InspectTemplate) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// flattenInspectTemplateInspectConfigMinLikelihoodEnumMap flattens the contents of InspectTemplateInspectConfigMinLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigMinLikelihoodEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigMinLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigMinLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigMinLikelihoodEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigMinLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigMinLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigMinLikelihoodEnumSlice flattens the contents of InspectTemplateInspectConfigMinLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigMinLikelihoodEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigMinLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigMinLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigMinLikelihoodEnum{}
	}

	items := make([]InspectTemplateInspectConfigMinLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigMinLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigMinLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigMinLikelihoodEnum with the same value as that string.
func flattenInspectTemplateInspectConfigMinLikelihoodEnum(i interface{}) *InspectTemplateInspectConfigMinLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigMinLikelihoodEnumRef(s)
}

// flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum with the same value as that string.
func flattenInspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum(i interface{}) *InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigCustomInfoTypesLikelihoodEnumRef(s)
}

// flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumMap flattens the contents of InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumSlice flattens the contents of InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum{}
	}

	items := make([]InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum with the same value as that string.
func flattenInspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum(i interface{}) *InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigCustomInfoTypesExclusionTypeEnumRef(s)
}

// flattenInspectTemplateInspectConfigContentOptionsEnumMap flattens the contents of InspectTemplateInspectConfigContentOptionsEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigContentOptionsEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigContentOptionsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigContentOptionsEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigContentOptionsEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigContentOptionsEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigContentOptionsEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigContentOptionsEnumSlice flattens the contents of InspectTemplateInspectConfigContentOptionsEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigContentOptionsEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigContentOptionsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigContentOptionsEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigContentOptionsEnum{}
	}

	items := make([]InspectTemplateInspectConfigContentOptionsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigContentOptionsEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigContentOptionsEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigContentOptionsEnum with the same value as that string.
func flattenInspectTemplateInspectConfigContentOptionsEnum(i interface{}) *InspectTemplateInspectConfigContentOptionsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigContentOptionsEnumRef(s)
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum with the same value as that string.
func flattenInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum(i interface{}) *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihoodEnumRef(s)
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumMap flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumMap(c *Client, i interface{}, res *InspectTemplate) map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	items := make(map[string]InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum)
	for k, item := range a {
		items[k] = *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(item.(interface{}))
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumSlice flattens the contents of InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum from a JSON
// response object.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumSlice(c *Client, i interface{}, res *InspectTemplate) []InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	if len(a) == 0 {
		return []InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum{}
	}

	items := make([]InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(item.(interface{})))
	}

	return items
}

// flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum asserts that an interface is a string, and returns a
// pointer to a *InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum with the same value as that string.
func flattenInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum(i interface{}) *InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return InspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingTypeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *InspectTemplate) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalInspectTemplate(b, c, r)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
			return false
		}
		if nr.Parent == nil && ncr.Parent == nil {
			c.Config.Logger.Info("Both Parent fields null - considering equal.")
		} else if nr.Parent == nil || ncr.Parent == nil {
			c.Config.Logger.Info("Only one Parent field is null - considering unequal.")
			return false
		} else if *nr.Parent != *ncr.Parent {
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

type inspectTemplateDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         inspectTemplateApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToInspectTemplateDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]inspectTemplateDiff, error) {
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
	var diffs []inspectTemplateDiff
	// For each operation name, create a inspectTemplateDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := inspectTemplateDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToInspectTemplateApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToInspectTemplateApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (inspectTemplateApiOperation, error) {
	switch opName {

	case "updateInspectTemplateUpdateInspectTemplateOperation":
		return &updateInspectTemplateUpdateInspectTemplateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractInspectTemplateFields(r *InspectTemplate) error {
	vInspectConfig := r.InspectConfig
	if vInspectConfig == nil {
		// note: explicitly not the empty object.
		vInspectConfig = &InspectTemplateInspectConfig{}
	}
	if err := extractInspectTemplateInspectConfigFields(r, vInspectConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectConfig) {
		r.InspectConfig = vInspectConfig
	}
	return nil
}
func extractInspectTemplateInspectConfigFields(r *InspectTemplate, o *InspectTemplateInspectConfig) error {
	vLimits := o.Limits
	if vLimits == nil {
		// note: explicitly not the empty object.
		vLimits = &InspectTemplateInspectConfigLimits{}
	}
	if err := extractInspectTemplateInspectConfigLimitsFields(r, vLimits); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLimits) {
		o.Limits = vLimits
	}
	return nil
}
func extractInspectTemplateInspectConfigInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigInfoTypes) error {
	return nil
}
func extractInspectTemplateInspectConfigLimitsFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimits) error {
	return nil
}
func extractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}
	if err := extractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	return nil
}
func extractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypes) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vSurrogateType := o.SurrogateType
	if vSurrogateType == nil {
		// note: explicitly not the empty object.
		vSurrogateType = &InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeFields(r, vSurrogateType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSurrogateType) {
		o.SurrogateType = vSurrogateType
	}
	vStoredType := o.StoredType
	if vStoredType == nil {
		// note: explicitly not the empty object.
		vStoredType = &InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesStoredTypeFields(r, vStoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStoredType) {
		o.StoredType = vStoredType
	}
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesInfoType) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesDictionaryFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesRegex) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) error {
	return nil
}
func extractInspectTemplateInspectConfigCustomInfoTypesStoredTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesStoredType) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSet) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetInfoTypes) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	vExclusionRule := o.ExclusionRule
	if vExclusionRule == nil {
		// note: explicitly not the empty object.
		vExclusionRule = &InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleFields(r, vExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExclusionRule) {
		o.ExclusionRule = vExclusionRule
	}
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRule) error {
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vExcludeInfoTypes := o.ExcludeInfoTypes
	if vExcludeInfoTypes == nil {
		// note: explicitly not the empty object.
		vExcludeInfoTypes = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r, vExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExcludeInfoTypes) {
		o.ExcludeInfoTypes = vExcludeInfoTypes
	}
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) error {
	return nil
}
func extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) error {
	return nil
}

func postReadExtractInspectTemplateFields(r *InspectTemplate) error {
	vInspectConfig := r.InspectConfig
	if vInspectConfig == nil {
		// note: explicitly not the empty object.
		vInspectConfig = &InspectTemplateInspectConfig{}
	}
	if err := postReadExtractInspectTemplateInspectConfigFields(r, vInspectConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInspectConfig) {
		r.InspectConfig = vInspectConfig
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigFields(r *InspectTemplate, o *InspectTemplateInspectConfig) error {
	vLimits := o.Limits
	if vLimits == nil {
		// note: explicitly not the empty object.
		vLimits = &InspectTemplateInspectConfigLimits{}
	}
	if err := extractInspectTemplateInspectConfigLimitsFields(r, vLimits); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLimits) {
		o.Limits = vLimits
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigInfoTypes) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigLimitsFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimits) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoType) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType{}
	}
	if err := extractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypes) error {
	vInfoType := o.InfoType
	if vInfoType == nil {
		// note: explicitly not the empty object.
		vInfoType = &InspectTemplateInspectConfigCustomInfoTypesInfoType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesInfoTypeFields(r, vInfoType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vInfoType) {
		o.InfoType = vInfoType
	}
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &InspectTemplateInspectConfigCustomInfoTypesDictionary{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &InspectTemplateInspectConfigCustomInfoTypesRegex{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vSurrogateType := o.SurrogateType
	if vSurrogateType == nil {
		// note: explicitly not the empty object.
		vSurrogateType = &InspectTemplateInspectConfigCustomInfoTypesSurrogateType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeFields(r, vSurrogateType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSurrogateType) {
		o.SurrogateType = vSurrogateType
	}
	vStoredType := o.StoredType
	if vStoredType == nil {
		// note: explicitly not the empty object.
		vStoredType = &InspectTemplateInspectConfigCustomInfoTypesStoredType{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesStoredTypeFields(r, vStoredType); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vStoredType) {
		o.StoredType = vStoredType
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesInfoTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesInfoType) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesDictionaryFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath{}
	}
	if err := extractInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionaryWordList) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesRegex) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesSurrogateType) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigCustomInfoTypesStoredTypeFields(r *InspectTemplate, o *InspectTemplateInspectConfigCustomInfoTypesStoredType) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSet) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetInfoTypes) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRules) error {
	vHotwordRule := o.HotwordRule
	if vHotwordRule == nil {
		// note: explicitly not the empty object.
		vHotwordRule = &InspectTemplateInspectConfigRuleSetRulesHotwordRule{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleFields(r, vHotwordRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRule) {
		o.HotwordRule = vHotwordRule
	}
	vExclusionRule := o.ExclusionRule
	if vExclusionRule == nil {
		// note: explicitly not the empty object.
		vExclusionRule = &InspectTemplateInspectConfigRuleSetRulesExclusionRule{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleFields(r, vExclusionRule); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExclusionRule) {
		o.ExclusionRule = vExclusionRule
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesHotwordRuleFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRule) error {
	vHotwordRegex := o.HotwordRegex
	if vHotwordRegex == nil {
		// note: explicitly not the empty object.
		vHotwordRegex = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r, vHotwordRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHotwordRegex) {
		o.HotwordRegex = vHotwordRegex
	}
	vProximity := o.Proximity
	if vProximity == nil {
		// note: explicitly not the empty object.
		vProximity = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityFields(r, vProximity); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vProximity) {
		o.Proximity = vProximity
	}
	vLikelihoodAdjustment := o.LikelihoodAdjustment
	if vLikelihoodAdjustment == nil {
		// note: explicitly not the empty object.
		vLikelihoodAdjustment = &InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r, vLikelihoodAdjustment); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLikelihoodAdjustment) {
		o.LikelihoodAdjustment = vLikelihoodAdjustment
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRule) error {
	vDictionary := o.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		o.Dictionary = vDictionary
	}
	vRegex := o.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		o.Regex = vRegex
	}
	vExcludeInfoTypes := o.ExcludeInfoTypes
	if vExcludeInfoTypes == nil {
		// note: explicitly not the empty object.
		vExcludeInfoTypes = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r, vExcludeInfoTypes); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vExcludeInfoTypes) {
		o.ExcludeInfoTypes = vExcludeInfoTypes
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath{}
	}
	if err := extractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes) error {
	return nil
}
func postReadExtractInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesFields(r *InspectTemplate, o *InspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes) error {
	return nil
}
