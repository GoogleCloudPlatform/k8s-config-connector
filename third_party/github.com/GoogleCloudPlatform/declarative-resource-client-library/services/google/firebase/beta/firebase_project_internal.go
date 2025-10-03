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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *FirebaseProject) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Resources) {
		if err := r.Resources.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FirebaseProjectResources) validate() error {
	return nil
}
func (r *FirebaseProject) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://firebase.googleapis.com/v1beta1/", params)
}

func (r *FirebaseProject) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}", nr.basePath(), userBasePath, params), nil
}

func (r *FirebaseProject) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{}
	return dcl.URL("projects", nr.basePath(), userBasePath, params), nil

}

func (r *FirebaseProject) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}:addFirebase", nr.basePath(), userBasePath, params), nil

}

// firebaseProjectApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type firebaseProjectApiOperation interface {
	do(context.Context, *FirebaseProject, *Client) error
}

// newUpdateFirebaseProjectUpdateFirebaseProjectRequest creates a request for an
// FirebaseProject resource's UpdateFirebaseProject update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFirebaseProjectUpdateFirebaseProjectRequest(ctx context.Context, f *FirebaseProject, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Annotations; !dcl.IsEmptyValueIndirect(v) {
		req["annotations"] = v
	}
	b, err := c.getFirebaseProjectRaw(ctx, f)
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

// marshalUpdateFirebaseProjectUpdateFirebaseProjectRequest converts the update into
// the final JSON request body.
func marshalUpdateFirebaseProjectUpdateFirebaseProjectRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFirebaseProjectUpdateFirebaseProjectOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateFirebaseProjectUpdateFirebaseProjectOperation) do(ctx context.Context, r *FirebaseProject, c *Client) error {
	_, err := c.GetFirebaseProject(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateFirebaseProject")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateFirebaseProjectUpdateFirebaseProjectRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateFirebaseProjectUpdateFirebaseProjectRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listFirebaseProjectRaw(ctx context.Context, r *FirebaseProject, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != FirebaseProjectMaxPage {
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

type listFirebaseProjectOperation struct {
	Results []map[string]interface{} `json:"results"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listFirebaseProject(ctx context.Context, r *FirebaseProject, pageToken string, pageSize int32) ([]*FirebaseProject, string, error) {
	b, err := c.listFirebaseProjectRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listFirebaseProjectOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*FirebaseProject
	for _, v := range m.Results {
		res, err := unmarshalMapFirebaseProject(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFirebaseProjectOperation struct {
	response map[string]interface{}
}

func (op *createFirebaseProjectOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createFirebaseProjectOperation) do(ctx context.Context, r *FirebaseProject, c *Client) error {
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

	if _, err := c.GetFirebaseProject(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getFirebaseProjectRaw(ctx context.Context, r *FirebaseProject) ([]byte, error) {

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

func (c *Client) firebaseProjectDiffsForRawDesired(ctx context.Context, rawDesired *FirebaseProject, opts ...dcl.ApplyOption) (initial, desired *FirebaseProject, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FirebaseProject
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FirebaseProject); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected FirebaseProject, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFirebaseProject(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a FirebaseProject resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FirebaseProject resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that FirebaseProject resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFirebaseProjectDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for FirebaseProject: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for FirebaseProject: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFirebaseProjectFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFirebaseProjectInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for FirebaseProject: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFirebaseProjectDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for FirebaseProject: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFirebaseProject(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFirebaseProjectInitialState(rawInitial, rawDesired *FirebaseProject) (*FirebaseProject, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFirebaseProjectDesiredState(rawDesired, rawInitial *FirebaseProject, opts ...dcl.ApplyOption) (*FirebaseProject, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Resources = canonicalizeFirebaseProjectResources(rawDesired.Resources, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &FirebaseProject{}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.IsZeroValue(rawDesired.Annotations) || (dcl.IsEmptyValueIndirect(rawDesired.Annotations) && dcl.IsEmptyValueIndirect(rawInitial.Annotations)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Annotations = rawInitial.Annotations
	} else {
		canonicalDesired.Annotations = rawDesired.Annotations
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	return canonicalDesired, nil
}

func canonicalizeFirebaseProjectNewState(c *Client, rawNew, rawDesired *FirebaseProject) (*FirebaseProject, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ProjectId) && dcl.IsEmptyValueIndirect(rawDesired.ProjectId) {
		rawNew.ProjectId = rawDesired.ProjectId
	} else {
		if dcl.StringCanonicalize(rawDesired.ProjectId, rawNew.ProjectId) {
			rawNew.ProjectId = rawDesired.ProjectId
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.ProjectNumber) && dcl.IsEmptyValueIndirect(rawDesired.ProjectNumber) {
		rawNew.ProjectNumber = rawDesired.ProjectNumber
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
		rawNew.Resources = canonicalizeNewFirebaseProjectResources(c, rawDesired.Resources, rawNew.Resources)
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Annotations) && dcl.IsEmptyValueIndirect(rawDesired.Annotations) {
		rawNew.Annotations = rawDesired.Annotations
	} else {
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

func canonicalizeFirebaseProjectResources(des, initial *FirebaseProjectResources, opts ...dcl.ApplyOption) *FirebaseProjectResources {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FirebaseProjectResources{}

	return cDes
}

func canonicalizeFirebaseProjectResourcesSlice(des, initial []FirebaseProjectResources, opts ...dcl.ApplyOption) []FirebaseProjectResources {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FirebaseProjectResources, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFirebaseProjectResources(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FirebaseProjectResources, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFirebaseProjectResources(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFirebaseProjectResources(c *Client, des, nw *FirebaseProjectResources) *FirebaseProjectResources {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FirebaseProjectResources while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.HostingSite, nw.HostingSite) {
		nw.HostingSite = des.HostingSite
	}
	if dcl.StringCanonicalize(des.RealtimeDatabaseInstance, nw.RealtimeDatabaseInstance) {
		nw.RealtimeDatabaseInstance = des.RealtimeDatabaseInstance
	}
	if dcl.StringCanonicalize(des.StorageBucket, nw.StorageBucket) {
		nw.StorageBucket = des.StorageBucket
	}
	if dcl.StringCanonicalize(des.LocationId, nw.LocationId) {
		nw.LocationId = des.LocationId
	}

	return nw
}

func canonicalizeNewFirebaseProjectResourcesSet(c *Client, des, nw []FirebaseProjectResources) []FirebaseProjectResources {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FirebaseProjectResources
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFirebaseProjectResourcesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFirebaseProjectResources(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFirebaseProjectResourcesSlice(c *Client, des, nw []FirebaseProjectResources) []FirebaseProjectResources {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FirebaseProjectResources
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFirebaseProjectResources(c, &d, &n))
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
func diffFirebaseProject(c *Client, desired, actual *FirebaseProject, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectNumber, actual.ProjectNumber, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectNumber")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFirebaseProjectUpdateFirebaseProjectOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Resources, actual.Resources, dcl.DiffInfo{OutputOnly: true, ObjectFunction: compareFirebaseProjectResourcesNewStyle, EmptyObject: EmptyFirebaseProjectResources, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Resources")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Annotations, actual.Annotations, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFirebaseProjectUpdateFirebaseProjectOperation")}, fn.AddNest("Annotations")); len(ds) != 0 || err != nil {
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
func compareFirebaseProjectResourcesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FirebaseProjectResources)
	if !ok {
		desiredNotPointer, ok := d.(FirebaseProjectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirebaseProjectResources or *FirebaseProjectResources", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FirebaseProjectResources)
	if !ok {
		actualNotPointer, ok := a.(FirebaseProjectResources)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FirebaseProjectResources", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.HostingSite, actual.HostingSite, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HostingSite")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RealtimeDatabaseInstance, actual.RealtimeDatabaseInstance, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RealtimeDatabaseInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageBucket, actual.StorageBucket, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LocationId, actual.LocationId, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LocationId")); len(ds) != 0 || err != nil {
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
func (r *FirebaseProject) urlNormalized() *FirebaseProject {
	normalized := dcl.Copy(*r).(FirebaseProject)
	normalized.ProjectId = dcl.SelfLinkToName(r.ProjectId)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *FirebaseProject) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateFirebaseProject" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
		}
		return dcl.URL("projects/{{project}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FirebaseProject resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FirebaseProject) marshal(c *Client) ([]byte, error) {
	m, err := expandFirebaseProject(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FirebaseProject: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFirebaseProject decodes JSON responses into the FirebaseProject resource schema.
func unmarshalFirebaseProject(b []byte, c *Client, res *FirebaseProject) (*FirebaseProject, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFirebaseProject(m, c, res)
}

func unmarshalMapFirebaseProject(m map[string]interface{}, c *Client, res *FirebaseProject) (*FirebaseProject, error) {

	flattened := flattenFirebaseProject(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFirebaseProject expands FirebaseProject into a JSON request object.
func expandFirebaseProject(c *Client, f *FirebaseProject) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Annotations; dcl.ValueShouldBeSent(v) {
		m["annotations"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}

	return m, nil
}

// flattenFirebaseProject flattens FirebaseProject from a JSON request object into the
// FirebaseProject type.
func flattenFirebaseProject(c *Client, i interface{}, res *FirebaseProject) *FirebaseProject {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &FirebaseProject{}
	resultRes.ProjectId = dcl.FlattenString(m["projectId"])
	resultRes.ProjectNumber = dcl.FlattenInteger(m["projectNumber"])
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Resources = flattenFirebaseProjectResources(c, m["resources"], res)
	resultRes.State = flattenFirebaseProjectStateEnum(m["state"])
	resultRes.Annotations = dcl.FlattenKeyValuePairs(m["annotations"])
	resultRes.Project = dcl.FlattenString(m["project"])

	return resultRes
}

// expandFirebaseProjectResourcesMap expands the contents of FirebaseProjectResources into a JSON
// request object.
func expandFirebaseProjectResourcesMap(c *Client, f map[string]FirebaseProjectResources, res *FirebaseProject) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFirebaseProjectResources(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFirebaseProjectResourcesSlice expands the contents of FirebaseProjectResources into a JSON
// request object.
func expandFirebaseProjectResourcesSlice(c *Client, f []FirebaseProjectResources, res *FirebaseProject) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFirebaseProjectResources(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFirebaseProjectResourcesMap flattens the contents of FirebaseProjectResources from a JSON
// response object.
func flattenFirebaseProjectResourcesMap(c *Client, i interface{}, res *FirebaseProject) map[string]FirebaseProjectResources {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirebaseProjectResources{}
	}

	if len(a) == 0 {
		return map[string]FirebaseProjectResources{}
	}

	items := make(map[string]FirebaseProjectResources)
	for k, item := range a {
		items[k] = *flattenFirebaseProjectResources(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFirebaseProjectResourcesSlice flattens the contents of FirebaseProjectResources from a JSON
// response object.
func flattenFirebaseProjectResourcesSlice(c *Client, i interface{}, res *FirebaseProject) []FirebaseProjectResources {
	a, ok := i.([]interface{})
	if !ok {
		return []FirebaseProjectResources{}
	}

	if len(a) == 0 {
		return []FirebaseProjectResources{}
	}

	items := make([]FirebaseProjectResources, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirebaseProjectResources(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFirebaseProjectResources expands an instance of FirebaseProjectResources into a JSON
// request object.
func expandFirebaseProjectResources(c *Client, f *FirebaseProjectResources, res *FirebaseProject) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})

	return m, nil
}

// flattenFirebaseProjectResources flattens an instance of FirebaseProjectResources from a JSON
// response object.
func flattenFirebaseProjectResources(c *Client, i interface{}, res *FirebaseProject) *FirebaseProjectResources {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FirebaseProjectResources{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFirebaseProjectResources
	}
	r.HostingSite = dcl.FlattenString(m["hostingSite"])
	r.RealtimeDatabaseInstance = dcl.FlattenString(m["realtimeDatabaseInstance"])
	r.StorageBucket = dcl.FlattenString(m["storageBucket"])
	r.LocationId = dcl.FlattenString(m["locationId"])

	return r
}

// flattenFirebaseProjectStateEnumMap flattens the contents of FirebaseProjectStateEnum from a JSON
// response object.
func flattenFirebaseProjectStateEnumMap(c *Client, i interface{}, res *FirebaseProject) map[string]FirebaseProjectStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FirebaseProjectStateEnum{}
	}

	if len(a) == 0 {
		return map[string]FirebaseProjectStateEnum{}
	}

	items := make(map[string]FirebaseProjectStateEnum)
	for k, item := range a {
		items[k] = *flattenFirebaseProjectStateEnum(item.(interface{}))
	}

	return items
}

// flattenFirebaseProjectStateEnumSlice flattens the contents of FirebaseProjectStateEnum from a JSON
// response object.
func flattenFirebaseProjectStateEnumSlice(c *Client, i interface{}, res *FirebaseProject) []FirebaseProjectStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FirebaseProjectStateEnum{}
	}

	if len(a) == 0 {
		return []FirebaseProjectStateEnum{}
	}

	items := make([]FirebaseProjectStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFirebaseProjectStateEnum(item.(interface{})))
	}

	return items
}

// flattenFirebaseProjectStateEnum asserts that an interface is a string, and returns a
// pointer to a *FirebaseProjectStateEnum with the same value as that string.
func flattenFirebaseProjectStateEnum(i interface{}) *FirebaseProjectStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FirebaseProjectStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FirebaseProject) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFirebaseProject(b, c, r)
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

type firebaseProjectDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         firebaseProjectApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFirebaseProjectDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]firebaseProjectDiff, error) {
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
	var diffs []firebaseProjectDiff
	// For each operation name, create a firebaseProjectDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := firebaseProjectDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFirebaseProjectApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFirebaseProjectApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (firebaseProjectApiOperation, error) {
	switch opName {

	case "updateFirebaseProjectUpdateFirebaseProjectOperation":
		return &updateFirebaseProjectUpdateFirebaseProjectOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFirebaseProjectFields(r *FirebaseProject) error {
	vResources := r.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &FirebaseProjectResources{}
	}
	if err := extractFirebaseProjectResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		r.Resources = vResources
	}
	return nil
}
func extractFirebaseProjectResourcesFields(r *FirebaseProject, o *FirebaseProjectResources) error {
	return nil
}

func postReadExtractFirebaseProjectFields(r *FirebaseProject) error {
	vResources := r.Resources
	if vResources == nil {
		// note: explicitly not the empty object.
		vResources = &FirebaseProjectResources{}
	}
	if err := postReadExtractFirebaseProjectResourcesFields(r, vResources); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vResources) {
		r.Resources = vResources
	}
	return nil
}
func postReadExtractFirebaseProjectResourcesFields(r *FirebaseProject, o *FirebaseProjectResources) error {
	return nil
}
