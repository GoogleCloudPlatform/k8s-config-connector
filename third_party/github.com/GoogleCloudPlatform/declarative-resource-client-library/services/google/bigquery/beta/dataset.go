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

type Dataset struct {
	Etag                           *string                                `json:"etag"`
	Id                             *string                                `json:"id"`
	SelfLink                       *string                                `json:"selfLink"`
	Name                           *string                                `json:"name"`
	Project                        *string                                `json:"project"`
	FriendlyName                   *string                                `json:"friendlyName"`
	Description                    *string                                `json:"description"`
	DefaultTableExpirationMs       *string                                `json:"defaultTableExpirationMs"`
	DefaultPartitionExpirationMs   *string                                `json:"defaultPartitionExpirationMs"`
	Labels                         map[string]string                      `json:"labels"`
	Access                         []DatasetAccess                        `json:"access"`
	CreationTime                   *int64                                 `json:"creationTime"`
	LastModifiedTime               *int64                                 `json:"lastModifiedTime"`
	Location                       *string                                `json:"location"`
	Published                      *bool                                  `json:"published"`
	DefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration `json:"defaultEncryptionConfiguration"`
}

func (r *Dataset) String() string {
	return dcl.SprintResource(r)
}

type DatasetAccess struct {
	empty        bool                  `json:"-"`
	Role         *string               `json:"role"`
	UserByEmail  *string               `json:"userByEmail"`
	GroupByEmail *string               `json:"groupByEmail"`
	Domain       *string               `json:"domain"`
	SpecialGroup *string               `json:"specialGroup"`
	IamMember    *string               `json:"iamMember"`
	View         *DatasetAccessView    `json:"view"`
	Routine      *DatasetAccessRoutine `json:"routine"`
}

type jsonDatasetAccess DatasetAccess

func (r *DatasetAccess) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccess
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccess
	} else {

		r.Role = res.Role

		r.UserByEmail = res.UserByEmail

		r.GroupByEmail = res.GroupByEmail

		r.Domain = res.Domain

		r.SpecialGroup = res.SpecialGroup

		r.IamMember = res.IamMember

		r.View = res.View

		r.Routine = res.Routine

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccess is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDatasetAccess *DatasetAccess = &DatasetAccess{empty: true}

func (r *DatasetAccess) Empty() bool {
	return r.empty
}

func (r *DatasetAccess) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccess) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetAccessView struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	TableId   *string `json:"tableId"`
}

type jsonDatasetAccessView DatasetAccessView

func (r *DatasetAccessView) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccessView
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccessView
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.TableId = res.TableId

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccessView is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDatasetAccessView *DatasetAccessView = &DatasetAccessView{empty: true}

func (r *DatasetAccessView) Empty() bool {
	return r.empty
}

func (r *DatasetAccessView) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccessView) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetAccessRoutine struct {
	empty     bool    `json:"-"`
	ProjectId *string `json:"projectId"`
	DatasetId *string `json:"datasetId"`
	RoutineId *string `json:"routineId"`
}

type jsonDatasetAccessRoutine DatasetAccessRoutine

func (r *DatasetAccessRoutine) UnmarshalJSON(data []byte) error {
	var res jsonDatasetAccessRoutine
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetAccessRoutine
	} else {

		r.ProjectId = res.ProjectId

		r.DatasetId = res.DatasetId

		r.RoutineId = res.RoutineId

	}
	return nil
}

// This object is used to assert a desired state where this DatasetAccessRoutine is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDatasetAccessRoutine *DatasetAccessRoutine = &DatasetAccessRoutine{empty: true}

func (r *DatasetAccessRoutine) Empty() bool {
	return r.empty
}

func (r *DatasetAccessRoutine) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetAccessRoutine) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type DatasetDefaultEncryptionConfiguration struct {
	empty      bool    `json:"-"`
	KmsKeyName *string `json:"kmsKeyName"`
}

type jsonDatasetDefaultEncryptionConfiguration DatasetDefaultEncryptionConfiguration

func (r *DatasetDefaultEncryptionConfiguration) UnmarshalJSON(data []byte) error {
	var res jsonDatasetDefaultEncryptionConfiguration
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyDatasetDefaultEncryptionConfiguration
	} else {

		r.KmsKeyName = res.KmsKeyName

	}
	return nil
}

// This object is used to assert a desired state where this DatasetDefaultEncryptionConfiguration is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyDatasetDefaultEncryptionConfiguration *DatasetDefaultEncryptionConfiguration = &DatasetDefaultEncryptionConfiguration{empty: true}

func (r *DatasetDefaultEncryptionConfiguration) Empty() bool {
	return r.empty
}

func (r *DatasetDefaultEncryptionConfiguration) String() string {
	return dcl.SprintResource(r)
}

func (r *DatasetDefaultEncryptionConfiguration) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Dataset) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "bigquery",
		Type:    "Dataset",
		Version: "beta",
	}
}

func (r *Dataset) ID() (string, error) {
	if err := extractDatasetFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"etag":                             dcl.ValueOrEmptyString(nr.Etag),
		"id":                               dcl.ValueOrEmptyString(nr.Id),
		"self_link":                        dcl.ValueOrEmptyString(nr.SelfLink),
		"name":                             dcl.ValueOrEmptyString(nr.Name),
		"project":                          dcl.ValueOrEmptyString(nr.Project),
		"friendly_name":                    dcl.ValueOrEmptyString(nr.FriendlyName),
		"description":                      dcl.ValueOrEmptyString(nr.Description),
		"default_table_expiration_ms":      dcl.ValueOrEmptyString(nr.DefaultTableExpirationMs),
		"default_partition_expiration_ms":  dcl.ValueOrEmptyString(nr.DefaultPartitionExpirationMs),
		"labels":                           dcl.ValueOrEmptyString(nr.Labels),
		"access":                           dcl.ValueOrEmptyString(nr.Access),
		"creation_time":                    dcl.ValueOrEmptyString(nr.CreationTime),
		"last_modified_time":               dcl.ValueOrEmptyString(nr.LastModifiedTime),
		"location":                         dcl.ValueOrEmptyString(nr.Location),
		"published":                        dcl.ValueOrEmptyString(nr.Published),
		"default_encryption_configuration": dcl.ValueOrEmptyString(nr.DefaultEncryptionConfiguration),
	}
	return dcl.Nprintf("projects/{{project}}/datasets/{{name}}", params), nil
}

const DatasetMaxPage = -1

type DatasetList struct {
	Items []*Dataset

	nextToken string

	pageSize int32

	resource *Dataset
}

func (l *DatasetList) HasNext() bool {
	return l.nextToken != ""
}

func (l *DatasetList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listDataset(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListDataset(ctx context.Context, project string) (*DatasetList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListDatasetWithMaxResults(ctx, project, DatasetMaxPage)

}

func (c *Client) ListDatasetWithMaxResults(ctx context.Context, project string, pageSize int32) (*DatasetList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Dataset{
		Project: &project,
	}
	items, token, err := c.listDataset(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &DatasetList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetDataset(ctx context.Context, r *Dataset) (*Dataset, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractDatasetFields(r)

	b, err := c.getDatasetRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalDataset(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeDatasetNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractDatasetFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteDataset(ctx context.Context, r *Dataset) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Dataset resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Dataset...")
	deleteOp := deleteDatasetOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllDataset deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllDataset(ctx context.Context, project string, filter func(*Dataset) bool) error {
	listObj, err := c.ListDataset(ctx, project)
	if err != nil {
		return err
	}

	err = c.deleteAllDataset(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllDataset(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyDataset(ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Dataset
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyDatasetHelper(c, ctx, rawDesired, opts...)
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

func applyDatasetHelper(c *Client, ctx context.Context, rawDesired *Dataset, opts ...dcl.ApplyOption) (*Dataset, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyDataset...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractDatasetFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.datasetDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToDatasetDiffs(c.Config, fieldDiffs, opts)
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
	var ops []datasetApiOperation
	if create {
		ops = append(ops, &createDatasetOperation{})
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
	return applyDatasetDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyDatasetDiff(c *Client, ctx context.Context, desired *Dataset, rawDesired *Dataset, ops []datasetApiOperation, opts ...dcl.ApplyOption) (*Dataset, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetDataset(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createDatasetOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapDataset(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeDatasetNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeDatasetNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeDatasetDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractDatasetFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractDatasetFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffDataset(c, newDesired, newState)
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
