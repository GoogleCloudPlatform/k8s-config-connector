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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Backup) validate() error {

	if err := dcl.RequiredParameter(r.Name, "Name"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sourceInstance"); err != nil {
		return err
	}
	if err := dcl.Required(r, "sourceFileShare"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Project, "Project"); err != nil {
		return err
	}
	if err := dcl.RequiredParameter(r.Location, "Location"); err != nil {
		return err
	}
	return nil
}
func (r *Backup) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://file.googleapis.com/v1beta1/", params)
}

func (r *Backup) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Backup) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups", nr.basePath(), userBasePath, params), nil

}

func (r *Backup) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups?backupId={{name}}", nr.basePath(), userBasePath, params), nil

}

func (r *Backup) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project":  dcl.ValueOrEmptyString(nr.Project),
		"location": dcl.ValueOrEmptyString(nr.Location),
		"name":     dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", nr.basePath(), userBasePath, params), nil
}

// backupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type backupApiOperation interface {
	do(context.Context, *Backup, *Client) error
}

// newUpdateBackupUpdateBackupRequest creates a request for an
// Backup resource's UpdateBackup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateBackupUpdateBackupRequest(ctx context.Context, f *Backup, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}
	res := f
	_ = res

	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	return req, nil
}

// marshalUpdateBackupUpdateBackupRequest converts the update into
// the final JSON request body.
func marshalUpdateBackupUpdateBackupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateBackupUpdateBackupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateBackupUpdateBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {
	_, err := c.GetBackup(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateBackup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateBackupUpdateBackupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateBackupUpdateBackupRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listBackupRaw(ctx context.Context, r *Backup, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BackupMaxPage {
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

type listBackupOperation struct {
	Backups []map[string]interface{} `json:"backups"`
	Token   string                   `json:"nextPageToken"`
}

func (c *Client) listBackup(ctx context.Context, r *Backup, pageToken string, pageSize int32) ([]*Backup, string, error) {
	b, err := c.listBackupRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBackupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Backup
	for _, v := range m.Backups {
		res, err := unmarshalMapBackup(v, c, r)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = r.Project
		res.Location = r.Location
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllBackup(ctx context.Context, f func(*Backup) bool, resources []*Backup) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteBackup(ctx, res)
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

type deleteBackupOperation struct{}

func (op *deleteBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {
	r, err := c.GetBackup(ctx, r)
	if err != nil {
		if dcl.IsNotFound(err) {
			c.Config.Logger.InfoWithContextf(ctx, "Backup not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetBackup checking for existence. error: %v", err)
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
		_, err := c.GetBackup(ctx, r)
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
type createBackupOperation struct {
	response map[string]interface{}
}

func (op *createBackupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBackupOperation) do(ctx context.Context, r *Backup, c *Client) error {
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

	if _, err := c.GetBackup(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBackupRaw(ctx context.Context, r *Backup) ([]byte, error) {

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

func (c *Client) backupDiffsForRawDesired(ctx context.Context, rawDesired *Backup, opts ...dcl.ApplyOption) (initial, desired *Backup, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Backup
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Backup); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Backup, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBackup(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Backup resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Backup resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Backup resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBackupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Backup: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Backup: %v", rawDesired)

	// The Get call applies postReadExtract and so the result may contain fields that are not part of API version.
	if err := extractBackupFields(rawInitial); err != nil {
		return nil, nil, nil, err
	}

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBackupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Backup: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBackupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Backup: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBackup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBackupInitialState(rawInitial, rawDesired *Backup) (*Backup, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBackupDesiredState(rawDesired, rawInitial *Backup, opts ...dcl.ApplyOption) (*Backup, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Backup{}
	if dcl.NameToSelfLink(rawDesired.Name, rawInitial.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) || (dcl.IsEmptyValueIndirect(rawDesired.Labels) && dcl.IsEmptyValueIndirect(rawInitial.Labels)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	if dcl.IsZeroValue(rawDesired.SourceInstance) || (dcl.IsEmptyValueIndirect(rawDesired.SourceInstance) && dcl.IsEmptyValueIndirect(rawInitial.SourceInstance)) {
		// Desired and initial values are equivalent, so set canonical desired value to initial value.
		canonicalDesired.SourceInstance = rawInitial.SourceInstance
	} else {
		canonicalDesired.SourceInstance = rawDesired.SourceInstance
	}
	if dcl.StringCanonicalize(rawDesired.SourceFileShare, rawInitial.SourceFileShare) {
		canonicalDesired.SourceFileShare = rawInitial.SourceFileShare
	} else {
		canonicalDesired.SourceFileShare = rawDesired.SourceFileShare
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

func canonicalizeBackupNewState(c *Client, rawNew, rawDesired *Backup) (*Backup, error) {

	rawNew.Name = rawDesired.Name

	if dcl.IsEmptyValueIndirect(rawNew.Description) && dcl.IsEmptyValueIndirect(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.State) && dcl.IsEmptyValueIndirect(rawDesired.State) {
		rawNew.State = rawDesired.State
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CreateTime) && dcl.IsEmptyValueIndirect(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.Labels) && dcl.IsEmptyValueIndirect(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.CapacityGb) && dcl.IsEmptyValueIndirect(rawDesired.CapacityGb) {
		rawNew.CapacityGb = rawDesired.CapacityGb
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.StorageBytes) && dcl.IsEmptyValueIndirect(rawDesired.StorageBytes) {
		rawNew.StorageBytes = rawDesired.StorageBytes
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceInstance) && dcl.IsEmptyValueIndirect(rawDesired.SourceInstance) {
		rawNew.SourceInstance = rawDesired.SourceInstance
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceFileShare) && dcl.IsEmptyValueIndirect(rawDesired.SourceFileShare) {
		rawNew.SourceFileShare = rawDesired.SourceFileShare
	} else {
		if dcl.StringCanonicalize(rawDesired.SourceFileShare, rawNew.SourceFileShare) {
			rawNew.SourceFileShare = rawDesired.SourceFileShare
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SourceInstanceTier) && dcl.IsEmptyValueIndirect(rawDesired.SourceInstanceTier) {
		rawNew.SourceInstanceTier = rawDesired.SourceInstanceTier
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.DownloadBytes) && dcl.IsEmptyValueIndirect(rawDesired.DownloadBytes) {
		rawNew.DownloadBytes = rawDesired.DownloadBytes
	} else {
	}

	rawNew.Project = rawDesired.Project

	rawNew.Location = rawDesired.Location

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffBackup(c *Client, desired, actual *Backup, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
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

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBackupUpdateBackupOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
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

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updateBackupUpdateBackupOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CapacityGb, actual.CapacityGb, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CapacityGb")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StorageBytes, actual.StorageBytes, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("StorageBytes")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceInstance, actual.SourceInstance, dcl.DiffInfo{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceInstance")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceFileShare, actual.SourceFileShare, dcl.DiffInfo{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceFileShare")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SourceInstanceTier, actual.SourceInstanceTier, dcl.DiffInfo{OutputOnly: true, Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SourceInstanceTier")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DownloadBytes, actual.DownloadBytes, dcl.DiffInfo{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("DownloadBytes")); len(ds) != 0 || err != nil {
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

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Backup) urlNormalized() *Backup {
	normalized := dcl.Copy(*r).(Backup)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SourceInstance = dcl.SelfLinkToName(r.SourceInstance)
	normalized.SourceFileShare = dcl.SelfLinkToName(r.SourceFileShare)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	normalized.Location = dcl.SelfLinkToName(r.Location)
	return &normalized
}

func (r *Backup) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateBackup" {
		fields := map[string]interface{}{
			"project":  dcl.ValueOrEmptyString(nr.Project),
			"location": dcl.ValueOrEmptyString(nr.Location),
			"name":     dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("projects/{{project}}/locations/{{location}}/backups/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Backup resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Backup) marshal(c *Client) ([]byte, error) {
	m, err := expandBackup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Backup: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBackup decodes JSON responses into the Backup resource schema.
func unmarshalBackup(b []byte, c *Client, res *Backup) (*Backup, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBackup(m, c, res)
}

func unmarshalMapBackup(m map[string]interface{}, c *Client, res *Backup) (*Backup, error) {

	flattened := flattenBackup(c, m, res)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandBackup expands Backup into a JSON request object.
func expandBackup(c *Client, f *Backup) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	res := f
	_ = res
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v := f.SourceInstance; dcl.ValueShouldBeSent(v) {
		m["sourceInstance"] = v
	}
	if v := f.SourceFileShare; dcl.ValueShouldBeSent(v) {
		m["sourceFileShare"] = v
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

// flattenBackup flattens Backup from a JSON request object into the
// Backup type.
func flattenBackup(c *Client, i interface{}, res *Backup) *Backup {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	resultRes := &Backup{}
	resultRes.Name = dcl.FlattenString(m["name"])
	resultRes.Description = dcl.FlattenString(m["description"])
	resultRes.State = flattenBackupStateEnum(m["state"])
	resultRes.CreateTime = dcl.FlattenString(m["createTime"])
	resultRes.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	resultRes.CapacityGb = dcl.FlattenInteger(m["capacityGb"])
	resultRes.StorageBytes = dcl.FlattenInteger(m["storageBytes"])
	resultRes.SourceInstance = dcl.FlattenString(m["sourceInstance"])
	resultRes.SourceFileShare = dcl.FlattenString(m["sourceFileShare"])
	resultRes.SourceInstanceTier = flattenBackupSourceInstanceTierEnum(m["sourceInstanceTier"])
	resultRes.DownloadBytes = dcl.FlattenInteger(m["downloadBytes"])
	resultRes.Project = dcl.FlattenString(m["project"])
	resultRes.Location = dcl.FlattenString(m["location"])

	return resultRes
}

// flattenBackupStateEnumMap flattens the contents of BackupStateEnum from a JSON
// response object.
func flattenBackupStateEnumMap(c *Client, i interface{}, res *Backup) map[string]BackupStateEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackupStateEnum{}
	}

	if len(a) == 0 {
		return map[string]BackupStateEnum{}
	}

	items := make(map[string]BackupStateEnum)
	for k, item := range a {
		items[k] = *flattenBackupStateEnum(item.(interface{}))
	}

	return items
}

// flattenBackupStateEnumSlice flattens the contents of BackupStateEnum from a JSON
// response object.
func flattenBackupStateEnumSlice(c *Client, i interface{}, res *Backup) []BackupStateEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackupStateEnum{}
	}

	if len(a) == 0 {
		return []BackupStateEnum{}
	}

	items := make([]BackupStateEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackupStateEnum(item.(interface{})))
	}

	return items
}

// flattenBackupStateEnum asserts that an interface is a string, and returns a
// pointer to a *BackupStateEnum with the same value as that string.
func flattenBackupStateEnum(i interface{}) *BackupStateEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BackupStateEnumRef(s)
}

// flattenBackupSourceInstanceTierEnumMap flattens the contents of BackupSourceInstanceTierEnum from a JSON
// response object.
func flattenBackupSourceInstanceTierEnumMap(c *Client, i interface{}, res *Backup) map[string]BackupSourceInstanceTierEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]BackupSourceInstanceTierEnum{}
	}

	if len(a) == 0 {
		return map[string]BackupSourceInstanceTierEnum{}
	}

	items := make(map[string]BackupSourceInstanceTierEnum)
	for k, item := range a {
		items[k] = *flattenBackupSourceInstanceTierEnum(item.(interface{}))
	}

	return items
}

// flattenBackupSourceInstanceTierEnumSlice flattens the contents of BackupSourceInstanceTierEnum from a JSON
// response object.
func flattenBackupSourceInstanceTierEnumSlice(c *Client, i interface{}, res *Backup) []BackupSourceInstanceTierEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []BackupSourceInstanceTierEnum{}
	}

	if len(a) == 0 {
		return []BackupSourceInstanceTierEnum{}
	}

	items := make([]BackupSourceInstanceTierEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenBackupSourceInstanceTierEnum(item.(interface{})))
	}

	return items
}

// flattenBackupSourceInstanceTierEnum asserts that an interface is a string, and returns a
// pointer to a *BackupSourceInstanceTierEnum with the same value as that string.
func flattenBackupSourceInstanceTierEnum(i interface{}) *BackupSourceInstanceTierEnum {
	s, ok := i.(string)
	if !ok {
		return nil
	}

	return BackupSourceInstanceTierEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Backup) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBackup(b, c, r)
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

type backupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         backupApiOperation
	FieldName        string // used for error logging
}

func convertFieldDiffsToBackupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]backupDiff, error) {
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
	var diffs []backupDiff
	// For each operation name, create a backupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		// Use the first field diff's field name for logging required recreate error.
		diff := backupDiff{FieldName: fieldDiffs[0].FieldName}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBackupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBackupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (backupApiOperation, error) {
	switch opName {

	case "updateBackupUpdateBackupOperation":
		return &updateBackupUpdateBackupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractBackupFields(r *Backup) error {
	return nil
}

func postReadExtractBackupFields(r *Backup) error {
	return nil
}
