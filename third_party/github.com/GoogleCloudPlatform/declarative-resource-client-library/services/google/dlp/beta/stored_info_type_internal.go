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
)

func (r *StoredInfoType) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"LargeCustomDictionary", "Dictionary", "Regex"}, r.LargeCustomDictionary, r.Dictionary, r.Regex); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Parent, "Parent"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.LargeCustomDictionary) {
		if err := r.LargeCustomDictionary.validate(); err != nil {
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
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionary) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"CloudStorageFileSet", "BigQueryField"}, r.CloudStorageFileSet, r.BigQueryField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.OutputPath) {
		if err := r.OutputPath.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.CloudStorageFileSet) {
		if err := r.CloudStorageFileSet.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.BigQueryField) {
		if err := r.BigQueryField.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionaryOutputPath) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) validate() error {
	if err := dcl.Required(r, "url"); err != nil {
		return err
	}
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionaryBigQueryField) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Table) {
		if err := r.Table.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Field) {
		if err := r.Field.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) validate() error {
	return nil
}
func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) validate() error {
	return nil
}
func (r *StoredInfoTypeDictionary) validate() error {
	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"WordList", "CloudStoragePath"}, r.WordList, r.CloudStoragePath); err != nil {
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
func (r *StoredInfoTypeDictionaryWordList) validate() error {
	if err := dcl.Required(r, "words"); err != nil {
		return err
	}
	return nil
}
func (r *StoredInfoTypeDictionaryCloudStoragePath) validate() error {
	if err := dcl.Required(r, "path"); err != nil {
		return err
	}
	return nil
}
func (r *StoredInfoTypeRegex) validate() error {
	if err := dcl.Required(r, "pattern"); err != nil {
		return err
	}
	return nil
}
func (r *StoredInfoType) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://dlp.googleapis.com/v2/", params)
}

func (r *StoredInfoType) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *StoredInfoType) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/storedInfoTypes", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/storedInfoTypes", nr.basePath(), userBasePath, params), nil

}

func (r *StoredInfoType) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/storedInfoTypes", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/storedInfoTypes", nr.basePath(), userBasePath, params), nil

}

func (r *StoredInfoType) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"location": dcl.ValueOrEmptyString(nr.Location),
		"parent":   dcl.ValueOrEmptyString(nr.Parent),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.URL("{{parent}}/locations/{{location}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, params), nil
	}

	return dcl.URL("{{parent}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, params), nil
}

// storedInfoTypeApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type storedInfoTypeApiOperation interface {
	do(context.Context, *StoredInfoType, *Client) error
}

// newUpdateStoredInfoTypeUpdateStoredInfoTypeRequest creates a request for an
// StoredInfoType resource's UpdateStoredInfoType update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateStoredInfoTypeUpdateStoredInfoTypeRequest(ctx context.Context, f *StoredInfoType, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v, err := expandStoredInfoTypeLargeCustomDictionary(c, f.LargeCustomDictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding LargeCustomDictionary into largeCustomDictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["largeCustomDictionary"] = v
	}
	if v, err := expandStoredInfoTypeDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["dictionary"] = v
	}
	if v, err := expandStoredInfoTypeRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["regex"] = v
	}
	return req, nil
}

// marshalUpdateStoredInfoTypeUpdateStoredInfoTypeRequest converts the update into
// the final JSON request body.
func marshalUpdateStoredInfoTypeUpdateStoredInfoTypeRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateStoredInfoTypeUpdateStoredInfoTypeOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) listStoredInfoTypeRaw(ctx context.Context, r *StoredInfoType, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != StoredInfoTypeMaxPage {
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

type listStoredInfoTypeOperation struct {
	StoredInfoTypes []map[string]interface{} `json:"storedInfoTypes"`
	Token           string                   `json:"nextPageToken"`
}

func (c *Client) listStoredInfoType(ctx context.Context, r *StoredInfoType, pageToken string, pageSize int32) ([]*StoredInfoType, string, error) {
	b, err := c.listStoredInfoTypeRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listStoredInfoTypeOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*StoredInfoType
	for _, v := range m.StoredInfoTypes {
		res, err := unmarshalMapStoredInfoType(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Location = r.Location
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllStoredInfoType(ctx context.Context, f func(*StoredInfoType) bool, resources []*StoredInfoType) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteStoredInfoType(ctx, res)
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

type deleteStoredInfoTypeOperation struct{}

func (op *deleteStoredInfoTypeOperation) do(ctx context.Context, r *StoredInfoType, c *Client) error {
	r, err := c.GetStoredInfoType(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "StoredInfoType not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetStoredInfoType checking for existence. error: %v", err)
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
		return fmt.Errorf("failed to delete StoredInfoType: %w", err)
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createStoredInfoTypeOperation struct {
	response map[string]interface{}
}

func (op *createStoredInfoTypeOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createStoredInfoTypeOperation) do(ctx context.Context, r *StoredInfoType, c *Client) error {
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
		// Allowing creation to continue with Name set could result in a StoredInfoType with the wrong Name.
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

	if _, err := c.GetStoredInfoType(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getStoredInfoTypeRaw(ctx context.Context, r *StoredInfoType) ([]byte, error) {

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

func (c *Client) storedInfoTypeDiffsForRawDesired(ctx context.Context, rawDesired *StoredInfoType, opts ...dcl.ApplyOption) (initial, desired *StoredInfoType, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *StoredInfoType
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*StoredInfoType); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected StoredInfoType, got %T", sh)
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
		desired, err := canonicalizeStoredInfoTypeDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetStoredInfoType(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a StoredInfoType resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve StoredInfoType resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that StoredInfoType resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeStoredInfoTypeDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for StoredInfoType: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for StoredInfoType: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractStoredInfoTypeFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeStoredInfoTypeInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for StoredInfoType: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeStoredInfoTypeDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for StoredInfoType: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffStoredInfoType(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func storedInfoTypeFromParent(r *StoredInfoType, parent map[string]interface{}) (map[string]interface{}, error) {
	var err error
	for _, path := range [][]string{[]string{"currentVersion", "config"}, []string{"pendingVersions[0]", "config"}} {
		// Try each possible path where the resource could be.
		itemAtQuery, err := dcl.GetMapEntry(parent, path)
		if err != nil {
			// Resource not found at path, look at next path.
			continue
		}
		return itemAtQuery.(map[string]interface{}), nil
	}
	return nil, dcl.NotFoundError{Cause: err}
}

func canonicalizeStoredInfoTypeInitialState(rawInitial, rawDesired *StoredInfoType) (*StoredInfoType, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.LargeCustomDictionary) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.Dictionary, rawInitial.Regex) {
			rawInitial.LargeCustomDictionary = EmptyStoredInfoTypeLargeCustomDictionary
		}
	}

	if !dcl.IsZeroValue(rawInitial.Dictionary) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.LargeCustomDictionary, rawInitial.Regex) {
			rawInitial.Dictionary = EmptyStoredInfoTypeDictionary
		}
	}

	if !dcl.IsZeroValue(rawInitial.Regex) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.LargeCustomDictionary, rawInitial.Dictionary) {
			rawInitial.Regex = EmptyStoredInfoTypeRegex
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

func canonicalizeStoredInfoTypeDesiredState(rawDesired, rawInitial *StoredInfoType, opts ...dcl.ApplyOption) (*StoredInfoType, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.LargeCustomDictionary = canonicalizeStoredInfoTypeLargeCustomDictionary(rawDesired.LargeCustomDictionary, nil, opts...)
		rawDesired.Dictionary = canonicalizeStoredInfoTypeDictionary(rawDesired.Dictionary, nil, opts...)
		rawDesired.Regex = canonicalizeStoredInfoTypeRegex(rawDesired.Regex, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &StoredInfoType{}
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
	canonicalDesired.LargeCustomDictionary = canonicalizeStoredInfoTypeLargeCustomDictionary(rawDesired.LargeCustomDictionary, rawInitial.LargeCustomDictionary, opts...)
	canonicalDesired.Dictionary = canonicalizeStoredInfoTypeDictionary(rawDesired.Dictionary, rawInitial.Dictionary, opts...)
	canonicalDesired.Regex = canonicalizeStoredInfoTypeRegex(rawDesired.Regex, rawInitial.Regex, opts...)
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

	if canonicalDesired.LargeCustomDictionary != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.Dictionary, rawDesired.Regex) {
			canonicalDesired.LargeCustomDictionary = EmptyStoredInfoTypeLargeCustomDictionary
		}
	}

	if canonicalDesired.Dictionary != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.LargeCustomDictionary, rawDesired.Regex) {
			canonicalDesired.Dictionary = EmptyStoredInfoTypeDictionary
		}
	}

	if canonicalDesired.Regex != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.LargeCustomDictionary, rawDesired.Dictionary) {
			canonicalDesired.Regex = EmptyStoredInfoTypeRegex
		}
	}

	return canonicalDesired, nil
}

func canonicalizeStoredInfoTypeNewState(c *Client, rawNew, rawDesired *StoredInfoType) (*StoredInfoType, error) {

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

	if dcl.IsEmptyValueIndirect(rawNew.LargeCustomDictionary) && dcl.IsEmptyValueIndirect(rawDesired.LargeCustomDictionary) {
		rawNew.LargeCustomDictionary = rawDesired.LargeCustomDictionary
	} else {
		rawNew.LargeCustomDictionary = canonicalizeNewStoredInfoTypeLargeCustomDictionary(c, rawDesired.LargeCustomDictionary, rawNew.LargeCustomDictionary)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Dictionary) && dcl.IsEmptyValueIndirect(rawDesired.Dictionary) {
		rawNew.Dictionary = rawDesired.Dictionary
	} else {
		rawNew.Dictionary = canonicalizeNewStoredInfoTypeDictionary(c, rawDesired.Dictionary, rawNew.Dictionary)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Regex) && dcl.IsEmptyValueIndirect(rawDesired.Regex) {
		rawNew.Regex = rawDesired.Regex
	} else {
		rawNew.Regex = canonicalizeNewStoredInfoTypeRegex(c, rawDesired.Regex, rawNew.Regex)
	}

	rawNew.Parent = rawDesired.Parent

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeStoredInfoTypeLargeCustomDictionary(des, initial *StoredInfoTypeLargeCustomDictionary, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if des.CloudStorageFileSet != nil || (initial != nil && initial.CloudStorageFileSet != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.BigQueryField) {
			des.CloudStorageFileSet = nil
			if initial != nil {
				initial.CloudStorageFileSet = nil
			}
		}
	}

	if des.BigQueryField != nil || (initial != nil && initial.BigQueryField != nil) {
		// Check if anything else is set.
		if dcl.AnySet(des.CloudStorageFileSet) {
			des.BigQueryField = nil
			if initial != nil {
				initial.BigQueryField = nil
			}
		}
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionary{}

	cDes.OutputPath = canonicalizeStoredInfoTypeLargeCustomDictionaryOutputPath(des.OutputPath, initial.OutputPath, opts...)
	cDes.CloudStorageFileSet = canonicalizeStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(des.CloudStorageFileSet, initial.CloudStorageFileSet, opts...)
	cDes.BigQueryField = canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryField(des.BigQueryField, initial.BigQueryField, opts...)

	return cDes
}

func canonicalizeStoredInfoTypeLargeCustomDictionarySlice(des, initial []StoredInfoTypeLargeCustomDictionary, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionary(c *Client, des, nw *StoredInfoTypeLargeCustomDictionary) *StoredInfoTypeLargeCustomDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.OutputPath = canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPath(c, des.OutputPath, nw.OutputPath)
	nw.CloudStorageFileSet = canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, des.CloudStorageFileSet, nw.CloudStorageFileSet)
	nw.BigQueryField = canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryField(c, des.BigQueryField, nw.BigQueryField)

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionarySet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionary) []StoredInfoTypeLargeCustomDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionarySlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionary) []StoredInfoTypeLargeCustomDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryOutputPath(des, initial *StoredInfoTypeLargeCustomDictionaryOutputPath, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionaryOutputPath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionaryOutputPath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryOutputPathSlice(des, initial []StoredInfoTypeLargeCustomDictionaryOutputPath, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionaryOutputPath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionaryOutputPath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionaryOutputPath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryOutputPath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionaryOutputPath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPath(c *Client, des, nw *StoredInfoTypeLargeCustomDictionaryOutputPath) *StoredInfoTypeLargeCustomDictionaryOutputPath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionaryOutputPath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPathSet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryOutputPath) []StoredInfoTypeLargeCustomDictionaryOutputPath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionaryOutputPath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryOutputPathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPathSlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryOutputPath) []StoredInfoTypeLargeCustomDictionaryOutputPath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionaryOutputPath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryOutputPath(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(des, initial *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}

	if dcl.StringCanonicalize(des.Url, initial.Url) || dcl.IsZeroValue(des.Url) {
		cDes.Url = initial.Url
	} else {
		cDes.Url = des.Url
	}

	return cDes
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice(des, initial []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c *Client, des, nw *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Url, nw.Url) {
		nw.Url = des.Url
	}

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryField(des, initial *StoredInfoTypeLargeCustomDictionaryBigQueryField, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionaryBigQueryField{}

	cDes.Table = canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(des.Table, initial.Table, opts...)
	cDes.Field = canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(des.Field, initial.Field, opts...)

	return cDes
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice(des, initial []StoredInfoTypeLargeCustomDictionaryBigQueryField, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryField, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryField(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryField, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryField(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryField(c *Client, des, nw *StoredInfoTypeLargeCustomDictionaryBigQueryField) *StoredInfoTypeLargeCustomDictionaryBigQueryField {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionaryBigQueryField while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Table = canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, des.Table, nw.Table)
	nw.Field = canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, des.Field, nw.Field)

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldSet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryField) []StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionaryBigQueryField
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryField(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryField) []StoredInfoTypeLargeCustomDictionaryBigQueryField {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionaryBigQueryField
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryField(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(des, initial *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}

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

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice(des, initial []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c *Client, des, nw *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(des, initial *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, opts ...dcl.ApplyOption) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}

	return cDes
}

func canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice(des, initial []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, opts ...dcl.ApplyOption) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c *Client, des, nw *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeLargeCustomDictionaryBigQueryFieldField while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}

	return nw
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSet(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice(c *Client, des, nw []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeDictionary(des, initial *StoredInfoTypeDictionary, opts ...dcl.ApplyOption) *StoredInfoTypeDictionary {
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

	cDes := &StoredInfoTypeDictionary{}

	cDes.WordList = canonicalizeStoredInfoTypeDictionaryWordList(des.WordList, initial.WordList, opts...)
	cDes.CloudStoragePath = canonicalizeStoredInfoTypeDictionaryCloudStoragePath(des.CloudStoragePath, initial.CloudStoragePath, opts...)

	return cDes
}

func canonicalizeStoredInfoTypeDictionarySlice(des, initial []StoredInfoTypeDictionary, opts ...dcl.ApplyOption) []StoredInfoTypeDictionary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeDictionary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeDictionary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeDictionary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeDictionary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeDictionary(c *Client, des, nw *StoredInfoTypeDictionary) *StoredInfoTypeDictionary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeDictionary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.WordList = canonicalizeNewStoredInfoTypeDictionaryWordList(c, des.WordList, nw.WordList)
	nw.CloudStoragePath = canonicalizeNewStoredInfoTypeDictionaryCloudStoragePath(c, des.CloudStoragePath, nw.CloudStoragePath)

	return nw
}

func canonicalizeNewStoredInfoTypeDictionarySet(c *Client, des, nw []StoredInfoTypeDictionary) []StoredInfoTypeDictionary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeDictionary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeDictionaryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeDictionary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeDictionarySlice(c *Client, des, nw []StoredInfoTypeDictionary) []StoredInfoTypeDictionary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeDictionary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeDictionary(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeDictionaryWordList(des, initial *StoredInfoTypeDictionaryWordList, opts ...dcl.ApplyOption) *StoredInfoTypeDictionaryWordList {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeDictionaryWordList{}

	if dcl.StringArrayCanonicalize(des.Words, initial.Words) {
		cDes.Words = initial.Words
	} else {
		cDes.Words = des.Words
	}

	return cDes
}

func canonicalizeStoredInfoTypeDictionaryWordListSlice(des, initial []StoredInfoTypeDictionaryWordList, opts ...dcl.ApplyOption) []StoredInfoTypeDictionaryWordList {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeDictionaryWordList, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeDictionaryWordList(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeDictionaryWordList, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeDictionaryWordList(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeDictionaryWordList(c *Client, des, nw *StoredInfoTypeDictionaryWordList) *StoredInfoTypeDictionaryWordList {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeDictionaryWordList while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.Words, nw.Words) {
		nw.Words = des.Words
	}

	return nw
}

func canonicalizeNewStoredInfoTypeDictionaryWordListSet(c *Client, des, nw []StoredInfoTypeDictionaryWordList) []StoredInfoTypeDictionaryWordList {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeDictionaryWordList
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeDictionaryWordListNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeDictionaryWordList(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeDictionaryWordListSlice(c *Client, des, nw []StoredInfoTypeDictionaryWordList) []StoredInfoTypeDictionaryWordList {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeDictionaryWordList
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeDictionaryWordList(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeDictionaryCloudStoragePath(des, initial *StoredInfoTypeDictionaryCloudStoragePath, opts ...dcl.ApplyOption) *StoredInfoTypeDictionaryCloudStoragePath {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeDictionaryCloudStoragePath{}

	if dcl.StringCanonicalize(des.Path, initial.Path) || dcl.IsZeroValue(des.Path) {
		cDes.Path = initial.Path
	} else {
		cDes.Path = des.Path
	}

	return cDes
}

func canonicalizeStoredInfoTypeDictionaryCloudStoragePathSlice(des, initial []StoredInfoTypeDictionaryCloudStoragePath, opts ...dcl.ApplyOption) []StoredInfoTypeDictionaryCloudStoragePath {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeDictionaryCloudStoragePath, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeDictionaryCloudStoragePath(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeDictionaryCloudStoragePath, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeDictionaryCloudStoragePath(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeDictionaryCloudStoragePath(c *Client, des, nw *StoredInfoTypeDictionaryCloudStoragePath) *StoredInfoTypeDictionaryCloudStoragePath {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeDictionaryCloudStoragePath while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Path, nw.Path) {
		nw.Path = des.Path
	}

	return nw
}

func canonicalizeNewStoredInfoTypeDictionaryCloudStoragePathSet(c *Client, des, nw []StoredInfoTypeDictionaryCloudStoragePath) []StoredInfoTypeDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeDictionaryCloudStoragePath
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeDictionaryCloudStoragePathNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeDictionaryCloudStoragePath(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeDictionaryCloudStoragePathSlice(c *Client, des, nw []StoredInfoTypeDictionaryCloudStoragePath) []StoredInfoTypeDictionaryCloudStoragePath {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeDictionaryCloudStoragePath
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeDictionaryCloudStoragePath(c, &d, &n))
	}

	return items
}

func canonicalizeStoredInfoTypeRegex(des, initial *StoredInfoTypeRegex, opts ...dcl.ApplyOption) *StoredInfoTypeRegex {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &StoredInfoTypeRegex{}

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

func canonicalizeStoredInfoTypeRegexSlice(des, initial []StoredInfoTypeRegex, opts ...dcl.ApplyOption) []StoredInfoTypeRegex {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]StoredInfoTypeRegex, 0, len(des))
		for _, d := range des {
			cd := canonicalizeStoredInfoTypeRegex(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]StoredInfoTypeRegex, 0, len(des))
	for i, d := range des {
		cd := canonicalizeStoredInfoTypeRegex(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewStoredInfoTypeRegex(c *Client, des, nw *StoredInfoTypeRegex) *StoredInfoTypeRegex {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for StoredInfoTypeRegex while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Pattern, nw.Pattern) {
		nw.Pattern = des.Pattern
	}

	return nw
}

func canonicalizeNewStoredInfoTypeRegexSet(c *Client, des, nw []StoredInfoTypeRegex) []StoredInfoTypeRegex {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []StoredInfoTypeRegex
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareStoredInfoTypeRegexNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewStoredInfoTypeRegex(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewStoredInfoTypeRegexSlice(c *Client, des, nw []StoredInfoTypeRegex) []StoredInfoTypeRegex {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []StoredInfoTypeRegex
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewStoredInfoTypeRegex(c, &d, &n))
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
func diffStoredInfoType(c *Client, desired, actual *StoredInfoType, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LargeCustomDictionary, actual.LargeCustomDictionary, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionary, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("LargeCustomDictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Dictionary, actual.Dictionary, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeDictionaryNewStyle, EmptyObject: EmptyStoredInfoTypeDictionary, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Dictionary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Regex, actual.Regex, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeRegexNewStyle, EmptyObject: EmptyStoredInfoTypeRegex, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Regex")); len(ds) != 0 || err != nil {
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
func compareStoredInfoTypeLargeCustomDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionary)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionary or *StoredInfoTypeLargeCustomDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionary)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.OutputPath, actual.OutputPath, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryOutputPathNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionaryOutputPath, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("OutputPath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStorageFileSet, actual.CloudStorageFileSet, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("CloudStorageFileSet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.BigQueryField, actual.BigQueryField, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionaryBigQueryField, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("BigQueryField")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeLargeCustomDictionaryOutputPathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionaryOutputPath)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionaryOutputPath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryOutputPath or *StoredInfoTypeLargeCustomDictionaryOutputPath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionaryOutputPath)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionaryOutputPath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryOutputPath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet or *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Url, actual.Url, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Url")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionaryBigQueryField)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionaryBigQueryField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryField or *StoredInfoTypeLargeCustomDictionaryBigQueryField", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionaryBigQueryField)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionaryBigQueryField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryField", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Table, actual.Table, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Table")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Field, actual.Field, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldNewStyle, EmptyObject: EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldField, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Field")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable or *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DatasetId, actual.DatasetId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("DatasetId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TableId, actual.TableId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("TableId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField or *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeLargeCustomDictionaryBigQueryFieldField", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeDictionaryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeDictionary)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionary or *StoredInfoTypeDictionary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeDictionary)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeDictionary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.WordList, actual.WordList, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeDictionaryWordListNewStyle, EmptyObject: EmptyStoredInfoTypeDictionaryWordList, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("WordList")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CloudStoragePath, actual.CloudStoragePath, dcl.DiffInfo{ObjectFunction: compareStoredInfoTypeDictionaryCloudStoragePathNewStyle, EmptyObject: EmptyStoredInfoTypeDictionaryCloudStoragePath, OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("CloudStoragePath")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeDictionaryWordListNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeDictionaryWordList)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionaryWordList or *StoredInfoTypeDictionaryWordList", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeDictionaryWordList)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeDictionaryWordList)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionaryWordList", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Words, actual.Words, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Words")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeDictionaryCloudStoragePathNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeDictionaryCloudStoragePath)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionaryCloudStoragePath or *StoredInfoTypeDictionaryCloudStoragePath", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeDictionaryCloudStoragePath)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeDictionaryCloudStoragePath)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeDictionaryCloudStoragePath", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Path, actual.Path, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Path")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareStoredInfoTypeRegexNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*StoredInfoTypeRegex)
	if !ok {
		desiredNotPointer, ok := d.(StoredInfoTypeRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeRegex or *StoredInfoTypeRegex", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*StoredInfoTypeRegex)
	if !ok {
		actualNotPointer, ok := a.(StoredInfoTypeRegex)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a StoredInfoTypeRegex", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Pattern, actual.Pattern, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("Pattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupIndexes, actual.GroupIndexes, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateStoredInfoTypeUpdateStoredInfoTypeOperation")}, fn.AddNest("GroupIndexes")); len(ds) != 0 || err != nil {
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
func (r *StoredInfoType) urlNormalized() *StoredInfoType {
	normalized := dcl.Copy(*r).(StoredInfoType)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.Parent = r.Parent
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *StoredInfoType) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateStoredInfoType" {
		fields := map[string]interface{}{
			"location": dcl.ValueOrEmptyString(nr.Location),
			"parent":   dcl.ValueOrEmptyString(nr.Parent),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		if dcl.IsRegion(nr.Location) {
			return dcl.URL("{{parent}}/locations/{{location}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, fields), nil
		}

		return dcl.URL("{{parent}}/storedInfoTypes/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the StoredInfoType resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *StoredInfoType) marshal(c *Client) ([]byte, error) {
	m, err := expandStoredInfoType(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling StoredInfoType: %w", err)
	}
	m = encodeStoredInfoTypeCreateRequest(m)

	return json.Marshal(m)
}

// unmarshalStoredInfoType decodes JSON responses into the StoredInfoType resource schema.
func unmarshalStoredInfoType(b []byte, c *Client, res *StoredInfoType) (*StoredInfoType, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapStoredInfoType(m, c, res)
}

func unmarshalMapStoredInfoType(m map[string]interface{}, c *Client, res *StoredInfoType) (*StoredInfoType, error) {

	flattened := flattenStoredInfoType(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandStoredInfoType expands StoredInfoType into a JSON request object.
func expandStoredInfoType(c *Client, f *StoredInfoType) (map[string]interface{}, error) {
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
	if v, err := expandStoredInfoTypeLargeCustomDictionary(c, f.LargeCustomDictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding LargeCustomDictionary into largeCustomDictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["largeCustomDictionary"] = v
	}
	if v, err := expandStoredInfoTypeDictionary(c, f.Dictionary, res); err != nil {
		return nil, fmt.Errorf("error expanding Dictionary into dictionary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["dictionary"] = v
	}
	if v, err := expandStoredInfoTypeRegex(c, f.Regex, res); err != nil {
		return nil, fmt.Errorf("error expanding Regex into regex: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["regex"] = v
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

// flattenStoredInfoType flattens StoredInfoType from a JSON request object into the
// StoredInfoType type.
func flattenStoredInfoType(c *Client, i interface{}, res *StoredInfoType) *StoredInfoType {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &StoredInfoType{}
	resultRes.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	resultRes.DisplayName = dcl.FlattenString(m["displayName"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.LargeCustomDictionary = flattenStoredInfoTypeLargeCustomDictionary(c, m["largeCustomDictionary"], res)
	resultRes.Dictionary = flattenStoredInfoTypeDictionary(c, m["dictionary"], res)
	resultRes.Regex = flattenStoredInfoTypeRegex(c, m["regex"], res)
	resultRes.Parent = dcl.FlattenString(m["parent"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandStoredInfoTypeLargeCustomDictionaryMap expands the contents of StoredInfoTypeLargeCustomDictionary into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionary, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionarySlice expands the contents of StoredInfoTypeLargeCustomDictionary into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionarySlice(c *Client, f []StoredInfoTypeLargeCustomDictionary, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryMap flattens the contents of StoredInfoTypeLargeCustomDictionary from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionary{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionary{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionary)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionarySlice flattens the contents of StoredInfoTypeLargeCustomDictionary from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionarySlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionary{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionary{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionary expands an instance of StoredInfoTypeLargeCustomDictionary into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionary(c *Client, f *StoredInfoTypeLargeCustomDictionary, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandStoredInfoTypeLargeCustomDictionaryOutputPath(c, f.OutputPath, res); err != nil {
		return nil, fmt.Errorf("error expanding OutputPath into outputPath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["outputPath"] = v
	}
	if v, err := expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, f.CloudStorageFileSet, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStorageFileSet into cloudStorageFileSet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStorageFileSet"] = v
	}
	if v, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryField(c, f.BigQueryField, res); err != nil {
		return nil, fmt.Errorf("error expanding BigQueryField into bigQueryField: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["bigQueryField"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionary flattens an instance of StoredInfoTypeLargeCustomDictionary from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionary(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionary
	}
	r.OutputPath = flattenStoredInfoTypeLargeCustomDictionaryOutputPath(c, m["outputPath"], res)
	r.CloudStorageFileSet = flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, m["cloudStorageFileSet"], res)
	r.BigQueryField = flattenStoredInfoTypeLargeCustomDictionaryBigQueryField(c, m["bigQueryField"], res)

	return r
}

// expandStoredInfoTypeLargeCustomDictionaryOutputPathMap expands the contents of StoredInfoTypeLargeCustomDictionaryOutputPath into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryOutputPathMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionaryOutputPath, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryOutputPath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionaryOutputPathSlice expands the contents of StoredInfoTypeLargeCustomDictionaryOutputPath into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryOutputPathSlice(c *Client, f []StoredInfoTypeLargeCustomDictionaryOutputPath, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryOutputPath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryOutputPathMap flattens the contents of StoredInfoTypeLargeCustomDictionaryOutputPath from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryOutputPathMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionaryOutputPath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionaryOutputPath)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionaryOutputPath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionaryOutputPathSlice flattens the contents of StoredInfoTypeLargeCustomDictionaryOutputPath from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryOutputPathSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionaryOutputPath {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryOutputPath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionaryOutputPath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionaryOutputPath expands an instance of StoredInfoTypeLargeCustomDictionaryOutputPath into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryOutputPath(c *Client, f *StoredInfoTypeLargeCustomDictionaryOutputPath, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryOutputPath flattens an instance of StoredInfoTypeLargeCustomDictionaryOutputPath from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryOutputPath(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionaryOutputPath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionaryOutputPath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionaryOutputPath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetMap expands the contents of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice expands the contents of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice(c *Client, f []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetMap flattens the contents of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice flattens the contents of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet expands an instance of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c *Client, f *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Url; !dcl.IsEmptyValueIndirect(v) {
		m["url"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet flattens an instance of StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	}
	r.Url = dcl.FlattenString(m["url"])

	return r
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldMap expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionaryBigQueryField, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryField(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice(c *Client, f []StoredInfoTypeLargeCustomDictionaryBigQueryField, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryField(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldMap flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionaryBigQueryField {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionaryBigQueryField)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionaryBigQueryField(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionaryBigQueryField {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryField, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionaryBigQueryField(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryField expands an instance of StoredInfoTypeLargeCustomDictionaryBigQueryField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryField(c *Client, f *StoredInfoTypeLargeCustomDictionaryBigQueryField, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, f.Table, res); err != nil {
		return nil, fmt.Errorf("error expanding Table into table: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["table"] = v
	}
	if v, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, f.Field, res); err != nil {
		return nil, fmt.Errorf("error expanding Field into field: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["field"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryField flattens an instance of StoredInfoTypeLargeCustomDictionaryBigQueryField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryField(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionaryBigQueryField {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionaryBigQueryField{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionaryBigQueryField
	}
	r.Table = flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, m["table"], res)
	r.Field = flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, m["field"], res)

	return r
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableMap expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice(c *Client, f []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableMap flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable expands an instance of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c *Client, f *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.ProjectId); err != nil {
		return nil, fmt.Errorf("error expanding ProjectId into projectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.DatasetId); err != nil {
		return nil, fmt.Errorf("error expanding DatasetId into datasetId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["datasetId"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.TableId); err != nil {
		return nil, fmt.Errorf("error expanding TableId into tableId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["tableId"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable flattens an instance of StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	}
	r.ProjectId = dcl.FlattenString(m["projectId"])
	r.DatasetId = dcl.FlattenString(m["datasetId"])
	r.TableId = dcl.FlattenString(m["tableId"])

	return r
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldMap expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldMap(c *Client, f map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice expands the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice(c *Client, f []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldMap flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}

	items := make(map[string]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice flattens the contents of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}

	items := make([]StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldField expands an instance of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField into a JSON
// request object.
func expandStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c *Client, f *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldField flattens an instance of StoredInfoTypeLargeCustomDictionaryBigQueryFieldField from a JSON
// response object.
func flattenStoredInfoTypeLargeCustomDictionaryBigQueryFieldField(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	}
	r.Name = dcl.FlattenString(m["name"])

	return r
}

// expandStoredInfoTypeDictionaryMap expands the contents of StoredInfoTypeDictionary into a JSON
// request object.
func expandStoredInfoTypeDictionaryMap(c *Client, f map[string]StoredInfoTypeDictionary, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeDictionarySlice expands the contents of StoredInfoTypeDictionary into a JSON
// request object.
func expandStoredInfoTypeDictionarySlice(c *Client, f []StoredInfoTypeDictionary, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeDictionary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeDictionaryMap flattens the contents of StoredInfoTypeDictionary from a JSON
// response object.
func flattenStoredInfoTypeDictionaryMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeDictionary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeDictionary{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeDictionary{}
	}

	items := make(map[string]StoredInfoTypeDictionary)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeDictionary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeDictionarySlice flattens the contents of StoredInfoTypeDictionary from a JSON
// response object.
func flattenStoredInfoTypeDictionarySlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeDictionary {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeDictionary{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeDictionary{}
	}

	items := make([]StoredInfoTypeDictionary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeDictionary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeDictionary expands an instance of StoredInfoTypeDictionary into a JSON
// request object.
func expandStoredInfoTypeDictionary(c *Client, f *StoredInfoTypeDictionary, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandStoredInfoTypeDictionaryWordList(c, f.WordList, res); err != nil {
		return nil, fmt.Errorf("error expanding WordList into wordList: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["wordList"] = v
	}
	if v, err := expandStoredInfoTypeDictionaryCloudStoragePath(c, f.CloudStoragePath, res); err != nil {
		return nil, fmt.Errorf("error expanding CloudStoragePath into cloudStoragePath: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["cloudStoragePath"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeDictionary flattens an instance of StoredInfoTypeDictionary from a JSON
// response object.
func flattenStoredInfoTypeDictionary(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeDictionary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeDictionary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeDictionary
	}
	r.WordList = flattenStoredInfoTypeDictionaryWordList(c, m["wordList"], res)
	r.CloudStoragePath = flattenStoredInfoTypeDictionaryCloudStoragePath(c, m["cloudStoragePath"], res)

	return r
}

// expandStoredInfoTypeDictionaryWordListMap expands the contents of StoredInfoTypeDictionaryWordList into a JSON
// request object.
func expandStoredInfoTypeDictionaryWordListMap(c *Client, f map[string]StoredInfoTypeDictionaryWordList, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeDictionaryWordListSlice expands the contents of StoredInfoTypeDictionaryWordList into a JSON
// request object.
func expandStoredInfoTypeDictionaryWordListSlice(c *Client, f []StoredInfoTypeDictionaryWordList, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeDictionaryWordList(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeDictionaryWordListMap flattens the contents of StoredInfoTypeDictionaryWordList from a JSON
// response object.
func flattenStoredInfoTypeDictionaryWordListMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeDictionaryWordList {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeDictionaryWordList{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeDictionaryWordList{}
	}

	items := make(map[string]StoredInfoTypeDictionaryWordList)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeDictionaryWordList(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeDictionaryWordListSlice flattens the contents of StoredInfoTypeDictionaryWordList from a JSON
// response object.
func flattenStoredInfoTypeDictionaryWordListSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeDictionaryWordList {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeDictionaryWordList{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeDictionaryWordList{}
	}

	items := make([]StoredInfoTypeDictionaryWordList, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeDictionaryWordList(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeDictionaryWordList expands an instance of StoredInfoTypeDictionaryWordList into a JSON
// request object.
func expandStoredInfoTypeDictionaryWordList(c *Client, f *StoredInfoTypeDictionaryWordList, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Words; v != nil {
		m["words"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeDictionaryWordList flattens an instance of StoredInfoTypeDictionaryWordList from a JSON
// response object.
func flattenStoredInfoTypeDictionaryWordList(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeDictionaryWordList {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeDictionaryWordList{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeDictionaryWordList
	}
	r.Words = dcl.FlattenStringSlice(m["words"])

	return r
}

// expandStoredInfoTypeDictionaryCloudStoragePathMap expands the contents of StoredInfoTypeDictionaryCloudStoragePath into a JSON
// request object.
func expandStoredInfoTypeDictionaryCloudStoragePathMap(c *Client, f map[string]StoredInfoTypeDictionaryCloudStoragePath, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeDictionaryCloudStoragePathSlice expands the contents of StoredInfoTypeDictionaryCloudStoragePath into a JSON
// request object.
func expandStoredInfoTypeDictionaryCloudStoragePathSlice(c *Client, f []StoredInfoTypeDictionaryCloudStoragePath, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeDictionaryCloudStoragePath(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeDictionaryCloudStoragePathMap flattens the contents of StoredInfoTypeDictionaryCloudStoragePath from a JSON
// response object.
func flattenStoredInfoTypeDictionaryCloudStoragePathMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeDictionaryCloudStoragePath {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeDictionaryCloudStoragePath{}
	}

	items := make(map[string]StoredInfoTypeDictionaryCloudStoragePath)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeDictionaryCloudStoragePath(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeDictionaryCloudStoragePathSlice flattens the contents of StoredInfoTypeDictionaryCloudStoragePath from a JSON
// response object.
func flattenStoredInfoTypeDictionaryCloudStoragePathSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeDictionaryCloudStoragePath {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeDictionaryCloudStoragePath{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeDictionaryCloudStoragePath{}
	}

	items := make([]StoredInfoTypeDictionaryCloudStoragePath, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeDictionaryCloudStoragePath(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeDictionaryCloudStoragePath expands an instance of StoredInfoTypeDictionaryCloudStoragePath into a JSON
// request object.
func expandStoredInfoTypeDictionaryCloudStoragePath(c *Client, f *StoredInfoTypeDictionaryCloudStoragePath, res *StoredInfoType) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Path; !dcl.IsEmptyValueIndirect(v) {
		m["path"] = v
	}

	return m, nil
}

// flattenStoredInfoTypeDictionaryCloudStoragePath flattens an instance of StoredInfoTypeDictionaryCloudStoragePath from a JSON
// response object.
func flattenStoredInfoTypeDictionaryCloudStoragePath(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeDictionaryCloudStoragePath {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeDictionaryCloudStoragePath{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeDictionaryCloudStoragePath
	}
	r.Path = dcl.FlattenString(m["path"])

	return r
}

// expandStoredInfoTypeRegexMap expands the contents of StoredInfoTypeRegex into a JSON
// request object.
func expandStoredInfoTypeRegexMap(c *Client, f map[string]StoredInfoTypeRegex, res *StoredInfoType) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandStoredInfoTypeRegex(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandStoredInfoTypeRegexSlice expands the contents of StoredInfoTypeRegex into a JSON
// request object.
func expandStoredInfoTypeRegexSlice(c *Client, f []StoredInfoTypeRegex, res *StoredInfoType) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandStoredInfoTypeRegex(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenStoredInfoTypeRegexMap flattens the contents of StoredInfoTypeRegex from a JSON
// response object.
func flattenStoredInfoTypeRegexMap(c *Client, i interface{}, res *StoredInfoType) map[string]StoredInfoTypeRegex {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]StoredInfoTypeRegex{}
	}

	if len(a) == 0 {
		return map[string]StoredInfoTypeRegex{}
	}

	items := make(map[string]StoredInfoTypeRegex)
	for k, item := range a {
		items[k] = *flattenStoredInfoTypeRegex(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenStoredInfoTypeRegexSlice flattens the contents of StoredInfoTypeRegex from a JSON
// response object.
func flattenStoredInfoTypeRegexSlice(c *Client, i interface{}, res *StoredInfoType) []StoredInfoTypeRegex {
	a, ok := i.([]interface{})
	if !ok {
		return []StoredInfoTypeRegex{}
	}

	if len(a) == 0 {
		return []StoredInfoTypeRegex{}
	}

	items := make([]StoredInfoTypeRegex, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenStoredInfoTypeRegex(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandStoredInfoTypeRegex expands an instance of StoredInfoTypeRegex into a JSON
// request object.
func expandStoredInfoTypeRegex(c *Client, f *StoredInfoTypeRegex, res *StoredInfoType) (map[string]interface{}, error) {
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

// flattenStoredInfoTypeRegex flattens an instance of StoredInfoTypeRegex from a JSON
// response object.
func flattenStoredInfoTypeRegex(c *Client, i interface{}, res *StoredInfoType) *StoredInfoTypeRegex {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &StoredInfoTypeRegex{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyStoredInfoTypeRegex
	}
	r.Pattern = dcl.FlattenString(m["pattern"])
	r.GroupIndexes = dcl.FlattenIntSlice(m["groupIndexes"])

	return r
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *StoredInfoType) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalStoredInfoType(b, c, r)
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

type storedInfoTypeDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         storedInfoTypeApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToStoredInfoTypeDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]storedInfoTypeDiff, error) {
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
	var diffs []storedInfoTypeDiff
	// For each operation name, create a storedInfoTypeDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := storedInfoTypeDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToStoredInfoTypeApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToStoredInfoTypeApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (storedInfoTypeApiOperation, error) {
	switch opName {

	case "updateStoredInfoTypeUpdateStoredInfoTypeOperation":
		return &updateStoredInfoTypeUpdateStoredInfoTypeOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractStoredInfoTypeFields(r *StoredInfoType) error {
	vLargeCustomDictionary := r.LargeCustomDictionary
	if vLargeCustomDictionary == nil {
		// note: explicitly not the empty object.
		vLargeCustomDictionary = &StoredInfoTypeLargeCustomDictionary{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryFields(r, vLargeCustomDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLargeCustomDictionary) {
		r.LargeCustomDictionary = vLargeCustomDictionary
	}
	vDictionary := r.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &StoredInfoTypeDictionary{}
	}
	if err := extractStoredInfoTypeDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		r.Dictionary = vDictionary
	}
	vRegex := r.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &StoredInfoTypeRegex{}
	}
	if err := extractStoredInfoTypeRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		r.Regex = vRegex
	}
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionary) error {
	vOutputPath := o.OutputPath
	if vOutputPath == nil {
		// note: explicitly not the empty object.
		vOutputPath = &StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryOutputPathFields(r, vOutputPath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOutputPath) {
		o.OutputPath = vOutputPath
	}
	vCloudStorageFileSet := o.CloudStorageFileSet
	if vCloudStorageFileSet == nil {
		// note: explicitly not the empty object.
		vCloudStorageFileSet = &StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetFields(r, vCloudStorageFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStorageFileSet) {
		o.CloudStorageFileSet = vCloudStorageFileSet
	}
	vBigQueryField := o.BigQueryField
	if vBigQueryField == nil {
		// note: explicitly not the empty object.
		vBigQueryField = &StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFields(r, vBigQueryField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigQueryField) {
		o.BigQueryField = vBigQueryField
	}
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryOutputPathFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryOutputPath) error {
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) error {
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryField) error {
	vTable := o.Table
	if vTable == nil {
		// note: explicitly not the empty object.
		vTable = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableFields(r, vTable); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTable) {
		o.Table = vTable
	}
	vField := o.Field
	if vField == nil {
		// note: explicitly not the empty object.
		vField = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldFields(r, vField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vField) {
		o.Field = vField
	}
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) error {
	return nil
}
func extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) error {
	return nil
}
func extractStoredInfoTypeDictionaryFields(r *StoredInfoType, o *StoredInfoTypeDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &StoredInfoTypeDictionaryWordList{}
	}
	if err := extractStoredInfoTypeDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &StoredInfoTypeDictionaryCloudStoragePath{}
	}
	if err := extractStoredInfoTypeDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func extractStoredInfoTypeDictionaryWordListFields(r *StoredInfoType, o *StoredInfoTypeDictionaryWordList) error {
	return nil
}
func extractStoredInfoTypeDictionaryCloudStoragePathFields(r *StoredInfoType, o *StoredInfoTypeDictionaryCloudStoragePath) error {
	return nil
}
func extractStoredInfoTypeRegexFields(r *StoredInfoType, o *StoredInfoTypeRegex) error {
	return nil
}

func postReadExtractStoredInfoTypeFields(r *StoredInfoType) error {
	vLargeCustomDictionary := r.LargeCustomDictionary
	if vLargeCustomDictionary == nil {
		// note: explicitly not the empty object.
		vLargeCustomDictionary = &StoredInfoTypeLargeCustomDictionary{}
	}
	if err := postReadExtractStoredInfoTypeLargeCustomDictionaryFields(r, vLargeCustomDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vLargeCustomDictionary) {
		r.LargeCustomDictionary = vLargeCustomDictionary
	}
	vDictionary := r.Dictionary
	if vDictionary == nil {
		// note: explicitly not the empty object.
		vDictionary = &StoredInfoTypeDictionary{}
	}
	if err := postReadExtractStoredInfoTypeDictionaryFields(r, vDictionary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vDictionary) {
		r.Dictionary = vDictionary
	}
	vRegex := r.Regex
	if vRegex == nil {
		// note: explicitly not the empty object.
		vRegex = &StoredInfoTypeRegex{}
	}
	if err := postReadExtractStoredInfoTypeRegexFields(r, vRegex); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vRegex) {
		r.Regex = vRegex
	}
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionary) error {
	vOutputPath := o.OutputPath
	if vOutputPath == nil {
		// note: explicitly not the empty object.
		vOutputPath = &StoredInfoTypeLargeCustomDictionaryOutputPath{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryOutputPathFields(r, vOutputPath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOutputPath) {
		o.OutputPath = vOutputPath
	}
	vCloudStorageFileSet := o.CloudStorageFileSet
	if vCloudStorageFileSet == nil {
		// note: explicitly not the empty object.
		vCloudStorageFileSet = &StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetFields(r, vCloudStorageFileSet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStorageFileSet) {
		o.CloudStorageFileSet = vCloudStorageFileSet
	}
	vBigQueryField := o.BigQueryField
	if vBigQueryField == nil {
		// note: explicitly not the empty object.
		vBigQueryField = &StoredInfoTypeLargeCustomDictionaryBigQueryField{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFields(r, vBigQueryField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBigQueryField) {
		o.BigQueryField = vBigQueryField
	}
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryOutputPathFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryOutputPath) error {
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) error {
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryField) error {
	vTable := o.Table
	if vTable == nil {
		// note: explicitly not the empty object.
		vTable = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableFields(r, vTable); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTable) {
		o.Table = vTable
	}
	vField := o.Field
	if vField == nil {
		// note: explicitly not the empty object.
		vField = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{}
	}
	if err := extractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldFields(r, vField); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vField) {
		o.Field = vField
	}
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) error {
	return nil
}
func postReadExtractStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldFields(r *StoredInfoType, o *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) error {
	return nil
}
func postReadExtractStoredInfoTypeDictionaryFields(r *StoredInfoType, o *StoredInfoTypeDictionary) error {
	vWordList := o.WordList
	if vWordList == nil {
		// note: explicitly not the empty object.
		vWordList = &StoredInfoTypeDictionaryWordList{}
	}
	if err := extractStoredInfoTypeDictionaryWordListFields(r, vWordList); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vWordList) {
		o.WordList = vWordList
	}
	vCloudStoragePath := o.CloudStoragePath
	if vCloudStoragePath == nil {
		// note: explicitly not the empty object.
		vCloudStoragePath = &StoredInfoTypeDictionaryCloudStoragePath{}
	}
	if err := extractStoredInfoTypeDictionaryCloudStoragePathFields(r, vCloudStoragePath); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vCloudStoragePath) {
		o.CloudStoragePath = vCloudStoragePath
	}
	return nil
}
func postReadExtractStoredInfoTypeDictionaryWordListFields(r *StoredInfoType, o *StoredInfoTypeDictionaryWordList) error {
	return nil
}
func postReadExtractStoredInfoTypeDictionaryCloudStoragePathFields(r *StoredInfoType, o *StoredInfoTypeDictionaryCloudStoragePath) error {
	return nil
}
func postReadExtractStoredInfoTypeRegexFields(r *StoredInfoType, o *StoredInfoTypeRegex) error {
	return nil
}
