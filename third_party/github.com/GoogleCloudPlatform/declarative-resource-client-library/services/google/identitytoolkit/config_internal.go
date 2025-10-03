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
package identitytoolkit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Config) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.SignIn) {
		if err := r.SignIn.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Notification) {
		if err := r.Notification.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Quota) {
		if err := r.Quota.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Monitoring) {
		if err := r.Monitoring.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.MultiTenant) {
		if err := r.MultiTenant.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Client) {
		if err := r.Client.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Mfa) {
		if err := r.Mfa.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BlockingFunctions) {
		if err := r.BlockingFunctions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigSignIn) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Email) {
		if err := r.Email.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PhoneNumber) {
		if err := r.PhoneNumber.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Anonymous) {
		if err := r.Anonymous.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HashConfig) {
		if err := r.HashConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigSignInEmail) validate() error {
	if !dcl.IsEmptyValueIndirect(r.HashConfig) {
		if err := r.HashConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigSignInEmailHashConfig) validate() error {
	return nil
}
func (r *ConfigSignInPhoneNumber) validate() error {
	return nil
}
func (r *ConfigSignInAnonymous) validate() error {
	return nil
}
func (r *ConfigSignInHashConfig) validate() error {
	return nil
}
func (r *ConfigNotification) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SendEmail) {
		if err := r.SendEmail.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.SendSms) {
		if err := r.SendSms.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigNotificationSendEmail) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Smtp) {
		if err := r.Smtp.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ResetPasswordTemplate) {
		if err := r.ResetPasswordTemplate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.VerifyEmailTemplate) {
		if err := r.VerifyEmailTemplate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.ChangeEmailTemplate) {
		if err := r.ChangeEmailTemplate.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DnsInfo) {
		if err := r.DnsInfo.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.RevertSecondFactorAdditionTemplate) {
		if err := r.RevertSecondFactorAdditionTemplate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigNotificationSendEmailSmtp) validate() error {
	return nil
}
func (r *ConfigNotificationSendEmailResetPasswordTemplate) validate() error {
	return nil
}
func (r *ConfigNotificationSendEmailVerifyEmailTemplate) validate() error {
	return nil
}
func (r *ConfigNotificationSendEmailChangeEmailTemplate) validate() error {
	return nil
}
func (r *ConfigNotificationSendEmailDnsInfo) validate() error {
	return nil
}
func (r *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) validate() error {
	return nil
}
func (r *ConfigNotificationSendSms) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SmsTemplate) {
		if err := r.SmsTemplate.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigNotificationSendSmsSmsTemplate) validate() error {
	return nil
}
func (r *ConfigQuota) validate() error {
	if !dcl.IsEmptyValueIndirect(r.SignUpQuotaConfig) {
		if err := r.SignUpQuotaConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigQuotaSignUpQuotaConfig) validate() error {
	return nil
}
func (r *ConfigMonitoring) validate() error {
	if !dcl.IsEmptyValueIndirect(r.RequestLogging) {
		if err := r.RequestLogging.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigMonitoringRequestLogging) validate() error {
	return nil
}
func (r *ConfigMultiTenant) validate() error {
	return nil
}
func (r *ConfigClient) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Permissions) {
		if err := r.Permissions.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConfigClientPermissions) validate() error {
	return nil
}
func (r *ConfigMfa) validate() error {
	return nil
}
func (r *ConfigBlockingFunctions) validate() error {
	return nil
}
func (r *ConfigBlockingFunctionsTriggers) validate() error {
	return nil
}
func (r *Config) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://identitytoolkit.googleapis.com/admin/v2", params)
}

func (r *Config) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/config", nr.basePath(), userBasePath, params), nil
}

// configApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type configApiOperation interface {
	do(context.Context, *Config, *Client) error
}

// newUpdateConfigUpdateProjectConfigRequest creates a request for an
// Config resource's UpdateProjectConfig update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateConfigUpdateProjectConfigRequest(ctx context.Context, f *Config, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandConfigSignIn(c, f.SignIn, res); err != nil {
		return nil, fmt.Errorf("error expanding SignIn into signIn: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["signIn"] = v
	}
	if v, err := expandConfigNotification(c, f.Notification, res); err != nil {
		return nil, fmt.Errorf("error expanding Notification into notification: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["notification"] = v
	}
	if v, err := expandConfigQuota(c, f.Quota, res); err != nil {
		return nil, fmt.Errorf("error expanding Quota into quota: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["quota"] = v
	}
	if v, err := expandConfigMonitoring(c, f.Monitoring, res); err != nil {
		return nil, fmt.Errorf("error expanding Monitoring into monitoring: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["monitoring"] = v
	}
	if v, err := expandConfigMultiTenant(c, f.MultiTenant, res); err != nil {
		return nil, fmt.Errorf("error expanding MultiTenant into multiTenant: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["multiTenant"] = v
	}
	if v := f.AuthorizedDomains; v != nil {
		req["authorizedDomains"] = v
	}
	if v, err := expandConfigClient(c, f.Client, res); err != nil {
		return nil, fmt.Errorf("error expanding Client into client: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["client"] = v
	}
	if v, err := expandConfigMfa(c, f.Mfa, res); err != nil {
		return nil, fmt.Errorf("error expanding Mfa into mfa: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mfa"] = v
	}
	if v, err := expandConfigBlockingFunctions(c, f.BlockingFunctions, res); err != nil {
		return nil, fmt.Errorf("error expanding BlockingFunctions into blockingFunctions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["blockingFunctions"] = v
	}
	return req, nil
}

// marshalUpdateConfigUpdateProjectConfigRequest converts the update into
// the final JSON request body.
func marshalUpdateConfigUpdateProjectConfigRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateConfigUpdateProjectConfigOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateConfigUpdateProjectConfigOperation) do(ctx context.Context, r *Config, c *Client) error {
	_, err := c.GetConfig(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateProjectConfig")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateConfigUpdateProjectConfigRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateConfigUpdateProjectConfigRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createConfigOperation struct {
	response map[string]interface{}
}

func (op *createConfigOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getConfigRaw(ctx context.Context, r *Config) ([]byte, error) {

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

func (c *Client) configDiffsForRawDesired(ctx context.Context, rawDesired *Config, opts ...dcl.ApplyOption) (initial, desired *Config, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Config
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Config); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Config, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetConfig(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Config resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Config resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Config resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeConfigDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Config: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Config: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractConfigFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeConfigInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Config: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeConfigDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Config: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffConfig(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeConfigInitialState(rawInitial, rawDesired *Config) (*Config, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeConfigDesiredState(rawDesired, rawInitial *Config, opts ...dcl.ApplyOption) (*Config, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.SignIn = canonicalizeConfigSignIn(rawDesired.SignIn, nil, opts...)
		rawDesired.Notification = canonicalizeConfigNotification(rawDesired.Notification, nil, opts...)
		rawDesired.Quota = canonicalizeConfigQuota(rawDesired.Quota, nil, opts...)
		rawDesired.Monitoring = canonicalizeConfigMonitoring(rawDesired.Monitoring, nil, opts...)
		rawDesired.MultiTenant = canonicalizeConfigMultiTenant(rawDesired.MultiTenant, nil, opts...)
		rawDesired.Client = canonicalizeConfigClient(rawDesired.Client, nil, opts...)
		rawDesired.Mfa = canonicalizeConfigMfa(rawDesired.Mfa, nil, opts...)
		rawDesired.BlockingFunctions = canonicalizeConfigBlockingFunctions(rawDesired.BlockingFunctions, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Config{}
	canonicalDesired.SignIn = canonicalizeConfigSignIn(rawDesired.SignIn, rawInitial.SignIn, opts...)
	canonicalDesired.Notification = canonicalizeConfigNotification(rawDesired.Notification, rawInitial.Notification, opts...)
	canonicalDesired.Quota = canonicalizeConfigQuota(rawDesired.Quota, rawInitial.Quota, opts...)
	canonicalDesired.Monitoring = canonicalizeConfigMonitoring(rawDesired.Monitoring, rawInitial.Monitoring, opts...)
	canonicalDesired.MultiTenant = canonicalizeConfigMultiTenant(rawDesired.MultiTenant, rawInitial.MultiTenant, opts...)
	if dcl.StringArrayCanonicalize(rawDesired.AuthorizedDomains, rawInitial.AuthorizedDomains) {
		canonicalDesired.AuthorizedDomains = rawInitial.AuthorizedDomains
	} else {
		canonicalDesired.AuthorizedDomains = rawDesired.AuthorizedDomains
	}
	canonicalDesired.Client = canonicalizeConfigClient(rawDesired.Client, rawInitial.Client, opts...)
	canonicalDesired.Mfa = canonicalizeConfigMfa(rawDesired.Mfa, rawInitial.Mfa, opts...)
	canonicalDesired.BlockingFunctions = canonicalizeConfigBlockingFunctions(rawDesired.BlockingFunctions, rawInitial.BlockingFunctions, opts...)
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeConfigNewState(c *Client, rawNew, rawDesired *Config) (*Config, error) {

	if dcl.IsEmptyValueIndirect(rawNew.SignIn) && dcl.IsEmptyValueIndirect(rawDesired.SignIn) {
		rawNew.SignIn = rawDesired.SignIn
	} else {
		rawNew.SignIn = canonicalizeNewConfigSignIn(c, rawDesired.SignIn, rawNew.SignIn)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Notification) && dcl.IsEmptyValueIndirect(rawDesired.Notification) {
		rawNew.Notification = rawDesired.Notification
	} else {
		rawNew.Notification = canonicalizeNewConfigNotification(c, rawDesired.Notification, rawNew.Notification)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Quota) && dcl.IsEmptyValueIndirect(rawDesired.Quota) {
		rawNew.Quota = rawDesired.Quota
	} else {
		rawNew.Quota = canonicalizeNewConfigQuota(c, rawDesired.Quota, rawNew.Quota)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Monitoring) && dcl.IsEmptyValueIndirect(rawDesired.Monitoring) {
		rawNew.Monitoring = rawDesired.Monitoring
	} else {
		rawNew.Monitoring = canonicalizeNewConfigMonitoring(c, rawDesired.Monitoring, rawNew.Monitoring)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MultiTenant) && dcl.IsEmptyValueIndirect(rawDesired.MultiTenant) {
		rawNew.MultiTenant = rawDesired.MultiTenant
	} else {
		rawNew.MultiTenant = canonicalizeNewConfigMultiTenant(c, rawDesired.MultiTenant, rawNew.MultiTenant)
	}

	if dcl.IsEmptyValueIndirect(rawNew.AuthorizedDomains) && dcl.IsEmptyValueIndirect(rawDesired.AuthorizedDomains) {
		rawNew.AuthorizedDomains = rawDesired.AuthorizedDomains
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.AuthorizedDomains, rawNew.AuthorizedDomains) {
			rawNew.AuthorizedDomains = rawDesired.AuthorizedDomains
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subtype) && dcl.IsEmptyValueIndirect(rawDesired.Subtype) {
		rawNew.Subtype = rawDesired.Subtype
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Client) && dcl.IsEmptyValueIndirect(rawDesired.Client) {
		rawNew.Client = rawDesired.Client
	} else {
		rawNew.Client = canonicalizeNewConfigClient(c, rawDesired.Client, rawNew.Client)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Mfa) && dcl.IsEmptyValueIndirect(rawDesired.Mfa) {
		rawNew.Mfa = rawDesired.Mfa
	} else {
		rawNew.Mfa = canonicalizeNewConfigMfa(c, rawDesired.Mfa, rawNew.Mfa)
	}

	if dcl.IsEmptyValueIndirect(rawNew.BlockingFunctions) && dcl.IsEmptyValueIndirect(rawDesired.BlockingFunctions) {
		rawNew.BlockingFunctions = rawDesired.BlockingFunctions
	} else {
		rawNew.BlockingFunctions = canonicalizeNewConfigBlockingFunctions(c, rawDesired.BlockingFunctions, rawNew.BlockingFunctions)
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeConfigSignIn(des, initial *ConfigSignIn, opts ...dcl.ApplyOption) *ConfigSignIn {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignIn{}

	cDes.Email = canonicalizeConfigSignInEmail(des.Email, initial.Email, opts...)
	cDes.PhoneNumber = canonicalizeConfigSignInPhoneNumber(des.PhoneNumber, initial.PhoneNumber, opts...)
	cDes.Anonymous = canonicalizeConfigSignInAnonymous(des.Anonymous, initial.Anonymous, opts...)
	if dcl.BoolCanonicalize(des.AllowDuplicateEmails, initial.AllowDuplicateEmails) || dcl.IsZeroValue(des.AllowDuplicateEmails) {
		cDes.AllowDuplicateEmails = initial.AllowDuplicateEmails
	} else {
		cDes.AllowDuplicateEmails = des.AllowDuplicateEmails
	}

	return cDes
}

func canonicalizeConfigSignInSlice(des, initial []ConfigSignIn, opts ...dcl.ApplyOption) []ConfigSignIn {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignIn, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignIn(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignIn, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignIn(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignIn(c *Client, des, nw *ConfigSignIn) *ConfigSignIn {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignIn while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Email = canonicalizeNewConfigSignInEmail(c, des.Email, nw.Email)
	nw.PhoneNumber = canonicalizeNewConfigSignInPhoneNumber(c, des.PhoneNumber, nw.PhoneNumber)
	nw.Anonymous = canonicalizeNewConfigSignInAnonymous(c, des.Anonymous, nw.Anonymous)
	if dcl.BoolCanonicalize(des.AllowDuplicateEmails, nw.AllowDuplicateEmails) {
		nw.AllowDuplicateEmails = des.AllowDuplicateEmails
	}
	nw.HashConfig = canonicalizeNewConfigSignInHashConfig(c, des.HashConfig, nw.HashConfig)

	return nw
}

func canonicalizeNewConfigSignInSet(c *Client, des, nw []ConfigSignIn) []ConfigSignIn {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignIn
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignIn(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInSlice(c *Client, des, nw []ConfigSignIn) []ConfigSignIn {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignIn
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignIn(c, &d, &n))
	}

	return items
}

func canonicalizeConfigSignInEmail(des, initial *ConfigSignInEmail, opts ...dcl.ApplyOption) *ConfigSignInEmail {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignInEmail{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.PasswordRequired, initial.PasswordRequired) || dcl.IsZeroValue(des.PasswordRequired) {
		cDes.PasswordRequired = initial.PasswordRequired
	} else {
		cDes.PasswordRequired = des.PasswordRequired
	}

	return cDes
}

func canonicalizeConfigSignInEmailSlice(des, initial []ConfigSignInEmail, opts ...dcl.ApplyOption) []ConfigSignInEmail {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignInEmail, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignInEmail(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignInEmail, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignInEmail(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignInEmail(c *Client, des, nw *ConfigSignInEmail) *ConfigSignInEmail {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignInEmail while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.PasswordRequired, nw.PasswordRequired) {
		nw.PasswordRequired = des.PasswordRequired
	}
	nw.HashConfig = canonicalizeNewConfigSignInEmailHashConfig(c, des.HashConfig, nw.HashConfig)

	return nw
}

func canonicalizeNewConfigSignInEmailSet(c *Client, des, nw []ConfigSignInEmail) []ConfigSignInEmail {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignInEmail
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInEmailNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignInEmail(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInEmailSlice(c *Client, des, nw []ConfigSignInEmail) []ConfigSignInEmail {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignInEmail
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignInEmail(c, &d, &n))
	}

	return items
}

func canonicalizeConfigSignInEmailHashConfig(des, initial *ConfigSignInEmailHashConfig, opts ...dcl.ApplyOption) *ConfigSignInEmailHashConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignInEmailHashConfig{}

	return cDes
}

func canonicalizeConfigSignInEmailHashConfigSlice(des, initial []ConfigSignInEmailHashConfig, opts ...dcl.ApplyOption) []ConfigSignInEmailHashConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignInEmailHashConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignInEmailHashConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignInEmailHashConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignInEmailHashConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignInEmailHashConfig(c *Client, des, nw *ConfigSignInEmailHashConfig) *ConfigSignInEmailHashConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignInEmailHashConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SignerKey, nw.SignerKey) {
		nw.SignerKey = des.SignerKey
	}
	if dcl.StringCanonicalize(des.SaltSeparator, nw.SaltSeparator) {
		nw.SaltSeparator = des.SaltSeparator
	}

	return nw
}

func canonicalizeNewConfigSignInEmailHashConfigSet(c *Client, des, nw []ConfigSignInEmailHashConfig) []ConfigSignInEmailHashConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignInEmailHashConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInEmailHashConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignInEmailHashConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInEmailHashConfigSlice(c *Client, des, nw []ConfigSignInEmailHashConfig) []ConfigSignInEmailHashConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignInEmailHashConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignInEmailHashConfig(c, &d, &n))
	}

	return items
}

func canonicalizeConfigSignInPhoneNumber(des, initial *ConfigSignInPhoneNumber, opts ...dcl.ApplyOption) *ConfigSignInPhoneNumber {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignInPhoneNumber{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}
	if canonicalizeConfigTestPhoneNumbers(des.TestPhoneNumbers, initial.TestPhoneNumbers) || dcl.IsZeroValue(des.TestPhoneNumbers) {
		cDes.TestPhoneNumbers = initial.TestPhoneNumbers
	} else {
		cDes.TestPhoneNumbers = des.TestPhoneNumbers
	}

	return cDes
}

func canonicalizeConfigSignInPhoneNumberSlice(des, initial []ConfigSignInPhoneNumber, opts ...dcl.ApplyOption) []ConfigSignInPhoneNumber {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignInPhoneNumber, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignInPhoneNumber(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignInPhoneNumber, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignInPhoneNumber(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignInPhoneNumber(c *Client, des, nw *ConfigSignInPhoneNumber) *ConfigSignInPhoneNumber {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignInPhoneNumber while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if canonicalizeConfigTestPhoneNumbers(des.TestPhoneNumbers, nw.TestPhoneNumbers) {
		nw.TestPhoneNumbers = des.TestPhoneNumbers
	}

	return nw
}

func canonicalizeNewConfigSignInPhoneNumberSet(c *Client, des, nw []ConfigSignInPhoneNumber) []ConfigSignInPhoneNumber {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignInPhoneNumber
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInPhoneNumberNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignInPhoneNumber(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInPhoneNumberSlice(c *Client, des, nw []ConfigSignInPhoneNumber) []ConfigSignInPhoneNumber {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignInPhoneNumber
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignInPhoneNumber(c, &d, &n))
	}

	return items
}

func canonicalizeConfigSignInAnonymous(des, initial *ConfigSignInAnonymous, opts ...dcl.ApplyOption) *ConfigSignInAnonymous {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignInAnonymous{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeConfigSignInAnonymousSlice(des, initial []ConfigSignInAnonymous, opts ...dcl.ApplyOption) []ConfigSignInAnonymous {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignInAnonymous, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignInAnonymous(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignInAnonymous, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignInAnonymous(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignInAnonymous(c *Client, des, nw *ConfigSignInAnonymous) *ConfigSignInAnonymous {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignInAnonymous while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewConfigSignInAnonymousSet(c *Client, des, nw []ConfigSignInAnonymous) []ConfigSignInAnonymous {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignInAnonymous
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInAnonymousNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignInAnonymous(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInAnonymousSlice(c *Client, des, nw []ConfigSignInAnonymous) []ConfigSignInAnonymous {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignInAnonymous
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignInAnonymous(c, &d, &n))
	}

	return items
}

func canonicalizeConfigSignInHashConfig(des, initial *ConfigSignInHashConfig, opts ...dcl.ApplyOption) *ConfigSignInHashConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigSignInHashConfig{}

	return cDes
}

func canonicalizeConfigSignInHashConfigSlice(des, initial []ConfigSignInHashConfig, opts ...dcl.ApplyOption) []ConfigSignInHashConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigSignInHashConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigSignInHashConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigSignInHashConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigSignInHashConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigSignInHashConfig(c *Client, des, nw *ConfigSignInHashConfig) *ConfigSignInHashConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigSignInHashConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SignerKey, nw.SignerKey) {
		nw.SignerKey = des.SignerKey
	}
	if dcl.StringCanonicalize(des.SaltSeparator, nw.SaltSeparator) {
		nw.SaltSeparator = des.SaltSeparator
	}

	return nw
}

func canonicalizeNewConfigSignInHashConfigSet(c *Client, des, nw []ConfigSignInHashConfig) []ConfigSignInHashConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigSignInHashConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigSignInHashConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigSignInHashConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigSignInHashConfigSlice(c *Client, des, nw []ConfigSignInHashConfig) []ConfigSignInHashConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigSignInHashConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigSignInHashConfig(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotification(des, initial *ConfigNotification, opts ...dcl.ApplyOption) *ConfigNotification {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotification{}

	cDes.SendEmail = canonicalizeConfigNotificationSendEmail(des.SendEmail, initial.SendEmail, opts...)
	cDes.SendSms = canonicalizeConfigNotificationSendSms(des.SendSms, initial.SendSms, opts...)
	if dcl.StringCanonicalize(des.DefaultLocale, initial.DefaultLocale) || dcl.IsZeroValue(des.DefaultLocale) {
		cDes.DefaultLocale = initial.DefaultLocale
	} else {
		cDes.DefaultLocale = des.DefaultLocale
	}

	return cDes
}

func canonicalizeConfigNotificationSlice(des, initial []ConfigNotification, opts ...dcl.ApplyOption) []ConfigNotification {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotification, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotification(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotification, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotification(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotification(c *Client, des, nw *ConfigNotification) *ConfigNotification {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotification while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SendEmail = canonicalizeNewConfigNotificationSendEmail(c, des.SendEmail, nw.SendEmail)
	nw.SendSms = canonicalizeNewConfigNotificationSendSms(c, des.SendSms, nw.SendSms)
	if dcl.StringCanonicalize(des.DefaultLocale, nw.DefaultLocale) {
		nw.DefaultLocale = des.DefaultLocale
	}

	return nw
}

func canonicalizeNewConfigNotificationSet(c *Client, des, nw []ConfigNotification) []ConfigNotification {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotification
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotification(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSlice(c *Client, des, nw []ConfigNotification) []ConfigNotification {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotification
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotification(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmail(des, initial *ConfigNotificationSendEmail, opts ...dcl.ApplyOption) *ConfigNotificationSendEmail {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmail{}

	if dcl.IsZeroValue(des.Method) || (dcl.IsEmptyValueIndirect(des.Method) && dcl.IsEmptyValueIndirect(initial.Method)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Method = initial.Method
	} else {
		cDes.Method = des.Method
	}
	cDes.Smtp = canonicalizeConfigNotificationSendEmailSmtp(des.Smtp, initial.Smtp, opts...)
	cDes.ResetPasswordTemplate = canonicalizeConfigNotificationSendEmailResetPasswordTemplate(des.ResetPasswordTemplate, initial.ResetPasswordTemplate, opts...)
	cDes.VerifyEmailTemplate = canonicalizeConfigNotificationSendEmailVerifyEmailTemplate(des.VerifyEmailTemplate, initial.VerifyEmailTemplate, opts...)
	cDes.ChangeEmailTemplate = canonicalizeConfigNotificationSendEmailChangeEmailTemplate(des.ChangeEmailTemplate, initial.ChangeEmailTemplate, opts...)
	if dcl.StringCanonicalize(des.CallbackUri, initial.CallbackUri) || dcl.IsZeroValue(des.CallbackUri) {
		cDes.CallbackUri = initial.CallbackUri
	} else {
		cDes.CallbackUri = des.CallbackUri
	}
	cDes.DnsInfo = canonicalizeConfigNotificationSendEmailDnsInfo(des.DnsInfo, initial.DnsInfo, opts...)
	cDes.RevertSecondFactorAdditionTemplate = canonicalizeConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(des.RevertSecondFactorAdditionTemplate, initial.RevertSecondFactorAdditionTemplate, opts...)

	return cDes
}

func canonicalizeConfigNotificationSendEmailSlice(des, initial []ConfigNotificationSendEmail, opts ...dcl.ApplyOption) []ConfigNotificationSendEmail {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmail, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmail(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmail, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmail(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmail(c *Client, des, nw *ConfigNotificationSendEmail) *ConfigNotificationSendEmail {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmail while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Smtp = canonicalizeNewConfigNotificationSendEmailSmtp(c, des.Smtp, nw.Smtp)
	nw.ResetPasswordTemplate = canonicalizeNewConfigNotificationSendEmailResetPasswordTemplate(c, des.ResetPasswordTemplate, nw.ResetPasswordTemplate)
	nw.VerifyEmailTemplate = canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplate(c, des.VerifyEmailTemplate, nw.VerifyEmailTemplate)
	nw.ChangeEmailTemplate = canonicalizeNewConfigNotificationSendEmailChangeEmailTemplate(c, des.ChangeEmailTemplate, nw.ChangeEmailTemplate)
	if dcl.StringCanonicalize(des.CallbackUri, nw.CallbackUri) {
		nw.CallbackUri = des.CallbackUri
	}
	nw.DnsInfo = canonicalizeNewConfigNotificationSendEmailDnsInfo(c, des.DnsInfo, nw.DnsInfo)
	nw.RevertSecondFactorAdditionTemplate = canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, des.RevertSecondFactorAdditionTemplate, nw.RevertSecondFactorAdditionTemplate)

	return nw
}

func canonicalizeNewConfigNotificationSendEmailSet(c *Client, des, nw []ConfigNotificationSendEmail) []ConfigNotificationSendEmail {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmail
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmail(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailSlice(c *Client, des, nw []ConfigNotificationSendEmail) []ConfigNotificationSendEmail {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmail
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmail(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailSmtp(des, initial *ConfigNotificationSendEmailSmtp, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailSmtp {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailSmtp{}

	if dcl.StringCanonicalize(des.SenderEmail, initial.SenderEmail) || dcl.IsZeroValue(des.SenderEmail) {
		cDes.SenderEmail = initial.SenderEmail
	} else {
		cDes.SenderEmail = des.SenderEmail
	}
	if dcl.StringCanonicalize(des.Host, initial.Host) || dcl.IsZeroValue(des.Host) {
		cDes.Host = initial.Host
	} else {
		cDes.Host = des.Host
	}
	if dcl.IsZeroValue(des.Port) || (dcl.IsEmptyValueIndirect(des.Port) && dcl.IsEmptyValueIndirect(initial.Port)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Port = initial.Port
	} else {
		cDes.Port = des.Port
	}
	if dcl.StringCanonicalize(des.Username, initial.Username) || dcl.IsZeroValue(des.Username) {
		cDes.Username = initial.Username
	} else {
		cDes.Username = des.Username
	}
	if dcl.StringCanonicalize(des.Password, initial.Password) || dcl.IsZeroValue(des.Password) {
		cDes.Password = initial.Password
	} else {
		cDes.Password = des.Password
	}
	if dcl.IsZeroValue(des.SecurityMode) || (dcl.IsEmptyValueIndirect(des.SecurityMode) && dcl.IsEmptyValueIndirect(initial.SecurityMode)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.SecurityMode = initial.SecurityMode
	} else {
		cDes.SecurityMode = des.SecurityMode
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailSmtpSlice(des, initial []ConfigNotificationSendEmailSmtp, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailSmtp {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailSmtp, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailSmtp(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailSmtp, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailSmtp(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailSmtp(c *Client, des, nw *ConfigNotificationSendEmailSmtp) *ConfigNotificationSendEmailSmtp {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailSmtp while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SenderEmail, nw.SenderEmail) {
		nw.SenderEmail = des.SenderEmail
	}
	if dcl.StringCanonicalize(des.Host, nw.Host) {
		nw.Host = des.Host
	}
	if dcl.StringCanonicalize(des.Username, nw.Username) {
		nw.Username = des.Username
	}
	nw.Password = des.Password

	return nw
}

func canonicalizeNewConfigNotificationSendEmailSmtpSet(c *Client, des, nw []ConfigNotificationSendEmailSmtp) []ConfigNotificationSendEmailSmtp {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailSmtp
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailSmtpNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailSmtp(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailSmtpSlice(c *Client, des, nw []ConfigNotificationSendEmailSmtp) []ConfigNotificationSendEmailSmtp {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailSmtp
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailSmtp(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailResetPasswordTemplate(des, initial *ConfigNotificationSendEmailResetPasswordTemplate, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailResetPasswordTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailResetPasswordTemplate{}

	if dcl.StringCanonicalize(des.SenderLocalPart, initial.SenderLocalPart) || dcl.IsZeroValue(des.SenderLocalPart) {
		cDes.SenderLocalPart = initial.SenderLocalPart
	} else {
		cDes.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, initial.Subject) || dcl.IsZeroValue(des.Subject) {
		cDes.Subject = initial.Subject
	} else {
		cDes.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, initial.SenderDisplayName) || dcl.IsZeroValue(des.SenderDisplayName) {
		cDes.SenderDisplayName = initial.SenderDisplayName
	} else {
		cDes.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}
	if dcl.IsZeroValue(des.BodyFormat) || (dcl.IsEmptyValueIndirect(des.BodyFormat) && dcl.IsEmptyValueIndirect(initial.BodyFormat)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BodyFormat = initial.BodyFormat
	} else {
		cDes.BodyFormat = des.BodyFormat
	}
	if dcl.StringCanonicalize(des.ReplyTo, initial.ReplyTo) || dcl.IsZeroValue(des.ReplyTo) {
		cDes.ReplyTo = initial.ReplyTo
	} else {
		cDes.ReplyTo = des.ReplyTo
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailResetPasswordTemplateSlice(des, initial []ConfigNotificationSendEmailResetPasswordTemplate, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailResetPasswordTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailResetPasswordTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailResetPasswordTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailResetPasswordTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailResetPasswordTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailResetPasswordTemplate(c *Client, des, nw *ConfigNotificationSendEmailResetPasswordTemplate) *ConfigNotificationSendEmailResetPasswordTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailResetPasswordTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SenderLocalPart, nw.SenderLocalPart) {
		nw.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, nw.Subject) {
		nw.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, nw.SenderDisplayName) {
		nw.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	if dcl.StringCanonicalize(des.ReplyTo, nw.ReplyTo) {
		nw.ReplyTo = des.ReplyTo
	}
	if dcl.BoolCanonicalize(des.Customized, nw.Customized) {
		nw.Customized = des.Customized
	}

	return nw
}

func canonicalizeNewConfigNotificationSendEmailResetPasswordTemplateSet(c *Client, des, nw []ConfigNotificationSendEmailResetPasswordTemplate) []ConfigNotificationSendEmailResetPasswordTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailResetPasswordTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailResetPasswordTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailResetPasswordTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailResetPasswordTemplateSlice(c *Client, des, nw []ConfigNotificationSendEmailResetPasswordTemplate) []ConfigNotificationSendEmailResetPasswordTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailResetPasswordTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailResetPasswordTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailVerifyEmailTemplate(des, initial *ConfigNotificationSendEmailVerifyEmailTemplate, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailVerifyEmailTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailVerifyEmailTemplate{}

	if dcl.StringCanonicalize(des.SenderLocalPart, initial.SenderLocalPart) || dcl.IsZeroValue(des.SenderLocalPart) {
		cDes.SenderLocalPart = initial.SenderLocalPart
	} else {
		cDes.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, initial.Subject) || dcl.IsZeroValue(des.Subject) {
		cDes.Subject = initial.Subject
	} else {
		cDes.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, initial.SenderDisplayName) || dcl.IsZeroValue(des.SenderDisplayName) {
		cDes.SenderDisplayName = initial.SenderDisplayName
	} else {
		cDes.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}
	if dcl.IsZeroValue(des.BodyFormat) || (dcl.IsEmptyValueIndirect(des.BodyFormat) && dcl.IsEmptyValueIndirect(initial.BodyFormat)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BodyFormat = initial.BodyFormat
	} else {
		cDes.BodyFormat = des.BodyFormat
	}
	if dcl.StringCanonicalize(des.ReplyTo, initial.ReplyTo) || dcl.IsZeroValue(des.ReplyTo) {
		cDes.ReplyTo = initial.ReplyTo
	} else {
		cDes.ReplyTo = des.ReplyTo
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailVerifyEmailTemplateSlice(des, initial []ConfigNotificationSendEmailVerifyEmailTemplate, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailVerifyEmailTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailVerifyEmailTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailVerifyEmailTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailVerifyEmailTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailVerifyEmailTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplate(c *Client, des, nw *ConfigNotificationSendEmailVerifyEmailTemplate) *ConfigNotificationSendEmailVerifyEmailTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailVerifyEmailTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SenderLocalPart, nw.SenderLocalPart) {
		nw.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, nw.Subject) {
		nw.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, nw.SenderDisplayName) {
		nw.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	if dcl.StringCanonicalize(des.ReplyTo, nw.ReplyTo) {
		nw.ReplyTo = des.ReplyTo
	}
	if dcl.BoolCanonicalize(des.Customized, nw.Customized) {
		nw.Customized = des.Customized
	}

	return nw
}

func canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplateSet(c *Client, des, nw []ConfigNotificationSendEmailVerifyEmailTemplate) []ConfigNotificationSendEmailVerifyEmailTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailVerifyEmailTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailVerifyEmailTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplateSlice(c *Client, des, nw []ConfigNotificationSendEmailVerifyEmailTemplate) []ConfigNotificationSendEmailVerifyEmailTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailVerifyEmailTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailVerifyEmailTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailChangeEmailTemplate(des, initial *ConfigNotificationSendEmailChangeEmailTemplate, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailChangeEmailTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailChangeEmailTemplate{}

	if dcl.StringCanonicalize(des.SenderLocalPart, initial.SenderLocalPart) || dcl.IsZeroValue(des.SenderLocalPart) {
		cDes.SenderLocalPart = initial.SenderLocalPart
	} else {
		cDes.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, initial.Subject) || dcl.IsZeroValue(des.Subject) {
		cDes.Subject = initial.Subject
	} else {
		cDes.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, initial.SenderDisplayName) || dcl.IsZeroValue(des.SenderDisplayName) {
		cDes.SenderDisplayName = initial.SenderDisplayName
	} else {
		cDes.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}
	if dcl.IsZeroValue(des.BodyFormat) || (dcl.IsEmptyValueIndirect(des.BodyFormat) && dcl.IsEmptyValueIndirect(initial.BodyFormat)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BodyFormat = initial.BodyFormat
	} else {
		cDes.BodyFormat = des.BodyFormat
	}
	if dcl.StringCanonicalize(des.ReplyTo, initial.ReplyTo) || dcl.IsZeroValue(des.ReplyTo) {
		cDes.ReplyTo = initial.ReplyTo
	} else {
		cDes.ReplyTo = des.ReplyTo
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailChangeEmailTemplateSlice(des, initial []ConfigNotificationSendEmailChangeEmailTemplate, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailChangeEmailTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailChangeEmailTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailChangeEmailTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailChangeEmailTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailChangeEmailTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailChangeEmailTemplate(c *Client, des, nw *ConfigNotificationSendEmailChangeEmailTemplate) *ConfigNotificationSendEmailChangeEmailTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailChangeEmailTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SenderLocalPart, nw.SenderLocalPart) {
		nw.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, nw.Subject) {
		nw.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, nw.SenderDisplayName) {
		nw.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	if dcl.StringCanonicalize(des.ReplyTo, nw.ReplyTo) {
		nw.ReplyTo = des.ReplyTo
	}
	if dcl.BoolCanonicalize(des.Customized, nw.Customized) {
		nw.Customized = des.Customized
	}

	return nw
}

func canonicalizeNewConfigNotificationSendEmailChangeEmailTemplateSet(c *Client, des, nw []ConfigNotificationSendEmailChangeEmailTemplate) []ConfigNotificationSendEmailChangeEmailTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailChangeEmailTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailChangeEmailTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailChangeEmailTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailChangeEmailTemplateSlice(c *Client, des, nw []ConfigNotificationSendEmailChangeEmailTemplate) []ConfigNotificationSendEmailChangeEmailTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailChangeEmailTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailChangeEmailTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailDnsInfo(des, initial *ConfigNotificationSendEmailDnsInfo, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailDnsInfo {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailDnsInfo{}

	if dcl.BoolCanonicalize(des.UseCustomDomain, initial.UseCustomDomain) || dcl.IsZeroValue(des.UseCustomDomain) {
		cDes.UseCustomDomain = initial.UseCustomDomain
	} else {
		cDes.UseCustomDomain = des.UseCustomDomain
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailDnsInfoSlice(des, initial []ConfigNotificationSendEmailDnsInfo, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailDnsInfo {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailDnsInfo, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailDnsInfo(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailDnsInfo, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailDnsInfo(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailDnsInfo(c *Client, des, nw *ConfigNotificationSendEmailDnsInfo) *ConfigNotificationSendEmailDnsInfo {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailDnsInfo while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.CustomDomain, nw.CustomDomain) {
		nw.CustomDomain = des.CustomDomain
	}
	if dcl.BoolCanonicalize(des.UseCustomDomain, nw.UseCustomDomain) {
		nw.UseCustomDomain = des.UseCustomDomain
	}
	if dcl.StringCanonicalize(des.PendingCustomDomain, nw.PendingCustomDomain) {
		nw.PendingCustomDomain = des.PendingCustomDomain
	}

	return nw
}

func canonicalizeNewConfigNotificationSendEmailDnsInfoSet(c *Client, des, nw []ConfigNotificationSendEmailDnsInfo) []ConfigNotificationSendEmailDnsInfo {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailDnsInfo
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailDnsInfoNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailDnsInfo(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailDnsInfoSlice(c *Client, des, nw []ConfigNotificationSendEmailDnsInfo) []ConfigNotificationSendEmailDnsInfo {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailDnsInfo
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailDnsInfo(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(des, initial *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, opts ...dcl.ApplyOption) *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}

	if dcl.StringCanonicalize(des.SenderLocalPart, initial.SenderLocalPart) || dcl.IsZeroValue(des.SenderLocalPart) {
		cDes.SenderLocalPart = initial.SenderLocalPart
	} else {
		cDes.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, initial.Subject) || dcl.IsZeroValue(des.Subject) {
		cDes.Subject = initial.Subject
	} else {
		cDes.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, initial.SenderDisplayName) || dcl.IsZeroValue(des.SenderDisplayName) {
		cDes.SenderDisplayName = initial.SenderDisplayName
	} else {
		cDes.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, initial.Body) || dcl.IsZeroValue(des.Body) {
		cDes.Body = initial.Body
	} else {
		cDes.Body = des.Body
	}
	if dcl.IsZeroValue(des.BodyFormat) || (dcl.IsEmptyValueIndirect(des.BodyFormat) && dcl.IsEmptyValueIndirect(initial.BodyFormat)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.BodyFormat = initial.BodyFormat
	} else {
		cDes.BodyFormat = des.BodyFormat
	}
	if dcl.StringCanonicalize(des.ReplyTo, initial.ReplyTo) || dcl.IsZeroValue(des.ReplyTo) {
		cDes.ReplyTo = initial.ReplyTo
	} else {
		cDes.ReplyTo = des.ReplyTo
	}

	return cDes
}

func canonicalizeConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice(des, initial []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, opts ...dcl.ApplyOption) []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c *Client, des, nw *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SenderLocalPart, nw.SenderLocalPart) {
		nw.SenderLocalPart = des.SenderLocalPart
	}
	if dcl.StringCanonicalize(des.Subject, nw.Subject) {
		nw.Subject = des.Subject
	}
	if dcl.StringCanonicalize(des.SenderDisplayName, nw.SenderDisplayName) {
		nw.SenderDisplayName = des.SenderDisplayName
	}
	if dcl.StringCanonicalize(des.Body, nw.Body) {
		nw.Body = des.Body
	}
	if dcl.StringCanonicalize(des.ReplyTo, nw.ReplyTo) {
		nw.ReplyTo = des.ReplyTo
	}
	if dcl.BoolCanonicalize(des.Customized, nw.Customized) {
		nw.Customized = des.Customized
	}

	return nw
}

func canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSet(c *Client, des, nw []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendEmailRevertSecondFactorAdditionTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice(c *Client, des, nw []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendSms(des, initial *ConfigNotificationSendSms, opts ...dcl.ApplyOption) *ConfigNotificationSendSms {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendSms{}

	if dcl.BoolCanonicalize(des.UseDeviceLocale, initial.UseDeviceLocale) || dcl.IsZeroValue(des.UseDeviceLocale) {
		cDes.UseDeviceLocale = initial.UseDeviceLocale
	} else {
		cDes.UseDeviceLocale = des.UseDeviceLocale
	}

	return cDes
}

func canonicalizeConfigNotificationSendSmsSlice(des, initial []ConfigNotificationSendSms, opts ...dcl.ApplyOption) []ConfigNotificationSendSms {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendSms, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendSms(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendSms, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendSms(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendSms(c *Client, des, nw *ConfigNotificationSendSms) *ConfigNotificationSendSms {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendSms while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.UseDeviceLocale, nw.UseDeviceLocale) {
		nw.UseDeviceLocale = des.UseDeviceLocale
	}
	nw.SmsTemplate = canonicalizeNewConfigNotificationSendSmsSmsTemplate(c, des.SmsTemplate, nw.SmsTemplate)

	return nw
}

func canonicalizeNewConfigNotificationSendSmsSet(c *Client, des, nw []ConfigNotificationSendSms) []ConfigNotificationSendSms {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendSms
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendSmsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendSms(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendSmsSlice(c *Client, des, nw []ConfigNotificationSendSms) []ConfigNotificationSendSms {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendSms
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendSms(c, &d, &n))
	}

	return items
}

func canonicalizeConfigNotificationSendSmsSmsTemplate(des, initial *ConfigNotificationSendSmsSmsTemplate, opts ...dcl.ApplyOption) *ConfigNotificationSendSmsSmsTemplate {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigNotificationSendSmsSmsTemplate{}

	return cDes
}

func canonicalizeConfigNotificationSendSmsSmsTemplateSlice(des, initial []ConfigNotificationSendSmsSmsTemplate, opts ...dcl.ApplyOption) []ConfigNotificationSendSmsSmsTemplate {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigNotificationSendSmsSmsTemplate, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigNotificationSendSmsSmsTemplate(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigNotificationSendSmsSmsTemplate, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigNotificationSendSmsSmsTemplate(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigNotificationSendSmsSmsTemplate(c *Client, des, nw *ConfigNotificationSendSmsSmsTemplate) *ConfigNotificationSendSmsSmsTemplate {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigNotificationSendSmsSmsTemplate while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Content, nw.Content) {
		nw.Content = des.Content
	}

	return nw
}

func canonicalizeNewConfigNotificationSendSmsSmsTemplateSet(c *Client, des, nw []ConfigNotificationSendSmsSmsTemplate) []ConfigNotificationSendSmsSmsTemplate {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigNotificationSendSmsSmsTemplate
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigNotificationSendSmsSmsTemplateNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigNotificationSendSmsSmsTemplate(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigNotificationSendSmsSmsTemplateSlice(c *Client, des, nw []ConfigNotificationSendSmsSmsTemplate) []ConfigNotificationSendSmsSmsTemplate {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigNotificationSendSmsSmsTemplate
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigNotificationSendSmsSmsTemplate(c, &d, &n))
	}

	return items
}

func canonicalizeConfigQuota(des, initial *ConfigQuota, opts ...dcl.ApplyOption) *ConfigQuota {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigQuota{}

	cDes.SignUpQuotaConfig = canonicalizeConfigQuotaSignUpQuotaConfig(des.SignUpQuotaConfig, initial.SignUpQuotaConfig, opts...)

	return cDes
}

func canonicalizeConfigQuotaSlice(des, initial []ConfigQuota, opts ...dcl.ApplyOption) []ConfigQuota {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigQuota, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigQuota(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigQuota, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigQuota(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigQuota(c *Client, des, nw *ConfigQuota) *ConfigQuota {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigQuota while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.SignUpQuotaConfig = canonicalizeNewConfigQuotaSignUpQuotaConfig(c, des.SignUpQuotaConfig, nw.SignUpQuotaConfig)

	return nw
}

func canonicalizeNewConfigQuotaSet(c *Client, des, nw []ConfigQuota) []ConfigQuota {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigQuota
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigQuotaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigQuota(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigQuotaSlice(c *Client, des, nw []ConfigQuota) []ConfigQuota {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigQuota
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigQuota(c, &d, &n))
	}

	return items
}

func canonicalizeConfigQuotaSignUpQuotaConfig(des, initial *ConfigQuotaSignUpQuotaConfig, opts ...dcl.ApplyOption) *ConfigQuotaSignUpQuotaConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigQuotaSignUpQuotaConfig{}

	if dcl.IsZeroValue(des.Quota) || (dcl.IsEmptyValueIndirect(des.Quota) && dcl.IsEmptyValueIndirect(initial.Quota)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Quota = initial.Quota
	} else {
		cDes.Quota = des.Quota
	}
	if dcl.IsZeroValue(des.StartTime) || (dcl.IsEmptyValueIndirect(des.StartTime) && dcl.IsEmptyValueIndirect(initial.StartTime)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.StartTime = initial.StartTime
	} else {
		cDes.StartTime = des.StartTime
	}
	if dcl.StringCanonicalize(des.QuotaDuration, initial.QuotaDuration) || dcl.IsZeroValue(des.QuotaDuration) {
		cDes.QuotaDuration = initial.QuotaDuration
	} else {
		cDes.QuotaDuration = des.QuotaDuration
	}

	return cDes
}

func canonicalizeConfigQuotaSignUpQuotaConfigSlice(des, initial []ConfigQuotaSignUpQuotaConfig, opts ...dcl.ApplyOption) []ConfigQuotaSignUpQuotaConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigQuotaSignUpQuotaConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigQuotaSignUpQuotaConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigQuotaSignUpQuotaConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigQuotaSignUpQuotaConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigQuotaSignUpQuotaConfig(c *Client, des, nw *ConfigQuotaSignUpQuotaConfig) *ConfigQuotaSignUpQuotaConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigQuotaSignUpQuotaConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.QuotaDuration, nw.QuotaDuration) {
		nw.QuotaDuration = des.QuotaDuration
	}

	return nw
}

func canonicalizeNewConfigQuotaSignUpQuotaConfigSet(c *Client, des, nw []ConfigQuotaSignUpQuotaConfig) []ConfigQuotaSignUpQuotaConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigQuotaSignUpQuotaConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigQuotaSignUpQuotaConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigQuotaSignUpQuotaConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigQuotaSignUpQuotaConfigSlice(c *Client, des, nw []ConfigQuotaSignUpQuotaConfig) []ConfigQuotaSignUpQuotaConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigQuotaSignUpQuotaConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigQuotaSignUpQuotaConfig(c, &d, &n))
	}

	return items
}

func canonicalizeConfigMonitoring(des, initial *ConfigMonitoring, opts ...dcl.ApplyOption) *ConfigMonitoring {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigMonitoring{}

	cDes.RequestLogging = canonicalizeConfigMonitoringRequestLogging(des.RequestLogging, initial.RequestLogging, opts...)

	return cDes
}

func canonicalizeConfigMonitoringSlice(des, initial []ConfigMonitoring, opts ...dcl.ApplyOption) []ConfigMonitoring {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigMonitoring, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigMonitoring(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigMonitoring, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigMonitoring(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigMonitoring(c *Client, des, nw *ConfigMonitoring) *ConfigMonitoring {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigMonitoring while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.RequestLogging = canonicalizeNewConfigMonitoringRequestLogging(c, des.RequestLogging, nw.RequestLogging)

	return nw
}

func canonicalizeNewConfigMonitoringSet(c *Client, des, nw []ConfigMonitoring) []ConfigMonitoring {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigMonitoring
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigMonitoringNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigMonitoring(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigMonitoringSlice(c *Client, des, nw []ConfigMonitoring) []ConfigMonitoring {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigMonitoring
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigMonitoring(c, &d, &n))
	}

	return items
}

func canonicalizeConfigMonitoringRequestLogging(des, initial *ConfigMonitoringRequestLogging, opts ...dcl.ApplyOption) *ConfigMonitoringRequestLogging {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigMonitoringRequestLogging{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeConfigMonitoringRequestLoggingSlice(des, initial []ConfigMonitoringRequestLogging, opts ...dcl.ApplyOption) []ConfigMonitoringRequestLogging {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigMonitoringRequestLogging, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigMonitoringRequestLogging(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigMonitoringRequestLogging, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigMonitoringRequestLogging(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigMonitoringRequestLogging(c *Client, des, nw *ConfigMonitoringRequestLogging) *ConfigMonitoringRequestLogging {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigMonitoringRequestLogging while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewConfigMonitoringRequestLoggingSet(c *Client, des, nw []ConfigMonitoringRequestLogging) []ConfigMonitoringRequestLogging {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigMonitoringRequestLogging
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigMonitoringRequestLoggingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigMonitoringRequestLogging(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigMonitoringRequestLoggingSlice(c *Client, des, nw []ConfigMonitoringRequestLogging) []ConfigMonitoringRequestLogging {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigMonitoringRequestLogging
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigMonitoringRequestLogging(c, &d, &n))
	}

	return items
}

func canonicalizeConfigMultiTenant(des, initial *ConfigMultiTenant, opts ...dcl.ApplyOption) *ConfigMultiTenant {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigMultiTenant{}

	if dcl.BoolCanonicalize(des.AllowTenants, initial.AllowTenants) || dcl.IsZeroValue(des.AllowTenants) {
		cDes.AllowTenants = initial.AllowTenants
	} else {
		cDes.AllowTenants = des.AllowTenants
	}
	if dcl.IsZeroValue(des.DefaultTenantLocation) || (dcl.IsEmptyValueIndirect(des.DefaultTenantLocation) && dcl.IsEmptyValueIndirect(initial.DefaultTenantLocation)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DefaultTenantLocation = initial.DefaultTenantLocation
	} else {
		cDes.DefaultTenantLocation = des.DefaultTenantLocation
	}

	return cDes
}

func canonicalizeConfigMultiTenantSlice(des, initial []ConfigMultiTenant, opts ...dcl.ApplyOption) []ConfigMultiTenant {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigMultiTenant, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigMultiTenant(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigMultiTenant, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigMultiTenant(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigMultiTenant(c *Client, des, nw *ConfigMultiTenant) *ConfigMultiTenant {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigMultiTenant while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.AllowTenants, nw.AllowTenants) {
		nw.AllowTenants = des.AllowTenants
	}

	return nw
}

func canonicalizeNewConfigMultiTenantSet(c *Client, des, nw []ConfigMultiTenant) []ConfigMultiTenant {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigMultiTenant
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigMultiTenantNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigMultiTenant(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigMultiTenantSlice(c *Client, des, nw []ConfigMultiTenant) []ConfigMultiTenant {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigMultiTenant
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigMultiTenant(c, &d, &n))
	}

	return items
}

func canonicalizeConfigClient(des, initial *ConfigClient, opts ...dcl.ApplyOption) *ConfigClient {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigClient{}

	cDes.Permissions = canonicalizeConfigClientPermissions(des.Permissions, initial.Permissions, opts...)

	return cDes
}

func canonicalizeConfigClientSlice(des, initial []ConfigClient, opts ...dcl.ApplyOption) []ConfigClient {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigClient, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigClient(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigClient, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigClient(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigClient(c *Client, des, nw *ConfigClient) *ConfigClient {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigClient while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.ApiKey, nw.ApiKey) {
		nw.ApiKey = des.ApiKey
	}
	nw.Permissions = canonicalizeNewConfigClientPermissions(c, des.Permissions, nw.Permissions)
	if dcl.StringCanonicalize(des.FirebaseSubdomain, nw.FirebaseSubdomain) {
		nw.FirebaseSubdomain = des.FirebaseSubdomain
	}

	return nw
}

func canonicalizeNewConfigClientSet(c *Client, des, nw []ConfigClient) []ConfigClient {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigClient
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigClientNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigClient(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigClientSlice(c *Client, des, nw []ConfigClient) []ConfigClient {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigClient
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigClient(c, &d, &n))
	}

	return items
}

func canonicalizeConfigClientPermissions(des, initial *ConfigClientPermissions, opts ...dcl.ApplyOption) *ConfigClientPermissions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigClientPermissions{}

	if dcl.BoolCanonicalize(des.DisabledUserSignup, initial.DisabledUserSignup) || dcl.IsZeroValue(des.DisabledUserSignup) {
		cDes.DisabledUserSignup = initial.DisabledUserSignup
	} else {
		cDes.DisabledUserSignup = des.DisabledUserSignup
	}
	if dcl.BoolCanonicalize(des.DisabledUserDeletion, initial.DisabledUserDeletion) || dcl.IsZeroValue(des.DisabledUserDeletion) {
		cDes.DisabledUserDeletion = initial.DisabledUserDeletion
	} else {
		cDes.DisabledUserDeletion = des.DisabledUserDeletion
	}

	return cDes
}

func canonicalizeConfigClientPermissionsSlice(des, initial []ConfigClientPermissions, opts ...dcl.ApplyOption) []ConfigClientPermissions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigClientPermissions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigClientPermissions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigClientPermissions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigClientPermissions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigClientPermissions(c *Client, des, nw *ConfigClientPermissions) *ConfigClientPermissions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigClientPermissions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.DisabledUserSignup, nw.DisabledUserSignup) {
		nw.DisabledUserSignup = des.DisabledUserSignup
	}
	if dcl.BoolCanonicalize(des.DisabledUserDeletion, nw.DisabledUserDeletion) {
		nw.DisabledUserDeletion = des.DisabledUserDeletion
	}

	return nw
}

func canonicalizeNewConfigClientPermissionsSet(c *Client, des, nw []ConfigClientPermissions) []ConfigClientPermissions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigClientPermissions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigClientPermissionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigClientPermissions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigClientPermissionsSlice(c *Client, des, nw []ConfigClientPermissions) []ConfigClientPermissions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigClientPermissions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigClientPermissions(c, &d, &n))
	}

	return items
}

func canonicalizeConfigMfa(des, initial *ConfigMfa, opts ...dcl.ApplyOption) *ConfigMfa {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigMfa{}

	if dcl.IsZeroValue(des.State) || (dcl.IsEmptyValueIndirect(des.State) && dcl.IsEmptyValueIndirect(initial.State)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.State = initial.State
	} else {
		cDes.State = des.State
	}

	return cDes
}

func canonicalizeConfigMfaSlice(des, initial []ConfigMfa, opts ...dcl.ApplyOption) []ConfigMfa {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigMfa, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigMfa(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigMfa, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigMfa(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigMfa(c *Client, des, nw *ConfigMfa) *ConfigMfa {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigMfa while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewConfigMfaSet(c *Client, des, nw []ConfigMfa) []ConfigMfa {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigMfa
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigMfaNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigMfa(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigMfaSlice(c *Client, des, nw []ConfigMfa) []ConfigMfa {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigMfa
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigMfa(c, &d, &n))
	}

	return items
}

func canonicalizeConfigBlockingFunctions(des, initial *ConfigBlockingFunctions, opts ...dcl.ApplyOption) *ConfigBlockingFunctions {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigBlockingFunctions{}

	if dcl.IsZeroValue(des.Triggers) || (dcl.IsEmptyValueIndirect(des.Triggers) && dcl.IsEmptyValueIndirect(initial.Triggers)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Triggers = initial.Triggers
	} else {
		cDes.Triggers = des.Triggers
	}

	return cDes
}

func canonicalizeConfigBlockingFunctionsSlice(des, initial []ConfigBlockingFunctions, opts ...dcl.ApplyOption) []ConfigBlockingFunctions {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigBlockingFunctions, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigBlockingFunctions(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigBlockingFunctions, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigBlockingFunctions(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigBlockingFunctions(c *Client, des, nw *ConfigBlockingFunctions) *ConfigBlockingFunctions {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigBlockingFunctions while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewConfigBlockingFunctionsSet(c *Client, des, nw []ConfigBlockingFunctions) []ConfigBlockingFunctions {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigBlockingFunctions
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigBlockingFunctionsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigBlockingFunctions(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigBlockingFunctionsSlice(c *Client, des, nw []ConfigBlockingFunctions) []ConfigBlockingFunctions {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigBlockingFunctions
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigBlockingFunctions(c, &d, &n))
	}

	return items
}

func canonicalizeConfigBlockingFunctionsTriggers(des, initial *ConfigBlockingFunctionsTriggers, opts ...dcl.ApplyOption) *ConfigBlockingFunctionsTriggers {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConfigBlockingFunctionsTriggers{}

	if dcl.IsZeroValue(des.FunctionUri) || (dcl.IsEmptyValueIndirect(des.FunctionUri) && dcl.IsEmptyValueIndirect(initial.FunctionUri)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.FunctionUri = initial.FunctionUri
	} else {
		cDes.FunctionUri = des.FunctionUri
	}

	return cDes
}

func canonicalizeConfigBlockingFunctionsTriggersSlice(des, initial []ConfigBlockingFunctionsTriggers, opts ...dcl.ApplyOption) []ConfigBlockingFunctionsTriggers {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConfigBlockingFunctionsTriggers, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConfigBlockingFunctionsTriggers(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConfigBlockingFunctionsTriggers, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConfigBlockingFunctionsTriggers(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConfigBlockingFunctionsTriggers(c *Client, des, nw *ConfigBlockingFunctionsTriggers) *ConfigBlockingFunctionsTriggers {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConfigBlockingFunctionsTriggers while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewConfigBlockingFunctionsTriggersSet(c *Client, des, nw []ConfigBlockingFunctionsTriggers) []ConfigBlockingFunctionsTriggers {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConfigBlockingFunctionsTriggers
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConfigBlockingFunctionsTriggersNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConfigBlockingFunctionsTriggers(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConfigBlockingFunctionsTriggersSlice(c *Client, des, nw []ConfigBlockingFunctionsTriggers) []ConfigBlockingFunctionsTriggers {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConfigBlockingFunctionsTriggers
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConfigBlockingFunctionsTriggers(c, &d, &n))
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
func diffConfig(c *Client, desired, actual *Config, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.SignIn, actual.SignIn, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigSignInNewStyle, EmptyObject: EmptyConfigSignIn, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SignIn")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Notification, actual.Notification, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigNotificationNewStyle, EmptyObject: EmptyConfigNotification, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Notification")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Quota, actual.Quota, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigQuotaNewStyle, EmptyObject: EmptyConfigQuota, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Quota")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Monitoring, actual.Monitoring, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigMonitoringNewStyle, EmptyObject: EmptyConfigMonitoring, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Monitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MultiTenant, actual.MultiTenant, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigMultiTenantNewStyle, EmptyObject: EmptyConfigMultiTenant, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("MultiTenant")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuthorizedDomains, actual.AuthorizedDomains, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("AuthorizedDomains")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subtype, actual.Subtype, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subtype")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Client, actual.Client, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigClientNewStyle, EmptyObject: EmptyConfigClient, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Client")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Mfa, actual.Mfa, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigMfaNewStyle, EmptyObject: EmptyConfigMfa, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Mfa")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BlockingFunctions, actual.BlockingFunctions, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareConfigBlockingFunctionsNewStyle, EmptyObject: EmptyConfigBlockingFunctions, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("BlockingFunctions")); len(ds) != 0 || err != nil {
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

	if len(newDiffs) > 0 {
		c.Config.Logger.Infof("Diff function found diffs: %v", newDiffs)
	}
	return newDiffs, nil
}
func compareConfigSignInNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignIn)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignIn)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignIn or *ConfigSignIn", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignIn)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignIn)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignIn", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Email, actual.Email, dcl.DiffInfo{ObjectFunction: compareConfigSignInEmailNewStyle, EmptyObject: EmptyConfigSignInEmail, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Email")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PhoneNumber, actual.PhoneNumber, dcl.DiffInfo{ObjectFunction: compareConfigSignInPhoneNumberNewStyle, EmptyObject: EmptyConfigSignInPhoneNumber, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("PhoneNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Anonymous, actual.Anonymous, dcl.DiffInfo{ObjectFunction: compareConfigSignInAnonymousNewStyle, EmptyObject: EmptyConfigSignInAnonymous, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Anonymous")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AllowDuplicateEmails, actual.AllowDuplicateEmails, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("AllowDuplicateEmails")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HashConfig, actual.HashConfig, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareConfigSignInHashConfigNewStyle, EmptyObject: EmptyConfigSignInHashConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HashConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigSignInEmailNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignInEmail)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignInEmail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInEmail or *ConfigSignInEmail", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignInEmail)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignInEmail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInEmail", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PasswordRequired, actual.PasswordRequired, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("PasswordRequired")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HashConfig, actual.HashConfig, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareConfigSignInEmailHashConfigNewStyle, EmptyObject: EmptyConfigSignInEmailHashConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HashConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigSignInEmailHashConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignInEmailHashConfig)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignInEmailHashConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInEmailHashConfig or *ConfigSignInEmailHashConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignInEmailHashConfig)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignInEmailHashConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInEmailHashConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Algorithm, actual.Algorithm, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Algorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignerKey, actual.SignerKey, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SignerKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SaltSeparator, actual.SaltSeparator, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SaltSeparator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rounds, actual.Rounds, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Rounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemoryCost, actual.MemoryCost, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemoryCost")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigSignInPhoneNumberNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignInPhoneNumber)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignInPhoneNumber)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInPhoneNumber or *ConfigSignInPhoneNumber", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignInPhoneNumber)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignInPhoneNumber)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInPhoneNumber", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TestPhoneNumbers, actual.TestPhoneNumbers, dcl.DiffInfo{CustomDiff: canonicalizeConfigTestPhoneNumbers, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("TestPhoneNumbers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigSignInAnonymousNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignInAnonymous)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignInAnonymous)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInAnonymous or *ConfigSignInAnonymous", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignInAnonymous)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignInAnonymous)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInAnonymous", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigSignInHashConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigSignInHashConfig)
	if !ok {
		desiredNotPointer, ok := d.(ConfigSignInHashConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInHashConfig or *ConfigSignInHashConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigSignInHashConfig)
	if !ok {
		actualNotPointer, ok := a.(ConfigSignInHashConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigSignInHashConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Algorithm, actual.Algorithm, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Algorithm")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SignerKey, actual.SignerKey, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SignerKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SaltSeparator, actual.SaltSeparator, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SaltSeparator")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Rounds, actual.Rounds, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Rounds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MemoryCost, actual.MemoryCost, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MemoryCost")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotification)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotification)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotification or *ConfigNotification", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotification)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotification)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotification", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SendEmail, actual.SendEmail, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailNewStyle, EmptyObject: EmptyConfigNotificationSendEmail, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SendEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SendSms, actual.SendSms, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendSmsNewStyle, EmptyObject: EmptyConfigNotificationSendSms, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SendSms")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultLocale, actual.DefaultLocale, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("DefaultLocale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmail)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmail or *ConfigNotificationSendEmail", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmail)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmail)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmail", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Method, actual.Method, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Method")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Smtp, actual.Smtp, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailSmtpNewStyle, EmptyObject: EmptyConfigNotificationSendEmailSmtp, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Smtp")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResetPasswordTemplate, actual.ResetPasswordTemplate, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailResetPasswordTemplateNewStyle, EmptyObject: EmptyConfigNotificationSendEmailResetPasswordTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ResetPasswordTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.VerifyEmailTemplate, actual.VerifyEmailTemplate, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailVerifyEmailTemplateNewStyle, EmptyObject: EmptyConfigNotificationSendEmailVerifyEmailTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("VerifyEmailTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ChangeEmailTemplate, actual.ChangeEmailTemplate, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailChangeEmailTemplateNewStyle, EmptyObject: EmptyConfigNotificationSendEmailChangeEmailTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ChangeEmailTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CallbackUri, actual.CallbackUri, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("CallbackUri")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DnsInfo, actual.DnsInfo, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailDnsInfoNewStyle, EmptyObject: EmptyConfigNotificationSendEmailDnsInfo, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("DnsInfo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RevertSecondFactorAdditionTemplate, actual.RevertSecondFactorAdditionTemplate, dcl.DiffInfo{ObjectFunction: compareConfigNotificationSendEmailRevertSecondFactorAdditionTemplateNewStyle, EmptyObject: EmptyConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RevertSecondFactorAdditionTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailSmtpNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailSmtp)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailSmtp)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailSmtp or *ConfigNotificationSendEmailSmtp", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailSmtp)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailSmtp)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailSmtp", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SenderEmail, actual.SenderEmail, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Host, actual.Host, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Host")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Port, actual.Port, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Port")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Username, actual.Username, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Username")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Password, actual.Password, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Password")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecurityMode, actual.SecurityMode, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SecurityMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailResetPasswordTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailResetPasswordTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailResetPasswordTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailResetPasswordTemplate or *ConfigNotificationSendEmailResetPasswordTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailResetPasswordTemplate)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailResetPasswordTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailResetPasswordTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SenderLocalPart, actual.SenderLocalPart, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderLocalPart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SenderDisplayName, actual.SenderDisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderDisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BodyFormat, actual.BodyFormat, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("BodyFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplyTo, actual.ReplyTo, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("ReplyTo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Customized, actual.Customized, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Customized")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailVerifyEmailTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailVerifyEmailTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailVerifyEmailTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailVerifyEmailTemplate or *ConfigNotificationSendEmailVerifyEmailTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailVerifyEmailTemplate)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailVerifyEmailTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailVerifyEmailTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SenderLocalPart, actual.SenderLocalPart, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderLocalPart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SenderDisplayName, actual.SenderDisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderDisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BodyFormat, actual.BodyFormat, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("BodyFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplyTo, actual.ReplyTo, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("ReplyTo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Customized, actual.Customized, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Customized")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailChangeEmailTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailChangeEmailTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailChangeEmailTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailChangeEmailTemplate or *ConfigNotificationSendEmailChangeEmailTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailChangeEmailTemplate)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailChangeEmailTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailChangeEmailTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SenderLocalPart, actual.SenderLocalPart, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderLocalPart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SenderDisplayName, actual.SenderDisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderDisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BodyFormat, actual.BodyFormat, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("BodyFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplyTo, actual.ReplyTo, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("ReplyTo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Customized, actual.Customized, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Customized")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailDnsInfoNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailDnsInfo)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailDnsInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailDnsInfo or *ConfigNotificationSendEmailDnsInfo", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailDnsInfo)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailDnsInfo)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailDnsInfo", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.CustomDomain, actual.CustomDomain, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UseCustomDomain, actual.UseCustomDomain, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("UseCustomDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PendingCustomDomain, actual.PendingCustomDomain, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PendingCustomDomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CustomDomainState, actual.CustomDomainState, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CustomDomainState")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DomainVerificationRequestTime, actual.DomainVerificationRequestTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DomainVerificationRequestTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendEmailRevertSecondFactorAdditionTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate or *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SenderLocalPart, actual.SenderLocalPart, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderLocalPart")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subject, actual.Subject, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Subject")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SenderDisplayName, actual.SenderDisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SenderDisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Body, actual.Body, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Body")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BodyFormat, actual.BodyFormat, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("BodyFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReplyTo, actual.ReplyTo, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("ReplyTo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Customized, actual.Customized, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Customized")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendSmsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendSms)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendSms)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendSms or *ConfigNotificationSendSms", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendSms)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendSms)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendSms", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.UseDeviceLocale, actual.UseDeviceLocale, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("UseDeviceLocale")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SmsTemplate, actual.SmsTemplate, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareConfigNotificationSendSmsSmsTemplateNewStyle, EmptyObject: EmptyConfigNotificationSendSmsSmsTemplate, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SmsTemplate")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigNotificationSendSmsSmsTemplateNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigNotificationSendSmsSmsTemplate)
	if !ok {
		desiredNotPointer, ok := d.(ConfigNotificationSendSmsSmsTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendSmsSmsTemplate or *ConfigNotificationSendSmsSmsTemplate", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigNotificationSendSmsSmsTemplate)
	if !ok {
		actualNotPointer, ok := a.(ConfigNotificationSendSmsSmsTemplate)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigNotificationSendSmsSmsTemplate", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Content, actual.Content, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Content")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigQuotaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigQuota)
	if !ok {
		desiredNotPointer, ok := d.(ConfigQuota)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigQuota or *ConfigQuota", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigQuota)
	if !ok {
		actualNotPointer, ok := a.(ConfigQuota)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigQuota", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SignUpQuotaConfig, actual.SignUpQuotaConfig, dcl.DiffInfo{ObjectFunction: compareConfigQuotaSignUpQuotaConfigNewStyle, EmptyObject: EmptyConfigQuotaSignUpQuotaConfig, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("SignUpQuotaConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigQuotaSignUpQuotaConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigQuotaSignUpQuotaConfig)
	if !ok {
		desiredNotPointer, ok := d.(ConfigQuotaSignUpQuotaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigQuotaSignUpQuotaConfig or *ConfigQuotaSignUpQuotaConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigQuotaSignUpQuotaConfig)
	if !ok {
		actualNotPointer, ok := a.(ConfigQuotaSignUpQuotaConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigQuotaSignUpQuotaConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Quota, actual.Quota, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Quota")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StartTime, actual.StartTime, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("StartTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.QuotaDuration, actual.QuotaDuration, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("QuotaDuration")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigMonitoringNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigMonitoring)
	if !ok {
		desiredNotPointer, ok := d.(ConfigMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMonitoring or *ConfigMonitoring", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigMonitoring)
	if !ok {
		actualNotPointer, ok := a.(ConfigMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMonitoring", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.RequestLogging, actual.RequestLogging, dcl.DiffInfo{ObjectFunction: compareConfigMonitoringRequestLoggingNewStyle, EmptyObject: EmptyConfigMonitoringRequestLogging, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("RequestLogging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigMonitoringRequestLoggingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigMonitoringRequestLogging)
	if !ok {
		desiredNotPointer, ok := d.(ConfigMonitoringRequestLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMonitoringRequestLogging or *ConfigMonitoringRequestLogging", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigMonitoringRequestLogging)
	if !ok {
		actualNotPointer, ok := a.(ConfigMonitoringRequestLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMonitoringRequestLogging", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigMultiTenantNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigMultiTenant)
	if !ok {
		desiredNotPointer, ok := d.(ConfigMultiTenant)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMultiTenant or *ConfigMultiTenant", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigMultiTenant)
	if !ok {
		actualNotPointer, ok := a.(ConfigMultiTenant)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMultiTenant", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.AllowTenants, actual.AllowTenants, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("AllowTenants")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultTenantLocation, actual.DefaultTenantLocation, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("DefaultTenantLocation")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigClientNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigClient)
	if !ok {
		desiredNotPointer, ok := d.(ConfigClient)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigClient or *ConfigClient", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigClient)
	if !ok {
		actualNotPointer, ok := a.(ConfigClient)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigClient", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ApiKey, actual.ApiKey, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApiKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Permissions, actual.Permissions, dcl.DiffInfo{ObjectFunction: compareConfigClientPermissionsNewStyle, EmptyObject: EmptyConfigClientPermissions, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Permissions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FirebaseSubdomain, actual.FirebaseSubdomain, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("FirebaseSubdomain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigClientPermissionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigClientPermissions)
	if !ok {
		desiredNotPointer, ok := d.(ConfigClientPermissions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigClientPermissions or *ConfigClientPermissions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigClientPermissions)
	if !ok {
		actualNotPointer, ok := a.(ConfigClientPermissions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigClientPermissions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.DisabledUserSignup, actual.DisabledUserSignup, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("DisabledUserSignup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisabledUserDeletion, actual.DisabledUserDeletion, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("DisabledUserDeletion")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigMfaNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigMfa)
	if !ok {
		desiredNotPointer, ok := d.(ConfigMfa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMfa or *ConfigMfa", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigMfa)
	if !ok {
		actualNotPointer, ok := a.(ConfigMfa)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigMfa", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigBlockingFunctionsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigBlockingFunctions)
	if !ok {
		desiredNotPointer, ok := d.(ConfigBlockingFunctions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigBlockingFunctions or *ConfigBlockingFunctions", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigBlockingFunctions)
	if !ok {
		actualNotPointer, ok := a.(ConfigBlockingFunctions)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigBlockingFunctions", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Triggers, actual.Triggers, dcl.DiffInfo{ObjectFunction: compareConfigBlockingFunctionsTriggersNewStyle, EmptyObject: EmptyConfigBlockingFunctionsTriggers, OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("Triggers")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareConfigBlockingFunctionsTriggersNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConfigBlockingFunctionsTriggers)
	if !ok {
		desiredNotPointer, ok := d.(ConfigBlockingFunctionsTriggers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigBlockingFunctionsTriggers or *ConfigBlockingFunctionsTriggers", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConfigBlockingFunctionsTriggers)
	if !ok {
		actualNotPointer, ok := a.(ConfigBlockingFunctionsTriggers)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConfigBlockingFunctionsTriggers", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.FunctionUri, actual.FunctionUri, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateConfigUpdateProjectConfigOperation")}, fn.AddNest("FunctionUri")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Config) urlNormalized() *Config {
	normalized := dcl.Copy(*r).(Config)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Config) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateProjectConfig" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
		}
		return dcl.URL("projects/{{project}}/config", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Config resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Config) marshal(c *Client) ([]byte, error) {
	m, err := expandConfig(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Config: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalConfig decodes JSON responses into the Config resource schema.
func unmarshalConfig(b []byte, c *Client, res *Config) (*Config, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapConfig(m, c, res)
}

func unmarshalMapConfig(m map[string]interface{}, c *Client, res *Config) (*Config, error) {

	flattened := flattenConfig(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandConfig expands Config into a JSON request object.
func expandConfig(c *Client, f *Config) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := expandConfigSignIn(c, f.SignIn, res); err != nil {
		return nil, fmt.Errorf("error expanding SignIn into signIn: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["signIn"] = v
	}
	if v, err := expandConfigNotification(c, f.Notification, res); err != nil {
		return nil, fmt.Errorf("error expanding Notification into notification: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["notification"] = v
	}
	if v, err := expandConfigQuota(c, f.Quota, res); err != nil {
		return nil, fmt.Errorf("error expanding Quota into quota: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["quota"] = v
	}
	if v, err := expandConfigMonitoring(c, f.Monitoring, res); err != nil {
		return nil, fmt.Errorf("error expanding Monitoring into monitoring: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["monitoring"] = v
	}
	if v, err := expandConfigMultiTenant(c, f.MultiTenant, res); err != nil {
		return nil, fmt.Errorf("error expanding MultiTenant into multiTenant: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["multiTenant"] = v
	}
	if v := f.AuthorizedDomains; v != nil {
		m["authorizedDomains"] = v
	}
	if v, err := expandConfigClient(c, f.Client, res); err != nil {
		return nil, fmt.Errorf("error expanding Client into client: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["client"] = v
	}
	if v, err := expandConfigMfa(c, f.Mfa, res); err != nil {
		return nil, fmt.Errorf("error expanding Mfa into mfa: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mfa"] = v
	}
	if v, err := expandConfigBlockingFunctions(c, f.BlockingFunctions, res); err != nil {
		return nil, fmt.Errorf("error expanding BlockingFunctions into blockingFunctions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["blockingFunctions"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenConfig flattens Config from a JSON request object into the
// Config type.
func flattenConfig(c *Client, i interface{}, res *Config) *Config {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Config{}
	resultRes.SignIn = flattenConfigSignIn(c, m["signIn"], res)
	resultRes.Notification = flattenConfigNotification(c, m["notification"], res)
	resultRes.Quota = flattenConfigQuota(c, m["quota"], res)
	resultRes.Monitoring = flattenConfigMonitoring(c, m["monitoring"], res)
	resultRes.MultiTenant = flattenConfigMultiTenant(c, m["multiTenant"], res)
	resultRes.AuthorizedDomains = dcl.FlattenStringSlice(m["authorizedDomains"])
	resultRes.Subtype = flattenConfigSubtypeEnum(m["subtype"])
	resultRes.Client = flattenConfigClient(c, m["client"], res)
	resultRes.Mfa = flattenConfigMfa(c, m["mfa"], res)
	resultRes.BlockingFunctions = flattenConfigBlockingFunctions(c, m["blockingFunctions"], res)
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandConfigSignInMap expands the contents of ConfigSignIn into a JSON
// request object.
func expandConfigSignInMap(c *Client, f map[string]ConfigSignIn, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignIn(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInSlice expands the contents of ConfigSignIn into a JSON
// request object.
func expandConfigSignInSlice(c *Client, f []ConfigSignIn, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignIn(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInMap flattens the contents of ConfigSignIn from a JSON
// response object.
func flattenConfigSignInMap(c *Client, i interface{}, res *Config) map[string]ConfigSignIn {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignIn{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignIn{}
	}

	items := make(map[string]ConfigSignIn)
	for k, item := range a {
		items[k] = *flattenConfigSignIn(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInSlice flattens the contents of ConfigSignIn from a JSON
// response object.
func flattenConfigSignInSlice(c *Client, i interface{}, res *Config) []ConfigSignIn {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignIn{}
	}

	if len(a) == 0 {
		return []ConfigSignIn{}
	}

	items := make([]ConfigSignIn, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignIn(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignIn expands an instance of ConfigSignIn into a JSON
// request object.
func expandConfigSignIn(c *Client, f *ConfigSignIn, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigSignInEmail(c, f.Email, res); err != nil {
		return nil, fmt.Errorf("error expanding Email into email: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["email"] = v
	}
	if v, err := expandConfigSignInPhoneNumber(c, f.PhoneNumber, res); err != nil {
		return nil, fmt.Errorf("error expanding PhoneNumber into phoneNumber: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["phoneNumber"] = v
	}
	if v, err := expandConfigSignInAnonymous(c, f.Anonymous, res); err != nil {
		return nil, fmt.Errorf("error expanding Anonymous into anonymous: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["anonymous"] = v
	}
	if v := f.AllowDuplicateEmails; !dcl.IsEmptyValueIndirect(v) {
		m["allowDuplicateEmails"] = v
	}

	return m, nil
}

// flattenConfigSignIn flattens an instance of ConfigSignIn from a JSON
// response object.
func flattenConfigSignIn(c *Client, i interface{}, res *Config) *ConfigSignIn {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignIn{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignIn
	}
	r.Email = flattenConfigSignInEmail(c, m["email"], res)
	r.PhoneNumber = flattenConfigSignInPhoneNumber(c, m["phoneNumber"], res)
	r.Anonymous = flattenConfigSignInAnonymous(c, m["anonymous"], res)
	r.AllowDuplicateEmails = dcl.FlattenBool(m["allowDuplicateEmails"])
	r.HashConfig = flattenConfigSignInHashConfig(c, m["hashConfig"], res)

	return r
}

// expandConfigSignInEmailMap expands the contents of ConfigSignInEmail into a JSON
// request object.
func expandConfigSignInEmailMap(c *Client, f map[string]ConfigSignInEmail, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignInEmail(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInEmailSlice expands the contents of ConfigSignInEmail into a JSON
// request object.
func expandConfigSignInEmailSlice(c *Client, f []ConfigSignInEmail, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignInEmail(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInEmailMap flattens the contents of ConfigSignInEmail from a JSON
// response object.
func flattenConfigSignInEmailMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInEmail {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInEmail{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInEmail{}
	}

	items := make(map[string]ConfigSignInEmail)
	for k, item := range a {
		items[k] = *flattenConfigSignInEmail(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInEmailSlice flattens the contents of ConfigSignInEmail from a JSON
// response object.
func flattenConfigSignInEmailSlice(c *Client, i interface{}, res *Config) []ConfigSignInEmail {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInEmail{}
	}

	if len(a) == 0 {
		return []ConfigSignInEmail{}
	}

	items := make([]ConfigSignInEmail, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInEmail(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignInEmail expands an instance of ConfigSignInEmail into a JSON
// request object.
func expandConfigSignInEmail(c *Client, f *ConfigSignInEmail, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.PasswordRequired; !dcl.IsEmptyValueIndirect(v) {
		m["passwordRequired"] = v
	}

	return m, nil
}

// flattenConfigSignInEmail flattens an instance of ConfigSignInEmail from a JSON
// response object.
func flattenConfigSignInEmail(c *Client, i interface{}, res *Config) *ConfigSignInEmail {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignInEmail{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignInEmail
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.PasswordRequired = dcl.FlattenBool(m["passwordRequired"])
	r.HashConfig = flattenConfigSignInEmailHashConfig(c, m["hashConfig"], res)

	return r
}

// expandConfigSignInEmailHashConfigMap expands the contents of ConfigSignInEmailHashConfig into a JSON
// request object.
func expandConfigSignInEmailHashConfigMap(c *Client, f map[string]ConfigSignInEmailHashConfig, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignInEmailHashConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInEmailHashConfigSlice expands the contents of ConfigSignInEmailHashConfig into a JSON
// request object.
func expandConfigSignInEmailHashConfigSlice(c *Client, f []ConfigSignInEmailHashConfig, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignInEmailHashConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInEmailHashConfigMap flattens the contents of ConfigSignInEmailHashConfig from a JSON
// response object.
func flattenConfigSignInEmailHashConfigMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInEmailHashConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInEmailHashConfig{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInEmailHashConfig{}
	}

	items := make(map[string]ConfigSignInEmailHashConfig)
	for k, item := range a {
		items[k] = *flattenConfigSignInEmailHashConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInEmailHashConfigSlice flattens the contents of ConfigSignInEmailHashConfig from a JSON
// response object.
func flattenConfigSignInEmailHashConfigSlice(c *Client, i interface{}, res *Config) []ConfigSignInEmailHashConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInEmailHashConfig{}
	}

	if len(a) == 0 {
		return []ConfigSignInEmailHashConfig{}
	}

	items := make([]ConfigSignInEmailHashConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInEmailHashConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignInEmailHashConfig expands an instance of ConfigSignInEmailHashConfig into a JSON
// request object.
func expandConfigSignInEmailHashConfig(c *Client, f *ConfigSignInEmailHashConfig, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenConfigSignInEmailHashConfig flattens an instance of ConfigSignInEmailHashConfig from a JSON
// response object.
func flattenConfigSignInEmailHashConfig(c *Client, i interface{}, res *Config) *ConfigSignInEmailHashConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignInEmailHashConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignInEmailHashConfig
	}
	r.Algorithm = flattenConfigSignInEmailHashConfigAlgorithmEnum(m["algorithm"])
	r.SignerKey = dcl.FlattenString(m["signerKey"])
	r.SaltSeparator = dcl.FlattenString(m["saltSeparator"])
	r.Rounds = dcl.FlattenInteger(m["rounds"])
	r.MemoryCost = dcl.FlattenInteger(m["memoryCost"])

	return r
}

// expandConfigSignInPhoneNumberMap expands the contents of ConfigSignInPhoneNumber into a JSON
// request object.
func expandConfigSignInPhoneNumberMap(c *Client, f map[string]ConfigSignInPhoneNumber, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignInPhoneNumber(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInPhoneNumberSlice expands the contents of ConfigSignInPhoneNumber into a JSON
// request object.
func expandConfigSignInPhoneNumberSlice(c *Client, f []ConfigSignInPhoneNumber, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignInPhoneNumber(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInPhoneNumberMap flattens the contents of ConfigSignInPhoneNumber from a JSON
// response object.
func flattenConfigSignInPhoneNumberMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInPhoneNumber {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInPhoneNumber{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInPhoneNumber{}
	}

	items := make(map[string]ConfigSignInPhoneNumber)
	for k, item := range a {
		items[k] = *flattenConfigSignInPhoneNumber(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInPhoneNumberSlice flattens the contents of ConfigSignInPhoneNumber from a JSON
// response object.
func flattenConfigSignInPhoneNumberSlice(c *Client, i interface{}, res *Config) []ConfigSignInPhoneNumber {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInPhoneNumber{}
	}

	if len(a) == 0 {
		return []ConfigSignInPhoneNumber{}
	}

	items := make([]ConfigSignInPhoneNumber, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInPhoneNumber(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignInPhoneNumber expands an instance of ConfigSignInPhoneNumber into a JSON
// request object.
func expandConfigSignInPhoneNumber(c *Client, f *ConfigSignInPhoneNumber, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.TestPhoneNumbers; !dcl.IsEmptyValueIndirect(v) {
		m["testPhoneNumbers"] = v
	}

	return m, nil
}

// flattenConfigSignInPhoneNumber flattens an instance of ConfigSignInPhoneNumber from a JSON
// response object.
func flattenConfigSignInPhoneNumber(c *Client, i interface{}, res *Config) *ConfigSignInPhoneNumber {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignInPhoneNumber{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignInPhoneNumber
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.TestPhoneNumbers = dcl.FlattenKeyValuePairs(m["testPhoneNumbers"])

	return r
}

// expandConfigSignInAnonymousMap expands the contents of ConfigSignInAnonymous into a JSON
// request object.
func expandConfigSignInAnonymousMap(c *Client, f map[string]ConfigSignInAnonymous, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignInAnonymous(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInAnonymousSlice expands the contents of ConfigSignInAnonymous into a JSON
// request object.
func expandConfigSignInAnonymousSlice(c *Client, f []ConfigSignInAnonymous, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignInAnonymous(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInAnonymousMap flattens the contents of ConfigSignInAnonymous from a JSON
// response object.
func flattenConfigSignInAnonymousMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInAnonymous {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInAnonymous{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInAnonymous{}
	}

	items := make(map[string]ConfigSignInAnonymous)
	for k, item := range a {
		items[k] = *flattenConfigSignInAnonymous(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInAnonymousSlice flattens the contents of ConfigSignInAnonymous from a JSON
// response object.
func flattenConfigSignInAnonymousSlice(c *Client, i interface{}, res *Config) []ConfigSignInAnonymous {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInAnonymous{}
	}

	if len(a) == 0 {
		return []ConfigSignInAnonymous{}
	}

	items := make([]ConfigSignInAnonymous, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInAnonymous(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignInAnonymous expands an instance of ConfigSignInAnonymous into a JSON
// request object.
func expandConfigSignInAnonymous(c *Client, f *ConfigSignInAnonymous, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenConfigSignInAnonymous flattens an instance of ConfigSignInAnonymous from a JSON
// response object.
func flattenConfigSignInAnonymous(c *Client, i interface{}, res *Config) *ConfigSignInAnonymous {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignInAnonymous{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignInAnonymous
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandConfigSignInHashConfigMap expands the contents of ConfigSignInHashConfig into a JSON
// request object.
func expandConfigSignInHashConfigMap(c *Client, f map[string]ConfigSignInHashConfig, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigSignInHashConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigSignInHashConfigSlice expands the contents of ConfigSignInHashConfig into a JSON
// request object.
func expandConfigSignInHashConfigSlice(c *Client, f []ConfigSignInHashConfig, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigSignInHashConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigSignInHashConfigMap flattens the contents of ConfigSignInHashConfig from a JSON
// response object.
func flattenConfigSignInHashConfigMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInHashConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInHashConfig{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInHashConfig{}
	}

	items := make(map[string]ConfigSignInHashConfig)
	for k, item := range a {
		items[k] = *flattenConfigSignInHashConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigSignInHashConfigSlice flattens the contents of ConfigSignInHashConfig from a JSON
// response object.
func flattenConfigSignInHashConfigSlice(c *Client, i interface{}, res *Config) []ConfigSignInHashConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInHashConfig{}
	}

	if len(a) == 0 {
		return []ConfigSignInHashConfig{}
	}

	items := make([]ConfigSignInHashConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInHashConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigSignInHashConfig expands an instance of ConfigSignInHashConfig into a JSON
// request object.
func expandConfigSignInHashConfig(c *Client, f *ConfigSignInHashConfig, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenConfigSignInHashConfig flattens an instance of ConfigSignInHashConfig from a JSON
// response object.
func flattenConfigSignInHashConfig(c *Client, i interface{}, res *Config) *ConfigSignInHashConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigSignInHashConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigSignInHashConfig
	}
	r.Algorithm = flattenConfigSignInHashConfigAlgorithmEnum(m["algorithm"])
	r.SignerKey = dcl.FlattenString(m["signerKey"])
	r.SaltSeparator = dcl.FlattenString(m["saltSeparator"])
	r.Rounds = dcl.FlattenInteger(m["rounds"])
	r.MemoryCost = dcl.FlattenInteger(m["memoryCost"])

	return r
}

// expandConfigNotificationMap expands the contents of ConfigNotification into a JSON
// request object.
func expandConfigNotificationMap(c *Client, f map[string]ConfigNotification, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotification(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSlice expands the contents of ConfigNotification into a JSON
// request object.
func expandConfigNotificationSlice(c *Client, f []ConfigNotification, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotification(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationMap flattens the contents of ConfigNotification from a JSON
// response object.
func flattenConfigNotificationMap(c *Client, i interface{}, res *Config) map[string]ConfigNotification {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotification{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotification{}
	}

	items := make(map[string]ConfigNotification)
	for k, item := range a {
		items[k] = *flattenConfigNotification(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSlice flattens the contents of ConfigNotification from a JSON
// response object.
func flattenConfigNotificationSlice(c *Client, i interface{}, res *Config) []ConfigNotification {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotification{}
	}

	if len(a) == 0 {
		return []ConfigNotification{}
	}

	items := make([]ConfigNotification, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotification(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotification expands an instance of ConfigNotification into a JSON
// request object.
func expandConfigNotification(c *Client, f *ConfigNotification, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigNotificationSendEmail(c, f.SendEmail, res); err != nil {
		return nil, fmt.Errorf("error expanding SendEmail into sendEmail: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sendEmail"] = v
	}
	if v, err := expandConfigNotificationSendSms(c, f.SendSms, res); err != nil {
		return nil, fmt.Errorf("error expanding SendSms into sendSms: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["sendSms"] = v
	}
	if v := f.DefaultLocale; !dcl.IsEmptyValueIndirect(v) {
		m["defaultLocale"] = v
	}

	return m, nil
}

// flattenConfigNotification flattens an instance of ConfigNotification from a JSON
// response object.
func flattenConfigNotification(c *Client, i interface{}, res *Config) *ConfigNotification {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotification{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotification
	}
	r.SendEmail = flattenConfigNotificationSendEmail(c, m["sendEmail"], res)
	r.SendSms = flattenConfigNotificationSendSms(c, m["sendSms"], res)
	r.DefaultLocale = dcl.FlattenString(m["defaultLocale"])

	return r
}

// expandConfigNotificationSendEmailMap expands the contents of ConfigNotificationSendEmail into a JSON
// request object.
func expandConfigNotificationSendEmailMap(c *Client, f map[string]ConfigNotificationSendEmail, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmail(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailSlice expands the contents of ConfigNotificationSendEmail into a JSON
// request object.
func expandConfigNotificationSendEmailSlice(c *Client, f []ConfigNotificationSendEmail, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmail(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailMap flattens the contents of ConfigNotificationSendEmail from a JSON
// response object.
func flattenConfigNotificationSendEmailMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmail {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmail{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmail{}
	}

	items := make(map[string]ConfigNotificationSendEmail)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmail(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailSlice flattens the contents of ConfigNotificationSendEmail from a JSON
// response object.
func flattenConfigNotificationSendEmailSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmail {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmail{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmail{}
	}

	items := make([]ConfigNotificationSendEmail, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmail(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmail expands an instance of ConfigNotificationSendEmail into a JSON
// request object.
func expandConfigNotificationSendEmail(c *Client, f *ConfigNotificationSendEmail, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Method; !dcl.IsEmptyValueIndirect(v) {
		m["method"] = v
	}
	if v, err := expandConfigNotificationSendEmailSmtp(c, f.Smtp, res); err != nil {
		return nil, fmt.Errorf("error expanding Smtp into smtp: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["smtp"] = v
	}
	if v, err := expandConfigNotificationSendEmailResetPasswordTemplate(c, f.ResetPasswordTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding ResetPasswordTemplate into resetPasswordTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["resetPasswordTemplate"] = v
	}
	if v, err := expandConfigNotificationSendEmailVerifyEmailTemplate(c, f.VerifyEmailTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding VerifyEmailTemplate into verifyEmailTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["verifyEmailTemplate"] = v
	}
	if v, err := expandConfigNotificationSendEmailChangeEmailTemplate(c, f.ChangeEmailTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding ChangeEmailTemplate into changeEmailTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["changeEmailTemplate"] = v
	}
	if v := f.CallbackUri; !dcl.IsEmptyValueIndirect(v) {
		m["callbackUri"] = v
	}
	if v, err := expandConfigNotificationSendEmailDnsInfo(c, f.DnsInfo, res); err != nil {
		return nil, fmt.Errorf("error expanding DnsInfo into dnsInfo: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dnsInfo"] = v
	}
	if v, err := expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, f.RevertSecondFactorAdditionTemplate, res); err != nil {
		return nil, fmt.Errorf("error expanding RevertSecondFactorAdditionTemplate into revertSecondFactorAdditionTemplate: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["revertSecondFactorAdditionTemplate"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmail flattens an instance of ConfigNotificationSendEmail from a JSON
// response object.
func flattenConfigNotificationSendEmail(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmail {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmail{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmail
	}
	r.Method = flattenConfigNotificationSendEmailMethodEnum(m["method"])
	r.Smtp = flattenConfigNotificationSendEmailSmtp(c, m["smtp"], res)
	r.ResetPasswordTemplate = flattenConfigNotificationSendEmailResetPasswordTemplate(c, m["resetPasswordTemplate"], res)
	r.VerifyEmailTemplate = flattenConfigNotificationSendEmailVerifyEmailTemplate(c, m["verifyEmailTemplate"], res)
	r.ChangeEmailTemplate = flattenConfigNotificationSendEmailChangeEmailTemplate(c, m["changeEmailTemplate"], res)
	r.CallbackUri = dcl.FlattenString(m["callbackUri"])
	r.DnsInfo = flattenConfigNotificationSendEmailDnsInfo(c, m["dnsInfo"], res)
	r.RevertSecondFactorAdditionTemplate = flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, m["revertSecondFactorAdditionTemplate"], res)

	return r
}

// expandConfigNotificationSendEmailSmtpMap expands the contents of ConfigNotificationSendEmailSmtp into a JSON
// request object.
func expandConfigNotificationSendEmailSmtpMap(c *Client, f map[string]ConfigNotificationSendEmailSmtp, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailSmtp(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailSmtpSlice expands the contents of ConfigNotificationSendEmailSmtp into a JSON
// request object.
func expandConfigNotificationSendEmailSmtpSlice(c *Client, f []ConfigNotificationSendEmailSmtp, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailSmtp(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailSmtpMap flattens the contents of ConfigNotificationSendEmailSmtp from a JSON
// response object.
func flattenConfigNotificationSendEmailSmtpMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailSmtp {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailSmtp{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailSmtp{}
	}

	items := make(map[string]ConfigNotificationSendEmailSmtp)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailSmtp(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailSmtpSlice flattens the contents of ConfigNotificationSendEmailSmtp from a JSON
// response object.
func flattenConfigNotificationSendEmailSmtpSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailSmtp {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailSmtp{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailSmtp{}
	}

	items := make([]ConfigNotificationSendEmailSmtp, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailSmtp(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailSmtp expands an instance of ConfigNotificationSendEmailSmtp into a JSON
// request object.
func expandConfigNotificationSendEmailSmtp(c *Client, f *ConfigNotificationSendEmailSmtp, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SenderEmail; !dcl.IsEmptyValueIndirect(v) {
		m["senderEmail"] = v
	}
	if v := f.Host; !dcl.IsEmptyValueIndirect(v) {
		m["host"] = v
	}
	if v := f.Port; !dcl.IsEmptyValueIndirect(v) {
		m["port"] = v
	}
	if v := f.Username; !dcl.IsEmptyValueIndirect(v) {
		m["username"] = v
	}
	if v := f.Password; !dcl.IsEmptyValueIndirect(v) {
		m["password"] = v
	}
	if v := f.SecurityMode; !dcl.IsEmptyValueIndirect(v) {
		m["securityMode"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailSmtp flattens an instance of ConfigNotificationSendEmailSmtp from a JSON
// response object.
func flattenConfigNotificationSendEmailSmtp(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailSmtp {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailSmtp{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailSmtp
	}
	r.SenderEmail = dcl.FlattenString(m["senderEmail"])
	r.Host = dcl.FlattenString(m["host"])
	r.Port = dcl.FlattenInteger(m["port"])
	r.Username = dcl.FlattenString(m["username"])
	r.Password = dcl.FlattenString(m["password"])
	r.SecurityMode = flattenConfigNotificationSendEmailSmtpSecurityModeEnum(m["securityMode"])

	return r
}

// expandConfigNotificationSendEmailResetPasswordTemplateMap expands the contents of ConfigNotificationSendEmailResetPasswordTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailResetPasswordTemplateMap(c *Client, f map[string]ConfigNotificationSendEmailResetPasswordTemplate, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailResetPasswordTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailResetPasswordTemplateSlice expands the contents of ConfigNotificationSendEmailResetPasswordTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailResetPasswordTemplateSlice(c *Client, f []ConfigNotificationSendEmailResetPasswordTemplate, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailResetPasswordTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailResetPasswordTemplateMap flattens the contents of ConfigNotificationSendEmailResetPasswordTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailResetPasswordTemplateMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailResetPasswordTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailResetPasswordTemplate{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailResetPasswordTemplate{}
	}

	items := make(map[string]ConfigNotificationSendEmailResetPasswordTemplate)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailResetPasswordTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailResetPasswordTemplateSlice flattens the contents of ConfigNotificationSendEmailResetPasswordTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailResetPasswordTemplateSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailResetPasswordTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailResetPasswordTemplate{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailResetPasswordTemplate{}
	}

	items := make([]ConfigNotificationSendEmailResetPasswordTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailResetPasswordTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailResetPasswordTemplate expands an instance of ConfigNotificationSendEmailResetPasswordTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailResetPasswordTemplate(c *Client, f *ConfigNotificationSendEmailResetPasswordTemplate, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SenderLocalPart; !dcl.IsEmptyValueIndirect(v) {
		m["senderLocalPart"] = v
	}
	if v := f.Subject; !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v := f.SenderDisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["senderDisplayName"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v := f.BodyFormat; !dcl.IsEmptyValueIndirect(v) {
		m["bodyFormat"] = v
	}
	if v := f.ReplyTo; !dcl.IsEmptyValueIndirect(v) {
		m["replyTo"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailResetPasswordTemplate flattens an instance of ConfigNotificationSendEmailResetPasswordTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailResetPasswordTemplate(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailResetPasswordTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailResetPasswordTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailResetPasswordTemplate
	}
	r.SenderLocalPart = dcl.FlattenString(m["senderLocalPart"])
	r.Subject = dcl.FlattenString(m["subject"])
	r.SenderDisplayName = dcl.FlattenString(m["senderDisplayName"])
	r.Body = dcl.FlattenString(m["body"])
	r.BodyFormat = flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(m["bodyFormat"])
	r.ReplyTo = dcl.FlattenString(m["replyTo"])
	r.Customized = dcl.FlattenBool(m["customized"])

	return r
}

// expandConfigNotificationSendEmailVerifyEmailTemplateMap expands the contents of ConfigNotificationSendEmailVerifyEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailVerifyEmailTemplateMap(c *Client, f map[string]ConfigNotificationSendEmailVerifyEmailTemplate, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailVerifyEmailTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailVerifyEmailTemplateSlice expands the contents of ConfigNotificationSendEmailVerifyEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailVerifyEmailTemplateSlice(c *Client, f []ConfigNotificationSendEmailVerifyEmailTemplate, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailVerifyEmailTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailVerifyEmailTemplateMap flattens the contents of ConfigNotificationSendEmailVerifyEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailVerifyEmailTemplateMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailVerifyEmailTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailVerifyEmailTemplate{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailVerifyEmailTemplate{}
	}

	items := make(map[string]ConfigNotificationSendEmailVerifyEmailTemplate)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailVerifyEmailTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailVerifyEmailTemplateSlice flattens the contents of ConfigNotificationSendEmailVerifyEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailVerifyEmailTemplateSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailVerifyEmailTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailVerifyEmailTemplate{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailVerifyEmailTemplate{}
	}

	items := make([]ConfigNotificationSendEmailVerifyEmailTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailVerifyEmailTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailVerifyEmailTemplate expands an instance of ConfigNotificationSendEmailVerifyEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailVerifyEmailTemplate(c *Client, f *ConfigNotificationSendEmailVerifyEmailTemplate, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SenderLocalPart; !dcl.IsEmptyValueIndirect(v) {
		m["senderLocalPart"] = v
	}
	if v := f.Subject; !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v := f.SenderDisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["senderDisplayName"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v := f.BodyFormat; !dcl.IsEmptyValueIndirect(v) {
		m["bodyFormat"] = v
	}
	if v := f.ReplyTo; !dcl.IsEmptyValueIndirect(v) {
		m["replyTo"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailVerifyEmailTemplate flattens an instance of ConfigNotificationSendEmailVerifyEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailVerifyEmailTemplate(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailVerifyEmailTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailVerifyEmailTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailVerifyEmailTemplate
	}
	r.SenderLocalPart = dcl.FlattenString(m["senderLocalPart"])
	r.Subject = dcl.FlattenString(m["subject"])
	r.SenderDisplayName = dcl.FlattenString(m["senderDisplayName"])
	r.Body = dcl.FlattenString(m["body"])
	r.BodyFormat = flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(m["bodyFormat"])
	r.ReplyTo = dcl.FlattenString(m["replyTo"])
	r.Customized = dcl.FlattenBool(m["customized"])

	return r
}

// expandConfigNotificationSendEmailChangeEmailTemplateMap expands the contents of ConfigNotificationSendEmailChangeEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailChangeEmailTemplateMap(c *Client, f map[string]ConfigNotificationSendEmailChangeEmailTemplate, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailChangeEmailTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailChangeEmailTemplateSlice expands the contents of ConfigNotificationSendEmailChangeEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailChangeEmailTemplateSlice(c *Client, f []ConfigNotificationSendEmailChangeEmailTemplate, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailChangeEmailTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailChangeEmailTemplateMap flattens the contents of ConfigNotificationSendEmailChangeEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailChangeEmailTemplateMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailChangeEmailTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailChangeEmailTemplate{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailChangeEmailTemplate{}
	}

	items := make(map[string]ConfigNotificationSendEmailChangeEmailTemplate)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailChangeEmailTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailChangeEmailTemplateSlice flattens the contents of ConfigNotificationSendEmailChangeEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailChangeEmailTemplateSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailChangeEmailTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailChangeEmailTemplate{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailChangeEmailTemplate{}
	}

	items := make([]ConfigNotificationSendEmailChangeEmailTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailChangeEmailTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailChangeEmailTemplate expands an instance of ConfigNotificationSendEmailChangeEmailTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailChangeEmailTemplate(c *Client, f *ConfigNotificationSendEmailChangeEmailTemplate, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SenderLocalPart; !dcl.IsEmptyValueIndirect(v) {
		m["senderLocalPart"] = v
	}
	if v := f.Subject; !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v := f.SenderDisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["senderDisplayName"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v := f.BodyFormat; !dcl.IsEmptyValueIndirect(v) {
		m["bodyFormat"] = v
	}
	if v := f.ReplyTo; !dcl.IsEmptyValueIndirect(v) {
		m["replyTo"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailChangeEmailTemplate flattens an instance of ConfigNotificationSendEmailChangeEmailTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailChangeEmailTemplate(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailChangeEmailTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailChangeEmailTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailChangeEmailTemplate
	}
	r.SenderLocalPart = dcl.FlattenString(m["senderLocalPart"])
	r.Subject = dcl.FlattenString(m["subject"])
	r.SenderDisplayName = dcl.FlattenString(m["senderDisplayName"])
	r.Body = dcl.FlattenString(m["body"])
	r.BodyFormat = flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(m["bodyFormat"])
	r.ReplyTo = dcl.FlattenString(m["replyTo"])
	r.Customized = dcl.FlattenBool(m["customized"])

	return r
}

// expandConfigNotificationSendEmailDnsInfoMap expands the contents of ConfigNotificationSendEmailDnsInfo into a JSON
// request object.
func expandConfigNotificationSendEmailDnsInfoMap(c *Client, f map[string]ConfigNotificationSendEmailDnsInfo, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailDnsInfo(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailDnsInfoSlice expands the contents of ConfigNotificationSendEmailDnsInfo into a JSON
// request object.
func expandConfigNotificationSendEmailDnsInfoSlice(c *Client, f []ConfigNotificationSendEmailDnsInfo, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailDnsInfo(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailDnsInfoMap flattens the contents of ConfigNotificationSendEmailDnsInfo from a JSON
// response object.
func flattenConfigNotificationSendEmailDnsInfoMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailDnsInfo {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailDnsInfo{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailDnsInfo{}
	}

	items := make(map[string]ConfigNotificationSendEmailDnsInfo)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailDnsInfo(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailDnsInfoSlice flattens the contents of ConfigNotificationSendEmailDnsInfo from a JSON
// response object.
func flattenConfigNotificationSendEmailDnsInfoSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailDnsInfo {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailDnsInfo{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailDnsInfo{}
	}

	items := make([]ConfigNotificationSendEmailDnsInfo, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailDnsInfo(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailDnsInfo expands an instance of ConfigNotificationSendEmailDnsInfo into a JSON
// request object.
func expandConfigNotificationSendEmailDnsInfo(c *Client, f *ConfigNotificationSendEmailDnsInfo, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseCustomDomain; !dcl.IsEmptyValueIndirect(v) {
		m["useCustomDomain"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailDnsInfo flattens an instance of ConfigNotificationSendEmailDnsInfo from a JSON
// response object.
func flattenConfigNotificationSendEmailDnsInfo(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailDnsInfo {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailDnsInfo{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailDnsInfo
	}
	r.CustomDomain = dcl.FlattenString(m["customDomain"])
	r.UseCustomDomain = dcl.FlattenBool(m["useCustomDomain"])
	r.PendingCustomDomain = dcl.FlattenString(m["pendingCustomDomain"])
	r.CustomDomainState = flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(m["customDomainState"])
	r.DomainVerificationRequestTime = dcl.FlattenString(m["domainVerificationRequestTime"])

	return r
}

// expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplateMap expands the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplateMap(c *Client, f map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice expands the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice(c *Client, f []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateMap flattens the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}

	items := make(map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice flattens the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}

	items := make([]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplate expands an instance of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate into a JSON
// request object.
func expandConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c *Client, f *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SenderLocalPart; !dcl.IsEmptyValueIndirect(v) {
		m["senderLocalPart"] = v
	}
	if v := f.Subject; !dcl.IsEmptyValueIndirect(v) {
		m["subject"] = v
	}
	if v := f.SenderDisplayName; !dcl.IsEmptyValueIndirect(v) {
		m["senderDisplayName"] = v
	}
	if v := f.Body; !dcl.IsEmptyValueIndirect(v) {
		m["body"] = v
	}
	if v := f.BodyFormat; !dcl.IsEmptyValueIndirect(v) {
		m["bodyFormat"] = v
	}
	if v := f.ReplyTo; !dcl.IsEmptyValueIndirect(v) {
		m["replyTo"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplate flattens an instance of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate from a JSON
// response object.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplate(c *Client, i interface{}, res *Config) *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendEmailRevertSecondFactorAdditionTemplate
	}
	r.SenderLocalPart = dcl.FlattenString(m["senderLocalPart"])
	r.Subject = dcl.FlattenString(m["subject"])
	r.SenderDisplayName = dcl.FlattenString(m["senderDisplayName"])
	r.Body = dcl.FlattenString(m["body"])
	r.BodyFormat = flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(m["bodyFormat"])
	r.ReplyTo = dcl.FlattenString(m["replyTo"])
	r.Customized = dcl.FlattenBool(m["customized"])

	return r
}

// expandConfigNotificationSendSmsMap expands the contents of ConfigNotificationSendSms into a JSON
// request object.
func expandConfigNotificationSendSmsMap(c *Client, f map[string]ConfigNotificationSendSms, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendSms(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendSmsSlice expands the contents of ConfigNotificationSendSms into a JSON
// request object.
func expandConfigNotificationSendSmsSlice(c *Client, f []ConfigNotificationSendSms, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendSms(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendSmsMap flattens the contents of ConfigNotificationSendSms from a JSON
// response object.
func flattenConfigNotificationSendSmsMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendSms {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendSms{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendSms{}
	}

	items := make(map[string]ConfigNotificationSendSms)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendSms(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendSmsSlice flattens the contents of ConfigNotificationSendSms from a JSON
// response object.
func flattenConfigNotificationSendSmsSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendSms {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendSms{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendSms{}
	}

	items := make([]ConfigNotificationSendSms, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendSms(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendSms expands an instance of ConfigNotificationSendSms into a JSON
// request object.
func expandConfigNotificationSendSms(c *Client, f *ConfigNotificationSendSms, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.UseDeviceLocale; !dcl.IsEmptyValueIndirect(v) {
		m["useDeviceLocale"] = v
	}

	return m, nil
}

// flattenConfigNotificationSendSms flattens an instance of ConfigNotificationSendSms from a JSON
// response object.
func flattenConfigNotificationSendSms(c *Client, i interface{}, res *Config) *ConfigNotificationSendSms {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendSms{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendSms
	}
	r.UseDeviceLocale = dcl.FlattenBool(m["useDeviceLocale"])
	r.SmsTemplate = flattenConfigNotificationSendSmsSmsTemplate(c, m["smsTemplate"], res)

	return r
}

// expandConfigNotificationSendSmsSmsTemplateMap expands the contents of ConfigNotificationSendSmsSmsTemplate into a JSON
// request object.
func expandConfigNotificationSendSmsSmsTemplateMap(c *Client, f map[string]ConfigNotificationSendSmsSmsTemplate, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigNotificationSendSmsSmsTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigNotificationSendSmsSmsTemplateSlice expands the contents of ConfigNotificationSendSmsSmsTemplate into a JSON
// request object.
func expandConfigNotificationSendSmsSmsTemplateSlice(c *Client, f []ConfigNotificationSendSmsSmsTemplate, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigNotificationSendSmsSmsTemplate(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigNotificationSendSmsSmsTemplateMap flattens the contents of ConfigNotificationSendSmsSmsTemplate from a JSON
// response object.
func flattenConfigNotificationSendSmsSmsTemplateMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendSmsSmsTemplate {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendSmsSmsTemplate{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendSmsSmsTemplate{}
	}

	items := make(map[string]ConfigNotificationSendSmsSmsTemplate)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendSmsSmsTemplate(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigNotificationSendSmsSmsTemplateSlice flattens the contents of ConfigNotificationSendSmsSmsTemplate from a JSON
// response object.
func flattenConfigNotificationSendSmsSmsTemplateSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendSmsSmsTemplate {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendSmsSmsTemplate{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendSmsSmsTemplate{}
	}

	items := make([]ConfigNotificationSendSmsSmsTemplate, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendSmsSmsTemplate(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigNotificationSendSmsSmsTemplate expands an instance of ConfigNotificationSendSmsSmsTemplate into a JSON
// request object.
func expandConfigNotificationSendSmsSmsTemplate(c *Client, f *ConfigNotificationSendSmsSmsTemplate, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenConfigNotificationSendSmsSmsTemplate flattens an instance of ConfigNotificationSendSmsSmsTemplate from a JSON
// response object.
func flattenConfigNotificationSendSmsSmsTemplate(c *Client, i interface{}, res *Config) *ConfigNotificationSendSmsSmsTemplate {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigNotificationSendSmsSmsTemplate{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigNotificationSendSmsSmsTemplate
	}
	r.Content = dcl.FlattenString(m["content"])

	return r
}

// expandConfigQuotaMap expands the contents of ConfigQuota into a JSON
// request object.
func expandConfigQuotaMap(c *Client, f map[string]ConfigQuota, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigQuota(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigQuotaSlice expands the contents of ConfigQuota into a JSON
// request object.
func expandConfigQuotaSlice(c *Client, f []ConfigQuota, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigQuota(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigQuotaMap flattens the contents of ConfigQuota from a JSON
// response object.
func flattenConfigQuotaMap(c *Client, i interface{}, res *Config) map[string]ConfigQuota {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigQuota{}
	}

	if len(a) == 0 {
		return map[string]ConfigQuota{}
	}

	items := make(map[string]ConfigQuota)
	for k, item := range a {
		items[k] = *flattenConfigQuota(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigQuotaSlice flattens the contents of ConfigQuota from a JSON
// response object.
func flattenConfigQuotaSlice(c *Client, i interface{}, res *Config) []ConfigQuota {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigQuota{}
	}

	if len(a) == 0 {
		return []ConfigQuota{}
	}

	items := make([]ConfigQuota, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigQuota(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigQuota expands an instance of ConfigQuota into a JSON
// request object.
func expandConfigQuota(c *Client, f *ConfigQuota, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigQuotaSignUpQuotaConfig(c, f.SignUpQuotaConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding SignUpQuotaConfig into signUpQuotaConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["signUpQuotaConfig"] = v
	}

	return m, nil
}

// flattenConfigQuota flattens an instance of ConfigQuota from a JSON
// response object.
func flattenConfigQuota(c *Client, i interface{}, res *Config) *ConfigQuota {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigQuota{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigQuota
	}
	r.SignUpQuotaConfig = flattenConfigQuotaSignUpQuotaConfig(c, m["signUpQuotaConfig"], res)

	return r
}

// expandConfigQuotaSignUpQuotaConfigMap expands the contents of ConfigQuotaSignUpQuotaConfig into a JSON
// request object.
func expandConfigQuotaSignUpQuotaConfigMap(c *Client, f map[string]ConfigQuotaSignUpQuotaConfig, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigQuotaSignUpQuotaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigQuotaSignUpQuotaConfigSlice expands the contents of ConfigQuotaSignUpQuotaConfig into a JSON
// request object.
func expandConfigQuotaSignUpQuotaConfigSlice(c *Client, f []ConfigQuotaSignUpQuotaConfig, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigQuotaSignUpQuotaConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigQuotaSignUpQuotaConfigMap flattens the contents of ConfigQuotaSignUpQuotaConfig from a JSON
// response object.
func flattenConfigQuotaSignUpQuotaConfigMap(c *Client, i interface{}, res *Config) map[string]ConfigQuotaSignUpQuotaConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigQuotaSignUpQuotaConfig{}
	}

	if len(a) == 0 {
		return map[string]ConfigQuotaSignUpQuotaConfig{}
	}

	items := make(map[string]ConfigQuotaSignUpQuotaConfig)
	for k, item := range a {
		items[k] = *flattenConfigQuotaSignUpQuotaConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigQuotaSignUpQuotaConfigSlice flattens the contents of ConfigQuotaSignUpQuotaConfig from a JSON
// response object.
func flattenConfigQuotaSignUpQuotaConfigSlice(c *Client, i interface{}, res *Config) []ConfigQuotaSignUpQuotaConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigQuotaSignUpQuotaConfig{}
	}

	if len(a) == 0 {
		return []ConfigQuotaSignUpQuotaConfig{}
	}

	items := make([]ConfigQuotaSignUpQuotaConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigQuotaSignUpQuotaConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigQuotaSignUpQuotaConfig expands an instance of ConfigQuotaSignUpQuotaConfig into a JSON
// request object.
func expandConfigQuotaSignUpQuotaConfig(c *Client, f *ConfigQuotaSignUpQuotaConfig, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Quota; !dcl.IsEmptyValueIndirect(v) {
		m["quota"] = v
	}
	if v := f.StartTime; !dcl.IsEmptyValueIndirect(v) {
		m["startTime"] = v
	}
	if v := f.QuotaDuration; !dcl.IsEmptyValueIndirect(v) {
		m["quotaDuration"] = v
	}

	return m, nil
}

// flattenConfigQuotaSignUpQuotaConfig flattens an instance of ConfigQuotaSignUpQuotaConfig from a JSON
// response object.
func flattenConfigQuotaSignUpQuotaConfig(c *Client, i interface{}, res *Config) *ConfigQuotaSignUpQuotaConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigQuotaSignUpQuotaConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigQuotaSignUpQuotaConfig
	}
	r.Quota = dcl.FlattenInteger(m["quota"])
	r.StartTime = dcl.FlattenString(m["startTime"])
	r.QuotaDuration = dcl.FlattenString(m["quotaDuration"])

	return r
}

// expandConfigMonitoringMap expands the contents of ConfigMonitoring into a JSON
// request object.
func expandConfigMonitoringMap(c *Client, f map[string]ConfigMonitoring, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigMonitoringSlice expands the contents of ConfigMonitoring into a JSON
// request object.
func expandConfigMonitoringSlice(c *Client, f []ConfigMonitoring, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigMonitoringMap flattens the contents of ConfigMonitoring from a JSON
// response object.
func flattenConfigMonitoringMap(c *Client, i interface{}, res *Config) map[string]ConfigMonitoring {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigMonitoring{}
	}

	if len(a) == 0 {
		return map[string]ConfigMonitoring{}
	}

	items := make(map[string]ConfigMonitoring)
	for k, item := range a {
		items[k] = *flattenConfigMonitoring(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigMonitoringSlice flattens the contents of ConfigMonitoring from a JSON
// response object.
func flattenConfigMonitoringSlice(c *Client, i interface{}, res *Config) []ConfigMonitoring {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigMonitoring{}
	}

	if len(a) == 0 {
		return []ConfigMonitoring{}
	}

	items := make([]ConfigMonitoring, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigMonitoring(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigMonitoring expands an instance of ConfigMonitoring into a JSON
// request object.
func expandConfigMonitoring(c *Client, f *ConfigMonitoring, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigMonitoringRequestLogging(c, f.RequestLogging, res); err != nil {
		return nil, fmt.Errorf("error expanding RequestLogging into requestLogging: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["requestLogging"] = v
	}

	return m, nil
}

// flattenConfigMonitoring flattens an instance of ConfigMonitoring from a JSON
// response object.
func flattenConfigMonitoring(c *Client, i interface{}, res *Config) *ConfigMonitoring {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigMonitoring{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigMonitoring
	}
	r.RequestLogging = flattenConfigMonitoringRequestLogging(c, m["requestLogging"], res)

	return r
}

// expandConfigMonitoringRequestLoggingMap expands the contents of ConfigMonitoringRequestLogging into a JSON
// request object.
func expandConfigMonitoringRequestLoggingMap(c *Client, f map[string]ConfigMonitoringRequestLogging, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigMonitoringRequestLogging(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigMonitoringRequestLoggingSlice expands the contents of ConfigMonitoringRequestLogging into a JSON
// request object.
func expandConfigMonitoringRequestLoggingSlice(c *Client, f []ConfigMonitoringRequestLogging, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigMonitoringRequestLogging(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigMonitoringRequestLoggingMap flattens the contents of ConfigMonitoringRequestLogging from a JSON
// response object.
func flattenConfigMonitoringRequestLoggingMap(c *Client, i interface{}, res *Config) map[string]ConfigMonitoringRequestLogging {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigMonitoringRequestLogging{}
	}

	if len(a) == 0 {
		return map[string]ConfigMonitoringRequestLogging{}
	}

	items := make(map[string]ConfigMonitoringRequestLogging)
	for k, item := range a {
		items[k] = *flattenConfigMonitoringRequestLogging(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigMonitoringRequestLoggingSlice flattens the contents of ConfigMonitoringRequestLogging from a JSON
// response object.
func flattenConfigMonitoringRequestLoggingSlice(c *Client, i interface{}, res *Config) []ConfigMonitoringRequestLogging {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigMonitoringRequestLogging{}
	}

	if len(a) == 0 {
		return []ConfigMonitoringRequestLogging{}
	}

	items := make([]ConfigMonitoringRequestLogging, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigMonitoringRequestLogging(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigMonitoringRequestLogging expands an instance of ConfigMonitoringRequestLogging into a JSON
// request object.
func expandConfigMonitoringRequestLogging(c *Client, f *ConfigMonitoringRequestLogging, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenConfigMonitoringRequestLogging flattens an instance of ConfigMonitoringRequestLogging from a JSON
// response object.
func flattenConfigMonitoringRequestLogging(c *Client, i interface{}, res *Config) *ConfigMonitoringRequestLogging {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigMonitoringRequestLogging{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigMonitoringRequestLogging
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandConfigMultiTenantMap expands the contents of ConfigMultiTenant into a JSON
// request object.
func expandConfigMultiTenantMap(c *Client, f map[string]ConfigMultiTenant, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigMultiTenant(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigMultiTenantSlice expands the contents of ConfigMultiTenant into a JSON
// request object.
func expandConfigMultiTenantSlice(c *Client, f []ConfigMultiTenant, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigMultiTenant(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigMultiTenantMap flattens the contents of ConfigMultiTenant from a JSON
// response object.
func flattenConfigMultiTenantMap(c *Client, i interface{}, res *Config) map[string]ConfigMultiTenant {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigMultiTenant{}
	}

	if len(a) == 0 {
		return map[string]ConfigMultiTenant{}
	}

	items := make(map[string]ConfigMultiTenant)
	for k, item := range a {
		items[k] = *flattenConfigMultiTenant(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigMultiTenantSlice flattens the contents of ConfigMultiTenant from a JSON
// response object.
func flattenConfigMultiTenantSlice(c *Client, i interface{}, res *Config) []ConfigMultiTenant {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigMultiTenant{}
	}

	if len(a) == 0 {
		return []ConfigMultiTenant{}
	}

	items := make([]ConfigMultiTenant, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigMultiTenant(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigMultiTenant expands an instance of ConfigMultiTenant into a JSON
// request object.
func expandConfigMultiTenant(c *Client, f *ConfigMultiTenant, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.AllowTenants; !dcl.IsEmptyValueIndirect(v) {
		m["allowTenants"] = v
	}
	if v := f.DefaultTenantLocation; !dcl.IsEmptyValueIndirect(v) {
		m["defaultTenantLocation"] = v
	}

	return m, nil
}

// flattenConfigMultiTenant flattens an instance of ConfigMultiTenant from a JSON
// response object.
func flattenConfigMultiTenant(c *Client, i interface{}, res *Config) *ConfigMultiTenant {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigMultiTenant{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigMultiTenant
	}
	r.AllowTenants = dcl.FlattenBool(m["allowTenants"])
	r.DefaultTenantLocation = dcl.FlattenString(m["defaultTenantLocation"])

	return r
}

// expandConfigClientMap expands the contents of ConfigClient into a JSON
// request object.
func expandConfigClientMap(c *Client, f map[string]ConfigClient, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigClient(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigClientSlice expands the contents of ConfigClient into a JSON
// request object.
func expandConfigClientSlice(c *Client, f []ConfigClient, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigClient(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigClientMap flattens the contents of ConfigClient from a JSON
// response object.
func flattenConfigClientMap(c *Client, i interface{}, res *Config) map[string]ConfigClient {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigClient{}
	}

	if len(a) == 0 {
		return map[string]ConfigClient{}
	}

	items := make(map[string]ConfigClient)
	for k, item := range a {
		items[k] = *flattenConfigClient(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigClientSlice flattens the contents of ConfigClient from a JSON
// response object.
func flattenConfigClientSlice(c *Client, i interface{}, res *Config) []ConfigClient {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigClient{}
	}

	if len(a) == 0 {
		return []ConfigClient{}
	}

	items := make([]ConfigClient, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigClient(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigClient expands an instance of ConfigClient into a JSON
// request object.
func expandConfigClient(c *Client, f *ConfigClient, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigClientPermissions(c, f.Permissions, res); err != nil {
		return nil, fmt.Errorf("error expanding Permissions into permissions: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["permissions"] = v
	}

	return m, nil
}

// flattenConfigClient flattens an instance of ConfigClient from a JSON
// response object.
func flattenConfigClient(c *Client, i interface{}, res *Config) *ConfigClient {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigClient{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigClient
	}
	r.ApiKey = dcl.FlattenString(m["apiKey"])
	r.Permissions = flattenConfigClientPermissions(c, m["permissions"], res)
	r.FirebaseSubdomain = dcl.FlattenString(m["firebaseSubdomain"])

	return r
}

// expandConfigClientPermissionsMap expands the contents of ConfigClientPermissions into a JSON
// request object.
func expandConfigClientPermissionsMap(c *Client, f map[string]ConfigClientPermissions, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigClientPermissions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigClientPermissionsSlice expands the contents of ConfigClientPermissions into a JSON
// request object.
func expandConfigClientPermissionsSlice(c *Client, f []ConfigClientPermissions, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigClientPermissions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigClientPermissionsMap flattens the contents of ConfigClientPermissions from a JSON
// response object.
func flattenConfigClientPermissionsMap(c *Client, i interface{}, res *Config) map[string]ConfigClientPermissions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigClientPermissions{}
	}

	if len(a) == 0 {
		return map[string]ConfigClientPermissions{}
	}

	items := make(map[string]ConfigClientPermissions)
	for k, item := range a {
		items[k] = *flattenConfigClientPermissions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigClientPermissionsSlice flattens the contents of ConfigClientPermissions from a JSON
// response object.
func flattenConfigClientPermissionsSlice(c *Client, i interface{}, res *Config) []ConfigClientPermissions {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigClientPermissions{}
	}

	if len(a) == 0 {
		return []ConfigClientPermissions{}
	}

	items := make([]ConfigClientPermissions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigClientPermissions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigClientPermissions expands an instance of ConfigClientPermissions into a JSON
// request object.
func expandConfigClientPermissions(c *Client, f *ConfigClientPermissions, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.DisabledUserSignup; !dcl.IsEmptyValueIndirect(v) {
		m["disabledUserSignup"] = v
	}
	if v := f.DisabledUserDeletion; !dcl.IsEmptyValueIndirect(v) {
		m["disabledUserDeletion"] = v
	}

	return m, nil
}

// flattenConfigClientPermissions flattens an instance of ConfigClientPermissions from a JSON
// response object.
func flattenConfigClientPermissions(c *Client, i interface{}, res *Config) *ConfigClientPermissions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigClientPermissions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigClientPermissions
	}
	r.DisabledUserSignup = dcl.FlattenBool(m["disabledUserSignup"])
	r.DisabledUserDeletion = dcl.FlattenBool(m["disabledUserDeletion"])

	return r
}

// expandConfigMfaMap expands the contents of ConfigMfa into a JSON
// request object.
func expandConfigMfaMap(c *Client, f map[string]ConfigMfa, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigMfa(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigMfaSlice expands the contents of ConfigMfa into a JSON
// request object.
func expandConfigMfaSlice(c *Client, f []ConfigMfa, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigMfa(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigMfaMap flattens the contents of ConfigMfa from a JSON
// response object.
func flattenConfigMfaMap(c *Client, i interface{}, res *Config) map[string]ConfigMfa {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigMfa{}
	}

	if len(a) == 0 {
		return map[string]ConfigMfa{}
	}

	items := make(map[string]ConfigMfa)
	for k, item := range a {
		items[k] = *flattenConfigMfa(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigMfaSlice flattens the contents of ConfigMfa from a JSON
// response object.
func flattenConfigMfaSlice(c *Client, i interface{}, res *Config) []ConfigMfa {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigMfa{}
	}

	if len(a) == 0 {
		return []ConfigMfa{}
	}

	items := make([]ConfigMfa, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigMfa(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigMfa expands an instance of ConfigMfa into a JSON
// request object.
func expandConfigMfa(c *Client, f *ConfigMfa, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.State; !dcl.IsEmptyValueIndirect(v) {
		m["state"] = v
	}

	return m, nil
}

// flattenConfigMfa flattens an instance of ConfigMfa from a JSON
// response object.
func flattenConfigMfa(c *Client, i interface{}, res *Config) *ConfigMfa {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigMfa{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigMfa
	}
	r.State = flattenConfigMfaStateEnum(m["state"])

	return r
}

// expandConfigBlockingFunctionsMap expands the contents of ConfigBlockingFunctions into a JSON
// request object.
func expandConfigBlockingFunctionsMap(c *Client, f map[string]ConfigBlockingFunctions, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigBlockingFunctions(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigBlockingFunctionsSlice expands the contents of ConfigBlockingFunctions into a JSON
// request object.
func expandConfigBlockingFunctionsSlice(c *Client, f []ConfigBlockingFunctions, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigBlockingFunctions(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigBlockingFunctionsMap flattens the contents of ConfigBlockingFunctions from a JSON
// response object.
func flattenConfigBlockingFunctionsMap(c *Client, i interface{}, res *Config) map[string]ConfigBlockingFunctions {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigBlockingFunctions{}
	}

	if len(a) == 0 {
		return map[string]ConfigBlockingFunctions{}
	}

	items := make(map[string]ConfigBlockingFunctions)
	for k, item := range a {
		items[k] = *flattenConfigBlockingFunctions(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigBlockingFunctionsSlice flattens the contents of ConfigBlockingFunctions from a JSON
// response object.
func flattenConfigBlockingFunctionsSlice(c *Client, i interface{}, res *Config) []ConfigBlockingFunctions {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigBlockingFunctions{}
	}

	if len(a) == 0 {
		return []ConfigBlockingFunctions{}
	}

	items := make([]ConfigBlockingFunctions, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigBlockingFunctions(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigBlockingFunctions expands an instance of ConfigBlockingFunctions into a JSON
// request object.
func expandConfigBlockingFunctions(c *Client, f *ConfigBlockingFunctions, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandConfigBlockingFunctionsTriggersMap(c, f.Triggers, res); err != nil {
		return nil, fmt.Errorf("error expanding Triggers into triggers: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["triggers"] = v
	}

	return m, nil
}

// flattenConfigBlockingFunctions flattens an instance of ConfigBlockingFunctions from a JSON
// response object.
func flattenConfigBlockingFunctions(c *Client, i interface{}, res *Config) *ConfigBlockingFunctions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigBlockingFunctions{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigBlockingFunctions
	}
	r.Triggers = flattenConfigBlockingFunctionsTriggersMap(c, m["triggers"], res)

	return r
}

// expandConfigBlockingFunctionsTriggersMap expands the contents of ConfigBlockingFunctionsTriggers into a JSON
// request object.
func expandConfigBlockingFunctionsTriggersMap(c *Client, f map[string]ConfigBlockingFunctionsTriggers, res *Config) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConfigBlockingFunctionsTriggers(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConfigBlockingFunctionsTriggersSlice expands the contents of ConfigBlockingFunctionsTriggers into a JSON
// request object.
func expandConfigBlockingFunctionsTriggersSlice(c *Client, f []ConfigBlockingFunctionsTriggers, res *Config) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConfigBlockingFunctionsTriggers(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConfigBlockingFunctionsTriggersMap flattens the contents of ConfigBlockingFunctionsTriggers from a JSON
// response object.
func flattenConfigBlockingFunctionsTriggersMap(c *Client, i interface{}, res *Config) map[string]ConfigBlockingFunctionsTriggers {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigBlockingFunctionsTriggers{}
	}

	if len(a) == 0 {
		return map[string]ConfigBlockingFunctionsTriggers{}
	}

	items := make(map[string]ConfigBlockingFunctionsTriggers)
	for k, item := range a {
		items[k] = *flattenConfigBlockingFunctionsTriggers(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConfigBlockingFunctionsTriggersSlice flattens the contents of ConfigBlockingFunctionsTriggers from a JSON
// response object.
func flattenConfigBlockingFunctionsTriggersSlice(c *Client, i interface{}, res *Config) []ConfigBlockingFunctionsTriggers {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigBlockingFunctionsTriggers{}
	}

	if len(a) == 0 {
		return []ConfigBlockingFunctionsTriggers{}
	}

	items := make([]ConfigBlockingFunctionsTriggers, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigBlockingFunctionsTriggers(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConfigBlockingFunctionsTriggers expands an instance of ConfigBlockingFunctionsTriggers into a JSON
// request object.
func expandConfigBlockingFunctionsTriggers(c *Client, f *ConfigBlockingFunctionsTriggers, res *Config) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.FunctionUri; !dcl.IsEmptyValueIndirect(v) {
		m["functionUri"] = v
	}

	return m, nil
}

// flattenConfigBlockingFunctionsTriggers flattens an instance of ConfigBlockingFunctionsTriggers from a JSON
// response object.
func flattenConfigBlockingFunctionsTriggers(c *Client, i interface{}, res *Config) *ConfigBlockingFunctionsTriggers {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConfigBlockingFunctionsTriggers{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConfigBlockingFunctionsTriggers
	}
	r.FunctionUri = dcl.FlattenString(m["functionUri"])
	r.UpdateTime = dcl.FlattenString(m["updateTime"])

	return r
}

// flattenConfigSignInEmailHashConfigAlgorithmEnumMap flattens the contents of ConfigSignInEmailHashConfigAlgorithmEnum from a JSON
// response object.
func flattenConfigSignInEmailHashConfigAlgorithmEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInEmailHashConfigAlgorithmEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInEmailHashConfigAlgorithmEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInEmailHashConfigAlgorithmEnum{}
	}

	items := make(map[string]ConfigSignInEmailHashConfigAlgorithmEnum)
	for k, item := range a {
		items[k] = *flattenConfigSignInEmailHashConfigAlgorithmEnum(item.(interface{}))
	}

	return items
}

// flattenConfigSignInEmailHashConfigAlgorithmEnumSlice flattens the contents of ConfigSignInEmailHashConfigAlgorithmEnum from a JSON
// response object.
func flattenConfigSignInEmailHashConfigAlgorithmEnumSlice(c *Client, i interface{}, res *Config) []ConfigSignInEmailHashConfigAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInEmailHashConfigAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []ConfigSignInEmailHashConfigAlgorithmEnum{}
	}

	items := make([]ConfigSignInEmailHashConfigAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInEmailHashConfigAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenConfigSignInEmailHashConfigAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigSignInEmailHashConfigAlgorithmEnum with the same value as that string.
func flattenConfigSignInEmailHashConfigAlgorithmEnum(i interface{}) *ConfigSignInEmailHashConfigAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigSignInEmailHashConfigAlgorithmEnumRef(s)
}

// flattenConfigSignInHashConfigAlgorithmEnumMap flattens the contents of ConfigSignInHashConfigAlgorithmEnum from a JSON
// response object.
func flattenConfigSignInHashConfigAlgorithmEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigSignInHashConfigAlgorithmEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSignInHashConfigAlgorithmEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigSignInHashConfigAlgorithmEnum{}
	}

	items := make(map[string]ConfigSignInHashConfigAlgorithmEnum)
	for k, item := range a {
		items[k] = *flattenConfigSignInHashConfigAlgorithmEnum(item.(interface{}))
	}

	return items
}

// flattenConfigSignInHashConfigAlgorithmEnumSlice flattens the contents of ConfigSignInHashConfigAlgorithmEnum from a JSON
// response object.
func flattenConfigSignInHashConfigAlgorithmEnumSlice(c *Client, i interface{}, res *Config) []ConfigSignInHashConfigAlgorithmEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSignInHashConfigAlgorithmEnum{}
	}

	if len(a) == 0 {
		return []ConfigSignInHashConfigAlgorithmEnum{}
	}

	items := make([]ConfigSignInHashConfigAlgorithmEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSignInHashConfigAlgorithmEnum(item.(interface{})))
	}

	return items
}

// flattenConfigSignInHashConfigAlgorithmEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigSignInHashConfigAlgorithmEnum with the same value as that string.
func flattenConfigSignInHashConfigAlgorithmEnum(i interface{}) *ConfigSignInHashConfigAlgorithmEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigSignInHashConfigAlgorithmEnumRef(s)
}

// flattenConfigNotificationSendEmailMethodEnumMap flattens the contents of ConfigNotificationSendEmailMethodEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailMethodEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailMethodEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailMethodEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailMethodEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailMethodEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailMethodEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailMethodEnumSlice flattens the contents of ConfigNotificationSendEmailMethodEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailMethodEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailMethodEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailMethodEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailMethodEnum{}
	}

	items := make([]ConfigNotificationSendEmailMethodEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailMethodEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailMethodEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailMethodEnum with the same value as that string.
func flattenConfigNotificationSendEmailMethodEnum(i interface{}) *ConfigNotificationSendEmailMethodEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailMethodEnumRef(s)
}

// flattenConfigNotificationSendEmailSmtpSecurityModeEnumMap flattens the contents of ConfigNotificationSendEmailSmtpSecurityModeEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailSmtpSecurityModeEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailSmtpSecurityModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailSmtpSecurityModeEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailSmtpSecurityModeEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailSmtpSecurityModeEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailSmtpSecurityModeEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailSmtpSecurityModeEnumSlice flattens the contents of ConfigNotificationSendEmailSmtpSecurityModeEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailSmtpSecurityModeEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailSmtpSecurityModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailSmtpSecurityModeEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailSmtpSecurityModeEnum{}
	}

	items := make([]ConfigNotificationSendEmailSmtpSecurityModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailSmtpSecurityModeEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailSmtpSecurityModeEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailSmtpSecurityModeEnum with the same value as that string.
func flattenConfigNotificationSendEmailSmtpSecurityModeEnum(i interface{}) *ConfigNotificationSendEmailSmtpSecurityModeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailSmtpSecurityModeEnumRef(s)
}

// flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumMap flattens the contents of ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumSlice flattens the contents of ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum{}
	}

	items := make([]ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum with the same value as that string.
func flattenConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum(i interface{}) *ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailResetPasswordTemplateBodyFormatEnumRef(s)
}

// flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumMap flattens the contents of ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumSlice flattens the contents of ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum{}
	}

	items := make([]ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum with the same value as that string.
func flattenConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum(i interface{}) *ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailVerifyEmailTemplateBodyFormatEnumRef(s)
}

// flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumMap flattens the contents of ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumSlice flattens the contents of ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum{}
	}

	items := make([]ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum with the same value as that string.
func flattenConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum(i interface{}) *ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailChangeEmailTemplateBodyFormatEnumRef(s)
}

// flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnumMap flattens the contents of ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnumSlice flattens the contents of ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum{}
	}

	items := make([]ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum with the same value as that string.
func flattenConfigNotificationSendEmailDnsInfoCustomDomainStateEnum(i interface{}) *ConfigNotificationSendEmailDnsInfoCustomDomainStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailDnsInfoCustomDomainStateEnumRef(s)
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumMap flattens the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum{}
	}

	items := make(map[string]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum)
	for k, item := range a {
		items[k] = *flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(item.(interface{}))
	}

	return items
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumSlice flattens the contents of ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum from a JSON
// response object.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumSlice(c *Client, i interface{}, res *Config) []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum{}
	}

	if len(a) == 0 {
		return []ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum{}
	}

	items := make([]ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(item.(interface{})))
	}

	return items
}

// flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum with the same value as that string.
func flattenConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum(i interface{}) *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigNotificationSendEmailRevertSecondFactorAdditionTemplateBodyFormatEnumRef(s)
}

// flattenConfigSubtypeEnumMap flattens the contents of ConfigSubtypeEnum from a JSON
// response object.
func flattenConfigSubtypeEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigSubtypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigSubtypeEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigSubtypeEnum{}
	}

	items := make(map[string]ConfigSubtypeEnum)
	for k, item := range a {
		items[k] = *flattenConfigSubtypeEnum(item.(interface{}))
	}

	return items
}

// flattenConfigSubtypeEnumSlice flattens the contents of ConfigSubtypeEnum from a JSON
// response object.
func flattenConfigSubtypeEnumSlice(c *Client, i interface{}, res *Config) []ConfigSubtypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigSubtypeEnum{}
	}

	if len(a) == 0 {
		return []ConfigSubtypeEnum{}
	}

	items := make([]ConfigSubtypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigSubtypeEnum(item.(interface{})))
	}

	return items
}

// flattenConfigSubtypeEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigSubtypeEnum with the same value as that string.
func flattenConfigSubtypeEnum(i interface{}) *ConfigSubtypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigSubtypeEnumRef(s)
}

// flattenConfigMfaStateEnumMap flattens the contents of ConfigMfaStateEnum from a JSON
// response object.
func flattenConfigMfaStateEnumMap(c *Client, i interface{}, res *Config) map[string]ConfigMfaStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConfigMfaStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ConfigMfaStateEnum{}
	}

	items := make(map[string]ConfigMfaStateEnum)
	for k, item := range a {
		items[k] = *flattenConfigMfaStateEnum(item.(interface{}))
	}

	return items
}

// flattenConfigMfaStateEnumSlice flattens the contents of ConfigMfaStateEnum from a JSON
// response object.
func flattenConfigMfaStateEnumSlice(c *Client, i interface{}, res *Config) []ConfigMfaStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConfigMfaStateEnum{}
	}

	if len(a) == 0 {
		return []ConfigMfaStateEnum{}
	}

	items := make([]ConfigMfaStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConfigMfaStateEnum(item.(interface{})))
	}

	return items
}

// flattenConfigMfaStateEnum asserts that an interface is a string, and returns a
// pointer to a *ConfigMfaStateEnum with the same value as that string.
func flattenConfigMfaStateEnum(i interface{}) *ConfigMfaStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConfigMfaStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Config) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalConfig(b, c, r)
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

type configDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         configApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToConfigDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]configDiff, error) {
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
	var diffs []configDiff
	// For each operation name, create a configDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := configDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToConfigApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToConfigApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (configApiOperation, error) {
	switch opName {

	case "updateConfigUpdateProjectConfigOperation":
		return &updateConfigUpdateProjectConfigOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractConfigFields(r *Config) error {
	vSignIn := r.SignIn
	if vSignIn == nil {
		// note: explicitly not the empty object.
		vSignIn = &ConfigSignIn{}
	}
	if err := extractConfigSignInFields(r, vSignIn); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignIn) {
		r.SignIn = vSignIn
	}
	vNotification := r.Notification
	if vNotification == nil {
		// note: explicitly not the empty object.
		vNotification = &ConfigNotification{}
	}
	if err := extractConfigNotificationFields(r, vNotification); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotification) {
		r.Notification = vNotification
	}
	vQuota := r.Quota
	if vQuota == nil {
		// note: explicitly not the empty object.
		vQuota = &ConfigQuota{}
	}
	if err := extractConfigQuotaFields(r, vQuota); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vQuota) {
		r.Quota = vQuota
	}
	vMonitoring := r.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &ConfigMonitoring{}
	}
	if err := extractConfigMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		r.Monitoring = vMonitoring
	}
	vMultiTenant := r.MultiTenant
	if vMultiTenant == nil {
		// note: explicitly not the empty object.
		vMultiTenant = &ConfigMultiTenant{}
	}
	if err := extractConfigMultiTenantFields(r, vMultiTenant); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMultiTenant) {
		r.MultiTenant = vMultiTenant
	}
	vClient := r.Client
	if vClient == nil {
		// note: explicitly not the empty object.
		vClient = &ConfigClient{}
	}
	if err := extractConfigClientFields(r, vClient); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vClient) {
		r.Client = vClient
	}
	vMfa := r.Mfa
	if vMfa == nil {
		// note: explicitly not the empty object.
		vMfa = &ConfigMfa{}
	}
	if err := extractConfigMfaFields(r, vMfa); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMfa) {
		r.Mfa = vMfa
	}
	vBlockingFunctions := r.BlockingFunctions
	if vBlockingFunctions == nil {
		// note: explicitly not the empty object.
		vBlockingFunctions = &ConfigBlockingFunctions{}
	}
	if err := extractConfigBlockingFunctionsFields(r, vBlockingFunctions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBlockingFunctions) {
		r.BlockingFunctions = vBlockingFunctions
	}
	return nil
}
func extractConfigSignInFields(r *Config, o *ConfigSignIn) error {
	vEmail := o.Email
	if vEmail == nil {
		// note: explicitly not the empty object.
		vEmail = &ConfigSignInEmail{}
	}
	if err := extractConfigSignInEmailFields(r, vEmail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEmail) {
		o.Email = vEmail
	}
	vPhoneNumber := o.PhoneNumber
	if vPhoneNumber == nil {
		// note: explicitly not the empty object.
		vPhoneNumber = &ConfigSignInPhoneNumber{}
	}
	if err := extractConfigSignInPhoneNumberFields(r, vPhoneNumber); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPhoneNumber) {
		o.PhoneNumber = vPhoneNumber
	}
	vAnonymous := o.Anonymous
	if vAnonymous == nil {
		// note: explicitly not the empty object.
		vAnonymous = &ConfigSignInAnonymous{}
	}
	if err := extractConfigSignInAnonymousFields(r, vAnonymous); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAnonymous) {
		o.Anonymous = vAnonymous
	}
	vHashConfig := o.HashConfig
	if vHashConfig == nil {
		// note: explicitly not the empty object.
		vHashConfig = &ConfigSignInHashConfig{}
	}
	if err := extractConfigSignInHashConfigFields(r, vHashConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHashConfig) {
		o.HashConfig = vHashConfig
	}
	return nil
}
func extractConfigSignInEmailFields(r *Config, o *ConfigSignInEmail) error {
	vHashConfig := o.HashConfig
	if vHashConfig == nil {
		// note: explicitly not the empty object.
		vHashConfig = &ConfigSignInEmailHashConfig{}
	}
	if err := extractConfigSignInEmailHashConfigFields(r, vHashConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHashConfig) {
		o.HashConfig = vHashConfig
	}
	return nil
}
func extractConfigSignInEmailHashConfigFields(r *Config, o *ConfigSignInEmailHashConfig) error {
	return nil
}
func extractConfigSignInPhoneNumberFields(r *Config, o *ConfigSignInPhoneNumber) error {
	return nil
}
func extractConfigSignInAnonymousFields(r *Config, o *ConfigSignInAnonymous) error {
	return nil
}
func extractConfigSignInHashConfigFields(r *Config, o *ConfigSignInHashConfig) error {
	return nil
}
func extractConfigNotificationFields(r *Config, o *ConfigNotification) error {
	vSendEmail := o.SendEmail
	if vSendEmail == nil {
		// note: explicitly not the empty object.
		vSendEmail = &ConfigNotificationSendEmail{}
	}
	if err := extractConfigNotificationSendEmailFields(r, vSendEmail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSendEmail) {
		o.SendEmail = vSendEmail
	}
	vSendSms := o.SendSms
	if vSendSms == nil {
		// note: explicitly not the empty object.
		vSendSms = &ConfigNotificationSendSms{}
	}
	if err := extractConfigNotificationSendSmsFields(r, vSendSms); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSendSms) {
		o.SendSms = vSendSms
	}
	return nil
}
func extractConfigNotificationSendEmailFields(r *Config, o *ConfigNotificationSendEmail) error {
	vSmtp := o.Smtp
	if vSmtp == nil {
		// note: explicitly not the empty object.
		vSmtp = &ConfigNotificationSendEmailSmtp{}
	}
	if err := extractConfigNotificationSendEmailSmtpFields(r, vSmtp); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSmtp) {
		o.Smtp = vSmtp
	}
	vResetPasswordTemplate := o.ResetPasswordTemplate
	if vResetPasswordTemplate == nil {
		// note: explicitly not the empty object.
		vResetPasswordTemplate = &ConfigNotificationSendEmailResetPasswordTemplate{}
	}
	if err := extractConfigNotificationSendEmailResetPasswordTemplateFields(r, vResetPasswordTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResetPasswordTemplate) {
		o.ResetPasswordTemplate = vResetPasswordTemplate
	}
	vVerifyEmailTemplate := o.VerifyEmailTemplate
	if vVerifyEmailTemplate == nil {
		// note: explicitly not the empty object.
		vVerifyEmailTemplate = &ConfigNotificationSendEmailVerifyEmailTemplate{}
	}
	if err := extractConfigNotificationSendEmailVerifyEmailTemplateFields(r, vVerifyEmailTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVerifyEmailTemplate) {
		o.VerifyEmailTemplate = vVerifyEmailTemplate
	}
	vChangeEmailTemplate := o.ChangeEmailTemplate
	if vChangeEmailTemplate == nil {
		// note: explicitly not the empty object.
		vChangeEmailTemplate = &ConfigNotificationSendEmailChangeEmailTemplate{}
	}
	if err := extractConfigNotificationSendEmailChangeEmailTemplateFields(r, vChangeEmailTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vChangeEmailTemplate) {
		o.ChangeEmailTemplate = vChangeEmailTemplate
	}
	vDnsInfo := o.DnsInfo
	if vDnsInfo == nil {
		// note: explicitly not the empty object.
		vDnsInfo = &ConfigNotificationSendEmailDnsInfo{}
	}
	if err := extractConfigNotificationSendEmailDnsInfoFields(r, vDnsInfo); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDnsInfo) {
		o.DnsInfo = vDnsInfo
	}
	vRevertSecondFactorAdditionTemplate := o.RevertSecondFactorAdditionTemplate
	if vRevertSecondFactorAdditionTemplate == nil {
		// note: explicitly not the empty object.
		vRevertSecondFactorAdditionTemplate = &ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}
	if err := extractConfigNotificationSendEmailRevertSecondFactorAdditionTemplateFields(r, vRevertSecondFactorAdditionTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRevertSecondFactorAdditionTemplate) {
		o.RevertSecondFactorAdditionTemplate = vRevertSecondFactorAdditionTemplate
	}
	return nil
}
func extractConfigNotificationSendEmailSmtpFields(r *Config, o *ConfigNotificationSendEmailSmtp) error {
	return nil
}
func extractConfigNotificationSendEmailResetPasswordTemplateFields(r *Config, o *ConfigNotificationSendEmailResetPasswordTemplate) error {
	return nil
}
func extractConfigNotificationSendEmailVerifyEmailTemplateFields(r *Config, o *ConfigNotificationSendEmailVerifyEmailTemplate) error {
	return nil
}
func extractConfigNotificationSendEmailChangeEmailTemplateFields(r *Config, o *ConfigNotificationSendEmailChangeEmailTemplate) error {
	return nil
}
func extractConfigNotificationSendEmailDnsInfoFields(r *Config, o *ConfigNotificationSendEmailDnsInfo) error {
	return nil
}
func extractConfigNotificationSendEmailRevertSecondFactorAdditionTemplateFields(r *Config, o *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) error {
	return nil
}
func extractConfigNotificationSendSmsFields(r *Config, o *ConfigNotificationSendSms) error {
	vSmsTemplate := o.SmsTemplate
	if vSmsTemplate == nil {
		// note: explicitly not the empty object.
		vSmsTemplate = &ConfigNotificationSendSmsSmsTemplate{}
	}
	if err := extractConfigNotificationSendSmsSmsTemplateFields(r, vSmsTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSmsTemplate) {
		o.SmsTemplate = vSmsTemplate
	}
	return nil
}
func extractConfigNotificationSendSmsSmsTemplateFields(r *Config, o *ConfigNotificationSendSmsSmsTemplate) error {
	return nil
}
func extractConfigQuotaFields(r *Config, o *ConfigQuota) error {
	vSignUpQuotaConfig := o.SignUpQuotaConfig
	if vSignUpQuotaConfig == nil {
		// note: explicitly not the empty object.
		vSignUpQuotaConfig = &ConfigQuotaSignUpQuotaConfig{}
	}
	if err := extractConfigQuotaSignUpQuotaConfigFields(r, vSignUpQuotaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignUpQuotaConfig) {
		o.SignUpQuotaConfig = vSignUpQuotaConfig
	}
	return nil
}
func extractConfigQuotaSignUpQuotaConfigFields(r *Config, o *ConfigQuotaSignUpQuotaConfig) error {
	return nil
}
func extractConfigMonitoringFields(r *Config, o *ConfigMonitoring) error {
	vRequestLogging := o.RequestLogging
	if vRequestLogging == nil {
		// note: explicitly not the empty object.
		vRequestLogging = &ConfigMonitoringRequestLogging{}
	}
	if err := extractConfigMonitoringRequestLoggingFields(r, vRequestLogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestLogging) {
		o.RequestLogging = vRequestLogging
	}
	return nil
}
func extractConfigMonitoringRequestLoggingFields(r *Config, o *ConfigMonitoringRequestLogging) error {
	return nil
}
func extractConfigMultiTenantFields(r *Config, o *ConfigMultiTenant) error {
	return nil
}
func extractConfigClientFields(r *Config, o *ConfigClient) error {
	vPermissions := o.Permissions
	if vPermissions == nil {
		// note: explicitly not the empty object.
		vPermissions = &ConfigClientPermissions{}
	}
	if err := extractConfigClientPermissionsFields(r, vPermissions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPermissions) {
		o.Permissions = vPermissions
	}
	return nil
}
func extractConfigClientPermissionsFields(r *Config, o *ConfigClientPermissions) error {
	return nil
}
func extractConfigMfaFields(r *Config, o *ConfigMfa) error {
	return nil
}
func extractConfigBlockingFunctionsFields(r *Config, o *ConfigBlockingFunctions) error {
	return nil
}
func extractConfigBlockingFunctionsTriggersFields(r *Config, o *ConfigBlockingFunctionsTriggers) error {
	return nil
}

func postReadExtractConfigFields(r *Config) error {
	vSignIn := r.SignIn
	if vSignIn == nil {
		// note: explicitly not the empty object.
		vSignIn = &ConfigSignIn{}
	}
	if err := postReadExtractConfigSignInFields(r, vSignIn); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignIn) {
		r.SignIn = vSignIn
	}
	vNotification := r.Notification
	if vNotification == nil {
		// note: explicitly not the empty object.
		vNotification = &ConfigNotification{}
	}
	if err := postReadExtractConfigNotificationFields(r, vNotification); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vNotification) {
		r.Notification = vNotification
	}
	vQuota := r.Quota
	if vQuota == nil {
		// note: explicitly not the empty object.
		vQuota = &ConfigQuota{}
	}
	if err := postReadExtractConfigQuotaFields(r, vQuota); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vQuota) {
		r.Quota = vQuota
	}
	vMonitoring := r.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &ConfigMonitoring{}
	}
	if err := postReadExtractConfigMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		r.Monitoring = vMonitoring
	}
	vMultiTenant := r.MultiTenant
	if vMultiTenant == nil {
		// note: explicitly not the empty object.
		vMultiTenant = &ConfigMultiTenant{}
	}
	if err := postReadExtractConfigMultiTenantFields(r, vMultiTenant); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMultiTenant) {
		r.MultiTenant = vMultiTenant
	}
	vClient := r.Client
	if vClient == nil {
		// note: explicitly not the empty object.
		vClient = &ConfigClient{}
	}
	if err := postReadExtractConfigClientFields(r, vClient); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vClient) {
		r.Client = vClient
	}
	vMfa := r.Mfa
	if vMfa == nil {
		// note: explicitly not the empty object.
		vMfa = &ConfigMfa{}
	}
	if err := postReadExtractConfigMfaFields(r, vMfa); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMfa) {
		r.Mfa = vMfa
	}
	vBlockingFunctions := r.BlockingFunctions
	if vBlockingFunctions == nil {
		// note: explicitly not the empty object.
		vBlockingFunctions = &ConfigBlockingFunctions{}
	}
	if err := postReadExtractConfigBlockingFunctionsFields(r, vBlockingFunctions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBlockingFunctions) {
		r.BlockingFunctions = vBlockingFunctions
	}
	return nil
}
func postReadExtractConfigSignInFields(r *Config, o *ConfigSignIn) error {
	vEmail := o.Email
	if vEmail == nil {
		// note: explicitly not the empty object.
		vEmail = &ConfigSignInEmail{}
	}
	if err := extractConfigSignInEmailFields(r, vEmail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vEmail) {
		o.Email = vEmail
	}
	vPhoneNumber := o.PhoneNumber
	if vPhoneNumber == nil {
		// note: explicitly not the empty object.
		vPhoneNumber = &ConfigSignInPhoneNumber{}
	}
	if err := extractConfigSignInPhoneNumberFields(r, vPhoneNumber); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPhoneNumber) {
		o.PhoneNumber = vPhoneNumber
	}
	vAnonymous := o.Anonymous
	if vAnonymous == nil {
		// note: explicitly not the empty object.
		vAnonymous = &ConfigSignInAnonymous{}
	}
	if err := extractConfigSignInAnonymousFields(r, vAnonymous); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAnonymous) {
		o.Anonymous = vAnonymous
	}
	vHashConfig := o.HashConfig
	if vHashConfig == nil {
		// note: explicitly not the empty object.
		vHashConfig = &ConfigSignInHashConfig{}
	}
	if err := extractConfigSignInHashConfigFields(r, vHashConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHashConfig) {
		o.HashConfig = vHashConfig
	}
	return nil
}
func postReadExtractConfigSignInEmailFields(r *Config, o *ConfigSignInEmail) error {
	vHashConfig := o.HashConfig
	if vHashConfig == nil {
		// note: explicitly not the empty object.
		vHashConfig = &ConfigSignInEmailHashConfig{}
	}
	if err := extractConfigSignInEmailHashConfigFields(r, vHashConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHashConfig) {
		o.HashConfig = vHashConfig
	}
	return nil
}
func postReadExtractConfigSignInEmailHashConfigFields(r *Config, o *ConfigSignInEmailHashConfig) error {
	return nil
}
func postReadExtractConfigSignInPhoneNumberFields(r *Config, o *ConfigSignInPhoneNumber) error {
	return nil
}
func postReadExtractConfigSignInAnonymousFields(r *Config, o *ConfigSignInAnonymous) error {
	return nil
}
func postReadExtractConfigSignInHashConfigFields(r *Config, o *ConfigSignInHashConfig) error {
	return nil
}
func postReadExtractConfigNotificationFields(r *Config, o *ConfigNotification) error {
	vSendEmail := o.SendEmail
	if vSendEmail == nil {
		// note: explicitly not the empty object.
		vSendEmail = &ConfigNotificationSendEmail{}
	}
	if err := extractConfigNotificationSendEmailFields(r, vSendEmail); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSendEmail) {
		o.SendEmail = vSendEmail
	}
	vSendSms := o.SendSms
	if vSendSms == nil {
		// note: explicitly not the empty object.
		vSendSms = &ConfigNotificationSendSms{}
	}
	if err := extractConfigNotificationSendSmsFields(r, vSendSms); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSendSms) {
		o.SendSms = vSendSms
	}
	return nil
}
func postReadExtractConfigNotificationSendEmailFields(r *Config, o *ConfigNotificationSendEmail) error {
	vSmtp := o.Smtp
	if vSmtp == nil {
		// note: explicitly not the empty object.
		vSmtp = &ConfigNotificationSendEmailSmtp{}
	}
	if err := extractConfigNotificationSendEmailSmtpFields(r, vSmtp); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSmtp) {
		o.Smtp = vSmtp
	}
	vResetPasswordTemplate := o.ResetPasswordTemplate
	if vResetPasswordTemplate == nil {
		// note: explicitly not the empty object.
		vResetPasswordTemplate = &ConfigNotificationSendEmailResetPasswordTemplate{}
	}
	if err := extractConfigNotificationSendEmailResetPasswordTemplateFields(r, vResetPasswordTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResetPasswordTemplate) {
		o.ResetPasswordTemplate = vResetPasswordTemplate
	}
	vVerifyEmailTemplate := o.VerifyEmailTemplate
	if vVerifyEmailTemplate == nil {
		// note: explicitly not the empty object.
		vVerifyEmailTemplate = &ConfigNotificationSendEmailVerifyEmailTemplate{}
	}
	if err := extractConfigNotificationSendEmailVerifyEmailTemplateFields(r, vVerifyEmailTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVerifyEmailTemplate) {
		o.VerifyEmailTemplate = vVerifyEmailTemplate
	}
	vChangeEmailTemplate := o.ChangeEmailTemplate
	if vChangeEmailTemplate == nil {
		// note: explicitly not the empty object.
		vChangeEmailTemplate = &ConfigNotificationSendEmailChangeEmailTemplate{}
	}
	if err := extractConfigNotificationSendEmailChangeEmailTemplateFields(r, vChangeEmailTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vChangeEmailTemplate) {
		o.ChangeEmailTemplate = vChangeEmailTemplate
	}
	vDnsInfo := o.DnsInfo
	if vDnsInfo == nil {
		// note: explicitly not the empty object.
		vDnsInfo = &ConfigNotificationSendEmailDnsInfo{}
	}
	if err := extractConfigNotificationSendEmailDnsInfoFields(r, vDnsInfo); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDnsInfo) {
		o.DnsInfo = vDnsInfo
	}
	vRevertSecondFactorAdditionTemplate := o.RevertSecondFactorAdditionTemplate
	if vRevertSecondFactorAdditionTemplate == nil {
		// note: explicitly not the empty object.
		vRevertSecondFactorAdditionTemplate = &ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate{}
	}
	if err := extractConfigNotificationSendEmailRevertSecondFactorAdditionTemplateFields(r, vRevertSecondFactorAdditionTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRevertSecondFactorAdditionTemplate) {
		o.RevertSecondFactorAdditionTemplate = vRevertSecondFactorAdditionTemplate
	}
	return nil
}
func postReadExtractConfigNotificationSendEmailSmtpFields(r *Config, o *ConfigNotificationSendEmailSmtp) error {
	return nil
}
func postReadExtractConfigNotificationSendEmailResetPasswordTemplateFields(r *Config, o *ConfigNotificationSendEmailResetPasswordTemplate) error {
	return nil
}
func postReadExtractConfigNotificationSendEmailVerifyEmailTemplateFields(r *Config, o *ConfigNotificationSendEmailVerifyEmailTemplate) error {
	return nil
}
func postReadExtractConfigNotificationSendEmailChangeEmailTemplateFields(r *Config, o *ConfigNotificationSendEmailChangeEmailTemplate) error {
	return nil
}
func postReadExtractConfigNotificationSendEmailDnsInfoFields(r *Config, o *ConfigNotificationSendEmailDnsInfo) error {
	return nil
}
func postReadExtractConfigNotificationSendEmailRevertSecondFactorAdditionTemplateFields(r *Config, o *ConfigNotificationSendEmailRevertSecondFactorAdditionTemplate) error {
	return nil
}
func postReadExtractConfigNotificationSendSmsFields(r *Config, o *ConfigNotificationSendSms) error {
	vSmsTemplate := o.SmsTemplate
	if vSmsTemplate == nil {
		// note: explicitly not the empty object.
		vSmsTemplate = &ConfigNotificationSendSmsSmsTemplate{}
	}
	if err := extractConfigNotificationSendSmsSmsTemplateFields(r, vSmsTemplate); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSmsTemplate) {
		o.SmsTemplate = vSmsTemplate
	}
	return nil
}
func postReadExtractConfigNotificationSendSmsSmsTemplateFields(r *Config, o *ConfigNotificationSendSmsSmsTemplate) error {
	return nil
}
func postReadExtractConfigQuotaFields(r *Config, o *ConfigQuota) error {
	vSignUpQuotaConfig := o.SignUpQuotaConfig
	if vSignUpQuotaConfig == nil {
		// note: explicitly not the empty object.
		vSignUpQuotaConfig = &ConfigQuotaSignUpQuotaConfig{}
	}
	if err := extractConfigQuotaSignUpQuotaConfigFields(r, vSignUpQuotaConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSignUpQuotaConfig) {
		o.SignUpQuotaConfig = vSignUpQuotaConfig
	}
	return nil
}
func postReadExtractConfigQuotaSignUpQuotaConfigFields(r *Config, o *ConfigQuotaSignUpQuotaConfig) error {
	return nil
}
func postReadExtractConfigMonitoringFields(r *Config, o *ConfigMonitoring) error {
	vRequestLogging := o.RequestLogging
	if vRequestLogging == nil {
		// note: explicitly not the empty object.
		vRequestLogging = &ConfigMonitoringRequestLogging{}
	}
	if err := extractConfigMonitoringRequestLoggingFields(r, vRequestLogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRequestLogging) {
		o.RequestLogging = vRequestLogging
	}
	return nil
}
func postReadExtractConfigMonitoringRequestLoggingFields(r *Config, o *ConfigMonitoringRequestLogging) error {
	return nil
}
func postReadExtractConfigMultiTenantFields(r *Config, o *ConfigMultiTenant) error {
	return nil
}
func postReadExtractConfigClientFields(r *Config, o *ConfigClient) error {
	vPermissions := o.Permissions
	if vPermissions == nil {
		// note: explicitly not the empty object.
		vPermissions = &ConfigClientPermissions{}
	}
	if err := extractConfigClientPermissionsFields(r, vPermissions); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPermissions) {
		o.Permissions = vPermissions
	}
	return nil
}
func postReadExtractConfigClientPermissionsFields(r *Config, o *ConfigClientPermissions) error {
	return nil
}
func postReadExtractConfigMfaFields(r *Config, o *ConfigMfa) error {
	return nil
}
func postReadExtractConfigBlockingFunctionsFields(r *Config, o *ConfigBlockingFunctions) error {
	return nil
}
func postReadExtractConfigBlockingFunctionsTriggersFields(r *Config, o *ConfigBlockingFunctionsTriggers) error {
	return nil
}
