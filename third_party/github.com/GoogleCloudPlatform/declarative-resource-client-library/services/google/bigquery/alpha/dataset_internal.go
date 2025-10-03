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

func (r *Dataset) validate() error {

	if err := dcl.Required(r, "name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "project"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultEncryptionConfiguration) {
		if err := r.DefaultEncryptionConfiguration.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DatasetAccess) validate() error {
	if err := dcl.Required(r, "role"); err != nil {
		return err
	}
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"UserByEmail", "GroupByEmail", "Domain", "SpecialGroup", "IamMember", "View", "Routine"}, r.UserByEmail, r.GroupByEmail, r.Domain, r.SpecialGroup, r.IamMember, r.View, r.Routine); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.View) {
		if err := r.View.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Routine) {
		if err := r.Routine.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *DatasetAccessView) validate() error {
	if err := dcl.Required(r, "projectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "datasetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "tableId"); err != nil {
		return err
	}
	return nil
}
func (r *DatasetAccessRoutine) validate() error {
	if err := dcl.Required(r, "projectId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "datasetId"); err != nil {
		return err
	}
	if err := dcl.Required(r, "routineId"); err != nil {
		return err
	}
	return nil
}
func (r *DatasetDefaultEncryptionConfiguration) validate() error {
	return nil
}
func (r *Dataset) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://bigquery.googleapis.com/bigquery/v2/", params)
}

func (r *Dataset) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Dataset) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/datasets", nr.basePath(), userBasePath, params), nil

}

func (r *Dataset) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/datasets", nr.basePath(), userBasePath, params), nil

}

func (r *Dataset) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
		"name":    dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/datasets/{{name}}", nr.basePath(), userBasePath, params), nil
}

// datasetApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type datasetApiOperation interface {
	do(context.Context, *Dataset, *Client) error
}

// newUpdateDatasetPatchDatasetRequest creates a request for an
// Dataset resource's PatchDataset update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateDatasetPatchDatasetRequest(ctx context.Context, f *Dataset, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.FriendlyName; !dcl.IsEmptyValueIndirect(v) {
		req["friendlyName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.DefaultTableExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		req["defaultTableExpirationMs"] = v
	}
	if v := f.DefaultPartitionExpirationMs; !dcl.IsEmptyValueIndirect(v) {
		req["defaultPartitionExpirationMs"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandDatasetAccessSlice(c, f.Access, res); err != nil {
		return nil, fmt.Errorf("error expanding Access into access: %w", err)
	} else if v != nil {
		req["access"] = v
	}
	b, err := c.getDatasetRaw(ctx, f)
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

// marshalUpdateDatasetPatchDatasetRequest converts the update into
// the final JSON request body.
func marshalUpdateDatasetPatchDatasetRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"datasetReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"datasetReference", "projectId"},
	)
	return json.Marshal(m)
}

type updateDatasetPatchDatasetOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateDatasetPatchDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {
	_, err := c.GetDataset(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "PatchDataset")
	if err != nil {
		return err
	}

	req, err := newUpdateDatasetPatchDatasetRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateDatasetPatchDatasetRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listDatasetRaw(ctx context.Context, r *Dataset, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != DatasetMaxPage {
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

type listDatasetOperation struct {
	Datasets []map[string]interface{} `json:"datasets"`
	Token    string                   `json:"nextPageToken"`
}

func (c *Client) listDataset(ctx context.Context, r *Dataset, pageToken string, pageSize int32) ([]*Dataset, string, error) {
	b, err := c.listDatasetRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listDatasetOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Dataset
	for _, v := range m.Datasets {
		res, err := unmarshalMapDataset(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllDataset(ctx context.Context, f func(*Dataset) bool, resources []*Dataset) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteDataset(ctx, res)
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

type deleteDatasetOperation struct{}

func (op *deleteDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {
	r, err := c.GetDataset(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Dataset not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetDataset checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete Dataset: %w", err)
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetDataset(ctx, r)
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
type createDatasetOperation struct {
	response map[string]interface{}
}

func (op *createDatasetOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createDatasetOperation) do(ctx context.Context, r *Dataset, c *Client) error {
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

	if _, err := c.GetDataset(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getDatasetRaw(ctx context.Context, r *Dataset) ([]byte, error) {

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

func (c *Client) datasetDiffsForRawDesired(ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (initial, desired *Dataset, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Dataset
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Dataset); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Dataset, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetDataset(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Dataset resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Dataset resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Dataset resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeDatasetDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Dataset: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Dataset: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractDatasetFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeDatasetInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Dataset: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeDatasetDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Dataset: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffDataset(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeDatasetInitialState(rawInitial, rawDesired *Dataset) (*Dataset, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeDatasetDesiredState(rawDesired, rawInitial *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultEncryptionConfiguration = canonicalizeDatasetDefaultEncryptionConfiguration(rawDesired.DefaultEncryptionConfiguration, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Dataset{}
	if dcl.StringCanonicalize(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Project) || (dcl.IsEmptyValueIndirect(rawDesired.Project) && dcl.IsEmptyValueIndirect(rawInitial.Project)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.StringCanonicalize(rawDesired.FriendlyName, rawInitial.FriendlyName) {
		canonicalDesired.FriendlyName = rawInitial.FriendlyName
	} else {
		canonicalDesired.FriendlyName = rawDesired.FriendlyName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.StringCanonicalize(rawDesired.DefaultTableExpirationMs, rawInitial.DefaultTableExpirationMs) {
		canonicalDesired.DefaultTableExpirationMs = rawInitial.DefaultTableExpirationMs
	} else {
		canonicalDesired.DefaultTableExpirationMs = rawDesired.DefaultTableExpirationMs
	}
	if dcl.StringCanonicalize(rawDesired.DefaultPartitionExpirationMs, rawInitial.DefaultPartitionExpirationMs) {
		canonicalDesired.DefaultPartitionExpirationMs = rawInitial.DefaultPartitionExpirationMs
	} else {
		canonicalDesired.DefaultPartitionExpirationMs = rawDesired.DefaultPartitionExpirationMs
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.Access = canonicalizeDatasetAccessSlice(rawDesired.Access, rawInitial.Access, opts...)
	if dcl.StringCanonicalize(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	if dcl.BoolCanonicalize(rawDesired.Published, rawInitial.Published) {
		canonicalDesired.Published = rawInitial.Published
	} else {
		canonicalDesired.Published = rawDesired.Published
	}
	canonicalDesired.DefaultEncryptionConfiguration = canonicalizeDatasetDefaultEncryptionConfiguration(rawDesired.DefaultEncryptionConfiguration, rawInitial.DefaultEncryptionConfiguration, opts...)
	return canonicalDesired, nil
}

func canonicalizeDatasetNewState(c *Client, rawNew, rawDesired *Dataset) (*Dataset, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Etag) && dcl.IsEmptyValueIndirect(rawDesired.Etag) {
		rawNew.Etag = rawDesired.Etag
	} else {
		if dcl.StringCanonicalize(rawDesired.Etag, rawNew.Etag) {
			rawNew.Etag = rawDesired.Etag
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Id) && dcl.IsEmptyValueIndirect(rawDesired.Id) {
		rawNew.Id = rawDesired.Id
	} else {
		if dcl.StringCanonicalize(rawDesired.Id, rawNew.Id) {
			rawNew.Id = rawDesired.Id
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SelfLink) && dcl.IsEmptyValueIndirect(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.StringCanonicalize(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
		if dcl.StringCanonicalize(rawDesired.Name, rawNew.Name) {
			rawNew.Name = rawDesired.Name
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Project) && dcl.IsEmptyValueIndirect(rawDesired.Project) {
		rawNew.Project = rawDesired.Project
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.FriendlyName) && dcl.IsEmptyValueIndirect(rawDesired.FriendlyName) {
		rawNew.FriendlyName = rawDesired.FriendlyName
	} else {
		if dcl.StringCanonicalize(rawDesired.FriendlyName, rawNew.FriendlyName) {
			rawNew.FriendlyName = rawDesired.FriendlyName
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultTableExpirationMs) && dcl.IsEmptyValueIndirect(rawDesired.DefaultTableExpirationMs) {
		rawNew.DefaultTableExpirationMs = rawDesired.DefaultTableExpirationMs
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultTableExpirationMs, rawNew.DefaultTableExpirationMs) {
			rawNew.DefaultTableExpirationMs = rawDesired.DefaultTableExpirationMs
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultPartitionExpirationMs) && dcl.IsEmptyValueIndirect(rawDesired.DefaultPartitionExpirationMs) {
		rawNew.DefaultPartitionExpirationMs = rawDesired.DefaultPartitionExpirationMs
	} else {
		if dcl.StringCanonicalize(rawDesired.DefaultPartitionExpirationMs, rawNew.DefaultPartitionExpirationMs) {
			rawNew.DefaultPartitionExpirationMs = rawDesired.DefaultPartitionExpirationMs
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Access) && dcl.IsEmptyValueIndirect(rawDesired.Access) {
		rawNew.Access = rawDesired.Access
	} else {
		rawNew.Access = canonicalizeNewDatasetAccessSet(c, rawDesired.Access, rawNew.Access)
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreationTime) && dcl.IsEmptyValueIndirect(rawDesired.CreationTime) {
		rawNew.CreationTime = rawDesired.CreationTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.LastModifiedTime) && dcl.IsEmptyValueIndirect(rawDesired.LastModifiedTime) {
		rawNew.LastModifiedTime = rawDesired.LastModifiedTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Location) && dcl.IsEmptyValueIndirect(rawDesired.Location) {
		rawNew.Location = rawDesired.Location
	} else {
		if dcl.StringCanonicalize(rawDesired.Location, rawNew.Location) {
			rawNew.Location = rawDesired.Location
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Published) && dcl.IsEmptyValueIndirect(rawDesired.Published) {
		rawNew.Published = rawDesired.Published
	} else {
		if dcl.BoolCanonicalize(rawDesired.Published, rawNew.Published) {
			rawNew.Published = rawDesired.Published
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.DefaultEncryptionConfiguration) && dcl.IsEmptyValueIndirect(rawDesired.DefaultEncryptionConfiguration) {
		rawNew.DefaultEncryptionConfiguration = rawDesired.DefaultEncryptionConfiguration
	} else {
		rawNew.DefaultEncryptionConfiguration = canonicalizeNewDatasetDefaultEncryptionConfiguration(c, rawDesired.DefaultEncryptionConfiguration, rawNew.DefaultEncryptionConfiguration)
	}

	return rawNew, nil
}

func canonicalizeDatasetAccess(des, initial *DatasetAccess, opts ...dcl.ApplyOption) *DatasetAccess {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.UserByEmail != nil || (initial != nil && initial.UserByEmail != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.GroupByEmail, des.Domain, des.SpecialGroup, des.IamMember, des.View, des.Routine) {
			des.UserByEmail = nil
			if initial != nil {
				initial.UserByEmail = nil
			}
		}
	}

	if des.GroupByEmail != nil || (initial != nil && initial.GroupByEmail != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.Domain, des.SpecialGroup, des.IamMember, des.View, des.Routine) {
			des.GroupByEmail = nil
			if initial != nil {
				initial.GroupByEmail = nil
			}
		}
	}

	if des.Domain != nil || (initial != nil && initial.Domain != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.GroupByEmail, des.SpecialGroup, des.IamMember, des.View, des.Routine) {
			des.Domain = nil
			if initial != nil {
				initial.Domain = nil
			}
		}
	}

	if des.SpecialGroup != nil || (initial != nil && initial.SpecialGroup != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.GroupByEmail, des.Domain, des.IamMember, des.View, des.Routine) {
			des.SpecialGroup = nil
			if initial != nil {
				initial.SpecialGroup = nil
			}
		}
	}

	if des.IamMember != nil || (initial != nil && initial.IamMember != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.GroupByEmail, des.Domain, des.SpecialGroup, des.View, des.Routine) {
			des.IamMember = nil
			if initial != nil {
				initial.IamMember = nil
			}
		}
	}

	if des.View != nil || (initial != nil && initial.View != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.GroupByEmail, des.Domain, des.SpecialGroup, des.IamMember, des.Routine) {
			des.View = nil
			if initial != nil {
				initial.View = nil
			}
		}
	}

	if des.Routine != nil || (initial != nil && initial.Routine != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.UserByEmail, des.GroupByEmail, des.Domain, des.SpecialGroup, des.IamMember, des.View) {
			des.Routine = nil
			if initial != nil {
				initial.Routine = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &DatasetAccess{}

	if canonicalizeDatasetAccessRole(des.Role, initial.Role) || dcl.IsZeroValue(des.Role) {
		cDes.Role = initial.Role
	} else {
		cDes.Role = des.Role
	}
	if dcl.StringCanonicalize(des.UserByEmail, initial.UserByEmail) || dcl.IsZeroValue(des.UserByEmail) {
		cDes.UserByEmail = initial.UserByEmail
	} else {
		cDes.UserByEmail = des.UserByEmail
	}
	if dcl.StringCanonicalize(des.GroupByEmail, initial.GroupByEmail) || dcl.IsZeroValue(des.GroupByEmail) {
		cDes.GroupByEmail = initial.GroupByEmail
	} else {
		cDes.GroupByEmail = des.GroupByEmail
	}
	if dcl.StringCanonicalize(des.Domain, initial.Domain) || dcl.IsZeroValue(des.Domain) {
		cDes.Domain = initial.Domain
	} else {
		cDes.Domain = des.Domain
	}
	if dcl.StringCanonicalize(des.SpecialGroup, initial.SpecialGroup) || dcl.IsZeroValue(des.SpecialGroup) {
		cDes.SpecialGroup = initial.SpecialGroup
	} else {
		cDes.SpecialGroup = des.SpecialGroup
	}
	if dcl.StringCanonicalize(des.IamMember, initial.IamMember) || dcl.IsZeroValue(des.IamMember) {
		cDes.IamMember = initial.IamMember
	} else {
		cDes.IamMember = des.IamMember
	}
	cDes.View = canonicalizeDatasetAccessView(des.View, initial.View, opts...)
	cDes.Routine = canonicalizeDatasetAccessRoutine(des.Routine, initial.Routine, opts...)

	return cDes
}

func canonicalizeDatasetAccessSlice(des, initial []DatasetAccess, opts ...dcl.ApplyOption) []DatasetAccess {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DatasetAccess, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDatasetAccess(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DatasetAccess, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDatasetAccess(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDatasetAccess(c *Client, des, nw *DatasetAccess) *DatasetAccess {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DatasetAccess while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if canonicalizeDatasetAccessRole(des.Role, nw.Role) {
		nw.Role = des.Role
	}
	if dcl.StringCanonicalize(des.UserByEmail, nw.UserByEmail) {
		nw.UserByEmail = des.UserByEmail
	}
	if dcl.StringCanonicalize(des.GroupByEmail, nw.GroupByEmail) {
		nw.GroupByEmail = des.GroupByEmail
	}
	if dcl.StringCanonicalize(des.Domain, nw.Domain) {
		nw.Domain = des.Domain
	}
	if dcl.StringCanonicalize(des.SpecialGroup, nw.SpecialGroup) {
		nw.SpecialGroup = des.SpecialGroup
	}
	if dcl.StringCanonicalize(des.IamMember, nw.IamMember) {
		nw.IamMember = des.IamMember
	}
	nw.View = canonicalizeNewDatasetAccessView(c, des.View, nw.View)
	nw.Routine = canonicalizeNewDatasetAccessRoutine(c, des.Routine, nw.Routine)

	return nw
}

func canonicalizeNewDatasetAccessSet(c *Client, des, nw []DatasetAccess) []DatasetAccess {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DatasetAccess
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDatasetAccessNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDatasetAccess(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDatasetAccessSlice(c *Client, des, nw []DatasetAccess) []DatasetAccess {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccess
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccess(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetAccessView(des, initial *DatasetAccessView, opts ...dcl.ApplyOption) *DatasetAccessView {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DatasetAccessView{}

	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.IsZeroValue(des.DatasetId) || (dcl.IsEmptyValueIndirect(des.DatasetId) && dcl.IsEmptyValueIndirect(initial.DatasetId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DatasetId = initial.DatasetId
	} else {
		cDes.DatasetId = des.DatasetId
	}
	if dcl.IsZeroValue(des.TableId) || (dcl.IsEmptyValueIndirect(des.TableId) && dcl.IsEmptyValueIndirect(initial.TableId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.TableId = initial.TableId
	} else {
		cDes.TableId = des.TableId
	}

	return cDes
}

func canonicalizeDatasetAccessViewSlice(des, initial []DatasetAccessView, opts ...dcl.ApplyOption) []DatasetAccessView {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DatasetAccessView, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDatasetAccessView(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DatasetAccessView, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDatasetAccessView(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDatasetAccessView(c *Client, des, nw *DatasetAccessView) *DatasetAccessView {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DatasetAccessView while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDatasetAccessViewSet(c *Client, des, nw []DatasetAccessView) []DatasetAccessView {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DatasetAccessView
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDatasetAccessViewNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDatasetAccessView(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDatasetAccessViewSlice(c *Client, des, nw []DatasetAccessView) []DatasetAccessView {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccessView
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccessView(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetAccessRoutine(des, initial *DatasetAccessRoutine, opts ...dcl.ApplyOption) *DatasetAccessRoutine {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DatasetAccessRoutine{}

	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}
	if dcl.IsZeroValue(des.DatasetId) || (dcl.IsEmptyValueIndirect(des.DatasetId) && dcl.IsEmptyValueIndirect(initial.DatasetId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.DatasetId = initial.DatasetId
	} else {
		cDes.DatasetId = des.DatasetId
	}
	if dcl.IsZeroValue(des.RoutineId) || (dcl.IsEmptyValueIndirect(des.RoutineId) && dcl.IsEmptyValueIndirect(initial.RoutineId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.RoutineId = initial.RoutineId
	} else {
		cDes.RoutineId = des.RoutineId
	}

	return cDes
}

func canonicalizeDatasetAccessRoutineSlice(des, initial []DatasetAccessRoutine, opts ...dcl.ApplyOption) []DatasetAccessRoutine {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DatasetAccessRoutine, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDatasetAccessRoutine(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DatasetAccessRoutine, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDatasetAccessRoutine(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDatasetAccessRoutine(c *Client, des, nw *DatasetAccessRoutine) *DatasetAccessRoutine {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DatasetAccessRoutine while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDatasetAccessRoutineSet(c *Client, des, nw []DatasetAccessRoutine) []DatasetAccessRoutine {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DatasetAccessRoutine
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDatasetAccessRoutineNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDatasetAccessRoutine(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDatasetAccessRoutineSlice(c *Client, des, nw []DatasetAccessRoutine) []DatasetAccessRoutine {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetAccessRoutine
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetAccessRoutine(c, &d, &n))
	}

	return items
}

func canonicalizeDatasetDefaultEncryptionConfiguration(des, initial *DatasetDefaultEncryptionConfiguration, opts ...dcl.ApplyOption) *DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &DatasetDefaultEncryptionConfiguration{}

	if dcl.IsZeroValue(des.KmsKeyName) || (dcl.IsEmptyValueIndirect(des.KmsKeyName) && dcl.IsEmptyValueIndirect(initial.KmsKeyName)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.KmsKeyName = initial.KmsKeyName
	} else {
		cDes.KmsKeyName = des.KmsKeyName
	}

	return cDes
}

func canonicalizeDatasetDefaultEncryptionConfigurationSlice(des, initial []DatasetDefaultEncryptionConfiguration, opts ...dcl.ApplyOption) []DatasetDefaultEncryptionConfiguration {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]DatasetDefaultEncryptionConfiguration, 0, len(des))
		for _, d := range des {
			cd := canonicalizeDatasetDefaultEncryptionConfiguration(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]DatasetDefaultEncryptionConfiguration, 0, len(des))
	for i, d := range des {
		cd := canonicalizeDatasetDefaultEncryptionConfiguration(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewDatasetDefaultEncryptionConfiguration(c *Client, des, nw *DatasetDefaultEncryptionConfiguration) *DatasetDefaultEncryptionConfiguration {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for DatasetDefaultEncryptionConfiguration while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewDatasetDefaultEncryptionConfigurationSet(c *Client, des, nw []DatasetDefaultEncryptionConfiguration) []DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []DatasetDefaultEncryptionConfiguration
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareDatasetDefaultEncryptionConfigurationNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewDatasetDefaultEncryptionConfiguration(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewDatasetDefaultEncryptionConfigurationSlice(c *Client, des, nw []DatasetDefaultEncryptionConfiguration) []DatasetDefaultEncryptionConfiguration {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []DatasetDefaultEncryptionConfiguration
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewDatasetDefaultEncryptionConfiguration(c, &d, &n))
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
func diffDataset(c *Client, desired, actual *Dataset, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Etag, actual.Etag, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Etag")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SelfLink")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.FriendlyName, actual.FriendlyName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("FriendlyName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultTableExpirationMs, actual.DefaultTableExpirationMs, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("DefaultTableExpirationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultPartitionExpirationMs, actual.DefaultPartitionExpirationMs, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("DefaultPartitionExpirationMs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Access, actual.Access, dcl.DiffInfo{Type: "Set", ObjectFunction: compareDatasetAccessNewStyle, EmptyObject: EmptyDatasetAccess, OperationSelector: dcl.TriggersOperation("updateDatasetPatchDatasetOperation")}, fn.AddNest("Access")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreationTime, actual.CreationTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreationTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LastModifiedTime, actual.LastModifiedTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LastModifiedTime")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Published, actual.Published, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Published")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultEncryptionConfiguration, actual.DefaultEncryptionConfiguration, dcl.DiffInfo{ObjectFunction: compareDatasetDefaultEncryptionConfigurationNewStyle, EmptyObject: EmptyDatasetDefaultEncryptionConfiguration, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DefaultEncryptionConfiguration")); len(ds) != 0 || err != nil {
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
func compareDatasetAccessNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccess)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccess or *DatasetAccess", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccess)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccess)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccess", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Role, actual.Role, dcl.DiffInfo{CustomDiff: canonicalizeDatasetAccessRole, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Role")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UserByEmail, actual.UserByEmail, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UserByEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupByEmail, actual.GroupByEmail, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupByEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Domain, actual.Domain, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Domain")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SpecialGroup, actual.SpecialGroup, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SpecialGroup")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IamMember, actual.IamMember, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IamMember")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.View, actual.View, dcl.DiffInfo{ObjectFunction: compareDatasetAccessViewNewStyle, EmptyObject: EmptyDatasetAccessView, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("View")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Routine, actual.Routine, dcl.DiffInfo{ObjectFunction: compareDatasetAccessRoutineNewStyle, EmptyObject: EmptyDatasetAccessRoutine, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Routine")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetAccessViewNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccessView)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccessView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessView or *DatasetAccessView", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccessView)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccessView)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessView", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableId, actual.TableId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TableId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetAccessRoutineNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetAccessRoutine)
	if !ok {
		desiredNotPointer, ok := d.(DatasetAccessRoutine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessRoutine or *DatasetAccessRoutine", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetAccessRoutine)
	if !ok {
		actualNotPointer, ok := a.(DatasetAccessRoutine)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetAccessRoutine", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RoutineId, actual.RoutineId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("RoutineId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareDatasetDefaultEncryptionConfigurationNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*DatasetDefaultEncryptionConfiguration)
	if !ok {
		desiredNotPointer, ok := d.(DatasetDefaultEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetDefaultEncryptionConfiguration or *DatasetDefaultEncryptionConfiguration", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*DatasetDefaultEncryptionConfiguration)
	if !ok {
		actualNotPointer, ok := a.(DatasetDefaultEncryptionConfiguration)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a DatasetDefaultEncryptionConfiguration", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.KmsKeyName, actual.KmsKeyName, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("KmsKeyName")); len(ds) != 0 || err != nil {
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
func (r *Dataset) urlNormalized() *Dataset {
	normalized := dcl.Copy(*r).(Dataset)
	normalized.Etag = dcl.SelfLinkToName(r.Etag)
	normalized.Id = dcl.SelfLinkToName(r.Id)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.FriendlyName = dcl.SelfLinkToName(r.FriendlyName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.DefaultTableExpirationMs = dcl.SelfLinkToName(r.DefaultTableExpirationMs)
	normalized.DefaultPartitionExpirationMs = dcl.SelfLinkToName(r.DefaultPartitionExpirationMs)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Dataset) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "PatchDataset" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
			"name":    dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/datasets/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Dataset resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Dataset) marshal(c *Client) ([]byte, error) {
	m, err := expandDataset(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Dataset: %w", err)
	}
	dcl.MoveMapEntry(
		m,
		[]string{"name"},
		[]string{"datasetReference", "datasetId"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"project"},
		[]string{"datasetReference", "projectId"},
	)

	return json.Marshal(m)
}

// unmarshalDataset decodes JSON responses into the Dataset resource schema.
func unmarshalDataset(b []byte, c *Client, res *Dataset) (*Dataset, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapDataset(m, c, res)
}

func unmarshalMapDataset(m map[string]interface{}, c *Client, res *Dataset) (*Dataset, error) {
	dcl.MoveMapEntry(
		m,
		[]string{"datasetReference", "datasetId"},
		[]string{"name"},
	)
	dcl.MoveMapEntry(
		m,
		[]string{"datasetReference", "projectId"},
		[]string{"project"},
	)

	flattened := flattenDataset(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandDataset expands Dataset into a JSON request object.
func expandDataset(c *Client, f *Dataset) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v := f.Name; dcl.ValueShouldBeSent(v) {
		m["name"] = v
	}
	if v := f.Project; dcl.ValueShouldBeSent(v) {
		m["project"] = v
	}
	if v := f.FriendlyName; dcl.ValueShouldBeSent(v) {
		m["friendlyName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.DefaultTableExpirationMs; dcl.ValueShouldBeSent(v) {
		m["defaultTableExpirationMs"] = v
	}
	if v := f.DefaultPartitionExpirationMs; dcl.ValueShouldBeSent(v) {
		m["defaultPartitionExpirationMs"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandDatasetAccessSlice(c, f.Access, res); err != nil {
		return nil, fmt.Errorf("error expanding Access into access: %w", err)
	} else if v != nil {
		m["access"] = v
	}
	if v := f.Location; dcl.ValueShouldBeSent(v) {
		m["location"] = v
	}
	if v := f.Published; dcl.ValueShouldBeSent(v) {
		m["published"] = v
	}
	if v, err := expandDatasetDefaultEncryptionConfiguration(c, f.DefaultEncryptionConfiguration, res); err != nil {
		return nil, fmt.Errorf("error expanding DefaultEncryptionConfiguration into defaultEncryptionConfiguration: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["defaultEncryptionConfiguration"] = v
	}

	return m, nil
}

// flattenDataset flattens Dataset from a JSON request object into the
// Dataset type.
func flattenDataset(c *Client, i interface{}, res *Dataset) *Dataset {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Dataset{}
	resultRes.Etag = dcl.FlattenString(m["etag"])
	resultRes.Id = dcl.FlattenString(m["id"])
	resultRes.SelfLink = dcl.FlattenString(m["selfLink"])
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.FriendlyName = dcl.FlattenString(m["friendlyName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.DefaultTableExpirationMs = dcl.FlattenString(m["defaultTableExpirationMs"])
	resultRes.DefaultPartitionExpirationMs = dcl.FlattenString(m["defaultPartitionExpirationMs"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.Access = flattenDatasetAccessSlice(c, m["access"], res)
	resultRes.CreationTime = dcl.FlattenInteger(m["creationTime"])
	resultRes.LastModifiedTime = dcl.FlattenInteger(m["lastModifiedTime"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Published = dcl.FlattenBool(m["published"])
	resultRes.DefaultEncryptionConfiguration = flattenDatasetDefaultEncryptionConfiguration(c, m["defaultEncryptionConfiguration"], res)

	return resultRes
}

// expandDatasetAccessMap expands the contents of DatasetAccess into a JSON
// request object.
func expandDatasetAccessMap(c *Client, f map[string]DatasetAccess, res *Dataset) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccess(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessSlice expands the contents of DatasetAccess into a JSON
// request object.
func expandDatasetAccessSlice(c *Client, f []DatasetAccess, res *Dataset) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccess(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessMap flattens the contents of DatasetAccess from a JSON
// response object.
func flattenDatasetAccessMap(c *Client, i interface{}, res *Dataset) map[string]DatasetAccess {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccess{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccess{}
	}

	items := make(map[string]DatasetAccess)
	for k, item := range a {
		items[k] = *flattenDatasetAccess(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDatasetAccessSlice flattens the contents of DatasetAccess from a JSON
// response object.
func flattenDatasetAccessSlice(c *Client, i interface{}, res *Dataset) []DatasetAccess {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccess{}
	}

	if len(a) == 0 {
		return []DatasetAccess{}
	}

	items := make([]DatasetAccess, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccess(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDatasetAccess expands an instance of DatasetAccess into a JSON
// request object.
func expandDatasetAccess(c *Client, f *DatasetAccess, res *Dataset) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Role; !dcl.IsEmptyValueIndirect(v) {
		m["role"] = v
	}
	if v := f.UserByEmail; !dcl.IsEmptyValueIndirect(v) {
		m["userByEmail"] = v
	}
	if v := f.GroupByEmail; !dcl.IsEmptyValueIndirect(v) {
		m["groupByEmail"] = v
	}
	if v := f.Domain; !dcl.IsEmptyValueIndirect(v) {
		m["domain"] = v
	}
	if v := f.SpecialGroup; !dcl.IsEmptyValueIndirect(v) {
		m["specialGroup"] = v
	}
	if v := f.IamMember; !dcl.IsEmptyValueIndirect(v) {
		m["iamMember"] = v
	}
	if v, err := expandDatasetAccessView(c, f.View, res); err != nil {
		return nil, fmt.Errorf("error expanding View into view: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["view"] = v
	}
	if v, err := expandDatasetAccessRoutine(c, f.Routine, res); err != nil {
		return nil, fmt.Errorf("error expanding Routine into routine: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["routine"] = v
	}

	return m, nil
}

// flattenDatasetAccess flattens an instance of DatasetAccess from a JSON
// response object.
func flattenDatasetAccess(c *Client, i interface{}, res *Dataset) *DatasetAccess {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccess{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccess
	}
	r.Role = dcl.FlattenString(m["role"])
	r.UserByEmail = dcl.FlattenString(m["userByEmail"])
	r.GroupByEmail = dcl.FlattenString(m["groupByEmail"])
	r.Domain = dcl.FlattenString(m["domain"])
	r.SpecialGroup = dcl.FlattenString(m["specialGroup"])
	r.IamMember = dcl.FlattenString(m["iamMember"])
	r.View = flattenDatasetAccessView(c, m["view"], res)
	r.Routine = flattenDatasetAccessRoutine(c, m["routine"], res)

	return r
}

// expandDatasetAccessViewMap expands the contents of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessViewMap(c *Client, f map[string]DatasetAccessView, res *Dataset) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccessView(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessViewSlice expands the contents of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessViewSlice(c *Client, f []DatasetAccessView, res *Dataset) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccessView(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessViewMap flattens the contents of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessViewMap(c *Client, i interface{}, res *Dataset) map[string]DatasetAccessView {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccessView{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccessView{}
	}

	items := make(map[string]DatasetAccessView)
	for k, item := range a {
		items[k] = *flattenDatasetAccessView(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDatasetAccessViewSlice flattens the contents of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessViewSlice(c *Client, i interface{}, res *Dataset) []DatasetAccessView {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccessView{}
	}

	if len(a) == 0 {
		return []DatasetAccessView{}
	}

	items := make([]DatasetAccessView, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccessView(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDatasetAccessView expands an instance of DatasetAccessView into a JSON
// request object.
func expandDatasetAccessView(c *Client, f *DatasetAccessView, res *Dataset) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.DatasetId; !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v := f.TableId; !dcl.IsEmptyValueIndirect(v) {
		m["tableId"] = v
	}

	return m, nil
}

// flattenDatasetAccessView flattens an instance of DatasetAccessView from a JSON
// response object.
func flattenDatasetAccessView(c *Client, i interface{}, res *Dataset) *DatasetAccessView {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccessView{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccessView
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.TableId = dcl.FlattenString(m["tableId"])

	return r
}

// expandDatasetAccessRoutineMap expands the contents of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutineMap(c *Client, f map[string]DatasetAccessRoutine, res *Dataset) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetAccessRoutine(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetAccessRoutineSlice expands the contents of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutineSlice(c *Client, f []DatasetAccessRoutine, res *Dataset) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetAccessRoutine(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetAccessRoutineMap flattens the contents of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutineMap(c *Client, i interface{}, res *Dataset) map[string]DatasetAccessRoutine {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetAccessRoutine{}
	}

	if len(a) == 0 {
		return map[string]DatasetAccessRoutine{}
	}

	items := make(map[string]DatasetAccessRoutine)
	for k, item := range a {
		items[k] = *flattenDatasetAccessRoutine(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDatasetAccessRoutineSlice flattens the contents of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutineSlice(c *Client, i interface{}, res *Dataset) []DatasetAccessRoutine {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetAccessRoutine{}
	}

	if len(a) == 0 {
		return []DatasetAccessRoutine{}
	}

	items := make([]DatasetAccessRoutine, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetAccessRoutine(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDatasetAccessRoutine expands an instance of DatasetAccessRoutine into a JSON
// request object.
func expandDatasetAccessRoutine(c *Client, f *DatasetAccessRoutine, res *Dataset) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ProjectId; !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v := f.DatasetId; !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v := f.RoutineId; !dcl.IsEmptyValueIndirect(v) {
		m["routineId"] = v
	}

	return m, nil
}

// flattenDatasetAccessRoutine flattens an instance of DatasetAccessRoutine from a JSON
// response object.
func flattenDatasetAccessRoutine(c *Client, i interface{}, res *Dataset) *DatasetAccessRoutine {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetAccessRoutine{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetAccessRoutine
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.RoutineId = dcl.FlattenString(m["routineId"])

	return r
}

// expandDatasetDefaultEncryptionConfigurationMap expands the contents of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfigurationMap(c *Client, f map[string]DatasetDefaultEncryptionConfiguration, res *Dataset) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandDatasetDefaultEncryptionConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandDatasetDefaultEncryptionConfigurationSlice expands the contents of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfigurationSlice(c *Client, f []DatasetDefaultEncryptionConfiguration, res *Dataset) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandDatasetDefaultEncryptionConfiguration(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenDatasetDefaultEncryptionConfigurationMap flattens the contents of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfigurationMap(c *Client, i interface{}, res *Dataset) map[string]DatasetDefaultEncryptionConfiguration {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]DatasetDefaultEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return map[string]DatasetDefaultEncryptionConfiguration{}
	}

	items := make(map[string]DatasetDefaultEncryptionConfiguration)
	for k, item := range a {
		items[k] = *flattenDatasetDefaultEncryptionConfiguration(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenDatasetDefaultEncryptionConfigurationSlice flattens the contents of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfigurationSlice(c *Client, i interface{}, res *Dataset) []DatasetDefaultEncryptionConfiguration {
	a, ok := i.([]interface{})
	if !ok {
		return []DatasetDefaultEncryptionConfiguration{}
	}

	if len(a) == 0 {
		return []DatasetDefaultEncryptionConfiguration{}
	}

	items := make([]DatasetDefaultEncryptionConfiguration, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenDatasetDefaultEncryptionConfiguration(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandDatasetDefaultEncryptionConfiguration expands an instance of DatasetDefaultEncryptionConfiguration into a JSON
// request object.
func expandDatasetDefaultEncryptionConfiguration(c *Client, f *DatasetDefaultEncryptionConfiguration, res *Dataset) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.KmsKeyName; !dcl.IsEmptyValueIndirect(v) {
		m["kmsKeyName"] = v
	}

	return m, nil
}

// flattenDatasetDefaultEncryptionConfiguration flattens an instance of DatasetDefaultEncryptionConfiguration from a JSON
// response object.
func flattenDatasetDefaultEncryptionConfiguration(c *Client, i interface{}, res *Dataset) *DatasetDefaultEncryptionConfiguration {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &DatasetDefaultEncryptionConfiguration{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyDatasetDefaultEncryptionConfiguration
	}
	r.KmsKeyName = dcl.FlattenString(m["kmsKeyName"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Dataset) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalDataset(b, c, r)
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

type datasetDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         datasetApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToDatasetDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]datasetDiff, error) {
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
	var diffs []datasetDiff
	// For each operation name, create a datasetDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := datasetDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToDatasetApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToDatasetApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (datasetApiOperation, error) {
	switch opName {

	case "updateDatasetPatchDatasetOperation":
		return &updateDatasetPatchDatasetOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractDatasetFields(r *Dataset) error {
	vDefaultEncryptionConfiguration := r.DefaultEncryptionConfiguration
	if vDefaultEncryptionConfiguration == nil {
		// note: explicitly not the empty object.
		vDefaultEncryptionConfiguration = &DatasetDefaultEncryptionConfiguration{}
	}
	if err := extractDatasetDefaultEncryptionConfigurationFields(r, vDefaultEncryptionConfiguration); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultEncryptionConfiguration) {
		r.DefaultEncryptionConfiguration = vDefaultEncryptionConfiguration
	}
	return nil
}
func extractDatasetAccessFields(r *Dataset, o *DatasetAccess) error {
	vView := o.View
	if vView == nil {
		// note: explicitly not the empty object.
		vView = &DatasetAccessView{}
	}
	if err := extractDatasetAccessViewFields(r, vView); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vView) {
		o.View = vView
	}
	vRoutine := o.Routine
	if vRoutine == nil {
		// note: explicitly not the empty object.
		vRoutine = &DatasetAccessRoutine{}
	}
	if err := extractDatasetAccessRoutineFields(r, vRoutine); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRoutine) {
		o.Routine = vRoutine
	}
	return nil
}
func extractDatasetAccessViewFields(r *Dataset, o *DatasetAccessView) error {
	return nil
}
func extractDatasetAccessRoutineFields(r *Dataset, o *DatasetAccessRoutine) error {
	return nil
}
func extractDatasetDefaultEncryptionConfigurationFields(r *Dataset, o *DatasetDefaultEncryptionConfiguration) error {
	return nil
}

func postReadExtractDatasetFields(r *Dataset) error {
	vDefaultEncryptionConfiguration := r.DefaultEncryptionConfiguration
	if vDefaultEncryptionConfiguration == nil {
		// note: explicitly not the empty object.
		vDefaultEncryptionConfiguration = &DatasetDefaultEncryptionConfiguration{}
	}
	if err := postReadExtractDatasetDefaultEncryptionConfigurationFields(r, vDefaultEncryptionConfiguration); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDefaultEncryptionConfiguration) {
		r.DefaultEncryptionConfiguration = vDefaultEncryptionConfiguration
	}
	return nil
}
func postReadExtractDatasetAccessFields(r *Dataset, o *DatasetAccess) error {
	vView := o.View
	if vView == nil {
		// note: explicitly not the empty object.
		vView = &DatasetAccessView{}
	}
	if err := extractDatasetAccessViewFields(r, vView); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vView) {
		o.View = vView
	}
	vRoutine := o.Routine
	if vRoutine == nil {
		// note: explicitly not the empty object.
		vRoutine = &DatasetAccessRoutine{}
	}
	if err := extractDatasetAccessRoutineFields(r, vRoutine); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRoutine) {
		o.Routine = vRoutine
	}
	return nil
}
func postReadExtractDatasetAccessViewFields(r *Dataset, o *DatasetAccessView) error {
	return nil
}
func postReadExtractDatasetAccessRoutineFields(r *Dataset, o *DatasetAccessRoutine) error {
	return nil
}
func postReadExtractDatasetDefaultEncryptionConfigurationFields(r *Dataset, o *DatasetDefaultEncryptionConfiguration) error {
	return nil
}
