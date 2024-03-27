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
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type GrpcRoute struct {
	Name        *string           `json:"name"`
	CreateTime  *string           `json:"createTime"`
	UpdateTime  *string           `json:"updateTime"`
	Labels      map[string]string `json:"labels"`
	Description *string           `json:"description"`
	Hostnames   []string          `json:"hostnames"`
	Meshes      []string          `json:"meshes"`
	Gateways    []string          `json:"gateways"`
	Rules       []GrpcRouteRules  `json:"rules"`
	Project     *string           `json:"project"`
	Location    *string           `json:"location"`
	SelfLink    *string           `json:"selfLink"`
}

func (r *GrpcRoute) String() string {
	return dcl.SprintResource(r)
}

// The enum GrpcRouteRulesMatchesMethodTypeEnum.
type GrpcRouteRulesMatchesMethodTypeEnum string

// GrpcRouteRulesMatchesMethodTypeEnumRef returns a *GrpcRouteRulesMatchesMethodTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GrpcRouteRulesMatchesMethodTypeEnumRef(s string) *GrpcRouteRulesMatchesMethodTypeEnum {
	v := GrpcRouteRulesMatchesMethodTypeEnum(s)
	return &v
}

func (v GrpcRouteRulesMatchesMethodTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"TYPE_UNSPECIFIED", "EXACT", "REGULAR_EXPRESSION"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GrpcRouteRulesMatchesMethodTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum GrpcRouteRulesMatchesHeadersTypeEnum.
type GrpcRouteRulesMatchesHeadersTypeEnum string

// GrpcRouteRulesMatchesHeadersTypeEnumRef returns a *GrpcRouteRulesMatchesHeadersTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func GrpcRouteRulesMatchesHeadersTypeEnumRef(s string) *GrpcRouteRulesMatchesHeadersTypeEnum {
	v := GrpcRouteRulesMatchesHeadersTypeEnum(s)
	return &v
}

func (v GrpcRouteRulesMatchesHeadersTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MATCH_TYPE_UNSPECIFIED", "MATCH_ANY", "MATCH_ALL"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "GrpcRouteRulesMatchesHeadersTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type GrpcRouteRules struct {
	empty   bool                    `json:"-"`
	Matches []GrpcRouteRulesMatches `json:"matches"`
	Action  *GrpcRouteRulesAction   `json:"action"`
}

type jsonGrpcRouteRules GrpcRouteRules

func (r *GrpcRouteRules) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRules
	} else {

		r.Matches = res.Matches

		r.Action = res.Action

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRules *GrpcRouteRules = &GrpcRouteRules{empty: true}

func (r *GrpcRouteRules) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRules) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesMatches struct {
	empty   bool                           `json:"-"`
	Method  *GrpcRouteRulesMatchesMethod   `json:"method"`
	Headers []GrpcRouteRulesMatchesHeaders `json:"headers"`
}

type jsonGrpcRouteRulesMatches GrpcRouteRulesMatches

func (r *GrpcRouteRulesMatches) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesMatches
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesMatches
	} else {

		r.Method = res.Method

		r.Headers = res.Headers

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesMatches is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesMatches *GrpcRouteRulesMatches = &GrpcRouteRulesMatches{empty: true}

func (r *GrpcRouteRulesMatches) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesMatches) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesMatches) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesMatchesMethod struct {
	empty         bool                                 `json:"-"`
	Type          *GrpcRouteRulesMatchesMethodTypeEnum `json:"type"`
	GrpcService   *string                              `json:"grpcService"`
	GrpcMethod    *string                              `json:"grpcMethod"`
	CaseSensitive *bool                                `json:"caseSensitive"`
}

type jsonGrpcRouteRulesMatchesMethod GrpcRouteRulesMatchesMethod

func (r *GrpcRouteRulesMatchesMethod) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesMatchesMethod
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesMatchesMethod
	} else {

		r.Type = res.Type

		r.GrpcService = res.GrpcService

		r.GrpcMethod = res.GrpcMethod

		r.CaseSensitive = res.CaseSensitive

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesMatchesMethod is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesMatchesMethod *GrpcRouteRulesMatchesMethod = &GrpcRouteRulesMatchesMethod{empty: true}

func (r *GrpcRouteRulesMatchesMethod) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesMatchesMethod) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesMatchesMethod) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesMatchesHeaders struct {
	empty bool                                  `json:"-"`
	Type  *GrpcRouteRulesMatchesHeadersTypeEnum `json:"type"`
	Key   *string                               `json:"key"`
	Value *string                               `json:"value"`
}

type jsonGrpcRouteRulesMatchesHeaders GrpcRouteRulesMatchesHeaders

func (r *GrpcRouteRulesMatchesHeaders) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesMatchesHeaders
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesMatchesHeaders
	} else {

		r.Type = res.Type

		r.Key = res.Key

		r.Value = res.Value

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesMatchesHeaders is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesMatchesHeaders *GrpcRouteRulesMatchesHeaders = &GrpcRouteRulesMatchesHeaders{empty: true}

func (r *GrpcRouteRulesMatchesHeaders) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesMatchesHeaders) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesMatchesHeaders) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesAction struct {
	empty                bool                                      `json:"-"`
	Destinations         []GrpcRouteRulesActionDestinations        `json:"destinations"`
	FaultInjectionPolicy *GrpcRouteRulesActionFaultInjectionPolicy `json:"faultInjectionPolicy"`
	Timeout              *string                                   `json:"timeout"`
	RetryPolicy          *GrpcRouteRulesActionRetryPolicy          `json:"retryPolicy"`
}

type jsonGrpcRouteRulesAction GrpcRouteRulesAction

func (r *GrpcRouteRulesAction) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesAction
	} else {

		r.Destinations = res.Destinations

		r.FaultInjectionPolicy = res.FaultInjectionPolicy

		r.Timeout = res.Timeout

		r.RetryPolicy = res.RetryPolicy

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesAction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesAction *GrpcRouteRulesAction = &GrpcRouteRulesAction{empty: true}

func (r *GrpcRouteRulesAction) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesAction) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesActionDestinations struct {
	empty       bool    `json:"-"`
	Weight      *int64  `json:"weight"`
	ServiceName *string `json:"serviceName"`
}

type jsonGrpcRouteRulesActionDestinations GrpcRouteRulesActionDestinations

func (r *GrpcRouteRulesActionDestinations) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesActionDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesActionDestinations
	} else {

		r.Weight = res.Weight

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesActionDestinations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesActionDestinations *GrpcRouteRulesActionDestinations = &GrpcRouteRulesActionDestinations{empty: true}

func (r *GrpcRouteRulesActionDestinations) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesActionDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesActionDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesActionFaultInjectionPolicy struct {
	empty bool                                           `json:"-"`
	Delay *GrpcRouteRulesActionFaultInjectionPolicyDelay `json:"delay"`
	Abort *GrpcRouteRulesActionFaultInjectionPolicyAbort `json:"abort"`
}

type jsonGrpcRouteRulesActionFaultInjectionPolicy GrpcRouteRulesActionFaultInjectionPolicy

func (r *GrpcRouteRulesActionFaultInjectionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesActionFaultInjectionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesActionFaultInjectionPolicy
	} else {

		r.Delay = res.Delay

		r.Abort = res.Abort

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesActionFaultInjectionPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesActionFaultInjectionPolicy *GrpcRouteRulesActionFaultInjectionPolicy = &GrpcRouteRulesActionFaultInjectionPolicy{empty: true}

func (r *GrpcRouteRulesActionFaultInjectionPolicy) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesActionFaultInjectionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesActionFaultInjectionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesActionFaultInjectionPolicyDelay struct {
	empty      bool    `json:"-"`
	FixedDelay *string `json:"fixedDelay"`
	Percentage *int64  `json:"percentage"`
}

type jsonGrpcRouteRulesActionFaultInjectionPolicyDelay GrpcRouteRulesActionFaultInjectionPolicyDelay

func (r *GrpcRouteRulesActionFaultInjectionPolicyDelay) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesActionFaultInjectionPolicyDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesActionFaultInjectionPolicyDelay
	} else {

		r.FixedDelay = res.FixedDelay

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesActionFaultInjectionPolicyDelay is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesActionFaultInjectionPolicyDelay *GrpcRouteRulesActionFaultInjectionPolicyDelay = &GrpcRouteRulesActionFaultInjectionPolicyDelay{empty: true}

func (r *GrpcRouteRulesActionFaultInjectionPolicyDelay) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesActionFaultInjectionPolicyDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesActionFaultInjectionPolicyDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesActionFaultInjectionPolicyAbort struct {
	empty      bool   `json:"-"`
	HttpStatus *int64 `json:"httpStatus"`
	Percentage *int64 `json:"percentage"`
}

type jsonGrpcRouteRulesActionFaultInjectionPolicyAbort GrpcRouteRulesActionFaultInjectionPolicyAbort

func (r *GrpcRouteRulesActionFaultInjectionPolicyAbort) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesActionFaultInjectionPolicyAbort
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesActionFaultInjectionPolicyAbort
	} else {

		r.HttpStatus = res.HttpStatus

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesActionFaultInjectionPolicyAbort is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesActionFaultInjectionPolicyAbort *GrpcRouteRulesActionFaultInjectionPolicyAbort = &GrpcRouteRulesActionFaultInjectionPolicyAbort{empty: true}

func (r *GrpcRouteRulesActionFaultInjectionPolicyAbort) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesActionFaultInjectionPolicyAbort) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesActionFaultInjectionPolicyAbort) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type GrpcRouteRulesActionRetryPolicy struct {
	empty           bool     `json:"-"`
	RetryConditions []string `json:"retryConditions"`
	NumRetries      *int64   `json:"numRetries"`
}

type jsonGrpcRouteRulesActionRetryPolicy GrpcRouteRulesActionRetryPolicy

func (r *GrpcRouteRulesActionRetryPolicy) UnmarshalJSON(data []byte) error {
	var res jsonGrpcRouteRulesActionRetryPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyGrpcRouteRulesActionRetryPolicy
	} else {

		r.RetryConditions = res.RetryConditions

		r.NumRetries = res.NumRetries

	}
	return nil
}

// This object is used to assert a desired state where this GrpcRouteRulesActionRetryPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyGrpcRouteRulesActionRetryPolicy *GrpcRouteRulesActionRetryPolicy = &GrpcRouteRulesActionRetryPolicy{empty: true}

func (r *GrpcRouteRulesActionRetryPolicy) Empty() bool {
	return r.empty
}

func (r *GrpcRouteRulesActionRetryPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *GrpcRouteRulesActionRetryPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *GrpcRoute) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "GrpcRoute",
		Version: "alpha",
	}
}

func (r *GrpcRoute) ID() (string, error) {
	if err := extractGrpcRouteFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":      dcl.ValueOrEmptyString(nr.Labels),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"hostnames":   dcl.ValueOrEmptyString(nr.Hostnames),
		"meshes":      dcl.ValueOrEmptyString(nr.Meshes),
		"gateways":    dcl.ValueOrEmptyString(nr.Gateways),
		"rules":       dcl.ValueOrEmptyString(nr.Rules),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
		"self_link":   dcl.ValueOrEmptyString(nr.SelfLink),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/grpcRoutes/{{name}}", params), nil
}

const GrpcRouteMaxPage = -1

type GrpcRouteList struct {
	Items []*GrpcRoute

	nextToken string

	pageSize int32

	resource *GrpcRoute
}

func (l *GrpcRouteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *GrpcRouteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listGrpcRoute(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListGrpcRoute(ctx context.Context, project, location string) (*GrpcRouteList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListGrpcRouteWithMaxResults(ctx, project, location, GrpcRouteMaxPage)

}

func (c *Client) ListGrpcRouteWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*GrpcRouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &GrpcRoute{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listGrpcRoute(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &GrpcRouteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetGrpcRoute(ctx context.Context, r *GrpcRoute) (*GrpcRoute, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractGrpcRouteFields(r)

	b, err := c.getGrpcRouteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalGrpcRoute(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeGrpcRouteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractGrpcRouteFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteGrpcRoute(ctx context.Context, r *GrpcRoute) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("GrpcRoute resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting GrpcRoute...")
	deleteOp := deleteGrpcRouteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllGrpcRoute deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllGrpcRoute(ctx context.Context, project, location string, filter func(*GrpcRoute) bool) error {
	listObj, err := c.ListGrpcRoute(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllGrpcRoute(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllGrpcRoute(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyGrpcRoute(ctx context.Context, rawDesired *GrpcRoute, opts ...dcl.ApplyOption) (*GrpcRoute, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *GrpcRoute
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyGrpcRouteHelper(c, ctx, rawDesired, opts...)
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

func applyGrpcRouteHelper(c *Client, ctx context.Context, rawDesired *GrpcRoute, opts ...dcl.ApplyOption) (*GrpcRoute, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyGrpcRoute...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractGrpcRouteFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.grpcRouteDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToGrpcRouteDiffs(c.Config, fieldDiffs, opts)
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
	var ops []grpcRouteApiOperation
	if create {
		ops = append(ops, &createGrpcRouteOperation{})
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
	return applyGrpcRouteDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyGrpcRouteDiff(c *Client, ctx context.Context, desired *GrpcRoute, rawDesired *GrpcRoute, ops []grpcRouteApiOperation, opts ...dcl.ApplyOption) (*GrpcRoute, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetGrpcRoute(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createGrpcRouteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapGrpcRoute(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeGrpcRouteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeGrpcRouteNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeGrpcRouteDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractGrpcRouteFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractGrpcRouteFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffGrpcRoute(c, newDesired, newState)
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
