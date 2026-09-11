// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lustre

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	gcp "cloud.google.com/go/lustre/apiv1"
	lustrepb "cloud.google.com/go/lustre/apiv1/lustrepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/lustre/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.LustreInstanceGVK, NewInstanceModel)
}

func NewInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelInstance{config: *config}, nil
}

var _ directbase.Model = &modelInstance{}

type modelInstance struct {
	config config.ControllerConfig
}

func (m *modelInstance) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Instance client: %w", err)
	}
	return gcpClient, err
}

func (m *modelInstance) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LustreInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewInstanceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Resolve all resource references (e.g. networkRef)
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get lustre GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	httpClient, err := m.config.NewAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building authenticated HTTP client: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := LustreInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(obj.GetLabels())

	return &InstanceAdapter{
		id:                 id,
		gcpClient:          gcpClient,
		httpClient:         httpClient,
		desired:            desired,
		desiredAccessRules: obj.Spec.AccessRulesOptions,
	}, nil
}

func (m *modelInstance) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type InstanceAdapter struct {
	id                 *krm.LustreInstanceIdentity
	gcpClient          *gcp.Client
	httpClient         *http.Client
	desired            *lustrepb.Instance
	actual             *lustrepb.Instance
	desiredAccessRules *krm.AccessRulesOptions
	actualAccessRules  *krm.AccessRulesOptions
}

var _ directbase.Adapter = &InstanceAdapter{}

// Find retrieves the GCP resource.
func (a *InstanceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Instance", "name", a.id)

	req := &lustrepb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Instance %q: %w", a.id, err)
	}

	a.actual = instancepb

	accessRules, err := a.getAccessRulesOptions(ctx)
	if err != nil {
		log.V(2).Info("could not fetch accessRulesOptions (will proceed without them)", "err", err)
	} else {
		a.actualAccessRules = accessRules
	}

	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Instance", "name", a.id)

	a.desired.Name = a.id.String()

	req := &lustrepb.CreateInstanceRequest{
		Parent:     a.id.Parent().String(),
		InstanceId: a.id.ID(),
		Instance:   a.desired,
	}
	op, err := a.gcpClient.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Instance %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Instance %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Instance", "name", a.id)

	if a.desiredAccessRules != nil {
		log.V(2).Info("applying initial accessRulesOptions after creation", "name", a.id)
		if err := a.patchAccessRulesOptions(ctx, a.desiredAccessRules); err != nil {
			return fmt.Errorf("setting accessRulesOptions for %s after creation: %w", a.id, err)
		}
	}

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetInstance(ctx, &lustrepb.GetInstanceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching Instance %s after creation: %w", a.id, err)
	}
	latestAccessRules, err := a.getAccessRulesOptions(ctx)
	if err != nil {
		log.V(2).Info("could not fetch accessRulesOptions after create", "err", err)
	}

	return a.updateStatus(ctx, createOp, latest, latestAccessRules)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *InstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Instance", "name", a.id)

	a.desired.Name = a.id.String()

	paths, err := common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	accessRulesChanged := !accessRulesEqual(a.desiredAccessRules, a.actualAccessRules)

	if len(paths) == 0 && !accessRulesChanged {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual, a.actualAccessRules)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	if accessRulesChanged {
		report.AddField("spec.accessRulesOptions", a.actualAccessRules, a.desiredAccessRules)
	}
	structuredreporting.ReportDiff(ctx, report)

	if len(paths) > 0 {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		req := &lustrepb.UpdateInstanceRequest{
			UpdateMask: updateMask,
			Instance:   a.desired,
		}
		op, err := a.gcpClient.UpdateInstance(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Instance %s: %w", a.id, err)
		}
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Instance %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Instance proto fields", "name", a.id)
	}

	if accessRulesChanged {
		log.V(2).Info("accessRulesOptions need update", "name", a.id)
		if err := a.patchAccessRulesOptions(ctx, a.desiredAccessRules); err != nil {
			return fmt.Errorf("updating accessRulesOptions for Instance %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated accessRulesOptions", "name", a.id)
	}

	// Fetch fully-populated resource after update
	latest, err := a.gcpClient.GetInstance(ctx, &lustrepb.GetInstanceRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching Instance %s after update: %w", a.id, err)
	}
	latestAccessRules, err := a.getAccessRulesOptions(ctx)
	if err != nil {
		log.V(2).Info("could not fetch accessRulesOptions after update", "err", err)
	}

	return a.updateStatus(ctx, updateOp, latest, latestAccessRules)
}

func (a *InstanceAdapter) patchAccessRulesOptions(ctx context.Context, accessRules *krm.AccessRulesOptions) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("patching accessRulesOptions", "name", a.id)

	bodyObj := map[string]interface{}{}
	if accessRules != nil {
		bodyObj["accessRulesOptions"] = accessRules
	} else {
		bodyObj["accessRulesOptions"] = nil
	}

	bodyBytes, err := json.Marshal(bodyObj)
	if err != nil {
		return fmt.Errorf("marshaling accessRulesOptions: %w", err)
	}

	url := fmt.Sprintf("https://lustre.googleapis.com/v1/%s?updateMask=accessRulesOptions", a.id.String())
	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("building PATCH request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing PATCH request for accessRulesOptions: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading PATCH response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("PATCH %s returned %d: %s", url, resp.StatusCode, string(respBody))
	}

	var opResp struct {
		Name string `json:"name"`
		Done bool   `json:"done"`
	}
	if err := json.Unmarshal(respBody, &opResp); err != nil {
		return fmt.Errorf("parsing operation response: %w", err)
	}

	if opResp.Name != "" && !opResp.Done {
		op := a.gcpClient.UpdateInstanceOperation(opResp.Name)
		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for accessRulesOptions update operation %s: %w", opResp.Name, err)
		}
	}
	return nil
}

func (a *InstanceAdapter) getAccessRulesOptions(ctx context.Context) (*krm.AccessRulesOptions, error) {
	url := fmt.Sprintf("https://lustre.googleapis.com/v1/%s", a.id.String())
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GET %s returned %d: %s", url, resp.StatusCode, string(body))
	}

	var raw struct {
		AccessRulesOptions *krm.AccessRulesOptions `json:"accessRulesOptions,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decoding instance JSON: %w", err)
	}
	return raw.AccessRulesOptions, nil
}

func (a *InstanceAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *lustrepb.Instance, accessRules *krm.AccessRulesOptions) error {
	mapCtx := &direct.MapContext{}
	status := &krm.LustreInstanceStatus{}
	status.ObservedState = LustreInstanceObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ObservedState.AccessRulesOptions = accessRules
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func accessRulesEqual(desired, actual *krm.AccessRulesOptions) bool {
	if desired == nil {
		return true
	}
	if actual == nil {
		return false
	}
	if direct.ValueOf(desired.DefaultSquashMode) != direct.ValueOf(actual.DefaultSquashMode) {
		return false
	}
	if direct.ValueOf(desired.DefaultSquashUid) != direct.ValueOf(actual.DefaultSquashUid) {
		return false
	}
	if direct.ValueOf(desired.DefaultSquashGid) != direct.ValueOf(actual.DefaultSquashGid) {
		return false
	}
	if len(desired.AccessRules) != len(actual.AccessRules) {
		return false
	}
	for i := range desired.AccessRules {
		r1 := &desired.AccessRules[i]
		r2 := &actual.AccessRules[i]
		if direct.ValueOf(r1.Name) != direct.ValueOf(r2.Name) {
			return false
		}
		if direct.ValueOf(r1.SquashMode) != direct.ValueOf(r2.SquashMode) {
			return false
		}
		if !slicesEqual(r1.IpAddressRanges, r2.IpAddressRanges) {
			return false
		}
	}
	return true
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *InstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LustreInstance{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LustreInstanceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	if a.actualAccessRules != nil {
		obj.Spec.AccessRulesOptions = a.actualAccessRules
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.LustreInstanceGVK)
	u.Object = uObj

	export.SetProjectID(u, a.id.Parent().ProjectID)
	export.SetLabels(u, a.actual.Labels)
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *InstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Instance", "name", a.id)

	req := &lustrepb.DeleteInstanceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Instance, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Instance %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Instance %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Instance", "name", a.id)
	return true, nil
}
