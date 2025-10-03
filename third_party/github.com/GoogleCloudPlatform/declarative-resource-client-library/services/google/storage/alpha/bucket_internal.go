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

func (r *Bucket) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.Required(r, "location"); err != nil {
		return err
	}
	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Lifecycle) {
		if err := r.Lifecycle.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Logging) {
		if err := r.Logging.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Versioning) {
		if err := r.Versioning.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Website) {
		if err := r.Website.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BucketCors) validate() error {
	return nil
}
func (r *BucketLifecycle) validate() error {
	return nil
}
func (r *BucketLifecycleRule) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Action) {
		if err := r.Action.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Condition) {
		if err := r.Condition.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *BucketLifecycleRuleAction) validate() error {
	return nil
}
func (r *BucketLifecycleRuleCondition) validate() error {
	return nil
}
func (r *BucketLogging) validate() error {
	return nil
}
func (r *BucketVersioning) validate() error {
	return nil
}
func (r *BucketWebsite) validate() error {
	return nil
}
func (r *Bucket) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://www.googleapis.com/storage/v1/", params)
}

func (r *Bucket) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("b/{{name}}?userProject={{project}}", nr.basePath(), userBasePath, params), nil
}

func (r *Bucket) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("b?project={{project}}", nr.basePath(), userBasePath, params), nil

}

func (r *Bucket) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("b?project={{project}}", nr.basePath(), userBasePath, params), nil

}

func (r *Bucket) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("b/{{name}}?userProject={{project}}", nr.basePath(), userBasePath, params), nil
}

func (r *Bucket) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *nr.Name,
	}
	return dcl.URL("b/{{name}}/iam", nr.basePath(), userBasePath, fields)
}

func (r *Bucket) SetPolicyVerb() string {
	return "PUT"
}

func (r *Bucket) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"name": *nr.Name,
	}
	return dcl.URL("b/{{name}}/iam", nr.basePath(), userBasePath, fields)
}

func (r *Bucket) IAMPolicyVersion() int {
	return 3
}

// bucketApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type bucketApiOperation interface {
	do(context.Context, *Bucket, *Client) error
}

// newUpdateBucketUpdateRequest creates a request for an
// Bucket resource's update update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBucketUpdateRequest(ctx context.Context, f *Bucket, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandBucketCorsSlice(c, f.Cors, res); err != nil {
		return nil, fmt.Errorf("error expanding Cors into cors: %w", err)
	} else if v != nil {
		req["cors"] = v
	}
	if v, err := expandBucketLifecycle(c, f.Lifecycle, res); err != nil {
		return nil, fmt.Errorf("error expanding Lifecycle into lifecycle: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["lifecycle"] = v
	}
	if v, err := expandBucketLogging(c, f.Logging, res); err != nil {
		return nil, fmt.Errorf("error expanding Logging into logging: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["logging"] = v
	}
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		req["storageClass"] = v
	}
	if v, err := expandBucketVersioning(c, f.Versioning, res); err != nil {
		return nil, fmt.Errorf("error expanding Versioning into versioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["versioning"] = v
	}
	if v, err := expandBucketWebsite(c, f.Website, res); err != nil {
		return nil, fmt.Errorf("error expanding Website into website: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["website"] = v
	}
	return req, nil
}

// marshalUpdateBucketUpdateRequest converts the update into
// the final JSON request body.
func marshalUpdateBucketUpdateRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBucketUpdateOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBucketUpdateOperation) do(ctx context.Context, r *Bucket, c *Client) error {
	_, err := c.GetBucket(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "update")
	if err != nil {
		return err
	}

	req, err := newUpdateBucketUpdateRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateBucketUpdateRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBucketRaw(ctx context.Context, r *Bucket, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BucketMaxPage {
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

type listBucketOperation struct {
	Items []map[string]interface{} `json:"items"`
	Token string                   `json:"nextPageToken"`
}

func (c *Client) listBucket(ctx context.Context, r *Bucket, pageToken string, pageSize int32) ([]*Bucket, string, error) {
	b, err := c.listBucketRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBucketOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Bucket
	for _, v := range m.Items {
		res, err := unmarshalMapBucket(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBucket(ctx context.Context, f func(*Bucket) bool, resources []*Bucket) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBucket(ctx, res)
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

type deleteBucketOperation struct{}

func (op *deleteBucketOperation) do(ctx context.Context, r *Bucket, c *Client) error {
	r, err := c.GetBucket(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Bucket not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetBucket checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Bucket: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetBucket(ctx, r)
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
type createBucketOperation struct {
	response map[string]interface{}
}

func (op *createBucketOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBucketOperation) do(ctx context.Context, r *Bucket, c *Client) error {
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

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	if _, err := c.GetBucket(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBucketRaw(ctx context.Context, r *Bucket) ([]byte, error) {

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

func (c *Client) bucketDiffsForRawDesired(ctx context.Context, rawDesired *Bucket, opts ...dcl.ApplyOption) (initial, desired *Bucket, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Bucket
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Bucket); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Bucket, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBucket(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Bucket resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Bucket resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Bucket resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBucketDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Bucket: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Bucket: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractBucketFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBucketInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Bucket: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBucketDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Bucket: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBucket(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBucketInitialState(rawInitial, rawDesired *Bucket) (*Bucket, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBucketDesiredState(rawDesired, rawInitial *Bucket, opts ...dcl.ApplyOption) (*Bucket, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Lifecycle = canonicalizeBucketLifecycle(rawDesired.Lifecycle, nil, opts...)
		rawDesired.Logging = canonicalizeBucketLogging(rawDesired.Logging, nil, opts...)
		rawDesired.Versioning = canonicalizeBucketVersioning(rawDesired.Versioning, nil, opts...)
		rawDesired.Website = canonicalizeBucketWebsite(rawDesired.Website, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Bucket{}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.Cors = canonicalizeBucketCorsSlice(rawDesired.Cors, rawInitial.Cors, opts...)
	canonicalDesired.Lifecycle = canonicalizeBucketLifecycle(rawDesired.Lifecycle, rawInitial.Lifecycle, opts...)
	canonicalDesired.Logging = canonicalizeBucketLogging(rawDesired.Logging, rawInitial.Logging, opts...)
	if dcl.IsZeroValue(rawDesired.StorageClass) || (dcl.IsEmptyValueIndirect(rawDesired.StorageClass) && dcl.IsEmptyValueIndirect(rawInitial.StorageClass)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.StorageClass = rawInitial.StorageClass
	} else {
		canonicalDesired.StorageClass = rawDesired.StorageClass
	}
	canonicalDesired.Versioning = canonicalizeBucketVersioning(rawDesired.Versioning, rawInitial.Versioning, opts...)
	canonicalDesired.Website = canonicalizeBucketWebsite(rawDesired.Website, rawInitial.Website, opts...)
	return canonicalDesired, nil
}

func canonicalizeBucketNewState(c *Client, rawNew, rawDesired *Bucket) (*Bucket, error) {

	rawNew.Project = rawDesired.Project

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Cors) && dcl.IsEmptyValueIndirect(rawDesired.Cors) {
		rawNew.Cors = rawDesired.Cors
	} else {
		rawNew.Cors = canonicalizeNewBucketCorsSlice(c, rawDesired.Cors, rawNew.Cors)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Lifecycle) && dcl.IsEmptyValueIndirect(rawDesired.Lifecycle) {
		rawNew.Lifecycle = rawDesired.Lifecycle
	} else {
		rawNew.Lifecycle = canonicalizeNewBucketLifecycle(c, rawDesired.Lifecycle, rawNew.Lifecycle)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Logging) && dcl.IsEmptyValueIndirect(rawDesired.Logging) {
		rawNew.Logging = rawDesired.Logging
	} else {
		rawNew.Logging = canonicalizeNewBucketLogging(c, rawDesired.Logging, rawNew.Logging)
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageClass) && dcl.IsEmptyValueIndirect(rawDesired.StorageClass) {
		rawNew.StorageClass = rawDesired.StorageClass
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Versioning) && dcl.IsEmptyValueIndirect(rawDesired.Versioning) {
		rawNew.Versioning = rawDesired.Versioning
	} else {
		rawNew.Versioning = canonicalizeNewBucketVersioning(c, rawDesired.Versioning, rawNew.Versioning)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Website) && dcl.IsEmptyValueIndirect(rawDesired.Website) {
		rawNew.Website = rawDesired.Website
	} else {
		rawNew.Website = canonicalizeNewBucketWebsite(c, rawDesired.Website, rawNew.Website)
	}

	return rawNew, nil
}

func canonicalizeBucketCors(des, initial *BucketCors, opts ...dcl.ApplyOption) *BucketCors {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketCors{}

	if dcl.IsZeroValue(des.MaxAgeSeconds) || (dcl.IsEmptyValueIndirect(des.MaxAgeSeconds) && dcl.IsEmptyValueIndirect(initial.MaxAgeSeconds)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MaxAgeSeconds = initial.MaxAgeSeconds
	} else {
		cDes.MaxAgeSeconds = des.MaxAgeSeconds
	}
	if dcl.StringArrayCanonicalize(des.Method, initial.Method) {
		cDes.Method = initial.Method
	} else {
		cDes.Method = des.Method
	}
	if dcl.StringArrayCanonicalize(des.Origin, initial.Origin) {
		cDes.Origin = initial.Origin
	} else {
		cDes.Origin = des.Origin
	}
	if dcl.StringArrayCanonicalize(des.ResponseHeader, initial.ResponseHeader) {
		cDes.ResponseHeader = initial.ResponseHeader
	} else {
		cDes.ResponseHeader = des.ResponseHeader
	}

	return cDes
}

func canonicalizeBucketCorsSlice(des, initial []BucketCors, opts ...dcl.ApplyOption) []BucketCors {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketCors, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketCors(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketCors, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketCors(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketCors(c *Client, des, nw *BucketCors) *BucketCors {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketCors while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Method, nw.Method) {
		nw.Method = des.Method
	}
	if dcl.StringArrayCanonicalize(des.Origin, nw.Origin) {
		nw.Origin = des.Origin
	}
	if dcl.StringArrayCanonicalize(des.ResponseHeader, nw.ResponseHeader) {
		nw.ResponseHeader = des.ResponseHeader
	}

	return nw
}

func canonicalizeNewBucketCorsSet(c *Client, des, nw []BucketCors) []BucketCors {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketCors
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketCorsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketCors(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketCorsSlice(c *Client, des, nw []BucketCors) []BucketCors {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketCors
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketCors(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycle(des, initial *BucketLifecycle, opts ...dcl.ApplyOption) *BucketLifecycle {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketLifecycle{}

	cDes.Rule = canonicalizeBucketLifecycleRuleSlice(des.Rule, initial.Rule, opts...)

	return cDes
}

func canonicalizeBucketLifecycleSlice(des, initial []BucketLifecycle, opts ...dcl.ApplyOption) []BucketLifecycle {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketLifecycle, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketLifecycle(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketLifecycle, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketLifecycle(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketLifecycle(c *Client, des, nw *BucketLifecycle) *BucketLifecycle {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketLifecycle while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Rule = canonicalizeNewBucketLifecycleRuleSlice(c, des.Rule, nw.Rule)

	return nw
}

func canonicalizeNewBucketLifecycleSet(c *Client, des, nw []BucketLifecycle) []BucketLifecycle {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketLifecycle
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketLifecycleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketLifecycle(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketLifecycleSlice(c *Client, des, nw []BucketLifecycle) []BucketLifecycle {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycle
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycle(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRule(des, initial *BucketLifecycleRule, opts ...dcl.ApplyOption) *BucketLifecycleRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketLifecycleRule{}

	cDes.Action = canonicalizeBucketLifecycleRuleAction(des.Action, initial.Action, opts...)
	cDes.Condition = canonicalizeBucketLifecycleRuleCondition(des.Condition, initial.Condition, opts...)

	return cDes
}

func canonicalizeBucketLifecycleRuleSlice(des, initial []BucketLifecycleRule, opts ...dcl.ApplyOption) []BucketLifecycleRule {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketLifecycleRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketLifecycleRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketLifecycleRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketLifecycleRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketLifecycleRule(c *Client, des, nw *BucketLifecycleRule) *BucketLifecycleRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketLifecycleRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Action = canonicalizeNewBucketLifecycleRuleAction(c, des.Action, nw.Action)
	nw.Condition = canonicalizeNewBucketLifecycleRuleCondition(c, des.Condition, nw.Condition)

	return nw
}

func canonicalizeNewBucketLifecycleRuleSet(c *Client, des, nw []BucketLifecycleRule) []BucketLifecycleRule {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketLifecycleRule
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketLifecycleRule(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketLifecycleRuleSlice(c *Client, des, nw []BucketLifecycleRule) []BucketLifecycleRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRule(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRuleAction(des, initial *BucketLifecycleRuleAction, opts ...dcl.ApplyOption) *BucketLifecycleRuleAction {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketLifecycleRuleAction{}

	if dcl.StringCanonicalize(des.StorageClass, initial.StorageClass) || dcl.IsZeroValue(des.StorageClass) {
		cDes.StorageClass = initial.StorageClass
	} else {
		cDes.StorageClass = des.StorageClass
	}
	if dcl.IsZeroValue(des.Type) || (dcl.IsEmptyValueIndirect(des.Type) && dcl.IsEmptyValueIndirect(initial.Type)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Type = initial.Type
	} else {
		cDes.Type = des.Type
	}

	return cDes
}

func canonicalizeBucketLifecycleRuleActionSlice(des, initial []BucketLifecycleRuleAction, opts ...dcl.ApplyOption) []BucketLifecycleRuleAction {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketLifecycleRuleAction, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketLifecycleRuleAction(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketLifecycleRuleAction, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketLifecycleRuleAction(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketLifecycleRuleAction(c *Client, des, nw *BucketLifecycleRuleAction) *BucketLifecycleRuleAction {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketLifecycleRuleAction while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.StorageClass, nw.StorageClass) {
		nw.StorageClass = des.StorageClass
	}

	return nw
}

func canonicalizeNewBucketLifecycleRuleActionSet(c *Client, des, nw []BucketLifecycleRuleAction) []BucketLifecycleRuleAction {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketLifecycleRuleAction
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleActionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketLifecycleRuleAction(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketLifecycleRuleActionSlice(c *Client, des, nw []BucketLifecycleRuleAction) []BucketLifecycleRuleAction {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRuleAction
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRuleAction(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLifecycleRuleCondition(des, initial *BucketLifecycleRuleCondition, opts ...dcl.ApplyOption) *BucketLifecycleRuleCondition {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketLifecycleRuleCondition{}

	if dcl.IsZeroValue(des.Age) || (dcl.IsEmptyValueIndirect(des.Age) && dcl.IsEmptyValueIndirect(initial.Age)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Age = initial.Age
	} else {
		cDes.Age = des.Age
	}
	if dcl.IsZeroValue(des.CreatedBefore) || (dcl.IsEmptyValueIndirect(des.CreatedBefore) && dcl.IsEmptyValueIndirect(initial.CreatedBefore)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.CreatedBefore = initial.CreatedBefore
	} else {
		cDes.CreatedBefore = des.CreatedBefore
	}
	if dcl.IsZeroValue(des.WithState) || (dcl.IsEmptyValueIndirect(des.WithState) && dcl.IsEmptyValueIndirect(initial.WithState)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.WithState = initial.WithState
	} else {
		cDes.WithState = des.WithState
	}
	if dcl.StringArrayCanonicalize(des.MatchesStorageClass, initial.MatchesStorageClass) {
		cDes.MatchesStorageClass = initial.MatchesStorageClass
	} else {
		cDes.MatchesStorageClass = des.MatchesStorageClass
	}
	if dcl.IsZeroValue(des.NumNewerVersions) || (dcl.IsEmptyValueIndirect(des.NumNewerVersions) && dcl.IsEmptyValueIndirect(initial.NumNewerVersions)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.NumNewerVersions = initial.NumNewerVersions
	} else {
		cDes.NumNewerVersions = des.NumNewerVersions
	}

	return cDes
}

func canonicalizeBucketLifecycleRuleConditionSlice(des, initial []BucketLifecycleRuleCondition, opts ...dcl.ApplyOption) []BucketLifecycleRuleCondition {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketLifecycleRuleCondition, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketLifecycleRuleCondition(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketLifecycleRuleCondition, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketLifecycleRuleCondition(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketLifecycleRuleCondition(c *Client, des, nw *BucketLifecycleRuleCondition) *BucketLifecycleRuleCondition {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketLifecycleRuleCondition while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.MatchesStorageClass, nw.MatchesStorageClass) {
		nw.MatchesStorageClass = des.MatchesStorageClass
	}

	return nw
}

func canonicalizeNewBucketLifecycleRuleConditionSet(c *Client, des, nw []BucketLifecycleRuleCondition) []BucketLifecycleRuleCondition {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketLifecycleRuleCondition
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketLifecycleRuleConditionNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketLifecycleRuleCondition(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketLifecycleRuleConditionSlice(c *Client, des, nw []BucketLifecycleRuleCondition) []BucketLifecycleRuleCondition {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLifecycleRuleCondition
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLifecycleRuleCondition(c, &d, &n))
	}

	return items
}

func canonicalizeBucketLogging(des, initial *BucketLogging, opts ...dcl.ApplyOption) *BucketLogging {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketLogging{}

	if dcl.StringCanonicalize(des.LogBucket, initial.LogBucket) || dcl.IsZeroValue(des.LogBucket) {
		cDes.LogBucket = initial.LogBucket
	} else {
		cDes.LogBucket = des.LogBucket
	}
	if dcl.StringCanonicalize(des.LogObjectPrefix, initial.LogObjectPrefix) || dcl.IsZeroValue(des.LogObjectPrefix) {
		cDes.LogObjectPrefix = initial.LogObjectPrefix
	} else {
		cDes.LogObjectPrefix = des.LogObjectPrefix
	}

	return cDes
}

func canonicalizeBucketLoggingSlice(des, initial []BucketLogging, opts ...dcl.ApplyOption) []BucketLogging {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketLogging, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketLogging(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketLogging, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketLogging(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketLogging(c *Client, des, nw *BucketLogging) *BucketLogging {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketLogging while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.LogBucket, nw.LogBucket) {
		nw.LogBucket = des.LogBucket
	}
	if dcl.StringCanonicalize(des.LogObjectPrefix, nw.LogObjectPrefix) {
		nw.LogObjectPrefix = des.LogObjectPrefix
	}

	return nw
}

func canonicalizeNewBucketLoggingSet(c *Client, des, nw []BucketLogging) []BucketLogging {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketLogging
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketLoggingNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketLogging(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketLoggingSlice(c *Client, des, nw []BucketLogging) []BucketLogging {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketLogging
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketLogging(c, &d, &n))
	}

	return items
}

func canonicalizeBucketVersioning(des, initial *BucketVersioning, opts ...dcl.ApplyOption) *BucketVersioning {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketVersioning{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeBucketVersioningSlice(des, initial []BucketVersioning, opts ...dcl.ApplyOption) []BucketVersioning {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketVersioning, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketVersioning(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketVersioning, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketVersioning(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketVersioning(c *Client, des, nw *BucketVersioning) *BucketVersioning {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketVersioning while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewBucketVersioningSet(c *Client, des, nw []BucketVersioning) []BucketVersioning {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketVersioning
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketVersioningNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketVersioning(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketVersioningSlice(c *Client, des, nw []BucketVersioning) []BucketVersioning {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketVersioning
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketVersioning(c, &d, &n))
	}

	return items
}

func canonicalizeBucketWebsite(des, initial *BucketWebsite, opts ...dcl.ApplyOption) *BucketWebsite {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &BucketWebsite{}

	if dcl.StringCanonicalize(des.MainPageSuffix, initial.MainPageSuffix) || dcl.IsZeroValue(des.MainPageSuffix) {
		cDes.MainPageSuffix = initial.MainPageSuffix
	} else {
		cDes.MainPageSuffix = des.MainPageSuffix
	}
	if dcl.StringCanonicalize(des.NotFoundPage, initial.NotFoundPage) || dcl.IsZeroValue(des.NotFoundPage) {
		cDes.NotFoundPage = initial.NotFoundPage
	} else {
		cDes.NotFoundPage = des.NotFoundPage
	}

	return cDes
}

func canonicalizeBucketWebsiteSlice(des, initial []BucketWebsite, opts ...dcl.ApplyOption) []BucketWebsite {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]BucketWebsite, 0, len(des))
		for _, d := range des {
			cd := canonicalizeBucketWebsite(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]BucketWebsite, 0, len(des))
	for i, d := range des {
		cd := canonicalizeBucketWebsite(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewBucketWebsite(c *Client, des, nw *BucketWebsite) *BucketWebsite {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for BucketWebsite while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.MainPageSuffix, nw.MainPageSuffix) {
		nw.MainPageSuffix = des.MainPageSuffix
	}
	if dcl.StringCanonicalize(des.NotFoundPage, nw.NotFoundPage) {
		nw.NotFoundPage = des.NotFoundPage
	}

	return nw
}

func canonicalizeNewBucketWebsiteSet(c *Client, des, nw []BucketWebsite) []BucketWebsite {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []BucketWebsite
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareBucketWebsiteNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewBucketWebsite(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewBucketWebsiteSlice(c *Client, des, nw []BucketWebsite) []BucketWebsite {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []BucketWebsite
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewBucketWebsite(c, &d, &n))
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
func diffBucket(c *Client, desired, actual *Bucket, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
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

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Cors, actual.Cors, dcl.DiffInfo{ObjectFunction: compareBucketCorsNewStyle, EmptyObject: EmptyBucketCors, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Cors")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Lifecycle, actual.Lifecycle, dcl.DiffInfo{ObjectFunction: compareBucketLifecycleNewStyle, EmptyObject: EmptyBucketLifecycle, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Lifecycle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Logging, actual.Logging, dcl.DiffInfo{ObjectFunction: compareBucketLoggingNewStyle, EmptyObject: EmptyBucketLogging, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Logging")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageClass, actual.StorageClass, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("StorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Versioning, actual.Versioning, dcl.DiffInfo{ObjectFunction: compareBucketVersioningNewStyle, EmptyObject: EmptyBucketVersioning, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Versioning")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Website, actual.Website, dcl.DiffInfo{ObjectFunction: compareBucketWebsiteNewStyle, EmptyObject: EmptyBucketWebsite, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Website")); len(ds) != 0 || err != nil {
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
func compareBucketCorsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketCors)
	if !ok {
		desiredNotPointer, ok := d.(BucketCors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketCors or *BucketCors", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketCors)
	if !ok {
		actualNotPointer, ok := a.(BucketCors)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketCors", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MaxAgeSeconds, actual.MaxAgeSeconds, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MaxAgeSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Method, actual.Method, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Method")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Origin, actual.Origin, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Origin")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ResponseHeader, actual.ResponseHeader, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("ResponseHeader")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycle)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycle)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycle or *BucketLifecycle", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycle)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycle)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycle", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Rule, actual.Rule, dcl.DiffInfo{ObjectFunction: compareBucketLifecycleRuleNewStyle, EmptyObject: EmptyBucketLifecycleRule, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Rule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRule)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRule or *BucketLifecycleRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRule)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Action, actual.Action, dcl.DiffInfo{ObjectFunction: compareBucketLifecycleRuleActionNewStyle, EmptyObject: EmptyBucketLifecycleRuleAction, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Action")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Condition, actual.Condition, dcl.DiffInfo{ObjectFunction: compareBucketLifecycleRuleConditionNewStyle, EmptyObject: EmptyBucketLifecycleRuleCondition, OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Condition")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleActionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRuleAction)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRuleAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleAction or *BucketLifecycleRuleAction", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRuleAction)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRuleAction)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleAction", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.StorageClass, actual.StorageClass, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("StorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Type, actual.Type, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Type")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLifecycleRuleConditionNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLifecycleRuleCondition)
	if !ok {
		desiredNotPointer, ok := d.(BucketLifecycleRuleCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleCondition or *BucketLifecycleRuleCondition", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLifecycleRuleCondition)
	if !ok {
		actualNotPointer, ok := a.(BucketLifecycleRuleCondition)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLifecycleRuleCondition", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Age, actual.Age, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Age")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreatedBefore, actual.CreatedBefore, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("CreatedBefore")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.WithState, actual.WithState, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("IsLive")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MatchesStorageClass, actual.MatchesStorageClass, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MatchesStorageClass")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NumNewerVersions, actual.NumNewerVersions, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("NumNewerVersions")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketLoggingNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketLogging)
	if !ok {
		desiredNotPointer, ok := d.(BucketLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLogging or *BucketLogging", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketLogging)
	if !ok {
		actualNotPointer, ok := a.(BucketLogging)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketLogging", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.LogBucket, actual.LogBucket, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("LogBucket")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogObjectPrefix, actual.LogObjectPrefix, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("LogObjectPrefix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketVersioningNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketVersioning)
	if !ok {
		desiredNotPointer, ok := d.(BucketVersioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketVersioning or *BucketVersioning", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketVersioning)
	if !ok {
		actualNotPointer, ok := a.(BucketVersioning)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketVersioning", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareBucketWebsiteNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*BucketWebsite)
	if !ok {
		desiredNotPointer, ok := d.(BucketWebsite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketWebsite or *BucketWebsite", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*BucketWebsite)
	if !ok {
		actualNotPointer, ok := a.(BucketWebsite)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a BucketWebsite", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.MainPageSuffix, actual.MainPageSuffix, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("MainPageSuffix")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.NotFoundPage, actual.NotFoundPage, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBucketUpdateOperation")}, fn.AddNest("NotFoundPage")); len(ds) != 0 || err != nil {
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
func (r *Bucket) urlNormalized() *Bucket {
	normalized := dcl.Copy(*r).(Bucket)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	return &normalized
}

func (r *Bucket) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "update" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("b/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Bucket resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Bucket) marshal(c *Client) ([]byte, error) {
	m, err := expandBucket(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Bucket: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBucket decodes JSON responses into the Bucket resource schema.
func unmarshalBucket(b []byte, c *Client, res *Bucket) (*Bucket, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBucket(m, c, res)
}

func unmarshalMapBucket(m map[string]interface{}, c *Client, res *Bucket) (*Bucket, error) {

	flattened := flattenBucket(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandBucket expands Bucket into a JSON request object.
func expandBucket(c *Client, f *Bucket) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v := f.Location; dcl.ValueShouldBeSent(v) {
		m["location"] = v
	}
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v, err := expandBucketCorsSlice(c, f.Cors, res); err != nil {
		return nil, fmt.Errorf("error expanding Cors into cors: %w", err)
	} else if v != nil {
		m["cors"] = v
	}
	if v, err := expandBucketLifecycle(c, f.Lifecycle, res); err != nil {
		return nil, fmt.Errorf("error expanding Lifecycle into lifecycle: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["lifecycle"] = v
	}
	if v, err := expandBucketLogging(c, f.Logging, res); err != nil {
		return nil, fmt.Errorf("error expanding Logging into logging: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["logging"] = v
	}
	if v := f.StorageClass; dcl.ValueShouldBeSent(v) {
		m["storageClass"] = v
	}
	if v, err := expandBucketVersioning(c, f.Versioning, res); err != nil {
		return nil, fmt.Errorf("error expanding Versioning into versioning: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["versioning"] = v
	}
	if v, err := expandBucketWebsite(c, f.Website, res); err != nil {
		return nil, fmt.Errorf("error expanding Website into website: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["website"] = v
	}

	return m, nil
}

// flattenBucket flattens Bucket from a JSON request object into the
// Bucket type.
func flattenBucket(c *Client, i interface{}, res *Bucket) *Bucket {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Bucket{}
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Cors = flattenBucketCorsSlice(c, m["cors"], res)
	resultRes.Lifecycle = flattenBucketLifecycle(c, m["lifecycle"], res)
	resultRes.Logging = flattenBucketLogging(c, m["logging"], res)
	resultRes.StorageClass = flattenBucketStorageClassEnum(m["storageClass"])
	resultRes.Versioning = flattenBucketVersioning(c, m["versioning"], res)
	resultRes.Website = flattenBucketWebsite(c, m["website"], res)

	return resultRes
}

// expandBucketCorsMap expands the contents of BucketCors into a JSON
// request object.
func expandBucketCorsMap(c *Client, f map[string]BucketCors, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketCors(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketCorsSlice expands the contents of BucketCors into a JSON
// request object.
func expandBucketCorsSlice(c *Client, f []BucketCors, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketCors(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketCorsMap flattens the contents of BucketCors from a JSON
// response object.
func flattenBucketCorsMap(c *Client, i interface{}, res *Bucket) map[string]BucketCors {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketCors{}
	}

	if len(a) == 0 {
		return map[string]BucketCors{}
	}

	items := make(map[string]BucketCors)
	for k, item := range a {
		items[k] = *flattenBucketCors(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketCorsSlice flattens the contents of BucketCors from a JSON
// response object.
func flattenBucketCorsSlice(c *Client, i interface{}, res *Bucket) []BucketCors {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketCors{}
	}

	if len(a) == 0 {
		return []BucketCors{}
	}

	items := make([]BucketCors, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketCors(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketCors expands an instance of BucketCors into a JSON
// request object.
func expandBucketCors(c *Client, f *BucketCors, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MaxAgeSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["maxAgeSeconds"] = v
	}
	if v := f.Method; v != nil {
		m["method"] = v
	}
	if v := f.Origin; v != nil {
		m["origin"] = v
	}
	if v := f.ResponseHeader; v != nil {
		m["responseHeader"] = v
	}

	return m, nil
}

// flattenBucketCors flattens an instance of BucketCors from a JSON
// response object.
func flattenBucketCors(c *Client, i interface{}, res *Bucket) *BucketCors {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketCors{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketCors
	}
	r.MaxAgeSeconds = dcl.FlattenInteger(m["maxAgeSeconds"])
	r.Method = dcl.FlattenStringSlice(m["method"])
	r.Origin = dcl.FlattenStringSlice(m["origin"])
	r.ResponseHeader = dcl.FlattenStringSlice(m["responseHeader"])

	return r
}

// expandBucketLifecycleMap expands the contents of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycleMap(c *Client, f map[string]BucketLifecycle, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycle(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleSlice expands the contents of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycleSlice(c *Client, f []BucketLifecycle, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycle(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleMap flattens the contents of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycleMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycle {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycle{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycle{}
	}

	items := make(map[string]BucketLifecycle)
	for k, item := range a {
		items[k] = *flattenBucketLifecycle(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketLifecycleSlice flattens the contents of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycleSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycle {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycle{}
	}

	if len(a) == 0 {
		return []BucketLifecycle{}
	}

	items := make([]BucketLifecycle, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycle(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketLifecycle expands an instance of BucketLifecycle into a JSON
// request object.
func expandBucketLifecycle(c *Client, f *BucketLifecycle, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBucketLifecycleRuleSlice(c, f.Rule, res); err != nil {
		return nil, fmt.Errorf("error expanding Rule into rule: %w", err)
	} else if v != nil {
		m["rule"] = v
	}

	return m, nil
}

// flattenBucketLifecycle flattens an instance of BucketLifecycle from a JSON
// response object.
func flattenBucketLifecycle(c *Client, i interface{}, res *Bucket) *BucketLifecycle {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycle{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycle
	}
	r.Rule = flattenBucketLifecycleRuleSlice(c, m["rule"], res)

	return r
}

// expandBucketLifecycleRuleMap expands the contents of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRuleMap(c *Client, f map[string]BucketLifecycleRule, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRule(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleSlice expands the contents of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRuleSlice(c *Client, f []BucketLifecycleRule, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRule(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleMap flattens the contents of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRuleMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycleRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRule{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRule{}
	}

	items := make(map[string]BucketLifecycleRule)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRule(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketLifecycleRuleSlice flattens the contents of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRuleSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycleRule {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRule{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRule{}
	}

	items := make([]BucketLifecycleRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRule(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketLifecycleRule expands an instance of BucketLifecycleRule into a JSON
// request object.
func expandBucketLifecycleRule(c *Client, f *BucketLifecycleRule, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandBucketLifecycleRuleAction(c, f.Action, res); err != nil {
		return nil, fmt.Errorf("error expanding Action into action: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["action"] = v
	}
	if v, err := expandBucketLifecycleRuleCondition(c, f.Condition, res); err != nil {
		return nil, fmt.Errorf("error expanding Condition into condition: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["condition"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRule flattens an instance of BucketLifecycleRule from a JSON
// response object.
func flattenBucketLifecycleRule(c *Client, i interface{}, res *Bucket) *BucketLifecycleRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRule
	}
	r.Action = flattenBucketLifecycleRuleAction(c, m["action"], res)
	r.Condition = flattenBucketLifecycleRuleCondition(c, m["condition"], res)

	return r
}

// expandBucketLifecycleRuleActionMap expands the contents of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleActionMap(c *Client, f map[string]BucketLifecycleRuleAction, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRuleAction(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleActionSlice expands the contents of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleActionSlice(c *Client, f []BucketLifecycleRuleAction, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRuleAction(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleActionMap flattens the contents of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleActionMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycleRuleAction {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleAction{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleAction{}
	}

	items := make(map[string]BucketLifecycleRuleAction)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleAction(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketLifecycleRuleActionSlice flattens the contents of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleActionSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycleRuleAction {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleAction{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleAction{}
	}

	items := make([]BucketLifecycleRuleAction, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleAction(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketLifecycleRuleAction expands an instance of BucketLifecycleRuleAction into a JSON
// request object.
func expandBucketLifecycleRuleAction(c *Client, f *BucketLifecycleRuleAction, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.StorageClass; !dcl.IsEmptyValueIndirect(v) {
		m["storageClass"] = v
	}
	if v := f.Type; !dcl.IsEmptyValueIndirect(v) {
		m["type"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRuleAction flattens an instance of BucketLifecycleRuleAction from a JSON
// response object.
func flattenBucketLifecycleRuleAction(c *Client, i interface{}, res *Bucket) *BucketLifecycleRuleAction {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRuleAction{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRuleAction
	}
	r.StorageClass = dcl.FlattenString(m["storageClass"])
	r.Type = flattenBucketLifecycleRuleActionTypeEnum(m["type"])

	return r
}

// expandBucketLifecycleRuleConditionMap expands the contents of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleConditionMap(c *Client, f map[string]BucketLifecycleRuleCondition, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLifecycleRuleCondition(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLifecycleRuleConditionSlice expands the contents of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleConditionSlice(c *Client, f []BucketLifecycleRuleCondition, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLifecycleRuleCondition(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLifecycleRuleConditionMap flattens the contents of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleConditionMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycleRuleCondition {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleCondition{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleCondition{}
	}

	items := make(map[string]BucketLifecycleRuleCondition)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleCondition(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketLifecycleRuleConditionSlice flattens the contents of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleConditionSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycleRuleCondition {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleCondition{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleCondition{}
	}

	items := make([]BucketLifecycleRuleCondition, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleCondition(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketLifecycleRuleCondition expands an instance of BucketLifecycleRuleCondition into a JSON
// request object.
func expandBucketLifecycleRuleCondition(c *Client, f *BucketLifecycleRuleCondition, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Age; !dcl.IsEmptyValueIndirect(v) {
		m["age"] = v
	}
	if v := f.CreatedBefore; !dcl.IsEmptyValueIndirect(v) {
		m["createdBefore"] = v
	}
	if v, err := expandStorageBucketLifecycleWithState(c, f.WithState, res); err != nil {
		return nil, fmt.Errorf("error expanding WithState into isLive: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["isLive"] = v
	}
	if v := f.MatchesStorageClass; v != nil {
		m["matchesStorageClass"] = v
	}
	if v := f.NumNewerVersions; !dcl.IsEmptyValueIndirect(v) {
		m["numNewerVersions"] = v
	}

	return m, nil
}

// flattenBucketLifecycleRuleCondition flattens an instance of BucketLifecycleRuleCondition from a JSON
// response object.
func flattenBucketLifecycleRuleCondition(c *Client, i interface{}, res *Bucket) *BucketLifecycleRuleCondition {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLifecycleRuleCondition{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLifecycleRuleCondition
	}
	r.Age = dcl.FlattenInteger(m["age"])
	r.CreatedBefore = dcl.FlattenString(m["createdBefore"])
	r.WithState = flattenStorageBucketLifecycleWithState(c, m["isLive"], res)
	r.MatchesStorageClass = dcl.FlattenStringSlice(m["matchesStorageClass"])
	r.NumNewerVersions = dcl.FlattenInteger(m["numNewerVersions"])

	return r
}

// expandBucketLoggingMap expands the contents of BucketLogging into a JSON
// request object.
func expandBucketLoggingMap(c *Client, f map[string]BucketLogging, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketLogging(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketLoggingSlice expands the contents of BucketLogging into a JSON
// request object.
func expandBucketLoggingSlice(c *Client, f []BucketLogging, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketLogging(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketLoggingMap flattens the contents of BucketLogging from a JSON
// response object.
func flattenBucketLoggingMap(c *Client, i interface{}, res *Bucket) map[string]BucketLogging {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLogging{}
	}

	if len(a) == 0 {
		return map[string]BucketLogging{}
	}

	items := make(map[string]BucketLogging)
	for k, item := range a {
		items[k] = *flattenBucketLogging(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketLoggingSlice flattens the contents of BucketLogging from a JSON
// response object.
func flattenBucketLoggingSlice(c *Client, i interface{}, res *Bucket) []BucketLogging {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLogging{}
	}

	if len(a) == 0 {
		return []BucketLogging{}
	}

	items := make([]BucketLogging, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLogging(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketLogging expands an instance of BucketLogging into a JSON
// request object.
func expandBucketLogging(c *Client, f *BucketLogging, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.LogBucket; !dcl.IsEmptyValueIndirect(v) {
		m["logBucket"] = v
	}
	if v := f.LogObjectPrefix; !dcl.IsEmptyValueIndirect(v) {
		m["logObjectPrefix"] = v
	}

	return m, nil
}

// flattenBucketLogging flattens an instance of BucketLogging from a JSON
// response object.
func flattenBucketLogging(c *Client, i interface{}, res *Bucket) *BucketLogging {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketLogging{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketLogging
	}
	r.LogBucket = dcl.FlattenString(m["logBucket"])
	r.LogObjectPrefix = dcl.FlattenString(m["logObjectPrefix"])

	return r
}

// expandBucketVersioningMap expands the contents of BucketVersioning into a JSON
// request object.
func expandBucketVersioningMap(c *Client, f map[string]BucketVersioning, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketVersioning(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketVersioningSlice expands the contents of BucketVersioning into a JSON
// request object.
func expandBucketVersioningSlice(c *Client, f []BucketVersioning, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketVersioning(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketVersioningMap flattens the contents of BucketVersioning from a JSON
// response object.
func flattenBucketVersioningMap(c *Client, i interface{}, res *Bucket) map[string]BucketVersioning {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketVersioning{}
	}

	if len(a) == 0 {
		return map[string]BucketVersioning{}
	}

	items := make(map[string]BucketVersioning)
	for k, item := range a {
		items[k] = *flattenBucketVersioning(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketVersioningSlice flattens the contents of BucketVersioning from a JSON
// response object.
func flattenBucketVersioningSlice(c *Client, i interface{}, res *Bucket) []BucketVersioning {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketVersioning{}
	}

	if len(a) == 0 {
		return []BucketVersioning{}
	}

	items := make([]BucketVersioning, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketVersioning(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketVersioning expands an instance of BucketVersioning into a JSON
// request object.
func expandBucketVersioning(c *Client, f *BucketVersioning, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenBucketVersioning flattens an instance of BucketVersioning from a JSON
// response object.
func flattenBucketVersioning(c *Client, i interface{}, res *Bucket) *BucketVersioning {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketVersioning{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketVersioning
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandBucketWebsiteMap expands the contents of BucketWebsite into a JSON
// request object.
func expandBucketWebsiteMap(c *Client, f map[string]BucketWebsite, res *Bucket) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandBucketWebsite(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandBucketWebsiteSlice expands the contents of BucketWebsite into a JSON
// request object.
func expandBucketWebsiteSlice(c *Client, f []BucketWebsite, res *Bucket) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandBucketWebsite(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenBucketWebsiteMap flattens the contents of BucketWebsite from a JSON
// response object.
func flattenBucketWebsiteMap(c *Client, i interface{}, res *Bucket) map[string]BucketWebsite {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketWebsite{}
	}

	if len(a) == 0 {
		return map[string]BucketWebsite{}
	}

	items := make(map[string]BucketWebsite)
	for k, item := range a {
		items[k] = *flattenBucketWebsite(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenBucketWebsiteSlice flattens the contents of BucketWebsite from a JSON
// response object.
func flattenBucketWebsiteSlice(c *Client, i interface{}, res *Bucket) []BucketWebsite {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketWebsite{}
	}

	if len(a) == 0 {
		return []BucketWebsite{}
	}

	items := make([]BucketWebsite, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketWebsite(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandBucketWebsite expands an instance of BucketWebsite into a JSON
// request object.
func expandBucketWebsite(c *Client, f *BucketWebsite, res *Bucket) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.MainPageSuffix; !dcl.IsEmptyValueIndirect(v) {
		m["mainPageSuffix"] = v
	}
	if v := f.NotFoundPage; !dcl.IsEmptyValueIndirect(v) {
		m["notFoundPage"] = v
	}

	return m, nil
}

// flattenBucketWebsite flattens an instance of BucketWebsite from a JSON
// response object.
func flattenBucketWebsite(c *Client, i interface{}, res *Bucket) *BucketWebsite {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &BucketWebsite{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyBucketWebsite
	}
	r.MainPageSuffix = dcl.FlattenString(m["mainPageSuffix"])
	r.NotFoundPage = dcl.FlattenString(m["notFoundPage"])

	return r
}

// flattenBucketLifecycleRuleActionTypeEnumMap flattens the contents of BucketLifecycleRuleActionTypeEnum from a JSON
// response object.
func flattenBucketLifecycleRuleActionTypeEnumMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycleRuleActionTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleActionTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleActionTypeEnum{}
	}

	items := make(map[string]BucketLifecycleRuleActionTypeEnum)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleActionTypeEnum(item.(interface{}))
	}

	return items
}

// flattenBucketLifecycleRuleActionTypeEnumSlice flattens the contents of BucketLifecycleRuleActionTypeEnum from a JSON
// response object.
func flattenBucketLifecycleRuleActionTypeEnumSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycleRuleActionTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleActionTypeEnum{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleActionTypeEnum{}
	}

	items := make([]BucketLifecycleRuleActionTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleActionTypeEnum(item.(interface{})))
	}

	return items
}

// flattenBucketLifecycleRuleActionTypeEnum asserts that an interface is a string, and returns a
// pointer to a *BucketLifecycleRuleActionTypeEnum with the same value as that string.
func flattenBucketLifecycleRuleActionTypeEnum(i interface{}) *BucketLifecycleRuleActionTypeEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BucketLifecycleRuleActionTypeEnumRef(s)
}

// flattenBucketLifecycleRuleConditionWithStateEnumMap flattens the contents of BucketLifecycleRuleConditionWithStateEnum from a JSON
// response object.
func flattenBucketLifecycleRuleConditionWithStateEnumMap(c *Client, i interface{}, res *Bucket) map[string]BucketLifecycleRuleConditionWithStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketLifecycleRuleConditionWithStateEnum{}
	}

	if len(a) == 0 {
		return map[string]BucketLifecycleRuleConditionWithStateEnum{}
	}

	items := make(map[string]BucketLifecycleRuleConditionWithStateEnum)
	for k, item := range a {
		items[k] = *flattenBucketLifecycleRuleConditionWithStateEnum(item.(interface{}))
	}

	return items
}

// flattenBucketLifecycleRuleConditionWithStateEnumSlice flattens the contents of BucketLifecycleRuleConditionWithStateEnum from a JSON
// response object.
func flattenBucketLifecycleRuleConditionWithStateEnumSlice(c *Client, i interface{}, res *Bucket) []BucketLifecycleRuleConditionWithStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketLifecycleRuleConditionWithStateEnum{}
	}

	if len(a) == 0 {
		return []BucketLifecycleRuleConditionWithStateEnum{}
	}

	items := make([]BucketLifecycleRuleConditionWithStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketLifecycleRuleConditionWithStateEnum(item.(interface{})))
	}

	return items
}

// flattenBucketLifecycleRuleConditionWithStateEnum asserts that an interface is a string, and returns a
// pointer to a *BucketLifecycleRuleConditionWithStateEnum with the same value as that string.
func flattenBucketLifecycleRuleConditionWithStateEnum(i interface{}) *BucketLifecycleRuleConditionWithStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BucketLifecycleRuleConditionWithStateEnumRef(s)
}

// flattenBucketStorageClassEnumMap flattens the contents of BucketStorageClassEnum from a JSON
// response object.
func flattenBucketStorageClassEnumMap(c *Client, i interface{}, res *Bucket) map[string]BucketStorageClassEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BucketStorageClassEnum{}
	}

	if len(a) == 0 {
		return map[string]BucketStorageClassEnum{}
	}

	items := make(map[string]BucketStorageClassEnum)
	for k, item := range a {
		items[k] = *flattenBucketStorageClassEnum(item.(interface{}))
	}

	return items
}

// flattenBucketStorageClassEnumSlice flattens the contents of BucketStorageClassEnum from a JSON
// response object.
func flattenBucketStorageClassEnumSlice(c *Client, i interface{}, res *Bucket) []BucketStorageClassEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BucketStorageClassEnum{}
	}

	if len(a) == 0 {
		return []BucketStorageClassEnum{}
	}

	items := make([]BucketStorageClassEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBucketStorageClassEnum(item.(interface{})))
	}

	return items
}

// flattenBucketStorageClassEnum asserts that an interface is a string, and returns a
// pointer to a *BucketStorageClassEnum with the same value as that string.
func flattenBucketStorageClassEnum(i interface{}) *BucketStorageClassEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BucketStorageClassEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Bucket) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBucket(b, c, r)
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

type bucketDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         bucketApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToBucketDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]bucketDiff, error) {
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
	var diffs []bucketDiff
	// For each operation name, create a bucketDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := bucketDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBucketApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBucketApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (bucketApiOperation, error) {
	switch opName {

	case "updateBucketUpdateOperation":
		return &updateBucketUpdateOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractBucketFields(r *Bucket) error {
	vLifecycle := r.Lifecycle
	if vLifecycle == nil {
		// note: explicitly not the empty object.
		vLifecycle = &BucketLifecycle{}
	}
	if err := extractBucketLifecycleFields(r, vLifecycle); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLifecycle) {
		r.Lifecycle = vLifecycle
	}
	vLogging := r.Logging
	if vLogging == nil {
		// note: explicitly not the empty object.
		vLogging = &BucketLogging{}
	}
	if err := extractBucketLoggingFields(r, vLogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLogging) {
		r.Logging = vLogging
	}
	vVersioning := r.Versioning
	if vVersioning == nil {
		// note: explicitly not the empty object.
		vVersioning = &BucketVersioning{}
	}
	if err := extractBucketVersioningFields(r, vVersioning); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVersioning) {
		r.Versioning = vVersioning
	}
	vWebsite := r.Website
	if vWebsite == nil {
		// note: explicitly not the empty object.
		vWebsite = &BucketWebsite{}
	}
	if err := extractBucketWebsiteFields(r, vWebsite); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWebsite) {
		r.Website = vWebsite
	}
	return nil
}
func extractBucketCorsFields(r *Bucket, o *BucketCors) error {
	return nil
}
func extractBucketLifecycleFields(r *Bucket, o *BucketLifecycle) error {
	return nil
}
func extractBucketLifecycleRuleFields(r *Bucket, o *BucketLifecycleRule) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &BucketLifecycleRuleAction{}
	}
	if err := extractBucketLifecycleRuleActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	vCondition := o.Condition
	if vCondition == nil {
		// note: explicitly not the empty object.
		vCondition = &BucketLifecycleRuleCondition{}
	}
	if err := extractBucketLifecycleRuleConditionFields(r, vCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCondition) {
		o.Condition = vCondition
	}
	return nil
}
func extractBucketLifecycleRuleActionFields(r *Bucket, o *BucketLifecycleRuleAction) error {
	return nil
}
func extractBucketLifecycleRuleConditionFields(r *Bucket, o *BucketLifecycleRuleCondition) error {
	return nil
}
func extractBucketLoggingFields(r *Bucket, o *BucketLogging) error {
	return nil
}
func extractBucketVersioningFields(r *Bucket, o *BucketVersioning) error {
	return nil
}
func extractBucketWebsiteFields(r *Bucket, o *BucketWebsite) error {
	return nil
}

func postReadExtractBucketFields(r *Bucket) error {
	vLifecycle := r.Lifecycle
	if vLifecycle == nil {
		// note: explicitly not the empty object.
		vLifecycle = &BucketLifecycle{}
	}
	if err := postReadExtractBucketLifecycleFields(r, vLifecycle); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLifecycle) {
		r.Lifecycle = vLifecycle
	}
	vLogging := r.Logging
	if vLogging == nil {
		// note: explicitly not the empty object.
		vLogging = &BucketLogging{}
	}
	if err := postReadExtractBucketLoggingFields(r, vLogging); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLogging) {
		r.Logging = vLogging
	}
	vVersioning := r.Versioning
	if vVersioning == nil {
		// note: explicitly not the empty object.
		vVersioning = &BucketVersioning{}
	}
	if err := postReadExtractBucketVersioningFields(r, vVersioning); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vVersioning) {
		r.Versioning = vVersioning
	}
	vWebsite := r.Website
	if vWebsite == nil {
		// note: explicitly not the empty object.
		vWebsite = &BucketWebsite{}
	}
	if err := postReadExtractBucketWebsiteFields(r, vWebsite); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWebsite) {
		r.Website = vWebsite
	}
	return nil
}
func postReadExtractBucketCorsFields(r *Bucket, o *BucketCors) error {
	return nil
}
func postReadExtractBucketLifecycleFields(r *Bucket, o *BucketLifecycle) error {
	return nil
}
func postReadExtractBucketLifecycleRuleFields(r *Bucket, o *BucketLifecycleRule) error {
	vAction := o.Action
	if vAction == nil {
		// note: explicitly not the empty object.
		vAction = &BucketLifecycleRuleAction{}
	}
	if err := extractBucketLifecycleRuleActionFields(r, vAction); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vAction) {
		o.Action = vAction
	}
	vCondition := o.Condition
	if vCondition == nil {
		// note: explicitly not the empty object.
		vCondition = &BucketLifecycleRuleCondition{}
	}
	if err := extractBucketLifecycleRuleConditionFields(r, vCondition); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCondition) {
		o.Condition = vCondition
	}
	return nil
}
func postReadExtractBucketLifecycleRuleActionFields(r *Bucket, o *BucketLifecycleRuleAction) error {
	return nil
}
func postReadExtractBucketLifecycleRuleConditionFields(r *Bucket, o *BucketLifecycleRuleCondition) error {
	return nil
}
func postReadExtractBucketLoggingFields(r *Bucket, o *BucketLogging) error {
	return nil
}
func postReadExtractBucketVersioningFields(r *Bucket, o *BucketVersioning) error {
	return nil
}
func postReadExtractBucketWebsiteFields(r *Bucket, o *BucketWebsite) error {
	return nil
}
