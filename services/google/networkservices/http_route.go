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
package networkservices

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type HttpRoute struct {
	Name        *string           `json:"name"`
	Description *string           `json:"description"`
	CreateTime  *string           `json:"createTime"`
	UpdateTime  *string           `json:"updateTime"`
	Hostnames   []string          `json:"hostnames"`
	Meshes      []string          `json:"meshes"`
	Gateways    []string          `json:"gateways"`
	Labels      map[string]string `json:"labels"`
	Rules       []HttpRouteRules  `json:"rules"`
	Project     *string           `json:"project"`
	Location    *string           `json:"location"`
	SelfLink    *string           `json:"selfLink"`
}

func (r *HttpRoute) String() string {
	return dcl.SprintResource(r)
}

// The enum HttpRouteRulesActionRedirectResponseCodeEnum.
type HttpRouteRulesActionRedirectResponseCodeEnum string

// HttpRouteRulesActionRedirectResponseCodeEnumRef returns a *HttpRouteRulesActionRedirectResponseCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func HttpRouteRulesActionRedirectResponseCodeEnumRef(s string) *HttpRouteRulesActionRedirectResponseCodeEnum {
	v := HttpRouteRulesActionRedirectResponseCodeEnum(s)
	return &v
}

func (v HttpRouteRulesActionRedirectResponseCodeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "HttpRouteRulesActionRedirectResponseCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type HttpRouteRules struct {
	empty   bool                    `json:"-"`
	Matches []HttpRouteRulesMatches `json:"matches"`
	Action  *HttpRouteRulesAction   `json:"action"`
}

type jsonHttpRouteRules HttpRouteRules

func (r *HttpRouteRules) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRules
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRules
	} else {

		r.Matches = res.Matches

		r.Action = res.Action

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRules is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRules *HttpRouteRules = &HttpRouteRules{empty: true}

func (r *HttpRouteRules) Empty() bool {
	return r.empty
}

func (r *HttpRouteRules) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRules) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesMatches struct {
	empty           bool                                   `json:"-"`
	FullPathMatch   *string                                `json:"fullPathMatch"`
	PrefixMatch     *string                                `json:"prefixMatch"`
	RegexMatch      *string                                `json:"regexMatch"`
	IgnoreCase      *bool                                  `json:"ignoreCase"`
	Headers         []HttpRouteRulesMatchesHeaders         `json:"headers"`
	QueryParameters []HttpRouteRulesMatchesQueryParameters `json:"queryParameters"`
}

type jsonHttpRouteRulesMatches HttpRouteRulesMatches

func (r *HttpRouteRulesMatches) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesMatches
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesMatches
	} else {

		r.FullPathMatch = res.FullPathMatch

		r.PrefixMatch = res.PrefixMatch

		r.RegexMatch = res.RegexMatch

		r.IgnoreCase = res.IgnoreCase

		r.Headers = res.Headers

		r.QueryParameters = res.QueryParameters

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesMatches is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesMatches *HttpRouteRulesMatches = &HttpRouteRulesMatches{empty: true}

func (r *HttpRouteRulesMatches) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesMatches) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesMatches) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesMatchesHeaders struct {
	empty        bool                                    `json:"-"`
	Header       *string                                 `json:"header"`
	ExactMatch   *string                                 `json:"exactMatch"`
	RegexMatch   *string                                 `json:"regexMatch"`
	PrefixMatch  *string                                 `json:"prefixMatch"`
	PresentMatch *bool                                   `json:"presentMatch"`
	SuffixMatch  *string                                 `json:"suffixMatch"`
	RangeMatch   *HttpRouteRulesMatchesHeadersRangeMatch `json:"rangeMatch"`
	InvertMatch  *bool                                   `json:"invertMatch"`
}

type jsonHttpRouteRulesMatchesHeaders HttpRouteRulesMatchesHeaders

func (r *HttpRouteRulesMatchesHeaders) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesMatchesHeaders
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesMatchesHeaders
	} else {

		r.Header = res.Header

		r.ExactMatch = res.ExactMatch

		r.RegexMatch = res.RegexMatch

		r.PrefixMatch = res.PrefixMatch

		r.PresentMatch = res.PresentMatch

		r.SuffixMatch = res.SuffixMatch

		r.RangeMatch = res.RangeMatch

		r.InvertMatch = res.InvertMatch

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesMatchesHeaders is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesMatchesHeaders *HttpRouteRulesMatchesHeaders = &HttpRouteRulesMatchesHeaders{empty: true}

func (r *HttpRouteRulesMatchesHeaders) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesMatchesHeaders) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesMatchesHeaders) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesMatchesHeadersRangeMatch struct {
	empty bool   `json:"-"`
	Start *int64 `json:"start"`
	End   *int64 `json:"end"`
}

type jsonHttpRouteRulesMatchesHeadersRangeMatch HttpRouteRulesMatchesHeadersRangeMatch

func (r *HttpRouteRulesMatchesHeadersRangeMatch) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesMatchesHeadersRangeMatch
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesMatchesHeadersRangeMatch
	} else {

		r.Start = res.Start

		r.End = res.End

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesMatchesHeadersRangeMatch is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesMatchesHeadersRangeMatch *HttpRouteRulesMatchesHeadersRangeMatch = &HttpRouteRulesMatchesHeadersRangeMatch{empty: true}

func (r *HttpRouteRulesMatchesHeadersRangeMatch) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesMatchesHeadersRangeMatch) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesMatchesHeadersRangeMatch) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesMatchesQueryParameters struct {
	empty          bool    `json:"-"`
	QueryParameter *string `json:"queryParameter"`
	ExactMatch     *string `json:"exactMatch"`
	RegexMatch     *string `json:"regexMatch"`
	PresentMatch   *bool   `json:"presentMatch"`
}

type jsonHttpRouteRulesMatchesQueryParameters HttpRouteRulesMatchesQueryParameters

func (r *HttpRouteRulesMatchesQueryParameters) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesMatchesQueryParameters
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesMatchesQueryParameters
	} else {

		r.QueryParameter = res.QueryParameter

		r.ExactMatch = res.ExactMatch

		r.RegexMatch = res.RegexMatch

		r.PresentMatch = res.PresentMatch

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesMatchesQueryParameters is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesMatchesQueryParameters *HttpRouteRulesMatchesQueryParameters = &HttpRouteRulesMatchesQueryParameters{empty: true}

func (r *HttpRouteRulesMatchesQueryParameters) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesMatchesQueryParameters) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesMatchesQueryParameters) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesAction struct {
	empty                  bool                                        `json:"-"`
	Destinations           []HttpRouteRulesActionDestinations          `json:"destinations"`
	Redirect               *HttpRouteRulesActionRedirect               `json:"redirect"`
	FaultInjectionPolicy   *HttpRouteRulesActionFaultInjectionPolicy   `json:"faultInjectionPolicy"`
	RequestHeaderModifier  *HttpRouteRulesActionRequestHeaderModifier  `json:"requestHeaderModifier"`
	ResponseHeaderModifier *HttpRouteRulesActionResponseHeaderModifier `json:"responseHeaderModifier"`
	UrlRewrite             *HttpRouteRulesActionUrlRewrite             `json:"urlRewrite"`
	Timeout                *string                                     `json:"timeout"`
	RetryPolicy            *HttpRouteRulesActionRetryPolicy            `json:"retryPolicy"`
	RequestMirrorPolicy    *HttpRouteRulesActionRequestMirrorPolicy    `json:"requestMirrorPolicy"`
	CorsPolicy             *HttpRouteRulesActionCorsPolicy             `json:"corsPolicy"`
}

type jsonHttpRouteRulesAction HttpRouteRulesAction

func (r *HttpRouteRulesAction) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesAction
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesAction
	} else {

		r.Destinations = res.Destinations

		r.Redirect = res.Redirect

		r.FaultInjectionPolicy = res.FaultInjectionPolicy

		r.RequestHeaderModifier = res.RequestHeaderModifier

		r.ResponseHeaderModifier = res.ResponseHeaderModifier

		r.UrlRewrite = res.UrlRewrite

		r.Timeout = res.Timeout

		r.RetryPolicy = res.RetryPolicy

		r.RequestMirrorPolicy = res.RequestMirrorPolicy

		r.CorsPolicy = res.CorsPolicy

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesAction is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesAction *HttpRouteRulesAction = &HttpRouteRulesAction{empty: true}

func (r *HttpRouteRulesAction) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesAction) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesAction) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionDestinations struct {
	empty       bool    `json:"-"`
	Weight      *int64  `json:"weight"`
	ServiceName *string `json:"serviceName"`
}

type jsonHttpRouteRulesActionDestinations HttpRouteRulesActionDestinations

func (r *HttpRouteRulesActionDestinations) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionDestinations
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionDestinations
	} else {

		r.Weight = res.Weight

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionDestinations is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionDestinations *HttpRouteRulesActionDestinations = &HttpRouteRulesActionDestinations{empty: true}

func (r *HttpRouteRulesActionDestinations) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionDestinations) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionDestinations) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionRedirect struct {
	empty         bool                                          `json:"-"`
	HostRedirect  *string                                       `json:"hostRedirect"`
	PathRedirect  *string                                       `json:"pathRedirect"`
	PrefixRewrite *string                                       `json:"prefixRewrite"`
	ResponseCode  *HttpRouteRulesActionRedirectResponseCodeEnum `json:"responseCode"`
	HttpsRedirect *bool                                         `json:"httpsRedirect"`
	StripQuery    *bool                                         `json:"stripQuery"`
	PortRedirect  *int64                                        `json:"portRedirect"`
}

type jsonHttpRouteRulesActionRedirect HttpRouteRulesActionRedirect

func (r *HttpRouteRulesActionRedirect) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionRedirect
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionRedirect
	} else {

		r.HostRedirect = res.HostRedirect

		r.PathRedirect = res.PathRedirect

		r.PrefixRewrite = res.PrefixRewrite

		r.ResponseCode = res.ResponseCode

		r.HttpsRedirect = res.HttpsRedirect

		r.StripQuery = res.StripQuery

		r.PortRedirect = res.PortRedirect

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionRedirect is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionRedirect *HttpRouteRulesActionRedirect = &HttpRouteRulesActionRedirect{empty: true}

func (r *HttpRouteRulesActionRedirect) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionRedirect) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionRedirect) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionFaultInjectionPolicy struct {
	empty bool                                           `json:"-"`
	Delay *HttpRouteRulesActionFaultInjectionPolicyDelay `json:"delay"`
	Abort *HttpRouteRulesActionFaultInjectionPolicyAbort `json:"abort"`
}

type jsonHttpRouteRulesActionFaultInjectionPolicy HttpRouteRulesActionFaultInjectionPolicy

func (r *HttpRouteRulesActionFaultInjectionPolicy) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionFaultInjectionPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionFaultInjectionPolicy
	} else {

		r.Delay = res.Delay

		r.Abort = res.Abort

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionFaultInjectionPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionFaultInjectionPolicy *HttpRouteRulesActionFaultInjectionPolicy = &HttpRouteRulesActionFaultInjectionPolicy{empty: true}

func (r *HttpRouteRulesActionFaultInjectionPolicy) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionFaultInjectionPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionFaultInjectionPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionFaultInjectionPolicyDelay struct {
	empty      bool    `json:"-"`
	FixedDelay *string `json:"fixedDelay"`
	Percentage *int64  `json:"percentage"`
}

type jsonHttpRouteRulesActionFaultInjectionPolicyDelay HttpRouteRulesActionFaultInjectionPolicyDelay

func (r *HttpRouteRulesActionFaultInjectionPolicyDelay) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionFaultInjectionPolicyDelay
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionFaultInjectionPolicyDelay
	} else {

		r.FixedDelay = res.FixedDelay

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionFaultInjectionPolicyDelay is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionFaultInjectionPolicyDelay *HttpRouteRulesActionFaultInjectionPolicyDelay = &HttpRouteRulesActionFaultInjectionPolicyDelay{empty: true}

func (r *HttpRouteRulesActionFaultInjectionPolicyDelay) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionFaultInjectionPolicyDelay) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionFaultInjectionPolicyDelay) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionFaultInjectionPolicyAbort struct {
	empty      bool   `json:"-"`
	HttpStatus *int64 `json:"httpStatus"`
	Percentage *int64 `json:"percentage"`
}

type jsonHttpRouteRulesActionFaultInjectionPolicyAbort HttpRouteRulesActionFaultInjectionPolicyAbort

func (r *HttpRouteRulesActionFaultInjectionPolicyAbort) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionFaultInjectionPolicyAbort
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionFaultInjectionPolicyAbort
	} else {

		r.HttpStatus = res.HttpStatus

		r.Percentage = res.Percentage

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionFaultInjectionPolicyAbort is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionFaultInjectionPolicyAbort *HttpRouteRulesActionFaultInjectionPolicyAbort = &HttpRouteRulesActionFaultInjectionPolicyAbort{empty: true}

func (r *HttpRouteRulesActionFaultInjectionPolicyAbort) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionFaultInjectionPolicyAbort) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionFaultInjectionPolicyAbort) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionRequestHeaderModifier struct {
	empty  bool              `json:"-"`
	Set    map[string]string `json:"set"`
	Add    map[string]string `json:"add"`
	Remove []string          `json:"remove"`
}

type jsonHttpRouteRulesActionRequestHeaderModifier HttpRouteRulesActionRequestHeaderModifier

func (r *HttpRouteRulesActionRequestHeaderModifier) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionRequestHeaderModifier
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionRequestHeaderModifier
	} else {

		r.Set = res.Set

		r.Add = res.Add

		r.Remove = res.Remove

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionRequestHeaderModifier is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionRequestHeaderModifier *HttpRouteRulesActionRequestHeaderModifier = &HttpRouteRulesActionRequestHeaderModifier{empty: true}

func (r *HttpRouteRulesActionRequestHeaderModifier) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionRequestHeaderModifier) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionRequestHeaderModifier) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionResponseHeaderModifier struct {
	empty  bool              `json:"-"`
	Set    map[string]string `json:"set"`
	Add    map[string]string `json:"add"`
	Remove []string          `json:"remove"`
}

type jsonHttpRouteRulesActionResponseHeaderModifier HttpRouteRulesActionResponseHeaderModifier

func (r *HttpRouteRulesActionResponseHeaderModifier) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionResponseHeaderModifier
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionResponseHeaderModifier
	} else {

		r.Set = res.Set

		r.Add = res.Add

		r.Remove = res.Remove

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionResponseHeaderModifier is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionResponseHeaderModifier *HttpRouteRulesActionResponseHeaderModifier = &HttpRouteRulesActionResponseHeaderModifier{empty: true}

func (r *HttpRouteRulesActionResponseHeaderModifier) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionResponseHeaderModifier) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionResponseHeaderModifier) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionUrlRewrite struct {
	empty             bool    `json:"-"`
	PathPrefixRewrite *string `json:"pathPrefixRewrite"`
	HostRewrite       *string `json:"hostRewrite"`
}

type jsonHttpRouteRulesActionUrlRewrite HttpRouteRulesActionUrlRewrite

func (r *HttpRouteRulesActionUrlRewrite) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionUrlRewrite
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionUrlRewrite
	} else {

		r.PathPrefixRewrite = res.PathPrefixRewrite

		r.HostRewrite = res.HostRewrite

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionUrlRewrite is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionUrlRewrite *HttpRouteRulesActionUrlRewrite = &HttpRouteRulesActionUrlRewrite{empty: true}

func (r *HttpRouteRulesActionUrlRewrite) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionUrlRewrite) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionUrlRewrite) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionRetryPolicy struct {
	empty           bool     `json:"-"`
	RetryConditions []string `json:"retryConditions"`
	NumRetries      *int64   `json:"numRetries"`
	PerTryTimeout   *string  `json:"perTryTimeout"`
}

type jsonHttpRouteRulesActionRetryPolicy HttpRouteRulesActionRetryPolicy

func (r *HttpRouteRulesActionRetryPolicy) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionRetryPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionRetryPolicy
	} else {

		r.RetryConditions = res.RetryConditions

		r.NumRetries = res.NumRetries

		r.PerTryTimeout = res.PerTryTimeout

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionRetryPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionRetryPolicy *HttpRouteRulesActionRetryPolicy = &HttpRouteRulesActionRetryPolicy{empty: true}

func (r *HttpRouteRulesActionRetryPolicy) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionRetryPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionRetryPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionRequestMirrorPolicy struct {
	empty       bool                                                `json:"-"`
	Destination *HttpRouteRulesActionRequestMirrorPolicyDestination `json:"destination"`
}

type jsonHttpRouteRulesActionRequestMirrorPolicy HttpRouteRulesActionRequestMirrorPolicy

func (r *HttpRouteRulesActionRequestMirrorPolicy) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionRequestMirrorPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionRequestMirrorPolicy
	} else {

		r.Destination = res.Destination

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionRequestMirrorPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionRequestMirrorPolicy *HttpRouteRulesActionRequestMirrorPolicy = &HttpRouteRulesActionRequestMirrorPolicy{empty: true}

func (r *HttpRouteRulesActionRequestMirrorPolicy) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionRequestMirrorPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionRequestMirrorPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionRequestMirrorPolicyDestination struct {
	empty       bool    `json:"-"`
	Weight      *int64  `json:"weight"`
	ServiceName *string `json:"serviceName"`
}

type jsonHttpRouteRulesActionRequestMirrorPolicyDestination HttpRouteRulesActionRequestMirrorPolicyDestination

func (r *HttpRouteRulesActionRequestMirrorPolicyDestination) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionRequestMirrorPolicyDestination
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionRequestMirrorPolicyDestination
	} else {

		r.Weight = res.Weight

		r.ServiceName = res.ServiceName

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionRequestMirrorPolicyDestination is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionRequestMirrorPolicyDestination *HttpRouteRulesActionRequestMirrorPolicyDestination = &HttpRouteRulesActionRequestMirrorPolicyDestination{empty: true}

func (r *HttpRouteRulesActionRequestMirrorPolicyDestination) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionRequestMirrorPolicyDestination) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionRequestMirrorPolicyDestination) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type HttpRouteRulesActionCorsPolicy struct {
	empty              bool     `json:"-"`
	AllowOrigins       []string `json:"allowOrigins"`
	AllowOriginRegexes []string `json:"allowOriginRegexes"`
	AllowMethods       []string `json:"allowMethods"`
	AllowHeaders       []string `json:"allowHeaders"`
	ExposeHeaders      []string `json:"exposeHeaders"`
	MaxAge             *string  `json:"maxAge"`
	AllowCredentials   *bool    `json:"allowCredentials"`
	Disabled           *bool    `json:"disabled"`
}

type jsonHttpRouteRulesActionCorsPolicy HttpRouteRulesActionCorsPolicy

func (r *HttpRouteRulesActionCorsPolicy) UnmarshalJSON(data []byte) error {
	var res jsonHttpRouteRulesActionCorsPolicy
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyHttpRouteRulesActionCorsPolicy
	} else {

		r.AllowOrigins = res.AllowOrigins

		r.AllowOriginRegexes = res.AllowOriginRegexes

		r.AllowMethods = res.AllowMethods

		r.AllowHeaders = res.AllowHeaders

		r.ExposeHeaders = res.ExposeHeaders

		r.MaxAge = res.MaxAge

		r.AllowCredentials = res.AllowCredentials

		r.Disabled = res.Disabled

	}
	return nil
}

// This object is used to assert a desired state where this HttpRouteRulesActionCorsPolicy is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyHttpRouteRulesActionCorsPolicy *HttpRouteRulesActionCorsPolicy = &HttpRouteRulesActionCorsPolicy{empty: true}

func (r *HttpRouteRulesActionCorsPolicy) Empty() bool {
	return r.empty
}

func (r *HttpRouteRulesActionCorsPolicy) String() string {
	return dcl.SprintResource(r)
}

func (r *HttpRouteRulesActionCorsPolicy) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *HttpRoute) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_services",
		Type:    "HttpRoute",
		Version: "networkservices",
	}
}

func (r *HttpRoute) ID() (string, error) {
	if err := extractHttpRouteFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":        dcl.ValueOrEmptyString(nr.Name),
		"description": dcl.ValueOrEmptyString(nr.Description),
		"create_time": dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time": dcl.ValueOrEmptyString(nr.UpdateTime),
		"hostnames":   dcl.ValueOrEmptyString(nr.Hostnames),
		"meshes":      dcl.ValueOrEmptyString(nr.Meshes),
		"gateways":    dcl.ValueOrEmptyString(nr.Gateways),
		"labels":      dcl.ValueOrEmptyString(nr.Labels),
		"rules":       dcl.ValueOrEmptyString(nr.Rules),
		"project":     dcl.ValueOrEmptyString(nr.Project),
		"location":    dcl.ValueOrEmptyString(nr.Location),
		"self_link":   dcl.ValueOrEmptyString(nr.SelfLink),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/httpRoutes/{{name}}", params), nil
}

const HttpRouteMaxPage = -1

type HttpRouteList struct {
	Items []*HttpRoute

	nextToken string

	pageSize int32

	resource *HttpRoute
}

func (l *HttpRouteList) HasNext() bool {
	return l.nextToken != ""
}

func (l *HttpRouteList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listHttpRoute(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListHttpRoute(ctx context.Context, project, location string) (*HttpRouteList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListHttpRouteWithMaxResults(ctx, project, location, HttpRouteMaxPage)

}

func (c *Client) ListHttpRouteWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*HttpRouteList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &HttpRoute{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listHttpRoute(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &HttpRouteList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetHttpRoute(ctx context.Context, r *HttpRoute) (*HttpRoute, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractHttpRouteFields(r)

	b, err := c.getHttpRouteRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalHttpRoute(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeHttpRouteNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractHttpRouteFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteHttpRoute(ctx context.Context, r *HttpRoute) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("HttpRoute resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting HttpRoute...")
	deleteOp := deleteHttpRouteOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllHttpRoute deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllHttpRoute(ctx context.Context, project, location string, filter func(*HttpRoute) bool) error {
	listObj, err := c.ListHttpRoute(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllHttpRoute(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllHttpRoute(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyHttpRoute(ctx context.Context, rawDesired *HttpRoute, opts ...dcl.ApplyOption) (*HttpRoute, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *HttpRoute
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyHttpRouteHelper(c, ctx, rawDesired, opts...)
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

func applyHttpRouteHelper(c *Client, ctx context.Context, rawDesired *HttpRoute, opts ...dcl.ApplyOption) (*HttpRoute, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyHttpRoute...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractHttpRouteFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.httpRouteDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToHttpRouteDiffs(c.Config, fieldDiffs, opts)
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
	var ops []httpRouteApiOperation
	if create {
		ops = append(ops, &createHttpRouteOperation{})
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
	return applyHttpRouteDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyHttpRouteDiff(c *Client, ctx context.Context, desired *HttpRoute, rawDesired *HttpRoute, ops []httpRouteApiOperation, opts ...dcl.ApplyOption) (*HttpRoute, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetHttpRoute(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createHttpRouteOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapHttpRoute(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeHttpRouteNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeHttpRouteNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeHttpRouteDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractHttpRouteFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractHttpRouteFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffHttpRoute(c, newDesired, newState)
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
