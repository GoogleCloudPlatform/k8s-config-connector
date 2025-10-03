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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *FeatureMembership) validate() error {

	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Feature, "Feature"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Membership, "Membership"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.Mesh) {
		if err := r.Mesh.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Configmanagement) {
		if err := r.Configmanagement.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Policycontroller) {
		if err := r.Policycontroller.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipMesh) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagement) validate() error {
	if !dcl.IsEmptyValueIndirect(r.ConfigSync) {
		if err := r.ConfigSync.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PolicyController) {
		if err := r.PolicyController.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Binauthz) {
		if err := r.Binauthz.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.HierarchyController) {
		if err := r.HierarchyController.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagementConfigSync) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Git) {
		if err := r.Git.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.Oci) {
		if err := r.Oci.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagementConfigSyncGit) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementConfigSyncOci) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementPolicyController) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Monitoring) {
		if err := r.Monitoring.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipConfigmanagementPolicyControllerMonitoring) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementBinauthz) validate() error {
	return nil
}
func (r *FeatureMembershipConfigmanagementHierarchyController) validate() error {
	return nil
}
func (r *FeatureMembershipPolicycontroller) validate() error {
	if err := dcl.Required(r, "policyControllerHubConfig"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.PolicyControllerHubConfig) {
		if err := r.PolicyControllerHubConfig.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Monitoring) {
		if err := r.Monitoring.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.PolicyContent) {
		if err := r.PolicyContent.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) validate() error {
	return nil
}
func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) validate() error {
	if !dcl.IsEmptyValueIndirect(r.TemplateLibrary) {
		if err := r.TemplateLibrary.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) validate() error {
	return nil
}
func (r *FeatureMembership) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://gkehub.googleapis.com/v1beta1/", params)
}

func (r *FeatureMembership) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"feature":  dcl.ValueOrEmptyString(nr.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", nr.basePath(), userBasePath, params), nil
}

func (r *FeatureMembership) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"feature":  dcl.ValueOrEmptyString(nr.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", nr.basePath(), userBasePath, params), nil

}

func (r *FeatureMembership) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"feature":  dcl.ValueOrEmptyString(nr.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", nr.basePath(), userBasePath, params), nil

}

func (r *FeatureMembership) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"feature":  dcl.ValueOrEmptyString(nr.Feature),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", nr.basePath(), userBasePath, params), nil
}

// featureMembershipApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type featureMembershipApiOperation interface {
	do(context.Context, *FeatureMembership, *Client) error
}

// newUpdateFeatureMembershipUpdateFeatureMembershipRequest creates a request for an
// FeatureMembership resource's UpdateFeatureMembership update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateFeatureMembershipUpdateFeatureMembershipRequest(ctx context.Context, f *FeatureMembership, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v, err := expandFeatureMembershipMesh(c, f.Mesh, res); err != nil {
		return nil, fmt.Errorf("error expanding Mesh into mesh: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["mesh"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagement(c, f.Configmanagement, res); err != nil {
		return nil, fmt.Errorf("error expanding Configmanagement into configmanagement: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["configmanagement"] = v
	}
	if v, err := expandFeatureMembershipPolicycontroller(c, f.Policycontroller, res); err != nil {
		return nil, fmt.Errorf("error expanding Policycontroller into policycontroller: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["policycontroller"] = v
	}
	return req, nil
}

// marshalUpdateFeatureMembershipUpdateFeatureMembershipRequest converts the update into
// the final JSON request body.
func marshalUpdateFeatureMembershipUpdateFeatureMembershipRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateFeatureMembershipUpdateFeatureMembershipOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (c *Client) deleteAllFeatureMembership(ctx context.Context, f func(*FeatureMembership) bool, resources []*FeatureMembership) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteFeatureMembership(ctx, res)
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

type deleteFeatureMembershipOperation struct{}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createFeatureMembershipOperation struct {
	response map[string]interface{}
}

func (op *createFeatureMembershipOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) featureMembershipDiffsForRawDesired(ctx context.Context, rawDesired *FeatureMembership, opts ...dcl.ApplyOption) (initial, desired *FeatureMembership, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *FeatureMembership
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*FeatureMembership); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected FeatureMembership, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetFeatureMembership(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a FeatureMembership resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve FeatureMembership resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that FeatureMembership resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for FeatureMembership: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for FeatureMembership: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractFeatureMembershipFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeFeatureMembershipInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for FeatureMembership: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for FeatureMembership: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffFeatureMembership(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeFeatureMembershipInitialState(rawInitial, rawDesired *FeatureMembership) (*FeatureMembership, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeFeatureMembershipDesiredState(rawDesired, rawInitial *FeatureMembership, opts ...dcl.ApplyOption) (*FeatureMembership, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.Mesh = canonicalizeFeatureMembershipMesh(rawDesired.Mesh, nil, opts...)
		rawDesired.Configmanagement = canonicalizeFeatureMembershipConfigmanagement(rawDesired.Configmanagement, nil, opts...)
		rawDesired.Policycontroller = canonicalizeFeatureMembershipPolicycontroller(rawDesired.Policycontroller, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &FeatureMembership{}
	canonicalDesired.Mesh = canonicalizeFeatureMembershipMesh(rawDesired.Mesh, rawInitial.Mesh, opts...)
	canonicalDesired.Configmanagement = canonicalizeFeatureMembershipConfigmanagement(rawDesired.Configmanagement, rawInitial.Configmanagement, opts...)
	canonicalDesired.Policycontroller = canonicalizeFeatureMembershipPolicycontroller(rawDesired.Policycontroller, rawInitial.Policycontroller, opts...)
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
	if dcl.NameToSelfLink(rawDesired.Feature, rawInitial.Feature) {
		canonicalDesired.Feature = rawInitial.Feature
	} else {
		canonicalDesired.Feature = rawDesired.Feature
	}
	if dcl.NameToSelfLink(rawDesired.Membership, rawInitial.Membership) {
		canonicalDesired.Membership = rawInitial.Membership
	} else {
		canonicalDesired.Membership = rawDesired.Membership
	}
	if dcl.NameToSelfLink(rawDesired.MembershipLocation, rawInitial.MembershipLocation) {
		canonicalDesired.MembershipLocation = rawInitial.MembershipLocation
	} else {
		canonicalDesired.MembershipLocation = rawDesired.MembershipLocation
	}
	return canonicalDesired, nil
}

func canonicalizeFeatureMembershipNewState(c *Client, rawNew, rawDesired *FeatureMembership) (*FeatureMembership, error) {

	if dcl.IsEmptyValueIndirect(rawNew.Mesh) && dcl.IsEmptyValueIndirect(rawDesired.Mesh) {
		rawNew.Mesh = rawDesired.Mesh
	} else {
		rawNew.Mesh = canonicalizeNewFeatureMembershipMesh(c, rawDesired.Mesh, rawNew.Mesh)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Configmanagement) && dcl.IsEmptyValueIndirect(rawDesired.Configmanagement) {
		rawNew.Configmanagement = rawDesired.Configmanagement
	} else {
		rawNew.Configmanagement = canonicalizeNewFeatureMembershipConfigmanagement(c, rawDesired.Configmanagement, rawNew.Configmanagement)
	}

	if dcl.IsEmptyValueIndirect(rawNew.Policycontroller) && dcl.IsEmptyValueIndirect(rawDesired.Policycontroller) {
		rawNew.Policycontroller = rawDesired.Policycontroller
	} else {
		rawNew.Policycontroller = canonicalizeNewFeatureMembershipPolicycontroller(c, rawDesired.Policycontroller, rawNew.Policycontroller)
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	rawNew.Feature = rawDesired.Feature

	rawNew.Membership = rawDesired.Membership

	rawNew.MembershipLocation = rawDesired.MembershipLocation

	return rawNew, nil
}

func canonicalizeFeatureMembershipMesh(des, initial *FeatureMembershipMesh, opts ...dcl.ApplyOption) *FeatureMembershipMesh {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipMesh{}

	if dcl.IsZeroValue(des.Management) || (dcl.IsEmptyValueIndirect(des.Management) && dcl.IsEmptyValueIndirect(initial.Management)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Management = initial.Management
	} else {
		cDes.Management = des.Management
	}
	if dcl.IsZeroValue(des.ControlPlane) || (dcl.IsEmptyValueIndirect(des.ControlPlane) && dcl.IsEmptyValueIndirect(initial.ControlPlane)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ControlPlane = initial.ControlPlane
	} else {
		cDes.ControlPlane = des.ControlPlane
	}

	return cDes
}

func canonicalizeFeatureMembershipMeshSlice(des, initial []FeatureMembershipMesh, opts ...dcl.ApplyOption) []FeatureMembershipMesh {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipMesh, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipMesh(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipMesh, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipMesh(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipMesh(c *Client, des, nw *FeatureMembershipMesh) *FeatureMembershipMesh {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipMesh while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureMembershipMeshSet(c *Client, des, nw []FeatureMembershipMesh) []FeatureMembershipMesh {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipMesh
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipMeshNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipMesh(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipMeshSlice(c *Client, des, nw []FeatureMembershipMesh) []FeatureMembershipMesh {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipMesh
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipMesh(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagement(des, initial *FeatureMembershipConfigmanagement, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagement {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagement{}

	cDes.ConfigSync = canonicalizeFeatureMembershipConfigmanagementConfigSync(des.ConfigSync, initial.ConfigSync, opts...)
	cDes.PolicyController = canonicalizeFeatureMembershipConfigmanagementPolicyController(des.PolicyController, initial.PolicyController, opts...)
	cDes.Binauthz = canonicalizeFeatureMembershipConfigmanagementBinauthz(des.Binauthz, initial.Binauthz, opts...)
	cDes.HierarchyController = canonicalizeFeatureMembershipConfigmanagementHierarchyController(des.HierarchyController, initial.HierarchyController, opts...)
	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementSlice(des, initial []FeatureMembershipConfigmanagement, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagement {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagement, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagement(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagement, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagement(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagement(c *Client, des, nw *FeatureMembershipConfigmanagement) *FeatureMembershipConfigmanagement {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagement while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.ConfigSync = canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c, des.ConfigSync, nw.ConfigSync)
	nw.PolicyController = canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c, des.PolicyController, nw.PolicyController)
	nw.Binauthz = canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c, des.Binauthz, nw.Binauthz)
	nw.HierarchyController = canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c, des.HierarchyController, nw.HierarchyController)
	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementSet(c *Client, des, nw []FeatureMembershipConfigmanagement) []FeatureMembershipConfigmanagement {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagement
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagement(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementSlice(c *Client, des, nw []FeatureMembershipConfigmanagement) []FeatureMembershipConfigmanagement {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagement
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagement(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementConfigSync(des, initial *FeatureMembershipConfigmanagementConfigSync, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementConfigSync{}

	cDes.Git = canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(des.Git, initial.Git, opts...)
	if dcl.StringCanonicalize(des.SourceFormat, initial.SourceFormat) || dcl.IsZeroValue(des.SourceFormat) {
		cDes.SourceFormat = initial.SourceFormat
	} else {
		cDes.SourceFormat = des.SourceFormat
	}
	if dcl.BoolCanonicalize(des.PreventDrift, initial.PreventDrift) || dcl.IsZeroValue(des.PreventDrift) {
		cDes.PreventDrift = initial.PreventDrift
	} else {
		cDes.PreventDrift = des.PreventDrift
	}
	if dcl.IsZeroValue(des.MetricsGcpServiceAccountEmail) || (dcl.IsEmptyValueIndirect(des.MetricsGcpServiceAccountEmail) && dcl.IsEmptyValueIndirect(initial.MetricsGcpServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.MetricsGcpServiceAccountEmail = initial.MetricsGcpServiceAccountEmail
	} else {
		cDes.MetricsGcpServiceAccountEmail = des.MetricsGcpServiceAccountEmail
	}
	cDes.Oci = canonicalizeFeatureMembershipConfigmanagementConfigSyncOci(des.Oci, initial.Oci, opts...)

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncSlice(des, initial []FeatureMembershipConfigmanagementConfigSync, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementConfigSync, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementConfigSync(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementConfigSync, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementConfigSync(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c *Client, des, nw *FeatureMembershipConfigmanagementConfigSync) *FeatureMembershipConfigmanagementConfigSync {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementConfigSync while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Git = canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c, des.Git, nw.Git)
	if dcl.StringCanonicalize(des.SourceFormat, nw.SourceFormat) {
		nw.SourceFormat = des.SourceFormat
	}
	if dcl.BoolCanonicalize(des.PreventDrift, nw.PreventDrift) {
		nw.PreventDrift = des.PreventDrift
	}
	nw.Oci = canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOci(c, des.Oci, nw.Oci)

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncSet(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSync) []FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementConfigSync
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementConfigSyncNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSync) []FeatureMembershipConfigmanagementConfigSync {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementConfigSync
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSync(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(des, initial *FeatureMembershipConfigmanagementConfigSyncGit, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementConfigSyncGit{}

	if dcl.StringCanonicalize(des.SyncRepo, initial.SyncRepo) || dcl.IsZeroValue(des.SyncRepo) {
		cDes.SyncRepo = initial.SyncRepo
	} else {
		cDes.SyncRepo = des.SyncRepo
	}
	if dcl.StringCanonicalize(des.SyncBranch, initial.SyncBranch) || dcl.IsZeroValue(des.SyncBranch) {
		cDes.SyncBranch = initial.SyncBranch
	} else {
		cDes.SyncBranch = des.SyncBranch
	}
	if dcl.StringCanonicalize(des.PolicyDir, initial.PolicyDir) || dcl.IsZeroValue(des.PolicyDir) {
		cDes.PolicyDir = initial.PolicyDir
	} else {
		cDes.PolicyDir = des.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, initial.SyncWaitSecs) || dcl.IsZeroValue(des.SyncWaitSecs) {
		cDes.SyncWaitSecs = initial.SyncWaitSecs
	} else {
		cDes.SyncWaitSecs = des.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SyncRev, initial.SyncRev) || dcl.IsZeroValue(des.SyncRev) {
		cDes.SyncRev = initial.SyncRev
	} else {
		cDes.SyncRev = des.SyncRev
	}
	if dcl.StringCanonicalize(des.SecretType, initial.SecretType) || dcl.IsZeroValue(des.SecretType) {
		cDes.SecretType = initial.SecretType
	} else {
		cDes.SecretType = des.SecretType
	}
	if dcl.StringCanonicalize(des.HttpsProxy, initial.HttpsProxy) || dcl.IsZeroValue(des.HttpsProxy) {
		cDes.HttpsProxy = initial.HttpsProxy
	} else {
		cDes.HttpsProxy = des.HttpsProxy
	}
	if dcl.IsZeroValue(des.GcpServiceAccountEmail) || (dcl.IsEmptyValueIndirect(des.GcpServiceAccountEmail) && dcl.IsEmptyValueIndirect(initial.GcpServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GcpServiceAccountEmail = initial.GcpServiceAccountEmail
	} else {
		cDes.GcpServiceAccountEmail = des.GcpServiceAccountEmail
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncGitSlice(des, initial []FeatureMembershipConfigmanagementConfigSyncGit, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementConfigSyncGit {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementConfigSyncGit, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementConfigSyncGit, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementConfigSyncGit(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c *Client, des, nw *FeatureMembershipConfigmanagementConfigSyncGit) *FeatureMembershipConfigmanagementConfigSyncGit {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementConfigSyncGit while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SyncRepo, nw.SyncRepo) {
		nw.SyncRepo = des.SyncRepo
	}
	if dcl.StringCanonicalize(des.SyncBranch, nw.SyncBranch) {
		nw.SyncBranch = des.SyncBranch
	}
	if dcl.StringCanonicalize(des.PolicyDir, nw.PolicyDir) {
		nw.PolicyDir = des.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, nw.SyncWaitSecs) {
		nw.SyncWaitSecs = des.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SyncRev, nw.SyncRev) {
		nw.SyncRev = des.SyncRev
	}
	if dcl.StringCanonicalize(des.SecretType, nw.SecretType) {
		nw.SecretType = des.SecretType
	}
	if dcl.StringCanonicalize(des.HttpsProxy, nw.HttpsProxy) {
		nw.HttpsProxy = des.HttpsProxy
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGitSet(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncGit) []FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementConfigSyncGit
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncGit) []FeatureMembershipConfigmanagementConfigSyncGit {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementConfigSyncGit
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSyncGit(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncOci(des, initial *FeatureMembershipConfigmanagementConfigSyncOci, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementConfigSyncOci {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementConfigSyncOci{}

	if dcl.StringCanonicalize(des.SyncRepo, initial.SyncRepo) || dcl.IsZeroValue(des.SyncRepo) {
		cDes.SyncRepo = initial.SyncRepo
	} else {
		cDes.SyncRepo = des.SyncRepo
	}
	if dcl.StringCanonicalize(des.PolicyDir, initial.PolicyDir) || dcl.IsZeroValue(des.PolicyDir) {
		cDes.PolicyDir = initial.PolicyDir
	} else {
		cDes.PolicyDir = des.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, initial.SyncWaitSecs) || dcl.IsZeroValue(des.SyncWaitSecs) {
		cDes.SyncWaitSecs = initial.SyncWaitSecs
	} else {
		cDes.SyncWaitSecs = des.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SecretType, initial.SecretType) || dcl.IsZeroValue(des.SecretType) {
		cDes.SecretType = initial.SecretType
	} else {
		cDes.SecretType = des.SecretType
	}
	if dcl.IsZeroValue(des.GcpServiceAccountEmail) || (dcl.IsEmptyValueIndirect(des.GcpServiceAccountEmail) && dcl.IsEmptyValueIndirect(initial.GcpServiceAccountEmail)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.GcpServiceAccountEmail = initial.GcpServiceAccountEmail
	} else {
		cDes.GcpServiceAccountEmail = des.GcpServiceAccountEmail
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementConfigSyncOciSlice(des, initial []FeatureMembershipConfigmanagementConfigSyncOci, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementConfigSyncOci {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementConfigSyncOci, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementConfigSyncOci(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementConfigSyncOci, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementConfigSyncOci(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOci(c *Client, des, nw *FeatureMembershipConfigmanagementConfigSyncOci) *FeatureMembershipConfigmanagementConfigSyncOci {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementConfigSyncOci while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.SyncRepo, nw.SyncRepo) {
		nw.SyncRepo = des.SyncRepo
	}
	if dcl.StringCanonicalize(des.PolicyDir, nw.PolicyDir) {
		nw.PolicyDir = des.PolicyDir
	}
	if dcl.StringCanonicalize(des.SyncWaitSecs, nw.SyncWaitSecs) {
		nw.SyncWaitSecs = des.SyncWaitSecs
	}
	if dcl.StringCanonicalize(des.SecretType, nw.SecretType) {
		nw.SecretType = des.SecretType
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOciSet(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncOci) []FeatureMembershipConfigmanagementConfigSyncOci {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementConfigSyncOci
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementConfigSyncOciNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOci(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOciSlice(c *Client, des, nw []FeatureMembershipConfigmanagementConfigSyncOci) []FeatureMembershipConfigmanagementConfigSyncOci {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementConfigSyncOci
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementConfigSyncOci(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementPolicyController(des, initial *FeatureMembershipConfigmanagementPolicyController, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementPolicyController{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}
	if dcl.StringArrayCanonicalize(des.ExemptableNamespaces, initial.ExemptableNamespaces) {
		cDes.ExemptableNamespaces = initial.ExemptableNamespaces
	} else {
		cDes.ExemptableNamespaces = des.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, initial.ReferentialRulesEnabled) || dcl.IsZeroValue(des.ReferentialRulesEnabled) {
		cDes.ReferentialRulesEnabled = initial.ReferentialRulesEnabled
	} else {
		cDes.ReferentialRulesEnabled = des.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, initial.LogDeniesEnabled) || dcl.IsZeroValue(des.LogDeniesEnabled) {
		cDes.LogDeniesEnabled = initial.LogDeniesEnabled
	} else {
		cDes.LogDeniesEnabled = des.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.MutationEnabled, initial.MutationEnabled) || dcl.IsZeroValue(des.MutationEnabled) {
		cDes.MutationEnabled = initial.MutationEnabled
	} else {
		cDes.MutationEnabled = des.MutationEnabled
	}
	cDes.Monitoring = canonicalizeFeatureMembershipConfigmanagementPolicyControllerMonitoring(des.Monitoring, initial.Monitoring, opts...)
	if dcl.BoolCanonicalize(des.TemplateLibraryInstalled, initial.TemplateLibraryInstalled) || dcl.IsZeroValue(des.TemplateLibraryInstalled) {
		cDes.TemplateLibraryInstalled = initial.TemplateLibraryInstalled
	} else {
		cDes.TemplateLibraryInstalled = des.TemplateLibraryInstalled
	}
	if dcl.StringCanonicalize(des.AuditIntervalSeconds, initial.AuditIntervalSeconds) || dcl.IsZeroValue(des.AuditIntervalSeconds) {
		cDes.AuditIntervalSeconds = initial.AuditIntervalSeconds
	} else {
		cDes.AuditIntervalSeconds = des.AuditIntervalSeconds
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementPolicyControllerSlice(des, initial []FeatureMembershipConfigmanagementPolicyController, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementPolicyController {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementPolicyController, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementPolicyController(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementPolicyController, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementPolicyController(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c *Client, des, nw *FeatureMembershipConfigmanagementPolicyController) *FeatureMembershipConfigmanagementPolicyController {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementPolicyController while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.StringArrayCanonicalize(des.ExemptableNamespaces, nw.ExemptableNamespaces) {
		nw.ExemptableNamespaces = des.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, nw.ReferentialRulesEnabled) {
		nw.ReferentialRulesEnabled = des.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, nw.LogDeniesEnabled) {
		nw.LogDeniesEnabled = des.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.MutationEnabled, nw.MutationEnabled) {
		nw.MutationEnabled = des.MutationEnabled
	}
	nw.Monitoring = canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, des.Monitoring, nw.Monitoring)
	if dcl.BoolCanonicalize(des.TemplateLibraryInstalled, nw.TemplateLibraryInstalled) {
		nw.TemplateLibraryInstalled = des.TemplateLibraryInstalled
	}
	if dcl.StringCanonicalize(des.AuditIntervalSeconds, nw.AuditIntervalSeconds) {
		nw.AuditIntervalSeconds = des.AuditIntervalSeconds
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerSet(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyController) []FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementPolicyController
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementPolicyControllerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyController) []FeatureMembershipConfigmanagementPolicyController {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementPolicyController
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementPolicyController(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementPolicyControllerMonitoring(des, initial *FeatureMembershipConfigmanagementPolicyControllerMonitoring, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementPolicyControllerMonitoring{}

	if dcl.IsZeroValue(des.Backends) || (dcl.IsEmptyValueIndirect(des.Backends) && dcl.IsEmptyValueIndirect(initial.Backends)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Backends = initial.Backends
	} else {
		cDes.Backends = des.Backends
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice(des, initial []FeatureMembershipConfigmanagementPolicyControllerMonitoring, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementPolicyControllerMonitoring, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementPolicyControllerMonitoring(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementPolicyControllerMonitoring, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementPolicyControllerMonitoring(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoring(c *Client, des, nw *FeatureMembershipConfigmanagementPolicyControllerMonitoring) *FeatureMembershipConfigmanagementPolicyControllerMonitoring {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementPolicyControllerMonitoring while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoringSet(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyControllerMonitoring) []FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementPolicyControllerMonitoring
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementPolicyControllerMonitoringNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice(c *Client, des, nw []FeatureMembershipConfigmanagementPolicyControllerMonitoring) []FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementPolicyControllerMonitoring
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementBinauthz(des, initial *FeatureMembershipConfigmanagementBinauthz, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementBinauthz{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementBinauthzSlice(des, initial []FeatureMembershipConfigmanagementBinauthz, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementBinauthz {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementBinauthz, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementBinauthz(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementBinauthz, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementBinauthz(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c *Client, des, nw *FeatureMembershipConfigmanagementBinauthz) *FeatureMembershipConfigmanagementBinauthz {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementBinauthz while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementBinauthzSet(c *Client, des, nw []FeatureMembershipConfigmanagementBinauthz) []FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementBinauthz
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementBinauthzNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementBinauthzSlice(c *Client, des, nw []FeatureMembershipConfigmanagementBinauthz) []FeatureMembershipConfigmanagementBinauthz {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementBinauthz
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementBinauthz(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipConfigmanagementHierarchyController(des, initial *FeatureMembershipConfigmanagementHierarchyController, opts ...dcl.ApplyOption) *FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipConfigmanagementHierarchyController{}

	if dcl.BoolCanonicalize(des.Enabled, initial.Enabled) || dcl.IsZeroValue(des.Enabled) {
		cDes.Enabled = initial.Enabled
	} else {
		cDes.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.EnablePodTreeLabels, initial.EnablePodTreeLabels) || dcl.IsZeroValue(des.EnablePodTreeLabels) {
		cDes.EnablePodTreeLabels = initial.EnablePodTreeLabels
	} else {
		cDes.EnablePodTreeLabels = des.EnablePodTreeLabels
	}
	if dcl.BoolCanonicalize(des.EnableHierarchicalResourceQuota, initial.EnableHierarchicalResourceQuota) || dcl.IsZeroValue(des.EnableHierarchicalResourceQuota) {
		cDes.EnableHierarchicalResourceQuota = initial.EnableHierarchicalResourceQuota
	} else {
		cDes.EnableHierarchicalResourceQuota = des.EnableHierarchicalResourceQuota
	}

	return cDes
}

func canonicalizeFeatureMembershipConfigmanagementHierarchyControllerSlice(des, initial []FeatureMembershipConfigmanagementHierarchyController, opts ...dcl.ApplyOption) []FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipConfigmanagementHierarchyController, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipConfigmanagementHierarchyController(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipConfigmanagementHierarchyController, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipConfigmanagementHierarchyController(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c *Client, des, nw *FeatureMembershipConfigmanagementHierarchyController) *FeatureMembershipConfigmanagementHierarchyController {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipConfigmanagementHierarchyController while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.BoolCanonicalize(des.Enabled, nw.Enabled) {
		nw.Enabled = des.Enabled
	}
	if dcl.BoolCanonicalize(des.EnablePodTreeLabels, nw.EnablePodTreeLabels) {
		nw.EnablePodTreeLabels = des.EnablePodTreeLabels
	}
	if dcl.BoolCanonicalize(des.EnableHierarchicalResourceQuota, nw.EnableHierarchicalResourceQuota) {
		nw.EnableHierarchicalResourceQuota = des.EnableHierarchicalResourceQuota
	}

	return nw
}

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyControllerSet(c *Client, des, nw []FeatureMembershipConfigmanagementHierarchyController) []FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipConfigmanagementHierarchyController
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, des, nw []FeatureMembershipConfigmanagementHierarchyController) []FeatureMembershipConfigmanagementHierarchyController {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipConfigmanagementHierarchyController
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipConfigmanagementHierarchyController(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipPolicycontroller(des, initial *FeatureMembershipPolicycontroller, opts ...dcl.ApplyOption) *FeatureMembershipPolicycontroller {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipPolicycontroller{}

	if dcl.StringCanonicalize(des.Version, initial.Version) || dcl.IsZeroValue(des.Version) {
		cDes.Version = initial.Version
	} else {
		cDes.Version = des.Version
	}
	cDes.PolicyControllerHubConfig = canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfig(des.PolicyControllerHubConfig, initial.PolicyControllerHubConfig, opts...)

	return cDes
}

func canonicalizeFeatureMembershipPolicycontrollerSlice(des, initial []FeatureMembershipPolicycontroller, opts ...dcl.ApplyOption) []FeatureMembershipPolicycontroller {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipPolicycontroller, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipPolicycontroller(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipPolicycontroller, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipPolicycontroller(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipPolicycontroller(c *Client, des, nw *FeatureMembershipPolicycontroller) *FeatureMembershipPolicycontroller {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipPolicycontroller while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Version, nw.Version) {
		nw.Version = des.Version
	}
	nw.PolicyControllerHubConfig = canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, des.PolicyControllerHubConfig, nw.PolicyControllerHubConfig)

	return nw
}

func canonicalizeNewFeatureMembershipPolicycontrollerSet(c *Client, des, nw []FeatureMembershipPolicycontroller) []FeatureMembershipPolicycontroller {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipPolicycontroller
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipPolicycontrollerNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipPolicycontroller(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipPolicycontrollerSlice(c *Client, des, nw []FeatureMembershipPolicycontroller) []FeatureMembershipPolicycontroller {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipPolicycontroller
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipPolicycontroller(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfig(des, initial *FeatureMembershipPolicycontrollerPolicyControllerHubConfig, opts ...dcl.ApplyOption) *FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}

	if dcl.IsZeroValue(des.InstallSpec) || (dcl.IsEmptyValueIndirect(des.InstallSpec) && dcl.IsEmptyValueIndirect(initial.InstallSpec)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.InstallSpec = initial.InstallSpec
	} else {
		cDes.InstallSpec = des.InstallSpec
	}
	if dcl.StringArrayCanonicalize(des.ExemptableNamespaces, initial.ExemptableNamespaces) {
		cDes.ExemptableNamespaces = initial.ExemptableNamespaces
	} else {
		cDes.ExemptableNamespaces = des.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, initial.ReferentialRulesEnabled) || dcl.IsZeroValue(des.ReferentialRulesEnabled) {
		cDes.ReferentialRulesEnabled = initial.ReferentialRulesEnabled
	} else {
		cDes.ReferentialRulesEnabled = des.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, initial.LogDeniesEnabled) || dcl.IsZeroValue(des.LogDeniesEnabled) {
		cDes.LogDeniesEnabled = initial.LogDeniesEnabled
	} else {
		cDes.LogDeniesEnabled = des.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.MutationEnabled, initial.MutationEnabled) || dcl.IsZeroValue(des.MutationEnabled) {
		cDes.MutationEnabled = initial.MutationEnabled
	} else {
		cDes.MutationEnabled = des.MutationEnabled
	}
	cDes.Monitoring = canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(des.Monitoring, initial.Monitoring, opts...)
	if dcl.IsZeroValue(des.AuditIntervalSeconds) || (dcl.IsEmptyValueIndirect(des.AuditIntervalSeconds) && dcl.IsEmptyValueIndirect(initial.AuditIntervalSeconds)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.AuditIntervalSeconds = initial.AuditIntervalSeconds
	} else {
		cDes.AuditIntervalSeconds = des.AuditIntervalSeconds
	}
	if dcl.IsZeroValue(des.ConstraintViolationLimit) || (dcl.IsEmptyValueIndirect(des.ConstraintViolationLimit) && dcl.IsEmptyValueIndirect(initial.ConstraintViolationLimit)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.ConstraintViolationLimit = initial.ConstraintViolationLimit
	} else {
		cDes.ConstraintViolationLimit = des.ConstraintViolationLimit
	}
	cDes.PolicyContent = canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(des.PolicyContent, initial.PolicyContent, opts...)

	return cDes
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice(des, initial []FeatureMembershipPolicycontrollerPolicyControllerHubConfig, opts ...dcl.ApplyOption) []FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfig, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfig(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfig, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfig(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c *Client, des, nw *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) *FeatureMembershipPolicycontrollerPolicyControllerHubConfig {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipPolicycontrollerPolicyControllerHubConfig while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.ExemptableNamespaces, nw.ExemptableNamespaces) {
		nw.ExemptableNamespaces = des.ExemptableNamespaces
	}
	if dcl.BoolCanonicalize(des.ReferentialRulesEnabled, nw.ReferentialRulesEnabled) {
		nw.ReferentialRulesEnabled = des.ReferentialRulesEnabled
	}
	if dcl.BoolCanonicalize(des.LogDeniesEnabled, nw.LogDeniesEnabled) {
		nw.LogDeniesEnabled = des.LogDeniesEnabled
	}
	if dcl.BoolCanonicalize(des.MutationEnabled, nw.MutationEnabled) {
		nw.MutationEnabled = des.MutationEnabled
	}
	nw.Monitoring = canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, des.Monitoring, nw.Monitoring)
	nw.PolicyContent = canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, des.PolicyContent, nw.PolicyContent)

	return nw
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigSet(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfig) []FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfig
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfig) []FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfig
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(des, initial *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, opts ...dcl.ApplyOption) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}

	if dcl.IsZeroValue(des.Backends) || (dcl.IsEmptyValueIndirect(des.Backends) && dcl.IsEmptyValueIndirect(initial.Backends)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Backends = initial.Backends
	} else {
		cDes.Backends = des.Backends
	}

	return cDes
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice(des, initial []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, opts ...dcl.ApplyOption) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c *Client, des, nw *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSet(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(des, initial *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, opts ...dcl.ApplyOption) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}

	cDes.TemplateLibrary = canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(des.TemplateLibrary, initial.TemplateLibrary, opts...)

	return cDes
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice(des, initial []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, opts ...dcl.ApplyOption) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c *Client, des, nw *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.TemplateLibrary = canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, des.TemplateLibrary, nw.TemplateLibrary)

	return nw
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSet(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, &d, &n))
	}

	return items
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(des, initial *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, opts ...dcl.ApplyOption) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}

	if dcl.IsZeroValue(des.Installation) || (dcl.IsEmptyValueIndirect(des.Installation) && dcl.IsEmptyValueIndirect(initial.Installation)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		cDes.Installation = initial.Installation
	} else {
		cDes.Installation = des.Installation
	}

	return cDes
}

func canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice(des, initial []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, opts ...dcl.ApplyOption) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if dcl.IsEmptyValueIndirect(des) {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, 0, len(des))
		for _, d := range des {
			cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, 0, len(des))
	for i, d := range des {
		cd := canonicalizeFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c *Client, des, nw *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsEmptyValueIndirect(des) {
			c.Config.Logger.Info("Found explicitly empty value for FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySet(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if des == nil {
		return nw
	}

	// Find the elements in des that are also in nw and canonicalize them. Remove matched elements from nw.
	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
	for _, d := range des {
		matchedIndex := -1
		for i, n := range nw {
			if diffs, _ := compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedIndex = i
				break
			}
		}
		if matchedIndex != -1 {
			items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, &d, &nw[matchedIndex]))
			nw = append(nw[:matchedIndex], nw[matchedIndex+1:]...)
		}
	}
	// Also include elements in nw that are not matched in des.
	items = append(items, nw...)

	return items
}

func canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice(c *Client, des, nw []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, &d, &n))
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
func diffFeatureMembership(c *Client, desired, actual *FeatureMembership, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Mesh, actual.Mesh, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipMeshNewStyle, EmptyObject: EmptyFeatureMembershipMesh, OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("Mesh")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Configmanagement, actual.Configmanagement, dcl.DiffInfo{MergeNestedDiffs: true, ObjectFunction: compareFeatureMembershipConfigmanagementNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagement, OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("Configmanagement")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Policycontroller, actual.Policycontroller, dcl.DiffInfo{MergeNestedDiffs: true, ObjectFunction: compareFeatureMembershipPolicycontrollerNewStyle, EmptyObject: EmptyFeatureMembershipPolicycontroller, OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("Policycontroller")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.Feature, actual.Feature, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Feature")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Membership, actual.Membership, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Membership")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MembershipLocation, actual.MembershipLocation, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MembershipLocation")); len(ds) != 0 || err != nil {
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
func compareFeatureMembershipMeshNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipMesh)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipMesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipMesh or *FeatureMembershipMesh", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipMesh)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipMesh)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipMesh", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Management, actual.Management, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("Management")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ControlPlane, actual.ControlPlane, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("ControlPlane")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagement)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagement or *FeatureMembershipConfigmanagement", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagement)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagement)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagement", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ConfigSync, actual.ConfigSync, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipConfigmanagementConfigSyncNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementConfigSync, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConfigSync")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyController, actual.PolicyController, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipConfigmanagementPolicyControllerNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementPolicyController, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyController")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Binauthz, actual.Binauthz, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareFeatureMembershipConfigmanagementBinauthzNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementBinauthz, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Binauthz")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HierarchyController, actual.HierarchyController, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementHierarchyController, CustomDiff: emptyHNCSameAsAllFalse, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HierarchyController")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementConfigSyncNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementConfigSync)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementConfigSync)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSync or *FeatureMembershipConfigmanagementConfigSync", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementConfigSync)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementConfigSync)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSync", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Git, actual.Git, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementConfigSyncGit, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Git")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceFormat, actual.SourceFormat, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceFormat")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PreventDrift, actual.PreventDrift, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PreventDrift")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MetricsGcpServiceAccountEmail, actual.MetricsGcpServiceAccountEmail, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MetricsGcpServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Oci, actual.Oci, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipConfigmanagementConfigSyncOciNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementConfigSyncOci, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Oci")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementConfigSyncGitNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementConfigSyncGit)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementConfigSyncGit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncGit or *FeatureMembershipConfigmanagementConfigSyncGit", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementConfigSyncGit)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementConfigSyncGit)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncGit", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SyncRepo, actual.SyncRepo, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncRepo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncBranch, actual.SyncBranch, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncBranch")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyDir, actual.PolicyDir, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyDir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncWaitSecs, actual.SyncWaitSecs, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncWaitSecs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncRev, actual.SyncRev, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SyncRev")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretType, actual.SecretType, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SecretType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.HttpsProxy, actual.HttpsProxy, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("HttpsProxy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcpServiceAccountEmail, actual.GcpServiceAccountEmail, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GcpServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementConfigSyncOciNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementConfigSyncOci)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementConfigSyncOci)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncOci or *FeatureMembershipConfigmanagementConfigSyncOci", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementConfigSyncOci)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementConfigSyncOci)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementConfigSyncOci", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.SyncRepo, actual.SyncRepo, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("SyncRepo")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyDir, actual.PolicyDir, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("PolicyDir")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SyncWaitSecs, actual.SyncWaitSecs, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("SyncWaitSecs")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SecretType, actual.SecretType, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("SecretType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GcpServiceAccountEmail, actual.GcpServiceAccountEmail, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateFeatureMembershipUpdateFeatureMembershipOperation")}, fn.AddNest("GcpServiceAccountEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementPolicyControllerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementPolicyController)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementPolicyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyController or *FeatureMembershipConfigmanagementPolicyController", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementPolicyController)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementPolicyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyController", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemptableNamespaces, actual.ExemptableNamespaces, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemptableNamespaces")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReferentialRulesEnabled, actual.ReferentialRulesEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReferentialRulesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogDeniesEnabled, actual.LogDeniesEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LogDeniesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MutationEnabled, actual.MutationEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutationEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Monitoring, actual.Monitoring, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareFeatureMembershipConfigmanagementPolicyControllerMonitoringNewStyle, EmptyObject: EmptyFeatureMembershipConfigmanagementPolicyControllerMonitoring, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Monitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.TemplateLibraryInstalled, actual.TemplateLibraryInstalled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TemplateLibraryInstalled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuditIntervalSeconds, actual.AuditIntervalSeconds, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuditIntervalSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementPolicyControllerMonitoringNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementPolicyControllerMonitoring)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementPolicyControllerMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyControllerMonitoring or *FeatureMembershipConfigmanagementPolicyControllerMonitoring", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementPolicyControllerMonitoring)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementPolicyControllerMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementPolicyControllerMonitoring", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Backends, actual.Backends, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Backends")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementBinauthzNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementBinauthz)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementBinauthz)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementBinauthz or *FeatureMembershipConfigmanagementBinauthz", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementBinauthz)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementBinauthz)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementBinauthz", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipConfigmanagementHierarchyControllerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController or *FeatureMembershipConfigmanagementHierarchyController", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipConfigmanagementHierarchyController)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipConfigmanagementHierarchyController)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipConfigmanagementHierarchyController", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Enabled, actual.Enabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Enabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnablePodTreeLabels, actual.EnablePodTreeLabels, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnablePodTreeLabels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnableHierarchicalResourceQuota, actual.EnableHierarchicalResourceQuota, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("EnableHierarchicalResourceQuota")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipPolicycontrollerNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipPolicycontroller)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipPolicycontroller)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontroller or *FeatureMembershipPolicycontroller", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipPolicycontroller)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipPolicycontroller)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontroller", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Version, actual.Version, dcl.DiffInfo{ServerDefault: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Version")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyControllerHubConfig, actual.PolicyControllerHubConfig, dcl.DiffInfo{ObjectFunction: compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigNewStyle, EmptyObject: EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfig, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyControllerHubConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfig)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipPolicycontrollerPolicyControllerHubConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfig or *FeatureMembershipPolicycontrollerPolicyControllerHubConfig", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfig)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipPolicycontrollerPolicyControllerHubConfig)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfig", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.InstallSpec, actual.InstallSpec, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InstallSpec")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ExemptableNamespaces, actual.ExemptableNamespaces, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ExemptableNamespaces")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ReferentialRulesEnabled, actual.ReferentialRulesEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ReferentialRulesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.LogDeniesEnabled, actual.LogDeniesEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("LogDeniesEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.MutationEnabled, actual.MutationEnabled, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("MutationEnabled")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Monitoring, actual.Monitoring, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringNewStyle, EmptyObject: EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Monitoring")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AuditIntervalSeconds, actual.AuditIntervalSeconds, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AuditIntervalSeconds")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ConstraintViolationLimit, actual.ConstraintViolationLimit, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ConstraintViolationLimit")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PolicyContent, actual.PolicyContent, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentNewStyle, EmptyObject: EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("PolicyContent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring or *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Backends, actual.Backends, dcl.DiffInfo{ServerDefault: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Backends")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent or *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.TemplateLibrary, actual.TemplateLibrary, dcl.DiffInfo{ServerDefault: true, ObjectFunction: compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryNewStyle, EmptyObject: EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("TemplateLibrary")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary)
	if !ok {
		desiredNotPointer, ok := d.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary or *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary)
	if !ok {
		actualNotPointer, ok := a.(FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Installation, actual.Installation, dcl.DiffInfo{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Installation")); len(ds) != 0 || err != nil {
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
func (r *FeatureMembership) urlNormalized() *FeatureMembership {
	normalized := dcl.Copy(*r).(FeatureMembership)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	normalized.Feature = dcl.SelfLinkToName(r.Feature)
	normalized.Membership = dcl.SelfLinkToName(r.Membership)
	normalized.MembershipLocation = dcl.SelfLinkToName(r.MembershipLocation)
	return &normalized
}

func (r *FeatureMembership) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateFeatureMembership" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"feature":  dcl.ValueOrEmptyString(nr.Feature),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/features/{{feature}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the FeatureMembership resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *FeatureMembership) marshal(c *Client) ([]byte, error) {
	m, err := expandFeatureMembership(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling FeatureMembership: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalFeatureMembership decodes JSON responses into the FeatureMembership resource schema.
func unmarshalFeatureMembership(b []byte, c *Client, res *FeatureMembership) (*FeatureMembership, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapFeatureMembership(m, c, res)
}

func unmarshalMapFeatureMembership(m map[string]interface{}, c *Client, res *FeatureMembership) (*FeatureMembership, error) {

	flattened := flattenFeatureMembership(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandFeatureMembership expands FeatureMembership into a JSON request object.
func expandFeatureMembership(c *Client, f *FeatureMembership) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := expandFeatureMembershipMesh(c, f.Mesh, res); err != nil {
		return nil, fmt.Errorf("error expanding Mesh into mesh: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["mesh"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagement(c, f.Configmanagement, res); err != nil {
		return nil, fmt.Errorf("error expanding Configmanagement into configmanagement: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["configmanagement"] = v
	}
	if v, err := expandFeatureMembershipPolicycontroller(c, f.Policycontroller, res); err != nil {
		return nil, fmt.Errorf("error expanding Policycontroller into policycontroller: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policycontroller"] = v
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
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Feature into feature: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["feature"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Membership into membership: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["membership"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding MembershipLocation into membershipLocation: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["membershipLocation"] = v
	}

	return m, nil
}

// flattenFeatureMembership flattens FeatureMembership from a JSON request object into the
// FeatureMembership type.
func flattenFeatureMembership(c *Client, i interface{}, res *FeatureMembership) *FeatureMembership {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &FeatureMembership{}
	resultRes.Mesh = flattenFeatureMembershipMesh(c, m["mesh"], res)
	resultRes.Configmanagement = flattenFeatureMembershipConfigmanagement(c, m["configmanagement"], res)
	resultRes.Policycontroller = flattenFeatureMembershipPolicycontroller(c, m["policycontroller"], res)
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])
	resultRes.Feature = dcl.FlattenString(m["feature"])
	resultRes.Membership = dcl.FlattenString(m["membership"])
	resultRes.MembershipLocation = dcl.FlattenString(m["membershipLocation"])

	return resultRes
}

// expandFeatureMembershipMeshMap expands the contents of FeatureMembershipMesh into a JSON
// request object.
func expandFeatureMembershipMeshMap(c *Client, f map[string]FeatureMembershipMesh, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipMesh(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipMeshSlice expands the contents of FeatureMembershipMesh into a JSON
// request object.
func expandFeatureMembershipMeshSlice(c *Client, f []FeatureMembershipMesh, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipMesh(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipMeshMap flattens the contents of FeatureMembershipMesh from a JSON
// response object.
func flattenFeatureMembershipMeshMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipMesh {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipMesh{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipMesh{}
	}

	items := make(map[string]FeatureMembershipMesh)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipMesh(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipMeshSlice flattens the contents of FeatureMembershipMesh from a JSON
// response object.
func flattenFeatureMembershipMeshSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipMesh {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipMesh{}
	}

	if len(a) == 0 {
		return []FeatureMembershipMesh{}
	}

	items := make([]FeatureMembershipMesh, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipMesh(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipMesh expands an instance of FeatureMembershipMesh into a JSON
// request object.
func expandFeatureMembershipMesh(c *Client, f *FeatureMembershipMesh, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Management; !dcl.IsEmptyValueIndirect(v) {
		m["management"] = v
	}
	if v := f.ControlPlane; !dcl.IsEmptyValueIndirect(v) {
		m["controlPlane"] = v
	}

	return m, nil
}

// flattenFeatureMembershipMesh flattens an instance of FeatureMembershipMesh from a JSON
// response object.
func flattenFeatureMembershipMesh(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipMesh {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipMesh{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipMesh
	}
	r.Management = flattenFeatureMembershipMeshManagementEnum(m["management"])
	r.ControlPlane = flattenFeatureMembershipMeshControlPlaneEnum(m["controlPlane"])

	return r
}

// expandFeatureMembershipConfigmanagementMap expands the contents of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagementMap(c *Client, f map[string]FeatureMembershipConfigmanagement, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagement(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementSlice expands the contents of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagementSlice(c *Client, f []FeatureMembershipConfigmanagement, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagement(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementMap flattens the contents of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagement {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagement{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagement{}
	}

	items := make(map[string]FeatureMembershipConfigmanagement)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagement(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementSlice flattens the contents of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagement {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagement{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagement{}
	}

	items := make([]FeatureMembershipConfigmanagement, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagement(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagement expands an instance of FeatureMembershipConfigmanagement into a JSON
// request object.
func expandFeatureMembershipConfigmanagement(c *Client, f *FeatureMembershipConfigmanagement, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureMembershipConfigmanagementConfigSync(c, f.ConfigSync, res); err != nil {
		return nil, fmt.Errorf("error expanding ConfigSync into configSync: %w", err)
	} else if v != nil {
		m["configSync"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementPolicyController(c, f.PolicyController, res); err != nil {
		return nil, fmt.Errorf("error expanding PolicyController into policyController: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyController"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementBinauthz(c, f.Binauthz, res); err != nil {
		return nil, fmt.Errorf("error expanding Binauthz into binauthz: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["binauthz"] = v
	}
	if v, err := expandHierarchyControllerConfig(c, f.HierarchyController, res); err != nil {
		return nil, fmt.Errorf("error expanding HierarchyController into hierarchyController: %w", err)
	} else if v != nil {
		m["hierarchyController"] = v
	}
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagement flattens an instance of FeatureMembershipConfigmanagement from a JSON
// response object.
func flattenFeatureMembershipConfigmanagement(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagement {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagement{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagement
	}
	r.ConfigSync = flattenFeatureMembershipConfigmanagementConfigSync(c, m["configSync"], res)
	r.PolicyController = flattenFeatureMembershipConfigmanagementPolicyController(c, m["policyController"], res)
	r.Binauthz = flattenFeatureMembershipConfigmanagementBinauthz(c, m["binauthz"], res)
	r.HierarchyController = flattenHierarchyControllerConfig(c, m["hierarchyController"], res)
	r.Version = dcl.FlattenString(m["version"])

	return r
}

// expandFeatureMembershipConfigmanagementConfigSyncMap expands the contents of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncMap(c *Client, f map[string]FeatureMembershipConfigmanagementConfigSync, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSync(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementConfigSyncSlice expands the contents of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, f []FeatureMembershipConfigmanagementConfigSync, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSync(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncMap flattens the contents of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementConfigSync {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementConfigSync{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementConfigSync{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementConfigSync)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementConfigSync(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementConfigSyncSlice flattens the contents of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementConfigSync {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementConfigSync{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementConfigSync{}
	}

	items := make([]FeatureMembershipConfigmanagementConfigSync, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementConfigSync(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementConfigSync expands an instance of FeatureMembershipConfigmanagementConfigSync into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSync(c *Client, f *FeatureMembershipConfigmanagementConfigSync, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, f.Git, res); err != nil {
		return nil, fmt.Errorf("error expanding Git into git: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["git"] = v
	}
	if v := f.SourceFormat; !dcl.IsEmptyValueIndirect(v) {
		m["sourceFormat"] = v
	}
	if v := f.PreventDrift; !dcl.IsEmptyValueIndirect(v) {
		m["preventDrift"] = v
	}
	if v := f.MetricsGcpServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["metricsGcpServiceAccountEmail"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementConfigSyncOci(c, f.Oci, res); err != nil {
		return nil, fmt.Errorf("error expanding Oci into oci: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["oci"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementConfigSync flattens an instance of FeatureMembershipConfigmanagementConfigSync from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSync(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementConfigSync {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementConfigSync{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementConfigSync
	}
	r.Git = flattenFeatureMembershipConfigmanagementConfigSyncGit(c, m["git"], res)
	r.SourceFormat = dcl.FlattenString(m["sourceFormat"])
	r.PreventDrift = dcl.FlattenBool(m["preventDrift"])
	r.MetricsGcpServiceAccountEmail = dcl.FlattenString(m["metricsGcpServiceAccountEmail"])
	r.Oci = flattenFeatureMembershipConfigmanagementConfigSyncOci(c, m["oci"], res)

	return r
}

// expandFeatureMembershipConfigmanagementConfigSyncGitMap expands the contents of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGitMap(c *Client, f map[string]FeatureMembershipConfigmanagementConfigSyncGit, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementConfigSyncGitSlice expands the contents of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, f []FeatureMembershipConfigmanagementConfigSyncGit, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncGit(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncGitMap flattens the contents of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGitMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementConfigSyncGit {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementConfigSyncGit)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementConfigSyncGit(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementConfigSyncGitSlice flattens the contents of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGitSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementConfigSyncGit {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementConfigSyncGit{}
	}

	items := make([]FeatureMembershipConfigmanagementConfigSyncGit, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementConfigSyncGit(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementConfigSyncGit expands an instance of FeatureMembershipConfigmanagementConfigSyncGit into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncGit(c *Client, f *FeatureMembershipConfigmanagementConfigSyncGit, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SyncRepo; !dcl.IsEmptyValueIndirect(v) {
		m["syncRepo"] = v
	}
	if v := f.SyncBranch; !dcl.IsEmptyValueIndirect(v) {
		m["syncBranch"] = v
	}
	if v := f.PolicyDir; !dcl.IsEmptyValueIndirect(v) {
		m["policyDir"] = v
	}
	if v := f.SyncWaitSecs; !dcl.IsEmptyValueIndirect(v) {
		m["syncWaitSecs"] = v
	}
	if v := f.SyncRev; !dcl.IsEmptyValueIndirect(v) {
		m["syncRev"] = v
	}
	if v := f.SecretType; !dcl.IsEmptyValueIndirect(v) {
		m["secretType"] = v
	}
	if v := f.HttpsProxy; !dcl.IsEmptyValueIndirect(v) {
		m["httpsProxy"] = v
	}
	if v := f.GcpServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["gcpServiceAccountEmail"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncGit flattens an instance of FeatureMembershipConfigmanagementConfigSyncGit from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncGit(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementConfigSyncGit {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementConfigSyncGit{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementConfigSyncGit
	}
	r.SyncRepo = dcl.FlattenString(m["syncRepo"])
	r.SyncBranch = dcl.FlattenString(m["syncBranch"])
	r.PolicyDir = dcl.FlattenString(m["policyDir"])
	r.SyncWaitSecs = dcl.FlattenString(m["syncWaitSecs"])
	r.SyncRev = dcl.FlattenString(m["syncRev"])
	r.SecretType = dcl.FlattenString(m["secretType"])
	r.HttpsProxy = dcl.FlattenString(m["httpsProxy"])
	r.GcpServiceAccountEmail = dcl.FlattenString(m["gcpServiceAccountEmail"])

	return r
}

// expandFeatureMembershipConfigmanagementConfigSyncOciMap expands the contents of FeatureMembershipConfigmanagementConfigSyncOci into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncOciMap(c *Client, f map[string]FeatureMembershipConfigmanagementConfigSyncOci, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncOci(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementConfigSyncOciSlice expands the contents of FeatureMembershipConfigmanagementConfigSyncOci into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncOciSlice(c *Client, f []FeatureMembershipConfigmanagementConfigSyncOci, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementConfigSyncOci(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncOciMap flattens the contents of FeatureMembershipConfigmanagementConfigSyncOci from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncOciMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementConfigSyncOci {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementConfigSyncOci{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementConfigSyncOci{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementConfigSyncOci)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementConfigSyncOci(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementConfigSyncOciSlice flattens the contents of FeatureMembershipConfigmanagementConfigSyncOci from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncOciSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementConfigSyncOci {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementConfigSyncOci{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementConfigSyncOci{}
	}

	items := make([]FeatureMembershipConfigmanagementConfigSyncOci, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementConfigSyncOci(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementConfigSyncOci expands an instance of FeatureMembershipConfigmanagementConfigSyncOci into a JSON
// request object.
func expandFeatureMembershipConfigmanagementConfigSyncOci(c *Client, f *FeatureMembershipConfigmanagementConfigSyncOci, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.SyncRepo; !dcl.IsEmptyValueIndirect(v) {
		m["syncRepo"] = v
	}
	if v := f.PolicyDir; !dcl.IsEmptyValueIndirect(v) {
		m["policyDir"] = v
	}
	if v := f.SyncWaitSecs; !dcl.IsEmptyValueIndirect(v) {
		m["syncWaitSecs"] = v
	}
	if v := f.SecretType; !dcl.IsEmptyValueIndirect(v) {
		m["secretType"] = v
	}
	if v := f.GcpServiceAccountEmail; !dcl.IsEmptyValueIndirect(v) {
		m["gcpServiceAccountEmail"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementConfigSyncOci flattens an instance of FeatureMembershipConfigmanagementConfigSyncOci from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementConfigSyncOci(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementConfigSyncOci {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementConfigSyncOci{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementConfigSyncOci
	}
	r.SyncRepo = dcl.FlattenString(m["syncRepo"])
	r.PolicyDir = dcl.FlattenString(m["policyDir"])
	r.SyncWaitSecs = dcl.FlattenString(m["syncWaitSecs"])
	r.SecretType = dcl.FlattenString(m["secretType"])
	r.GcpServiceAccountEmail = dcl.FlattenString(m["gcpServiceAccountEmail"])

	return r
}

// expandFeatureMembershipConfigmanagementPolicyControllerMap expands the contents of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerMap(c *Client, f map[string]FeatureMembershipConfigmanagementPolicyController, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyController(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementPolicyControllerSlice expands the contents of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, f []FeatureMembershipConfigmanagementPolicyController, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyController(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMap flattens the contents of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementPolicyController {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementPolicyController{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementPolicyController{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementPolicyController)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementPolicyController(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementPolicyControllerSlice flattens the contents of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementPolicyController {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementPolicyController{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementPolicyController{}
	}

	items := make([]FeatureMembershipConfigmanagementPolicyController, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementPolicyController(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementPolicyController expands an instance of FeatureMembershipConfigmanagementPolicyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyController(c *Client, f *FeatureMembershipConfigmanagementPolicyController, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}
	if v := f.ExemptableNamespaces; v != nil {
		m["exemptableNamespaces"] = v
	}
	if v := f.ReferentialRulesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["referentialRulesEnabled"] = v
	}
	if v := f.LogDeniesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["logDeniesEnabled"] = v
	}
	if v := f.MutationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["mutationEnabled"] = v
	}
	if v, err := expandFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, f.Monitoring, res); err != nil {
		return nil, fmt.Errorf("error expanding Monitoring into monitoring: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["monitoring"] = v
	}
	if v := f.TemplateLibraryInstalled; !dcl.IsEmptyValueIndirect(v) {
		m["templateLibraryInstalled"] = v
	}
	if v := f.AuditIntervalSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["auditIntervalSeconds"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementPolicyController flattens an instance of FeatureMembershipConfigmanagementPolicyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyController(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementPolicyController {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementPolicyController{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementPolicyController
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.ExemptableNamespaces = dcl.FlattenStringSlice(m["exemptableNamespaces"])
	r.ReferentialRulesEnabled = dcl.FlattenBool(m["referentialRulesEnabled"])
	r.LogDeniesEnabled = dcl.FlattenBool(m["logDeniesEnabled"])
	r.MutationEnabled = dcl.FlattenBool(m["mutationEnabled"])
	r.Monitoring = flattenFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, m["monitoring"], res)
	r.TemplateLibraryInstalled = dcl.FlattenBool(m["templateLibraryInstalled"])
	r.AuditIntervalSeconds = dcl.FlattenString(m["auditIntervalSeconds"])

	return r
}

// expandFeatureMembershipConfigmanagementPolicyControllerMonitoringMap expands the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoring into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerMonitoringMap(c *Client, f map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoring, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice expands the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoring into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice(c *Client, f []FeatureMembershipConfigmanagementPolicyControllerMonitoring, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringMap flattens the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoring from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoring)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice flattens the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoring from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}

	items := make([]FeatureMembershipConfigmanagementPolicyControllerMonitoring, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementPolicyControllerMonitoring(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementPolicyControllerMonitoring expands an instance of FeatureMembershipConfigmanagementPolicyControllerMonitoring into a JSON
// request object.
func expandFeatureMembershipConfigmanagementPolicyControllerMonitoring(c *Client, f *FeatureMembershipConfigmanagementPolicyControllerMonitoring, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Backends; v != nil {
		m["backends"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoring flattens an instance of FeatureMembershipConfigmanagementPolicyControllerMonitoring from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoring(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementPolicyControllerMonitoring {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementPolicyControllerMonitoring{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementPolicyControllerMonitoring
	}
	r.Backends = flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumSlice(c, m["backends"], res)

	return r
}

// expandFeatureMembershipConfigmanagementBinauthzMap expands the contents of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthzMap(c *Client, f map[string]FeatureMembershipConfigmanagementBinauthz, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementBinauthz(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementBinauthzSlice expands the contents of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthzSlice(c *Client, f []FeatureMembershipConfigmanagementBinauthz, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementBinauthz(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementBinauthzMap flattens the contents of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthzMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementBinauthz {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementBinauthz{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementBinauthz{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementBinauthz)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementBinauthz(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementBinauthzSlice flattens the contents of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthzSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementBinauthz {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementBinauthz{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementBinauthz{}
	}

	items := make([]FeatureMembershipConfigmanagementBinauthz, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementBinauthz(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementBinauthz expands an instance of FeatureMembershipConfigmanagementBinauthz into a JSON
// request object.
func expandFeatureMembershipConfigmanagementBinauthz(c *Client, f *FeatureMembershipConfigmanagementBinauthz, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; !dcl.IsEmptyValueIndirect(v) {
		m["enabled"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementBinauthz flattens an instance of FeatureMembershipConfigmanagementBinauthz from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementBinauthz(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementBinauthz {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementBinauthz{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementBinauthz
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])

	return r
}

// expandFeatureMembershipConfigmanagementHierarchyControllerMap expands the contents of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyControllerMap(c *Client, f map[string]FeatureMembershipConfigmanagementHierarchyController, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipConfigmanagementHierarchyController(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipConfigmanagementHierarchyControllerSlice expands the contents of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, f []FeatureMembershipConfigmanagementHierarchyController, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipConfigmanagementHierarchyController(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipConfigmanagementHierarchyControllerMap flattens the contents of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyControllerMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementHierarchyController {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementHierarchyController{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementHierarchyController{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementHierarchyController)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementHierarchyController(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipConfigmanagementHierarchyControllerSlice flattens the contents of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyControllerSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementHierarchyController {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementHierarchyController{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementHierarchyController{}
	}

	items := make([]FeatureMembershipConfigmanagementHierarchyController, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementHierarchyController(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipConfigmanagementHierarchyController expands an instance of FeatureMembershipConfigmanagementHierarchyController into a JSON
// request object.
func expandFeatureMembershipConfigmanagementHierarchyController(c *Client, f *FeatureMembershipConfigmanagementHierarchyController, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Enabled; v != nil {
		m["enabled"] = v
	}
	if v := f.EnablePodTreeLabels; v != nil {
		m["enablePodTreeLabels"] = v
	}
	if v := f.EnableHierarchicalResourceQuota; v != nil {
		m["enableHierarchicalResourceQuota"] = v
	}

	return m, nil
}

// flattenFeatureMembershipConfigmanagementHierarchyController flattens an instance of FeatureMembershipConfigmanagementHierarchyController from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementHierarchyController(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipConfigmanagementHierarchyController {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipConfigmanagementHierarchyController{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipConfigmanagementHierarchyController
	}
	r.Enabled = dcl.FlattenBool(m["enabled"])
	r.EnablePodTreeLabels = dcl.FlattenBool(m["enablePodTreeLabels"])
	r.EnableHierarchicalResourceQuota = dcl.FlattenBool(m["enableHierarchicalResourceQuota"])

	return r
}

// expandFeatureMembershipPolicycontrollerMap expands the contents of FeatureMembershipPolicycontroller into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerMap(c *Client, f map[string]FeatureMembershipPolicycontroller, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipPolicycontroller(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipPolicycontrollerSlice expands the contents of FeatureMembershipPolicycontroller into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerSlice(c *Client, f []FeatureMembershipPolicycontroller, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipPolicycontroller(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipPolicycontrollerMap flattens the contents of FeatureMembershipPolicycontroller from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontroller {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontroller{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontroller{}
	}

	items := make(map[string]FeatureMembershipPolicycontroller)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontroller(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerSlice flattens the contents of FeatureMembershipPolicycontroller from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontroller {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontroller{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontroller{}
	}

	items := make([]FeatureMembershipPolicycontroller, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontroller(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipPolicycontroller expands an instance of FeatureMembershipPolicycontroller into a JSON
// request object.
func expandFeatureMembershipPolicycontroller(c *Client, f *FeatureMembershipPolicycontroller, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Version; !dcl.IsEmptyValueIndirect(v) {
		m["version"] = v
	}
	if v, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, f.PolicyControllerHubConfig, res); err != nil {
		return nil, fmt.Errorf("error expanding PolicyControllerHubConfig into policyControllerHubConfig: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyControllerHubConfig"] = v
	}

	return m, nil
}

// flattenFeatureMembershipPolicycontroller flattens an instance of FeatureMembershipPolicycontroller from a JSON
// response object.
func flattenFeatureMembershipPolicycontroller(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipPolicycontroller {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipPolicycontroller{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipPolicycontroller
	}
	r.Version = dcl.FlattenString(m["version"])
	r.PolicyControllerHubConfig = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, m["policyControllerHubConfig"], res)

	return r
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMap expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfig into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMap(c *Client, f map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfig, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfig into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice(c *Client, f []FeatureMembershipPolicycontrollerPolicyControllerHubConfig, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfig from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfig)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfig from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfig, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfig expands an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfig into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c *Client, f *FeatureMembershipPolicycontrollerPolicyControllerHubConfig, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.InstallSpec; !dcl.IsEmptyValueIndirect(v) {
		m["installSpec"] = v
	}
	if v := f.ExemptableNamespaces; v != nil {
		m["exemptableNamespaces"] = v
	}
	if v := f.ReferentialRulesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["referentialRulesEnabled"] = v
	}
	if v := f.LogDeniesEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["logDeniesEnabled"] = v
	}
	if v := f.MutationEnabled; !dcl.IsEmptyValueIndirect(v) {
		m["mutationEnabled"] = v
	}
	if v, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, f.Monitoring, res); err != nil {
		return nil, fmt.Errorf("error expanding Monitoring into monitoring: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["monitoring"] = v
	}
	if v := f.AuditIntervalSeconds; !dcl.IsEmptyValueIndirect(v) {
		m["auditIntervalSeconds"] = v
	}
	if v := f.ConstraintViolationLimit; !dcl.IsEmptyValueIndirect(v) {
		m["constraintViolationLimit"] = v
	}
	if v, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, f.PolicyContent, res); err != nil {
		return nil, fmt.Errorf("error expanding PolicyContent into policyContent: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["policyContent"] = v
	}

	return m, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfig flattens an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfig from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfig(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipPolicycontrollerPolicyControllerHubConfig {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfig
	}
	r.InstallSpec = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(m["installSpec"])
	r.ExemptableNamespaces = dcl.FlattenStringSlice(m["exemptableNamespaces"])
	r.ReferentialRulesEnabled = dcl.FlattenBool(m["referentialRulesEnabled"])
	r.LogDeniesEnabled = dcl.FlattenBool(m["logDeniesEnabled"])
	r.MutationEnabled = dcl.FlattenBool(m["mutationEnabled"])
	r.Monitoring = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, m["monitoring"], res)
	r.AuditIntervalSeconds = dcl.FlattenInteger(m["auditIntervalSeconds"])
	r.ConstraintViolationLimit = dcl.FlattenInteger(m["constraintViolationLimit"])
	r.PolicyContent = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, m["policyContent"], res)

	return r
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringMap expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringMap(c *Client, f map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice(c *Client, f []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring expands an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c *Client, f *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Backends; v != nil {
		m["backends"] = v
	}

	return m, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring flattens an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring
	}
	r.Backends = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumSlice(c, m["backends"], res)

	return r
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentMap expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentMap(c *Client, f map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice(c *Client, f []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent expands an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c *Client, f *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, f.TemplateLibrary, res); err != nil {
		return nil, fmt.Errorf("error expanding TemplateLibrary into templateLibrary: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["templateLibrary"] = v
	}

	return m, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent flattens an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent
	}
	r.TemplateLibrary = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, m["templateLibrary"], res)

	return r
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryMap expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryMap(c *Client, f map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, res *FeatureMembership) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, &item, res)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice expands the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice(c *Client, f []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, res *FeatureMembership) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, &item, res)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, item.(map[string]interface{}), res)
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrarySlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c, item.(map[string]interface{}), res))
	}

	return items
}

// expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary expands an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary into a JSON
// request object.
func expandFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c *Client, f *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary, res *FeatureMembership) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Installation; !dcl.IsEmptyValueIndirect(v) {
		m["installation"] = v
	}

	return m, nil
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary flattens an instance of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary(c *Client, i interface{}, res *FeatureMembership) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary
	}
	r.Installation = flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(m["installation"])

	return r
}

// flattenFeatureMembershipMeshManagementEnumMap flattens the contents of FeatureMembershipMeshManagementEnum from a JSON
// response object.
func flattenFeatureMembershipMeshManagementEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipMeshManagementEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipMeshManagementEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipMeshManagementEnum{}
	}

	items := make(map[string]FeatureMembershipMeshManagementEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipMeshManagementEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipMeshManagementEnumSlice flattens the contents of FeatureMembershipMeshManagementEnum from a JSON
// response object.
func flattenFeatureMembershipMeshManagementEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipMeshManagementEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipMeshManagementEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipMeshManagementEnum{}
	}

	items := make([]FeatureMembershipMeshManagementEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipMeshManagementEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipMeshManagementEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipMeshManagementEnum with the same value as that string.
func flattenFeatureMembershipMeshManagementEnum(i interface{}) *FeatureMembershipMeshManagementEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipMeshManagementEnumRef(s)
}

// flattenFeatureMembershipMeshControlPlaneEnumMap flattens the contents of FeatureMembershipMeshControlPlaneEnum from a JSON
// response object.
func flattenFeatureMembershipMeshControlPlaneEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipMeshControlPlaneEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipMeshControlPlaneEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipMeshControlPlaneEnum{}
	}

	items := make(map[string]FeatureMembershipMeshControlPlaneEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipMeshControlPlaneEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipMeshControlPlaneEnumSlice flattens the contents of FeatureMembershipMeshControlPlaneEnum from a JSON
// response object.
func flattenFeatureMembershipMeshControlPlaneEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipMeshControlPlaneEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipMeshControlPlaneEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipMeshControlPlaneEnum{}
	}

	items := make([]FeatureMembershipMeshControlPlaneEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipMeshControlPlaneEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipMeshControlPlaneEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipMeshControlPlaneEnum with the same value as that string.
func flattenFeatureMembershipMeshControlPlaneEnum(i interface{}) *FeatureMembershipMeshControlPlaneEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipMeshControlPlaneEnumRef(s)
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumMap flattens the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum{}
	}

	items := make(map[string]FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumSlice flattens the contents of FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum from a JSON
// response object.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum{}
	}

	items := make([]FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum with the same value as that string.
func flattenFeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum(i interface{}) *FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnumRef(s)
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum with the same value as that string.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum(i interface{}) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipPolicycontrollerPolicyControllerHubConfigInstallSpecEnumRef(s)
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum with the same value as that string.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum(i interface{}) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringBackendsEnumRef(s)
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumMap flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumMap(c *Client, i interface{}, res *FeatureMembership) map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum{}
	}

	if len(a) == 0 {
		return map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum{}
	}

	items := make(map[string]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum)
	for k, item := range a {
		items[k] = *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(item.(interface{}))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumSlice flattens the contents of FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum from a JSON
// response object.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumSlice(c *Client, i interface{}, res *FeatureMembership) []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum{}
	}

	if len(a) == 0 {
		return []FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum{}
	}

	items := make([]FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(item.(interface{})))
	}

	return items
}

// flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum asserts that an interface is a string, and returns a
// pointer to a *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum with the same value as that string.
func flattenFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum(i interface{}) *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryInstallationEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *FeatureMembership) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalFeatureMembership(b, c, r)
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
		if nr.Feature == nil && ncr.Feature == nil {
			c.Config.Logger.Info("Both Feature fields null - considering equal.")
		} else if nr.Feature == nil || ncr.Feature == nil {
			c.Config.Logger.Info("Only one Feature field is null - considering unequal.")
			return false
		} else if *nr.Feature != *ncr.Feature {
			return false
		}
		return true
	}
}

type featureMembershipDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         featureMembershipApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToFeatureMembershipDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]featureMembershipDiff, error) {
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
	var diffs []featureMembershipDiff
	// For each operation name, create a featureMembershipDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := featureMembershipDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToFeatureMembershipApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToFeatureMembershipApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (featureMembershipApiOperation, error) {
	switch opName {

	case "updateFeatureMembershipUpdateFeatureMembershipOperation":
		return &updateFeatureMembershipUpdateFeatureMembershipOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractFeatureMembershipFields(r *FeatureMembership) error {
	vMesh := r.Mesh
	if vMesh == nil {
		// note: explicitly not the empty object.
		vMesh = &FeatureMembershipMesh{}
	}
	if err := extractFeatureMembershipMeshFields(r, vMesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMesh) {
		r.Mesh = vMesh
	}
	vConfigmanagement := r.Configmanagement
	if vConfigmanagement == nil {
		// note: explicitly not the empty object.
		vConfigmanagement = &FeatureMembershipConfigmanagement{}
	}
	if err := extractFeatureMembershipConfigmanagementFields(r, vConfigmanagement); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigmanagement) {
		r.Configmanagement = vConfigmanagement
	}
	vPolicycontroller := r.Policycontroller
	if vPolicycontroller == nil {
		// note: explicitly not the empty object.
		vPolicycontroller = &FeatureMembershipPolicycontroller{}
	}
	if err := extractFeatureMembershipPolicycontrollerFields(r, vPolicycontroller); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicycontroller) {
		r.Policycontroller = vPolicycontroller
	}
	return nil
}
func extractFeatureMembershipMeshFields(r *FeatureMembership, o *FeatureMembershipMesh) error {
	return nil
}
func extractFeatureMembershipConfigmanagementFields(r *FeatureMembership, o *FeatureMembershipConfigmanagement) error {
	vConfigSync := o.ConfigSync
	if vConfigSync == nil {
		// note: explicitly not the empty object.
		vConfigSync = &FeatureMembershipConfigmanagementConfigSync{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncFields(r, vConfigSync); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigSync) {
		o.ConfigSync = vConfigSync
	}
	vPolicyController := o.PolicyController
	if vPolicyController == nil {
		// note: explicitly not the empty object.
		vPolicyController = &FeatureMembershipConfigmanagementPolicyController{}
	}
	if err := extractFeatureMembershipConfigmanagementPolicyControllerFields(r, vPolicyController); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyController) {
		o.PolicyController = vPolicyController
	}
	vBinauthz := o.Binauthz
	if vBinauthz == nil {
		// note: explicitly not the empty object.
		vBinauthz = &FeatureMembershipConfigmanagementBinauthz{}
	}
	if err := extractFeatureMembershipConfigmanagementBinauthzFields(r, vBinauthz); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinauthz) {
		o.Binauthz = vBinauthz
	}
	vHierarchyController := o.HierarchyController
	if vHierarchyController == nil {
		// note: explicitly not the empty object.
		vHierarchyController = &FeatureMembershipConfigmanagementHierarchyController{}
	}
	if err := extractFeatureMembershipConfigmanagementHierarchyControllerFields(r, vHierarchyController); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHierarchyController) {
		o.HierarchyController = vHierarchyController
	}
	return nil
}
func extractFeatureMembershipConfigmanagementConfigSyncFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSync) error {
	vGit := o.Git
	if vGit == nil {
		// note: explicitly not the empty object.
		vGit = &FeatureMembershipConfigmanagementConfigSyncGit{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncGitFields(r, vGit); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGit) {
		o.Git = vGit
	}
	vOci := o.Oci
	if vOci == nil {
		// note: explicitly not the empty object.
		vOci = &FeatureMembershipConfigmanagementConfigSyncOci{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncOciFields(r, vOci); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOci) {
		o.Oci = vOci
	}
	return nil
}
func extractFeatureMembershipConfigmanagementConfigSyncGitFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSyncGit) error {
	return nil
}
func extractFeatureMembershipConfigmanagementConfigSyncOciFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSyncOci) error {
	return nil
}
func extractFeatureMembershipConfigmanagementPolicyControllerFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementPolicyController) error {
	vMonitoring := o.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}
	if err := extractFeatureMembershipConfigmanagementPolicyControllerMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		o.Monitoring = vMonitoring
	}
	return nil
}
func extractFeatureMembershipConfigmanagementPolicyControllerMonitoringFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementPolicyControllerMonitoring) error {
	return nil
}
func extractFeatureMembershipConfigmanagementBinauthzFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementBinauthz) error {
	return nil
}
func extractFeatureMembershipConfigmanagementHierarchyControllerFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementHierarchyController) error {
	return nil
}
func extractFeatureMembershipPolicycontrollerFields(r *FeatureMembership, o *FeatureMembershipPolicycontroller) error {
	vPolicyControllerHubConfig := o.PolicyControllerHubConfig
	if vPolicyControllerHubConfig == nil {
		// note: explicitly not the empty object.
		vPolicyControllerHubConfig = &FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigFields(r, vPolicyControllerHubConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyControllerHubConfig) {
		o.PolicyControllerHubConfig = vPolicyControllerHubConfig
	}
	return nil
}
func extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) error {
	vMonitoring := o.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		o.Monitoring = vMonitoring
	}
	vPolicyContent := o.PolicyContent
	if vPolicyContent == nil {
		// note: explicitly not the empty object.
		vPolicyContent = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentFields(r, vPolicyContent); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyContent) {
		o.PolicyContent = vPolicyContent
	}
	return nil
}
func extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) error {
	return nil
}
func extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) error {
	vTemplateLibrary := o.TemplateLibrary
	if vTemplateLibrary == nil {
		// note: explicitly not the empty object.
		vTemplateLibrary = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryFields(r, vTemplateLibrary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplateLibrary) {
		o.TemplateLibrary = vTemplateLibrary
	}
	return nil
}
func extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) error {
	return nil
}

func postReadExtractFeatureMembershipFields(r *FeatureMembership) error {
	vMesh := r.Mesh
	if vMesh == nil {
		// note: explicitly not the empty object.
		vMesh = &FeatureMembershipMesh{}
	}
	if err := postReadExtractFeatureMembershipMeshFields(r, vMesh); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMesh) {
		r.Mesh = vMesh
	}
	vConfigmanagement := r.Configmanagement
	if vConfigmanagement == nil {
		// note: explicitly not the empty object.
		vConfigmanagement = &FeatureMembershipConfigmanagement{}
	}
	if err := postReadExtractFeatureMembershipConfigmanagementFields(r, vConfigmanagement); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigmanagement) {
		r.Configmanagement = vConfigmanagement
	}
	vPolicycontroller := r.Policycontroller
	if vPolicycontroller == nil {
		// note: explicitly not the empty object.
		vPolicycontroller = &FeatureMembershipPolicycontroller{}
	}
	if err := postReadExtractFeatureMembershipPolicycontrollerFields(r, vPolicycontroller); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicycontroller) {
		r.Policycontroller = vPolicycontroller
	}
	return nil
}
func postReadExtractFeatureMembershipMeshFields(r *FeatureMembership, o *FeatureMembershipMesh) error {
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementFields(r *FeatureMembership, o *FeatureMembershipConfigmanagement) error {
	vConfigSync := o.ConfigSync
	if vConfigSync == nil {
		// note: explicitly not the empty object.
		vConfigSync = &FeatureMembershipConfigmanagementConfigSync{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncFields(r, vConfigSync); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vConfigSync) {
		o.ConfigSync = vConfigSync
	}
	vPolicyController := o.PolicyController
	if vPolicyController == nil {
		// note: explicitly not the empty object.
		vPolicyController = &FeatureMembershipConfigmanagementPolicyController{}
	}
	if err := extractFeatureMembershipConfigmanagementPolicyControllerFields(r, vPolicyController); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyController) {
		o.PolicyController = vPolicyController
	}
	vBinauthz := o.Binauthz
	if vBinauthz == nil {
		// note: explicitly not the empty object.
		vBinauthz = &FeatureMembershipConfigmanagementBinauthz{}
	}
	if err := extractFeatureMembershipConfigmanagementBinauthzFields(r, vBinauthz); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vBinauthz) {
		o.Binauthz = vBinauthz
	}
	vHierarchyController := o.HierarchyController
	if vHierarchyController == nil {
		// note: explicitly not the empty object.
		vHierarchyController = &FeatureMembershipConfigmanagementHierarchyController{}
	}
	if err := extractFeatureMembershipConfigmanagementHierarchyControllerFields(r, vHierarchyController); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vHierarchyController) {
		o.HierarchyController = vHierarchyController
	}
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementConfigSyncFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSync) error {
	vGit := o.Git
	if vGit == nil {
		// note: explicitly not the empty object.
		vGit = &FeatureMembershipConfigmanagementConfigSyncGit{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncGitFields(r, vGit); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vGit) {
		o.Git = vGit
	}
	vOci := o.Oci
	if vOci == nil {
		// note: explicitly not the empty object.
		vOci = &FeatureMembershipConfigmanagementConfigSyncOci{}
	}
	if err := extractFeatureMembershipConfigmanagementConfigSyncOciFields(r, vOci); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vOci) {
		o.Oci = vOci
	}
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementConfigSyncGitFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSyncGit) error {
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementConfigSyncOciFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementConfigSyncOci) error {
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementPolicyControllerFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementPolicyController) error {
	vMonitoring := o.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &FeatureMembershipConfigmanagementPolicyControllerMonitoring{}
	}
	if err := extractFeatureMembershipConfigmanagementPolicyControllerMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		o.Monitoring = vMonitoring
	}
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementPolicyControllerMonitoringFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementPolicyControllerMonitoring) error {
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementBinauthzFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementBinauthz) error {
	return nil
}
func postReadExtractFeatureMembershipConfigmanagementHierarchyControllerFields(r *FeatureMembership, o *FeatureMembershipConfigmanagementHierarchyController) error {
	return nil
}
func postReadExtractFeatureMembershipPolicycontrollerFields(r *FeatureMembership, o *FeatureMembershipPolicycontroller) error {
	vPolicyControllerHubConfig := o.PolicyControllerHubConfig
	if vPolicyControllerHubConfig == nil {
		// note: explicitly not the empty object.
		vPolicyControllerHubConfig = &FeatureMembershipPolicycontrollerPolicyControllerHubConfig{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigFields(r, vPolicyControllerHubConfig); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyControllerHubConfig) {
		o.PolicyControllerHubConfig = vPolicyControllerHubConfig
	}
	return nil
}
func postReadExtractFeatureMembershipPolicycontrollerPolicyControllerHubConfigFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfig) error {
	vMonitoring := o.Monitoring
	if vMonitoring == nil {
		// note: explicitly not the empty object.
		vMonitoring = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringFields(r, vMonitoring); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vMonitoring) {
		o.Monitoring = vMonitoring
	}
	vPolicyContent := o.PolicyContent
	if vPolicyContent == nil {
		// note: explicitly not the empty object.
		vPolicyContent = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentFields(r, vPolicyContent); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vPolicyContent) {
		o.PolicyContent = vPolicyContent
	}
	return nil
}
func postReadExtractFeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoringFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigMonitoring) error {
	return nil
}
func postReadExtractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContent) error {
	vTemplateLibrary := o.TemplateLibrary
	if vTemplateLibrary == nil {
		// note: explicitly not the empty object.
		vTemplateLibrary = &FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary{}
	}
	if err := extractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryFields(r, vTemplateLibrary); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(vTemplateLibrary) {
		o.TemplateLibrary = vTemplateLibrary
	}
	return nil
}
func postReadExtractFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibraryFields(r *FeatureMembership, o *FeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary) error {
	return nil
}
