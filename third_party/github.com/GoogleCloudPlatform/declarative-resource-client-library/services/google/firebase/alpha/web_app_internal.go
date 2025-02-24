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

func (r *WebApp) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	return nil
}
func (r *WebApp) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://firebase.googleapis.com/v1alpha/", params)
}

func (r *WebApp) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/webApps/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *WebApp) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/webApps", nr.basePath(), userBasePath, params), nil

}

func (r *WebApp) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/webApps", nr.basePath(), userBasePath, params), nil

}

func (r *WebApp) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/webApps/{{name}}:remove", nr.basePath(), userBasePath, params), nil
}

// webAppApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type webAppApiOperation interface {
	do(context.Context, *WebApp, *Client) error
}

// newUpdateWebAppUpdateWebAppRequest creates a request for an
// WebApp resource's UpdateWebApp update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateWebAppUpdateWebAppRequest(ctx context.Context, f *WebApp, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.AppUrls; v != nil {
		req["appUrls"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.ApiKeyId); err != nil {
		return nil, fmt.Errorf("error expanding ApiKeyId into apiKeyId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["apiKeyId"] = v
	}
	b, err := c.getWebAppRaw(ctx, f)
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

// marshalUpdateWebAppUpdateWebAppRequest converts the update into
// the final JSON request body.
func marshalUpdateWebAppUpdateWebAppRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateWebAppUpdateWebAppOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateWebAppUpdateWebAppOperation) do(ctx context.Context, r *WebApp, c *Client) error {
	_, err := c.GetWebApp(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateWebApp")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateWebAppUpdateWebAppRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateWebAppUpdateWebAppRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listWebAppRaw(ctx context.Context, r *WebApp, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != WebAppMaxPage {
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

type listWebAppOperation struct {
	Apps  []map[string]interface{} `json:"apps"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listWebApp(ctx context.Context, r *WebApp, pageToken string, pageSize int32) ([]*WebApp, string, error) {
	b, err := c.listWebAppRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listWebAppOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*WebApp
	for _, v := range m.Apps {
		res, err := unmarshalMapWebApp(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllWebApp(ctx context.Context, f func(*WebApp) bool, resources []*WebApp) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteWebApp(ctx, res)
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

type deleteWebAppOperation struct{}

func (op *deleteWebAppOperation) do(ctx context.Context, r *WebApp, c *Client) error {
	r, err := c.GetWebApp(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "WebApp not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetWebApp checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, body, c.Config.RetryProvider)
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
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createWebAppOperation struct {
	response map[string]interface{}
}

func (op *createWebAppOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createWebAppOperation) do(ctx context.Context, r *WebApp, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a WebApp with the wrong Name.
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

	if _, err := c.GetWebApp(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getWebAppRaw(ctx context.Context, r *WebApp) ([]byte, error) {

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

func (c *Client) webAppDiffsForRawDesired(ctx context.Context, rawDesired *WebApp, opts ...dcl.ApplyOption) (initial, desired *WebApp, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *WebApp
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*WebApp); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected WebApp, got %T", sh)
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
		desired, err := canonicalizeWebAppDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetWebApp(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a WebApp resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve WebApp resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that WebApp resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeWebAppDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for WebApp: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for WebApp: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractWebAppFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeWebAppInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for WebApp: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeWebAppDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for WebApp: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffWebApp(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeWebAppInitialState(rawInitial, rawDesired *WebApp) (*WebApp, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeWebAppDesiredState(rawDesired, rawInitial *WebApp, opts ...dcl.ApplyOption) (*WebApp, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &WebApp{}
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
	if dcl.StringArrayCanonicalize(rawDesired.AppUrls, rawInitial.AppUrls) {
		canonicalDesired.AppUrls = rawInitial.AppUrls
	} else {
		canonicalDesired.AppUrls = rawDesired.AppUrls
	}
	if dcl.IsZeroValue(rawDesired.ApiKeyId) || (dcl.IsEmptyValueIndirect(rawDesired.ApiKeyId) && dcl.IsEmptyValueIndirect(rawInitial.ApiKeyId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.ApiKeyId = rawInitial.ApiKeyId
	} else {
		canonicalDesired.ApiKeyId = rawDesired.ApiKeyId
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeWebAppNewState(c *Client, rawNew, rawDesired *WebApp) (*WebApp, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.AppId) && dcl.IsEmptyValueIndirect(rawDesired.AppId) {
		rawNew.AppId = rawDesired.AppId
	} else {
		if dcl.StringCanonicalize(rawDesired.AppId, rawNew.AppId) {
			rawNew.AppId = rawDesired.AppId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DisplayName) && dcl.IsEmptyValueIndirect(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectId) && dcl.IsEmptyValueIndirect(rawDesired.ProjectId) {
		rawNew.ProjectId = rawDesired.ProjectId
	} else {
		if dcl.StringCanonicalize(rawDesired.ProjectId, rawNew.ProjectId) {
			rawNew.ProjectId = rawDesired.ProjectId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.AppUrls) && dcl.IsEmptyValueIndirect(rawDesired.AppUrls) {
		rawNew.AppUrls = rawDesired.AppUrls
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.AppUrls, rawNew.AppUrls) {
			rawNew.AppUrls = rawDesired.AppUrls
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.WebId) && dcl.IsEmptyValueIndirect(rawDesired.WebId) {
		rawNew.WebId = rawDesired.WebId
	} else {
		if dcl.StringCanonicalize(rawDesired.WebId, rawNew.WebId) {
			rawNew.WebId = rawDesired.WebId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ApiKeyId) && dcl.IsEmptyValueIndirect(rawDesired.ApiKeyId) {
		rawNew.ApiKeyId = rawDesired.ApiKeyId
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffWebApp(c *Client, desired, actual *WebApp, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.AppId, actual.AppId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AppId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWebAppUpdateWebAppOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AppUrls, actual.AppUrls, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateWebAppUpdateWebAppOperation")}, fn.AddNest("AppUrls")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WebId, actual.WebId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("WebId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ApiKeyId, actual.ApiKeyId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateWebAppUpdateWebAppOperation")}, fn.AddNest("ApiKeyId")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *WebApp) urlNormalized() *WebApp {
	normalized := dcl.Copy(*r).(WebApp)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.AppId = dcl.SelfLinkToName(r.AppId)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.ProjectId = dcl.SelfLinkToName(r.ProjectId)
	normalized.WebId = dcl.SelfLinkToName(r.WebId)
	normalized.ApiKeyId = dcl.SelfLinkToName(r.ApiKeyId)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *WebApp) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateWebApp" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/webApps/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the WebApp resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *WebApp) marshal(c *Client) ([]byte, error) {
	m, err := expandWebApp(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling WebApp: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalWebApp decodes JSON responses into the WebApp resource schema.
func unmarshalWebApp(b []byte, c *Client, res *WebApp) (*WebApp, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapWebApp(m, c, res)
}

func unmarshalMapWebApp(m map[string]interface{}, c *Client, res *WebApp) (*WebApp, error) {

	flattened := flattenWebApp(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandWebApp expands WebApp into a JSON request object.
func expandWebApp(c *Client, f *WebApp) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.DeriveField("projects/%s/webApps/%s", f.Name, dcl.SelfLinkToName(f.Project), dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.AppUrls; v != nil {
		m["appUrls"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.ApiKeyId); err != nil {
		return nil, fmt.Errorf("error expanding ApiKeyId into apiKeyId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["apiKeyId"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenWebApp flattens WebApp from a JSON request object into the
// WebApp type.
func flattenWebApp(c *Client, i interface{}, res *WebApp) *WebApp {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &WebApp{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.AppId = dcl.FlattenString(m["appId"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.ProjectId = dcl.FlattenString(m["projectId"])
	resultRes.AppUrls = dcl.FlattenStringSlice(m["appUrls"])
	resultRes.WebId = dcl.FlattenString(m["webId"])
	resultRes.ApiKeyId = dcl.FlattenString(m["apiKeyId"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *WebApp) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalWebApp(b, c, r)
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

type webAppDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         webAppApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToWebAppDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]webAppDiff, error) {
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
	var diffs []webAppDiff
	// For each operation name, create a webAppDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := webAppDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToWebAppApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToWebAppApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (webAppApiOperation, error) {
	switch opName {

	case "updateWebAppUpdateWebAppOperation":
		return &updateWebAppUpdateWebAppOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractWebAppFields(r *WebApp) error {
	return nil
}

func postReadExtractWebAppFields(r *WebApp) error {
	return nil
}
