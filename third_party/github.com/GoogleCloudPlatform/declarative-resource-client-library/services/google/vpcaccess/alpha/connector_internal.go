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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Connector) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Subnet) {
		if err := r.Subnet.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *ConnectorSubnet) validate() error {
	return nil
}
func (r *Connector) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://vpcaccess.googleapis.com/v1/", params)
}

func (r *Connector) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Connector) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors", nr.basePath(), userBasePath, params), nil

}

func (r *Connector) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors?connectorId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Connector) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/connectors/{{name}}", nr.basePath(), userBasePath, params), nil
}

// connectorApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type connectorApiOperation interface {
	do(context.Context, *Connector, *Client) error
}

func (c *Client) listConnectorRaw(ctx context.Context, r *Connector, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != ConnectorMaxPage {
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

type listConnectorOperation struct {
	Connectors []map[string]interface{} `json:"connectors"`
	Token      string                   `json:"nextPageToken"`
}

func (c *Client) listConnector(ctx context.Context, r *Connector, pageToken string, pageSize int32) ([]*Connector, string, error) {
	b, err := c.listConnectorRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listConnectorOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Connector
	for _, v := range m.Connectors {
		res, err := unmarshalMapConnector(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllConnector(ctx context.Context, f func(*Connector) bool, resources []*Connector) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteConnector(ctx, res)
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

type deleteConnectorOperation struct{}

func (op *deleteConnectorOperation) do(ctx context.Context, r *Connector, c *Client) error {
	r, err := c.GetConnector(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Connector not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetConnector checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// We saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// This is the reason we are adding retry to handle that case.
	retriesRemaining := 10
	dcl.Do(ctx, func(ctx context.Context) (*dcl.RetryDetails, error) {
		_, err := c.GetConnector(ctx, r)
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
type createConnectorOperation struct {
	response map[string]interface{}
}

func (op *createConnectorOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createConnectorOperation) do(ctx context.Context, r *Connector, c *Client) error {
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
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	if _, err := c.GetConnector(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getConnectorRaw(ctx context.Context, r *Connector) ([]byte, error) {

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

func (c *Client) connectorDiffsForRawDesired(ctx context.Context, rawDesired *Connector, opts ...dcl.ApplyOption) (initial, desired *Connector, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Connector
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Connector); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Connector, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetConnector(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Connector resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Connector resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Connector resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeConnectorDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Connector: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Connector: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractConnectorFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeConnectorInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Connector: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeConnectorDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Connector: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffConnector(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeConnectorInitialState(rawInitial, rawDesired *Connector) (*Connector, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeConnectorDesiredState(rawDesired, rawInitial *Connector, opts ...dcl.ApplyOption) (*Connector, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Subnet = canonicalizeConnectorSubnet(rawDesired.Subnet, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Connector{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.IsZeroValue(rawDesired.Network) || (dcl.IsEmptyValueIndirect(rawDesired.Network) && dcl.IsEmptyValueIndirect(rawInitial.Network)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Network = rawInitial.Network
	} else {
		canonicalDesired.Network = rawDesired.Network
	}
	if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawInitial.IPCidrRange) {
		canonicalDesired.IPCidrRange = rawInitial.IPCidrRange
	} else {
		canonicalDesired.IPCidrRange = rawDesired.IPCidrRange
	}
	if dcl.IsZeroValue(rawDesired.MinThroughput) || (dcl.IsEmptyValueIndirect(rawDesired.MinThroughput) && dcl.IsEmptyValueIndirect(rawInitial.MinThroughput)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.MinThroughput = rawInitial.MinThroughput
	} else {
		canonicalDesired.MinThroughput = rawDesired.MinThroughput
	}
	if dcl.IsZeroValue(rawDesired.MaxThroughput) || (dcl.IsEmptyValueIndirect(rawDesired.MaxThroughput) && dcl.IsEmptyValueIndirect(rawInitial.MaxThroughput)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.MaxThroughput = rawInitial.MaxThroughput
	} else {
		canonicalDesired.MaxThroughput = rawDesired.MaxThroughput
	}
	canonicalDesired.Subnet = canonicalizeConnectorSubnet(rawDesired.Subnet, rawInitial.Subnet, opts...)
	if dcl.StringCanonicalize(rawDesired.MachineType, rawInitial.MachineType) {
		canonicalDesired.MachineType = rawInitial.MachineType
	} else {
		canonicalDesired.MachineType = rawDesired.MachineType
	}
	if dcl.IsZeroValue(rawDesired.MinInstances) || (dcl.IsEmptyValueIndirect(rawDesired.MinInstances) && dcl.IsEmptyValueIndirect(rawInitial.MinInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.MinInstances = rawInitial.MinInstances
	} else {
		canonicalDesired.MinInstances = rawDesired.MinInstances
	}
	if dcl.IsZeroValue(rawDesired.MaxInstances) || (dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) && dcl.IsEmptyValueIndirect(rawInitial.MaxInstances)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.MaxInstances = rawInitial.MaxInstances
	} else {
		canonicalDesired.MaxInstances = rawDesired.MaxInstances
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}
	if dcl.NameToSelfLink(rawDesired.Location, rawInitial.Location) {
		canonicalDesired.Location = rawInitial.Location
	} else {
		canonicalDesired.Location = rawDesired.Location
	}
	return canonicalDesired, nil
}

func canonicalizeConnectorNewState(c *Client, rawNew, rawDesired *Connector) (*Connector, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Network) && dcl.IsEmptyValueIndirect(rawDesired.Network) {
		rawNew.Network = rawDesired.Network
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.IPCidrRange) && dcl.IsEmptyValueIndirect(rawDesired.IPCidrRange) {
		rawNew.IPCidrRange = rawDesired.IPCidrRange
	} else {
		if dcl.StringCanonicalize(rawDesired.IPCidrRange, rawNew.IPCidrRange) {
			rawNew.IPCidrRange = rawDesired.IPCidrRange
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinThroughput) && dcl.IsEmptyValueIndirect(rawDesired.MinThroughput) {
		rawNew.MinThroughput = rawDesired.MinThroughput
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxThroughput) && dcl.IsEmptyValueIndirect(rawDesired.MaxThroughput) {
		rawNew.MaxThroughput = rawDesired.MaxThroughput
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.ConnectedProjects) && dcl.IsEmptyValueIndirect(rawDesired.ConnectedProjects) {
		rawNew.ConnectedProjects = rawDesired.ConnectedProjects
	} else {
		if dcl.StringArrayCanonicalize(rawDesired.ConnectedProjects, rawNew.ConnectedProjects) {
			rawNew.ConnectedProjects = rawDesired.ConnectedProjects
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Subnet) && dcl.IsEmptyValueIndirect(rawDesired.Subnet) {
		rawNew.Subnet = rawDesired.Subnet
	} else {
		rawNew.Subnet = canonicalizeNewConnectorSubnet(c, rawDesired.Subnet, rawNew.Subnet)
	}

	if dcl.IsEmptyValueIndirect(rawNew.MachineType) && dcl.IsEmptyValueIndirect(rawDesired.MachineType) {
		rawNew.MachineType = rawDesired.MachineType
	} else {
		if dcl.StringCanonicalize(rawDesired.MachineType, rawNew.MachineType) {
			rawNew.MachineType = rawDesired.MachineType
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.MinInstances) && dcl.IsEmptyValueIndirect(rawDesired.MinInstances) {
		rawNew.MinInstances = rawDesired.MinInstances
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.MaxInstances) && dcl.IsEmptyValueIndirect(rawDesired.MaxInstances) {
		rawNew.MaxInstances = rawDesired.MaxInstances
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

func canonicalizeConnectorSubnet(des, initial *ConnectorSubnet, opts ...dcl.ApplyOption) *ConnectorSubnet {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &ConnectorSubnet{}

	if dcl.IsZeroValue(des.Name) || (dcl.IsEmptyValueIndirect(des.Name) && dcl.IsEmptyValueIndirect(initial.Name)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.IsZeroValue(des.ProjectId) || (dcl.IsEmptyValueIndirect(des.ProjectId) && dcl.IsEmptyValueIndirect(initial.ProjectId)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ProjectId = initial.ProjectId
	} else {
		cDes.ProjectId = des.ProjectId
	}

	return cDes
}

func canonicalizeConnectorSubnetSlice(des, initial []ConnectorSubnet, opts ...dcl.ApplyOption) []ConnectorSubnet {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]ConnectorSubnet, 0, len(des))
		for _, d := range des {
			cd := canonicalizeConnectorSubnet(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]ConnectorSubnet, 0, len(des))
	for i, d := range des {
		cd := canonicalizeConnectorSubnet(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewConnectorSubnet(c *Client, des, nw *ConnectorSubnet) *ConnectorSubnet {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for ConnectorSubnet while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewConnectorSubnetSet(c *Client, des, nw []ConnectorSubnet) []ConnectorSubnet {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []ConnectorSubnet
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareConnectorSubnetNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewConnectorSubnet(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewConnectorSubnetSlice(c *Client, des, nw []ConnectorSubnet) []ConnectorSubnet {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []ConnectorSubnet
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewConnectorSubnet(c, &d, &n))
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
func diffConnector(c *Client, desired, actual *Connector, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Network, actual.Network, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Network")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IPCidrRange, actual.IPCidrRange, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("IpCidrRange")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.State, actual.State, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("State")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinThroughput, actual.MinThroughput, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinThroughput")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxThroughput, actual.MaxThroughput, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxThroughput")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConnectedProjects, actual.ConnectedProjects, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConnectedProjects")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Subnet, actual.Subnet, dcl.DiffInfo{ObjectFunction: compareConnectorSubnetNewStyle, EmptyObject: EmptyConnectorSubnet, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Subnet")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MachineType, actual.MachineType, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MachineType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MinInstances, actual.MinInstances, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MinInstances")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MaxInstances, actual.MaxInstances, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MaxInstances")); len(ds) != 0 || err != nil {
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
func compareConnectorSubnetNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*ConnectorSubnet)
	if !ok {
		desiredNotPointer, ok := d.(ConnectorSubnet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectorSubnet or *ConnectorSubnet", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*ConnectorSubnet)
	if !ok {
		actualNotPointer, ok := a.(ConnectorSubnet)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a ConnectorSubnet", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ProjectId, actual.ProjectId, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ProjectId")); len(ds) != 0 || err != nil {
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
func (r *Connector) urlNormalized() *Connector {
	normalized := dcl.Copy(*r).(Connector)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Network = dcl.SelfLinkToName(r.Network)
	normalized.IPCidrRange = dcl.SelfLinkToName(r.IPCidrRange)
	normalized.MachineType = dcl.SelfLinkToName(r.MachineType)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Connector) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Connector resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Connector) marshal(c *Client) ([]byte, error) {
	m, err := expandConnector(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Connector: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalConnector decodes JSON responses into the Connector resource schema.
func unmarshalConnector(b []byte, c *Client, res *Connector) (*Connector, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapConnector(m, c, res)
}

func unmarshalMapConnector(m map[string]interface{}, c *Client, res *Connector) (*Connector, error) {

	flattened := flattenConnector(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandConnector expands Connector into a JSON request object.
func expandConnector(c *Client, f *Connector) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.Network); err != nil {
		return nil, fmt.Errorf("error expanding Network into network: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["network"] = v
	}
	if v := f.IPCidrRange; dcl.ValueShouldBeSent(v) {
		m["ipCidrRange"] = v
	}
	if v := f.MinThroughput; dcl.ValueShouldBeSent(v) {
		m["minThroughput"] = v
	}
	if v := f.MaxThroughput; dcl.ValueShouldBeSent(v) {
		m["maxThroughput"] = v
	}
	if v, err := expandConnectorSubnet(c, f.Subnet, res); err != nil {
		return nil, fmt.Errorf("error expanding Subnet into subnet: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["subnet"] = v
	}
	if v := f.MachineType; dcl.ValueShouldBeSent(v) {
		m["machineType"] = v
	}
	if v := f.MinInstances; dcl.ValueShouldBeSent(v) {
		m["minInstances"] = v
	}
	if v := f.MaxInstances; dcl.ValueShouldBeSent(v) {
		m["maxInstances"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["project"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Location into location: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["location"] = v
	}

	return m, nil
}

// flattenConnector flattens Connector from a JSON request object into the
// Connector type.
func flattenConnector(c *Client, i interface{}, res *Connector) *Connector {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Connector{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Network = dcl.FlattenString(m["network"])
	resultRes.IPCidrRange = dcl.FlattenString(m["ipCidrRange"])
	resultRes.State = flattenConnectorStateEnum(m["state"])
	resultRes.MinThroughput = dcl.FlattenInteger(m["minThroughput"])
	resultRes.MaxThroughput = dcl.FlattenInteger(m["maxThroughput"])
	resultRes.ConnectedProjects = dcl.FlattenStringSlice(m["connectedProjects"])
	resultRes.Subnet = flattenConnectorSubnet(c, m["subnet"], res)
	resultRes.MachineType = dcl.FlattenString(m["machineType"])
	resultRes.MinInstances = dcl.FlattenInteger(m["minInstances"])
	resultRes.MaxInstances = dcl.FlattenInteger(m["maxInstances"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// expandConnectorSubnetMap expands the contents of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnetMap(c *Client, f map[string]ConnectorSubnet, res *Connector) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandConnectorSubnet(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandConnectorSubnetSlice expands the contents of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnetSlice(c *Client, f []ConnectorSubnet, res *Connector) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandConnectorSubnet(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenConnectorSubnetMap flattens the contents of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnetMap(c *Client, i interface{}, res *Connector) map[string]ConnectorSubnet {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectorSubnet{}
	}

	if len(a) == 0 {
		return map[string]ConnectorSubnet{}
	}

	items := make(map[string]ConnectorSubnet)
	for k, item := range a {
		items[k] = *flattenConnectorSubnet(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenConnectorSubnetSlice flattens the contents of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnetSlice(c *Client, i interface{}, res *Connector) []ConnectorSubnet {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectorSubnet{}
	}

	if len(a) == 0 {
		return []ConnectorSubnet{}
	}

	items := make([]ConnectorSubnet, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectorSubnet(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandConnectorSubnet expands an instance of ConnectorSubnet into a JSON
// request object.
func expandConnectorSubnet(c *Client, f *ConnectorSubnet, res *Connector) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := dcl.SelfLinkToNameExpander(f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v, err := dcl.SelfLinkToNameExpander(f.ProjectId); err != nil {
		return nil, fmt.Errorf("error expanding ProjectId into projectId: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["projectId"] = v
	}

	return m, nil
}

// flattenConnectorSubnet flattens an instance of ConnectorSubnet from a JSON
// response object.
func flattenConnectorSubnet(c *Client, i interface{}, res *Connector) *ConnectorSubnet {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &ConnectorSubnet{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyConnectorSubnet
	}
	r.Name = dcl.FlattenString(m["name"])
	r.ProjectId = dcl.FlattenString(m["projectId"])

	return r
}

// flattenConnectorStateEnumMap flattens the contents of ConnectorStateEnum from a JSON
// response object.
func flattenConnectorStateEnumMap(c *Client, i interface{}, res *Connector) map[string]ConnectorStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]ConnectorStateEnum{}
	}

	if len(a) == 0 {
		return map[string]ConnectorStateEnum{}
	}

	items := make(map[string]ConnectorStateEnum)
	for k, item := range a {
		items[k] = *flattenConnectorStateEnum(item.(interface{}))
	}

	return items
}

// flattenConnectorStateEnumSlice flattens the contents of ConnectorStateEnum from a JSON
// response object.
func flattenConnectorStateEnumSlice(c *Client, i interface{}, res *Connector) []ConnectorStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []ConnectorStateEnum{}
	}

	if len(a) == 0 {
		return []ConnectorStateEnum{}
	}

	items := make([]ConnectorStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenConnectorStateEnum(item.(interface{})))
	}

	return items
}

// flattenConnectorStateEnum asserts that an interface is a string, and returns a
// pointer to a *ConnectorStateEnum with the same value as that string.
func flattenConnectorStateEnum(i interface{}) *ConnectorStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return ConnectorStateEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Connector) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalConnector(b, c, r)
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
		if nr.Location == nil && ncr.Location == nil {
			c.Config.Logger.Info("Both Location fields null - considering equal.")
		} else if nr.Location == nil || ncr.Location == nil {
			c.Config.Logger.Info("Only one Location field is null - considering unequal.")
			return false
		} else if *nr.Location != *ncr.Location {
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

type connectorDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         connectorApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToConnectorDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]connectorDiff, error) {
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
	var diffs []connectorDiff
	// For each operation name, create a connectorDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := connectorDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToConnectorApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToConnectorApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (connectorApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractConnectorFields(r *Connector) error {
	vSubnet := r.Subnet
	if vSubnet == nil {
		// note: explicitly not the empty object.
		vSubnet = &ConnectorSubnet{}
	}
	if err := extractConnectorSubnetFields(r, vSubnet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSubnet) {
		r.Subnet = vSubnet
	}
	return nil
}
func extractConnectorSubnetFields(r *Connector, o *ConnectorSubnet) error {
	return nil
}

func postReadExtractConnectorFields(r *Connector) error {
	vSubnet := r.Subnet
	if vSubnet == nil {
		// note: explicitly not the empty object.
		vSubnet = &ConnectorSubnet{}
	}
	if err := postReadExtractConnectorSubnetFields(r, vSubnet); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vSubnet) {
		r.Subnet = vSubnet
	}
	return nil
}
func postReadExtractConnectorSubnetFields(r *Connector, o *ConnectorSubnet) error {
	return nil
}
