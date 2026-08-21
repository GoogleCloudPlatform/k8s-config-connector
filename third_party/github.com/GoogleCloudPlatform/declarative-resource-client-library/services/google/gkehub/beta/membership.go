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
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

type Membership struct {
	Endpoint           *MembershipEndpoint               `json:"endpoint"`
	Name               *string                           `json:"name"`
	Labels             map[string]string                 `json:"labels"`
	Description        *string                           `json:"description"`
	State              *MembershipState                  `json:"state"`
	CreateTime         *string                           `json:"createTime"`
	UpdateTime         *string                           `json:"updateTime"`
	DeleteTime         *string                           `json:"deleteTime"`
	ExternalId         *string                           `json:"externalId"`
	LastConnectionTime *string                           `json:"lastConnectionTime"`
	UniqueId           *string                           `json:"uniqueId"`
	Authority          *MembershipAuthority              `json:"authority"`
	InfrastructureType *MembershipInfrastructureTypeEnum `json:"infrastructureType"`
	Project            *string                           `json:"project"`
	Location           *string                           `json:"location"`
}

func (r *Membership) String() string {
	return dcl.SprintResource(r)
}

// The enum MembershipStateCodeEnum.
type MembershipStateCodeEnum string

// MembershipStateCodeEnumRef returns a *MembershipStateCodeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MembershipStateCodeEnumRef(s string) *MembershipStateCodeEnum {
	v := MembershipStateCodeEnum(s)
	return &v
}

func (v MembershipStateCodeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"CODE_UNSPECIFIED", "CREATING", "READY", "DELETING", "UPDATING", "SERVICE_UPDATING"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MembershipStateCodeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

// The enum MembershipInfrastructureTypeEnum.
type MembershipInfrastructureTypeEnum string

// MembershipInfrastructureTypeEnumRef returns a *MembershipInfrastructureTypeEnum with the value of string s
// If the empty string is provided, nil is returned.
func MembershipInfrastructureTypeEnumRef(s string) *MembershipInfrastructureTypeEnum {
	v := MembershipInfrastructureTypeEnum(s)
	return &v
}

func (v MembershipInfrastructureTypeEnum) Validate() error {
	if string(v) == "" {
		// Empty enum is okay.
		return nil
	}
	for _, s := range []string{"INFRASTRUCTURE_TYPE_UNSPECIFIED", "ON_PREM", "MULTI_CLOUD"} {
		if string(v) == s {
			return nil
		}
	}
	return &dcl.EnumInvalidError{
		Enum:  "MembershipInfrastructureTypeEnum",
		Value: string(v),
		Valid: []string{},
	}
}

type MembershipEndpoint struct {
	empty              bool                                  `json:"-"`
	GkeCluster         *MembershipEndpointGkeCluster         `json:"gkeCluster"`
	KubernetesMetadata *MembershipEndpointKubernetesMetadata `json:"kubernetesMetadata"`
	KubernetesResource *MembershipEndpointKubernetesResource `json:"kubernetesResource"`
}

type jsonMembershipEndpoint MembershipEndpoint

func (r *MembershipEndpoint) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpoint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpoint
	} else {

		r.GkeCluster = res.GkeCluster

		r.KubernetesMetadata = res.KubernetesMetadata

		r.KubernetesResource = res.KubernetesResource

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpoint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpoint *MembershipEndpoint = &MembershipEndpoint{empty: true}

func (r *MembershipEndpoint) Empty() bool {
	return r.empty
}

func (r *MembershipEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointGkeCluster struct {
	empty        bool    `json:"-"`
	ResourceLink *string `json:"resourceLink"`
}

type jsonMembershipEndpointGkeCluster MembershipEndpointGkeCluster

func (r *MembershipEndpointGkeCluster) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointGkeCluster
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointGkeCluster
	} else {

		r.ResourceLink = res.ResourceLink

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointGkeCluster is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointGkeCluster *MembershipEndpointGkeCluster = &MembershipEndpointGkeCluster{empty: true}

func (r *MembershipEndpointGkeCluster) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointGkeCluster) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointGkeCluster) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointKubernetesMetadata struct {
	empty                      bool    `json:"-"`
	KubernetesApiServerVersion *string `json:"kubernetesApiServerVersion"`
	NodeProviderId             *string `json:"nodeProviderId"`
	NodeCount                  *int64  `json:"nodeCount"`
	VcpuCount                  *int64  `json:"vcpuCount"`
	MemoryMb                   *int64  `json:"memoryMb"`
	UpdateTime                 *string `json:"updateTime"`
}

type jsonMembershipEndpointKubernetesMetadata MembershipEndpointKubernetesMetadata

func (r *MembershipEndpointKubernetesMetadata) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointKubernetesMetadata
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointKubernetesMetadata
	} else {

		r.KubernetesApiServerVersion = res.KubernetesApiServerVersion

		r.NodeProviderId = res.NodeProviderId

		r.NodeCount = res.NodeCount

		r.VcpuCount = res.VcpuCount

		r.MemoryMb = res.MemoryMb

		r.UpdateTime = res.UpdateTime

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointKubernetesMetadata is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointKubernetesMetadata *MembershipEndpointKubernetesMetadata = &MembershipEndpointKubernetesMetadata{empty: true}

func (r *MembershipEndpointKubernetesMetadata) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointKubernetesMetadata) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointKubernetesMetadata) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointKubernetesResource struct {
	empty                bool                                                      `json:"-"`
	MembershipCrManifest *string                                                   `json:"membershipCrManifest"`
	MembershipResources  []MembershipEndpointKubernetesResourceMembershipResources `json:"membershipResources"`
	ConnectResources     []MembershipEndpointKubernetesResourceConnectResources    `json:"connectResources"`
	ResourceOptions      *MembershipEndpointKubernetesResourceResourceOptions      `json:"resourceOptions"`
}

type jsonMembershipEndpointKubernetesResource MembershipEndpointKubernetesResource

func (r *MembershipEndpointKubernetesResource) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointKubernetesResource
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointKubernetesResource
	} else {

		r.MembershipCrManifest = res.MembershipCrManifest

		r.MembershipResources = res.MembershipResources

		r.ConnectResources = res.ConnectResources

		r.ResourceOptions = res.ResourceOptions

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointKubernetesResource is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointKubernetesResource *MembershipEndpointKubernetesResource = &MembershipEndpointKubernetesResource{empty: true}

func (r *MembershipEndpointKubernetesResource) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointKubernetesResource) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointKubernetesResource) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointKubernetesResourceMembershipResources struct {
	empty         bool    `json:"-"`
	Manifest      *string `json:"manifest"`
	ClusterScoped *bool   `json:"clusterScoped"`
}

type jsonMembershipEndpointKubernetesResourceMembershipResources MembershipEndpointKubernetesResourceMembershipResources

func (r *MembershipEndpointKubernetesResourceMembershipResources) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointKubernetesResourceMembershipResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointKubernetesResourceMembershipResources
	} else {

		r.Manifest = res.Manifest

		r.ClusterScoped = res.ClusterScoped

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointKubernetesResourceMembershipResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointKubernetesResourceMembershipResources *MembershipEndpointKubernetesResourceMembershipResources = &MembershipEndpointKubernetesResourceMembershipResources{empty: true}

func (r *MembershipEndpointKubernetesResourceMembershipResources) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointKubernetesResourceMembershipResources) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointKubernetesResourceMembershipResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointKubernetesResourceConnectResources struct {
	empty         bool    `json:"-"`
	Manifest      *string `json:"manifest"`
	ClusterScoped *bool   `json:"clusterScoped"`
}

type jsonMembershipEndpointKubernetesResourceConnectResources MembershipEndpointKubernetesResourceConnectResources

func (r *MembershipEndpointKubernetesResourceConnectResources) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointKubernetesResourceConnectResources
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointKubernetesResourceConnectResources
	} else {

		r.Manifest = res.Manifest

		r.ClusterScoped = res.ClusterScoped

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointKubernetesResourceConnectResources is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointKubernetesResourceConnectResources *MembershipEndpointKubernetesResourceConnectResources = &MembershipEndpointKubernetesResourceConnectResources{empty: true}

func (r *MembershipEndpointKubernetesResourceConnectResources) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointKubernetesResourceConnectResources) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointKubernetesResourceConnectResources) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipEndpointKubernetesResourceResourceOptions struct {
	empty          bool    `json:"-"`
	ConnectVersion *string `json:"connectVersion"`
	V1Beta1Crd     *bool   `json:"v1beta1Crd"`
}

type jsonMembershipEndpointKubernetesResourceResourceOptions MembershipEndpointKubernetesResourceResourceOptions

func (r *MembershipEndpointKubernetesResourceResourceOptions) UnmarshalJSON(data []byte) error {
	var res jsonMembershipEndpointKubernetesResourceResourceOptions
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipEndpointKubernetesResourceResourceOptions
	} else {

		r.ConnectVersion = res.ConnectVersion

		r.V1Beta1Crd = res.V1Beta1Crd

	}
	return nil
}

// This object is used to assert a desired state where this MembershipEndpointKubernetesResourceResourceOptions is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipEndpointKubernetesResourceResourceOptions *MembershipEndpointKubernetesResourceResourceOptions = &MembershipEndpointKubernetesResourceResourceOptions{empty: true}

func (r *MembershipEndpointKubernetesResourceResourceOptions) Empty() bool {
	return r.empty
}

func (r *MembershipEndpointKubernetesResourceResourceOptions) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipEndpointKubernetesResourceResourceOptions) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipState struct {
	empty bool                     `json:"-"`
	Code  *MembershipStateCodeEnum `json:"code"`
}

type jsonMembershipState MembershipState

func (r *MembershipState) UnmarshalJSON(data []byte) error {
	var res jsonMembershipState
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipState
	} else {

		r.Code = res.Code

	}
	return nil
}

// This object is used to assert a desired state where this MembershipState is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipState *MembershipState = &MembershipState{empty: true}

func (r *MembershipState) Empty() bool {
	return r.empty
}

func (r *MembershipState) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipState) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type MembershipAuthority struct {
	empty                bool    `json:"-"`
	Issuer               *string `json:"issuer"`
	WorkloadIdentityPool *string `json:"workloadIdentityPool"`
	IdentityProvider     *string `json:"identityProvider"`
}

type jsonMembershipAuthority MembershipAuthority

func (r *MembershipAuthority) UnmarshalJSON(data []byte) error {
	var res jsonMembershipAuthority
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyMembershipAuthority
	} else {

		r.Issuer = res.Issuer

		r.WorkloadIdentityPool = res.WorkloadIdentityPool

		r.IdentityProvider = res.IdentityProvider

	}
	return nil
}

// This object is used to assert a desired state where this MembershipAuthority is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyMembershipAuthority *MembershipAuthority = &MembershipAuthority{empty: true}

func (r *MembershipAuthority) Empty() bool {
	return r.empty
}

func (r *MembershipAuthority) String() string {
	return dcl.SprintResource(r)
}

func (r *MembershipAuthority) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *Membership) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "gke_hub",
		Type:    "Membership",
		Version: "beta",
	}
}

func (r *Membership) ID() (string, error) {
	if err := extractMembershipFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"endpoint":             dcl.ValueOrEmptyString(nr.Endpoint),
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"labels":               dcl.ValueOrEmptyString(nr.Labels),
		"description":          dcl.ValueOrEmptyString(nr.Description),
		"state":                dcl.ValueOrEmptyString(nr.State),
		"create_time":          dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":          dcl.ValueOrEmptyString(nr.UpdateTime),
		"delete_time":          dcl.ValueOrEmptyString(nr.DeleteTime),
		"external_id":          dcl.ValueOrEmptyString(nr.ExternalId),
		"last_connection_time": dcl.ValueOrEmptyString(nr.LastConnectionTime),
		"unique_id":            dcl.ValueOrEmptyString(nr.UniqueId),
		"authority":            dcl.ValueOrEmptyString(nr.Authority),
		"infrastructure_type":  dcl.ValueOrEmptyString(nr.InfrastructureType),
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/memberships/{{name}}", params), nil
}

const MembershipMaxPage = -1

type MembershipList struct {
	Items []*Membership

	nextToken string

	pageSize int32

	resource *Membership
}

func (l *MembershipList) HasNext() bool {
	return l.nextToken != ""
}

func (l *MembershipList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listMembership(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListMembership(ctx context.Context, project, location string) (*MembershipList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListMembershipWithMaxResults(ctx, project, location, MembershipMaxPage)

}

func (c *Client) ListMembershipWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*MembershipList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &Membership{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listMembership(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &MembershipList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetMembership(ctx context.Context, r *Membership) (*Membership, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractMembershipFields(r)

	b, err := c.getMembershipRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalMembership(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeMembershipNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractMembershipFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteMembership(ctx context.Context, r *Membership) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("Membership resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting Membership...")
	deleteOp := deleteMembershipOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllMembership deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllMembership(ctx context.Context, project, location string, filter func(*Membership) bool) error {
	listObj, err := c.ListMembership(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllMembership(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllMembership(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyMembership(ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (*Membership, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *Membership
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyMembershipHelper(c, ctx, rawDesired, opts...)
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

func applyMembershipHelper(c *Client, ctx context.Context, rawDesired *Membership, opts ...dcl.ApplyOption) (*Membership, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyMembership...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractMembershipFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.membershipDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToMembershipDiffs(c.Config, fieldDiffs, opts)
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
	var ops []membershipApiOperation
	if create {
		ops = append(ops, &createMembershipOperation{})
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
	return applyMembershipDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyMembershipDiff(c *Client, ctx context.Context, desired *Membership, rawDesired *Membership, ops []membershipApiOperation, opts ...dcl.ApplyOption) (*Membership, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetMembership(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createMembershipOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapMembership(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeMembershipNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeMembershipNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeMembershipDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractMembershipFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractMembershipFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffMembership(c, newDesired, newState)
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

func (r *Membership) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"optionsRequestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "", body, nil
}
