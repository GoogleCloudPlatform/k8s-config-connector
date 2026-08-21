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

type ClientTlsPolicy struct {
	Name               *string                             `json:"name"`
	Description        *string                             `json:"description"`
	CreateTime         *string                             `json:"createTime"`
	UpdateTime         *string                             `json:"updateTime"`
	Labels             map[string]string                   `json:"labels"`
	Sni                *string                             `json:"sni"`
	ClientCertificate  *ClientTlsPolicyClientCertificate   `json:"clientCertificate"`
	ServerValidationCa []ClientTlsPolicyServerValidationCa `json:"serverValidationCa"`
	Project            *string                             `json:"project"`
	Location           *string                             `json:"location"`
}

func (r *ClientTlsPolicy) String() string {
	return dcl.SprintResource(r)
}

type ClientTlsPolicyClientCertificate struct {
	empty                       bool                                                         `json:"-"`
	LocalFilepath               *ClientTlsPolicyClientCertificateLocalFilepath               `json:"localFilepath"`
	GrpcEndpoint                *ClientTlsPolicyClientCertificateGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ClientTlsPolicyClientCertificateCertificateProviderInstance `json:"certificateProviderInstance"`
}

type jsonClientTlsPolicyClientCertificate ClientTlsPolicyClientCertificate

func (r *ClientTlsPolicyClientCertificate) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyClientCertificate
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyClientCertificate
	} else {

		r.LocalFilepath = res.LocalFilepath

		r.GrpcEndpoint = res.GrpcEndpoint

		r.CertificateProviderInstance = res.CertificateProviderInstance

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificate is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificate *ClientTlsPolicyClientCertificate = &ClientTlsPolicyClientCertificate{empty: true}

func (r *ClientTlsPolicyClientCertificate) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyClientCertificate) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificate) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyClientCertificateLocalFilepath struct {
	empty           bool    `json:"-"`
	CertificatePath *string `json:"certificatePath"`
	PrivateKeyPath  *string `json:"privateKeyPath"`
}

type jsonClientTlsPolicyClientCertificateLocalFilepath ClientTlsPolicyClientCertificateLocalFilepath

func (r *ClientTlsPolicyClientCertificateLocalFilepath) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyClientCertificateLocalFilepath
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyClientCertificateLocalFilepath
	} else {

		r.CertificatePath = res.CertificatePath

		r.PrivateKeyPath = res.PrivateKeyPath

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificateLocalFilepath is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificateLocalFilepath *ClientTlsPolicyClientCertificateLocalFilepath = &ClientTlsPolicyClientCertificateLocalFilepath{empty: true}

func (r *ClientTlsPolicyClientCertificateLocalFilepath) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyClientCertificateLocalFilepath) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificateLocalFilepath) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyClientCertificateGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

type jsonClientTlsPolicyClientCertificateGrpcEndpoint ClientTlsPolicyClientCertificateGrpcEndpoint

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyClientCertificateGrpcEndpoint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyClientCertificateGrpcEndpoint
	} else {

		r.TargetUri = res.TargetUri

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificateGrpcEndpoint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificateGrpcEndpoint *ClientTlsPolicyClientCertificateGrpcEndpoint = &ClientTlsPolicyClientCertificateGrpcEndpoint{empty: true}

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificateGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyClientCertificateCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

type jsonClientTlsPolicyClientCertificateCertificateProviderInstance ClientTlsPolicyClientCertificateCertificateProviderInstance

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyClientCertificateCertificateProviderInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyClientCertificateCertificateProviderInstance
	} else {

		r.PluginInstance = res.PluginInstance

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyClientCertificateCertificateProviderInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyClientCertificateCertificateProviderInstance *ClientTlsPolicyClientCertificateCertificateProviderInstance = &ClientTlsPolicyClientCertificateCertificateProviderInstance{empty: true}

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyClientCertificateCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCa struct {
	empty                       bool                                                          `json:"-"`
	CaCertPath                  *string                                                       `json:"caCertPath"`
	GrpcEndpoint                *ClientTlsPolicyServerValidationCaGrpcEndpoint                `json:"grpcEndpoint"`
	CertificateProviderInstance *ClientTlsPolicyServerValidationCaCertificateProviderInstance `json:"certificateProviderInstance"`
}

type jsonClientTlsPolicyServerValidationCa ClientTlsPolicyServerValidationCa

func (r *ClientTlsPolicyServerValidationCa) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyServerValidationCa
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyServerValidationCa
	} else {

		r.CaCertPath = res.CaCertPath

		r.GrpcEndpoint = res.GrpcEndpoint

		r.CertificateProviderInstance = res.CertificateProviderInstance

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCa is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCa *ClientTlsPolicyServerValidationCa = &ClientTlsPolicyServerValidationCa{empty: true}

func (r *ClientTlsPolicyServerValidationCa) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyServerValidationCa) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCa) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCaGrpcEndpoint struct {
	empty     bool    `json:"-"`
	TargetUri *string `json:"targetUri"`
}

type jsonClientTlsPolicyServerValidationCaGrpcEndpoint ClientTlsPolicyServerValidationCaGrpcEndpoint

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyServerValidationCaGrpcEndpoint
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyServerValidationCaGrpcEndpoint
	} else {

		r.TargetUri = res.TargetUri

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCaGrpcEndpoint is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCaGrpcEndpoint *ClientTlsPolicyServerValidationCaGrpcEndpoint = &ClientTlsPolicyServerValidationCaGrpcEndpoint{empty: true}

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCaGrpcEndpoint) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

type ClientTlsPolicyServerValidationCaCertificateProviderInstance struct {
	empty          bool    `json:"-"`
	PluginInstance *string `json:"pluginInstance"`
}

type jsonClientTlsPolicyServerValidationCaCertificateProviderInstance ClientTlsPolicyServerValidationCaCertificateProviderInstance

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) UnmarshalJSON(data []byte) error {
	var res jsonClientTlsPolicyServerValidationCaCertificateProviderInstance
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	var m map[string]interface{}
	json.Unmarshal(data, &m)

	if len(m) == 0 {
		*r = *EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance
	} else {

		r.PluginInstance = res.PluginInstance

	}
	return nil
}

// This object is used to assert a desired state where this ClientTlsPolicyServerValidationCaCertificateProviderInstance is
// empty. Go lacks global const objects, but this object should be treated
// as one. Modifying this object will have undesirable results.
var EmptyClientTlsPolicyServerValidationCaCertificateProviderInstance *ClientTlsPolicyServerValidationCaCertificateProviderInstance = &ClientTlsPolicyServerValidationCaCertificateProviderInstance{empty: true}

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) Empty() bool {
	return r.empty
}

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) String() string {
	return dcl.SprintResource(r)
}

func (r *ClientTlsPolicyServerValidationCaCertificateProviderInstance) HashCode() string {
	// Placeholder for a more complex hash method that handles ordering, etc
	// Hash resource body for easy comparison later
	hash := sha256.Sum256([]byte(r.String()))
	return fmt.Sprintf("%x", hash)
}

// Describe returns a simple description of this resource to ensure that automated tools
// can identify it.
func (r *ClientTlsPolicy) Describe() dcl.ServiceTypeVersion {
	return dcl.ServiceTypeVersion{
		Service: "network_security",
		Type:    "ClientTlsPolicy",
		Version: "alpha",
	}
}

func (r *ClientTlsPolicy) ID() (string, error) {
	if err := extractClientTlsPolicyFields(r); err != nil {
		return "", err
	}
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name":                 dcl.ValueOrEmptyString(nr.Name),
		"description":          dcl.ValueOrEmptyString(nr.Description),
		"create_time":          dcl.ValueOrEmptyString(nr.CreateTime),
		"update_time":          dcl.ValueOrEmptyString(nr.UpdateTime),
		"labels":               dcl.ValueOrEmptyString(nr.Labels),
		"sni":                  dcl.ValueOrEmptyString(nr.Sni),
		"client_certificate":   dcl.ValueOrEmptyString(nr.ClientCertificate),
		"server_validation_ca": dcl.ValueOrEmptyString(nr.ServerValidationCa),
		"project":              dcl.ValueOrEmptyString(nr.Project),
		"location":             dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.Nprintf("projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}", params), nil
}

const ClientTlsPolicyMaxPage = -1

type ClientTlsPolicyList struct {
	Items []*ClientTlsPolicy

	nextToken string

	pageSize int32

	resource *ClientTlsPolicy
}

func (l *ClientTlsPolicyList) HasNext() bool {
	return l.nextToken != ""
}

func (l *ClientTlsPolicyList) Next(ctx context.Context, c *Client) error {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if !l.HasNext() {
		return fmt.Errorf("no next page")
	}
	items, token, err := c.listClientTlsPolicy(ctx, l.resource, l.nextToken, l.pageSize)
	if err != nil {
		return err
	}
	l.Items = items
	l.nextToken = token
	return err
}

func (c *Client) ListClientTlsPolicy(ctx context.Context, project, location string) (*ClientTlsPolicyList, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	return c.ListClientTlsPolicyWithMaxResults(ctx, project, location, ClientTlsPolicyMaxPage)

}

func (c *Client) ListClientTlsPolicyWithMaxResults(ctx context.Context, project, location string, pageSize int32) (*ClientTlsPolicyList, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// Create a resource object so that we can use proper url normalization methods.
	r := &ClientTlsPolicy{
		Project:  &project,
		Location: &location,
	}
	items, token, err := c.listClientTlsPolicy(ctx, r, "", pageSize)
	if err != nil {
		return nil, err
	}
	return &ClientTlsPolicyList{
		Items:     items,
		nextToken: token,
		pageSize:  pageSize,
		resource:  r,
	}, nil
}

func (c *Client) GetClientTlsPolicy(ctx context.Context, r *ClientTlsPolicy) (*ClientTlsPolicy, error) {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	// This is *purposefully* supressing errors.
	// This function is used with url-normalized values + not URL normalized values.
	// URL Normalized values will throw unintentional errors, since those values are not of the proper parent form.
	extractClientTlsPolicyFields(r)

	b, err := c.getClientTlsPolicyRaw(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil, &googleapi.Error{
				Code:    404,
				Message: err.Error(),
			}
		}
		return nil, err
	}
	result, err := unmarshalClientTlsPolicy(b, c, r)
	if err != nil {
		return nil, err
	}
	result.Project = r.Project
	result.Location = r.Location
	result.Name = r.Name

	c.Config.Logger.InfoWithContextf(ctx, "Retrieved raw result state: %v", result)
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with specified state: %v", r)
	result, err = canonicalizeClientTlsPolicyNewState(c, result, r)
	if err != nil {
		return nil, err
	}
	if err := postReadExtractClientTlsPolicyFields(result); err != nil {
		return result, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Created result state: %v", result)

	return result, nil
}

func (c *Client) DeleteClientTlsPolicy(ctx context.Context, r *ClientTlsPolicy) error {
	ctx = dcl.ContextWithRequestID(ctx)
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	if r == nil {
		return fmt.Errorf("ClientTlsPolicy resource is nil")
	}
	c.Config.Logger.InfoWithContext(ctx, "Deleting ClientTlsPolicy...")
	deleteOp := deleteClientTlsPolicyOperation{}
	return deleteOp.do(ctx, r, c)
}

// DeleteAllClientTlsPolicy deletes all resources that the filter functions returns true on.
func (c *Client) DeleteAllClientTlsPolicy(ctx context.Context, project, location string, filter func(*ClientTlsPolicy) bool) error {
	listObj, err := c.ListClientTlsPolicy(ctx, project, location)
	if err != nil {
		return err
	}

	err = c.deleteAllClientTlsPolicy(ctx, filter, listObj.Items)
	if err != nil {
		return err
	}
	for listObj.HasNext() {
		err = listObj.Next(ctx, c)
		if err != nil {
			return nil
		}
		err = c.deleteAllClientTlsPolicy(ctx, filter, listObj.Items)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) ApplyClientTlsPolicy(ctx context.Context, rawDesired *ClientTlsPolicy, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Config.TimeoutOr(0*time.Second))
	defer cancel()

	ctx = dcl.ContextWithRequestID(ctx)
	var resultNewState *ClientTlsPolicy
	err := dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		newState, err := applyClientTlsPolicyHelper(c, ctx, rawDesired, opts...)
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

func applyClientTlsPolicyHelper(c *Client, ctx context.Context, rawDesired *ClientTlsPolicy, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {
	c.Config.Logger.InfoWithContext(ctx, "Beginning ApplyClientTlsPolicy...")
	c.Config.Logger.InfoWithContextf(ctx, "User specified desired state: %v", rawDesired)

	// 1.1: Validation of user-specified fields in desired state.
	if err := rawDesired.validate(); err != nil {
		return nil, err
	}

	if err := extractClientTlsPolicyFields(rawDesired); err != nil {
		return nil, err
	}

	initial, desired, fieldDiffs, err := c.clientTlsPolicyDiffsForRawDesired(ctx, rawDesired, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a diff: %w", err)
	}

	diffs, err := convertFieldDiffsToClientTlsPolicyDiffs(c.Config, fieldDiffs, opts)
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
	var ops []clientTlsPolicyApiOperation
	if create {
		ops = append(ops, &createClientTlsPolicyOperation{})
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
	return applyClientTlsPolicyDiff(c, ctx, desired, rawDesired, ops, opts...)
}

func applyClientTlsPolicyDiff(c *Client, ctx context.Context, desired *ClientTlsPolicy, rawDesired *ClientTlsPolicy, ops []clientTlsPolicyApiOperation, opts ...dcl.ApplyOption) (*ClientTlsPolicy, error) {
	// 3.1, 3.2a Retrieval of raw new state & canonicalization with desired state
	c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state...")
	rawNew, err := c.GetClientTlsPolicy(ctx, desired)
	if err != nil {
		return nil, err
	}
	// Get additional values from the first response.
	// These values should be merged into the newState above.
	if len(ops) > 0 {
		lastOp := ops[len(ops)-1]
		if o, ok := lastOp.(*createClientTlsPolicyOperation); ok {
			if r, hasR := o.FirstResponse(); hasR {

				c.Config.Logger.InfoWithContext(ctx, "Retrieving raw new state from operation...")

				fullResp, err := unmarshalMapClientTlsPolicy(r, c, rawDesired)
				if err != nil {
					return nil, err
				}

				rawNew, err = canonicalizeClientTlsPolicyNewState(c, rawNew, fullResp)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	c.Config.Logger.InfoWithContextf(ctx, "Canonicalizing with raw desired state: %v", rawDesired)
	// 3.2b Canonicalization of raw new state using raw desired state
	newState, err := canonicalizeClientTlsPolicyNewState(c, rawNew, rawDesired)
	if err != nil {
		return rawNew, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created canonical new state: %v", newState)
	// 3.3 Comparison of the new state and raw desired state.
	// TODO(magic-modules-eng): EVENTUALLY_CONSISTENT_UPDATE
	newDesired, err := canonicalizeClientTlsPolicyDesiredState(rawDesired, newState)
	if err != nil {
		return newState, err
	}

	if err := postReadExtractClientTlsPolicyFields(newState); err != nil {
		return newState, err
	}

	// Need to ensure any transformations made here match acceptably in differ.
	if err := postReadExtractClientTlsPolicyFields(newDesired); err != nil {
		return newState, err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Diffing using canonicalized desired state: %v", newDesired)
	newDiffs, err := diffClientTlsPolicy(c, newDesired, newState)
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

func (r *ClientTlsPolicy) GetPolicy(basePath string) (string, string, *bytes.Buffer, error) {
	u := r.getPolicyURL(basePath)
	body := &bytes.Buffer{}
	u, err := dcl.AddQueryParams(u, map[string]string{"options.requestedPolicyVersion": fmt.Sprintf("%d", r.IAMPolicyVersion())})
	if err != nil {
		return "", "", nil, err
	}
	return u, "GET", body, nil
}
