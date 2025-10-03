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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Bucket struct {
	Project      *string                 `json:"project"`
	Location     *string                 `json:"location"`
	Name         *string                 `json:"name"`
	Cors         []BucketCors            `json:"cors"`
	Lifecycle    *BucketLifecycle        `json:"lifecycle"`
	Logging      *BucketLogging          `json:"logging"`
	StorageClass *BucketStorageClassEnum `json:"storageClass"`
	Versioning   *BucketVersioning       `json:"versioning"`
	Website      *BucketWebsite          `json:"website"`
}

func (r *Bucket) String() string {
	return dcl.SprintResource(r)
}

// The enum BucketLifecycleRuleActionTypeEnum.
type BucketLifecycleRuleActionTypeEnum string

// BucketLifecycleRuleActionTypeEnumRef returns a *BucketLifecycleRuleActionTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketLifecycleRuleActionTypeEnumRef(s string) *BucketLifecycleRuleActionTypeEnum {
	v := BucketLifecycleRuleActionTypeEnum(s)
	return &v
}

func (v BucketLifecycleRuleActionTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"Delete", "SetStorageClass"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketLifecycleRuleActionTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BucketLifecycleRuleConditionWithStateEnum.
type BucketLifecycleRuleConditionWithStateEnum string

// BucketLifecycleRuleConditionWithStateEnumRef returns a *BucketLifecycleRuleConditionWithStateEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketLifecycleRuleConditionWithStateEnumRef(s string) *BucketLifecycleRuleConditionWithStateEnum {
	v := BucketLifecycleRuleConditionWithStateEnum(s)
	return &v
}

func (v BucketLifecycleRuleConditionWithStateEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"LIVE", "ARCHIVED", "ANY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketLifecycleRuleConditionWithStateEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum BucketStorageClassEnum.
type BucketStorageClassEnum string

// BucketStorageClassEnumRef returns a *BucketStorageClassEnum with the value of string s
// If the empty string is provided, nil is returned.
func BucketStorageClassEnumRef(s string) *BucketStorageClassEnum {
	v := BucketStorageClassEnum(s)
	return &v
}

func (v BucketStorageClassEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MULTI_REGIONAL", "REGIONAL", "STANDARD", "NEARLINE", "COLDLINE", "ARCHIVE", "DURABLE_REDUCED_AVAILABILITY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "BucketStorageClassEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type BucketCors struct {
	empty          bool     `json:"-"`
	MaxAgeSeconds  *int64   `json:"maxAgeSeconds"`
	Method         []string `json:"method"`
	Origin         []string `json:"origin"`
	ResponseHeader []string `json:"responseHeader"`
}

type jsonBucketCors BucketCors

func (r *BucketCors) UnmarshalJSON(data []byte) error {
	var res jsonBucketCors
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketCors
	} else {

		r.MaxAgeSeconds = res.MaxAgeSeconds

		r.Method = res.Method

		r.Origin = res.Origin

		r.ResponseHeader = res.ResponseHeader

	}
	return nil
}

// This object is used to assert a desired state where this BucketCors is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketCors *BucketCors = &BucketCors{empty: true}

func (r *BucketCors) Empty() bool {
	return r.empty
}

func (r *BucketCors) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketCors) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycle struct {
	empty bool                  `json:"-"`
	Rule  []BucketLifecycleRule `json:"rule"`
}

type jsonBucketLifecycle BucketLifecycle

func (r *BucketLifecycle) UnmarshalJSON(data []byte) error {
	var res jsonBucketLifecycle
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketLifecycle
	} else {

		r.Rule = res.Rule

	}
	return nil
}

// This object is used to assert a desired state where this BucketLifecycle is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketLifecycle *BucketLifecycle = &BucketLifecycle{empty: true}

func (r *BucketLifecycle) Empty() bool {
	return r.empty
}

func (r *BucketLifecycle) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycle) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRule struct {
	empty     bool                          `json:"-"`
	Action    *BucketLifecycleRuleAction    `json:"action"`
	Condition *BucketLifecycleRuleCondition `json:"condition"`
}

type jsonBucketLifecycleRule BucketLifecycleRule

func (r *BucketLifecycleRule) UnmarshalJSON(data []byte) error {
	var res jsonBucketLifecycleRule
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketLifecycleRule
	} else {

		r.Action = res.Action

		r.Condition = res.Condition

	}
	return nil
}

// This object is used to assert a desired state where this BucketLifecycleRule is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketLifecycleRule *BucketLifecycleRule = &BucketLifecycleRule{empty: true}

func (r *BucketLifecycleRule) Empty() bool {
	return r.empty
}

func (r *BucketLifecycleRule) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRule) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRuleAction struct {
	empty        bool                               `json:"-"`
	StorageClass *string                            `json:"storageClass"`
	Type         *BucketLifecycleRuleActionTypeEnum `json:"type"`
}

type jsonBucketLifecycleRuleAction BucketLifecycleRuleAction

func (r *BucketLifecycleRuleAction) UnmarshalJSON(data []byte) error {
	var res jsonBucketLifecycleRuleAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketLifecycleRuleAction
	} else {

		r.StorageClass = res.StorageClass

		r.Type = res.Type

	}
	return nil
}

// This object is used to assert a desired state where this BucketLifecycleRuleAction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketLifecycleRuleAction *BucketLifecycleRuleAction = &BucketLifecycleRuleAction{empty: true}

func (r *BucketLifecycleRuleAction) Empty() bool {
	return r.empty
}

func (r *BucketLifecycleRuleAction) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRuleAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLifecycleRuleCondition struct {
	empty               bool                                       `json:"-"`
	Age                 *int64                                     `json:"age"`
	CreatedBefore       *string                                    `json:"createdBefore"`
	WithState           *BucketLifecycleRuleConditionWithStateEnum `json:"withState"`
	MatchesStorageClass []string                                   `json:"matchesStorageClass"`
	NumNewerVersions    *int64                                     `json:"numNewerVersions"`
}

type jsonBucketLifecycleRuleCondition BucketLifecycleRuleCondition

func (r *BucketLifecycleRuleCondition) UnmarshalJSON(data []byte) error {
	var res jsonBucketLifecycleRuleCondition
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketLifecycleRuleCondition
	} else {

		r.Age = res.Age

		r.CreatedBefore = res.CreatedBefore

		r.WithState = res.WithState

		r.MatchesStorageClass = res.MatchesStorageClass

		r.NumNewerVersions = res.NumNewerVersions

	}
	return nil
}

// This object is used to assert a desired state where this BucketLifecycleRuleCondition is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketLifecycleRuleCondition *BucketLifecycleRuleCondition = &BucketLifecycleRuleCondition{empty: true}

func (r *BucketLifecycleRuleCondition) Empty() bool {
	return r.empty
}

func (r *BucketLifecycleRuleCondition) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLifecycleRuleCondition) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketLogging struct {
	empty           bool    `json:"-"`
	LogBucket       *string `json:"logBucket"`
	LogObjectPrefix *string `json:"logObjectPrefix"`
}

type jsonBucketLogging BucketLogging

func (r *BucketLogging) UnmarshalJSON(data []byte) error {
	var res jsonBucketLogging
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketLogging
	} else {

		r.LogBucket = res.LogBucket

		r.LogObjectPrefix = res.LogObjectPrefix

	}
	return nil
}

// This object is used to assert a desired state where this BucketLogging is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketLogging *BucketLogging = &BucketLogging{empty: true}

func (r *BucketLogging) Empty() bool {
	return r.empty
}

func (r *BucketLogging) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketLogging) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketVersioning struct {
	empty   bool  `json:"-"`
	Enabled *bool `json:"enabled"`
}

type jsonBucketVersioning BucketVersioning

func (r *BucketVersioning) UnmarshalJSON(data []byte) error {
	var res jsonBucketVersioning
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketVersioning
	} else {

		r.Enabled = res.Enabled

	}
	return nil
}

// This object is used to assert a desired state where this BucketVersioning is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketVersioning *BucketVersioning = &BucketVersioning{empty: true}

func (r *BucketVersioning) Empty() bool {
	return r.empty
}

func (r *BucketVersioning) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketVersioning) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type BucketWebsite struct {
	empty          bool    `json:"-"`
	MainPageSuffix *string `json:"mainPageSuffix"`
	NotFoundPage   *string `json:"notFoundPage"`
}

type jsonBucketWebsite BucketWebsite

func (r *BucketWebsite) UnmarshalJSON(data []byte) error {
	var res jsonBucketWebsite
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyBucketWebsite
	} else {

		r.MainPageSuffix = res.MainPageSuffix

		r.NotFoundPage = res.NotFoundPage

	}
	return nil
}

// This object is used to assert a desired state where this BucketWebsite is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyBucketWebsite *BucketWebsite = &BucketWebsite{empty: true}

func (r *BucketWebsite) Empty() bool {
	return r.empty
}

func (r *BucketWebsite) String() string {
	return dcl.SprintResource(r)
}

func (r *BucketWebsite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Bucket) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "storage",
		Type:    "Bucket",
		Version: "alpha",
	}
}

func (r *Bucket) ID() (string, error) {
	if err := extractBucketFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":       dcl.ValueOrEmptyString(nr.Project),
		"location":      dcl.ValueOrEmptyString(nr.Location),
		"name":          dcl.ValueOrEmptyString(nr.Name),
		"cors":          dcl.ValueOrEmptyString(nr.Cors),
		"lifecycle":     dcl.ValueOrEmptyString(nr.Lifecycle),
		"logging":       dcl.ValueOrEmptyString(nr.Logging),
		"storage_class": dcl.ValueOrEmptyString(nr.StorageClass),
		"versioning":    dcl.ValueOrEmptyString(nr.Versioning),
		"website":       dcl.ValueOrEmptyString(nr.Website),
	}
	return dcl.Nprintf("b/{{name}}?userProject={{project}}", params), nil
}

const BucketMaxPage = -1

type BucketList struct {
	Items []*Bucket

	nextToken string

	pageSize int32

	resource *Bucket
}

func (l *BucketList) HasNext() bool {
	return l.nextToken != ""
}

func (l *BucketList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listBucket(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListBucket(ctx context.Context, project string) (*BucketList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListBucketWithMaxResults(ctx, project, BucketMaxPage)

}

func (c *Client) ListBucketWithMaxResults(ctx context.Context, project string, pageSize int32) (*BucketList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Bucket{
		Project: &project,
	}
	items, token, err := c.listBucket(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &BucketList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetBucket(ctx context.Context, r *Bucket) (*Bucket, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractBucketFields(r)

	b, err := c.getBucketRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalBucket(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeBucketNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractBucketFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteBucket(ctx context.Context, r *Bucket) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Bucket resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Bucket...")
	deleteOp := deleteBucketOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllBucket deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllBucket(ctx context.Context, project string, filter func(*Bucket) bool) error {
	listObj, err := c.ListBucket(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllBucket(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllBucket(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyBucket(ctx context.Context, rawDesired *Bucket, opts ...dcl.ApplyOption) (*Bucket, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Bucket
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyBucketHelper(c, ctx, rawDesired, opts...)
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

func applyBucketHelper(c *Client, ctx context.Context, rawDesired *Bucket, opts ...dcl.ApplyOption) (*Bucket, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyBucket...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractBucketFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.bucketDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToBucketDiffs(c.Config, fieldDiffs, opts)
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
	var ops []bucketApiOperation
	if create {
		ops = append(ops, &createBucketOperation{})
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
	return applyBucketDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyBucketDiff(c *Client, ctx context.Context, desired *Bucket, rawDesired *Bucket, ops []bucketApiOperation, opts ...dcl.ApplyOption) (*Bucket, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetBucket(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createBucketOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapBucket(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeBucketNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeBucketNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeBucketDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractBucketFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractBucketFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffBucket(c, newDesired, newState)
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

func (r *Bucket) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
