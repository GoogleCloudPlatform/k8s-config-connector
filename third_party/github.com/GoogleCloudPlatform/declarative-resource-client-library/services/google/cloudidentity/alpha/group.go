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
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Group struct {
	Name                     *string                        `json:"name"`
	GroupKey                 *GroupGroupKey                 `json:"groupKey"`
	AdditionalGroupKeys      []GroupAdditionalGroupKeys     `json:"additionalGroupKeys"`
	Parent                   *string                        `json:"parent"`
	DisplayName              *string                        `json:"displayName"`
	Description              *string                        `json:"description"`
	CreateTime               *string                        `json:"createTime"`
	UpdateTime               *string                        `json:"updateTime"`
	Labels                   map[string]string              `json:"labels"`
	DirectMemberCount        *int64                         `json:"directMemberCount"`
	DirectMemberCountPerType *GroupDirectMemberCountPerType `json:"directMemberCountPerType"`
	DerivedAliases           []GroupDerivedAliases          `json:"derivedAliases"`
	DynamicGroupMetadata     *GroupDynamicGroupMetadata     `json:"dynamicGroupMetadata"`
	PosixGroups              []GroupPosixGroups             `json:"posixGroups"`
	InitialGroupConfig       *GroupInitialGroupConfigEnum   `json:"initialGroupConfig"`
}

func (r *Group) String() string {
	return dcl.SprintResource(r)
}

// The enum GroupDynamicGroupMetadataQueriesResourceTypeEnum.
type GroupDynamicGroupMetadataQueriesResourceTypeEnum string

// GroupDynamicGroupMetadataQueriesResourceTypeEnumRef returns a *GroupDynamicGroupMetadataQueriesResourceTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GroupDynamicGroupMetadataQueriesResourceTypeEnumRef(s string) *GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	v := GroupDynamicGroupMetadataQueriesResourceTypeEnum(s)
	return &v
}

func (v GroupDynamicGroupMetadataQueriesResourceTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"RESOURCE_TYPE_UNSPECIFIED", "USER"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GroupDynamicGroupMetadataQueriesResourceTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GroupDynamicGroupMetadataStatusStatusEnum.
type GroupDynamicGroupMetadataStatusStatusEnum string

// GroupDynamicGroupMetadataStatusStatusEnumRef returns a *GroupDynamicGroupMetadataStatusStatusEnum with the value of string s
// If the empty string is provided, nil is returned.
func GroupDynamicGroupMetadataStatusStatusEnumRef(s string) *GroupDynamicGroupMetadataStatusStatusEnum {
	v := GroupDynamicGroupMetadataStatusStatusEnum(s)
	return &v
}

func (v GroupDynamicGroupMetadataStatusStatusEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"STATUS_UNSPECIFIED", "UP_TO_DATE", "UPDATING_MEMBERSHIPS", "INVALID_QUERY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GroupDynamicGroupMetadataStatusStatusEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GroupInitialGroupConfigEnum.
type GroupInitialGroupConfigEnum string

// GroupInitialGroupConfigEnumRef returns a *GroupInitialGroupConfigEnum with the value of string s
// If the empty string is provided, nil is returned.
func GroupInitialGroupConfigEnumRef(s string) *GroupInitialGroupConfigEnum {
	v := GroupInitialGroupConfigEnum(s)
	return &v
}

func (v GroupInitialGroupConfigEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INITIAL_GROUP_CONFIG_UNSPECIFIED", "WITH_INITIAL_OWNER", "EMPTY"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GroupInitialGroupConfigEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type GroupGroupKey struct {
	empty     bool    `json:"-"`
	Id        *string `json:"id"`
	Namespace *string `json:"namespace"`
}

type jsonGroupGroupKey GroupGroupKey

func (r *GroupGroupKey) UnmarshalJSON(data []byte) error {
	var res jsonGroupGroupKey
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupGroupKey
	} else {

		r.Id = res.Id

		r.Namespace = res.Namespace

	}
	return nil
}

// This object is used to assert a desired state where this GroupGroupKey is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupGroupKey *GroupGroupKey = &GroupGroupKey{empty: true}

func (r *GroupGroupKey) Empty() bool {
	return r.empty
}

func (r *GroupGroupKey) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupGroupKey) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupAdditionalGroupKeys struct {
	empty     bool    `json:"-"`
	Id        *string `json:"id"`
	Namespace *string `json:"namespace"`
}

type jsonGroupAdditionalGroupKeys GroupAdditionalGroupKeys

func (r *GroupAdditionalGroupKeys) UnmarshalJSON(data []byte) error {
	var res jsonGroupAdditionalGroupKeys
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupAdditionalGroupKeys
	} else {

		r.Id = res.Id

		r.Namespace = res.Namespace

	}
	return nil
}

// This object is used to assert a desired state where this GroupAdditionalGroupKeys is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupAdditionalGroupKeys *GroupAdditionalGroupKeys = &GroupAdditionalGroupKeys{empty: true}

func (r *GroupAdditionalGroupKeys) Empty() bool {
	return r.empty
}

func (r *GroupAdditionalGroupKeys) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupAdditionalGroupKeys) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupDirectMemberCountPerType struct {
	empty      bool   `json:"-"`
	UserCount  *int64 `json:"userCount"`
	GroupCount *int64 `json:"groupCount"`
}

type jsonGroupDirectMemberCountPerType GroupDirectMemberCountPerType

func (r *GroupDirectMemberCountPerType) UnmarshalJSON(data []byte) error {
	var res jsonGroupDirectMemberCountPerType
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupDirectMemberCountPerType
	} else {

		r.UserCount = res.UserCount

		r.GroupCount = res.GroupCount

	}
	return nil
}

// This object is used to assert a desired state where this GroupDirectMemberCountPerType is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupDirectMemberCountPerType *GroupDirectMemberCountPerType = &GroupDirectMemberCountPerType{empty: true}

func (r *GroupDirectMemberCountPerType) Empty() bool {
	return r.empty
}

func (r *GroupDirectMemberCountPerType) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupDirectMemberCountPerType) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupDerivedAliases struct {
	empty     bool    `json:"-"`
	Id        *string `json:"id"`
	Namespace *string `json:"namespace"`
}

type jsonGroupDerivedAliases GroupDerivedAliases

func (r *GroupDerivedAliases) UnmarshalJSON(data []byte) error {
	var res jsonGroupDerivedAliases
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupDerivedAliases
	} else {

		r.Id = res.Id

		r.Namespace = res.Namespace

	}
	return nil
}

// This object is used to assert a desired state where this GroupDerivedAliases is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupDerivedAliases *GroupDerivedAliases = &GroupDerivedAliases{empty: true}

func (r *GroupDerivedAliases) Empty() bool {
	return r.empty
}

func (r *GroupDerivedAliases) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupDerivedAliases) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupDynamicGroupMetadata struct {
	empty   bool                               `json:"-"`
	Queries []GroupDynamicGroupMetadataQueries `json:"queries"`
	Status  *GroupDynamicGroupMetadataStatus   `json:"status"`
}

type jsonGroupDynamicGroupMetadata GroupDynamicGroupMetadata

func (r *GroupDynamicGroupMetadata) UnmarshalJSON(data []byte) error {
	var res jsonGroupDynamicGroupMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupDynamicGroupMetadata
	} else {

		r.Queries = res.Queries

		r.Status = res.Status

	}
	return nil
}

// This object is used to assert a desired state where this GroupDynamicGroupMetadata is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupDynamicGroupMetadata *GroupDynamicGroupMetadata = &GroupDynamicGroupMetadata{empty: true}

func (r *GroupDynamicGroupMetadata) Empty() bool {
	return r.empty
}

func (r *GroupDynamicGroupMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupDynamicGroupMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupDynamicGroupMetadataQueries struct {
	empty        bool                                              `json:"-"`
	ResourceType *GroupDynamicGroupMetadataQueriesResourceTypeEnum `json:"resourceType"`
	Query        *string                                           `json:"query"`
}

type jsonGroupDynamicGroupMetadataQueries GroupDynamicGroupMetadataQueries

func (r *GroupDynamicGroupMetadataQueries) UnmarshalJSON(data []byte) error {
	var res jsonGroupDynamicGroupMetadataQueries
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupDynamicGroupMetadataQueries
	} else {

		r.ResourceType = res.ResourceType

		r.Query = res.Query

	}
	return nil
}

// This object is used to assert a desired state where this GroupDynamicGroupMetadataQueries is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupDynamicGroupMetadataQueries *GroupDynamicGroupMetadataQueries = &GroupDynamicGroupMetadataQueries{empty: true}

func (r *GroupDynamicGroupMetadataQueries) Empty() bool {
	return r.empty
}

func (r *GroupDynamicGroupMetadataQueries) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupDynamicGroupMetadataQueries) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupDynamicGroupMetadataStatus struct {
	empty      bool                                       `json:"-"`
	Status     *GroupDynamicGroupMetadataStatusStatusEnum `json:"status"`
	StatusTime *string                                    `json:"statusTime"`
}

type jsonGroupDynamicGroupMetadataStatus GroupDynamicGroupMetadataStatus

func (r *GroupDynamicGroupMetadataStatus) UnmarshalJSON(data []byte) error {
	var res jsonGroupDynamicGroupMetadataStatus
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupDynamicGroupMetadataStatus
	} else {

		r.Status = res.Status

		r.StatusTime = res.StatusTime

	}
	return nil
}

// This object is used to assert a desired state where this GroupDynamicGroupMetadataStatus is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupDynamicGroupMetadataStatus *GroupDynamicGroupMetadataStatus = &GroupDynamicGroupMetadataStatus{empty: true}

func (r *GroupDynamicGroupMetadataStatus) Empty() bool {
	return r.empty
}

func (r *GroupDynamicGroupMetadataStatus) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupDynamicGroupMetadataStatus) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GroupPosixGroups struct {
	empty    bool    `json:"-"`
	Name     *string `json:"name"`
	Gid      *string `json:"gid"`
	SystemId *string `json:"systemId"`
}

type jsonGroupPosixGroups GroupPosixGroups

func (r *GroupPosixGroups) UnmarshalJSON(data []byte) error {
	var res jsonGroupPosixGroups
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGroupPosixGroups
	} else {

		r.Name = res.Name

		r.Gid = res.Gid

		r.SystemId = res.SystemId

	}
	return nil
}

// This object is used to assert a desired state where this GroupPosixGroups is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGroupPosixGroups *GroupPosixGroups = &GroupPosixGroups{empty: true}

func (r *GroupPosixGroups) Empty() bool {
	return r.empty
}

func (r *GroupPosixGroups) String() string {
	return dcl.SprintResource(r)
}

func (r *GroupPosixGroups) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Group) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "cloudidentity",
		Type:    "Group",
		Version: "alpha",
	}
}

func (r *Group) ID() (string, error) {
	if err := extractGroupFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                         dcl.ValueOrEmptyString(nr.Name),
		"group_key":                    dcl.ValueOrEmptyString(nr.GroupKey),
		"additional_group_keys":        dcl.ValueOrEmptyString(nr.AdditionalGroupKeys),
		"parent":                       dcl.ValueOrEmptyString(nr.Parent),
		"display_name":                 dcl.ValueOrEmptyString(nr.DisplayName),
		"description":                  dcl.ValueOrEmptyString(nr.Description),
		"create_time":                  dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":                  dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":                       dcl.ValueOrEmptyString(nr.Labels),
		"direct_member_count":          dcl.ValueOrEmptyString(nr.DirectMemberCount),
		"direct_member_count_per_type": dcl.ValueOrEmptyString(nr.DirectMemberCountPerType),
		"derived_aliases":              dcl.ValueOrEmptyString(nr.DerivedAliases),
		"dynamic_group_metadata":       dcl.ValueOrEmptyString(nr.DynamicGroupMetadata),
		"posix_groups":                 dcl.ValueOrEmptyString(nr.PosixGroups),
		"initial_group_config":         dcl.ValueOrEmptyString(nr.InitialGroupConfig),
	}
	return dcl.Nprintf("groups/{{name}}", params), nil
}

const GroupMaxPage = -1

type GroupList struct {
	Items []*Group

	nextToken string

	pageSize int32

	resource *Group
}

func (l *GroupList) HasNext() bool {
	return l.nextToken != ""
}

func (l *GroupList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listGroup(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListGroup(ctx context.Context, parent string) (*GroupList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListGroupWithMaxResults(ctx, parent, GroupMaxPage)

}

func (c *Client) ListGroupWithMaxResults(ctx context.Context, parent string, pageSize int32) (*GroupList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Group{
		Parent: &parent,
	}
	items, token, err := c.listGroup(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &GroupList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetGroup(ctx context.Context, r *Group) (*Group, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractGroupFields(r)

	b, err := c.getGroupRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalGroup(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeGroupNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractGroupFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteGroup(ctx context.Context, r *Group) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Group resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Group...")
	deleteOp := deleteGroupOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllGroup deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllGroup(ctx context.Context, parent string, filter func(*Group) bool) error {
	listObj, err := c.ListGroup(ctx, parent)
	if err != nil {
		return err
	}

	err = c.deleteAllGroup(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllGroup(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyGroup(ctx context.Context, rawDesired *Group, opts ...dcl.ApplyOption) (*Group, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Group
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyGroupHelper(c, ctx, rawDesired, opts...)
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

func applyGroupHelper(c *Client, ctx context.Context, rawDesired *Group, opts ...dcl.ApplyOption) (*Group, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyGroup...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractGroupFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.groupDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToGroupDiffs(c.Config, fieldDiffs, opts)
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
	var ops []groupApiOperation
	if create {
		ops = append(ops, &createGroupOperation{})
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

	// Handle diff for create request.
	if initial == nil {
		return applyGroupDiff(c, ctx, desired, rawDesired, ops, opts...)
	}

	// Handle eventually consistent update request.
	maxDiffAttempts := 5
	diffRetryIntervalSeconds := 2

	var newState *Group
	var diffErr error
	for i := 1; i <= maxDiffAttempts; i++ {
		newState, diffErr = applyGroupDiff(c, ctx, desired, rawDesired, ops, opts...)
		if diffErr == nil {
			return newState, nil
		} else if !errors.As(diffErr, &dcl.DiffAfterApplyError{}) {
			return newState, diffErr
		}

		t := time.NewTimer(time.Duration(diffRetryIntervalSeconds) * time.Second)
		select {
		case <-ctx.Done():
			t.Stop()
			c.Config.Logger.InfoWithContext(ctx, "Diff operation canceled by context")
			return newState, diffErr
		case <-t.C:
		}
	} // end of diffAttempt loop
	return newState, diffErr
}

func applyGroupDiff(c *Client, ctx context.Context, desired *Group, rawDesired *Group, ops []groupApiOperation, opts ...dcl.ApplyOption) (*Group, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetGroup(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createGroupOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapGroup(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeGroupNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeGroupNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeGroupDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractGroupFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractGroupFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffGroup(c, newDesired, newState)
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
