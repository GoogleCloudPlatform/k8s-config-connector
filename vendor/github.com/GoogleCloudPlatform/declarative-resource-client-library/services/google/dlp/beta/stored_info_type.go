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
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type StoredInfoType struct {
	Name                  *string                              `json:"name"`
	DisplayName           *string                              `json:"displayName"`
	Description           *string                              `json:"description"`
	LargeCustomDictionary *StoredInfoTypeLargeCustomDictionary `json:"largeCustomDictionary"`
	Dictionary            *StoredInfoTypeDictionary            `json:"dictionary"`
	Regex                 *StoredInfoTypeRegex                 `json:"regex"`
	Parent                *string                              `json:"parent"`
	Location              *string                              `json:"location"`
}

func (r *StoredInfoType) String() string {
	return dcl.SprintResource(r)
}

type StoredInfoTypeLargeCustomDictionary struct {
	empty               bool                                                    `json:"-"`
	OutputPath          *StoredInfoTypeLargeCustomDictionaryOutputPath          `json:"outputPath"`
	CloudStorageFileSet *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet `json:"cloudStorageFileSet"`
	BigQueryField       *StoredInfoTypeLargeCustomDictionaryBigQueryField       `json:"bigQueryField"`
}

type jsonStoredInfoTypeLargeCustomDictionary StoredInfoTypeLargeCustomDictionary

func (r *StoredInfoTypeLargeCustomDictionary) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionary
	} else {

		r.OutputPath = res.OutputPath

		r.CloudStorageFileSet = res.CloudStorageFileSet

		r.BigQueryField = res.BigQueryField

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionary *StoredInfoTypeLargeCustomDictionary = &StoredInfoTypeLargeCustomDictionary{empty: true}

func (r *StoredInfoTypeLargeCustomDictionary) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeLargeCustomDictionaryOutputPath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonStoredInfoTypeLargeCustomDictionaryOutputPath StoredInfoTypeLargeCustomDictionaryOutputPath

func (r *StoredInfoTypeLargeCustomDictionaryOutputPath) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionaryOutputPath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionaryOutputPath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionaryOutputPath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionaryOutputPath *StoredInfoTypeLargeCustomDictionaryOutputPath = &StoredInfoTypeLargeCustomDictionaryOutputPath{empty: true}

func (r *StoredInfoTypeLargeCustomDictionaryOutputPath) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionaryOutputPath) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionaryOutputPath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet struct {
	empty bool    `json:"-"`
	Url   *string `json:"url"`
}

type jsonStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet

func (r *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	} else {

		r.Url = res.Url

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet = &StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet{empty: true}

func (r *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeLargeCustomDictionaryBigQueryField struct {
	empty bool                                                   `json:"-"`
	Table *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable `json:"table"`
	Field *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField `json:"field"`
}

type jsonStoredInfoTypeLargeCustomDictionaryBigQueryField StoredInfoTypeLargeCustomDictionaryBigQueryField

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryField) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionaryBigQueryField
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionaryBigQueryField
	} else {

		r.Table = res.Table

		r.Field = res.Field

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionaryBigQueryField is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionaryBigQueryField *StoredInfoTypeLargeCustomDictionaryBigQueryField = &StoredInfoTypeLargeCustomDictionaryBigQueryField{empty: true}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryField) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryField) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryField) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	TableId   *string `json:"tableId"`
}

type jsonStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.TableId = res.TableId

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable{empty: true}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeLargeCustomDictionaryBigQueryFieldField struct {
	empty bool    `json:"-"`
	Name  *string `json:"name"`
}

type jsonStoredInfoTypeLargeCustomDictionaryBigQueryFieldField StoredInfoTypeLargeCustomDictionaryBigQueryFieldField

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	} else {

		r.Name = res.Name

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeLargeCustomDictionaryBigQueryFieldField is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeLargeCustomDictionaryBigQueryFieldField *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField = &StoredInfoTypeLargeCustomDictionaryBigQueryFieldField{empty: true}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeLargeCustomDictionaryBigQueryFieldField) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeDictionary struct {
	empty            bool                                      `json:"-"`
	WordList         *StoredInfoTypeDictionaryWordList         `json:"wordList"`
	CloudStoragePath *StoredInfoTypeDictionaryCloudStoragePath `json:"cloudStoragePath"`
}

type jsonStoredInfoTypeDictionary StoredInfoTypeDictionary

func (r *StoredInfoTypeDictionary) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeDictionary
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeDictionary
	} else {

		r.WordList = res.WordList

		r.CloudStoragePath = res.CloudStoragePath

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeDictionary is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeDictionary *StoredInfoTypeDictionary = &StoredInfoTypeDictionary{empty: true}

func (r *StoredInfoTypeDictionary) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeDictionary) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeDictionary) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeDictionaryWordList struct {
	empty bool     `json:"-"`
	Words []string `json:"words"`
}

type jsonStoredInfoTypeDictionaryWordList StoredInfoTypeDictionaryWordList

func (r *StoredInfoTypeDictionaryWordList) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeDictionaryWordList
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeDictionaryWordList
	} else {

		r.Words = res.Words

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeDictionaryWordList is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeDictionaryWordList *StoredInfoTypeDictionaryWordList = &StoredInfoTypeDictionaryWordList{empty: true}

func (r *StoredInfoTypeDictionaryWordList) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeDictionaryWordList) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeDictionaryWordList) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeDictionaryCloudStoragePath struct {
	empty bool    `json:"-"`
	Path  *string `json:"path"`
}

type jsonStoredInfoTypeDictionaryCloudStoragePath StoredInfoTypeDictionaryCloudStoragePath

func (r *StoredInfoTypeDictionaryCloudStoragePath) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeDictionaryCloudStoragePath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeDictionaryCloudStoragePath
	} else {

		r.Path = res.Path

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeDictionaryCloudStoragePath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeDictionaryCloudStoragePath *StoredInfoTypeDictionaryCloudStoragePath = &StoredInfoTypeDictionaryCloudStoragePath{empty: true}

func (r *StoredInfoTypeDictionaryCloudStoragePath) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeDictionaryCloudStoragePath) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeDictionaryCloudStoragePath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type StoredInfoTypeRegex struct {
	empty        bool    `json:"-"`
	Pattern      *string `json:"pattern"`
	GroupIndexes []int64 `json:"groupIndexes"`
}

type jsonStoredInfoTypeRegex StoredInfoTypeRegex

func (r *StoredInfoTypeRegex) UnmarshalJSON(data []byte) error {
	var res jsonStoredInfoTypeRegex
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyStoredInfoTypeRegex
	} else {

		r.Pattern = res.Pattern

		r.GroupIndexes = res.GroupIndexes

	}
	return nil
}

// This object is used to assert a desired state where this StoredInfoTypeRegex is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyStoredInfoTypeRegex *StoredInfoTypeRegex = &StoredInfoTypeRegex{empty: true}

func (r *StoredInfoTypeRegex) Empty() bool {
	return r.empty
}

func (r *StoredInfoTypeRegex) String() string {
	return dcl.SprintResource(r)
}

func (r *StoredInfoTypeRegex) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *StoredInfoType) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "dlp",
		Type:    "StoredInfoType",
		Version: "beta",
	}
}

func (r *StoredInfoType) ID() (string, error) {
	if err := extractStoredInfoTypeFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                    dcl.ValueOrEmptyString(nr.Name),
		"display_name":            dcl.ValueOrEmptyString(nr.DisplayName),
		"description":             dcl.ValueOrEmptyString(nr.Description),
		"large_custom_dictionary": dcl.ValueOrEmptyString(nr.LargeCustomDictionary),
		"dictionary":              dcl.ValueOrEmptyString(nr.Dictionary),
		"regex":                   dcl.ValueOrEmptyString(nr.Regex),
		"parent":                  dcl.ValueOrEmptyString(nr.Parent),
		"location":                dcl.ValueOrEmptyString(nr.Location),
	}
	if dcl.IsRegion(nr.Location) {
		return dcl.Nprintf("{{parent}}/locations/{{location}}/storedInfoTypes/{{name}}", params), nil
	}

	return dcl.Nprintf("{{parent}}/storedInfoTypes/{{name}}", params), nil
}

const StoredInfoTypeMaxPage = -1

type StoredInfoTypeList struct {
	Items []*StoredInfoType

	nextToken string

	pageSize int32

	resource *StoredInfoType
}

func (l *StoredInfoTypeList) HasNext() bool {
	return l.nextToken != ""
}

func (l *StoredInfoTypeList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listStoredInfoType(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListStoredInfoType(ctx context.Context, location, parent string) (*StoredInfoTypeList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListStoredInfoTypeWithMaxResults(ctx, location, parent, StoredInfoTypeMaxPage)

}

func (c *Client) ListStoredInfoTypeWithMaxResults(ctx context.Context, location, parent string, pageSize int32) (*StoredInfoTypeList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &StoredInfoType{
		Location: &location,
		Parent:   &parent,
	}
	items, token, err := c.listStoredInfoType(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &StoredInfoTypeList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetStoredInfoType(ctx context.Context, r *StoredInfoType) (*StoredInfoType, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractStoredInfoTypeFields(r)

	b, err := c.getStoredInfoTypeRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	item, err := storedInfoTypeFromParent(r, m)
	if err != nil {
		return nil, err
	}
	b, err = json.Marshal(item)
	if err != nil {
		return nil, err
	}

	result, err := unmarshalStoredInfoType(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Location = r.Location
	result.Parent = r.Parent
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeStoredInfoTypeNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractStoredInfoTypeFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteStoredInfoType(ctx context.Context, r *StoredInfoType) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("StoredInfoType resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting StoredInfoType...")
	deleteOp := deleteStoredInfoTypeOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllStoredInfoType deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllStoredInfoType(ctx context.Context, location, parent string, filter func(*StoredInfoType) bool) error {
	listObj, err := c.ListStoredInfoType(ctx, location, parent)
	if err != nil {
		return err
	}

	err = c.deleteAllStoredInfoType(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllStoredInfoType(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyStoredInfoType(ctx context.Context, rawDesired *StoredInfoType, opts ...dcl.ApplyOption) (*StoredInfoType, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *StoredInfoType
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyStoredInfoTypeHelper(c, ctx, rawDesired, opts...)
		resultNewState = newState
		if err != nil {
			// If the error is 409, there is conflict in resource update.
			// Here we want to apply changes based on latest state.
			if dcl.IsConflictError(err) {
				return &dcl.RetryDetails{}, dcl.OperationNotDone{Err: err}
			}
			return nil, err
		}
		return nil, nil
	}, c.Config.RetryProvider)
	return resultNewState, err
}

func applyStoredInfoTypeHelper(c *Client, ctx context.Context, rawDesired *StoredInfoType, opts ...dcl.ApplyOption) (*StoredInfoType, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyStoredInfoType...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractStoredInfoTypeFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.storedInfoTypeDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToStoredInfoTypeDiffs(c.Config, fieldDiffs, opts)
	if err != nil {
		return nil, err
	}

	// TODO(magic-modules-eng): 2.2 Feasibility check (all updates are feasible so far).

	// 2.3: Lifecycle Directive Check
	var create bool
	lp := dcl.FetchLifecycleParams(opts)
	if initial == nil {
		if dcl.HasLifecycleParam(lp, dcl.BlockCreation) {
			return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Creation blocked by lifecycle params: %#v.", desired)}
		}
		create = true
	} else if dcl.HasLifecycleParam(lp, dcl.BlockAcquire) {
		return nil, dcl.ApplyInfeasibleError{
			Message: fmt.Sprintf("Resource already exists - apply blocked by lifecycle params: %#v.", initial),
		}
	} else {
		for _, d := range diffs {
			if d.RequiresRecreate {
				return nil, dcl.ApplyInfeasibleError{
					Message: fmt.Sprintf("infeasible update: (%v) would require recreation", d),
				}
			}
			if dcl.HasLifecycleParam(lp, dcl.BlockModification) {
				return nil, dcl.ApplyInfeasibleError{Message: fmt.Sprintf("Modification blocked, diff (%v) unresolvable.", d)}
			}
		}
	}

	// 2.4 Imperative Request Planning
	var ops []storedInfoTypeApiOperation
	if create {
		ops = append(ops, &createStoredInfoTypeOperation{})
	} else {
		for _, d := range diffs {
			ops = append(ops, d.UpdateOp)
		}
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created plan: %#v", ops)

	// 2.5 Request Actuation
	for _, op := range ops {
		c.Config.Logger.InfoWithContextf(ctx, "Performing operation %T %+v", op, op)
		if err := op.do(ctx, desired, c); err != nil {
			c.Config.Logger.InfoWithContextf(ctx, "Failed operation %T %+v: %v", op, op, err)
			return nil, err
		}
		c.Config.Logger.InfoWithContextf(ctx, "Finished operation %T %+v", op, op)
	}
	return applyStoredInfoTypeDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyStoredInfoTypeDiff(c *Client, ctx context.Context, desired *StoredInfoType, rawDesired *StoredInfoType, ops []storedInfoTypeApiOperation, opts ...dcl.ApplyOption) (*StoredInfoType, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetStoredInfoType(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createStoredInfoTypeOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapStoredInfoType(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeStoredInfoTypeNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeStoredInfoTypeNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeStoredInfoTypeDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractStoredInfoTypeFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractStoredInfoTypeFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffStoredInfoType(c, newDesired, newState)
	if err != nil {
		return newState, err
	}

	if len(newDiffs) == 0 {
		c.Config.Logger.InfoWithContext(ctx, "No diffs found. Apply was successful.")
	} else {
		c.Config.Logger.InfoWithContextf(ctx, "Found diffs: %v", newDiffs)
		diffMessages := make([]string, len(newDiffs))
		for i, d := range newDiffs {
			diffMessages[i] = fmt.Sprintf("%v", d)
		}
		return newState, dcl.DiffAfterApplyError{Diffs: diffMessages}
	}
	c.Config.Logger.InfoWithContext(ctx, "Done Apply.")
	return newState, nil
}
